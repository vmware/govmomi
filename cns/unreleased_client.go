// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cns

import (
	"context"

	"github.com/vmware/govmomi/cns/methods"
	cnstypes "github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/object"
)

// UpdateVolumeCrypto updates container volumes with given crypto specifications.
func (c *Client) UpdateVolumeCrypto(ctx context.Context, updateSpecList []cnstypes.CnsVolumeCryptoUpdateSpec) (*object.Task, error) {
	req := cnstypes.CnsUpdateVolumeCrypto{
		This:        CnsVolumeManagerInstance,
		UpdateSpecs: updateSpecList,
	}
	res, err := methods.CnsUpdateVolumeCrypto(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return object.NewTask(c.vim25Client, res.Returnval), nil
}
