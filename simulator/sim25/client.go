// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package sim25

import (
	"context"
	"time"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

// SetSessionTimeout changes the default session timeout.
func SetSessionTimeout(ctx context.Context, c *vim25.Client, timeout time.Duration) error {
	// Note that real vCenter supports the same OptionValue Key, but only with a Value of whole seconds.
	req := types.UpdateOptions{
		This: *c.ServiceContent.Setting,
		ChangedValue: []types.BaseOptionValue{
			&types.OptionValue{
				Key:   "config.vmacore.soap.sessionTimeout",
				Value: timeout.String(),
			},
		},
	}

	_, err := methods.UpdateOptions(ctx, c, &req)
	return err
}
