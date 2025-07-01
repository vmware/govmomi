// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package pbm_test

import (
	"context"
	"os"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/pbm"
	pbmsim "github.com/vmware/govmomi/pbm/simulator"
	"github.com/vmware/govmomi/pbm/types"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
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

	pc, err := pbm.NewClient(ctx, c.Client)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("PBM version=%s", pc.ServiceContent.AboutInfo.Version)

	rtype := types.PbmProfileResourceType{
		ResourceType: string(types.PbmProfileResourceTypeEnumSTORAGE),
	}

	category := types.PbmProfileCategoryEnumREQUIREMENT

	// 1. Query all the profiles on the vCenter.
	ids, err := pc.QueryProfile(ctx, rtype, string(category))
	if err != nil {
		t.Fatal(err)
	}

	var qids []string

	for _, id := range ids {
		qids = append(qids, id.UniqueId)
	}

	var cids []string

	// 2. Retrieve the content of all profiles.
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

	// Check whether ids retrieved from QueryProfile and RetrieveContent
	// are identical.
	if !reflect.DeepEqual(qids, cids) {
		t.Error("ids mismatch")
	}

	// 3. Get list of datastores in a cluster if cluster name is specified.
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

		err = v.RetrieveWithFilter(ctx, kind, []string{"datastore"}, &cluster, property.Match{"name": clusterName})
		if err != nil {
			t.Fatal(err)
		}

		datastores = cluster.Datastore
	}

	_ = v.Destroy(ctx)

	t.Logf("checking %d datatores for compatibility results", len(datastores))

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

	// 4. Get the compatibility results for all the profiles on the vCenter.
	res, err := pc.CheckRequirements(ctx, hubs, nil, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("CheckRequirements results: %d", len(res))

	// user spec for the profile.
	// VSAN profile with 2 capability instances - hostFailuresToTolerate = 2, stripeWidth = 1
	pbmCreateSpecForVSAN := pbm.CapabilityProfileCreateSpec{
		Name:        "Kubernetes-VSAN-TestPolicy",
		Description: "VSAN Test policy create",
		Category:    string(types.PbmProfileCategoryEnumREQUIREMENT),
		CapabilityList: []pbm.Capability{
			{
				ID:        "hostFailuresToTolerate",
				Namespace: "VSAN",
				PropertyList: []pbm.Property{
					{
						ID:       "hostFailuresToTolerate",
						Value:    "2",
						DataType: "int",
					},
				},
			},
			{
				ID:        "stripeWidth",
				Namespace: "VSAN",
				PropertyList: []pbm.Property{
					{
						ID:       "stripeWidth",
						Value:    "1",
						DataType: "int",
					},
				},
			},
		},
	}

	// Create PBM capability spec for the above defined user spec.
	createSpecVSAN, err := pbm.CreateCapabilityProfileSpec(pbmCreateSpecForVSAN)
	if err != nil {
		t.Fatal(err)
	}

	// 5. Create SPBM VSAN profile.
	vsanProfileID, err := pc.CreateProfile(ctx, *createSpecVSAN)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("VSAN Profile: %q successfully created", vsanProfileID.UniqueId)

	// 6. Verify if profile created exists by issuing a RetrieveContent request.
	_, err = pc.RetrieveContent(ctx, []types.PbmProfileId{*vsanProfileID})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Profile: %q exists on vCenter", vsanProfileID.UniqueId)

	// 7. Get compatible datastores for the VSAN profile.
	compatibleDatastores := res.CompatibleDatastores()
	t.Logf("Found %d compatible-datastores for profile: %q", len(compatibleDatastores), vsanProfileID.UniqueId)

	// 8. Get non-compatible datastores for the VSAN profile.
	nonCompatibleDatastores := res.NonCompatibleDatastores()
	t.Logf("Found %d non-compatible datastores for profile: %q", len(nonCompatibleDatastores), vsanProfileID.UniqueId)

	// Check whether count of compatible and non-compatible datastores match the total number of datastores.
	if (len(nonCompatibleDatastores) + len(compatibleDatastores)) != len(datastores) {
		t.Error("datastore count mismatch")
	}

	// user spec for the profile.
	// VSAN profile with 2 capability instances - stripeWidth = 1 and an SIOC profile.
	pbmCreateSpecVSANandSIOC := pbm.CapabilityProfileCreateSpec{
		Name:        "Kubernetes-VSAN-SIOC-TestPolicy",
		Description: "VSAN-SIOC-Test policy create",
		Category:    string(types.PbmProfileCategoryEnumREQUIREMENT),
		CapabilityList: []pbm.Capability{
			{
				ID:        "stripeWidth",
				Namespace: "VSAN",
				PropertyList: []pbm.Property{
					{
						ID:       "stripeWidth",
						Value:    "1",
						DataType: "int",
					},
				},
			},
			{
				ID:        "spm@DATASTOREIOCONTROL",
				Namespace: "spm",
				PropertyList: []pbm.Property{
					{
						ID:       "limit",
						Value:    "200",
						DataType: "int",
					},
					{
						ID:       "reservation",
						Value:    "1000",
						DataType: "int",
					},
					{
						ID:       "shares",
						Value:    "2000",
						DataType: "int",
					},
				},
			},
		},
	}

	// Create PBM capability spec for the above defined user spec.
	createSpecVSANandSIOC, err := pbm.CreateCapabilityProfileSpec(pbmCreateSpecVSANandSIOC)
	if err != nil {
		t.Fatal(err)
	}

	// 9. Create SPBM VSAN profile.
	vsansiocProfileID, err := pc.CreateProfile(ctx, *createSpecVSANandSIOC)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("VSAN-SIOC Profile: %q successfully created", vsansiocProfileID.UniqueId)

	// 9. Get ProfileID by Name
	profileID, err := pc.ProfileIDByName(ctx, pbmCreateSpecVSANandSIOC.Name)
	if err != nil {
		t.Fatal(err)
	}

	if vsansiocProfileID.UniqueId != profileID {
		t.Errorf("vsan-sioc profile: %q and retrieved profileID: %q did not match", vsansiocProfileID.UniqueId, profileID)
	}

	t.Logf("VSAN-SIOC profile: %q and retrieved profileID: %q successfully matched", vsansiocProfileID.UniqueId, profileID)

	// 10. Get ProfileName by ID
	profileName, err := pc.GetProfileNameByID(ctx, profileID)
	if err != nil {
		t.Fatal(err)
	}

	if profileName != pbmCreateSpecVSANandSIOC.Name {
		t.Errorf("vsan-sioc profile: %q and retrieved profileName: %q did not match", pbmCreateSpecVSANandSIOC.Name, profileName)
	}

	t.Logf("vsan-sioc profile: %q and retrieved profileName: %q successfully matched", pbmCreateSpecVSANandSIOC.Name, profileName)

	// 11. Delete VSAN and VSAN-SIOC profile.
	_, err = pc.DeleteProfile(ctx, []types.PbmProfileId{*vsanProfileID, *vsansiocProfileID})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Profile: %+v successfully deleted", []types.PbmProfileId{*vsanProfileID, *vsansiocProfileID})
}

