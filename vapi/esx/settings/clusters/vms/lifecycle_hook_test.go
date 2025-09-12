// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vms

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHookList(t *testing.T) {
	in := []byte(`{
    "hooks": [
        {
            "lifecycle_state": "POST_POWER_ON",
            "hook_activated": "2025-09-04T04:44:08.425Z",
            "configuration": {
                "timeout": 0
            },
            "vm": "vm-1048",
            "dynamic_update_processed": false
        }
    ]
}`)

	var out HookListResult
	require.NoError(t, json.Unmarshal(in, &out))
	require.Len(t, out.Hooks, 1)
	require.Equal(t, out.Hooks[0].LifecycleState, PostPowerOn)
	require.Equal(t, out.Hooks[0].Vm, "vm-1048")
	require.False(t, out.Hooks[0].DynamicUpdateProcessed)
}
