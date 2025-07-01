// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object_test

import (
	"context"
	"reflect"
	"sync"
	"testing"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func TestExtensionMangerUpdates(t *testing.T) {
	extension := types.Extension{
		Description: &types.Description{
			Label:   "govmomi-test",
			Summary: "Extension Manager test",
		},
		Key:                    t.Name(),
		Version:                "0.0.1",
		ShownInSolutionManager: types.NewBool(false),
	}

	description := extension.Description.GetDescription()

	f := func(item string) string {
		return (&mo.Field{Path: "extensionList", Key: extension.Key, Item: item}).String()
	}

	tests := []types.PropertyChange{
		{Name: f(""), Val: extension, Op: types.PropertyChangeOpAdd},
		{Name: f(""), Val: extension, Op: types.PropertyChangeOpAssign},
		{Name: f("description"), Val: *description, Op: types.PropertyChangeOpAssign},
		{Name: f("description.label"), Val: description.Label, Op: types.PropertyChangeOpAssign},
		{Name: f(""), Val: nil, Op: types.PropertyChangeOpRemove},
	}

	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		m := object.NewExtensionManager(c)
		pc := property.DefaultCollector(c)

		for _, test := range tests {
			t.Logf("%s: %s", test.Op, test.Name)
			update := make(chan bool)
			parked := sync.OnceFunc(func() { update <- true })

			var change *types.PropertyChange
			cb := func(p []types.PropertyChange) bool {
				parked()
				change = &p[0]
				if change.Op != test.Op {
					t.Logf("ignore: change Op=%s, test Op=%s", change.Op, test.Op)
					return false
				}
				return true
			}

			go func() {
				werr := property.Wait(ctx, pc, m.Reference(), []string{test.Name}, cb)
				if werr != nil {
					t.Log(werr)
				}
				update <- true
			}()
			<-update // wait until above go func is parked in WaitForUpdatesEx()

			switch test.Op {
			case types.PropertyChangeOpAdd:
				if err := m.Register(ctx, extension); err != nil {
					t.Fatal(err)
				}
			case types.PropertyChangeOpAssign:
				if err := m.Update(ctx, extension); err != nil {
					t.Fatal(err)
				}
			case types.PropertyChangeOpRemove:
				if err := m.Unregister(ctx, extension.Key); err != nil {
					t.Fatal(err)
				}
			}
			<-update // wait until update is received (cb returns true)

			if change == nil {
				t.Fatal("no change")
			}

			if change.Name != test.Name {
				t.Errorf("Name: %s", change.Name)
			}

			if change.Op != test.Op {
				t.Errorf("Op: %s", change.Op)
			}

			if !reflect.DeepEqual(change.Val, test.Val) {
				t.Errorf("change.Val: %#v", change.Val)
				t.Errorf("test.Val:   %#v", test.Val)
			}
		}
	})
}
