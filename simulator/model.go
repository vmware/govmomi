// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/simulator/vpx"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

type DelayConfig struct {
	// Delay specifies the number of milliseconds to delay serving a SOAP call. 0 means no delay.
	// This can be used to simulate a poorly performing vCenter or network lag.
	Delay int

	// Delay specifies the number of milliseconds to delay serving a specific method.
	// Each entry in the map represents the name of a method and its associated delay in milliseconds,
	// This can be used to simulate a poorly performing vCenter or network lag.
	MethodDelay map[string]int

	// DelayJitter defines the delay jitter as a coefficient of variation (stddev/mean).
	// This can be used to simulate unpredictable delay. 0 means no jitter, i.e. all invocations get the same delay.
	DelayJitter float64
}

// Model is used to populate a Model with an initial set of managed entities.
// This is a simple helper for tests running against a simulator, to populate an inventory
// with commonly used models.
// The inventory names generated by a Model have a string prefix per-type and integer suffix per-instance.
// The names are concatenated with their ancestor names and delimited by '_', making the generated names unique.
type Model struct {
	Service *Service `json:"-"`

	ServiceContent types.ServiceContent `json:"-"`
	RootFolder     mo.Folder            `json:"-"`

	// Autostart will power on Model created VMs when true
	Autostart bool `json:"-"`

	// Datacenter specifies the number of Datacenter entities to create
	// Name prefix: DC, vcsim flag: -dc
	Datacenter int `json:"datacenter"`

	// Portgroup specifies the number of DistributedVirtualPortgroup entities to create per Datacenter
	// Name prefix: DVPG, vcsim flag: -pg
	Portgroup int `json:"portgroup"`

	// PortgroupNSX specifies the number NSX backed DistributedVirtualPortgroup entities to create per Datacenter
	// Name prefix: NSXPG, vcsim flag: -nsx-pg
	PortgroupNSX int `json:"portgroupNSX"`

	// OpaqueNetwork specifies the number of OpaqueNetwork entities to create per Datacenter,
	// with Summary.OpaqueNetworkType set to nsx.LogicalSwitch and Summary.OpaqueNetworkId to a random uuid.
	// Name prefix: NSX, vcsim flag: -nsx
	OpaqueNetwork int `json:"opaqueNetwork"`

	// Host specifies the number of standalone HostSystems entities to create per Datacenter
	// Name prefix: H, vcsim flag: -standalone-host
	Host int `json:"host,omitempty"`

	// Cluster specifies the number of ClusterComputeResource entities to create per Datacenter
	// Name prefix: C, vcsim flag: -cluster
	Cluster int `json:"cluster"`

	// ClusterHost specifies the number of HostSystems entities to create within a Cluster
	// Name prefix: H, vcsim flag: -host
	ClusterHost int `json:"clusterHost,omitempty"`

	// Pool specifies the number of ResourcePool entities to create per Cluster
	// Note that every cluster has a root ResourcePool named "Resources", as real vCenter does.
	// For example: /DC0/host/DC0_C0/Resources
	// The root ResourcePool is named "RP0" within other object names.
	// When Model.Pool is set to 1 or higher, this creates child ResourcePools under the root pool.
	// Note that this flag is not effective on standalone hosts.
	// For example: /DC0/host/DC0_C0/Resources/DC0_C0_RP1
	// Name prefix: RP, vcsim flag: -pool
	Pool int `json:"pool"`

	// Datastore specifies the number of Datastore entities to create
	// Each Datastore will have temporary local file storage and will be mounted
	// on every HostSystem created by the ModelConfig
	// Name prefix: LocalDS, vcsim flag: -ds
	Datastore int `json:"datastore"`

	// Machine specifies the number of VirtualMachine entities to create per
	// ResourcePool. If the pool flag is specified, the specified number of virtual
	// machines will be deployed to each child pool and prefixed with the child
	// resource pool name. Otherwise they are deployed into the root resource pool,
	// prefixed with RP0. On standalone hosts, machines are always deployed into the
	// root resource pool without any prefix.
	// Name prefix: VM, vcsim flag: -vm
	Machine int `json:"machine"`

	// Folder specifies the number of Datacenter to place within a Folder.
	// This includes a folder for the Datacenter itself and its host, vm, network and datastore folders.
	// All resources for the Datacenter are placed within these folders, rather than the top-level folders.
	// Name prefix: F, vcsim flag: -folder
	Folder int `json:"folder"`

	// App specifies the number of VirtualApp to create per Cluster
	// Name prefix: APP, vcsim flag: -app
	App int `json:"app"`

	// Pod specifies the number of StoragePod to create per Cluster
	// Name prefix: POD, vcsim flag: -pod
	Pod int `json:"pod"`

	// Delay configurations
	DelayConfig DelayConfig `json:"-"`

	// total number of inventory objects, set by Count()
	total int

	dirs []string
}

