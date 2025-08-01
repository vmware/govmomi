// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package progress

import (
	"container/list"
	"context"
	"fmt"
	"io"
	"sync/atomic"
	"time"
)

type readerReport struct {
	pos  int64   // Keep first to ensure 64-bit alignment
	size int64   // Keep first to ensure 64-bit alignment
	bps  *uint64 // Keep first to ensure 64-bit alignment

	t time.Time

	err error
}

func (r readerReport) Percentage() float32 {
	if r.size <= 0 {
		return 0
	}
	return 100.0 * float32(r.pos) / float32(r.size)
}

func (r readerReport) Detail() string {
	const (
		KiB = 1024
		MiB = 1024 * KiB
		GiB = 1024 * MiB
	)

	// Use the reader's bps field, so this report returns an up-to-date number.
	//
	// For example: if there hasn't been progress for the last 5 seconds, the
	// most recent report should return "0B/s".
	//
	bps := atomic.LoadUint64(r.bps)

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

func (p readerReport) Error() error {
	return p.err
}

// reader wraps an io.Reader and sends a progress report over a channel for
// every read it handles.
type reader struct {
	r io.Reader

	pos  int64
	size int64
	bps  uint64

	ch  chan<- Report
	ctx context.Context
}

func NewReader(ctx context.Context, s Sinker, r io.Reader, size int64) *reader {
	pr := reader{
		r:    r,
		ctx:  ctx,
		size: size,
	}

	// Reports must be sent downstream and to the bps computation loop.
	pr.ch = Tee(s, newBpsLoop(&pr.bps)).Sink()

	return &pr
}

// Read calls the Read function on the underlying io.Reader. Additionally,
// every read causes a progress report to be sent to the progress reader's
// underlying channel.
func (r *reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	r.pos += int64(n)

	if err != nil && err != io.EOF {
		return n, err
	}

	q := readerReport{
		t:    time.Now(),
		pos:  r.pos,
		size: r.size,
		bps:  &r.bps,
	}

	select {
	case r.ch <- q:
	case <-r.ctx.Done():
	}

	return n, err
}

// Done marks the progress reader as done, optionally including an error in the
// progress report. After sending it, the underlying channel is closed.
func (r *reader) Done(err error) {
	q := readerReport{
		t:    time.Now(),
		pos:  r.pos,
		size: r.size,
		bps:  &r.bps,
		err:  err,
	}

	select {
	case r.ch <- q:
		close(r.ch)
	case <-r.ctx.Done():
	}
}

// newBpsLoop returns a sink that monitors and stores throughput.
func newBpsLoop(dst *uint64) SinkFunc {
	fn := func() chan<- Report {
		sink := make(chan Report)
		go bpsLoop(sink, dst)
		return sink
	}

	return fn
}

func bpsLoop(ch <-chan Report, dst *uint64) {
	l := list.New()

	for {
		var tch <-chan time.Time

		// Setup timer for front of list to become stale.
		if e := l.Front(); e != nil {
			dt := time.Second - time.Since(e.Value.(readerReport).t)
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
			atomic.StoreUint64(dst, 0)
		} else {
			f := l.Front().Value.(readerReport)
			b := l.Back().Value.(readerReport)
			atomic.StoreUint64(dst, uint64(b.pos-f.pos))
		}
	}
}
