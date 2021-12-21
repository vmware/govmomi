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
	"os"

	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
	"github.com/vmware/govmomi/vim25/xml/util"
)

func ExampleVimObjectsFromString() {
	const vimObjects = `
<obj xmlns:vim25="urn:vim25"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="vim25:VirtualMachineConfigSpec">
    <name>go-vm</name>
    <numCPUs>2</numCPUs>
    <memoryMB>2048</memoryMB>
</obj>
<obj xmlns:vim25="urn:vim25"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="vim25:VirtualMachineConfigInfo">
    <name>go-vm</name>
    <version>vmx-07</version>
    <guestId>dosGuest</guestId>
    <hardware>
        <numCPU>2</numCPU>
        <memoryMB>2048</memoryMB>
    </hardware>
</obj>
`

	// Decode all of the objects from the XML string.
	chanObj, chanErr := util.VimObjectsFromString(vimObjects)

	// For each object we see, print its type and information about it, along
	// with any errors that occur along the way.
	for {
		select {
		case obj, ok := <-chanObj:
			if !ok {
				return
			}
			fmt.Printf("%T\n", obj)
			switch t := obj.(type) {
			case *types.VirtualMachineConfigInfo:
				fmt.Printf("  name:    %s\n", t.Name)
				fmt.Printf("  numCPU:  %d\n", t.Hardware.NumCPU)
				fmt.Printf("  memMiB:  %d\n", t.Hardware.MemoryMB)
				fmt.Printf("  version: %s\n", t.Version)
				fmt.Printf("  guestID: %s\n", t.GuestId)
			case *types.VirtualMachineConfigSpec:
				fmt.Printf("  name:   %s\n", t.Name)
				fmt.Printf("  numCPU: %d\n", t.NumCPUs)
				fmt.Printf("  memMiB: %d\n", t.MemoryMB)
			}
		case err, ok := <-chanErr:
			if !ok {
				return
			}
			fmt.Println(err)
		}
	}

	// Output:
	// *types.VirtualMachineConfigSpec
	//   name:   go-vm
	//   numCPU: 2
	//   memMiB: 2048
	// *types.VirtualMachineConfigInfo
	//   name:    go-vm
	//   numCPU:  2
	//   memMiB:  2048
	//   version: vmx-07
	//   guestID: dosGuest
}

func ExampleVimObjectsToEncoder() {
	vimObjs := []interface{}{
		&types.VirtualMachineConfigInfo{
			Name:    "MyFirstVM",
			GuestId: "dosGuest",
			Version: "vmx-07",
			Hardware: types.VirtualHardware{
				NumCPU:   2,
				MemoryMB: 2048,
			},
		},
		&types.VirtualMachineConfigSpec{
			Name:     "MyFirstVM",
			NumCPUs:  2,
			MemoryMB: 2048,
		},
	}

	// Create an XML encoder so we can use pretty-printing for the XML.
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")

	// Encode the objects to stdout.
	util.VimObjectsToEncoder(enc, vimObjs...)

	// Print a trailing \n character.
	os.Stdout.Write([]byte{10})

	// Output:
	// <obj xmlns:vim25="urn:vim25" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="vim25:VirtualMachineConfigInfo">
	//   <changeVersion></changeVersion>
	//   <modified>0001-01-01T00:00:00Z</modified>
	//   <name>MyFirstVM</name>
	//   <guestFullName></guestFullName>
	//   <version>vmx-07</version>
	//   <uuid></uuid>
	//   <template>false</template>
	//   <guestId>dosGuest</guestId>
	//   <alternateGuestName></alternateGuestName>
	//   <files></files>
	//   <flags></flags>
	//   <defaultPowerOps></defaultPowerOps>
	//   <hardware>
	//     <numCPU>2</numCPU>
	//     <memoryMB>2048</memoryMB>
	//   </hardware>
	// </obj>
	// <obj xmlns:vim25="urn:vim25" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:type="vim25:VirtualMachineConfigSpec">
	//   <name>MyFirstVM</name>
	//   <numCPUs>2</numCPUs>
	//   <memoryMB>2048</memoryMB>
	// </obj>
}
