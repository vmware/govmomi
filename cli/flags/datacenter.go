// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/types"
)

type DatacenterFlag struct {
	common

	*ClientFlag
	*OutputFlag

	Name   string
	dc     *object.Datacenter
	finder *find.Finder
	err    error
}

var datacenterFlagKey = flagKey("datacenter")

func NewDatacenterFlag(ctx context.Context) (*DatacenterFlag, context.Context) {
	if v := ctx.Value(datacenterFlagKey); v != nil {
		return v.(*DatacenterFlag), ctx
	}

	v := &DatacenterFlag{}
	v.ClientFlag, ctx = NewClientFlag(ctx)
	v.OutputFlag, ctx = NewOutputFlag(ctx)
	ctx = context.WithValue(ctx, datacenterFlagKey, v)
	return v, ctx
}

func (flag *DatacenterFlag) Register(ctx context.Context, f *flag.FlagSet) {
	flag.RegisterOnce(func() {
		flag.ClientFlag.Register(ctx, f)
		flag.OutputFlag.Register(ctx, f)

		env := "GOVC_DATACENTER"
		value := os.Getenv(env)
		usage := fmt.Sprintf("Datacenter [%s]", env)
		f.StringVar(&flag.Name, "dc", value, usage)
	})
}

func (flag *DatacenterFlag) Process(ctx context.Context) error {
	return flag.ProcessOnce(func() error {
		if err := flag.ClientFlag.Process(ctx); err != nil {
			return err
		}
		if err := flag.OutputFlag.Process(ctx); err != nil {
			return err
		}
		return nil
	})
}

func (flag *DatacenterFlag) Finder(all ...bool) (*find.Finder, error) {
	if flag.finder != nil {
		return flag.finder, nil
	}

	c, err := flag.Client()
	if err != nil {
		return nil, err
	}

	allFlag := false
	if len(all) == 1 {
		allFlag = all[0]
	}
	finder := find.NewFinder(c, allFlag)

	// Datacenter is not required (ls command for example).
	// Set for relative func if dc flag is given or
	// if there is a single (default) Datacenter
	ctx := context.TODO()
	if flag.Name == "" {
		flag.dc, flag.err = finder.DefaultDatacenter(ctx)
	} else {
		if flag.dc, err = finder.Datacenter(ctx, flag.Name); err != nil {
			return nil, err
		}
	}

	finder.SetDatacenter(flag.dc)

	flag.finder = finder

	return flag.finder, nil
}

func (flag *DatacenterFlag) Datacenter() (*object.Datacenter, error) {
	if flag.dc != nil {
		return flag.dc, nil
	}

	_, err := flag.Finder()
	if err != nil {
		return nil, err
	}

	if flag.err != nil {
		// Should only happen if no dc is specified and len(dcs) > 1
		return nil, flag.err
	}

	return flag.dc, err
}

func (flag *DatacenterFlag) DatacenterIfSpecified() (*object.Datacenter, error) {
	if flag.Name == "" {
		return nil, nil
	}
	return flag.Datacenter()
}

func (flag *DatacenterFlag) ManagedObject(ctx context.Context, arg string) (types.ManagedObjectReference, error) {
	var ref types.ManagedObjectReference

	finder, err := flag.Finder()
	if err != nil {
		return ref, err
	}

	if ref.FromString(arg) {
		if strings.HasPrefix(ref.Type, "com.vmware.content.") {
			return ref, nil // special case for content library
		}
		pc := property.DefaultCollector(flag.client)
		var content []types.ObjectContent
		err = pc.RetrieveOne(ctx, ref, []string{"name"}, &content)
		if err == nil {
			return ref, nil
		}
	}

	l, err := finder.ManagedObjectList(ctx, arg)
	if err != nil {
		return ref, err
	}

	switch len(l) {
	case 0:
		return ref, fmt.Errorf("%s not found", arg)
	case 1:
		return l[0].Object.Reference(), nil
	default:
		var objs []types.ManagedObjectReference
		for _, o := range l {
			objs = append(objs, o.Object.Reference())
		}
		return ref, fmt.Errorf("%d objects match %q: %s (unique inventory path required)", len(l), arg, objs)
	}
}

func (flag *DatacenterFlag) ManagedObjects(ctx context.Context, args []string) ([]types.ManagedObjectReference, error) {
	var refs []types.ManagedObjectReference

	c, err := flag.Client()
	if err != nil {
		return nil, err
	}

	if len(args) == 0 {
		refs = append(refs, c.ServiceContent.RootFolder)
		return refs, nil
	}

	finder, err := flag.Finder()
	if err != nil {
		return nil, err
	}

	for _, arg := range args {
		elements, err := finder.ManagedObjectList(ctx, arg)
		if err != nil {
			return nil, err
		}

		if len(elements) == 0 {
			return nil, fmt.Errorf("object '%s' not found", arg)
		}

		if len(elements) > 1 && !strings.Contains(arg, "/") {
			return nil, fmt.Errorf("%q must be qualified with a path", arg)
		}

		for _, e := range elements {
			refs = append(refs, e.Object.Reference())
		}
	}

	return refs, nil
}
