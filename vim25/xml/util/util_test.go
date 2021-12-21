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

// Package util implements some helper functions for marshaling/unmarshaling
// vim objects to/from XML.
package util_test

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
	"github.com/vmware/govmomi/vim25/xml/util"
)

type testCaseKind uint8

const (
	testCaseFrom testCaseKind = 1 << iota
	testCaseTo
)

var (
	testCases = []struct {
		name string
		data string
		objs []interface{}
		errs []error
		kind testCaseKind
	}{
		{
			name: "Single valid vim.vm.ConfigSpec",
			kind: testCaseFrom | testCaseTo,
			data: vimVmConfigSpecSmall,
			objs: []interface{}{
				&types.VirtualMachineConfigSpec{
					NumCPUs:  2,
					MemoryMB: 2048,
				},
			},
		},
		{
			name: "Multiple valid vim.vm.ConfigSpec",
			kind: testCaseFrom | testCaseTo,
			data: vimVmConfigSpecSmallAndMed,
			objs: []interface{}{
				&types.VirtualMachineConfigSpec{
					NumCPUs:  2,
					MemoryMB: 2048,
				},
				&types.VirtualMachineConfigSpec{
					NumCPUs:  4,
					MemoryMB: 4096,
				},
			},
		},
		{
			name: "Single invalid vim.vm.ConfigSpec",
			kind: testCaseFrom,
			data: vimVmConfigSpecInvalidEndElement,
			errs: []error{
				&xml.SyntaxError{
					Line: 6,
					Msg:  "element <obj> closed by </obj1>",
				},
			},
		},
		{
			name: "Single valid vim.vm.ConfigInfo",
			kind: testCaseFrom | testCaseTo,
			data: vimVmConfigInfoSmall,
			objs: []interface{}{
				&types.VirtualMachineConfigInfo{
					Name:    "MyFirstVM",
					GuestId: "dosGuest",
					Version: "vmx-07",
					Hardware: types.VirtualHardware{
						NumCPU:   2,
						MemoryMB: 2048,
						Device: []types.BaseVirtualDevice{
							&types.VirtualSCSIController{
								VirtualController: types.VirtualController{
									BusNumber: 1,
									VirtualDevice: types.VirtualDevice{
										Key: -1,
									},
								},
								ScsiCtlrUnitNumber: 1,
								SharedBus:          types.VirtualSCSISharingNoSharing,
							},
							&types.VirtualDisk{
								VirtualDevice: types.VirtualDevice{
									Key:           -2,
									ControllerKey: -1,
									Backing:       &types.VirtualDiskFlatVer2BackingInfo{},
								},
								CapacityInKB: 51200,
							},
						},
					},
				},
			},
		},
	}
)

func TestVimObjectsFrom(t *testing.T) {

	// Please note only the FromBytes and FromString variants are directly
	// tested. This is because they both in turn invoke FromReader which in
	// turn invokes FromDecoder. This means both FromBytes and FromString
	// provide coverage for FromReader and FromDecoder.
	methods := []struct {
		name string
	}{
		{name: "Bytes"},
		{name: "String"},
	}

	for i := range methods {
		m := methods[i]

		t.Run(m.name, func(t *testing.T) {

			for j := range testCases {
				c := testCases[j]

				// Only run supported test cases.
				if c.kind&testCaseFrom == 0 {
					continue
				}

				t.Run(c.name, func(t *testing.T) {
					var (
						chanObj <-chan interface{}
						chanErr <-chan error
					)
					switch m.name {
					case "Bytes":
						chanObj, chanErr = util.VimObjectsFromBytes([]byte(c.data))
					case "String":
						chanObj, chanErr = util.VimObjectsFromString(c.data)

					}

					// Get the decoded objects and any errors that occurred.
					objs, errs := getVimObjectsAndErrors(chanObj, chanErr)

					// Validate the expected errors.
					if e, a := len(c.errs), len(errs); e != a {
						t.Errorf("invalid error count: exp=%d, act=%d", e, a)
					}
					for i := range c.errs {
						if e, a := c.errs[i].Error(), errs[i].Error(); e != a {
							t.Errorf("unxpected error: exp=%s, act=%s", e, a)
						}
					}

					// Validate the expected objects.
					if e, a := len(c.objs), len(objs); e != a {
						t.Errorf("invalid object count: exp=%d, act=%d", e, a)
					}
					for i := range c.objs {
						e, a := c.objs[i], objs[i]
						if err := compareVimObjects(e, a); err != nil {
							t.Error(err)
						}
					}
				})
			}
		})
	}
}

