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

package object

import (
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type StorageResourceManager struct {
	Common
}

func NewStorageResourceManager(c *vim25.Client) *StorageResourceManager {
	sr := StorageResourceManager{
		Common: NewCommon(c, *c.ServiceContent.StorageResourceManager),
	}

	return &sr
}

func (sr StorageResourceManager) ApplyStorageDrsRecommendation(key []string) (*Task, error) {
	req := types.ApplyStorageDrsRecommendation_Task{
		This: sr.Reference(),
		Key:  key,
	}

	res, err := methods.ApplyStorageDrsRecommendation_Task(context.TODO(), sr.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(sr.c, res.Returnval), nil
}

func (sr StorageResourceManager) ApplyStorageDrsRecommendationToPod(pod *StoragePod, key string) (*Task, error) {
	req := types.ApplyStorageDrsRecommendationToPod_Task{
		This: sr.Reference(),
		Key:  key,
	}

	if pod != nil {
		req.Pod = pod.Reference()
	}

	res, err := methods.ApplyStorageDrsRecommendationToPod_Task(context.TODO(), sr.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(sr.c, res.Returnval), nil
}

func (sr StorageResourceManager) CancelStorageDrsRecommendation(key []string) error {
	req := types.CancelStorageDrsRecommendation{
		This: sr.Reference(),
		Key:  key,
	}

	_, err := methods.CancelStorageDrsRecommendation(context.TODO(), sr.c, &req)

	return err
}

func (sr StorageResourceManager) ConfigureDatastoreIORM(datastore *Datastore, spec types.StorageIORMConfigSpec, key string) (*Task, error) {
	req := types.ConfigureDatastoreIORM_Task{
		This: sr.Reference(),
		Spec: spec,
	}

	if datastore != nil {
		req.Datastore = datastore.Reference()
	}

	res, err := methods.ConfigureDatastoreIORM_Task(context.TODO(), sr.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(sr.c, res.Returnval), nil
}

func (sr StorageResourceManager) ConfigureStorageDrsForPod(pod *StoragePod, spec types.StorageDrsConfigSpec, modify bool) (*Task, error) {
	req := types.ConfigureStorageDrsForPod_Task{
		This:   sr.Reference(),
		Spec:   spec,
		Modify: modify,
	}

	if pod != nil {
		req.Pod = pod.Reference()
	}

	res, err := methods.ConfigureStorageDrsForPod_Task(context.TODO(), sr.c, &req)
	if err != nil {
		return nil, err
	}

	return NewTask(sr.c, res.Returnval), nil
}

func (sr StorageResourceManager) QueryDatastorePerformanceSummary(datastore *Datastore) ([]types.StoragePerformanceSummary, error) {
	req := types.QueryDatastorePerformanceSummary{
		This: sr.Reference(),
	}

	if datastore != nil {
		req.Datastore = datastore.Reference()
	}

	res, err := methods.QueryDatastorePerformanceSummary(context.TODO(), sr.c, &req)
	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

func (sr StorageResourceManager) QueryIORMConfigOption(host *HostSystem) (*types.StorageIORMConfigOption, error) {
	req := types.QueryIORMConfigOption{
		This: sr.Reference(),
	}

	if host != nil {
		req.Host = host.Reference()
	}

	res, err := methods.QueryIORMConfigOption(context.TODO(), sr.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (sr StorageResourceManager) RecommendDatastores(storageSpec types.StoragePlacementSpec) (*types.StoragePlacementResult, error) {
	req := types.RecommendDatastores{
		This:        sr.Reference(),
		StorageSpec: storageSpec,
	}

	res, err := methods.RecommendDatastores(context.TODO(), sr.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

func (sr StorageResourceManager) RefreshStorageDrsRecommendation(pod *StoragePod) error {
	req := types.RefreshStorageDrsRecommendation{
		This: sr.Reference(),
	}

	if pod != nil {
		req.Pod = pod.Reference()
	}

	_, err := methods.RefreshStorageDrsRecommendation(context.TODO(), sr.c, &req)

	return err
}
