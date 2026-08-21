// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"net"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/toolbox"
	"github.com/vmware/govmomi/vim25/types"
)

// serveOnePipe runs serveConn against one end of a net.Pipe and returns the
// caller's end. "tools.set-state" is used throughout because dispatch answers
// it without touching the VM object, keeping these tests about framing only.
func serveOnePipe(t *testing.T) net.Conn {
	t.Helper()

	srv := newGuestRPCServer(&VirtualMachine{}, "")
	client, server := net.Pipe()

	srv.wg.Add(1)
	go srv.serveConn(server)

	t.Cleanup(func() { _ = client.Close() })
	_ = client.SetDeadline(time.Now().Add(10 * time.Second))

	return client
}

// TestServeConn_DataMapAcrossPayloadSizes guards against reintroducing a
// length-based framing heuristic.
//
// This socket previously carried two framings and chose between them by
// comparing the 4-byte length prefix against a threshold. Because a DataMap
// length is big-endian and the other framing's was little-endian, their
// readings overlapped: 255 entries lengths below 1 MiB were routed to the
// wrong parser, and the guest selects the length by padding its payload. The
// sizes below are the ones measured to misroute; they must all round-trip.
func TestServeConn_DataMapAcrossPayloadSizes(t *testing.T) {
	for _, payloadLen := range []int{16, 32, 228, 484, 740, 996, 4096} {
		cmd := "tools.set-state " + strings.Repeat("x", payloadLen-len("tools.set-state "))
		if len(cmd) != payloadLen {
			t.Fatalf("test setup: built a %d-byte command, wanted %d", len(cmd), payloadLen)
		}

		client := serveOnePipe(t)

		if err := toolbox.WriteDataMapPacket(client, []byte(cmd), true); err != nil {
			t.Fatalf("payload %d: write: %v", payloadLen, err)
		}
		reply, _, err := toolbox.ReadDataMapPacket(client)
		if err != nil {
			t.Fatalf("payload %d: response not DataMap-framed: %v", payloadLen, err)
		}
		if got := strings.TrimRight(string(reply), "\x00"); got != "1 " {
			t.Errorf("payload %d: got reply %q, want %q", payloadLen, got, "1 ")
		}
	}
}

// TestServeConn_PersistentConnection pins that the server keeps serving a
// connection across exchanges when the client does not set fastClose.
// UnixChannel opens once in Start and reuses the connection, so a server that
// closed per request would break on the second exchange.
func TestServeConn_PersistentConnection(t *testing.T) {
	client := serveOnePipe(t)

	for i := 0; i < 3; i++ {
		if err := toolbox.WriteDataMapPacket(client, []byte("tools.set-state powerOn"), false); err != nil {
			t.Fatalf("exchange %d: write: %v", i, err)
		}
		reply, _, err := toolbox.ReadDataMapPacket(client)
		if err != nil {
			t.Fatalf("exchange %d: read: %v", i, err)
		}
		if got := strings.TrimRight(string(reply), "\x00"); got != "1 " {
			t.Errorf("exchange %d: got %q, want %q", i, got, "1 ")
		}
	}
}

// TestInfoSet_RejectsNonGuestinfoNamespace pins the write namespace.
//
// The socket directory is bind-mounted read-write into a container that may run
// an arbitrary image, and vcsim reads RUN.* keys as host-side container
// directives, so a guest able to write them could reconfigure its own backing.
func TestInfoSet_RejectsNonGuestinfoNamespace(t *testing.T) {
	srv := newGuestRPCServer(&VirtualMachine{}, "")

	for _, key := range []string{
		"RUN.container",
		"RUN.vmci",
		"RUN.volume.evil",
		"RUN.nestedContainers",
		"guestinfoNoDot",
	} {
		if got := srv.dispatch("info-set " + key + " value"); !strings.HasPrefix(got, "0 ") {
			t.Errorf("info-set %s: got %q, want a rejection", key, got)
		}
	}
}

// TestInfoSet_RejectsSlashInKey pins real VMX's second guest-write
// restriction: a key containing "/" is read-only from the guest even when it
// is otherwise under guestinfo.*. This is distinct from the namespace check
// above — a key can pass the guestinfo. prefix and still be unwritable.
func TestInfoSet_RejectsSlashInKey(t *testing.T) {
	srv := newGuestRPCServer(&VirtualMachine{}, "")

	for _, key := range []string{
		"guestinfo.foo/bar",
		"guestinfo./tools/state",
	} {
		if got := srv.dispatch("info-set " + key + " value"); !strings.HasPrefix(got, "0 ") {
			t.Errorf("info-set %s: got %q, want a rejection", key, got)
		}
	}
}

// TestInfoGet_RejectsNonGuestinfoNamespace mirrors
// TestInfoSet_RejectsNonGuestinfoNamespace for the read side. Real VMX
// restricts info-get to guestinfo.* the same as info-set; before this test
// existed, info-get had no namespace check at all and would return the value
// of any ExtraConfig key present on the VM, including RUN.* host-config
// directives.
func TestInfoGet_RejectsNonGuestinfoNamespace(t *testing.T) {
	srv := newGuestRPCServer(&VirtualMachine{}, "")

	for _, key := range []string{
		"RUN.container",
		"RUN.vmci",
		"RUN.volume.evil",
		"guestinfoNoDot",
	} {
		if got := srv.dispatch("info-get " + key); !strings.HasPrefix(got, "0 ") {
			t.Errorf("info-get %s: got %q, want a rejection", key, got)
		}
	}
}

// TestInfoGet_AllowsSlashInKey pins the asymmetry: unlike info-set, info-get
// does NOT reject a slash-containing guestinfo.* key. Real VMX makes such
// keys read-only from the guest, not unreadable — the host (e.g. via
// ReconfigVM) can still set one, and the guest can still read it back.
func TestInfoGet_AllowsSlashInKey(t *testing.T) {
	env := newVCSIMEnv(t)

	vmObj := firstVM(env.simCtx)
	require.NotNil(t, vmObj, "no VMs in test inventory")

	const key = "guestinfo.foo/bar"
	const val = "host-value"

	// Host-side write, bypassing GuestRPC entirely (as ReconfigVM would).
	// Only the guest info-set path enforces the slash restriction.
	env.simCtx.AutoUpdate(vmObj, func() {
		vmObj.Config.ExtraConfig = append(vmObj.Config.ExtraConfig,
			&types.OptionValue{Key: key, Value: val})
	})

	socketPath := GuestRPCSocketPath(vmObj.Self.Value)
	t.Cleanup(func() { _ = os.Remove(socketPath) })

	srv := newGuestRPCServer(vmObj, socketPath)
	require.NoError(t, srv.Start(env.simCtx))
	t.Cleanup(srv.Stop)

	out := dialGuestRPC(t, socketPath)
	reply, err := out.Request([]byte("info-get " + key))
	require.NoError(t, err, "info-get")
	require.Equal(t, val, strings.TrimSpace(string(reply)))
}
