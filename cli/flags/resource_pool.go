// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
)

type ResourcePoolFlag struct {
	common

	*DatacenterFlag

	name string
	pool *object.ResourcePool
}

var resourcePoolFlagKey = flagKey("resourcePool")

func NewResourcePoolFlag(ctx context.Context) (*ResourcePoolFlag, context.Context) {
	if v := ctx.Value(resourcePoolFlagKey); v != nil {
		return v.(*ResourcePoolFlag), ctx
	}

	v := &ResourcePoolFlag{}
	v.DatacenterFlag, ctx = NewDatacenterFlag(ctx)
	ctx = context.WithValue(ctx, resourcePoolFlagKey, v)
	return v, ctx
}

func (flag *ResourcePoolFlag) Register(ctx context.Context, f *flag.FlagSet) {
	flag.RegisterOnce(func() {
		flag.DatacenterFlag.Register(ctx, f)

		env := "GOVC_RESOURCE_POOL"
		value := os.Getenv(env)
		usage := fmt.Sprintf("Resource pool [%s]", env)
		f.StringVar(&flag.name, "pool", value, usage)
	})
}

func (flag *ResourcePoolFlag) Process(ctx context.Context) error {
	return flag.ProcessOnce(func() error {
		if err := flag.DatacenterFlag.Process(ctx); err != nil {
			return err
		}
		return nil
	})
}

func (flag *ResourcePoolFlag) IsSet() bool {
	return flag.name != ""
}

func (flag *ResourcePoolFlag) ResourcePool() (*object.ResourcePool, error) {
	if flag.pool != nil {
		return flag.pool, nil
	}

	finder, err := flag.Finder()
	if err != nil {
		return nil, err
	}

	flag.pool, err = finder.ResourcePoolOrDefault(context.TODO(), flag.name)
	if err != nil {
		if _, ok := err.(*find.NotFoundError); ok {
			vapp, verr := finder.VirtualApp(context.TODO(), flag.name)
			if verr != nil {
				return nil, err
			}
			flag.pool = vapp.ResourcePool
		} else {
			return nil, err
		}
	}

	return flag.pool, nil
}

func (flag *ResourcePoolFlag) ResourcePoolIfSpecified() (*object.ResourcePool, error) {
	if flag.name == "" {
		return nil, nil
	}
	return flag.ResourcePool()
}
