/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package simulator_test

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

// Custom username + password authentication
func Example_usernamePasswordLogin() {
	model := simulator.VPX()

	defer model.Remove()
	err := model.Create()
	if err != nil {
		log.Fatal(err)
	}

	model.Service.Listen = &url.URL{
		User: url.UserPassword("my-username", "my-password"),
	}

	s := model.Service.NewServer()
	defer s.Close()

	c, err := govmomi.NewClient(context.Background(), s.URL, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("login to %s as %s", c.Client.ServiceContent.About.ApiType, s.URL.User)
	// Output: login to VirtualCenter as my-username:my-password
}

// Set VM properties that the API cannot change in a real vCenter.
func Example_setVirtualMachineProperties() {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		vm, err := find.NewFinder(c).VirtualMachine(ctx, "DC0_H0_VM0")
		if err != nil {
			log.Fatal(err)
		}

		spec := types.VirtualMachineConfigSpec{
			ExtraConfig: []types.BaseOptionValue{
				&types.OptionValue{Key: "SET.guest.ipAddress", Value: "10.0.0.42"},
			},
		}

		task, _ := vm.Reconfigure(ctx, spec)

		_ = task.Wait(ctx)

		ip, _ := vm.WaitForIP(ctx)
		fmt.Printf("ip is %s", ip)
	})
	// Output: ip is 10.0.0.42
}

// Tie a docker container to the lifecycle of a vcsim VM
func Example_runContainer() {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		if _, err := exec.LookPath("docker"); err != nil {
			fmt.Println("0 diff")
			return // docker is required
		}

		finder := find.NewFinder(c)
		pool, _ := finder.ResourcePool(ctx, "DC0_H0/Resources")
		dc, err := finder.Datacenter(ctx, "DC0")
		if err != nil {
			log.Fatal(err)
		}
		f, _ := dc.Folders(ctx)
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		args := fmt.Sprintf("-v '%s:/usr/share/nginx/html:ro' nginx", dir)

		spec := types.VirtualMachineConfigSpec{
			Name: "nginx",
			Files: &types.VirtualMachineFileInfo{
				VmPathName: "[LocalDS_0] nginx",
			},
			ExtraConfig: []types.BaseOptionValue{
				&types.OptionValue{Key: "RUN.container", Value: args}, // run nginx
			},
		}

		// Create a new VM
		task, err := f.VmFolder.CreateVM(ctx, spec, pool, nil)
		if err != nil {
			log.Fatal(err)
		}
		info, err := task.WaitForResult(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}
		vm := object.NewVirtualMachine(c, info.Result.(types.ManagedObjectReference))

		// PowerOn VM starts the nginx container
		task, _ = vm.PowerOn(ctx)
		err = task.Wait(ctx)
		if err != nil {
			log.Fatal(err)
		}

		ip, _ := vm.WaitForIP(ctx, true) // Returns the docker container's IP

		// Count the number of bytes in feature_test.go via nginx
		res, err := http.Get(fmt.Sprintf("http://%s/feature_test.go", ip))
		if err != nil {
			log.Fatal(err)
		}
		n, err := io.Copy(ioutil.Discard, res.Body)
		if err != nil {
			log.Print(err)
		}

		// PowerOff stops the container
		task, _ = vm.PowerOff(ctx)
		_ = task.Wait(ctx)
		// Destroy deletes the container
		task, _ = vm.Destroy(ctx)
		_ = task.Wait(ctx)

		st, _ := os.Stat("feature_test.go")
		fmt.Printf("%d diff", n-st.Size())
	})
	// Output: 0 diff
}
