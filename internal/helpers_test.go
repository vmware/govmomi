/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.
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

package internal_test

import (
	"testing"

	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/simulator/esx"
)

func TestHostSystemManagementIPs(t *testing.T) {
	ips := internal.HostSystemManagementIPs(esx.HostSystem.Config.VirtualNicManagerInfo.NetConfig)

	if len(ips) != 1 {
		t.Fatalf("no mgmt ip found")
	}
	if ips[0].String() != "127.0.0.1" {
		t.Fatalf("Expected management ip %s, got %s", "127.0.0.1", ips[0].String())
	}
}
