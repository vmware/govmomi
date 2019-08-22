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

*/

package vm

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"

	"reflect"
)

type custwin struct {
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
	specname string
	// AutoLogon & Count
	al  bool
	alc int

	// Change the domain
	domain string

	// computerName PrefixNameGenerator
	png string

	// TimeZone
	tz int

	custip     string
	custmask   string
	custgw     string
	custdns1   string
	custdns2   string
	waitForIP  bool
	annotation string
	snapshot   string
	link       bool

	Client         *vim25.Client
	Datacenter     *object.Datacenter
	Datastore      *object.Datastore
	StoragePod     *object.StoragePod
	ResourcePool   *object.ResourcePool
	HostSystem     *object.HostSystem
	Folder         *object.Folder
	VirtualMachine *object.VirtualMachine
}

func init() {
	cli.Register("vm.custwin", &custwin{})
}

func (cmd *custwin) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.StoragePodFlag, ctx = flags.NewStoragePodFlag(ctx)
	cmd.StoragePodFlag.Register(ctx, f)

	cmd.ResourcePoolFlag, ctx = flags.NewResourcePoolFlag(ctx)
	cmd.ResourcePoolFlag.Register(ctx, f)

	cmd.HostSystemFlag, ctx = flags.NewHostSystemFlag(ctx)
	cmd.HostSystemFlag.Register(ctx, f)

	cmd.NetworkFlag, ctx = flags.NewNetworkFlag(ctx)
	cmd.NetworkFlag.Register(ctx, f)

	cmd.FolderFlag, ctx = flags.NewFolderFlag(ctx)
	cmd.FolderFlag.Register(ctx, f)

	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)

	f.IntVar(&cmd.memory, "m", 0, "Size in MB of memory")
	f.IntVar(&cmd.cpus, "c", 0, "Number of CPUs")
	f.BoolVar(&cmd.on, "on", true, "Power on VM")
	f.BoolVar(&cmd.force, "force", false, "Create VM if vmx already exists")
	f.BoolVar(&cmd.template, "template", false, "Create a Template")
	f.StringVar(&cmd.specname, "specname", "", "Customization Specification Name")
	f.StringVar(&cmd.custip, "custip", "", "Customization IPAddress")

	// AutoLogon
	f.BoolVar(&cmd.al, "al", false, "Set AutoLogon Enable")
	f.IntVar(&cmd.alc, "alc", -1, "AutoLogonCount - Number of autoLogons")

	// Change domain Name
	f.StringVar(&cmd.domain, "domain", "", "Change Domain name - Uses current user/pass in spec")

	// use the PrefixNameGenerator - specifiy prefix
	f.StringVar(&cmd.png, "png", "", "Name Generator -> Prefix ")

	// Timezone
	f.IntVar(&cmd.tz, "tz", -1, "TimeZone - DECIMAL: https://https://support.microsoft.com/en-us/help/973627/microsoft-time-zone-index-values")

	f.StringVar(&cmd.custgw, "custgw", "", "Customization Gateway")
	f.StringVar(&cmd.custdns1, "custdns1", "", "Customization DNS 1")
	f.StringVar(&cmd.custdns2, "custdns2", "", "Customization DNS 2")
	f.StringVar(&cmd.custmask, "custmask", "", "Customization Netmask")
	f.BoolVar(&cmd.waitForIP, "waitip", false, "Wait for VM to acquire IP address")
	f.StringVar(&cmd.annotation, "annotation", "", "VM description")
	f.StringVar(&cmd.snapshot, "snapshot", "", "Snapshot name to clone from")
	f.BoolVar(&cmd.link, "link", false, "Creates a linked clone from snapshot or source VM")
}

func (cmd *custwin) Usage() string {
	return "NAME"
}

func (cmd *custwin) Description() string {
	return `Clone Windows Template to VM and Customize.

Examples:
  Clone vm and specifiy CustomSpec and overrides, includes normal clone functionality (vm.clone) 
	govc vm.custwin -vm TEMPLATENAME -specname SPECNAME -vm NEWVM
	govc vm.custwin -vm TEMPLATENAME -specname SPECNAME -custgw GATEWAY -custmask NETMASK -custdns1 DNS1 -custdns2 DNS2 NEWVM
		*Note: Panic if CustomSpec is setup for DHCP
	govc vm.custwin -vm TEMPLATENAME -specname SPECNAME -al -alc 3 NEWVM
	govc vm.custwin -vm TEMPLATENAME -specname SPECNAME -png demo NEWVM
		*Note: Above overrides the Name Generator Prefix but does not rename the VC vm
	govc vm.custwin -vm TEMPLATENAME -specname SPECNAME -tz TIMEZONE 
		*Note: TimeZone is DECIMAL via: https://support.microsoft.com/en-us/help/973627/microsoft-time-zone-index-values
`
}

