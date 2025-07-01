// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package toolbox

import (
	"net"
	"reflect"
	"testing"
)

func TestDefaultGuestNicProto(t *testing.T) {
	p := DefaultGuestNicInfo()

	info := p.V3

	for _, nic := range info.Nics {
		if len(nic.MacAddress) == 0 {
			continue
		}
		_, err := net.ParseMAC(nic.MacAddress)
		if err != nil {
			t.Errorf("invalid MAC %s: %s", nic.MacAddress, err)
		}
	}

	b, err := EncodeXDR(p)
	if err != nil {
		t.Fatal(err)
	}

	var dp GuestNicInfo
	err = DecodeXDR(b, &dp)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(p, &dp) {
		t.Error("decode mismatch")
	}
}

func TestMaxGuestNic(t *testing.T) {
	p := DefaultGuestNicInfo()

	maxNics = len(p.V3.Nics)

	a, _ := net.Interfaces()
	a = append(a, a...) // double the number of interfaces returned
	netInterfaces = func() ([]net.Interface, error) {
		return a, nil
	}

	p = DefaultGuestNicInfo()

	l := len(p.V3.Nics)
	if l != maxNics {
		t.Errorf("Nics=%d", l)
	}
}
