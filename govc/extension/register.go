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

package extension

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"time"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type register struct {
	*flags.ClientFlag

	update bool
}

func init() {
	cli.Register("extension.register", &register{})
}

func (cmd *register) Register(f *flag.FlagSet) {
	f.BoolVar(&cmd.update, "update", false, "Update extension")
}

func (cmd *register) Process() error { return nil }

func (cmd *register) Run(f *flag.FlagSet) error {
	ctx := context.TODO()

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := object.GetExtensionManager(c)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	var e types.Extension
	e.Description = new(types.Description)

	if err = json.Unmarshal(b, &e); err != nil {
		return err
	}

	e.LastHeartbeatTime = time.Now().UTC()

	if cmd.update {
		return m.Update(ctx, e)
	}

	return m.Register(ctx, e)
}
