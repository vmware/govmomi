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

package esxcli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

type esxcli struct {
	*flags.HostSystemFlag
}

type Request struct {
	SoapFault *soap.Fault
	Args      []string
	Flags     map[string]string
}

type Values map[string][]string

type Response struct {
	SoapFault *soap.Fault
	Values    []Values
}

func init() {
	cli.Register("host.esxcli", &esxcli{})
}

func (cmd *esxcli) Register(f *flag.FlagSet) {}

func (cmd *esxcli) Process() error { return nil }

func (cmd *esxcli) Run(f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return nil
	}

	host, err := cmd.HostSystem()
	if err != nil {
		return err
	}

	// TODO: we could use internal vmodl types to support VC
	if host.Value != "ha-host" {
		return errors.New("esxcli not supported through vCenter")
	}

	req := Request{}
	req.ParseArgs(f.Args())

	var res Response

	if err := c.RoundTrip(&req, &res); err != nil {
		return err
	}

	if len(res.Values) == 0 {
		return nil
	}

	var keys []string
	for key := range res.Values[0] {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	// TODO: OutputFlag / format options
	for i, rv := range res.Values {
		if i > 0 {
			fmt.Fprintln(tw)
			_ = tw.Flush()
		}
		for _, key := range keys {
			fmt.Fprintf(tw, "%s:\t%s\n", key, strings.Join(rv[key], ", "))
		}
	}

	return tw.Flush()
}

func (r *Request) ParseArgs(args []string) {
	r.Flags = make(map[string]string)

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if strings.HasPrefix(arg, "--") {
			i++
			r.Flags[arg[2:]] = args[i]
		} else {
			r.Args = append(r.Args, arg)
		}
	}

}

func (r *Request) Fault() *soap.Fault { return r.SoapFault }

func (r *Request) EncodeToken(e *xml.Encoder, t xml.Token) {
	if err := e.EncodeToken(t); err != nil {
		panic(err)
	}
}

func (r *Request) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	ns := r.Args[:len(r.Args)-1]

	ref := types.ManagedObjectReference{
		Type:  "VimEsxCLI" + strings.Join(ns, ""),
		Value: "ha-cli-handler-" + strings.Join(ns, "-"),
	}

	this := xml.StartElement{
		Name: xml.Name{Local: "_this"},
		Attr: []xml.Attr{
			xml.Attr{
				Name:  xml.Name{Local: "type"},
				Value: ref.Type,
			},
		},
	}

	cmd := xml.StartElement{
		Name: xml.Name{
			Space: "urn:vim25",
			Local: "VimEsxCLI" + strings.Join(r.Args, ""),
		},
	}

	r.EncodeToken(e, start)

	r.EncodeToken(e, cmd)

	r.EncodeToken(e, this)
	r.EncodeToken(e, xml.CharData(ref.Value))
	r.EncodeToken(e, xml.EndElement{Name: this.Name})

	for key, val := range r.Flags {
		opt := xml.StartElement{
			Name: xml.Name{Local: key},
		}

		r.EncodeToken(e, opt)
		r.EncodeToken(e, xml.CharData(val))
		r.EncodeToken(e, xml.EndElement{Name: opt.Name})
	}

	r.EncodeToken(e, xml.EndElement{Name: cmd.Name})

	r.EncodeToken(e, xml.EndElement{Name: start.Name})

	return nil
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

func (r *Response) Fault() *soap.Fault { return r.SoapFault }

func (r *Response) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		t, err := d.Token()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		if s, ok := t.(xml.StartElement); ok {
			switch strings.ToLower(s.Name.Local) {
			case "returnval":
				v := Values{}
				if err := d.DecodeElement(&v, &s); err != nil {
					return err
				}
				r.Values = append(r.Values, v)
			case "fault":
				r.SoapFault = new(soap.Fault)
				if err := d.DecodeElement(r.SoapFault, &s); err != nil {
					return err
				}
			}
		}
	}
}
