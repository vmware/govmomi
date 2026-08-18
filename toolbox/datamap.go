// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"encoding/binary"
	"fmt"
	"io"
)

// DataMap field types (from open-vm-tools lib/dataMap/dataMap.h).
const (
	DMFieldTypeInt64  = 1
	DMFieldTypeString = 2
)

// GuestRPC DataMap field IDs (from open-vm-tools lib/guestRpc/tclodefs.h).
const (
	GuestRPCFieldType      = 1
	GuestRPCFieldPayload   = 2
	GuestRPCFieldFastClose = 3

	GuestRPCTypeData = 1
)

// ReadDataMapPacket reads one DataMap-framed packet from r and returns the
// GUESTRPCPKT_FIELD_PAYLOAD bytes and whether GUESTRPCPKT_FIELD_FAST_CLOSE was set.
//
// Wire format (all integers big-endian):
//
//	[4B: entries_length]
//	[entries: repeated { fieldType(4B) fieldID(4B) value... }]
//
//	int64 value:  low32(4B) high32(4B)
//	string value: strLen(4B) bytes(strLen)
func ReadDataMapPacket(r io.Reader) (payload []byte, fastClose bool, err error) {
	var hdr [4]byte
	if _, err = io.ReadFull(r, hdr[:]); err != nil {
		return
	}
	entriesLen := binary.BigEndian.Uint32(hdr[:])
	if entriesLen == 0 {
		return nil, false, nil
	}

	entries := make([]byte, entriesLen)
	if _, err = io.ReadFull(r, entries); err != nil {
		return
	}

	buf := entries
	for len(buf) >= 8 {
		fieldType := binary.BigEndian.Uint32(buf[:4])
		fieldID := binary.BigEndian.Uint32(buf[4:8])
		buf = buf[8:]

		switch fieldType {
		case DMFieldTypeInt64:
			if len(buf) < 8 {
				err = fmt.Errorf("DataMap: truncated int64 field %d", fieldID)
				return
			}
			low := binary.BigEndian.Uint32(buf[:4])
			// high32 (buf[4:8]) is intentionally ignored: all DataMap int64
			// fields we care about (FIELD_TYPE, FIELD_FAST_CLOSE) fit in uint32.
			buf = buf[8:]
			if fieldID == GuestRPCFieldFastClose && low != 0 {
				fastClose = true
			}

		case DMFieldTypeString:
			if len(buf) < 4 {
				err = fmt.Errorf("DataMap: truncated string length for field %d", fieldID)
				return
			}
			strLen := binary.BigEndian.Uint32(buf[:4])
			buf = buf[4:]
			if uint32(len(buf)) < strLen {
				err = fmt.Errorf("DataMap: truncated string body field %d (want %d have %d)",
					fieldID, strLen, len(buf))
				return
			}
			if fieldID == GuestRPCFieldPayload {
				payload = buf[:strLen]
			}
			buf = buf[strLen:]

		default:
			// Unknown field type — stop parsing to avoid spinning on malformed input.
			return
		}
	}
	return
}

// WriteDataMapPacket encodes payload as a DataMap packet and writes it to w.
// When fastClose is true, GUESTRPCPKT_FIELD_FAST_CLOSE=1 is appended (used by
// one-shot RPCI callers like vmtoolsd --cmd; set false for server responses).
func WriteDataMapPacket(w io.Writer, payload []byte, fastClose bool) error {
	// Entry layout sizes:
	//   FIELD_TYPE  (int64): type(4) + fieldID(4) + low32(4) + high32(4) = 16 bytes
	//   FIELD_PAYLOAD (str): type(4) + fieldID(4) + strlen(4) + bytes     = 12 + len(payload)
	//   FIELD_FAST_CLOSE (int64, optional): 16 bytes
	entriesLen := 16 + 12 + len(payload)
	if fastClose {
		entriesLen += 16
	}

	buf := make([]byte, 4+entriesLen)
	off := 0
	put32 := func(v uint32) {
		binary.BigEndian.PutUint32(buf[off:], v)
		off += 4
	}

	put32(uint32(entriesLen)) // header: length of entries section

	// FIELD_TYPE = TYPE_DATA
	put32(DMFieldTypeInt64)
	put32(GuestRPCFieldType)
	put32(GuestRPCTypeData) // low32
	put32(0)                // high32

	// FIELD_PAYLOAD = payload string
	put32(DMFieldTypeString)
	put32(GuestRPCFieldPayload)
	put32(uint32(len(payload)))
	// String body is variable-length so it can't use put32; copy + advance manually.
	copy(buf[off:], payload)
	off += len(payload)

	if fastClose {
		// FIELD_FAST_CLOSE = 1
		put32(DMFieldTypeInt64)
		put32(GuestRPCFieldFastClose)
		put32(1) // low32 = 1
		put32(0) // high32
	}

	_, err := w.Write(buf)
	return err
}
