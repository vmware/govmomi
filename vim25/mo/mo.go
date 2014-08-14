/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package mo

import (
	"time"

	"github.com/vmware/govmomi/vim25/types"
)

type Alarm struct {
	ExtensibleManagedObject

	Info types.AlarmInfo `mo:"info"`
}

type AlarmManager struct {
	DefaultExpression []types.BaseAlarmExpression `mo:"defaultExpression"`
	Description       types.AlarmDescription      `mo:"description"`
}

type AuthorizationManager struct {
	PrivilegeList []types.AuthorizationPrivilege `mo:"privilegeList"`
	RoleList      []types.AuthorizationRole      `mo:"roleList"`
	Description   types.AuthorizationDescription `mo:"description"`
}

type ClusterComputeResource struct {
	ComputeResource

	Configuration     types.ClusterConfigInfo          `mo:"configuration"`
	Recommendation    []types.ClusterRecommendation    `mo:"recommendation"`
	DrsRecommendation []types.ClusterDrsRecommendation `mo:"drsRecommendation"`
	MigrationHistory  []types.ClusterDrsMigration      `mo:"migrationHistory"`
	ActionHistory     []types.ClusterActionHistory     `mo:"actionHistory"`
	DrsFault          []types.ClusterDrsFaults         `mo:"drsFault"`
}

type ComputeResource struct {
	ManagedEntity

	ResourcePool       *types.ManagedObjectReference       `mo:"resourcePool"`
	Host               []types.ManagedObjectReference      `mo:"host"`
	Datastore          []types.ManagedObjectReference      `mo:"datastore"`
	Network            []types.ManagedObjectReference      `mo:"network"`
	Summary            types.BaseComputeResourceSummary    `mo:"summary"`
	EnvironmentBrowser *types.ManagedObjectReference       `mo:"environmentBrowser"`
	ConfigurationEx    types.BaseComputeResourceConfigInfo `mo:"configurationEx"`
}

type ContainerView struct {
	ManagedObjectView

	Container types.ManagedObjectReference `mo:"container"`
	Type      []string                     `mo:"type"`
	Recursive bool                         `mo:"recursive"`
}

type CustomFieldsManager struct {
	Field []types.CustomFieldDef `mo:"field"`
}

type CustomizationSpecManager struct {
	Info          []types.CustomizationSpecInfo `mo:"info"`
	EncryptionKey []byte                        `mo:"encryptionKey"`
}

type Datacenter struct {
	ManagedEntity

	VmFolder        types.ManagedObjectReference   `mo:"vmFolder"`
	HostFolder      types.ManagedObjectReference   `mo:"hostFolder"`
	DatastoreFolder types.ManagedObjectReference   `mo:"datastoreFolder"`
	NetworkFolder   types.ManagedObjectReference   `mo:"networkFolder"`
	Datastore       []types.ManagedObjectReference `mo:"datastore"`
	Network         []types.ManagedObjectReference `mo:"network"`
	Configuration   types.DatacenterConfigInfo     `mo:"configuration"`
}

type Datastore struct {
	ManagedEntity

	Info              types.BaseDatastoreInfo        `mo:"info"`
	Summary           types.DatastoreSummary         `mo:"summary"`
	Host              []types.DatastoreHostMount     `mo:"host"`
	Vm                []types.ManagedObjectReference `mo:"vm"`
	Browser           types.ManagedObjectReference   `mo:"browser"`
	Capability        types.DatastoreCapability      `mo:"capability"`
	IormConfiguration *types.StorageIORMInfo         `mo:"iormConfiguration"`
}

type DistributedVirtualPortgroup struct {
	Network

	Key      string                      `mo:"key"`
	Config   types.DVPortgroupConfigInfo `mo:"config"`
	PortKeys []string                    `mo:"portKeys"`
}

type DistributedVirtualSwitch struct {
	ManagedEntity

	Uuid                string                         `mo:"uuid"`
	Capability          types.DVSCapability            `mo:"capability"`
	Summary             types.DVSSummary               `mo:"summary"`
	Config              types.BaseDVSConfigInfo        `mo:"config"`
	NetworkResourcePool []types.DVSNetworkResourcePool `mo:"networkResourcePool"`
	Portgroup           []types.ManagedObjectReference `mo:"portgroup"`
	Runtime             *types.DVSRuntimeInfo          `mo:"runtime"`
}

