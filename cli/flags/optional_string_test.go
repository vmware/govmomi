/*
Copyright (c) 2023-2023 VMware, Inc. All Rights Reserved.

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

package flags

import (
	"flag"
	"testing"
)

func TestOptionalString(t *testing.T) {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	var val *string

	fs.Var(NewOptionalString(&val), "ostring", "optional string")

	s := fs.Lookup("ostring")

	if s.DefValue != "<nil>" {
		t.Fail()
	}

	if s.Value.String() != "<nil>" {
		t.Fail()
	}

	if s.Value.(flag.Getter).Get() != nil {
		t.Fail()
	}

	s.Value.Set("test")

	if s.Value.String() != "test" {
		t.Fail()
	}

	if s.Value.(flag.Getter).Get() != "test" {
		t.Fail()
	}

	if val == nil || *val != "test" {
		t.Fail()
	}
}
