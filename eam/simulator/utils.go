// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"

	"github.com/vmware/govmomi/eam/types"
	vimobj "github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator"
	vimmo "github.com/vmware/govmomi/vim25/mo"
	vim "github.com/vmware/govmomi/vim25/types"
)

var (
	// fsOrHTTPRx matches a URI that is either on the local file
	// system or an HTTP endpoint.
	fsOrHTTPRx = regexp.MustCompile(`^(?:\.?/)|(?:http[s]?:)`)
)

var (
	agentVmDatastoreEmptyErr     = errors.New("AgentVmDatastore is empty")
	agentVmNetworkEmptyErr       = errors.New("AgentVmNetwork is empty")
	foldersEmptyErr              = errors.New("Folders is empty")
	scopeComputeResourceEmptyErr = errors.New("Scope.ComputeResource is empty")
	poolRefNilErr                = errors.New("Unable to determine ResourcePool from ComputeResource")
)

func getAgentVMPlacementOptions(
	ctx *simulator.Context,
	reg *simulator.Registry,
	r *rand.Rand, i int,
	baseAgencyConfig types.BaseAgencyConfigInfo) (AgentVMPlacementOptions, error) {

	var opts AgentVMPlacementOptions
	agencyConfig := baseAgencyConfig.GetAgencyConfigInfo()
	if l := len(agencyConfig.AgentVmDatastore); l == 0 {
		return opts, agentVmDatastoreEmptyErr
	} else if l == len(agencyConfig.AgentConfig) {
		opts.datastore = agencyConfig.AgentVmDatastore[i]
	} else {
		opts.datastore = agencyConfig.AgentVmDatastore[r.Intn(l)]
	}

	if l := len(agencyConfig.AgentVmNetwork); l == 0 {
		return opts, agentVmNetworkEmptyErr
	} else if l == len(agencyConfig.AgentConfig) {
		opts.network = agencyConfig.AgentVmNetwork[i]
	} else {
		opts.network = agencyConfig.AgentVmNetwork[r.Intn(l)]
	}

	if l := len(agencyConfig.Folders); l == 0 {
		return opts, foldersEmptyErr
	} else if l == len(agencyConfig.AgentConfig) {
		opts.folder = agencyConfig.Folders[i].FolderId
		opts.datacenter = agencyConfig.Folders[i].DatacenterId
	} else {
		j := r.Intn(l)
		opts.folder = agencyConfig.Folders[j].FolderId
		opts.datacenter = agencyConfig.Folders[j].DatacenterId
	}

	if l := len(agencyConfig.ResourcePools); l == 0 {
		switch tscope := agencyConfig.Scope.(type) {
		case *types.AgencyComputeResourceScope:
			crRefs := tscope.ComputeResource
			if l := len(crRefs); l == 0 {
				return opts, scopeComputeResourceEmptyErr
			} else if l == len(agencyConfig.AgentConfig) {
				opts.computeResource = crRefs[i]
			} else {
				opts.computeResource = crRefs[r.Intn(l)]
			}
			poolRef, err := getPoolFromComputeResource(ctx, reg, opts.computeResource)
			if err != nil {
				return opts, err
			}
			opts.pool = *poolRef
		}
	} else if l == len(agencyConfig.AgentConfig) {
		opts.pool = agencyConfig.ResourcePools[i].ResourcePoolId
		opts.computeResource = agencyConfig.ResourcePools[i].ComputeResourceId
	} else {
		j := r.Intn(l)
		opts.pool = agencyConfig.ResourcePools[j].ResourcePoolId
		opts.computeResource = agencyConfig.ResourcePools[j].ComputeResourceId
	}

	hosts := getHostsFromPool(ctx, reg, opts.pool)
	if l := len(hosts); l == 0 {
		return opts, fmt.Errorf("HostSystems not found for %v", opts.pool)
	} else if l == len(agencyConfig.AgentConfig) {
		opts.host = hosts[i]
	} else {
		opts.host = hosts[r.Intn(l)]
	}

	return opts, nil
}

func getPoolFromComputeResource(
	ctx *simulator.Context,
	reg *simulator.Registry,
	computeResource vim.ManagedObjectReference) (*vim.ManagedObjectReference, error) {

	var poolRef *vim.ManagedObjectReference
	crObj := reg.Get(computeResource)
	if crObj == nil {
		return nil, fmt.Errorf("%v not in registry", computeResource)
	}
	ctx.WithLock(crObj, func() {
		switch cr := crObj.(type) {
		case *vimmo.ComputeResource:
			poolRef = cr.ResourcePool
		case *simulator.ClusterComputeResource:
			poolRef = cr.ResourcePool
		default:
			panic(fmt.Errorf(
				"%v is not a %s or %s",
				crObj,
				"*mo.ComputeResource",
				"*simulator.ClusterComputeResource",
			))
		}
	})
	if poolRef == nil {
		return nil, poolRefNilErr
	}
	return poolRef, nil
}

// getHostsFromPool returns the host(s) for the provided compute resource.
func getHostsFromPool(
	ctx *simulator.Context,
	reg *simulator.Registry,
	poolRef vim.ManagedObjectReference) []vim.ManagedObjectReference {

	pool := reg.Get(poolRef).(vimmo.Entity)
	cr := getEntityComputeResource(reg, pool)

	var hosts []vim.ManagedObjectReference

	ctx.WithLock(cr, func() {
		switch cr := cr.(type) {
		case *vimmo.ComputeResource:
			hosts = cr.Host
		case *simulator.ClusterComputeResource:
			hosts = cr.Host
		}
	})

	return hosts
}

// hostsWithDatastore returns hosts that have access to the given datastore path
func hostsWithDatastore( // nolint:unused nolint:deadcode
	reg *simulator.Registry,
	hosts []vim.ManagedObjectReference, path string) []vim.ManagedObjectReference {

	attached := hosts[:0]
	var p vimobj.DatastorePath
	p.FromString(path)

	for _, host := range hosts {
		h := reg.Get(host).(*simulator.HostSystem)
		if reg.FindByName(p.Datastore, h.Datastore) != nil {
			attached = append(attached, host)
		}
	}

	return attached
}

// getEntityComputeResource returns the ComputeResource parent for the given item.
// A ResourcePool for example may have N Parents of type ResourcePool, but the top
// most Parent pool is always a ComputeResource child.
func getEntityComputeResource(
	reg *simulator.Registry,
	item vimmo.Entity) vimmo.Entity {

	for {
		parent := item.Entity().Parent
		item = reg.Get(*parent).(vimmo.Entity)
		switch item.Reference().Type {
		case "ComputeResource":
			return item
		case "ClusterComputeResource":
			return item
		}
	}
}
