// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package task_test

import (
	"context"
	"errors"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/task"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

func TestFaultString(t *testing.T) {
	msgs := func(m ...string) types.VimFault {
		var v types.VimFault
		for _, s := range m {
			v.FaultMessage = append(v.FaultMessage, types.LocalizableMessage{Key: "k", Message: s})
		}
		return v
	}

	tests := []struct {
		name string
		in   *types.LocalizedMethodFault
		want string
	}{
		{"nil", nil, ""},
		{"nil fault", &types.LocalizedMethodFault{LocalizedMessage: "boom"}, "boom"},
		{
			"no details",
			&types.LocalizedMethodFault{Fault: &types.InvalidArgument{}, LocalizedMessage: "boom"},
			"boom",
		},
		{
			"fault messages",
			&types.LocalizedMethodFault{
				Fault:            &types.GenericVmConfigFault{VmConfigFault: types.VmConfigFault{VimFault: msgs("a", "", "[context]x", "b", "boom")}},
				LocalizedMessage: "boom",
			},
			"boom: a; b",
		},
		{
			"empty localized message",
			&types.LocalizedMethodFault{Fault: &types.InvalidArgument{RuntimeFault: types.RuntimeFault{MethodFault: types.MethodFault{
				FaultMessage: []types.LocalizableMessage{{Message: "a"}},
			}}}},
			"a",
		},
		{
			"nested NoCompatibleHost",
			&types.LocalizedMethodFault{
				LocalizedMessage: "Operation failed due to: ",
				Fault: &types.NoCompatibleHost{
					VimFault: msgs("top"),
					Error: []types.LocalizedMethodFault{
						{
							LocalizedMessage: "host1 bad",
							Fault: &types.NoCompatibleHost{
								VimFault: msgs("deep"),
								Error:    []types.LocalizedMethodFault{{LocalizedMessage: "inner", Fault: &types.InvalidArgument{}}},
							},
						},
						{LocalizedMessage: "host2 bad", Fault: &types.InvalidArgument{}},
					},
				},
			},
			"Operation failed due to: top; host1 bad; deep; inner; host2 bad",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := task.FaultString(test.in); got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}

// failingVM injects a task failure with a generic message and specific details.
type failingVM struct {
	*simulator.VirtualMachine
}

func (vm *failingVM) PowerOffVMTask(ctx *simulator.Context, req *types.PowerOffVM_Task) soap.HasFault {
	t := simulator.CreateTask(req.This, "powerOff", func(*simulator.Task) (types.AnyType, types.BaseMethodFault) {
		return nil, &types.NoCompatibleHost{
			VimFault: types.VimFault{MethodFault: types.MethodFault{
				FaultMessage: []types.LocalizableMessage{{Key: "vim.fault.Detail", Message: "disk is locked"}},
			}},
			Error: []types.LocalizedMethodFault{{LocalizedMessage: "host-1 is in maintenance", Fault: &types.InvalidArgument{}}},
		}
	})

	return &methods.PowerOffVM_TaskBody{
		Res: &types.PowerOffVM_TaskResponse{Returnval: t.Run(ctx)},
	}
}

func TestTaskErrorVcsim(t *testing.T) {
	model := simulator.VPX()
	defer model.Remove()
	if err := model.Create(); err != nil {
		t.Fatal(err)
	}

	s := model.Service.NewServer()
	defer s.Close()

	ctx := context.Background()
	c, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	obj := model.Map().Any("VirtualMachine").(*simulator.VirtualMachine)
	model.Map().Put(&failingVM{obj})
	vm := object.NewVirtualMachine(c.Client, obj.Reference())

	check := func(err error) {
		t.Helper()
		if err == nil {
			t.Fatal("expected error")
		}
		var terr task.Error
		if !errors.As(err, &terr) {
			t.Fatalf("expected task.Error, got %T", err)
		}
		const want = "*types.NoCompatibleHost: disk is locked; host-1 is in maintenance"
		if err.Error() != want {
			t.Errorf("got %q, want %q", err.Error(), want)
		}
		if _, ok := terr.Fault().(*types.NoCompatibleHost); !ok {
			t.Errorf("unexpected fault %T", terr.Fault())
		}
		if terr.LocalizedMessage != "*types.NoCompatibleHost" {
			t.Errorf("LocalizedMessage changed: %q", terr.LocalizedMessage)
		}
	}

	tk, err := vm.PowerOff(ctx)
	if err != nil {
		t.Fatal(err)
	}
	check(tk.Wait(ctx))

	tk, _ = vm.PowerOff(ctx)
	_, err = tk.WaitForResult(ctx)
	check(err)
}
