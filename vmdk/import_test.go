/*
Copyright (c) 2017-2024 VMware, Inc. All Rights Reserved.

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

package vmdk_test

import (
	"os"
	"testing"

	"github.com/vmware/govmomi/vmdk"
)

func TestDiskInfo(t *testing.T) {
	di, err := vmdk.Stat("../govc/test/images/ttylinux-pc_i486-16.1-disk1.vmdk")
	if err != nil {
		if os.IsNotExist(err) {
			t.SkipNow()
		}
		t.Fatal(err)
	}

	_, err = di.OVF()
	if err != nil {
		t.Fatal(err)
	}
}

func TestDiskInvalid(t *testing.T) {
	_, err := vmdk.Stat("import_test.go")
	if err != vmdk.ErrInvalidFormat {
		t.Errorf("expected ErrInvalidFormat: %s", err)
	}
}
