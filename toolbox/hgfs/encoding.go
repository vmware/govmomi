// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package hgfs

import (
	"bytes"
	"encoding"
	"encoding/binary"
)

// MarshalBinary is a wrapper around binary.Write
func MarshalBinary(fields ...any) ([]byte, error) {
	buf := new(bytes.Buffer)

	for _, p := range fields {
		switch m := p.(type) {
		case encoding.BinaryMarshaler:
			data, err := m.MarshalBinary()
			if err != nil {
				return nil, ProtocolError(err)
			}

			_, _ = buf.Write(data)
		case []byte:
			_, _ = buf.Write(m)
		case string:
			_, _ = buf.WriteString(m)
		default:
			err := binary.Write(buf, binary.LittleEndian, p)
			if err != nil {
				return nil, ProtocolError(err)
			}
		}
	}

	return buf.Bytes(), nil
}

// UnmarshalBinary is a wrapper around binary.Read
func UnmarshalBinary(data []byte, fields ...any) error {
	buf := bytes.NewBuffer(data)

	for _, p := range fields {
		switch m := p.(type) {
		case encoding.BinaryUnmarshaler:
			return m.UnmarshalBinary(buf.Bytes())
		case *[]byte:
			*m = buf.Bytes()
			return nil
		default:
			err := binary.Read(buf, binary.LittleEndian, p)
			if err != nil {
				return ProtocolError(err)
			}
		}
	}

	return nil
}
