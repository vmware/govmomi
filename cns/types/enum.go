/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

type CnsVolumeType string

const (
	CnsVolumeTypeBlock = CnsVolumeType("BLOCK")
	CnsVolumeTypeFile  = CnsVolumeType("FILE")
)

func init() {
	types.Add("CnsVolumeType", reflect.TypeOf((*CnsVolumeType)(nil)).Elem())
}

type CnsClusterFlavor string

const (
	CnsClusterFlavorVanilla  = CnsClusterFlavor("VANILLA")
	CnsClusterFlavorWorkload = CnsClusterFlavor("WORKLOAD")
	CnsClusterFlavorGuest    = CnsClusterFlavor("GUEST_CLUSTER")
	CnsClusterFlavorUnknown  = CnsClusterFlavor("ClusterFlavor_Unknown")
)

func init() {
	types.Add("CnsClusterFlavor", reflect.TypeOf((*CnsClusterFlavor)(nil)).Elem())
}

type QuerySelectionNameType string

const (
	QuerySelectionNameTypeVolumeType             = QuerySelectionNameType("VOLUME_TYPE")
	QuerySelectionNameTypeVolumeName             = QuerySelectionNameType("VOLUME_NAME")
	QuerySelectionNameTypeBackingObjectDetails   = QuerySelectionNameType("BACKING_OBJECT_DETAILS")
	QuerySelectionNameTypeComplianceStatus       = QuerySelectionNameType("COMPLIANCE_STATUS")
	QuerySelectionNameTypeDataStoreAccessibility = QuerySelectionNameType("DATASTORE_ACCESSIBILITY_STATUS")
	QuerySelectionNameTypeHealthStatus           = QuerySelectionNameType("HEALTH_STATUS")
)

func init() {
	types.Add("QuerySelectionNameType", reflect.TypeOf((*QuerySelectionNameType)(nil)).Elem())
}

type CnsClusterType string

const (
	CnsClusterTypeKubernetes = CnsClusterType("KUBERNETES")
)

func init() {
	types.Add("CnsClusterType", reflect.TypeOf((*CnsClusterType)(nil)).Elem())
}

type CnsKubernetesEntityType string

const (
	CnsKubernetesEntityTypePVC = CnsKubernetesEntityType("PERSISTENT_VOLUME_CLAIM")
	CnsKubernetesEntityTypePV  = CnsKubernetesEntityType("PERSISTENT_VOLUME")
	CnsKubernetesEntityTypePOD = CnsKubernetesEntityType("POD")
)

type CnsQuerySelectionNameType string

const (
	CnsQuerySelectionName_VOLUME_NAME                    = CnsQuerySelectionNameType("VOLUME_NAME")
	CnsQuerySelectionName_VOLUME_TYPE                    = CnsQuerySelectionNameType("VOLUME_TYPE")
	CnsQuerySelectionName_BACKING_OBJECT_DETAILS         = CnsQuerySelectionNameType("BACKING_OBJECT_DETAILS")
	CnsQuerySelectionName_COMPLIANCE_STATUS              = CnsQuerySelectionNameType("COMPLIANCE_STATUS")
	CnsQuerySelectionName_DATASTORE_ACCESSIBILITY_STATUS = CnsQuerySelectionNameType("DATASTORE_ACCESSIBILITY_STATUS")
)

func init() {
	types.Add("CnsKubernetesEntityType", reflect.TypeOf((*CnsKubernetesEntityType)(nil)).Elem())
}
