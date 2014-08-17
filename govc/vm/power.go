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

package vm

import (
	"errors"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
)

type power struct {
	*flags.ClientFlag
	*flags.DatacenterFlag

	On  bool
	Off bool
}

func init() {
	cli.Register(&power{})
}

func (c *power) Register(f *flag.FlagSet) {
	f.BoolVar(&c.On, "on", false, "Power on")
	f.BoolVar(&c.Off, "off", false, "Power off")
}

func (c *power) Process() error {
	if !c.On && !c.Off || (c.On && c.Off) {
		return errors.New("specify -on OR -off")
	}
	return nil
}

func (c *power) PowerToString() string {
	switch {
	case c.On:
		return "on"
	case c.Off:
		return "off"
	default:
		panic("invalid state")
	}
}

func (c *power) Run(f *flag.FlagSet) error {
	if len(f.Args()) != 1 {
		return flag.ErrHelp
	}

	fmt.Printf("power %s VM %s\n", c.PowerToString(), f.Arg(0))

	return nil
}
