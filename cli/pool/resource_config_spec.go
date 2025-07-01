// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package pool

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

func NewResourceConfigSpecFlag() *ResourceConfigSpecFlag {
	return &ResourceConfigSpecFlag{types.DefaultResourceConfigSpec(), nil}
}

type ResourceConfigSpecFlag struct {
	types.ResourceConfigSpec
	*flags.ResourceAllocationFlag
}

func (s *ResourceConfigSpecFlag) Register(ctx context.Context, f *flag.FlagSet) {
	s.ResourceAllocationFlag = flags.NewResourceAllocationFlag(&s.CpuAllocation, &s.MemoryAllocation)
	s.ResourceAllocationFlag.Register(ctx, f)
}

func (s *ResourceConfigSpecFlag) Process(ctx context.Context) error {
	return nil
}