type EnvironmentBrowser struct {
	DatastoreBrowser *types.ManagedObjectReference `mo:"datastoreBrowser"`
}

type EventHistoryCollector struct {
	HistoryCollector

	LatestPage []types.BaseEvent `mo:"latestPage"`
}

type EventManager struct {
	Description  types.EventDescription `mo:"description"`
	LatestEvent  types.BaseEvent        `mo:"latestEvent"`
	MaxCollector int                    `mo:"maxCollector"`
}

type ExtensibleManagedObject struct {
	Value          []types.BaseCustomFieldValue `mo:"value"`
	AvailableField []types.CustomFieldDef       `mo:"availableField"`
}

type ExtensionManager struct {
	ExtensionList []types.Extension `mo:"extensionList"`
}

type Folder struct {
	ManagedEntity

	ChildType   []string                       `mo:"childType"`
	ChildEntity []types.ManagedObjectReference `mo:"childEntity"`
}

type GuestOperationsManager struct {
	AuthManager    *types.ManagedObjectReference `mo:"authManager"`
	FileManager    *types.ManagedObjectReference `mo:"fileManager"`
	ProcessManager *types.ManagedObjectReference `mo:"processManager"`
}

type HistoryCollector struct {
	Filter types.AnyType `mo:"filter"`
}

type HostAuthenticationManager struct {
	Info           types.HostAuthenticationManagerInfo `mo:"info"`
	SupportedStore []types.ManagedObjectReference      `mo:"supportedStore"`
}

type HostAuthenticationStore struct {
	Info types.BaseHostAuthenticationStoreInfo `mo:"info"`
}

type HostAutoStartManager struct {
	Config types.HostAutoStartManagerConfig `mo:"config"`
}

type HostCacheConfigurationManager struct {
	CacheConfigurationInfo []types.HostCacheConfigurationInfo `mo:"cacheConfigurationInfo"`
}

type HostCpuSchedulerSystem struct {
	ExtensibleManagedObject

	HyperthreadInfo *types.HostHyperThreadScheduleInfo `mo:"hyperthreadInfo"`
}

type HostDatastoreBrowser struct {
	Datastore     []types.ManagedObjectReference `mo:"datastore"`
	SupportedType []types.BaseFileQuery          `mo:"supportedType"`
}

type HostDatastoreSystem struct {
	Datastore    []types.ManagedObjectReference        `mo:"datastore"`
	Capabilities types.HostDatastoreSystemCapabilities `mo:"capabilities"`
}

type HostDateTimeSystem struct {
	DateTimeInfo types.HostDateTimeInfo `mo:"dateTimeInfo"`
}

type HostDiagnosticSystem struct {
	ActivePartition *types.HostDiagnosticPartition `mo:"activePartition"`
}

type HostEsxAgentHostManager struct {
	ConfigInfo types.HostEsxAgentHostManagerConfigInfo `mo:"configInfo"`
}

type HostFirewallSystem struct {
	ExtensibleManagedObject

	FirewallInfo *types.HostFirewallInfo `mo:"firewallInfo"`
}

type HostGraphicsManager struct {
	ExtensibleManagedObject

	GraphicsInfo []types.HostGraphicsInfo `mo:"graphicsInfo"`
}

type HostHealthStatusSystem struct {
	Runtime types.HealthSystemRuntime `mo:"runtime"`
}

type HostMemorySystem struct {
	ExtensibleManagedObject

	ConsoleReservationInfo        *types.ServiceConsoleReservationInfo       `mo:"consoleReservationInfo"`
	VirtualMachineReservationInfo *types.VirtualMachineMemoryReservationInfo `mo:"virtualMachineReservationInfo"`
}

