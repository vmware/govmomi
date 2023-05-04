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

package pnic

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type pnicChange struct {
	*flags.HostSystemFlag

	chassisID  string
	portID     string
	systemName string
	nicName    string
}

func init() {
	cli.Register("host.pnic.lldp", &pnicChange{})
}

func (cmd *pnicChange) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	f.StringVar(&cmd.nicName, "nic", "vmnic0", "nic name")
	f.StringVar(&cmd.chassisID, "chassisID", "", "chassis id")
	f.StringVar(&cmd.portID, "portID", "", "port id")
	f.StringVar(&cmd.systemName, "systemName", "", "system name")
}

func (cmd *pnicChange) Process(ctx context.Context) error {
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *pnicChange) Usage() string {
	return "DEVICE"
}

func (cmd *pnicChange) Description() string {
	return `Connect pnic to lldp endpoint.

Examples:
  govc host.pnic.lldp -nic vmnic0 -chassisID 11:AB:CD:E1:23:12 -portID Eth12 -systemName LeafSwitch10 -host DC0_H0`
}

func (cmd *pnicChange) Connect(ctx context.Context, parent *object.HostNetworkSystem) error {
	var spec types.HostPnicLLDPEndpointSpec
	spec.ChassisID = cmd.chassisID
	spec.PortID = cmd.portID
	spec.NicName = cmd.nicName
	spec.SystemName = cmd.systemName

	req := types.ConnectPnicLLDPEndpoint_Task{
		This: parent.Reference(),
		Spec: spec,
	}

	res, err := methods.ConnectPnicLLDPEndpoint_Task(ctx, parent.Client(), &req)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("connect pnic to lldp %v\n", cmd.chassisID))
	defer logger.Wait()

	task := object.NewTask(parent.Client(), res.Returnval)
	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *pnicChange) Run(ctx context.Context, f *flag.FlagSet) error {
	fmt.Printf("args %v : %v : %v : %v\n", f.Arg(0), f.Arg(1), f.Arg(2), f.Arg(3))
	fmt.Printf("cmd %+v\n", cmd)
	hs, err := cmd.HostSystem()
	if err != nil {
		return err
	}
	fmt.Printf("host system %+v\n", hs)

	var host mo.HostSystem
	err = hs.Properties(ctx, hs.Reference(), nil, &host)
	if err != nil {
		return err
	}
	fmt.Printf("-----> host: %+v\n", host.Config.Network.Pnic)
	found := false
	for _, pnic := range host.Config.Network.Pnic {
		if pnic.Device == cmd.nicName {
			found = true
		}
	}
	if !found {
		return fmt.Errorf("nic %v is not preset in host %v", cmd.nicName, host)
	}

	ns, err := cmd.HostNetworkSystem()
	if err != nil {
		return err
	}

	return cmd.Connect(ctx, ns)
}
