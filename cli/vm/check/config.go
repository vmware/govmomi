/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package check

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/types"
)

type config struct {
	checkFlag
}

func init() {
	cli.Register("vm.check.config", &config{}, true)
}

func (cmd *config) Description() string {
	return `Check if VM config spec can be applied.

Examples:
  govc vm.create -spec ... | govc vm.check.config -pool $pool`
}

func (cmd *config) Run(ctx context.Context, f *flag.FlagSet) error {
	var spec types.VirtualMachineConfigSpec

	if err := cmd.Spec(&spec); err != nil {
		return err
	}

	checker, err := cmd.compatChecker()
	if err != nil {
		return err
	}

	res, err := checker.CheckVmConfig(ctx, spec, cmd.Machine, cmd.Host, cmd.Pool, cmd.testTypes...)
	if err != nil {
		return err
	}

	return cmd.result(ctx, res)
}
