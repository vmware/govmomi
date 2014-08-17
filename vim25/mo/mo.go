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
	"reflect"
	"time"

	"github.com/vmware/govmomi/vim25/types"
)

type Alarm struct {
	ExtensibleManagedObject

	Info types.AlarmInfo `mo:"info"`
}

func init() {
	t["Alarm"] = reflect.TypeOf((*Alarm)(nil)).Elem()
}

type AlarmManager struct {
	Ref types.ManagedObjectReference

	DefaultExpression []types.BaseAlarmExpression `mo:"defaultExpression"`
	Description       types.AlarmDescription      `mo:"description"`
}

func (m AlarmManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["AlarmManager"] = reflect.TypeOf((*AlarmManager)(nil)).Elem()
}

type AuthorizationManager struct {
	Ref types.ManagedObjectReference

	PrivilegeList []types.AuthorizationPrivilege `mo:"privilegeList"`
	RoleList      []types.AuthorizationRole      `mo:"roleList"`
	Description   types.AuthorizationDescription `mo:"description"`
}

func (m AuthorizationManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["AuthorizationManager"] = reflect.TypeOf((*AuthorizationManager)(nil)).Elem()
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

func init() {
	t["ClusterComputeResource"] = reflect.TypeOf((*ClusterComputeResource)(nil)).Elem()
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

func init() {
	t["ComputeResource"] = reflect.TypeOf((*ComputeResource)(nil)).Elem()
}

type ContainerView struct {
	ManagedObjectView

	Container types.ManagedObjectReference `mo:"container"`
	Type      []string                     `mo:"type"`
	Recursive bool                         `mo:"recursive"`
}

func init() {
	t["ContainerView"] = reflect.TypeOf((*ContainerView)(nil)).Elem()
}

type CustomFieldsManager struct {
	Ref types.ManagedObjectReference

	Field []types.CustomFieldDef `mo:"field"`
}

func (m CustomFieldsManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["CustomFieldsManager"] = reflect.TypeOf((*CustomFieldsManager)(nil)).Elem()
}

type CustomizationSpecManager struct {
	Ref types.ManagedObjectReference

	Info          []types.CustomizationSpecInfo `mo:"info"`
	EncryptionKey []byte                        `mo:"encryptionKey"`
}

func (m CustomizationSpecManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["CustomizationSpecManager"] = reflect.TypeOf((*CustomizationSpecManager)(nil)).Elem()
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

func init() {
	t["Datacenter"] = reflect.TypeOf((*Datacenter)(nil)).Elem()
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

func init() {
	t["Datastore"] = reflect.TypeOf((*Datastore)(nil)).Elem()
}

type DistributedVirtualPortgroup struct {
	Network

	Key      string                      `mo:"key"`
	Config   types.DVPortgroupConfigInfo `mo:"config"`
	PortKeys []string                    `mo:"portKeys"`
}

func init() {
	t["DistributedVirtualPortgroup"] = reflect.TypeOf((*DistributedVirtualPortgroup)(nil)).Elem()
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

func init() {
	t["DistributedVirtualSwitch"] = reflect.TypeOf((*DistributedVirtualSwitch)(nil)).Elem()
}

type EnvironmentBrowser struct {
	Ref types.ManagedObjectReference

	DatastoreBrowser *types.ManagedObjectReference `mo:"datastoreBrowser"`
}

func (m EnvironmentBrowser) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["EnvironmentBrowser"] = reflect.TypeOf((*EnvironmentBrowser)(nil)).Elem()
}

type EventHistoryCollector struct {
	HistoryCollector

	LatestPage []types.BaseEvent `mo:"latestPage"`
}

func init() {
	t["EventHistoryCollector"] = reflect.TypeOf((*EventHistoryCollector)(nil)).Elem()
}

type EventManager struct {
	Ref types.ManagedObjectReference

	Description  types.EventDescription `mo:"description"`
	LatestEvent  types.BaseEvent        `mo:"latestEvent"`
	MaxCollector int                    `mo:"maxCollector"`
}

func (m EventManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["EventManager"] = reflect.TypeOf((*EventManager)(nil)).Elem()
}

type ExtensibleManagedObject struct {
	Ref types.ManagedObjectReference

	Value          []types.BaseCustomFieldValue `mo:"value"`
	AvailableField []types.CustomFieldDef       `mo:"availableField"`
}

func (m ExtensibleManagedObject) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["ExtensibleManagedObject"] = reflect.TypeOf((*ExtensibleManagedObject)(nil)).Elem()
}

type ExtensionManager struct {
	Ref types.ManagedObjectReference

	ExtensionList []types.Extension `mo:"extensionList"`
}

func (m ExtensionManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["ExtensionManager"] = reflect.TypeOf((*ExtensionManager)(nil)).Elem()
}

type Folder struct {
	ManagedEntity

	ChildType   []string                       `mo:"childType"`
	ChildEntity []types.ManagedObjectReference `mo:"childEntity"`
}

func init() {
	t["Folder"] = reflect.TypeOf((*Folder)(nil)).Elem()
}

type GuestOperationsManager struct {
	Ref types.ManagedObjectReference

	AuthManager    *types.ManagedObjectReference `mo:"authManager"`
	FileManager    *types.ManagedObjectReference `mo:"fileManager"`
	ProcessManager *types.ManagedObjectReference `mo:"processManager"`
}

func (m GuestOperationsManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["GuestOperationsManager"] = reflect.TypeOf((*GuestOperationsManager)(nil)).Elem()
}

type HistoryCollector struct {
	Ref types.ManagedObjectReference

	Filter types.AnyType `mo:"filter"`
}

func (m HistoryCollector) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HistoryCollector"] = reflect.TypeOf((*HistoryCollector)(nil)).Elem()
}

type HostAuthenticationManager struct {
	Ref types.ManagedObjectReference

	Info           types.HostAuthenticationManagerInfo `mo:"info"`
	SupportedStore []types.ManagedObjectReference      `mo:"supportedStore"`
}

func (m HostAuthenticationManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HostAuthenticationManager"] = reflect.TypeOf((*HostAuthenticationManager)(nil)).Elem()
}

type HostAuthenticationStore struct {
	Ref types.ManagedObjectReference

	Info types.BaseHostAuthenticationStoreInfo `mo:"info"`
}

func (m HostAuthenticationStore) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HostAuthenticationStore"] = reflect.TypeOf((*HostAuthenticationStore)(nil)).Elem()
}

type HostAutoStartManager struct {
	Ref types.ManagedObjectReference

	Config types.HostAutoStartManagerConfig `mo:"config"`
}

func (m HostAutoStartManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HostAutoStartManager"] = reflect.TypeOf((*HostAutoStartManager)(nil)).Elem()
}

type HostCacheConfigurationManager struct {
	Ref types.ManagedObjectReference

	CacheConfigurationInfo []types.HostCacheConfigurationInfo `mo:"cacheConfigurationInfo"`
}

func (m HostCacheConfigurationManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HostCacheConfigurationManager"] = reflect.TypeOf((*HostCacheConfigurationManager)(nil)).Elem()
}

type HostCpuSchedulerSystem struct {
	ExtensibleManagedObject

	HyperthreadInfo *types.HostHyperThreadScheduleInfo `mo:"hyperthreadInfo"`
}

func init() {
	t["HostCpuSchedulerSystem"] = reflect.TypeOf((*HostCpuSchedulerSystem)(nil)).Elem()
}

type HostDatastoreBrowser struct {
	Ref types.ManagedObjectReference

	Datastore     []types.ManagedObjectReference `mo:"datastore"`
	SupportedType []types.BaseFileQuery          `mo:"supportedType"`
}

func (m HostDatastoreBrowser) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HostDatastoreBrowser"] = reflect.TypeOf((*HostDatastoreBrowser)(nil)).Elem()
}

