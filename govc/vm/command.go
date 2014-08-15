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
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
)

type vm struct {
	flag.FlagSet
}

type create struct {
	vm

	pool      string
	host      string
	datastore string
	memory    int
	cpus      int
	guestID   string
}

type power struct {
	vm
}

func init() {
	cli.Register(&create{})
	cli.Register(&power{})
}

func (c *vm) Name() string {
	return c.Args()[0]
}

func (c *create) Parse(args []string) error {
	c.StringVar(&c.pool, "p", "", "Resource pool")
	c.StringVar(&c.host, "o", "", "Host")
	c.StringVar(&c.datastore, "d", "", "Datastore")
	c.IntVar(&c.memory, "m", 128, "Size in MB of memory")
	c.IntVar(&c.cpus, "c", 1, "Number of CPUs")
	c.StringVar(&c.guestID, "g", "otherGuest", "Guest OS")

	return c.FlagSet.Parse(args)
}

func (c *create) Run() error {
	if len(c.Args()) != 1 {
		return flag.ErrHelp
	}

	fmt.Printf("create %s VM '%s' with %d MB memory\n", c.guestID, c.Name(), c.memory)

	return nil
}

func (c *power) Run() error {
	if len(c.Args()) != 1 {
		return flag.ErrHelp
	}

	fmt.Printf("power %s VM %s\n", c.Name(), c.Arg(1))

	return nil
}
