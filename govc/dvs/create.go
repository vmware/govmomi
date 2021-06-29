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

package dvs

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type create struct {
	*flags.FolderFlag

	types.DVSCreateSpec

	configSpec *types.VMwareDVSConfigSpec

	dProtocol string

	numUplinkPorts int32
}

func init() {
	cli.Register("dvs.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)

	cmd.configSpec = new(types.VMwareDVSConfigSpec)

	cmd.DVSCreateSpec.ConfigSpec = cmd.configSpec
	cmd.DVSCreateSpec.ProductInfo = new(types.DistributedVirtualSwitchProductSpec)

	cmd.numUplinkPorts = 2 // Set Default Uplink ports to 2

	f.StringVar(&cmd.ProductInfo.Version, "product-version", "", "DVS product version")
	f.Var(flags.NewInt32(&cmd.configSpec.MaxMtu), "mtu", "DVS Max MTU")
	f.StringVar(&cmd.dProtocol, "discovery-protocol", "", "Link Discovery Protocol")
	f.Var(flags.NewInt32(&cmd.numUplinkPorts), "num-uplinks", "Number of Uplinks")
}

func (cmd *create) Usage() string {
	return "DVS"
}

func (cmd *create) Description() string {
	return `Create DVS (DistributedVirtualSwitch) in datacenter.

The dvs is added to the folder specified by the 'folder' flag. If not given,
this defaults to the network folder in the specified or default datacenter.

Examples:
  govc dvs.create DSwitch
  govc dvs.create -product-version 5.5.0 DSwitch
  govc dvs.create -mtu 9000 DSwitch
  govc dvs.create -discovery-protocol [lldp|cdp] DSwitch`
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.FolderFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	name := f.Arg(0)

	cmd.configSpec.Name = name

	if cmd.dProtocol != "" {
		cmd.configSpec.LinkDiscoveryProtocolConfig = &types.LinkDiscoveryProtocolConfig{
			Protocol:  cmd.dProtocol,
			Operation: "listen",
		}
	}

	numUplinkPorts := int(cmd.numUplinkPorts)

	if numUplinkPorts > 0 {
		var policy types.DVSNameArrayUplinkPortPolicy
		for i := 0; i < numUplinkPorts; i++ {
			policy.UplinkPortName = append(policy.UplinkPortName, fmt.Sprintf("Uplink %d", i+1))
		}
		cmd.configSpec.UplinkPortPolicy = &policy
	} else {
		return errors.New("Must specify more than one uplink port")
	}

	folder, err := cmd.FolderOrDefault("network")
	if err != nil {
		return err
	}

	task, err := folder.CreateDVS(ctx, cmd.DVSCreateSpec)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("adding %s to folder %s... ", name, folder.InventoryPath))
	defer logger.Wait()

	_, err = task.WaitForResult(ctx, logger)
	return err
}
