// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package progress

import (
	"context"
	"io"
	"strings"
	"sync"
	"testing"
	"testing/iotest"
)

func TestReader(t *testing.T) {
	s := "helloworld"
	ch := make(chan Report, 1)
	pr := NewReader(context.Background(), &dummySinker{ch}, strings.NewReader(s), int64(len(s)))

	var buf [10]byte
	var q Report
	var n int
	var err error

	// Read first byte
	n, err = pr.Read(buf[0:1])
	if n != 1 {
		t.Errorf("Expected n=1, but got: %d", n)
	}

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	q = <-ch
	if q.Error() != nil {
		t.Errorf("Error: %s", err)
	}

	if f := q.Percentage(); f != 10.0 {
		t.Errorf("Expected percentage after 1 byte to be 10%%, but got: %.0f%%", f)
	}

	// Read remaining bytes
	n, err = pr.Read(buf[:])
	if n != 9 {
		t.Errorf("Expected n=1, but got: %d", n)
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	q = <-ch
	if q.Error() != nil {
		t.Errorf("Error: %s", err)
	}

	if f := q.Percentage(); f != 100.0 {
		t.Errorf("Expected percentage after 10 bytes to be 100%%, but got: %.0f%%", f)
	}

	// Read EOF
	_, err = pr.Read(buf[:])
	<-ch
	if err != io.EOF {
		t.Errorf("Expected io.EOF, but got: %s", err)
	}

	// Mark progress reader as done
	pr.Done(io.EOF)
	<-ch
	if err != io.EOF {
		t.Errorf("Expected io.EOF, but got: %s", err)
	}

	// Progress channel should be closed after progress reader is marked done
	_, ok := <-ch
	if ok {
		t.Errorf("Expected channel to be closed")
	}
}

// Read may be called by the http transport after Done.
func TestReaderReadAfterDone(t *testing.T) {
	s := "helloworld"
	ch := make(chan Report, 2)
	pr := NewReader(context.Background(), &dummySinker{ch}, strings.NewReader(s), int64(len(s)))

	pr.Done(nil)
	pr.Done(nil) // must not panic on double close

	var buf [10]byte
	n, err := pr.Read(buf[:]) // must not panic: send on closed channel
	if n != len(s) || err != nil {
		t.Errorf("n=%d err=%v", n, err)
	}
}

// Exercise Read and Done concurrently; run with -race.
func TestReaderConcurrentReadDone(t *testing.T) {
	for i := 0; i < 100; i++ {
		ch := make(chan Report, 1024)
		pr := NewReader(context.Background(), &dummySinker{ch}, io.LimitReader(zeroReader{}, 1<<20), 1<<20)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			var buf [512]byte
			for {
				if _, err := pr.Read(buf[:]); err != nil {
					return
				}
			}
		}()
		pr.Done(nil)
		wg.Wait()
	}
}

// Read errors must be returned without sending a report.
func TestReaderReadError(t *testing.T) {
	ch := make(chan Report, 1)
	pr := NewReader(context.Background(), &dummySinker{ch}, iotest.ErrReader(io.ErrUnexpectedEOF), 10)

	if _, err := pr.Read(make([]byte, 1)); err != io.ErrUnexpectedEOF {
		t.Errorf("err=%v", err)
	}
	select {
	case <-ch:
		t.Error("unexpected report")
	default:
	}
}

type zeroReader struct{}

func (zeroReader) Read(b []byte) (int, error) { clear(b); return len(b), nil }
