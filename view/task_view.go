// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package view

import (
	"context"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/types"
)

// TaskView extends ListView such that it can follow a ManagedEntity's recentTask updates.
type TaskView struct {
	*ListView

	Follow bool

	Watch *types.ManagedObjectReference
}

// CreateTaskView creates a new ListView that optionally watches for a ManagedEntity's recentTask updates.
func (m Manager) CreateTaskView(ctx context.Context, watch *types.ManagedObjectReference) (*TaskView, error) {
	l, err := m.CreateListView(ctx, nil)
	if err != nil {
		return nil, err
	}

	tv := &TaskView{
		ListView: l,
		Watch:    watch,
	}

	return tv, nil
}

// Collect calls function f for each Task update.
func (v TaskView) Collect(ctx context.Context, f func([]types.TaskInfo)) error {
	// Using TaskHistoryCollector would be less clunky, but it isn't supported on ESX at all.
	ref := v.Reference()
	filter := new(property.WaitFilter).Add(ref, "Task", []string{"info"}, v.TraversalSpec())

	if v.Watch != nil {
		filter.Add(*v.Watch, v.Watch.Type, []string{"recentTask"})
	}

	pc := property.DefaultCollector(v.Client())

	completed := make(map[string]bool)

	return property.WaitForUpdates(ctx, pc, filter, func(updates []types.ObjectUpdate) bool {
		var infos []types.TaskInfo
		var prune []types.ManagedObjectReference
		var tasks []types.ManagedObjectReference
		var reset func()

		for _, update := range updates {
			for _, change := range update.ChangeSet {
				if change.Name == "recentTask" {
					tasks = change.Val.(types.ArrayOfManagedObjectReference).ManagedObjectReference
					if len(tasks) != 0 {
						reset = func() {
							_, _ = v.Reset(ctx, tasks)

							// Remember any tasks we've reported as complete already,
							// to avoid reporting multiple times when Reset is triggered.
							rtasks := make(map[string]bool)
							for i := range tasks {
								if _, ok := completed[tasks[i].Value]; ok {
									rtasks[tasks[i].Value] = true
								}
							}
							completed = rtasks
						}
					}

					continue
				}

				info, ok := change.Val.(types.TaskInfo)
				if !ok {
					continue
				}

				if !completed[info.Task.Value] {
					infos = append(infos, info)
				}

				if v.Follow && info.CompleteTime != nil {
					prune = append(prune, info.Task)
					completed[info.Task.Value] = true
				}
			}
		}

		if len(infos) != 0 {
			f(infos)
		}

		if reset != nil {
			reset()
		} else if len(prune) != 0 {
			_, _ = v.Remove(ctx, prune)
		}

		if len(tasks) != 0 && len(infos) == 0 {
			return false
		}

		return !v.Follow
	})
}
