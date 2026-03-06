// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import (
	"fmt"
	"strings"

	"github.com/vmware/govmomi/vim25/xml"
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
	Category string `xml:"Category,omitempty" json:"category,omitempty"`
}

type AnnotationSection struct {
	Section

	Annotation string `xml:"Annotation" json:"annotation,omitempty"`
}

// ProductSectionItem represents one element in the Category|Property sequence
// per DSP0243 9.5.1: "The set of Property elements grouped by a Category element
// is the sequence of Property elements following the Category element."
type ProductSectionItem struct {
	Category *string   `xml:"Category,omitempty" json:"category,omitempty"`
	Property *Property `xml:"Property,omitempty" json:"property,omitempty"`
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
	Property    []Property `xml:"Property" json:"property,omitempty"` // populated from Items for backward compat

	// Items preserves the order of Category and Property elements (DSP0243 9.5.1).
	// When unmarshaling, both Items and Property are populated.
	Items []ProductSectionItem `xml:"-" json:"items,omitempty"`
}

// UnmarshalXML decodes ProductSection and preserves the order of Category and
// Property elements per DSP0243 9.5.1 (category grouping).
func (p *ProductSection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, a := range start.Attr {
		switch a.Name.Local {
		case "class":
			p.Class = &a.Value
		case "instance":
			p.Instance = &a.Value
		case "required":
			v := strings.EqualFold(a.Value, "true")
			p.Required = &v
		}
	}
	var items []ProductSectionItem
	var properties []Property
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch t := tok.(type) {
		case xml.EndElement:
			p.Items = items
			p.Property = properties
			if p.Section.Category == "" && len(items) > 0 {
				for i := len(items) - 1; i >= 0; i-- {
					if items[i].Category != nil {
						p.Section.Category = *items[i].Category
						break
					}
				}
			}
			return nil
		case xml.StartElement:
			switch t.Name.Local {
			case "Info":
				var s string
				if err := d.DecodeElement(&s, &t); err != nil {
					return err
				}
				p.Section.Info = s
			case "Category":
				var s string
				if err := d.DecodeElement(&s, &t); err != nil {
					return err
				}
				s = strings.TrimSpace(s)
				p.Section.Category = s
				items = append(items, ProductSectionItem{Category: &s})
			case "Property":
				var prop Property
				if err := d.DecodeElement(&prop, &t); err != nil {
					return err
				}
				properties = append(properties, prop)
				items = append(items, ProductSectionItem{Property: &prop})
			case "Product":
				if err := d.DecodeElement(&p.Product, &t); err != nil {
					return err
				}
			case "Vendor":
				if err := d.DecodeElement(&p.Vendor, &t); err != nil {
					return err
				}
			case "Version":
				if err := d.DecodeElement(&p.Version, &t); err != nil {
					return err
				}
			case "FullVersion":
				if err := d.DecodeElement(&p.FullVersion, &t); err != nil {
					return err
				}
			case "ProductUrl":
				if err := d.DecodeElement(&p.ProductURL, &t); err != nil {
					return err
				}
			case "VendorUrl":
				if err := d.DecodeElement(&p.VendorURL, &t); err != nil {
					return err
				}
			case "AppUrl":
				if err := d.DecodeElement(&p.AppURL, &t); err != nil {
					return err
				}
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}
		}
	}
}

// PropertiesWithCategory returns properties in document order with the
// Category that groups each one (DSP0243 9.5.1: the Category element
// preceding the property in the sequence).
func (p ProductSection) PropertiesWithCategory() []struct {
	Category string
	Property Property
} {
	var out []struct {
		Category string
		Property Property
	}
	var curCategory string
	for _, item := range p.Items {
		if item.Category != nil {
			curCategory = *item.Category
		}
		if item.Property != nil {
			out = append(out, struct {
				Category string
				Property Property
			}{curCategory, *item.Property})
		}
	}
	return out
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
