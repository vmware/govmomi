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

package flags

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/vmware/govmomi/object"
	"golang.org/x/net/context"
)

type VirtualAppFlag struct {
	*DatacenterFlag
	*SearchFlag

	register sync.Once
	name     string
	app      *object.VirtualApp
}

func (flag *VirtualAppFlag) Register(f *flag.FlagSet) {
	flag.SearchFlag = NewSearchFlag(SearchVirtualApps)

	flag.register.Do(func() {
		env := "GOVC_VAPP"
		value := os.Getenv(env)
		usage := fmt.Sprintf("Virtual App [%s]", env)
		f.StringVar(&flag.name, "vapp", value, usage)
	})
}

func (flag *VirtualAppFlag) Process() error { return nil }

func (flag *VirtualAppFlag) VirtualApp() (*object.VirtualApp, error) {
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

	flag.app, err = finder.VirtualApp(context.TODO(), flag.name)
	return flag.app, err
}
