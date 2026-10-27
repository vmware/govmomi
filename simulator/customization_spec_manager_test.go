// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

func TestDeleteCustomizationSpec(t *testing.T) {
	Test(func(ctx context.Context, c *vim25.Client) {
		m := object.NewCustomizationSpecManager(c)
		name := DefaultCustomizationSpec[0].Info.Name

		err := m.DeleteCustomizationSpec(ctx, name)
		require.NoError(t, err)

		exists, err := m.DoesCustomizationSpecExist(ctx, name)
		require.NoError(t, err)
		require.False(t, exists)

		err = m.DeleteCustomizationSpec(ctx, name)
		require.True(t, fault.Is(err, &types.NotFound{}))

		// the package level defaults must not be modified
		require.Equal(t, name, DefaultCustomizationSpec[0].Info.Name)
	})
}
