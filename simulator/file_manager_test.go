// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"path"
	"sync"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25/types"
)

// newCopyTestClient starts an ESX simulator and returns a client for it. The
// returned func tears the simulator down.
func newCopyTestClient(t *testing.T) (*govmomi.Client, func()) {
	t.Helper()

	m := ESX()
	if err := m.Create(); err != nil {
		t.Fatal(err)
	}

	s := m.Service.NewServer()
	ctx := context.Background()
	c, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		s.Close()
		m.Remove()
		t.Fatal(err)
	}

	if err = object.NewFileManager(c.Client).MakeDirectory(ctx, "[LocalDS_0] dst", nil, true); err != nil {
		s.Close()
		m.Remove()
		t.Fatal(err)
	}

	return c, func() {
		s.Close()
		m.Remove()
	}
}

// createCopyTestDisk creates a thin disk at name with the given capacity. The
// capacity differs between disks so that their contents differ.
func createCopyTestDisk(ctx context.Context, t *testing.T, c *govmomi.Client, name string, capacityKb int64) {
	t.Helper()

	fm := object.NewFileManager(c.Client)
	if err := fm.MakeDirectory(ctx, path.Dir(name), nil, true); err != nil {
		t.Fatal(err)
	}

	dm := object.NewVirtualDiskManager(c.Client)
	spec := &types.FileBackedVirtualDiskSpec{
		VirtualDiskSpec: types.VirtualDiskSpec{
			AdapterType: string(types.VirtualDiskAdapterTypeLsiLogic),
			DiskType:    string(types.VirtualDiskTypeThin),
		},
		CapacityKb: capacityKb,
	}

	task, err := dm.CreateVirtualDisk(ctx, name, nil, spec)
	if err != nil {
		t.Fatal(err)
	}
	if err = task.Wait(ctx); err != nil {
		t.Fatal(err)
	}
}

// copyTestDisk starts a non-force CopyVirtualDisk from src to dst and waits for it.
func copyTestDisk(ctx context.Context, c *govmomi.Client, src, dst string) error {
	dm := object.NewVirtualDiskManager(c.Client)

	task, err := dm.CopyVirtualDisk(ctx, src, nil, dst, nil, nil, false)
	if err != nil {
		return err
	}

	return task.Wait(ctx)
}

// isFileAlreadyExists reports whether err is a task fault of type FileAlreadyExists.
func isFileAlreadyExists(err error) bool {
	te, ok := err.(task.Error)
	if !ok {
		return false
	}

	_, ok = te.Fault().(*types.FileAlreadyExists)
	return ok
}

func TestCopyVirtualDiskConcurrentIdentical(t *testing.T) {
	ctx := context.Background()

	c, cleanup := newCopyTestClient(t)
	defer cleanup()

	src := "[LocalDS_0] src/src.vmdk"
	dst := "[LocalDS_0] dst/dst.vmdk"
	createCopyTestDisk(ctx, t, c, src, 1024*1024)

	// Several identical non-force copies to one destination, as vm-operator issues
	// when concurrent reconciles cache the same image disk. Each must succeed, and
	// the destination must end up with the source's content.
	const n = 4
	errs := make([]error, n)
	var wg sync.WaitGroup
	for i := range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			errs[i] = copyTestDisk(ctx, c, src, dst)
		}()
	}
	wg.Wait()

	for i, err := range errs {
		if err != nil {
			t.Errorf("copy %d: %v", i, err)
		}
	}

	// A later identical copy, after the racers finished, must also succeed.
	if err := copyTestDisk(ctx, c, src, dst); err != nil {
		t.Errorf("sequential identical copy: %v", err)
	}
}

func TestCopyVirtualDiskExistingDifferentFile(t *testing.T) {
	ctx := context.Background()

	c, cleanup := newCopyTestClient(t)
	defer cleanup()

	src := "[LocalDS_0] src/src.vmdk"
	other := "[LocalDS_0] other/other.vmdk"
	dst := "[LocalDS_0] dst/dst.vmdk"
	createCopyTestDisk(ctx, t, c, src, 1024*1024)
	createCopyTestDisk(ctx, t, c, other, 2*1024*1024)

	if err := copyTestDisk(ctx, c, other, dst); err != nil {
		t.Fatal(err)
	}

	// dst holds other's content, so a copy from src must not overwrite it.
	err := copyTestDisk(ctx, c, src, dst)
	if !isFileAlreadyExists(err) {
		t.Errorf("copy over a different existing file: got %v, want FileAlreadyExists", err)
	}

	// A copy from other, whose content dst already holds, still succeeds.
	if err := copyTestDisk(ctx, c, other, dst); err != nil {
		t.Errorf("identical copy over existing dst: %v", err)
	}
}

func TestCopyVirtualDiskConcurrentMixedSources(t *testing.T) {
	ctx := context.Background()

	c, cleanup := newCopyTestClient(t)
	defer cleanup()

	src := "[LocalDS_0] src/src.vmdk"
	other := "[LocalDS_0] other/other.vmdk"
	dst := "[LocalDS_0] dst/dst.vmdk"
	createCopyTestDisk(ctx, t, c, src, 1024*1024)
	createCopyTestDisk(ctx, t, c, other, 2*1024*1024)

	// Whichever source wins the destination, the losers from the other source
	// must get FileAlreadyExists, and the identical copies must succeed.
	type result struct {
		from string
		err  error
	}
	var (
		wg      sync.WaitGroup
		results = make([]result, 0, 6)
		mu      sync.Mutex
	)
	for _, from := range []string{src, src, src, other, other, other} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := copyTestDisk(ctx, c, from, dst)
			mu.Lock()
			results = append(results, result{from, err})
			mu.Unlock()
		}()
	}
	wg.Wait()

	var winner string
	for _, r := range results {
		if r.err == nil {
			if winner != "" && winner != r.from {
				t.Fatalf("copies from both %s and %s succeeded", winner, r.from)
			}
			winner = r.from
			continue
		}
		if !isFileAlreadyExists(r.err) {
			t.Errorf("copy from %s: %v", r.from, r.err)
		}
	}
	if winner == "" {
		t.Fatal("no copy succeeded")
	}
	for _, r := range results {
		if r.err != nil && r.from == winner {
			t.Errorf("identical copy from winner %s failed: %v", winner, r.err)
		}
	}
}
