/*
Copyright (c) 2015-2024 VMware, Inc. All Rights Reserved.

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

package ovf

import (
	"fmt"
)

// Envelope is defined according to
// https://www.dmtf.org/sites/default/files/standards/documents/DSP0243_2.1.1.pdf.
//
// Section 9 describes the parent/child relationships.
//
// A VirtualSystem may have zero or more VirtualHardware sections.
type Envelope struct {
	References []File `xml:"References>File" json:"references,omitempty"`

	// Package level meta-data
	Disk             *DiskSection             `xml:"DiskSection,omitempty" json:"diskSection,omitempty"`
	Network          *NetworkSection          `xml:"NetworkSection,omitempty" json:"networkSection,omitempty"`
	DeploymentOption *DeploymentOptionSection `xml:"DeploymentOptionSection,omitempty" json:"deploymentOptionSection,omitempty"`

	// Content: A VirtualSystem or a VirtualSystemCollection
	VirtualSystem           *VirtualSystem           `xml:"VirtualSystem,omitempty" json:"virtualSystem,omitempty"`
	VirtualSystemCollection *VirtualSystemCollection `xml:"VirtualSystemCollection,omitempty" json:"virtualSystemCollection,omitempty"`
}

type VirtualSystem struct {
	Content

	Annotation      *AnnotationSection       `xml:"AnnotationSection,omitempty" json:"annotationSection,omitempty"`
	Product         []ProductSection         `xml:"ProductSection,omitempty" json:"productSection,omitempty"`
	Eula            []EulaSection            `xml:"EulaSection,omitempty" json:"eulaSection,omitempty"`
	OperatingSystem *OperatingSystemSection  `xml:"OperatingSystemSection,omitempty" json:"operatingSystemSection,omitempty"`
	VirtualHardware []VirtualHardwareSection `xml:"VirtualHardwareSection,omitempty" json:"virtualHardwareSection,omitempty"`
}

type VirtualSystemCollection struct {
	Content

	// Collection level meta-data
	ResourceAllocation *ResourceAllocationSection `xml:"ResourceAllocationSection,omitempty" json:"resourceAllocationSection,omitempty"`
	Annotation         *AnnotationSection         `xml:"AnnotationSection,omitempty" json:"annotationSection,omitempty"`
	Product            []ProductSection           `xml:"ProductSection,omitempty" json:"productSection,omitempty"`
	Eula               []EulaSection              `xml:"EulaSection,omitempty" json:"eulaSection,omitempty"`

	// Content: One or more VirtualSystems
	VirtualSystem []VirtualSystem `xml:"VirtualSystem,omitempty" json:"virtualSystem,omitempty"`
}

type File struct {
	ID          string  `xml:"id,attr" json:"id,omitempty"`
	Href        string  `xml:"href,attr" json:"href,omitempty"`
	Size        uint    `xml:"size,attr" json:"size,omitempty"`
	Compression *string `xml:"compression,attr" json:"compression,omitempty"`
	ChunkSize   *int    `xml:"chunkSize,attr" json:"chunkSize,omitempty"`
}

type Content struct {
	ID   string  `xml:"id,attr" json:"id,omitempty"`
	Info string  `xml:"Info" json:"info,omitempty"`
	Name *string `xml:"Name" json:"name,omitempty"`
}

type Section struct {
	Required *bool  `xml:"required,attr" json:"required,omitempty"`
	Info     string `xml:"Info" json:"info,omitempty"`
	Category string `xml:"Category" json:"category,omitempty"`
}

type AnnotationSection struct {
	Section

	Annotation string `xml:"Annotation" json:"annotation,omitempty"`
}

type ProductSection struct {
	Section

	Class    *string `xml:"class,attr" json:"class,omitempty"`
	Instance *string `xml:"instance,attr" json:"instance,omitempty"`

	Product     string     `xml:"Product" json:"product,omitempty"`
	Vendor      string     `xml:"Vendor" json:"vendor,omitempty"`
	Version     string     `xml:"Version" json:"version,omitempty"`
	FullVersion string     `xml:"FullVersion" json:"fullVersion,omitempty"`
	ProductURL  string     `xml:"ProductUrl" json:"productUrl,omitempty"`
	VendorURL   string     `xml:"VendorUrl" json:"vendorUrl,omitempty"`
	AppURL      string     `xml:"AppUrl" json:"appUrl,omitempty"`
	Property    []Property `xml:"Property" json:"property,omitempty"`
}

func (p ProductSection) Key(prop Property) string {
	// From OVF spec, section 9.5.1:
	// key-value-env = [class-value "."] key-value-prod ["." instance-value]

	k := prop.Key
	if p.Class != nil {
		k = fmt.Sprintf("%s.%s", *p.Class, k)
	}
	if p.Instance != nil {
		k = fmt.Sprintf("%s.%s", k, *p.Instance)
	}
	return k
}

type Property struct {
	Key              string  `xml:"key,attr" json:"key,omitempty"`
	Type             string  `xml:"type,attr" json:"type,omitempty"`
	Qualifiers       *string `xml:"qualifiers,attr" json:"qualifiers,omitempty"`
	UserConfigurable *bool   `xml:"userConfigurable,attr" json:"userConfigurable,omitempty"`
	Default          *string `xml:"value,attr" json:"default,omitempty"`
	Password         *bool   `xml:"password,attr" json:"password,omitempty"`
	Configuration    *string `xml:"configuration,attr" json:"configuration,omitempty"`

	Label       *string `xml:"Label" json:"label,omitempty"`
	Description *string `xml:"Description" json:"description,omitempty"`

	Values []PropertyConfigurationValue `xml:"Value" json:"value,omitempty"`
}

type PropertyConfigurationValue struct {
	Value         string  `xml:"value,attr" json:"value,omitempty"`
	Configuration *string `xml:"configuration,attr" json:"configuration,omitempty"`
}

type NetworkSection struct {
	Section

	Networks []Network `xml:"Network" json:"network,omitempty"`
}

type Network struct {
	Name string `xml:"name,attr" json:"name,omitempty"`

	Description string `xml:"Description" json:"description,omitempty"`
}

type DiskSection struct {
	Section

	Disks []VirtualDiskDesc `xml:"Disk" json:"disk,omitempty"`
}

type VirtualDiskDesc struct {
	DiskID                  string  `xml:"diskId,attr" json:"diskId,omitempty"`
	FileRef                 *string `xml:"fileRef,attr" json:"fileRef,omitempty"`
	Capacity                string  `xml:"capacity,attr" json:"capacity,omitempty"`
	CapacityAllocationUnits *string `xml:"capacityAllocationUnits,attr" json:"capacityAllocationUnits,omitempty"`
	Format                  *string `xml:"format,attr" json:"format,omitempty"`
	PopulatedSize           *int    `xml:"populatedSize,attr" json:"populatedSize,omitempty"`
	ParentRef               *string `xml:"parentRef,attr" json:"parentRef,omitempty"`
}

type OperatingSystemSection struct {
	Section

	ID      int16   `xml:"id,attr" json:"id"`
	Version *string `xml:"version,attr" json:"version,omitempty"`
	OSType  *string `xml:"osType,attr" json:"osType,omitempty"`

	Description *string `xml:"Description" json:"description,omitempty"`
}

type EulaSection struct {
	Section

	License string `xml:"License" json:"license,omitempty"`
}

type Config struct {
	Required *bool  `xml:"required,attr" json:"required,omitempty"`
	Key      string `xml:"key,attr" json:"key,omitempty"`
	Value    string `xml:"value,attr" json:"value,omitempty"`
}

type VirtualHardwareSection struct {
	Section

	ID        *string `xml:"id,attr" json:"id"`
	Transport *string `xml:"transport,attr" json:"transport,omitempty"`

	System      *VirtualSystemSettingData       `xml:"System" json:"system,omitempty"`
	Item        []ResourceAllocationSettingData `xml:"Item" json:"item,omitempty"`
	StorageItem []StorageAllocationSettingData  `xml:"StorageItem" json:"storageItem,omitempty"`
	Config      []Config                        `xml:"Config" json:"config,omitempty"`
	ExtraConfig []Config                        `xml:"ExtraConfig" json:"extraConfig,omitempty"`
}

type VirtualSystemSettingData struct {
	CIMVirtualSystemSettingData
}

type ResourceAllocationSettingData struct {
	CIMResourceAllocationSettingData

	Required       *bool           `xml:"required,attr" json:"required,omitempty"`
	Configuration  *string         `xml:"configuration,attr" json:"configuration,omitempty"`
	Bound          *string         `xml:"bound,attr" json:"bound,omitempty"`
	Config         []Config        `xml:"Config" json:"config,omitempty"`
	CoresPerSocket *CoresPerSocket `xml:"CoresPerSocket" json:"coresPerSocket,omitempty"`
}

type StorageAllocationSettingData struct {
	CIMStorageAllocationSettingData

	Required      *bool   `xml:"required,attr" json:"required,omitempty"`
	Configuration *string `xml:"configuration,attr" json:"configuration,omitempty"`
	Bound         *string `xml:"bound,attr" json:"bound,omitempty"`
}

type ResourceAllocationSection struct {
	Section

	Item []ResourceAllocationSettingData `xml:"Item" json:"item,omitempty"`
}

type DeploymentOptionSection struct {
	Section

	Configuration []DeploymentOptionConfiguration `xml:"Configuration" json:"configuration,omitempty"`
}

type DeploymentOptionConfiguration struct {
	ID      string `xml:"id,attr" json:"id"`
	Default *bool  `xml:"default,attr" json:"default,omitempty"`

	Label       string `xml:"Label" json:"label,omitempty"`
	Description string `xml:"Description" json:"description,omitempty"`
}

type CoresPerSocket struct {
	Required *bool `xml:"required,attr" json:"required,omitempty"`
	Value    int32 `xml:",chardata" json:"value,omitempty"`
}
