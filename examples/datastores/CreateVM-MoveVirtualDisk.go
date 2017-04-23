/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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

/*
This example program shows how the `finder` and `property` packages can
be used to navigate a vSphere inventory structure using govmomi.
*/

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

const (
	ProviderName              = "vsphere"
	ActivePowerState          = "poweredOn"
	SCSIControllerType        = "scsi"
	LSILogicControllerType    = "lsiLogic"
	BusLogicControllerType    = "busLogic"
	PVSCSIControllerType      = "pvscsi"
	LSILogicSASControllerType = "lsiLogic-sas"
	SCSIControllerLimit       = 4
	SCSIControllerDeviceLimit = 15
	SCSIDeviceSlots           = 16
	SCSIReservedSlot          = 7
	ThinDiskType              = "thin"
	PreallocatedDiskType      = "preallocated"
	EagerZeroedThickDiskType  = "eagerZeroedThick"
	ZeroedThickDiskType       = "zeroedThick"
	VolDir                    = "kubevols"
	RoundTripperDefaultCount  = 3
)

var ErrNoDiskUUIDFound = errors.New("No disk UUID found")
var ErrNoDiskIDFound = errors.New("No vSphere disk ID found")
var ErrNoDevicesFound = errors.New("No devices found")
var ErrNonSupportedControllerType = errors.New("Disk is attached to non-supported controller type")
var ErrFileAlreadyExist = errors.New("File requested already exist")

// GetEnvString returns string from environment variable.
func GetEnvString(v string, def string) string {
	r := os.Getenv(v)
	if r == "" {
		return def
	}

	return r
}

// GetEnvBool returns boolean from environment variable.
func GetEnvBool(v string, def bool) bool {
	r := os.Getenv(v)
	if r == "" {
		return def
	}

	switch strings.ToLower(r[0:1]) {
	case "t", "y", "1":
		return true
	}

	return false
}

const (
	envURL      = "GOVMOMI_URL"
	envUserName = "GOVMOMI_USERNAME"
	envPassword = "GOVMOMI_PASSWORD"
	envInsecure = "GOVMOMI_INSECURE"
)

var urlDescription = fmt.Sprintf("ESX or vCenter URL [%s]", envURL)
var urlFlag = flag.String("url", GetEnvString(envURL, "https://username:password@host/sdk"), urlDescription)

var insecureDescription = fmt.Sprintf("Don't verify the server's certificate chain [%s]", envInsecure)
var insecureFlag = flag.Bool("insecure", GetEnvBool(envInsecure, false), insecureDescription)

