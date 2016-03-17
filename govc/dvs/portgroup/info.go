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

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type info struct {
	*flags.DatacenterFlag

	dvsPath  string
	dvpgPath string

	active     bool
	connected  bool
	inside     bool
	uplinkPort bool
	vlanId     int
	count      uint
}

func init() {
	cli.Register("dvs.portgroup.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.StringVar(&cmd.dvsPath, "dvs", "", "Distributed Virtual Switch path (required)")
	f.StringVar(&cmd.dvpgPath, "dvpg", "", "Distributed Virtual Portgroup path")

	f.BoolVar(&cmd.active, "active", false, "Filter by port active or inactive status")
	f.BoolVar(&cmd.connected, "connected", false, "Filter by port connected or disconnected status")
	f.BoolVar(&cmd.inside, "inside", true, "Filter by port inside or outside status")
	f.BoolVar(&cmd.uplinkPort, "uplinkPort", false, "Filter for uplink ports")
	f.IntVar(&cmd.vlanId, "vlanId", 0, "Filter by VLAN ID (0 = unfiltered)")
	f.UintVar(&cmd.count, "count", 0, "Number of matches to return (0 = unlimited)")
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	if len(cmd.dvsPath) == 0 {
		return flag.ErrHelp
	}

	client, err := cmd.Client()
	if err != nil {
		return err
	}

	// Find the DVS
	dvsFinder, err := cmd.Finder()
	if err != nil {
		return err
	}

	dvsNet, err := dvsFinder.Network(ctx, cmd.dvsPath)
	if err != nil {
		return err
	}

	dvs, ok := dvsNet.(*object.DistributedVirtualSwitch)
	if !ok {
		return fmt.Errorf("%s (%T) is not of type %T", cmd.dvsPath, dvsNet, dvs)
	}

	// Set base search criteria
	criteria := types.DistributedVirtualSwitchPortCriteria{
		Connected:  types.NewBool(cmd.connected),
		Active:     types.NewBool(cmd.active),
		UplinkPort: types.NewBool(cmd.uplinkPort),
		Inside:     types.NewBool(cmd.inside),
	}

	// If a distributed virtual portgroup path is set, then add its portgroup key to the base criteria
	if len(cmd.dvpgPath) > 0 {
		// This creates a new recursive finder to populate the properties of the managed object
		dvpgFinder := find.NewFinder(client, true)

		es, err := dvpgFinder.ManagedObjectListChildren(context.TODO(), cmd.dvpgPath)
		if err != nil {
			return err
		}

		dvpgObj := es[0].Object.(mo.DistributedVirtualPortgroup)
		criteria.PortgroupKey = []string{dvpgObj.Config.Key}
	}

	// Prepare request
	req := types.FetchDVPorts{
		This:     dvs.Reference(),
		Criteria: &criteria,
	}

	// Fetch ports
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
