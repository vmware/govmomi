/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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
This example program shows how the `finder` and `property` packages can
be used to navigate a vSphere inventory structure using govmomi.
*/

package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

// GetEnvString returns string from environment variable.
func GetEnvString(v string, def string) string {
	r := os.Getenv(v)
	if r == "" {
		return def
	}

	return r
}

// GetEnvBool returns boolean from environment variable.
func GetEnvBool(v string, def bool) bool {
	r := os.Getenv(v)
	if r == "" {
		return def
	}

	switch strings.ToLower(r[0:1]) {
	case "t", "y", "1":
		return true
	}

	return false
}

// Humanize converts a number in bytes to a more readable format.
func Humanize(v int64) string {
	const KB = 1024
	const MB = 1024 * KB
	const GB = 1024 * MB
	const TB = 1024 * GB
	const PB = 1024 * TB

	switch {
	case v < KB:
		return fmt.Sprintf("%dB", v)
	case v < MB:
		return fmt.Sprintf("%.1fKB", float32(v)/KB)
	case v < GB:
		return fmt.Sprintf("%.1fMB", float32(v)/MB)
	case v < TB:
		return fmt.Sprintf("%.1fGB", float32(v)/GB)
	case v < PB:
		return fmt.Sprintf("%.1fTB", float32(v)/TB)
	default:
		return "a lot"
	}
}

var urlVar = "GOVMOMI_URL"
var urlDescription = fmt.Sprintf("ESX or vCenter URL [%s]", urlVar)
var urlFlag = flag.String("url", GetEnvString(urlVar, "https://username:password@host/sdk"), urlDescription)

var insecureVar = "GOVMOMI_INSECURE"
var insecureDescription = fmt.Sprintf("Don't verify the server's certificate chain [%s]", insecureVar)
var insecureFlag = flag.Bool("insecure", GetEnvBool(insecureVar, false), insecureDescription)

func exit(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()

	// Parse URL from string
	u, err := url.Parse(*urlFlag)
	if err != nil {
		exit(err)
	}

	// Connect and log in to ESX or vCenter
	c, err := govmomi.NewClient(ctx, u, *insecureFlag)
	if err != nil {
		exit(err)
	}

	f := find.NewFinder(c.Client, true)

	// Find one and only datacenter
	dc, err := f.DefaultDatacenter(ctx)
	if err != nil {
		exit(err)
	}

	// Make future calls local to this datacenter
	f.SetDatacenter(dc)

	// Find datastores in datacenter
	dss, err := f.DatastoreList(ctx, "*")
	if err != nil {
		exit(err)
	}

	pc := property.DefaultCollector(c.Client)

	// Convert datastores into list of references
	var refs []types.ManagedObjectReference
	for _, ds := range dss {
		refs = append(refs, ds.Reference())
	}

	// Retrieve summary property for all datastores
	var dst []mo.Datastore
	err = pc.Retrieve(ctx, refs, []string{"summary"}, &dst)
	if err != nil {
		exit(err)
	}

	// Print summary per datastore
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "Name:\tType:\tCapacity:\tFree:\n")
	for _, ds := range dst {
		fmt.Fprintf(tw, "%s\t", ds.Summary.Name)
		fmt.Fprintf(tw, "%s\t", ds.Summary.Type)
		fmt.Fprintf(tw, "%s\t", Humanize(ds.Summary.Capacity))
		fmt.Fprintf(tw, "%s\t", Humanize(ds.Summary.FreeSpace))
		fmt.Fprintf(tw, "\n")
	}
	tw.Flush()
}
