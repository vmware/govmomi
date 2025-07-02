// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package soap

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi/vim25/types"
)

type AcquireMksTicketBody struct {
	Req    *types.AcquireMksTicket         `xml:"urn:vim25 AcquireMksTicket,omitempty"`
	Res    *types.AcquireMksTicketResponse `xml:"AcquireMksTicketResponse,omitempty"`
	Fault_ *Fault                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *AcquireMksTicketBody) Fault() *Fault { return b.Fault_ }

func TestUnpack(t *testing.T) {
	req := &types.AcquireMksTicket{
		This: types.ManagedObjectReference{
			Type:  "Folder",
			Value: "group-d1",
		},
	}
	var reqBody HasFault = &AcquireMksTicketBody{
		Req: req,
	}
	this, method, params, err := unpackSOAPRequest(reqBody)
	if err != nil {
		t.Fatalf("Cannot unpack request %v", err)
	}
	if method != "AcquireMksTicket" {
		t.Fatalf("Expected 'AcquireMksTicket' methods but got: %v", method)
	}
	if this.Type != "Folder" {
		t.Fatalf("Expected 'Folder' type in this but got: %v", this.Type)
	}
	if this.Value != "group-d1" {
		t.Fatalf("Expected 'Folder' type in this but got: %v", this.Type)
	}
	if params != req {
		t.Fatal("Expected pointer to the req struct")
	}
}

func TestGetResultPtr(t *testing.T) {
	var respBody = &AcquireMksTicketBody{
		Res: &types.AcquireMksTicketResponse{
			Returnval: types.VirtualMachineMksTicket{},
		},
	}
	resPtr2, err := getSOAPResultPtr(respBody)
	if err != nil {
		t.Fatal(err)
	}
	if &respBody.Res.Returnval != resPtr2 {
		t.Fatal("Expected the original pointer to the result value but got different one.")
	}
}

func TestErrorUnmarshal(t *testing.T) {
	headers := http.Header{}
	headers.Set("content-type", "application/json")
	body := io.NopCloser(strings.NewReader(`{
		"_typeName": "InvalidLogin",
		"faultstring": "Cannot complete login due to an incorrect user name or password."
	}`))
	resp := &http.Response{
		StatusCode: 500,
		Header:     headers,
		Body:       body,
	}

	addr, _ := url.Parse("https://localhost/test")
	c := NewClient(addr, true)
	c.Namespace = "urn:vim25"

	var result any
	unmarshaler := c.responseUnmarshaler(&result)

	err := unmarshaler(resp)
	if err == nil {
		t.Error("Expected non nil error")
	}
	if !strings.HasPrefix(err.Error(), "ServerFaultCode: Cannot complete login") {
		t.Error("Unexpected error:", err)
	}
	if !IsSoapFault(err) {
		t.Error("Expected SOAP fault.")
	}
	fault := ToSoapFault(err)

	if fault.String != "Cannot complete login due to an incorrect user name or password." {
		t.Error("Unexpected faultstring message. Got:", fault.String)
	}
	if fault.Code != "ServerFaultCode" {
		t.Error("Unexpected faultcode. Got:", fault.Code)

	}
	// This may be incorrect. It is more of a speculation
	if fault.XMLName.Space != "urn:vim25" {
		t.Error("Expected vim25 namespace obtained from the client. Got:",
			fault.XMLName.Space)
	}
	if fault.XMLName.Local != "InvalidLoginFault" {
		t.Error("ExpectedInvalidLoginFault type. Got:", fault.XMLName.Local)
	}
	if _, ok := fault.Detail.Fault.(types.InvalidLogin); !ok {
		t.Error("Expected InvalidLogin nested fault type. Got:",
			reflect.TypeOf(fault.Detail.Fault).Name())
	}

}

