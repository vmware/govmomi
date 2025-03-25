// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import (
	"bytes"
	"fmt"

	"github.com/vmware/govmomi/vim25/xml"
)

const (
	ovfEnvHeader = `<Environment
		xmlns="http://schemas.dmtf.org/ovf/environment/1"
		xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
		xmlns:oe="http://schemas.dmtf.org/ovf/environment/1"
		xmlns:ve="http://www.vmware.com/schema/ovfenv"
		oe:id=""
		ve:esxId="%s">`
	ovfEnvPlatformSection = `<PlatformSection>
		<Kind>%s</Kind>
		<Version>%s</Version>
		<Vendor>%s</Vendor>
		<Locale>%s</Locale>
		</PlatformSection>`
	ovfEnvPropertyHeader = `<PropertySection>`
	ovfEnvPropertyEntry  = `<Property oe:key="%s" oe:value="%s"/>`
	ovfEnvPropertyFooter = `</PropertySection>`
	ovfEnvFooter         = `</Environment>`
)

type Env struct {
	XMLName xml.Name `xml:"http://schemas.dmtf.org/ovf/environment/1 Environment" json:"xmlName"`
	ID      string   `xml:"id,attr" json:"id"`
	EsxID   string   `xml:"http://www.vmware.com/schema/ovfenv esxId,attr" json:"esxID"`

	Platform *PlatformSection `xml:"PlatformSection" json:"platformSection,omitempty"`
	Property *PropertySection `xml:"PropertySection" json:"propertySection,omitempty"`
}

type PlatformSection struct {
	Kind    string `xml:"Kind" json:"kind,omitempty"`
	Version string `xml:"Version" json:"version,omitempty"`
	Vendor  string `xml:"Vendor" json:"vendor,omitempty"`
	Locale  string `xml:"Locale" json:"locale,omitempty"`
}

type PropertySection struct {
	Properties []EnvProperty `xml:"Property" json:"property,omitempty"`
}

type EnvProperty struct {
	Key   string `xml:"key,attr" json:"key"`
	Value string `xml:"value,attr" json:"value,omitempty"`
}

// Marshal marshals Env to xml by using xml.Marshal.
func (e Env) Marshal() (string, error) {
	x, err := xml.Marshal(e)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s", xml.Header, x), nil
}

// MarshalManual manually marshals Env to xml suitable for a vApp guest.
// It exists to overcome the lack of expressiveness in Go's XML namespaces.
func (e Env) MarshalManual() string {
	var buffer bytes.Buffer

	buffer.WriteString(xml.Header)
	buffer.WriteString(fmt.Sprintf(ovfEnvHeader, e.EsxID))
	buffer.WriteString(fmt.Sprintf(ovfEnvPlatformSection, e.Platform.Kind, e.Platform.Version, e.Platform.Vendor, e.Platform.Locale))

	buffer.WriteString(fmt.Sprint(ovfEnvPropertyHeader))
	for _, p := range e.Property.Properties {
		buffer.WriteString(fmt.Sprintf(ovfEnvPropertyEntry, p.Key, p.Value))
	}
	buffer.WriteString(fmt.Sprint(ovfEnvPropertyFooter))

	buffer.WriteString(fmt.Sprint(ovfEnvFooter))

	return buffer.String()
}
