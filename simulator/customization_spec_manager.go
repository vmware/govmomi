// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"encoding/pem"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/vmware/govmomi/simulator/internal"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

var DefaultCustomizationSpec = []types.CustomizationSpecItem{
	{
		Info: types.CustomizationSpecInfo{
			Name:           "vcsim-linux",
			Description:    "",
			Type:           "Linux",
			ChangeVersion:  "1569965707",
			LastUpdateTime: types.NewTime(time.Now()),
		},
		Spec: types.CustomizationSpec{
			Options: &types.CustomizationLinuxOptions{},
			Identity: &types.CustomizationLinuxPrep{
				CustomizationIdentitySettings: types.CustomizationIdentitySettings{},
				HostName:                      &types.CustomizationVirtualMachineName{},
				Domain:                        "eng.vmware.com",
				TimeZone:                      "Pacific/Apia",
				HwClockUTC:                    types.NewBool(true),
			},
			GlobalIPSettings: types.CustomizationGlobalIPSettings{
				DnsSuffixList: nil,
				DnsServerList: []string{"127.0.1.1"},
			},
			NicSettingMap: []types.CustomizationAdapterMapping{
				{
					MacAddress: "",
					Adapter: types.CustomizationIPSettings{
						Ip:            &types.CustomizationDhcpIpGenerator{},
						SubnetMask:    "",
						Gateway:       nil,
						IpV6Spec:      (*types.CustomizationIPSettingsIpV6AddressSpec)(nil),
						DnsServerList: nil,
						DnsDomain:     "",
						PrimaryWINS:   "",
						SecondaryWINS: "",
						NetBIOS:       "",
					},
				},
			},
			EncryptionKey: nil,
		},
	},
	{
		Info: types.CustomizationSpecInfo{
			Name:           "vcsim-linux-static",
			Description:    "",
			Type:           "Linux",
			ChangeVersion:  "1569969598",
			LastUpdateTime: types.NewTime(time.Now()),
		},
		Spec: types.CustomizationSpec{
			Options: &types.CustomizationLinuxOptions{},
			Identity: &types.CustomizationLinuxPrep{
				CustomizationIdentitySettings: types.CustomizationIdentitySettings{},
				HostName: &types.CustomizationPrefixName{
					CustomizationName: types.CustomizationName{},
					Base:              "vcsim",
				},
				Domain:     "eng.vmware.com",
				TimeZone:   "Africa/Cairo",
				HwClockUTC: types.NewBool(true),
			},
			GlobalIPSettings: types.CustomizationGlobalIPSettings{
				DnsSuffixList: nil,
				DnsServerList: []string{"127.0.1.1"},
			},
			NicSettingMap: []types.CustomizationAdapterMapping{
				{
					MacAddress: "",
					Adapter: types.CustomizationIPSettings{
						Ip:            &types.CustomizationUnknownIpGenerator{},
						SubnetMask:    "255.255.255.0",
						Gateway:       []string{"10.0.0.1"},
						IpV6Spec:      (*types.CustomizationIPSettingsIpV6AddressSpec)(nil),
						DnsServerList: nil,
						DnsDomain:     "",
						PrimaryWINS:   "",
						SecondaryWINS: "",
						NetBIOS:       "",
					},
				},
			},
			EncryptionKey: nil,
		},
	},
	{
		Info: types.CustomizationSpecInfo{
			Name:           "vcsim-windows-static",
			Description:    "",
			Type:           "Windows",
			ChangeVersion:  "1569978029",
			LastUpdateTime: types.NewTime(time.Now()),
		},
		Spec: types.CustomizationSpec{
			Options: &types.CustomizationWinOptions{
				CustomizationOptions: types.CustomizationOptions{},
				ChangeSID:            true,
				DeleteAccounts:       false,
				Reboot:               "",
			},
			Identity: &types.CustomizationSysprep{
				CustomizationIdentitySettings: types.CustomizationIdentitySettings{},
				GuiUnattended: types.CustomizationGuiUnattended{
					Password:       (*types.CustomizationPassword)(nil),
					TimeZone:       2,
					AutoLogon:      false,
					AutoLogonCount: 1,
				},
				UserData: types.CustomizationUserData{
					FullName:     "vcsim",
					OrgName:      "VMware",
					ComputerName: &types.CustomizationVirtualMachineName{},
					ProductId:    "",
				},
				GuiRunOnce: (*types.CustomizationGuiRunOnce)(nil),
				Identification: types.CustomizationIdentification{
					JoinWorkgroup:       "WORKGROUP",
					JoinDomain:          "",
					DomainAdmin:         "",
					DomainAdminPassword: (*types.CustomizationPassword)(nil),
				},
				LicenseFilePrintData: &types.CustomizationLicenseFilePrintData{
					AutoMode:  "perServer",
					AutoUsers: 5,
				},
			},
			GlobalIPSettings: types.CustomizationGlobalIPSettings{},
			NicSettingMap: []types.CustomizationAdapterMapping{
				{
					MacAddress: "",
					Adapter: types.CustomizationIPSettings{
						Ip:            &types.CustomizationUnknownIpGenerator{},
						SubnetMask:    "255.255.255.0",
						Gateway:       []string{"10.0.0.1"},
						IpV6Spec:      (*types.CustomizationIPSettingsIpV6AddressSpec)(nil),
						DnsServerList: nil,
						DnsDomain:     "",
						PrimaryWINS:   "",
						SecondaryWINS: "",
						NetBIOS:       "",
					},
				},
			},
			EncryptionKey: nil,
		},
	},
	{
		Info: types.CustomizationSpecInfo{
			Name:           "vcsim-windows-domain",
			Description:    "",
			Type:           "Windows",
			ChangeVersion:  "1569970234",
			LastUpdateTime: types.NewTime(time.Now()),
		},
		Spec: types.CustomizationSpec{
			Options: &types.CustomizationWinOptions{
				CustomizationOptions: types.CustomizationOptions{},
				ChangeSID:            true,
				DeleteAccounts:       false,
				Reboot:               "",
			},
			Identity: &types.CustomizationSysprep{
				CustomizationIdentitySettings: types.CustomizationIdentitySettings{},
				GuiUnattended: types.CustomizationGuiUnattended{
					Password: &types.CustomizationPassword{
						Value:     "3Gs...==",
						PlainText: false,
					},
					TimeZone:       15,
					AutoLogon:      false,
					AutoLogonCount: 1,
				},
				UserData: types.CustomizationUserData{
					FullName:     "dougm",
					OrgName:      "VMware",
					ComputerName: &types.CustomizationVirtualMachineName{},
					ProductId:    "",
				},
				GuiRunOnce: (*types.CustomizationGuiRunOnce)(nil),
				Identification: types.CustomizationIdentification{
					JoinWorkgroup: "",
					JoinDomain:    "DOMAIN",
					DomainAdmin:   "vcsim",
					DomainAdminPassword: &types.CustomizationPassword{
						Value:     "H3g...==",
						PlainText: false,
					},
				},
				LicenseFilePrintData: &types.CustomizationLicenseFilePrintData{
					AutoMode:  "perServer",
					AutoUsers: 5,
				},
			},
			GlobalIPSettings: types.CustomizationGlobalIPSettings{},
			NicSettingMap: []types.CustomizationAdapterMapping{
				{
					MacAddress: "",
					Adapter: types.CustomizationIPSettings{
						Ip:            &types.CustomizationUnknownIpGenerator{},
						SubnetMask:    "255.255.255.0",
						Gateway:       []string{"10.0.0.1"},
						IpV6Spec:      (*types.CustomizationIPSettingsIpV6AddressSpec)(nil),
						DnsServerList: nil,
						DnsDomain:     "",
						PrimaryWINS:   "",
						SecondaryWINS: "",
						NetBIOS:       "",
					},
				},
			},
			EncryptionKey: nil,
		},
	},
}

