// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package task

import (
	"context"

	"github.com/vmware/govmomi/history"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

// HistoryCollector provides a mechanism for retrieving historical data and
// updates when the server appends new tasks.
type HistoryCollector struct {
	*history.Collector
}

func newHistoryCollector(c *vim25.Client, ref types.ManagedObjectReference) *HistoryCollector {
	return &HistoryCollector{
		Collector: history.NewCollector(c, ref),
	}
}

// LatestPage returns items in the 'viewable latest page' of the task history collector.
// As new tasks that match the collector's TaskFilterSpec are created,
// they are added to this page, and the oldest tasks are removed from the collector to keep
// the size of the page to that allowed by SetCollectorPageSize.
// The "oldest task" is the one with the oldest creation time. The tasks in the returned page are unordered.
func (h HistoryCollector) LatestPage(ctx context.Context) ([]types.TaskInfo, error) {
	var o mo.TaskHistoryCollector

	err := h.Properties(ctx, h.Reference(), []string{"latestPage"}, &o)
	if err != nil {
		return nil, err
	}

	return o.LatestPage, nil
}

// ReadNextTasks reads the scrollable view from the current position. The
// scrollable position is moved to the next newer page after the read. No item
// is returned when the end of the collector is reached.
func (h HistoryCollector) ReadNextTasks(ctx context.Context, maxCount int32) ([]types.TaskInfo, error) {
	req := types.ReadNextTasks{
		This:     h.Reference(),
		MaxCount: maxCount,
	}

	res, err := methods.ReadNextTasks(ctx, h.Client(), &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

// ReadPreviousTasks reads the scrollable view from the current position. The
// scrollable position is then moved to the next older page after the read. No
// item is returned when the head of the collector is reached.
func (h HistoryCollector) ReadPreviousTasks(ctx context.Context, maxCount int32) ([]types.TaskInfo, error) {
	req := types.ReadPreviousTasks{
		This:     h.Reference(),
		MaxCount: maxCount,
	}

	res, err := methods.ReadPreviousTasks(ctx, h.Client(), &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}
