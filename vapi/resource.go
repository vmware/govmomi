// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vapi

import (
	"strings"

	"github.com/vmware/govmomi/vim25/types"
)

const (
	// Path is the new-style endpoint for API resources. It supersedes /rest.
	Path = "/api"
)

func Task(id string) types.ManagedObjectReference {
	return types.ManagedObjectReference{
		Type:  "Task",
		Value: strings.SplitN(id, ":", 2)[0],
	}
}
