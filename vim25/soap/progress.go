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

package soap

import (
	"container/list"
	"fmt"
	"io"
	"sync/atomic"
	"time"

	"github.com/vmware/govmomi/vim25"
)

const (
	KiB = 1024
	MiB = 1024 * KiB
	GiB = 1024 * MiB
)

type Progress struct {
	t time.Time

	pos  int64
	size int64
	bps  *uint64

	err error
}

func (p Progress) Percentage() float32 {
	return 100.0 * float32(p.pos) / float32(p.size)
}

func (p Progress) Detail() string {
	// Refer to the progress reader's bps field, so that this function always
	// returns an up-to-date number.
	//
	// For example: if there hasn't been progress for the last 5 seconds, the
	// most recent progress report should report "0B/s".
	//
	bps := atomic.LoadUint64(p.bps)

	switch {
	case bps >= GiB:
		return fmt.Sprintf("%.1fGiB/s", float32(bps)/float32(GiB))
	case bps >= MiB:
		return fmt.Sprintf("%.1fMiB/s", float32(bps)/float32(MiB))
	case bps >= KiB:
		return fmt.Sprintf("%.1fKiB/s", float32(bps)/float32(KiB))
	default:
		return fmt.Sprintf("%dB/s", bps)
	}
}

func (p Progress) Error() error {
	return p.err
}

// progressReader wraps a io.Reader and sends a progress report over a channel
// for every read it handles.
type progressReader struct {
	r io.Reader

	pos  int64
	size int64
	bps  uint64

	ch    chan<- vim25.Progress
	bpsch chan<- Progress
}

// Read calls the Read function on the underlying io.Reader. Additionally,
// every read causes a progress report to be sent to the progress reader's
// underlying channel. This progress report is sent optimistically; it is
// dropped if it cannot be received immediately.
func (p *progressReader) Read(b []byte) (int, error) {
	n, err := p.r.Read(b)
	if err != nil {
		return n, err
	}

	p.pos += int64(n)
	q := Progress{
		t:    time.Now(),
		pos:  p.pos,
		size: p.size,
		bps:  &p.bps,
	}

	// Start bps computation if not already running
	if p.bpsch == nil {
		ch := make(chan Progress)
		p.bpsch = ch
		go p.bpsLoop(ch)
	}

	// Don't care if this is dropped
	select {
	case p.ch <- q:
	default:
	}

	// Don't care if this is dropped
	select {
	case p.bpsch <- q:
	default:
	}

	return n, err
}

// Done marks the progress reader as done, optionally including an error in the
// progress report. This final progress report is never dropped on the sending
// side. After sending it, the underlying channel is closed.
func (p *progressReader) Done(err error) {
	q := Progress{
		t:    time.Now(),
		pos:  p.pos,
		size: p.size,
		err:  err,
	}

	// Last one must always be delivered
	p.ch <- q
	close(p.ch)

	// Stop bps computation if running
	if p.bpsch != nil {
		close(p.bpsch)
	}
}

// bpsLoop computes the reader's throughput.
func (p *progressReader) bpsLoop(ch chan Progress) {
	l := list.New()

	for {
		var tch <-chan time.Time

		// Setup timer for front of list to become stale.
		if e := l.Front(); e != nil {
			dt := time.Second - time.Now().Sub(e.Value.(Progress).t)
			tch = time.After(dt)
		}

		select {
		case q, ok := <-ch:
			if !ok {
				return
			}

			l.PushBack(q)
		case <-tch:
			l.Remove(l.Front())
		}

		// Compute new bps
		if l.Len() == 0 {
			atomic.StoreUint64(&p.bps, 0)
		} else {
			f := l.Front().Value.(Progress)
			b := l.Back().Value.(Progress)
			atomic.StoreUint64(&p.bps, uint64(b.pos-f.pos))
		}
	}
}
