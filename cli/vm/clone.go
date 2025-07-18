// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vm

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type clone struct {
	*flags.ClientFlag
	*flags.ClusterFlag
	*flags.DatacenterFlag
	*flags.DatastoreFlag
	*flags.StoragePodFlag
	*flags.ResourcePoolFlag
	*flags.HostSystemFlag
	*flags.NetworkFlag
	*flags.FolderFlag
	*flags.VirtualMachineFlag

	name          string
	memory        int
	cpus          int
	on            bool
	force         bool
	template      bool
	customization string
	waitForIP     bool
	annotation    string
	snapshot      string
	link          bool

	Client         *vim25.Client
	Cluster        *object.ClusterComputeResource
	Datacenter     *object.Datacenter
	Datastore      *object.Datastore
	StoragePod     *object.StoragePod
	ResourcePool   *object.ResourcePool
	HostSystem     *object.HostSystem
	Folder         *object.Folder
	VirtualMachine *object.VirtualMachine
}

func init() {
	cli.Register("vm.clone", &clone{})
}

func (cmd *clone) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.RegisterPlacement(ctx, f)

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
	f.StringVar(&cmd.customization, "customization", "", "Customization Specification Name")
	f.BoolVar(&cmd.waitForIP, "waitip", false, "Wait for VM to acquire IP address")
	f.StringVar(&cmd.annotation, "annotation", "", "VM description")
	f.StringVar(&cmd.snapshot, "snapshot", "", "Snapshot name to clone from")
	f.BoolVar(&cmd.link, "link", false, "Creates a linked clone from snapshot or source VM")
}

func (cmd *clone) Usage() string {
	return "NAME"
}

func (cmd *clone) Description() string {
	return `Clone VM or template to NAME.

Examples:
  govc vm.clone -vm template-vm new-vm
  govc vm.clone -vm template-vm -link new-vm
  govc vm.clone -vm template-vm -snapshot s-name new-vm
  govc vm.clone -vm template-vm -link -snapshot s-name new-vm
  govc vm.clone -vm template-vm -cluster cluster1 new-vm # use compute cluster placement
  govc vm.clone -vm template-vm -datastore-cluster dscluster new-vm # use datastore cluster placement
  govc vm.clone -vm template-vm -snapshot $(govc snapshot.tree -vm template-vm -C) new-vm
  govc vm.clone -vm template-vm -template new-template # clone a VM template
  govc vm.clone -vm=/ClusterName/vm/FolderName/VM_templateName -on=true -host=myesxi01 -ds=datastore01 myVM_name`
}

func (cmd *clone) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.ClusterFlag.Process(ctx); err != nil {
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

func (cmd *clone) Run(ctx context.Context, f *flag.FlagSet) error {
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

	cmd.Cluster, err = cmd.ClusterFlag.ClusterIfSpecified()
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
	} else if cmd.Cluster == nil {
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
		if cmd.Cluster == nil {
			// -host is optional
			if cmd.ResourcePool, err = cmd.ResourcePoolFlag.ResourcePool(); err != nil {
				return err
			}
		} else {
			if cmd.ResourcePool, err = cmd.Cluster.ResourcePool(ctx); err != nil {
				return err
			}
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
	if cmd.Spec {
		return nil
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

	if cmd.template {
		return nil
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

func (cmd *clone) cloneVM(ctx context.Context) (*object.VirtualMachine, error) {
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
	var poolref *types.ManagedObjectReference
	if cmd.ResourcePool != nil {
		poolref = types.NewReference(cmd.ResourcePool.Reference())
	}

	relocateSpec := types.VirtualMachineRelocateSpec{
		DeviceChange: configSpecs,
		Folder:       &folderref,
		Pool:         poolref,
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
	vmref := cmd.VirtualMachine.Reference()

	// clone to storage pod
	datastoreref := types.ManagedObjectReference{}
	if cmd.StoragePod != nil && cmd.Datastore == nil {
		storagePod := cmd.StoragePod.Reference()

		// Build pod selection spec from config spec
		podSelectionSpec := types.StorageDrsPodSelectionSpec{
			StoragePod: &storagePod,
		}

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
			return nil, fmt.Errorf("no datastore-cluster recommendations")
		}

		// Get the first recommendation
		datastoreref = recommendations[0].Action[0].(*types.StoragePlacementAction).Destination
	} else if cmd.StoragePod == nil && cmd.Datastore != nil {
		datastoreref = cmd.Datastore.Reference()
	} else if cmd.Cluster != nil {
		spec := types.PlacementSpec{
			PlacementType: string(types.PlacementSpecPlacementTypeClone),
			CloneName:     cmd.name,
			CloneSpec:     cloneSpec,
			RelocateSpec:  &cloneSpec.Location,
			Vm:            &vmref,
		}
		result, err := cmd.Cluster.PlaceVm(ctx, spec)
		if err != nil {
			return nil, err
		}

		recs := result.Recommendations
		if len(recs) == 0 {
			return nil, fmt.Errorf("no cluster recommendations")
		}

		rspec := *recs[0].Action[0].(*types.PlacementAction).RelocateSpec
		cloneSpec.Location.Host = rspec.Host
		cloneSpec.Location.Datastore = rspec.Datastore
		datastoreref = *rspec.Datastore
	} else {
		return nil, fmt.Errorf("please provide either a cluster, datastore or datastore-cluster")
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
			dsPath := datastore.Path(vmxPath)
			return nil, fmt.Errorf("file %s already exists", dsPath)
		}
	}

	// check if customization specification requested
	if len(cmd.customization) > 0 {
		// get the customization spec manager
		customizationSpecManager := object.NewCustomizationSpecManager(cmd.Client)
		// check if customization specification exists
		exists, err := customizationSpecManager.DoesCustomizationSpecExist(ctx, cmd.customization)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, fmt.Errorf("customization specification %s does not exists", cmd.customization)
		}
		// get the customization specification
		customSpecItem, err := customizationSpecManager.GetCustomizationSpec(ctx, cmd.customization)
		if err != nil {
			return nil, err
		}
		customSpec := customSpecItem.Spec
		// set the customization
		cloneSpec.Customization = &customSpec
	}

	if cmd.Spec {
		return nil, cmd.WriteAny(cloneSpec)
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
