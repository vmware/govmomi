/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package pool

import (
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type create struct {
	*flags.ResourcePoolFlag
	*ResourceConfigSpecFlag
}

func init() {
	spec := NewResourceConfigSpecFlag()
	spec.SetAllocation(func(a *types.ResourceAllocationInfo) {
		a.Shares.Level = types.SharesLevelNormal
		a.ExpandableReservation = true
	})

	cli.Register("pool.create", &create{ResourceConfigSpecFlag: spec})
}

func (cmd *create) Register(f *flag.FlagSet) {}

func (cmd *create) Process() error { return nil }

func (cmd *create) Run(f *flag.FlagSet) error {
	parent, err := cmd.ResourcePool()
	if err != nil {
		return err
	}

	_, err = parent.Create(f.Arg(0), cmd.ResourceConfigSpec)
	return err
}