func processOverride(u *url.URL) {
	envUsername := os.Getenv(envUserName)
	envPassword := os.Getenv(envPassword)

	// Override username if provided
	if envUsername != "" {
		var password string
		var ok bool

		if u.User != nil {
			password, ok = u.User.Password()
		}

		if ok {
			u.User = url.UserPassword(envUsername, password)
		} else {
			u.User = url.User(envUsername)
		}
	}

	// Override password if provided
	if envPassword != "" {
		var username string

		if u.User != nil {
			username = u.User.Username()
		}

		u.User = url.UserPassword(username, envPassword)
	}
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

func getSCSIController(vmDevices object.VirtualDeviceList, scsiType string) *types.VirtualController {
	// get virtual scsi controller of passed argument type
	for _, device := range vmDevices {
		devType := vmDevices.Type(device)
		if devType == scsiType {
			if c, ok := device.(types.BaseVirtualController); ok {
				return c.GetVirtualController()
			}
		}
	}
	return nil
}

func getSCSIControllers(vmDevices object.VirtualDeviceList) []*types.VirtualController {
	// get all virtual scsi controllers
	var scsiControllers []*types.VirtualController
	for _, device := range vmDevices {
		devType := vmDevices.Type(device)
		switch devType {
		case SCSIControllerType, strings.ToLower(LSILogicControllerType), strings.ToLower(BusLogicControllerType), PVSCSIControllerType, strings.ToLower(LSILogicSASControllerType):
			if c, ok := device.(types.BaseVirtualController); ok {
				scsiControllers = append(scsiControllers, c.GetVirtualController())
			}
		}
	}
	return scsiControllers
}

func getSCSIControllersOfType(vmDevices object.VirtualDeviceList, scsiType string) []*types.VirtualController {
	// get virtual scsi controllers of passed argument type
	var scsiControllers []*types.VirtualController
	for _, device := range vmDevices {
		devType := vmDevices.Type(device)
		if devType == scsiType {
			if c, ok := device.(types.BaseVirtualController); ok {
				scsiControllers = append(scsiControllers, c.GetVirtualController())
			}
		}
	}
	return scsiControllers
}

func getAvailableSCSIController(scsiControllers []*types.VirtualController) *types.VirtualController {
	// get SCSI controller which has space for adding more devices
	for _, controller := range scsiControllers {
		if len(controller.Device) < SCSIControllerDeviceLimit {
			return controller
		}
	}
	return nil
}

// Removes SCSI controller which is latest attached to VM.
func cleanUpController(newSCSIController types.BaseVirtualDevice, vmDevices object.VirtualDeviceList, vm *object.VirtualMachine, ctx context.Context) error {
	if newSCSIController == nil || vmDevices == nil || vm == nil {
		return nil
	}
	ctls := vmDevices.SelectByType(newSCSIController)
	if len(ctls) < 1 {
		return ErrNoDevicesFound
	}
	newScsi := ctls[len(ctls)-1]
	err := vm.RemoveDevice(ctx, true, newScsi)
	if err != nil {
		return err
	}
	return nil
}

func getNextUnitNumber(devices object.VirtualDeviceList, c types.BaseVirtualController) (int32, error) {
	// get next available SCSI controller unit number
	var takenUnitNumbers [SCSIDeviceSlots]bool
	takenUnitNumbers[SCSIReservedSlot] = true
	key := c.GetVirtualController().Key

	for _, device := range devices {
		d := device.GetVirtualDevice()
		if d.ControllerKey == key {
			if d.UnitNumber != nil {
				takenUnitNumbers[*d.UnitNumber] = true
			}
		}
	}
	for unitNumber, takenUnitNumber := range takenUnitNumbers {
		if !takenUnitNumber {
			return int32(unitNumber), nil
		}
	}
	return -1, fmt.Errorf("SCSI Controller with key=%d does not have any available slots (LUN).", key)
}

func formatVirtualDiskUUID(uuid string) string {
	uuidwithNoSpace := strings.Replace(uuid, " ", "", -1)
	uuidWithNoHypens := strings.Replace(uuidwithNoSpace, "-", "", -1)
	return strings.ToLower(uuidWithNoHypens)
}

// Returns formatted UUID for a virtual disk device.
func getVirtualDiskUUID(newDevice types.BaseVirtualDevice) (string, error) {
	vd := newDevice.GetVirtualDevice()

	if b, ok := vd.Backing.(*types.VirtualDiskFlatVer2BackingInfo); ok {
		uuid := formatVirtualDiskUUID(b.Uuid)
		return uuid, nil
	}
	return "", ErrNoDiskUUIDFound
}

// Get VM disk path.
func getVMDiskPath(c *govmomi.Client, virtualMachine *object.VirtualMachine, diskName string) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pc := property.DefaultCollector(c.Client)

	// Convert virtualMachines into list of references
	var vmRefs []types.ManagedObjectReference
	vmRefs = append(vmRefs, virtualMachine.Reference())

	// Retrieve layoutEx.file property for the given datastore
	var vmMorefs []mo.VirtualMachine
	err := pc.Retrieve(ctx, vmRefs, []string{"layoutEx.file"}, &vmMorefs)
	if err != nil {
		return "", err
	}

	for _, vm := range vmMorefs {
		fileLayoutInfo := vm.LayoutEx.File
		for _, fileInfo := range fileLayoutInfo {
			fmt.Println(fileInfo.)
			// Search for the diskName in the VM file layout
			if strings.HasSuffix(fileInfo.Name, diskName) {
				return fileInfo.Name, nil
			}
		}
	}
	return "", nil
}

