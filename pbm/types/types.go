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
	"time"

	"github.com/vmware/govmomi/vim25/types"
)

// A boxed array of `PbmCapabilityConstraintInstance`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmCapabilityConstraintInstance struct {
	PbmCapabilityConstraintInstance []PbmCapabilityConstraintInstance `xml:"PbmCapabilityConstraintInstance,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmCapabilityConstraintInstance", reflect.TypeOf((*ArrayOfPbmCapabilityConstraintInstance)(nil)).Elem())
}

// A boxed array of `PbmCapabilityInstance`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmCapabilityInstance struct {
	PbmCapabilityInstance []PbmCapabilityInstance `xml:"PbmCapabilityInstance,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmCapabilityInstance", reflect.TypeOf((*ArrayOfPbmCapabilityInstance)(nil)).Elem())
}

// A boxed array of `PbmCapabilityMetadata`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmCapabilityMetadata struct {
	PbmCapabilityMetadata []PbmCapabilityMetadata `xml:"PbmCapabilityMetadata,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmCapabilityMetadata", reflect.TypeOf((*ArrayOfPbmCapabilityMetadata)(nil)).Elem())
}

// A boxed array of `PbmCapabilityMetadataPerCategory`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmCapabilityMetadataPerCategory struct {
	PbmCapabilityMetadataPerCategory []PbmCapabilityMetadataPerCategory `xml:"PbmCapabilityMetadataPerCategory,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmCapabilityMetadataPerCategory", reflect.TypeOf((*ArrayOfPbmCapabilityMetadataPerCategory)(nil)).Elem())
}

// A boxed array of `PbmCapabilityPropertyInstance`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmCapabilityPropertyInstance struct {
	PbmCapabilityPropertyInstance []PbmCapabilityPropertyInstance `xml:"PbmCapabilityPropertyInstance,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmCapabilityPropertyInstance", reflect.TypeOf((*ArrayOfPbmCapabilityPropertyInstance)(nil)).Elem())
}

// A boxed array of `PbmCapabilityPropertyMetadata`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmCapabilityPropertyMetadata struct {
	PbmCapabilityPropertyMetadata []PbmCapabilityPropertyMetadata `xml:"PbmCapabilityPropertyMetadata,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmCapabilityPropertyMetadata", reflect.TypeOf((*ArrayOfPbmCapabilityPropertyMetadata)(nil)).Elem())
}

// A boxed array of `PbmCapabilitySchema`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmCapabilitySchema struct {
	PbmCapabilitySchema []PbmCapabilitySchema `xml:"PbmCapabilitySchema,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmCapabilitySchema", reflect.TypeOf((*ArrayOfPbmCapabilitySchema)(nil)).Elem())
}

// A boxed array of `PbmCapabilitySubProfile`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmCapabilitySubProfile struct {
	PbmCapabilitySubProfile []PbmCapabilitySubProfile `xml:"PbmCapabilitySubProfile,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmCapabilitySubProfile", reflect.TypeOf((*ArrayOfPbmCapabilitySubProfile)(nil)).Elem())
}

// A boxed array of `PbmCapabilityVendorNamespaceInfo`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmCapabilityVendorNamespaceInfo struct {
	PbmCapabilityVendorNamespaceInfo []PbmCapabilityVendorNamespaceInfo `xml:"PbmCapabilityVendorNamespaceInfo,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmCapabilityVendorNamespaceInfo", reflect.TypeOf((*ArrayOfPbmCapabilityVendorNamespaceInfo)(nil)).Elem())
}

// A boxed array of `PbmCapabilityVendorResourceTypeInfo`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmCapabilityVendorResourceTypeInfo struct {
	PbmCapabilityVendorResourceTypeInfo []PbmCapabilityVendorResourceTypeInfo `xml:"PbmCapabilityVendorResourceTypeInfo,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmCapabilityVendorResourceTypeInfo", reflect.TypeOf((*ArrayOfPbmCapabilityVendorResourceTypeInfo)(nil)).Elem())
}

// A boxed array of `PbmCompliancePolicyStatus`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmCompliancePolicyStatus struct {
	PbmCompliancePolicyStatus []PbmCompliancePolicyStatus `xml:"PbmCompliancePolicyStatus,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmCompliancePolicyStatus", reflect.TypeOf((*ArrayOfPbmCompliancePolicyStatus)(nil)).Elem())
}

// A boxed array of `PbmComplianceResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmComplianceResult struct {
	PbmComplianceResult []PbmComplianceResult `xml:"PbmComplianceResult,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmComplianceResult", reflect.TypeOf((*ArrayOfPbmComplianceResult)(nil)).Elem())
}

// A boxed array of `PbmDatastoreSpaceStatistics`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmDatastoreSpaceStatistics struct {
	PbmDatastoreSpaceStatistics []PbmDatastoreSpaceStatistics `xml:"PbmDatastoreSpaceStatistics,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmDatastoreSpaceStatistics", reflect.TypeOf((*ArrayOfPbmDatastoreSpaceStatistics)(nil)).Elem())
}

// A boxed array of `PbmDefaultProfileInfo`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmDefaultProfileInfo struct {
	PbmDefaultProfileInfo []PbmDefaultProfileInfo `xml:"PbmDefaultProfileInfo,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmDefaultProfileInfo", reflect.TypeOf((*ArrayOfPbmDefaultProfileInfo)(nil)).Elem())
}

// A boxed array of `PbmFaultNoPermissionEntityPrivileges`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmFaultNoPermissionEntityPrivileges struct {
	PbmFaultNoPermissionEntityPrivileges []PbmFaultNoPermissionEntityPrivileges `xml:"PbmFaultNoPermissionEntityPrivileges,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmFaultNoPermissionEntityPrivileges", reflect.TypeOf((*ArrayOfPbmFaultNoPermissionEntityPrivileges)(nil)).Elem())
}

// A boxed array of `PbmLoggingConfiguration`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmLoggingConfiguration struct {
	PbmLoggingConfiguration []PbmLoggingConfiguration `xml:"PbmLoggingConfiguration,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmLoggingConfiguration", reflect.TypeOf((*ArrayOfPbmLoggingConfiguration)(nil)).Elem())
}

// A boxed array of `PbmPlacementCompatibilityResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmPlacementCompatibilityResult struct {
	PbmPlacementCompatibilityResult []PbmPlacementCompatibilityResult `xml:"PbmPlacementCompatibilityResult,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmPlacementCompatibilityResult", reflect.TypeOf((*ArrayOfPbmPlacementCompatibilityResult)(nil)).Elem())
}

// A boxed array of `PbmPlacementHub`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmPlacementHub struct {
	PbmPlacementHub []PbmPlacementHub `xml:"PbmPlacementHub,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmPlacementHub", reflect.TypeOf((*ArrayOfPbmPlacementHub)(nil)).Elem())
}

// A boxed array of `PbmPlacementMatchingResources`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmPlacementMatchingResources struct {
	PbmPlacementMatchingResources []BasePbmPlacementMatchingResources `xml:"PbmPlacementMatchingResources,omitempty,typeattr" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmPlacementMatchingResources", reflect.TypeOf((*ArrayOfPbmPlacementMatchingResources)(nil)).Elem())
}

// A boxed array of `PbmPlacementRequirement`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmPlacementRequirement struct {
	PbmPlacementRequirement []BasePbmPlacementRequirement `xml:"PbmPlacementRequirement,omitempty,typeattr" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmPlacementRequirement", reflect.TypeOf((*ArrayOfPbmPlacementRequirement)(nil)).Elem())
}

// A boxed array of `PbmPlacementResourceUtilization`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmPlacementResourceUtilization struct {
	PbmPlacementResourceUtilization []PbmPlacementResourceUtilization `xml:"PbmPlacementResourceUtilization,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmPlacementResourceUtilization", reflect.TypeOf((*ArrayOfPbmPlacementResourceUtilization)(nil)).Elem())
}

// A boxed array of `PbmProfile`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmProfile struct {
	PbmProfile []BasePbmProfile `xml:"PbmProfile,omitempty,typeattr" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmProfile", reflect.TypeOf((*ArrayOfPbmProfile)(nil)).Elem())
}

// A boxed array of `PbmProfileId`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmProfileId struct {
	PbmProfileId []PbmProfileId `xml:"PbmProfileId,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmProfileId", reflect.TypeOf((*ArrayOfPbmProfileId)(nil)).Elem())
}

// A boxed array of `PbmProfileOperationOutcome`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmProfileOperationOutcome struct {
	PbmProfileOperationOutcome []PbmProfileOperationOutcome `xml:"PbmProfileOperationOutcome,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmProfileOperationOutcome", reflect.TypeOf((*ArrayOfPbmProfileOperationOutcome)(nil)).Elem())
}

// A boxed array of `PbmProfileResourceType`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmProfileResourceType struct {
	PbmProfileResourceType []PbmProfileResourceType `xml:"PbmProfileResourceType,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmProfileResourceType", reflect.TypeOf((*ArrayOfPbmProfileResourceType)(nil)).Elem())
}

// A boxed array of `PbmProfileType`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmProfileType struct {
	PbmProfileType []PbmProfileType `xml:"PbmProfileType,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmProfileType", reflect.TypeOf((*ArrayOfPbmProfileType)(nil)).Elem())
}

// A boxed array of `PbmQueryProfileResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmQueryProfileResult struct {
	PbmQueryProfileResult []PbmQueryProfileResult `xml:"PbmQueryProfileResult,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmQueryProfileResult", reflect.TypeOf((*ArrayOfPbmQueryProfileResult)(nil)).Elem())
}

// A boxed array of `PbmQueryReplicationGroupResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmQueryReplicationGroupResult struct {
	PbmQueryReplicationGroupResult []PbmQueryReplicationGroupResult `xml:"PbmQueryReplicationGroupResult,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmQueryReplicationGroupResult", reflect.TypeOf((*ArrayOfPbmQueryReplicationGroupResult)(nil)).Elem())
}

// A boxed array of `PbmRollupComplianceResult`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmRollupComplianceResult struct {
	PbmRollupComplianceResult []PbmRollupComplianceResult `xml:"PbmRollupComplianceResult,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmRollupComplianceResult", reflect.TypeOf((*ArrayOfPbmRollupComplianceResult)(nil)).Elem())
}

// A boxed array of `PbmServerObjectRef`. To be used in `Any` placeholders.
//
// This structure may be used only with operations rendered under `/pbm`.
type ArrayOfPbmServerObjectRef struct {
	PbmServerObjectRef []PbmServerObjectRef `xml:"PbmServerObjectRef,omitempty" json:"_value"`
}

func init() {
	types.Add("pbm:ArrayOfPbmServerObjectRef", reflect.TypeOf((*ArrayOfPbmServerObjectRef)(nil)).Elem())
}

// The `PbmAboutInfo` data object stores identifying data
// about the Storage Policy Server.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmAboutInfo struct {
	types.DynamicData

	// Name of the server.
	Name string `xml:"name" json:"name"`
	// Version number.
	Version string `xml:"version" json:"version"`
	// Globally unique identifier associated with this server instance.
	InstanceUuid string `xml:"instanceUuid" json:"instanceUuid"`
}

func init() {
	types.Add("pbm:PbmAboutInfo", reflect.TypeOf((*PbmAboutInfo)(nil)).Elem())
}

// An AlreadyExists fault is thrown when an attempt is made to add an element to
// a collection, if the element's key, name, or identifier already exists in
// that collection.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmAlreadyExists struct {
	PbmFault

	Name string `xml:"name,omitempty" json:"name,omitempty"`
}

func init() {
	types.Add("pbm:PbmAlreadyExists", reflect.TypeOf((*PbmAlreadyExists)(nil)).Elem())
}

type PbmAlreadyExistsFault PbmAlreadyExists

func init() {
	types.Add("pbm:PbmAlreadyExistsFault", reflect.TypeOf((*PbmAlreadyExistsFault)(nil)).Elem())
}

type PbmAssignDefaultRequirementProfile PbmAssignDefaultRequirementProfileRequestType

