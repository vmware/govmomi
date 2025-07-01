// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import "testing"

var _ Channel = new(backdoorChannel)

func TestBackdoorChannel(t *testing.T) {
	in := NewBackdoorChannelIn()
	out := NewBackdoorChannelOut()

	funcs := []func() error{
		in.Start,
		out.Start,
		in.Stop,
		out.Stop,
	}

	for _, f := range funcs {
		err := f()

		if err != nil {
			if err == ErrNotVirtualWorld {
				t.SkipNow()
			}
			t.Fatal(err)
		}
	}

	// expect an error if we don't specify the protocol
	err := new(backdoorChannel).Start()
	if err == nil {
		t.Error("expected error")
	}
}
