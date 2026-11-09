// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package simulator

// Container-backed integration tests for the govmomi/toolbox binary acting as
// vmtoolsd/vmware-rpctool, exercising GuestRPC over the Component-A Unix
// socket. A seccomp-based AF_VSOCK intercept ("Component B") was prototyped
// separately and is not part of this change; it remains a candidate for
// future re-enablement if a caller needs a real AF_VSOCK socket() call to
// succeed inside the container rather than falling back to the Unix socket.
//
// Requires: Docker (or podman aliased as docker) on Linux.
// Skip gate: test.HasDocker().

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/test"
	"github.com/vmware/govmomi/vim25/types"
)

// buildToolboxBinary builds the govmomi/toolbox binary as a static linux/amd64
// binary.  It is the canonical vmtoolsd / vmware-rpctool replacement for
// container-backed VMs in vcsim tests.
// buildToolboxBinary returns the path to the cached govmomi/toolbox binary,
// sharing the build with the simulator's own buildToolboxArtifact (sync.Once).
// This avoids a second independent "go build" when RUN.vmci=true is set,
// which would also trigger buildToolboxArtifact during PowerOn.
func buildToolboxBinary(t *testing.T) string {
	t.Helper()
	path, err := buildToolboxArtifact()
	if err != nil {
		t.Skipf("toolbox artifact build failed: %v", err)
	}
	if path == "" {
		t.Skip("toolbox binary unavailable (govmomi/toolbox build may have been skipped)")
	}
	return path
}

// dockerExec runs "docker exec containerID args..." and returns trimmed stdout.
// Stderr from docker exec (including podman compatibility warnings) is isolated
// and only included in the failure message, not in the return value.
func dockerExec(t *testing.T, ctx context.Context, containerID string, args ...string) string {
	t.Helper()
	cmd := exec.CommandContext(ctx, "docker",
		append([]string{"exec", containerID}, args...)...)
	out, err := cmd.Output()
	if err != nil {
		var stderr []byte
		if ee, ok := err.(*exec.ExitError); ok {
			stderr = ee.Stderr
		}
		t.Fatalf("docker exec %v: %v\nstderr: %s", args, err, stderr)
	}
	return strings.TrimSpace(string(out))
}

// containerIDFor returns the Docker container ID for vmObj's running container.
func containerIDFor(t *testing.T, simCtx *Context, vmObj *VirtualMachine) string {
	t.Helper()
	var id string
	simCtx.WithLock(vmObj, func() {
		if vmObj.svm != nil && vmObj.svm.c != nil {
			id = vmObj.svm.c.id
		}
	})
	require.NotEmpty(t, id, "container ID must be set after PowerOn")
	return id
}

// readGuestInfoKey reads KEY directly from vmObj's ExtraConfig without going
// through the PropertyCollector.  Must be called with no lock held.
func readGuestInfoKey(simCtx *Context, vmObj *VirtualMachine, key string) string {
	var val string
	simCtx.WithLock(vmObj, func() {
		for _, opt := range vmObj.Config.ExtraConfig {
			ov := opt.GetOptionValue()
			if ov.Key == key {
				val = fmt.Sprintf("%v", ov.Value)
				return
			}
		}
	})
	return val
}

