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

package govmomi

import (
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type GuestProcessManager struct {
	types.ManagedObjectReference

	c *Client
}

func (m GuestProcessManager) Reference() types.ManagedObjectReference {
	return m.ManagedObjectReference
}

func (m GuestProcessManager) ListProcessesInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, pids []int64) ([]types.GuestProcessInfo, error) {
	req := types.ListProcessesInGuest{
		This: m.Reference(),
		Vm:   vm.Reference(),
		Auth: auth,
		Pids: pids,
	}

	res, err := methods.ListProcessesInGuest(m.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, err
}

func (m GuestProcessManager) ReadEnvironmentVariableInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, names []string) ([]string, error) {
	req := types.ReadEnvironmentVariableInGuest{
		This:  m.Reference(),
		Vm:    vm.Reference(),
		Auth:  auth,
		Names: names,
	}

	res, err := methods.ReadEnvironmentVariableInGuest(m.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, err
}

func (m GuestProcessManager) StartProgramInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, spec types.BaseGuestProgramSpec) (int64, error) {
	req := types.StartProgramInGuest{
		This: m.Reference(),
		Vm:   vm.Reference(),
		Auth: auth,
		Spec: spec,
	}

	res, err := methods.StartProgramInGuest(m.c, &req)
	if err != nil {
		return 0, err
	}

	return res.Returnval, err
}

func (m GuestProcessManager) TerminateProcessInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication, pid int64) error {
	req := types.TerminateProcessInGuest{
		This: m.Reference(),
		Vm:   vm.Reference(),
		Auth: auth,
		Pid:  pid,
	}

	_, err := methods.TerminateProcessInGuest(m.c, &req)
	return err
}
