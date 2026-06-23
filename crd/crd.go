// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// Package crd is used to nil-import the API packages so that generators may
// run against them successfully.
package crd

import (
	_ "github.com/vmware/govmomi/crd/pkg/vim/api/v1alpha1"
)
