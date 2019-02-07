
### 0.22.2 (2020-02-04)

* Expose soap client default transport

### 0.22.1 (2020-01-13)

* Fix SAML token auth using Holder-of-Key with delegated Bearer identity against 6.7 U3b+

### 0.22.0 (2020-01-10)

* Add OVF properties to library.Deploy method

* Add retry support for HTTP status codes

* Use cs.identity service type for sts endpoint lookups

* Add Content Library VM template APIs

* Add SearchIndex FindAllByDnsName and FindAllByIp methods

* Fix HostSystem.ManagementIPs to use SelectedVnic

* Change generated ResourceReductionToToleratePercent to pointer type

* Add DistributedVirtualSwitch.ReconfigureDVPort method

* Add VirtualMachine.IsTemplate method

* Add GetInventoryPath to NetworkReference interface

* Support HoK tokens with Interactive Users

* Replace mo.LoadRetrievePropertiesResponse with mo.LoadObjectContent

* Add VirtualHardwareSection.StorageItem

* Add ResourcePool.Owner method

* Add VirtualMachine.QueryChangedDiskAreas method

* Update generated code to vSphere 6.7u3

* Add option to propagate MissingSet faults in property.WaitForUpdates

* Add content library subscription support

* Fix deadlock for keep alive handlers that attempt log in

* Add CNS API bindings

* Add FetchCapabilityMetadata method to Pbm client

* Add v4 option to VirtualMachine.WaitForIP

* Add VirtualHardwareSection.StorageItem

### 0.21.0 (2019-07-24)

* Add vsan package

* Add vslm (FCD) global catalog support

* Add content library support

### 0.20.0 (2019-02-06)

### Build

* Refactored Travis-CI to use containers

### Govc

* fix object.collect error for multiple objects with same path
* add device name match support to device.ls and device.remove
* add vm.disk.attach -mode flag
* add category option to relevant tags commands
* add vm.create -version option
* fields.set can now add missing fields
* add fields.info command

### Vcsa

* bump to 6.7.0 U1

### Vcsim

