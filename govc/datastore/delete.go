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

package datastore

import (
	"errors"
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
)

type delete struct {
	*flags.DatastoreFlag
}

func init() {
	cli.Register(&delete{})
}

func (cmd *delete) Register(f *flag.FlagSet) {}

func (cmd *delete) Process() error { return nil }

func (cmd *delete) Run(f *flag.FlagSet) error {
	args := f.Args()
	if len(args) == 0 {
		return errors.New("missing operand")
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	dc, err := cmd.Datacenter()
	if err != nil {
		return err
	}

	// TODO(PN): Accept multiple args
	path, err := cmd.DatastorePath(args[0])
	if err != nil {
		return err
	}

	task, err := c.FileManager().DeleteDatastoreFile(path, dc)
	if err != nil {
		return err
	}

	return task.Wait()
}