// ESX is the default Model for a standalone ESX instance
func ESX() *Model {
	return &Model{
		ServiceContent: esx.ServiceContent,
		RootFolder:     esx.RootFolder,
		Autostart:      true,
		Datastore:      1,
		Machine:        2,
		DelayConfig: DelayConfig{
			Delay:       0,
			DelayJitter: 0,
			MethodDelay: nil,
		},
	}
}

// VPX is the default Model for a vCenter instance
func VPX() *Model {
	return &Model{
		ServiceContent: vpx.ServiceContent,
		RootFolder:     vpx.RootFolder,
		Autostart:      true,
		Datacenter:     1,
		Portgroup:      1,
		Host:           1,
		Cluster:        1,
		ClusterHost:    3,
		Datastore:      1,
		Machine:        2,
		DelayConfig: DelayConfig{
			Delay:       0,
			DelayJitter: 0,
			MethodDelay: nil,
		},
	}
}

// Map returns the Model.Service.Context.Map
func (m *Model) Map() *Registry {
	return m.Service.Context.Map
}

// Count returns a Model with total number of each existing type
func (m *Model) Count() Model {
	count := Model{}

	for ref, obj := range m.Map().objects {
		if _, ok := obj.(mo.Entity); !ok {
			continue
		}

		count.total++

		switch ref.Type {
		case "Datacenter":
			count.Datacenter++
		case "DistributedVirtualPortgroup":
			count.Portgroup++
		case "ClusterComputeResource":
			count.Cluster++
		case "Datastore":
			count.Datastore++
		case "HostSystem":
			count.Host++
		case "VirtualMachine":
			count.Machine++
		case "ResourcePool":
			count.Pool++
		case "VirtualApp":
			count.App++
		case "Folder":
			count.Folder++
		case "StoragePod":
			count.Pod++
		case "OpaqueNetwork":
			count.OpaqueNetwork++
		}
	}

	return count
}

func (*Model) fmtName(prefix string, num int) string {
	return fmt.Sprintf("%s%d", prefix, num)
}