* require authentication in vapi simulator
* Resolve issue making device changes on clone (resolves [#1355](https://github.com/vmware/govmomi/issues/1355))
* fix SearchDatastore task info entity
* add EnvironmentBrowser support
* avoid zero IP address in GOVC_URL output
* avoid panic when event template is not defined
* implement RefreshStorageInfo method for virtual machine
* configure HostSystem port
* datastore.upload now creates missing directories in destination path.
* add option to run container as vm
* add SessionIsActive support
* fix fault detail encoding
* support base types in property filter
* PropertyCollector should not require PathSet
* allow '.' in vm name
* populate VM guest.net field
* add SearchIndex FindByDnsName support
* correct property update in RemoveSnapshotTask
* update VM snapshot methods to change VM properties with UpdateObject
* support setting vm fields via extraConfig
* update VM configureDevices method to change VM properties with UpdateObject
* update VM device add operation - stricter key generation, new InvalidDeviceSpec condition
* add PBM support
* put VM into registry earlier during CreateVM
* add datastore access check for vm host placement
* add task_manager description property templates
* fix defaults when generating vmdk paths
* fix custom_fields_manager test
* replace HostSystem template IP with vcsim listen address
* Change CustomFieldsManager SetField to use ctx.WithLock and add InvalidArgument fault check.
* update DVS methods to use UpdateObject instead of setting fields directly
* add vslm support
* add setCustomValue support
* add fault message to PropertyCollector RetrieveProperties
* add HistoryCollector scrollable view support


<a name="v0.19.0"></a>
## [v0.19.0](https://github.com/vmware/govmomi/compare/v0.18.0...v0.19.0) (2018-09-30)

### Example

* uniform unit for host memory

### Govc

* fix test case for new cluster.rule.info command
* add new command cluster.rule.info

### README

* Fix path to LICENSE.txt file

### Vcsa

* bump to 6.7.0d
* bump to 6.7.0c release
* bump to 6.7.0a release

### Vcsim

* add dvpg networks to HostSystem.Parent
* add support for tags API
* Logout should not unregister PropertyCollector singleton
* add ResetVM and SuspendVM support
* add support for PropertyCollector incremental updates
* do not include DVS in HostSystem.Network


<a name="v0.18.0"></a>
## [v0.18.0](https://github.com/vmware/govmomi/compare/v0.17.1...v0.18.0) (2018-05-24)

### Govc

* import.ovf pool flag should be optional if host is specified
* avoid Login() attempt if username is not set
* add json support to find command
* fix host.esxcli error handling

### Vcsim

* add STS simulator
* use VirtualDisk CapacityInKB for device summary
* add property collector field type mapping for integer arrays


<a name="v0.17.1"></a>
## [v0.17.1](https://github.com/vmware/govmomi/compare/v0.17.0...v0.17.1) (2018-03-19)

### Vcsim

* add Destroy method for Folder and Datacenter types
* add EventManager.QueryEvents


<a name="v0.17.0"></a>
## [v0.17.0](https://github.com/vmware/govmomi/compare/v0.16.0...v0.17.0) (2018-02-28)

### Govc

* fix vm.clone to use -net flag when source does not have a NIC
* object.collect support for raw filters
* fix host.info CPU usage
* add -cluster flag to license.assign command
* allow columns in guest login password ([#972](https://github.com/vmware/govmomi/issues/972))

### Object

* Return correct helper object for OpaqueNetwork

### Toolbox

* validate request offset in ListFiles ([#946](https://github.com/vmware/govmomi/issues/946))

### Vcsim

* add simulator.Datastore type
* set VirtualMachine summary.config.instanceUuid
* update HostSystem.Summary.Host reference
* add EventManager support
* stats related fixes
* avoid data races
* respect VirtualDeviceConfigSpec FileOperation
* avoid keeping the VM log file open
* add UpdateOptions support
* add session support
* Add VM.MarkAsTemplate support
* more input spec honored in ReConfig VM
* Initialize VM fields properly
* Honor the input spec in ReConfig VM
* Add HostLocalAccountManager
* workaround xml ns issue with pyvsphere ([#958](https://github.com/vmware/govmomi/issues/958))
* add MakeDirectoryResponse ([#938](https://github.com/vmware/govmomi/issues/938))
* copy RoleList for AuthorizationManager ([#932](https://github.com/vmware/govmomi/issues/932))
* apply vm spec NumCoresPerSocket ([#930](https://github.com/vmware/govmomi/issues/930))
* Configure dvs with the dvs config spec
* Add VirtualMachine guest ID validation ([#921](https://github.com/vmware/govmomi/issues/921))
* add QueryVirtualDiskUuid ([#920](https://github.com/vmware/govmomi/issues/920))
* update ServiceContent to 6.5 ([#917](https://github.com/vmware/govmomi/issues/917))


<a name="v0.16.0"></a>
## [v0.16.0](https://github.com/vmware/govmomi/compare/v0.15.0...v0.16.0) (2017-11-08)

### Govc

* Fix VM clone when source doesn't have vNics
* add tasks and task.cancel commands
* add reboot option to host.shutdown

### Readme

* fix formatting of listing ([#908](https://github.com/vmware/govmomi/issues/908))

### Toolbox

* avoid race when closing channels on stop
* reset session when invalidated by the vmx
* default to tar format for directory archives
* make gzip optional for directory archive transfer
* avoid blocking the RPC channel when transferring process IO
* use host management IP for guest file transfer
* add Client Upload and Download methods
* support single file download via archive handler
* SendGuestInfo before the vmx asks us to
* update vmw-guestinfo
* remove receiver from DefaultStartCommand
* map exec.ErrNotFound to vix.FileNotFound
* pass URL to ArchiveHandler Read/Write methods
* make directory archive read/write customizable
* add http and exec round trippers
* fix ListFiles when given a symlink
* support transferring /proc files from guest

### Vcsim

* preserve order in QueryIpPools ([#914](https://github.com/vmware/govmomi/issues/914))
* return moref from Task.Run ([#913](https://github.com/vmware/govmomi/issues/913))
* Implement IpPoolManager lifecycle
* add autostart option to power on VMs ([#906](https://github.com/vmware/govmomi/issues/906))
* use soapenv namespace for Fault types
* various property additions
* Generate similar ref value like VC
* Add moref to vm's summary
* validate authz privilege ids
* AuthorizationManager additions
* Add IpPoolManager
* VirtualDisk file backing datastore is optional
* add PerformanceManager
* Implement add/update/remove roles
* Generate device filename in CreateVM
* add AuthorizationManager
* populate vm snapshot fields
* Add UpdateNetworkConfig to HostNetworkSystem
* Implement virtual machine snapshot
* set VirtualDisk backing datastore
* Implement enter/exit maintenance mode
* Implement add/remove license
* add portgroup related operations
* add fields support
* remove use of df program for datastore info
* add FileQuery support to datastore search
* add HostConfigInfo template
* add HostSystem hardware property
* Fix merging of default devices
* Add cdrom and scsi controller to Model VMs

### Vim25

* Move internal stuff to internal package

### Vscim

* Implement UserDirectory


<a name="v0.15.0"></a>
## [v0.15.0](https://github.com/vmware/govmomi/compare/v0.14.0...v0.15.0) (2017-06-19)


<a name="v0.14.0"></a>
## [v0.14.0](https://github.com/vmware/govmomi/compare/v0.13.0...v0.14.0) (2017-04-08)

### Emacs

* add metric select

### Finder

* support changing object root in find mode


<a name="v0.13.0"></a>
## [v0.13.0](https://github.com/vmware/govmomi/compare/v0.12.1...v0.13.0) (2017-03-02)

### Emacs

* add govc-command-history

### Finder

* support automatic Folder recursion ([#663](https://github.com/vmware/govmomi/issues/663))

### Vcsim

* esxcli FirewallInfo fixes ([#661](https://github.com/vmware/govmomi/issues/661))


<a name="v0.12.1"></a>
## [v0.12.1](https://github.com/vmware/govmomi/compare/v0.12.0...v0.12.1) (2016-12-19)


<a name="v0.12.0"></a>
## [v0.12.0](https://github.com/vmware/govmomi/compare/v0.11.4...v0.12.0) (2016-12-01)


<a name="v0.11.4"></a>
## [v0.11.4](https://github.com/vmware/govmomi/compare/v0.11.3...v0.11.4) (2016-11-15)


<a name="v0.11.3"></a>
## [v0.11.3](https://github.com/vmware/govmomi/compare/v0.11.2...v0.11.3) (2016-11-08)


<a name="v0.11.2"></a>
## [v0.11.2](https://github.com/vmware/govmomi/compare/v0.11.1...v0.11.2) (2016-11-01)


<a name="v0.11.1"></a>
## [v0.11.1](https://github.com/vmware/govmomi/compare/v0.11.0...v0.11.1) (2016-10-27)


<a name="v0.11.0"></a>
## [v0.11.0](https://github.com/vmware/govmomi/compare/v0.10.0...v0.11.0) (2016-10-25)


<a name="v0.10.0"></a>
## [v0.10.0](https://github.com/vmware/govmomi/compare/v0.9.0...v0.10.0) (2016-10-20)


<a name="v0.9.0"></a>
## [v0.9.0](https://github.com/vmware/govmomi/compare/v0.8.0...v0.9.0) (2016-09-09)

### Emacs

* add govc-session-network
* add govc json diff


<a name="v0.8.0"></a>
## [v0.8.0](https://github.com/vmware/govmomi/compare/v0.7.1...v0.8.0) (2016-06-30)


<a name="v0.7.1"></a>
## [v0.7.1](https://github.com/vmware/govmomi/compare/v0.7.0...v0.7.1) (2016-06-03)


<a name="v0.7.0"></a>
## [v0.7.0](https://github.com/vmware/govmomi/compare/v0.6.2...v0.7.0) (2016-06-02)


<a name="v0.6.2"></a>
## [v0.6.2](https://github.com/vmware/govmomi/compare/v0.6.1...v0.6.2) (2016-05-13)


<a name="v0.6.1"></a>
## [v0.6.1](https://github.com/vmware/govmomi/compare/v0.6.0...v0.6.1) (2016-04-30)


<a name="v0.6.0"></a>
## [v0.6.0](https://github.com/vmware/govmomi/compare/v0.5.0...v0.6.0) (2016-04-29)


<a name="v0.5.0"></a>
## [v0.5.0](https://github.com/vmware/govmomi/compare/v0.4.0...v0.5.0) (2016-03-30)


<a name="v0.4.0"></a>
## [v0.4.0](https://github.com/vmware/govmomi/compare/v0.3.0...v0.4.0) (2016-02-26)


<a name="v0.3.0"></a>
## [v0.3.0](https://github.com/vmware/govmomi/compare/v0.2.0...v0.3.0) (2016-01-15)

### VirtualMachine

* Add Customize function on object.VirtualMachine


<a name="v0.2.0"></a>
## [v0.2.0](https://github.com/vmware/govmomi/compare/prerelease-v0.1.0-73-gfc131d4...v0.2.0) (2015-09-15)

### Reverts

* Add Host information to vm.info


<a name="prerelease-v0.1.0-73-gfc131d4"></a>
## [prerelease-v0.1.0-73-gfc131d4](https://github.com/vmware/govmomi/compare/prerelease-v0.1.0-62-g7734772...prerelease-v0.1.0-73-gfc131d4) (2015-07-13)


<a name="prerelease-v0.1.0-62-g7734772"></a>
## [prerelease-v0.1.0-62-g7734772](https://github.com/vmware/govmomi/compare/prerelease-v0.1.0-52-g871f5d4...prerelease-v0.1.0-62-g7734772) (2015-07-06)


<a name="prerelease-v0.1.0-52-g871f5d4"></a>
## [prerelease-v0.1.0-52-g871f5d4](https://github.com/vmware/govmomi/compare/v0.1.0...prerelease-v0.1.0-52-g871f5d4) (2015-06-16)

### Reverts

* Fix git dirty status error in build script


<a name="v0.1.0"></a>
## [v0.1.0](https://github.com/vmware/govmomi/compare/test...v0.1.0) (2015-03-17)