type HostDatastoreSystem struct {
	Ref types.ManagedObjectReference

	Datastore    []types.ManagedObjectReference        `mo:"datastore"`
	Capabilities types.HostDatastoreSystemCapabilities `mo:"capabilities"`
}

func (m HostDatastoreSystem) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HostDatastoreSystem"] = reflect.TypeOf((*HostDatastoreSystem)(nil)).Elem()
}

type HostDateTimeSystem struct {
	Ref types.ManagedObjectReference

	DateTimeInfo types.HostDateTimeInfo `mo:"dateTimeInfo"`
}

func (m HostDateTimeSystem) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HostDateTimeSystem"] = reflect.TypeOf((*HostDateTimeSystem)(nil)).Elem()
}

type HostDiagnosticSystem struct {
	Ref types.ManagedObjectReference

	ActivePartition *types.HostDiagnosticPartition `mo:"activePartition"`
}

func (m HostDiagnosticSystem) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HostDiagnosticSystem"] = reflect.TypeOf((*HostDiagnosticSystem)(nil)).Elem()
}

type HostEsxAgentHostManager struct {
	Ref types.ManagedObjectReference

	ConfigInfo types.HostEsxAgentHostManagerConfigInfo `mo:"configInfo"`
}

func (m HostEsxAgentHostManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HostEsxAgentHostManager"] = reflect.TypeOf((*HostEsxAgentHostManager)(nil)).Elem()
}

