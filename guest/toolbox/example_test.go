/*
   Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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

package toolbox_test

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/guest/toolbox"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/test"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func ExampleClient_Run() {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		if !test.HasDocker() {
			fmt.Println("Linux")
			return nil
		}
		vm, err := find.NewFinder(c).VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			return err
		}
		err = simulator.RunContainer(ctx, c, vm, "nginx")
		if err != nil {
			return err
		}

		tools, err := toolbox.NewClient(ctx, c, vm, &types.NamePasswordAuthentication{
			Username: "user",
			Password: "pass",
		})

		cmd := &exec.Cmd{
			Path:   "uname",
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}

		return tools.Run(ctx, cmd)
	})
	// Output:
	// Linux
}
