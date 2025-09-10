// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package folder

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

var allTypes = []string{}

var createAndPowerOnTypes = []string{
	string(types.PlaceVmsXClusterSpecPlacementTypeCreateAndPowerOn),
}

var relocateTypes = []string{
	string(types.PlaceVmsXClusterSpecPlacementTypeRelocate),
}

var reconfigureTypes = []string{
	string(types.PlaceVmsXClusterSpecPlacementTypeReconfigure),
}

func init() {
	allTypes = append(allTypes, createAndPowerOnTypes...)
	allTypes = append(allTypes, relocateTypes...)
	allTypes = append(allTypes, reconfigureTypes...)
}

type typeFlag string

func (t *typeFlag) Set(s string) error {
	s = strings.ToLower(s)
	for _, e := range allTypes {
		if s == strings.ToLower(e) {
			*t = typeFlag(e)
			return nil
		}
	}

	return fmt.Errorf("unknown type")
}

func (t *typeFlag) String() string {
	return string(*t)
}

func (t *typeFlag) partOf(m []string) bool {
	for _, e := range m {
		if t.String() == e {
			return true
		}
	}
	return false
}

func (t *typeFlag) IsCreateAndPowerOnType() bool {
	return t.partOf(createAndPowerOnTypes)
}

func (t *typeFlag) IsRelocateType() bool {
	return t.partOf(relocateTypes)
}

func (t *typeFlag) IsReconfigureType() bool {
	return t.partOf(reconfigureTypes)
}

type place struct {
	*flags.ClientFlag
	*flags.DatacenterFlag
	*flags.VirtualMachineFlag
	*flags.OutputFlag

	pool              flags.StringList
	Type              typeFlag
	CandidateNetworks flags.StringList
}

func init() {
	cli.Register("folder.place", &place{}, false)
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
	f.Var(&cmd.Type, "type", fmt.Sprintf("Placement type (%s)", strings.Join(allTypes, "|")))
	f.Var(&cmd.CandidateNetworks, "candidate-networks", "Candidate network names (repeat for multiple nics)")
}

func (cmd *place) Usage() string {
	return "PATH..."
}

func (cmd *place) Description() string {
	return `Get a placement recommendation for an existing VM

Examples:
  govc folder.place -rp $rp1Name -rp $rp2Name -rp $rp3Name -vm $vmName -type relocate -candidate-networks "netA|netB" -candidate-networks "netC"
  Each use of the "-candidate-networks" flag corresponds to one NIC of the VM.
  Within the value of each flag, use "|" to specify multiple candidate networks for that NIC.
  For example:
      -candidate-networks "netA|netB" → NIC 0 can connect to either netA or netB.
      -candidate-networks "netC"      → NIC 1 can connect only to netC.
`
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
	client, err := cmd.Client()
	if err != nil {
		return err
	}
	// Use latest version to pick up latest PlaceVmsXCluster API.
	err = client.UseServiceVersion()
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

	var relocateSpec *types.VirtualMachineRelocateSpec

	// TODO: Support createAndPowerOn and reconfigure.
	switch {
	case cmd.Type.IsRelocateType():
		relocateSpec = &types.VirtualMachineRelocateSpec{}
		break
	case cmd.Type.IsReconfigureType():
		return errors.New("reconfigure is currently an unsupported placement type")
	case cmd.Type.IsCreateAndPowerOnType():
		return errors.New("createAndPowerOn is currently an unsupported placement type")
	default:
		return errors.New("please specify a valid type")
	}
	var candidateNetworks []types.PlaceVmsXClusterSpecVmPlacementSpecCandidateNetworks
	if len(cmd.CandidateNetworks) > 0 {
		client, err := cmd.Client()
		if err != nil {
			return err
		}
		finder := find.NewFinder(client, false)

		dc, err := cmd.Datacenter()
		if err != nil {
			return err
		}
		finder.SetDatacenter(dc)

		for _, nic := range cmd.CandidateNetworks {
			// Each 'nic' string is like "netA|netB"
			netNames := strings.Split(nic, "|")
			var refs []types.ManagedObjectReference

			for _, name := range netNames {
				fmt.Printf(">>> Looking up network: %q\n", name)
				netObj, err := finder.Network(ctx, name)
				if err != nil {
					return fmt.Errorf("network %q not found: %w", name, err)
				}
				refs = append(refs, netObj.Reference())
			}
			candidateNetworks = append(candidateNetworks, types.PlaceVmsXClusterSpecVmPlacementSpecCandidateNetworks{
				Networks: refs,
			})
		}
	}

	// PlaceVMsXCluster is only valid against the root folder.
	folder := object.NewRootFolder(client)

	refs := make([]types.ManagedObjectReference, 0, len(cmd.pool))

	for _, arg := range cmd.pool {
		rp, err := finder.ResourcePool(ctx, arg)
		if err != nil {
			return err
		}

		refs = append(refs, rp.Reference())
	}

	vmPlacementSpecs := []types.PlaceVmsXClusterSpecVmPlacementSpec{{
		Vm:                types.NewReference(vm.Reference()),
		ConfigSpec:        types.VirtualMachineConfigSpec{},
		RelocateSpec:      relocateSpec,
		CandidateNetworks: candidateNetworks,
	}}

	placementSpec := types.PlaceVmsXClusterSpec{
		ResourcePools:           refs,
		PlacementType:           cmd.Type.String(),
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

func (res *placementResult) Dump() any {
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

	// Display the available network references from the placement recommendation.
	if act, ok := any(action).(*types.ClusterClusterInitialPlacementAction); ok {
		if len(act.AvailableNetworks) > 0 {
			fmt.Fprintf(w, "  AvailableNetworks:\n")
			for _, net := range act.AvailableNetworks {
				path, err := find.InventoryPath(res.ctx, res.vimClient, net)
				if err != nil {
					return err
				}
				fmt.Fprintf(w, "\t- %s\n", path)
			}
		}
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

	// Display the available network references from the placement recommendation.
	if len(action.AvailableNetworks) > 0 {
		fmt.Fprintf(w, "  AvailableNetworks:\n")
		for _, net := range action.AvailableNetworks {
			path, err := find.InventoryPath(res.ctx, res.vimClient, net)
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "\t- %s\n", path)
		}
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