type CustomizationSpecManager struct {
	mo.CustomizationSpecManager

	items []types.CustomizationSpecItem
}

func (m *CustomizationSpecManager) init(r *Registry) {
	m.items = DefaultCustomizationSpec

	// Real VC is different DN, X509v3 extensions, etc.
	// This is still useful for testing []byte of DER encoded cert over SOAP
	if len(m.EncryptionKey) == 0 {
		block, _ := pem.Decode(internal.LocalhostCert)
		m.EncryptionKey = block.Bytes
	}
}

var customizeNameCounter uint64

func customizeName(vm *VirtualMachine, base types.BaseCustomizationName) string {
	n := atomic.AddUint64(&customizeNameCounter, 1)

	switch name := base.(type) {
	case *types.CustomizationPrefixName:
		return fmt.Sprintf("%s-%d", name.Base, n)
	case *types.CustomizationCustomName:
		return fmt.Sprintf("%s-%d", name.Argument, n)
	case *types.CustomizationFixedName:
		return name.Name
	case *types.CustomizationUnknownName:
		return ""
	case *types.CustomizationVirtualMachineName:
		return fmt.Sprintf("%s-%d", vm.Name, n)
	default:
		return ""
	}
}

func (m *CustomizationSpecManager) DoesCustomizationSpecExist(ctx *Context, req *types.DoesCustomizationSpecExist) soap.HasFault {
	exists := false

	for _, item := range m.items {
		if item.Info.Name == req.Name {
			exists = true
			break
		}
	}

	return &methods.DoesCustomizationSpecExistBody{
		Res: &types.DoesCustomizationSpecExistResponse{
			Returnval: exists,
		},
	}
}

func (m *CustomizationSpecManager) GetCustomizationSpec(ctx *Context, req *types.GetCustomizationSpec) soap.HasFault {
	body := new(methods.GetCustomizationSpecBody)

	for _, item := range m.items {
		if item.Info.Name == req.Name {
			body.Res = &types.GetCustomizationSpecResponse{
				Returnval: item,
			}
			return body
		}
	}

	body.Fault_ = Fault("", new(types.NotFound))

	return body
}

func (m *CustomizationSpecManager) CreateCustomizationSpec(ctx *Context, req *types.CreateCustomizationSpec) soap.HasFault {
	body := new(methods.CreateCustomizationSpecBody)

	for _, item := range m.items {
		if item.Info.Name == req.Item.Info.Name {
			body.Fault_ = Fault("", &types.AlreadyExists{Name: req.Item.Info.Name})
			return body
		}
	}

	m.items = append(m.items, req.Item)
	body.Res = new(types.CreateCustomizationSpecResponse)

	return body
}

func (m *CustomizationSpecManager) OverwriteCustomizationSpec(ctx *Context, req *types.OverwriteCustomizationSpec) soap.HasFault {
	body := new(methods.OverwriteCustomizationSpecBody)

	for i, item := range m.items {
		if item.Info.Name == req.Item.Info.Name {
			m.items[i] = req.Item
			body.Res = new(types.OverwriteCustomizationSpecResponse)
			return body
		}
	}

	body.Fault_ = Fault("", new(types.NotFound))

	return body
}

func (m *CustomizationSpecManager) Get() mo.Reference {
	clone := *m

	for i := range clone.items {
		clone.Info = append(clone.Info, clone.items[i].Info)
	}

	return &clone
}
