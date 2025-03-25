// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/vmware/govmomi/object"
)

type VirtualAppFlag struct {
	common

	*DatacenterFlag
	*SearchFlag

	name string
	app  *object.VirtualApp
}

var virtualAppFlagKey = flagKey("virtualApp")

func NewVirtualAppFlag(ctx context.Context) (*VirtualAppFlag, context.Context) {
	if v := ctx.Value(virtualAppFlagKey); v != nil {
		return v.(*VirtualAppFlag), ctx
	}

	v := &VirtualAppFlag{}
	v.DatacenterFlag, ctx = NewDatacenterFlag(ctx)
	v.SearchFlag, ctx = NewSearchFlag(ctx, SearchVirtualApps)
	ctx = context.WithValue(ctx, virtualAppFlagKey, v)
	return v, ctx
}

func (flag *VirtualAppFlag) Register(ctx context.Context, f *flag.FlagSet) {
	flag.RegisterOnce(func() {
		flag.DatacenterFlag.Register(ctx, f)
		flag.SearchFlag.Register(ctx, f)

		env := "GOVC_VAPP"
		value := os.Getenv(env)
		usage := fmt.Sprintf("Virtual App [%s]", env)
		f.StringVar(&flag.name, "vapp", value, usage)
	})
}

func (flag *VirtualAppFlag) Process(ctx context.Context) error {
	return flag.ProcessOnce(func() error {
		if err := flag.DatacenterFlag.Process(ctx); err != nil {
			return err
		}
		if err := flag.SearchFlag.Process(ctx); err != nil {
			return err
		}
		return nil
	})
}

func (flag *VirtualAppFlag) VirtualApp() (*object.VirtualApp, error) {
	ctx := context.TODO()

	if flag.app != nil {
		return flag.app, nil
	}

	// Use search flags if specified.
	if flag.SearchFlag.IsSet() {
		app, err := flag.SearchFlag.VirtualApp()
		if err != nil {
			return nil, err
		}

		flag.app = app
		return flag.app, nil
	}

	// Never look for a default virtual app.
	if flag.name == "" {
		return nil, nil
	}

	finder, err := flag.Finder()
	if err != nil {
		return nil, err
	}

	flag.app, err = finder.VirtualApp(ctx, flag.name)
	return flag.app, err
}
