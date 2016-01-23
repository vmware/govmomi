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

package object

import (
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

// HistoryCollector represents a history collector client
type HistoryCollector struct {
	Common
}

// NewHistoryCollector creates a new history collector client instance
func NewHistoryCollector(c *vim25.Client, ref types.ManagedObjectReference) *HistoryCollector {
	return &HistoryCollector{
		Common: NewCommon(c, ref),
	}
}

// Destroy this history collector
func (h HistoryCollector) Destroy(ctx context.Context) error {
	req := types.DestroyCollector{
		This: h.Reference(),
	}

	_, err := methods.DestroyCollector(ctx, h.c, &req)
	return err
}

// Reset this history collector
func (h HistoryCollector) Reset(ctx context.Context) error {
	req := types.ResetCollector{
		This: h.Reference(),
	}

	_, err := methods.ResetCollector(ctx, h.c, &req)
	return err
}

// Rewind this history collector
func (h HistoryCollector) Rewind(ctx context.Context) error {
	req := types.RewindCollector{
		This: h.Reference(),
	}

	_, err := methods.RewindCollector(ctx, h.c, &req)
	return err
}

// SetPageSize for this history collector
func (h HistoryCollector) SetPageSize(ctx context.Context, maxCount int) error {
	req := types.SetCollectorPageSize{
		This:     h.Reference(),
		MaxCount: maxCount,
	}

	_, err := methods.SetCollectorPageSize(ctx, h.c, &req)
	return err
}
