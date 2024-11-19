/*
Copyright (c) 2023-2023 VMware, Inc. All Rights Reserved.

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

package target

import (
	"context"
	"flag"
	"fmt"
	"io"
	"reflect"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type capls struct {
	flags.EnvBrowser
}

func init() {
	cli.Register("vm.target.cap.ls", &capls{})
}

func (cmd *capls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.EnvBrowser.Register(ctx, f)
}

func (cmd *capls) Description() string {
	return `List VM config target capabilities.

The config target data contains capabilities about the execution environment for a VM
in the given CLUSTER, and optionally for a specific HOST.

Examples:
  govc vm.target.cap.ls -cluster C0
  govc vm.target.cap.ls -host my_hostname
  govc vm.target.cap.ls -vm my_vm`
}

func (cmd *capls) Run(ctx context.Context, f *flag.FlagSet) error {
	b, err := cmd.Browser(ctx)
	if err != nil {
		return err
	}

	host, err := cmd.HostSystemIfSpecified()
	if err != nil {
		return err
	}

	cap, err := b.QueryTargetCapabilities(ctx, host)
	if err != nil {
		return err
	}

	return cmd.VirtualMachineFlag.WriteResult(&caplsResult{cap})
}

type caplsResult struct {
	*types.HostCapability
}

func (r *caplsResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	cap := reflect.ValueOf(r.HostCapability).Elem()
	kind := cap.Type()

	for i := 0; i < cap.NumField(); i++ {
		field := cap.Field(i)

		if kind.Field(i).Anonymous {
			continue
		}
		if field.Kind() == reflect.Pointer {
			if field.IsNil() {
				continue
			}
			field = field.Elem()
		}

		fmt.Fprintf(tw, "%s:\t%v\n", kind.Field(i).Name, field.Interface())
	}

	return tw.Flush()
}

func (r *caplsResult) Dump() interface{} {
	return r.HostCapability
}
