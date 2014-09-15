package esxcli

import (
	"io"
	"strings"

	"github.com/vmware/govmomi/vim25/xml"
)

type Request struct {
	Args  []string
	Flags map[string]string
}

type Values map[string][]string

type Response struct {
	Values []Values
}

func (r *Request) ParseArgs(args []string) {
	r.Flags = make(map[string]string)

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if strings.HasPrefix(arg, "--") {
			i++
			arg = strings.Replace(arg, "-", "", -1)
			r.Flags[arg] = args[i]
		} else {
			r.Args = append(r.Args, arg)
		}
	}
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
