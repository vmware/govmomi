// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package soap

import (
	"bytes"
	"testing"

	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

func TestEmptyEnvelope(t *testing.T) {
	env := Envelope{}

	b, err := xml.Marshal(env)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	expected := `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/"></Envelope>`
	actual := string(b)
	if expected != actual {
		t.Fatalf("expected: %s, actual: %s", expected, actual)
	}
}

func TestNonEmptyHeader(t *testing.T) {
	env := Envelope{
		Header: &Header{
			ID: "foo",
		},
	}

	b, err := xml.Marshal(env)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	env = Envelope{}
	err = xml.Unmarshal(b, &env)
	if err != nil {
		t.Errorf("error: %s", err)
	}

	if env.Header.ID != "foo" {
		t.Errorf("ID=%s", env.Header.ID)
	}
}

func TestDecodeEnvelopeWithXsPrefix(t *testing.T) {
	data := []byte(`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
    <soap:Body>
        <ReconfigVM_Task xmlns="urn:vim25" xmlns:ns2="urn:vsan">
            <_this type="VirtualMachine">vm-1518</_this>
            <spec>
                <extraConfig>
                    <key>SET.guest.customizationInfo.customizationStatus</key>
                    <value xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xs="http://www.w3.org/2001/XMLSchema" xsi:type="xs:string">TEST_STATUS</value>
                </extraConfig>
            </spec>
        </ReconfigVM_Task>
    </soap:Body>
</soap:Envelope>`)

	var req types.ReconfigVM_Task
	var body struct {
		Req *types.ReconfigVM_Task `xml:"urn:vim25 ReconfigVM_Task"`
	}
	body.Req = &req
	env := Envelope{Body: &body}

	dec := xml.NewDecoder(bytes.NewReader(data))
	dec.TypeFunc = types.TypeFunc()
	err := dec.Decode(&env)
	if err != nil {
		t.Fatal(err)
	}

	if len(req.Spec.ExtraConfig) == 0 {
		t.Fatal("ExtraConfig is empty")
	}

	opt := req.Spec.ExtraConfig[0].GetOptionValue()
	if opt == nil {
		t.Fatal("ExtraConfig element is not an OptionValue")
	}

	if opt.Key != "SET.guest.customizationInfo.customizationStatus" {
		t.Errorf("Unexpected key: %s", opt.Key)
	}

	if opt.Value == nil {
		t.Fatal("Value is nil")
	}

	val, ok := opt.Value.(string)
	if !ok {
		t.Fatalf("Unexpected value type: %T (expected string)", opt.Value)
	}

	if val != "TEST_STATUS" {
		t.Errorf("Unexpected value: %s", val)
	}
}
