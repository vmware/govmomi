// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"reflect"
	"testing"
)

func TestTypeFunc(t *testing.T) {
	var ok bool

	fn := TypeFunc()

	_, ok = fn("unknown")
	if ok {
		t.Errorf("Expected ok==false")
	}

	actual, ok := fn("UserProfile")
	if !ok {
		t.Errorf("Expected ok==true")
	}

	expected := reflect.TypeOf(UserProfile{})
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %#v, actual: %#v", expected, actual)
	}
}
