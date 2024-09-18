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

package kms

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/crypto"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type setdefault struct {
	*flags.DatacenterFlag

	entity string
}

func init() {
	cli.Register("kms.default", &setdefault{})
}

func (cmd *setdefault) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.StringVar(&cmd.entity, "e", "", "Set entity default KMS cluster (cluster or host folder)")
}

func (cmd *setdefault) Usage() string {
	return "NAME"
}

func (cmd *setdefault) Description() string {
	return `Set default KMS cluster.

Examples:
  govc kms.default my-kp
  govc kms.default - # clear default
  govc kms.default -e /dc/host/cluster my-kp
  govc kms.default -e /dc/host/cluster my-kp - # clear default`
}

func (cmd *setdefault) Run(ctx context.Context, f *flag.FlagSet) error {
	id := f.Arg(0)
	if id == "" {
		return flag.ErrHelp
	}

	if id == "-" {
		id = "" // clear default
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := crypto.GetManagerKmip(c)
	if err != nil {
		return err
	}

	var entity *types.ManagedObjectReference

	if cmd.entity != "" {
		obj, err := cmd.ManagedObject(ctx, cmd.entity)
		if err != nil {
			return err
		}

		entity = types.NewReference(obj.Reference())
	}

	return m.SetDefaultKmsClusterId(ctx, id, entity)
}
