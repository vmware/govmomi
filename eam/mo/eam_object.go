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

package mo

import (
	"github.com/vmware/govmomi/eam/types"
	vim "github.com/vmware/govmomi/vim25/types"
)

// EamObject contains the fields common to all EAM objects.
type EamObject struct {
	Self  vim.ManagedObjectReference `json:"self"`
	Issue []types.BaseIssue          `json:"issue,omitempty"`
}

func (m EamObject) String() string {
	return m.Self.String()
}

func (m EamObject) Reference() vim.ManagedObjectReference {
	return m.Self
}
