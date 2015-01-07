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

type change struct {
	*flags.ResourcePoolFlag
	*ResourceConfigSpecFlag
	name string
}

func init() {
	spec := NewResourceConfigSpecFlag()
	cli.Register("pool.change", &change{ResourceConfigSpecFlag: spec})
}

func (cmd *change) Register(f *flag.FlagSet) {
	f.StringVar(&cmd.name, "name", "", "Resource pool name")
}

func (cmd *change) Process() error { return nil }

func (cmd *change) Run(f *flag.FlagSet) error {
	pool, err := cmd.ResourcePool()
	if err != nil {
		return err
	}

	cmd.SetAllocation(func(a *types.ResourceAllocationInfo) {
		if a.Shares.Level == "" {
			a.Shares = nil
		}
	})

	return pool.UpdateConfig(cmd.name, &cmd.ResourceConfigSpec)
}
