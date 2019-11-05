/*
Copyright (c) 2015-2017 VMware, Inc. All Rights Reserved.

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

package dvs

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type change struct {
	*flags.FolderFlag

	types.DVSCreateSpec

	configSpec *types.VMwareDVSConfigSpec
}

func init() {
	cli.Register("dvs.change", &change{})
}

func (cmd *change) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)

	cmd.configSpec = new(types.VMwareDVSConfigSpec)

	cmd.DVSCreateSpec.ConfigSpec = cmd.configSpec
	cmd.DVSCreateSpec.ProductInfo = new(types.DistributedVirtualSwitchProductSpec)

	f.StringVar(&cmd.ProductInfo.Version, "product-version", "", "DVS product version")
	f.Var(flags.NewInt32(&cmd.configSpec.MaxMtu), "mtu", "DVS Max MTU")
}

func (cmd *change) Usage() string {
	return "DVS"
}

func (cmd *change) Description() string {
	return `Change DVS (DistributedVirtualSwitch) in datacenter.

Examples:
  govc dvs.change
  govc dvs.change -product-version 5.5.0 DSwitch
  govc dvs.change -mtu 9000 DSwitch`
}

func (cmd *change) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *change) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	name := f.Arg(0)

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	folder, err := cmd.FolderOrDefault("network")
	if err != nil {
		return err
	}

	networks, err := finder.NetworkList(ctx, folder.InventoryPath+"/"+name)
	if err != nil {
		return err
	}

	for _, net := range networks {
		if dvs, ok := net.(*object.DistributedVirtualSwitch); ok {
			var s mo.DistributedVirtualSwitch
			err = dvs.Properties(ctx, dvs.Reference(), []string{"config"}, &s)
			if err != nil {
				return err
			}

			cmd.configSpec.ConfigVersion = s.Config.GetDVSConfigInfo().ConfigVersion
			task, err := dvs.Reconfigure(ctx, cmd.ConfigSpec)
			if err != nil {
				return err
			}

			logger := cmd.ProgressLogger(fmt.Sprintf("updating %s in folder %s... ", name, folder.InventoryPath))
			defer logger.Wait()

			_, err = task.WaitForResult(ctx, logger)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
