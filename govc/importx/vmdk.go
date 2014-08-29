/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package importx

import (
	"errors"
	"flag"
	"fmt"
	"path"
	"reflect"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/govc/util"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type vmdk struct {
	*flags.DatastoreFlag
	*flags.ResourcePoolFlag
	*flags.OutputFlag

	upload bool
	force  bool
	keep   bool

	Client       *govmomi.Client
	Datacenter   *govmomi.Datacenter
	Datastore    *govmomi.Datastore
	ResourcePool *govmomi.ResourcePool
}

func init() {
	cli.Register("import.vmdk", &vmdk{})
	cli.Alias("import.vmdk", "datastore.import")
}

func (cmd *vmdk) Register(f *flag.FlagSet) {
	f.BoolVar(&cmd.upload, "upload", true, "Upload specified disk")
	f.BoolVar(&cmd.force, "force", false, "Overwrite existing disk")
	f.BoolVar(&cmd.keep, "keep", false, "Keep uploaded disk after import")
}

func (cmd *vmdk) Process() error { return nil }

func (cmd *vmdk) Run(f *flag.FlagSet) error {
	var err error

	args := f.Args()
	if len(args) != 1 {
		return errors.New("no file to import")
	}

	file := importable(f.Arg(0))

	cmd.Client, err = cmd.DatastoreFlag.Client()
	if err != nil {
		return err
	}

	cmd.Datacenter, err = cmd.DatastoreFlag.Datacenter()
	if err != nil {
		return err
	}

	cmd.Datastore, err = cmd.DatastoreFlag.Datastore()
	if err != nil {
		return err
	}

	cmd.ResourcePool, err = cmd.ResourcePoolFlag.ResourcePool()
	if err != nil {
		return err
	}

	if cmd.upload {
		err = cmd.Upload(file)
		if err != nil {
			return err
		}
	}

	return cmd.Import(file)
}

func (cmd *vmdk) Import(i importable) error {
	err := cmd.Copy(i)
	if err != nil {
		return err
	}

	if !cmd.keep {
		err = cmd.Delete(path.Dir(i.RemoteVMDK()))
		if err != nil {
			return err
		}
	}

	return nil
}

func (cmd *vmdk) Upload(i importable) error {
	u, err := cmd.Datastore.URL(cmd.Client, cmd.Datacenter, i.RemoteVMDK())
	if err != nil {
		return err
	}

	p := soap.DefaultUpload
	if cmd.OutputFlag.TTY {
		ch := make(chan vim25.Progress)
		wg := cmd.ProgressLogger("Uploading... ", ch)
		defer wg.Wait()

		p.ProgressCh = ch
	}

	return cmd.Client.Client.UploadFile(string(i), u, &p)
}

func (cmd *vmdk) Copy(i importable) error {
	var err error

	pa := util.NewProgressAggregator(1)
	wg := cmd.ProgressLogger("Importing... ", pa.C)
	switch p := cmd.Client.ServiceContent.About.ApiType; p {
	case "HostAgent":
		err = cmd.CopyHostAgent(i, pa)
	case "VirtualCenter":
		err = cmd.CopyVirtualCenter(i, pa)
	default:
		return fmt.Errorf("unsupported product line: %s", p)
	}

	pa.Done()
	wg.Wait()

	return err
}

type basicProgressWrapper struct {
	detail string
	err    error
}

func (b basicProgressWrapper) Percentage() float32 {
	return 0.0
}

func (b basicProgressWrapper) Detail() string {
	return b.detail
}

func (b basicProgressWrapper) Error() error {
	return b.err
}

