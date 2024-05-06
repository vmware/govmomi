/*
Copyright (c) 2014-2024 VMware, Inc. All Rights Reserved.

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

type PbmAssociateAndApplyPolicyStatusPolicyStatus string

const (
	// Policy applied successfully.
	PbmAssociateAndApplyPolicyStatusPolicyStatusSuccess = PbmAssociateAndApplyPolicyStatusPolicyStatus("success")
	// Policy cannot be applied
	PbmAssociateAndApplyPolicyStatusPolicyStatusFailed = PbmAssociateAndApplyPolicyStatusPolicyStatus("failed")
	// Policy cannot be applied
	PbmAssociateAndApplyPolicyStatusPolicyStatusInvalid = PbmAssociateAndApplyPolicyStatusPolicyStatus("invalid")
)

func (e PbmAssociateAndApplyPolicyStatusPolicyStatus) Values() []PbmAssociateAndApplyPolicyStatusPolicyStatus {
	return []PbmAssociateAndApplyPolicyStatusPolicyStatus{
		PbmAssociateAndApplyPolicyStatusPolicyStatusSuccess,
		PbmAssociateAndApplyPolicyStatusPolicyStatusFailed,
		PbmAssociateAndApplyPolicyStatusPolicyStatusInvalid,
	}
}

func (e PbmAssociateAndApplyPolicyStatusPolicyStatus) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmAssociateAndApplyPolicyStatusPolicyStatus", reflect.TypeOf((*PbmAssociateAndApplyPolicyStatusPolicyStatus)(nil)).Elem())
}

// The `PbmBuiltinGenericType_enum` enumerated type defines the list
// of builtin generic datatypes.
//
// See
// `PbmCapabilityGenericTypeInfo*.*PbmCapabilityGenericTypeInfo.genericTypeName`.
//
// A generic datatype indicates how to interpret a collection of values
// of a specific datatype (`PbmCapabilityTypeInfo.typeName`).
type PbmBuiltinGenericType string

const (
	// Indicates a full or partial range of values (`PbmCapabilityRange`).
	//
	// A full range specifies both <code>min</code> and <code>max</code> values.
	// A partial range specifies one or the other, <code>min</code> or <code>max</code>.
	PbmBuiltinGenericTypeVMW_RANGE = PbmBuiltinGenericType("VMW_RANGE")
	// Indicates a single value or a discrete set of values
	// (`PbmCapabilityDiscreteSet`).
	PbmBuiltinGenericTypeVMW_SET = PbmBuiltinGenericType("VMW_SET")
)

func (e PbmBuiltinGenericType) Values() []PbmBuiltinGenericType {
	return []PbmBuiltinGenericType{
		PbmBuiltinGenericTypeVMW_RANGE,
		PbmBuiltinGenericTypeVMW_SET,
	}
}

func (e PbmBuiltinGenericType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmBuiltinGenericType", reflect.TypeOf((*PbmBuiltinGenericType)(nil)).Elem())
}

// The `PbmBuiltinType_enum` enumerated type defines datatypes
// for storage profiles.
//
// Property metadata
// (`PbmCapabilityPropertyMetadata`) uses the builtin types
// to define data types for storage capabilities and requirements.
// It may also specify the semantics that are applied to a collection
// of builtin type values. See `PbmCapabilityTypeInfo`.
// These semantics are specified as a generic builtin type.
// See `PbmCapabilityGenericTypeInfo`.
// The type information determines how capability constraints are interpreted
// `PbmCapabilityPropertyInstance.value`).
type PbmBuiltinType string

const (
	// Unsigned long value.
	//
	// This datatype supports the following constraint values.
	//   - Single value
	//   - Full or partial range of values (`PbmCapabilityRange`)
	//   - Discrete set of values (`PbmCapabilityDiscreteSet`)
	PbmBuiltinTypeXSD_LONG = PbmBuiltinType("XSD_LONG")
	// Datatype not supported.
	PbmBuiltinTypeXSD_SHORT = PbmBuiltinType("XSD_SHORT")
	// Datatype not supported.
	//
	// Use XSD\_INT instead.
	PbmBuiltinTypeXSD_INTEGER = PbmBuiltinType("XSD_INTEGER")
	// Integer value.
	//
	// This datatype supports the following constraint values.
	//   - Single value
	//   - Full or partial range of values (`PbmCapabilityRange`)
	//   - Discrete set of values (`PbmCapabilityDiscreteSet`)
	PbmBuiltinTypeXSD_INT = PbmBuiltinType("XSD_INT")
	// String value.
	//
	// This datatype supports a single value
	// or a discrete set of values (`PbmCapabilityDiscreteSet`).
	PbmBuiltinTypeXSD_STRING = PbmBuiltinType("XSD_STRING")
	// Boolean value.
	PbmBuiltinTypeXSD_BOOLEAN = PbmBuiltinType("XSD_BOOLEAN")
	// Double precision floating point value.
	//
	// This datatype supports the following
	// constraint values.
	//   - Single value
	//   - Full or partial range of values (`PbmCapabilityRange`)
	//   - Discrete set of values (`PbmCapabilityDiscreteSet`)
	PbmBuiltinTypeXSD_DOUBLE = PbmBuiltinType("XSD_DOUBLE")
	// Date and time value.
	PbmBuiltinTypeXSD_DATETIME = PbmBuiltinType("XSD_DATETIME")
	// Timespan value (`PbmCapabilityTimeSpan`).
	//
	// This datatype supports
	// the following constraint values.
	//   - Single value
	//   - Full or partial range of values (`PbmCapabilityRange`)
	//   - Discrete set of values (`PbmCapabilityDiscreteSet`)
	PbmBuiltinTypeVMW_TIMESPAN = PbmBuiltinType("VMW_TIMESPAN")
	PbmBuiltinTypeVMW_POLICY   = PbmBuiltinType("VMW_POLICY")
)

func (e PbmBuiltinType) Values() []PbmBuiltinType {
	return []PbmBuiltinType{
		PbmBuiltinTypeXSD_LONG,
		PbmBuiltinTypeXSD_SHORT,
		PbmBuiltinTypeXSD_INTEGER,
		PbmBuiltinTypeXSD_INT,
		PbmBuiltinTypeXSD_STRING,
		PbmBuiltinTypeXSD_BOOLEAN,
		PbmBuiltinTypeXSD_DOUBLE,
		PbmBuiltinTypeXSD_DATETIME,
		PbmBuiltinTypeVMW_TIMESPAN,
		PbmBuiltinTypeVMW_POLICY,
	}
}

func (e PbmBuiltinType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmBuiltinType", reflect.TypeOf((*PbmBuiltinType)(nil)).Elem())
}

// List of operators that are supported for constructing policy.
//
// Currently only tag based properties can use this operator.
// Other operators can be added as required.
type PbmCapabilityOperator string

const (
	PbmCapabilityOperatorNOT = PbmCapabilityOperator("NOT")
)

func (e PbmCapabilityOperator) Values() []PbmCapabilityOperator {
	return []PbmCapabilityOperator{
		PbmCapabilityOperatorNOT,
	}
}

func (e PbmCapabilityOperator) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmCapabilityOperator", reflect.TypeOf((*PbmCapabilityOperator)(nil)).Elem())
}

// The `PbmCapabilityTimeUnitType_enum` enumeration type
// defines the supported list of time units for profiles that specify
// time span capabilities and constraints.
//
// See `PbmCapabilityTimeSpan`.
type PbmCapabilityTimeUnitType string

const (
	// Constraints and capabilities expressed in units of seconds.
	PbmCapabilityTimeUnitTypeSECONDS = PbmCapabilityTimeUnitType("SECONDS")
	// Constraints and capabilities expressed in units of minutes.
	PbmCapabilityTimeUnitTypeMINUTES = PbmCapabilityTimeUnitType("MINUTES")
	// Constraints and capabilities expressed in units of hours.
	PbmCapabilityTimeUnitTypeHOURS = PbmCapabilityTimeUnitType("HOURS")
	// Constraints and capabilities expressed in units of days.
	PbmCapabilityTimeUnitTypeDAYS = PbmCapabilityTimeUnitType("DAYS")
	// Constraints and capabilities expressed in units of weeks.
	PbmCapabilityTimeUnitTypeWEEKS = PbmCapabilityTimeUnitType("WEEKS")
	// Constraints and capabilities expressed in units of months.
	PbmCapabilityTimeUnitTypeMONTHS = PbmCapabilityTimeUnitType("MONTHS")
	// Constraints and capabilities expressed in units of years.
	PbmCapabilityTimeUnitTypeYEARS = PbmCapabilityTimeUnitType("YEARS")
)

func (e PbmCapabilityTimeUnitType) Values() []PbmCapabilityTimeUnitType {
	return []PbmCapabilityTimeUnitType{
		PbmCapabilityTimeUnitTypeSECONDS,
		PbmCapabilityTimeUnitTypeMINUTES,
		PbmCapabilityTimeUnitTypeHOURS,
		PbmCapabilityTimeUnitTypeDAYS,
		PbmCapabilityTimeUnitTypeWEEKS,
		PbmCapabilityTimeUnitTypeMONTHS,
		PbmCapabilityTimeUnitTypeYEARS,
	}
}

func (e PbmCapabilityTimeUnitType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmCapabilityTimeUnitType", reflect.TypeOf((*PbmCapabilityTimeUnitType)(nil)).Elem())
}

// The `PbmComplianceResultComplianceTaskStatus_enum`
// enumeration type defines the set of task status for compliance
// operations.
//
// See `PbmComplianceResult` and
// `PbmRollupComplianceResult`.
type PbmComplianceResultComplianceTaskStatus string

const (
	// Compliance calculation is in progress.
	PbmComplianceResultComplianceTaskStatusInProgress = PbmComplianceResultComplianceTaskStatus("inProgress")
	// Compliance calculation has succeeded.
	PbmComplianceResultComplianceTaskStatusSuccess = PbmComplianceResultComplianceTaskStatus("success")
	// Compliance calculation failed due to some exception.
	PbmComplianceResultComplianceTaskStatusFailed = PbmComplianceResultComplianceTaskStatus("failed")
)

func (e PbmComplianceResultComplianceTaskStatus) Values() []PbmComplianceResultComplianceTaskStatus {
	return []PbmComplianceResultComplianceTaskStatus{
		PbmComplianceResultComplianceTaskStatusInProgress,
		PbmComplianceResultComplianceTaskStatusSuccess,
		PbmComplianceResultComplianceTaskStatusFailed,
	}
}

func (e PbmComplianceResultComplianceTaskStatus) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmComplianceResultComplianceTaskStatus", reflect.TypeOf((*PbmComplianceResultComplianceTaskStatus)(nil)).Elem())
}

// The `PbmComplianceStatus_enum`
// enumeration type defines the set of status values
// for compliance operations.
//
// See `PbmComplianceResult` and
// `PbmRollupComplianceResult`.
type PbmComplianceStatus string

const (
	// Entity is in compliance.
	PbmComplianceStatusCompliant = PbmComplianceStatus("compliant")
	// Entity is out of compliance.
	PbmComplianceStatusNonCompliant = PbmComplianceStatus("nonCompliant")
	// Compliance status of the entity is not known.
	PbmComplianceStatusUnknown = PbmComplianceStatus("unknown")
	// Compliance computation is not applicable for this entity,
	// because it does not have any storage requirements that
	// apply to the object-based datastore on which this entity is placed.
	PbmComplianceStatusNotApplicable = PbmComplianceStatus("notApplicable")
	// This is the same as `PbmComplianceResult.mismatch`
	// variable.
	//
	// Compliance status becomes out-of-date when the profile
	// associated with the entity is edited and not applied. The compliance
	// status will remain in out-of-date compliance status until the latest
	// policy is applied to the entity.
	PbmComplianceStatusOutOfDate = PbmComplianceStatus("outOfDate")
)

func (e PbmComplianceStatus) Values() []PbmComplianceStatus {
	return []PbmComplianceStatus{
		PbmComplianceStatusCompliant,
		PbmComplianceStatusNonCompliant,
		PbmComplianceStatusUnknown,
		PbmComplianceStatusNotApplicable,
		PbmComplianceStatusOutOfDate,
	}
}

func (e PbmComplianceStatus) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmComplianceStatus", reflect.TypeOf((*PbmComplianceStatus)(nil)).Elem())
}

// This enum corresponds to the keystores used by
// sps.
type PbmDebugManagerKeystoreName string

const (
	// Refers to SMS keystore
	PbmDebugManagerKeystoreNameSMS = PbmDebugManagerKeystoreName("SMS")
	// Refers to TRUSTED\_ROOTS keystore.
	PbmDebugManagerKeystoreNameTRUSTED_ROOTS = PbmDebugManagerKeystoreName("TRUSTED_ROOTS")
)

func (e PbmDebugManagerKeystoreName) Values() []PbmDebugManagerKeystoreName {
	return []PbmDebugManagerKeystoreName{
		PbmDebugManagerKeystoreNameSMS,
		PbmDebugManagerKeystoreNameTRUSTED_ROOTS,
	}
}

func (e PbmDebugManagerKeystoreName) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmDebugManagerKeystoreName", reflect.TypeOf((*PbmDebugManagerKeystoreName)(nil)).Elem())
}

// The enumeration type defines the set of health status values for an entity
// that is part of entity health operation.
type PbmHealthStatusForEntity string

const (
	// For file share: 'red' if the file server for this file share is in error
	// state or any of its backing vSAN objects are degraded.
	//
	// For FCD: 'red' if the datastore on which the FCD resides is not
	// accessible from any of the hosts it is mounted.
	PbmHealthStatusForEntityRed = PbmHealthStatusForEntity("red")
	// For file share: 'yellow' if some backing objects are repairing, i.e.
	//
	// warning state.
	// For FCD: 'yellow' if the datastore on which the entity resides is
	// accessible only from some of the hosts it is mounted but not all.
	PbmHealthStatusForEntityYellow = PbmHealthStatusForEntity("yellow")
	// For file share: 'green' if the file server for this file share is
	// running properly and all its backing vSAN objects are healthy.
	//
	// For FCD: 'green' if the datastore on which the entity resides
	// is accessible from all the hosts it is mounted.
	PbmHealthStatusForEntityGreen = PbmHealthStatusForEntity("green")
	// If the health status of a file share is unknown, not valid for FCD.
	PbmHealthStatusForEntityUnknown = PbmHealthStatusForEntity("unknown")
)

func (e PbmHealthStatusForEntity) Values() []PbmHealthStatusForEntity {
	return []PbmHealthStatusForEntity{
		PbmHealthStatusForEntityRed,
		PbmHealthStatusForEntityYellow,
		PbmHealthStatusForEntityGreen,
		PbmHealthStatusForEntityUnknown,
	}
}

func (e PbmHealthStatusForEntity) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmHealthStatusForEntity", reflect.TypeOf((*PbmHealthStatusForEntity)(nil)).Elem())
}

// Recognized types of an IO Filter.
//
// String constant used in `IofilterInfo#filterType`.
// These should match(upper case) the IO Filter classes as defined by IO Filter framework.
// See https://opengrok.eng.vmware.com/source/xref/vmcore-main.perforce.1666/bora/scons/apps/esx/iofilterApps.sc#33
type PbmIofilterInfoFilterType string

const (
	PbmIofilterInfoFilterTypeINSPECTION         = PbmIofilterInfoFilterType("INSPECTION")
	PbmIofilterInfoFilterTypeCOMPRESSION        = PbmIofilterInfoFilterType("COMPRESSION")
	PbmIofilterInfoFilterTypeENCRYPTION         = PbmIofilterInfoFilterType("ENCRYPTION")
	PbmIofilterInfoFilterTypeREPLICATION        = PbmIofilterInfoFilterType("REPLICATION")
	PbmIofilterInfoFilterTypeCACHE              = PbmIofilterInfoFilterType("CACHE")
	PbmIofilterInfoFilterTypeDATAPROVIDER       = PbmIofilterInfoFilterType("DATAPROVIDER")
	PbmIofilterInfoFilterTypeDATASTOREIOCONTROL = PbmIofilterInfoFilterType("DATASTOREIOCONTROL")
)

func (e PbmIofilterInfoFilterType) Values() []PbmIofilterInfoFilterType {
	return []PbmIofilterInfoFilterType{
		PbmIofilterInfoFilterTypeINSPECTION,
		PbmIofilterInfoFilterTypeCOMPRESSION,
		PbmIofilterInfoFilterTypeENCRYPTION,
		PbmIofilterInfoFilterTypeREPLICATION,
		PbmIofilterInfoFilterTypeCACHE,
		PbmIofilterInfoFilterTypeDATAPROVIDER,
		PbmIofilterInfoFilterTypeDATASTOREIOCONTROL,
	}
}

func (e PbmIofilterInfoFilterType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmIofilterInfoFilterType", reflect.TypeOf((*PbmIofilterInfoFilterType)(nil)).Elem())
}

// Denotes the line of service of a schema.
type PbmLineOfServiceInfoLineOfServiceEnum string

const (
	PbmLineOfServiceInfoLineOfServiceEnumINSPECTION           = PbmLineOfServiceInfoLineOfServiceEnum("INSPECTION")
	PbmLineOfServiceInfoLineOfServiceEnumCOMPRESSION          = PbmLineOfServiceInfoLineOfServiceEnum("COMPRESSION")
	PbmLineOfServiceInfoLineOfServiceEnumENCRYPTION           = PbmLineOfServiceInfoLineOfServiceEnum("ENCRYPTION")
	PbmLineOfServiceInfoLineOfServiceEnumREPLICATION          = PbmLineOfServiceInfoLineOfServiceEnum("REPLICATION")
	PbmLineOfServiceInfoLineOfServiceEnumCACHING              = PbmLineOfServiceInfoLineOfServiceEnum("CACHING")
	PbmLineOfServiceInfoLineOfServiceEnumPERSISTENCE          = PbmLineOfServiceInfoLineOfServiceEnum("PERSISTENCE")
	PbmLineOfServiceInfoLineOfServiceEnumDATA_PROVIDER        = PbmLineOfServiceInfoLineOfServiceEnum("DATA_PROVIDER")
	PbmLineOfServiceInfoLineOfServiceEnumDATASTORE_IO_CONTROL = PbmLineOfServiceInfoLineOfServiceEnum("DATASTORE_IO_CONTROL")
	PbmLineOfServiceInfoLineOfServiceEnumDATA_PROTECTION      = PbmLineOfServiceInfoLineOfServiceEnum("DATA_PROTECTION")
	PbmLineOfServiceInfoLineOfServiceEnumSTRETCHED_CLUSTER    = PbmLineOfServiceInfoLineOfServiceEnum("STRETCHED_CLUSTER")
)

func (e PbmLineOfServiceInfoLineOfServiceEnum) Values() []PbmLineOfServiceInfoLineOfServiceEnum {
	return []PbmLineOfServiceInfoLineOfServiceEnum{
		PbmLineOfServiceInfoLineOfServiceEnumINSPECTION,
		PbmLineOfServiceInfoLineOfServiceEnumCOMPRESSION,
		PbmLineOfServiceInfoLineOfServiceEnumENCRYPTION,
		PbmLineOfServiceInfoLineOfServiceEnumREPLICATION,
		PbmLineOfServiceInfoLineOfServiceEnumCACHING,
		PbmLineOfServiceInfoLineOfServiceEnumPERSISTENCE,
		PbmLineOfServiceInfoLineOfServiceEnumDATA_PROVIDER,
		PbmLineOfServiceInfoLineOfServiceEnumDATASTORE_IO_CONTROL,
		PbmLineOfServiceInfoLineOfServiceEnumDATA_PROTECTION,
		PbmLineOfServiceInfoLineOfServiceEnumSTRETCHED_CLUSTER,
	}
}

func (e PbmLineOfServiceInfoLineOfServiceEnum) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmLineOfServiceInfoLineOfServiceEnum", reflect.TypeOf((*PbmLineOfServiceInfoLineOfServiceEnum)(nil)).Elem())
}

// This enum corresponds to the different packages whose logging
// is configured independently by sps service.
type PbmLoggingConfigurationComponent string

const (
	// Modifies logging level of com.vmware.pbm package.
	PbmLoggingConfigurationComponentPbm = PbmLoggingConfigurationComponent("pbm")
	// Modifies logging level of com.vmware.vslm package.
	PbmLoggingConfigurationComponentVslm = PbmLoggingConfigurationComponent("vslm")
	// Modifies logging level of com.vmware.vim.sms package.
	PbmLoggingConfigurationComponentSms = PbmLoggingConfigurationComponent("sms")
	// Modifies logging level of com.vmware.spbm package.
	PbmLoggingConfigurationComponentSpbm = PbmLoggingConfigurationComponent("spbm")
	// Modifies logging level of com.vmware.sps package.
	PbmLoggingConfigurationComponentSps = PbmLoggingConfigurationComponent("sps")
	// Modifies logging level of httpclient wire header.
	PbmLoggingConfigurationComponentHttpclient_header = PbmLoggingConfigurationComponent("httpclient_header")
	// Modifies logging level of httpclient wire content.
	PbmLoggingConfigurationComponentHttpclient_content = PbmLoggingConfigurationComponent("httpclient_content")
	// Modifies logging level of com.vmware.vim.vmomi package.
	PbmLoggingConfigurationComponentVmomi = PbmLoggingConfigurationComponent("vmomi")
)

func (e PbmLoggingConfigurationComponent) Values() []PbmLoggingConfigurationComponent {
	return []PbmLoggingConfigurationComponent{
		PbmLoggingConfigurationComponentPbm,
		PbmLoggingConfigurationComponentVslm,
		PbmLoggingConfigurationComponentSms,
		PbmLoggingConfigurationComponentSpbm,
		PbmLoggingConfigurationComponentSps,
		PbmLoggingConfigurationComponentHttpclient_header,
		PbmLoggingConfigurationComponentHttpclient_content,
		PbmLoggingConfigurationComponentVmomi,
	}
}

func (e PbmLoggingConfigurationComponent) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmLoggingConfigurationComponent", reflect.TypeOf((*PbmLoggingConfigurationComponent)(nil)).Elem())
}

// This enum corresponds to the different log levels supported
// by sps service.
type PbmLoggingConfigurationLogLevel string

const (
	// Refers to INFO level logging
	PbmLoggingConfigurationLogLevelINFO = PbmLoggingConfigurationLogLevel("INFO")
	// Refers to DEBUG level logging.
	PbmLoggingConfigurationLogLevelDEBUG = PbmLoggingConfigurationLogLevel("DEBUG")
	// Refers to TRACE level logging.
	PbmLoggingConfigurationLogLevelTRACE = PbmLoggingConfigurationLogLevel("TRACE")
)

func (e PbmLoggingConfigurationLogLevel) Values() []PbmLoggingConfigurationLogLevel {
	return []PbmLoggingConfigurationLogLevel{
		PbmLoggingConfigurationLogLevelINFO,
		PbmLoggingConfigurationLogLevelDEBUG,
		PbmLoggingConfigurationLogLevelTRACE,
	}
}

func (e PbmLoggingConfigurationLogLevel) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmLoggingConfigurationLogLevel", reflect.TypeOf((*PbmLoggingConfigurationLogLevel)(nil)).Elem())
}

// The `PbmObjectType_enum` enumerated type
// defines vSphere Server object types that are known
// to the Storage Policy Server.
//
// See `PbmServerObjectRef*.*PbmServerObjectRef.objectType`.
type PbmObjectType string

const (
	// Indicates a virtual machine, not including the disks, identified by the virtual machine
	// identifier _virtual-machine-mor_.
	PbmObjectTypeVirtualMachine = PbmObjectType("virtualMachine")
	// Indicates the virtual machine and all its disks, identified by the virtual machine
	// identifier _virtual-machine-mor_.
	PbmObjectTypeVirtualMachineAndDisks = PbmObjectType("virtualMachineAndDisks")
	// Indicates a virtual disk, identified by disk key
	// (_virtual-machine-mor_:_disk-key_).
	PbmObjectTypeVirtualDiskId = PbmObjectType("virtualDiskId")
	// Indicates a virtual disk, identified by UUID - for First Class Storage Object support.
	PbmObjectTypeVirtualDiskUUID = PbmObjectType("virtualDiskUUID")
	// Indicates a datastore.
	PbmObjectTypeDatastore = PbmObjectType("datastore")
	// Indicates a VSAN object
	PbmObjectTypeVsanObjectId = PbmObjectType("vsanObjectId")
	// Indicates a file service
	PbmObjectTypeFileShareId = PbmObjectType("fileShareId")
	// Unknown object type.
	PbmObjectTypeUnknown = PbmObjectType("unknown")
)

func (e PbmObjectType) Values() []PbmObjectType {
	return []PbmObjectType{
		PbmObjectTypeVirtualMachine,
		PbmObjectTypeVirtualMachineAndDisks,
		PbmObjectTypeVirtualDiskId,
		PbmObjectTypeVirtualDiskUUID,
		PbmObjectTypeDatastore,
		PbmObjectTypeVsanObjectId,
		PbmObjectTypeFileShareId,
		PbmObjectTypeUnknown,
	}
}

func (e PbmObjectType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmObjectType", reflect.TypeOf((*PbmObjectType)(nil)).Elem())
}

// The `PbmOperation_enum` enumerated type
// defines the provisioning operation being performed on the entity like FCD, virtual machine.
type PbmOperation string

const (
	// Indicates create operation of an entity.
	PbmOperationCREATE = PbmOperation("CREATE")
	// Indicates register operation of an entity.
	PbmOperationREGISTER = PbmOperation("REGISTER")
	// Indicates reconfigure operation of an entity.
	PbmOperationRECONFIGURE = PbmOperation("RECONFIGURE")
	// Indicates migrate operation of an entity.
	PbmOperationMIGRATE = PbmOperation("MIGRATE")
	// Indicates clone operation of an entity.
	PbmOperationCLONE = PbmOperation("CLONE")
)

func (e PbmOperation) Values() []PbmOperation {
	return []PbmOperation{
		PbmOperationCREATE,
		PbmOperationREGISTER,
		PbmOperationRECONFIGURE,
		PbmOperationMIGRATE,
		PbmOperationCLONE,
	}
}

func (e PbmOperation) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmOperation", reflect.TypeOf((*PbmOperation)(nil)).Elem())
}

// Volume allocation type constants.
type PbmPolicyAssociationVolumeAllocationType string

const (
	// Space required is fully allocated and initialized.
	//
	// It is wiped clean of any previous content on the
	// physical media. Gives faster runtime IO performance.
	PbmPolicyAssociationVolumeAllocationTypeFullyInitialized = PbmPolicyAssociationVolumeAllocationType("FullyInitialized")
	// Space required is fully allocated.
	//
	// It may contain
	// stale data on the physical media.
	PbmPolicyAssociationVolumeAllocationTypeReserveSpace = PbmPolicyAssociationVolumeAllocationType("ReserveSpace")
	// Space required is allocated and zeroed on demand
	// as the space is used.
	PbmPolicyAssociationVolumeAllocationTypeConserveSpaceWhenPossible = PbmPolicyAssociationVolumeAllocationType("ConserveSpaceWhenPossible")
)

func (e PbmPolicyAssociationVolumeAllocationType) Values() []PbmPolicyAssociationVolumeAllocationType {
	return []PbmPolicyAssociationVolumeAllocationType{
		PbmPolicyAssociationVolumeAllocationTypeFullyInitialized,
		PbmPolicyAssociationVolumeAllocationTypeReserveSpace,
		PbmPolicyAssociationVolumeAllocationTypeConserveSpaceWhenPossible,
	}
}

func (e PbmPolicyAssociationVolumeAllocationType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmPolicyAssociationVolumeAllocationType", reflect.TypeOf((*PbmPolicyAssociationVolumeAllocationType)(nil)).Elem())
}

// The `PbmProfileCategoryEnum_enum`
// enumerated type defines the profile categories for a capability-based
// storage profile.
//
// See
// `PbmCapabilityProfile`.
type PbmProfileCategoryEnum string

const (
	// Indicates a storage requirement.
	//
	// Requirements are based on
	// storage capabilities.
	PbmProfileCategoryEnumREQUIREMENT = PbmProfileCategoryEnum("REQUIREMENT")
	// Indicates a storage capability.
	//
	// Storage capabilities
	// are defined by storage providers.
	PbmProfileCategoryEnumRESOURCE = PbmProfileCategoryEnum("RESOURCE")
	// Indicates a data service policy that can be embedded into
	// another storage policy.
	//
	// Policies of this type can't be assigned to
	// Virtual Machines or Virtual Disks.
	PbmProfileCategoryEnumDATA_SERVICE_POLICY = PbmProfileCategoryEnum("DATA_SERVICE_POLICY")
)

func (e PbmProfileCategoryEnum) Values() []PbmProfileCategoryEnum {
	return []PbmProfileCategoryEnum{
		PbmProfileCategoryEnumREQUIREMENT,
		PbmProfileCategoryEnumRESOURCE,
		PbmProfileCategoryEnumDATA_SERVICE_POLICY,
	}
}

func (e PbmProfileCategoryEnum) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmProfileCategoryEnum", reflect.TypeOf((*PbmProfileCategoryEnum)(nil)).Elem())
}

// The `PbmProfileResourceTypeEnum_enum` enumerated type defines the set of resource
// types that are supported for profile management.
//
// See `PbmProfileResourceType`.
type PbmProfileResourceTypeEnum string

const (
	// Indicates resources that support storage profiles.
	PbmProfileResourceTypeEnumSTORAGE = PbmProfileResourceTypeEnum("STORAGE")
)

func (e PbmProfileResourceTypeEnum) Values() []PbmProfileResourceTypeEnum {
	return []PbmProfileResourceTypeEnum{
		PbmProfileResourceTypeEnumSTORAGE,
	}
}

func (e PbmProfileResourceTypeEnum) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmProfileResourceTypeEnum", reflect.TypeOf((*PbmProfileResourceTypeEnum)(nil)).Elem())
}

// System pre-created profile types.
type PbmSystemCreatedProfileType string

const (
	// Indicates the system pre-created editable VSAN default profile.
	PbmSystemCreatedProfileTypeVsanDefaultProfile = PbmSystemCreatedProfileType("VsanDefaultProfile")
	// Indicates the system pre-created non-editable default profile
	// for VVOL datastores.
	PbmSystemCreatedProfileTypeVVolDefaultProfile = PbmSystemCreatedProfileType("VVolDefaultProfile")
	// Indicates the system pre-created non-editable default profile
	// for PMem datastores
	PbmSystemCreatedProfileTypePmemDefaultProfile = PbmSystemCreatedProfileType("PmemDefaultProfile")
	// Indicates the system pre-created non-editable VMC default profile.
	PbmSystemCreatedProfileTypeVmcManagementProfile = PbmSystemCreatedProfileType("VmcManagementProfile")
	// Indicates the system pre-created non-editable VSANMAX default profile.
	PbmSystemCreatedProfileTypeVsanMaxDefaultProfile = PbmSystemCreatedProfileType("VsanMaxDefaultProfile")
)

func (e PbmSystemCreatedProfileType) Values() []PbmSystemCreatedProfileType {
	return []PbmSystemCreatedProfileType{
		PbmSystemCreatedProfileTypeVsanDefaultProfile,
		PbmSystemCreatedProfileTypeVVolDefaultProfile,
		PbmSystemCreatedProfileTypePmemDefaultProfile,
		PbmSystemCreatedProfileTypeVmcManagementProfile,
		PbmSystemCreatedProfileTypeVsanMaxDefaultProfile,
	}
}

func (e PbmSystemCreatedProfileType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmSystemCreatedProfileType", reflect.TypeOf((*PbmSystemCreatedProfileType)(nil)).Elem())
}

// The `PbmVmOperation_enum` enumerated type
// defines the provisioning operation being performed on the virtual machine.
type PbmVmOperation string

const (
	// Indicates create operation of a virtual machine.
	PbmVmOperationCREATE = PbmVmOperation("CREATE")
	// Indicates reconfigure operation of a virtual machine.
	PbmVmOperationRECONFIGURE = PbmVmOperation("RECONFIGURE")
	// Indicates migrate operation of a virtual machine.
	PbmVmOperationMIGRATE = PbmVmOperation("MIGRATE")
	// Indicates clone operation of a virtual machine.
	PbmVmOperationCLONE = PbmVmOperation("CLONE")
)

func (e PbmVmOperation) Values() []PbmVmOperation {
	return []PbmVmOperation{
		PbmVmOperationCREATE,
		PbmVmOperationRECONFIGURE,
		PbmVmOperationMIGRATE,
		PbmVmOperationCLONE,
	}
}

func (e PbmVmOperation) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmVmOperation", reflect.TypeOf((*PbmVmOperation)(nil)).Elem())
}

// The `PbmVvolType_enum` enumeration type
// defines VVOL types.
//
// VvolType's are referenced to specify which objectType
// to fetch for default capability.
type PbmVvolType string

const (
	// meta-data volume
	PbmVvolTypeConfig = PbmVvolType("Config")
	// vmdk volume
	PbmVvolTypeData = PbmVvolType("Data")
	// swap volume
	PbmVvolTypeSwap = PbmVvolType("Swap")
)

func (e PbmVvolType) Values() []PbmVvolType {
	return []PbmVvolType{
		PbmVvolTypeConfig,
		PbmVvolTypeData,
		PbmVvolTypeSwap,
	}
}

func (e PbmVvolType) Strings() []string {
	return types.EnumValuesAsStrings(e.Values())
}

func init() {
	types.Add("pbm:PbmVvolType", reflect.TypeOf((*PbmVvolType)(nil)).Elem())
}
