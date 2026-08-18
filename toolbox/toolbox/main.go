// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// toolbox is a GuestRPC agent binary built on the govmomi/toolbox library.
//
// It is intended to be embedded in guest processes (custom init processes,
// container-backed VMs, test harnesses) that need VMware GuestRPC / guestinfo
// functionality without linking to the full open-vm-tools stack.
//
// # Transport selection
//
// All modes (daemon, one-shot, rpctool) share the same selection logic:
//
//  1. VMX_RPC_SOCK Unix socket — preferred when $VMX_RPC_SOCK is set.
//     govmomi/simulator sets this env var on container-backed VMs (RUN.vmci=true)
//     and bind-mounts the socket into the container. The Unix socket persists for
//     the full container lifetime, unlike the AF_VSOCK seccomp intercept which is
//     tied to the initial crun process and cannot be re-established after systemd
//     (or a nested container runtime) exec-replaces that process. Use this transport
//     for any service that runs AFTER initial container boot (e.g. systemd services).
//
//  2. AF_VSOCK with DataMap framing.  Status:
//     • works in real VMs with the vmci kernel module loaded.
//     • govmomi/simulator container-backed VMs - support pending via seccomp intercept.
//
//  3. x86 backdoor channel (in eax,dx / port 0x5658) — fallback when neither
//     VMX_RPC_SOCK nor AF_VSOCK is available. Real VMs only.
//
// Errors from channel open are always reported to stderr; the binary never
// fails silently.
//
// # Invocation modes
//
// Daemon mode (no --cmd):
//
//	toolbox            — register capabilities + poll TCLO (vsock or backdoor)
//
// One-shot RPCI mode (vmtoolsd --cmd semantics):
//
//	toolbox --cmd "info-get guestinfo.KEY"
//	toolbox --cmd "info-set guestinfo.KEY VALUE"
//	toolbox --cmd "RPCI_COMMAND"
//
//	  • info-get: print bare value to stdout; exit 0 on found, exit 1 on missing.
//	  • info-set: exit 0 on success, exit 1 on failure.
//	  • capability / state RPCs: exit 0 immediately (no server round-trip needed).
//	  • other RPCs: exit 0 on "1 ..." response, exit 1 on "0 ..." response.
//
// vmware-rpctool compatibility (invoked as "vmware-rpctool"):
//
//	vmware-rpctool 'RPCI_COMMAND'
//
//	Sends the raw command, prints the raw "1 VALUE" or "0 reason" response, and
//	exits 0 on a successful channel round-trip (including "0 No value found").
//	Exits 1 only on channel failure.  This matches the real vmware-rpctool
//	contract that cloud-init DataSourceVMware relies on.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/vmware/govmomi/toolbox"
)

