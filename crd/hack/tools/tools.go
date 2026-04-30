//go:build govmomi_tools
// +build govmomi_tools

// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

// Package tools manages the version of tooling used to build this project.
package tools

import (
	_ "k8s.io/code-generator"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
