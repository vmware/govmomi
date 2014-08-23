/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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
	"errors"
	"flag"
	"fmt"

	"github.com/vmware/govmomi"
)

const (
	SearchVirtualMachines = iota + 1
	SearchHosts
)

type SearchFlag struct {
	*ClientFlag
	*DatacenterFlag
	*ListFlag

	t int

	byDatastorePath string
	byDNSName       string
	byInventoryPath string
	byIP            string
	byUUID          string

	isset bool
}

func NewSearchFlag(t int) *SearchFlag {
	s := SearchFlag{
		t: t,
	}

	return &s
}

func (flag *SearchFlag) Register(f *flag.FlagSet) {
	if flag.t == 0 {
		panic("search type not set")
	}

	switch flag.t {
	case SearchVirtualMachines:
		f.StringVar(&flag.byDatastorePath, "path", "", "Find VM by path to .vmx file")
	}

	switch flag.t {
	case SearchVirtualMachines, SearchHosts:
		f.StringVar(&flag.byDNSName, "dns", "", "Find entity by FQDN")
		f.StringVar(&flag.byIP, "ip", "", "Find entity by IP address")
		f.StringVar(&flag.byUUID, "uuid", "", "Find entity by instance UUID")
	}

	f.StringVar(&flag.byInventoryPath, "ipath", "", "Find entity by inventory path")
}

func (flag *SearchFlag) Process() error {
	flags := []string{
		flag.byDatastorePath,
		flag.byDNSName,
		flag.byInventoryPath,
		flag.byIP,
		flag.byUUID,
	}

	flag.isset = false
	for _, f := range flags {
		if f != "" {
			if flag.isset {
				return errors.New("cannot use more than one search flag")
			}
			flag.isset = true
		}
	}

	return nil
}

func (flag *SearchFlag) Isset() bool {
	return flag.isset
}

func (flag *SearchFlag) searchByDatastorePath(c *govmomi.Client, dc *govmomi.Datacenter) (govmomi.Reference, error) {
	switch flag.t {
	case SearchVirtualMachines:
		return c.SearchIndex().FindByDatastorePath(dc, flag.byDatastorePath)
	default:
		panic("unsupported type")
	}
}

func (flag *SearchFlag) searchByDNSName(c *govmomi.Client, dc *govmomi.Datacenter) (govmomi.Reference, error) {
	switch flag.t {
	case SearchVirtualMachines:
		return c.SearchIndex().FindByDnsName(dc, flag.byDNSName, true)
	case SearchHosts:
		return c.SearchIndex().FindByDnsName(dc, flag.byDNSName, false)
	default:
		panic("unsupported type")
	}
}

func (flag *SearchFlag) searchByInventoryPath(c *govmomi.Client, dc *govmomi.Datacenter) (govmomi.Reference, error) {
	// TODO(PN): The datacenter flag should not be set because it is ignored.
	return c.SearchIndex().FindByInventoryPath(flag.byInventoryPath)
}

func (flag *SearchFlag) searchByIP(c *govmomi.Client, dc *govmomi.Datacenter) (govmomi.Reference, error) {
	switch flag.t {
	case SearchVirtualMachines:
		return c.SearchIndex().FindByIp(dc, flag.byIP, true)
	case SearchHosts:
		return c.SearchIndex().FindByIp(dc, flag.byIP, false)
	default:
		panic("unsupported type")
	}
}

func (flag *SearchFlag) searchByUUID(c *govmomi.Client, dc *govmomi.Datacenter) (govmomi.Reference, error) {
	switch flag.t {
	case SearchVirtualMachines:
		return c.SearchIndex().FindByUuid(dc, flag.byUUID, true)
	case SearchHosts:
		return c.SearchIndex().FindByUuid(dc, flag.byUUID, false)
	default:
		panic("unsupported type")
	}
}

func (flag *SearchFlag) search() (govmomi.Reference, error) {
	var ref govmomi.Reference
	var err error

	c, err := flag.Client()
	if err != nil {
		return nil, err
	}

	dc, err := flag.Datacenter()
	if err != nil {
		return nil, err
	}

	switch {
	case flag.byDatastorePath != "":
		ref, err = flag.searchByDatastorePath(c, dc)
	case flag.byDNSName != "":
		ref, err = flag.searchByDNSName(c, dc)
	case flag.byInventoryPath != "":
		ref, err = flag.searchByInventoryPath(c, dc)
	case flag.byIP != "":
		ref, err = flag.searchByIP(c, dc)
	case flag.byUUID != "":
		ref, err = flag.searchByUUID(c, dc)
	default:
		err = errors.New("no search flag specified")
	}

	if err != nil {
		return nil, err
	}

	if ref == nil {
		return nil, fmt.Errorf("not found")
	}

	return ref, nil
}

func (flag *SearchFlag) relativeTo() (*govmomi.DatacenterFolders, error) {
	c, err := flag.Client()
	if err != nil {
		return nil, err
	}

	dc, err := flag.Datacenter()
	if err != nil {
		return nil, err
	}

	f, err := dc.Folders(c)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (flag *SearchFlag) relativeToVmFolder() (govmomi.Reference, error) {
	f, err := flag.relativeTo()
	if err != nil {
		return nil, err
	}

	return govmomi.Folder{f.VmFolder.Reference()}, nil
}

func (flag *SearchFlag) VirtualMachine() (*govmomi.VirtualMachine, error) {
	ref, err := flag.search()
	if err != nil {
		return nil, err
	}

	vm, ok := ref.(*govmomi.VirtualMachine)
	if !ok {
		return nil, fmt.Errorf("expected VirtualMachine entity, got %s", ref.Reference().Type)
	}

	return vm, nil
}

func (flag *SearchFlag) VirtualMachines(args []string) ([]*govmomi.VirtualMachine, error) {
	var out []*govmomi.VirtualMachine

	if flag.Isset() {
		vm, err := flag.VirtualMachine()
		if err != nil {
			return nil, err
		}

		out = append(out, vm)
		return out, nil
	}

	// List virtual machines
	if len(args) == 0 {
		return nil, errors.New("no argument")
	}

	es, err := flag.ListSlice(args, false, flag.relativeToVmFolder)
	if err != nil {
		return nil, err
	}

	// Filter non-VMs
	for _, e := range es {
		ref := e.Object.Reference()
		if ref.Type == "VirtualMachine" {
			out = append(out, &govmomi.VirtualMachine{ref})
		}
	}

	return out, nil
}

func (flag *SearchFlag) HostSystem() (*govmomi.HostSystem, error) {
	ref, err := flag.search()
	if err != nil {
		return nil, err
	}

	host, ok := ref.(*govmomi.HostSystem)
	if !ok {
		return nil, fmt.Errorf("expected HostSystem entity, got %s", ref.Reference().Type)
	}

	return host, nil
}
