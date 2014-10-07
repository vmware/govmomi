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

package govmomi

import (
	"reflect"
	"testing"

	"github.com/vmware/govmomi/test"
	"github.com/vmware/govmomi/vim25/mo"
)

func TestSearch(t *testing.T) {
	u := test.URL()
	if u == nil {
		t.SkipNow()
	}

	c, err := NewClient(*u, true)
	if err != nil {
		t.Error(err)
	}

	s := c.SearchIndex()

	ref, err := s.FindChild(c.RootFolder(), "ha-datacenter")
	if err != nil {
		t.Fatal(err)
	}

	dc, ok := ref.(*Datacenter)
	if !ok {
		t.Errorf("Expected Datacenter: %#v", ref)
	}

	folders, err := dc.Folders()
	if err != nil {
		t.Fatal(err)
	}

	ref, err = s.FindChild(folders.DatastoreFolder, "datastore1")
	if err != nil {
		t.Fatal(err)
	}

	_, ok = ref.(*Datastore)
	if !ok {
		t.Errorf("Expected Datastore: %#v", ref)
	}

	ref, err = s.FindByInventoryPath("/ha-datacenter/network/VM Network")
	if err != nil {
		t.Fatal(err)
	}

	_, ok = ref.(*Network)
	if !ok {
		t.Errorf("Expected Network: %#v", ref)
	}

	crs, err := folders.HostFolder.Children()
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
	if len(crs) != 0 {
		var cr mo.ComputeResource
		ref = crs[0]
		err = c.Properties(ref.Reference(), []string{"host"}, &cr)
		if err != nil {
			t.Fatal(err)
		}

		var host mo.HostSystem
		ref = NewHostSystem(c, cr.Host[0])
		err = c.Properties(ref.Reference(), []string{"name", "hardware"}, &host)
		if err != nil {
			t.Fatal(err)
		}

		shost, err := s.FindByDnsName(dc, host.Name, false) // TODO: get name/ip from nic manager
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(ref, shost) {
			t.Errorf("%#v != %#v\n", ref, shost)
		}

		shost, err = s.FindByUuid(dc, host.Hardware.SystemInfo.Uuid, false)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(ref, shost) {
			t.Errorf("%#v != %#v\n", ref, shost)
		}
	}

	vms, err := folders.VmFolder.Children()
	if err != nil {
		t.Fatal(err)
	}
	if len(vms) != 0 {
		var vm mo.VirtualMachine
		ref = vms[0]
		err = c.Properties(ref.Reference(), []string{"config", "guest"}, &vm)
		if err != nil {
			t.Fatal(err)
		}
		svm, err := s.FindByDatastorePath(dc, vm.Config.Files.VmPathName)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(ref, svm) {
			t.Errorf("%#v != %#v\n", ref, svm)
		}

		svm, err = s.FindByUuid(dc, vm.Config.Uuid, true)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(ref, svm) {
			t.Errorf("%#v != %#v\n", ref, svm)
		}

		if vm.Guest.HostName != "" {
			svm, err := s.FindByDnsName(dc, vm.Guest.HostName, true)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(ref, svm) {
				t.Errorf("%#v != %#v\n", ref, svm)
			}
		}

		if vm.Guest.IpAddress != "" {
			svm, err := s.FindByIp(dc, vm.Guest.IpAddress, true)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(ref, svm) {
				t.Errorf("%#v != %#v\n", ref, svm)
			}
		}
	}
}
