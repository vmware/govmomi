/*
Copyright (c) 2023-2023 VMware, Inc. All Rights Reserved.

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

package importx_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/vmware/govmomi/govc/importx"
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
	var opts importx.Options
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