type HostFirewallSystem struct {
	ExtensibleManagedObject

	FirewallInfo *types.HostFirewallInfo `mo:"firewallInfo"`
}

func init() {
	t["HostFirewallSystem"] = reflect.TypeOf((*HostFirewallSystem)(nil)).Elem()
}

type HostGraphicsManager struct {
	ExtensibleManagedObject

	GraphicsInfo []types.HostGraphicsInfo `mo:"graphicsInfo"`
}

func init() {
	t["HostGraphicsManager"] = reflect.TypeOf((*HostGraphicsManager)(nil)).Elem()
}

type HostHealthStatusSystem struct {
	Ref types.ManagedObjectReference

	Runtime types.HealthSystemRuntime `mo:"runtime"`
}

func (m HostHealthStatusSystem) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HostHealthStatusSystem"] = reflect.TypeOf((*HostHealthStatusSystem)(nil)).Elem()
}

type HostMemorySystem struct {
	ExtensibleManagedObject

	ConsoleReservationInfo        *types.ServiceConsoleReservationInfo       `mo:"consoleReservationInfo"`
	VirtualMachineReservationInfo *types.VirtualMachineMemoryReservationInfo `mo:"virtualMachineReservationInfo"`
}

func init() {
	t["HostMemorySystem"] = reflect.TypeOf((*HostMemorySystem)(nil)).Elem()
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

func init() {
	t["HostNetworkSystem"] = reflect.TypeOf((*HostNetworkSystem)(nil)).Elem()
}

type HostPciPassthruSystem struct {
	ExtensibleManagedObject

	PciPassthruInfo []types.BaseHostPciPassthruInfo `mo:"pciPassthruInfo"`
}

func init() {
	t["HostPciPassthruSystem"] = reflect.TypeOf((*HostPciPassthruSystem)(nil)).Elem()
}

type HostPowerSystem struct {
	Ref types.ManagedObjectReference

	Capability types.PowerSystemCapability `mo:"capability"`
	Info       types.PowerSystemInfo       `mo:"info"`
}

func (m HostPowerSystem) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HostPowerSystem"] = reflect.TypeOf((*HostPowerSystem)(nil)).Elem()
}

type HostProfile struct {
	Profile

	ReferenceHost *types.ManagedObjectReference `mo:"referenceHost"`
}

func init() {
	t["HostProfile"] = reflect.TypeOf((*HostProfile)(nil)).Elem()
}

type HostServiceSystem struct {
	ExtensibleManagedObject

	ServiceInfo types.HostServiceInfo `mo:"serviceInfo"`
}

func init() {
	t["HostServiceSystem"] = reflect.TypeOf((*HostServiceSystem)(nil)).Elem()
}

