/*
Copyright (c) 2014-2016 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

gressd@vmware.com 08/24/19 Cfreated
# govc vm.customize -vm new-vm -ip x.x.x.x -dns-server x.x.x.x -dns-suffix example.com customization-name

*/

package vm

import (
	"context"
	"flag"
	"fmt"
	"strconv"

	"reflect"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

type customize struct {
	*flags.ClientFlag
	*flags.DatacenterFlag
	*flags.DatastoreFlag
	*flags.StoragePodFlag
	*flags.ResourcePoolFlag
	*flags.HostSystemFlag
	*flags.NetworkFlag
	*flags.FolderFlag
	*flags.VirtualMachineFlag

	name     string
	memory   int
	cpus     int
	on       bool
	force    bool
	template bool
	// AutoLogon & Count
	al  bool
	alc int

	// Change the domain
	domain string

	// computerName PrefixNameGenerator
	png string

	// TimeZone
	tz       string
	specname string

	ip         string
	netmask    string
	gateway    string
	dnsserver1 string
	dnsserver2 string
	waitForIP  bool
	annotation string
	snapshot   string
	link       bool

	Client         *vim25.Client
	VirtualMachine *object.VirtualMachine
}

func init() {
	cli.Register("vm.customize", &customize{})
}

func (cmd *customize) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.StringVar(&cmd.ip, "ip", "", "Customization IPAddress")

	// AutoLogon
	f.BoolVar(&cmd.al, "al", false, "Set AutoLogon Enable")
	f.IntVar(&cmd.alc, "alc", -1, "AutoLogonCount - Number of autoLogons")

	// Change domain Name
	f.StringVar(&cmd.domain, "domain", "", "Change Domain name - Uses current user/pass in spec")

	// use the PrefixNameGenerator - specifiy prefix
	f.StringVar(&cmd.png, "png", "", "Name Generator -> Prefix ")

	// Timezone Linux
	f.StringVar(&cmd.tz, "tz", "", "TimeZone - String: ex:  Windoes: 35: Linux: America/New_York ")

	f.StringVar(&cmd.gateway, "gateway", "", "Customization Gateway")
	f.StringVar(&cmd.dnsserver1, "dns-server1", "", "Customization DNS 1")
	f.StringVar(&cmd.dnsserver2, "dns-server2", "", "Customization DNS 2")
	f.StringVar(&cmd.netmask, "netmask", "", "Customization Netmask")
	/*
		f.BoolVar(&cmd.waitForIP, "waitip", false, "Wait for VM to acquire IP address")
		f.BoolVar(&cmd.on, "on", true, "Power on VM")
	*/
}

func (cmd *customize) Usage() string {
	return "NAME"
}

func (cmd *customize) Description() string {
	return `Clone Windows Template to VM and Customize.

Examples:
  Clone vm and specifiy CustomSpec and overrides, includes normal clone functionality (vm.clone) 
	govc vm.customize -vm NEWVM CUSTOMSPEC-NAME
	govc vm.customize -vm NEWVM -gateway GATEWAY -netmask NETMASK -ip NEWIP -dns-server DNS1 -dns-server2 DNS2 CUSTOMSPEC-NAME
		*Note: Panic if CustomSpec is setup for DHCP
	govc vm.customize -vm NEWVM -al -alc 3 CUSTOMSPEC-NAME
	govc vm.customize -vm NEWVM -png demo CUSTOMSPEC-NAME
		*Note: Above overrides the Name Generator Prefix but does not rename the VC vm
	govc vm.customize -vm NEWVM -tz TIMEZONE 
		*Note: Windows TimeZone is DECIMAL via: https://support.microsoft.com/en-us/help/973627/microsoft-time-zone-index-values
		       Linux   TimeZone is string representation:  America/New_York
`
}

func (cmd *customize) Process(ctx context.Context) error {

	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *customize) Run(ctx context.Context, f *flag.FlagSet) error {
	var err error

	if len(f.Args()) != 1 {
		return flag.ErrHelp
	}

	cmd.specname = f.Arg(0)
	if cmd.specname == "" {
		return flag.ErrHelp
	}

	cmd.Client, err = cmd.ClientFlag.Client()
	if err != nil {
		return err
	}

	if cmd.VirtualMachine, err = cmd.VirtualMachineFlag.VirtualMachine(); err != nil {
		return err
	}

	if cmd.VirtualMachine == nil {
		return flag.ErrHelp
	}

	_, err = cmd.customizeVM(ctx, *cmd.VirtualMachine)
	if err != nil {
		return err
	}

	return nil
}

