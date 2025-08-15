// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
	QuerySelectionNameTypeDataStoreUrl           = QuerySelectionNameType("DATASTORE_URL")
	QuerySelectionNameTypePolicyId               = QuerySelectionNameType("POLICY_ID")
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
	CnsQuerySelectionName_HEALTH_STATUS                  = CnsQuerySelectionNameType("HEALTH_STATUS")
	CnsQuerySelectionName_DATASTORE_URL                  = CnsQuerySelectionNameType("DATASTORE_URL")
	CnsQuerySelectionName_POLICY_ID                      = CnsQuerySelectionNameType("POLICY_ID")
)

func init() {
	types.Add("CnsKubernetesEntityType", reflect.TypeOf((*CnsKubernetesEntityType)(nil)).Elem())
}

type CnsSyncVolumeMode string

const (
	CnsSyncVolumeModeSPACE_USAGE = CnsSyncVolumeMode("SPACE_USAGE")
)

func init() {
	types.Add("vsan:CnsSyncVolumeMode", reflect.TypeOf((*CnsSyncVolumeMode)(nil)).Elem())
}

type CnsUnregisterTargetVolumeType string

const (
	CnsUnregisterTargetVolumeTypeFCD         = CnsUnregisterTargetVolumeType("FCD")
	CnsUnregisterTargetVolumeTypeLEGACY_DISK = CnsUnregisterTargetVolumeType("LEGACY_DISK")
)

func init() {
	types.Add("vsan:CnsUnregisterTargetVolumeType", reflect.TypeOf((*CnsUnregisterTargetVolumeType)(nil)).Elem())
}
