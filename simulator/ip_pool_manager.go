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

package simulator

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

var ipPool = MustNewIpPool(&types.IpPool{
	Id:   1,
	Name: "ip-pool",
	AvailableIpv4Addresses: 250,
	AvailableIpv6Addresses: 250,
	AllocatedIpv6Addresses: 0,
	AllocatedIpv4Addresses: 0,
	Ipv4Config: &types.IpPoolIpPoolConfigInfo{
		Netmask:       "10.10.10.255",
		Gateway:       "10.10.10.1",
		SubnetAddress: "10.10.10.0",
		Range:         "10.10.10.2#250",
	},
	Ipv6Config: &types.IpPoolIpPoolConfigInfo{
		Netmask:       "2001:4860:0:2001::ff",
		Gateway:       "2001:4860:0:2001::1",
		SubnetAddress: "2001:4860:0:2001::0",
		Range:         "2001:4860:0:2001::2#250",
	},
})

type IpPoolManager struct {
	mo.IpPoolManager

	pools      map[int32]*IpPool
	nextPoolId int32
}

func NewIpPoolManager(ref types.ManagedObjectReference) *IpPoolManager {
	m := &IpPoolManager{}
	m.Self = ref

	m.pools = map[int32]*IpPool{
		1: ipPool,
	}
	m.nextPoolId = 2

	return m
}

func (m *IpPoolManager) CreateIpPool(req *types.CreateIpPool) soap.HasFault {
	return &methods.CreateIpPoolBody{
		Fault_: Fault("not implemented", &types.RuntimeFault{}),
	}
}

func (m *IpPoolManager) DestroyIpPool(req *types.DestroyIpPool) soap.HasFault {
	return &methods.DestroyIpPoolBody{
		Fault_: Fault("not implemented", &types.RuntimeFault{}),
	}
}

func (m *IpPoolManager) QueryIpPools(req *types.QueryIpPools) soap.HasFault {
	return &methods.QueryIpPoolsBody{
		Fault_: Fault("not implemented", &types.RuntimeFault{}),
	}
}

func (m *IpPoolManager) UpdateIpPool(req *types.UpdateIpPool) soap.HasFault {
	return &methods.UpdateIpPoolBody{
		Fault_: Fault("not implemented", &types.RuntimeFault{}),
	}
}

func (m *IpPoolManager) AllocateIpv4Address(req *types.AllocateIpv4Address) soap.HasFault {
	body := &methods.AllocateIpv4AddressBody{}

	pool, ok := m.pools[req.PoolId]
	if !ok {
		body.Fault_ = Fault("", &types.InvalidArgument{})
		return body
	}

	ip, err := pool.AllocateIPv4(req.AllocationId)
	if err != nil {
		body.Fault_ = Fault(err.Error(), &types.RuntimeFault{})
		return body
	}

	body.Res = &types.AllocateIpv4AddressResponse{
		Returnval: ip,
	}

	return body
}

func (m *IpPoolManager) AllocateIpv6Address(req *types.AllocateIpv6Address) soap.HasFault {
	body := &methods.AllocateIpv6AddressBody{}

	pool, ok := m.pools[req.PoolId]
	if !ok {
		body.Fault_ = Fault("", &types.InvalidArgument{})
		return body
	}

	ip, err := pool.AllocateIPv6(req.AllocationId)
	if err != nil {
		body.Fault_ = Fault(err.Error(), &types.RuntimeFault{})
		return body
	}

	body.Res = &types.AllocateIpv6AddressResponse{
		Returnval: ip,
	}

	return body
}

func (m *IpPoolManager) ReleaseIpAllocation(req *types.ReleaseIpAllocation) soap.HasFault {
	body := &methods.ReleaseIpAllocationBody{}

	pool, ok := m.pools[req.PoolId]
	if !ok {
		body.Fault_ = Fault("", &types.InvalidArgument{})
		return body
	}

	pool.ReleaseIpv4(req.AllocationId)
	pool.ReleaseIpv6(req.AllocationId)

	body.Res = &types.ReleaseIpAllocationResponse{}

	return body
}

func (m *IpPoolManager) QueryIPAllocations(req *types.QueryIPAllocations) soap.HasFault {
	body := &methods.QueryIPAllocationsBody{}

	pool, ok := m.pools[req.PoolId]
	if !ok {
		body.Fault_ = Fault("", &types.InvalidArgument{})
		return body
	}

	body.Res = &types.QueryIPAllocationsResponse{}

	ipv4, ok := pool.ipv4Allocation[req.ExtensionKey]
	if ok {
		body.Res.Returnval = append(body.Res.Returnval, types.IpPoolManagerIpAllocation{
			IpAddress:    ipv4,
			AllocationId: req.ExtensionKey,
		})
	}

	ipv6, ok := pool.ipv6Allocation[req.ExtensionKey]
	if ok {
		body.Res.Returnval = append(body.Res.Returnval, types.IpPoolManagerIpAllocation{
			IpAddress:    ipv6,
			AllocationId: req.ExtensionKey,
		})
	}

	return body
}

