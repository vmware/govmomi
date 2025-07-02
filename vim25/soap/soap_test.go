// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package soap

import (
	"testing"

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