// kinds maps managed object types to their vcsim wrapper types
var kinds = map[string]reflect.Type{
	"Alarm":                              reflect.TypeOf((*Alarm)(nil)).Elem(),
	"AlarmManager":                       reflect.TypeOf((*AlarmManager)(nil)).Elem(),
	"AuthorizationManager":               reflect.TypeOf((*AuthorizationManager)(nil)).Elem(),
	"ClusterComputeResource":             reflect.TypeOf((*ClusterComputeResource)(nil)).Elem(),
	"CustomFieldsManager":                reflect.TypeOf((*CustomFieldsManager)(nil)).Elem(),
	"CustomizationSpecManager":           reflect.TypeOf((*CustomizationSpecManager)(nil)).Elem(),
	"CryptoManagerKmip":                  reflect.TypeOf((*CryptoManagerKmip)(nil)).Elem(),
	"Datacenter":                         reflect.TypeOf((*Datacenter)(nil)).Elem(),
	"Datastore":                          reflect.TypeOf((*Datastore)(nil)).Elem(),
	"DatastoreNamespaceManager":          reflect.TypeOf((*DatastoreNamespaceManager)(nil)).Elem(),
	"DistributedVirtualPortgroup":        reflect.TypeOf((*DistributedVirtualPortgroup)(nil)).Elem(),
	"DistributedVirtualSwitch":           reflect.TypeOf((*DistributedVirtualSwitch)(nil)).Elem(),
	"DistributedVirtualSwitchManager":    reflect.TypeOf((*DistributedVirtualSwitchManager)(nil)).Elem(),
	"EnvironmentBrowser":                 reflect.TypeOf((*EnvironmentBrowser)(nil)).Elem(),
	"EventManager":                       reflect.TypeOf((*EventManager)(nil)).Elem(),
	"ExtensionManager":                   reflect.TypeOf((*ExtensionManager)(nil)).Elem(),
	"FileManager":                        reflect.TypeOf((*FileManager)(nil)).Elem(),
	"Folder":                             reflect.TypeOf((*Folder)(nil)).Elem(),
	"GuestOperationsManager":             reflect.TypeOf((*GuestOperationsManager)(nil)).Elem(),
	"HostDatastoreBrowser":               reflect.TypeOf((*HostDatastoreBrowser)(nil)).Elem(),
	"HostLocalAccountManager":            reflect.TypeOf((*HostLocalAccountManager)(nil)).Elem(),
	"HostNetworkSystem":                  reflect.TypeOf((*HostNetworkSystem)(nil)).Elem(),
	"HostCertificateManager":             reflect.TypeOf((*HostCertificateManager)(nil)).Elem(),
	"HostSystem":                         reflect.TypeOf((*HostSystem)(nil)).Elem(),
	"IpPoolManager":                      reflect.TypeOf((*IpPoolManager)(nil)).Elem(),
	"LicenseManager":                     reflect.TypeOf((*LicenseManager)(nil)).Elem(),
	"LicenseAssignmentManager":           reflect.TypeOf((*LicenseAssignmentManager)(nil)).Elem(),
	"OptionManager":                      reflect.TypeOf((*OptionManager)(nil)).Elem(),
	"OvfManager":                         reflect.TypeOf((*OvfManager)(nil)).Elem(),
	"PerformanceManager":                 reflect.TypeOf((*PerformanceManager)(nil)).Elem(),
	"PropertyCollector":                  reflect.TypeOf((*PropertyCollector)(nil)).Elem(),
	"ResourcePool":                       reflect.TypeOf((*ResourcePool)(nil)).Elem(),
	"SearchIndex":                        reflect.TypeOf((*SearchIndex)(nil)).Elem(),
	"SessionManager":                     reflect.TypeOf((*SessionManager)(nil)).Elem(),
	"StoragePod":                         reflect.TypeOf((*StoragePod)(nil)).Elem(),
	"StorageResourceManager":             reflect.TypeOf((*StorageResourceManager)(nil)).Elem(),
	"TaskManager":                        reflect.TypeOf((*TaskManager)(nil)).Elem(),
	"TenantTenantManager":                reflect.TypeOf((*TenantManager)(nil)).Elem(),
	"UserDirectory":                      reflect.TypeOf((*UserDirectory)(nil)).Elem(),
	"VcenterVStorageObjectManager":       reflect.TypeOf((*VcenterVStorageObjectManager)(nil)).Elem(),
	"ViewManager":                        reflect.TypeOf((*ViewManager)(nil)).Elem(),
	"VirtualApp":                         reflect.TypeOf((*VirtualApp)(nil)).Elem(),
	"VirtualDiskManager":                 reflect.TypeOf((*VirtualDiskManager)(nil)).Elem(),
	"VirtualMachine":                     reflect.TypeOf((*VirtualMachine)(nil)).Elem(),
	"VirtualMachineCompatibilityChecker": reflect.TypeOf((*VmCompatibilityChecker)(nil)).Elem(),
	"VirtualMachineProvisioningChecker":  reflect.TypeOf((*VmProvisioningChecker)(nil)).Elem(),
	"VmwareDistributedVirtualSwitch":     reflect.TypeOf((*DistributedVirtualSwitch)(nil)).Elem(),
}

