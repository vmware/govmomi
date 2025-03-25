// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	esxcli "github.com/vmware/govmomi/cli/esx"
	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

type DynamicTypeManager struct {
	types.ManagedObjectReference
}

type ManagedMethodExecuter struct {
	types.ManagedObjectReference

	h *HostSystem
}

var esxcliNotFound = &types.LocalizedMethodFault{
	Fault:            new(types.NotFound),
	LocalizedMessage: "The object or item referred to could not be found.",
}

func esxcliFault(msg ...string) *types.LocalizedMethodFault {
	return &types.LocalizedMethodFault{
		Fault:            &internal.VimEsxCLICLIFault{ErrMsg: msg},
		LocalizedMessage: "EsxCLI.CLIFault.summary",
	}
}

func (h *HostSystem) RetrieveManagedMethodExecuter(ctx *Context, req *internal.RetrieveManagedMethodExecuterRequest) soap.HasFault {
	if h.mme == nil {
		h.mme = &ManagedMethodExecuter{
			types.ManagedObjectReference{Type: "ManagedMethodExecuter", Value: h.Self.Value},
			h,
		}
		ctx.Map.Put(h.mme)
	}

	return &internal.RetrieveManagedMethodExecuterBody{
		Res: &internal.RetrieveManagedMethodExecuterResponse{
			Returnval: &internal.ReflectManagedMethodExecuter{
				ManagedObjectReference: h.mme.Reference(),
			},
		},
	}
}

func (h *HostSystem) RetrieveDynamicTypeManager(ctx *Context, req *internal.RetrieveDynamicTypeManagerRequest) soap.HasFault {
	if h.dtm == nil {
		h.dtm = &DynamicTypeManager{
			types.ManagedObjectReference{Type: "DynamicTypeManager", Value: h.Self.Value},
		}
		ctx.Map.Put(h.dtm)
	}

	return &internal.RetrieveDynamicTypeManagerBody{
		Res: &internal.RetrieveDynamicTypeManagerResponse{
			Returnval: &internal.InternalDynamicTypeManager{
				ManagedObjectReference: h.dtm.Reference(),
			},
		},
	}
}

func (*DynamicTypeManager) DynamicTypeMgrQueryTypeInfo(ctx *Context, req *internal.DynamicTypeMgrQueryTypeInfoRequest) soap.HasFault {
	all := esx.TypeInfo

	if spec, ok := req.FilterSpec.(*internal.DynamicTypeMgrTypeFilterSpec); ok {
		all = internal.DynamicTypeMgrAllTypeInfo{}

		for _, info := range esx.TypeInfo.DataTypeInfo {
			if strings.Contains(info.Name, spec.TypeSubstr) {
				all.DataTypeInfo = append(all.DataTypeInfo, info)
			}
		}

		for _, info := range esx.TypeInfo.EnumTypeInfo {
			if strings.Contains(info.Name, spec.TypeSubstr) {
				all.EnumTypeInfo = append(all.EnumTypeInfo, info)
			}
		}

		for _, info := range esx.TypeInfo.ManagedTypeInfo {
			if strings.Contains(info.Name, spec.TypeSubstr) {
				all.ManagedTypeInfo = append(all.ManagedTypeInfo, info)
			}
		}
	}

	body := &internal.DynamicTypeMgrQueryTypeInfoBody{
		Res: &internal.DynamicTypeMgrQueryTypeInfoResponse{
			Returnval: all,
		},
	}

	return body
}

func (m *DynamicTypeManager) DynamicTypeMgrQueryMoInstances(ctx *Context, req *internal.DynamicTypeMgrQueryMoInstancesRequest) soap.HasFault {
	body := &internal.DynamicTypeMgrQueryMoInstancesBody{
		Res: &internal.DynamicTypeMgrQueryMoInstancesResponse{
			Returnval: nil,
		},
	}

	return body
}

