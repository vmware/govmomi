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
	"strings"

	"github.com/vmware/govmomi"
)

const (
	SearchVirtualMachines = iota + 1
	SearchHosts
)

type SearchFlag struct {
	*ClientFlag
	*DatacenterFlag

	t      int
	entity string

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

	switch t {
	case SearchVirtualMachines:
		s.entity = "VM"
	case SearchHosts:
		s.entity = "host"
	default:
		panic("invalid search type")
	}

	return &s
}

func (flag *SearchFlag) Register(fs *flag.FlagSet) {
	register := func(v *string, f string, d string) {
		f = fmt.Sprintf("%s.%s", strings.ToLower(flag.entity), f)
		d = fmt.Sprintf(d, flag.entity)
		fs.StringVar(v, f, "", d)
	}

	switch flag.t {
	case SearchVirtualMachines:
		register(&flag.byDatastorePath, "path", "Find %s by path to .vmx file")
	}

	switch flag.t {
	case SearchVirtualMachines, SearchHosts:
		register(&flag.byDNSName, "dns", "Find %s by FQDN")
		register(&flag.byIP, "ip", "Find %s by IP address")
		register(&flag.byUUID, "uuid", "Find %s by instance UUID")
	}

	register(&flag.byInventoryPath, "ipath", "Find %s by inventory path")
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

func (flag *SearchFlag) IsSet() bool {
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
		return nil, fmt.Errorf("no such %s", flag.entity)
	}

	return ref, nil
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

	if flag.IsSet() {
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

	finder, err := flag.Finder()
	if err != nil {
		return nil, err
	}

	return finder.VirtualMachineList(args...)
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

func (flag *SearchFlag) HostSystems(args []string) ([]*govmomi.HostSystem, error) {
	var out []*govmomi.HostSystem

	if flag.IsSet() {
		host, err := flag.HostSystem()
		if err != nil {
			return nil, err
		}

		out = append(out, host)
		return out, nil
	}

	// List host system
	if len(args) == 0 {
		return nil, errors.New("no argument")
	}

	finder, err := flag.Finder()
	if err != nil {
		return nil, err
	}

	return finder.HostSystemList(args...)
}