func loadObject(ctx *Context, content types.ObjectContent) (mo.Reference, error) {
	var obj mo.Reference
	id := content.Obj

	kind, ok := kinds[id.Type]
	if ok {
		obj = reflect.New(kind).Interface().(mo.Reference)
	}

	if obj == nil {
		// No vcsim wrapper for this type, e.g. IoFilterManager
		x, err := mo.ObjectContentToType(content, true)
		if err != nil {
			return nil, err
		}
		obj = x.(mo.Reference)
	} else {
		if len(content.PropSet) == 0 {
			// via NewServiceInstance()
			ctx.Map.setReference(obj, id)
		} else {
			// via Model.Load()
			dst := getManagedObject(obj).Addr().Interface().(mo.Reference)
			err := mo.LoadObjectContent([]types.ObjectContent{content}, dst)
			if err != nil {
				return nil, err
			}
		}

		if x, ok := obj.(interface{ init(*Registry) }); ok {
			x.init(ctx.Map)
		}
	}

	return obj, nil
}

// resolveReferences attempts to resolve any object references that were not included via Load()
// example: Load's dir only contains a single OpaqueNetwork, we need to create a Datacenter and
// place the OpaqueNetwork in the Datacenter's network folder.
func (m *Model) resolveReferences(ctx *Context) error {
	dc, ok := ctx.Map.Any("Datacenter").(*Datacenter)
	if !ok {
		// Need to have at least 1 Datacenter
		root := ctx.Map.Get(ctx.Map.content().RootFolder).(*Folder)
		ref := root.CreateDatacenter(ctx, &types.CreateDatacenter{
			This: root.Self,
			Name: "DC0",
		}).(*methods.CreateDatacenterBody).Res.Returnval
		dc = ctx.Map.Get(ref).(*Datacenter)
	}

	for ref, val := range ctx.Map.objects {
		me, ok := val.(mo.Entity)
		if !ok {
			continue
		}
		e := me.Entity()
		if e.Parent == nil || ref.Type == "Folder" {
			continue
		}
		if ctx.Map.Get(*e.Parent) == nil {
			// object was loaded without its parent, attempt to foster with another parent
			switch e.Parent.Type {
			case "Folder":
				folder := dc.folder(ctx, me)
				e.Parent = &folder.Self
				log.Printf("%s adopted %s", e.Parent, ref)
				folderPutChild(ctx, folder, me)
			default:
				return fmt.Errorf("unable to foster %s with parent type=%s", ref, e.Parent.Type)
			}
		}
		// TODO: resolve any remaining orphan references via mo.References()
	}

	return nil
}

func (m *Model) decode(path string, data interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	dec := xml.NewDecoder(f)
	dec.TypeFunc = types.TypeFunc()
	err = dec.Decode(data)
	_ = f.Close()
	return err
}

func (m *Model) loadMethod(obj mo.Reference, dir string) error {
	dir = filepath.Join(dir, obj.Reference().Encode())

	info, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	zero := reflect.Value{}
	for _, x := range info {
		name := strings.TrimSuffix(x.Name(), ".xml") + "Response"
		path := filepath.Join(dir, x.Name())
		response := reflect.ValueOf(obj).Elem().FieldByName(name)
		if response == zero {
			return fmt.Errorf("field %T.%s not found", obj, name)
		}
		if err = m.decode(path, response.Addr().Interface()); err != nil {
			return err
		}
	}

	return nil
}

// NewContext initializes a Context with a NewRegistry
func NewContext() *Context {
	r := NewRegistry()

	return &Context{
		Context: context.Background(),
		Session: &Session{
			UserSession: types.UserSession{
				Key: uuid.New().String(),
			},
			Registry: NewRegistry(),
			Map:      r,
		},
		Map: r,
	}
}