// TestVMCI_ToolboxBinary_GuestInfoRoundTrip validates that the govmomi/toolbox
// binary, when volume-overlaid as /usr/bin/vmtoolsd and /usr/bin/vmware-rpctool
// in a container-backed VM, correctly reads and writes guestinfo through the
// vcsim GuestRPC Unix socket server.
//
// Subtests:
//  1. vmtoolsd-info-set: info-set KEY VAL → value visible in ExtraConfig.
//  2. vmtoolsd-info-get: info-get KEY → bare value printed to stdout, exit 0.
//  3. vmware-rpctool-info-get: raw "1 VALUE" response, exit 0 (cloud-init contract).
//  4. vmtoolsd-info-get-missing: missing key → exit 1, no stdout (cloud-init contract).
//  5. stop: PowerOff completes cleanly.
func TestVMCI_ToolboxBinary_GuestInfoRoundTrip(t *testing.T) {
	if !test.HasDocker() {
		t.Skip("requires docker or podman (aliased as docker) on linux")
	}

	// ── Build the toolbox binary ──────────────────────────────────────────
	toolboxBin := buildToolboxBinary(t)

	// ── Stand up vcsim ────────────────────────────────────────────────────
	env := newVCSIMEnv(t)

	// ── Create container-backed VM ────────────────────────────────────────
	// Overlay the toolbox binary as /usr/bin/vmtoolsd (explicit RUN.volume).
	// /usr/bin/vmware-rpctool is auto-injected by buildToolboxArtifact (RUN.vmci=true).
	// Both dispatch on filepath.Base(os.Args[0]).
	const key = "guestinfo.toolbox-roundtrip"
	const val = "hello-from-toolbox"

	spec := types.VirtualMachineConfigSpec{
		Name:  "toolbox-roundtrip-test",
		Files: &types.VirtualMachineFileInfo{VmPathName: "[LocalDS_0] toolbox-roundtrip-test"},
		ExtraConfig: []types.BaseOptionValue{
			&types.OptionValue{
				Key:   ContainerBackingOptionKey,
				Value: `["alpine","sh","-c","trap exit SIGTERM; sleep infinity & wait"]`,
			},
			&types.OptionValue{Key: "RUN.vmci", Value: "true"},
			// Inject the toolbox binary as /usr/bin/vmtoolsd so vmtoolsd --cmd
			// dispatches on filepath.Base(os.Args[0]) = "vmtoolsd".
			// /usr/bin/vmware-rpctool is auto-injected by buildToolboxArtifact when
			// RUN.vmci=true; no explicit RUN.volume entry needed for it.
			&types.OptionValue{
				Key:   "RUN.volume.vmtoolsd",
				Value: toolboxBin + ":/usr/bin/vmtoolsd:ro",
			},
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

	vmObj := env.simCtx.Map.Get(vmRef).(*VirtualMachine)
	containerID := containerIDFor(t, env.simCtx, vmObj)

	// ── TEST 1: vmtoolsd --cmd info-set ───────────────────────────────────
	t.Run("vmtoolsd-info-set", func(t *testing.T) {
		// info-set exits 0 and produces no stdout on success.
		dockerExec(t, env.goCtx, containerID,
			"vmtoolsd", "--cmd", "info-set "+key+" "+val)

		got := readGuestInfoKey(env.simCtx, vmObj, key)
		require.Equal(t, val, got,
			"ExtraConfig key %q must be set after vmtoolsd --cmd info-set", key)
	})

	// ── TEST 2: vmtoolsd --cmd info-get ───────────────────────────────────
	t.Run("vmtoolsd-info-get", func(t *testing.T) {
		// info-get exits 0 and prints the bare value (no "1 " prefix).
		out := dockerExec(t, env.goCtx, containerID,
			"vmtoolsd", "--cmd", "info-get "+key)
		require.Equal(t, val, out,
			"vmtoolsd --cmd info-get must print bare value")
	})

	// ── TEST 3: vmware-rpctool info-get ───────────────────────────────────
	t.Run("vmware-rpctool-info-get", func(t *testing.T) {
		// vmware-rpctool prints the raw "1 VALUE" or "0 ..." response and
		// always exits 0 on a successful channel round-trip.  cloud-init
		// parses the "1 " prefix itself.
		out := dockerExec(t, env.goCtx, containerID,
			"vmware-rpctool", "info-get "+key)
		require.Equal(t, "1 "+val, out,
			"vmware-rpctool info-get must print raw '1 VALUE' response")
	})

	// ── TEST 4: vmtoolsd info-get missing key → exit 1 ───────────────────
	// Verifies the exit-code contract that cloud-init DataSourceVMware depends on:
	// exit 0 = key found, exit 1 = key missing or channel error.
	t.Run("vmtoolsd-info-get-missing", func(t *testing.T) {
		cmd := exec.CommandContext(env.goCtx, "docker",
			"exec", containerID,
			"vmtoolsd", "--cmd", "info-get guestinfo.nonexistent.key")
		out, err := cmd.Output()

		var exitErr *exec.ExitError
		require.ErrorAs(t, err, &exitErr,
			"vmtoolsd info-get for a missing key must exit non-zero")
		require.Equal(t, 1, exitErr.ExitCode(),
			"vmtoolsd info-get for a missing key must exit 1")
		require.Empty(t, strings.TrimSpace(string(out)),
			"vmtoolsd info-get for a missing key must produce no stdout")
	})

	// ── TEST 5: Stop path ─────────────────────────────────────────────────
	t.Run("stop", func(t *testing.T) {
		stopCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		offTask, err := vm.PowerOff(stopCtx)
		require.NoError(t, err)
		require.NoError(t, offTask.Wait(stopCtx),
			"PowerOff must complete within 60 s")

		t.Logf("PowerOff completed")
	})
}

// TestVMCI_SystemdInit_GuestInfoRoundTrip validates that a systemd-managed
// one-shot service can write a guestinfo key via the govmomi toolbox binary
// when systemd is the container's PID 1.
//
// # Design rationale
//
// The fix in toolbox/toolbox/main.go (openRPCChannel) makes vmware-rpctool
// prefer $VMX_RPC_SOCK (Component A — the GuestRPC Unix socket bind-mounted by
// vcsim) over AF_VSOCK. Component A is accessible for the full container
// lifetime; it is not affected by the crun→systemd exec transition.
//
// The test uses the dedicated image simulator/testdata/vmci-systemd-init, which
// is a minimal debian+systemd container with the vmci-test-writer.service unit
// baked in (via 'systemctl enable' in the Dockerfile). Baking the service into
// the image avoids the Docker/Podman bind-mount pitfall where mounting a host
// file to a non-existent container path creates a directory, not a file.
//
// This test does NOT use RUN.nestedContainers=true because --tmpfs /run
// (added by that flag) hides the /run/vmware/rpc.sock bind mount. The
// dedicated image has no nested container runtime so --tmpfs /run is not
// needed. RUN.privileged=true provides --privileged for systemd without the
// tmpfs mounts.
//
// # Test layout
//
//   - Image: docker.io/library/vmci-systemd-init:latest (minimal debian+systemd).
//     Build with: docker build -t docker.io/library/vmci-systemd-init:latest
//     ./simulator/testdata/vmci-systemd-init/
//   - RUN.vmci=true: Component A (GuestRPC Unix socket at $VMX_RPC_SOCK) +
//     toolbox binary injection.
//   - RUN.privileged=true: --privileged for systemd, WITHOUT --tmpfs /run.
//   - vmci-test-writer.service is baked into the image. It runs after
//     basic.target and calls vmware-rpctool to set guestinfo.vmci-systemd-test.
//   - Test polls ExtraConfig for guestinfo.vmci-systemd-test within 90 s.
//
// # Expected result: PASS.
// If this test fails, check:
//  1. VMX_RPC_SOCK is set in the container env and propagated via PassEnvironment=.
//  2. /run/vmware/rpc.sock is accessible (not hidden by --tmpfs /run).
//  3. openRPCChannel() in toolbox/toolbox/main.go checks VMX_RPC_SOCK first.
func TestVMCI_SystemdInit_GuestInfoRoundTrip(t *testing.T) {
	if !test.HasDocker() {
		t.Skip("requires docker or podman (aliased as docker) on linux")
	}

	const image = "docker.io/library/vmci-systemd-init:latest"
	if err := exec.Command("docker", "image", "inspect", image).Run(); err != nil {
		t.Skipf("image %s not available locally; "+
			"build with: docker build -t %s ./simulator/testdata/vmci-systemd-init/",
			image, image)
	}

	// ── Build the toolbox binary ──────────────────────────────────────────
	// RUN.vmci=true auto-injects the toolbox binary as /usr/bin/vmware-rpctool.
	// buildToolboxBinary ensures the binary is compiled and cached before the
	// container starts, so the injection bind-mount is available at PowerOn.
	buildToolboxBinary(t)

	// ── Stand up vcsim ────────────────────────────────────────────────────
	env := newVCSIMEnv(t)

	// ── Create container-backed VM ────────────────────────────────────────
	const guestInfoKey = "guestinfo.vmci-systemd-test"

	spec := types.VirtualMachineConfigSpec{
		Name:  "systemd-init-test",
		Files: &types.VirtualMachineFileInfo{VmPathName: "[LocalDS_0] systemd-init-test"},
		ExtraConfig: []types.BaseOptionValue{
			// vmci-systemd-init uses CMD ["/sbin/init"], which boots systemd as PID 1.
			&types.OptionValue{Key: ContainerBackingOptionKey, Value: `["` + image + `"]`},
			// RUN.vmci activates:
			//   - Component A: GuestRPC Unix socket (VMX_RPC_SOCK=/run/vmware/rpc.sock)
			//   - toolbox binary auto-injected as /usr/bin/vmware-rpctool
			&types.OptionValue{Key: "RUN.vmci", Value: "true"},
			// RUN.privileged: adds --privileged so systemd can mount cgroups.
			// Does NOT add --tmpfs /run, keeping /run/vmware/rpc.sock accessible.
			&types.OptionValue{Key: "RUN.privileged", Value: "true"},
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

	vmObj := env.simCtx.Map.Get(vmRef).(*VirtualMachine)
	containerID := containerIDFor(t, env.simCtx, vmObj)

	// ── Poll for guestinfo write ───────────────────────────────────────────
	// The service runs after network.target; systemd boot in a container
	// typically takes 10–30 s. Poll every 2 s with a 90 s ceiling.
	t.Logf("container %s: polling for %q (up to 90 s)…", containerID, guestInfoKey)
	const pollInterval = 2 * time.Second
	const pollTimeout = 90 * time.Second
	deadline := time.Now().Add(pollTimeout)

	var got string
	for time.Now().Before(deadline) {
		got = readGuestInfoKey(env.simCtx, vmObj, guestInfoKey)
		if got != "" {
			break
		}
		time.Sleep(pollInterval)
	}

	t.Logf("guestinfo value after polling: %q", got)
	require.Equal(t, "pass", got,
		"vmci-test-writer.service did not write guestinfo via vmware-rpctool within %s", pollTimeout)

	// ── Stop path ─────────────────────────────────────────────────────────
	stopCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	offTask, err := vm.PowerOff(stopCtx)
	require.NoError(t, err)
	require.NoError(t, offTask.Wait(stopCtx),
		"PowerOff must complete within 60 s")
}

// TestVMCI_NestedContainers_GuestInfoRoundTrip is the nestedContainers=true variant
// of TestVMCI_SystemdInit_GuestInfoRoundTrip. It exercises the full production
// configuration: systemd PID 1 + RUN.vmci=true + RUN.nestedContainers=true.
//
// # Why this test exists
//
// The "privileged, no tmpfs" variant (TestVMCI_SystemdInit_GuestInfoRoundTrip)
// confirms the basic GuestRPC path. This test adds:
//   - --tmpfs /run (from nestedContainers), which hides /run/vmware/rpc.sock
//     unless the bind-mount is a directory mount (see guestRPCVolumeMount fix).
//   - --cgroupns=private and --privileged (needed for containerd/kubelet inside).
//
// Its purpose is to characterise exactly where the nestedContainers combination
// fails or passes. It is expected to PASS once the directory bind-mount fix in
// guestrpc_server.go (guestRPCVolumeMount) is in place. If it still fails,
// the failure log will identify the next gap.
//
// # Known failure modes before the directory-mount fix
//
//   - guestinfo.vmci-systemd-test never set: vmware-rpctool cannot connect to the
//     GuestRPC socket because /run/vmware/rpc.sock is hidden by --tmpfs /run.
//     Fixed by bind-mounting the directory /run/vmware/ instead of the file.
//
// # Build prerequisite
//
// Build the test image first if not already present:
//
//	docker build -t docker.io/library/vmci-systemd-init:latest \
//	  ./simulator/testdata/vmci-systemd-init/
func TestVMCI_NestedContainers_GuestInfoRoundTrip(t *testing.T) {
	if !test.HasDocker() {
		t.Skip("requires docker or podman (aliased as docker) on linux")
	}

	const image = "docker.io/library/vmci-systemd-init:latest"
	if err := exec.Command("docker", "image", "inspect", image).Run(); err != nil {
		t.Skipf("image %s not available locally; "+
			"build with: docker build -t %s ./simulator/testdata/vmci-systemd-init/",
			image, image)
	}

	buildToolboxBinary(t)

	env := newVCSIMEnv(t)

	const guestInfoKey = "guestinfo.vmci-systemd-test"

	spec := types.VirtualMachineConfigSpec{
		Name:  "nested-containers-test",
		Files: &types.VirtualMachineFileInfo{VmPathName: "[LocalDS_0] nested-containers-test"},
		ExtraConfig: []types.BaseOptionValue{
			&types.OptionValue{Key: ContainerBackingOptionKey, Value: `["` + image + `"]`},
			&types.OptionValue{Key: "RUN.vmci", Value: "true"},
			// Use nestedContainers=true (not RUN.privileged). This adds --tmpfs /run,
			// which is the scenario we need to validate. The directory bind-mount fix
			// in guestRPCVolumeMount should keep /run/vmware/rpc.sock accessible.
			&types.OptionValue{Key: "RUN.nestedContainers", Value: "true"},
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

	vmObj := env.simCtx.Map.Get(vmRef).(*VirtualMachine)
	containerID := containerIDFor(t, env.simCtx, vmObj)

	t.Logf("container %s: polling for %q (up to 120 s)…", containerID, guestInfoKey)
	const pollInterval = 2 * time.Second
	const pollTimeout = 120 * time.Second
	deadline := time.Now().Add(pollTimeout)

	var got string
	for time.Now().Before(deadline) {
		got = readGuestInfoKey(env.simCtx, vmObj, guestInfoKey)
		if got != "" {
			break
		}
		time.Sleep(pollInterval)
	}

	t.Logf("guestinfo value after polling: %q", got)
	require.Equal(t, "pass", got,
		"vmci-test-writer.service did not write guestinfo within %s "+
			"(RUN.nestedContainers=true + --tmpfs /run)", pollTimeout)

	stopCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	offTask, err := vm.PowerOff(stopCtx)
	require.NoError(t, err)
	require.NoError(t, offTask.Wait(stopCtx))
}