func init() {
	types.Add("pbm:PbmAssignDefaultRequirementProfile", reflect.TypeOf((*PbmAssignDefaultRequirementProfile)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmAssignDefaultRequirementProfile`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmAssignDefaultRequirementProfileRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The profile that needs to be made default profile.
	Profile PbmProfileId `xml:"profile" json:"profile"`
	// The datastores for which the profile needs to be made as default profile.
	Datastores []PbmPlacementHub `xml:"datastores" json:"datastores"`
}

func init() {
	types.Add("pbm:PbmAssignDefaultRequirementProfileRequestType", reflect.TypeOf((*PbmAssignDefaultRequirementProfileRequestType)(nil)).Elem())
}

type PbmAssignDefaultRequirementProfileResponse struct {
}

// Constraints on the properties for a single occurrence of a capability.
//
// All properties must satisfy their respective constraints to be compliant.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityConstraintInstance struct {
	types.DynamicData

	// Property instance array for this constraint
	PropertyInstance []PbmCapabilityPropertyInstance `xml:"propertyInstance" json:"propertyInstance"`
}

func init() {
	types.Add("pbm:PbmCapabilityConstraintInstance", reflect.TypeOf((*PbmCapabilityConstraintInstance)(nil)).Elem())
}

// The `PbmCapabilityConstraints` data object is the base
// object for capability subprofile constraints.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityConstraints struct {
	types.DynamicData
}

func init() {
	types.Add("pbm:PbmCapabilityConstraints", reflect.TypeOf((*PbmCapabilityConstraints)(nil)).Elem())
}

// A property value with description.
//
// It can be repeated under DiscreteSet.
// E.g., set of tags, each with description and tag name.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityDescription struct {
	types.DynamicData

	// Description of the property value
	Description PbmExtendedElementDescription `xml:"description" json:"description"`
	// Values for the set.
	//
	// must be one of the supported datatypes as
	// defined in `PbmBuiltinType_enum`
	// Must only contain unique values to comply with the Set semantics
	Value types.AnyType `xml:"value,typeattr" json:"value"`
}

func init() {
	types.Add("pbm:PbmCapabilityDescription", reflect.TypeOf((*PbmCapabilityDescription)(nil)).Elem())
}

// The `PbmCapabilityDiscreteSet` data object defines a set of values
// for storage profile property instances (`PbmCapabilityPropertyInstance`).
//
// Use the discrete set type to define a set of values of a supported builtin type
// (`PbmBuiltinType_enum`), for example a set of integers
// (XSD\_INT) or a set of unsigned long values (XSD\_LONG).
// See `PbmBuiltinGenericType_enum*.*VMW_SET`.
//
// A discrete set of values is declared as an array of <code>xsd:anyType</code> values.
//   - When you define a property instance for a storage profile requirement
//     and pass an array of values to the Server, you must set the array elements
//     to values of the appropriate datatype.
//   - When you read a discrete set from a property instance for a storage profile
//     capability, you must cast the <code>xsd:anyType</code> array element values
//     to the appropriate datatype.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityDiscreteSet struct {
	types.DynamicData

	// Array of values for the set.
	//
	// The values must be one of the supported datatypes
	// as defined in `PbmBuiltinType_enum` or `PbmBuiltinGenericType_enum`.
	Values []types.AnyType `xml:"values,typeattr" json:"values"`
}

func init() {
	types.Add("pbm:PbmCapabilityDiscreteSet", reflect.TypeOf((*PbmCapabilityDiscreteSet)(nil)).Elem())
}

// Generic type definition for capabilities.
//
// Indicates how a collection of values of a specific datatype
// (`PbmCapabilityTypeInfo.typeName`)
// will be interpreted.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityGenericTypeInfo struct {
	PbmCapabilityTypeInfo

	// Name of the generic type.
	//
	// Must correspond to one of the values defined in
	// `PbmBuiltinGenericType_enum`.
	GenericTypeName string `xml:"genericTypeName" json:"genericTypeName"`
}

func init() {
	types.Add("pbm:PbmCapabilityGenericTypeInfo", reflect.TypeOf((*PbmCapabilityGenericTypeInfo)(nil)).Elem())
}

// The `PbmCapabilityInstance` data object defines a storage capability instance.
//
// Metadata for the capability is described in `PbmCapabilityMetadata`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityInstance struct {
	types.DynamicData

	// Identifier for the capability.
	//
	// The identifier value corresponds to
	// `PbmCapabilityMetadata*.*PbmCapabilityMetadata.id`.
	Id PbmCapabilityMetadataUniqueId `xml:"id" json:"id"`
	// Constraints on the properties that comprise this capability.
	//
	// Each entry represents a constraint on one or more of the properties that
	// constitute this capability. A datum must meet one of the
	// constraints to be compliant.
	Constraint []PbmCapabilityConstraintInstance `xml:"constraint" json:"constraint"`
}

func init() {
	types.Add("pbm:PbmCapabilityInstance", reflect.TypeOf((*PbmCapabilityInstance)(nil)).Elem())
}

// Metadata for a single unique setting defined by a provider.
//
// A simple setting is a setting with one property.
// A complex setting contains more than one property.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityMetadata struct {
	types.DynamicData

	// Unique identifier for the capability.
	Id PbmCapabilityMetadataUniqueId `xml:"id" json:"id"`
	// Capability name and description
	Summary PbmExtendedElementDescription `xml:"summary" json:"summary"`
	// Indicates whether incorporating given capability is mandatory during creation of
	// profile.
	Mandatory *bool `xml:"mandatory" json:"mandatory,omitempty"`
	// The flag hint dictates the interpretation of constraints specified for this capability
	// in a storage policy profile.
	//
	// If hint is false, then constraints will affect placement.
	// If hint is true, constraints will not affect placement,
	// but will still be passed to provisioning operations if the provider understands the
	// relevant namespace. Optional property, false if not set.
	Hint *bool `xml:"hint" json:"hint,omitempty"`
	// Property Id of the key property, if this capability represents a key
	// value pair.
	//
	// Value is empty string if not set.
	KeyId string `xml:"keyId,omitempty" json:"keyId,omitempty"`
	// Flag to indicate if multiple constraints are allowed in the capability
	// instance.
	//
	// False if not set.
	AllowMultipleConstraints *bool `xml:"allowMultipleConstraints" json:"allowMultipleConstraints,omitempty"`
	// Metadata for the properties that comprise this capability.
	PropertyMetadata []PbmCapabilityPropertyMetadata `xml:"propertyMetadata" json:"propertyMetadata"`
}

func init() {
	types.Add("pbm:PbmCapabilityMetadata", reflect.TypeOf((*PbmCapabilityMetadata)(nil)).Elem())
}

// The `PbmCapabilityMetadataPerCategory`
// data object defines capability metadata for a profile subcategory.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityMetadataPerCategory struct {
	types.DynamicData

	// Profile subcategory to which the capability metadata belongs.
	//
	// The subcategory is specified by the storage provider.
	SubCategory string `xml:"subCategory" json:"subCategory"`
	// Capability metadata for this category
	CapabilityMetadata []PbmCapabilityMetadata `xml:"capabilityMetadata" json:"capabilityMetadata"`
}

func init() {
	types.Add("pbm:PbmCapabilityMetadataPerCategory", reflect.TypeOf((*PbmCapabilityMetadataPerCategory)(nil)).Elem())
}

// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityMetadataUniqueId struct {
	types.DynamicData

	// Namespace to which this capability belongs.
	//
	// Must be the same as
	// { @link CapabilityObjectSchema#namespace } defined for this
	// capability
	Namespace string `xml:"namespace" json:"namespace"`
	// unique identifier for this capability within given namespace
	Id string `xml:"id" json:"id"`
}

func init() {
	types.Add("pbm:PbmCapabilityMetadataUniqueId", reflect.TypeOf((*PbmCapabilityMetadataUniqueId)(nil)).Elem())
}

// Name space information for the capability metadata schema.
//
// NOTE: Name spaces are required to be globally unique across resource types.
// A same vendor can register multiple name spaces for same resource type or
// for different resource type, but the schema namespace URL must be unique
// for each of these cases.
// A CapabilityMetadata object is uniquely identified based on the namespace
// it belongs to and it's unique identifier within that namespace.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityNamespaceInfo struct {
	types.DynamicData

	// Schema version
	Version string `xml:"version" json:"version"`
	// Schema namespace.
	Namespace string                         `xml:"namespace" json:"namespace"`
	Info      *PbmExtendedElementDescription `xml:"info,omitempty" json:"info,omitempty"`
}

func init() {
	types.Add("pbm:PbmCapabilityNamespaceInfo", reflect.TypeOf((*PbmCapabilityNamespaceInfo)(nil)).Elem())
}

// The `PbmCapabilityProfile` data object defines
// capability-based profiles.
//
// A capability-based profile is derived
// from tag-based storage capabilities or from vSAN storage capabilities.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityProfile struct {
	PbmProfile

	// Indicates whether the profile is requirement
	// profile, a resource profile or a data service profile.
	//
	// The <code>profileCategory</code>
	// is a string value that corresponds to one of the
	// `PbmProfileCategoryEnum_enum` values.
	//   - REQUIREMENT profile - Defines the storage constraints applied
	//     to virtual machine placement. Requirements are defined by
	//     the user and can be associated with virtual machines and virtual
	//     disks. During provisioning, you can use a requirements profile
	//     for compliance and placement checking to support
	//     selection and configuration of resources.
	//   - RESOURCE profile - Specifies system-defined storage capabilities.
	//     You cannot modify a resource profile. You cannot associate a resource
	//     profile with vSphere entities, use it during provisioning, or target
	//     entities for resource selection or configuration.
	//     This type of profile gives the user visibility into the capabilities
	//     supported by the storage provider.
	//   - DATA\_SERVICE\_POLICY - Indicates a data service policy that can
	//     be embedded into another storage policy. Policies of this type can't
	//     be assigned to Virtual Machines or Virtual Disks. This policy cannot
	//     be used for compliance checking.
	ProfileCategory string `xml:"profileCategory" json:"profileCategory"`
	// Type of the target resource to which the capability information applies.
	//
	// A fixed enum that defines resource types for which capabilities can be defined
	// see `PbmProfileResourceType`, `PbmProfileResourceTypeEnum_enum`
	ResourceType PbmProfileResourceType `xml:"resourceType" json:"resourceType"`
	// Subprofiles that describe storage requirements or storage provider capabilities,
	// depending on the profile category (REQUIREMENT or RESOURCE).
	Constraints BasePbmCapabilityConstraints `xml:"constraints,typeattr" json:"constraints"`
	// Generation ID is used to communicate the current version of the profile to VASA
	// providers.
	//
	// It is only applicable to REQUIREMENT profile types. Every time a
	// requirement profile is edited, the Server will increment the generationId. You
	// do not need to set the generationID. When an object is created (or
	// reconfigured), the Server will send the requirement profile content, profile ID and
	// the generationID to VASA provider.
	GenerationId int64 `xml:"generationId,omitempty" json:"generationId,omitempty"`
	// Deprecated since it is not supported.
	//
	// Not supported in this release.
	IsDefault bool `xml:"isDefault" json:"isDefault"`
	// Indicates the type of system pre-created default profile.
	//
	// This will be set only for system pre-created default profiles. And
	// this is not set for RESOURCE profiles.
	SystemCreatedProfileType string `xml:"systemCreatedProfileType,omitempty" json:"systemCreatedProfileType,omitempty"`
	// This property is set only for data service policy.
	//
	// Indicates the line of service
	// `PbmLineOfServiceInfoLineOfServiceEnum_enum` of the data service policy.
	LineOfService string `xml:"lineOfService,omitempty" json:"lineOfService,omitempty"`
}

func init() {
	types.Add("pbm:PbmCapabilityProfile", reflect.TypeOf((*PbmCapabilityProfile)(nil)).Elem())
}

// The `PbmCapabilityProfileCreateSpec` describes storage requirements.
//
// Use this data object to create a `PbmCapabilityProfile`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityProfileCreateSpec struct {
	types.DynamicData

	// Name of the capability based profile to be created.
	//
	// The maximum length of the name is 80 characters.
	Name string `xml:"name" json:"name"`
	// Text description associated with the profile.
	Description string `xml:"description,omitempty" json:"description,omitempty"`
	// Category specifies the type of policy to be created.
	//
	// This can be REQUIREMENT from
	// `PbmProfileCategoryEnum_enum`
	// or null when creating a storage policy. And it can be DATA\_SERVICE\_POLICY from
	// `PbmProfileCategoryEnum_enum`
	// when creating a data service policy. RESOURCE from `PbmProfileCategoryEnum_enum`
	// is not allowed as resource profile is created by the system.
	Category string `xml:"category,omitempty" json:"category,omitempty"`
	// Deprecated as of vSphere API 6.5.
	//
	// Specifies the type of resource to which the profile applies.
	//
	// The only legal value is STORAGE - deprecated.
	ResourceType PbmProfileResourceType `xml:"resourceType" json:"resourceType"`
	// Set of subprofiles that define the storage requirements.
	//
	// A subprofile corresponds to a rule set in the vSphere Web Client.
	Constraints BasePbmCapabilityConstraints `xml:"constraints,typeattr" json:"constraints"`
}

func init() {
	types.Add("pbm:PbmCapabilityProfileCreateSpec", reflect.TypeOf((*PbmCapabilityProfileCreateSpec)(nil)).Elem())
}

// Fault used when a datastore doesnt match the capability profile property instance in requirements profile.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityProfilePropertyMismatchFault struct {
	PbmPropertyMismatchFault

	// The property instance in the resource profile that does not match.
	ResourcePropertyInstance PbmCapabilityPropertyInstance `xml:"resourcePropertyInstance" json:"resourcePropertyInstance"`
}

func init() {
	types.Add("pbm:PbmCapabilityProfilePropertyMismatchFault", reflect.TypeOf((*PbmCapabilityProfilePropertyMismatchFault)(nil)).Elem())
}

type PbmCapabilityProfilePropertyMismatchFaultFault BasePbmCapabilityProfilePropertyMismatchFault

func init() {
	types.Add("pbm:PbmCapabilityProfilePropertyMismatchFaultFault", reflect.TypeOf((*PbmCapabilityProfilePropertyMismatchFaultFault)(nil)).Elem())
}

// The `PbmCapabilityProfileUpdateSpec` data object
// contains data that you use to update a storage profile.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityProfileUpdateSpec struct {
	types.DynamicData

	// Specifies a new profile name.
	Name string `xml:"name,omitempty" json:"name,omitempty"`
	// Specifies a new profile description.
	Description string `xml:"description,omitempty" json:"description,omitempty"`
	// Specifies one or more subprofiles.
	//
	// A subprofile defines one or more
	// storage requirements.
	Constraints BasePbmCapabilityConstraints `xml:"constraints,omitempty,typeattr" json:"constraints,omitempty"`
}

func init() {
	types.Add("pbm:PbmCapabilityProfileUpdateSpec", reflect.TypeOf((*PbmCapabilityProfileUpdateSpec)(nil)).Elem())
}

// The `PbmCapabilityPropertyInstance` data object describes a virtual machine
// storage requirement.
//
// A storage requirement is based on the storage capability
// described in the `PbmCapabilityPropertyMetadata` and in the
// datastore profile property instance.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityPropertyInstance struct {
	types.DynamicData

	// Requirement property identifier.
	//
	// This identifier corresponds to the
	// storage capability metadata identifier
	// (`PbmCapabilityPropertyMetadata*.*PbmCapabilityPropertyMetadata.id`).
	Id string `xml:"id" json:"id"`
	// Operator for the values.
	//
	// Currently only support NOT operator for
	// tag namespace
	// See operator definition in (`PbmCapabilityOperator_enum`).
	Operator string `xml:"operator,omitempty" json:"operator,omitempty"`
	// Property value.
	//
	// You must specify the value.
	// A property value is one value or a collection of values.
	//   - A single property value is expressed as a scalar value.
	//   - A collection of values is expressed as a `PbmCapabilityDiscreteSet`
	//     or a `PbmCapabilityRange` of values.
	//
	// The datatype of each value must be one of the
	// `PbmBuiltinType_enum` datatypes.
	// If the property consists of a collection of values,
	// the interpretation of those values is determined by the
	// `PbmCapabilityGenericTypeInfo`.
	//
	// Type information for a property instance is described in the property metadata
	// (`PbmCapabilityPropertyMetadata*.*PbmCapabilityPropertyMetadata.type`).
	Value types.AnyType `xml:"value,typeattr" json:"value"`
}

func init() {
	types.Add("pbm:PbmCapabilityPropertyInstance", reflect.TypeOf((*PbmCapabilityPropertyInstance)(nil)).Elem())
}

// The `PbmCapabilityPropertyMetadata` data object describes storage capability.
//
// An instance of property metadata may apply to many property instances
// (`PbmCapabilityPropertyInstance`).
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityPropertyMetadata struct {
	types.DynamicData

	// Property identifier.
	//
	// Should be unique within the definition of the
	// capability. Property instances refer to this identifier
	// (`PbmCapabilityPropertyInstance*.*PbmCapabilityPropertyInstance.id`).
	Id string `xml:"id" json:"id"`
	// Property name and description.
	//   - The <code>summary.label</code> property
	//     (`PbmExtendedElementDescription.label`)
	//     contains property 'name' in server locale.
	//   - The <code>summary.summary</code> property
	//     (`PbmExtendedElementDescription.summary`)
	//     contains property 'description' in server locale.
	//   - The <code>summary.messageCatalogKeyPrefix</code> property
	//     (`PbmExtendedElementDescription.messageCatalogKeyPrefix`)
	//     contains unique prefix for this property within given message catalog.
	//     Prefix format: &lt;capability\_unique\_identifier&gt;.&lt;property\_id&gt;
	//     capability\_unique\_identifier -- string representation of
	//     `PbmCapabilityMetadataUniqueId` which globally identifies given
	//     capability metadata definition uniquely.
	//     property\_id -- 'id' of this property `PbmCapabilityPropertyMetadata.id`
	//     Eg www.emc.com.storage.Recovery.Recovery\_site
	//     www.emc.com.storage.Recovery.RPO
	//     www.emc.com.storage.Recovery.RTO
	Summary PbmExtendedElementDescription `xml:"summary" json:"summary"`
	// Indicates whether incorporating given capability is mandatory during creation of
	// profile.
	Mandatory bool `xml:"mandatory" json:"mandatory"`
	// Type information for the capability.
	//
	// The type of a property value
	// (`PbmCapabilityPropertyInstance*.*PbmCapabilityPropertyInstance.value`)
	// is specified as a builtin datatype and may also specify the interpretation of a
	// collection of values of that datatype.
	//   - `PbmCapabilityPropertyMetadata.type*.*PbmCapabilityTypeInfo.typeName`
	//     specifies the `PbmBuiltinType_enum`.
	//   - `PbmCapabilityPropertyMetadata.type*.*PbmCapabilityGenericTypeInfo.genericTypeName`
	//     indicates how a collection of values of the specified datatype will be interpreted
	//     (`PbmBuiltinGenericType_enum`).
	Type BasePbmCapabilityTypeInfo `xml:"type,omitempty,typeattr" json:"type,omitempty"`
	// Default value, if any, that the property will assume when not
	// constrained by requirements.
	//
	// This object must be of the
	// `PbmCapabilityPropertyMetadata.type`
	// defined for the property.
	DefaultValue types.AnyType `xml:"defaultValue,omitempty,typeattr" json:"defaultValue,omitempty"`
	// All legal values that the property may take on, across all
	// implementations of the property.
	//
	// This definition of legal values is not
	// determined by any particular resource configuration; rather it is
	// inherent to the definition of the property. If undefined, then any value
	// of the correct type is legal. This object must be a generic container for
	// the `PbmCapabilityPropertyMetadata.type`
	// defined for the property;
	// see `PbmBuiltinGenericType_enum`
	// for the supported generic container types.
	AllowedValue types.AnyType `xml:"allowedValue,omitempty,typeattr" json:"allowedValue,omitempty"`
	// A hint for data-driven systems that assist in authoring requirements
	// constraints.
	//
	// Acceptable values defined by
	// `PbmBuiltinGenericType_enum`.
	// A property will typically only have constraints of a given type in
	// requirement profiles, even if it is likely to use constraints of
	// different types across capability profiles. This value, if specified,
	// specifies the expected kind of constraint used in requirement profiles.
	// Considerations for using this information:
	//   - This is only a hint; any properly formed constraint
	//     (see `PbmCapabilityPropertyInstance.value`)
	//     is still valid for a requirement profile.
	//   - If VMW\_SET is hinted, then a single value matching the property metadata type is
	//     also an expected form of constraint, as the latter is an allowed convenience
	//     for expressing a single-member set.
	//   - If this hint is not specified, then the authoring system may default to a form of
	//     constraint determined by its own criteria.
	RequirementsTypeHint string `xml:"requirementsTypeHint,omitempty" json:"requirementsTypeHint,omitempty"`
}

func init() {
	types.Add("pbm:PbmCapabilityPropertyMetadata", reflect.TypeOf((*PbmCapabilityPropertyMetadata)(nil)).Elem())
}

// The `PbmCapabilityRange` data object defines a range of values for storage property
// instances (`PbmCapabilityPropertyInstance`).
//
// Use the range type to define a range of values of a supported builtin type,
// for example range&lt;int&gt;, range&lt;long&gt;, or range&lt;timespan&gt;.
// You can specify a partial range by omitting one of the properties, min or max.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityRange struct {
	types.DynamicData

	// Minimum value of range.
	//
	// Must be one of the supported
	// datatypes as defined in `PbmBuiltinType_enum`.
	// Must be the same datatype as min.
	Min types.AnyType `xml:"min,typeattr" json:"min"`
	// Maximum value of range.
	//
	// Must be one of the supported
	// datatypes as defined in `PbmBuiltinType_enum`.
	// Must be the same datatype as max.
	Max types.AnyType `xml:"max,typeattr" json:"max"`
}

func init() {
	types.Add("pbm:PbmCapabilityRange", reflect.TypeOf((*PbmCapabilityRange)(nil)).Elem())
}

// Capability Schema information
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilitySchema struct {
	types.DynamicData

	VendorInfo    PbmCapabilitySchemaVendorInfo `xml:"vendorInfo" json:"vendorInfo"`
	NamespaceInfo PbmCapabilityNamespaceInfo    `xml:"namespaceInfo" json:"namespaceInfo"`
	// Service type for the schema.
	//
	// Do not use Category as each service needs to have its own schema version.
	//
	// If omitted, this schema specifies persistence capabilities.
	LineOfService BasePbmLineOfServiceInfo `xml:"lineOfService,omitempty,typeattr" json:"lineOfService,omitempty"`
	// Capability metadata organized by category
	CapabilityMetadataPerCategory []PbmCapabilityMetadataPerCategory `xml:"capabilityMetadataPerCategory" json:"capabilityMetadataPerCategory"`
}

func init() {
	types.Add("pbm:PbmCapabilitySchema", reflect.TypeOf((*PbmCapabilitySchema)(nil)).Elem())
}

// Information about vendor/owner of the capability metadata schema
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilitySchemaVendorInfo struct {
	types.DynamicData

	// Unique identifier for the vendor who owns the given capability
	// schema definition
	VendorUuid string `xml:"vendorUuid" json:"vendorUuid"`
	// Captures name and description information about the vendor/owner of
	// the schema.
	//   - The <code>summary.label</code> property
	//     (`PbmExtendedElementDescription.label`)
	//     contains vendor name information in server locale.
	//   - The <code>summary.summary</code> property
	//     (`PbmExtendedElementDescription.summary`)
	//     contains vendor description string in server locale.
	//   - The <code>summary.messageCatalogKeyPrefix</code> property
	//     (`PbmExtendedElementDescription.messageCatalogKeyPrefix`)
	//     contains unique prefix for the vendor information within given message
	//     catalog.
	Info PbmExtendedElementDescription `xml:"info" json:"info"`
}

func init() {
	types.Add("pbm:PbmCapabilitySchemaVendorInfo", reflect.TypeOf((*PbmCapabilitySchemaVendorInfo)(nil)).Elem())
}

// A `PbmCapabilitySubProfile`
// is a section within a profile that aggregates one or more capability
// instances.
//
// Capability instances define storage constraints.
//
// All constraints within a subprofile are ANDed by default.
// When you perform compliance checking on a virtual machine or virtual
// disk, all of the constraints must be satisfied by the storage capabilities.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilitySubProfile struct {
	types.DynamicData

	// Subprofile name.
	Name string `xml:"name" json:"name"`
	// List of capability instances.
	Capability []PbmCapabilityInstance `xml:"capability" json:"capability"`
	// Indicates whether the source policy profile allows creating a virtual machine
	// or virtual disk that may be non-compliant.
	ForceProvision *bool `xml:"forceProvision" json:"forceProvision,omitempty"`
}

func init() {
	types.Add("pbm:PbmCapabilitySubProfile", reflect.TypeOf((*PbmCapabilitySubProfile)(nil)).Elem())
}

// The `PbmCapabilitySubProfileConstraints` data object defines a group
// of storage subprofiles.
//
// Subprofile usage depends on the type of profile
// (`PbmCapabilityProfile*.*PbmCapabilityProfile.profileCategory`).
//   - For a REQUIREMENTS profile, each subprofile defines storage requirements.
//     A Storage Policy API requirements subprofile corresponds to a vSphere Web Client
//     rule set.
//   - For a RESOURCE profile, each subprofile defines storage capabilities.
//     Storage capabilities are read-only.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilitySubProfileConstraints struct {
	PbmCapabilityConstraints

	// Aggregation of one or more subprofiles.
	//
	// The relationship among all subprofiles is "OR". When you perform
	// compliance checking on a profile that contains more than one subprofile,
	// a non-compliant result for any one of the subprofiles will produce a
	// non-compliant result for the operation.
	SubProfiles []PbmCapabilitySubProfile `xml:"subProfiles" json:"subProfiles"`
}

func init() {
	types.Add("pbm:PbmCapabilitySubProfileConstraints", reflect.TypeOf((*PbmCapabilitySubProfileConstraints)(nil)).Elem())
}

// The `PbmCapabilityTimeSpan` data object defines a time value and time unit,
// for example 10 hours or 5 minutes.
//
// See
// `PbmBuiltinType_enum*.*VMW_TIMESPAN`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityTimeSpan struct {
	types.DynamicData

	// Time value.
	//
	// Must be a positive integer.
	Value int32 `xml:"value" json:"value"`
	// Unit value for time.
	//
	// The string value must correspond
	// to one of the `PbmCapabilityTimeUnitType_enum` values.
	Unit string `xml:"unit" json:"unit"`
}

func init() {
	types.Add("pbm:PbmCapabilityTimeSpan", reflect.TypeOf((*PbmCapabilityTimeSpan)(nil)).Elem())
}

// The `PbmCapabilityTypeInfo` data object defines the datatype for a requirement
// or capability property.
//
// See `PbmCapabilityPropertyMetadata`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityTypeInfo struct {
	types.DynamicData

	// Datatype for a property.
	//
	// Must be one of the types defined
	// in `PbmBuiltinType_enum`.
	//
	// A property value might consist of a collection of values of the specified
	// datatype. The interpretation of the collection is determined by the
	// generic type (`PbmCapabilityGenericTypeInfo.genericTypeName`).
	// The generic type indicates how a collection of values
	// of the specified datatype will be interpreted. See the descriptions of the
	// `PbmBuiltinType_enum` definitions.
	TypeName string `xml:"typeName" json:"typeName"`
}

func init() {
	types.Add("pbm:PbmCapabilityTypeInfo", reflect.TypeOf((*PbmCapabilityTypeInfo)(nil)).Elem())
}

// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityVendorNamespaceInfo struct {
	types.DynamicData

	VendorInfo    PbmCapabilitySchemaVendorInfo `xml:"vendorInfo" json:"vendorInfo"`
	NamespaceInfo PbmCapabilityNamespaceInfo    `xml:"namespaceInfo" json:"namespaceInfo"`
}

func init() {
	types.Add("pbm:PbmCapabilityVendorNamespaceInfo", reflect.TypeOf((*PbmCapabilityVendorNamespaceInfo)(nil)).Elem())
}

// This structure may be used only with operations rendered under `/pbm`.
type PbmCapabilityVendorResourceTypeInfo struct {
	types.DynamicData

	// Resource type for which given vendor has registered given namespace
	// along with capability metadata that belongs to the namespace.
	//
	// Must match one of the values for enum `PbmProfileResourceTypeEnum_enum`
	ResourceType string `xml:"resourceType" json:"resourceType"`
	// List of all vendorInfo &lt;--&gt; namespaceInfo tuples that are registered for
	// given resource type
	VendorNamespaceInfo []PbmCapabilityVendorNamespaceInfo `xml:"vendorNamespaceInfo" json:"vendorNamespaceInfo"`
}

func init() {
	types.Add("pbm:PbmCapabilityVendorResourceTypeInfo", reflect.TypeOf((*PbmCapabilityVendorResourceTypeInfo)(nil)).Elem())
}

type PbmCheckCompatibility PbmCheckCompatibilityRequestType

func init() {
	types.Add("pbm:PbmCheckCompatibility", reflect.TypeOf((*PbmCheckCompatibility)(nil)).Elem())
}

// The parameters of `PbmPlacementSolver.PbmCheckCompatibility`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCheckCompatibilityRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Candidate list of hubs, either datastores or storage pods or a
	// mix. If this parameter is not specified, the Server uses all
	// of the datastores and storage pods for placement compatibility
	// checking.
	HubsToSearch []PbmPlacementHub `xml:"hubsToSearch,omitempty" json:"hubsToSearch,omitempty"`
	// Storage requirement profile.
	Profile PbmProfileId `xml:"profile" json:"profile"`
}

func init() {
	types.Add("pbm:PbmCheckCompatibilityRequestType", reflect.TypeOf((*PbmCheckCompatibilityRequestType)(nil)).Elem())
}

type PbmCheckCompatibilityResponse struct {
	Returnval []PbmPlacementCompatibilityResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmCheckCompatibilityWithSpec PbmCheckCompatibilityWithSpecRequestType

func init() {
	types.Add("pbm:PbmCheckCompatibilityWithSpec", reflect.TypeOf((*PbmCheckCompatibilityWithSpec)(nil)).Elem())
}

// The parameters of `PbmPlacementSolver.PbmCheckCompatibilityWithSpec`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCheckCompatibilityWithSpecRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Candidate list of hubs, either datastores or storage pods
	// or a mix. If this parameter is not specified, the Server uses all of the
	// datastores and storage pods for placement compatibility checking.
	HubsToSearch []PbmPlacementHub `xml:"hubsToSearch,omitempty" json:"hubsToSearch,omitempty"`
	// Specification for a capability based profile.
	ProfileSpec PbmCapabilityProfileCreateSpec `xml:"profileSpec" json:"profileSpec"`
}

func init() {
	types.Add("pbm:PbmCheckCompatibilityWithSpecRequestType", reflect.TypeOf((*PbmCheckCompatibilityWithSpecRequestType)(nil)).Elem())
}

type PbmCheckCompatibilityWithSpecResponse struct {
	Returnval []PbmPlacementCompatibilityResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmCheckCompliance PbmCheckComplianceRequestType

func init() {
	types.Add("pbm:PbmCheckCompliance", reflect.TypeOf((*PbmCheckCompliance)(nil)).Elem())
}

// The parameters of `PbmComplianceManager.PbmCheckCompliance`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCheckComplianceRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// One or more references to storage entities.
	// You can specify virtual machines and virtual disks
	// A maximum of 1000 virtual machines and/or virtual disks can be specified
	// in a call. The results of calling the checkCompliance API with
	// more than a 1000 entities is undefined.
	//   - If the list of entities also contains datastores, the Server
	//     will ignore the datastores.
	//   - If the list contains valid and invalid entities, the Server ignores
	//     the invalid entities and returns results for the valid entities.
	//     Invalid entities are entities that are not in the vCenter inventory.
	//   - If the list contains only datastores, the method throws
	//     an <code>InvalidArgument</code> fault.
	//   - If the list contains virtual machines and disks and the entities
	//     are invalid or have been deleted by the time of the request, the method
	//     throws an <code>InvalidArgument</code> fault.
	//
	// If an entity does not have an associated storage profile, the entity
	// is removed from the list.
	Entities []PbmServerObjectRef `xml:"entities" json:"entities"`
	// Not used. If specified, the Server ignores the value.
	// The Server uses the profiles associated with the specified entities.
	Profile *PbmProfileId `xml:"profile,omitempty" json:"profile,omitempty"`
}

func init() {
	types.Add("pbm:PbmCheckComplianceRequestType", reflect.TypeOf((*PbmCheckComplianceRequestType)(nil)).Elem())
}

type PbmCheckComplianceResponse struct {
	Returnval []PbmComplianceResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmCheckRequirements PbmCheckRequirementsRequestType

func init() {
	types.Add("pbm:PbmCheckRequirements", reflect.TypeOf((*PbmCheckRequirements)(nil)).Elem())
}

// The parameters of `PbmPlacementSolver.PbmCheckRequirements`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCheckRequirementsRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Candidate list of hubs, either datastores or storage pods
	// or a mix. If this parameter is not specified, the Server uses all of the
	// datastores and storage pods for placement compatibility checking.
	HubsToSearch []PbmPlacementHub `xml:"hubsToSearch,omitempty" json:"hubsToSearch,omitempty"`
	// reference to the object being placed. Should be null when a new
	// object is being provisioned. Should be specified when placement compatibility is being checked
	// for an existing object. Supported objects are
	// `virtualMachine`,
	// `virtualMachineAndDisks`,
	// `virtualDiskId`,
	// `virtualDiskUUID`
	PlacementSubjectRef *PbmServerObjectRef `xml:"placementSubjectRef,omitempty" json:"placementSubjectRef,omitempty"`
	// Requirements including the policy requirements, compute
	// requirements and capacity requirements. It is invalid to specify no requirements. It is also
	// invalid to specify duplicate requirements or multiple conflicting requirements such as
	// specifying both `PbmPlacementCapabilityConstraintsRequirement` and
	// `PbmPlacementCapabilityProfileRequirement`.
	PlacementSubjectRequirement []BasePbmPlacementRequirement `xml:"placementSubjectRequirement,omitempty,typeattr" json:"placementSubjectRequirement,omitempty"`
}

func init() {
	types.Add("pbm:PbmCheckRequirementsRequestType", reflect.TypeOf((*PbmCheckRequirementsRequestType)(nil)).Elem())
}

type PbmCheckRequirementsResponse struct {
	Returnval []PbmPlacementCompatibilityResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmCheckRollupCompliance PbmCheckRollupComplianceRequestType

func init() {
	types.Add("pbm:PbmCheckRollupCompliance", reflect.TypeOf((*PbmCheckRollupCompliance)(nil)).Elem())
}

// The parameters of `PbmComplianceManager.PbmCheckRollupCompliance`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCheckRollupComplianceRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// One or more references to virtual machines.
	// A maximum of 1000 virtual machines can be specified
	// in a call. The results of calling the checkRollupCompliance API with
	// more than a 1000 entities is undefined.
	Entity []PbmServerObjectRef `xml:"entity" json:"entity"`
}

func init() {
	types.Add("pbm:PbmCheckRollupComplianceRequestType", reflect.TypeOf((*PbmCheckRollupComplianceRequestType)(nil)).Elem())
}

type PbmCheckRollupComplianceResponse struct {
	Returnval []PbmRollupComplianceResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

// Super class for all compatibility check faults.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCompatibilityCheckFault struct {
	PbmFault

	// Placement Hub
	Hub PbmPlacementHub `xml:"hub" json:"hub"`
}

func init() {
	types.Add("pbm:PbmCompatibilityCheckFault", reflect.TypeOf((*PbmCompatibilityCheckFault)(nil)).Elem())
}

type PbmCompatibilityCheckFaultFault BasePbmCompatibilityCheckFault

func init() {
	types.Add("pbm:PbmCompatibilityCheckFaultFault", reflect.TypeOf((*PbmCompatibilityCheckFaultFault)(nil)).Elem())
}

// Additional information on the effects of backend resources and
// operations on the storage object.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmComplianceOperationalStatus struct {
	types.DynamicData

	// Whether the object is currently affected by the failure of backend
	// storage resources.
	//
	// Optional property.
	Healthy *bool `xml:"healthy" json:"healthy,omitempty"`
	// Estimated completion time of a backend operation affecting the object.
	//
	// If set, then "transitional" will be true.
	// Optional property.
	OperationETA *time.Time `xml:"operationETA" json:"operationETA,omitempty"`
	// Percent progress of a backend operation affecting the object.
	//
	// If set, then "transitional" will be true.
	// Optional property.
	OperationProgress int64 `xml:"operationProgress,omitempty" json:"operationProgress,omitempty"`
	// Whether an object is undergoing a backend operation that may affect
	// its performance.
	//
	// This may be a rebalancing the resources of a healthy
	// object or recovery tasks for an unhealthy object.
	// Optional property.
	Transitional *bool `xml:"transitional" json:"transitional,omitempty"`
}

func init() {
	types.Add("pbm:PbmComplianceOperationalStatus", reflect.TypeOf((*PbmComplianceOperationalStatus)(nil)).Elem())
}

// The `PbmCompliancePolicyStatus` data object provides information
// when compliance checking produces non-compliant results.
//
// See
// `PbmComplianceResult*.*PbmComplianceResult.violatedPolicies`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCompliancePolicyStatus struct {
	types.DynamicData

	// Expected storage capability values of profile policies defined
	// by a storage provider.
	ExpectedValue PbmCapabilityInstance `xml:"expectedValue" json:"expectedValue"`
	// Current storage requirement values of the profile policies
	// specified for the virtual machine or virtual disk.
	CurrentValue *PbmCapabilityInstance `xml:"currentValue,omitempty" json:"currentValue,omitempty"`
}

func init() {
	types.Add("pbm:PbmCompliancePolicyStatus", reflect.TypeOf((*PbmCompliancePolicyStatus)(nil)).Elem())
}

// The `PbmComplianceResult` data object describes the results of profile compliance
// checking for a virtual machine or virtual disk.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmComplianceResult struct {
	types.DynamicData

	// Time when the compliance was checked.
	CheckTime time.Time `xml:"checkTime" json:"checkTime"`
	// Virtual machine or virtual disk for which compliance was checked.
	Entity PbmServerObjectRef `xml:"entity" json:"entity"`
	// Requirement profile with which the compliance was checked.
	Profile *PbmProfileId `xml:"profile,omitempty" json:"profile,omitempty"`
	// Status of the current running compliance operation.
	//
	// If there is no
	// compliance check operation triggered, this indicates the last compliance
	// task status. <code>complianceTaskStatus</code> is a string value that
	// corresponds to one of the
	// `PbmComplianceResultComplianceTaskStatus_enum` values.
	ComplianceTaskStatus string `xml:"complianceTaskStatus,omitempty" json:"complianceTaskStatus,omitempty"`
	// Status of the compliance operation.
	//
	// <code>complianceStatus</code> is a
	// string value that corresponds to one of the
	// `PbmComplianceStatus_enum` values.
	//
	// When you perform compliance checking on an entity whose associated profile
	// contains more than one subprofile (
	// `PbmCapabilityProfile` .
	// `PbmCapabilityProfile.constraints`), a compliant
	// result for any one of the subprofiles will produce a compliant result
	// for the operation.
	ComplianceStatus string `xml:"complianceStatus" json:"complianceStatus"`
	// Deprecated as of vSphere 2016, use
	// `PbmComplianceStatus_enum` to
	// know if a mismatch has occurred. If
	// `PbmComplianceResult.complianceStatus` value
	// is outOfDate, mismatch has occurred.
	//
	// Set to true if there is a profile version mismatch between the Storage
	// Profile Server and the storage provider.
	//
	// If you receive a result that
	// indicates a mismatch, you must use the vSphere API to update the profile
	// associated with the virtual machine or virtual disk.
	Mismatch bool `xml:"mismatch" json:"mismatch"`
	// Values for capabilities that are known to be non-compliant with the specified constraints.
	ViolatedPolicies []PbmCompliancePolicyStatus `xml:"violatedPolicies,omitempty" json:"violatedPolicies,omitempty"`
	// This property is set if the compliance task fails with errors.
	//
	// There can be
	// more than one error since a policy containing multiple blobs can return
	// multiple failures, one for each blob.
	ErrorCause []types.LocalizedMethodFault `xml:"errorCause,omitempty" json:"errorCause,omitempty"`
	// Additional information on the effects of backend resources and
	// operations on the storage object.
	OperationalStatus *PbmComplianceOperationalStatus `xml:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	// Informational localized messages provided by the VASA provider in
	// addition to the <code>violatedPolicy</code>.
	Info *PbmExtendedElementDescription `xml:"info,omitempty" json:"info,omitempty"`
}

func init() {
	types.Add("pbm:PbmComplianceResult", reflect.TypeOf((*PbmComplianceResult)(nil)).Elem())
}

type PbmCreate PbmCreateRequestType

func init() {
	types.Add("pbm:PbmCreate", reflect.TypeOf((*PbmCreate)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmCreate`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmCreateRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Capability-based profile specification.
	CreateSpec PbmCapabilityProfileCreateSpec `xml:"createSpec" json:"createSpec"`
}

func init() {
	types.Add("pbm:PbmCreateRequestType", reflect.TypeOf((*PbmCreateRequestType)(nil)).Elem())
}

type PbmCreateResponse struct {
	Returnval PbmProfileId `xml:"returnval" json:"returnval"`
}

// DataServiceToProfilesMap maps the data service policy to the parent storage policies
// if referred.
//
// This is returned from the API call
// `ProfileManager#queryParentStoragePolicies(ProfileId[])`
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmDataServiceToPoliciesMap struct {
	types.DynamicData

	// Denotes a Data Service Policy Id.
	DataServicePolicy PbmProfileId `xml:"dataServicePolicy" json:"dataServicePolicy"`
	// Storage Policies that refer to the Data Service Policy given by
	// `PbmDataServiceToPoliciesMap.dataServicePolicy`.
	ParentStoragePolicies []PbmProfileId `xml:"parentStoragePolicies,omitempty" json:"parentStoragePolicies,omitempty"`
	// The fault is set in case of error conditions and this property will
	// have the reason.
	Fault *types.LocalizedMethodFault `xml:"fault,omitempty" json:"fault,omitempty"`
}

func init() {
	types.Add("pbm:PbmDataServiceToPoliciesMap", reflect.TypeOf((*PbmDataServiceToPoliciesMap)(nil)).Elem())
}

// Space stats for datastore
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmDatastoreSpaceStatistics struct {
	types.DynamicData

	// Capability profile id.
	//
	// It is null when the statistics are for the entire
	// datastore.
	ProfileId string `xml:"profileId,omitempty" json:"profileId,omitempty"`
	// Total physical space in MB.
	PhysicalTotalInMB int64 `xml:"physicalTotalInMB" json:"physicalTotalInMB"`
	// Total physical free space in MB.
	PhysicalFreeInMB int64 `xml:"physicalFreeInMB" json:"physicalFreeInMB"`
	// Used physical storage space in MB.
	PhysicalUsedInMB int64 `xml:"physicalUsedInMB" json:"physicalUsedInMB"`
	// Logical space limit set by the storage admin in MB.
	//
	// Omitted if there is no Logical space limit.
	LogicalLimitInMB int64 `xml:"logicalLimitInMB,omitempty" json:"logicalLimitInMB,omitempty"`
	// Free logical storage space in MB.
	LogicalFreeInMB int64 `xml:"logicalFreeInMB" json:"logicalFreeInMB"`
	// Used logical storage space in MB.
	LogicalUsedInMB int64 `xml:"logicalUsedInMB" json:"logicalUsedInMB"`
}

func init() {
	types.Add("pbm:PbmDatastoreSpaceStatistics", reflect.TypeOf((*PbmDatastoreSpaceStatistics)(nil)).Elem())
}

// Not supported in this release.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmDefaultCapabilityProfile struct {
	PbmCapabilityProfile

	// Not supported in this release.
	VvolType []string `xml:"vvolType" json:"vvolType"`
	// Not supported in this release.
	ContainerId string `xml:"containerId" json:"containerId"`
}

func init() {
	types.Add("pbm:PbmDefaultCapabilityProfile", reflect.TypeOf((*PbmDefaultCapabilityProfile)(nil)).Elem())
}

// Warning fault used to indicate that the vendor specific datastore matches the tag in the
// requirements profile that does not have a vendor specific rule set.
//
// In such case,
// an empty blob is sent to the vendor specific datastore and the default profile would apply.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmDefaultProfileAppliesFault struct {
	PbmCompatibilityCheckFault
}

func init() {
	types.Add("pbm:PbmDefaultProfileAppliesFault", reflect.TypeOf((*PbmDefaultProfileAppliesFault)(nil)).Elem())
}

type PbmDefaultProfileAppliesFaultFault PbmDefaultProfileAppliesFault

func init() {
	types.Add("pbm:PbmDefaultProfileAppliesFaultFault", reflect.TypeOf((*PbmDefaultProfileAppliesFaultFault)(nil)).Elem())
}

// Data structure that stores the default profile for datastores.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmDefaultProfileInfo struct {
	types.DynamicData

	// Datastores
	Datastores []PbmPlacementHub `xml:"datastores" json:"datastores"`
	// Default requirements profile.
	//
	// It is set to null if the datastores are not associated with any default profile.
	DefaultProfile BasePbmProfile `xml:"defaultProfile,omitempty,typeattr" json:"defaultProfile,omitempty"`
	// NoPermission fault if default profile is not permitted.
	MethodFault *types.LocalizedMethodFault `xml:"methodFault,omitempty" json:"methodFault,omitempty"`
}

func init() {
	types.Add("pbm:PbmDefaultProfileInfo", reflect.TypeOf((*PbmDefaultProfileInfo)(nil)).Elem())
}

type PbmDelete PbmDeleteRequestType

func init() {
	types.Add("pbm:PbmDelete", reflect.TypeOf((*PbmDelete)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmDelete`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmDeleteRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Array of profile identifiers.
	ProfileId []PbmProfileId `xml:"profileId" json:"profileId"`
}

func init() {
	types.Add("pbm:PbmDeleteRequestType", reflect.TypeOf((*PbmDeleteRequestType)(nil)).Elem())
}

type PbmDeleteResponse struct {
	Returnval []PbmProfileOperationOutcome `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

// A DuplicateName exception is thrown because a name already exists
// in the same name space.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmDuplicateName struct {
	PbmFault

	// The name that is already bound in the name space.
	Name string `xml:"name" json:"name"`
}

func init() {
	types.Add("pbm:PbmDuplicateName", reflect.TypeOf((*PbmDuplicateName)(nil)).Elem())
}

type PbmDuplicateNameFault PbmDuplicateName

func init() {
	types.Add("pbm:PbmDuplicateNameFault", reflect.TypeOf((*PbmDuplicateNameFault)(nil)).Elem())
}

// This structure may be used only with operations rendered under `/pbm`.
type PbmExtendedElementDescription struct {
	types.DynamicData

	// Display label.
	Label string `xml:"label" json:"label"`
	// Summary description.
	Summary string `xml:"summary" json:"summary"`
	// Enumeration or literal ID being described.
	Key string `xml:"key" json:"key"`
	// Key to the localized message string in the catalog.
	//
	// If the localized string contains parameters, values to the
	// parameters will be provided in #messageArg.
	// E.g: If the message in the catalog is
	// "IP address is {address}", value for "address"
	// will be provided by #messageArg.
	// Both summary and label in ElementDescription will have a corresponding
	// entry in the message catalog with the keys
	// &lt;messageCatalogKeyPrefix&gt;.summary and &lt;messageCatalogKeyPrefix&gt;.label
	// respectively.
	// ElementDescription.summary and ElementDescription.label will contain
	// the strings in server locale.
	MessageCatalogKeyPrefix string `xml:"messageCatalogKeyPrefix" json:"messageCatalogKeyPrefix"`
	// Provides named arguments that can be used to localize the
	// message in the catalog.
	MessageArg []types.KeyAnyValue `xml:"messageArg,omitempty" json:"messageArg,omitempty"`
}

func init() {
	types.Add("pbm:PbmExtendedElementDescription", reflect.TypeOf((*PbmExtendedElementDescription)(nil)).Elem())
}

// The super class for all pbm faults.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmFault struct {
	types.MethodFault
}

func init() {
	types.Add("pbm:PbmFault", reflect.TypeOf((*PbmFault)(nil)).Elem())
}

type PbmFaultFault BasePbmFault

func init() {
	types.Add("pbm:PbmFaultFault", reflect.TypeOf((*PbmFaultFault)(nil)).Elem())
}

// Thrown when login fails due to token not provided or token could not be
// validated.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmFaultInvalidLogin struct {
	PbmFault
}

func init() {
	types.Add("pbm:PbmFaultInvalidLogin", reflect.TypeOf((*PbmFaultInvalidLogin)(nil)).Elem())
}

type PbmFaultInvalidLoginFault PbmFaultInvalidLogin

func init() {
	types.Add("pbm:PbmFaultInvalidLoginFault", reflect.TypeOf((*PbmFaultInvalidLoginFault)(nil)).Elem())
}

// Thrown when an operation is denied because of a privilege
// not held on a storage profile.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmFaultNoPermission struct {
	types.SecurityError

	// List of profile ids and missing privileges for each profile
	MissingPrivileges []PbmFaultNoPermissionEntityPrivileges `xml:"missingPrivileges,omitempty" json:"missingPrivileges,omitempty"`
}

func init() {
	types.Add("pbm:PbmFaultNoPermission", reflect.TypeOf((*PbmFaultNoPermission)(nil)).Elem())
}

// This structure may be used only with operations rendered under `/pbm`.
type PbmFaultNoPermissionEntityPrivileges struct {
	types.DynamicData

	ProfileId    *PbmProfileId `xml:"profileId,omitempty" json:"profileId,omitempty"`
	PrivilegeIds []string      `xml:"privilegeIds,omitempty" json:"privilegeIds,omitempty"`
}

func init() {
	types.Add("pbm:PbmFaultNoPermissionEntityPrivileges", reflect.TypeOf((*PbmFaultNoPermissionEntityPrivileges)(nil)).Elem())
}

type PbmFaultNoPermissionFault PbmFaultNoPermission

func init() {
	types.Add("pbm:PbmFaultNoPermissionFault", reflect.TypeOf((*PbmFaultNoPermissionFault)(nil)).Elem())
}

// A NotFound error occurs when a referenced component of a managed
// object cannot be found.
//
// The referenced component can be a data
// object type (such as a role or permission) or a primitive
// (such as a string).
//
// For example, if the missing referenced component is a data object, such as
// VirtualSwitch, the NotFound error is
// thrown. The NotFound error is also thrown if the data object is found, but the referenced name
// (for example, "vswitch0") is not.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmFaultNotFound struct {
	PbmFault
}

