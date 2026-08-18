// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

// Phase 3 integration tests: real container writes guestinfo via the GuestRPC
// unix socket; PropertyCollector sees the change.
//
// Requires Docker (or podman aliased as docker) on Linux.
// Skip gate: test.HasDocker().

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/test"
	"github.com/vmware/govmomi/vim25/types"
)

// buildStaticBinary cross-compiles pkg as a static linux/amd64 binary written
// to a temp dir.  Returns the output path.
//
// Skips the test if the host OS is not Linux.  Fails (not skips) for build
// errors on Linux, since those indicate a genuine regression rather than a
// missing cross-compilation toolchain.
func buildStaticBinary(t *testing.T, pkg, name string) string {
	t.Helper()

	if runtime.GOOS != "linux" {
		t.Skip("container integration test requires Linux")
	}

	outDir := t.TempDir()
	outPath := filepath.Join(outDir, name)

	cmd := exec.Command("go", "build",
		"-o", outPath,
		"-ldflags", "-w -s",
		pkg)
	cmd.Env = append(os.Environ(),
		"CGO_ENABLED=0",
		"GOOS=linux",
		"GOARCH=amd64",
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("build %s: %v\n%s", name, err, out)
	}

	require.NoError(t, os.Chmod(outPath, fs.FileMode(0755)))
	return outPath
}

// buildVsockTestBinary builds the Component A (unix socket) test binary.
func buildVsockTestBinary(t *testing.T) string {
	return buildStaticBinary(t, "github.com/vmware/govmomi/simulator/testdata/vsock-test", "vsock-test")
}

// TestContainerGuestRPC_RoundTrip is the Phase 3 integration gate:
// a real container process writes guestinfo via the GuestRPC unix socket, and
// the test verifies the value appears in ExtraConfig via PropertyCollector.
//
// Container setup:
//   - image:   alpine (small, available, shell included)
//   - command: sh -c "/vsock-test --roundtrip guestinfo.from.container hello && sleep 9999"
//   - volumes: <vsock-test binary>:/vsock-test:ro  (injected)
//   - RUN.vmci=true  → GuestRPC server started; socket bind-mounted at /run/vmware/rpc.sock
func TestContainerGuestRPC_RoundTrip(t *testing.T) {
	if !test.HasDocker() {
		t.Skip("requires docker or podman (aliased as docker) on linux")
	}

	vsockTestBin := buildVsockTestBinary(t)

	env := newVCSIMEnv(t)

	const key = "guestinfo.from.container"
	const val = "hello-from-container"
	cmd := fmt.Sprintf(
		`/vsock-test --socket %s --roundtrip %s %s && sleep 9999`,
		GuestRPCSocketName, key, val,
	)

	spec := types.VirtualMachineConfigSpec{
		Name: "guestrpc-integration-test",
		Files: &types.VirtualMachineFileInfo{
			VmPathName: "[LocalDS_0] guestrpc-integration-test",
		},
		ExtraConfig: []types.BaseOptionValue{
			&types.OptionValue{Key: ContainerBackingOptionKey, Value: fmt.Sprintf(`["alpine","sh","-c",%q]`, cmd)},
			&types.OptionValue{Key: "RUN.vmci", Value: "true"},
			&types.OptionValue{Key: "RUN.volume.vsock-test", Value: vsockTestBin + ":/vsock-test:ro"},
		},
	}

	require.NoError(t, test.ApplyContainerRuntimeDefaults(&spec))

	task, err := env.folder.VmFolder.CreateVM(env.goCtx, spec, env.pool, nil)
	require.NoError(t, err)

	info, err := task.WaitForResult(env.goCtx, nil)
	require.NoError(t, err)

	vmRef := info.Result.(types.ManagedObjectReference)
	vm := object.NewVirtualMachine(env.client.Client, vmRef)

	powerTask, err := vm.PowerOn(env.goCtx)
	require.NoError(t, err)
	require.NoError(t, powerTask.Wait(env.goCtx))

	pollExtraConfig(t, env, vmRef, key, val, 30*time.Second)

	offTask, err := vm.PowerOff(env.goCtx)
	require.NoError(t, err)
	require.NoError(t, offTask.Wait(env.goCtx))
}

// TestContainerGuestRPC_MultiVM verifies that N concurrent container-backed VMs
// each write independent guestinfo values and neither VM can read the other's data.
func TestContainerGuestRPC_MultiVM(t *testing.T) {
	if !test.HasDocker() {
		t.Skip("requires docker or podman (aliased as docker) on linux")
	}

	vsockTestBin := buildVsockTestBinary(t)

	env := newVCSIMEnv(t)

	const n = 2
	type vmEntry struct {
		ref types.ManagedObjectReference
		vm  *object.VirtualMachine
		key string
		val string
	}

	var vms []vmEntry
	for i := 0; i < n; i++ {
		key := fmt.Sprintf("guestinfo.multivm.vm%d", i)
		val := fmt.Sprintf("vm%d-hello", i)
		cmd := fmt.Sprintf(
			`/vsock-test --socket %s --roundtrip %s %s && sleep 9999`,
			GuestRPCSocketName, key, val,
		)
		spec := types.VirtualMachineConfigSpec{
			Name: fmt.Sprintf("guestrpc-multivm-%d", i),
			Files: &types.VirtualMachineFileInfo{
				VmPathName: fmt.Sprintf("[LocalDS_0] guestrpc-multivm-%d", i),
			},
			ExtraConfig: []types.BaseOptionValue{
				&types.OptionValue{Key: ContainerBackingOptionKey, Value: fmt.Sprintf(`["alpine","sh","-c",%q]`, cmd)},
				&types.OptionValue{Key: "RUN.vmci", Value: "true"},
				&types.OptionValue{Key: "RUN.volume.vsock-test", Value: vsockTestBin + ":/vsock-test:ro"},
			},
		}
		require.NoError(t, test.ApplyContainerRuntimeDefaults(&spec))

		task, err := env.folder.VmFolder.CreateVM(env.goCtx, spec, env.pool, nil)
		require.NoError(t, err)
		info, err := task.WaitForResult(env.goCtx, nil)
		require.NoError(t, err)

		ref := info.Result.(types.ManagedObjectReference)
		vm := object.NewVirtualMachine(env.client.Client, ref)
		powerTask, err := vm.PowerOn(env.goCtx)
		require.NoError(t, err)
		require.NoError(t, powerTask.Wait(env.goCtx))

		vms = append(vms, vmEntry{ref: ref, vm: vm, key: key, val: val})
	}

	for _, entry := range vms {
		pollExtraConfig(t, env, entry.ref, entry.key, entry.val, 30*time.Second)
	}

	for _, entry := range vms {
		offTask, err := entry.vm.PowerOff(env.goCtx)
		require.NoError(t, err)
		require.NoError(t, offTask.Wait(env.goCtx))
	}
}
