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
	"testing"
)

func TestHardwareVersion(t *testing.T) {
	testCases := []struct {
		name            string
		in              string
		expectedIsValid bool
		expectedInt     int
		expectedString  string
	}{
		{
			name:            "empty string",
			in:              "",
			expectedIsValid: false,
			expectedInt:     0,
			expectedString:  "",
		},
		{
			name:            "vmx prefix sans number",
			in:              "vmx-",
			expectedIsValid: false,
			expectedInt:     0,
			expectedString:  "",
		},
		{
			name:            "number sans vmx prefix",
			in:              "13",
			expectedIsValid: false,
			expectedInt:     0,
			expectedString:  "",
		},
		{
			name:            "vmx-13",
			in:              "vmx-13",
			expectedIsValid: true,
			expectedInt:     13,
			expectedString:  "vmx-13",
		},
		{
			name:            "VMX-18",
			in:              "VMX-18",
			expectedIsValid: true,
			expectedInt:     18,
			expectedString:  "vmx-18",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			hv := HardwareVersion(tc.in)
			if a, e := string(hv), tc.in; a != e {
				t.Errorf("unexpected type assert: a=%v, e=%v", a, e)
			}
			if a, e := hv.IsValid(), tc.expectedIsValid; a != e {
				t.Errorf("unexpected IsValid: a=%v, e=%v", a, e)
			}
			if a, e := hv.Int(), tc.expectedInt; a != e {
				t.Errorf("unexpected Int: a=%v, e=%v", a, e)
			}
			if a, e := hv.String(), tc.expectedString; a != e {
				t.Errorf("unexpected String: a=%v, e=%v", a, e)
			}
		})
	}
}
