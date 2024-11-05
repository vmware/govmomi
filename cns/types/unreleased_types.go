/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package types

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/types"
)

// CnsBlockCreateSpec is the specification for creating block volumes.
type CnsBlockCreateSpec struct {
	CnsBaseCreateSpec

	// Crypto specifies the encryption settings for the volume to be created.
	// Works with block volumes only.
	CryptoSpec *types.CryptoSpec `xml:"cryptoSpec,omitempty"`
}

func init() {
	types.Add("CnsBlockCreateSpec", reflect.TypeOf((*CnsBlockCreateSpec)(nil)).Elem())
}