func (cmd *customize) customizeVM(ctx context.Context, vm object.VirtualMachine) (*object.VirtualMachine, error) {

	// Windows or Linux
	var targetWindows = false
	var err error

	// check if specname specification requested
	if len(cmd.specname) > 0 {
		// get the specname spec manager
		specnameSpecManager := object.NewCustomizationSpecManager(cmd.Client)
		// check if specname specification exists
		exists, err := specnameSpecManager.DoesCustomizationSpecExist(ctx, cmd.specname)
		if err != nil {
			return nil, err
		}
		if !exists {
			errMsg := fmt.Sprintf("specname specification %s does not exists", cmd.specname)
			panic(errMsg)
			//return nil, fmt.Errorf("specname specification %s does not exists", cmd.specname)
		}
		// get the specname specification
		customSpecItem, err := specnameSpecManager.GetCustomizationSpec(ctx, cmd.specname)
		if err != nil {
			return nil, err
		}

		customInfo := customSpecItem.Info

		// Is it a windows or linux spec specified
		if customInfo.Type == "Windows" {
			targetWindows = true
		}

		customSpec := customSpecItem.Spec

		// Change Domain
		if cmd.domain != "" {
			chkDomain := customSpec.Identity.(*types.CustomizationSysprep).Identification.JoinDomain
			if len(chkDomain) > 0 {
				customSpec.Identity.(*types.CustomizationSysprep).Identification.JoinDomain = cmd.domain
			} else {
				errMsg := fmt.Sprintf("ERROR: Spec specified [%s] is not part of a Domain specified as [%s]", cmd.specname, cmd.domain)
				panic(errMsg)
			}
		}

		// Change Prefix ... Must have PrefixNameGenerator specified in Template
		if cmd.png != "" {
			if chckIsOfType(customSpec.Identity.(*types.CustomizationSysprep).UserData.ComputerName) {
				if len(cmd.png) > 0 {
					currentPrefix := customSpec.Identity.(*types.CustomizationSysprep).UserData.ComputerName
					currentPrefix.(*types.CustomizationPrefixName).Base = cmd.png
				} else {
					errMsg := fmt.Sprintf("ERROR: Spec specified [%s] is not part of PrefixNameGenerator specified", cmd.specname)
					panic(errMsg)
				}
			}
		}

		// If al (autoLogon) is True, && AutoLogonCount (alc) != 0, set value
		if cmd.al {
			if !targetWindows {
				errMsg := fmt.Sprintf("Error: You specified AutoLogon, windows option only! SpecName[%s] is of type [%s]", cmd.specname, customInfo.Type)
				panic(errMsg)
			}
			customSpec.Identity.(*types.CustomizationSysprep).GuiUnattended.AutoLogon = cmd.al
			if cmd.alc > 0 {
				customSpec.Identity.(*types.CustomizationSysprep).GuiUnattended.AutoLogonCount = int32(cmd.alc)
			}
		}

		// If TimeZone is different than spec, use thatError
		if targetWindows {
			if cmd.tz != "" {
				// Windows set as int
				wintz, err := strconv.Atoi(cmd.tz)
				if err == nil {
					errMsg := fmt.Sprintf("Error: Windows spec needs decimal Timezone.. specified[%s] err[%s", cmd.tz, err)
					panic(errMsg)
				}
				customSpec.Identity.(*types.CustomizationSysprep).GuiUnattended.TimeZone = int32(wintz)
			}
		} else {
			// Linux set as string
			// vim.vm.customization.LinuxPrep
			if cmd.tz != "" {
				customSpec.Identity.(*types.CustomizationLinuxPrep).TimeZone = cmd.tz
			}
		}

		// If FixedIP, allow the overrides .....
		if chckIsOfType(customSpec.NicSettingMap[0].Adapter.Ip) {

			ptrIp := customSpec.NicSettingMap[0].Adapter.Ip
			if len(cmd.ip) > 0 {
				strEle := reflect.ValueOf(ptrIp).Elem()
				iPaddr := strEle.FieldByName("IpAddress")
				iPaddr.SetString(cmd.ip)
			}

			if len(cmd.netmask) > 0 {
				customSpec.NicSettingMap[0].Adapter.SubnetMask = cmd.netmask
			}
			if len(cmd.gateway) > 0 {
				customSpec.NicSettingMap[0].Adapter.Gateway[0] = cmd.gateway
			}

			// Windows sets dns different than windows
			if targetWindows {
				if len(cmd.dnsserver1) > 0 {
					customSpec.NicSettingMap[0].Adapter.DnsServerList[0] = cmd.dnsserver1
				}
				if len(cmd.dnsserver2) > 0 {
					customSpec.NicSettingMap[0].Adapter.DnsServerList[1] = cmd.dnsserver2
				}
			} else {
				if len(cmd.dnsserver1) > 0 {
					winSpec := customSpec.GlobalIPSettings
					winSpec.DnsServerList[0] = cmd.dnsserver1
				}
				if len(cmd.dnsserver2) > 0 {
					winSpec := customSpec.GlobalIPSettings
					winSpec.DnsServerList[1] = cmd.dnsserver2
				}
				// Tertiary

			}
		} else {
			// DHCP in spec, if trying to override with non dhcp, get out of here
			if len(cmd.ip) > 0 || len(cmd.netmask) > 0 || len(cmd.gateway) > 0 || len(cmd.dnsserver1) > 0 || len(cmd.dnsserver2) > 0 {
				panic("Error: DHCP spec detected and trying to use Fixed overrides")
			}
		}
		/* Issue the Customize */
		task, err := vm.Customize(ctx, customSpec)
		if err != nil {
			return nil, err
		}

		//TaskInfo return
		_, err = task.WaitForResult(ctx, nil)
		if err != nil {
			return nil, err
		}

		/* Include PowerOn?
			TODO:  One item with customize is auto Poweron/off and dhcp ip's being presented
				   should add waitCustIp (Wait for Cusomization IP returned...)

		if cmd.on {
			task, err := vm.PowerOn(ctx)
			if err != nil {
				return nil, err
			}

			_, err = task.WaitForResult(ctx, nil)
			if err != nil {
				return nil, err
			}

			if cmd.waitForIP {
				theIp, err = vm.WaitForIP(ctx)
				if err != nil {
					return nil, err
				}
			}
		}
		*/
	} else {
		errMsg := "Error: A specificationn-name must be specified !"
		panic(errMsg)
	}

	return nil, err
}

// Check to see of the xml section exist of type
func chckIsOfType(t interface{}) bool {
	switch t.(type) {
	case *types.CustomizationFixedIp:
		return true
	case *types.CustomizationPrefixName:
		return true
	default:
		//fmt.Println("->Return Defualt: false:")
		return false
	}
}
