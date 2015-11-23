/*
Copyright (c) 2014-2015 VMware, Inc. All Rights Reserved.

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

package flags

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type DatacenterFlag struct {
	*ClientFlag
	*OutputFlag

	register sync.Once
	path     string
	dc       *object.Datacenter
	finder   *find.Finder
	err      error
}

func (flag *DatacenterFlag) Register(f *flag.FlagSet) {
	flag.register.Do(func() {
		env := "GOVC_DATACENTER"
		value := os.Getenv(env)
		usage := fmt.Sprintf("Datacenter [%s]", env)
		f.StringVar(&flag.path, "dc", value, usage)
	})
}

func (flag *DatacenterFlag) Process() error { return nil }

func (flag *DatacenterFlag) Finder() (*find.Finder, error) {
	if flag.finder != nil {
		return flag.finder, nil
	}

	c, err := flag.Client()
	if err != nil {
		return nil, err
	}

	finder := find.NewFinder(c, flag.JSON)

	// Datacenter is not required (ls command for example).
	// Set for relative func if dc flag is given or
	// if there is a single (default) Datacenter
	if flag.path == "" {
		flag.dc, flag.err = finder.DefaultDatacenter(context.TODO())
	} else {
		if flag.dc, err = finder.Datacenter(context.TODO(), flag.path); err != nil {
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

		for _, e := range elements {
			refs = append(refs, e.Object.Reference())
		}
	}

	return refs, nil
}
