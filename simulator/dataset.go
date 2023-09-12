/*
Copyright (c) 2023-2023 VMware, Inc. All Rights Reserved.

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
	"github.com/vmware/govmomi/vapi/vm/dataset"
)

type DataSet struct {
	*dataset.Info
	ID      string
	Entries map[string]string
}

func copyDataSetsForVmClone(src map[string]*DataSet) map[string]*DataSet {
	copy := make(map[string]*DataSet, len(src))
	for k, v := range src {
		if v.OmitFromSnapshotAndClone {
			continue
		}
		copy[k] = copyDataSet(v)
	}
	return copy
}

func copyDataSet(src *DataSet) *DataSet {
	if src == nil {
		return nil
	}
	copy := &DataSet{
		Info: &dataset.Info{
			Name:                     src.Name,
			Description:              src.Description,
			Host:                     src.Host,
			Guest:                    src.Guest,
			Used:                     src.Used,
			OmitFromSnapshotAndClone: src.OmitFromSnapshotAndClone,
		},
		ID:      src.ID,
		Entries: copyEntries(src.Entries),
	}
	return copy
}

func copyEntries(src map[string]string) map[string]string {
	copy := make(map[string]string, len(src))
	for k, v := range src {
		copy[k] = v
	}
	return copy
}
