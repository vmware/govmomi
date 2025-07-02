// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/test"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

// takes a content string to serve from the container and returns ExtraConfig options
// to construct container
// content - the contents of index.html
// port - the port to forward to the container port 80
func constructNginxBacking(t *testing.T, content string, port int) []types.BaseOptionValue {
	dir := t.TempDir()
	// experience shows that a parent directory created as part of the TempDir call may not have
	// o+rx, preventing use within a container that doesn't have the same uid
	for dirpart := dir; dirpart != "/"; dirpart = filepath.Dir(dirpart) {
		os.Chmod(dirpart, 0755)
		stat, err := os.Stat(dirpart)
		require.Nil(t, err, "must be able to check file and directory permissions")
		require.NotZero(t, stat.Mode()&0005, "does not have o+rx permissions", dirpart)
	}

	fpath := filepath.Join(dir, "index.html")
	err := os.WriteFile(fpath, []byte(content), 0644)
	require.Nil(t, err, "Expected to cleanly write content to file: %s", err)

	// just in case umask gets in the way
	err = os.Chmod(fpath, 0644)
	require.Nil(t, err, "Expected to cleanly set file permissions on content: %s", err)

	args := fmt.Sprintf("-v '%s:/usr/share/nginx/html:ro' nginx", dir)

	return []types.BaseOptionValue{
		&types.OptionValue{Key: ContainerBackingOptionKey, Value: args}, // run nginx
		&types.OptionValue{Key: "RUN.port.80", Value: "8888"},           // test port remap
	}
}

// validates the VM is serving the expected content on the expected ports
// pairs with constructNginxBacking
func validateNginxContainer(t *testing.T, vm *object.VirtualMachine, expected string, port int) error {
	ip, _ := vm.WaitForIP(context.Background(), true) // Returns the docker container's IP

	// Count the number of bytes in feature_test.go via nginx going direct to the container
	cmd := exec.Command("docker", "run", "--rm", "curlimages/curl", "curl", "-f", fmt.Sprintf("http://%s:80", ip))
	var buf bytes.Buffer
	cmd.Stdout = &buf
	err := cmd.Run()
	res := buf.String()

	if err != nil || strings.TrimSpace(res) != expected {
		// we use Fail not Fatal because we want to clean up
		t.Fail()
		t.Log(err, buf.String())
		fmt.Printf("%d diff", buf.Len()-len(expected))
	}

	// Count the number of bytes in feature_test.go via nginx going via port remap on host
	cmd = exec.Command("curl", "-f", fmt.Sprintf("http://localhost:%d", port))
	buf.Reset()
	cmd.Stdout = &buf
	err = cmd.Run()
	res = buf.String()
	if err != nil || strings.TrimSpace(res) != expected {
		t.Fail()
		t.Log(err, buf.String())
		fmt.Printf("%d diff", buf.Len()-len(expected))
	}

	return nil
}

// 1. Construct ExtraConfig args for container backing
// 2. Create VM using that ExtraConfig
// 3. Confirm docker container present that matches expectations
func TestCreateVMWithContainerBacking(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		if !test.HasDocker() {
			t.Skip("requires docker on linux")
			return
		}

		finder := find.NewFinder(c)
		pool, _ := finder.ResourcePool(ctx, "DC0_H0/Resources")
		dc, err := finder.Datacenter(ctx, "DC0")
		if err != nil {
			log.Fatal(err)
		}

		content := "foo"
		port := 8888

		spec := types.VirtualMachineConfigSpec{
			Name: "nginx-container-backed-from-creation",
			Files: &types.VirtualMachineFileInfo{
				VmPathName: "[LocalDS_0] nginx",
			},
			ExtraConfig: constructNginxBacking(t, content, port),
		}

		f, _ := dc.Folders(ctx)
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

		err = validateNginxContainer(t, vm, content, port)
		if err != nil {
			log.Fatal(err)
		}

		spec2 := types.VirtualMachineConfigSpec{
			ExtraConfig: []types.BaseOptionValue{
				&types.OptionValue{Key: ContainerBackingOptionKey, Value: ""},
			},
		}

		task, err = vm.Reconfigure(ctx, spec2)
		if err != nil {
			log.Fatal(err)
		}

		info, err = task.WaitForResult(ctx, nil)
		if err != nil {
			log.Fatal(info, err)
		}

		// PowerOff stops the container
		task, _ = vm.PowerOff(ctx)
		_ = task.Wait(ctx)
		// Destroy deletes the container
		task, _ = vm.Destroy(ctx)
		_ = task.Wait(ctx)
	})
}

// 1. Create VM without ExtraConfig args for container backing
// 2. Construct ExtraConfig args for container backing
// 3. Update VM with ExtraConfig
// 4. Confirm docker container present that matches expectations
func TestUpdateVMAddContainerBacking(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		if !test.HasDocker() {
			t.Skip("requires docker on linux")
			return
		}

		finder := find.NewFinder(c)
		pool, _ := finder.ResourcePool(ctx, "DC0_H0/Resources")
		dc, err := finder.Datacenter(ctx, "DC0")
		if err != nil {
			log.Fatal(err)
		}

		content := "foo"
		port := 8888

		spec := types.VirtualMachineConfigSpec{
			Name: "nginx-container-after-reconfig",
			Files: &types.VirtualMachineFileInfo{
				VmPathName: "[LocalDS_0] nginx",
			},
		}

		f, _ := dc.Folders(ctx)
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

		spec2 := types.VirtualMachineConfigSpec{
			ExtraConfig: constructNginxBacking(t, content, port),
		}

		task, err = vm.Reconfigure(ctx, spec2)
		if err != nil {
			log.Fatal(err)
		}

		info, err = task.WaitForResult(ctx, nil)
		if err != nil {
			log.Fatal(info, err)
		}

		err = validateNginxContainer(t, vm, content, port)
		if err != nil {
			log.Fatal(err)
		}

		// PowerOff stops the container
		task, _ = vm.PowerOff(ctx)
		_ = task.Wait(ctx)
		// Destroy deletes the container
		task, _ = vm.Destroy(ctx)
		_ = task.Wait(ctx)
	})
}