var (
	errNoIpAvailable     = errors.New("no ip address available")
	errInvalidAllocation = errors.New("allocation id not recognized")
)

type IpPool struct {
	config         *types.IpPool
	ipv4Allocation map[string]string
	ipv6Allocation map[string]string
	ipv4Pool       []string
	ipv6Pool       []string
}

func MustNewIpPool(config *types.IpPool) *IpPool {
	pool, err := NewIpPool(config)
	if err != nil {
		panic(err)
	}

	return pool
}

func NewIpPool(config *types.IpPool) (*IpPool, error) {
	pool := &IpPool{
		config:         config,
		ipv4Allocation: make(map[string]string),
		ipv6Allocation: make(map[string]string),
	}

	return pool, pool.init()
}

func (p *IpPool) init() error {
	// IPv4 range
	if p.config.Ipv4Config != nil {
		ranges := strings.Split(p.config.Ipv4Config.Range, ",")
		for _, r := range ranges {
			sp := strings.Split(r, "#")
			if len(sp) != 2 {
				return fmt.Errorf("format of range should be ip#number; got %q", r)
			}

			ip := net.ParseIP(strings.TrimSpace(sp[0])).To4()
			if ip == nil {
				return fmt.Errorf("bad ip format: %q", sp[0])
			}

			length, err := strconv.Atoi(sp[1])
			if err != nil {
				return err
			}

			for i := 0; i < length; i++ {
				p.ipv4Pool = append(p.ipv4Pool, net.IPv4(ip[0], ip[1], ip[2], ip[3]+byte(i)).String())
			}
		}
	}

	// IPv6 range
	if p.config.Ipv6Config != nil {
		ranges := strings.Split(p.config.Ipv6Config.Range, ",")
		for _, r := range ranges {
			sp := strings.Split(r, "#")
			if len(sp) != 2 {
				return fmt.Errorf("format of range should be ip#number; got %q", r)
			}

			ip := net.ParseIP(strings.TrimSpace(sp[0])).To16()
			if ip == nil {
				return fmt.Errorf("bad ip format: %q", sp[0])
			}

			length, err := strconv.Atoi(sp[1])
			if err != nil {
				return err
			}

			for i := 0; i < length; i++ {
				var ipv6 [16]byte
				copy(ipv6[:], ip)
				ipv6[15] += byte(i)
				p.ipv6Pool = append(p.ipv6Pool, net.IP(ipv6[:]).String())
			}
		}
	}

	return nil
}

func (p *IpPool) AllocateIPv4(allocation string) (string, error) {
	if ip, ok := p.ipv4Allocation[allocation]; ok {
		return ip, nil
	}

	l := len(p.ipv4Pool)
	if l == 0 {
		return "", errNoIpAvailable
	}

	ip := p.ipv4Pool[l-1]

	p.config.AvailableIpv4Addresses--
	p.ipv4Pool = p.ipv4Pool[:l-1]
	p.ipv4Allocation[allocation] = ip

	return ip, nil
}

func (p *IpPool) ReleaseIpv4(allocation string) error {
	ip, ok := p.ipv4Allocation[allocation]
	if !ok {
		return errInvalidAllocation
	}

	delete(p.ipv4Allocation, allocation)
	p.config.AvailableIpv4Addresses++
	p.ipv4Pool = append(p.ipv4Pool, ip)

	return nil
}

func (p *IpPool) AllocateIPv6(allocation string) (string, error) {
	if ip, ok := p.ipv6Allocation[allocation]; ok {
		return ip, nil
	}

	l := len(p.ipv6Pool)
	if l == 0 {
		return "", errNoIpAvailable
	}

	ip := p.ipv6Pool[l-1]

	p.config.AvailableIpv6Addresses--
	p.ipv6Pool = p.ipv6Pool[:l-1]
	p.ipv6Allocation[allocation] = ip

	return ip, nil
}

func (p *IpPool) ReleaseIpv6(allocation string) error {
	ip, ok := p.ipv6Allocation[allocation]
	if !ok {
		return errInvalidAllocation
	}

	delete(p.ipv6Allocation, allocation)
	p.config.AvailableIpv6Addresses++
	p.ipv6Pool = append(p.ipv6Pool, ip)

	return nil
}