// PrepareDestination makes sure that the destination VMDK does not yet exist.
// If the force flag is passed, it removes the existing VMDK. This functions
// exists to give a meaningful error if the remote VMDK already exists.
//
// CopyVirtualDisk can return a "<src> file does not exist" error while in fact
// the source file *does* exist and the *destination* file also exist.
//
func (cmd *vmdk) PrepareDestination(i importable) error {
	b, err := cmd.Datastore.Browser(cmd.Client)
	if err != nil {
		return err
	}

	vmdkPath := i.RemoteDst()
	spec := types.HostDatastoreBrowserSearchSpec{
		Details: &types.FileQueryFlags{
			FileType:  true,
			FileOwner: true, // TODO: omitempty is generated, but seems to be required
		},
		MatchPattern: []string{path.Base(vmdkPath)},
	}

	dsPath := cmd.Datastore.Path(path.Dir(vmdkPath))
	task, err := b.SearchDatastore(cmd.Client, dsPath, &spec)
	if err != nil {
		return err
	}

	// Don't use progress aggregator here; an error may be a good thing.
	info, err := task.WaitForResult(nil)
	if err != nil {
		if info.Error != nil {
			_, ok := info.Error.Fault.(*types.FileNotFound)
			if ok {
				// FileNotFound means the base path doesn't exist. Create it.
				dsPath := cmd.Datastore.Path(path.Dir(vmdkPath))
				return cmd.Client.FileManager().MakeDirectory(dsPath, cmd.Datacenter, true)
			}
		}

		return err
	}

	res := info.Result.(types.HostDatastoreBrowserSearchResults)
	if len(res.File) == 0 {
		// Destination path doesn't exist; all good to continue with import.
		return nil
	}

	// Check that the returned entry has the right type.
	switch res.File[0].(type) {
	case *types.VmDiskFileInfo:
	default:
		expected := "VmDiskFileInfo"
		actual := reflect.TypeOf(res.File[0])
		panic(fmt.Sprintf("Expected: %s, actual: %s", expected, actual))
	}

	if !cmd.force {
		dsPath := cmd.Datastore.Path(vmdkPath)
		err = fmt.Errorf("File %s already exists", dsPath)
		return err
	}

	// Delete existing disk.
	vdm := cmd.Client.VirtualDiskManager()
	task, err = vdm.DeleteVirtualDisk(cmd.Datastore.Path(vmdkPath), cmd.Datacenter)
	if err != nil {
		return err
	}

	return task.Wait()
}

func (cmd *vmdk) CopyHostAgent(i importable, pa *util.ProgressAggregator) error {
	pch := pa.NewChannel("preparing destination")
	pch <- basicProgressWrapper{}
	err := cmd.PrepareDestination(i)
	pch <- basicProgressWrapper{err: err}
	close(pch)
	if err != nil {
		return err
	}

	spec := &types.VirtualDiskSpec{
		AdapterType: "lsiLogic",
		DiskType:    "thin",
	}

	dc := cmd.Datacenter
	src := cmd.Datastore.Path(i.RemoteVMDK())
	dst := cmd.Datastore.Path(i.RemoteDst())
	vdm := cmd.Client.VirtualDiskManager()
	task, err := vdm.CopyVirtualDisk(src, dc, dst, dc, spec, false)
	if err != nil {
		return err
	}

	pch = pa.NewChannel("copying disk")
	_, err = task.WaitForResult(pch)
	if err != nil {
		return err
	}

	return nil
}

func (cmd *vmdk) CopyVirtualCenter(i importable, pa *util.ProgressAggregator) error {
	var err error

	dstName := path.Dir(i.RemoteDst())
	srcName := dstName + "-src"

	spec := &configSpec{
		Name:    srcName,
		GuestId: "otherGuest",
		Files: &types.VirtualMachineFileInfo{
			VmPathName: fmt.Sprintf("[%s]", cmd.Datastore.Name()),
		},
	}

	spec.AddDisk(cmd.Datastore, i.RemoteVMDK())

	src, err := cmd.CreateVM(spec)
	if err != nil {
		return err
	}

	dst, err := cmd.CloneVM(src, dstName)
	if err != nil {
		return err
	}

	err = cmd.DestroyVM(src)
	if err != nil {
		return err
	}

	err = cmd.DestroyVM(dst)
	if err != nil {
		return err
	}

	return nil
}

func (cmd *vmdk) Move(src, dst string) error {
	fm := cmd.Client.FileManager()
	dsSrc := cmd.Datastore.Path(src)
	dsDst := cmd.Datastore.Path(dst)
	task, err := fm.MoveDatastoreFile(dsSrc, cmd.Datacenter, dsDst, cmd.Datacenter, true)
	if err != nil {
		return err
	}

	return task.Wait()
}

func (cmd *vmdk) Delete(path string) error {
	fm := cmd.Client.FileManager()
	dsPath := cmd.Datastore.Path(path)
	task, err := fm.DeleteDatastoreFile(dsPath, cmd.Datacenter)
	if err != nil {
		return err
	}

	return task.Wait()
}

