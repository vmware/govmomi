// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

// TestContainerDestroyRetry uses SOAP dispatch and tasks with an engine stub;
// an internal container.remove retry alone would miss lost VM registration.
func TestContainerDestroyRetry(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("requires a POSIX shell")
	}
	for _, failure := range []string{"rm", "volume rm", "network rm"} {
		t.Run(failure, func(t *testing.T) {
			dir := stubDestroyEngine(t)
			require.NoError(t, os.WriteFile(filepath.Join(dir, "fail"), []byte(failure), 0600))

			model, client, backing, ctx := containerDestroyFixture(t)
			defer os.Remove(filepath.Join(dir, "fail"))
			vm := object.NewVirtualMachine(client.Client, backing.Self)
			devices := len(backing.Config.Hardware.Device)
			originalID := backing.svm.c.id

			first, err := vm.Destroy(ctx)
			require.NoError(t, err)
			require.Error(t, first.Wait(ctx))
			var state mo.VirtualMachine
			require.NoError(t, vm.Properties(ctx, vm.Reference(), []string{"name", "config.hardware.device"}, &state), "cleanup error must leave the same MoRef publicly reachable")
			require.Len(t, state.Config.Hardware.Device, devices)
			require.Equal(t, originalID, backing.svm.c.id)
			for _, obj := range model.Map().AllReference("Task") {
				require.NotEqual(t, "FileManager.deleteDatastoreFile", obj.(*Task).Info.DescriptionId, "cleanup error must not request file deletion")
			}

			require.NoError(t, os.Remove(filepath.Join(dir, "fail")))
			second, err := vm.Destroy(ctx)
			require.NoError(t, err, "the retry must pass through the original public MoRef")
			require.NoError(t, second.Wait(ctx))
			require.Nil(t, model.Map().Get(vm.Reference()))
			require.Empty(t, backing.svm.c.id)
			for _, resource := range []string{"container-removed", "volume-removed", "network-removed"} {
				_, err := os.Stat(filepath.Join(dir, resource))
				require.NoError(t, err)
			}
			commands, err := os.ReadFile(filepath.Join(dir, "commands"))
			require.NoError(t, err)
			require.Equal(t, 2, strings.Count(string(commands), "rm\t-v\t-f\towned-id\n"), "both public attempts must preserve exact container identity")
			_, err = vm.Destroy(ctx)
			require.True(t, soap.IsSoapFault(err))
			require.IsType(t, types.ManagedObjectNotFound{}, soap.ToSoapFault(err).VimFault())
		})
	}
}

func TestContainerDestroyPoweredOn(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("requires a POSIX shell")
	}
	dir := stubDestroyEngine(t)
	model, client, backing, ctx := containerDestroyFixture(t)
	backing.Runtime.PowerState = types.VirtualMachinePowerStatePoweredOn
	vm := object.NewVirtualMachine(client.Client, backing.Self)
	result, err := vm.Destroy(ctx)
	require.NoError(t, err)
	err = result.Wait(ctx)
	var fault task.Error
	require.True(t, errors.As(err, &fault))
	require.IsType(t, &types.InvalidPowerState{}, fault.Fault())
	require.Same(t, backing, model.Map().Get(vm.Reference()))
	require.Equal(t, "owned-id", backing.svm.c.id)
	_, err = os.Stat(filepath.Join(dir, "commands"))
	require.True(t, os.IsNotExist(err), "powered-on rejection must not invoke the engine")
}

func containerDestroyFixture(t *testing.T) (*Model, *govmomi.Client, *VirtualMachine, context.Context) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	t.Cleanup(cancel)
	model := VPX()
	model.Cluster = 0
	model.Machine = 1
	model.Autostart = false
	t.Cleanup(model.Remove)
	require.NoError(t, model.Create())
	server := model.Service.NewServer()
	t.Cleanup(server.Close)
	client, err := govmomi.NewClient(ctx, server.URL, true)
	require.NoError(t, err)
	vm := model.Map().Any("VirtualMachine").(*VirtualMachine)
	vm.Runtime.PowerState = types.VirtualMachinePowerStatePoweredOff
	vm.svm = &simVM{vm: vm, c: &container{name: "owned-container", id: "owned-id"}}
	return model, client, vm, ctx
}

func stubDestroyEngine(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	script := `#!/bin/sh
case "$1" in
  rm) step=rm ;;
  *) step="$1 $2" ;;
esac
printf '%s' "$1" >> "$CONTROL_DIR/commands"
shift
printf '\t%s' "$@" >> "$CONTROL_DIR/commands"
printf '\n' >> "$CONTROL_DIR/commands"
if [ -f "$CONTROL_DIR/fail" ]; then
  failure=$(cat "$CONTROL_DIR/fail")
  if [ "$step" = "$failure" ]; then exit 1; fi
fi
case "$step" in
  rm) : > "$CONTROL_DIR/container-removed" ;;
  "volume ls") [ -f "$CONTROL_DIR/volume-removed" ] || printf 'volume-a\n' ;;
  "network ls") [ -f "$CONTROL_DIR/network-removed" ] || printf 'network-a\n' ;;
  "volume rm") : > "$CONTROL_DIR/volume-removed" ;;
  "network rm") : > "$CONTROL_DIR/network-removed" ;;
  *) exit 1 ;;
esac
`
	require.NoError(t, os.WriteFile(filepath.Join(dir, "docker"), []byte(script), 0700))
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("CONTROL_DIR", dir)
	return dir
}

// A request can resolve the VM before destruction, then wait before acquiring
// its lock. It must not invoke the captured handler after that VM is removed.
func TestContainerDestroyQueuedRequest(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("requires a POSIX shell")
	}
	for _, name := range []string{"Destroy_Task", "PowerOnVM_Task", "UnregisterVM"} {
		t.Run(name, func(t *testing.T) {
			stubDestroyEngine(t)
			model, client, backing, ctx := containerDestroyFixture(t)
			vm := object.NewVirtualMachine(client.Client, backing.Self)
			ready, resume := make(chan struct{}), make(chan struct{})
			var release sync.Once
			unblock := func() { release.Do(func() { close(resume) }) }
			defer unblock()
			var captured atomic.Bool
			model.Map().Handler = func(_ *Context, method *Method) (mo.Reference, types.BaseMethodFault) {
				if method.This == vm.Reference() && method.Name == name && captured.CompareAndSwap(false, true) {
					close(ready)
					<-resume
				}
				return nil, nil
			}
			done := make(chan error, 1)
			go func() {
				var err error
				switch name {
				case "Destroy_Task":
					_, err = vm.Destroy(ctx)
				case "PowerOnVM_Task":
					_, err = vm.PowerOn(ctx)
				case "UnregisterVM":
					err = vm.Unregister(ctx)
				}
				done <- err
			}()
			select {
			case <-ready:
			case <-ctx.Done():
				t.Fatal("queued request never reached handler lookup")
			}
			result, err := vm.Destroy(ctx)
			require.NoError(t, err)
			require.NoError(t, result.Wait(ctx))
			unblock()
			select {
			case err = <-done:
			case <-ctx.Done():
				t.Fatal("queued request did not complete")
			}
			require.True(t, soap.IsSoapFault(err), "queued request must reject the removed VM")
			require.IsType(t, types.ManagedObjectNotFound{}, soap.ToSoapFault(err).VimFault())
			require.Nil(t, model.Map().Get(vm.Reference()))
		})
	}
}
