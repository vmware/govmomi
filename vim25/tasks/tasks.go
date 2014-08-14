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

package tasks

import (
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

func AddDVPortgroup(r soap.RoundTripper, req *types.AddDVPortgroup_Task) (Task, error) {
	res, err := methods.AddDVPortgroup_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func AddDisks(r soap.RoundTripper, req *types.AddDisks_Task) (Task, error) {
	res, err := methods.AddDisks_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func AddHost(r soap.RoundTripper, req *types.AddHost_Task) (Task, error) {
	res, err := methods.AddHost_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func AddStandaloneHost(r soap.RoundTripper, req *types.AddStandaloneHost_Task) (Task, error) {
	res, err := methods.AddStandaloneHost_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ApplyHostConfig(r soap.RoundTripper, req *types.ApplyHostConfig_Task) (Task, error) {
	res, err := methods.ApplyHostConfig_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ApplyStorageDrsRecommendationToPod(r soap.RoundTripper, req *types.ApplyStorageDrsRecommendationToPod_Task) (Task, error) {
	res, err := methods.ApplyStorageDrsRecommendationToPod_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ApplyStorageDrsRecommendation(r soap.RoundTripper, req *types.ApplyStorageDrsRecommendation_Task) (Task, error) {
	res, err := methods.ApplyStorageDrsRecommendation_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CheckAnswerFileStatus(r soap.RoundTripper, req *types.CheckAnswerFileStatus_Task) (Task, error) {
	res, err := methods.CheckAnswerFileStatus_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CheckCompatibility(r soap.RoundTripper, req *types.CheckCompatibility_Task) (Task, error) {
	res, err := methods.CheckCompatibility_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CheckCompliance(r soap.RoundTripper, req *types.CheckCompliance_Task) (Task, error) {
	res, err := methods.CheckCompliance_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CheckHostPatch(r soap.RoundTripper, req *types.CheckHostPatch_Task) (Task, error) {
	res, err := methods.CheckHostPatch_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CheckMigrate(r soap.RoundTripper, req *types.CheckMigrate_Task) (Task, error) {
	res, err := methods.CheckMigrate_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CheckProfileCompliance(r soap.RoundTripper, req *types.CheckProfileCompliance_Task) (Task, error) {
	res, err := methods.CheckProfileCompliance_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CheckRelocate(r soap.RoundTripper, req *types.CheckRelocate_Task) (Task, error) {
	res, err := methods.CheckRelocate_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CloneVApp(r soap.RoundTripper, req *types.CloneVApp_Task) (Task, error) {
	res, err := methods.CloneVApp_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CloneVM(r soap.RoundTripper, req *types.CloneVM_Task) (Task, error) {
	res, err := methods.CloneVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ConfigureDatastoreIORM(r soap.RoundTripper, req *types.ConfigureDatastoreIORM_Task) (Task, error) {
	res, err := methods.ConfigureDatastoreIORM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ConfigureHostCache(r soap.RoundTripper, req *types.ConfigureHostCache_Task) (Task, error) {
	res, err := methods.ConfigureHostCache_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ConfigureStorageDrsForPod(r soap.RoundTripper, req *types.ConfigureStorageDrsForPod_Task) (Task, error) {
	res, err := methods.ConfigureStorageDrsForPod_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ConfigureVFlashResourceEx(r soap.RoundTripper, req *types.ConfigureVFlashResourceEx_Task) (Task, error) {
	res, err := methods.ConfigureVFlashResourceEx_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ConsolidateVMDisks(r soap.RoundTripper, req *types.ConsolidateVMDisks_Task) (Task, error) {
	res, err := methods.ConsolidateVMDisks_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CopyDatastoreFile(r soap.RoundTripper, req *types.CopyDatastoreFile_Task) (Task, error) {
	res, err := methods.CopyDatastoreFile_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CopyVirtualDisk(r soap.RoundTripper, req *types.CopyVirtualDisk_Task) (Task, error) {
	res, err := methods.CopyVirtualDisk_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CreateChildVM(r soap.RoundTripper, req *types.CreateChildVM_Task) (Task, error) {
	res, err := methods.CreateChildVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CreateDVPortgroup(r soap.RoundTripper, req *types.CreateDVPortgroup_Task) (Task, error) {
	res, err := methods.CreateDVPortgroup_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CreateDVS(r soap.RoundTripper, req *types.CreateDVS_Task) (Task, error) {
	res, err := methods.CreateDVS_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CreateScreenshot(r soap.RoundTripper, req *types.CreateScreenshot_Task) (Task, error) {
	res, err := methods.CreateScreenshot_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CreateSecondaryVM(r soap.RoundTripper, req *types.CreateSecondaryVM_Task) (Task, error) {
	res, err := methods.CreateSecondaryVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CreateSnapshot(r soap.RoundTripper, req *types.CreateSnapshot_Task) (Task, error) {
	res, err := methods.CreateSnapshot_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CreateVM(r soap.RoundTripper, req *types.CreateVM_Task) (Task, error) {
	res, err := methods.CreateVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CreateVirtualDisk(r soap.RoundTripper, req *types.CreateVirtualDisk_Task) (Task, error) {
	res, err := methods.CreateVirtualDisk_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func CustomizeVM(r soap.RoundTripper, req *types.CustomizeVM_Task) (Task, error) {
	res, err := methods.CustomizeVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func DVPortgroupRollback(r soap.RoundTripper, req *types.DVPortgroupRollback_Task) (Task, error) {
	res, err := methods.DVPortgroupRollback_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func DVSManagerExportEntity(r soap.RoundTripper, req *types.DVSManagerExportEntity_Task) (Task, error) {
	res, err := methods.DVSManagerExportEntity_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func DVSManagerImportEntity(r soap.RoundTripper, req *types.DVSManagerImportEntity_Task) (Task, error) {
	res, err := methods.DVSManagerImportEntity_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func DVSRollback(r soap.RoundTripper, req *types.DVSRollback_Task) (Task, error) {
	res, err := methods.DVSRollback_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func DatastoreExitMaintenanceMode(r soap.RoundTripper, req *types.DatastoreExitMaintenanceMode_Task) (Task, error) {
	res, err := methods.DatastoreExitMaintenanceMode_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func DefragmentVirtualDisk(r soap.RoundTripper, req *types.DefragmentVirtualDisk_Task) (Task, error) {
	res, err := methods.DefragmentVirtualDisk_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func DeleteDatastoreFile(r soap.RoundTripper, req *types.DeleteDatastoreFile_Task) (Task, error) {
	res, err := methods.DeleteDatastoreFile_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func DeleteVirtualDisk(r soap.RoundTripper, req *types.DeleteVirtualDisk_Task) (Task, error) {
	res, err := methods.DeleteVirtualDisk_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func Destroy(r soap.RoundTripper, req *types.Destroy_Task) (Task, error) {
	res, err := methods.Destroy_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func DisableSecondaryVM(r soap.RoundTripper, req *types.DisableSecondaryVM_Task) (Task, error) {
	res, err := methods.DisableSecondaryVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func DisconnectHost(r soap.RoundTripper, req *types.DisconnectHost_Task) (Task, error) {
	res, err := methods.DisconnectHost_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func EagerZeroVirtualDisk(r soap.RoundTripper, req *types.EagerZeroVirtualDisk_Task) (Task, error) {
	res, err := methods.EagerZeroVirtualDisk_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func EnableSecondaryVM(r soap.RoundTripper, req *types.EnableSecondaryVM_Task) (Task, error) {
	res, err := methods.EnableSecondaryVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func EnterMaintenanceMode(r soap.RoundTripper, req *types.EnterMaintenanceMode_Task) (Task, error) {
	res, err := methods.EnterMaintenanceMode_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func EstimateStorageForConsolidateSnapshots(r soap.RoundTripper, req *types.EstimateStorageForConsolidateSnapshots_Task) (Task, error) {
	res, err := methods.EstimateStorageForConsolidateSnapshots_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ExitMaintenanceMode(r soap.RoundTripper, req *types.ExitMaintenanceMode_Task) (Task, error) {
	res, err := methods.ExitMaintenanceMode_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ExportAnswerFile(r soap.RoundTripper, req *types.ExportAnswerFile_Task) (Task, error) {
	res, err := methods.ExportAnswerFile_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ExtendVirtualDisk(r soap.RoundTripper, req *types.ExtendVirtualDisk_Task) (Task, error) {
	res, err := methods.ExtendVirtualDisk_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func GenerateHostProfileTaskList(r soap.RoundTripper, req *types.GenerateHostProfileTaskList_Task) (Task, error) {
	res, err := methods.GenerateHostProfileTaskList_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func GenerateLogBundles(r soap.RoundTripper, req *types.GenerateLogBundles_Task) (Task, error) {
	res, err := methods.GenerateLogBundles_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ImportCertificateForCAM(r soap.RoundTripper, req *types.ImportCertificateForCAM_Task) (Task, error) {
	res, err := methods.ImportCertificateForCAM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func InflateVirtualDisk(r soap.RoundTripper, req *types.InflateVirtualDisk_Task) (Task, error) {
	res, err := methods.InflateVirtualDisk_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func InitializeDisks(r soap.RoundTripper, req *types.InitializeDisks_Task) (Task, error) {
	res, err := methods.InitializeDisks_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func InstallHostPatchV2(r soap.RoundTripper, req *types.InstallHostPatchV2_Task) (Task, error) {
	res, err := methods.InstallHostPatchV2_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func InstallHostPatch(r soap.RoundTripper, req *types.InstallHostPatch_Task) (Task, error) {
	res, err := methods.InstallHostPatch_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func JoinDomainWithCAM(r soap.RoundTripper, req *types.JoinDomainWithCAM_Task) (Task, error) {
	res, err := methods.JoinDomainWithCAM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func JoinDomain(r soap.RoundTripper, req *types.JoinDomain_Task) (Task, error) {
	res, err := methods.JoinDomain_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func LeaveCurrentDomain(r soap.RoundTripper, req *types.LeaveCurrentDomain_Task) (Task, error) {
	res, err := methods.LeaveCurrentDomain_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func MakePrimaryVM(r soap.RoundTripper, req *types.MakePrimaryVM_Task) (Task, error) {
	res, err := methods.MakePrimaryVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func MergeDvs(r soap.RoundTripper, req *types.MergeDvs_Task) (Task, error) {
	res, err := methods.MergeDvs_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func MigrateVM(r soap.RoundTripper, req *types.MigrateVM_Task) (Task, error) {
	res, err := methods.MigrateVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func MoveDVPort(r soap.RoundTripper, req *types.MoveDVPort_Task) (Task, error) {
	res, err := methods.MoveDVPort_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func MoveDatastoreFile(r soap.RoundTripper, req *types.MoveDatastoreFile_Task) (Task, error) {
	res, err := methods.MoveDatastoreFile_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func MoveHostInto(r soap.RoundTripper, req *types.MoveHostInto_Task) (Task, error) {
	res, err := methods.MoveHostInto_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func MoveIntoFolder(r soap.RoundTripper, req *types.MoveIntoFolder_Task) (Task, error) {
	res, err := methods.MoveIntoFolder_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func MoveInto(r soap.RoundTripper, req *types.MoveInto_Task) (Task, error) {
	res, err := methods.MoveInto_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func MoveVirtualDisk(r soap.RoundTripper, req *types.MoveVirtualDisk_Task) (Task, error) {
	res, err := methods.MoveVirtualDisk_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func PerformDvsProductSpecOperation(r soap.RoundTripper, req *types.PerformDvsProductSpecOperation_Task) (Task, error) {
	res, err := methods.PerformDvsProductSpecOperation_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func PowerDownHostToStandBy(r soap.RoundTripper, req *types.PowerDownHostToStandBy_Task) (Task, error) {
	res, err := methods.PowerDownHostToStandBy_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func PowerOffVApp(r soap.RoundTripper, req *types.PowerOffVApp_Task) (Task, error) {
	res, err := methods.PowerOffVApp_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func PowerOffVM(r soap.RoundTripper, req *types.PowerOffVM_Task) (Task, error) {
	res, err := methods.PowerOffVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func PowerOnMultiVM(r soap.RoundTripper, req *types.PowerOnMultiVM_Task) (Task, error) {
	res, err := methods.PowerOnMultiVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func PowerOnVApp(r soap.RoundTripper, req *types.PowerOnVApp_Task) (Task, error) {
	res, err := methods.PowerOnVApp_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func PowerOnVM(r soap.RoundTripper, req *types.PowerOnVM_Task) (Task, error) {
	res, err := methods.PowerOnVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func PowerUpHostFromStandBy(r soap.RoundTripper, req *types.PowerUpHostFromStandBy_Task) (Task, error) {
	res, err := methods.PowerUpHostFromStandBy_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func PromoteDisks(r soap.RoundTripper, req *types.PromoteDisks_Task) (Task, error) {
	res, err := methods.PromoteDisks_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func QueryHostPatch(r soap.RoundTripper, req *types.QueryHostPatch_Task) (Task, error) {
	res, err := methods.QueryHostPatch_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func QueryVMotionCompatibilityEx(r soap.RoundTripper, req *types.QueryVMotionCompatibilityEx_Task) (Task, error) {
	res, err := methods.QueryVMotionCompatibilityEx_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func RebootHost(r soap.RoundTripper, req *types.RebootHost_Task) (Task, error) {
	res, err := methods.RebootHost_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ReconfigVM(r soap.RoundTripper, req *types.ReconfigVM_Task) (Task, error) {
	res, err := methods.ReconfigVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ReconfigureCluster(r soap.RoundTripper, req *types.ReconfigureCluster_Task) (Task, error) {
	res, err := methods.ReconfigureCluster_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ReconfigureComputeResource(r soap.RoundTripper, req *types.ReconfigureComputeResource_Task) (Task, error) {
	res, err := methods.ReconfigureComputeResource_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ReconfigureDVPort(r soap.RoundTripper, req *types.ReconfigureDVPort_Task) (Task, error) {
	res, err := methods.ReconfigureDVPort_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ReconfigureDVPortgroup(r soap.RoundTripper, req *types.ReconfigureDVPortgroup_Task) (Task, error) {
	res, err := methods.ReconfigureDVPortgroup_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ReconfigureDatacenter(r soap.RoundTripper, req *types.ReconfigureDatacenter_Task) (Task, error) {
	res, err := methods.ReconfigureDatacenter_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ReconfigureDvs(r soap.RoundTripper, req *types.ReconfigureDvs_Task) (Task, error) {
	res, err := methods.ReconfigureDvs_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ReconfigureHostForDAS(r soap.RoundTripper, req *types.ReconfigureHostForDAS_Task) (Task, error) {
	res, err := methods.ReconfigureHostForDAS_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ReconnectHost(r soap.RoundTripper, req *types.ReconnectHost_Task) (Task, error) {
	res, err := methods.ReconnectHost_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func RectifyDvsHost(r soap.RoundTripper, req *types.RectifyDvsHost_Task) (Task, error) {
	res, err := methods.RectifyDvsHost_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func RectifyDvsOnHost(r soap.RoundTripper, req *types.RectifyDvsOnHost_Task) (Task, error) {
	res, err := methods.RectifyDvsOnHost_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func RegisterChildVM(r soap.RoundTripper, req *types.RegisterChildVM_Task) (Task, error) {
	res, err := methods.RegisterChildVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func RegisterVM(r soap.RoundTripper, req *types.RegisterVM_Task) (Task, error) {
	res, err := methods.RegisterVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func RelocateVM(r soap.RoundTripper, req *types.RelocateVM_Task) (Task, error) {
	res, err := methods.RelocateVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func RemoveAllSnapshots(r soap.RoundTripper, req *types.RemoveAllSnapshots_Task) (Task, error) {
	res, err := methods.RemoveAllSnapshots_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func RemoveDiskMapping(r soap.RoundTripper, req *types.RemoveDiskMapping_Task) (Task, error) {
	res, err := methods.RemoveDiskMapping_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func RemoveDisk(r soap.RoundTripper, req *types.RemoveDisk_Task) (Task, error) {
	res, err := methods.RemoveDisk_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func RemoveSnapshot(r soap.RoundTripper, req *types.RemoveSnapshot_Task) (Task, error) {
	res, err := methods.RemoveSnapshot_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func Rename(r soap.RoundTripper, req *types.Rename_Task) (Task, error) {
	res, err := methods.Rename_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ResetVM(r soap.RoundTripper, req *types.ResetVM_Task) (Task, error) {
	res, err := methods.ResetVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ResignatureUnresolvedVmfsVolume(r soap.RoundTripper, req *types.ResignatureUnresolvedVmfsVolume_Task) (Task, error) {
	res, err := methods.ResignatureUnresolvedVmfsVolume_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ResolveMultipleUnresolvedVmfsVolumesEx(r soap.RoundTripper, req *types.ResolveMultipleUnresolvedVmfsVolumesEx_Task) (Task, error) {
	res, err := methods.ResolveMultipleUnresolvedVmfsVolumesEx_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func RevertToCurrentSnapshot(r soap.RoundTripper, req *types.RevertToCurrentSnapshot_Task) (Task, error) {
	res, err := methods.RevertToCurrentSnapshot_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func RevertToSnapshot(r soap.RoundTripper, req *types.RevertToSnapshot_Task) (Task, error) {
	res, err := methods.RevertToSnapshot_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ScanHostPatchV2(r soap.RoundTripper, req *types.ScanHostPatchV2_Task) (Task, error) {
	res, err := methods.ScanHostPatchV2_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ScanHostPatch(r soap.RoundTripper, req *types.ScanHostPatch_Task) (Task, error) {
	res, err := methods.ScanHostPatch_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func SearchDatastoreSubFolders(r soap.RoundTripper, req *types.SearchDatastoreSubFolders_Task) (Task, error) {
	res, err := methods.SearchDatastoreSubFolders_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func SearchDatastore(r soap.RoundTripper, req *types.SearchDatastore_Task) (Task, error) {
	res, err := methods.SearchDatastore_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ShrinkVirtualDisk(r soap.RoundTripper, req *types.ShrinkVirtualDisk_Task) (Task, error) {
	res, err := methods.ShrinkVirtualDisk_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ShutdownHost(r soap.RoundTripper, req *types.ShutdownHost_Task) (Task, error) {
	res, err := methods.ShutdownHost_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func StageHostPatch(r soap.RoundTripper, req *types.StageHostPatch_Task) (Task, error) {
	res, err := methods.StageHostPatch_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func StartRecording(r soap.RoundTripper, req *types.StartRecording_Task) (Task, error) {
	res, err := methods.StartRecording_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func StartReplaying(r soap.RoundTripper, req *types.StartReplaying_Task) (Task, error) {
	res, err := methods.StartReplaying_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func StopRecording(r soap.RoundTripper, req *types.StopRecording_Task) (Task, error) {
	res, err := methods.StopRecording_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func StopReplaying(r soap.RoundTripper, req *types.StopReplaying_Task) (Task, error) {
	res, err := methods.StopReplaying_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func SuspendVApp(r soap.RoundTripper, req *types.SuspendVApp_Task) (Task, error) {
	res, err := methods.SuspendVApp_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func SuspendVM(r soap.RoundTripper, req *types.SuspendVM_Task) (Task, error) {
	res, err := methods.SuspendVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func TerminateFaultTolerantVM(r soap.RoundTripper, req *types.TerminateFaultTolerantVM_Task) (Task, error) {
	res, err := methods.TerminateFaultTolerantVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func TurnOffFaultToleranceForVM(r soap.RoundTripper, req *types.TurnOffFaultToleranceForVM_Task) (Task, error) {
	res, err := methods.TurnOffFaultToleranceForVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func UninstallHostPatch(r soap.RoundTripper, req *types.UninstallHostPatch_Task) (Task, error) {
	res, err := methods.UninstallHostPatch_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func UnregisterAndDestroy(r soap.RoundTripper, req *types.UnregisterAndDestroy_Task) (Task, error) {
	res, err := methods.UnregisterAndDestroy_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func UpdateAnswerFile(r soap.RoundTripper, req *types.UpdateAnswerFile_Task) (Task, error) {
	res, err := methods.UpdateAnswerFile_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func UpdateDVSHealthCheckConfig(r soap.RoundTripper, req *types.UpdateDVSHealthCheckConfig_Task) (Task, error) {
	res, err := methods.UpdateDVSHealthCheckConfig_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func UpdateDVSLacpGroupConfig(r soap.RoundTripper, req *types.UpdateDVSLacpGroupConfig_Task) (Task, error) {
	res, err := methods.UpdateDVSLacpGroupConfig_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func UpdateVirtualMachineFiles(r soap.RoundTripper, req *types.UpdateVirtualMachineFiles_Task) (Task, error) {
	res, err := methods.UpdateVirtualMachineFiles_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func UpdateVsan(r soap.RoundTripper, req *types.UpdateVsan_Task) (Task, error) {
	res, err := methods.UpdateVsan_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func UpgradeTools(r soap.RoundTripper, req *types.UpgradeTools_Task) (Task, error) {
	res, err := methods.UpgradeTools_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func UpgradeVM(r soap.RoundTripper, req *types.UpgradeVM_Task) (Task, error) {
	res, err := methods.UpgradeVM_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}

func ZeroFillVirtualDisk(r soap.RoundTripper, req *types.ZeroFillVirtualDisk_Task) (Task, error) {
	res, err := methods.ZeroFillVirtualDisk_Task(r, req)
	if err != nil {
		return nil, err
	}

	t := newTask(r, res.Returnval)
	return t, nil
}
