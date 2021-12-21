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

package main

import (
	"log"
	"os"

	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
	"github.com/vmware/govmomi/vim25/xml/util"
)

func main() {
	var objToXml interface{}

	if len(os.Args) > 1 && os.Args[1] == "-" {
		chanObj, chanErr := util.VimObjectsFromReader(os.Stdin)
		select {
		case obj := <-chanObj:
			objToXml = obj
		case err := <-chanErr:
			log.Fatal(err)
		}
	} else {
		objToXml = &types.VirtualMachineConfigSpec{
			Name:     "go-vm",
			NumCPUs:  2,
			MemoryMB: 2048,
		}
	}

	// Create an XML encoder so we can use pretty-printing for the XML.
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")

	if err := util.VimObjectsToEncoder(enc, objToXml); err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write([]byte{10})
}
