// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package history

import (
	"context"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type Collector struct {
	r types.ManagedObjectReference
	c *vim25.Client
}

func NewCollector(c *vim25.Client, ref types.ManagedObjectReference) *Collector {
	return &Collector{
		r: ref,
		c: c,
	}
}

// Reference returns the managed object reference of this collector
func (c Collector) Reference() types.ManagedObjectReference {
	return c.r
}

// Client returns the vim25 client used by this collector
func (c Collector) Client() *vim25.Client {
	return c.c
}

// Properties wraps property.DefaultCollector().RetrieveOne() and returns
// properties for the specified managed object reference
func (c Collector) Properties(ctx context.Context, r types.ManagedObjectReference, ps []string, dst any) error {
	return property.DefaultCollector(c.c).RetrieveOne(ctx, r, ps, dst)
}

func (c Collector) Destroy(ctx context.Context) error {
	req := types.DestroyCollector{
		This: c.r,
	}

	_, err := methods.DestroyCollector(ctx, c.c, &req)
	return err
}

func (c Collector) Reset(ctx context.Context) error {
	req := types.ResetCollector{
		This: c.r,
	}

	_, err := methods.ResetCollector(ctx, c.c, &req)
	return err
}

func (c Collector) Rewind(ctx context.Context) error {
	req := types.RewindCollector{
		This: c.r,
	}

	_, err := methods.RewindCollector(ctx, c.c, &req)
	return err
}

func (c Collector) SetPageSize(ctx context.Context, maxCount int32) error {
	req := types.SetCollectorPageSize{
		This:     c.r,
		MaxCount: maxCount,
	}

	_, err := methods.SetCollectorPageSize(ctx, c.c, &req)
	return err
}