type HostSnmpSystem struct {
	Ref types.ManagedObjectReference

	Configuration types.HostSnmpConfigSpec        `mo:"configuration"`
	Limits        types.HostSnmpSystemAgentLimits `mo:"limits"`
}

func (m HostSnmpSystem) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HostSnmpSystem"] = reflect.TypeOf((*HostSnmpSystem)(nil)).Elem()
}

type HostStorageSystem struct {
	ExtensibleManagedObject

	StorageDeviceInfo    *types.HostStorageDeviceInfo   `mo:"storageDeviceInfo"`
	FileSystemVolumeInfo types.HostFileSystemVolumeInfo `mo:"fileSystemVolumeInfo"`
	SystemFile           []string                       `mo:"systemFile"`
	MultipathStateInfo   *types.HostMultipathStateInfo  `mo:"multipathStateInfo"`
}

func init() {
	t["HostStorageSystem"] = reflect.TypeOf((*HostStorageSystem)(nil)).Elem()
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

func init() {
	t["HostSystem"] = reflect.TypeOf((*HostSystem)(nil)).Elem()
}

type HostVFlashManager struct {
	Ref types.ManagedObjectReference

	VFlashConfigInfo *types.HostVFlashManagerVFlashConfigInfo `mo:"vFlashConfigInfo"`
}

func (m HostVFlashManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HostVFlashManager"] = reflect.TypeOf((*HostVFlashManager)(nil)).Elem()
}

type HostVMotionSystem struct {
	ExtensibleManagedObject

	NetConfig *types.HostVMotionNetConfig `mo:"netConfig"`
	IpConfig  *types.HostIpConfig         `mo:"ipConfig"`
}

func init() {
	t["HostVMotionSystem"] = reflect.TypeOf((*HostVMotionSystem)(nil)).Elem()
}

type HostVirtualNicManager struct {
	ExtensibleManagedObject

	Info types.HostVirtualNicManagerInfo `mo:"info"`
}

func init() {
	t["HostVirtualNicManager"] = reflect.TypeOf((*HostVirtualNicManager)(nil)).Elem()
}

type HostVsanSystem struct {
	Ref types.ManagedObjectReference

	Config types.VsanHostConfigInfo `mo:"config"`
}

func (m HostVsanSystem) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HostVsanSystem"] = reflect.TypeOf((*HostVsanSystem)(nil)).Elem()
}

type HttpNfcLease struct {
	Ref types.ManagedObjectReference

	InitializeProgress int                         `mo:"initializeProgress"`
	Info               *types.HttpNfcLeaseInfo     `mo:"info"`
	State              types.HttpNfcLeaseState     `mo:"state"`
	Error              *types.LocalizedMethodFault `mo:"error"`
}

func (m HttpNfcLease) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["HttpNfcLease"] = reflect.TypeOf((*HttpNfcLease)(nil)).Elem()
}

type LicenseManager struct {
	Ref types.ManagedObjectReference

	Source                   types.BaseLicenseSource            `mo:"source"`
	SourceAvailable          bool                               `mo:"sourceAvailable"`
	Diagnostics              *types.LicenseDiagnostics          `mo:"diagnostics"`
	FeatureInfo              []types.LicenseFeatureInfo         `mo:"featureInfo"`
	LicensedEdition          string                             `mo:"licensedEdition"`
	Licenses                 []types.LicenseManagerLicenseInfo  `mo:"licenses"`
	LicenseAssignmentManager *types.ManagedObjectReference      `mo:"licenseAssignmentManager"`
	Evaluation               types.LicenseManagerEvaluationInfo `mo:"evaluation"`
}

func (m LicenseManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["LicenseManager"] = reflect.TypeOf((*LicenseManager)(nil)).Elem()
}

type LocalizationManager struct {
	Ref types.ManagedObjectReference

	Catalog []types.LocalizationManagerMessageCatalog `mo:"catalog"`
}

