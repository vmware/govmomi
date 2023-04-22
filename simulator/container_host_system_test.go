/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package simulator

import (
	"testing"

	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/types"
)

func TestSimHost(t *testing.T) {
	m := ESX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	hs := NewHostSystem(esx.HostSystem)
	if hs.Summary.Runtime != &hs.Runtime {
		t.Fatal("expected hs.Summary.Runtime == &hs.Runtime; got !=")
	}

	hs.HostSystem.Config.Option = append(hs.HostSystem.Config.Option, &types.OptionValue{Key: "RUN.container"})
	hs.HostSystem.Config.FileSystemVolume = nil

	hs.configure(SpoofContext(), types.HostConnectSpec{}, true)

	hs.sh.remove(SpoofContext())

}
