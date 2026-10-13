// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"bytes"
	"encoding/binary"
	"strings"
	"testing"
)

// TestDataMapRoundTrip covers the payload sizes a GuestRPC exchange actually
// uses, including the small regime where a length-based framing heuristic used
// to misroute packets.
func TestDataMapRoundTrip(t *testing.T) {
	for _, n := range []int{0, 1, 2, 8, 16, 228, 484, 740, 4096} {
		payload := []byte(strings.Repeat("x", n))

		var buf bytes.Buffer
		if err := WriteDataMapPacket(&buf, payload, false); err != nil {
			t.Fatalf("WriteDataMapPacket(%d bytes): %v", n, err)
		}
		got, fastClose, err := ReadDataMapPacket(&buf)
		if err != nil {
			t.Fatalf("ReadDataMapPacket(%d bytes): %v", n, err)
		}
		if fastClose {
			t.Errorf("payload %d: fastClose set on a packet written with false", n)
		}
		if n > 0 && !bytes.Equal(got, payload) {
			t.Errorf("payload %d: round-trip mismatch", n)
		}
	}
}

// TestDataMapFastCloseRoundTrips pins the flag, since UnixChannel relies on
// NOT setting it (its connection is reused) while VsockChannel relies on it.
func TestDataMapFastCloseRoundTrips(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteDataMapPacket(&buf, []byte("info-get guestinfo.x"), true); err != nil {
		t.Fatalf("WriteDataMapPacket: %v", err)
	}
	_, fastClose, err := ReadDataMapPacket(&buf)
	if err != nil {
		t.Fatalf("ReadDataMapPacket: %v", err)
	}
	if !fastClose {
		t.Error("fastClose was set on write but not reported on read")
	}
}

// TestReadDataMapPacketRejectsOversizeLength pins the bound on the entries
// length prefix.
//
// The prefix is peer-supplied: vcsim bind-mounts the GuestRPC socket directory
// read-write into a container that may run an arbitrary image, so an unbounded
// length would let four bytes from the guest size an allocation in the host
// process. Without the bound the header below requests 4 GiB.
func TestReadDataMapPacketRejectsOversizeLength(t *testing.T) {
	for _, n := range []uint32{MaxDataMapEntries + 1, 1 << 30, ^uint32(0)} {
		hdr := make([]byte, 4)
		binary.BigEndian.PutUint32(hdr, n)

		payload, _, err := ReadDataMapPacket(bytes.NewReader(hdr))
		if err == nil {
			t.Errorf("entries length %d: expected an error, got %d bytes", n, len(payload))
		}
		if payload != nil {
			t.Errorf("entries length %d: expected no allocation, got %d bytes", n, len(payload))
		}
	}
}
