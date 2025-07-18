// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vpx

import (
	"github.com/google/uuid"

	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/vim25/types"
)

// ServiceContent is the default template for the ServiceInstance content property.
// Capture method:
// govc object.collect -s -dump - content
var ServiceContent = types.ServiceContent{
	RootFolder:        types.ManagedObjectReference{Type: "Folder", Value: "group-d1"},
	PropertyCollector: types.ManagedObjectReference{Type: "PropertyCollector", Value: "propertyCollector"},
	ViewManager:       &types.ManagedObjectReference{Type: "ViewManager", Value: "ViewManager"},
	About: types.AboutInfo{
		Name:                  "VMware vCenter Server",
		FullName:              "VMware vCenter Server 6.5.0 build-5973321",
		Vendor:                "VMware, Inc.",
		Version:               "6.5.0",
		Build:                 "5973321",
		LocaleVersion:         "INTL",
		LocaleBuild:           "000",
		OsType:                "linux-x64",
		ProductLineId:         "vpx",
		ApiType:               "VirtualCenter",
		ApiVersion:            "6.5",
		InstanceUuid:          internal.OID(string(uuid.NodeID())).String(),
		LicenseProductName:    "VMware VirtualCenter Server",
		LicenseProductVersion: "6.0",
	},
	Setting:                     &types.ManagedObjectReference{Type: "OptionManager", Value: "VpxSettings"},
	UserDirectory:               &types.ManagedObjectReference{Type: "UserDirectory", Value: "UserDirectory"},
	SessionManager:              &types.ManagedObjectReference{Type: "SessionManager", Value: "SessionManager"},
	AuthorizationManager:        &types.ManagedObjectReference{Type: "AuthorizationManager", Value: "AuthorizationManager"},
	ServiceManager:              &types.ManagedObjectReference{Type: "ServiceManager", Value: "ServiceMgr"},
	PerfManager:                 &types.ManagedObjectReference{Type: "PerformanceManager", Value: "PerfMgr"},
	ScheduledTaskManager:        &types.ManagedObjectReference{Type: "ScheduledTaskManager", Value: "ScheduledTaskManager"},
	AlarmManager:                &types.ManagedObjectReference{Type: "AlarmManager", Value: "AlarmManager"},
	EventManager:                &types.ManagedObjectReference{Type: "EventManager", Value: "EventManager"},
	TaskManager:                 &types.ManagedObjectReference{Type: "TaskManager", Value: "TaskManager"},
	ExtensionManager:            &types.ManagedObjectReference{Type: "ExtensionManager", Value: "ExtensionManager"},
	CustomizationSpecManager:    &types.ManagedObjectReference{Type: "CustomizationSpecManager", Value: "CustomizationSpecManager"},
	CustomFieldsManager:         &types.ManagedObjectReference{Type: "CustomFieldsManager", Value: "CustomFieldsManager"},
	AccountManager:              (*types.ManagedObjectReference)(nil),
	DiagnosticManager:           &types.ManagedObjectReference{Type: "DiagnosticManager", Value: "DiagMgr"},
	LicenseManager:              &types.ManagedObjectReference{Type: "LicenseManager", Value: "LicenseManager"},
	SearchIndex:                 &types.ManagedObjectReference{Type: "SearchIndex", Value: "SearchIndex"},
	FileManager:                 &types.ManagedObjectReference{Type: "FileManager", Value: "FileManager"},
	DatastoreNamespaceManager:   &types.ManagedObjectReference{Type: "DatastoreNamespaceManager", Value: "DatastoreNamespaceManager"},
	VirtualDiskManager:          &types.ManagedObjectReference{Type: "VirtualDiskManager", Value: "virtualDiskManager"},
	VirtualizationManager:       (*types.ManagedObjectReference)(nil),
	SnmpSystem:                  &types.ManagedObjectReference{Type: "HostSnmpSystem", Value: "SnmpSystem"},
	VmProvisioningChecker:       &types.ManagedObjectReference{Type: "VirtualMachineProvisioningChecker", Value: "ProvChecker"},
	VmCompatibilityChecker:      &types.ManagedObjectReference{Type: "VirtualMachineCompatibilityChecker", Value: "CompatChecker"},
	OvfManager:                  &types.ManagedObjectReference{Type: "OvfManager", Value: "OvfManager"},
	IpPoolManager:               &types.ManagedObjectReference{Type: "IpPoolManager", Value: "IpPoolManager"},
	DvSwitchManager:             &types.ManagedObjectReference{Type: "DistributedVirtualSwitchManager", Value: "DVSManager"},
	HostProfileManager:          &types.ManagedObjectReference{Type: "HostProfileManager", Value: "HostProfileManager"},
	ClusterProfileManager:       &types.ManagedObjectReference{Type: "ClusterProfileManager", Value: "ClusterProfileManager"},
	ComplianceManager:           &types.ManagedObjectReference{Type: "ProfileComplianceManager", Value: "MoComplianceManager"},
	LocalizationManager:         &types.ManagedObjectReference{Type: "LocalizationManager", Value: "LocalizationManager"},
	StorageResourceManager:      &types.ManagedObjectReference{Type: "StorageResourceManager", Value: "StorageResourceManager"},
	GuestOperationsManager:      &types.ManagedObjectReference{Type: "GuestOperationsManager", Value: "guestOperationsManager"},
	OverheadMemoryManager:       &types.ManagedObjectReference{Type: "OverheadMemoryManager", Value: "OverheadMemoryManager"},
	CertificateManager:          &types.ManagedObjectReference{Type: "CertificateManager", Value: "certificateManager"},
	IoFilterManager:             &types.ManagedObjectReference{Type: "IoFilterManager", Value: "IoFilterManager"},
	VStorageObjectManager:       &types.ManagedObjectReference{Type: "VcenterVStorageObjectManager", Value: "VStorageObjectManager"},
	HostSpecManager:             &types.ManagedObjectReference{Type: "HostSpecificationManager", Value: "HostSpecificationManager"},
	CryptoManager:               &types.ManagedObjectReference{Type: "CryptoManagerKmip", Value: "CryptoManager"},
	HealthUpdateManager:         &types.ManagedObjectReference{Type: "HealthUpdateManager", Value: "HealthUpdateManager"},
	FailoverClusterConfigurator: &types.ManagedObjectReference{Type: "FailoverClusterConfigurator", Value: "FailoverClusterConfigurator"},
	FailoverClusterManager:      &types.ManagedObjectReference{Type: "FailoverClusterManager", Value: "FailoverClusterManager"},
	TenantManager:               &types.ManagedObjectReference{Type: "TenantTenantManager", Value: "TenantManager"},
}
