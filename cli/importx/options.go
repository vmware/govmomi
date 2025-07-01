// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package importx

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/ovf"
	"github.com/vmware/govmomi/ovf/importer"
	"github.com/vmware/govmomi/vim25/types"
)

type OptionsFlag struct {
	Options importer.Options

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
