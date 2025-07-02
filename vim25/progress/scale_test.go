// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package progress

import "testing"

func TestScaleMany(t *testing.T) {
	ch := make(chan Report)
	a := NewAggregator(dummySinker{ch})
	defer a.Done()

	s := Scale(a, 5)

	go func() {
		for i := 0; i < 5; i++ {
			go func(ch chan<- Report) {
				ch <- dummyReport{p: 0.0}
				ch <- dummyReport{p: 50.0}
				close(ch)
			}(s.Sink())
		}
	}()

	// Expect percentages to be scaled across sinks
	for p := float32(0.0); p < 100.0; p += 10.0 {
		r := <-ch
		if r.Percentage() != p {
			t.Errorf("Expected percentage to be: %.0f%%", p)
		}
	}
}