// Returns vSphere objects virtual machine, virtual device list, datastore and datacenter.
func getVirtualMachineDevices(ctx context.Context, c *govmomi.Client, name string) (*object.VirtualMachine, object.VirtualDeviceList, *object.Datacenter, error) {
	// Create a new finder
	f := find.NewFinder(c.Client, true)

	// Fetch and set data center
	dc, err := f.DefaultDatacenter(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	f.SetDatacenter(dc)

	vmRegex := name

	vm, err := f.VirtualMachine(ctx, vmRegex)
	if err != nil {
		return nil, nil, nil, err
	}

	// Get devices from VM
	vmDevices, err := vm.Device(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	return vm, vmDevices, dc, nil
}

func getVirtualDiskUUIDByPath(volPath string, dc *object.Datacenter, client *govmomi.Client) (string, error) {
	if len(volPath) > 0 && filepath.Ext(volPath) != ".vmdk" {
		volPath += ".vmdk"
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// VirtualDiskManager provides a way to manage and manipulate virtual disks on vmware datastores.
	vdm := object.NewVirtualDiskManager(client.Client)
	// Returns uuid of vmdk virtual disk
	diskUUID, err := vdm.QueryVirtualDiskUuid(ctx, volPath, dc)

	if err != nil {
		return "", ErrNoDiskUUIDFound
	}

	diskUUID = formatVirtualDiskUUID(diskUUID)

	return diskUUID, nil
}

// Returns a device id which is internal vSphere API identifier for the attached virtual disk.
func getVirtualDiskID(volPath string, vmDevices object.VirtualDeviceList, dc *object.Datacenter, client *govmomi.Client) (string, error) {
	volumeUUID, err := getVirtualDiskUUIDByPath(volPath, dc, client)

	if err != nil {
		return "", err
	}

	// filter vm devices to retrieve disk ID for the given vmdk file
	for _, device := range vmDevices {
		if vmDevices.TypeName(device) == "VirtualDisk" {
			diskUUID, _ := getVirtualDiskUUID(device)
			if diskUUID == volumeUUID {
				return vmDevices.Name(device), nil
			}
		}
	}
	return "", ErrNoDiskIDFound
}

// DetachDisk detaches given virtual disk volume from the compute running kubelet.
func detachDisk(ctx context.Context, c *govmomi.Client, volPath string, nodeName string) error {
	vm, vmDevices, dc, err := getVirtualMachineDevices(ctx, c, nodeName)
	fmt.Println(vm)

	if err != nil {
		return err
	}

	diskID, err := getVirtualDiskID(volPath, vmDevices, dc, c)
	if err != nil {
		return err
	}

	// Gets virtual disk device
	device := vmDevices.Find(diskID)
	fmt.Println(device)

	if device == nil {
		return fmt.Errorf("device '%s' not found", diskID)
	}

	// Detach disk from VM
	err = vm.RemoveDevice(ctx, true, device)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()

	// Parse URL from string
	u, err := url.Parse(*urlFlag)
	if err != nil {
		exit(err)
	}

	// Override username and/or password as required
	processOverride(u)

	// Connect and log in to ESX or vCenter
	c, err := govmomi.NewClient(ctx, u, *insecureFlag)
	if err != nil {
		exit(err)
	}

	f := find.NewFinder(c.Client, true)

	// Find one and only datacenter
	dc, err := f.DefaultDatacenter(ctx)
	if err != nil {
		exit(err)
	}

	// Make future calls local to this datacenter
	f.SetDatacenter(dc)

	vmDiskPath1 := "[vsanDatastore] kubevols/redis-master.vmdk"

	datastorePathObj := new(object.DatastorePath)
	isSuccess := datastorePathObj.FromString(vmDiskPath1)
	if !isSuccess {
		exit(errors.New("Failed to parse vmDiskPath1"))
	}
	// ds, err := f.Datastore(ctx, datastorePathObj.Datastore)
	// var dsRef types.ManagedObjectReference = ds.Reference
	// new code
	dummyVMName := "MARVEL-checkyou-prashanth157"
	// vmVirtualMachineConfigSpec := types.VirtualMachineConfigSpec{
	// 	Name: dummyVMName,
	// 	Files: &types.VirtualMachineFileInfo{
	// 		VmPathName: "[" + datastorePathObj.Datastore + "]",
	// 	},
	// 	NumCPUs:  1,
	// 	MemoryMB: 2048,
	// }

	// scsiDeviceConfigSpec := &types.VirtualDeviceConfigSpec{
	// 	Operation: types.VirtualDeviceConfigSpecOperationAdd,
	// 	Device: &types.VirtualLsiLogicController{
	// 		types.VirtualSCSIController{
	// 			SharedBus: types.VirtualSCSISharingNoSharing,
	// 			VirtualController: types.VirtualController{
	// 				BusNumber: 0,
	// 				VirtualDevice: types.VirtualDevice{
	// 					Key: 1000,
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// diskConfigSpec := &types.VirtualDeviceConfigSpec{
	// 	Operation:     types.VirtualDeviceConfigSpecOperationAdd,
	// 	FileOperation: types.VirtualDeviceConfigSpecFileOperationCreate,
	// 	Device: &types.VirtualDisk{
	// 		VirtualDevice: types.VirtualDevice{
	// 			Key:           0,
	// 			ControllerKey: 1000,
	// 			UnitNumber:    new(int32), // zero default value
	// 			Backing: &types.VirtualDiskFlatVer2BackingInfo{
	// 				DiskMode:        string(types.VirtualDiskModePersistent),
	// 				ThinProvisioned: types.NewBool(true),
	// 				VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
	// 					FileName: "[" + datastorePathObj.Datastore + "] " + "youandMe.vmdk",
	// 				},
	// 			},
	// 		},
	// 		CapacityInKB: 4000000,
	// 	},
	// }

	// storagePolicySpec := &types.VirtualMachineDefinedProfileSpec{
	// 	ProfileId: "",
	// 	ProfileData: &types.VirtualMachineProfileRawData{
	// 		ExtensionKey: "com.vmware.vim.sps",
	// 		ObjectData:   "((\"stripeWidth\" i1))",
	// 	},
	// }

	// diskConfigSpec.Profile = append(diskConfigSpec.Profile, storagePolicySpec)
	// vmVirtualMachineConfigSpec.DeviceChange = append(vmVirtualMachineConfigSpec.DeviceChange, scsiDeviceConfigSpec)
	// vmVirtualMachineConfigSpec.DeviceChange = append(vmVirtualMachineConfigSpec.DeviceChange, diskConfigSpec)
	// dcFolders, err := dc.Folders(ctx)
	// resourcepool, err := f.ResourcePool(ctx, "*"+"cluster-vsan-1"+"/Resources")
	// if err != nil {
	// 	exit(err)
	// }

	// var refs []types.ManagedObjectReference
	// vmfolders, _ := dcFolders.VmFolder.Children(ctx)
	// for _, kk := range vmfolders {
	// 	refs = append(refs, kk.Reference())
	// }

	// // var ll []mo.ComputeResource
	// // pc := property.DefaultCollector(c.Client)
	// // err = pc.Retrieve(ctx, refs, []string{"name"}, &ll)
	// // if err != nil {
	// // 	exit(err)
	// // }

	// var folderRefs []types.ManagedObjectReference
	// for _, vm := range refs {
	// 	// fmt.Println(vm.Value + " " + vm.Type)
	// 	if vm.Type == "Folder" {
	// 		folderRefs = append(folderRefs, vm)
	// 	}
	// }

	// var myFolder *object.Folder
	// pc := property.DefaultCollector(c.Client)
	// for _, ref := range folderRefs {
	// 	var refs []types.ManagedObjectReference
	// 	var folderMorefs []mo.Folder
	// 	refs = append(refs, ref)
	// 	err = pc.Retrieve(ctx, refs, []string{"name"}, &folderMorefs)
	// 	for _, mref := range folderMorefs {
	// 		if mref.Name == "k8s-mine" {
	// 			myFolder = object.NewFolder(c.Client, ref)
	// 		}
	// 	}
	// }

	// task, err := myFolder.CreateVM(ctx, vmVirtualMachineConfigSpec, resourcepool, nil)
	// if err != nil {
	// 	fmt.Println("vm create failure")
	// 	exit(err)
	// }

	// err = task.Wait(ctx)
	// if err != nil {
	// 	fmt.Println("vm create wait failure")
	// 	exit(err)
	// }
	// // // end of new code

	vmRegex := "k8s-mine/" + dummyVMName

	dummyVM, err := f.VirtualMachine(ctx, vmRegex)
	if err != nil {
		fmt.Println("vm failure")
		exit(err)
	}

	vmDiskPath, err := getVMDiskPath(c, dummyVM, "youandMe.vmdk")
	if err != nil {
		fmt.Println("vmDiskPath failure")
		exit(err)
	}
	if vmDiskPath == "" {
		fmt.Println("vmDiskPath empty")
		exit(err)
	}

	fmt.Println(vmDiskPath)
	vmDiskPath2, err := getVMDiskPath(c, dummyVM, dummyVMName+".vmdk")
	if err != nil {
		exit(err)
	}
	if vmDiskPath2 == "" {
		exit(err)
	}

	fmt.Println("vmDiskPath2 is :" + vmDiskPath2)

	err = detachDisk(ctx, c, vmDiskPath, vmRegex)
	if err != nil {
		exit(err)
	}

	err = detachDisk(ctx, c, vmDiskPath2, vmRegex)
	if err != nil {
		exit(err)
	}

	destroyTask, err := dummyVM.Destroy(ctx)
	if err != nil {
		exit(err)
	}

	err = destroyTask.Wait(ctx)
	if err != nil {
		exit(err)
	}
	// --
	virtualDiskManager := object.NewVirtualDiskManager(c.Client)
	destPath := "[" + datastorePathObj.Datastore + "] kubevols/youandMe.vmdk"
	task1, err := virtualDiskManager.MoveVirtualDisk(ctx, vmDiskPath, dc, destPath, dc, true)
	if err != nil {
		exit(err)
	}

	err = task1.Wait(ctx)
	if err != nil {
		exit(err)
	}

	// i := strings.Index(vmDiskPath, "/")
	// folderPath := ""
	// if i > -1 {
	// 	folderPath = vmDiskPath[:i+1]
	// }
	// fmt.Println("folder path is: " + folderPath)
	// fileManager := object.NewFileManager(c.Client)
	// task, err = fileManager.DeleteDatastoreFile(ctx, folderPath, dc)
	// if err != nil {
	// 	exit(err)
	// }

	// err = task.Wait(ctx)
	// if err != nil {
	// 	exit(err)
	// }
	//--
	// fmt.Println("\n" + vmDiskPath + "\n")

	// err = detachDisk(ctx, c, vmDiskPath, dummyVMName)
	// if err != nil {
	// 	exit(err)
	// }

	// // Delete the dummy VM
	// destroyTask, err := dummyVM.Destroy(ctx)
	// if err != nil {
	// 	exit(err)
	// }

	// err = destroyTask.Wait(ctx)
	// if err != nil {
	// 	exit(err)
	// }

	// virtualDiskManager := object.NewVirtualDiskManager(c.Client)
	// task, err = virtualDiskManager.DeleteVirtualDisk(ctx, vmDiskPath, dc)
	// if err != nil {
	// 	exit(err)
	// }

	// err = task.Wait(ctx)
	// if err != nil {
	// 	exit(err)
	// }

	// i := strings.Index(vmDiskPath, "/")
	// var chars string
	// if i > -1 {
	// 	chars = vmDiskPath[:i+1]
	// } else {
	// 	fmt.Println("Index not found")
	// 	fmt.Println(vmDiskPath)
	// 	exit(err)
	// }
	// fileManager := object.NewFileManager(c.Client)
	// task, err = fileManager.DeleteDatastoreFile(ctx, chars, dc)

	// if err != nil {
	// 	exit(err)
	// }

	// err = task.Wait(ctx)
	// if err != nil {
	// 	exit(err)
	// }
	//-------------------
	// vmDevices, err := vm.Device(ctx)
	// if err != nil {
	// 	exit(err)
	// }

	// var diskControllerType = "pvscsi"
	// // find SCSI controller of particular type from VM devices
	// allSCSIControllers := getSCSIControllers(vmDevices)
	// scsiControllersOfRequiredType := getSCSIControllersOfType(vmDevices, diskControllerType)
	// scsiController := getAvailableSCSIController(scsiControllersOfRequiredType)

	// var newSCSICreated = false
	// var newSCSIController types.BaseVirtualDevice

	// // creating a scsi controller as there is none found of controller type defined
	// if scsiController == nil {
	// 	if len(allSCSIControllers) >= 4 {
	// 		fmt.Fprintf(os.Stderr, "Error: we reached the maximum number of controllers we can attach\n")
	// 		os.Exit(1)
	// 	}
	// 	newSCSIController, err := vmDevices.CreateSCSIController(diskControllerType)
	// 	if err != nil {
	// 		log.Printf("unable to find controller")
	// 		exit(err)
	// 	}
	// 	configNewSCSIController := newSCSIController.(types.BaseVirtualSCSIController).GetVirtualSCSIController()
	// 	hotAndRemove := true
	// 	configNewSCSIController.HotAddRemove = &hotAndRemove
	// 	configNewSCSIController.SharedBus = types.VirtualSCSISharing(types.VirtualSCSISharingNoSharing)

	// 	// add the scsi controller to virtual machine
	// 	err = vm.AddDevice(context.TODO(), newSCSIController)
	// 	if err != nil {
	// 		// attempt clean up of scsi controller
	// 		if vmDevices, err := vm.Device(ctx); err == nil {
	// 			cleanUpController(newSCSIController, vmDevices, vm, ctx)
	// 		}
	// 		exit(err)
	// 	}

	// 	// verify scsi controller in virtual machine
	// 	vmDevices, err = vm.Device(ctx)
	// 	if err != nil {
	// 		// cannot cleanup if there is no device list
	// 		exit(err)
	// 	}

	// 	scsiController = getSCSIController(vmDevices, "pvscsi")
	// 	if scsiController == nil {
	// 		// attempt clean up of scsi controller
	// 		cleanUpController(newSCSIController, vmDevices, vm, ctx)
	// 		fmt.Fprintf(os.Stderr, "Error: no scsi  controller\n")
	// 		os.Exit(1)
	// 	}
	// 	newSCSICreated = true
	// }

	// disk := vmDevices.CreateDisk(scsiController, ds.Reference(), vmDiskPath)
	// unitNumber, err := getNextUnitNumber(vmDevices, scsiController)
	// if err != nil {
	// 	exit(err)
	// }
	// *disk.UnitNumber = unitNumber

	// backing := disk.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
	// backing.DiskMode = string(types.VirtualDiskModeIndependent_persistent)

	// // Reconfigure VM
	// virtualMachineConfigSpec := types.VirtualMachineConfigSpec{}
	// deviceConfigSpec := &types.VirtualDeviceConfigSpec{
	// 	Device:    disk,
	// 	Operation: types.VirtualDeviceConfigSpecOperationAdd,
	// 	// FileOperation: types.VirtualDeviceConfigSpecFileOperationCreate,
	// }

	// storageProfileSpec := &types.VirtualMachineDefinedProfileSpec{
	// 	ProfileId: "8f4d6cae-2c27-4ecb-935e-1a3f4b823386",
	// }

	// deviceConfigSpec.Profile = append(deviceConfigSpec.Profile, storageProfileSpec)
	// virtualMachineConfigSpec.DeviceChange = append(virtualMachineConfigSpec.DeviceChange, deviceConfigSpec)
	// task, err = vm.Reconfigure(ctx, virtualMachineConfigSpec)
	// if err != nil {
	// 	exit(err)
	// }

	// err = task.Wait(ctx)
	// if err != nil {
	// 	log.Printf("hi")
	// 	exit(err)
	// }

	// vmDevices, err = vm.Device(ctx)
	// if err != nil {
	// 	if newSCSICreated {
	// 		cleanUpController(newSCSIController, vmDevices, vm, ctx)
	// 	}
	// 	exit(err)
	// }

	// devices := vmDevices.SelectByType(disk)
	// if len(devices) < 1 {
	// 	if newSCSICreated {
	// 		cleanUpController(newSCSIController, vmDevices, vm, ctx)
	// 	}
	// 	exit(errors.New("No devices found"))
	// }

	// // get new disk id
	// newDevice := devices[len(devices)-1]
	// deviceName := devices.Name(newDevice)

	// // get device uuid
	// diskUUID, err := getVirtualDiskUUID(newDevice)
	// if err != nil {
	// 	if newSCSICreated {
	// 		cleanUpController(newSCSIController, vmDevices, vm, ctx)
	// 	}
	// 	exit(err)
	// }

	// tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
	// fmt.Fprintf(tw, "DeviceName:\tDiskUUID:\n")
	// fmt.Fprintf(tw, "%s\t", deviceName)
	// fmt.Fprintf(tw, "%s\t", diskUUID)
	// fmt.Fprintf(tw, "\n")
	// tw.Flush()
	// ----------------------------
	// err = vm.Unregister(ctx)
	// if err != nil {
	// 	exit(err)
	// }
}
