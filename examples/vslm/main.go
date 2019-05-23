/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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

/*
This example program shows how the `view` and `property` packages can
be used to navigate a vSphere inventory structure using govmomi.
*/

package main

import (
	"context"
	"fmt"
	"log"
	//	"os"
	//	"text/tabwriter"

	"github.com/vmware/govmomi/examples"
	//	"github.com/vmware/govmomi/units"
	//	"github.com/vmware/govmomi/view"
	//	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vslm"
)

func main() {
	ctx := context.Background()

	// Connect and login to vCenter (vpxd)
	c, err := examples.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	vslmClient, err := vslm.NewClient(ctx, c.Client)
	
	if err != nil {
		log.Fatal(err)
	}
	objectManager := vslm.NewVSLMObjectManager(vslmClient)
	res, err := objectManager.ListVStorageObjectForSpec(ctx, nil, 1000)
	
	fmt.Println("Got ", len(res.Id), " IDs returned")
	defer c.Logout(ctx)
}