func TestSuccessUnmarshal(t *testing.T) {
	headers := http.Header{}
	headers.Set("content-type", "application/json")
	body := io.NopCloser(strings.NewReader(`{
		"_typeName": "UserSession",
		"key": "527025b6-f0f4-144e-0d36-a7aaf2eb21da",
		"userName": "VSPHERE.LOCAL\\Administrator",
		"fullName": "Administrator vsphere.local",
		"loginTime": "2023-02-14T22:12:48.753169Z",
		"lastActiveTime": "2023-02-14T22:12:48.753169Z",
		"locale": "en",
		"messageLocale": "en",
		"extensionSession": false,
		"ipAddress": "10.93.153.94",
		"userAgent": "curl/7.86.0",
		"callCount": 0
	}`))
	resp := &http.Response{
		StatusCode:    200,
		Header:        headers,
		Body:          body,
		ContentLength: 100, // Fake length to avoid check for empty body
	}

	addr, _ := url.Parse("https://localhost/test")
	c := NewClient(addr, true)
	c.Namespace = "urn:vim25"

	var result any
	unmarshaler := c.responseUnmarshaler(&result)

	err := unmarshaler(resp)
	if err != nil {
		t.Error("Expected nil error")
	}
	var us types.UserSession
	var ok bool
	if us, ok = result.(types.UserSession); !ok {
		t.Error("Expected user session")
	}
	if us.UserName != "VSPHERE.LOCAL\\Administrator" {
		t.Error("Unexpected user name", us.UserName)
	}
}

type mockHTTP struct {
	impl func(*http.Request) (*http.Response, error)
}

func (m *mockHTTP) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.impl == nil {
		return nil, fmt.Errorf("mock HTTP implementation is not set")
	}
	return m.impl(req)
}

func TestFullRequestCycle(t *testing.T) {
	ctx := context.Background()
	req := AcquireMksTicketBody{
		Req: &types.AcquireMksTicket{
			This: types.ManagedObjectReference{
				Type:  "VirtualMachine",
				Value: "vm-42",
			},
		},
	}
	addr, _ := url.Parse("https://localhost/test")
	c := NewClient(addr, true)
	c.Namespace = "urn:vim25"
	c.Version = "8.0.0.1"
	c.Cookie = func() *HeaderElement {
		return &HeaderElement{Value: "(original)"}
	}
	c.UseJSON(true)

	c.Transport = &mockHTTP{
		impl: func(r *http.Request) (*http.Response, error) {
			if r.URL.String() != "https://localhost/test/vim25/8.0.0.1/VirtualMachine/vm-42/AcquireMksTicket" {
				t.Error("Unexpected url value:", r.URL)
			}

			if r.Method != "POST" {
				t.Error("Unexpected method value:", r.Method)
			}
			if r.Header.Get("vmware-api-session-id") != "(original)" {
				t.Error("Unexpected authentication value:", r.Header.Get("vmware-api-session-id"))
			}
			reqBytes, _ := io.ReadAll(r.Body)
			reqBody := string(reqBytes)
			assert.JSONEq(t, `{"_typeName":"AcquireMksTicket"}`, reqBody)

			if t.Failed() {
				t.FailNow()
			}
			body := `{
				"_typeName": "VirtualMachineMksTicket",
				"ticket": "ticket_value",
				"cfgFile": "file.ini",
				"host": "localhost",
				"port": 8080
			}`
			headers := http.Header{}
			headers.Set("content-type", "application/json")

			return &http.Response{
				StatusCode:    200,
				Header:        headers,
				Body:          io.NopCloser(strings.NewReader(body)),
				ContentLength: 100, // Fake length to avoid check for empty body
			}, nil
		},
	}

	res := AcquireMksTicketBody{}

	e := c.RoundTrip(ctx, &req, &res)

	if e != nil {
		t.Fatal(e)
	}
	if res.Fault_ != nil {
		t.Fatal(res.Fault_)
	}
	if res.Res.Returnval.Ticket != "ticket_value" || res.Res.Returnval.Port != 8080 {
		t.Fatalf("Return value not unmarshaled correctly '%#v'", res.Res.Returnval)
	}
}
