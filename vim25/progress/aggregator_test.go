// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package progress

import (
	"testing"
	"time"
)

func TestAggregatorNoSinks(t *testing.T) {
	ch := make(chan Report)
	a := NewAggregator(dummySinker{ch})
	a.Done()

	_, ok := <-ch
	if ok {
		t.Errorf("Expected channel to be closed")
	}
}

func TestAggregatorMultipleSinks(t *testing.T) {
	ch := make(chan Report)
	a := NewAggregator(dummySinker{ch})

	for i := 0; i < 5; i++ {
		go func(ch chan<- Report) {
			ch <- dummyReport{}
			ch <- dummyReport{}
			close(ch)
		}(a.Sink())

		<-ch
		<-ch
	}

	a.Done()

	_, ok := <-ch
	if ok {
		t.Errorf("Expected channel to be closed")
	}
}

func TestAggregatorSinkInFlightOnDone(t *testing.T) {
	ch := make(chan Report)
	a := NewAggregator(dummySinker{ch})

	// Simulate upstream
	go func(ch chan<- Report) {
		time.Sleep(1 * time.Millisecond)
		ch <- dummyReport{}
		close(ch)
	}(a.Sink())

	// Drain downstream
	go func(ch <-chan Report) {
		<-ch
	}(ch)

	// This should wait for upstream to complete
	a.Done()

	_, ok := <-ch
	if ok {
		t.Errorf("Expected channel to be closed")
	}
}