func (m LocalizationManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["LocalizationManager"] = reflect.TypeOf((*LocalizationManager)(nil)).Elem()
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

func init() {
	t["ManagedEntity"] = reflect.TypeOf((*ManagedEntity)(nil)).Elem()
}

type ManagedObjectView struct {
	Ref types.ManagedObjectReference

	View []types.ManagedObjectReference `mo:"view"`
}

func (m ManagedObjectView) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["ManagedObjectView"] = reflect.TypeOf((*ManagedObjectView)(nil)).Elem()
}

type Network struct {
	ManagedEntity

	Name    string                         `mo:"name"`
	Summary types.NetworkSummary           `mo:"summary"`
	Host    []types.ManagedObjectReference `mo:"host"`
	Vm      []types.ManagedObjectReference `mo:"vm"`
}

func init() {
	t["Network"] = reflect.TypeOf((*Network)(nil)).Elem()
}

type OptionManager struct {
	Ref types.ManagedObjectReference

	SupportedOption []types.OptionDef   `mo:"supportedOption"`
	Setting         []types.OptionValue `mo:"setting"`
}

func (m OptionManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["OptionManager"] = reflect.TypeOf((*OptionManager)(nil)).Elem()
}

type OvfManager struct {
	Ref types.ManagedObjectReference

	OvfImportOption []types.OvfOptionInfo `mo:"ovfImportOption"`
	OvfExportOption []types.OvfOptionInfo `mo:"ovfExportOption"`
}

func (m OvfManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["OvfManager"] = reflect.TypeOf((*OvfManager)(nil)).Elem()
}

type PerformanceManager struct {
	Ref types.ManagedObjectReference

	Description        types.PerformanceDescription `mo:"description"`
	HistoricalInterval []types.PerfInterval         `mo:"historicalInterval"`
	PerfCounter        []types.PerfCounterInfo      `mo:"perfCounter"`
}

func (m PerformanceManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["PerformanceManager"] = reflect.TypeOf((*PerformanceManager)(nil)).Elem()
}

type Profile struct {
	Ref types.ManagedObjectReference

	Config           types.BaseProfileConfigInfo    `mo:"config"`
	Description      *types.ProfileDescription      `mo:"description"`
	Name             string                         `mo:"name"`
	CreatedTime      time.Time                      `mo:"createdTime"`
	ModifiedTime     time.Time                      `mo:"modifiedTime"`
	Entity           []types.ManagedObjectReference `mo:"entity"`
	ComplianceStatus string                         `mo:"complianceStatus"`
}

func (m Profile) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["Profile"] = reflect.TypeOf((*Profile)(nil)).Elem()
}

type ProfileManager struct {
	Ref types.ManagedObjectReference

	Profile []types.ManagedObjectReference `mo:"profile"`
}

func (m ProfileManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["ProfileManager"] = reflect.TypeOf((*ProfileManager)(nil)).Elem()
}

type PropertyCollector struct {
	Ref types.ManagedObjectReference

	Filter []types.ManagedObjectReference `mo:"filter"`
}

func (m PropertyCollector) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["PropertyCollector"] = reflect.TypeOf((*PropertyCollector)(nil)).Elem()
}

type PropertyFilter struct {
	Ref types.ManagedObjectReference

	Spec           types.PropertyFilterSpec `mo:"spec"`
	PartialUpdates bool                     `mo:"partialUpdates"`
}

func (m PropertyFilter) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["PropertyFilter"] = reflect.TypeOf((*PropertyFilter)(nil)).Elem()
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

func init() {
	t["ResourcePool"] = reflect.TypeOf((*ResourcePool)(nil)).Elem()
}

type ScheduledTask struct {
	ExtensibleManagedObject

	Info types.ScheduledTaskInfo `mo:"info"`
}

func init() {
	t["ScheduledTask"] = reflect.TypeOf((*ScheduledTask)(nil)).Elem()
}

type ScheduledTaskManager struct {
	Ref types.ManagedObjectReference

	ScheduledTask []types.ManagedObjectReference `mo:"scheduledTask"`
	Description   types.ScheduledTaskDescription `mo:"description"`
}

func (m ScheduledTaskManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["ScheduledTaskManager"] = reflect.TypeOf((*ScheduledTaskManager)(nil)).Elem()
}