type HostNetworkSystem struct {
	ExtensibleManagedObject

	Capabilities         *types.HostNetCapabilities        `mo:"capabilities"`
	NetworkInfo          *types.HostNetworkInfo            `mo:"networkInfo"`
	OffloadCapabilities  *types.HostNetOffloadCapabilities `mo:"offloadCapabilities"`
	NetworkConfig        *types.HostNetworkConfig          `mo:"networkConfig"`
	DnsConfig            types.BaseHostDnsConfig           `mo:"dnsConfig"`
	IpRouteConfig        types.BaseHostIpRouteConfig       `mo:"ipRouteConfig"`
	ConsoleIpRouteConfig types.BaseHostIpRouteConfig       `mo:"consoleIpRouteConfig"`
}

type HostPciPassthruSystem struct {
	ExtensibleManagedObject

	PciPassthruInfo []types.BaseHostPciPassthruInfo `mo:"pciPassthruInfo"`
}

type HostPowerSystem struct {
	Capability types.PowerSystemCapability `mo:"capability"`
	Info       types.PowerSystemInfo       `mo:"info"`
}

type HostProfile struct {
	Profile

	ReferenceHost *types.ManagedObjectReference `mo:"referenceHost"`
}

type HostServiceSystem struct {
	ExtensibleManagedObject

	ServiceInfo types.HostServiceInfo `mo:"serviceInfo"`
}

type HostSnmpSystem struct {
	Configuration types.HostSnmpConfigSpec        `mo:"configuration"`
	Limits        types.HostSnmpSystemAgentLimits `mo:"limits"`
}

type HostStorageSystem struct {
	ExtensibleManagedObject

	StorageDeviceInfo    *types.HostStorageDeviceInfo   `mo:"storageDeviceInfo"`
	FileSystemVolumeInfo types.HostFileSystemVolumeInfo `mo:"fileSystemVolumeInfo"`
	SystemFile           []string                       `mo:"systemFile"`
	MultipathStateInfo   *types.HostMultipathStateInfo  `mo:"multipathStateInfo"`
}

type HostSystem struct {
	ManagedEntity

	Runtime            types.HostRuntimeInfo            `mo:"runtime"`
	Summary            types.HostListSummary            `mo:"summary"`
	Hardware           *types.HostHardwareInfo          `mo:"hardware"`
	Capability         *types.HostCapability            `mo:"capability"`
	LicensableResource types.HostLicensableResourceInfo `mo:"licensableResource"`
	ConfigManager      types.HostConfigManager          `mo:"configManager"`
	Config             *types.HostConfigInfo            `mo:"config"`
	Vm                 []types.ManagedObjectReference   `mo:"vm"`
	Datastore          []types.ManagedObjectReference   `mo:"datastore"`
	Network            []types.ManagedObjectReference   `mo:"network"`
	DatastoreBrowser   types.ManagedObjectReference     `mo:"datastoreBrowser"`
	SystemResources    *types.HostSystemResourceInfo    `mo:"systemResources"`
}

type HostVFlashManager struct {
	VFlashConfigInfo *types.HostVFlashManagerVFlashConfigInfo `mo:"vFlashConfigInfo"`
}

type HostVMotionSystem struct {
	ExtensibleManagedObject

	NetConfig *types.HostVMotionNetConfig `mo:"netConfig"`
	IpConfig  *types.HostIpConfig         `mo:"ipConfig"`
}

type HostVirtualNicManager struct {
	ExtensibleManagedObject

	Info types.HostVirtualNicManagerInfo `mo:"info"`
}

type HostVsanSystem struct {
	Config types.VsanHostConfigInfo `mo:"config"`
}

type HttpNfcLease struct {
	InitializeProgress int                         `mo:"initializeProgress"`
	Info               *types.HttpNfcLeaseInfo     `mo:"info"`
	State              types.HttpNfcLeaseState     `mo:"state"`
	Error              *types.LocalizedMethodFault `mo:"error"`
}

type LicenseManager struct {
	Source                   types.BaseLicenseSource            `mo:"source"`
	SourceAvailable          bool                               `mo:"sourceAvailable"`
	Diagnostics              *types.LicenseDiagnostics          `mo:"diagnostics"`
	FeatureInfo              []types.LicenseFeatureInfo         `mo:"featureInfo"`
	LicensedEdition          string                             `mo:"licensedEdition"`
	Licenses                 []types.LicenseManagerLicenseInfo  `mo:"licenses"`
	LicenseAssignmentManager *types.ManagedObjectReference      `mo:"licenseAssignmentManager"`
	Evaluation               types.LicenseManagerEvaluationInfo `mo:"evaluation"`
}

