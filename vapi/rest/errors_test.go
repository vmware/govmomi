// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package rest

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrors(t *testing.T) {
	in := []byte(`{
    "error_type": "NOT_FOUND",
    "messages": [
        {
            "args": [
                "{cluster}",
                "cluster"
            ],
            "default_message": "Entity '{cluster}' does not exist or is not of type 'cluster'.",
            "localized": "Entity '{cluster}' does not exist or is not of type 'cluster'.",
            "id": "com.vmware.vcIntegrity.lifecycle.RequireAdminUserAuthz.EntityNotFound"
        }
    ]
}`)

	var out Error
	require.NoError(t, json.Unmarshal(in, &out))
	require.Equal(t, out.ErrorType, ErrNotFound)
	require.Len(t, out.Messages, 1)
	require.Equal(t, out.Messages[0].DefaultMessage, "Entity '{cluster}' does not exist or is not of type 'cluster'.")
	require.Equal(t, out.Messages[0].ID, "com.vmware.vcIntegrity.lifecycle.RequireAdminUserAuthz.EntityNotFound")

}
