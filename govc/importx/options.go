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

package importx

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/ovf"
	"github.com/vmware/govmomi/vim25/types"
)

type Property struct {
	types.KeyValue
	Spec *ovf.Property `json:",omitempty"`
}

type Network struct {
	Name    string
	Network string
}

type Options struct {
	AllDeploymentOptions []string `json:",omitempty"`
	Deployment           string   `json:",omitempty"`

	AllDiskProvisioningOptions []string `json:",omitempty"`
	DiskProvisioning           string

	AllIPAllocationPolicyOptions []string `json:",omitempty"`
	IPAllocationPolicy           string

	AllIPProtocolOptions []string `json:",omitempty"`
	IPProtocol           string

	PropertyMapping []Property `json:",omitempty"`

	NetworkMapping []Network `json:",omitempty"`

	Annotation string `json:",omitempty"`

	MarkAsTemplate bool
	PowerOn        bool
	InjectOvfEnv   bool
	WaitForIP      bool
	Name           *string
}

type OptionsFlag struct {
	Options Options

	path string
}

func newOptionsFlag(ctx context.Context) (*OptionsFlag, context.Context) {
	return &OptionsFlag{}, ctx
}

func (flag *OptionsFlag) Register(ctx context.Context, f *flag.FlagSet) {
	f.StringVar(&flag.path, "options", "", "Options spec file path for VM deployment")
}

func (flag *OptionsFlag) Process(ctx context.Context) error {
	if len(flag.path) == 0 {
		return nil
	}

	var err error
	in := os.Stdin

	if flag.path != "-" {
		in, err = os.Open(flag.path)
		if err != nil {
			return err
		}
		defer in.Close()
	}

	return json.NewDecoder(in).Decode(&flag.Options)
}

func (flag *OptionsFlag) powerOn(vm *object.VirtualMachine, out *flags.OutputFlag) error {
	if !flag.Options.PowerOn || flag.Options.MarkAsTemplate {
		return nil
	}

	out.Log("Powering on VM...\n")

	task, err := vm.PowerOn(context.Background())
	if err != nil {
		return err
	}

	return task.Wait(context.Background())
}

func (flag *OptionsFlag) markAsTemplate(vm *object.VirtualMachine, out *flags.OutputFlag) error {
	if !flag.Options.MarkAsTemplate {
		return nil
	}

	out.Log("Marking VM as template...\n")

	return vm.MarkAsTemplate(context.Background())
}

func (flag *OptionsFlag) injectOvfEnv(vm *object.VirtualMachine, out *flags.OutputFlag) error {
	if !flag.Options.InjectOvfEnv {
		return nil
	}

	out.Log("Injecting OVF environment...\n")

	var opts []types.BaseOptionValue

	a := vm.Client().ServiceContent.About

	// build up Environment in order to marshal to xml
	var props []ovf.EnvProperty
	for _, p := range flag.Options.PropertyMapping {
		props = append(props, ovf.EnvProperty{
			Key:   p.Key,
			Value: p.Value,
		})
	}

	env := ovf.Env{
		EsxID: vm.Reference().Value,
		Platform: &ovf.PlatformSection{
			Kind:    a.Name,
			Version: a.Version,
			Vendor:  a.Vendor,
			Locale:  "US",
		},
		Property: &ovf.PropertySection{
			Properties: props,
		},
	}

	opts = append(opts, &types.OptionValue{
		Key:   "guestinfo.ovfEnv",
		Value: env.MarshalManual(),
	})

	task, err := vm.Reconfigure(context.Background(), types.VirtualMachineConfigSpec{
		ExtraConfig: opts,
	})

	if err != nil {
		return err
	}

	return task.Wait(context.Background())
}

func (flag *OptionsFlag) waitForIP(vm *object.VirtualMachine, out *flags.OutputFlag) error {
	if !flag.Options.PowerOn || !flag.Options.WaitForIP || flag.Options.MarkAsTemplate {
		return nil
	}

	out.Log("Waiting for IP address...\n")
	ip, err := vm.WaitForIP(context.Background())
	if err != nil {
		return err
	}

	out.Log(fmt.Sprintf("Received IP address: %s\n", ip))
	return nil
}

func (flag *OptionsFlag) Deploy(vm *object.VirtualMachine, out *flags.OutputFlag) error {
	deploy := []func(*object.VirtualMachine, *flags.OutputFlag) error{
		flag.injectOvfEnv,
		flag.markAsTemplate,
		flag.powerOn,
		flag.waitForIP,
	}

	for _, step := range deploy {
		if err := step(vm, out); err != nil {
			return err
		}
	}

	return nil
}
