/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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

package object

import (
	"context"
	"testing"

	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/test"
)

func TestHostDatastoreSystemResignatureUnresolvedVmfsVolume(t *testing.T) {
	c := test.NewAuthenticatedClient(t)
	host := NewHostSystem(c, esx.HostSystem.Reference())

	ctx := context.Background()

	hds, err := host.ConfigManager().DatastoreSystem(ctx)
	if err != nil {
		t.Error(err)
	}
	hss, err := host.ConfigManager().StorageSystem(ctx)
	if err != nil {
		t.Error(err)
	}

	paths, err := hss.QueryUnresolvedVmfsVolumes(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(paths) == 0 {
		t.Skip("No unresolved vmfs volumes found")
	}

	task, err := hds.ResignatureUnresolvedVmfsVolumes(ctx, []string{paths[0].Extent[0].DevicePath})
	if err != nil {
		t.Fatal(err)
	}

	err = task.Wait(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