func init() {
	types.Add("pbm:PbmFaultNotFound", reflect.TypeOf((*PbmFaultNotFound)(nil)).Elem())
}

type PbmFaultNotFoundFault PbmFaultNotFound

func init() {
	types.Add("pbm:PbmFaultNotFoundFault", reflect.TypeOf((*PbmFaultNotFoundFault)(nil)).Elem())
}

// This structure may be used only with operations rendered under `/pbm`.
type PbmFaultProfileStorageFault struct {
	PbmFault
}

func init() {
	types.Add("pbm:PbmFaultProfileStorageFault", reflect.TypeOf((*PbmFaultProfileStorageFault)(nil)).Elem())
}

type PbmFaultProfileStorageFaultFault PbmFaultProfileStorageFault

func init() {
	types.Add("pbm:PbmFaultProfileStorageFaultFault", reflect.TypeOf((*PbmFaultProfileStorageFaultFault)(nil)).Elem())
}

type PbmFetchCapabilityMetadata PbmFetchCapabilityMetadataRequestType

func init() {
	types.Add("pbm:PbmFetchCapabilityMetadata", reflect.TypeOf((*PbmFetchCapabilityMetadata)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmFetchCapabilityMetadata`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmFetchCapabilityMetadataRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Type of profile resource. The Server supports the "STORAGE" resource
	// type only. If not specified, this method will return capability metadata for the storage
	// resources. Any other <code>resourceType</code> is considered invalid.
	ResourceType *PbmProfileResourceType `xml:"resourceType,omitempty" json:"resourceType,omitempty"`
	// Unique identifier for the vendor/owner of capability
	// metadata. The specified vendor ID must match
	// `PbmCapabilitySchemaVendorInfo*.*PbmCapabilitySchemaVendorInfo.vendorUuid`.
	// If omitted, the Server searchs all capability metadata registered with the system. If a
	// <code>vendorUuid</code> unknown to the Server is specified, empty results will be returned.
	VendorUuid string `xml:"vendorUuid,omitempty" json:"vendorUuid,omitempty"`
}

func init() {
	types.Add("pbm:PbmFetchCapabilityMetadataRequestType", reflect.TypeOf((*PbmFetchCapabilityMetadataRequestType)(nil)).Elem())
}

type PbmFetchCapabilityMetadataResponse struct {
	Returnval []PbmCapabilityMetadataPerCategory `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmFetchCapabilitySchema PbmFetchCapabilitySchemaRequestType

func init() {
	types.Add("pbm:PbmFetchCapabilitySchema", reflect.TypeOf((*PbmFetchCapabilitySchema)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmFetchCapabilitySchema`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmFetchCapabilitySchemaRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Unique identifier for the vendor/owner of capability metadata.
	// If omitted, the server searchs all capability metadata registered
	// with the system. The specified vendor ID must match
	// `PbmCapabilitySchemaVendorInfo*.*PbmCapabilitySchemaVendorInfo.vendorUuid`.
	VendorUuid string `xml:"vendorUuid,omitempty" json:"vendorUuid,omitempty"`
	// Optional line of service that must match `PbmLineOfServiceInfoLineOfServiceEnum_enum`.
	// If specified, the capability schema objects
	// are returned for the given lineOfServices. If null, then all
	// capability schema objects that may or may not have data service capabilities
	// are returned.
	LineOfService []string `xml:"lineOfService,omitempty" json:"lineOfService,omitempty"`
}

func init() {
	types.Add("pbm:PbmFetchCapabilitySchemaRequestType", reflect.TypeOf((*PbmFetchCapabilitySchemaRequestType)(nil)).Elem())
}

type PbmFetchCapabilitySchemaResponse struct {
	Returnval []PbmCapabilitySchema `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmFetchComplianceResult PbmFetchComplianceResultRequestType

func init() {
	types.Add("pbm:PbmFetchComplianceResult", reflect.TypeOf((*PbmFetchComplianceResult)(nil)).Elem())
}

// The parameters of `PbmComplianceManager.PbmFetchComplianceResult`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmFetchComplianceResultRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// One or more references to storage entities.
	// A maximum of 1000 virtual machines and/or virtual disks can be specified
	// in a call. The results of calling the fetchComplianceResult API with
	// more than a 1000 entities is undefined.
	//   - If the list of entities also contains datastores, the Server
	//     will ignore the datastores.
	//   - If the list contains valid and invalid entities, the Server ignores
	//     the invalid entities and returns results for the valid entities.
	//     Invalid entities are entities that are not in the vCenter inventory.
	//   - If the list contains only datastores, the method throws
	//     an <code>InvalidArgument</code> fault.
	Entities []PbmServerObjectRef `xml:"entities" json:"entities"`
	// Not used. if specified, the Server ignores the value.
	// The Server uses the profiles associated with the specified entities.
	Profile *PbmProfileId `xml:"profile,omitempty" json:"profile,omitempty"`
}

func init() {
	types.Add("pbm:PbmFetchComplianceResultRequestType", reflect.TypeOf((*PbmFetchComplianceResultRequestType)(nil)).Elem())
}

type PbmFetchComplianceResultResponse struct {
	Returnval []PbmComplianceResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

// The `PbmFetchEntityHealthStatusSpec` data object contains
// the arguments required for
// `PbmComplianceManager.PbmFetchEntityHealthStatusExt`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmFetchEntityHealthStatusSpec struct {
	types.DynamicData

	// `PbmServerObjectRef` for which the healthStatus is required
	ObjectRef PbmServerObjectRef `xml:"objectRef" json:"objectRef"`
	// BackingId for the ServerObjectRef
	// BackingId is mandatory for FCD on vSAN
	BackingId string `xml:"backingId,omitempty" json:"backingId,omitempty"`
}

func init() {
	types.Add("pbm:PbmFetchEntityHealthStatusSpec", reflect.TypeOf((*PbmFetchEntityHealthStatusSpec)(nil)).Elem())
}

type PbmFetchResourceType PbmFetchResourceTypeRequestType

func init() {
	types.Add("pbm:PbmFetchResourceType", reflect.TypeOf((*PbmFetchResourceType)(nil)).Elem())
}

type PbmFetchResourceTypeRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("pbm:PbmFetchResourceTypeRequestType", reflect.TypeOf((*PbmFetchResourceTypeRequestType)(nil)).Elem())
}

type PbmFetchResourceTypeResponse struct {
	Returnval []PbmProfileResourceType `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmFetchRollupComplianceResult PbmFetchRollupComplianceResultRequestType

func init() {
	types.Add("pbm:PbmFetchRollupComplianceResult", reflect.TypeOf((*PbmFetchRollupComplianceResult)(nil)).Elem())
}

// The parameters of `PbmComplianceManager.PbmFetchRollupComplianceResult`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmFetchRollupComplianceResultRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// One or more virtual machines.
	// A maximum of 1000 virtual machines can be specified
	// in a call. The results of calling the fetchRollupComplianceResult API with
	// more than a 1000 entity objects is undefined.
	Entity []PbmServerObjectRef `xml:"entity" json:"entity"`
}

func init() {
	types.Add("pbm:PbmFetchRollupComplianceResultRequestType", reflect.TypeOf((*PbmFetchRollupComplianceResultRequestType)(nil)).Elem())
}

type PbmFetchRollupComplianceResultResponse struct {
	Returnval []PbmRollupComplianceResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmFetchVendorInfo PbmFetchVendorInfoRequestType

func init() {
	types.Add("pbm:PbmFetchVendorInfo", reflect.TypeOf((*PbmFetchVendorInfo)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmFetchVendorInfo`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmFetchVendorInfoRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Specifies the resource type. The Server supports the STORAGE resource
	// type only. If not specified, server defaults to STORAGE resource type. Any other
	// <code>resourceType</code> is considered invalid.
	ResourceType *PbmProfileResourceType `xml:"resourceType,omitempty" json:"resourceType,omitempty"`
}

func init() {
	types.Add("pbm:PbmFetchVendorInfoRequestType", reflect.TypeOf((*PbmFetchVendorInfoRequestType)(nil)).Elem())
}

type PbmFetchVendorInfoResponse struct {
	Returnval []PbmCapabilityVendorResourceTypeInfo `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmFindApplicableDefaultProfile PbmFindApplicableDefaultProfileRequestType

func init() {
	types.Add("pbm:PbmFindApplicableDefaultProfile", reflect.TypeOf((*PbmFindApplicableDefaultProfile)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmFindApplicableDefaultProfile`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmFindApplicableDefaultProfileRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Datastores for which the default profile is found out. Note that
	// the datastore pods/clusters are not supported.
	Datastores []PbmPlacementHub `xml:"datastores" json:"datastores"`
}

func init() {
	types.Add("pbm:PbmFindApplicableDefaultProfileRequestType", reflect.TypeOf((*PbmFindApplicableDefaultProfileRequestType)(nil)).Elem())
}

type PbmFindApplicableDefaultProfileResponse struct {
	Returnval []BasePbmProfile `xml:"returnval,omitempty,typeattr" json:"returnval,omitempty"`
}

// Warning fault used to indicate that the vendor specific datastore matches the tag in the
// requirements profile but doesnt match the vendor specific rule set in the requirements profile.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmIncompatibleVendorSpecificRuleSet struct {
	PbmCapabilityProfilePropertyMismatchFault
}

func init() {
	types.Add("pbm:PbmIncompatibleVendorSpecificRuleSet", reflect.TypeOf((*PbmIncompatibleVendorSpecificRuleSet)(nil)).Elem())
}

type PbmIncompatibleVendorSpecificRuleSetFault PbmIncompatibleVendorSpecificRuleSet

func init() {
	types.Add("pbm:PbmIncompatibleVendorSpecificRuleSetFault", reflect.TypeOf((*PbmIncompatibleVendorSpecificRuleSetFault)(nil)).Elem())
}

// LegacyHubsNotSupported fault is thrown to indicate the legacy hubs that are not supported.
//
// For storage, legacy hubs or datastores are VMFS and NFS datastores.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmLegacyHubsNotSupported struct {
	PbmFault

	// Legacy hubs that are not supported.
	//
	// Only datastores will be populated in this fault. Datastore clusters
	// are not allowed.
	Hubs []PbmPlacementHub `xml:"hubs" json:"hubs"`
}

func init() {
	types.Add("pbm:PbmLegacyHubsNotSupported", reflect.TypeOf((*PbmLegacyHubsNotSupported)(nil)).Elem())
}

type PbmLegacyHubsNotSupportedFault PbmLegacyHubsNotSupported

func init() {
	types.Add("pbm:PbmLegacyHubsNotSupportedFault", reflect.TypeOf((*PbmLegacyHubsNotSupportedFault)(nil)).Elem())
}

// Describes Line of Service of a capability provider.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmLineOfServiceInfo struct {
	types.DynamicData

	// `PbmLineOfServiceInfoLineOfServiceEnum_enum` - must be one of the values
	// for enum `PbmLineOfServiceInfoLineOfServiceEnum_enum`.
	LineOfService string `xml:"lineOfService" json:"lineOfService"`
	// Name of the service - for informational
	// purposes only.
	Name PbmExtendedElementDescription `xml:"name" json:"name"`
	// Description of the service - for informational
	// purposes only.
	Description *PbmExtendedElementDescription `xml:"description,omitempty" json:"description,omitempty"`
}

func init() {
	types.Add("pbm:PbmLineOfServiceInfo", reflect.TypeOf((*PbmLineOfServiceInfo)(nil)).Elem())
}

// This structure may be used only with operations rendered under `/pbm`.
type PbmLoggingConfiguration struct {
	types.DynamicData

	Component string `xml:"component" json:"component"`
	LogLevel  string `xml:"logLevel" json:"logLevel"`
}

func init() {
	types.Add("pbm:PbmLoggingConfiguration", reflect.TypeOf((*PbmLoggingConfiguration)(nil)).Elem())
}

// NonExistentHubs is thrown to indicate that some non existent datastores are used.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmNonExistentHubs struct {
	PbmFault

	// Legacy hubs that do not exist.
	Hubs []PbmPlacementHub `xml:"hubs" json:"hubs"`
}

