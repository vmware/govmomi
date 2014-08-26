/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"io"
	"testing"

	"github.com/vmware/govmomi/vim25"
)

type pr struct {
	p float32
	d string
	e error
}

func (p pr) Percentage() float32 {
	return p.p
}

func (p pr) Detail() string {
	return p.d
}

func (p pr) Error() error {
	return p.e
}

func TestProgressAggregatorDone(t *testing.T) {
	pa := NewProgressAggregator(2)
	pa.Done()

	_, ok := <-pa.C
	if ok {
		t.Errorf("Expected channel to be closed")
	}
}

func TestProgressAggregatorPercentage(t *testing.T) {
	var ch chan<- vim25.Progress
	var p vim25.Progress

	pa := NewProgressAggregator(2)
	defer pa.Done()

	ch = pa.NewChannel("")
	ch <- pr{p: 20.0}
	close(ch)

	p = <-pa.C
	if pct := p.Percentage(); pct != (0.0/2.0)+(20.0/2.0) {
		t.Errorf("Expected 10%%, got %.0f%%", pct)
	}

	ch = pa.NewChannel("")
	ch <- pr{p: 20.0}
	close(ch)

	p = <-pa.C
	if pct := p.Percentage(); pct != (100.0/2.0)+(20.0/2.0) {
		t.Errorf("Expected 60%%, got %.0f%%", pct)
	}

	// Specified 2 as the total number. Check that the percentage
	// is adjusted if more than 2 are created through NewChannel.
	close(pa.NewChannel(""))
	close(pa.NewChannel(""))

	ch = pa.NewChannel("")
	ch <- pr{p: 20.0}
	close(ch)

	p = <-pa.C
	if pct := p.Percentage(); pct != (400.0/5.0)+(20.0/5.0) {
		t.Errorf("Expected 84%%, got %.0f%%", pct)
	}
}

func TestProgressAggregatorDetail(t *testing.T) {
	var ch chan<- vim25.Progress
	var p vim25.Progress

	pa := NewProgressAggregator(0)
	defer pa.Done()

	ch = pa.NewChannel("A")
	ch <- pr{d: ""} // No detail
	close(ch)

	p = <-pa.C
	if d := p.Detail(); d != "A" {
		t.Errorf("Expected %s, got %s", "A", d)
	}

	ch = pa.NewChannel("B")
	ch <- pr{d: "something"} // Some detail
	close(ch)

	p = <-pa.C
	if d := p.Detail(); d != "B: something" {
		t.Errorf("Expected %s, got %s", "B: something", d)
	}
}

func TestProgressAggregatorError(t *testing.T) {
	var ch chan<- vim25.Progress
	var p vim25.Progress

	pa := NewProgressAggregator(0)
	defer pa.Done()

	ch = pa.NewChannel("")
	ch <- pr{e: io.EOF} // Send error
	close(ch)

	p = <-pa.C
	if err := p.Error(); err != io.EOF {
		t.Errorf("Expected %s, got %s", io.EOF, err)
	}

	// ProgressAggregator should terminate on error.
	_, ok := <-pa.C
	if ok {
		t.Errorf("Expected channel to be closed")
	}
}
