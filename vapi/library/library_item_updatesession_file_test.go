/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package library

import (
	"strings"
	"testing"
)

func TestReadManifest(t *testing.T) {
	mf := `SHA1(ttylinux-pc_i486-16.1.ovf)= 344a536fab6782622b6beb923798e84134bd4cbd
SHA1(ttylinux-pc_i486-16.1-disk1.vmdk)= ed64564a37366bfe1c93af80e2ead0cbd398c3d3`

	sums, err := ReadManifest(strings.NewReader(mf))
	if err != nil {
		t.Error(err)
	}

	if len(sums) != 2 {
		t.Error(err)
	}
}