func init() {
	types.Add("pbm:PbmNonExistentHubs", reflect.TypeOf((*PbmNonExistentHubs)(nil)).Elem())
}

type PbmNonExistentHubsFault PbmNonExistentHubs

func init() {
	types.Add("pbm:PbmNonExistentHubsFault", reflect.TypeOf((*PbmNonExistentHubsFault)(nil)).Elem())
}

// Describes the data services provided by the storage arrays.
//
// In addition to storing bits, some VASA providers may also want to separate
// their capabilities into lines of service to let vSphere manage finer grain
// policies. For example an array may support replication natively, and may
// want vSphere policies to be defined for the replication aspect separately
// and compose them with persistence related policies.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmPersistenceBasedDataServiceInfo struct {
	PbmLineOfServiceInfo

	// This property should be set with compatible schema namespaces exposed by
	// the vendor provider.
	//
	// If not specified, vSphere assumes all Data Service
	// provider schemas are compatible with all persistence provider namespaces
	// advertised by the VASA provider.
	CompatiblePersistenceSchemaNamespace []string `xml:"compatiblePersistenceSchemaNamespace,omitempty" json:"compatiblePersistenceSchemaNamespace,omitempty"`
}

func init() {
	types.Add("pbm:PbmPersistenceBasedDataServiceInfo", reflect.TypeOf((*PbmPersistenceBasedDataServiceInfo)(nil)).Elem())
}

