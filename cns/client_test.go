// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cns

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/dougm/pretty"
	"github.com/google/uuid"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/pbm"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/debug"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"

	"github.com/vmware/govmomi"
	cnstypes "github.com/vmware/govmomi/cns/types"
	vim25types "github.com/vmware/govmomi/vim25/types"
	vsanfstypes "github.com/vmware/govmomi/vsan/vsanfs/types"
)

const VSphere70u3VersionInt = 703
const VSphere80u3VersionInt = 803
const VSphere91VersionInt = 910

func TestClient(t *testing.T) {
	// set CNS_DEBUG to true if you need to emit soap traces from these tests
	// soap traces will be emitted in the govmomi/cns/.soap directory
	// example export CNS_DEBUG='true'
	enableDebug := os.Getenv("CNS_DEBUG")
	soapTraceDirectory := ".soap"

	url := os.Getenv("CNS_VC_URL") // example: export CNS_VC_URL='https://username:password@vc-ip/sdk'
	datacenter := os.Getenv("CNS_DATACENTER")
	datastore := os.Getenv("CNS_DATASTORE")
	datastore2 := os.Getenv("CNS_DATASTORE2")
	spbmProfileName := os.Getenv("CNS_SPBM_PROFILE_NAME")

	// set CNS_RUN_TRANSACTION_TESTS environment to true, if you want to run CNS Transaction tests
	// example: export CNS_RUN_TRANSACTION_TESTS='true'
	run_cns_transaction_tests := os.Getenv("CNS_RUN_TRANSACTION_TESTS")
	// set CNS_RUN_SHARED_DISK_TESTS environment to true, if you want to run shared disk related tests
	// example: export CNS_RUN_SHARED_DISK_TESTS='true'
	run_shared_disk_tests_tests := os.Getenv("CNS_RUN_SHARED_DISK_TESTS")

	// set CNS_RUN_MULTICLUSTER_PER_ZONE_TESTS environment to true, if you want to run tests
	// on deployment with Zone with multiple vSphere Clusters
	// example: export CNS_RUN_MULTICLUSTER_PER_ZONE_TESTS='true'
	run_cns_multicluster_per_zone_tests := os.Getenv("CNS_RUN_MULTICLUSTER_PER_ZONE_TESTS")

	// set CNS_RUN_FILESHARE_TESTS environment to true, if your setup has vsanfileshare enabled.
	// when CNS_RUN_FILESHARE_TESTS is not set to true, vsan file share related tests are skipped.
	// example: export CNS_RUN_FILESHARE_TESTS='true'
	run_fileshare_tests := os.Getenv("CNS_RUN_FILESHARE_TESTS")

	// if backingDiskURLPath is not set, test for Creating Volume with setting BackingDiskUrlPath in the BackingObjectDetails of
	// CnsVolumeCreateSpec will be skipped.
	// example: export BACKING_DISK_URL_PATH='https://vc-ip/folder/vmdkfilePath.vmdk?dcPath=DataCenterPath&dsName=DataStoreName'
	backingDiskURLPath := os.Getenv("BACKING_DISK_URL_PATH")

	// set REMOTE_VC_URL, REMOTE_DATACENTER only if you want to test cross-VC CNS operations.
	// For instance, testing cross-VC volume migration.
	remoteVcUrl := os.Getenv("REMOTE_VC_URL")
	remoteDatacenter := os.Getenv("REMOTE_DATACENTER")

	// if datastoreForMigration is not set, test for CNS Relocate API of a volume to another datastore is skipped.
	// input format is same as CNS_DATASTORE. Format eg. "vSANDirect_10.92.217.162_mpx.vmhba0:C0:T2:L0"/ "vsandatastore"
	datastoreForMigration := os.Getenv("CNS_MIGRATION_DATASTORE")

	// if spbmPolicyId4Reconfig is not set, test for CnsReconfigVolumePolicy API will be skipped
	// example: export CNS_SPBM_POLICY_ID_4_RECONFIG=6f64d90e-2ad5-4c4d-8cbc-a3330ebc496c
	spbmPolicyId4Reconfig := os.Getenv("CNS_SPBM_POLICY_ID_4_RECONFIG")

	if url == "" || datacenter == "" || datastore == "" {
		t.Skip("CNS_VC_URL or CNS_DATACENTER or CNS_DATASTORE is not set")
	}
	resourcePoolPath := os.Getenv("CNS_RESOURCE_POOL_PATH") // example "/datacenter-name/host/host-ip/Resources" or  /datacenter-name/host/cluster-name/Resources
	u, err := soap.ParseURL(url)
	if err != nil {
		t.Fatal(err)
	}

	if enableDebug == "true" {
		if _, err := os.Stat(soapTraceDirectory); os.IsNotExist(err) {
			os.Mkdir(soapTraceDirectory, 0755)
		}
		p := debug.FileProvider{
			Path: soapTraceDirectory,
		}
		debug.SetProvider(&p)
	}

	ctx := context.Background()
	c, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		t.Fatal(err)
	}
	cnsClient, err := NewClient(ctx, c.Client)
	if err != nil {
		t.Fatal(err)
	}
	finder := find.NewFinder(cnsClient.vim25Client, false)
	dc, err := finder.Datacenter(ctx, datacenter)
	if err != nil {
		t.Fatal(err)
	}
	finder.SetDatacenter(dc)
	ds, err := finder.Datastore(ctx, datastore)
	if err != nil {
		t.Fatal(err)
	}

	props := []string{"info", "summary"}
	pc := property.DefaultCollector(c.Client)
	var dsSummaries []mo.Datastore
	err = pc.Retrieve(ctx, []vim25types.ManagedObjectReference{ds.Reference()}, props, &dsSummaries)
	if err != nil {
		t.Fatal(err)
	}
	dsUrl := dsSummaries[0].Summary.Url

	var dsList []vim25types.ManagedObjectReference
	dsList = append(dsList, ds.Reference())

	var containerClusterArray []cnstypes.CnsContainerCluster
	containerCluster := cnstypes.CnsContainerCluster{
		ClusterType:         string(cnstypes.CnsClusterTypeKubernetes),
		ClusterId:           "demo-cluster-id",
		VSphereUser:         "Administrator@vsphere.local",
		ClusterFlavor:       string(cnstypes.CnsClusterFlavorVanilla),
		ClusterDistribution: "OpenShift",
	}
	containerClusterArray = append(containerClusterArray, containerCluster)

	// Test CreateVolume API
	var cnsVolumeCreateSpecList []cnstypes.CnsVolumeCreateSpec
	cnsVolumeCreateSpec := cnstypes.CnsVolumeCreateSpec{
		Name:       "pvc-901e87eb-c2bd-11e9-806f-005056a0c9a0",
		VolumeType: string(cnstypes.CnsVolumeTypeBlock),
		Metadata: cnstypes.CnsVolumeMetadata{
			ContainerCluster: containerCluster,
		},
		BackingObjectDetails: &cnstypes.CnsBlockBackingDetails{
			CnsBackingObjectDetails: cnstypes.CnsBackingObjectDetails{
				CapacityInMb: 5120,
			},
		},
	}

	pvclaimUID := "901e87eb-c2bd-11e9-806f-005056a0c9a0"
	isvSphereVersion91orAbove := isvSphereVersion91orAbove(ctx, c.ServiceContent.About)
	if run_cns_transaction_tests == "true" && isvSphereVersion91orAbove {
		t.Logf("setting volumeID: %q in the cnsVolumeCreateSpec", pvclaimUID)
		cnsVolumeCreateSpec.VolumeId = &cnstypes.CnsVolumeId{
			Id: pvclaimUID,
		}
	}

	if run_cns_multicluster_per_zone_tests == "true" && isvSphereVersion91orAbove {
		var clusters []vim25types.ManagedObjectReference
		computeResources, err := finder.ComputeResourceList(ctx, "*")
		if err != nil {
			t.Fatal(err)
		}
		for _, computeResource := range computeResources {
			clusters = append(clusters, computeResource.Reference())
		}
		t.Logf("set cnsVolumeCreateSpec.ActiveClusters=%v", clusters)
		cnsVolumeCreateSpec.ActiveClusters = clusters

		pbmclient, err := pbm.NewClient(ctx, c.Client)
		if err != nil {
			t.Fatal(err)
		}
		storagePolicyID, err := pbmclient.ProfileIDByName(ctx, spbmProfileName)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("set cnsVolumeCreateSpec.Profile=%s", storagePolicyID)
		cnsVolumeCreateSpec.Profile = []vim25types.BaseVirtualMachineProfileSpec{
			&vim25types.VirtualMachineDefinedProfileSpec{
				ProfileId: storagePolicyID,
			},
		}
	} else {
		cnsVolumeCreateSpec.Datastores = dsList
	}
	cnsVolumeCreateSpecList = append(cnsVolumeCreateSpecList, cnsVolumeCreateSpec)
	t.Logf("Creating volume using the spec: %+v", pretty.Sprint(cnsVolumeCreateSpec))
	createTask, err := cnsClient.CreateVolume(ctx, cnsVolumeCreateSpecList)
	if err != nil {
		t.Errorf("Failed to create volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	createTaskInfo, err := GetTaskInfo(ctx, createTask)
	if err != nil {
		t.Errorf("Failed to create volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	createTaskResult, err := GetTaskResult(ctx, createTaskInfo)
	if err != nil {
		t.Errorf("Failed to create volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if createTaskResult == nil {
		t.Fatalf("Empty create task results")
		t.FailNow()
	}
	createVolumeOperationRes := createTaskResult.GetCnsVolumeOperationResult()
	if createVolumeOperationRes.Fault != nil {
		if cnsFault, ok := createVolumeOperationRes.Fault.Fault.(*cnstypes.CnsFault); ok {
			if cause := cnsFault.FaultCause; cause != nil {
				if inner, ok := cause.Fault.(*vim25types.NotSupported); ok {
					t.Logf("Caught NotSupported fault: %q", cause.LocalizedMessage)
				} else {
					t.Logf("Inner fault type: %T", inner)
				}
			}
		}
		t.Fatalf("Failed to create volume: fault=%+v", createVolumeOperationRes.Fault)
	}
	volumeId := createVolumeOperationRes.VolumeId.Id
	volumeCreateResult := (createTaskResult).(*cnstypes.CnsVolumeCreateResult)
	t.Logf("volumeCreateResult %+v", volumeCreateResult)
	if run_cns_transaction_tests == "true" && isvSphereVersion91orAbove {
		if volumeId != pvclaimUID {
			t.Fatalf("failed to create volume with supplied volume ID: %q volume "+
				"created with diffrent UUID: %q", pvclaimUID, volumeId)
		}
	}
	t.Logf("Volume created successfully. volumeId: %s", volumeId)

	if run_cns_multicluster_per_zone_tests == "true" && isvSphereVersion91orAbove {
		if len(volumeCreateResult.PlacementResults[0].Clusters) == 0 {
			t.Fatalf("clusters in the placement result in createvolumeresult can not be empty. "+
				"volumeCreateResult.PlacementResults :%q", pretty.Sprint(volumeCreateResult.PlacementResults))
		}
	}

	// Creating Volume with same ID again on different datastore
	// to observe CnsVolumeAlreadyExistsFault
	if run_cns_transaction_tests == "true" && isvSphereVersion91orAbove {
		if datastore2 != "" {
			ds2, err := finder.Datastore(ctx, datastore2)
			if err != nil {
				t.Fatal(err)
			}
			var ds2List []vim25types.ManagedObjectReference
			ds2List = append(ds2List, ds2.Reference())
			cnsVolumeCreateSpec.Datastores = ds2List
			var cnsVolumeCreateSpecList2 []cnstypes.CnsVolumeCreateSpec

			t.Logf("Creating Volume again on diffrent datastore using same generated Volume ID")
			cnsVolumeCreateSpecList = append(cnsVolumeCreateSpecList2, cnsVolumeCreateSpec)
			t.Logf("Creating volume using the spec: %+v", pretty.Sprint(cnsVolumeCreateSpec))
			createTask, err := cnsClient.CreateVolume(ctx, cnsVolumeCreateSpecList)
			if err != nil {
				t.Errorf("Failed to create volume. Error: %+v \n", err)
				t.Fatal(err)
			}
			createTaskInfo, err := GetTaskInfo(ctx, createTask)
			if err != nil {
				t.Errorf("Failed to create volume. Error: %+v \n", err)
				t.Fatal(err)
			}
			createTaskResult, err := GetTaskResult(ctx, createTaskInfo)
			if err != nil {
				t.Errorf("Failed to create volume. Error: %+v \n", err)
				t.Fatal(err)
			}
			if createTaskResult == nil {
				t.Fatalf("Empty create task results")
				t.FailNow()
			}
			createVolumeOperationRes := createTaskResult.GetCnsVolumeOperationResult()
			if createVolumeOperationRes.Fault != nil {
				t.Logf("createVolumeOperationRes.Fault: %+v", pretty.Sprint(createVolumeOperationRes))
				_, ok := createVolumeOperationRes.Fault.Fault.(*cnstypes.CnsVolumeAlreadyExistsFault)
				if !ok {
					t.Fatalf("Fault is not CnsVolumeAlreadyExistsFault")
				}
			} else {
				t.Fatalf("expecting CnsVolumeAlreadyExistsFault while creating volume with same ID on different datastore")
			}
		}
	}

	if cnsClient.Version != ReleaseVSAN67u3 {
		// Test creating static volume using existing CNS volume should fail
		var staticCnsVolumeCreateSpecList []cnstypes.CnsVolumeCreateSpec
		staticCnsVolumeCreateSpec := cnstypes.CnsVolumeCreateSpec{
			Name:       "pvc-901e87eb-c2bd-11e9-806f-005056a0c9a0",
			VolumeType: string(cnstypes.CnsVolumeTypeBlock),
			Metadata: cnstypes.CnsVolumeMetadata{
				ContainerCluster: containerCluster,
			},
			BackingObjectDetails: &cnstypes.CnsBlockBackingDetails{
				CnsBackingObjectDetails: cnstypes.CnsBackingObjectDetails{
					CapacityInMb: 5120,
				},
				BackingDiskId: volumeId,
			},
		}

		staticCnsVolumeCreateSpecList = append(staticCnsVolumeCreateSpecList, staticCnsVolumeCreateSpec)
		t.Logf("Creating volume using the spec: %+v", pretty.Sprint(staticCnsVolumeCreateSpec))
		recreateTask, err := cnsClient.CreateVolume(ctx, staticCnsVolumeCreateSpecList)
		if err != nil {
			t.Errorf("Failed to create volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		reCreateTaskInfo, err := GetTaskInfo(ctx, recreateTask)
		if err != nil {
			t.Errorf("Failed to create volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		reCreateTaskResult, err := GetTaskResult(ctx, reCreateTaskInfo)
		if err != nil {
			t.Errorf("Failed to create volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		if reCreateTaskResult == nil {
			t.Fatalf("Empty create task results")
			t.FailNow()
		}
		reCreateVolumeOperationRes := reCreateTaskResult.GetCnsVolumeOperationResult()
		t.Logf("reCreateVolumeOperationRes.: %+v", pretty.Sprint(reCreateVolumeOperationRes))
		if reCreateVolumeOperationRes.Fault == nil {
			t.Logf("re-create same volume did not fail as expected")
		} else {
			t.Fatalf("re-creating same volume failed with fault: %+v", pretty.Sprint(reCreateVolumeOperationRes.Fault))
		}
	}

	// Test QueryVolume API
	var queryFilter cnstypes.CnsQueryFilter
	var volumeIDList []cnstypes.CnsVolumeId
	volumeIDList = append(volumeIDList, cnstypes.CnsVolumeId{Id: volumeId})
	queryFilter.VolumeIds = volumeIDList
	t.Logf("Calling QueryVolume using queryFilter: %+v", pretty.Sprint(queryFilter))
	queryResult, err := cnsClient.QueryVolume(ctx, &queryFilter)
	if err != nil {
		t.Errorf("Failed to query volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	t.Logf("Successfully Queried Volumes. queryResult: %+v", pretty.Sprint(queryResult))

	// Test QueryVolumeInfo API
	// QueryVolumeInfo is not supported on ReleaseVSAN67u3 and ReleaseVSAN70
	// This API is available on vSphere 7.0u1 onward
	if cnsClient.Version != ReleaseVSAN67u3 && cnsClient.Version != ReleaseVSAN70 {
		t.Logf("Calling QueryVolumeInfo using: %+v", pretty.Sprint(volumeIDList))
		queryVolumeInfoTask, err := cnsClient.QueryVolumeInfo(ctx, volumeIDList)
		if err != nil {
			t.Errorf("Failed to query volumes with QueryVolumeInfo. Error: %+v \n", err)
			t.Fatal(err)
		}
		queryVolumeInfoTaskInfo, err := GetTaskInfo(ctx, queryVolumeInfoTask)
		if err != nil {
			t.Errorf("Failed to query volumes with QueryVolumeInfo. Error: %+v \n", err)
			t.Fatal(err)
		}
		queryVolumeInfoTaskResults, err := GetTaskResultArray(ctx, queryVolumeInfoTaskInfo)
		if err != nil {
			t.Errorf("Failed to query volumes with QueryVolumeInfo. Error: %+v \n", err)
			t.Fatal(err)
		}
		if queryVolumeInfoTaskResults == nil {
			t.Fatalf("Empty queryVolumeInfoTaskResult")
			t.FailNow()
		}
		for _, queryVolumeInfoTaskResult := range queryVolumeInfoTaskResults {
			queryVolumeInfoOperationRes := queryVolumeInfoTaskResult.GetCnsVolumeOperationResult()
			if queryVolumeInfoOperationRes.Fault != nil {
				t.Fatalf("Failed to query volumes with QueryVolumeInfo. fault=%+v", queryVolumeInfoOperationRes.Fault)
			}
			t.Logf("Successfully Queried Volumes. queryVolumeInfoTaskResult: %+v", pretty.Sprint(queryVolumeInfoTaskResult))
		}
	}

	// Test BackingDiskObjectId field only for vVol or vSAN volume type
	var queryFilterBackingDiskObjectIdTest cnstypes.CnsQueryFilter
	var volumeIDListBackingDiskObjectIdTest []cnstypes.CnsVolumeId
	volumeIDListBackingDiskObjectIdTest = append(volumeIDListBackingDiskObjectIdTest, cnstypes.CnsVolumeId{Id: volumeId})
	queryFilterBackingDiskObjectIdTest.VolumeIds = volumeIDListBackingDiskObjectIdTest
	t.Logf("Calling QueryVolume using queryFilter: %+v", pretty.Sprint(queryFilterBackingDiskObjectIdTest))
	queryResultBackingDiskObjectIdTest, err := cnsClient.QueryVolume(ctx, &queryFilterBackingDiskObjectIdTest)
	if err != nil {
		t.Errorf("Failed to query all volumes. Error: %+v \n", err)
		t.Fatal(err)
	}
	t.Logf("Successfully Queried Volumes. queryResultBackingDiskObjectIdTest: %+v", pretty.Sprint(queryResultBackingDiskObjectIdTest))
	t.Log("Checking backingDiskObjectId retieved")
	datastoreType, err := ds.Type(ctx)
	if err != nil {
		t.Errorf("Failed to get datastore type. Error: %+v \n", err)
		t.Fatal(err)
	}
	for _, vol := range queryResultBackingDiskObjectIdTest.Volumes {
		// BackingDiskObjectId is only for vsan/vvol, for other type this field is empty but test should not fail
		backingDiskObjectId := vol.BackingObjectDetails.(*cnstypes.CnsBlockBackingDetails).BackingDiskObjectId
		if backingDiskObjectId == "" {
			if datastoreType == vim25types.HostFileSystemVolumeFileSystemTypeVsan || datastoreType == vim25types.HostFileSystemVolumeFileSystemTypeVVOL {
				t.Errorf("Failed to get BackingDiskObjectId")
				t.FailNow()
			}
		}
	}

	// Test BackingDiskPath field
	var queryFilterBackingDiskPathTest cnstypes.CnsQueryFilter
	var volumeIDListBackingDiskPathTest []cnstypes.CnsVolumeId
	volumeIDListBackingDiskPathTest = append(volumeIDListBackingDiskPathTest, cnstypes.CnsVolumeId{Id: volumeId})
	queryFilterBackingDiskPathTest.VolumeIds = volumeIDListBackingDiskPathTest
	t.Logf("Calling QueryVolume using queryFilter: %+v", pretty.Sprint(queryFilterBackingDiskPathTest))
	queryResultBackingDiskPathTest, err := cnsClient.QueryVolume(ctx, &queryFilterBackingDiskPathTest)
	if err != nil {
		t.Errorf("Failed to query all volumes. Error: %+v \n", err)
		t.Fatal(err)
	}
	t.Logf("Successfully Queried Volumes. queryResultBackingDiskPathTest: %+v", pretty.Sprint(queryResultBackingDiskPathTest))
	t.Log("Checking backingDiskPath retrieved")
	for _, vol := range queryResultBackingDiskPathTest.Volumes {
		backingDiskPath := vol.BackingObjectDetails.(*cnstypes.CnsBlockBackingDetails).BackingDiskPath
		if backingDiskPath == "" {
			t.Errorf("Failed to get BackingDiskPath")
			t.FailNow()
		}
	}

	// Test QuerySnapshots API on 7.0 U3 or above
	var snapshotQueryFilter cnstypes.CnsSnapshotQueryFilter
	var querySnapshotsTaskResult *cnstypes.CnsSnapshotQueryResult
	var QuerySnapshotsFunc func(snapshotQueryFilter cnstypes.CnsSnapshotQueryFilter) *cnstypes.CnsSnapshotQueryResult

	if isvSphereVersion70U3orAbove(ctx, c.ServiceContent.About) {
		// Construct the CNS SnapshotQueryFilter and the function handler of QuerySnapshots
		QuerySnapshotsFunc = func(snapshotQueryFilter cnstypes.CnsSnapshotQueryFilter) *cnstypes.CnsSnapshotQueryResult {
			querySnapshotsTask, err := cnsClient.QuerySnapshots(ctx, snapshotQueryFilter)
			if err != nil {
				t.Fatalf("Failed to get the task of QuerySnapshots. Error: %+v \n", err)
			}
			querySnapshotsTaskInfo, err := GetTaskInfo(ctx, querySnapshotsTask)
			if err != nil {
				t.Fatalf("Failed to get the task info of QuerySnapshots. Error: %+v \n", err)
			}
			querySnapshotsTaskResult, err := GetQuerySnapshotsTaskResult(ctx, querySnapshotsTaskInfo)
			if err != nil {
				t.Fatalf("Failed to get the task result of QuerySnapshots. Error: %+v \n", err)
			}
			return querySnapshotsTaskResult
		}

		// Calls QuerySnapshots before CreateSnapshots
		snapshotQueryFilter = cnstypes.CnsSnapshotQueryFilter{
			SnapshotQuerySpecs: []cnstypes.CnsSnapshotQuerySpec{
				{
					VolumeId: cnstypes.CnsVolumeId{Id: volumeId},
				},
			},
		}
		t.Logf("QuerySnapshots before CreateSnapshots, snapshotQueryFilter %+v", snapshotQueryFilter)
		querySnapshotsTaskResult = QuerySnapshotsFunc(snapshotQueryFilter)
		t.Logf("snapshotQueryResult %+v", querySnapshotsTaskResult)
	}

	// Test CreateSnapshot API
	// Construct the CNS SnapshotCreateSpec list
	desc := "example-vanilla-block-snapshot"
	var cnsSnapshotCreateSpecList []cnstypes.CnsSnapshotCreateSpec
	cnsSnapshotCreateSpec := cnstypes.CnsSnapshotCreateSpec{
		VolumeId: cnstypes.CnsVolumeId{
			Id: volumeId,
		},
		Description: desc,
	}

	var generatedSnapshotUUIDFromClient string
	if run_cns_transaction_tests == "true" && isvSphereVersion91orAbove {
		generatedSnapshotUUIDFromClient = uuid.New().String()
		t.Logf("setting SnapshotID: %q in the cnsSnapshotCreateSpec", generatedSnapshotUUIDFromClient)
		cnsSnapshotCreateSpec.SnapshotId = &cnstypes.CnsSnapshotId{
			Id: generatedSnapshotUUIDFromClient,
		}
	}

	cnsSnapshotCreateSpecList = append(cnsSnapshotCreateSpecList, cnsSnapshotCreateSpec)
	t.Logf("Creating snapshot using the spec: %+v", pretty.Sprint(cnsSnapshotCreateSpecList))
	createSnapshotsTask, err := cnsClient.CreateSnapshots(ctx, cnsSnapshotCreateSpecList)
	if err != nil {
		t.Errorf("Failed to get the task of CreateSnapshots. Error: %+v \n", err)
		t.Fatal(err)
	}
	createSnapshotsTaskInfo, err := GetTaskInfo(ctx, createSnapshotsTask)
	if err != nil {
		t.Errorf("Failed to get the task info of CreateSnapshots. Error: %+v \n", err)
		t.Fatal(err)
	}
	createSnapshotsTaskResult, err := GetTaskResult(ctx, createSnapshotsTaskInfo)
	if err != nil {
		t.Errorf("Failed to get the task result of CreateSnapshots. Error: %+v \n", err)
		t.Fatal(err)
	}
	createSnapshotsOperationRes := createSnapshotsTaskResult.GetCnsVolumeOperationResult()
	if createSnapshotsOperationRes.Fault != nil {
		t.Fatalf("Failed to create snapshots: fault=%+v", createSnapshotsOperationRes.Fault)
	}

	snapshotCreateResult := any(createSnapshotsTaskResult).(*cnstypes.CnsSnapshotCreateResult)
	snapshotId := snapshotCreateResult.Snapshot.SnapshotId.Id
	snapshotCreateTime := snapshotCreateResult.Snapshot.CreateTime
	t.Logf("snapshotCreateResult: %+v", pretty.Sprint(snapshotCreateResult))
	if run_cns_transaction_tests == "true" && isvSphereVersion91orAbove {
		if snapshotId != generatedSnapshotUUIDFromClient {
			t.Fatalf("failed to create snapshot with snapshot id: %q snapshot "+
				"created with diffrent id: %q", generatedSnapshotUUIDFromClient, snapshotId)
		}
	}
	t.Logf("CreateSnapshots: Snapshot created successfully. volumeId: %q, snapshot id %q, time stamp %+v, opId: %q", volumeId, snapshotId, snapshotCreateTime, createSnapshotsTaskInfo.ActivationId)

	// Test QuerySnapshots API on 7.0 U3 or above
	if isvSphereVersion70U3orAbove(ctx, c.ServiceContent.About) {
		// Calls QuerySnapshots after CreateSnapshots
		snapshotQueryFilter = cnstypes.CnsSnapshotQueryFilter{
			SnapshotQuerySpecs: []cnstypes.CnsSnapshotQuerySpec{
				{
					VolumeId:   cnstypes.CnsVolumeId{Id: volumeId},
					SnapshotId: &cnstypes.CnsSnapshotId{Id: snapshotId},
				},
			},
		}
		t.Logf("QuerySnapshots after CreateSnapshots, snapshotQueryFilter %+v", snapshotQueryFilter)
		querySnapshotsTaskResult = QuerySnapshotsFunc(snapshotQueryFilter)
		t.Logf("snapshotQueryResult %+v", querySnapshotsTaskResult)
	}

	// Test CreateVolumeFromSnapshot functionality by calling CreateVolume with VolumeSource set
	// Query Volume for capacity
	var queryVolumeIDList []cnstypes.CnsVolumeId
	queryVolumeIDList = append(queryVolumeIDList, cnstypes.CnsVolumeId{Id: volumeId})
	queryFilter.VolumeIds = queryVolumeIDList
	t.Logf("CreateVolumeFromSnapshot: calling QueryVolume using queryFilter: %+v", pretty.Sprint(queryFilter))
	queryResult, err = cnsClient.QueryVolume(ctx, &queryFilter)
	if err != nil {
		t.Errorf("Failed to query volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	var snapshotSize int64
	if len(queryResult.Volumes) > 0 {
		snapshotSize = queryResult.Volumes[0].BackingObjectDetails.GetCnsBackingObjectDetails().CapacityInMb
	} else {
		msg := fmt.Sprintf("failed to get the snapshot size by querying volume: %q", volumeId)
		t.Fatal(msg)
	}
	t.Logf("CreateVolumeFromSnapshot: Successfully Queried Volumes. queryResult: %+v", pretty.Sprint(queryResult))

	// Test CnsSyncVolume API
	t.Logf("Calling syncVolume for volumeId: %v ...\n", volumeId)
	var cnsSyncVolumeSpecs []cnstypes.CnsSyncVolumeSpec
	dataStore := ds.Reference()
	cnsSyncVolumeSpec := cnstypes.CnsSyncVolumeSpec{
		VolumeId: cnstypes.CnsVolumeId{
			Id: volumeId,
		},
		Datastore: &dataStore,
		SyncMode:  []string{string(cnstypes.CnsSyncVolumeModeSPACE_USAGE)},
	}
	cnsSyncVolumeSpecs = append(cnsSyncVolumeSpecs, cnsSyncVolumeSpec)
	syncVolumeTask, err := cnsClient.SyncVolume(ctx, cnsSyncVolumeSpecs)
	if err != nil {
		t.Errorf("Failed to sync volume %v. Error: %+v \n", volumeId, err)
		t.Fatal(err)
	}
	syncVolumeTaskInfo, err := GetTaskInfo(ctx, syncVolumeTask)
	if err != nil {
		t.Errorf("Failed to get sync volume taskInfo. Error: %+v \n", err)
		t.Fatal(err)
	}
	if syncVolumeTaskInfo.State != vim25types.TaskInfoStateSuccess {
		t.Errorf("Failed to sync volume. Error: %+v \n", syncVolumeTaskInfo.Error)
		t.Fatalf("%+v", syncVolumeTaskInfo.Error)
	}
	t.Logf("sync volume for volumeId: %v successful...\n", volumeId)

	// Construct the CNS VolumeCreateSpec list
	cnsCreateVolumeFromSnapshotCreateSpec := cnstypes.CnsVolumeCreateSpec{
		Name:       "pvc-901e87eb-c2bd-11e9-806f-005056a0c9a0-create-from-snapshot",
		VolumeType: string(cnstypes.CnsVolumeTypeBlock),
		Datastores: dsList,
		Metadata: cnstypes.CnsVolumeMetadata{
			ContainerCluster: containerCluster,
		},
		BackingObjectDetails: &cnstypes.CnsBlockBackingDetails{
			CnsBackingObjectDetails: cnstypes.CnsBackingObjectDetails{
				CapacityInMb: snapshotSize,
			},
		},
		VolumeSource: &cnstypes.CnsSnapshotVolumeSource{
			VolumeId: cnstypes.CnsVolumeId{
				Id: volumeId,
			},
			SnapshotId: cnstypes.CnsSnapshotId{
				Id: snapshotId,
			},
		},
	}
	var cnsCreateVolumeFromSnapshotCreateSpecList []cnstypes.CnsVolumeCreateSpec
	cnsCreateVolumeFromSnapshotCreateSpecList = append(cnsCreateVolumeFromSnapshotCreateSpecList, cnsCreateVolumeFromSnapshotCreateSpec)
	t.Logf("Creating volume from snapshot using the spec: %+v", pretty.Sprint(cnsCreateVolumeFromSnapshotCreateSpec))
	createVolumeFromSnapshotTask, err := cnsClient.CreateVolume(ctx, cnsCreateVolumeFromSnapshotCreateSpecList)
	if err != nil {
		t.Errorf("Failed to create volume from snapshot. Error: %+v \n", err)
		t.Fatal(err)
	}
	createVolumeFromSnapshotTaskInfo, err := GetTaskInfo(ctx, createVolumeFromSnapshotTask)
	if err != nil {
		t.Errorf("Failed to create volume from snapshot. Error: %+v \n", err)
		t.Fatal(err)
	}
	createVolumeFromSnapshotTaskResult, err := GetTaskResult(ctx, createVolumeFromSnapshotTaskInfo)
	if err != nil {
		t.Errorf("Failed to create volume from snapshot. Error: %+v \n", err)
		t.Fatal(err)
	}
	if createVolumeFromSnapshotTaskResult == nil {
		t.Fatalf("Empty create task results")
		t.FailNow()
	}
	createVolumeFromSnapshotOperationRes := createVolumeFromSnapshotTaskResult.GetCnsVolumeOperationResult()
	if createVolumeFromSnapshotOperationRes.Fault != nil {
		t.Fatalf("Failed to create volume from snapshot: fault=%+v", createVolumeFromSnapshotOperationRes.Fault)
	}
	createVolumeFromSnapshotVolumeId := createVolumeFromSnapshotOperationRes.VolumeId.Id
	createVolumeFromSnapshotResult := (createVolumeFromSnapshotTaskResult).(*cnstypes.CnsVolumeCreateResult)
	t.Logf("createVolumeFromSnapshotResult %+v", createVolumeFromSnapshotResult)
	t.Logf("Volume created from snapshot %s successfully. volumeId: %s", snapshotId, createVolumeFromSnapshotVolumeId)

	//  Clean up volume created from snapshot above
	var deleteVolumeFromSnapshotVolumeIDList []cnstypes.CnsVolumeId
	deleteVolumeFromSnapshotVolumeIDList = append(deleteVolumeFromSnapshotVolumeIDList, cnstypes.CnsVolumeId{Id: createVolumeFromSnapshotVolumeId})
	t.Logf("Deleting volume: %+v", deleteVolumeFromSnapshotVolumeIDList)
	deleteVolumeFromSnapshotTask, err := cnsClient.DeleteVolume(ctx, deleteVolumeFromSnapshotVolumeIDList, true)
	if err != nil {
		t.Errorf("Failed to delete volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	deleteVolumeFromSnapshotTaskInfo, err := GetTaskInfo(ctx, deleteVolumeFromSnapshotTask)
	if err != nil {
		t.Errorf("Failed to delete volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	deleteVolumeFromSnapshotTaskResult, err := GetTaskResult(ctx, deleteVolumeFromSnapshotTaskInfo)
	if err != nil {
		t.Errorf("Failed to delete volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if deleteVolumeFromSnapshotTaskResult == nil {
		t.Fatalf("Empty delete task results")
		t.FailNow()
	}
	deleteVolumeFromSnapshotOperationRes := deleteVolumeFromSnapshotTaskResult.GetCnsVolumeOperationResult()
	if deleteVolumeFromSnapshotOperationRes.Fault != nil {
		t.Fatalf("Failed to delete volume: fault=%+v", deleteVolumeFromSnapshotOperationRes.Fault)
	}
	t.Logf("Volume: %q deleted successfully", createVolumeFromSnapshotVolumeId)

	// Test CreateLinkedClone functionality by calling CreateVolume with VolumeSource set to snapshot and linkedclone
	// flag set to true
	cnsCreateLinkedCloneCreateSpec := cnstypes.CnsVolumeCreateSpec{
		Name:       "pvc-901e87eb-c2bd-11e9-806f-005056a0c9a0-create-lc",
		VolumeType: string(cnstypes.CnsVolumeTypeBlock),
		Datastores: dsList,
		Metadata: cnstypes.CnsVolumeMetadata{
			ContainerCluster: containerCluster,
		},
		BackingObjectDetails: &cnstypes.CnsBlockBackingDetails{
			CnsBackingObjectDetails: cnstypes.CnsBackingObjectDetails{
				CapacityInMb: snapshotSize,
			},
		},
		VolumeSource: &cnstypes.CnsSnapshotVolumeSource{
			VolumeId: cnstypes.CnsVolumeId{
				Id: volumeId,
			},
			SnapshotId: cnstypes.CnsSnapshotId{
				Id: snapshotId,
			},
			LinkedClone: true,
		},
	}
	var cnsCreateLinkedCloneCreateSpecList []cnstypes.CnsVolumeCreateSpec
	cnsCreateLinkedCloneCreateSpecList = append(cnsCreateLinkedCloneCreateSpecList, cnsCreateLinkedCloneCreateSpec)
	t.Logf("Creating linkedclone using the spec: %+v", pretty.Sprint(cnsCreateLinkedCloneCreateSpec))
	createLinkedCloneTask, err := cnsClient.CreateVolume(ctx, cnsCreateLinkedCloneCreateSpecList)
	if err != nil {
		t.Errorf("Failed to create linkedclone. Error: %+v \n", err)
		t.Fatal(err)
	}
	createLinkedCloneTaskInfo, err := GetTaskInfo(ctx, createLinkedCloneTask)
	if err != nil {
		t.Errorf("Failed to linked clone. Error: %+v \n", err)
		t.Fatal(err)
	}
	createLinkedCloneTaskResult, err := GetTaskResult(ctx, createLinkedCloneTaskInfo)
	if err != nil {
		t.Errorf("Failed to create linkedclone. Error: %+v \n", err)
		t.Fatal(err)
	}
	if createLinkedCloneTaskResult == nil {
		t.Fatalf("Empty create task results")
		t.FailNow()
	}
	createLinkedCloneOperationRes := createLinkedCloneTaskResult.GetCnsVolumeOperationResult()
	if createLinkedCloneOperationRes.Fault != nil {
		t.Fatalf("Failed to create linkedclone: fault=%+v", createLinkedCloneOperationRes.Fault)
	}
	createLinkedCloneVolumeId := createLinkedCloneOperationRes.VolumeId.Id
	createLinkedCloneResult := (createLinkedCloneTaskResult).(*cnstypes.CnsVolumeCreateResult)
	t.Logf("createLinkedCloneResult %+v", createLinkedCloneResult)
	t.Logf("LinkedClone created from (volume %s snapshot %s) successfully. volumeId: %s", volumeId, snapshotId,
		createLinkedCloneVolumeId)

	//  Clean up linkedclone created above
	var deleteLinkedCloneVolumeIDList []cnstypes.CnsVolumeId
	deleteLinkedCloneVolumeIDList = append(deleteLinkedCloneVolumeIDList, cnstypes.CnsVolumeId{Id: createLinkedCloneVolumeId})
	t.Logf("Deleting volume: %+v", deleteLinkedCloneVolumeIDList)
	deleteLinkedCloneTask, err := cnsClient.DeleteVolume(ctx, deleteLinkedCloneVolumeIDList, true)
	if err != nil {
		t.Errorf("Failed to delete linkedclone. Error: %+v \n", err)
		t.Fatal(err)
	}
	deleteLinkedCloneTaskInfo, err := GetTaskInfo(ctx, deleteLinkedCloneTask)
	if err != nil {
		t.Errorf("Failed to delete linkedclone. Error: %+v \n", err)
		t.Fatal(err)
	}
	deleteLinkedCloneTaskResult, err := GetTaskResult(ctx, deleteLinkedCloneTaskInfo)
	if err != nil {
		t.Errorf("Failed to delete linkedclone. Error: %+v \n", err)
		t.Fatal(err)
	}
	if deleteLinkedCloneTaskResult == nil {
		t.Fatalf("Empty delete task results")
		t.FailNow()
	}
	deleteLinkedCloneOperationRes := deleteLinkedCloneTaskResult.GetCnsVolumeOperationResult()
	if deleteLinkedCloneOperationRes.Fault != nil {
		t.Fatalf("Failed to delete linkedclone: fault=%+v", deleteLinkedCloneOperationRes.Fault)
	}
	t.Logf("LinkedClone: %q deleted successfully", createLinkedCloneVolumeId)

	// Test DeleteSnapshot API
	// Construct the CNS SnapshotDeleteSpec list
	var cnsSnapshotDeleteSpecList []cnstypes.CnsSnapshotDeleteSpec
	cnsSnapshotDeleteSpec := cnstypes.CnsSnapshotDeleteSpec{
		VolumeId: cnstypes.CnsVolumeId{
			Id: volumeId,
		},
		SnapshotId: cnstypes.CnsSnapshotId{
			Id: snapshotId,
		},
	}
	cnsSnapshotDeleteSpecList = append(cnsSnapshotDeleteSpecList, cnsSnapshotDeleteSpec)
	t.Logf("Deleting snapshot using the spec: %+v", pretty.Sprint(cnsSnapshotDeleteSpecList))
	deleteSnapshotsTask, err := cnsClient.DeleteSnapshots(ctx, cnsSnapshotDeleteSpecList)
	if err != nil {
		t.Errorf("Failed to get the task of DeleteSnapshots. Error: %+v \n", err)
		t.Fatal(err)
	}
	deleteSnapshotsTaskInfo, err := GetTaskInfo(ctx, deleteSnapshotsTask)
	if err != nil {
		t.Errorf("Failed to get the task info of DeleteSnapshots. Error: %+v \n", err)
		t.Fatal(err)
	}

	deleteSnapshotsTaskResult, err := GetTaskResult(ctx, deleteSnapshotsTaskInfo)
	if err != nil {
		t.Errorf("Failed to get the task result of DeleteSnapshots. Error: %+v \n", err)
		t.Fatal(err)
	}

	deleteSnapshotsOperationRes := deleteSnapshotsTaskResult.GetCnsVolumeOperationResult()
	if deleteSnapshotsOperationRes.Fault != nil {
		t.Fatalf("Failed to delete snapshots: fault=%+v", deleteSnapshotsOperationRes.Fault)
	}

	snapshotDeleteResult := any(deleteSnapshotsTaskResult).(*cnstypes.CnsSnapshotDeleteResult)
	t.Logf("snapshotDeleteResult: %+v", pretty.Sprint(snapshotCreateResult))
	t.Logf("DeleteSnapshots: Snapshot deleted successfully. volumeId: %q, snapshot id %q, opId: %q", volumeId, snapshotDeleteResult.SnapshotId, deleteSnapshotsTaskInfo.ActivationId)

	// Test Relocate API
	// Relocate API is not supported on ReleaseVSAN67u3 and ReleaseVSAN70
	// This API is available on vSphere 7.0u1 onward
	if cnsClient.Version != ReleaseVSAN67u3 && cnsClient.Version != ReleaseVSAN70 &&
		datastoreForMigration != "" {

		var migrationDS *object.Datastore
		var serviceLocatorInstance *vim25types.ServiceLocator = nil

		// Cross-VC migration.
		// This is only supported on 8.0u3 onwards.
		if remoteVcUrl != "" && isvSphereVersion80U3orAbove(ctx, c.ServiceContent.About) {
			remoteUrl, err := soap.ParseURL(remoteVcUrl)
			if err != nil {
				t.Fatal(err)
			}
			remoteVcClient, err := govmomi.NewClient(ctx, remoteUrl, true)
			if err != nil {
				t.Fatal(err)
			}
			remoteCnsClient, err := NewClient(ctx, remoteVcClient.Client)
			if err != nil {
				t.Fatal(err)
			}
			remoteFinder := find.NewFinder(remoteCnsClient.vim25Client, false)
			remoteDc, err := remoteFinder.Datacenter(ctx, remoteDatacenter)
			if err != nil {
				t.Fatal(err)
			}
			remoteFinder.SetDatacenter(remoteDc)

			migrationDS, err = remoteFinder.Datastore(ctx, datastoreForMigration)
			if err != nil {
				t.Fatal(err)
			}

			// Get ServiceLocator instance for remote VC.
			userName := remoteUrl.User.Username()
			password, _ := remoteUrl.User.Password()
			serviceLocatorInstance, err = GetServiceLocatorInstance(ctx, userName, password, remoteVcClient)
			if err != nil {
				t.Fatal(err)
			}

		} else {
			// Same VC migration
			migrationDS, err = finder.Datastore(ctx, datastoreForMigration)
			if err != nil {
				t.Fatal(err)
			}
		}

		blockVolRelocateSpec := cnstypes.CnsBlockVolumeRelocateSpec{
			CnsVolumeRelocateSpec: cnstypes.CnsVolumeRelocateSpec{
				VolumeId: cnstypes.CnsVolumeId{
					Id: volumeId,
				},
				Datastore: migrationDS.Reference(),
			},
		}
		if serviceLocatorInstance != nil {
			blockVolRelocateSpec.ServiceLocator = serviceLocatorInstance
		}

		t.Logf("Relocating volume using the spec: %+v", pretty.Sprint(blockVolRelocateSpec))

		relocateTask, err := cnsClient.RelocateVolume(ctx, blockVolRelocateSpec)
		if err != nil {
			t.Errorf("Failed to migrate volume with Relocate API. Error: %+v \n", err)
			t.Fatal(err)
		}
		relocateTaskInfo, err := GetTaskInfo(ctx, relocateTask)
		if err != nil {
			t.Errorf("Failed to get info of task returned by Relocate API. Error: %+v \n", err)
			t.Fatal(err)
		}
		taskResults, err := GetTaskResultArray(ctx, relocateTaskInfo)
		if err != nil {
			t.Fatal(err)
		}
		for _, taskResult := range taskResults {
			res := taskResult.GetCnsVolumeOperationResult()
			if res.Fault != nil {
				t.Fatalf("Relocation failed due to fault: %+v", res.Fault)
			}
			t.Logf("Successfully Relocated volume. Relocate task info result: %+v", pretty.Sprint(taskResult))
		}
	}

	// Test ExtendVolume API
	var newCapacityInMb int64 = 10240
	var cnsVolumeExtendSpecList []cnstypes.CnsVolumeExtendSpec
	cnsVolumeExtendSpec := cnstypes.CnsVolumeExtendSpec{
		VolumeId: cnstypes.CnsVolumeId{
			Id: volumeId,
		},
		CapacityInMb: newCapacityInMb,
	}
	cnsVolumeExtendSpecList = append(cnsVolumeExtendSpecList, cnsVolumeExtendSpec)
	t.Logf("Extending volume using the spec: %+v", pretty.Sprint(cnsVolumeExtendSpecList))
	extendTask, err := cnsClient.ExtendVolume(ctx, cnsVolumeExtendSpecList)
	if err != nil {
		t.Errorf("Failed to extend volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	extendTaskInfo, err := GetTaskInfo(ctx, extendTask)
	if err != nil {
		t.Errorf("Failed to extend volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	extendTaskResult, err := GetTaskResult(ctx, extendTaskInfo)
	if err != nil {
		t.Errorf("Failed to extend volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if extendTaskResult == nil {
		t.Fatalf("Empty extend task results")
		t.FailNow()
	}
	extendVolumeOperationRes := extendTaskResult.GetCnsVolumeOperationResult()
	if extendVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to extend volume: fault=%+v", extendVolumeOperationRes.Fault)
	}
	extendVolumeId := extendVolumeOperationRes.VolumeId.Id
	t.Logf("Volume extended successfully. Volume ID: %s", extendVolumeId)

	// Verify volume is extended to the specified size
	t.Logf("Calling QueryVolume after ExtendVolume using queryFilter: %+v", queryFilter)
	queryResult, err = cnsClient.QueryVolume(ctx, &queryFilter)
	if err != nil {
		t.Errorf("Failed to query volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	t.Logf("Successfully Queried Volumes after ExtendVolume. queryResult: %+v", pretty.Sprint(queryResult))
	queryCapacity := queryResult.Volumes[0].BackingObjectDetails.(*cnstypes.CnsBlockBackingDetails).CapacityInMb
	if newCapacityInMb != queryCapacity {
		t.Errorf("After extend volume %s, expected new volume size is %d, but actual volume size is %d.", extendVolumeId, newCapacityInMb, queryCapacity)
	} else {
		t.Logf("Volume extended successfully to the new size. Volume ID: %s New Size: %d", extendVolumeId, newCapacityInMb)
	}

	// Test UpdateVolumeMetadata
	var updateSpecList []cnstypes.CnsVolumeMetadataUpdateSpec

	var metadataList []cnstypes.BaseCnsEntityMetadata
	newLabels := []vim25types.KeyValue{
		{
			Key:   "testLabel",
			Value: "testValue",
		},
	}
	pvmetadata := &cnstypes.CnsKubernetesEntityMetadata{
		CnsEntityMetadata: cnstypes.CnsEntityMetadata{
			DynamicData: vim25types.DynamicData{},
			EntityName:  "pvc-53465372-5c12-4818-96f8-0ace4f4fd116",
			Labels:      newLabels,
			Delete:      false,
			ClusterID:   "demo-cluster-id",
		},
		EntityType: string(cnstypes.CnsKubernetesEntityTypePV),
		Namespace:  "",
	}
	metadataList = append(metadataList, cnstypes.BaseCnsEntityMetadata(pvmetadata))

	pvcmetadata := &cnstypes.CnsKubernetesEntityMetadata{
		CnsEntityMetadata: cnstypes.CnsEntityMetadata{
			DynamicData: vim25types.DynamicData{},
			EntityName:  "example-vanilla-block-pvc",
			Labels:      newLabels,
			Delete:      false,
			ClusterID:   "demo-cluster-id",
		},
		EntityType: string(cnstypes.CnsKubernetesEntityTypePVC),
		Namespace:  "default",
		ReferredEntity: []cnstypes.CnsKubernetesEntityReference{
			{
				EntityType: string(cnstypes.CnsKubernetesEntityTypePV),
				EntityName: "pvc-53465372-5c12-4818-96f8-0ace4f4fd116",
				Namespace:  "",
				ClusterID:  "demo-cluster-id",
			},
		},
	}
	metadataList = append(metadataList, cnstypes.BaseCnsEntityMetadata(pvcmetadata))

	podmetadata := &cnstypes.CnsKubernetesEntityMetadata{
		CnsEntityMetadata: cnstypes.CnsEntityMetadata{
			DynamicData: vim25types.DynamicData{},
			EntityName:  "example-pod",
			Delete:      false,
			ClusterID:   "demo-cluster-id",
		},
		EntityType: string(cnstypes.CnsKubernetesEntityTypePOD),
		Namespace:  "default",
		ReferredEntity: []cnstypes.CnsKubernetesEntityReference{
			{
				EntityType: string(cnstypes.CnsKubernetesEntityTypePVC),
				EntityName: "example-vanilla-block-pvc",
				Namespace:  "default",
				ClusterID:  "demo-cluster-id",
			},
		},
	}
	metadataList = append(metadataList, cnstypes.BaseCnsEntityMetadata(podmetadata))

	cnsVolumeMetadataUpdateSpec := cnstypes.CnsVolumeMetadataUpdateSpec{
		VolumeId: cnstypes.CnsVolumeId{Id: volumeId},
		Metadata: cnstypes.CnsVolumeMetadata{
			DynamicData:           vim25types.DynamicData{},
			ContainerCluster:      containerCluster,
			EntityMetadata:        metadataList,
			ContainerClusterArray: containerClusterArray,
		},
	}
	t.Logf("Updating volume using the spec: %+v", cnsVolumeMetadataUpdateSpec)
	updateSpecList = append(updateSpecList, cnsVolumeMetadataUpdateSpec)
	updateTask, err := cnsClient.UpdateVolumeMetadata(ctx, updateSpecList)
	if err != nil {
		t.Errorf("Failed to update volume metadata. Error: %+v \n", err)
		t.Fatal(err)
	}
	updateTaskInfo, err := GetTaskInfo(ctx, updateTask)
	if err != nil {
		t.Errorf("Failed to update volume metadata. Error: %+v \n", err)
		t.Fatal(err)
	}
	updateTaskResult, err := GetTaskResult(ctx, updateTaskInfo)
	if err != nil {
		t.Errorf("Failed to update volume metadata. Error: %+v \n", err)
		t.Fatal(err)
	}
	if updateTaskResult == nil {
		t.Fatalf("Empty update task results")
		t.FailNow()
	}
	updateVolumeOperationRes := updateTaskResult.GetCnsVolumeOperationResult()
	if updateVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to update volume metadata: fault=%+v", updateVolumeOperationRes.Fault)
	} else {
		t.Logf("Successfully updated volume metadata")
	}

	t.Logf("Calling QueryVolume using queryFilter: %+v", pretty.Sprint(queryFilter))
	queryResult, err = cnsClient.QueryVolume(ctx, &queryFilter)
	if err != nil {
		t.Errorf("Failed to query volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	t.Logf("Successfully Queried Volumes. queryResult: %+v", pretty.Sprint(queryResult))

	// Test QueryVolume API with CnsKubernetesQueryFilter
	var k8sQueryFilter cnstypes.CnsKubernetesQueryFilter
	k8sQueryFilter.Namespaces = []string{"default"}
	k8sQueryFilter.CnsQueryFilter.ContainerClusterIds = []string{"demo-cluster-id"}
	t.Logf("Calling QueryVolume using k8squeryFilter: %+v", pretty.Sprint(k8sQueryFilter))
	queryResult, err = cnsClient.QueryVolume(ctx, &k8sQueryFilter)
	if err != nil {
		t.Errorf("Failed to query volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if len(queryResult.Volumes) > 0 {
		t.Logf("Successfully Queried Volumes using k8squeryFilter. queryResult: %+v", pretty.Sprint(queryResult))
	} else {
		t.Fatal("failed to query volumes using k8squeryFilter")
	}
	// Test QueryAll
	querySelection := cnstypes.CnsQuerySelection{
		Names: []string{
			string(cnstypes.CnsQuerySelectionName_VOLUME_NAME),
			string(cnstypes.CnsQuerySelectionName_VOLUME_TYPE),
			string(cnstypes.CnsQuerySelectionName_DATASTORE_URL),
			string(cnstypes.CnsQuerySelectionName_POLICY_ID),
			string(cnstypes.CnsQuerySelectionName_HEALTH_STATUS),
			string(cnstypes.CnsQuerySelectionName_BACKING_OBJECT_DETAILS),
			string(cnstypes.CnsQuerySelectionName_COMPLIANCE_STATUS),
			string(cnstypes.CnsQuerySelectionName_DATASTORE_ACCESSIBILITY_STATUS),
		},
	}
	queryResult, err = cnsClient.QueryAllVolume(ctx, cnstypes.CnsQueryFilter{}, querySelection)
	if err != nil {
		t.Errorf("Failed to query all volumes. Error: %+v \n", err)
		t.Fatal(err)
	}
	t.Logf("Successfully Queried all Volumes. queryResult: %+v", pretty.Sprint(queryResult))

	// Create a VM to test Attach Volume API.
	virtualMachineConfigSpec := vim25types.VirtualMachineConfigSpec{
		Name: "test-node-vm",
		Files: &vim25types.VirtualMachineFileInfo{
			VmPathName: "[" + datastore + "]",
		},
		NumCPUs:  1,
		MemoryMB: 4,
		DeviceChange: []vim25types.BaseVirtualDeviceConfigSpec{
			&vim25types.VirtualDeviceConfigSpec{
				Operation: vim25types.VirtualDeviceConfigSpecOperationAdd,
				Device: &vim25types.ParaVirtualSCSIController{
					VirtualSCSIController: vim25types.VirtualSCSIController{
						SharedBus: vim25types.VirtualSCSISharingNoSharing,
						VirtualController: vim25types.VirtualController{
							BusNumber: 0,
							VirtualDevice: vim25types.VirtualDevice{
								Key: 1000,
							},
						},
					},
				},
			},
		},
	}
	defaultFolder, err := finder.DefaultFolder(ctx)
	if err != nil {
		t.Fatal(err)
	}
	var resourcePool *object.ResourcePool
	if resourcePoolPath == "" {
		resourcePool, err = finder.DefaultResourcePool(ctx)
	} else {
		resourcePool, err = finder.ResourcePool(ctx, resourcePoolPath)
	}
	if err != nil {
		t.Errorf("Error occurred while getting DefaultResourcePool. err: %+v", err)
		t.Fatal(err)
	}
	task, err := defaultFolder.CreateVM(ctx, virtualMachineConfigSpec, resourcePool, nil)
	if err != nil {
		t.Errorf("Failed to create VM. Error: %+v \n", err)
		t.Fatal(err)
	}

	vmTaskInfo, err := task.WaitForResult(ctx, nil)
	if err != nil {
		t.Errorf("Error occurred while waiting for create VM task result. err: %+v", err)
		t.Fatal(err)
	}

	vmRef := vmTaskInfo.Result.(object.Reference)
	t.Logf("Node VM created successfully. vmRef: %+v", vmRef.Reference())

	nodeVM := object.NewVirtualMachine(cnsClient.vim25Client, vmRef.Reference())
	defer nodeVM.Destroy(ctx)

	// Test AttachVolume API
	var cnsVolumeAttachSpecList []cnstypes.CnsVolumeAttachDetachSpec
	cnsVolumeAttachSpec := cnstypes.CnsVolumeAttachDetachSpec{
		VolumeId: cnstypes.CnsVolumeId{
			Id: volumeId,
		},
		Vm: nodeVM.Reference(),
	}
	cnsVolumeAttachSpecList = append(cnsVolumeAttachSpecList, cnsVolumeAttachSpec)
	t.Logf("Attaching volume using the spec: %+v", cnsVolumeAttachSpec)
	attachTask, err := cnsClient.AttachVolume(ctx, cnsVolumeAttachSpecList)
	if err != nil {
		t.Errorf("Failed to attach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	attachTaskInfo, err := GetTaskInfo(ctx, attachTask)
	if err != nil {
		t.Errorf("Failed to attach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	attachTaskResult, err := GetTaskResult(ctx, attachTaskInfo)
	if err != nil {
		t.Errorf("Failed to attach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if attachTaskResult == nil {
		t.Fatalf("Empty attach task results")
		t.FailNow()
	}
	attachVolumeOperationRes := attachTaskResult.GetCnsVolumeOperationResult()
	if attachVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to attach volume: fault=%+v", attachVolumeOperationRes.Fault)
	}
	diskUUID := any(attachTaskResult).(*cnstypes.CnsVolumeAttachResult).DiskUUID
	t.Logf("Volume attached successfully. Disk UUID: %s", diskUUID)

	// Re-Attach same volume to the same node and expect ResourceInUse fault
	t.Logf("Re-Attaching volume using the spec: %+v", cnsVolumeAttachSpec)
	attachTask, err = cnsClient.AttachVolume(ctx, cnsVolumeAttachSpecList)
	if err != nil {
		t.Errorf("Failed to attach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	attachTaskInfo, err = GetTaskInfo(ctx, attachTask)
	if err != nil {
		t.Errorf("Failed to attach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	attachTaskResult, err = GetTaskResult(ctx, attachTaskInfo)
	if err != nil {
		t.Errorf("Failed to attach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if attachTaskResult == nil {
		t.Fatalf("Empty attach task results")
		t.FailNow()
	}
	reAttachVolumeOperationRes := attachTaskResult.GetCnsVolumeOperationResult()
	if isvSphereVersion91orAbove {
		if reAttachVolumeOperationRes.Fault != nil {
			t.Fatalf("re-attach same volume should not fail with ResourceInUse fault")
		}
	} else {
		if reAttachVolumeOperationRes.Fault != nil {
			t.Fatalf("re-attach same volume should not fail with ResourceInUse fault")
			t.Logf("reAttachVolumeOperationRes.Fault: %+v", pretty.Sprint(reAttachVolumeOperationRes.Fault))
			_, ok := reAttachVolumeOperationRes.Fault.Fault.(*vim25types.ResourceInUse)
			if !ok {
				t.Fatalf("Fault is not ResourceInUse")
			}
		} else {
			t.Fatalf("re-attach same volume should fail with ResourceInUse fault")
		}
	}

	// Test DetachVolume API
	var cnsVolumeDetachSpecList []cnstypes.CnsVolumeAttachDetachSpec
	cnsVolumeDetachSpec := cnstypes.CnsVolumeAttachDetachSpec{
		VolumeId: cnstypes.CnsVolumeId{
			Id: volumeId,
		},
		Vm: nodeVM.Reference(),
	}
	cnsVolumeDetachSpecList = append(cnsVolumeDetachSpecList, cnsVolumeDetachSpec)
	t.Logf("Detaching volume using the spec: %+v", cnsVolumeDetachSpec)
	detachTask, err := cnsClient.DetachVolume(ctx, cnsVolumeDetachSpecList)
	if err != nil {
		t.Errorf("Failed to detach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	detachTaskInfo, err := GetTaskInfo(ctx, detachTask)
	if err != nil {
		t.Errorf("Failed to detach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	detachTaskResult, err := GetTaskResult(ctx, detachTaskInfo)
	if err != nil {
		t.Errorf("Failed to detach volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if detachTaskResult == nil {
		t.Fatalf("Empty detach task results")
		t.FailNow()
	}
	detachVolumeOperationRes := detachTaskResult.GetCnsVolumeOperationResult()
	if detachVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to detach volume: fault=%+v", detachVolumeOperationRes.Fault)
	}
	t.Logf("Volume detached successfully")

	if run_shared_disk_tests_tests == "true" && isvSphereVersion91orAbove {
		t.Logf("Running shared disk tests")

		// Test AttachVolume API with multiwriter parameters
		var cnsVolumeAttachSpecList []cnstypes.CnsVolumeAttachDetachSpec
		cnsVolumeAttachSpec := cnstypes.CnsVolumeAttachDetachSpec{
			VolumeId: cnstypes.CnsVolumeId{
				Id: volumeId,
			},
			Vm:            nodeVM.Reference(),
			DiskMode:      "independent_persistent",
			Sharing:       "sharingMultiWriter",
			UnitNumber:    vim25types.NewInt32(23),
			ControllerKey: vim25types.NewInt32(1002),
		}

		cnsVolumeAttachSpecList = append(cnsVolumeAttachSpecList, cnsVolumeAttachSpec)
		t.Logf("Attaching volume using the spec: %+v", cnsVolumeAttachSpec)
		attachTask, err := cnsClient.AttachVolume(ctx, cnsVolumeAttachSpecList)
		if err != nil {
			t.Errorf("Failed to attach volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		attachTaskInfo, err := GetTaskInfo(ctx, attachTask)
		if err != nil {
			t.Errorf("Failed to attach volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		attachTaskResult, err := GetTaskResult(ctx, attachTaskInfo)
		if err != nil {
			t.Errorf("Failed to attach volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		if attachTaskResult == nil {
			t.Fatalf("Empty attach task results")
			t.FailNow()
		}
		attachVolumeOperationRes := attachTaskResult.GetCnsVolumeOperationResult()
		if attachVolumeOperationRes.Fault != nil {
			t.Fatalf("Failed to attach volume: fault=%+v", attachVolumeOperationRes.Fault)
		}
		diskUUID := any(attachTaskResult).(*cnstypes.CnsVolumeAttachResult).DiskUUID
		t.Logf("Volume attached sucessfully with shared disk params. Disk UUID: %s", diskUUID)

		// Test DetachVolume API
		var cnsVolumeDetachSpecList []cnstypes.CnsVolumeAttachDetachSpec
		cnsVolumeDetachSpec := cnstypes.CnsVolumeAttachDetachSpec{
			VolumeId: cnstypes.CnsVolumeId{
				Id: volumeId,
			},
			Vm: nodeVM.Reference(),
		}
		cnsVolumeDetachSpecList = append(cnsVolumeDetachSpecList, cnsVolumeDetachSpec)
		t.Logf("Detaching volume using the spec: %+v", cnsVolumeDetachSpec)
		detachTask, err := cnsClient.DetachVolume(ctx, cnsVolumeDetachSpecList)
		if err != nil {
			t.Errorf("Failed to detach volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		detachTaskInfo, err := GetTaskInfo(ctx, detachTask)
		if err != nil {
			t.Errorf("Failed to detach volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		detachTaskResult, err := GetTaskResult(ctx, detachTaskInfo)
		if err != nil {
			t.Errorf("Failed to detach volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		if detachTaskResult == nil {
			t.Fatalf("Empty detach task results")
			t.FailNow()
		}
		detachVolumeOperationRes := detachTaskResult.GetCnsVolumeOperationResult()
		if detachVolumeOperationRes.Fault != nil {
			t.Fatalf("Failed to detach volume: fault=%+v", detachVolumeOperationRes.Fault)
		}
		t.Logf("Volume detached sucessfully")
	}

	// Test QueryVolumeAsync API only for vSphere version 7.0.3 onwards
	if isvSphereVersion70U3orAbove(ctx, c.ServiceContent.About) {
		queryVolumeAsyncTask, err := cnsClient.QueryVolumeAsync(ctx, queryFilter, nil)
		if err != nil {
			t.Errorf("Failed to query volumes with QueryVolumeAsync. Error: %+v \n", err)
		}
		queryVolumeAsyncTaskInfo, err := GetTaskInfo(ctx, queryVolumeAsyncTask)
		if err != nil {
			t.Errorf("Failed to query volumes with QueryVolumeAsync. Error: %+v \n", err)
		}
		queryVolumeAsyncTaskResults, err := GetTaskResultArray(ctx, queryVolumeAsyncTaskInfo)
		if err != nil {
			t.Errorf("Failed to query volumes with QueryVolumeAsync. Error: %+v \n", err)
		}
		for _, queryVolumeAsyncTaskResult := range queryVolumeAsyncTaskResults {
			queryVolumeAsyncOperationRes := queryVolumeAsyncTaskResult.GetCnsVolumeOperationResult()
			if queryVolumeAsyncOperationRes.Fault != nil {
				t.Fatalf("Failed to query volumes with QueryVolumeAsync. fault=%+v", queryVolumeAsyncOperationRes.Fault)
			}
			t.Logf("Successfully queried Volume using queryAsync API. queryVolumeAsyncTaskResult: %+v", pretty.Sprint(queryVolumeAsyncTaskResult))
		}
	}

	// Test DeleteVolume API
	t.Logf("Deleting volume: %+v", volumeIDList)
	deleteTask, err := cnsClient.DeleteVolume(ctx, volumeIDList, true)
	if err != nil {
		t.Errorf("Failed to delete volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	deleteTaskInfo, err := GetTaskInfo(ctx, deleteTask)
	if err != nil {
		t.Errorf("Failed to delete volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	deleteTaskResult, err := GetTaskResult(ctx, deleteTaskInfo)
	if err != nil {
		t.Errorf("Failed to delete volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if deleteTaskResult == nil {
		t.Fatalf("Empty delete task results")
		t.FailNow()
	}
	deleteVolumeOperationRes := deleteTaskResult.GetCnsVolumeOperationResult()
	if deleteVolumeOperationRes.Fault != nil {
		t.Fatalf("Failed to delete volume: fault=%+v", deleteVolumeOperationRes.Fault)
	}
	t.Logf("Volume: %q deleted successfully", volumeId)

	if run_fileshare_tests == "true" && cnsClient.Version != ReleaseVSAN67u3 {
		// Test creating vSAN file-share Volume
		var cnsFileVolumeCreateSpecList []cnstypes.CnsVolumeCreateSpec
		vSANFileCreateSpec := &cnstypes.CnsVSANFileCreateSpec{
			SoftQuotaInMb: 5120,
			Permission: []vsanfstypes.VsanFileShareNetPermission{
				{
					Ips:         "*",
					Permissions: vsanfstypes.VsanFileShareAccessTypeREAD_WRITE,
					AllowRoot:   true,
				},
			},
		}

		cnsFileVolumeCreateSpec := cnstypes.CnsVolumeCreateSpec{
			Name:       "pvc-file-share-volume",
			VolumeType: string(cnstypes.CnsVolumeTypeFile),
			Datastores: dsList,
			Metadata: cnstypes.CnsVolumeMetadata{
				ContainerCluster:      containerCluster,
				ContainerClusterArray: containerClusterArray,
			},
			BackingObjectDetails: &cnstypes.CnsVsanFileShareBackingDetails{
				CnsFileBackingDetails: cnstypes.CnsFileBackingDetails{
					CnsBackingObjectDetails: cnstypes.CnsBackingObjectDetails{
						CapacityInMb: 5120,
					},
				},
			},
			CreateSpec: vSANFileCreateSpec,
		}
		cnsFileVolumeCreateSpecList = append(cnsFileVolumeCreateSpecList, cnsFileVolumeCreateSpec)
		t.Logf("Creating CNS file volume using the spec: %+v", cnsFileVolumeCreateSpec)
		createTask, err = cnsClient.CreateVolume(ctx, cnsFileVolumeCreateSpecList)
		if err != nil {
			t.Errorf("Failed to create vsan fileshare volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		createTaskInfo, err = GetTaskInfo(ctx, createTask)
		if err != nil {
			t.Errorf("Failed to create Fileshare volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		createTaskResult, err = GetTaskResult(ctx, createTaskInfo)
		if err != nil {
			t.Errorf("Failed to create Fileshare volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		if createTaskResult == nil {
			t.Fatalf("Empty create task results")
			t.FailNow()
		}
		createVolumeOperationRes = createTaskResult.GetCnsVolumeOperationResult()
		if createVolumeOperationRes.Fault != nil {
			t.Fatalf("Failed to create Fileshare volume: fault=%+v", createVolumeOperationRes.Fault)
		}
		filevolumeId := createVolumeOperationRes.VolumeId.Id
		t.Logf("Fileshare volume created successfully. filevolumeId: %s", filevolumeId)

		// Test QueryVolume API
		volumeIDList = []cnstypes.CnsVolumeId{{Id: filevolumeId}}
		queryFilter.VolumeIds = volumeIDList
		t.Logf("Calling QueryVolume using queryFilter: %+v", queryFilter)
		queryResult, err = cnsClient.QueryVolume(ctx, &queryFilter)
		if err != nil {
			t.Errorf("Failed to query volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		t.Logf("Successfully Queried Volumes. queryResult: %+v", queryResult)
		fileBackingInfo := queryResult.Volumes[0].BackingObjectDetails.(*cnstypes.CnsVsanFileShareBackingDetails)
		t.Logf("File Share Name: %s with accessPoints: %+v", fileBackingInfo.Name, fileBackingInfo.AccessPoints)

		// Test add read-only permissions using Configure ACLs
		netPerms := make([]vsanfstypes.VsanFileShareNetPermission, 0)
		netPerms = append(netPerms, vsanfstypes.VsanFileShareNetPermission{
			Ips:         "192.168.124.2",
			Permissions: "READ_ONLY",
		})

		vSanNFSACLEntry := make([]cnstypes.CnsNFSAccessControlSpec, 0)
		vSanNFSACLEntry = append(vSanNFSACLEntry, cnstypes.CnsNFSAccessControlSpec{
			Permission: netPerms,
		})

		volumeID := cnstypes.CnsVolumeId{
			Id: filevolumeId,
		}
		aclSpec := cnstypes.CnsVolumeACLConfigureSpec{
			VolumeId:              volumeID,
			AccessControlSpecList: vSanNFSACLEntry,
		}
		t.Logf("Invoking ConfigureVolumeACLs using the spec: %+v", pretty.Sprint(aclSpec))
		aclTask, err := cnsClient.ConfigureVolumeACLs(ctx, aclSpec)
		if err != nil {
			t.Errorf("Failed to configure VolumeACLs. Error: %+v", err)
			t.Fatal(err)
		}
		aclTaskInfo, err := GetTaskInfo(ctx, aclTask)
		if err != nil {
			t.Errorf("Failed to configure VolumeACLs. Error: %+v", err)
			t.Fatal(err)
		}
		aclTaskResult, err := GetTaskResult(ctx, aclTaskInfo)
		if err != nil {
			t.Errorf("Failed to configure VolumeACLs. Error: %+v", err)
			t.Fatal(err)
		}
		if aclTaskResult == nil {
			t.Fatalf("Empty configure VolumeACLs task results")
			t.FailNow()
		}

		// Test to revoke all permissions using Configure ACLs
		netPerms = make([]vsanfstypes.VsanFileShareNetPermission, 0)
		netPerms = append(netPerms, vsanfstypes.VsanFileShareNetPermission{
			Ips:         "192.168.124.2",
			Permissions: "READ_ONLY",
		})

		vSanNFSACLEntry = make([]cnstypes.CnsNFSAccessControlSpec, 0)
		vSanNFSACLEntry = append(vSanNFSACLEntry, cnstypes.CnsNFSAccessControlSpec{
			Permission: netPerms,
			Delete:     true,
		})

		aclSpec = cnstypes.CnsVolumeACLConfigureSpec{
			VolumeId:              volumeID,
			AccessControlSpecList: vSanNFSACLEntry,
		}
		t.Logf("Invoking ConfigureVolumeACLs using the spec: %+v", pretty.Sprint(aclSpec))
		aclTask, err = cnsClient.ConfigureVolumeACLs(ctx, aclSpec)
		if err != nil {
			t.Errorf("Failed to configure VolumeACLs. Error: %+v", err)
			t.Fatal(err)
		}
		aclTaskInfo, err = GetTaskInfo(ctx, aclTask)
		if err != nil {
			t.Errorf("Failed to configure VolumeACLs. Error: %+v", err)
			t.Fatal(err)
		}
		aclTaskResult, err = GetTaskResult(ctx, aclTaskInfo)
		if err != nil {
			t.Errorf("Failed to configure VolumeACLs. Error: %+v", err)
			t.Fatal(err)
		}
		if aclTaskResult == nil {
			t.Fatalf("Empty configure VolumeACLs task results")
			t.FailNow()
		}

		// Test Deleting vSAN file-share Volume
		var fileVolumeIDList []cnstypes.CnsVolumeId
		fileVolumeIDList = append(fileVolumeIDList, cnstypes.CnsVolumeId{Id: filevolumeId})
		t.Logf("Deleting fileshare volume: %+v", fileVolumeIDList)
		deleteTask, err = cnsClient.DeleteVolume(ctx, fileVolumeIDList, true)
		if err != nil {
			t.Errorf("Failed to delete fileshare volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		deleteTaskInfo, err = GetTaskInfo(ctx, deleteTask)
		if err != nil {
			t.Errorf("Failed to delete fileshare volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		deleteTaskResult, err = GetTaskResult(ctx, deleteTaskInfo)
		if err != nil {
			t.Errorf("Failed to delete fileshare volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		if deleteTaskResult == nil {
			t.Fatalf("Empty delete task results")
			t.FailNow()
		}
		deleteVolumeOperationRes = deleteTaskResult.GetCnsVolumeOperationResult()
		if deleteVolumeOperationRes.Fault != nil {
			t.Fatalf("Failed to delete fileshare volume: fault=%+v", deleteVolumeOperationRes.Fault)
		}
		t.Logf("fileshare volume:%q deleted successfully", filevolumeId)
	}
	if backingDiskURLPath != "" && cnsClient.Version != ReleaseVSAN67u3 && cnsClient.Version != ReleaseVSAN70 {
		// Test CreateVolume API with existing VMDK
		var cnsVolumeCreateSpecList []cnstypes.CnsVolumeCreateSpec
		cnsVolumeCreateSpec := cnstypes.CnsVolumeCreateSpec{
			Name:       "pvc-901e87eb-c2bd-11e9-806f-005056a0c9a0",
			VolumeType: string(cnstypes.CnsVolumeTypeBlock),
			Metadata: cnstypes.CnsVolumeMetadata{
				ContainerCluster: containerCluster,
			},
			BackingObjectDetails: &cnstypes.CnsBlockBackingDetails{
				BackingDiskUrlPath: backingDiskURLPath,
			},
		}
		cnsVolumeCreateSpecList = append(cnsVolumeCreateSpecList, cnsVolumeCreateSpec)
		t.Logf("Creating volume using the spec: %+v", pretty.Sprint(cnsVolumeCreateSpec))
		createTask, err := cnsClient.CreateVolume(ctx, cnsVolumeCreateSpecList)
		if err != nil {
			t.Errorf("Failed to create volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		createTaskInfo, err := GetTaskInfo(ctx, createTask)
		if err != nil {
			t.Errorf("Failed to create volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		createTaskResult, err := GetTaskResult(ctx, createTaskInfo)
		if err != nil {
			t.Errorf("Failed to create volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		if createTaskResult == nil {
			t.Fatalf("Empty create task results")
			t.FailNow()
		}
		createVolumeOperationRes := createTaskResult.GetCnsVolumeOperationResult()
		var volumeID string
		if createVolumeOperationRes.Fault != nil {
			t.Logf("Failed to create volume: fault=%+v", createVolumeOperationRes.Fault)
			fault, ok := createVolumeOperationRes.Fault.Fault.(*cnstypes.CnsAlreadyRegisteredFault)
			if !ok {
				t.Fatalf("Fault is not CnsAlreadyRegisteredFault")
			} else {
				t.Logf("Fault is CnsAlreadyRegisteredFault. backingDiskURLPath: %s is already registered", backingDiskURLPath)
				volumeID = fault.VolumeId.Id
			}
		} else {
			volumeID = createVolumeOperationRes.VolumeId.Id
			t.Logf("Volume created successfully with backingDiskURLPath: %s. volumeId: %s", backingDiskURLPath, volumeID)

			// Test re creating volume using BACKING_DISK_URL_PATH
			var reCreateCnsVolumeCreateSpecList []cnstypes.CnsVolumeCreateSpec
			reCreateCnsVolumeCreateSpec := cnstypes.CnsVolumeCreateSpec{
				Name:       "pvc-901e87eb-c2bd-11e9-806f-005056a0c9a0",
				VolumeType: string(cnstypes.CnsVolumeTypeBlock),
				Metadata: cnstypes.CnsVolumeMetadata{
					ContainerCluster: containerCluster,
				},
				BackingObjectDetails: &cnstypes.CnsBlockBackingDetails{
					BackingDiskUrlPath: backingDiskURLPath,
				},
			}

			reCreateCnsVolumeCreateSpecList = append(reCreateCnsVolumeCreateSpecList, reCreateCnsVolumeCreateSpec)
			t.Logf("Creating volume using the spec: %+v", pretty.Sprint(reCreateCnsVolumeCreateSpec))
			recreateTask, err := cnsClient.CreateVolume(ctx, reCreateCnsVolumeCreateSpecList)
			if err != nil {
				t.Errorf("Failed to create volume. Error: %+v \n", err)
				t.Fatal(err)
			}
			reCreateTaskInfo, err := GetTaskInfo(ctx, recreateTask)
			if err != nil {
				t.Errorf("Failed to create volume. Error: %+v \n", err)
				t.Fatal(err)
			}
			reCreateTaskResult, err := GetTaskResult(ctx, reCreateTaskInfo)
			if err != nil {
				t.Errorf("Failed to create volume. Error: %+v \n", err)
				t.Fatal(err)
			}
			if reCreateTaskResult == nil {
				t.Fatalf("Empty create task results")
				t.FailNow()
			}
			reCreateVolumeOperationRes := reCreateTaskResult.GetCnsVolumeOperationResult()
			t.Logf("reCreateVolumeOperationRes.: %+v", pretty.Sprint(reCreateVolumeOperationRes))
			if reCreateVolumeOperationRes.Fault != nil {
				t.Logf("Failed to create volume: fault=%+v", reCreateVolumeOperationRes.Fault)
				_, ok := reCreateVolumeOperationRes.Fault.Fault.(*cnstypes.CnsAlreadyRegisteredFault)
				if !ok {
					t.Fatalf("Fault is not CnsAlreadyRegisteredFault")
				} else {
					t.Logf("Fault is CnsAlreadyRegisteredFault. backingDiskURLPath: %q is already registered", backingDiskURLPath)
				}
			}
		}

		// Test QueryVolume API
		var queryFilter cnstypes.CnsQueryFilter
		var volumeIDList []cnstypes.CnsVolumeId
		volumeIDList = append(volumeIDList, cnstypes.CnsVolumeId{Id: volumeID})
		queryFilter.VolumeIds = volumeIDList
		t.Logf("Calling QueryVolume using queryFilter: %+v", pretty.Sprint(queryFilter))
		queryResult, err := cnsClient.QueryVolume(ctx, &queryFilter)
		if err != nil {
			t.Errorf("Failed to query volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		t.Logf("Successfully Queried Volumes. queryResult: %+v", pretty.Sprint(queryResult))

		t.Logf("Deleting CNS volume created above using BACKING_DISK_URL_PATH: %s with volume: %+v", backingDiskURLPath, volumeIDList)
		deleteTask, err = cnsClient.DeleteVolume(ctx, volumeIDList, true)
		if err != nil {
			t.Errorf("Failed to delete volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		deleteTaskInfo, err = GetTaskInfo(ctx, deleteTask)
		if err != nil {
			t.Errorf("Failed to delete volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		deleteTaskResult, err = GetTaskResult(ctx, deleteTaskInfo)
		if err != nil {
			t.Errorf("Failed to delete volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		if deleteTaskResult == nil {
			t.Fatalf("Empty delete task results")
			t.FailNow()
		}
		deleteVolumeOperationRes = deleteTaskResult.GetCnsVolumeOperationResult()
		if deleteVolumeOperationRes.Fault != nil {
			t.Fatalf("Failed to delete volume: fault=%+v", deleteVolumeOperationRes.Fault)
		}
		t.Logf("volume:%q deleted successfully", volumeID)
	}

	// Test CnsReconfigVolumePolicy API
	if spbmPolicyId4Reconfig != "" {
		var cnsVolumeCreateSpecList []cnstypes.CnsVolumeCreateSpec
		cnsVolumeCreateSpec := cnstypes.CnsVolumeCreateSpec{
			Name:       "pvc-901e87eb-c2bd-11e9-806f-005056a0c9a0-1",
			VolumeType: string(cnstypes.CnsVolumeTypeBlock),
			Datastores: dsList,
			Metadata: cnstypes.CnsVolumeMetadata{
				ContainerCluster: containerCluster,
			},
			BackingObjectDetails: &cnstypes.CnsBlockBackingDetails{
				CnsBackingObjectDetails: cnstypes.CnsBackingObjectDetails{
					CapacityInMb: 5120,
				},
			},
		}
		cnsVolumeCreateSpecList = append(cnsVolumeCreateSpecList, cnsVolumeCreateSpec)
		t.Logf("Creating volume using the spec: %+v", pretty.Sprint(cnsVolumeCreateSpecList))
		createTask, err = cnsClient.CreateVolume(ctx, cnsVolumeCreateSpecList)
		if err != nil {
			t.Errorf("Failed to create volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		createTaskInfo, err = GetTaskInfo(ctx, createTask)
		if err != nil {
			t.Errorf("Failed to create volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		createTaskResult, err = GetTaskResult(ctx, createTaskInfo)
		if err != nil {
			t.Errorf("Failed to create volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		if createTaskResult == nil {
			t.Fatalf("Empty create task results")
			t.FailNow()
		}
		createVolumeOperationRes = createTaskResult.GetCnsVolumeOperationResult()
		if createVolumeOperationRes.Fault != nil {
			t.Fatalf("Failed to create volume: fault=%+v", createVolumeOperationRes.Fault)
		}
		volumeId = createVolumeOperationRes.VolumeId.Id
		volumeCreateResult = (createTaskResult).(*cnstypes.CnsVolumeCreateResult)
		t.Logf("volumeCreateResult %+v", volumeCreateResult)
		t.Logf("Volume created successfully. volumeId: %s", volumeId)

		t.Logf("Calling reconfigpolicy on volume %v with policy %+v \n", volumeId, spbmPolicyId4Reconfig)
		reconfigSpecs := []cnstypes.CnsVolumePolicyReconfigSpec{
			{
				VolumeId: createVolumeOperationRes.VolumeId,
				Profile: []vim25types.BaseVirtualMachineProfileSpec{
					&vim25types.VirtualMachineDefinedProfileSpec{
						ProfileId: spbmPolicyId4Reconfig,
					},
				},
			},
		}
		reconfigTask, err := cnsClient.ReconfigVolumePolicy(ctx, reconfigSpecs)
		if err != nil {
			t.Errorf("Failed to reconfig policy %v on volume %v. Error: %+v \n", spbmPolicyId4Reconfig, volumeId, err)
			t.Fatal(err)
		}
		reconfigTaskInfo, err := GetTaskInfo(ctx, reconfigTask)
		if err != nil {
			t.Errorf("Failed to reconfig volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		reconfigTaskResult, err := GetTaskResult(ctx, reconfigTaskInfo)
		if err != nil {
			t.Errorf("Failed to reconfig volume. Error: %+v \n", err)
			t.Fatal(err)
		}
		if reconfigTaskResult == nil {
			t.Fatalf("Empty reconfig task results")
			t.FailNow()
		}
		reconfigVolumeOperationRes := reconfigTaskResult.GetCnsVolumeOperationResult()
		if reconfigVolumeOperationRes.Fault != nil {
			t.Fatalf("Failed to reconfig volume %v with policy %v: fault=%+v",
				volumeId, spbmPolicyId4Reconfig, reconfigVolumeOperationRes.Fault)
		}
		t.Logf("reconfigpolicy on volume %v with policy %+v successful\n", volumeId, spbmPolicyId4Reconfig)
	}

	// Test CnsSyncDatastore API
	t.Logf("Calling syncDatastore on %v ...\n", dsUrl)
	syncDatastoreTask, err := cnsClient.SyncDatastore(ctx, dsUrl, false)
	if err != nil {
		t.Errorf("Failed to sync datastore %v. Error: %+v \n", dsUrl, err)
		t.Fatal(err)
	}
	syncDatastoreTaskInfo, err := GetTaskInfo(ctx, syncDatastoreTask)
	if err != nil {
		t.Errorf("Failed to get sync datastore taskInfo. Error: %+v \n", err)
		t.Fatal(err)
	}
	if syncDatastoreTaskInfo.State != vim25types.TaskInfoStateSuccess {
		t.Errorf("Failed to sync datastore. Error: %+v \n", syncDatastoreTaskInfo.Error)
		t.Fatalf("%+v", syncDatastoreTaskInfo.Error)
	}
	t.Logf("syncDatastore on %v successful\n", dsUrl)

	t.Logf("Calling syncDatastore on %v with fullsync...\n", dsUrl)
	syncDatastoreTask, err = cnsClient.SyncDatastore(ctx, dsUrl, true)
	if err != nil {
		t.Errorf("Failed to sync datastore %v with full sync. Error: %+v \n", dsUrl, err)
		t.Fatal(err)
	}
	syncDatastoreTaskInfo, err = GetTaskInfo(ctx, syncDatastoreTask)
	if err != nil {
		t.Errorf("Failed to get sync datastore taskInfo with full sync. Error: %+v \n", err)
		t.Fatal(err)
	}
	if syncDatastoreTaskInfo.State != vim25types.TaskInfoStateSuccess {
		t.Errorf("Failed to sync datastore with full sync. Error: %+v \n", syncDatastoreTaskInfo.Error)
		t.Fatalf("%+v", syncDatastoreTaskInfo.Error)
	}
	t.Logf("syncDatastore on %v with full sync successful\n", dsUrl)
}

func TestUnregisterVolume(t *testing.T) {
	ctx := context.Background()
	// Setup: create a CNS client and a volume to unregister
	url := os.Getenv("CNS_VC_URL")
	if url == "" {
		t.Skip("CNS_VC_URL is not set")
	}
	datacenter := os.Getenv("CNS_DATACENTER")
	datastore := os.Getenv("CNS_DATASTORE")
	u, err := soap.ParseURL(url)
	if err != nil {
		t.Fatal(err)
	}
	c, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		t.Fatal(err)
	}

	if !isvSphereVersion91orAbove(context.Background(), c.ServiceContent.About) {
		t.Skip("This test requires vSphere 9.1 or above")
	}

	cnsClient, err := NewClient(ctx, c.Client)
	if err != nil {
		t.Fatal(err)
	}
	finder := find.NewFinder(cnsClient.vim25Client, false)
	dc, err := finder.Datacenter(ctx, datacenter)
	if err != nil {
		t.Fatal(err)
	}
	finder.SetDatacenter(dc)
	ds, err := finder.Datastore(ctx, datastore)
	if err != nil {
		t.Fatal(err)
	}

	props := []string{"info", "summary"}
	pc := property.DefaultCollector(c.Client)
	var dsSummaries []mo.Datastore
	err = pc.Retrieve(ctx, []vim25types.ManagedObjectReference{ds.Reference()}, props, &dsSummaries)
	if err != nil {
		t.Fatal(err)
	}

	var dsList []vim25types.ManagedObjectReference
	dsList = append(dsList, ds.Reference())

	var containerClusterArray []cnstypes.CnsContainerCluster
	containerCluster := cnstypes.CnsContainerCluster{
		ClusterType:         string(cnstypes.CnsClusterTypeKubernetes),
		ClusterId:           "demo-cluster-id",
		VSphereUser:         "Administrator@vsphere.local",
		ClusterFlavor:       string(cnstypes.CnsClusterFlavorVanilla),
		ClusterDistribution: "OpenShift",
	}
	containerClusterArray = append(containerClusterArray, containerCluster)

	// Test CreateVolume API
	volumeName := "pvc-" + uuid.New().String()
	var cnsVolumeCreateSpecList []cnstypes.CnsVolumeCreateSpec
	cnsVolumeCreateSpec := cnstypes.CnsVolumeCreateSpec{
		Name:       volumeName,
		VolumeType: string(cnstypes.CnsVolumeTypeBlock),
		Metadata: cnstypes.CnsVolumeMetadata{
			ContainerCluster: containerCluster,
		},
		BackingObjectDetails: &cnstypes.CnsBlockBackingDetails{
			CnsBackingObjectDetails: cnstypes.CnsBackingObjectDetails{
				CapacityInMb: 5120,
			},
		},
		Datastores: dsList,
	}

	cnsVolumeCreateSpecList = append(cnsVolumeCreateSpecList, cnsVolumeCreateSpec)

	createTask, err := cnsClient.CreateVolume(ctx, cnsVolumeCreateSpecList)
	if err != nil {
		t.Errorf("Failed to create volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	createTaskInfo, err := GetTaskInfo(ctx, createTask)
	if err != nil {
		t.Errorf("Failed to create volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	createTaskResult, err := GetTaskResult(ctx, createTaskInfo)
	if err != nil {
		t.Errorf("Failed to create volume. Error: %+v \n", err)
		t.Fatal(err)
	}
	if createTaskResult == nil {
		t.Fatalf("Empty create task results")
		t.FailNow()
	}
	createVolumeOperationRes := createTaskResult.GetCnsVolumeOperationResult()
	if createVolumeOperationRes.Fault != nil {
		if cnsFault, ok := createVolumeOperationRes.Fault.Fault.(*cnstypes.CnsFault); ok {
			if cause := cnsFault.FaultCause; cause != nil {
				if inner, ok := cause.Fault.(*vim25types.NotSupported); ok {
					t.Logf("Caught NotSupported fault: %q", cause.LocalizedMessage)
				} else {
					t.Logf("Inner fault type: %T", inner)
				}
			}
		}
		t.Fatalf("Failed to create volume: fault=%+v", createVolumeOperationRes.Fault)
	}
	volumeId := createVolumeOperationRes.VolumeId.Id
	volumeCreateResult := (createTaskResult).(*cnstypes.CnsVolumeCreateResult)
	t.Logf("volumeCreateResult %+v", volumeCreateResult)
	t.Logf("Volume created successfully. volumeId: %s", volumeId)

	spec := []cnstypes.CnsUnregisterVolumeSpec{
		{
			VolumeId:         cnstypes.CnsVolumeId{Id: volumeId},
			TargetVolumeType: string(cnstypes.CnsUnregisterTargetVolumeTypeFCD),
		},
	}
	task, err := cnsClient.UnregisterVolume(ctx, spec)
	if err != nil {
		t.Fatalf("UnregisterVolume failed: %+v", err)
	}
	taskInfo, err := GetTaskInfo(ctx, task)
	if err != nil {
		t.Fatalf("GetTaskInfo failed: %+v", err)
	}
	taskResult, err := GetTaskResult(ctx, taskInfo)
	if err != nil {
		t.Fatalf("GetTaskResult failed: %+v", err)
	}
	if taskResult == nil {
		t.Fatalf("Empty unregister task results")
	}
	operationRes := taskResult.GetCnsVolumeOperationResult()
	if operationRes.Fault != nil {
		t.Fatalf("Failed to unregister volume: fault=%+v", operationRes.Fault)
	}
	t.Logf("Volume unregistered successfully: %s", volumeId)
}

// isvSphereVersion70U3orAbove checks if specified version is 7.0 Update 3 or higher
// The method takes aboutInfo{} as input which contains details about
// VC version, build number and so on.
// If the version is 7.0 Update 3 or higher, the method returns true, else returns false
// along with appropriate errors during failure cases
func isvSphereVersion70U3orAbove(ctx context.Context, aboutInfo vim25types.AboutInfo) bool {
	items := strings.Split(aboutInfo.Version, ".")
	version := strings.Join(items[:], "")
	// Convert version string to string, Ex: "7.0.3" becomes 703, "7.0.3.1" becomes 703
	if len(version) >= 3 {
		vSphereVersionInt, err := strconv.Atoi(version[0:3])
		if err != nil {
			return false
		}
		// Check if the current vSphere version is 7.0.3 or higher
		if vSphereVersionInt >= VSphere70u3VersionInt {
			return true
		}
	}
	// For all other versions
	return false
}

// isvSphereVersion80U3orAbove checks if specified version is 8.0 Update 3 or higher
// The method takes aboutInfo{} as input which contains details about
// VC version, build number and so on.
// If the version is 8.0 Update 3 or higher, the method returns true, else returns false
// along with appropriate errors during failure cases
func isvSphereVersion80U3orAbove(ctx context.Context, aboutInfo vim25types.AboutInfo) bool {
	items := strings.Split(aboutInfo.Version, ".")
	version := strings.Join(items[:], "")
	// Convert version string to string, Ex: "8.0.3" becomes 803, "8.0.3.1" becomes 703
	if len(version) >= 3 {
		vSphereVersionInt, err := strconv.Atoi(version[0:3])
		if err != nil {
			return false
		}
		// Check if the current vSphere version is 8.0.3 or higher
		if vSphereVersionInt >= VSphere80u3VersionInt {
			return true
		}
	}
	// For all other versions
	return false
}

// isvSphereVersion91orAbove checks if specified version is 9.1 or higher
// The method takes aboutInfo{} as input which contains details about
// VC version, build number and so on.
// If the version is 9.1 higher, the method returns true, else returns false
// along with appropriate errors during failure cases
func isvSphereVersion91orAbove(ctx context.Context, aboutInfo vim25types.AboutInfo) bool {
	items := strings.Split(aboutInfo.Version, ".")
	version := strings.Join(items[:], "")
	if len(version) >= 3 {
		vSphereVersionInt, err := strconv.Atoi(version[0:3])
		if err != nil {
			return false
		}
		// Check if the current vSphere version is 9.1.0 or higher
		if vSphereVersionInt >= VSphere91VersionInt {
			return true
		}
	}
	// For all other versions
	return false
}