// Load Model from the given directory, as created by the 'govc object.save' command.
func (m *Model) Load(dir string) error {
	ctx := NewContext()
	var s *ServiceInstance

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if path == dir {
				return nil
			}
			return filepath.SkipDir
		}
		if filepath.Ext(path) != ".xml" {
			return nil
		}

		var content types.ObjectContent
		err = m.decode(path, &content)
		if err != nil {
			return err
		}

		if content.Obj == vim25.ServiceInstance {
			s = new(ServiceInstance)
			s.Self = content.Obj
			ctx.Map.Put(s)
			return mo.LoadObjectContent([]types.ObjectContent{content}, &s.ServiceInstance)
		}

		if s == nil {
			ctx, s = NewServiceInstance(ctx, m.ServiceContent, m.RootFolder)
		}

		obj, err := loadObject(ctx, content)
		if err != nil {
			return err
		}

		if x, ok := obj.(interface{ model(*Model) error }); ok {
			if err = x.model(m); err != nil {
				return err
			}
		}

		return m.loadMethod(ctx.Map.Put(obj), dir)
	})

	if err != nil {
		return err
	}

	m.Service = New(ctx, s)

	return m.resolveReferences(ctx)
}

// Create populates the Model with the given ModelConfig
func (m *Model) Create() error {
	ctx := NewContext()
	m.Service = New(NewServiceInstance(ctx, m.ServiceContent, m.RootFolder))
	return m.CreateInfrastructure(ctx)
}

