// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package esx

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

func load(name string) *Response {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var b Response

	dec := xml.NewDecoder(f)
	dec.TypeFunc = types.TypeFunc()
	if err := dec.Decode(&b); err != nil {
		panic(err)
	}

	return &b
}

func TestSystemHostnameGetResponse(t *testing.T) {
	res := load("fixtures/system_hostname_get.xml")

	expect := []Values{
		{
			"DomainName":               {"localdomain"},
			"FullyQualifiedDomainName": {"esxbox.localdomain"},
			"HostName":                 {"esxbox"},
		},
	}

	if !reflect.DeepEqual(res.Values, expect) {
		t.Errorf("%s != %s", res.Values, expect)
	}
}

func TestNetworkVmList(t *testing.T) {
	res := load("fixtures/network_vm_list.xml")

	expect := []Values{
		{
			"Name":     {"foo"},
			"Networks": {"VM Network", "dougm"},
			"NumPorts": {"2"},
			"WorldID":  {"98842"},
		},
		{
			"Name":     {"bar"},
			"Networks": {"VM Network"},
			"NumPorts": {"1"},
			"WorldID":  {"236235"},
		},
	}

	if !reflect.DeepEqual(res.Values, expect) {
		t.Errorf("%s != %s", res.Values, expect)
	}
}

func TestNetworkVmPortList(t *testing.T) {
	r := load("fixtures/network_vm_port_list.xml")

	expect := []Values{
		{
			"IPAddress":    {"192.168.247.149"},
			"MACAddress":   {"00:0c:29:12:b2:cf"},
			"PortID":       {"33554438"},
			"Portgroup":    {"VM Network"},
			"TeamUplink":   {"vmnic0"},
			"UplinkPortID": {"33554434"},
			"vSwitch":      {"vSwitch0"},
			"DVPortID":     {""},
		},
	}

	if !reflect.DeepEqual(r.Values, expect) {
		t.Errorf("%s != %s", r.Values, expect)
	}
}

func TestMarshalResponse(t *testing.T) {
	values := []Values{
		{
			"IPAddress":    {"192.168.247.149"},
			"MACAddress":   {"00:0c:29:12:b2:cf"},
			"PortID":       {"33554438"},
			"Portgroup":    {"VM Network"},
			"TeamUplink":   {"vmnic0"},
			"UplinkPortID": {"33554434"},
			"vSwitch":      {"vSwitch0"},
			"DVPortID":     {""},
			"Tags":         {"one", "two"},
		},
	}

	{
		res := &Response{
			Values: values,
		}

		var out bytes.Buffer

		fmt.Fprint(&out, xml.Header)
		e := xml.NewEncoder(&out)
		e.Indent("", "  ")
		err := e.Encode(res)
		if err != nil {
			t.Fatal(err)
		}
		_ = e.Flush()

		t.Log(out.String())
	}

	{
		res := &Response{
			String: "time",
		}

		var out bytes.Buffer

		fmt.Fprint(&out, xml.Header)
		e := xml.NewEncoder(&out)
		e.Indent("", "  ")
		err := e.Encode(res)
		if err != nil {
			t.Fatal(err)
		}
		_ = e.Flush()

		t.Log(out.String())
	}
}
