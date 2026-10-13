// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/toolbox"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

// vcSimEnv bundles the common vcsim test infrastructure so tests do not
// need to repeat the 10-line VPX setup block.  Call newVCSIMEnv(t) at the
// start of any test that needs a running vcsim instance.
type vcSimEnv struct {
	goCtx  context.Context
	m      *Model
	simCtx *Context
	client *govmomi.Client
	finder *find.Finder
	pool   *object.ResourcePool
	folder *object.DatacenterFolders
	pc     *property.Collector
}

// newVCSIMEnv creates a VPX model and HTTP server, registers cleanup via
// t.Cleanup, and returns a fully-initialised vcSimEnv.
func newVCSIMEnv(t *testing.T) *vcSimEnv {
	t.Helper()

	goCtx := context.Background()
	m := VPX()
	t.Cleanup(m.Remove)
	require.NoError(t, m.Create())
	s := m.Service.NewServer()
	t.Cleanup(s.Close)

	c, err := govmomi.NewClient(goCtx, s.URL, true)
	require.NoError(t, err)

	finder := find.NewFinder(c.Client)
	pool, err := finder.ResourcePool(goCtx, "DC0_H0/Resources")
	require.NoError(t, err)
	dc, err := finder.Datacenter(goCtx, "DC0")
	require.NoError(t, err)
	f, err := dc.Folders(goCtx)
	require.NoError(t, err)

	return &vcSimEnv{
		goCtx:  goCtx,
		m:      m,
		simCtx: m.Service.Context,
		client: c,
		finder: finder,
		pool:   pool,
		folder: f,
		pc:     property.DefaultCollector(c.Client),
	}
}

// pollExtraConfig blocks until KEY==WANT appears in vmRef's ExtraConfig via
// PropertyCollector, or the test fails with a timeout error.
func pollExtraConfig(t *testing.T, env *vcSimEnv, vmRef types.ManagedObjectReference, key, want string, timeout time.Duration) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		var mvm mo.VirtualMachine
		if err := env.pc.RetrieveOne(env.goCtx, vmRef, []string{"config.extraConfig"}, &mvm); err != nil {
			time.Sleep(200 * time.Millisecond)
			continue
		}
		for _, opt := range mvm.Config.ExtraConfig {
			v := opt.GetOptionValue()
			if v.Key == key {
				require.Equalf(t, want, v.Value,
					"guestinfo key %q: got %v want %q", key, v.Value, want)
				return
			}
		}
		time.Sleep(200 * time.Millisecond)
	}
	t.Fatalf("timeout after %s: %q not found in ExtraConfig", timeout, key)
}

// removeStaleGuestRPCSockets removes vcsim-rpc-*.sock stale files left by
// previously crashed test runs.  TestMain calls the equivalent sweep before
// the entire test binary runs; this function is available for tests that
// need an extra clean before a specific sensitive operation.
func removeStaleGuestRPCSockets() {
	for _, pattern := range []string{
		filepath.Join(os.TempDir(), "vcsim-rpc-*.sock"),
	} {
		if matches, err := filepath.Glob(pattern); err == nil {
			for _, f := range matches {
				_ = os.Remove(f)
			}
		}
	}
}

// dialGuestRPC connects a ChannelOut to the unix socket at socketPath.
func dialGuestRPC(t *testing.T, socketPath string) *toolbox.ChannelOut {
	t.Helper()
	ch := toolbox.NewUnixChannelOut(socketPath)
	require.NoError(t, ch.Start(), "dialGuestRPC: connect %s", socketPath)
	t.Cleanup(func() { _ = ch.Stop() })
	return &toolbox.ChannelOut{Channel: ch}
}

// firstVM returns the first VirtualMachine object from the registry.
func firstVM(simCtx *Context) *VirtualMachine {
	refs := simCtx.Map.AllReference("VirtualMachine")
	if len(refs) == 0 {
		return nil
	}
	return simCtx.Map.Get(refs[0].Reference()).(*VirtualMachine)
}

// TestGuestRPCServer_InfoGetSet exercises basic info-get and info-set.
func TestGuestRPCServer_InfoGetSet(t *testing.T) {
	env := newVCSIMEnv(t)

	vmObj := firstVM(env.simCtx)
	require.NotNil(t, vmObj, "no VMs in test inventory")

	socketPath := GuestRPCSocketPath(vmObj.Self.Value)
	t.Cleanup(func() { _ = os.Remove(socketPath) })

	srv := newGuestRPCServer(vmObj, socketPath)
	require.NoError(t, srv.Start(env.simCtx))
	t.Cleanup(srv.Stop)

	out := dialGuestRPC(t, socketPath)

	_, err := out.Request([]byte("info-set guestinfo.test.key hello"))
	require.NoError(t, err, "info-set")

	reply, err := out.Request([]byte("info-get guestinfo.test.key"))
	require.NoError(t, err, "info-get")
	require.Equal(t, "hello", strings.TrimSpace(string(reply)))
}

// TestGuestRPCServer_InfoSet_PropertyChange verifies info-set triggers a
// PropertyCollector notification.
func TestGuestRPCServer_InfoSet_PropertyChange(t *testing.T) {
	env := newVCSIMEnv(t)

	vmObj := firstVM(env.simCtx)
	require.NotNil(t, vmObj)

	socketPath := GuestRPCSocketPath(vmObj.Self.Value)
	t.Cleanup(func() { _ = os.Remove(socketPath) })

	srv := newGuestRPCServer(vmObj, socketPath)
	require.NoError(t, srv.Start(env.simCtx))
	t.Cleanup(srv.Stop)

	out := dialGuestRPC(t, socketPath)
	_, err := out.Request([]byte("info-set guestinfo.notify.test triggered"))
	require.NoError(t, err)

	pollExtraConfig(t, env, vmObj.Self, "guestinfo.notify.test", "triggered", 2*time.Second)
}