func (m *Model) CreateInfrastructure(ctx *Context) error {
	client := m.Service.client()
	root := object.NewRootFolder(client)

	// After all hosts are created, this var is used to mount the host datastores.
	var hosts []*object.HostSystem
	hostMap := make(map[string][]*object.HostSystem)

	// We need to defer VM creation until after the datastores are created.
	var vms []func() error
	// 1 DVS per DC, added to all hosts
	var dvs *object.DistributedVirtualSwitch
	// 1 NIC per VM, backed by a DVPG if Model.Portgroup > 0
	vmnet := esx.EthernetCard.Backing

	// addHost adds a cluster host or a standalone host.
	addHost := func(name string, f func(types.HostConnectSpec) (*object.Task, error)) (*object.HostSystem, error) {
		spec := types.HostConnectSpec{
			HostName: name,
		}

		task, err := f(spec)
		if err != nil {
			return nil, err
		}

		info, err := task.WaitForResult(context.Background(), nil)
		if err != nil {
			return nil, err
		}

		host := object.NewHostSystem(client, info.Result.(types.ManagedObjectReference))
		hosts = append(hosts, host)

		if dvs != nil {
			config := &types.DVSConfigSpec{
				Host: []types.DistributedVirtualSwitchHostMemberConfigSpec{{
					Operation: string(types.ConfigSpecOperationAdd),
					Host:      host.Reference(),
				}},
			}

			task, _ = dvs.Reconfigure(ctx, config)
			_, _ = task.WaitForResult(context.Background(), nil)
		}

		return host, nil
	}

	// addMachine returns a func to create a VM.
	addMachine := func(prefix string, host *object.HostSystem, pool *object.ResourcePool, folders *object.DatacenterFolders) {
		nic := esx.EthernetCard
		nic.Backing = vmnet
		ds := types.ManagedObjectReference{}

		f := func() error {
			for i := 0; i < m.Machine; i++ {
				name := m.fmtName(prefix+"_VM", i)

				config := types.VirtualMachineConfigSpec{
					Name:    name,
					GuestId: string(types.VirtualMachineGuestOsIdentifierOtherGuest),
					Files: &types.VirtualMachineFileInfo{
						VmPathName: "[LocalDS_0]",
					},
				}

				if pool == nil {
					pool, _ = host.ResourcePool(ctx)
				}

				var devices object.VirtualDeviceList

				scsi, _ := devices.CreateSCSIController("pvscsi")
				ide, _ := devices.CreateIDEController()
				cdrom, _ := devices.CreateCdrom(ide.(*types.VirtualIDEController))
				disk := devices.CreateDisk(scsi.(types.BaseVirtualController), ds,
					config.Files.VmPathName+" "+path.Join(name, "disk1.vmdk"))
				disk.CapacityInKB = int64(units.GB*10) / units.KB
				disk.StorageIOAllocation = &types.StorageIOAllocationInfo{Limit: types.NewInt64(-1)}

				devices = append(devices, scsi, cdrom, disk, &nic)

				config.DeviceChange, _ = devices.ConfigSpec(types.VirtualDeviceConfigSpecOperationAdd)

				task, err := folders.VmFolder.CreateVM(ctx, config, pool, host)
				if err != nil {
					return err
				}

				info, err := task.WaitForResult(ctx, nil)
				if err != nil {
					return err
				}

				vm := object.NewVirtualMachine(client, info.Result.(types.ManagedObjectReference))

				if m.Autostart {
					task, _ = vm.PowerOn(ctx)
					_, _ = task.WaitForResult(ctx, nil)
				}
			}

			return nil
		}

		vms = append(vms, f)
	}

	nfolder := 0

	for ndc := 0; ndc < m.Datacenter; ndc++ {
		dcName := m.fmtName("DC", ndc)
		folder := root
		fName := m.fmtName("F", nfolder)

		// If Datacenter > Folder, don't create folders for the first N DCs.
		if nfolder < m.Folder && ndc >= (m.Datacenter-m.Folder) {
			f, err := folder.CreateFolder(ctx, fName)
			if err != nil {
				return err
			}
			folder = f
		}

		dc, err := folder.CreateDatacenter(ctx, dcName)
		if err != nil {
			return err
		}

		folders, err := dc.Folders(ctx)
		if err != nil {
			return err
		}

		if m.Pod > 0 {
			for pod := 0; pod < m.Pod; pod++ {
				_, _ = folders.DatastoreFolder.CreateStoragePod(ctx, m.fmtName(dcName+"_POD", pod))
			}
		}

		if folder != root {
			// Create sub-folders and use them to create any resources that follow
			subs := []**object.Folder{&folders.DatastoreFolder, &folders.HostFolder, &folders.NetworkFolder, &folders.VmFolder}

			for _, sub := range subs {
				f, err := (*sub).CreateFolder(ctx, fName)
				if err != nil {
					return err
				}

				*sub = f
			}

			nfolder++
		}

		if m.Portgroup > 0 || m.PortgroupNSX > 0 {
			var spec types.DVSCreateSpec
			spec.ConfigSpec = &types.VMwareDVSConfigSpec{}
			spec.ConfigSpec.GetDVSConfigSpec().Name = m.fmtName("DVS", 0)

			task, err := folders.NetworkFolder.CreateDVS(ctx, spec)
			if err != nil {
				return err
			}

			info, err := task.WaitForResult(ctx, nil)
			if err != nil {
				return err
			}

			dvs = object.NewDistributedVirtualSwitch(client, info.Result.(types.ManagedObjectReference))
		}

		for npg := 0; npg < m.Portgroup; npg++ {
			name := m.fmtName(dcName+"_DVPG", npg)
			spec := types.DVPortgroupConfigSpec{
				Name:     name,
				Type:     string(types.DistributedVirtualPortgroupPortgroupTypeEarlyBinding),
				NumPorts: 1,
			}

			task, err := dvs.AddPortgroup(ctx, []types.DVPortgroupConfigSpec{spec})
			if err != nil {
				return err
			}
			if err = task.Wait(ctx); err != nil {
				return err
			}

			// Use the 1st DVPG for the VMs eth0 backing
			if npg == 0 {
				// AddPortgroup_Task does not return the moid, so we look it up by name
				net := ctx.Map.Get(folders.NetworkFolder.Reference()).(*Folder)
				pg := ctx.Map.FindByName(name, net.ChildEntity)

				vmnet, _ = object.NewDistributedVirtualPortgroup(client, pg.Reference()).EthernetCardBackingInfo(ctx)
			}
		}

		for npg := 0; npg < m.PortgroupNSX; npg++ {
			name := m.fmtName(dcName+"_NSXPG", npg)
			spec := types.DVPortgroupConfigSpec{
				Name:        name,
				Type:        string(types.DistributedVirtualPortgroupPortgroupTypeEarlyBinding),
				BackingType: string(types.DistributedVirtualPortgroupBackingTypeNsx),
			}

			task, err := dvs.AddPortgroup(ctx, []types.DVPortgroupConfigSpec{spec})
			if err != nil {
				return err
			}
			if err = task.Wait(ctx); err != nil {
				return err
			}
		}

		// Must use simulator methods directly for OpaqueNetwork
		networkFolder := ctx.Map.Get(folders.NetworkFolder.Reference()).(*Folder)

		for i := 0; i < m.OpaqueNetwork; i++ {
			var summary types.OpaqueNetworkSummary
			summary.Name = m.fmtName(dcName+"_NSX", i)
			err := networkFolder.AddOpaqueNetwork(ctx, summary)
			if err != nil {
				return err
			}
		}

		for nhost := 0; nhost < m.Host; nhost++ {
			name := m.fmtName(dcName+"_H", nhost)

			host, err := addHost(name, func(spec types.HostConnectSpec) (*object.Task, error) {
				return folders.HostFolder.AddStandaloneHost(ctx, spec, true, nil, nil)
			})
			if err != nil {
				return err
			}

			addMachine(name, host, nil, folders)
		}

		for ncluster := 0; ncluster < m.Cluster; ncluster++ {
			clusterName := m.fmtName(dcName+"_C", ncluster)

			cluster, err := folders.HostFolder.CreateCluster(ctx, clusterName, types.ClusterConfigSpecEx{})
			if err != nil {
				return err
			}

			for nhost := 0; nhost < m.ClusterHost; nhost++ {
				name := m.fmtName(clusterName+"_H", nhost)

				_, err = addHost(name, func(spec types.HostConnectSpec) (*object.Task, error) {
					return cluster.AddHost(ctx, spec, true, nil, nil)
				})
				if err != nil {
					return err
				}
			}

			rootRP, err := cluster.ResourcePool(ctx)
			if err != nil {
				return err
			}

			prefix := clusterName + "_RP"

			// put VMs in cluster RP if no child RP(s) configured
			if m.Pool == 0 {
				addMachine(prefix+"0", nil, rootRP, folders)
			}

			// create child RP(s) with VMs
			for childRP := 1; childRP <= m.Pool; childRP++ {
				spec := types.DefaultResourceConfigSpec()

				p, err := rootRP.Create(ctx, m.fmtName(prefix, childRP), spec)
				addMachine(m.fmtName(prefix, childRP), nil, p, folders)
				if err != nil {
					return err
				}
			}

			prefix = clusterName + "_APP"

			for napp := 0; napp < m.App; napp++ {
				rspec := types.DefaultResourceConfigSpec()
				vspec := NewVAppConfigSpec()
				name := m.fmtName(prefix, napp)

				vapp, err := rootRP.CreateVApp(ctx, name, rspec, vspec, nil)
				if err != nil {
					return err
				}

				addMachine(name, nil, vapp.ResourcePool, folders)
			}
		}

		hostMap[dcName] = hosts
		hosts = nil
	}

	if m.ServiceContent.RootFolder == esx.RootFolder.Reference() {
		// ESX model
		host := object.NewHostSystem(client, esx.HostSystem.Reference())

		dc := object.NewDatacenter(client, esx.Datacenter.Reference())
		folders, err := dc.Folders(ctx)
		if err != nil {
			return err
		}

		hostMap[dc.Reference().Value] = append(hosts, host)

		addMachine(host.Reference().Value, host, nil, folders)
	}

	for dc, dchosts := range hostMap {
		for i := 0; i < m.Datastore; i++ {
			err := m.createLocalDatastore(dc, m.fmtName("LocalDS_", i), dchosts)
			if err != nil {
				return err
			}
		}
	}

	for _, createVM := range vms {
		err := createVM()
		if err != nil {
			return err
		}
	}

	// Turn on delay AFTER we're done building the service content
	m.Service.delay = &m.DelayConfig

	return nil
}

