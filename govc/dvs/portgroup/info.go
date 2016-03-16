/*
Copyright (c) 2015-2016 VMware, Inc. All Rights Reserved.

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

package portgroup

import (
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type info struct {
	*flags.DatacenterFlag

	path string

	active       bool
	connected    bool
	inside       bool
	portgroupKey string
	portKey      string
	uplinkPort   bool
	vlanId       int
	count        uint
}

func init() {
	cli.Register("dvs.portgroup.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	//	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	//	cmd.ClientFlag.Register(ctx, f)

	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.StringVar(&cmd.path, "dvs", "", "DVS path")

	f.BoolVar(&cmd.active, "active", false, "If set, only the active ports are qualified")
	f.BoolVar(&cmd.connected, "connected", false, "If set, only the connected ports are qualified")
	f.BoolVar(&cmd.inside, "inside", false, "If unset, all ports in the switch are qualified. If set to true, only ports inside portgroupKey or any portgroup, if not set, are qualified. If set to false, only ports outside portgroupKey or any portgroup, if not set, are qualified")
	f.StringVar(&cmd.portgroupKey, "portgroupKey", "", "The keys of the portgroup that is used for the scope of inside. If this property is unset, it means any portgroup. If inside is unset, this property is ignored")
	f.StringVar(&cmd.portKey, "portKey", "", "If set, only the ports of which the key is in the array are qualified")
	f.BoolVar(&cmd.uplinkPort, "uplinkPort", false, "If set to true, only the uplink ports are qualified. If set to false, only non-uplink ports are qualified")
	f.IntVar(&cmd.vlanId, "vlanId", 0, "VLAN ID to filter")
	f.UintVar(&cmd.count, "count", 0, "Number of matches to return (0 = unlimited)")
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	portgroupKeySlice := []string{cmd.portgroupKey}
	portKeySlice := []string{cmd.portKey}

	criteria := types.DistributedVirtualSwitchPortCriteria{
		Connected:    types.NewBool(cmd.connected),
		Active:       types.NewBool(cmd.active),
		UplinkPort:   types.NewBool(cmd.uplinkPort),
		PortgroupKey: portgroupKeySlice,
		Inside:       types.NewBool(cmd.inside),
		PortKey:      portKeySlice,
	}

	client, err := cmd.Client()
	if err != nil {
		return err
	}

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	net, err := finder.Network(ctx, cmd.path)
	if err != nil {
		return err
	}

	dvs, ok := net.(*object.DistributedVirtualSwitch)
	if !ok {
		return fmt.Errorf("%s (%T) is not of type %T", cmd.path, net, dvs)
	}

	req := types.FetchDVPorts{
		This:     dvs.Reference(),
		Criteria: &criteria,
	}

	res, err := methods.FetchDVPorts(ctx, client, &req)
	if err != nil {
		return err
	}

	var returnedPorts uint = 0

	for _, port := range res.Returnval {

		portConfigSetting := *(port.Config.Setting.(*types.VMwareDVSPortSetting))
		portVlan := *(portConfigSetting.Vlan.(*types.VmwareDistributedVirtualSwitchVlanIdSpec))
		portVlanId := portVlan.VlanId

		if cmd.vlanId == 0 || portVlanId == cmd.vlanId {

			returnedPorts++

			fmt.Printf("PortgroupKey: %s\n", port.PortgroupKey)
			fmt.Printf("DvsUuid:      %s\n", port.DvsUuid)
			fmt.Printf("VlanId:       %d\n", portVlanId)
			fmt.Printf("PortKey:      %s\n\n", port.Key)

			if returnedPorts == cmd.count {
				break
			}
		}
	}

	return nil
}
