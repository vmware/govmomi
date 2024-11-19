/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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

package service

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type rm struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("namespace.service.rm", &rm{})
}

func (cmd *rm) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
}

func (cmd *rm) Description() string {
	return `Removes a vSphere Namespace Supervisor Service.

Examples:
  govc namespace.service.rm my-supervisor-service other-supervisor-service`
}

func (cmd *rm) Usage() string {
	return "NAME..."
}

func (cmd *rm) Run(ctx context.Context, f *flag.FlagSet) error {

	services := f.Args()
	if len(services) < 1 {
		return fmt.Errorf("at least one service must be passed as argument")
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := namespace.NewManager(c)
	for _, svc := range services {

		if err := m.RemoveSupervisorService(ctx, svc); err != nil {
			return err
		}
	}
	return nil
}