func (m *Model) createTempDir(dc string, name string) (string, error) {
	dir, err := os.MkdirTemp("", fmt.Sprintf("govcsim-%s-%s-", dc, name))
	if err == nil {
		m.dirs = append(m.dirs, dir)
	}
	return dir, err
}

func (m *Model) createLocalDatastore(dc string, name string, hosts []*object.HostSystem) error {
	ctx := context.Background()
	dir, err := m.createTempDir(dc, name)
	if err != nil {
		return err
	}

	for _, host := range hosts {
		dss, err := host.ConfigManager().DatastoreSystem(ctx)
		if err != nil {
			return err
		}

		_, err = dss.CreateLocalDatastore(ctx, name, dir)
		if err != nil {
			return err
		}
	}

	return nil
}

// Remove cleans up items created by the Model, such as local datastore directories
func (m *Model) Remove() {
	ctx := m.Service.Context
	// Remove associated vm containers, if any
	ctx.Map.m.Lock()
	for _, obj := range ctx.Map.objects {
		if vm, ok := obj.(*VirtualMachine); ok {
			vm.svm.remove(ctx)
		}
	}
	ctx.Map.m.Unlock()

	for _, dir := range m.dirs {
		_ = os.RemoveAll(dir)
	}
}

// Run calls f with a Client connected to a simulator server instance, which is stopped after f returns.
func (m *Model) Run(f func(context.Context, *vim25.Client) error) error {
	defer m.Remove()

	if m.Service == nil {
		err := m.Create()
		if err != nil {
			return err
		}
		// Only force TLS if the provided model didn't have any Service.
		m.Service.TLS = new(tls.Config)
	}

	m.Service.RegisterEndpoints = true

	s := m.Service.NewServer()
	defer s.Close()

	ctx := m.Service.Context
	c, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		return err
	}

	defer c.Logout(ctx)

	return f(ctx, c.Client)
}