type LocalizationManager struct {
	Catalog []types.LocalizationManagerMessageCatalog `mo:"catalog"`
}

type ManagedEntity struct {
	ExtensibleManagedObject

	Parent              *types.ManagedObjectReference  `mo:"parent"`
	CustomValue         []types.BaseCustomFieldValue   `mo:"customValue"`
	OverallStatus       types.ManagedEntityStatus      `mo:"overallStatus"`
	ConfigStatus        types.ManagedEntityStatus      `mo:"configStatus"`
	ConfigIssue         []types.BaseEvent              `mo:"configIssue"`
	EffectiveRole       []int                          `mo:"effectiveRole"`
	Permission          []types.Permission             `mo:"permission"`
	Name                string                         `mo:"name"`
	DisabledMethod      []string                       `mo:"disabledMethod"`
	RecentTask          []types.ManagedObjectReference `mo:"recentTask"`
	DeclaredAlarmState  []types.AlarmState             `mo:"declaredAlarmState"`
	TriggeredAlarmState []types.AlarmState             `mo:"triggeredAlarmState"`
	AlarmActionsEnabled *bool                          `mo:"alarmActionsEnabled"`
	Tag                 []types.Tag                    `mo:"tag"`
}

type ManagedObjectView struct {
	View []types.ManagedObjectReference `mo:"view"`
}

type Network struct {
	ManagedEntity

	Name    string                         `mo:"name"`
	Summary types.NetworkSummary           `mo:"summary"`
	Host    []types.ManagedObjectReference `mo:"host"`
	Vm      []types.ManagedObjectReference `mo:"vm"`
}

type OptionManager struct {
	SupportedOption []types.OptionDef   `mo:"supportedOption"`
	Setting         []types.OptionValue `mo:"setting"`
}

type OvfManager struct {
	OvfImportOption []types.OvfOptionInfo `mo:"ovfImportOption"`
	OvfExportOption []types.OvfOptionInfo `mo:"ovfExportOption"`
}

type PerformanceManager struct {
	Description        types.PerformanceDescription `mo:"description"`
	HistoricalInterval []types.PerfInterval         `mo:"historicalInterval"`
	PerfCounter        []types.PerfCounterInfo      `mo:"perfCounter"`
}

type Profile struct {
	Config           types.BaseProfileConfigInfo    `mo:"config"`
	Description      *types.ProfileDescription      `mo:"description"`
	Name             string                         `mo:"name"`
	CreatedTime      time.Time                      `mo:"createdTime"`
	ModifiedTime     time.Time                      `mo:"modifiedTime"`
	Entity           []types.ManagedObjectReference `mo:"entity"`
	ComplianceStatus string                         `mo:"complianceStatus"`
}

type ProfileManager struct {
	Profile []types.ManagedObjectReference `mo:"profile"`
}

type PropertyCollector struct {
	Filter []types.ManagedObjectReference `mo:"filter"`
}

type PropertyFilter struct {
	Spec           types.PropertyFilterSpec `mo:"spec"`
	PartialUpdates bool                     `mo:"partialUpdates"`
}

type ResourcePool struct {
	ManagedEntity

	Summary            types.BaseResourcePoolSummary  `mo:"summary"`
	Runtime            types.ResourcePoolRuntimeInfo  `mo:"runtime"`
	Owner              types.ManagedObjectReference   `mo:"owner"`
	ResourcePool       []types.ManagedObjectReference `mo:"resourcePool"`
	Vm                 []types.ManagedObjectReference `mo:"vm"`
	Config             types.ResourceConfigSpec       `mo:"config"`
	ChildConfiguration []types.ResourceConfigSpec     `mo:"childConfiguration"`
}

type ScheduledTask struct {
	ExtensibleManagedObject

	Info types.ScheduledTaskInfo `mo:"info"`
}

