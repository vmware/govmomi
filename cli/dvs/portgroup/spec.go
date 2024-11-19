/*
Copyright (c) 2016-2024 VMware, Inc. All Rights Reserved.

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
