/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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

package event

import (
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

// HistoryCollector collects event history
type HistoryCollector struct {
	*object.HistoryCollector
}

// NewHistoryCollector creates a new event history collector
func NewHistoryCollector(c *vim25.Client, ref types.ManagedObjectReference) *HistoryCollector {
	return &HistoryCollector{
		HistoryCollector: object.NewHistoryCollector(c, ref),
	}
}

// LatestPage gets the latest history page
func (h HistoryCollector) LatestPage(ctx context.Context) ([]types.BaseEvent, error) {
	var o mo.EventHistoryCollector

	err := h.Properties(ctx, h.Reference(), []string{"latestPage"}, &o)
	if err != nil {
		return nil, err
	}

	return o.LatestPage, nil
}

// ReadNextEvents reads the next events
func (h HistoryCollector) ReadNextEvents(ctx context.Context, maxCount int) ([]types.BaseEvent, error) {
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

// ReadPreviousEvents reads the previous event up to maxCount
func (h HistoryCollector) ReadPreviousEvents(ctx context.Context, maxCount int) ([]types.BaseEvent, error) {
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
