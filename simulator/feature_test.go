// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator_test

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/test"
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
		if !test.HasDocker() {
			fmt.Println("0 diff")
			return
		}

		finder := find.NewFinder(c)
		pool, _ := finder.ResourcePool(ctx, "DC0_H0/Resources")
		dc, err := finder.Datacenter(ctx, "DC0")
		if err != nil {
			log.Fatal(err)
		}
		f, _ := dc.Folders(ctx)
		dir, err := os.MkdirTemp("", "example")
		if err != nil {
			log.Fatal(err)
		}
		os.Chmod(dir, 0755)
		fpath := filepath.Join(dir, "index.html")
		fcontent := "foo"
		os.WriteFile(fpath, []byte(fcontent), 0644)
		// just in case umask gets in the way
		os.Chmod(fpath, 0644)
		defer os.RemoveAll(dir)

		args := fmt.Sprintf("-v '%s:/usr/share/nginx/html:ro' nginx", dir)

		spec := types.VirtualMachineConfigSpec{
			Name: "nginx",
			Files: &types.VirtualMachineFileInfo{
				VmPathName: "[LocalDS_0] nginx",
			},
			ExtraConfig: []types.BaseOptionValue{
				&types.OptionValue{Key: "RUN.container", Value: args}, // run nginx
				&types.OptionValue{Key: "RUN.port.80", Value: "8888"}, // test port remap
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

		// Count the number of bytes in feature_test.go via nginx going direct to the container
		cmd := exec.Command("docker", "run", "--rm", "curlimages/curl", "curl", "-f", fmt.Sprintf("http://%s", ip))
		var buf bytes.Buffer
		cmd.Stdout = &buf
		err = cmd.Run()
		res := buf.String()

		if err != nil || strings.TrimSpace(res) != fcontent {
			log.Fatal(err, buf.String())
		}

		// Count the number of bytes in feature_test.go via nginx going via port remap on host
		cmd = exec.Command("curl", "-f", "http://localhost:8888")
		buf.Reset()
		cmd.Stdout = &buf
		err = cmd.Run()
		res = buf.String()
		if err != nil || strings.TrimSpace(res) != fcontent {
			log.Fatal(err, buf.String())
		}

		// PowerOff stops the container
		task, _ = vm.PowerOff(ctx)
		_ = task.Wait(ctx)
		// Destroy deletes the container
		task, _ = vm.Destroy(ctx)
		_ = task.Wait(ctx)

		fmt.Printf("%d diff", buf.Len()-len(fcontent))
	})
	// Output: 0 diff
}