// Requirement type containing capability constraints
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmPlacementCapabilityConstraintsRequirement struct {
	PbmPlacementRequirement

	// Capability constraints
	Constraints BasePbmCapabilityConstraints `xml:"constraints,typeattr" json:"constraints"`
}

func init() {
	types.Add("pbm:PbmPlacementCapabilityConstraintsRequirement", reflect.TypeOf((*PbmPlacementCapabilityConstraintsRequirement)(nil)).Elem())
}

// A Requirement for a particular `PbmCapabilityProfile`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmPlacementCapabilityProfileRequirement struct {
	PbmPlacementRequirement

	// Reference to the capability profile being used as a requirement
	ProfileId PbmProfileId `xml:"profileId" json:"profileId"`
}

func init() {
	types.Add("pbm:PbmPlacementCapabilityProfileRequirement", reflect.TypeOf((*PbmPlacementCapabilityProfileRequirement)(nil)).Elem())
}

// The `PbmPlacementCompatibilityResult` data object
// contains the compatibility result of a placement request.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmPlacementCompatibilityResult struct {
	types.DynamicData

	// The <code>Datastore</code> or <code>StoragePod</code> under consideration
	// as a location for virtual machine files.
	Hub PbmPlacementHub `xml:"hub" json:"hub"`
	// Resources that match the policy.
	//
	// If populated, signifies that there are
	// specific resources that match the policy for `PbmPlacementCompatibilityResult.hub`. If null,
	// signifies that all resources (for example, hosts connected to the
	// datastore or storage pod) are compatible.
	MatchingResources []BasePbmPlacementMatchingResources `xml:"matchingResources,omitempty,typeattr" json:"matchingResources,omitempty"`
	// How many objects of the kind requested can be provisioned on this
	// `PbmPlacementCompatibilityResult.hub`.
	HowMany int64 `xml:"howMany,omitempty" json:"howMany,omitempty"`
	// This field is not populated if there is no size in the query, i.e.
	//
	// if the request carries only policy and no size requirements, this
	// will not be populated.
	Utilization []PbmPlacementResourceUtilization `xml:"utilization,omitempty" json:"utilization,omitempty"`
	// Array of faults that describe issues that may affect profile compatibility.
	//
	// Users should consider these issues before using this <code>Datastore</code>
	// or <code>StoragePod</code> and a connected <code>Host</code>s.
	Warning []types.LocalizedMethodFault `xml:"warning,omitempty" json:"warning,omitempty"`
	// Array of faults that prevent this datastore or storage pod from being compatible with the
	// specified profile, including if no host connected to this `PbmPlacementCompatibilityResult.hub` is compatible.
	Error []types.LocalizedMethodFault `xml:"error,omitempty" json:"error,omitempty"`
}

