// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import (
	"io"

	"github.com/vmware/govmomi/vim25/xml"
)

func Unmarshal(r io.Reader) (*Envelope, error) {
	var e Envelope

	dec := xml.NewDecoder(r)
	err := dec.Decode(&e)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

// Write satisfies the flags.OutputWriter interface.
func (e *Envelope) Write(w io.Writer) error {
	return xml.NewEncoder(w).Encode(e)
}