func (cmd *custwin) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.StoragePodFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ResourcePoolFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.HostSystemFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.NetworkFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.FolderFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}

	return nil
}

func (cmd *custwin) Run(ctx context.Context, f *flag.FlagSet) error {
	var err error

	if len(f.Args()) != 1 {
		return flag.ErrHelp
	}

	cmd.name = f.Arg(0)
	if cmd.name == "" {
		return flag.ErrHelp
	}

	cmd.Client, err = cmd.ClientFlag.Client()
	if err != nil {
		return err
	}

	cmd.Datacenter, err = cmd.DatacenterFlag.Datacenter()
	if err != nil {
		return err
	}

	if cmd.StoragePodFlag.Isset() {
		cmd.StoragePod, err = cmd.StoragePodFlag.StoragePod()
		if err != nil {
			return err
		}
	} else {
		cmd.Datastore, err = cmd.DatastoreFlag.Datastore()
		if err != nil {
			return err
		}
	}

	cmd.HostSystem, err = cmd.HostSystemFlag.HostSystemIfSpecified()
	if err != nil {
		return err
	}

	if cmd.HostSystem != nil {
		if cmd.ResourcePool, err = cmd.HostSystem.ResourcePool(ctx); err != nil {
			return err
		}
	} else {
		// -host is optional
		if cmd.ResourcePool, err = cmd.ResourcePoolFlag.ResourcePool(); err != nil {
			return err
		}
	}

	if cmd.Folder, err = cmd.FolderFlag.Folder(); err != nil {
		return err
	}

	if cmd.VirtualMachine, err = cmd.VirtualMachineFlag.VirtualMachine(); err != nil {
		return err
	}

	if cmd.VirtualMachine == nil {
		return flag.ErrHelp
	}

	vm, err := cmd.cloneVM(ctx)
	if err != nil {
		return err
	}

	if cmd.cpus > 0 || cmd.memory > 0 || cmd.annotation != "" {
		vmConfigSpec := types.VirtualMachineConfigSpec{}
		if cmd.cpus > 0 {
			vmConfigSpec.NumCPUs = int32(cmd.cpus)
		}
		if cmd.memory > 0 {
			vmConfigSpec.MemoryMB = int64(cmd.memory)
		}
		vmConfigSpec.Annotation = cmd.annotation
		task, err := vm.Reconfigure(ctx, vmConfigSpec)
		if err != nil {
			return err
		}
		_, err = task.WaitForResult(ctx, nil)
		if err != nil {
			return err
		}
	}

	if cmd.on {
		task, err := vm.PowerOn(ctx)
		if err != nil {
			return err
		}

		_, err = task.WaitForResult(ctx, nil)
		if err != nil {
			return err
		}

		if cmd.waitForIP {
			_, err = vm.WaitForIP(ctx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (cmd *custwin) cloneVM(ctx context.Context) (*object.VirtualMachine, error) {
	devices, err := cmd.VirtualMachine.Device(ctx)
	if err != nil {
		return nil, err
	}

	// prepare virtual device config spec for network card
	configSpecs := []types.BaseVirtualDeviceConfigSpec{}

	if cmd.NetworkFlag.IsSet() {
		op := types.VirtualDeviceConfigSpecOperationAdd
		card, derr := cmd.NetworkFlag.Device()
		if derr != nil {
			return nil, derr
		}
		// search for the first network card of the source
		for _, device := range devices {
			if _, ok := device.(types.BaseVirtualEthernetCard); ok {
				op = types.VirtualDeviceConfigSpecOperationEdit
				// set new backing info
				cmd.NetworkFlag.Change(device, card)
				card = device
				break
			}
		}

		configSpecs = append(configSpecs, &types.VirtualDeviceConfigSpec{
			Operation: op,
			Device:    card,
		})
	}

	folderref := cmd.Folder.Reference()
	poolref := cmd.ResourcePool.Reference()

	relocateSpec := types.VirtualMachineRelocateSpec{
		DeviceChange: configSpecs,
		Folder:       &folderref,
		Pool:         &poolref,
	}

	if cmd.HostSystem != nil {
		hostref := cmd.HostSystem.Reference()
		relocateSpec.Host = &hostref
	}

	cloneSpec := &types.VirtualMachineCloneSpec{
		PowerOn:  false,
		Template: cmd.template,
	}

	if cmd.snapshot == "" {
		if cmd.link {
			relocateSpec.DiskMoveType = string(types.VirtualMachineRelocateDiskMoveOptionsMoveAllDiskBackingsAndAllowSharing)
		}
	} else {
		if cmd.link {
			relocateSpec.DiskMoveType = string(types.VirtualMachineRelocateDiskMoveOptionsCreateNewChildDiskBacking)
		}

		ref, ferr := cmd.VirtualMachine.FindSnapshot(ctx, cmd.snapshot)
		if ferr != nil {
			return nil, ferr
		}

		cloneSpec.Snapshot = ref
	}

	cloneSpec.Location = relocateSpec

	// clone to storage pod
	datastoreref := types.ManagedObjectReference{}
	if cmd.StoragePod != nil && cmd.Datastore == nil {
		storagePod := cmd.StoragePod.Reference()

		// Build pod selection spec from config spec
		podSelectionSpec := types.StorageDrsPodSelectionSpec{
			StoragePod: &storagePod,
		}

		// Get the virtual machine reference
		vmref := cmd.VirtualMachine.Reference()

		// Build the placement spec
		storagePlacementSpec := types.StoragePlacementSpec{
			Folder:           &folderref,
			Vm:               &vmref,
			CloneName:        cmd.name,
			CloneSpec:        cloneSpec,
			PodSelectionSpec: podSelectionSpec,
			Type:             string(types.StoragePlacementSpecPlacementTypeClone),
		}

		// Get the storage placement result
		storageResourceManager := object.NewStorageResourceManager(cmd.Client)
		result, err := storageResourceManager.RecommendDatastores(ctx, storagePlacementSpec)
		if err != nil {
			return nil, err
		}

		// Get the recommendations
		recommendations := result.Recommendations
		if len(recommendations) == 0 {
			return nil, fmt.Errorf("no recommendations")
		}

		// Get the first recommendation
		datastoreref = recommendations[0].Action[0].(*types.StoragePlacementAction).Destination
	} else if cmd.StoragePod == nil && cmd.Datastore != nil {
		datastoreref = cmd.Datastore.Reference()
	} else {
		return nil, fmt.Errorf("please provide either a datastore or a storagepod")
	}

	// Set the destination datastore
	cloneSpec.Location.Datastore = &datastoreref

	// Check if vmx already exists
	if !cmd.force {
		vmxPath := fmt.Sprintf("%s/%s.vmx", cmd.name, cmd.name)

		var mds mo.Datastore
		err = property.DefaultCollector(cmd.Client).RetrieveOne(ctx, datastoreref, []string{"name"}, &mds)
		if err != nil {
			return nil, err
		}

		datastore := object.NewDatastore(cmd.Client, datastoreref)
		datastore.InventoryPath = mds.Name

		_, err := datastore.Stat(ctx, vmxPath)
		if err == nil {
			dsPath := cmd.Datastore.Path(vmxPath)
			return nil, fmt.Errorf("file %s already exists", dsPath)
		}
	}

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
		// fmt.Println("Type:", customInfo.Type)

		if customInfo.Type != "Windows" {
			errMsg := fmt.Sprintf("ERROR: Custom Spec [%s] has Type[%s] specified is not of Type[Windows]", cmd.specname)
			panic(errMsg)
		}

		customSpec := customSpecItem.Spec
		// set the specname
		cloneSpec.Customization = &customSpec

		// Change Domain
		if cmd.domain != "" {
			chkDomain := customSpec.Identity.(*types.CustomizationSysprep).Identification.JoinDomain
			if len(chkDomain) > 0 {
				//fmt.Println("current Domain:", customSpec.Identity.(*types.CustomizationSysprep).Identification.JoinDomain)
				//fmt.Println("set     Domain:", cmd.domain)
				customSpec.Identity.(*types.CustomizationSysprep).Identification.JoinDomain = cmd.domain
			} else {
				errMsg := fmt.Sprintf("ERROR: Spec specified [%s] is not part of a Domain specified as [%s]", cmd.specname, cmd.domain)
				panic(errMsg)
			}
		}

		// fmt.Println("<-PNG:", cmd.png)
		// Change Prefix ... Must have PrefixNameGenerator specified in Template
		if cmd.png != "" {
			//fmt.Println("--->CHECK:", customSpec.Identity.(*types.CustomizationSysprep).UserData.ComputerName)
			if isOfType(customSpec.Identity.(*types.CustomizationSysprep).UserData.ComputerName) {
				if len(cmd.png) > 0 {
					//TODO: If at some point, FullName and OrgName changes here
					//fname := customSpec.Identity.(*types.CustomizationSysprep).UserData.FullName
					//fmt.Printf("FullName: %s\n", fname)
					//orgName := customSpec.Identity.(*types.CustomizationSysprep).UserData.OrgName
					//fmt.Printf("OrgName: %s\n", orgName)

					currentPrefix := customSpec.Identity.(*types.CustomizationSysprep).UserData.ComputerName
					//					prefixName := currentPrefix.(*types.CustomizationPrefixName).Base
					// fmt.Println("current Prefix:", prefixName, ":New Prefix:", cmd.png)
					currentPrefix.(*types.CustomizationPrefixName).Base = cmd.png
					//customSpec.Identity.(*types.CustomizationSysprep).UserData.ComputerName.Base = cmd.png

				} else {
					errMsg := fmt.Sprintf("ERROR: Spec specified [%s] is not part of PrefixNameGenerator specified", cmd.specname)
					panic(errMsg)
				}
			}
		}

		// If al (autoLogon) is True, && AutoLogonCount (alc) != 0, set value
		if cmd.al {
			// fmt.Println("SET spec :autoLogon:", customSpec.Identity.(*types.CustomizationSysprep).GuiUnattended.AutoLogon,
			// 		":param:", cmd.al)
			customSpec.Identity.(*types.CustomizationSysprep).GuiUnattended.AutoLogon = cmd.al
			// If al (autoLogonCount) is not == to spec // set it

			if cmd.alc > 0 {
				//fmt.Println("SET spec :autoLogonCount:", customSpec.Identity.(*types.CustomizationSysprep).GuiUnattended.AutoLogonCount,
				// ":newCount:", cmd.alc)
				customSpec.Identity.(*types.CustomizationSysprep).GuiUnattended.AutoLogonCount = int32(cmd.alc)
			}
		}

		// If TimeZone is different than spec, use thatError
		if cmd.tz != -1 {
			if cmd.tz != int(customSpec.Identity.(*types.CustomizationSysprep).GuiUnattended.TimeZone) {
				//fmt.Println("SET spec :timeZone:", customSpec.Identity.(*types.CustomizationSysprep).GuiUnattended.TimeZone,
				//	":new:", cmd.tz)
				customSpec.Identity.(*types.CustomizationSysprep).GuiUnattended.TimeZone = int32(cmd.tz)
			}
		}

		// If FixedIP, allow the overrides .....
		if isOfType(customSpec.NicSettingMap[0].Adapter.Ip) {

			xx := customSpec.NicSettingMap[0].Adapter.Ip
			if len(cmd.custip) > 0 {
				aa := reflect.ValueOf(xx).Elem()

				// see if Dhcp and reject if so / can't custwin dhcp for static
				// if ipType == "vim.vm.specname.DhcpIpGenerator" {
				// 	return nil, fmt.Errorf("specname specification [%s] has DHCP specified", cmd.specname)
				// }

				bb := aa.FieldByName("IpAddress")
				bb.SetString(cmd.custip)
				//fmt.Println("<-Setup IP:",cmd.custip)
				//fmt.Println("<-Setup DONE IP:",cmd.custip)
			}

			if len(cmd.custmask) > 0 {
				customSpec.NicSettingMap[0].Adapter.SubnetMask = cmd.custmask
			}
			if len(cmd.custgw) > 0 {
				customSpec.NicSettingMap[0].Adapter.Gateway[0] = cmd.custgw
			}
			if len(cmd.custdns1) > 0 {
				customSpec.NicSettingMap[0].Adapter.DnsServerList[0] = cmd.custdns1
			}
			if len(cmd.custdns2) > 0 {
				customSpec.NicSettingMap[0].Adapter.DnsServerList[1] = cmd.custdns2
			}
		} else {
			// DHCP in spec, if trying to override with non dhcp, get out of here
			if len(cmd.custip) > 0 || len(cmd.custmask) > 0 || len(cmd.custgw) > 0 || len(cmd.custdns1) > 0 || len(cmd.custdns2) > 0 {
				panic("Error: DHCP spec detected and trying to use Fixed overrides")
			}
		}
	}

	task, err := cmd.VirtualMachine.Clone(ctx, cmd.Folder, cmd.name, *cloneSpec)
	if err != nil {
		return nil, err
	}

	logger := cmd.ProgressLogger(fmt.Sprintf("Cloning %s to %s...", cmd.VirtualMachine.InventoryPath, cmd.name))
	defer logger.Wait()

	info, err := task.WaitForResult(ctx, logger)
	if err != nil {
		return nil, err
	}

	return object.NewVirtualMachine(cmd.Client, info.Result.(types.ManagedObjectReference)), nil
}

func isOfType(t interface{}) bool {
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

/*
func isOfTypePrefix(t interface{}) bool {
	switch t.(type) {
	case *types.CustomizationPrefixName:
		return true
	default:
		//fmt.Println("->Return Defualt: false:")
		return false
	}
}
*/