func init() {
	types.Add("pbm:PbmPlacementCompatibilityResult", reflect.TypeOf((*PbmPlacementCompatibilityResult)(nil)).Elem())
}

// A `PbmPlacementHub` data object identifies a storage location
// where virtual machine files can be placed.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmPlacementHub struct {
	types.DynamicData

	// Type of the hub.
	//
	// Currently ManagedObject is the only supported type.
	HubType string `xml:"hubType" json:"hubType"`
	// Hub identifier; a ManagedObjectReference to a datastore or a storage pod.
	HubId string `xml:"hubId" json:"hubId"`
}

func init() {
	types.Add("pbm:PbmPlacementHub", reflect.TypeOf((*PbmPlacementHub)(nil)).Elem())
}

// Describes the collection of replication related resources that satisfy a
// policy, for a specific datastore.
//
// This class is returned only when the policy contains replication capabilities.
// For a storage pod, only those replication groups that are common across
// all datastores in the storage pod are considered compatible.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmPlacementMatchingReplicationResources struct {
	PbmPlacementMatchingResources

	// Replication groups that match the policy.
	ReplicationGroup []types.ReplicationGroupId `xml:"replicationGroup,omitempty" json:"replicationGroup,omitempty"`
}

func init() {
	types.Add("pbm:PbmPlacementMatchingReplicationResources", reflect.TypeOf((*PbmPlacementMatchingReplicationResources)(nil)).Elem())
}

// Describes the collection of resources (for example, hosts) that satisfy a
// policy, for a specific datastore.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmPlacementMatchingResources struct {
	types.DynamicData
}

func init() {
	types.Add("pbm:PbmPlacementMatchingResources", reflect.TypeOf((*PbmPlacementMatchingResources)(nil)).Elem())
}

// Defines a constraint for placing objects onto `PbmPlacementHub`s.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmPlacementRequirement struct {
	types.DynamicData
}

func init() {
	types.Add("pbm:PbmPlacementRequirement", reflect.TypeOf((*PbmPlacementRequirement)(nil)).Elem())
}

// Describes the resource utilization metrics of a datastore.
//
// These results are not to be treated as a guaranteed availability,
// they are useful to estimate the effects of a change of policy
// or the effects of a provisioning action.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmPlacementResourceUtilization struct {
	types.DynamicData

	// Name of the resource.
	Name PbmExtendedElementDescription `xml:"name" json:"name"`
	// Description of the resource.
	Description PbmExtendedElementDescription `xml:"description" json:"description"`
	// Currently available (i.e.
	//
	// before the provisioning step).
	AvailableBefore int64 `xml:"availableBefore,omitempty" json:"availableBefore,omitempty"`
	// Available after the provisioning step.
	AvailableAfter int64 `xml:"availableAfter,omitempty" json:"availableAfter,omitempty"`
	// Total resource availability
	Total int64 `xml:"total,omitempty" json:"total,omitempty"`
}

func init() {
	types.Add("pbm:PbmPlacementResourceUtilization", reflect.TypeOf((*PbmPlacementResourceUtilization)(nil)).Elem())
}

