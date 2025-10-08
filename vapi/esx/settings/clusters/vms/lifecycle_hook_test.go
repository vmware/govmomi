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

func TestDynamicUpdateSpec(t *testing.T) {
	// Test with AlternativeVmSpec
	in := []byte(`{
        "vm": "vm-123",
        "solution": "solution-456",
        "lifecycle_state": "POST_PROVISIONING",
        "alternative_vm_spec": {
            "selection_criteria": {
                "selection_type": "VM_EXTRA_CONFIG",
                "extra_config_value": "test-uuid"
            },
            "devices": "{\"deviceChange\":[]}"
        }
    }`)

	var out DynamicUpdateSpec
	require.NoError(t, json.Unmarshal(in, &out))
	require.Equal(t, out.Vm, "vm-123")
	require.Equal(t, out.Solution, "solution-456")
	require.Equal(t, out.LifecycleState, PostProvisioning)
	require.NotNil(t, out.AlternativeVmSpec)
	require.Equal(t, out.AlternativeVmSpec.SelectionCriteria.SelectionType, VmSelectionTypeVmExtraConfig)
	require.Equal(t, out.AlternativeVmSpec.SelectionCriteria.ExtraConfigValue, "test-uuid")
}

func TestDynamicUpdateSpecWithoutAlternativeVmSpec(t *testing.T) {
	// Test without AlternativeVmSpec
	in := []byte(`{
        "vm": "vm-789",
        "solution": "solution-101",
        "lifecycle_state": "POST_POWER_ON"
    }`)

	var out DynamicUpdateSpec
	require.NoError(t, json.Unmarshal(in, &out))
	require.Equal(t, out.Vm, "vm-789")
	require.Equal(t, out.Solution, "solution-101")
	require.Equal(t, out.LifecycleState, PostPowerOn)
	require.Nil(t, out.AlternativeVmSpec)
}
