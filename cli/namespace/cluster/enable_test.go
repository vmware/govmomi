// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplitCommaSeparated_Empty(t *testing.T) {
	list := splitCommaSeparatedList("")
	if list != nil || len(list) != 0 {
		t.Fatalf("Nonempty list produced from empty string: %#v", list)
	}
	list = splitCommaSeparatedList(",,,")
	if list != nil || len(list) != 0 {
		t.Fatalf("Nonempty list produced from empty list items: %#v", list)
	}
}

func TestSplitCommaSeparated_Values(t *testing.T) {
	list := splitCommaSeparatedList("0.0.0.0/0")
	if list == nil || len(list) != 1 || list[0] != "0.0.0.0/0" {
		t.Fatalf("Incorrect list produced for single value: %#v", list)
	}
	list = splitCommaSeparatedList(",0.0.0.0/0,1.1.1.1/1,2.2.2.2/2,")
	if list == nil || len(list) != 3 || list[0] != "0.0.0.0/0" || list[1] != "1.1.1.1/1" || list[2] != "2.2.2.2/2" {
		t.Fatalf("Incorrect list produced for multiple values: %#v", list)
	}
}

func TestSplitCidrList_Valid(t *testing.T) {
	list := []string{"0.0.0.0/0", "1.1.1.1/1", "2.2.2.2/2"}
	cidrs, err := splitCidrList(list)
	if err != nil || len(cidrs) != 3 {
		t.Fatalf("Wrong size list produced in cidr splitting: %#v", list)
	}
	for i, cidr := range cidrs {
		if cidr.Prefix != i || cidr.Address != fmt.Sprintf("%[1]d.%[1]d.%[1]d.%[1]d", i) {
			t.Fatalf("Invalid value set in cidr %d: %#v", i, cidr)
		}
	}
}

func TestSplitCidrList_Invalid(t *testing.T) {
	list := []string{"abc", "1.1.1.1/1", "2.2.2.2/2"}
	_, err := splitCidrList(list)
	if err == nil {
		t.Error("Error not produced trying to split an invalid cidr")
	} else if !strings.Contains(err.Error(), "invalid cidr") {
		t.Errorf("Unexpected error produced trying to split an invalid cidr: %s", err)
	}

	list = []string{"/24", "1.1.1.1/1", "2.2.2.2/2"}
	_, err = splitCidrList(list)
	if err == nil {
		t.Error("Error not produced trying to split an invalid cidr")
	} else if !strings.Contains(err.Error(), "parsing cidr") {
		t.Errorf("Unexpected error produced trying to split an invalid cidr: %s", err)
	}

	list = []string{"abc/abc", "1.1.1.1/1", "2.2.2.2/2"}
	_, err = splitCidrList(list)
	if err == nil {
		t.Error("Error not produced trying to split an invalid cidr")
	} else if !strings.Contains(err.Error(), "parsing cidr") {
		t.Errorf("Unexpected error produced trying to split an invalid cidr: %s", err)
	}

}