// The `PbmProfile` data object is the base object
// for storage capability profiles.
//
// This object defines metadata
// for the profile. The derived capability profile represents the
// user's intent for selection and configuration of storage resources
// and/or services that support deployment of virtual machines
// and virtual disks.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmProfile struct {
	types.DynamicData

	// Unique identifier for the profile.
	ProfileId PbmProfileId `xml:"profileId" json:"profileId"`
	Name      string       `xml:"name" json:"name"`
	// Profile description.
	Description string `xml:"description,omitempty" json:"description,omitempty"`
	// Time stamp of profile creation.
	CreationTime time.Time `xml:"creationTime" json:"creationTime"`
	// User name of the profile creator.
	//
	// Set during creation time.
	CreatedBy string `xml:"createdBy" json:"createdBy"`
	// Time stamp of latest modification to the profile.
	LastUpdatedTime time.Time `xml:"lastUpdatedTime" json:"lastUpdatedTime"`
	// Name of the user performing the latest modification of the profile.
	LastUpdatedBy string `xml:"lastUpdatedBy" json:"lastUpdatedBy"`
}

func init() {
	types.Add("pbm:PbmProfile", reflect.TypeOf((*PbmProfile)(nil)).Elem())
}

// Profile unique identifier.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmProfileId struct {
	types.DynamicData

	// Unique identifier of the profile.
	UniqueId string `xml:"uniqueId" json:"uniqueId"`
}

func init() {
	types.Add("pbm:PbmProfileId", reflect.TypeOf((*PbmProfileId)(nil)).Elem())
}

// The `PbmProfileOperationOutcome` data object describes the result
// of a `PbmProfileProfileManager` operation.
//
// If there was an
// error during the operation, the object identifies the fault.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmProfileOperationOutcome struct {
	types.DynamicData

	// Identifies the profile specified for the operation.
	ProfileId PbmProfileId `xml:"profileId" json:"profileId"`
	// One of the `PbmFault` objects.
	Fault *types.LocalizedMethodFault `xml:"fault,omitempty" json:"fault,omitempty"`
}

func init() {
	types.Add("pbm:PbmProfileOperationOutcome", reflect.TypeOf((*PbmProfileOperationOutcome)(nil)).Elem())
}

// The `PbmProfileResourceType` data object defines the vSphere resource type
// that is supported for profile management.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmProfileResourceType struct {
	types.DynamicData

	// Type of resource to which capability information applies.
	//
	// <code>resourceType</code> is a string value that corresponds to
	// a `PbmProfileResourceTypeEnum_enum` enumeration value.
	// Only the STORAGE resource type is supported.
	ResourceType string `xml:"resourceType" json:"resourceType"`
}

func init() {
	types.Add("pbm:PbmProfileResourceType", reflect.TypeOf((*PbmProfileResourceType)(nil)).Elem())
}

// The `PbmProfileType` identifier is defined by storage providers
// to distinguish between different types of profiles plugged into the system.
//
// An example of a system supported profile type is "CapabilityBasedProfileType"
// which will be the type used for all capability-based profiles created by
// the system using capability metadata information published to the system.
//
// For internal use only.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmProfileType struct {
	types.DynamicData

	// Unique type identifier for this profile type.
	//
	// eg "CapabilityBased", or other.
	UniqueId string `xml:"uniqueId" json:"uniqueId"`
}

func init() {
	types.Add("pbm:PbmProfileType", reflect.TypeOf((*PbmProfileType)(nil)).Elem())
}

// Fault used to indicate which property instance in requirements profile that does not
// match.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmPropertyMismatchFault struct {
	PbmCompatibilityCheckFault

	// Id of the CapabilityInstance in requirements profile that
	// does not match.
	CapabilityInstanceId PbmCapabilityMetadataUniqueId `xml:"capabilityInstanceId" json:"capabilityInstanceId"`
	// The property instance in requirement profile that does not match.
	RequirementPropertyInstance PbmCapabilityPropertyInstance `xml:"requirementPropertyInstance" json:"requirementPropertyInstance"`
}

func init() {
	types.Add("pbm:PbmPropertyMismatchFault", reflect.TypeOf((*PbmPropertyMismatchFault)(nil)).Elem())
}

type PbmPropertyMismatchFaultFault BasePbmPropertyMismatchFault

func init() {
	types.Add("pbm:PbmPropertyMismatchFaultFault", reflect.TypeOf((*PbmPropertyMismatchFaultFault)(nil)).Elem())
}

type PbmQueryAssociatedEntities PbmQueryAssociatedEntitiesRequestType

