// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// TestUnixFrameRoundTrip covers the ordinary path and the zero-length frame,
// which is used as a keepalive by the simulator's GuestRPC server.
func TestUnixFrameRoundTrip(t *testing.T) {
	for _, payload := range [][]byte{
		[]byte("info-get guestinfo.metadata"),
		[]byte("1 "),
		bytes.Repeat([]byte("x"), MaxUnixFrameSize),
	} {
		var buf bytes.Buffer
		if err := WriteUnixFrame(&buf, payload); err != nil {
			t.Fatalf("WriteUnixFrame(%d bytes): %v", len(payload), err)
		}
		got, err := ReadUnixFrame(&buf)
		if err != nil {
			t.Fatalf("ReadUnixFrame(%d bytes): %v", len(payload), err)
		}
		if !bytes.Equal(got, payload) {
			t.Errorf("round-trip of %d bytes did not match", len(payload))
		}
	}

	var buf bytes.Buffer
	if err := WriteUnixFrame(&buf, nil); err != nil {
		t.Fatalf("WriteUnixFrame(nil): %v", err)
	}
	got, err := ReadUnixFrame(&buf)
	if err != nil || got != nil {
		t.Errorf("zero-length frame: got %v, %v; want nil, nil", got, err)
	}
}

// TestReadUnixFrameRejectsOversizeLength pins the bound on the length prefix.
//
// The prefix is peer-supplied: vcsim bind-mounts the GuestRPC socket directory
// read-write into a container that may run an arbitrary image, so an unbounded
// length would let four bytes from the guest size an allocation in the host
// process. Without the bound, the header below requests 4 GiB.
func TestReadUnixFrameRejectsOversizeLength(t *testing.T) {
	for _, n := range []uint32{MaxUnixFrameSize + 1, 1 << 30, ^uint32(0)} {
		hdr := make([]byte, 4)
		binary.LittleEndian.PutUint32(hdr, n)

		buf, err := ReadUnixFrame(bytes.NewReader(hdr))
		if err == nil {
			t.Errorf("length %d: expected an error, got %d bytes", n, len(buf))
		}
		if buf != nil {
			t.Errorf("length %d: expected no allocation, got %d bytes", n, len(buf))
		}
	}
}