func (m *ManagedMethodExecuter) VimCLIInfoFetchCLIInfo(_ *Context, args esxcli.Values) (*esxcli.CommandInfo, *types.LocalizedMethodFault) {
	kind := args.Value("typeName")
	kind = strings.TrimPrefix(kind, "vim.EsxCLI.")

	for _, info := range esx.CommandInfo {
		if info.CommandInfoItem.Name == kind {
			return &info, nil
		}
	}

	return nil, esxcliNotFound
}

// sample from: govc host.esxcli -dump software vib get
var softwareVib = []esxcli.Values{
	{
		"AcceptanceLevel":         []string{"VMwareCertified"},
		"CreationDate":            []string{"2023-03-22"},
		"Depends":                 []string{"esx-version >= 5.0.0"},
		"Description":             []string{"An embedded web UI for ESXi"},
		"ID":                      []string{"VMware_bootbank_esx-ui_2.12.0-21482143"},
		"InstallDate":             []string{"2023-03-28"},
		"LiveInstallAllowed":      []string{"True"},
		"LiveRemoveAllowed":       []string{"True"},
		"MaintenanceModeRequired": []string{"False"},
		"Name":                    []string{"esx-ui"},
		"Overlay":                 []string{"False"},
		"Payloads":                []string{"esx-ui"},
		"Platforms":               []string{"host"},
		"StatelessReady":          []string{"True"},
		"Status":                  []string{""},
		"Summary":                 []string{"VMware Host Client"},
		"Type":                    []string{"bootbank"},
		"Vendor":                  []string{"VMware"},
		"Version":                 []string{"2.12.0-21482143"},
	},
	{
		"AcceptanceLevel":         []string{"VMwareCertified"},
		"CreationDate":            []string{"2023-03-25"},
		"Depends":                 []string{"vmkapi_2_11_0_0", "vmkapi_incompat_2_11_0_0"},
		"Description":             []string{"Intel DW GPIO controller driver"},
		"ID":                      []string{"VMW_bootbank_intelgpio_0.1-1vmw.801.0.0.21495797"},
		"InstallDate":             []string{"2023-03-28"},
		"LiveInstallAllowed":      []string{"False"},
		"LiveRemoveAllowed":       []string{"False"},
		"MaintenanceModeRequired": []string{"True"},
		"Name":                    []string{"intelgpio"},
		"Overlay":                 []string{"False"},
		"Payloads":                []string{"intelgpi"},
		"Platforms":               []string{"host"},
		"StatelessReady":          []string{"True"},
		"Status":                  []string{""},
		"Summary":                 []string{"VMware Esx VIB"},
		"Tags":                    []string{"RestrictStickyFiles", "module", "driver", "sdkname:esx", "sdkversion:8.0.1-21495797"},
		"Type":                    []string{"bootbank"},
		"Vendor":                  []string{"VMW"},
		"Version":                 []string{"0.1-1vmw.801.0.0.21495797"},
	},
}

func (m *ManagedMethodExecuter) VimEsxCLISoftwareVibGet(_ *Context, args esxcli.Values) (*esxcli.Response, *types.LocalizedMethodFault) {
	r := esxcli.Response{Kind: "VIBExt"}

	name := args.Value("vibname")

	if name != "" {
		for _, vib := range softwareVib {
			if vib.Value("Name") == name {
				r.Values = append(r.Values, vib)
				return &r, nil
			}
		}
		return nil, esxcliFault("[NoMatchError]", "id="+name)
	}

	r.Values = softwareVib

	return &r, nil
}

func (m *ManagedMethodExecuter) VimEsxCLISoftwareVibList(_ *Context, args esxcli.Values) (*esxcli.Response, *types.LocalizedMethodFault) {
	r := esxcli.Response{Kind: "SummaryExt"}

	r.Values = softwareVib // TODO: subset of VibGet fields

	return &r, nil
}

