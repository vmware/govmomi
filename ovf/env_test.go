// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import "testing"

func testEnv() Env {
	return Env{
		EsxID: "vm moref",
		Platform: &PlatformSection{
			Kind:    "VMware vCenter Server",
			Version: "5.5.0",
			Vendor:  "VMware, Inc.",
			Locale:  "US",
		},
		Property: &PropertySection{
			Properties: []EnvProperty{
				{"foo", "bar"},
				{"ham", "eggs"}}},
	}
}

func TestMarshalEnv(t *testing.T) {
	env := testEnv()

	xenv, err := env.Marshal()
	if err != nil {
		t.Fatalf("error marshalling environment %s", err)
	}
	if len(xenv) < 1 {
		t.Fatal("marshalled document is empty")
	}
}

func TestMarshalManualEnv(t *testing.T) {
	env := testEnv()

	xenv := env.MarshalManual()
	if len(xenv) < 1 {
		t.Fatal("marshal document is empty")
	}
}
