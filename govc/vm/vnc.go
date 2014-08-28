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

package vm

import (
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type vnc struct {
	*flags.VirtualMachineFlag

	Enable   bool
	Port     int
	Password string
}

func init() {
	cli.Register("vm.vnc", &vnc{})
}

func (cmd *vnc) Register(f *flag.FlagSet) {
	f.BoolVar(&cmd.Enable, "enable", true, "Enable VNC")
	f.IntVar(&cmd.Port, "port", -1, "VNC port")
	f.StringVar(&cmd.Password, "password", "", "VNC password")
}

func (cmd *vnc) Process() error {
	return nil
}

func (cmd *vnc) Run(f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	var mvm mo.VirtualMachine
	err = c.Properties(vm.ManagedObjectReference, nil, &mvm)
	if err != nil {
		return err
	}

	var mhs mo.HostSystem

	// TODO(PN): This call takes a looonnnggg time...
	err = c.Properties(*mvm.Runtime.Host, []string{"config"}, &mhs)
	if err != nil {
		return err
	}

	// Find IP address of the VM's host
	var ip string = "<unknown>"
	for _, nc := range mhs.Config.VirtualNicManagerInfo.NetConfig {
		if nc.NicType == "management" && len(nc.CandidateVnic) > 0 {
			ip = nc.CandidateVnic[0].Spec.Ip.IpAddress
			break
		}
	}

	curOptions := vncOptionsFromExtraConfig(mvm.Config.ExtraConfig)
	newOptions := make(vncOptions)

	if cmd.Enable {
		newOptions["enabled"] = "true"
		newOptions["port"] = fmt.Sprintf("%d", cmd.Port)
		newOptions["password"] = cmd.Password
	} else {
		newOptions["enabled"] = "false"
		newOptions["port"] = ""
		newOptions["password"] = ""
	}

	// Only update if configuration changed
	if !reflect.DeepEqual(curOptions, newOptions) {
		spec := types.VirtualMachineConfigSpec{
			ExtraConfig: newOptions.ToExtraConfig(),
		}

		task, err := vm.Reconfigure(c, spec)
		if err != nil {
			return err
		}

		err = task.Wait()
		if err != nil {
			return err
		}
	}

	if cmd.Enable {
		fmt.Printf("VNC is ENABLED at vnc://:%s@%s:%d\n", cmd.Password, ip, cmd.Port)
	} else {
		fmt.Printf("VNC is DISABLED\n")
	}

	return err
}

type vncOptions map[string]string

var vncPrefix = "RemoteDisplay.vnc."

func vncOptionsFromExtraConfig(ov []types.OptionValue) vncOptions {
	vo := make(vncOptions)
	for _, o := range ov {
		if strings.HasPrefix(o.Key, vncPrefix) {
			key := o.Key[len(vncPrefix):]
			if key != "key" {
				vo[key] = o.Value.(string)
			}
		}
	}
	return vo
}

func (vo vncOptions) ToExtraConfig() []types.OptionValue {
	ov := make([]types.OptionValue, 0, 0)
	for k, v := range vo {
		key := vncPrefix + k
		value := v

		o := types.OptionValue{
			Key:   key,
			Value: &value, // Pass pointer to avoid omitempty
		}

		ov = append(ov, o)
	}

	// Don't know how to deal with the key option, set it to be empty...
	o := types.OptionValue{
		Key:   vncPrefix + "key",
		Value: new(string), // Pass pointer to avoid omitempty
	}

	ov = append(ov, o)

	return ov
}
