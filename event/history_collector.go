// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package event

import (
	"context"

	"github.com/vmware/govmomi/history"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type HistoryCollector struct {
	*history.Collector
}

func newHistoryCollector(c *vim25.Client, ref types.ManagedObjectReference) *HistoryCollector {
	return &HistoryCollector{
		Collector: history.NewCollector(c, ref),
	}
}

func (h HistoryCollector) LatestPage(ctx context.Context) ([]types.BaseEvent, error) {
	var o mo.EventHistoryCollector

	err := h.Properties(ctx, h.Reference(), []string{"latestPage"}, &o)
	if err != nil {
		return nil, err
	}

	return o.LatestPage, nil
}

func (h HistoryCollector) ReadNextEvents(ctx context.Context, maxCount int32) ([]types.BaseEvent, error) {
	req := types.ReadNextEvents{
		This:     h.Reference(),
		MaxCount: maxCount,
	}

	res, err := methods.ReadNextEvents(ctx, h.Client(), &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (h HistoryCollector) ReadPreviousEvents(ctx context.Context, maxCount int32) ([]types.BaseEvent, error) {
	req := types.ReadPreviousEvents{
		This:     h.Reference(),
		MaxCount: maxCount,
	}

	res, err := methods.ReadPreviousEvents(ctx, h.Client(), &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}