type ServiceInstance struct {
	Ref types.ManagedObjectReference

	ServerClock time.Time            `mo:"serverClock"`
	Capability  types.Capability     `mo:"capability"`
	Content     types.ServiceContent `mo:"content"`
}

func (m ServiceInstance) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["ServiceInstance"] = reflect.TypeOf((*ServiceInstance)(nil)).Elem()
}

type ServiceManager struct {
	Ref types.ManagedObjectReference

	Service []types.ServiceManagerServiceInfo `mo:"service"`
}

func (m ServiceManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["ServiceManager"] = reflect.TypeOf((*ServiceManager)(nil)).Elem()
}

type SessionManager struct {
	Ref types.ManagedObjectReference

	SessionList         []types.UserSession `mo:"sessionList"`
	CurrentSession      *types.UserSession  `mo:"currentSession"`
	Message             *string             `mo:"message"`
	MessageLocaleList   []string            `mo:"messageLocaleList"`
	SupportedLocaleList []string            `mo:"supportedLocaleList"`
	DefaultLocale       string              `mo:"defaultLocale"`
}

func (m SessionManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["SessionManager"] = reflect.TypeOf((*SessionManager)(nil)).Elem()
}

type SimpleCommand struct {
	Ref types.ManagedObjectReference

	EncodingType types.SimpleCommandEncoding     `mo:"encodingType"`
	Entity       types.ServiceManagerServiceInfo `mo:"entity"`
}

func (m SimpleCommand) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["SimpleCommand"] = reflect.TypeOf((*SimpleCommand)(nil)).Elem()
}

type StoragePod struct {
	Folder

	Summary            *types.StoragePodSummary  `mo:"summary"`
	PodStorageDrsEntry *types.PodStorageDrsEntry `mo:"podStorageDrsEntry"`
}

func init() {
	t["StoragePod"] = reflect.TypeOf((*StoragePod)(nil)).Elem()
}

type Task struct {
	ExtensibleManagedObject

	Info types.TaskInfo `mo:"info"`
}

func init() {
	t["Task"] = reflect.TypeOf((*Task)(nil)).Elem()
}

type TaskHistoryCollector struct {
	HistoryCollector

	LatestPage []types.TaskInfo `mo:"latestPage"`
}

func init() {
	t["TaskHistoryCollector"] = reflect.TypeOf((*TaskHistoryCollector)(nil)).Elem()
}

type TaskManager struct {
	Ref types.ManagedObjectReference

	RecentTask   []types.ManagedObjectReference `mo:"recentTask"`
	Description  types.TaskDescription          `mo:"description"`
	MaxCollector int                            `mo:"maxCollector"`
}

func (m TaskManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["TaskManager"] = reflect.TypeOf((*TaskManager)(nil)).Elem()
}

type UserDirectory struct {
	Ref types.ManagedObjectReference

	DomainList []string `mo:"domainList"`
}

func (m UserDirectory) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["UserDirectory"] = reflect.TypeOf((*UserDirectory)(nil)).Elem()
}

type ViewManager struct {
	Ref types.ManagedObjectReference

	ViewList []types.ManagedObjectReference `mo:"viewList"`
}

func (m ViewManager) Reference() types.ManagedObjectReference {
	return m.Ref
}

func init() {
	t["ViewManager"] = reflect.TypeOf((*ViewManager)(nil)).Elem()
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

func init() {
	t["VirtualApp"] = reflect.TypeOf((*VirtualApp)(nil)).Elem()
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

func init() {
	t["VirtualMachine"] = reflect.TypeOf((*VirtualMachine)(nil)).Elem()
}

type VirtualMachineSnapshot struct {
	ExtensibleManagedObject

	Config        types.VirtualMachineConfigInfo `mo:"config"`
	ChildSnapshot []types.ManagedObjectReference `mo:"childSnapshot"`
}

func init() {
	t["VirtualMachineSnapshot"] = reflect.TypeOf((*VirtualMachineSnapshot)(nil)).Elem()
}
