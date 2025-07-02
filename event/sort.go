// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package event

import (
	"sort"

	"github.com/vmware/govmomi/vim25/types"
)

// Sort events in ascending order base on Key
// From the EventHistoryCollector.latestPage sdk docs:
//
//	The "oldest event" is the one with the smallest key (event ID).
//	The events in the returned page are unordered.
func Sort(events []types.BaseEvent) {
	sort.Sort(baseEvent(events))
}

type baseEvent []types.BaseEvent

func (d baseEvent) Len() int {
	return len(d)
}

func (d baseEvent) Less(i, j int) bool {
	return d[i].GetEvent().Key < d[j].GetEvent().Key
}

func (d baseEvent) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