func main() {
	// vmware-rpctool compatibility: when invoked as "vmware-rpctool", treat
	// the single positional argument as a raw GuestRPC command string.
	if strings.Contains(filepath.Base(os.Args[0]), "vmware-rpctool") {
		if len(os.Args) != 2 {
			fmt.Fprintln(os.Stderr, "usage: vmware-rpctool 'COMMAND'")
			os.Exit(2)
		}
		if err := rpcCmd(os.Args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "vmware-rpctool: %v\n", err)
			os.Exit(1)
		}
		return
	}

	cmdArg := flag.String("cmd", "", "send a one-shot RPCI command and exit (vmtoolsd --cmd semantics)")
	flag.Parse()

	if *cmdArg != "" {
		if err := vmtoolsdCmd(*cmdArg); err != nil {
			fmt.Fprintf(os.Stderr, "vmtoolsd: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Daemon mode: start the full toolbox service.
	in, out := selectChannels()
	service := toolbox.NewService(in, out)

	if os.Getuid() == 0 {
		service.Power.Halt.Handler = toolbox.Halt
		service.Power.Reboot.Handler = toolbox.Reboot
	}

	if err := service.Start(); err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		log.Printf("signal %s received", <-sig)
		service.Stop()
	}()

	service.Wait()
}

// selectChannels returns (in, out) channels for daemon mode.
// VMX_RPC_SOCK is NOT used here: the Unix socket server closes after each
// request, so a persistent daemon channel would break on the second exchange.
// Daemon mode relies on AF_VSOCK (which the seccomp intercept keeps alive
// during initial crun setup) or the x86 backdoor.
func selectChannels() (toolbox.Channel, toolbox.Channel) {
	if toolbox.IsVsockAvailable() {
		return toolbox.NewNoopChannelIn(), toolbox.NewVsockChannelOut()
	}
	return toolbox.NewBackdoorChannelIn(), toolbox.NewBackdoorChannelOut()
}

// wellKnownGuestRPCSock is the path at which govmomi/simulator bind-mounts the
// per-VM GuestRPC unix socket when RUN.vmci=true.  It is tried as a fallback
// when VMX_RPC_SOCK is not set — which happens when the toolbox binary is
// invoked from a systemd service unit that does not inherit the container's
// initial environment (systemd does not automatically propagate environment
// variables from PID 1's environment to child service units).
const wellKnownGuestRPCSock = "/run/vmware/rpc.sock"

// openRPCChannel returns the best available GuestRPC outbound channel.
//
// Selection order:
//  1. VMX_RPC_SOCK unix socket — set by govmomi/simulator on container-backed VMs.
//     The socket is accessible for the full container lifetime and is not affected
//     by the crun→systemd exec transition that invalidates the AF_VSOCK seccomp fd.
//  2. Well-known unix socket path (/run/vmware/rpc.sock) — same socket as (1)
//     but accessed by path when VMX_RPC_SOCK is not present in the process
//     environment (e.g. systemd service units invoked from the bootstrapper).
//  3. AF_VSOCK (DataMap framing) — real ESX VMs or RUN.vmci=true during crun phase.
//  4. x86 backdoor — VMware Workstation/Fusion fallback.
//
// Returns a non-nil error only if all four transports are unavailable.
func openRPCChannel() (toolbox.Channel, error) {
	// Unix socket: try VMX_RPC_SOCK env var first, then the well-known path.
	// If both resolve to the same path (or if VMX_RPC_SOCK is unset), the
	// extra attempt is two fast dial failures — acceptable cost for simplicity.
	for _, sockPath := range []string{os.Getenv("VMX_RPC_SOCK"), wellKnownGuestRPCSock} {
		if sockPath == "" {
			continue
		}
		ch := toolbox.NewUnixChannelOut(sockPath)
		if err := ch.Start(); err == nil {
			return ch, nil
		}
	}
	if toolbox.IsVsockAvailable() {
		ch := toolbox.NewVsockChannelOut()
		if err := ch.Start(); err == nil {
			return ch, nil
		}
	}
	ch := toolbox.NewBackdoorChannelOut()
	if err := ch.Start(); err != nil {
		return nil, fmt.Errorf("GuestRPC channel unavailable: unix socket not found at %s or %s, AF_VSOCK not present, and backdoor failed: %w",
			os.Getenv("VMX_RPC_SOCK"), wellKnownGuestRPCSock, err)
	}
	return ch, nil
}

// rpcCmd sends a raw GuestRPC command and prints the raw server response
// ("1 VALUE" or "0 reason") to stdout, then exits.
//
// Exit semantics match the real vmware-rpctool binary:
//   - Exit 0 + print response on a successful channel round-trip.
//   - Exit 1 + message on stderr on channel failure (connect error, I/O error).
//
// cloud-init DataSourceVMware calls vmware-rpctool and checks:
//   - "1 " prefix → VMware environment, datasource active.
//   - "0 " prefix or exit 1 → not VMware / key missing, datasource deactivates.
func rpcCmd(command string) error {
	ch, err := openRPCChannel()
	if err != nil {
		return err
	}
	defer ch.Stop()

	if err := ch.Send([]byte(command)); err != nil {
		return err
	}
	reply, err := ch.Receive()
	if err != nil {
		return err
	}
	fmt.Println(string(reply))
	return nil
}

// vmtoolsdCmd dispatches a vmtoolsd --cmd "COMMAND" argument.
//
// Exit semantics match the real vmtoolsd binary:
//   - info-get KEY: print bare value, exit 0 (found) or exit 1 (missing/error).
//   - info-set KEY VAL: exit 0 on success, exit 1 on failure.
//   - tools.* / Capabilities_Register / Set_Option: acknowledge, exit 0.
//   - other RPCs: exit 0 on "1 ..." response, exit 1 on "0 ..." response.
//   - channel failure: exit 1 + message on stderr.
func vmtoolsdCmd(command string) error {
	command = strings.TrimSpace(command)

	// Capability and state RPCs are acknowledged without a server round-trip.
	// The real vmtoolsd also short-circuits these in some configurations.
	if strings.HasPrefix(command, "tools.") ||
		strings.HasPrefix(command, "Capabilities_Register") ||
		strings.HasPrefix(command, "Set_Option ") {
		return nil
	}

	ch, err := openRPCChannel()
	if err != nil {
		return err
	}
	defer ch.Stop()
	rpci := &toolbox.ChannelOut{Channel: ch}

	if strings.HasPrefix(command, "info-get ") {
		// info-get: exit 0 + print bare value on success; exit 1 on missing.
		// ChannelOut.Request strips the "1 " prefix on success and returns
		// an error (wrapping the "0 ..." response) on failure.
		val, err := rpci.Request([]byte(command))
		if err != nil {
			return err
		}
		// Strip the trailing null byte that some callers append.
		fmt.Println(string(bytes.TrimRight(val, "\x00")))
		return nil
	}

	// info-set and everything else: success = "1 ...", failure = "0 ...".
	_, err = rpci.Request([]byte(command))
	return err
}
