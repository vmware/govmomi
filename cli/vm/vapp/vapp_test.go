// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vapp

import (
	"context"
	"flag"
	"net/url"
	"strings"
	"testing"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

const testVM = "DC0_H0_VM0"

// addDemoProperty simulates what deploying an OVF/OVA template does: define a
// vApp property on the VM. vcsim's stock VMs have none by default.
func addDemoProperty(ctx context.Context, c *vim25.Client, t *testing.T) {
	t.Helper()

	vm, err := find.NewFinder(c).VirtualMachine(ctx, testVM)
	if err != nil {
		t.Fatal(err)
	}

	spec := types.VirtualMachineConfigSpec{
		VAppConfig: &types.VmConfigSpec{
			Property: []types.VAppPropertySpec{
				{
					ArrayUpdateSpec: types.ArrayUpdateSpec{Operation: types.ArrayUpdateOperationAdd},
					Info: &types.VAppPropertyInfo{
						Key:   1,
						Id:    "guestinfo.hostname",
						Type:  "string",
						Label: "Hostname",
						Value: "initial-host",
					},
				},
			},
		},
	}

	task, err := vm.Reconfigure(ctx, spec)
	if err != nil {
		t.Fatal(err)
	}
	if err := task.Wait(ctx); err != nil {
		t.Fatal(err)
	}
}

func govcURL(c *vim25.Client) string {
	u := *c.URL()
	u.User = url.UserPassword("user", "pass")
	return u.String()
}

func runInfo(t *testing.T, ctx context.Context) string {
	t.Helper()

	cmd := &info{}
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	cmd.Register(ctx, fs)
	if err := fs.Parse([]string{"-vm", testVM}); err != nil {
		t.Fatal(err)
	}
	if err := cmd.Process(ctx); err != nil {
		t.Fatal(err)
	}

	var buf strings.Builder
	cmd.Out = &buf

	if err := cmd.Run(ctx, fs); err != nil {
		t.Fatal(err)
	}

	return buf.String()
}

func runSet(ctx context.Context, property string) error {
	cmd := &set{}
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	cmd.Register(ctx, fs)
	if err := fs.Parse([]string{"-vm", testVM, "-property", property}); err != nil {
		return err
	}
	if err := cmd.Process(ctx); err != nil {
		return err
	}

	return cmd.Run(ctx, fs)
}

func TestInfoEmpty(t *testing.T) {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		t.Setenv("GOVC_URL", govcURL(c))
		t.Setenv("GOVC_INSECURE", "true")

		out := runInfo(t, ctx)
		if strings.Contains(out, "guestinfo") {
			t.Fatalf("expected no vApp properties, got: %s", out)
		}

		return nil
	})
}

func TestSetUnknownProperty(t *testing.T) {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		t.Setenv("GOVC_URL", govcURL(c))
		t.Setenv("GOVC_INSECURE", "true")

		err := runSet(ctx, "does.not.exist=1")
		if err == nil {
			t.Fatal("expected error for unknown vApp property, got nil")
		}
		if !strings.Contains(err.Error(), "does.not.exist") {
			t.Fatalf("expected error to mention property id, got: %v", err)
		}

		return nil
	})
}

func TestSetExistingProperty(t *testing.T) {
	simulator.Run(func(ctx context.Context, c *vim25.Client) error {
		t.Setenv("GOVC_URL", govcURL(c))
		t.Setenv("GOVC_INSECURE", "true")

		addDemoProperty(ctx, c, t)

		before := runInfo(t, ctx)
		if !strings.Contains(before, "guestinfo.hostname") || !strings.Contains(before, "initial-host") {
			t.Fatalf("expected initial property in output, got: %s", before)
		}

		if err := runSet(ctx, "guestinfo.hostname=updated-host"); err != nil {
			t.Fatal(err)
		}

		after := runInfo(t, ctx)
		if !strings.Contains(after, "updated-host") {
			t.Fatalf("expected updated value in output, got: %s", after)
		}
		if strings.Contains(after, "initial-host") {
			t.Fatalf("expected old value to be gone, got: %s", after)
		}
		// label/type metadata must be preserved across the edit, not wiped.
		if !strings.Contains(after, "Hostname") || !strings.Contains(after, "string") {
			t.Fatalf("expected label/type metadata to be preserved, got: %s", after)
		}

		return nil
	})
}