func (cmd *vmdk) CreateVM(spec *configSpec) (*govmomi.VirtualMachine, error) {
	folders, err := cmd.Datacenter.Folders(cmd.Client)
	if err != nil {
		return nil, err
	}

	task, err := folders.VmFolder.CreateVM(cmd.Client, spec.ToSpec(), cmd.ResourcePool, nil)
	if err != nil {
		return nil, err
	}

	info, err := task.WaitForResult(nil)
	if err != nil {
		return nil, err
	}

	return govmomi.NewVirtualMachine(info.Result.(types.ManagedObjectReference)), nil
}

func (cmd *vmdk) CloneVM(vm *govmomi.VirtualMachine, name string) (*govmomi.VirtualMachine, error) {
	folders, err := cmd.Datacenter.Folders(cmd.Client)
	if err != nil {
		return nil, err
	}

	spec := types.VirtualMachineCloneSpec{
		Config:   &types.VirtualMachineConfigSpec{},
		Location: types.VirtualMachineRelocateSpec{},
	}

	task, err := vm.Clone(cmd.Client, folders.VmFolder, name, spec)
	if err != nil {
		return nil, err
	}

	info, err := task.WaitForResult(nil)
	if err != nil {
		return nil, err
	}

	return govmomi.NewVirtualMachine(info.Result.(types.ManagedObjectReference)), nil
}

func (cmd *vmdk) DestroyVM(vm *govmomi.VirtualMachine) error {
	var mvm mo.VirtualMachine

	err := cmd.Client.Properties(vm.Reference(), []string{"config.hardware"}, &mvm)
	if err != nil {
		return err
	}

	spec := new(configSpec)
	spec.RemoveDisks(&mvm)

	task, err := vm.Reconfigure(cmd.Client, spec.ToSpec())
	if err != nil {
		return err
	}

	err = task.Wait()
	if err != nil {
		return err
	}

	task, err = vm.Destroy(cmd.Client)
	if err != nil {
		return err
	}

	err = task.Wait()
	if err != nil {
		return err
	}

	return nil
}

type configSpec types.VirtualMachineConfigSpec

func (c *configSpec) ToSpec() types.VirtualMachineConfigSpec {
	return types.VirtualMachineConfigSpec(*c)
}

func (c *configSpec) AddChange(d types.BaseVirtualDeviceConfigSpec) {
	c.DeviceChange = append(c.DeviceChange, d)
}

func (c *configSpec) AddDisk(ds *govmomi.Datastore, path string) {
	controller := &types.VirtualLsiLogicController{
		VirtualSCSIController: types.VirtualSCSIController{
			SharedBus: types.VirtualSCSISharingNoSharing,
			VirtualController: types.VirtualController{
				BusNumber: 0,
				VirtualDevice: types.VirtualDevice{
					Key: -1,
				},
			},
		},
	}

	controllerSpec := &types.VirtualDeviceConfigSpec{
		Device:    controller,
		Operation: types.VirtualDeviceConfigSpecOperationAdd,
	}

	c.AddChange(controllerSpec)

	disk := &types.VirtualDisk{
		VirtualDevice: types.VirtualDevice{
			Key:           -1,
			ControllerKey: -1,
			UnitNumber:    -1,
			Backing: &types.VirtualDiskFlatVer2BackingInfo{
				VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
					FileName: ds.Path(path),
				},
				DiskMode:        string(types.VirtualDiskModePersistent),
				ThinProvisioned: true,
			},
		},
	}

	diskSpec := &types.VirtualDeviceConfigSpec{
		Device:    disk,
		Operation: types.VirtualDeviceConfigSpecOperationAdd,
	}

	c.AddChange(diskSpec)
}

func (c *configSpec) RemoveDisks(vm *mo.VirtualMachine) {
	for _, d := range vm.Config.Hardware.Device {
		switch device := d.(type) {
		case *types.VirtualDisk:
			removeOp := &types.VirtualDeviceConfigSpec{
				Operation: types.VirtualDeviceConfigSpecOperationRemove,
				Device:    device,
			}

			c.AddChange(removeOp)
		}
	}
}
