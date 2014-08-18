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

	t int

	byDatastorePath string
	byDNSName       string
	byInventoryPath string
	byIP            string
	byUUID          string
}

func NewSearchFlag(t int) *SearchFlag {
	s := SearchFlag{
		t: t,
	}

	return &s
}

func (s *SearchFlag) Register(f *flag.FlagSet) {
	switch s.t {
	case SearchVirtualMachines:
		f.StringVar(&s.byDatastorePath, "path", "", "Find VM by path to .vmx file")
	}

	switch s.t {
	case SearchVirtualMachines, SearchHosts:
		f.StringVar(&s.byDNSName, "dns", "", "Find entity by FQDN")
		f.StringVar(&s.byIP, "ip", "", "Find entity by IP address")
		f.StringVar(&s.byUUID, "uuid", "", "Find entity by instance UUID")
	}

	f.StringVar(&s.byInventoryPath, "ipath", "", "Find entity by inventory path")
}

func (s *SearchFlag) Process() error {
	flags := []string{
		s.byDatastorePath,
		s.byDNSName,
		s.byInventoryPath,
		s.byIP,
		s.byUUID,
	}

	isset := false
	for _, f := range flags {
		if f != "" {
			if isset {
				return errors.New("cannot use more than one search flag")
			}
			isset = true
		}
	}

	return nil
}

func (s *SearchFlag) searchByDatastorePath(c *govmomi.Client, dc *govmomi.Datacenter) (govmomi.Reference, error) {
	switch s.t {
	case SearchVirtualMachines:
		return c.SearchIndex().FindByDatastorePath(dc, s.byDatastorePath)
	default:
		panic("unsupported type")
	}
}

func (s *SearchFlag) searchByDNSName(c *govmomi.Client, dc *govmomi.Datacenter) (govmomi.Reference, error) {
	switch s.t {
	case SearchVirtualMachines:
		return c.SearchIndex().FindByDnsName(dc, s.byDNSName, true)
	case SearchHosts:
		return c.SearchIndex().FindByDnsName(dc, s.byDNSName, false)
	default:
		panic("unsupported type")
	}
}

func (s *SearchFlag) searchByInventoryPath(c *govmomi.Client, dc *govmomi.Datacenter) (govmomi.Reference, error) {
	// TODO(PN): The datacenter flag should not be set because it is ignored.
	return c.SearchIndex().FindByInventoryPath(s.byInventoryPath)
}

func (s *SearchFlag) searchByIP(c *govmomi.Client, dc *govmomi.Datacenter) (govmomi.Reference, error) {
	switch s.t {
	case SearchVirtualMachines:
		return c.SearchIndex().FindByIp(dc, s.byIP, true)
	case SearchHosts:
		return c.SearchIndex().FindByIp(dc, s.byIP, false)
	default:
		panic("unsupported type")
	}
}

func (s *SearchFlag) searchByUUID(c *govmomi.Client, dc *govmomi.Datacenter) (govmomi.Reference, error) {
	switch s.t {
	case SearchVirtualMachines:
		return c.SearchIndex().FindByUuid(dc, s.byUUID, true)
	case SearchHosts:
		return c.SearchIndex().FindByUuid(dc, s.byUUID, false)
	default:
		panic("unsupported type")
	}
}

func (s *SearchFlag) search() (govmomi.Reference, error) {
	var ref govmomi.Reference
	var err error

	c, err := s.Client()
	if err != nil {
		return nil, err
	}

	dc, err := s.Datacenter()
	if err != nil {
		return nil, err
	}

	switch {
	case s.byDatastorePath != "":
		ref, err = s.searchByDatastorePath(c, dc)
	case s.byDNSName != "":
		ref, err = s.searchByDNSName(c, dc)
	case s.byInventoryPath != "":
		ref, err = s.searchByInventoryPath(c, dc)
	case s.byIP != "":
		ref, err = s.searchByIP(c, dc)
	case s.byUUID != "":
		ref, err = s.searchByUUID(c, dc)
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

func (s *SearchFlag) VirtualMachine() (*govmomi.VirtualMachine, error) {
	ref, err := s.search()
	if err != nil {
		return nil, err
	}

	vm, ok := ref.(*govmomi.VirtualMachine)
	if !ok {
		return nil, fmt.Errorf("expected VirtualMachine entity, got %s", ref.Reference().Type)
	}

	return vm, nil
}

func (s *SearchFlag) HostSystem() (*govmomi.HostSystem, error) {
	ref, err := s.search()
	if err != nil {
		return nil, err
	}

	host, ok := ref.(*govmomi.HostSystem)
	if !ok {
		return nil, fmt.Errorf("expected HostSystem entity, got %s", ref.Reference().Type)
	}

	return host, nil
}
