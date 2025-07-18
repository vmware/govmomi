// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package nfc

import (
	"context"
	"log"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"github.com/vmware/govmomi/vim25/progress"
	"github.com/vmware/govmomi/vim25/types"
)

type FileItem struct {
	types.OvfFileItem

	URL        *url.URL
	Thumbprint string

	ch chan progress.Report
}

func NewFileItem(u *url.URL, item types.OvfFileItem) FileItem {
	return FileItem{
		OvfFileItem: item,
		URL:         u,
		ch:          make(chan progress.Report),
	}
}

func (o FileItem) Sink() chan<- progress.Report {
	return o.ch
}

// File converts the FileItem.OvfFileItem to an OvfFile
func (o FileItem) File() types.OvfFile {
	return types.OvfFile{
		DeviceId: o.DeviceId,
		Path:     o.Path,
		Size:     o.Size,
	}
}

type LeaseUpdater struct {
	pos   int64 // Number of bytes (keep first to ensure 64 bit alignment)
	total int64 // Total number of bytes (keep first to ensure 64 bit alignment)

	lease *Lease

	done chan struct{} // When lease updater should stop

	wg sync.WaitGroup // Track when update loop is done
}

func newLeaseUpdater(ctx context.Context, lease *Lease, info *LeaseInfo) *LeaseUpdater {
	l := LeaseUpdater{
		lease: lease,

		done: make(chan struct{}),
	}

	for _, item := range info.Items {
		l.total += item.Size
		go l.waitForProgress(item)
	}

	// Kickstart update loop
	l.wg.Add(1)
	go l.run()

	return &l
}

func (l *LeaseUpdater) waitForProgress(item FileItem) {
	var pos, total int64

	total = item.Size

	for {
		select {
		case <-l.done:
			return
		case p, ok := <-item.ch:
			// Return in case of error
			if ok && p.Error() != nil {
				return
			}

			if !ok {
				// Last element on the channel, add to total
				atomic.AddInt64(&l.pos, total-pos)
				return
			}

			// Approximate progress in number of bytes
			x := int64(float32(total) * (p.Percentage() / 100.0))
			atomic.AddInt64(&l.pos, x-pos)
			pos = x
		}
	}
}

func (l *LeaseUpdater) run() {
	defer l.wg.Done()

	tick := time.NewTicker(2 * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-l.done:
			return
		case <-tick.C:
			// From the vim api HttpNfcLeaseProgress(percent) doc, percent ==
			// "Completion status represented as an integer in the 0-100 range."
			// Always report the current value of percent, as it will renew the
			// lease even if the value hasn't changed or is 0.
			percent := int32(float32(100*atomic.LoadInt64(&l.pos)) / float32(l.total))
			err := l.lease.Progress(context.TODO(), percent)
			if err != nil {
				log.Printf("NFC lease progress: %s", err)
				return
			}
		}
	}
}

func (l *LeaseUpdater) Done() {
	close(l.done)
	l.wg.Wait()
}