// TestGuestRPCServer_RoundTrip tests host-write/guest-read and guest-write/host-read.
func TestGuestRPCServer_RoundTrip(t *testing.T) {
	env := newVCSIMEnv(t)

	vmObj := firstVM(env.simCtx)
	require.NotNil(t, vmObj)

	socketPath := GuestRPCSocketPath(vmObj.Self.Value)
	t.Cleanup(func() { _ = os.Remove(socketPath) })

	srv := newGuestRPCServer(vmObj, socketPath)
	require.NoError(t, srv.Start(env.simCtx))
	t.Cleanup(srv.Stop)

	env.simCtx.AutoUpdate(vmObj, func() {
		vmObj.Config.ExtraConfig = append(vmObj.Config.ExtraConfig,
			&types.OptionValue{Key: "guestinfo.from.host", Value: "hostvalue"})
	})

	out := dialGuestRPC(t, socketPath)

	reply, err := out.Request([]byte("info-get guestinfo.from.host"))
	require.NoError(t, err)
	require.Equal(t, "hostvalue", strings.TrimSpace(string(reply)))

	_, err = out.Request([]byte("info-set guestinfo.from.guest guestvalue"))
	require.NoError(t, err)

	var gotValue string
	env.simCtx.WithLock(vmObj, func() {
		for _, opt := range vmObj.Config.ExtraConfig {
			v := opt.GetOptionValue()
			if v.Key == "guestinfo.from.guest" {
				gotValue = fmt.Sprintf("%v", v.Value)
				break
			}
		}
	})
	require.Equal(t, "guestvalue", gotValue)
}

// TestGuestRPCServer_MultiVM verifies that two coexisting VMs each get an
// independent GuestRPC server, and that one VM's guestinfo keys are not
// visible to the other. The two VMs' servers are exercised sequentially, not
// concurrently; this test is about isolation, not concurrency.
func TestGuestRPCServer_MultiVM(t *testing.T) {
	env := newVCSIMEnv(t)

	refs := env.simCtx.Map.AllReference("VirtualMachine")
	if len(refs) < 2 {
		t.Skip("need at least 2 VMs")
	}

	type vmHandle struct {
		obj *VirtualMachine
		srv *GuestRPCServer
		out *toolbox.ChannelOut
	}

	var handles []vmHandle
	for i := 0; i < 2; i++ {
		obj := env.simCtx.Map.Get(refs[i].Reference()).(*VirtualMachine)
		sp := GuestRPCSocketPath(obj.Self.Value)
		t.Cleanup(func() { _ = os.Remove(sp) })

		srv := newGuestRPCServer(obj, sp)
		require.NoError(t, srv.Start(env.simCtx), "Start VM %d", i)
		t.Cleanup(srv.Stop)

		handles = append(handles, vmHandle{obj: obj, srv: srv, out: dialGuestRPC(t, sp)})
	}

	for i, h := range handles {
		key := fmt.Sprintf("guestinfo.concurrent.vm%d", i)
		val := fmt.Sprintf("vm%d-value", i)
		_, err := h.out.Request([]byte(fmt.Sprintf("info-set %s %s", key, val)))
		require.NoError(t, err, "VM %d info-set", i)
	}

	for i, h := range handles {
		key := fmt.Sprintf("guestinfo.concurrent.vm%d", i)
		want := fmt.Sprintf("vm%d-value", i)
		reply, err := h.out.Request([]byte("info-get " + key))
		require.NoError(t, err, "VM %d info-get own key", i)
		require.Equal(t, want, strings.TrimSpace(string(reply)))

		otherKey := fmt.Sprintf("guestinfo.concurrent.vm%d", 1-i)
		_, err2 := h.out.Request([]byte("info-get " + otherKey))
		require.Error(t, err2, "VM %d should not see VM %d's key %s", i, 1-i, otherKey)
	}
}

// TestGuestRPCServer_Cleanup verifies stale socket recovery.
func TestGuestRPCServer_Cleanup(t *testing.T) {
	env := newVCSIMEnv(t)

	vmObj := firstVM(env.simCtx)

	socketPath := GuestRPCSocketPath(vmObj.Self.Value + "-cleanup")
	socketDir := GuestRPCSocketDirPath(vmObj.Self.Value + "-cleanup")
	t.Cleanup(func() {
		_ = os.Remove(socketPath)
		_ = os.Remove(socketDir)
	})

	// Simulate a stale socket from a prior crash: pre-create the directory
	// (normally created by Start) then write a non-socket placeholder file.
	require.NoError(t, os.MkdirAll(socketDir, 0o700))
	require.NoError(t, os.WriteFile(socketPath, []byte("stale"), 0600))

	srv := newGuestRPCServer(vmObj, socketPath)
	require.NoError(t, srv.Start(env.simCtx), "Start with stale file")
	t.Cleanup(srv.Stop) // idempotent (sync.Once); also covers early-return on assertion failure below

	fi, err := os.Stat(socketPath)
	require.NoError(t, err)
	require.NotZero(t, fi.Mode()&os.ModeSocket, "expected socket mode, got %v", fi.Mode())

	srv.Stop()

	_, err = os.Stat(socketPath)
	require.True(t, os.IsNotExist(err), "socket file should be gone after Stop")
}
