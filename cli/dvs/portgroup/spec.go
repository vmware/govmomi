// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package portgroup

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type VLANMode string

const (
	VLAN      = VLANMode("vlan")
	TrunkVLAN = VLANMode("trunking")
)

type DVPortgroupConfigSpec struct {
	types.DVPortgroupConfigSpec
}

var (
	vlanMode  string
	vlanId    int32
	vlanRange string
	vlanSpec  types.BaseVmwareDistributedVirtualSwitchVlanSpec
)

func (spec *DVPortgroupConfigSpec) Register(ctx context.Context, f *flag.FlagSet) {
	ptypes := types.DistributedVirtualPortgroupPortgroupType("").Strings()

	vlanModes := []string{
		string(VLAN),
		string(TrunkVLAN),
	}

	f.StringVar(&spec.Type, "type", ptypes[0],
		fmt.Sprintf("Portgroup type (%s)", strings.Join(ptypes, "|")))
	f.Var(flags.NewInt32(&spec.NumPorts), "nports", "Number of ports")
	f.Var(flags.NewInt32(&vlanId), "vlan", "VLAN ID")
	f.StringVar(&vlanRange, "vlan-range", "0-4094", "VLAN Ranges with comma delimited")
	f.StringVar(&vlanMode, "vlan-mode", "vlan", fmt.Sprintf("vlan mode (%s)", strings.Join(vlanModes, "|")))
	f.Var(flags.NewOptionalBool(&spec.AutoExpand), "auto-expand", "Ignore the limit on the number of ports")
}

func getRange(vlanRange string) []types.NumericRange {
	nRanges := make([]types.NumericRange, 0)
	strRanges := strings.Split(vlanRange, ",")
	for _, v := range strRanges {
		vlans := strings.Split(v, "-")
		if len(vlans) != 2 {
			panic(fmt.Sprintf("range %s does not follow format with vlanId-vlanId", v))
		}
		start, err := strconv.Atoi(vlans[0])
		if err != nil {
			panic(fmt.Sprintf("range %s does not follow format with vlanId-vlanId error: %v", v, err))
		}
		end, err := strconv.Atoi(vlans[1])
		if err != nil {
			panic(fmt.Sprintf("range %s does not follow format with vlanId-vlanId error: %v", v, err))
		}
		nRanges = append(nRanges, types.NumericRange{
			Start: int32(start),
			End:   int32(end),
		})
	}
	return nRanges
}

func (spec *DVPortgroupConfigSpec) Spec() types.DVPortgroupConfigSpec {
	config := new(types.VMwareDVSPortSetting)
	switch VLANMode(vlanMode) {
	case VLAN:
		spec := new(types.VmwareDistributedVirtualSwitchVlanIdSpec)
		spec.VlanId = vlanId
		vlanSpec = spec
	case TrunkVLAN:
		spec := new(types.VmwareDistributedVirtualSwitchTrunkVlanSpec)
		spec.VlanId = getRange(vlanRange)
		vlanSpec = spec
	}
	config.Vlan = vlanSpec
	spec.DefaultPortConfig = config
	return spec.DVPortgroupConfigSpec
}
