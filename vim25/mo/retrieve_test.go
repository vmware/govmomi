/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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
	"os"
	"testing"

	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

func TestNotAuthenticated(t *testing.T) {
	f, err := os.Open("fixtures/not_authenticated.xml")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var b types.RetrievePropertiesResponse

	dec := xml.NewDecoder(f)
	dec.TypeFunc = types.TypeFunc()
	if err := dec.Decode(&b); err != nil {
		panic(err)
	}

	var s SessionManager

	err = LoadRetrievePropertiesResponse(&b, &s)
	if !soap.IsVimFault(err) {
		t.Errorf("Expected IsVimFault")
	}

	fault := soap.ToVimFault(err).(*types.NotAuthenticated)
	if fault.PrivilegeId != "System.View" {
		t.Errorf("Expected first fault to be returned")
	}
}