// Run calls Model.Run for each model and will panic if f returns an error.
// If no model is specified, the VPX Model is used by default.
func Run(f func(context.Context, *vim25.Client) error, model ...*Model) {
	m := model
	if len(m) == 0 {
		m = []*Model{VPX()}
	}

	for i := range m {
		err := m[i].Run(f)
		if err != nil {
			panic(err)
		}
	}
}

// Test calls Run and expects the caller propagate any errors, via testing.T for example.
func Test(f func(context.Context, *vim25.Client), model ...*Model) {
	Run(func(ctx context.Context, c *vim25.Client) error {
		f(ctx, c)
		return nil
	}, model...)
}

// RunContainer runs a vm container with the given args
func RunContainer(ctx context.Context, c *vim25.Client, vm mo.Reference, args string) error {
	obj, ok := vm.(*object.VirtualMachine)
	if !ok {
		obj = object.NewVirtualMachine(c, vm.Reference())
	}

	task, err := obj.PowerOff(ctx)
	if err != nil {
		return err
	}
	_ = task.Wait(ctx) // ignore InvalidPowerState if already off

	task, err = obj.Reconfigure(ctx, types.VirtualMachineConfigSpec{
		ExtraConfig: []types.BaseOptionValue{&types.OptionValue{Key: "RUN.container", Value: args}},
	})
	if err != nil {
		return err
	}
	if err = task.Wait(ctx); err != nil {
		return err
	}

	task, err = obj.PowerOn(ctx)
	if err != nil {
		return err
	}
	return task.Wait(ctx)
}

// delay sleeps according to DelayConfig. If no delay specified, returns immediately.
func (dc *DelayConfig) delay(method string) {
	d := 0
	if dc.Delay > 0 {
		d = dc.Delay
	}
	if md, ok := dc.MethodDelay[method]; ok {
		d += md
	}
	if dc.DelayJitter > 0 {
		d += int(rand.NormFloat64() * dc.DelayJitter * float64(d))
	}
	if d > 0 {
		//fmt.Printf("Delaying method %s %d ms\n", method, d)
		time.Sleep(time.Duration(d) * time.Millisecond)
	}
}
