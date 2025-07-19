// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/vmware/govmomi/test"
	"github.com/vmware/govmomi/vim25/mo"
)

func TestSearch(t *testing.T) {
	c := test.NewAuthenticatedClient(t)
	s := NewSearchIndex(c)

	ref, err := s.FindChild(context.Background(), NewRootFolder(c), "ha-datacenter")
	if err != nil {
		t.Fatal(err)
	}

	dc, ok := ref.(*Datacenter)
	if !ok {
		t.Errorf("Expected Datacenter: %#v", ref)
	}

	folders, err := dc.Folders(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	ref, err = s.FindChild(context.Background(), folders.DatastoreFolder, "datastore1")
	if err != nil {
		t.Fatal(err)
	}

	_, ok = ref.(*Datastore)
	if !ok {
		t.Errorf("Expected Datastore: %#v", ref)
	}

	ref, err = s.FindByInventoryPath(context.Background(), "/ha-datacenter/network/VM Network")
	if err != nil {
		t.Fatal(err)
	}

	network, ok := ref.(*Network)
	if !ok {
		t.Errorf("Expected Network: %#v", ref)
	}
	if network.GetInventoryPath() != "/ha-datacenter/network/VM Network" {
		t.Errorf("%q != %q\n", network.GetInventoryPath(), "/ha-datacenter/network/VM Network")
	}

	crs, err := folders.HostFolder.Children(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(crs) != 0 {
		var cr mo.ComputeResource
		ref = crs[0]
		err = s.Properties(context.Background(), ref.Reference(), []string{"host"}, &cr)
		if err != nil {
			t.Fatal(err)
		}

		var host mo.HostSystem
		ref = NewHostSystem(c, cr.Host[0])
		err = s.Properties(context.Background(), ref.Reference(), []string{"name", "hardware", "config"}, &host)
		if err != nil {
			t.Fatal(err)
		}

		dnsConfig := host.Config.Network.DnsConfig.GetHostDnsConfig()
		dnsName := fmt.Sprintf("%s.%s", dnsConfig.HostName, dnsConfig.DomainName)
		shost, err := s.FindByDnsName(context.Background(), dc, dnsName, false)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(ref, shost) {
			t.Errorf("%#v != %#v\n", ref, shost)
		}

		shost, err = s.FindByUuid(context.Background(), dc, host.Hardware.SystemInfo.Uuid, false, nil)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(ref, shost) {
			t.Errorf("%#v != %#v\n", ref, shost)
		}

		shosts, err := s.FindAllByUuid(context.Background(), dc, host.Hardware.SystemInfo.Uuid, false, nil)
		if err != nil {
			t.Fatal(err)
		}
		if len(shosts) != 1 {
			t.Errorf("len(shosts) != 1: %d\n", len(shosts))
		}
		if !reflect.DeepEqual(ref, shosts[0]) {
			t.Errorf("%#v != %#v\n", ref, shosts[0])
		}
	}

	vms, err := folders.VmFolder.Children(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(vms) != 0 {
		var vm mo.VirtualMachine
		ref = vms[0]
		err = s.Properties(context.Background(), ref.Reference(), []string{"config", "guest"}, &vm)
		if err != nil {
			t.Fatal(err)
		}
		svm, err := s.FindByDatastorePath(context.Background(), dc, vm.Config.Files.VmPathName)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(ref, svm) {
			t.Errorf("%#v != %#v\n", ref, svm)
		}

		svm, err = s.FindByUuid(context.Background(), dc, vm.Config.Uuid, true, nil)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(ref, svm) {
			t.Errorf("%#v != %#v\n", ref, svm)
		}

		svms, err := s.FindAllByUuid(context.Background(), dc, vm.Config.Uuid, true, nil)
		if err != nil {
			t.Fatal(err)
		}
		if len(svms) != 1 {
			t.Errorf("len(svms) != 1: %d\n", len(svms))
		}
		if !reflect.DeepEqual(ref, svms[0]) {
			t.Errorf("%#v != %#v\n", ref, svms[0])
		}

		if vm.Guest.HostName != "" {
			svm, err := s.FindByDnsName(context.Background(), dc, vm.Guest.HostName, true)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(ref, svm) {
				t.Errorf("%#v != %#v\n", ref, svm)
			}
		}

		if vm.Guest.IpAddress != "" {
			svm, err := s.FindByIp(context.Background(), dc, vm.Guest.IpAddress, true)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(ref, svm) {
				t.Errorf("%#v != %#v\n", ref, svm)
			}
		}
	}
}
