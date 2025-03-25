// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type AuthManager struct {
	types.ManagedObjectReference

	vm types.ManagedObjectReference

	c *vim25.Client
}

func (m AuthManager) Reference() types.ManagedObjectReference {
	return m.ManagedObjectReference
}

func (m AuthManager) AcquireCredentials(ctx context.Context, requestedAuth types.BaseGuestAuthentication, sessionID int64) (types.BaseGuestAuthentication, error) {
	req := types.AcquireCredentialsInGuest{
		This:          m.Reference(),
		Vm:            m.vm,
		RequestedAuth: requestedAuth,
		SessionID:     sessionID,
	}

	res, err := methods.AcquireCredentialsInGuest(ctx, m.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (m AuthManager) ReleaseCredentials(ctx context.Context, auth types.BaseGuestAuthentication) error {
	req := types.ReleaseCredentialsInGuest{
		This: m.Reference(),
		Vm:   m.vm,
		Auth: auth,
	}

	_, err := methods.ReleaseCredentialsInGuest(ctx, m.c, &req)

	return err
}

func (m AuthManager) ValidateCredentials(ctx context.Context, auth types.BaseGuestAuthentication) error {
	req := types.ValidateCredentialsInGuest{
		This: m.Reference(),
		Vm:   m.vm,
		Auth: auth,
	}

	_, err := methods.ValidateCredentialsInGuest(ctx, m.c, &req)
	if err != nil {
		return err
	}

	return nil
}