func (m *ManagedMethodExecuter) VimEsxCLIHardwareClockGet(_ *Context, _ esxcli.Values) (*esxcli.Response, *types.LocalizedMethodFault) {
	return &esxcli.Response{
		Kind:   "string",
		String: time.Now().UTC().Format(time.RFC3339),
	}, nil
}

func (m *ManagedMethodExecuter) VimEsxCLINetworkFirewallGet(ctx *Context, args esxcli.Values) (*esxcli.Response, *types.LocalizedMethodFault) {
	r := esxcli.Response{
		Kind: "Firewall",
		Values: []esxcli.Values{{
			"DefaultAction": []string{"DROP"},
			"Enabled":       []string{"false"},
			"Loaded":        []string{"true"},
		}},
	}

	return &r, nil
}

func (m *ManagedMethodExecuter) VimEsxCLINetworkVmList(ctx *Context, _ esxcli.Values) (*esxcli.Response, *types.LocalizedMethodFault) {
	r := esxcli.Response{Kind: "VM"}

	for _, ref := range m.h.Vm {
		vm := ctx.Map.Get(ref).(*VirtualMachine)

		var networks []string
		for _, ref := range vm.Network {
			name := entityName(ctx.Map.Get(ref).(mo.Entity))
			networks = append(networks, name)
		}

		r.Values = append(r.Values, esxcli.Values{
			"Name":     []string{vm.Name},
			"Networks": networks,
			"NumPorts": []string{strconv.Itoa(len(networks))},
			"WorldID":  []string{strconv.Itoa(vm.worldID())},
		})
	}

	return &r, nil
}

// sample from: govc host.esxcli -dump network ip connection list
var networkIpConnectionList = []esxcli.Values{
	{
		"CCAlgo":         []string{"newreno"},
		"ForeignAddress": []string{"0.0.0.0:0"},
		"LocalAddress":   []string{"0.0.0.0:443"},
		"Proto":          []string{"tcp"},
		"RecvQ":          []string{"0"},
		"SendQ":          []string{"0"},
		"State":          []string{"LISTEN"},
		"WorldID":        []string{"525276"},
		"WorldName":      []string{"envoy"},
	},
	{
		"CCAlgo":         []string{""},
		"ForeignAddress": []string{"0.0.0.0:0"},
		"LocalAddress":   []string{"127.0.0.1:123"},
		"Proto":          []string{"udp"},
		"RecvQ":          []string{"0"},
		"SendQ":          []string{"0"},
		"State":          []string{""},
		"WorldID":        []string{"530726"},
		"WorldName":      []string{"ntpd"},
	},
}

func (m *ManagedMethodExecuter) VimEsxCLINetworkIpConnectionList(_ *Context, args esxcli.Values) (*esxcli.Response, *types.LocalizedMethodFault) {
	r := esxcli.Response{Kind: "IpConnection"}

	kind := args.Value("type")
	if kind != "" && kind != "tcp" { // ip, tcp, udp, all
		return nil, esxcliFault("Invalid data constraint for parameter 'type'.")
	}

	r.Values = networkIpConnectionList

	return &r, nil
}

// sample from: govc host.esxcli -dump system settings advanced list
var systemSettingsAdvancedList = []esxcli.Values{
	{
		"DefaultIntValue": []string{"2"},
		"Description":     []string{"PShare salting allows for sharing isolation between multiple VM"},
		"HostSpecific":    []string{"false"},
		"Impact":          []string{"none"},
		"IntValue":        []string{"2"},
		"MaxValue":        []string{"2"},
		"MinValue":        []string{"0"},
		"Path":            []string{"/Mem/ShareForceSalting"},
		"Type":            []string{"integer"},
	},
	{
		"DefaultIntValue": []string{"0"},
		"Description":     []string{"Enable guest arp inspection IOChain to get IP"},
		"HostSpecific":    []string{"false"},
		"Impact":          []string{"none"},
		"IntValue":        []string{"1"},
		"MaxValue":        []string{"1"},
		"MinValue":        []string{"0"},
		"Path":            []string{"/Net/GuestIPHack"},
		"Type":            []string{"integer"},
	},
}