func TestSupportsEncryption(t *testing.T) {
	t.Run("valid profile id", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			pbmc, err := pbm.NewClient(ctx, c)
			assert.NoError(t, err)
			assert.NotNil(t, pbmc)

			ok, err := pbmc.SupportsEncryption(ctx, pbmsim.DefaultEncryptionProfileID)
			assert.NoError(t, err)
			assert.True(t, ok)
		})
	})
	t.Run("invalid profile id", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			pbmc, err := pbm.NewClient(ctx, c)
			assert.NoError(t, err)
			assert.NotNil(t, pbmc)

			ok, err := pbmc.SupportsEncryption(ctx, "invalid")
			assert.EqualError(t, err, "ServerFaultCode: Invalid profile ID")
			assert.True(t, fault.Is(err, &vim.RuntimeFault{}))
			assert.False(t, ok)
		})
	})
}

func TestClientCookie(t *testing.T) {
	simulator.Test(func(ctx context.Context, c *vim25.Client) {
		pbmc, err := pbm.NewClient(ctx, c)
		assert.NoError(t, err)
		assert.NotNil(t, pbmc)

		// Using default / expected Header.Cookie.XMLName = vcSessionCookie
		ok, err := pbmc.SupportsEncryption(ctx, pbmsim.DefaultEncryptionProfileID)
		assert.NoError(t, err)
		assert.True(t, ok)

		// Using invalid Header.Cookie.XMLName = myCookie
		myCookie := c.SessionCookie()
		myCookie.XMLName.Local = "myCookie"

		pbmc.Client.Cookie = func() *soap.HeaderElement {
			return myCookie
		}

		ok, err = pbmc.SupportsEncryption(ctx, pbmsim.DefaultEncryptionProfileID)
		assert.EqualError(t, err, "ServerFaultCode: NotAuthenticated")
		assert.False(t, ok)
	})
}

