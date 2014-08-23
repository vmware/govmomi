/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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
	"os"
	"path"
	"time"

	"github.com/vmware/govmomi/vim25/debug"
)

type DebugFlag struct {
	enable bool
}

func (d *DebugFlag) Register(f *flag.FlagSet) {
	f.BoolVar(&d.enable, "debug", false, "Store debug logs")
}

func (d *DebugFlag) Process() error {
	if d.enable {
		// Base path for storing debug logs.
		r := os.Getenv("GOVC_DEBUG_PATH")
		if r == "" {
			r = path.Join(os.Getenv("HOME"), ".govmomi")
		}
		r = path.Join(r, "debug")

		// Path for this particular run.
		now := time.Now().Format("2006-01-02T15-04-05.999999999")
		r = path.Join(r, now)

		err := os.MkdirAll(r, 0700)
		if err != nil {
			return err
		}

		p := debug.FileProvider{
			Path: r,
		}

		debug.SetProvider(&p)
	}

	return nil
}
