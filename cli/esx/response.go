// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package esx

import (
	"io"

	"github.com/vmware/govmomi/vim25/xml"
)

type Values map[string][]string

type Response struct {
	Info   *CommandInfoMethod `json:"info"`
	Values []Values           `json:"values"`
	String string             `json:"string"`
	Kind   string             `json:"-"`
}

func (v Values) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		t, err := d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		if s, ok := t.(xml.StartElement); ok {
			t, err = d.Token()
			if err != nil {
				return err
			}

			key := s.Name.Local
			var val string
			if c, ok := t.(xml.CharData); ok {
				val = string(c)
			}
			v[key] = append(v[key], val)
		}
	}
}

func (s Values) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tokens := []xml.Token{start}

	for key, val := range s {
		field := xml.StartElement{Name: xml.Name{Local: key}}
		for _, v := range val {
			tokens = append(tokens, field, xml.CharData(v), field.End())
		}
	}

	tokens = append(tokens, start.End())

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	return nil
}

func (v Values) Value(name string) string {
	if val, ok := v[name]; ok {
		if len(val) != 0 {
			return val[0]
		}
	}
	return ""
}

func (r *Response) Type(start xml.StartElement) string {
	for _, a := range start.Attr {
		if a.Name.Local == "type" {
			return a.Value
		}
	}
	return ""
}

func (r *Response) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	stype := r.Type(start)

	if stype != "ArrayOfDataObject" {
		switch stype {
		case "xsd:string", "xsd:boolean", "xsd:long":
			return d.DecodeElement(&r.String, &start)
		}
		v := Values{}
		if err := d.DecodeElement(&v, &start); err != nil {
			return err
		}
		r.Values = append(r.Values, v)
		return nil
	}

	for {
		t, err := d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		if s, ok := t.(xml.StartElement); ok {
			if s.Name.Local == "DataObject" {
				v := Values{}
				if err := d.DecodeElement(&v, &s); err != nil {
					return err
				}
				r.Values = append(r.Values, v)
			}
		}
	}
}

func (r *Response) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {
	kind := "ArrayOfDataObject"
	native := r.String != ""
	if native {
		kind = "xsd:" + r.Kind
	}

	start := xml.StartElement{
		Name: xml.Name{
			Space: "urn:vim25",
			Local: "obj",
		},
		Attr: []xml.Attr{
			{
				Name:  xml.Name{Local: "xmlns:xsd"},
				Value: "http://www.w3.org/2001/XMLSchema",
			},
			{
				Name:  xml.Name{Local: "xmlns:xsi"},
				Value: "http://www.w3.org/2001/XMLSchema-instance",
			},
			{
				Name:  xml.Name{Local: "xsi:type"},
				Value: kind,
			},
		},
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	var err error
	if native {
		err = e.EncodeToken(xml.CharData(r.String))
	} else {
		obj := xml.StartElement{
			Name: xml.Name{Local: "DataObject"},
			Attr: []xml.Attr{{
				Name:  xml.Name{Local: "xsi:type"},
				Value: r.Kind,
			}},
		}
		err = e.EncodeElement(r.Values, obj)
	}

	if err != nil {
		return err
	}

	return e.EncodeToken(start.End())
}
