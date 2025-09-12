// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vms

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/vapi/cis/tasks"
)

func TestClusterCompliance(t *testing.T) {
	in := []byte(`{
    "parent": "",
    "cancelable": false,
    "end_time": "2025-09-04T23:01:58.388Z",
    "description": {
        "args": [],
        "default_message": "Task created by VMware vSphere Lifecycle Manager",
        "localized": "Task created by VMware vSphere Lifecycle Manager",
        "id": "com.vmware.vcIntegrity.lifecycle.Task.Description"
    },
    "target": {
        "id": "domain-c1008",
        "type": "ClusterComputeResource"
    },
    "result": {
        "host_solutions_status": {
            "compliances": {},
            "status": "COMPLIANT"
        },
        "cluster_solutions_status": {
            "compliances": {
                "solution2": {
                    "compliances": {
                        "5f1df158-d6e6-4c4a-98e2-f60adfe8c47f": {
                            "notifications": {
                                "info": [
                                    {
                                        "id": "com.vmware.vcIntegrity.lifecycle.lccm.nonCompliance.inHook",
                                        "time": "1970-01-01T00:00:00.000Z",
                                        "message": {
                                            "args": [],
                                            "default_message": "Apply operation is waiting for a hook to be processed.",
                                            "localized": "Apply operation is waiting for a hook to be processed.",
                                            "id": "com.vmware.vcIntegrity.lifecycle.lccm.nonCompliance.inHook"
                                        },
                                        "type": "INFO"
                                    }
                                ]
                            },
                            "deployment": {
                                "lifecycle_hook": {
                                    "lifecycle_state": "POST_PROVISIONING",
                                    "hook_activated": "2025-09-04T23:01:16.780Z",
                                    "configuration": {
                                        "timeout": 0
                                    },
                                    "vm": "vm-1049"
                                },
                                "solution_info": {
                                    "hook_configurations": {
                                        "POST_PROVISIONING": {
                                            "timeout": 0
                                        }
                                    },
                                    "vm_clone_config": "NO_CLONES",
                                    "vm_name_template": {
                                        "prefix": "templatename",
                                        "suffix": "COUNTER"
                                    },
                                    "ovf_descriptor_properties": {},
                                    "ovf_resource": {
                                        "authentication_scheme": "NONE",
                                        "ssl_certificate_validation": "DISABLED",
                                        "location_type": "REMOTE_FILE",
                                        "url": "https://lvn-dvm-10-161-123-247.dvm.lvn.broadcom.net:5480/wcpagent/photon-ova.ovf"
                                    },
                                    "display_name": "food",
                                    "vm_storage_profiles": [
                                        "dbcb3f01-0930-4899-aa76-bf4c05476c34"
                                    ],
                                    "vm_folder": "group-v1044",
                                    "cluster_solution_info": {
                                        "vm_count": 1,
                                        "vm_networks": {},
                                        "devices": {
                                            "_typeName": "VirtualMachineConfigSpec",
                                            "deviceChange": [
                                                {
                                                    "_typeName": "VirtualDeviceConfigSpec",
                                                    "fileOperation": "create",
                                                    "device": {
                                                        "backing": {
                                                            "port": {
                                                                "switchUuid": "50 2f f4 4d 27 b4 2a e2-76 b0 57 7e 4c 34 46 5a",
                                                                "_typeName": "DistributedVirtualSwitchPortConnection",
                                                                "portgroupKey": "dvportgroup-1020"
                                                            },
                                                            "_typeName": "VirtualEthernetCardDistributedVirtualPortBackingInfo"
                                                        },
                                                        "_typeName": "VirtualVmxnet3",
                                                        "addressType": "generated",
                                                        "unitNumber": 7,
                                                        "deviceInfo": {
                                                            "summary": "VM Network",
                                                            "_typeName": "Description",
                                                            "label": "Network 1"
                                                        },
                                                        "key": 0
                                                    },
                                                    "operation": "add"
                                                }
                                            ]
                                        },
                                        "alternative_vm_specs": [],
                                        "vm_placement_policies": [
                                            "VM_VM_ANTI_AFFINITY"
                                        ],
                                        "vm_datastores": [],
                                        "remediation_policy": "SEQUENTIAL"
                                    },
                                    "deployment_type": "CLUSTER_VM_SET",
                                    "display_version": "1",
                                    "vm_disk_type": "THICK",
                                    "vm_resource_pool": "resgroup-1043",
                                    "vm_resource_spec": {
                                        "ovf_deployment_option": "small"
                                    },
                                    "vm_storage_policy": "PROFILE"
                                },
                                "vm": "vm-1049",
                                "issues": [],
                                "status": "IN_LIFECYCLE_HOOK"
                            },
                            "status": "NON_COMPLIANT"
                        }
                    },
                    "status": "NON_COMPLIANT"
                }
            },
            "status": "NON_COMPLIANT"
        },
        "status": "NON_COMPLIANT"
    },
    "start_time": "2025-09-04T23:01:58.353Z",
    "last_update_time": "2025-09-04T23:01:58.391Z",
    "service": "com.vmware.esx.settings.clusters.vms.solutions",
    "progress": {
        "total": 100,
        "completed": 100,
        "message": {
            "args": [],
            "default_message": "Current progress for task created by VMware vSphere Lifecycle Manager",
            "localized": "Current progress for task created by VMware vSphere Lifecycle Manager",
            "id": "com.vmware.vcIntegrity.lifecycle.Task.Progress"
        }
    },
    "operation": "check_compliance$task",
    "user": "Administrator@VSPHERE.LOCAL",
    "notifications": {
        "warnings": [],
        "errors": [],
        "info": []
    },
    "status": "SUCCEEDED"
}`)

	var task tasks.Info

	require.NoError(t, json.Unmarshal(in, &task))

	var c ClusterCompliance
	require.NoError(t, json.Unmarshal(task.Result, &c))

	_, ok := c.ClusterSolutionsStatus.Compliances["solution2"]
	require.True(t, ok)

	require.Equal(t, *c.ClusterSolutionsStatus.Compliances["solution2"].Compliances["5f1df158-d6e6-4c4a-98e2-f60adfe8c47f"].Deployment.Vm, "vm-1049")
}