func TestStoragePolicyK8sCompliantNames(t *testing.T) {
	t.Run("Validate K8sCompliantName set for existing storage profiles", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			pc, err := pbm.NewClient(ctx, c)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("PBM version=%s", pc.ServiceContent.AboutInfo.Version)

			// Query all the profiles on the vCenter and verify that all profiles have K8sCompliantName set.
			category := types.PbmProfileCategoryEnumREQUIREMENT
			profileDetails, err := pc.QueryProfileDetails(ctx, string(category), true)
			if err != nil {
				t.Fatal(err)
			}
			for _, profile := range profileDetails.Returnval {
				assert.True(t, profile.Profile.GetPbmCapabilityProfile().K8sCompliantName != "")
			}
		})
	})
	t.Run("Validate K8sCompliantName set after new policy creation", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			pc, err := pbm.NewClient(ctx, c)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("PBM version=%s", pc.ServiceContent.AboutInfo.Version)

			category := types.PbmProfileCategoryEnumREQUIREMENT
			validK8sCompliantName := "kubernetes-vsan-testpolicy"
			// Create one storage profile with some K8sCompliantName passed
			// VSAN profile with 2 capability instances - hostFailuresToTolerate = 2, stripeWidth = 1
			pbmCreateSpecForVSAN := pbm.CapabilityProfileCreateSpec{
				Name:        "Kubernetes-VSAN-TestPolicy",
				Description: "VSAN Test policy create",
				Category:    string(types.PbmProfileCategoryEnumREQUIREMENT),
				CapabilityList: []pbm.Capability{
					{
						ID:        "hostFailuresToTolerate",
						Namespace: "VSAN",
						PropertyList: []pbm.Property{
							{
								ID:       "hostFailuresToTolerate",
								Value:    "2",
								DataType: "int",
							},
						},
					},
					{
						ID:        "stripeWidth",
						Namespace: "VSAN",
						PropertyList: []pbm.Property{
							{
								ID:       "stripeWidth",
								Value:    "1",
								DataType: "int",
							},
						},
					},
				},
				K8sCompliantName: validK8sCompliantName,
			}

			createSpecVSAN, err := pbm.CreateCapabilityProfileSpec(pbmCreateSpecForVSAN)
			if err != nil {
				t.Fatal(err)
			}

			vsanProfileID, err := pc.CreateProfile(ctx, *createSpecVSAN)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("VSAN Profile: %q with Name %q successfully created", vsanProfileID.UniqueId, createSpecVSAN.Name)

			profileDetails, err := pc.QueryProfileDetails(ctx, string(category), true)
			if err != nil {
				t.Fatal(err)
			}
			for _, profile := range profileDetails.Returnval {
				if profile.Profile.GetPbmCapabilityProfile().ProfileId.UniqueId == vsanProfileID.UniqueId {
					assert.True(t, profile.Profile.GetPbmCapabilityProfile().K8sCompliantName != "")
					assert.True(t, profile.Profile.GetPbmCapabilityProfile().K8sCompliantName == validK8sCompliantName)
					t.Logf("Profile %q:%q has K8sCompliant Name %q set now on VCDB", createSpecVSAN.Name,
						vsanProfileID.UniqueId, profile.Profile.GetPbmCapabilityProfile().K8sCompliantName)
					break
				}
			}
		})
	})
	t.Run("Validate K8sCompliantName and OtherK8sCompliantNames set correctly after update", func(t *testing.T) {
		simulator.Test(func(ctx context.Context, c *vim25.Client) {
			pc, err := pbm.NewClient(ctx, c)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("PBM version=%s", pc.ServiceContent.AboutInfo.Version)

			category := types.PbmProfileCategoryEnumREQUIREMENT
			validK8sCompliantName := "kubernetes-vsan-testpolicy"
			validK8sCompliantName2 := "kubernetes-vsan-testpolicy-renamed"
			otherK8sCompliantNames := []string{"kubernetes-vsan-testpolicy-old"}

			// Create one storage profile with some K8sCompliantName passed
			// VSAN profile with 2 capability instances - hostFailuresToTolerate = 2, stripeWidth = 1
			pbmCreateSpecForVSAN := pbm.CapabilityProfileCreateSpec{
				Name:        "Kubernetes-VSAN-TestPolicy",
				Description: "VSAN Test policy create",
				Category:    string(types.PbmProfileCategoryEnumREQUIREMENT),
				CapabilityList: []pbm.Capability{
					{
						ID:        "hostFailuresToTolerate",
						Namespace: "VSAN",
						PropertyList: []pbm.Property{
							{
								ID:       "hostFailuresToTolerate",
								Value:    "2",
								DataType: "int",
							},
						},
					},
					{
						ID:        "stripeWidth",
						Namespace: "VSAN",
						PropertyList: []pbm.Property{
							{
								ID:       "stripeWidth",
								Value:    "1",
								DataType: "int",
							},
						},
					},
				},
				K8sCompliantName: validK8sCompliantName,
			}

			createSpecVSAN, err := pbm.CreateCapabilityProfileSpec(pbmCreateSpecForVSAN)
			if err != nil {
				t.Fatal(err)
			}

			vsanProfileID, err := pc.CreateProfile(ctx, *createSpecVSAN)
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("VSAN Profile: %q with Name %q successfully created", vsanProfileID.UniqueId, createSpecVSAN.Name)

			// Update K8sCompliantName and otherK8sCompliantNames of the policy, with both valid values
			err = pc.UpdateK8sCompliantNames(ctx, vsanProfileID.UniqueId, validK8sCompliantName, otherK8sCompliantNames)
			assert.NoError(t, err)

			profileDetails, err := pc.QueryProfileDetails(ctx, string(category), true)
			if err != nil {
				t.Fatal(err)
			}
			for _, profile := range profileDetails.Returnval {
				if profile.Profile.GetPbmCapabilityProfile().ProfileId.UniqueId == vsanProfileID.UniqueId {
					assert.True(t, profile.Profile.GetPbmCapabilityProfile().K8sCompliantName != "")
					assert.True(t, profile.Profile.GetPbmCapabilityProfile().K8sCompliantName == validK8sCompliantName)
					assert.EqualValues(t, profile.Profile.GetPbmCapabilityProfile().OtherK8sCompliantNames, otherK8sCompliantNames)
					break
				}
			}
			t.Logf("Profile %q:%q has updated otherK8sCompliant Names on VCDB", createSpecVSAN.Name, vsanProfileID.UniqueId)

			// Try to update K8sCompliantName, which should fail since the value is immutable
			err = pc.UpdateK8sCompliantNames(ctx, vsanProfileID.UniqueId, validK8sCompliantName2, otherK8sCompliantNames)
			assert.EqualError(t, err, "ServerFaultCode: Invalid Argument")

			// Try to update otherK8sCompliantNames with another unique name, which should succeed
			otherK8sCompliantNamesSuccess := []string{validK8sCompliantName + "-old-1"}
			otherK8sCompliantNamesSuccess = append(otherK8sCompliantNamesSuccess, otherK8sCompliantNames...)
			err = pc.UpdateK8sCompliantNames(ctx, vsanProfileID.UniqueId, validK8sCompliantName, otherK8sCompliantNamesSuccess)
			assert.NoError(t, err)

			// Try to update otherK8sCompliantNames with duplicate name, which should fail
			otherK8sCompliantNamesFail := []string{validK8sCompliantName}
			otherK8sCompliantNamesFail = append(otherK8sCompliantNamesFail, otherK8sCompliantNames...)
			err = pc.UpdateK8sCompliantNames(ctx, vsanProfileID.UniqueId, validK8sCompliantName, otherK8sCompliantNamesFail)
			assert.EqualError(t, err, "ServerFaultCode: Duplicate Name")

			// Resolve K8sCompliantName and otherK8sCompliantNames for the policy
			otherK8sCompliantNamesResolved := []string{}
			otherK8sCompliantNamesResolved = append(otherK8sCompliantNamesResolved, otherK8sCompliantNamesSuccess...)
			otherK8sCompliantNamesResolved = append(otherK8sCompliantNamesResolved, validK8sCompliantName+"-latebinding")
			err = pc.ResolveK8sCompliantNames(ctx)
			assert.NoError(t, err)

			profileDetails, err = pc.QueryProfileDetails(ctx, string(category), true)
			if err != nil {
				t.Fatal(err)
			}
			for _, profile := range profileDetails.Returnval {
				if profile.Profile.GetPbmCapabilityProfile().ProfileId.UniqueId == vsanProfileID.UniqueId {
					assert.True(t, profile.Profile.GetPbmCapabilityProfile().K8sCompliantName != "")
					assert.True(t, profile.Profile.GetPbmCapabilityProfile().K8sCompliantName == validK8sCompliantName)
					assert.EqualValues(t, profile.Profile.GetPbmCapabilityProfile().OtherK8sCompliantNames, otherK8sCompliantNamesResolved)
					break
				}
			}
			t.Logf("Profile %q:%q has resolved and updated otherK8sCompliant Names on VCDB", createSpecVSAN.Name, vsanProfileID.UniqueId)
		})
	})
}