func (m *ManagedMethodExecuter) VimEsxCLISystemSettingsAdvancedList(_ *Context, args esxcli.Values) (*esxcli.Response, *types.LocalizedMethodFault) {
	r := esxcli.Response{Kind: "SettingsAdvancedOption"}

	option := args.Value("option")
	if option != "" {
		for _, s := range systemSettingsAdvancedList {
			if s.Value("Path") == option {
				r.Values = append(r.Values, s)
				return &r, nil
			}
		}
		return nil, esxcliFault("Unable to find option")
	}

	r.Values = systemSettingsAdvancedList

	return &r, nil
}

func (m *ManagedMethodExecuter) VimEsxCLIVmProcessList(ctx *Context, _ esxcli.Values) (*esxcli.Response, *types.LocalizedMethodFault) {
	r := esxcli.Response{Kind: "VirtualMachine"}

	for _, ref := range m.h.Vm {
		vm := ctx.Map.Get(ref).(*VirtualMachine)

		r.Values = append(r.Values, esxcli.Values{
			"ConfigFile":  []string{vm.Config.Files.VmPathName},
			"DisplayName": []string{vm.Name},
			"ProcessID":   []string{"0"},
			"UUID":        []string{vm.uid.String()},
			"VMXCartelID": []string{strconv.Itoa(vm.worldID() + 1)},
			"WorldID":     []string{strconv.Itoa(vm.worldID())},
		})
	}

	return &r, nil
}

func (m *ManagedMethodExecuter) VimEsxCLIIscsiSoftwareGet(_ *Context, _ esxcli.Values) (*esxcli.Response, *types.LocalizedMethodFault) {
	return &esxcli.Response{Kind: "boolean", String: "false"}, nil
}

var boot = time.Now()

func (m *ManagedMethodExecuter) VimEsxCLISystemStatsUptimeGet(_ *Context, _ esxcli.Values) (*esxcli.Response, *types.LocalizedMethodFault) {
	uptime := fmt.Sprintf("%d", time.Since(boot))
	return &esxcli.Response{Kind: "long", String: uptime}, nil
}

func (_ *ManagedMethodExecuter) toXML(v any) string {
	var out bytes.Buffer

	err := xml.NewEncoder(&out).Encode(v)
	if err != nil {
		panic(err)
	}

	return out.String()
}

func (m *ManagedMethodExecuter) ExecuteSoap(ctx *Context, req *internal.ExecuteSoapRequest) soap.HasFault {
	res := new(internal.ReflectManagedMethodExecuterSoapResult)

	args := esxcli.Values{}
	for _, arg := range req.Argument {
		args[arg.Name] = arg.Value()
	}

	name := internal.EsxcliName(req.Method)
	method := reflect.ValueOf(m).MethodByName(name)

	var val types.AnyType
	err := esxcliNotFound

	if method.IsValid() {
		ret := method.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(args)})
		val = ret[0].Interface()
		err = ret[1].Interface().(*types.LocalizedMethodFault)
	}

	if err == nil {
		if r, ok := val.(*esxcli.Response); ok {
			if r.String == "" {
				// DataObject xsi:type has method name prefix
				r.Kind = strings.ReplaceAll(ucFirst(req.Method), ".", "") + r.Kind
			}
		}
		res.Response = m.toXML(val)
	} else {
		res.Fault = &internal.ReflectManagedMethodExecuterSoapFault{
			FaultMsg:    err.LocalizedMessage,
			FaultDetail: m.toXML(err),
		}
	}

	return &internal.ExecuteSoapBody{
		Res: &internal.ExecuteSoapResponse{
			Returnval: res,
		},
	}
}
