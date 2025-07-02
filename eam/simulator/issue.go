// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
