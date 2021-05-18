/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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
	"sync"

	"github.com/vmware/govmomi/eam/types"
)

var (
	activeIssueKeys   = map[int32]struct{}{}
	activeIssueKeysMu sync.RWMutex
)

func nextAvailableIssueKey() int32 {
	activeIssueKeysMu.Lock()
	defer activeIssueKeysMu.Unlock()
	for i := int32(1); ; i++ {
		if _, isActiveKey := activeIssueKeys[i]; !isActiveKey {
			activeIssueKeys[i] = struct{}{}
			return i
		}
	}
}

func freeIssueKey(i int32) {
	activeIssueKeysMu.Lock()
	defer activeIssueKeysMu.Unlock()
	delete(activeIssueKeys, i)
}

func issueType(issue types.BaseIssue) types.BaseIssue {
	switch typedIssue := issue.(type) {
	case types.BaseVmNotDeployed:
		return typedIssue.GetVmNotDeployed()
	case types.BaseVmDeployed:
		return typedIssue.GetVmDeployed()
	case types.BaseVmPoweredOff:
		return typedIssue.GetVmPoweredOff()
	case types.BaseVmIssue:
		return typedIssue.GetVmIssue()
	case types.BaseVibNotInstalled:
		return typedIssue.GetVibNotInstalled()
	case types.BaseVibIssue:
		return typedIssue.GetVibIssue()
	case types.BaseAgentIssue:
		return typedIssue.GetAgentIssue()
	case types.BaseAgencyIssue:
		return typedIssue.GetAgencyIssue()
	case types.BaseHostIssue:
		return typedIssue.GetHostIssue()
	default:
		return issue
	}
}