func TestVimObjectsTo(t *testing.T) {

	// Please note only the ToBytes and ToString variants are directly
	// tested. This is because they both in turn invoke ToWriter which in
	// turn invokes ToEncoder. This means both ToBytes and ToString
	// provide coverage for ToWriter and ToEncoder.
	methods := []struct {
		name string
	}{
		{name: "String"},
		{name: "Bytes"},
	}

	for i := range methods {
		m := methods[i]

		t.Run(m.name, func(t *testing.T) {

			for j := range testCases {
				c := testCases[j]

				// Only run supported test cases.
				if c.kind&testCaseTo == 0 {
					continue
				}

				t.Run(c.name, func(t *testing.T) {

					var (
						aval string
						aerr error
					)

					switch m.name {
					case "Bytes":
						var b []byte
						b, aerr = util.VimObjectsToBytes(c.objs...)
						if len(b) > 0 {
							aval = string(b)
						}
					case "String":
						aval, aerr = util.VimObjectsToString(c.objs...)
					}

					if len(c.errs) > 0 {
						if aerr == nil {
							t.Errorf("invalid error count: exp=1, act=0")
						} else if e, a := c.errs[0].Error(), aerr.Error(); e != a {
							t.Errorf("unxpected error: exp=%s, act=%s", e, a)
						}
					}

					if len(c.data) > 0 {
						if e, a := trimXMLWhiteSpace(c.data), aval; e != a {
							t.Errorf("invalid xml string: exp=%s, act=%s", e, a)
						}
					} else {
						t.Log(aval)
					}
				})
			}
		})
	}
}

func getVimObjectsAndErrors(
	chanObj <-chan interface{},
	chanErr <-chan error) (objs []interface{}, errs []error) {

	for {
		select {
		case v, ok := <-chanObj:
			if !ok {
				return
			}
			objs = append(objs, v)
		case v, ok := <-chanErr:
			if !ok {
				return
			}
			errs = append(errs, v)
		}
	}
}

func compareVimObjects(e, a interface{}) error {

	// eqFn is used to determine if e and a are equal and is dependent upon the
	// type of vim object.
	var eqFn func() error

	switch tA := a.(type) {
	case *types.VirtualMachineConfigInfo:
		eqFn = func() error {
			return compareVimVmConfigInfos(
				e.(*types.VirtualMachineConfigInfo),
				tA,
			)
		}
	case *types.VirtualMachineConfigSpec:
		eqFn = func() error {
			return compareVimVmConfigSpecs(
				e.(*types.VirtualMachineConfigSpec),
				tA,
			)
		}
	default:
		return fmt.Errorf("unexpected type: %T", a)
	}

	typeE, typeA := reflect.TypeOf(e), reflect.TypeOf(a)
	if typeE != typeA {
		return fmt.Errorf("invalid type: exp=%s, act=%s", typeE, typeA)
	}

	return eqFn()
}

func compareVimVmConfigSpecs(e, a *types.VirtualMachineConfigSpec) error {
	if e, a := e.NumCPUs, a.NumCPUs; e != a {
		return fmt.Errorf("invalid num cpu: exp: %d, act: %d", e, a)
	}
	if e, a := e.MemoryMB, a.MemoryMB; e != a {
		return fmt.Errorf("invalid mem mib: exp: %d, act: %d", e, a)
	}

	return nil
}

