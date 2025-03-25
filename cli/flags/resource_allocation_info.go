// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"context"
	"flag"
	"strconv"
	"strings"

	"github.com/vmware/govmomi/vim25/types"
)

type sharesInfo types.SharesInfo

func (s *sharesInfo) String() string {
	return string(s.Level)
}

func (s *sharesInfo) Set(val string) error {
	switch val {
	case string(types.SharesLevelNormal), string(types.SharesLevelLow), string(types.SharesLevelHigh):
		s.Level = types.SharesLevel(val)
	default:
		n, err := strconv.Atoi(val)
		if err != nil {
			return err
		}

		s.Level = types.SharesLevelCustom
		s.Shares = int32(n)
	}

	return nil
}

type ResourceAllocationFlag struct {
	cpu, mem              *types.ResourceAllocationInfo
	ExpandableReservation bool
}

func NewResourceAllocationFlag(cpu, mem *types.ResourceAllocationInfo) *ResourceAllocationFlag {
	return &ResourceAllocationFlag{cpu, mem, true}
}

func (r *ResourceAllocationFlag) Register(ctx context.Context, f *flag.FlagSet) {
	opts := []struct {
		name  string
		units string
		*types.ResourceAllocationInfo
	}{
		{"CPU", "MHz", r.cpu},
		{"Memory", "MB", r.mem},
	}

	for _, opt := range opts {
		prefix := strings.ToLower(opt.name)[:3]
		shares := (*sharesInfo)(opt.Shares)

		f.Var(NewOptionalInt64(&opt.Limit), prefix+".limit", opt.name+" limit in "+opt.units)
		f.Var(NewOptionalInt64(&opt.Reservation), prefix+".reservation", opt.name+" reservation in "+opt.units)
		if r.ExpandableReservation {
			f.Var(NewOptionalBool(&opt.ExpandableReservation), prefix+".expandable", opt.name+" expandable reservation")
		}
		f.Var(shares, prefix+".shares", opt.name+" shares level or number")
	}
}

func (s *ResourceAllocationFlag) Process(ctx context.Context) error {
	return nil
}
