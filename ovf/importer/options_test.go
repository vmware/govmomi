// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package importer_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/vmware/govmomi/ovf/importer"
)

func TestDecodeOptions(t *testing.T) {
	spec := []byte(`{
  "DiskProvisioning": "flat",
  "IPAllocationPolicy": "dhcpPolicy",
  "IPProtocol": "IPv4",
  "PropertyMapping": [
    {
      "Key": "ntp_server",
      "Value": "time.vmware.com"
    },
    {
      "key": "enable_ssh",
      "value": "True"
    }
  ],
  "NetworkMapping": [
    {
      "Name": "VM Network",
      "Network": ""
    }
  ],
  "MarkAsTemplate": false,
  "PowerOn": false,
  "InjectOvfEnv": false,
  "WaitForIP": false,
  "Name": null
}
`)
	var opts importer.Options
	err := json.NewDecoder(bytes.NewReader(spec)).Decode(&opts)
	if err != nil {
		t.Fatal(err)
	}

	// KeyValue are case insensitive
	for i, p := range opts.PropertyMapping {
		if p.Key == "" {
			t.Errorf("empty PropertyMapping[%d].Key", i)
		}
		if p.Value == "" {
			t.Errorf("empty PropertyMapping[%d].Value", i)
		}
	}
}
