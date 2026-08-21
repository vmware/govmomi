// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"bytes"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/vmware/govmomi/toolbox"
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

// TestServeConn_DataMapFramingNotMisrouted pins the framing discrimination
// against a regression that a magnitude threshold cannot avoid.
//
// Both framings open with a 4-byte length, one big-endian and one little-
// endian, so their readings overlap. The previous threshold routed to DataMap
// only when the little-endian reading exceeded 1 MiB, which is false whenever
// the DataMap entries length is a multiple of 256 below 3840 — an ordinary
// info-set command. Those packets were parsed as LE frames and answered in the
// wrong framing.
//
// The payload sizes below are exactly the ones measured to misroute; each
// produces an entries section of 256, 512 and 768 bytes respectively.
func TestServeConn_DataMapFramingNotMisrouted(t *testing.T) {
	for _, payloadLen := range []int{228, 484, 740} {
		cmd := "tools.set-state " + strings.Repeat("x", payloadLen-len("tools.set-state "))
		if len(cmd) != payloadLen {
			t.Fatalf("test setup: built a %d-byte command, wanted %d", len(cmd), payloadLen)
		}

		client := serveOnePipe(t)

		if err := toolbox.WriteDataMapPacket(client, []byte(cmd), true); err != nil {
			t.Fatalf("payload %d: WriteDataMapPacket: %v", payloadLen, err)
		}

		reply, _, err := toolbox.ReadDataMapPacket(client)
		if err != nil {
			t.Fatalf("payload %d: response was not DataMap-framed (misrouted to the LE parser): %v",
				payloadLen, err)
		}
		if got := strings.TrimRight(string(reply), "\x00"); got != "1 " {
			t.Errorf("payload %d: got reply %q, want %q", payloadLen, got, "1 ")
		}
	}
}

// TestServeConn_LEFramingStillRouted is the negative control: the structural
// discriminator must not pull ordinary LE frames into the DataMap parser.
func TestServeConn_LEFramingStillRouted(t *testing.T) {
	for _, cmd := range []string{
		"tools.set-state powerOn",
		"tools.capability.foo",
		"tools.set-state " + strings.Repeat("y", 4096),
	} {
		client := serveOnePipe(t)

		if err := toolbox.WriteUnixFrame(client, []byte(cmd)); err != nil {
			t.Fatalf("WriteUnixFrame(%d bytes): %v", len(cmd), err)
		}

		reply, err := toolbox.ReadUnixFrame(client)
		if err != nil {
			t.Fatalf("LE frame of %d bytes was not answered in LE framing: %v", len(cmd), err)
		}
		if got := strings.TrimRight(string(reply), "\x00"); got != "1 " {
			t.Errorf("LE frame of %d bytes: got reply %q, want %q", len(cmd), got, "1 ")
		}
	}
}

// TestServeConn_ZeroLengthLEFrame covers the keepalive frame, which carries no
// payload and so must be settled without reading further bytes.
func TestServeConn_ZeroLengthLEFrame(t *testing.T) {
	client := serveOnePipe(t)

	if err := toolbox.WriteUnixFrame(client, nil); err != nil {
		t.Fatalf("WriteUnixFrame(nil): %v", err)
	}

	reply, err := toolbox.ReadUnixFrame(client)
	if err != nil {
		t.Fatalf("zero-length frame not answered: %v", err)
	}
	if !bytes.Equal(reply, []byte("1 ")) {
		t.Errorf("zero-length frame: got %q, want %q", reply, "1 ")
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