func init() {
	types.Add("pbm:PbmQueryAssociatedEntities", reflect.TypeOf((*PbmQueryAssociatedEntities)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmQueryAssociatedEntities`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmQueryAssociatedEntitiesRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Storage policy array.
	Profiles []PbmProfileId `xml:"profiles,omitempty" json:"profiles,omitempty"`
}

func init() {
	types.Add("pbm:PbmQueryAssociatedEntitiesRequestType", reflect.TypeOf((*PbmQueryAssociatedEntitiesRequestType)(nil)).Elem())
}

type PbmQueryAssociatedEntitiesResponse struct {
	Returnval []PbmQueryProfileResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmQueryAssociatedEntity PbmQueryAssociatedEntityRequestType

func init() {
	types.Add("pbm:PbmQueryAssociatedEntity", reflect.TypeOf((*PbmQueryAssociatedEntity)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmQueryAssociatedEntity`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmQueryAssociatedEntityRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Profile identifier.
	Profile PbmProfileId `xml:"profile" json:"profile"`
	// If specified, the method returns only those entities
	// which match the type. The <code>entityType</code> string value must match
	// one of the `PbmObjectType_enum` values.
	// If not specified, the method returns all entities associated with the profile.
	EntityType string `xml:"entityType,omitempty" json:"entityType,omitempty"`
}

func init() {
	types.Add("pbm:PbmQueryAssociatedEntityRequestType", reflect.TypeOf((*PbmQueryAssociatedEntityRequestType)(nil)).Elem())
}

type PbmQueryAssociatedEntityResponse struct {
	Returnval []PbmServerObjectRef `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmQueryAssociatedProfile PbmQueryAssociatedProfileRequestType

func init() {
	types.Add("pbm:PbmQueryAssociatedProfile", reflect.TypeOf((*PbmQueryAssociatedProfile)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmQueryAssociatedProfile`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmQueryAssociatedProfileRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Reference to a virtual machine, virtual disk, or datastore.
	Entity PbmServerObjectRef `xml:"entity" json:"entity"`
}

func init() {
	types.Add("pbm:PbmQueryAssociatedProfileRequestType", reflect.TypeOf((*PbmQueryAssociatedProfileRequestType)(nil)).Elem())
}

type PbmQueryAssociatedProfileResponse struct {
	Returnval []PbmProfileId `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmQueryAssociatedProfiles PbmQueryAssociatedProfilesRequestType

func init() {
	types.Add("pbm:PbmQueryAssociatedProfiles", reflect.TypeOf((*PbmQueryAssociatedProfiles)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmQueryAssociatedProfiles`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmQueryAssociatedProfilesRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Array of server object references.
	Entities []PbmServerObjectRef `xml:"entities" json:"entities"`
}

func init() {
	types.Add("pbm:PbmQueryAssociatedProfilesRequestType", reflect.TypeOf((*PbmQueryAssociatedProfilesRequestType)(nil)).Elem())
}

type PbmQueryAssociatedProfilesResponse struct {
	Returnval []PbmQueryProfileResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmQueryByRollupComplianceStatus PbmQueryByRollupComplianceStatusRequestType

func init() {
	types.Add("pbm:PbmQueryByRollupComplianceStatus", reflect.TypeOf((*PbmQueryByRollupComplianceStatus)(nil)).Elem())
}

// The parameters of `PbmComplianceManager.PbmQueryByRollupComplianceStatus`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmQueryByRollupComplianceStatusRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// `PbmComplianceStatus_enum`
	Status string `xml:"status" json:"status"`
}

func init() {
	types.Add("pbm:PbmQueryByRollupComplianceStatusRequestType", reflect.TypeOf((*PbmQueryByRollupComplianceStatusRequestType)(nil)).Elem())
}

type PbmQueryByRollupComplianceStatusResponse struct {
	Returnval []PbmServerObjectRef `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmQueryDefaultRequirementProfile PbmQueryDefaultRequirementProfileRequestType

func init() {
	types.Add("pbm:PbmQueryDefaultRequirementProfile", reflect.TypeOf((*PbmQueryDefaultRequirementProfile)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmQueryDefaultRequirementProfile`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmQueryDefaultRequirementProfileRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Placement hub (i.e. datastore).
	Hub PbmPlacementHub `xml:"hub" json:"hub"`
}

func init() {
	types.Add("pbm:PbmQueryDefaultRequirementProfileRequestType", reflect.TypeOf((*PbmQueryDefaultRequirementProfileRequestType)(nil)).Elem())
}

type PbmQueryDefaultRequirementProfileResponse struct {
	Returnval *PbmProfileId `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmQueryDefaultRequirementProfiles PbmQueryDefaultRequirementProfilesRequestType

func init() {
	types.Add("pbm:PbmQueryDefaultRequirementProfiles", reflect.TypeOf((*PbmQueryDefaultRequirementProfiles)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmQueryDefaultRequirementProfiles`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmQueryDefaultRequirementProfilesRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// The datastores for which the default profiles are requested. For
	// legacy datastores we set
	// `DefaultProfileInfo.defaultProfile` to `null`.
	Datastores []PbmPlacementHub `xml:"datastores" json:"datastores"`
}

func init() {
	types.Add("pbm:PbmQueryDefaultRequirementProfilesRequestType", reflect.TypeOf((*PbmQueryDefaultRequirementProfilesRequestType)(nil)).Elem())
}

type PbmQueryDefaultRequirementProfilesResponse struct {
	Returnval []PbmDefaultProfileInfo `xml:"returnval" json:"returnval"`
}

type PbmQueryMatchingHub PbmQueryMatchingHubRequestType

func init() {
	types.Add("pbm:PbmQueryMatchingHub", reflect.TypeOf((*PbmQueryMatchingHub)(nil)).Elem())
}

// The parameters of `PbmPlacementSolver.PbmQueryMatchingHub`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmQueryMatchingHubRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Candidate list of hubs, either datastores or storage pods or a
	// mix. If this parameter is not specified, the Server uses all
	// of the datastores and storage pods.
	HubsToSearch []PbmPlacementHub `xml:"hubsToSearch,omitempty" json:"hubsToSearch,omitempty"`
	// Storage requirement profile.
	Profile PbmProfileId `xml:"profile" json:"profile"`
}

func init() {
	types.Add("pbm:PbmQueryMatchingHubRequestType", reflect.TypeOf((*PbmQueryMatchingHubRequestType)(nil)).Elem())
}

type PbmQueryMatchingHubResponse struct {
	Returnval []PbmPlacementHub `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmQueryMatchingHubWithSpec PbmQueryMatchingHubWithSpecRequestType

func init() {
	types.Add("pbm:PbmQueryMatchingHubWithSpec", reflect.TypeOf((*PbmQueryMatchingHubWithSpec)(nil)).Elem())
}

// The parameters of `PbmPlacementSolver.PbmQueryMatchingHubWithSpec`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmQueryMatchingHubWithSpecRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Candidate list of hubs, either datastores or storage
	// pods or a mix. If this parameter is not specified, the Server uses
	// all of the datastores and storage pods for placement compatibility checking.
	HubsToSearch []PbmPlacementHub `xml:"hubsToSearch,omitempty" json:"hubsToSearch,omitempty"`
	// Storage profile creation specification.
	CreateSpec PbmCapabilityProfileCreateSpec `xml:"createSpec" json:"createSpec"`
}

func init() {
	types.Add("pbm:PbmQueryMatchingHubWithSpecRequestType", reflect.TypeOf((*PbmQueryMatchingHubWithSpecRequestType)(nil)).Elem())
}

type PbmQueryMatchingHubWithSpecResponse struct {
	Returnval []PbmPlacementHub `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmQueryProfile PbmQueryProfileRequestType

func init() {
	types.Add("pbm:PbmQueryProfile", reflect.TypeOf((*PbmQueryProfile)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmQueryProfile`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmQueryProfileRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Type of resource. You can specify only STORAGE.
	ResourceType PbmProfileResourceType `xml:"resourceType" json:"resourceType"`
	// Profile category. The string value must correspond
	// to one of the `PbmProfileCategoryEnum_enum` values.
	// If you do not specify a profile category, the method returns profiles in all
	// categories.
	ProfileCategory string `xml:"profileCategory,omitempty" json:"profileCategory,omitempty"`
}

func init() {
	types.Add("pbm:PbmQueryProfileRequestType", reflect.TypeOf((*PbmQueryProfileRequestType)(nil)).Elem())
}

type PbmQueryProfileResponse struct {
	Returnval []PbmProfileId `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

// The `PbmQueryProfileResult` data object
// identifies a virtual machine, virtual disk, or datastore
// and it lists the identifier(s) for the associated profile(s).
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmQueryProfileResult struct {
	types.DynamicData

	// Reference to the virtual machine, virtual disk, or
	// datastore on which the query was performed.
	Object PbmServerObjectRef `xml:"object" json:"object"`
	// Array of identifiers for profiles which are associated with <code>object</code>.
	ProfileId []PbmProfileId `xml:"profileId,omitempty" json:"profileId,omitempty"`
	// Fault associated with the query, if there is one.
	Fault *types.LocalizedMethodFault `xml:"fault,omitempty" json:"fault,omitempty"`
}

func init() {
	types.Add("pbm:PbmQueryProfileResult", reflect.TypeOf((*PbmQueryProfileResult)(nil)).Elem())
}

// The `PbmQueryReplicationGroupResult` data object
// identifies a virtual machine, or a virtual disk and lists the identifier(s) for the associated
// replication group.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmQueryReplicationGroupResult struct {
	types.DynamicData

	// Reference to the virtual machine or virtual disk on which the query was performed.
	//
	// If the
	// query was performed for a virtual machine and all it's disks, this will reference each disk
	// and the virtual machine config individually.
	Object PbmServerObjectRef `xml:"object" json:"object"`
	// Replication group identifier which is associated with <code>object</code>.
	ReplicationGroupId *types.ReplicationGroupId `xml:"replicationGroupId,omitempty" json:"replicationGroupId,omitempty"`
	// Fault associated with the query, if there is one.
	Fault *types.LocalizedMethodFault `xml:"fault,omitempty" json:"fault,omitempty"`
}

func init() {
	types.Add("pbm:PbmQueryReplicationGroupResult", reflect.TypeOf((*PbmQueryReplicationGroupResult)(nil)).Elem())
}

type PbmQueryReplicationGroups PbmQueryReplicationGroupsRequestType

func init() {
	types.Add("pbm:PbmQueryReplicationGroups", reflect.TypeOf((*PbmQueryReplicationGroups)(nil)).Elem())
}

// The parameters of `PbmReplicationManager.PbmQueryReplicationGroups`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmQueryReplicationGroupsRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Array of server object references. Valid types are
	// `virtualMachine`,
	// `virtualMachineAndDisks`,
	// `virtualDiskId`,
	// `virtualDiskUUID`
	Entities []PbmServerObjectRef `xml:"entities,omitempty" json:"entities,omitempty"`
}

func init() {
	types.Add("pbm:PbmQueryReplicationGroupsRequestType", reflect.TypeOf((*PbmQueryReplicationGroupsRequestType)(nil)).Elem())
}

type PbmQueryReplicationGroupsResponse struct {
	Returnval []PbmQueryReplicationGroupResult `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmQuerySpaceStatsForStorageContainer PbmQuerySpaceStatsForStorageContainerRequestType

func init() {
	types.Add("pbm:PbmQuerySpaceStatsForStorageContainer", reflect.TypeOf((*PbmQuerySpaceStatsForStorageContainer)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmQuerySpaceStatsForStorageContainer`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmQuerySpaceStatsForStorageContainerRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Entity for which space statistics are being requested i.e datastore.
	Datastore PbmServerObjectRef `xml:"datastore" json:"datastore"`
	// \- capability profile Ids.
	// If omitted, the statistics for the container
	// as a whole would be returned.
	CapabilityProfileId []PbmProfileId `xml:"capabilityProfileId,omitempty" json:"capabilityProfileId,omitempty"`
}

func init() {
	types.Add("pbm:PbmQuerySpaceStatsForStorageContainerRequestType", reflect.TypeOf((*PbmQuerySpaceStatsForStorageContainerRequestType)(nil)).Elem())
}

type PbmQuerySpaceStatsForStorageContainerResponse struct {
	Returnval []PbmDatastoreSpaceStatistics `xml:"returnval,omitempty" json:"returnval,omitempty"`
}

type PbmResetDefaultRequirementProfile PbmResetDefaultRequirementProfileRequestType

func init() {
	types.Add("pbm:PbmResetDefaultRequirementProfile", reflect.TypeOf((*PbmResetDefaultRequirementProfile)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmResetDefaultRequirementProfile`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmResetDefaultRequirementProfileRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Profile to reset.
	Profile *PbmProfileId `xml:"profile,omitempty" json:"profile,omitempty"`
}

func init() {
	types.Add("pbm:PbmResetDefaultRequirementProfileRequestType", reflect.TypeOf((*PbmResetDefaultRequirementProfileRequestType)(nil)).Elem())
}

type PbmResetDefaultRequirementProfileResponse struct {
}

type PbmResetVSanDefaultProfile PbmResetVSanDefaultProfileRequestType

func init() {
	types.Add("pbm:PbmResetVSanDefaultProfile", reflect.TypeOf((*PbmResetVSanDefaultProfile)(nil)).Elem())
}

type PbmResetVSanDefaultProfileRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("pbm:PbmResetVSanDefaultProfileRequestType", reflect.TypeOf((*PbmResetVSanDefaultProfileRequestType)(nil)).Elem())
}

type PbmResetVSanDefaultProfileResponse struct {
}

// A ResourceInUse fault indicating that some error has occurred because a
// resource was in use.
//
// Information about the resource that is in use may
// be supplied.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmResourceInUse struct {
	PbmFault

	// Type of resource that is in use.
	Type string `xml:"type,omitempty" json:"type,omitempty"`
	// Name of the instance of the resource that is in use.
	Name string `xml:"name,omitempty" json:"name,omitempty"`
}

func init() {
	types.Add("pbm:PbmResourceInUse", reflect.TypeOf((*PbmResourceInUse)(nil)).Elem())
}

type PbmResourceInUseFault PbmResourceInUse

func init() {
	types.Add("pbm:PbmResourceInUseFault", reflect.TypeOf((*PbmResourceInUseFault)(nil)).Elem())
}

type PbmRetrieveContent PbmRetrieveContentRequestType

func init() {
	types.Add("pbm:PbmRetrieveContent", reflect.TypeOf((*PbmRetrieveContent)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmRetrieveContent`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmRetrieveContentRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Array of storage profile identifiers.
	ProfileIds []PbmProfileId `xml:"profileIds" json:"profileIds"`
}

func init() {
	types.Add("pbm:PbmRetrieveContentRequestType", reflect.TypeOf((*PbmRetrieveContentRequestType)(nil)).Elem())
}

type PbmRetrieveContentResponse struct {
	Returnval []BasePbmProfile `xml:"returnval,typeattr" json:"returnval"`
}

type PbmRetrieveServiceContent PbmRetrieveServiceContentRequestType

func init() {
	types.Add("pbm:PbmRetrieveServiceContent", reflect.TypeOf((*PbmRetrieveServiceContent)(nil)).Elem())
}

type PbmRetrieveServiceContentRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
}

func init() {
	types.Add("pbm:PbmRetrieveServiceContentRequestType", reflect.TypeOf((*PbmRetrieveServiceContentRequestType)(nil)).Elem())
}

type PbmRetrieveServiceContentResponse struct {
	Returnval PbmServiceInstanceContent `xml:"returnval" json:"returnval"`
}

// The `PbmRollupComplianceResult` data object identifies the virtual machine
// for which rollup compliance was checked, and it contains the overall status
// and a list of compliance result objects.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmRollupComplianceResult struct {
	types.DynamicData

	// Indicates the earliest time that compliance was checked for any
	// of the entities in the rollup compliance check.
	//
	// The compliance
	// check time for a single entity is represented in the
	// `PbmComplianceResult*.*PbmComplianceResult.checkTime`
	// property. If the `PbmComplianceResult.checkTime`
	// property is unset for any of the objects in the <code>results</code>
	// array, the <code>oldestCheckTime</code> property will be unset.
	OldestCheckTime time.Time `xml:"oldestCheckTime" json:"oldestCheckTime"`
	// Virtual machine for which the rollup compliance was checked.
	Entity PbmServerObjectRef `xml:"entity" json:"entity"`
	// Overall compliance status of the virtual machine and its virtual disks.
	//
	// <code>overallComplianceStatus</code> is a string value that
	// corresponds to one of the
	// `PbmComplianceResult*.*PbmComplianceResult.complianceStatus`
	// values.
	//
	// The overall compliance status is determined by the following rules, applied in the order
	// listed:
	//   - If all the entities are <code>compliant</code>, the overall status is
	//     <code>compliant</code>.
	//   - Else if any entity's status is <code>outOfDate</code>, the overall status is
	//     <code>outOfDate</code>.
	//   - Else if any entity's status is <code>nonCompliant</code>, the overall status is
	//     <code>nonCompliant</code>.
	//   - Else if any entity's status is <code>unknown</code>, the overall status is
	//     <code>unknown</code>.
	//   - Else if any entity's status is <code>notApplicable</code>, the overall status is
	//     <code>notApplicable</code>.
	OverallComplianceStatus string `xml:"overallComplianceStatus" json:"overallComplianceStatus"`
	// Overall compliance task status of the virtual machine and its virtual
	// disks.
	//
	// <code>overallComplianceTaskStatus</code> is a string value that
	// corresponds to one of the `PbmComplianceResult`.
	// `PbmComplianceResult.complianceTaskStatus` values.
	OverallComplianceTaskStatus string `xml:"overallComplianceTaskStatus,omitempty" json:"overallComplianceTaskStatus,omitempty"`
	// Individual compliance results that make up the rollup.
	Result []PbmComplianceResult `xml:"result,omitempty" json:"result,omitempty"`
	// This property is set if the overall compliance task fails with some error.
	//
	// This
	// property indicates the causes of error. If there are multiple failures, it stores
	// these failure in this array.
	ErrorCause []types.LocalizedMethodFault `xml:"errorCause,omitempty" json:"errorCause,omitempty"`
	// Deprecated as of vSphere 2016, use
	// `PbmRollupComplianceResult.overallComplianceStatus`
	// to know if profile mismatch has occurred. If
	// overallComplianceStatus value is outOfDate, it means
	// profileMismatch has occurred.
	//
	// True if and only if `PbmComplianceResult`.
	//
	// `PbmComplianceResult.mismatch` is true for at least one
	// entity in the rollup compliance check.
	ProfileMismatch bool `xml:"profileMismatch" json:"profileMismatch"`
}

func init() {
	types.Add("pbm:PbmRollupComplianceResult", reflect.TypeOf((*PbmRollupComplianceResult)(nil)).Elem())
}

// The `PbmServerObjectRef` data object identifies
// a virtual machine,
// virtual disk attached to a virtual machine,
// a first class storage object
// or a datastore.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmServerObjectRef struct {
	types.DynamicData

	// Type of vSphere Server object.
	//
	// The value of the <code>objectType</code> string
	// corresponds to one of the `PbmObjectType_enum`
	// enumerated type values.
	ObjectType string `xml:"objectType" json:"objectType"`
	// Unique identifier for the object.
	//
	// The value of <code>key</code> depends
	// on the <code>objectType</code>.
	//
	// <table border="1"cellpadding="5">
	// <tr><td>`*PbmObjectType**</td><td>*`key value**</td></tr>
	// <tr><td>virtualMachine</td><td>_virtual-machine-MOR_</td></tr>
	// <tr><td>virtualDiskId</td>
	// <td>_virtual-disk-MOR_:_VirtualDisk.key_</td></tr>
	// <tr><td>datastore</td><td>_datastore-MOR_</td></tr>
	// <tr><td colspan="2"align="right">MOR = ManagedObjectReference</td></tr>
	// </table>
	Key string `xml:"key" json:"key"`
	// vCenter Server UUID; the <code>ServiceContent.about.instanceUuid</code>
	// property in the vSphere API.
	ServerUuid string `xml:"serverUuid,omitempty" json:"serverUuid,omitempty"`
}

func init() {
	types.Add("pbm:PbmServerObjectRef", reflect.TypeOf((*PbmServerObjectRef)(nil)).Elem())
}

// The `PbmServiceInstanceContent` data object defines properties for the
// `PbmServiceInstance` managed object.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmServiceInstanceContent struct {
	types.DynamicData

	// Contains information that identifies the Storage Policy service.
	AboutInfo PbmAboutInfo `xml:"aboutInfo" json:"aboutInfo"`
	// For internal use.
	//
	// Refers instance of `PbmSessionManager`.
	SessionManager types.ManagedObjectReference `xml:"sessionManager" json:"sessionManager"`
	// For internal use.
	//
	// Refers instance of `PbmCapabilityMetadataManager`.
	CapabilityMetadataManager types.ManagedObjectReference `xml:"capabilityMetadataManager" json:"capabilityMetadataManager"`
	// Provides access to the Storage Policy ProfileManager.
	//
	// Refers instance of `PbmProfileProfileManager`.
	ProfileManager types.ManagedObjectReference `xml:"profileManager" json:"profileManager"`
	// Provides access to the Storage Policy ComplianceManager.
	//
	// Refers instance of `PbmComplianceManager`.
	ComplianceManager types.ManagedObjectReference `xml:"complianceManager" json:"complianceManager"`
	// Provides access to the Storage Policy PlacementSolver.
	//
	// Refers instance of `PbmPlacementSolver`.
	PlacementSolver types.ManagedObjectReference `xml:"placementSolver" json:"placementSolver"`
	// Provides access to the Storage Policy ReplicationManager.
	//
	// Refers instance of `PbmReplicationManager`.
	ReplicationManager *types.ManagedObjectReference `xml:"replicationManager,omitempty" json:"replicationManager,omitempty"`
}

func init() {
	types.Add("pbm:PbmServiceInstanceContent", reflect.TypeOf((*PbmServiceInstanceContent)(nil)).Elem())
}

type PbmUpdate PbmUpdateRequestType

func init() {
	types.Add("pbm:PbmUpdate", reflect.TypeOf((*PbmUpdate)(nil)).Elem())
}

// The parameters of `PbmProfileProfileManager.PbmUpdate`.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmUpdateRequestType struct {
	This types.ManagedObjectReference `xml:"_this" json:"_this"`
	// Profile identifier.
	ProfileId PbmProfileId `xml:"profileId" json:"profileId"`
	// Capability-based update specification.
	UpdateSpec PbmCapabilityProfileUpdateSpec `xml:"updateSpec" json:"updateSpec"`
}

func init() {
	types.Add("pbm:PbmUpdateRequestType", reflect.TypeOf((*PbmUpdateRequestType)(nil)).Elem())
}

type PbmUpdateResponse struct {
}

// Information about a supported data service provided using
// vSphere APIs for IO Filtering (VAIO) data service provider.
//
// This structure may be used only with operations rendered under `/pbm`.
type PbmVaioDataServiceInfo struct {
	PbmLineOfServiceInfo
}

func init() {
	types.Add("pbm:PbmVaioDataServiceInfo", reflect.TypeOf((*PbmVaioDataServiceInfo)(nil)).Elem())
}

type VersionURI string

func init() {
	types.Add("pbm:versionURI", reflect.TypeOf((*VersionURI)(nil)).Elem())
}
