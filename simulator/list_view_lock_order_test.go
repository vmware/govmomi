// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

// TestListViewModifyWaitForUpdatesLockOrder reproduces an ABBA deadlock between
// a ListView's object lock and a PropertyFilter's object lock, captured from a
// live vShim goroutine dump where a CSI driver watches CNS tasks through a
// ListView:
//
//   - WaitForUpdatesEx -> PropertyFilter.update (property_filter.go) takes the
//     filter's object lock, then PropertyFilter.collect traverses into the
//     ListView and takes the ListView's object lock. Order: filter -> ListView.
//   - ModifyListView (from another session) takes the ListView's object lock at
//     dispatch, then Registry.Update -> applyHandlers -> PropertyFilter.UpdateObject
//     takes the filter's object lock. Order: ListView -> filter.
//
// The two orders invert, so under concurrency they can each hold what the other
// wants and wedge forever. Every other ListView/PropertyCollector request then
// piles up behind the held ListView lock (in production: ~100 goroutines, CSI
// CreateVolume task completion never observed, PVC never binds, backup/restore
// hangs).
func TestListViewModifyWaitForUpdatesLockOrder(t *testing.T) {
	ctx := context.Background()

	m := VPX()
	defer m.Remove()
	if err := m.Create(); err != nil {
		t.Fatal(err)
	}

	s := m.Service.NewServer()
	defer s.Close()

	// One session (a CSI driver connection). The ListView is session-scoped;
	// the watch and the modify run as concurrent requests, so each gets its own
	// per-request *Context (the lock's SharedLockingContext) -- non-re-entrant.
	c, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	vmRef := m.Map().Any("VirtualMachine").Reference()

	lv, err := view.NewManager(c.Client).CreateListView(ctx, []types.ManagedObjectReference{vmRef})
	if err != nil {
		t.Fatal(err)
	}
	lvRef := lv.Reference()

	pc := property.DefaultCollector(c.Client)
	p, err := pc.Create(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Filter rooted at the ListView, traversing its members, so the filter's
	// refs include the ListView -- exactly the CSI task-watch shape.
	spec := types.PropertyFilterSpec{
		ObjectSet: []types.ObjectSpec{{
			Obj: lvRef,
			SelectSet: []types.BaseSelectionSpec{&types.TraversalSpec{
				Type: "ListView",
				Path: "view",
			}},
		}},
		PropSet: []types.PropertySpec{{Type: "VirtualMachine", PathSet: []string{"runtime.powerState"}}},
	}
	if _, err = methods.CreateFilter(ctx, c.Client, &types.CreateFilter{This: p.Reference(), Spec: spec}); err != nil {
		t.Fatal(err)
	}

	maxWait := int32(1)
	opts := &types.WaitOptions{MaxWaitSeconds: &maxWait}

	stop := make(chan struct{})
	var wg sync.WaitGroup

	// Watch: repeated WaitForUpdatesEx (filter -> ListView lock order).
	wg.Add(1)
	go func() {
		defer wg.Done()
		var version string
		for {
			select {
			case <-stop:
				return
			default:
			}
			res, werr := methods.WaitForUpdatesEx(ctx, c.Client, &types.WaitForUpdatesEx{
				This:    p.Reference(),
				Version: version,
				Options: opts,
			})
			if werr != nil {
				return
			}
			if res.Returnval != nil {
				version = res.Returnval.Version
			}
		}
	}()

	// Modify: hammer ModifyListView (ListView -> filter lock order).
	lvB := view.NewListView(c.Client, lvRef)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop:
				return
			default:
			}
			if _, rerr := lvB.Remove(ctx, []types.ManagedObjectReference{vmRef}); rerr != nil {
				return
			}
			if _, aerr := lvB.Add(ctx, []types.ManagedObjectReference{vmRef}); aerr != nil {
				return
			}
		}
	}()

	time.Sleep(3 * time.Second)
	close(stop)

	finished := make(chan struct{})
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case <-time.After(15 * time.Second):
		t.Fatal("deadlock: ModifyListView and WaitForUpdatesEx acquire the ListView and PropertyFilter object locks in opposite orders")
	}
}
