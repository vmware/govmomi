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
	"fmt"

	"github.com/vmware/govmomi/vim25"
)

// progressWrapper wraps an inbound progress type and decorates it with a
// scaled percentage and prefixed detail.
type progressWrapper struct {
	p      vim25.Progress
	index  int
	total  int
	detail string
}

func (p progressWrapper) Percentage() float32 {
	b := 100 * float32(p.index) / float32(p.total)
	return b + (p.p.Percentage() / float32(p.total))
}

func (p progressWrapper) Detail() string {
	s := p.detail
	if t := p.p.Detail(); t != "" {
		s = fmt.Sprintf("%s: %s", s, t)
	}
	return s
}

func (p progressWrapper) Error() error {
	return p.p.Error()
}

// progressChannell wraps an inbound progress channel together
// with a prefix that should be used for its detail message.
type progressChannel struct {
	detail string
	ch     <-chan vim25.Progress
}

// ProgressAggregator is a funnel for multiple progress channels. Can be used
// to group a sequence of independent but serial progress channels.
type ProgressAggregator struct {
	C <-chan vim25.Progress

	cs chan progressChannel
}

// NewProgressAggregator creates a new ProgressAggregor.
// The argument specifies the anticipated total number of progress channels.
// It is used to compute the relative completeness.
func NewProgressAggregator(total int) *ProgressAggregator {
	out := make(chan vim25.Progress)
	pa := ProgressAggregator{
		C:  out,
		cs: make(chan progressChannel, 100), // Large enough to not block calls to NewChannel.
	}

	go func() {
		var ok bool
		var err error

		defer close(out)

		for index := 0; err == nil; index++ {
			var pc progressChannel

			pc, ok = <-pa.cs
			if !ok {
				break
			}

			// Compensate total if it was inaccurate.
			if index >= total {
				total++
			}

			for err == nil {
				var p vim25.Progress

				p, ok = <-pc.ch
				if !ok {
					break
				}

				// Wrap it
				pw := progressWrapper{
					p:      p,
					index:  index,
					total:  total,
					detail: pc.detail,
				}

				// Forward it
				out <- pw

				// Store error so the loops break if there is one.
				err = p.Error()
			}
		}
	}()

	return &pa
}

func (pa *ProgressAggregator) NewChannel(detail string) chan<- vim25.Progress {
	ch := make(chan vim25.Progress, 1)
	pa.cs <- progressChannel{detail: detail, ch: ch}
	return ch
}

func (pa *ProgressAggregator) Done() {
	close(pa.cs)
}
