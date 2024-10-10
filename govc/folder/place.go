/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package folder

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

type place struct {
	*flags.ClientFlag
	*flags.DatacenterFlag
	*flags.VirtualMachineFlag
	*flags.OutputFlag

	pool          flags.StringList
	placementType string
}

func init() {
	cli.Register("folder.place", &place{}, true)
}

func (cmd *place) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.Var(&cmd.pool, "pool", "Resource Pools to use for placement.")
	f.StringVar(&cmd.placementType, "placementType", "", "Type of placement: relocate or reconfigure.")
}

func (cmd *place) Usage() string {
	return "PATH..."
}

func (cmd *place) Description() string {
	return `Get a placement recommendation for an existing VM

Examples:
  govc folder.place -rp $rp1Name -rp $rp2Name -rp $rp3Name-vm $vmName`
}

func (cmd *place) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *place) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	// Use latest version to pick up latest PlaceVmsXCluster API.
	err = c.UseServiceVersion()
	if err != nil {
		return err
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	placementType := types.PlaceVmsXClusterSpecPlacementType(cmd.placementType)

	var relocateSpec *types.VirtualMachineRelocateSpec

	// TODO: Support createAndPowerOn and reconfigure.
	switch placementType {
	case types.PlaceVmsXClusterSpecPlacementTypeRelocate:
		relocateSpec = &types.VirtualMachineRelocateSpec{}
		break
	default:
		return errors.New("please specify a valid and supported placement type")
	}

	// PlaceVMsXCluster is only valid against the root folder.
	folder := object.NewRootFolder(c)

	refs := make([]types.ManagedObjectReference, 0, len(cmd.pool))

	for _, arg := range cmd.pool {
		rp, err := finder.ResourcePool(ctx, arg)
		if err != nil {
			return err
		}

		refs = append(refs, rp.Reference())
	}

	vmPlacementSpecs := []types.PlaceVmsXClusterSpecVmPlacementSpec{{
		Vm:           types.NewReference(vm.Reference()),
		ConfigSpec:   types.VirtualMachineConfigSpec{},
		RelocateSpec: relocateSpec,
	}}

	placementSpec := types.PlaceVmsXClusterSpec{
		ResourcePools:           refs,
		PlacementType:           cmd.placementType,
		VmPlacementSpecs:        vmPlacementSpecs,
		HostRecommRequired:      types.NewBool(true),
		DatastoreRecommRequired: types.NewBool(true),
	}

	res, err := folder.PlaceVmsXCluster(ctx, placementSpec)
	if err != nil {
		return err
	}

	vimClient, err := cmd.ClientFlag.Client()
	if err != nil {
		return err
	}

	return cmd.WriteResult(&placementResult{res, vimClient, ctx, cmd.VirtualMachineFlag})
}

type placementResult struct {
	Result    *types.PlaceVmsXClusterResult `json:"result,omitempty"`
	vimClient *vim25.Client
	ctx       context.Context
	vm        *flags.VirtualMachineFlag
}

func (res *placementResult) Dump() interface{} {
	return res.Result
}

func (res *placementResult) initialPlacementAction(w io.Writer, pinfo types.PlaceVmsXClusterResultPlacementInfo, action *types.ClusterClusterInitialPlacementAction) error {

	spec := action.ConfigSpec
	if spec == nil {
		return nil
	}

	fields := []struct {
		name string
		moid *types.ManagedObjectReference
	}{
		{"Vm", pinfo.Vm},
		{"  Target", pinfo.Recommendation.Target},
		{"  TargetHost", action.TargetHost},
		{"  Pool", &action.Pool},
	}

	for _, f := range fields {
		if f.moid == nil {
			continue
		}
		path, err := find.InventoryPath(res.ctx, res.vimClient, *f.moid)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%s:\t%s\n", f.name, path)
	}

	return nil
}

func (res *placementResult) relocatePlacementAction(w io.Writer, pinfo types.PlaceVmsXClusterResultPlacementInfo, action *types.ClusterClusterRelocatePlacementAction) error {

	spec := action.RelocateSpec
	if spec == nil {
		return nil
	}

	fields := []struct {
		name string
		moid *types.ManagedObjectReference
	}{
		{"Vm", pinfo.Vm},
		{"  Target", pinfo.Recommendation.Target},
		{"  Folder", spec.Folder},
		{"  Datastore", spec.Datastore},
		{"  Pool", spec.Pool},
		{"  Host", spec.Host},
	}

	for _, f := range fields {
		if f.moid == nil {
			continue
		}
		path, err := find.InventoryPath(res.ctx, res.vimClient, *f.moid)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%s:\t%s\n", f.name, path)
	}

	return nil
}

func (res *placementResult) reconfigurePlacementAction(w io.Writer, pinfo types.PlaceVmsXClusterResultPlacementInfo, action *types.ClusterClusterReconfigurePlacementAction) error {

	spec := action.ConfigSpec
	if spec == nil {
		return nil
	}

	fields := []struct {
		name string
		moid *types.ManagedObjectReference
	}{
		{"Vm", pinfo.Vm},
		{"  Target", pinfo.Recommendation.Target},
		{"  TargetHost", action.TargetHost},
		{"  Pool", &action.Pool},
	}

	for _, f := range fields {
		if f.moid == nil {
			continue
		}
		path, err := find.InventoryPath(res.ctx, res.vimClient, *f.moid)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%s:\t%s\n", f.name, path)
	}

	return nil
}

func (res *placementResult) placementFault(w io.Writer, pfault types.PlaceVmsXClusterResultPlacementFaults, fault *types.LocalizedMethodFault) error {

	fields := []struct {
		name    string
		message string
		moid    *types.ManagedObjectReference
	}{
		{"Vm", "", pfault.Vm},
		{"  Message", fault.LocalizedMessage, nil},
	}

	for _, f := range fields {
		if f.moid == nil {
			if f.message != "" {
				fmt.Fprintf(w, "%s:\t%s\n", f.name, f.message)
			}
			continue
		}
		path, err := find.InventoryPath(res.ctx, res.vimClient, *f.moid)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%s:\t%s\n", f.name, path)
	}

	return nil
}

func (res placementResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, pinfo := range res.Result.PlacementInfos {

		for _, action := range pinfo.Recommendation.Action {

			if initPlaceAction, ok := action.(*types.ClusterClusterInitialPlacementAction); ok {
				err := res.initialPlacementAction(w, pinfo, initPlaceAction)
				if err != nil {
					return err
				}
			}

			if relocateAction, ok := action.(*types.ClusterClusterRelocatePlacementAction); ok {
				err := res.relocatePlacementAction(w, pinfo, relocateAction)
				if err != nil {
					return err
				}
			}

			if reconfigureAction, ok := action.(*types.ClusterClusterReconfigurePlacementAction); ok {
				err := res.reconfigurePlacementAction(w, pinfo, reconfigureAction)
				if err != nil {
					return err
				}
			}
		}
	}

	for _, pfault := range res.Result.Faults {

		for _, fault := range pfault.Faults {
			err := res.placementFault(w, pfault, &fault)
			if err != nil {
				return err
			}
		}
	}

	return tw.Flush()
}
