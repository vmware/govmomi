// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
