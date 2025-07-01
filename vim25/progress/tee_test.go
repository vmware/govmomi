// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package progress

import "testing"

func TestTee(t *testing.T) {
	var ok bool

	ch1 := make(chan Report)
	ch2 := make(chan Report)

	s := Tee(&dummySinker{ch: ch1}, &dummySinker{ch: ch2})

	in := s.Sink()
	in <- dummyReport{}
	close(in)

	// Receive dummy on both sinks
	<-ch1
	<-ch2

	_, ok = <-ch1
	if ok {
		t.Errorf("Expected channel to be closed")
	}

	_, ok = <-ch2
	if ok {
		t.Errorf("Expected channel to be closed")
	}
}