func compareVimVmConfigInfos(e, a *types.VirtualMachineConfigInfo) error {
	if e, a := e.Name, a.Name; e != a {
		return fmt.Errorf("invalid name: exp: %s, act: %s", e, a)
	}
	if e, a := e.GuestId, a.GuestId; e != a {
		return fmt.Errorf("invalid guestID: exp: %s, act: %s", e, a)
	}

	he, ha := e.Hardware, a.Hardware
	if e, a := he.NumCPU, ha.NumCPU; e != a {
		return fmt.Errorf("invalid hardware.numCPU: exp: %d, act: %d", e, a)
	}
	if e, a := he.MemoryMB, ha.MemoryMB; e != a {
		return fmt.Errorf("invalid hardware.memoryMB: exp: %d, act: %d", e, a)
	}

	if e, a := len(he.Device), len(ha.Device); e != a {
		return fmt.Errorf("invalid device count: exp=%d, act=%d", e, a)
	}
	for i := range he.Device {
		hed, had := he.Device[i], ha.Device[i]
		hev, hav := hed.GetVirtualDevice(), had.GetVirtualDevice()
		if hev == nil {
			return fmt.Errorf("exp virtual device is nil")
		}
		if hav == nil {
			return fmt.Errorf("act virtual device is nil")
		}
		if e, a := hev.Key, hav.Key; e != a {
			return fmt.Errorf("invalid deviceKey: exp: %d, act: %d", e, a)
		}
	}

	return nil
}

var (
	wsInBetweenLTGTRx   = regexp.MustCompile(`>\s+?<`)
	wsAtBeginningOfLine = regexp.MustCompile(`(?m)^\s{2,}`)
)

func trimXMLWhiteSpace(s string) string {
	s = wsAtBeginningOfLine.ReplaceAllString(s, " ")
	s = strings.ReplaceAll(s, "\n", "")
	s = wsInBetweenLTGTRx.ReplaceAllString(s, "><")
	return s
}

func ptrToInt64(i int64) *int64 {
	return &i
}

const (
	vimVmConfigSpecSmall = `
<obj xmlns:vim25="urn:vim25"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="vim25:VirtualMachineConfigSpec">
    <numCPUs>2</numCPUs>
    <memoryMB>2048</memoryMB>
</obj>
`
	vimVmConfigSpecMed = `
<obj xmlns:vim25="urn:vim25"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="vim25:VirtualMachineConfigSpec">
    <numCPUs>4</numCPUs>
    <memoryMB>4096</memoryMB>
</obj>
`

	vimVmConfigSpecSmallAndMed = vimVmConfigSpecSmall + vimVmConfigSpecMed

	vimVmConfigSpecInvalidEndElement = `
<obj xmlns:vim25="urn:vim25"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="vim25:VirtualMachineConfigSpec">
    <numCPUs>4</numCPUs>
    <memoryMB>4096</memoryMB>
</obj1>
`

	vimVmConfigInfoSmall = `
<obj xmlns:vim25="urn:vim25"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="vim25:VirtualMachineConfigInfo">
    <changeVersion></changeVersion>
    <modified>0001-01-01T00:00:00Z</modified>
    <name>MyFirstVM</name>
    <guestFullName></guestFullName>
    <version>vmx-07</version>
    <uuid></uuid>
    <template>false</template>
    <guestId>dosGuest</guestId>
    <alternateGuestName></alternateGuestName>
    <files></files>
    <flags></flags>
    <defaultPowerOps></defaultPowerOps>
    <hardware>
        <numCPU>2</numCPU>
        <memoryMB>2048</memoryMB>
        <device xmlns:XMLSchema-instance="http://www.w3.org/2001/XMLSchema-instance" XMLSchema-instance:type="VirtualSCSIController">
            <key>-1</key>
            <busNumber>1</busNumber>
            <sharedBus>noSharing</sharedBus>
            <scsiCtlrUnitNumber>1</scsiCtlrUnitNumber>
        </device>
        <device xmlns:XMLSchema-instance="http://www.w3.org/2001/XMLSchema-instance" XMLSchema-instance:type="VirtualDisk">
            <key>-2</key>
            <backing XMLSchema-instance:type="VirtualDiskFlatVer2BackingInfo">
                <fileName></fileName>
                <diskMode></diskMode>
            </backing>
            <controllerKey>-1</controllerKey>
            <capacityInKB>51200</capacityInKB>
        </device>
    </hardware>
</obj>
`
)
