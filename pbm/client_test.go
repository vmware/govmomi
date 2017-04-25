/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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

package pbm

import (
	"context"
	"os"
	"reflect"
	"sort"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/pbm/types"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	vim "github.com/vmware/govmomi/vim25/types"
)

func TestClient(t *testing.T) {
	url := os.Getenv("GOVMOMI_PBM_URL")
	if url == "" {
		t.SkipNow()
	}

	clusterName := os.Getenv("GOVMOMI_PBM_CLUSTER")

	u, err := soap.ParseURL(url)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	c, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		t.Fatal(err)
	}

	pc, err := NewClient(ctx, c.Client)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("PBM version=%s", pc.ServiceContent.AboutInfo.Version)

	rtype := types.PbmProfileResourceType{
		ResourceType: string(types.PbmProfileResourceTypeEnumSTORAGE),
	}

	category := types.PbmProfileCategoryEnumREQUIREMENT

	ids, err := pc.QueryProfile(ctx, rtype, string(category))
	if err != nil {
		t.Fatal(err)
	}

	var qids []string

	for _, id := range ids {
		qids = append(qids, id.UniqueId)
	}

	var cids []string

	policies, err := pc.RetrieveContent(ctx, ids)
	if err != nil {
		t.Fatal(err)
	}

	for i := range policies {
		profile := policies[i].GetPbmProfile()
		cids = append(cids, profile.ProfileId.UniqueId)
	}

	sort.Strings(qids)
	sort.Strings(cids)

	if !reflect.DeepEqual(qids, cids) {
		t.Error("ids mismatch")
	}

	root := c.ServiceContent.RootFolder
	var datastores []vim.ManagedObjectReference
	var kind []string

	if clusterName == "" {
		kind = []string{"Datastore"}
	} else {
		kind = []string{"ClusterComputeResource"}
	}

	m := view.NewManager(c.Client)

	v, err := m.CreateContainerView(ctx, root, kind, true)
	if err != nil {
		t.Fatal(err)
	}

	if clusterName == "" {
		datastores, err = v.Find(ctx, kind, nil)
		if err != nil {
			t.Fatal(err)
		}
	} else {
		var cluster mo.ClusterComputeResource

		err = v.RetrieveWithFilter(ctx, kind, []string{"datastore"}, &cluster, property.Filter{"name": clusterName})
		if err != nil {
			t.Fatal(err)
		}

		datastores = cluster.Datastore
	}

	_ = v.Destroy(ctx)

	t.Logf("checking %d datatores", len(datastores))

	var hubs []types.PbmPlacementHub

	for _, ds := range datastores {
		hubs = append(hubs, types.PbmPlacementHub{
			HubType: ds.Type,
			HubId:   ds.Value,
		})
	}

	var req []types.BasePbmPlacementRequirement

	for _, id := range ids {
		req = append(req, &types.PbmPlacementCapabilityProfileRequirement{
			ProfileId: id,
		})
	}

	res, err := pc.CheckRequirements(ctx, hubs, nil, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%d results", len(res))
}
