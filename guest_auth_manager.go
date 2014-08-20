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

type GuestAuthManager struct {
	types.ManagedObjectReference

	c *Client
}

func (m GuestAuthManager) Reference() types.ManagedObjectReference {
	return m.ManagedObjectReference
}

func (m GuestAuthManager) AcquireCredentialsInGuest(vm *VirtualMachine, requestedAuth types.BaseGuestAuthentication, sessionID int64) (types.BaseGuestAuthentication, error) {
	req := types.AcquireCredentialsInGuest{
		This:          m.Reference(),
		Vm:            vm.Reference(),
		RequestedAuth: requestedAuth,
		SessionID:     sessionID,
	}

	res, err := methods.AcquireCredentialsInGuest(m.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (m GuestAuthManager) ReleaseCredentialsInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication) error {
	req := types.ReleaseCredentialsInGuest{
		This: m.Reference(),
		Vm:   vm.Reference(),
		Auth: auth,
	}

	_, err := methods.ReleaseCredentialsInGuest(m.c, &req)

	return err
}

func (m GuestAuthManager) ValidateCredentialsInGuest(vm *VirtualMachine, auth types.BaseGuestAuthentication) error {
	req := types.ValidateCredentialsInGuest{
		This: m.Reference(),
		Vm:   vm.Reference(),
		Auth: auth,
	}

	_, err := methods.ValidateCredentialsInGuest(m.c, &req)
	if err != nil {
		return err
	}

	return nil
}