type ScheduledTaskManager struct {
	ScheduledTask []types.ManagedObjectReference `mo:"scheduledTask"`
	Description   types.ScheduledTaskDescription `mo:"description"`
}

type ServiceInstance struct {
	ServerClock time.Time            `mo:"serverClock"`
	Capability  types.Capability     `mo:"capability"`
	Content     types.ServiceContent `mo:"content"`
}

type ServiceManager struct {
	Service []types.ServiceManagerServiceInfo `mo:"service"`
}

type SessionManager struct {
	SessionList         []types.UserSession `mo:"sessionList"`
	CurrentSession      *types.UserSession  `mo:"currentSession"`
	Message             *string             `mo:"message"`
	MessageLocaleList   []string            `mo:"messageLocaleList"`
	SupportedLocaleList []string            `mo:"supportedLocaleList"`
	DefaultLocale       string              `mo:"defaultLocale"`
}

type SimpleCommand struct {
	EncodingType types.SimpleCommandEncoding     `mo:"encodingType"`
	Entity       types.ServiceManagerServiceInfo `mo:"entity"`
}

type StoragePod struct {
	Folder

	Summary            *types.StoragePodSummary  `mo:"summary"`
	PodStorageDrsEntry *types.PodStorageDrsEntry `mo:"podStorageDrsEntry"`
}

type Task struct {
	ExtensibleManagedObject

	Info types.TaskInfo `mo:"info"`
}

type TaskHistoryCollector struct {
	HistoryCollector

	LatestPage []types.TaskInfo `mo:"latestPage"`
}

type TaskManager struct {
	RecentTask   []types.ManagedObjectReference `mo:"recentTask"`
	Description  types.TaskDescription          `mo:"description"`
	MaxCollector int                            `mo:"maxCollector"`
}

type UserDirectory struct {
	DomainList []string `mo:"domainList"`
}

type ViewManager struct {
	ViewList []types.ManagedObjectReference `mo:"viewList"`
}

type VirtualApp struct {
	ResourcePool

	ParentFolder *types.ManagedObjectReference  `mo:"parentFolder"`
	Datastore    []types.ManagedObjectReference `mo:"datastore"`
	Network      []types.ManagedObjectReference `mo:"network"`
	VAppConfig   *types.VAppConfigInfo          `mo:"vAppConfig"`
	ParentVApp   *types.ManagedObjectReference  `mo:"parentVApp"`
	ChildLink    []types.VirtualAppLinkInfo     `mo:"childLink"`
}

type VirtualMachine struct {
	ManagedEntity

	Capability           types.VirtualMachineCapability    `mo:"capability"`
	Config               *types.VirtualMachineConfigInfo   `mo:"config"`
	Layout               *types.VirtualMachineFileLayout   `mo:"layout"`
	LayoutEx             *types.VirtualMachineFileLayoutEx `mo:"layoutEx"`
	Storage              *types.VirtualMachineStorageInfo  `mo:"storage"`
	EnvironmentBrowser   types.ManagedObjectReference      `mo:"environmentBrowser"`
	ResourcePool         *types.ManagedObjectReference     `mo:"resourcePool"`
	ParentVApp           *types.ManagedObjectReference     `mo:"parentVApp"`
	ResourceConfig       *types.ResourceConfigSpec         `mo:"resourceConfig"`
	Runtime              types.VirtualMachineRuntimeInfo   `mo:"runtime"`
	Guest                *types.GuestInfo                  `mo:"guest"`
	Summary              types.VirtualMachineSummary       `mo:"summary"`
	Datastore            []types.ManagedObjectReference    `mo:"datastore"`
	Network              []types.ManagedObjectReference    `mo:"network"`
	Snapshot             *types.VirtualMachineSnapshotInfo `mo:"snapshot"`
	RootSnapshot         []types.ManagedObjectReference    `mo:"rootSnapshot"`
	GuestHeartbeatStatus types.ManagedEntityStatus         `mo:"guestHeartbeatStatus"`
}

type VirtualMachineSnapshot struct {
	ExtensibleManagedObject

	Config        types.VirtualMachineConfigInfo `mo:"config"`
	ChildSnapshot []types.ManagedObjectReference `mo:"childSnapshot"`
}
