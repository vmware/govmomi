// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package license

import (
	"context"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type AssignmentManager struct {
	object.Common
}

func (m AssignmentManager) QueryAssigned(ctx context.Context, id string) ([]types.LicenseAssignmentManagerLicenseAssignment, error) {
	req := types.QueryAssignedLicenses{
		This:     m.Reference(),
		EntityId: id,
	}

	res, err := methods.QueryAssignedLicenses(ctx, m.Client(), &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (m AssignmentManager) Remove(ctx context.Context, id string) error {
	req := types.RemoveAssignedLicense{
		This:     m.Reference(),
		EntityId: id,
	}

	_, err := methods.RemoveAssignedLicense(ctx, m.Client(), &req)

	return err
}

func (m AssignmentManager) Update(ctx context.Context, id string, key string, name string) (*types.LicenseManagerLicenseInfo, error) {
	req := types.UpdateAssignedLicense{
		This:              m.Reference(),
		Entity:            id,
		LicenseKey:        key,
		EntityDisplayName: name,
	}

	res, err := methods.UpdateAssignedLicense(ctx, m.Client(), &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}
