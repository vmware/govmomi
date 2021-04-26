
<a name="v0.25.0"></a>
## [Release v0.25.0](https://github.com/vmware/govmomi/compare/v0.24.1...v0.25.0)

> Release Date: 2021-04-16

### 💫 `govc` (CLI)

- add vm.customize -dns-suffix flag
- update test images URL
- log invalid NetworkMapping.Name with import.ova command
- revert pretty print pruning optimization
- add library.update command
- add session.ls -S flag
- add tree command
- include sub task fault messages on failure
- remove device.boot -firmware default
- add '-trace' and '-verbose' flags

### 💫 `vcsim` (Simulator)

- put verbose logging behind '-trace' flag
- add moid value mapping mappings
- add vsan simulator

### 📖 Commits

- Fix folder write for govc container [6fe8d60a]
- docs: update for 0.25 release [e18b601f]
- govc: add vm.customize -dns-suffix flag [1ac314c3]
- Add Cron Docker Login Action [22d911f6]
- govc: update test images URL [60e0e895]
- Add action to automate release [3385b3e0]
- govc: log invalid NetworkMapping.Name with import.ova command [cdf3ace6]
- vcsim: put verbose logging behind '-trace' flag [27d8d2e4]
- govc: revert pretty print pruning optimization [f8b3d8a8]
- vcsim: add moid value mapping mappings [0ef4ae22]
- First step towards release automation [df08d4b2]
- export simulator.task.Wait() [f9b79a4f]
- Ensure lock hand-off to simulator.Task goroutine [917c4ec8]
- Simulator Task Delay [b45b228f]
- Make Simulator Tasks Async [4b59b652]
- Associate every registry lock with a Context. [bc52c793]
- Wait until VM creation completes before adding to folder [054971ee]
- govc: add library.update command [35481467]
- Fix race in simulator's PropertyCollector [7403b470]
- Add action to block WIP PRs [aadb2082]
- govc: add session.ls -S flag [749c2239]
- [3ad0f415] Update Dockerfiles and .goreleaser.yml [bc297330]
- vcsim: add vsan simulator [082f9927]
- Add a stretched cluster conversion command. [8c38d56d]
- gofmt [408c531a]
- Update govc/flags/output.go [e8a6b126]
- Add more badges [bf54a7c4]
- govc: add tree command [93245c1e]
- govc: include sub task fault messages on failure [790f9ce6]
- Use Github Actions Status Badges [07e6e923]
- govc: remove device.boot -firmware default [d2a353ba]
- Add chainable RoundTripper support to endpoint clients [4ed615f6]
- Add the vSAN stretched cluster reference. [bab95d26]
- Fix events example [6ff33db7]
- govc: add '-trace' and '-verbose' flags [de6032e0]
- Add support for calling vCenter for VLSM ExtendDisk and InflateDisk [7aae8dfb]
- Use Github Actions [2c57a8a3]
- Handling invalid reader size [0d155a61]
- Using progress reader in WriteFile [b70542a5]

<a name="v0.24.1"></a>
## [Release v0.24.1](https://github.com/vmware/govmomi/compare/v0.24.0...v0.24.1)

> Release Date: 2021-03-17

### 💫 `govc` (CLI)

- metric command enhancements and fixes
- fix vm.migrate search index flags
- fix cluster.usage Free field
- fix session curl when given a URL query
- validate license.remove
- validate required library.clone NAME arg
- note 'disk.ls -R' in volume.rm help
- add device.info examples to get disk UUID and vmdk
- fix vm.markasvm examples
- fix incorrect DeviceID value in device.pci.add
- add IPv6 support to vm.customize

### 💫 `vcsim` (Simulator)

- fix Task.Info.Entity in RevertToSnapshot_Task
- set VirtualMachine.Config.CreateDate property
- support EventFilterSpec.Time
- emit CustomizationSucceeded event from CustomizeVM
- add DistributedVirtualSwitchManager
- set VirtualDisk backing UUID
- move product suffix in ServiceContent.About
- use linked list for EventHistoryCollector
- escape datastore name
- record/replay EnvironmentBrowser.QueryConfigOption
- fix EventHistoryCollector fixes
- switch bats tests from esx to vcsim env
- fixes for PowerCLI Get-VirtualNetwork

### 📖 Commits

- Add client test file for vslm package to validate register disk and cns create volume [7a276bf6]
- Fix performance.Manager.SampleByName truncation [dc29aa29]
- Added UpdateServiceMessage to Session Manager [18b53fd2]
- govc: metric command enhancements and fixes [63bb5c1e]
- govc: fix vm.migrate search index flags [7844a8c2]
- Drop clusterDistribution from vSAN 7.0 update and create spec elements [7ab111bd]
- Marshal soapFaultError as JSON [52631496]
- fix tab indentation [f9e323a6]
- add tests and implement HA Ready Condition [ae129ba0]
- implement vSphere HA additional delay to VM HA overrides in govc [f34b3fa2]
- vcsim: fix Task.Info.Entity in RevertToSnapshot_Task [25970530]
- govc: fix cluster.usage Free field [5dacf627]
- use correct enum for vm restart priority [b7f9e034]
- Add support for snapshot size calculations [d3d49a36]
- Use a dash to indicate empty address [61bfa072]
- vcsim: set VirtualMachine.Config.CreateDate property [f0a045ac]
- vim25: fix race in TemporaryNetworkError retry func [4d9a9000]
- ovf: add Config and ExtraConfig to VirtualHardwareSection [2f14e4b2]
- Add vSAN 7.0U1 release constant [50328780]
- Update .goreleaser.yml [886573de]
- Change the address type to automatic [1cdb3164]
- Remove duplicate cns bindings from vsan directory [667a3791]
- govc: fix session curl when given a URL query [f71bcf25]
- Update volume ACL spec to add delete field [d92f41de]
- govc: validate license.remove [c954c2a5]
- Update ConfigureVolumeACLs bindings in cns types [2a4f8c8a]
- govc: validate required library.clone NAME arg [3b25c3f1]
- govc: note 'disk.ls -R' in volume.rm help [344b7a30]
- govc: add device.info examples to get disk UUID and vmdk [8942055a]
- govc: fix vm.markasvm examples [1b0af949]
- govc-env --save default [543e52ea]
- Little fix for "govc-env --save without config name" [0a5f2a99]
- gen: require nokogiri 1.11.0 or higher [4a7a0b45]
- govc: fix incorrect DeviceID value in device.pci.add [add8be5a]
- vcsim: support EventFilterSpec.Time [e51eb2b9]
- govc: add IPv6 support to vm.customize [1f4f5640]
- vcsim: emit CustomizationSucceeded event from CustomizeVM [8e45fa4a]
- vcsim: add DistributedVirtualSwitchManager [c000bd6e]
- vcsim: set VirtualDisk backing UUID [bcd5fa87]
- vcsim: move product suffix in ServiceContent.About [ccdcbe89]
- vcsim: use linked list for EventHistoryCollector [393e7330]
- vcsim: escape datastore name [9c4dc1a1]
- vcsim: record/replay EnvironmentBrowser.QueryConfigOption [9c2fe70f]
- vcsim: fix EventHistoryCollector fixes [5fd7e264]
- Skip tests that require docker on TravisCI [40a2cf0b]
- toolbox: skip tests that require Linux [00ee2911]
- vcsim: switch bats tests from esx to vcsim env [0b755a59]
- Updated projects to include VMware Event Broker Appliance [c6d5264a]
- ExampleCollector_Retrieve: Add missing err return [ae44a547]
- examples: add NetworkReference.EthernetCardBackingInfo [38da87ff]
- vcsim: fixes for PowerCLI Get-VirtualNetwork [3f1caf82]
- Fix DvsNetworkRuleQualifier interface [041a98b8]
- SHA-1 deprecated in 2011, sha256sum for releases [44e05fe4]

<a name="v0.24.0"></a>
## [Release v0.24.0](https://github.com/vmware/govmomi/compare/v0.23.1...v0.24.0)

> Release Date: 2020-12-21

### 💫 `govc` (CLI)

- fix build.sh git tag injection
- add cluster.usage command
- add volume.ls -ds option
- add device.boot -firmware option
- add dvs.portgroup.{add,change} '-auto-expand' flag
- fix object.collect ContainerView updates
- document vm.disk.attach -link behavior
- fix vm.clone panic when target VM already exists
- support sparse backing in vm.disk.change
- add CNS volume ls and rm commands
- add find -p flag
- add storage.policy commands
- add vm.console -wss flag
- support multi value flags in host.esxcli command
- add namespace.cluster.ls command

### 💫 `vcsim` (Simulator)

- include stderr in log message when volume import fails
- include stderr in log message when container fails to start
- rewrite vmfs path from saved model
- QueryConfigOptionEx Spec is optional
- support inventory updates in ContainerView
- set VirtualDevice.Connectable default for removable devices
- add AuthorizationManager methods
- set VirtualDisk backing option defaults
- add CloneVApp_Task support
- fix ListView.Modify
- avoid ViewManager.ModifyListView race
- add ListView to race test
- add mechanism for modeling methods
- fix save/load property collection for VmwareDistributedVirtualSwitch
- Honoring the instance uuid provided in spec by caller ([#2052](https://github.com/vmware/govmomi/issues/2052))

### 📖 Commits

- govc: fix build.sh git tag injection [1ec59a7c]
- Update docs for 0.24 release [164b9217]
- vcsim: include stderr in log message when volume import fails [bf80efab]
- Add batch APIs for multiple tags to object [4080e177]
- govc: add cluster.usage command [31c0836e]
- examples: add Folder.CreateVM [7178588c]
- Add test for vsan host config [2b962f3f]
- Add function to get vsan host config [165d7cb4]
- govc: add volume.ls -ds option [79514c81]
- Add Configure ACL go bindings [f7ff79df]
- vcsim: include stderr in log message when container fails to start [1f3fb17c]
- Add wrappers for retrieving vsan properties [3b83040a]
- Use gofmt [12e8969c]
- Add vSAN 7.0 API bindings [6454dbd4]
- Add vSAN 7.0 API bindings [6a216a52]
- Regenerate against vSphere 7.0U1 release [be15ad6c]
- govc: add device.boot -firmware option [5e57b3f6]
- vcsim: rewrite vmfs path from saved model [e1c4b06e]
- Change CnsCreateVolume to return PlacementResult for static volume provisioning. Also add unit test for this case. [26635452]
- govc: add dvs.portgroup.{add,change} '-auto-expand' flag [4d82f0ff]
- vcsim: QueryConfigOptionEx Spec is optional [bcdfb298]
- Add Placement object in CNS CreateVolume response. Add corresponding test. [8b194c23]
- Use available ctx in enable cluster network lookup [b085fc33]
- Cleanup some redundant code for cluster namespace enabling [f6f336ab]
- change negative one to rand neg int32 [d04f2b49]
- go binding for CNS RelocateVolume API [f819befd]
- fix the goimports validation error [ed93ea7d]
- support trunk mode port group [f402c0e1]
- change key default from -1 to rand neg int32 vsphere 7 introduced a key collision detection error when adding devices com.vmware.vim.vpxd.vmprov.duplicateDeviceKey which causes -1 keys to return an error of duplicate if you try and add two devices in the same AddDevice call [ff575977]
- Add option to disable secure cookies with non-TLS endpoints [39acef43]
- simulator: fix container vm example [ae19e30f]
- vcsim: support inventory updates in ContainerView [73e1af55]
- Add namespace.cluster.disable cmd + formatting fixes [593cd20d]
- Add namespace.cluster.enable cmd to govc [782ed95c]
- Make ListStorageProfiles public -> for enabling clusters in govc [e7403032]
- Adds support for enabling cluster namespaces via API [53965796]
- govc: fix object.collect ContainerView updates [4a1d05ac]
- govc: document vm.disk.attach -link behavior [e84d0d18]
- vcsim: set VirtualDevice.Connectable default for removable devices [a76123b2]
- examples: add ContainerView retrieve clusters [b4f7243b]
- vcsim: add AuthorizationManager methods [b195dd57]
- vcsim: set VirtualDisk backing option defaults [a71f6c77]
- examples: use session.Cache [1d21fff9]
- examples: add events [8af8cef6]
- Add ClusterDistribution field for CNS telemetry and Drop optional fields not known to the prior releases [3e2a8071]
- Fix for fatal error: concurrent map iteration and map write [4acfb726]
- Adding VsanQueryObjectIdentities and QueryVsanObjects [01610887]
- vcsim: add CloneVApp_Task support [fbde3866]
- govc: fix vm.clone panic when target VM already exists [70a9ced4]
- govc: support sparse backing in vm.disk.change [a97e6168]
- govc: add CNS volume ls and rm commands [3380cd30]
- sts: fix SignRequest bodyhash for non-empty request body [f9d7bfdf]
- vapi: add WCP support bundle bindings [7b4e997b]
- vcsim: fix ListView.Modify [aae78223]
- Add AuthorizationManager.HasUserPrivilegeOnEntities wrapper [0e4bce43]
- vim25/xml: sync with Go 1.15 encoding/xml [81207eab]
- govc: add find -p flag [f7170fd2]
- Add internal.InventoryPath helper [d49123c9]
- govc: add storage.policy commands [b40cdd8a]
- add / remove pci passthrough device for one VM [0c5cdd5d]
- govc: add vm.console -wss flag [d0111d28]
- Add sms generated types and methods [94bc8497]
- examples: fix simulator.RunContainer on MacOSX [e153061f]
- finder: simplify direct use of InventoryPath func [99fe9954]
- Added Instant Clone feature Resolves: [#1392](https://github.com/vmware/govmomi/issues/1392) [3760bd6c]
- govc: support multi value flags in host.esxcli command [86374ea2]
- vcsim: avoid ViewManager.ModifyListView race [9cca13ab]
- vcsim: add ListView to race test [156b1cb0]
- Add ExtendDisk and InflateDisk wrappers to vlsm/object_manager [f903d5da]
- Add AttachDisk and DetachDisk wrappers for the virtualMachine object. [073cc310]
- vapi: add tags.Manager.GetAttachedTagsOnObjects example [a0c7e829]
- Vsan Performance Data Collection API ([#2021](https://github.com/vmware/govmomi/issues/2021)) [378a24c4]
- vcsim: add mechanism for modeling methods [55f6f952]
- vcsim: fix save/load property collection for VmwareDistributedVirtualSwitch [69942fe2]
- bats: test fixes for running on MacOSX [fe3becfa]
- Merge branch 'master' into pc/HardwareInfoNotReplicatingInCloning [0422a070]
- vapi: add Content Library example [9f12aae4]
- vcsim: Honoring the instance uuid provided in spec by caller ([#2052](https://github.com/vmware/govmomi/issues/2052)) [33121b87]
- Setting hardware properties in clone VM spec from template VM [9a07942b]
- govc: add namespace.cluster.ls command [ebcfa3d4]
- vapi: add namespace management client and vcsim support [11d45e54]
- vapi: add helper support "/api" endpoint [cdc44d5e]

<a name="v0.23.1"></a>
## [Release v0.23.1](https://github.com/vmware/govmomi/compare/v0.23.0...v0.23.1)

> Release Date: 2020-07-02

### 💫 `vcsim` (Simulator)

- add HostNetworkSystem.QueryNetworkHint
- use HostNetworkSystem wrapper with -load flag
- set HostSystem IP in cluster AddHost_Task
- add PbmQueryAssociatedProfile method

### 📖 Commits

- check if config isn't nil before returning an uuid [b7add48c]
- added support for returning array of BaseCnsVolumeOperationResult for QueryVolumeInfo API [12955a6c]
- vcsim: add HostNetworkSystem.QueryNetworkHint [0697d33f]
- Merge branch 'master' into master [a5c9e1f0]
- adding in link to OPS [c14e3bc5]
- vcsim: use HostNetworkSystem wrapper with -load flag [d7f4bba6]
- vcsim: set HostSystem IP in cluster AddHost_Task [916b12e6]
- vcsim: add PbmQueryAssociatedProfile method [e63ec002]
- examples: add property.Collector.Retrieve example [0bbb6a7d]

<a name="v0.23.0"></a>
## [Release v0.23.0](https://github.com/vmware/govmomi/compare/prerelease-v0.22.1-247-g770fcba2...v0.23.0)

> Release Date: 2020-06-11

### 💫 `govc` (CLI)

- ipath search flag does not require a Datacenter

### 📖 Commits

- Update docs for 0.23 release [b639ab4c]
- vapi: use header authentication in file Upload/Download [be7742f2]
- provided examples for vm.clone and host.esxcli [50846878]
- Add appliance log forwarding config handler and govc verb ([#1994](https://github.com/vmware/govmomi/issues/1994)) [aa97c4d3]
- govc: ipath search flag does not require a Datacenter [4f19eb6d]

<a name="prerelease-v0.22.1-247-g770fcba2"></a>
## [Release prerelease-v0.22.1-247-g770fcba2](https://github.com/vmware/govmomi/compare/v0.22.2...prerelease-v0.22.1-247-g770fcba2)

> Release Date: 2020-05-29

### 💫 `govc` (CLI)

- support raw object references in import.ova NetworkMapping
- support find with -customValue filter
- support VirtualApp with -pool flag
- add -version flag to datastore.create command
- add session.login -X flag
- vm.clone ResourcePool is optional when -cluster is specified
- add REST support for session.login -cookie flag
- fix host.info CPU usage
- add session.ls -r flag
- add a VM template clone example
- ignore ManagedObjectNotFound errors in 'find' command
- remove ClientFlag.WithRestClient
- do not try to start a VM template
- add guest directory upload/download examples
- add vm.change -uuid flag
- enable library.checkout and library.checkin by default
- avoid truncation in object.collect
- add import.spec support for remote URLs
- support optional compute.policy.ls argument
- add vm.change '-memory-pin' flag
- support nested groups in sso.group.update
- add content library helpers
- add cluster.group.ls -l flag
- use OutputFlag for import.spec
- add library.clone -ovf flag
- fix doc for -g flag (guest id) choices
- add object.collect -o flag
- output formatting enhancements
- add find -l flag
- save sessions using sha256 ID

### 💫 `vcsim` (Simulator)

- CreateSnapshotTask now returns moref in result
- add lookup ServiceRegistration example
- add AuthorizationManager.HasPrivilegeOnEntities
- traverse configManager.datastoreSystem in object.save
- traverse configManager.virtualNicManager in object.save
- traverse configManager.networkSystem in object.save
- add extraConfigAlias table
- add EventHistoryCollector.ResetCollector implementation
- fixes for PowerCLI
- apply ExtraConfig after devices
- add another test/example for DVS host member validation
- validate DVS membership
- fix flaky library subscriber test
- avoid panic if ovf:capacityAllocationUnits is not present
- support QueryConfigOptionEx GuestId param
- VM templates do not have a ResourcePool
- validate session key in TerminateSession method
- unique MAC address for VM NICs
- create vmdk directory if needed
- support VMs with the same name
- support Folder in RelocateVM spec
- add guest operations support
- add HostStorageSystem support
- avoid possible panic in UnregisterVM_Task
- support tags with the same name
- add docs on generated inventory names
- add support for NSX backed networks

### 📖 Commits

- Finder: support DistributedVirtualSwitch traversal [7cdad997]
- govc: support raw object references in import.ova NetworkMapping [10c22fd1]
- vcsim: CreateSnapshotTask now returns moref in result [c3fe4f84]
- vcsim: add lookup ServiceRegistration example [b0af443c]
- simulator: fix handling of nil Reference in container walk [84f1b733]
- Adding sunProfileName in pbm.CapabilityProfileCreateSpec [b5b434b0]
- providing examples for govc guest.run [2111324a]
- Bump to vSphere version 7 [0eef3b29]
- go binding for CNS QueryVolumeInfo API [b277903e]
- Move simulator lookupservice registration into ServiceInstance [a048ea52]
- modify markdown link at simulator.Model [30f1a71a]
- Add REST session keep alive support [7881f541]
- vapi: sync access to rest.Client.SessionID [3aa9aaba]
- simulator: refactor folder children operations [0a53ac4b]
- simulator: relax ResourcePool constraint for createVM operation [b9152f85]
- simulator: relax typing condition on RP parent [70e9d821]
- simulator: relax ViewManager typing constraints [502b7efa]
- simulator: remove data race in VM creation flow [634fdde1]
- simulator: protect datastore freespace updates against data races [6eda0169]
- govc: support find with -customValue filter [414c548d]
- Add logic to return default HealthStatus in CnsCreateVolume. [487ca0d6]
- govc: support VirtualApp with -pool flag [0bf0e761]
- govc: add -version flag to datastore.create command [f1ae45f5]
- Add support for attach-tag-to-multiple-objects [d0751307]
- simulator: relax excessive type assertions in SearchIndex [5682b1f2]
- Modify parenthesis for markdown link [39a4da90]
- vcsim: add AuthorizationManager.HasPrivilegeOnEntities [34734712]
- 1. Add retry for CNS Create API with backing disk url 2. Fix binding for CnsAlreadyRegisteredFault [92d464b9]
- Add sample test for Create CNS API using backing disk Url path [235582fe]
- 1. Add BackingDiskUrlPath and CnsAlreadyFault go bindings to CNS APIs 2. Update CreateVolume CNS Util  to include BackingDiskUrlPath [b187863a]
- Add GetProfileNameByID functionality to PBM [409279fa]
- vcsim: traverse configManager.datastoreSystem in object.save [228e0a8f]
- vcsim: traverse configManager.virtualNicManager in object.save [8acac02a]
- vcsim: traverse configManager.networkSystem in object.save [8a4ab564]
- govc: add session.login -X flag [43e4f8c2]
- govc: vm.clone ResourcePool is optional when -cluster is specified [70b7e1b4]
- govc: add REST support for session.login -cookie flag [2c5ff385]
- Add guest.FileManager.TransferURL test [6ccaf303]
- Avoid possible nil pointer dereference in guest TransferURL [03c7611e]
- Fix delegated Holder-of-Key token signature [44a78f96]
- Update to vSphere 7 APIs [11b2aa1a]
- vcsim: add extraConfigAlias table [4b8a5988]
- vcsim: add EventHistoryCollector.ResetCollector implementation [a0fe825a]
- vcsim: fixes for PowerCLI [558747b3]
- vcsim: apply ExtraConfig after devices [9ae04495]
- govc: fix host.info CPU usage [7d66cf9a]
- vcsim: add another test/example for DVS host member validation [4286d7cd]
- Revert to using sha1 for session cache file names [515621d1]
- Default to separate session cache directories [f103a87a]
- vcsim: validate DVS membership [7e24bfcb]
- govc: add session.ls -r flag [244a8369]
- govc: add a VM template clone example [6c68ccf2]
- govc: ignore ManagedObjectNotFound errors in 'find' command [bb6ae4ab]
- vcsim: fix flaky library subscriber test [853656fd]
- Fix existing goimport issue [571f64e7]
- vcsim: avoid panic if ovf:capacityAllocationUnits is not present [7426e2fd]
- Add non-null HostLicensableResourceInfo to HostSystem [9e57f983]
- govc: remove ClientFlag.WithRestClient [210541fe]
- govc: do not try to start a VM template [75e9e80d]
- simulator: add interface for VirtualDiskManager [d9220e5d]
- vcsim: support QueryConfigOptionEx GuestId param [55599668]
- vcsim: VM templates do not have a ResourcePool [67d593cc]
- govc: add guest directory upload/download examples [667e6fbe]
- govc: add vm.change -uuid flag [167f5d83]
- govc: enable library.checkout and library.checkin by default [bcd06cee]
- Refactor govc session persistence into session/cache package [9d4faa6d]
- govc: avoid truncation in object.collect [6f087ded]
- Remove Task from function names in Task struct receiver methods [7a1fef65]
- Add SetTaskState SetTaskDescription UpdateProgress to object package [dd839655]
- vcsim: validate session key in TerminateSession method [469e11b9]
- Revert compute policy support [af41ae09]
- Fix the types of errors returned from VSLM tasks to be their originl vim faults rather than just wrappers of localized error msg [ad612b3e]
- Remove extra err check [9e82230f]
- govc: add import.spec support for remote URLs [e9bb4772]
- skip tests when env is not set [273aaf71]
- removing usage of spew package [159c423c]
- vapi: prefer header authn to cookie authn [76caec95]
- Dropping fields in entity metadata for 6.7u3 [6c04cfa0]
- using right version and namespace from sdk/vsanServiceVersions.xml for cns client. making cns/client.go backward compatible to vsan67u3 by dropping unknown elements [8d15081f]
- Add nil check for taskInfo result before typecasting CnsVolumeOperationBatchResult [8dfb29f5]
- fixing CnsFault go binding [d68bbf9b]
- syncing vmodl changes [5482bd07]
- fixing go binding for CnsVolumeOperationResult and CnsFault [3bcace84]
- Fixing govmomi binding for CNS as per latest VMODL for CnsVsanFileShareBackingDetails. Also fixed cns/client_test.go accordingly. [3c756cbd]
- Adding new API to get cluster configuration [4254df70]
- removing space before omitempty tag [0eacb4ed]
- Resolve bug in Simulator regarding BackingObjectDetails [59ce7e4a]
- Change the backingObjectDetails attribute to point to interface BaseCnsBackingObjectDetails [6ad7e87d]
- Add resize support [601f1ded]
- Updating go binding for vsan fileshare vmodl updates [56049aa4]
- Add CnsQuerySelectionNameType and CnsKubernetesEntityType back [af798c01]
- Add bindings for vSANFS and extend CNS bindings to support file volume [af2723fd]
- update taskClientVersion for vsphere 7.0 [4e7b9b00]
- govc: support optional compute.policy.ls argument [692c1008]
- Modified return type for Get policy [a7d4a77d]
- Compute Policy support [4007484e]
- vcsim: unique MAC address for VM NICs [88d298ff]
- govc: add vm.change '-memory-pin' flag [814e4e5c]
- reset all for recursive calls fix format error [de8bcf25]
- Fixed ContainerView.RetrieveWithFilter fetch all specs if empty list of properties given [57efe91f]
- Avoid possible panic in Filter.MatchProperty [5af5ac8d]
- Add vAPI create binding for compute policy [85889777]
- govc: support nested groups in sso.group.update [56e878a5]
- Added prefix toggle parameter to govc export.ovf [6f46ef8a]
- Disk mode should override default value in vm.disk.attach [6d3196e4]
- Replaced ClassOvfParams with ClassDeploymentOptionParams [4be7a425]
- vcsim: create vmdk directory if needed [c4f820dd]
- Add Content Library subscriptions support [1ab6fe09]
- vcsim: support VMs with the same name [488205f0]
- vcsim: support Folder in RelocateVM spec [68349a27]
- Update CONTRIBUTING to have more info about running CI tests, checks. [6a6a7875]
- Expose Soap client default transport (a.k.a. its http client default transport) [a73c0d4f]
- govc: add content library helpers [84346733]
- build(deps): bump nokogiri from 1.10.4 to 1.10.8 in /gen [a225a002]
- Avoid ServiceContent requirement in lookup.NewClient [b4395d65]
- fix blog links [c1e828cb]
- toolbox: bump test VM memory for current CoreOS release [863430ba]
- govc: add cluster.group.ls -l flag [0ccfd912]
- Add Namespace support to UseServiceVersion [1af6ec1d]
- vcsim: add guest operations support [ab1298d5]
- examples: Fixed error is not logging in example.go [0e4b487e]
- Add Content Library item copy support [f36e13fc]
- vcsim: add HostStorageSystem support [7ffb9255]
- govc: use OutputFlag for import.spec [ae84c494]
- govc: add library.clone -ovf flag [2dda4daa]
- vcsim: avoid possible panic in UnregisterVM_Task [77b31b84]
- govc: fix doc for -g flag (guest id) choices [519d302d]
- vcsim: support tags with the same name [617c18e7]
- govc: add object.collect -o flag [e582cbd1]
- Apply gomvomi vim25/xml changes [0c6eafc1]
- Simplify ObjectName method [4da54375]
- govc: output formatting enhancements [d2e6b7df]
- vcsim: add docs on generated inventory names [dfcf9437]
- govc: add find -l flag [e64c2423]
- govc: save sessions using sha256 ID [4db4430c]
- vcsim: add support for NSX backed networks [4cfc2905]
- examples: add ContainerView.Find [c17eb769]
- Import golang/go/src/encoding/xml v1.13.6 [36056ae6]
- Avoid encoding/xml import [346cf59a]
- fix simulator disk manager fault message. [9cbe57db]
- Add permissions for NoCryptoAdmin [7f685c23]

<a name="v0.22.2"></a>
## [Release v0.22.2](https://github.com/vmware/govmomi/compare/v0.22.1...v0.22.2)

> Release Date: 2020-02-13

### 📖 Commits

- Avoid ServiceContent requirement in lookup.NewClient [e7df0c11]

<a name="v0.22.1"></a>
## [Release v0.22.1](https://github.com/vmware/govmomi/compare/v0.22.0...v0.22.1)

> Release Date: 2020-01-13

### 📖 Commits

- Release version 0.22.1 [da368950]
- Fix AttributeValue.C14N for 6.7u3 [a62b12cf]
- Add finder example for MultipleFoundError [c3d102b1]
- vapi: add CreateTag example [802e5899]
- vapi: Add cluster modules client and simulator [15630b90]

<a name="v0.22.0"></a>
## [Release v0.22.0](https://github.com/vmware/govmomi/compare/v0.20.3...v0.22.0)

> Release Date: 2020-01-10

### 💫 `govc` (CLI)

- guest -i flag only applies to ProcessManager
- add 5.0 to vm.create hardware version map
- guest.run improvements
- add vm.customize multiple IP support
- fix library.info output formatting
- add optional library.info details
- handle xsd:string responses
- add library.info details
- fixup tasks formatting
- remove guest.run toolbox dependency
- default to simple esxcli format when hints fields is empty
- add datacenter create/delete examples
- fix vm.create doc regarding -on flag
- add device.boot -secure flag
- add doc on vm.info -r flag
- avoid env for -cluster placement flag
- add default library.create thumbprint
- add thumbprint flag to library.create
- add vm.power doc
- support vm.customize without a managed spec
- fixup usage suggestions
- add vm.customize command
- fix datacenter.info against nested folders
- add vm.change -latency flag
- validate moref argument
- add guest.df command
- support library paths in tags.attach commands
- add datastore.info -H flag
- add sso.group commands
- host.vnic.info -json support
- add context to LoadX509KeyPair error
- add vm.change hot-add options
- change logs.download -default=false
- increase guest.ps -X poll interval
- add -options support to library.deploy
- rename vcenter.deploy to library.deploy
- move library.item.update commands to library.session
- consolidate library commands
- export Archive Path field
- add vm.change vpmc-enabled flag
- fix vm.change against templates
- fix option.set for int32 type values
- add datastore.maintenance.{enter,exit} commands
- FCD workarounds
- add datastore.cluster.info Description
- add permission.remove -f flag

### 💫 `vcsim` (Simulator)

- propagate VirtualMachineCloneSpec.Template
- add -trace-file option
- Get IP address on non-default container network
- avoid possible panic in VirtualMachine.Destroy_Task
- automatically set Context.Caller
- remove container volumes
- bind mount BIOS UUID DMI files
- validate VirtualDisk UnitNumber
- add Floppy Drive support to OVF manager
- properly initialize portgroup portKeys field
- add vim25 client helper to vapi simulator
- use VMX_ prefix for guestinfo env vars
- don't allow duplicate names for Folder/StoragePod
- pass guestinfo vars as env vars to container vms
- add CustomizationSpecManager support
- simplify container vm arguments input
- update docs
- add record/playback functionality
- add VirtualMachine.Rename_Task support
- add feature examples
- Ensure that extraConfig from clone spec is added to VM being cloned
- use exported response helpers in vapi/simulator
- avoid ViewManager.ViewList
- avoid race in ViewManager
- use TLS in simulator.Run
- rename Example to Run
- add endpoint registration mechanism
- add PlaceVm support ([#1589](https://github.com/vmware/govmomi/issues/1589))
- DefaultDatastoreID is optional in library deploy
- add support to override credentials
- fix host uuid
- use stable UUIDs for inventory objects
- Press any key to exit
- Update NetworkInfo.Portgroup in simulator
- remove httptest.serve flag
- add library.deploy support
- add ovf manager
- fork httptest server package
- add content library support
- set guest.toolsRunningStatus property

### ⏮ Reverts

- gen: retain omitempty field tag with int pointer types

### 📖 Commits

- Update docs for 0.22 release [317707be]
- govc: guest -i flag only applies to ProcessManager [aed39212]
- Clarify DVS EthernetCardBackingInfo error message [22308123]
- Add Content Library synchronization support [a1c98f14]
- govc: add 5.0 to vm.create hardware version map [704b335f]
- Clarify System.Read privilege requirement for PortGroup backing [4e907d99]
- Fix guest.FileManager.TransferURL cache [554d9284]
- Remove toolbox specific guest run implementation [9b8da88a]
- govc: guest.run improvements [965109ae]
- govc: add vm.customize multiple IP support [ee28fcfd]
- Add OVF properties to library deploy ([#1755](https://github.com/vmware/govmomi/issues/1755)) [40001828]
- govc: fix library.info output formatting [68b3ea9f]
- vcsim: propagate VirtualMachineCloneSpec.Template [198b97ca]
- govc: add optional library.info details [5bb7f391]
- Added the missing RetrieveSnapshotDetails API in VSLM ([#1763](https://github.com/vmware/govmomi/issues/1763)) [2509e907]
- govc: handle xsd:string responses [d8ac7e51]
- Add library ItemType constants [45b3685d]
- Add retry support for HTTP status codes [f3e2c3ce]
- govc: add library.info details [31d3e357]
- govc: fixup tasks formatting [182c84a3]
- govc: remove guest.run toolbox dependency [08fb2b02]
- VSLM: fixed the missing param in the QueryChangedDiskArea API impl [b10bcbf3]
- vcsim: add -trace-file option [168a6a04]
- examples: output VM names in performance example [72b1cd92]
- vcsim: Get IP address on non-default container network [32eeeb24]
- Move to cs.identity service type for sso admin endpoint [f9f69237]
- vcsim: avoid possible panic in VirtualMachine.Destroy_Task [1427d581]
- vcsim: automatically set Context.Caller [067d58be]
- govc: default to simple esxcli format when hints fields is empty [a727283f]
- Move to cs.identity service type for sts endpoint [08adb5d6]
- vcsim: remove container volumes [9e8e9a5a]
- vcsim: bind mount BIOS UUID DMI files [6cc814b8]
- Content Library: add CheckOuts support [e793289c]
- Content Library: VM Template support [66c9b10c]
- examples: add Common.Rename [f4b3cda7]
- Pass vm.Config.Uuid into the "VM" container via an env var [19a726f7]
- govc: add datacenter create/delete examples [204af3c5]
- examples: add VirtualMachine.Customize [dab4ab0d]
- govc: fix vm.create doc regarding -on flag [f6c57ee7]
- govc: add device.boot -secure flag [8debfcc3]
- vcsim: validate VirtualDisk UnitNumber [9aec1386]
- Revert "gen: retain omitempty field tag with int pointer types" [7914609d]
- Add CustomizationSpecManager.Info method and example [9b2c5cc6]
- vcsim: add Floppy Drive support to OVF manager [d7e43b4e]
- Implement some missing methods ("*All*" variants) on SearchIndex MOB [0bf21ec2]
- govc: add doc on vm.info -r flag [2bb2a6ed]
- vcsim: properly initialize portgroup portKeys field [8646dace]
- govc: avoid env for -cluster placement flag [e50368c6]
- Add ability to set DVS discovery protocol on create and change [91b1e0a7]
- Move to Go 1.13 [1e130141]
- govc: add default library.create thumbprint [f16eb276]
- govc: add thumbprint flag to library.create [d8325f34]
- Fix hostsystem ManagementIPs call [62c20113]
- Update DVS change to use finder.Network for a single object [c4a3908f]
- Fix usage instructions [ee6fe09d]
- gen: retain omitempty field tag with int pointer types [5e6f5e3f]
- vcsim: add vim25 client helper to vapi simulator [286bd5e9]
- Add ability to change a vnic on a host [841386f1]
- Add ability to change the MTU on a DVS that has already been created [391dd80b]
- Change MTU param to use flags.NewInt32 as the type [26a45d61]
- Add MTU flag for DVS creation [dbcfc3a8]
- Generate pointer type for ResourceReductionToToleratePercent [0399353f]
- Add nil checks for all HostConfigManager references [3f6b8ef5]
- vcsim: use VMX_ prefix for guestinfo env vars [c3163247]
- Add option to follow all struct fields in mo.References [5381f171]
- Refactor session KeepAlive tests to use vcsim [04e4835c]
- Avoid possible deadlock in KeepAliveHandler [7391c241]
- build(deps): bump nokogiri from 1.6.3.1 to 1.10.4 in /gen [41422ea4]
- vcsim: don't allow duplicate names for Folder/StoragePod [a3a09c04]
- Add a method to update ports on a distributed virtual switch [4c72d2e9]
- govc: add vm.power doc [0bad2bc2]
- govc: support vm.customize without a managed spec [45d322ea]
- govc: fixup usage suggestions [0a058e0f]
- vcsim: pass guestinfo vars as env vars to container vms [a0a2296e]
- vcsim: add CustomizationSpecManager support [903fe182]
- vcsim: simplify container vm arguments input [eda6bf3b]
- vcsim: update docs [0ce9b0a1]
- adding managed obj type to table [c538d867]
- govc: add vm.customize command [3185f7bc]
- Include object.save directory in output [b2a7b47e]
- Initial support for hybrid Model.Load [e8281f87]
- vcsim: add record/playback functionality [7755fbda]
- set stable vsan client version [8a3fa4f2]
- Avoid empty principal in HoK token request [9eaac5cb]
- Allow sending multiple characters through -c and name the keys [4a8da68d]
- add simple command list filter [3e3d3515]
- vcsim: add VirtualMachine.Rename_Task support [fe000674]
- support two tags with the same name [9166bbdb]
- added log type and password scrubber [344653c1]
- vcsim: add feature examples [d87cd5ac]
- Report errors when cdrom.insert fails [30fc2225]
- vslm: fix to throw errors on tasks that are completed with error state [a94f2d3a]
- added IsTemplate vm helper [37054f03]
- Fix object.collect with moref argument [d7aeb628]
- add GetInventoryPath to NetworkReference interface [0765aa63]
- Fix description of vm.keystrokes [9fb975b0]
- vapi: support DeleteLibrary with subscribed libraries [234aaf53]
- vcsim: Ensure that extraConfig from clone spec is added to VM being cloned [2cc33fa8]
- vcsim: use exported response helpers in vapi/simulator [70ad060e]
- vapi: refactor for external API implementations [b069efc0]
- vcsim: avoid ViewManager.ViewList [1e7aa6c2]
- vcsim: avoid race in ViewManager [9b0db1c2]
- a failing testcase that triggers with -race test [bd298f43]
- vapi: expand internal path constants [03422dd2]
- Support HoK tokens with Interactive Users [d296a5f8]
- Fix error check in session.Secret [c6226542]
- vcsim: use TLS in simulator.Run [28b5fc6c]
- Replace LoadRetrievePropertiesResponse with LoadObjectContent [f9b4bb05]
- Add VirtualHardwareSection.StorageItem [d84679eb]
- Check whether there's a NIC before updating guest.ipAddress [a23a5cb1]
- Add interactiveSession flag [8a069c27]
- vm.keystrokes -s (Allow spaces) [25526b21]
- examples: add VirtualMachine.CreateSnapshot [1828eee9]
- vapi: return info with current session query [ca3763e7]
- vcsim: rename Example to Run [f962095f]
- vcsim: add endpoint registration mechanism [43d69860]
- govc: fix datacenter.info against nested folders [1b159e27]
- vcsim: add PlaceVm support ([#1589](https://github.com/vmware/govmomi/issues/1589)) [c183577b]
- Add ResourcePool.Owner method [3e71d6be]
- vcsim: DefaultDatastoreID is optional in library deploy [b17f3a51]
- Update generated code to vSphere 6.7u3 [68980704]
- Add VirtualMachine.QueryChangedDiskAreas(). [7416741c]
- Content library: support library ID in Finder [8ef87890]
- Add option to propagate MissingSet faults in property.WaitForUpdates [e373feb8]
- examples: fix flag parsing [6ff7040e]
- govc: add vm.change -latency flag [149ba7ad]
- govc: validate moref argument [c35a532d]
- Add content library subscription support [54df157b]
- Fix deadlock for keep alive handlers that attempt log in [b86466b7]
- CNS go bindings [9ad64557]
- Add simulator.Model.Run example [9de3b854]
- Include url in Client.Download error [4285b614]
- vcsa: update to 6.7 U3 [caf0b6b3]
- Update vcsim Readme.md [7ac56b64]
- Update README.md [48ef35df]
- Use gnu xargs in bats tests on Darwin [a40837d8]
- Add FetchCapabilityMetadata method to Pbm client [51ad97e1]
- Add v4 option to VirtualMachine.WaitForIP [d124bece]
- Add support for the cis session get method [a5a429c0]
- Don't limit library.Finder to local libraries [4513735f]
- examples: add ExampleVirtualMachine_Reconfigure [cad9a8e2]
- govc: add guest.df command [3fb02b52]
- Update docs for 0.21 release [a0fef816]
- Content library related cleanups [a38f6e87]
- Fix library AddLibraryItemFileFromURI fingerprint [e4024e9c]
- govc: support library paths in tags.attach commands [fa755779]
- Fixed type bug in global_object_manager Task.QueryResult [5e8cb495]
- govcsim: Support Default UplinkTeamingPolicy in DVSPG [4a67dc73]
- Added missing field in VslmExtendDisk_Task in ExtendDisk method [9da2362d]
- Add Juju to projects using govmomi [91377d77]
- VSLM FCD Global Object Manager client for 6.7U2+ [f9026a84]
- examples: add CustomFieldManager.Set [9495f0d8]
- govcsim: Create datastore as accessible [bb170705]
- Set the InventoryPath of the folder object in DefaultFolder ([#1515](https://github.com/vmware/govmomi/issues/1515)) [35d0b7d3]
- Add govmomi performance example [2d13a357]
- govc: add datastore.info -H flag [2ddfb86b]
- govcsim: Set datastore status as normal [55da29e5]
- Add various govmomi client examples [600e9f7c]
- Add http source support to library.import [5cccd732]
- Goreleaser update for multiple archives [99dd5947]
- govc: add sso.group commands [b3adfff2]
- tags API: add methods for association of multiple tags/objects [5889d091]
- govc: host.vnic.info -json support [b5372b0c]
- Add method that sets vim version to the endpoint service version [9b7688e0]
- Fix tls config in soap.NewServiceClient [fe3488f5]
- govc: add context to LoadX509KeyPair error [4c41c167]
- Support external PSC lookup service [d7430825]
- vcsim: add support to override credentials [774f3800]
- Fix HostNetworkSystem.QueryNetworkHint return value [47c9c070]
- govc: add vm.change hot-add options [910dac72]
- Fix json request tracing [4606125e]
- govc: change logs.download -default=false [746c314e]
- govc: increase guest.ps -X poll interval [05f946d4]
- Add library export support [77cb9df5]
- govc: add -options support to library.deploy [cc10a075]
- vcsim: fix host uuid [ecd7312b]
- vcsim: use stable UUIDs for inventory objects [c25c41c1]
- Fix pbm field type lookup [322d9629]
- vcsim: Press any key to exit [1345eeb8]
- Update examples to use examples.Run method [a4f58ac6]
- Add permanager example [a31db862]
- Fix port signature in REST endpoint token auth [384b1b95]
- Default to running against vcsim in examples [c222666f]
- Add generated vslm types and methods [199e737b]
- vcsim: Update NetworkInfo.Portgroup in simulator [ee14bd3d]
- Format import statement [dc631a2d]
- Fix paths in vsan/methods [f133c9e9]
- Update copy rights [d8e7cc75]
- Add vsan bindings [62412641]
- Support resignature of vmfs snapshots ([#1442](https://github.com/vmware/govmomi/issues/1442)) [fc3f0e9d]
- govc: rename vcenter.deploy to library.deploy [fe372923]
- govc: move library.item.update commands to library.session [436d7a04]
- govc: consolidate library commands [e6514757]
- govc: export Archive Path field [f8249ded]
- vcsa: bump to 6.7u2 [8a823c52]
- vcsim: remove httptest.serve flag [5b5eaa70]
- Update to vSphere 6.7u2 API [466dc5b2]
- Add error check to VirtualMachine.WaitForNetIP [e9f80882]
- Add ovftool support [5611aaa2]
- vcsim: add library.deploy support [20c1873e]
- vcsim: add ovf manager [0b1ad552]
- govc: add vm.change vpmc-enabled flag [d2ab2782]
- govc: fix vm.change against templates [e7b801c6]
- govc: fix option.set for int32 type values [8a856429]
- Typo and->an [9155093e]
- govc: add datastore.maintenance.{enter,exit} commands [81391309]
- Add support to reconcile FCD datastore inventory [1a857b94]
- govc: FCD workarounds [18cb9142]
- Fix staticcheck issues value of `XXX` is never used [499a8828]
- govc: add datastore.cluster.info Description [665affe5]
- Add error check for deferred functions [546e8897]
- Fix bug with multiple tags in category [367c8743]
- govc: add permission.remove -f flag [7b7f2013]
- Makefile: Fix govet target using go1.12 [87bc0c85]
- travis.yml: Update from golang 1.11 to 1.12 [791e5434]
- travis.yml: Update from Ubuntu Trusty to Xenial [a86a42a2]
- Report local Datastore back as type OTHER [d92ee75e]
- vcsim: fork httptest server package [6684016f]
- vcsim: add content library support [48c1e0a5]
- Make PostEvent TaskInfo param optional [69faa2de]
- Omit namespace tag in generated method body response types [608ad29f]
- Fix codespell issues [a7c03228]
- Fix a race in NewServer(). [728e77db]
- vcsim: set guest.toolsRunningStatus property [8543ea4f]
- Fix elseif gocritic issues [e3143407]
- Fix gocritic emptyStringTest issues [89b53312]
- Fix some trivial gocritic issues [63ba9232]
- simulator/host_datastore_browser.go: remove commented out code [0b8d0ee7]
- Fix some staticcheck issues [6c17d66c]
- Fix some gosimple issues [d45b5f34]
- Correct the year in the govc changelog [90e501a6]
- Update XDR to use fork [8082a261]
- govc/USAGE.md: Update documentation [e94ec246]
- snapshot.tree: Show snapshots description [3fde3319]
- Fix year in changelog [1d6f743b]
- support customize vm folder in ovf deploy [39b2c871]
- Use rest.Client for library uploads [3ad203d3]
- lib/finder: Support filenames with "/" [5d24c38c]
- govc library: use govc/flags for Datastore and ResourcePool [087f09f9]
- Remove nested progress.Tee usage [d1a7f491]
- govc/vm/*: Fix some gosec Errors unhandled issues [7312711e]
- vcsim/*: Fix Errors unhandled issues [88601bb7]
- session/*: Fix Errors unhandled issues [61d04b46]
- vmdk/*: Fix gosec Errors unhandled issues [f9a22349]
- Fix gosec Expect directory permissions to be 0750 or less issues [ca9b71a9]
- Fix gosec potential file inclusion via variable issues [6083e891]
- Build changes needed for content library [38091bf8]
- Content library additions/finder [885d4b44]
- Add support for content library [3fb72d1a]
- Fix API Version check. [64f2a5ea]
- govc/*: Fix some staticcheck issues [718331e3]
- Fix all staticcheck "error strings should not be capitalized" issues [ba7923ae]
- simulator/*: Fix some staticcheck issues [ed32a917]
- govc/vm/*: Fix staticcheck issues [f71d4efb]
- vim25/*: Fix staticcheck issues [3d77e2b1]
- .gitignore: add editor files *~ [d711005a]
- Fix [#1173](https://github.com/vmware/govmomi/issues/1173) [43ff04f1]
- Go Mod Support [562aa0db]

<a name="v0.20.3"></a>
## [Release v0.20.3](https://github.com/vmware/govmomi/compare/prerelease-v0.21.0-58-g8d28646...v0.20.3)

> Release Date: 2019-10-08

### 📖 Commits

- Fix tls config in soap.NewServiceClient [fdd27786]
- Set the InventoryPath of the folder object in DefaultFolder ([#1515](https://github.com/vmware/govmomi/issues/1515)) [bd9cfd18]
- Fix port signature in REST endpoint token auth [4514987f]

<a name="prerelease-v0.21.0-58-g8d28646"></a>
## [Release prerelease-v0.21.0-58-g8d28646](https://github.com/vmware/govmomi/compare/v0.21.0...prerelease-v0.21.0-58-g8d28646)

> Release Date: 2019-09-08

### 💫 `govc` (CLI)

- fix datacenter.info against nested folders
- add vm.change -latency flag
- validate moref argument
- add guest.df command

### 💫 `vcsim` (Simulator)

- rename Example to Run
- add endpoint registration mechanism
- add PlaceVm support ([#1589](https://github.com/vmware/govmomi/issues/1589))
- DefaultDatastoreID is optional in library deploy

### 📖 Commits

- examples: add VirtualMachine.CreateSnapshot [1828eee9]
- vapi: return info with current session query [ca3763e7]
- vcsim: rename Example to Run [f962095f]
- vcsim: add endpoint registration mechanism [43d69860]
- govc: fix datacenter.info against nested folders [1b159e27]
- vcsim: add PlaceVm support ([#1589](https://github.com/vmware/govmomi/issues/1589)) [c183577b]
- Add ResourcePool.Owner method [3e71d6be]
- vcsim: DefaultDatastoreID is optional in library deploy [b17f3a51]
- Update generated code to vSphere 6.7u3 [68980704]
- Add VirtualMachine.QueryChangedDiskAreas(). [7416741c]
- Content library: support library ID in Finder [8ef87890]
- Add option to propagate MissingSet faults in property.WaitForUpdates [e373feb8]
- examples: fix flag parsing [6ff7040e]
- govc: add vm.change -latency flag [149ba7ad]
- govc: validate moref argument [c35a532d]
- Add content library subscription support [54df157b]
- Fix deadlock for keep alive handlers that attempt log in [b86466b7]
- CNS go bindings [9ad64557]
- Add simulator.Model.Run example [9de3b854]
- Include url in Client.Download error [4285b614]
- vcsa: update to 6.7 U3 [caf0b6b3]
- Update vcsim Readme.md [7ac56b64]
- Update README.md [48ef35df]
- Use gnu xargs in bats tests on Darwin [a40837d8]
- Add FetchCapabilityMetadata method to Pbm client [51ad97e1]
- Add v4 option to VirtualMachine.WaitForIP [d124bece]
- Add support for the cis session get method [a5a429c0]
- Don't limit library.Finder to local libraries [4513735f]
- examples: add ExampleVirtualMachine_Reconfigure [cad9a8e2]
- govc: add guest.df command [3fb02b52]

<a name="v0.21.0"></a>
## [Release v0.21.0](https://github.com/vmware/govmomi/compare/v0.20.2...v0.21.0)

> Release Date: 2019-07-24

### 💫 `govc` (CLI)

- support library paths in tags.attach commands
- add datastore.info -H flag
- add sso.group commands
- host.vnic.info -json support
- add context to LoadX509KeyPair error
- add vm.change hot-add options
- change logs.download -default=false
- increase guest.ps -X poll interval
- add -options support to library.deploy
- rename vcenter.deploy to library.deploy
- move library.item.update commands to library.session
- consolidate library commands
- export Archive Path field
- add vm.change vpmc-enabled flag
- fix vm.change against templates
- fix option.set for int32 type values
- add datastore.maintenance.{enter,exit} commands
- FCD workarounds
- add datastore.cluster.info Description
- add permission.remove -f flag

### 💫 `vcsim` (Simulator)

- add support to override credentials
- fix host uuid
- use stable UUIDs for inventory objects
- Press any key to exit
- Update NetworkInfo.Portgroup in simulator
- remove httptest.serve flag
- add library.deploy support
- add ovf manager
- fork httptest server package
- add content library support
- set guest.toolsRunningStatus property

### 📖 Commits

- Update docs for 0.21 release [a0fef816]
- Content library related cleanups [a38f6e87]
- Fix library AddLibraryItemFileFromURI fingerprint [e4024e9c]
- govc: support library paths in tags.attach commands [fa755779]
- Fixed type bug in global_object_manager Task.QueryResult [5e8cb495]
- govcsim: Support Default UplinkTeamingPolicy in DVSPG [4a67dc73]
- Added missing field in VslmExtendDisk_Task in ExtendDisk method [9da2362d]
- Add Juju to projects using govmomi [91377d77]
- VSLM FCD Global Object Manager client for 6.7U2+ [f9026a84]
- examples: add CustomFieldManager.Set [9495f0d8]
- govcsim: Create datastore as accessible [bb170705]
- Set the InventoryPath of the folder object in DefaultFolder ([#1515](https://github.com/vmware/govmomi/issues/1515)) [35d0b7d3]
- Add govmomi performance example [2d13a357]
- govc: add datastore.info -H flag [2ddfb86b]
- govcsim: Set datastore status as normal [55da29e5]
- Add various govmomi client examples [600e9f7c]
- Add http source support to library.import [5cccd732]
- Goreleaser update for multiple archives [99dd5947]
- govc: add sso.group commands [b3adfff2]
- tags API: add methods for association of multiple tags/objects [5889d091]
- govc: host.vnic.info -json support [b5372b0c]
- Add method that sets vim version to the endpoint service version [9b7688e0]
- Fix tls config in soap.NewServiceClient [fe3488f5]
- govc: add context to LoadX509KeyPair error [4c41c167]
- Support external PSC lookup service [d7430825]
- vcsim: add support to override credentials [774f3800]
- Fix HostNetworkSystem.QueryNetworkHint return value [47c9c070]
- govc: add vm.change hot-add options [910dac72]
- Fix json request tracing [4606125e]
- govc: change logs.download -default=false [746c314e]
- govc: increase guest.ps -X poll interval [05f946d4]
- Add library export support [77cb9df5]
- govc: add -options support to library.deploy [cc10a075]
- vcsim: fix host uuid [ecd7312b]
- vcsim: use stable UUIDs for inventory objects [c25c41c1]
- Fix pbm field type lookup [322d9629]
- vcsim: Press any key to exit [1345eeb8]
- Update examples to use examples.Run method [a4f58ac6]
- Add permanager example [a31db862]
- Fix port signature in REST endpoint token auth [384b1b95]
- Default to running against vcsim in examples [c222666f]
- Add generated vslm types and methods [199e737b]
- vcsim: Update NetworkInfo.Portgroup in simulator [ee14bd3d]
- Format import statement [dc631a2d]
- Fix paths in vsan/methods [f133c9e9]
- Update copy rights [d8e7cc75]
- Add vsan bindings [62412641]
- Support resignature of vmfs snapshots ([#1442](https://github.com/vmware/govmomi/issues/1442)) [fc3f0e9d]
- govc: rename vcenter.deploy to library.deploy [fe372923]
- govc: move library.item.update commands to library.session [436d7a04]
- govc: consolidate library commands [e6514757]
- govc: export Archive Path field [f8249ded]
- vcsa: bump to 6.7u2 [8a823c52]
- vcsim: remove httptest.serve flag [5b5eaa70]
- Update to vSphere 6.7u2 API [466dc5b2]
- Add error check to VirtualMachine.WaitForNetIP [e9f80882]
- Add ovftool support [5611aaa2]
- vcsim: add library.deploy support [20c1873e]
- vcsim: add ovf manager [0b1ad552]
- govc: add vm.change vpmc-enabled flag [d2ab2782]
- govc: fix vm.change against templates [e7b801c6]
- govc: fix option.set for int32 type values [8a856429]
- Typo and->an [9155093e]
- govc: add datastore.maintenance.{enter,exit} commands [81391309]
- Add support to reconcile FCD datastore inventory [1a857b94]
- govc: FCD workarounds [18cb9142]
- Fix staticcheck issues value of `XXX` is never used [499a8828]
- govc: add datastore.cluster.info Description [665affe5]
- Add error check for deferred functions [546e8897]
- Fix bug with multiple tags in category [367c8743]
- govc: add permission.remove -f flag [7b7f2013]
- Makefile: Fix govet target using go1.12 [87bc0c85]
- travis.yml: Update from golang 1.11 to 1.12 [791e5434]
- travis.yml: Update from Ubuntu Trusty to Xenial [a86a42a2]
- Report local Datastore back as type OTHER [d92ee75e]
- vcsim: fork httptest server package [6684016f]
- vcsim: add content library support [48c1e0a5]
- Make PostEvent TaskInfo param optional [69faa2de]
- Omit namespace tag in generated method body response types [608ad29f]
- Fix codespell issues [a7c03228]
- Fix a race in NewServer(). [728e77db]
- vcsim: set guest.toolsRunningStatus property [8543ea4f]
- Fix elseif gocritic issues [e3143407]
- Fix gocritic emptyStringTest issues [89b53312]
- Fix some trivial gocritic issues [63ba9232]
- simulator/host_datastore_browser.go: remove commented out code [0b8d0ee7]
- Fix some staticcheck issues [6c17d66c]
- Fix some gosimple issues [d45b5f34]
- Correct the year in the govc changelog [90e501a6]
- Update XDR to use fork [8082a261]
- govc/USAGE.md: Update documentation [e94ec246]
- snapshot.tree: Show snapshots description [3fde3319]
- Fix year in changelog [1d6f743b]
- support customize vm folder in ovf deploy [39b2c871]
- Use rest.Client for library uploads [3ad203d3]
- lib/finder: Support filenames with "/" [5d24c38c]
- govc library: use govc/flags for Datastore and ResourcePool [087f09f9]
- Remove nested progress.Tee usage [d1a7f491]
- govc/vm/*: Fix some gosec Errors unhandled issues [7312711e]
- vcsim/*: Fix Errors unhandled issues [88601bb7]
- session/*: Fix Errors unhandled issues [61d04b46]
- vmdk/*: Fix gosec Errors unhandled issues [f9a22349]
- Fix gosec Expect directory permissions to be 0750 or less issues [ca9b71a9]
- Fix gosec potential file inclusion via variable issues [6083e891]
- Build changes needed for content library [38091bf8]
- Content library additions/finder [885d4b44]
- Add support for content library [3fb72d1a]
- Fix API Version check. [64f2a5ea]
- govc/*: Fix some staticcheck issues [718331e3]
- Fix all staticcheck "error strings should not be capitalized" issues [ba7923ae]
- simulator/*: Fix some staticcheck issues [ed32a917]
- govc/vm/*: Fix staticcheck issues [f71d4efb]
- vim25/*: Fix staticcheck issues [3d77e2b1]
- .gitignore: add editor files *~ [d711005a]
- Fix [#1173](https://github.com/vmware/govmomi/issues/1173) [43ff04f1]
- Go Mod Support [562aa0db]

<a name="v0.20.2"></a>
## [Release v0.20.2](https://github.com/vmware/govmomi/compare/v0.20.1...v0.20.2)

> Release Date: 2019-07-03

### 📖 Commits

- Set the InventoryPath of the folder object in DefaultFolder ([#1515](https://github.com/vmware/govmomi/issues/1515)) [bd9cfd18]

<a name="v0.20.1"></a>
## [Release v0.20.1](https://github.com/vmware/govmomi/compare/v0.20.0...v0.20.1)

> Release Date: 2019-05-20

### 📖 Commits

- Fix port signature in REST endpoint token auth [4514987f]

<a name="v0.20.0"></a>
## [Release v0.20.0](https://github.com/vmware/govmomi/compare/v0.19.0...v0.20.0)

> Release Date: 2019-02-06

### 💫 `govc` (CLI)

- fix object.collect error for multiple objects with same path
- add device name match support to device.ls and device.remove
- add vm.disk.attach -mode flag
- add category option to relevant tags commands
- add vm.create -version option
- fields.set can now add missing fields
- add fields.info command

### 💫 `vcsim` (Simulator)

- require authentication in vapi simulator
- Resolve issue making device changes on clone (resolves [#1355](https://github.com/vmware/govmomi/issues/1355))
- fix SearchDatastore task info entity
- add EnvironmentBrowser support
- avoid zero IP address in GOVC_URL output
- avoid panic when event template is not defined
- implement RefreshStorageInfo method for virtual machine
- configure HostSystem port
- datastore.upload now creates missing directories in destination path.
- add option to run container as vm
- add SessionIsActive support
- fix fault detail encoding
- support base types in property filter
- PropertyCollector should not require PathSet
- allow '.' in vm name
- populate VM guest.net field
- add SearchIndex FindByDnsName support
- correct property update in RemoveSnapshotTask
- update VM snapshot methods to change VM properties with UpdateObject
- support setting vm fields via extraConfig
- update VM configureDevices method to change VM properties with UpdateObject
- update VM device add operation - stricter key generation, new InvalidDeviceSpec condition
- add PBM support
- put VM into registry earlier during CreateVM
- add datastore access check for vm host placement
- add task_manager description property templates
- fix defaults when generating vmdk paths
- fix custom_fields_manager test
- replace HostSystem template IP with vcsim listen address
- Change CustomFieldsManager SetField to use ctx.WithLock and add InvalidArgument fault check.
- update DVS methods to use UpdateObject instead of setting fields directly
- add vslm support
- add setCustomValue support
- add fault message to PropertyCollector RetrieveProperties
- add HistoryCollector scrollable view support

### 📖 Commits

- Fix for govc/build.sh wrong dir [da7af247]
- Update docs for 0.20 release [90a863be]
- vcsim: require authentication in vapi simulator [957ef0f7]
- vcsim: Resolve issue making device changes on clone (resolves [#1355](https://github.com/vmware/govmomi/issues/1355)) [32148187]
- Use path id for tag-association requests [a7563c4d]
- vcsim: fix SearchDatastore task info entity [cbb4abc9]
- vcsim: add EnvironmentBrowser support [2682c021]
- vcsim: avoid zero IP address in GOVC_URL output [3b9a4c9f]
- Add 2x blog posts about vcsim [b261f25d]
- vcsim: avoid panic when event template is not defined [1921f73a]
- govc: fix object.collect error for multiple objects with same path [308dbf99]
- vcsim: implement RefreshStorageInfo method for virtual machine [d79013aa]
- vcsim: configure HostSystem port [69dfdd77]
- Fix of the missing http body close under soap client upload [4f50681f]
- vcsim: datastore.upload now creates missing directories in destination path. [bba50b40]
- Fixed 64-bit aligment issues with atomic counters [8ac7c5a8]
- fix device.info Write output [7ca12ea2]
- device.ls -json doesn't work for now [3a82237c]
- ssoadmin:create local group and add users to group ([#1327](https://github.com/vmware/govmomi/issues/1327)) [86f4ba29]
- Format with latest version of goimports [2d8ef2c6]
- govc: add device name match support to device.ls and device.remove [4635c1cc]
- Updated the examples for the correct format [d7857a13]
- Updated to reflect PR feedback [71e19136]
- vcsim: add option to run container as vm [d2506759]
- Added string support [61b7fe3e]
- Initial Support for PutUsbScanCodes [a72a4c42]
- vcsim: add SessionIsActive support [47284860]
- vcsim: fix fault detail encoding [c5ee00bf]
- Summary of changes:  1. Changing the pbm client's path as java client is expecting /pbm.  2. Added PbmRetrieveServiceContent method in the unauthorized list. [aaf83275]
- govc: add vm.disk.attach -mode flag [c36eb50f]
- vcsim: support base types in property filter [1284300c]
- vcsim: PropertyCollector should not require PathSet [25ae5c67]
- govc: add category option to relevant tags commands [b234cdbc]
- Makefiles for govc/vcsim; updates  govc/build.sh [138f30f8]
- vcsim: allow '.' in vm name [4f1c89e5]
- govc: add vm.create -version option [afe5f42d]
- vcsim: populate VM guest.net field [b8c04142]
- vcsim: add SearchIndex FindByDnsName support [223b2a2a]
- vcsim: correct property update in RemoveSnapshotTask [b26e10f0]
- vcsim: update VM snapshot methods to change VM properties with UpdateObject [693f3fb6]
- build: Refactored Travis-CI to use containers [e5948f44]
- vcsim: support setting vm fields via extraConfig [06e13bbe]
- Allow pointer values in mo.ApplyPropertyChange [651d4881]
- Tags support for First Class Disks [546a7df6]
- vcsim: update VM configureDevices method to change VM properties with UpdateObject [a4330365]
- vcsim: update VM device add operation - stricter key generation, new InvalidDeviceSpec condition [5f8acb7a]
- Merge branch 'master' into fields-info [86375ceb]
- Update govc/fields/add.go [bf962f18]
- Update govc/fields/add.go [98575e0c]
- govc: fields.set can now add missing fields [b733db99]
- govc: add fields.info command [cad627a6]
- vm.power: Make waiting for op completion optional [ed2a4cff]
- vcsim: add PBM support [846ae27a]
- vcsim: put VM into registry earlier during CreateVM [d41d18aa]
- Datastore Cluster placement support for First Class Disks [1926071e]
- vcsim: add datastore access check for vm host placement [89b4c2ce]
- vcsim: add task_manager description property templates [f9f9938e]
- vcsim: fix defaults when generating vmdk paths [9bb5bde2]
- vcsim: fix custom_fields_manager test [0b650fd3]
- vcsim: replace HostSystem template IP with vcsim listen address [588bc224]
- vcsim: Change CustomFieldsManager SetField to use ctx.WithLock and add InvalidArgument fault check. [7066f8dc]
- Display category name instead of ID in govc tags.info [ef517cae]
- goimports updates [d69c9787]
- vcsim: update DVS methods to use UpdateObject instead of setting fields directly [fe070811]
- vcsim: add vslm support [03939cce]
- Add vslm package and govc disk commands [accb2863]
- [doc] add an example for cpu and memory hotplug [478ebae6]
- vcsim: add setCustomValue support [c02efc3d]
- goimports updates [c3c79d16]
- vcsa: bump to 6.7.0 U1 [ce71b6c2]
- vcsim: add fault message to PropertyCollector RetrieveProperties [94804159]
- Removed NewWithDelay (not needed anymore) [1ad0d87d]
- Updated documentation [5900feef]
- Added delay functionality [5a87902b]
- Add LoginByToken to session KeepAliveHandler [c0518fd2]
- Update Ansible link in README [e0736431]
- vcsim: add HistoryCollector scrollable view support [36035f5b]
- Move govc tags rest.Client helper to ClientFlag [bc2636fe]
- Add SSO support for vAPI [54a181af]
- replace * by client's host+port [8817c27b]
- change hostname only if set to * and still set thumbprint [ac898b50]
- replace hostname only if unset [7a5cc6b7]

<a name="v0.19.0"></a>
## [Release v0.19.0](https://github.com/vmware/govmomi/compare/v0.18.0...v0.19.0)

> Release Date: 2018-09-30

### 💫 `govc` (CLI)

- fix test case for new cluster.rule.info command
- add new command cluster.rule.info

### 💫 `vcsim` (Simulator)

- add dvpg networks to HostSystem.Parent
- add support for tags API
- Logout should not unregister PropertyCollector singleton
- add ResetVM and SuspendVM support
- add support for PropertyCollector incremental updates
- do not include DVS in HostSystem.Network

### 📖 Commits

- Update docs for 0.19 release [3617f28d]
- vcsa: bump to 6.7.0d [4316838a]
- Added PerformanceManager simulator [64d875b9]
- vcsim: add dvpg networks to HostSystem.Parent [f3260968]
- Allowing the use of STS for exchanging tokens [862da065]
- Handle empty file name in import.spec [83ce863a]
- Bump travis golang version from 1.10 to 1.11 [a99f702d]
- Clean up unit test messaging [e4e8e2d6]
- Run goimports on go source files [8e04e3c7]
- Add mailmap for bruceadowns [2431ae00]
- Updates per dep ensure integration [e805b4ea]
- Add ignore of intellij project settings directory [70589fb6]
- Print action for dvs security groups [d114fa69]
- fix double err check [d458266a]
- remove providerSummary cache [3f0e0aa7]
- Avoid use of Finder all param in govc [cf9c16c4]
- Print DVS rules for dvportgroup [c4face4f]
- Finalize tags API [91a33dd4]
- README: Fix path to LICENSE.txt file [7d54bf9f]
- vcsim: add support for tags API [17352fce]
- vcsim: Logout should not unregister PropertyCollector singleton [c29d4b12]
- Fix format in test [8bda0ee1]
- Add test for WaitOption.MaxWaitSeconds == 0 behaviour in simulator [8be5207c]
- Fix the WaitOption.MaxWaitSeconds == 0 behaviour in simulator [900e1a35]
- vcsa: bump to 6.7.0c release [056ad0d4]
- govc: fix test case for new cluster.rule.info command [6b4a62b1]
- govc: add new command cluster.rule.info [1350eea6]
- add output in cluster.rule.ls -name for ClusterVmHostRuleInfo and ClusterDependencyRuleInfo rules, add -l Option to cluster.rule.ls [a05cd4b0]
- vcsim: add ResetVM and SuspendVM support [11fb0d58]
- Add ability to move multiple hosts into a cluster [3e6b2d6e]
- Add method to move host into cluster [e9f9920f]
- vcsim: add support for PropertyCollector incremental updates [39e6592d]
- Add testing support for govc tags commands [b7c270c6]
- vcsim: do not include DVS in HostSystem.Network [619fbe28]
- show rule details for ClusterVmHostRuleInfo rules in cluster.rule.ls [6b6060dc]
- Use govc find instead of ls to assign licenses [0c28a25d]
- Only test with Go 1.10 on Travis CI [c1377063]
- Avoid panic if fault detail is nil [4cfadda5]
- Upgrade for govc tags commands [d06874e1]
- Better documentation for VirtualMachine.UUID [fdfaec9c]
- Add UUID helper for VirtualMachine [e1285a03]
- Complete tags management APIs ([#1162](https://github.com/vmware/govmomi/issues/1162)) [919b728c]
- vcsa: bump to 6.7.0a release [b3251638]
- Optionally check root CAs for validity ([#1154](https://github.com/vmware/govmomi/issues/1154)) [a1fbb6ef]
- Fixed govc host.info logical CPU count [add38bed]
- Tags Categories cmd available  ([#1150](https://github.com/vmware/govmomi/issues/1150)) [1ddfb011]
- default MarkAsTemplate to false in import spec [83ae35fb]
- add option to mark VM as template on OVX import [49f0dea7]
- example: uniform unit for host memory [1f9e19f4]
- fix example output. [4cfd1376]

<a name="v0.18.0"></a>
## [Release v0.18.0](https://github.com/vmware/govmomi/compare/v0.17.1...v0.18.0)

> Release Date: 2018-05-24

### 💫 `govc` (CLI)

- import.ovf pool flag should be optional if host is specified
- avoid Login() attempt if username is not set
- add json support to find command
- fix host.esxcli error handling

### 💫 `vcsim` (Simulator)

- add STS simulator
- use VirtualDisk CapacityInKB for device summary
- add property collector field type mapping for integer arrays

### 📖 Commits

- Update docs for 0.18 release [e4b69fab]
- Bump versions [1dbfb317]
- govc: import.ovf pool flag should be optional if host is specified [b841ae01]
- Add -sharing option to vm.disk.create and vm.disk.attach [96a905c1]
- Add VirtualDiskManager wrapper to set UUID [4b4e2aaa]
- adjust datastore size when vm is added or updated or deleted [40a565b3]
- update datastore capacity and free space when it is started [7f6479ba]
- Avoid recursive root path search in govc find command [76dfefd3]
- Change key name according to Datacenter object [623c7fa9]
- added check for `InstanceUuid` when `VmSearch` is true in `FindByUuid` [24d0cf1b]
- Issue token if needed for govc sso commands [25fc474c]
- Fixed leading "/" requirement in FindByInventoryPath [822fd1c0]
- Add devbox scripts [59d9f6a0]
- Add -U option to sso.service.ls [fd45d81c]
- govc: avoid Login() attempt if username is not set [f5c84b98]
- vcsim: add STS simulator [8a5438b0]
- Fix govc vm.clone -annotation flag [93f7fbbd]
- save CapacityInKB in thousand delimited format [bcff5383]
- Avoid possible panic in portgroup EthernetCardBackingInfo [db12d4cb]
- Add STS support for token renewal [d120efcb]
- Add vmxnet2, pcnet32 and sriov to VirtualDeviceList.EthernetCardTypes [76b1ceaf]
- vcsim: use VirtualDisk CapacityInKB for device summary [c0337740]
- vcsim: add property collector field type mapping for integer arrays [3d7fbac2]
- Finder.DefaultHostSystem should find hosts in nested folders [42b30bb6]
- Avoid property.Filter matching against unset properties [b8323d6b]
- Update to vSphere 6.7 API [64788667]
- Bump vCenter and ESXi builds to the latest release [d3ae3004]
- Add ssoadmin client and commands [098fc449]
- vm.Snapshot should be 'nil' instead of an empty 'vim.vm.SnapshotInfo' when there are no snapshots [80a9c20e]
- added failing tests for when vm.Snapshot should / shouldn't be 'nil' [1b1b428e]
- Refactor LoginExtensionByCertificate tunnel usage [a34ab4ba]
- Lookup Service support [5b36033f]
- add empty fields, but don't return them in the case of 'RetrievePropertiesEx' [3f07eb74]
- added failing test case for issue 1061 [05bdabe0]
- SAML token authentication support [903e8644]
- govc: add json support to find command [d91fcbf4]
- govc: fix host.esxcli error handling [ba2d2323]
- Dep Support [ff687746]
- Add -firmware parameter to 'govc vm.create' with values bios|efi [5f701460]

<a name="v0.17.1"></a>
## [Release v0.17.1](https://github.com/vmware/govmomi/compare/v0.17.0...v0.17.1)

> Release Date: 2018-03-19

### 💫 `vcsim` (Simulator)

- add Destroy method for Folder and Datacenter types
- add EventManager.QueryEvents

### 📖 Commits

- govc release 0.17.1 [123ed177]
- Avoid possible panic in QueryVirtualDiskInfo [24d88451]
- Add goreleaser to automate release process [82129fb7]
- Fix dvs.portgroup.info filtering [ce88b296]
- vcsim: add Destroy method for Folder and Datacenter types [0502ee9b]
- In progress.Reader emit final report on EOF. [1620160d]
- vcsim: add EventManager.QueryEvents [0636dc8c]

<a name="v0.17.0"></a>
## [Release v0.17.0](https://github.com/vmware/govmomi/compare/v0.16.0...v0.17.0)

> Release Date: 2018-02-28

### 💫 `govc` (CLI)

- fix vm.clone to use -net flag when source does not have a NIC
- object.collect support for raw filters
- fix host.info CPU usage
- add -cluster flag to license.assign command
- allow columns in guest login password ([#972](https://github.com/vmware/govmomi/issues/972))

### 💫 `vcsim` (Simulator)

- add simulator.Datastore type
- set VirtualMachine summary.config.instanceUuid
- update HostSystem.Summary.Host reference
- add EventManager support
- stats related fixes
- avoid data races
- respect VirtualDeviceConfigSpec FileOperation
- avoid keeping the VM log file open
- add UpdateOptions support
- add session support
- Add VM.MarkAsTemplate support
- more input spec honored in ReConfig VM
- Initialize VM fields properly
- Honor the input spec in ReConfig VM
- Add HostLocalAccountManager
- workaround xml ns issue with pyvsphere ([#958](https://github.com/vmware/govmomi/issues/958))
- add MakeDirectoryResponse ([#938](https://github.com/vmware/govmomi/issues/938))
- copy RoleList for AuthorizationManager ([#932](https://github.com/vmware/govmomi/issues/932))
- apply vm spec NumCoresPerSocket ([#930](https://github.com/vmware/govmomi/issues/930))
- Configure dvs with the dvs config spec
- Add VirtualMachine guest ID validation ([#921](https://github.com/vmware/govmomi/issues/921))
- add QueryVirtualDiskUuid ([#920](https://github.com/vmware/govmomi/issues/920))
- update ServiceContent to 6.5 ([#917](https://github.com/vmware/govmomi/issues/917))

### 📖 Commits

- govc release 0.17 [1d63da8d]
- Print Table of Contents in usage.md Found good example of toc using markdown here: https://stackoverflow.com/a/448929/1572363 [3017acf8]
- Fix typo [ce54fe2c]
- Implement Destroy task for HostSystem [201fc601]
- Init PortKeys in DistributedVirtualPortgroup [92ce4244]
- Avoid json encoding error in Go 1.10 [795f2cc7]
- Add 'Type' field to device.info -json output [e805389e]
- Use VirtualDiskManager in datastore cp and mv commands [d622f149]
- object: Return correct helper object for OpaqueNetwork [f219bf3b]
- govc: fix vm.clone to use -net flag when source does not have a NIC [29498644]
- Fix build on Windows [43c95b21]
- Fix session persistence in session.login command [38124002]
- Add support for Datacenter.PowerOnMultiVM [144bb1cf]
- vcsim: add simulator.Datastore type [d2ba47d6]
- vcsim: set VirtualMachine summary.config.instanceUuid [937998a1]
- vcsim: update HostSystem.Summary.Host reference [1c76c63d]
- govc: object.collect support for raw filters [d12b8f25]
- vcsim: add EventManager support [274f3d63]
- vcsim: stats related fixes [cc21a5ab]
- Fix broken datastore link in VM [2d30cde3]
- Several context changes: [54b160b2]
- Leverage contexts in http uploads [f643f0ae]
- vcsim: avoid data races [fa2bee10]
- Remove omitempty tag from AffinitySet field [29bd00ec]
- vcsim: respect VirtualDeviceConfigSpec FileOperation [ca6f5d1d]
- vcsim: avoid keeping the VM log file open [7811dfce]
- govc: fix host.info CPU usage [6cb9fef8]
- govc: add -cluster flag to license.assign command [5786e7d2]
- Add datastore.disk.cp command [63c86f29]
- vcsim: add UpdateOptions support [828ce5ec]
- Bump vcsa scripts to use 6.5U1 EP5 [a13ad164]
- Add CloneSession support to govc and vcsim [c447244d]
- vcsim: add session support [d03f38fa]
- Added AttachScsiLun function ([#987](https://github.com/vmware/govmomi/issues/987)) [44e8d85e]
- vcsim: Add VM.MarkAsTemplate support [a3c9ed2b]
- Add cluster vm override commands ([#977](https://github.com/vmware/govmomi/issues/977)) [3f8349f3]
- Add option to filter events by type ([#976](https://github.com/vmware/govmomi/issues/976)) [91fbd1f7]
- User server clock in session.ls ([#973](https://github.com/vmware/govmomi/issues/973)) [1d8b92d9]
- vcsim: more input spec honored in ReConfig VM [50735461]
- vcsim: Initialize VM fields properly [638d972b]
- Add '-rescan-vmfs' option to host.storage.info ([#966](https://github.com/vmware/govmomi/issues/966)) [2892ed50]
- govc: allow columns in guest login password ([#972](https://github.com/vmware/govmomi/issues/972)) [d4ee331c]
- Use IsFileNotFound helper in Datastore.Stat ([#969](https://github.com/vmware/govmomi/issues/969)) [e15ff586]
- vcsim: Honor the input spec in ReConfig VM [aa0382c1]
- Hook AccountManager to UserDirectory [465bd948]
- Destroy event history collectors ([#962](https://github.com/vmware/govmomi/issues/962)) [aef2d795]
- vcsim: Add HostLocalAccountManager [42f9a133]
- vcsim: workaround xml ns issue with pyvsphere ([#958](https://github.com/vmware/govmomi/issues/958)) [76f376a3]
- Ignore AcquireLocalTicket errors ([#955](https://github.com/vmware/govmomi/issues/955)) [a1c49292]
- Add missing dependency in gen script [bb150d50]
- toolbox: validate request offset in ListFiles ([#946](https://github.com/vmware/govmomi/issues/946)) [0eacf959]
- Corrects datastore.disk usage which had not been generated ([#951](https://github.com/vmware/govmomi/issues/951)) [1d6aed22]
- Corrects vm.info usage with required args ([#950](https://github.com/vmware/govmomi/issues/950)) [de717389]
- Add datastore.disk inflate and shrink commands ([#943](https://github.com/vmware/govmomi/issues/943)) [c5ea3fb2]
- Corrects host.shutdown ([#939](https://github.com/vmware/govmomi/issues/939)) [adf4530b]
- vcsim: add MakeDirectoryResponse ([#938](https://github.com/vmware/govmomi/issues/938)) [45c5269b]
- vcsim: copy RoleList for AuthorizationManager ([#932](https://github.com/vmware/govmomi/issues/932)) [b4e77bd2]
- Fix [#933](https://github.com/vmware/govmomi/issues/933) ([#936](https://github.com/vmware/govmomi/issues/936)) [426a675a]
- Add cluster.group and cluster.rule commands ([#928](https://github.com/vmware/govmomi/issues/928)) [3be5f1d9]
- vcsim: apply vm spec NumCoresPerSocket ([#930](https://github.com/vmware/govmomi/issues/930)) [2a8a5168]
- vcsim: Configure dvs with the dvs config spec [3a61d85f]
- CreateChildDisk 6.7 support ([#926](https://github.com/vmware/govmomi/issues/926)) [3b25c720]
- Add VirtualDiskManager.CreateChildDisk ([#925](https://github.com/vmware/govmomi/issues/925)) [933ee3b2]
- vcsim: Add VirtualMachine guest ID validation ([#921](https://github.com/vmware/govmomi/issues/921)) [5f0f4004]
- vcsim: add QueryVirtualDiskUuid ([#920](https://github.com/vmware/govmomi/issues/920)) [ef571547]
- Implemened vm.upgrade operation. ([#918](https://github.com/vmware/govmomi/issues/918)) [0ea3b9bd]
- vcsim: update ServiceContent to 6.5 ([#917](https://github.com/vmware/govmomi/issues/917)) [27229ab7]
- Add support for cpu + mem allocation to vm.change command ([#916](https://github.com/vmware/govmomi/issues/916)) [46c79c93]

<a name="v0.16.0"></a>
## [Release v0.16.0](https://github.com/vmware/govmomi/compare/v0.15.0...v0.16.0)

> Release Date: 2017-11-08

### 💫 `govc` (CLI)

- Fix VM clone when source doesn't have vNics
- add tasks and task.cancel commands
- add reboot option to host.shutdown

### 💫 `vcsim` (Simulator)

- preserve order in QueryIpPools ([#914](https://github.com/vmware/govmomi/issues/914))
- return moref from Task.Run ([#913](https://github.com/vmware/govmomi/issues/913))
- Implement IpPoolManager lifecycle
- add autostart option to power on VMs ([#906](https://github.com/vmware/govmomi/issues/906))
- use soapenv namespace for Fault types
- various property additions
- Generate similar ref value like VC
- Add moref to vm's summary
- validate authz privilege ids
- AuthorizationManager additions
- Add IpPoolManager
- VirtualDisk file backing datastore is optional
- add PerformanceManager
- Implement add/update/remove roles
- Generate device filename in CreateVM
- add AuthorizationManager
- populate vm snapshot fields
- Add UpdateNetworkConfig to HostNetworkSystem
- Implement virtual machine snapshot
- set VirtualDisk backing datastore
- Implement enter/exit maintenance mode
- Implement add/remove license
- add portgroup related operations
- add fields support
- remove use of df program for datastore info
- add FileQuery support to datastore search
- add HostConfigInfo template
- add HostSystem hardware property
- Fix merging of default devices
- Add cdrom and scsi controller to Model VMs

### 📖 Commits

- Doc updates ([#915](https://github.com/vmware/govmomi/issues/915)) [7d879bac]
- vcsim: preserve order in QueryIpPools ([#914](https://github.com/vmware/govmomi/issues/914)) [4543f4b6]
- vcsim: return moref from Task.Run ([#913](https://github.com/vmware/govmomi/issues/913)) [b385183e]
- Remove tls-handshake-timeout flag ([#911](https://github.com/vmware/govmomi/issues/911)) [c8738903]
- vcsim: Implement IpPoolManager lifecycle [e29ab54a]
- Use ProgressLogger for vm.clone command ([#909](https://github.com/vmware/govmomi/issues/909)) [3619c1d9]
- readme: fix formatting of listing ([#908](https://github.com/vmware/govmomi/issues/908)) [13f2aba4]
- vcsim: add autostart option to power on VMs ([#906](https://github.com/vmware/govmomi/issues/906)) [b227a258]
- Add installation procedure in README.md ([#902](https://github.com/vmware/govmomi/issues/902)) [79934451]
- vcsim: use soapenv namespace for Fault types [ecde4a89]
- vcsim: various property additions [b1318195]
- Switch to kr/pretty package for the -dump flag [4d8737c9]
- Couple of fixes for import.spec result [e050b1b6]
- import.spec not to assign deploymentOption [017138ca]
- vcsim: Generate similar ref value like VC [c19ec714]
- govc: Fix VM clone when source doesn't have vNics [0295f1b0]
- vcsim: Add moref to vm's summary [f3046058]
- [govc] Introduce TLSHandshakeTimeout parameter ([#890](https://github.com/vmware/govmomi/issues/890)) [bfed5eea]
- Support import ova/ovf by URL [1c1291ca]
- Remove BaseResourceAllocationInfo [3cb5cc96]
- vcsim: validate authz privilege ids [5f3fba94]
- Add clone methods to session manager [c91b9605]
- vcsim: AuthorizationManager additions [c2caa6d7]
- vcsim: Add IpPoolManager [2cb741f2]
- Updates to vm.clone link + snapshot flags [644c1859]
- Add linked clone and snapshot support to vm.clone [cf624f1a]
- Fix govc events output [024c09fe]
- govc/events: read -json flag and output events as json [d4d94f44]
- Fix vm.register command template flag [24e71ea4]
- Fix object name suffix matching in Finder [5209daf2]
- vcsim: VirtualDisk file backing datastore is optional [a46ab163]
- vcsim: add PerformanceManager [d347175f]
- vcsim: Implement add/update/remove roles [df3763d5]
- Support clearing vm boot order [8d5c1558]
- vcsim: Generate device filename in CreateVM [ed18165d]
- Fix CustomFieldsManager.FindKey method signature [df93050a]
- vcsim: add AuthorizationManager [e8741bf0]
- vcsim: populate vm snapshot fields [8961efc1]
- Add method to find a CustomFieldDef by Key [17fb12a5]
- vscim: Implement UserDirectory [bc395ef0]
- vcsim: Add UpdateNetworkConfig to HostNetworkSystem [add0245e]
- vcsim: Implement virtual machine snapshot [2aa746c6]
- vcsim: set VirtualDisk backing datastore [104ddfb7]
- Add support for VM export [f3f51c58]
- vcsim: Implement enter/exit maintenance mode [505b5c65]
- vcsim: Implement add/remove license [a1f8a328]
- vcsim: add portgroup related operations [585cf5e1]
- vcsim: add fields support [a7e79a7e]
- vim25: Move internal stuff to internal package [e2944227]
- Add support for SOAP request operation ID header [c4cab690]
- vcsim: remove use of df program for datastore info [895573a5]
- Skip version check when using 6.7-dev API [4dd9a518]
- Change optional ResourceAllocationInfo fields to pointers [cc2ed7db]
- Use base type for DVS backing info [3f145230]
- Add vm.console command [df1c3132]
- Fixup recent tasks output [829b3f99]
- Add '-refresh' option to host.storage.info [c4e473af]
- toolbox: avoid race when closing channels on stop [3df440c8]
- toolbox: reset session when invalidated by the vmx [badad9a1]
- Include "Name" in device.info -json [a1a96c8f]
- vcsim: add FileQuery support to datastore search [defe810c]
- Default vm.migrate pool to the host pool [93f62ef7]
- vcsim: add HostConfigInfo template [5fcca79e]
- govc: add tasks and task.cancel commands [4fea6863]
- Use ovf to import vmdk [596e51a0]
- vcsim: add HostSystem hardware property [920a70c1]
- Add info about maintenance mode in host.info [9e2f8a78]
- Avoid panic if ova import task is canceled [78f3fc19]
- toolbox: default to tar format for directory archives [11827c7a]
- toolbox: make gzip optional for directory archive transfer [8811f9bf]
- toolbox: avoid blocking the RPC channel when transferring process IO [9703fe19]
- Add view and filter support to object.collect command [d6f60304]
- Tolerate repeated Close for file follower [3527a5f8]
- govc: add reboot option to host.shutdown [ddd32366]
- toolbox: use host management IP for guest file transfer [4d9061ac]
- toolbox: add Client Upload and Download methods [7d956b6b]
- toolbox: support single file download via archive handler [c7111c63]
- Use vcsim in bats tests [ebb77d7c]
- vCenter cluster testbed automation [4bb89668]
- toolbox: SendGuestInfo before the vmx asks us to [ad960e95]
- toolbox: update vmw-guestinfo [bdea7ff3]
- toolbox: remove receiver from DefaultStartCommand [51d12609]
- Add host thumbprint for use with guest file transfer [114329fc]
- Add FindByUuid method for Virtual Machine [5083a277]
- toolbox: map exec.ErrNotFound to vix.FileNotFound [e1ab84af]
- toolbox: pass URL to ArchiveHandler Read/Write methods [d1091087]
- toolbox: make directory archive read/write customizable [cddc353c]
- toolbox: add http and exec round trippers [ba6720ce]
- Handle object names containing a '/' [b35abbc8]
- toolbox: fix ListFiles when given a symlink [ac4891fb]
- Minor correction in README.md [60a6510f]
- toolbox: support transferring /proc files from guest [0c583dbc]
- vcsim: Fix merging of default devices [0833484e]
- Move toolbox from vmware/vic to govmomi [c9aaa3fa]
- vcsim: Add cdrom and scsi controller to Model VMs [f6a734f5]
- Move vcsim from vmware/vic to govmomi [9d47dd13]

<a name="v0.15.0"></a>
## [Release v0.15.0](https://github.com/vmware/govmomi/compare/v0.14.0...v0.15.0)

> Release Date: 2017-06-19

### 📖 Commits

- Release 0.15.0 [b63044e5]
- Add dvs.portgroup.info usage [3d357ef2]
- Add support for guest.FileManager directory download [72977afb]
- Update examples [94837bf7]
- Update wsdl generator [e1bbcf52]
- fix the WaitOptions struct, MaxWaitSeconds is optional, but we can set the value 0 [b16a3d81]
- Support removal of ExtraConfig entries [9ca7a2b5]
- Guest command updates [86cc210c]
- Doc updates [9c5f63e9]
- New examples: datastores, hosts and virtualmachines using view package [6d714f9e]
- update spew to be inline with testify [f48e1151]
- Adjust message slice passed to include [6f5c037c]
- Fix package name [48509bc3]
- Add host.shutdown command [6f635b73]
- Add doc on metric.sample instance flag ([#726](https://github.com/vmware/govmomi/issues/726)) [67b13b52]
- Fix tail n=0 case ([#725](https://github.com/vmware/govmomi/issues/725)) [8bff8355]
- Update copyright ([#723](https://github.com/vmware/govmomi/issues/723)) [10e6ced9]
- Allow caller to supply custom tail behavior ([#722](https://github.com/vmware/govmomi/issues/722)) [6f8ebd89]
- Add options to host.autostart.add ([#719](https://github.com/vmware/govmomi/issues/719)) [35caa01b]
- Add VC options command ([#717](https://github.com/vmware/govmomi/issues/717)) [2030458d]
- Exported FindSnapshot() Method ([#715](https://github.com/vmware/govmomi/issues/715)) [0ccad10c]
- Additional wrapper functions for SPBM [34202aca]
- Add AuthorizationManager {Enable,Disable}Methods [c7f718b1]
- Add PBM client and wrapper methods [d5e08cd2]
- Add generated types and methods for PBM [58019ca9]
- Regenerate against current vmodl.db [58960380]
- Support non-Go clients in xml decoder [f736458f]

<a name="v0.14.0"></a>
## [Release v0.14.0](https://github.com/vmware/govmomi/compare/v0.13.0...v0.14.0)

> Release Date: 2017-04-08

### 📖 Commits

- Release 0.14.0 [9bfdc5ce]
- Release 0.13.0 [3ba0eba5]
- Add object.find command [86063832]
- Adds FindManagedObject method. [0391e8eb]
- Include embedded fields in object.collect output [796e87c8]
- Use Duration flag for vm.ip -wait flag [2536e792]
- Merge commit 'b0b51b50f40da2752c35266b7535b5bbbc8659e3' into marema31/govc-vm-ip-wait [3aa64170]
- Implement EthernetCardBackingInfo for OpaqueNetwork [59466881]
- Finder: support changing object root in find mode [0d2e1b22]
- Add Bash completion script [9ded9d10]
- Add QueryVirtualDiskInfo [3bd4ab46]
- Emacs: add metric select [16f6aa4f]
- Add unit conversion to metric CSV [3763321e]
- Add -wait option to govc vm.ip to allow non-blocking query [b0b51b50]
- Add json support to metric ls and sample commands [f0d4774a]
- Add performance manager and govc metric commands [c9de0310]
- Add check for nil envelope [d758f694]
- Remove deferred Close() call in follower's Read() [ab595fb3]

<a name="v0.13.0"></a>
## [Release v0.13.0](https://github.com/vmware/govmomi/compare/v0.12.1...v0.13.0)

> Release Date: 2017-03-02

### 💫 `vcsim` (Simulator)

- esxcli FirewallInfo fixes ([#661](https://github.com/vmware/govmomi/issues/661))

### 📖 Commits

- Release 0.13.0 [b4a3f7a1]
- Add vm.guest.tools command [5bf03cb4]
- Host is optional for MarkAsVirtualMachine ([#675](https://github.com/vmware/govmomi/issues/675)) [b4ef3b73]
- Add vsan and disk commands / helpers ([#672](https://github.com/vmware/govmomi/issues/672)) [f4a3ffe5]
- Handle the case where e.VirtualSystem is nil ([#671](https://github.com/vmware/govmomi/issues/671)) [1f82c282]
- Remove object.ListView ([#669](https://github.com/vmware/govmomi/issues/669)) [dd346974]
- Wraps the ContainerView managed object. ([#667](https://github.com/vmware/govmomi/issues/667)) [4994038a]
- Handle nil TaskInfo in task.Wait callback [#2](https://github.com/vmware/govmomi/issues/2) ([#666](https://github.com/vmware/govmomi/issues/666)) [93064c06]
- Handle nil TaskInfo in task.Wait callback ([#665](https://github.com/vmware/govmomi/issues/665)) [f1f5b6cb]
- Support alternative './...' syntax for finder ([#664](https://github.com/vmware/govmomi/issues/664)) [f3cf126d]
- Finder: support automatic Folder recursion ([#663](https://github.com/vmware/govmomi/issues/663)) [9bda6c3e]
- Add a command line option to change an existing disk attached to a VM ([#658](https://github.com/vmware/govmomi/issues/658)) [0a28e595]
- Attach and list RDM/LUN ([#656](https://github.com/vmware/govmomi/issues/656)) [3e95cb11]
- vcsim: esxcli FirewallInfo fixes ([#661](https://github.com/vmware/govmomi/issues/661)) [5f7efaf1]
- Add device option to WaitForNetIP ([#660](https://github.com/vmware/govmomi/issues/660)) [17e6545f]
- Fix vm.change test [ba9e3f44]
- Add the option to describe a VM using the annotation option in ConfigSpec ([#657](https://github.com/vmware/govmomi/issues/657)) [e66c8344]
- Update doc [505fcf9c]
- Add support for reading and changing SyncTimeWithHost option ([#539](https://github.com/vmware/govmomi/issues/539)) [913c0eb4]
- Remove _Task suffix from vapp methods [682494e1]
- Emacs: add govc-command-history [733acc9e]
- Add object.collect command ([#652](https://github.com/vmware/govmomi/issues/652)) [ea52d587]
- Update email address for contributor Bruce Downs [f49782a8]

<a name="v0.12.1"></a>
## [Release v0.12.1](https://github.com/vmware/govmomi/compare/v0.12.0...v0.12.1)

> Release Date: 2016-12-19

### 📖 Commits

- Release 0.12.1 [6103db21]
- Note 6.5 support [45a53517]
- Add '-f' flag to logs command ([#643](https://github.com/vmware/govmomi/issues/643)) [fec40b21]
- govc.el: auth-source integration ([#648](https://github.com/vmware/govmomi/issues/648)) [40cf9f80]
- Add govc-command customization option ([#645](https://github.com/vmware/govmomi/issues/645)) [ca99f8de]
- Avoid Finder panic when SetDatacenter is not called ([#640](https://github.com/vmware/govmomi/issues/640)) [ad6e5634]
- Add storage support to vm.migrate ([#641](https://github.com/vmware/govmomi/issues/641)) [b5c807e3]
- govc/version: skip first char in git version mismatch error ([#642](https://github.com/vmware/govmomi/issues/642)) [1a7dc61e]
- Add Slack links [6bc730e1]
- Add DatastorePath helper ([#638](https://github.com/vmware/govmomi/issues/638)) [e152c355]
- Add support for file backed serialport devices ([#637](https://github.com/vmware/govmomi/issues/637)) [5b4d5215]
- Add vm.ip docs ([#636](https://github.com/vmware/govmomi/issues/636)) [f49bd564]

<a name="v0.12.0"></a>
## [Release v0.12.0](https://github.com/vmware/govmomi/compare/v0.11.4...v0.12.0)

> Release Date: 2016-12-01

### 📖 Commits

- Release 0.12.0 [ab40ac73]
- Disable use of service ticket for datastore HTTP access by default ([#635](https://github.com/vmware/govmomi/issues/635)) [e702e188]
- Attach context to HTTP requests for cancellations [1fba1af7]
- Support InjectOvfEnv without PowerOn when importing [79cb3d93]
- Support stdin as import options source [117118a2]
- Don't ignore version/manifest for existing sessions [b10f20f4]
- Add basic VirtualNVMEController support [82929d3f]
- re-generate vim25 using 6.5.0 [757a2d6d]

<a name="v0.11.4"></a>
## [Release v0.11.4](https://github.com/vmware/govmomi/compare/v0.11.3...v0.11.4)

> Release Date: 2016-11-15

### 📖 Commits

- Release 0.11.4 [b9bcc6f4]
- Add authz role helpers and commands [dbbf84e8]
- Add folder/pod examples [765b34dc]
- Add host.account examples [79cb52fd]
- Add host.portgroup.change examples [2a2cab2a]

<a name="v0.11.3"></a>
## [Release v0.11.3](https://github.com/vmware/govmomi/compare/v0.11.2...v0.11.3)

> Release Date: 2016-11-08

### 📖 Commits

- Release 0.11.3 [e16673dd]
- Add -product-version flag to dvs.create [629a573f]
- Allow DatastoreFile follower to drain current body [83028634]

<a name="v0.11.2"></a>
## [Release v0.11.2](https://github.com/vmware/govmomi/compare/v0.11.1...v0.11.2)

> Release Date: 2016-11-01

### 📖 Commits

- Release 0.11.2 [cd80b8e8]
- Avoid possible NPE in VirtualMachine.Device method [f15dcbdc]
- Add support for OpaqueNetwork type [128b352e]
- Add host account manager support for 5.5 [c5b9a266]

<a name="v0.11.1"></a>
## [Release v0.11.1](https://github.com/vmware/govmomi/compare/v0.11.0...v0.11.1)

> Release Date: 2016-10-27

### 📖 Commits

- Release 0.11.1 [1a7df5e3]
- Add support for VirtualApp in pool.change command [1ae858d1]
- Release script tweaks [91b2ad48]

<a name="v0.11.0"></a>
## [Release v0.11.0](https://github.com/vmware/govmomi/compare/v0.10.0...v0.11.0)

> Release Date: 2016-10-25

### 📖 Commits

- Release 0.11.0 [a16901d7]
- Add object destroy and rename commands [4fc9deb4]
- Add dvs.portgroup.change command [82634835]

<a name="v0.10.0"></a>
## [Release v0.10.0](https://github.com/vmware/govmomi/compare/v0.9.0...v0.10.0)

> Release Date: 2016-10-20

### 📖 Commits

- Release 0.10.0 [bb498f73]
- Release script updates [468a15af]
- Documentation updates [1c3499c4]
- Update contributors [1e52d88a]
- Fix snapshot.tree on vm with no snapshots [e3d59fd9]
- Add host.date info and change commands [711fdd9c]
- Add govc session ls and rm commands [16d7514a]
- Add HostConfigManager field checks [73c471a9]
- Improve cluster/host add thumbprint support [d7f94557]
- Add session.Locale var to change default locale [fea8955b]
- Add service ticket thumbprint validation [eefe6cc1]
- Set default locale to en_US [3a0a61a6]
- TLS enhancements [aa1a9a84]
- Treat DatastoreFile follower Close as "stop" [9f0e9654]
- Support typeattr for enum string types [838b2efa]
- Make vm.ip esxcli test optional [dcbc9d56]
- Remove vca references [9e20e0ae]
- Adding vSPC proxyURI to govc [7c708b2e]

<a name="v0.9.0"></a>
## [Release v0.9.0](https://github.com/vmware/govmomi/compare/v0.8.0...v0.9.0)

> Release Date: 2016-09-09

### 📖 Commits

- Release 0.9.0 [f9184c1d]
- Add govc -h flag [e050cb6d]
- Set default ScsiCtlrUnitNumber [a4343ea8]
- Add -R option to datastore.ls [a920d73d]
- Fix SCSI device unit number selection [f517decc]
- Add DatastoreFile helpers [abaf7597]
- Make Datastore ServiceTicket optional [7cfa7491]
- Add vm.migrate command [9ad57862]
- Add govc vm.{un}register commands [c66458f9]
- Checking result of reflect.TypeOf is not nil before continuing [54c0c6e5]
- Fix flags.NewOptionalBool panic [ea0189ea]
- Add govc guest command tests [a9cdf437]
- Add VirtualMachine.Unregister func [38dee111]
- make curl follow HTTP redirects [98b50d49]
- make goreportcard happy [8a27691f]
- Add govc vm snapshot commands [bf66f750]
- Validate vm.clone -vm flag value [eb02131a]
- Add device.usb.add command [62159d11]
- Remove a bunch of context.TODO() calls. [27e02431]
- Fixing tailing for events command [a9cee43a]
- Bump to 1.7 and start using new context pkg [4fa7b32a]
- Fix missing datastore name with vm.clone -force=false [4b7c59bf]
- Fix deletion of powered off vApp [e3642fce]
- Support stdin/stdout in datastore upload/download [63d60025]
- Emacs: add govc-session-network [e149909e]
- Emacs: add govc json diff [0ccc1788]
- Add host.portgroup.change command [f1d6e127]
- Add host.portgroup.info command [6f441a84]
- Add HostNetworkPolicy to host.vswitch.info [aaf40729]
- Add json support to host.vswitch.info command [5ccb0572]
- Support instance uuid in SearchFlag [9d19d1f7]
- Add json support to esxcli command [2d3bfc9f]
- Support multiple NICs with vm.ip -esxcli [bac04959]
- Add -unclaimed flag to host.storage.info command [b3177d23]
- govc - popualte 'Path' fiels in xxx.info output [b1234a90]
- Implemented additional ListView methods [7cab0ab6]
- Add 'Annotation' attribute to importx options. [498cb97d]
- Add NetworkMapping section to importx options. [223168f0]
- Remove vendor target from the Makefile [5c708f6b]
- Handle errors in QueryVirtualDiskUUid function ([#548](https://github.com/vmware/govmomi/issues/548)) [f8199eb8]
- vendor github.com/davecgh/go-spew/spew [73dcde2c]
- vendor golang.org/x/net/context [e1e407f7]
- Populate network mapping from ovf envelope ([#546](https://github.com/vmware/govmomi/issues/546)) [e3c3cd0a]
- Add QueryVirtualDiskUuid function ([#545](https://github.com/vmware/govmomi/issues/545)) [fa6668dc]
- Fixes panic in govc events [17682d5b]

<a name="v0.8.0"></a>
## [Release v0.8.0](https://github.com/vmware/govmomi/compare/v0.7.1...v0.8.0)

> Release Date: 2016-06-30

### 📖 Commits

- Release 0.8.0 [c0c7ce63]
- Disable datastore service ticket hostname usage [ce4b0be6]
- Add support for login via local ticket [3e44fe88]
- Add StoragePod support to govc folder.create [acf37905]
- Include StoragePod in Finder.FolderList [94d4e2c9]
- Avoid use of eval with govc env [473f3885]
- Add datacenter.create folder option [4fb7ad2e]
- Avoid vm.info panic against vcsim [77ea6f88]
- Session persistence improvements [95b2bc4d]
- Add type attribute to soap.Fault Detail [720bbd10]
- Add filtering for use of datastore service ticket [ff7b5b0d]
- Add support for Finder lookup via moref [fe9d7b52]
- Use ticket HostName for Datastore http access [c26c7976]
- Add govc/vm.markasvm command [bea2a43c]
- Add govc/vm.markastemplate command [9101528d]
- Add vm.markastemplate [982e64b8]

<a name="v0.7.1"></a>
## [Release v0.7.1](https://github.com/vmware/govmomi/compare/v0.7.0...v0.7.1)

> Release Date: 2016-06-03

### 📖 Commits

- Fix Datastore upload/download against VC [2cad28d0]

<a name="v0.7.0"></a>
## [Release v0.7.0](https://github.com/vmware/govmomi/compare/v0.6.2...v0.7.0)

> Release Date: 2016-06-02

### 📖 Commits

- Release 0.7.0 [6906d301]
- Move InventoryPath field to object.Common [558321df]
- Add -require flag to govc version command [4147a6ae]
- Add support for local type in datastore.create [d9fd9a4b]
- Fix vm.create disk scsi controller lookup [650b5800]
- Update changelog for govc to add datastore -namespace flag [9463b5e5]
- Update changelog with DatastoreNamespaceManager methods [4aab41b8]
- Support mkdir/rm of namespace on vsan [4d6ea358]
- InjectOvfEnv() should work with VSphere [bb7e2fd7]
- Add host.service command [91ca6bd5]
- Add host.storage.mark command [2f369a29]
- Add -rescan option to host.storage.info command [b001e05b]

<a name="v0.6.2"></a>
## [Release v0.6.2](https://github.com/vmware/govmomi/compare/v0.6.1...v0.6.2)

> Release Date: 2016-05-13

### 📖 Commits

- Release 0.6.2 [9051bd6b]
- Get complete file details in Datastore.Stat [3ab0d9b2]
- Convert types when possible [0c21607e]
- Avoid xsi:type overwriting type attribute [648d945a]
- adding remove all snapshots to vm objects [4e0680c1]

<a name="v0.6.1"></a>
## [Release v0.6.1](https://github.com/vmware/govmomi/compare/v0.6.0...v0.6.1)

> Release Date: 2016-04-30

### 📖 Commits

- Release 0.6.1 [18154e51]
- Fix mo.Entity interface [47098806]

<a name="v0.6.0"></a>
## [Release v0.6.0](https://github.com/vmware/govmomi/compare/v0.5.0...v0.6.0)

> Release Date: 2016-04-29

### 📖 Commits

- Release 0.6.0 [2c1d977a]
- Add folder.moveinto command [cc686c51]
- Add folder.{create,destroy,rename} methods [8e85a8d2]
- Add Common.Rename method [0ba22d24]
- Fix Finder.FolderList check [61792ed3]
- Restore optional DatacenterFlag [b6be92a1]
- Add OutputFlag support to govc about command [53903a3a]
- Add OptionManager and host.option commands [e66f7793]
- Add debug xmlformat script [9d69fe4b]
- Add option to use the same path for debug runs [f1786bec]
- Add folder.info command [99c8c5eb]
- Add datacenter.info command [eca4105a]
- Add mo.Entity interface [71484c40]
- Add helper to wait for multiple VM IPs [388df2f1]
- Add RevertToSnapshot [fc9f58d0]
- Add govc env command [a4aca111]
- Update CI config [ef17f4bd]
- Add host.account commands [fa91a600]
- Update release install instructions [44bb6d06]
- Leave AddressType empty in EthernetCardTypes [08ba4835]
- Add vm clone [f9704e39]
- Add datastore.Download method [e6969120]
- device.remove: add keep option [1aca660c]

<a name="v0.5.0"></a>
## [Release v0.5.0](https://github.com/vmware/govmomi/compare/v0.4.0...v0.5.0)

> Release Date: 2016-03-30

### 📖 Commits

- Release 0.5.0 [c1b29993]
- Use VirtualDeviceList for import.vmdk [b8549681]
- Remove debug flags from pool tests [cf96f70d]
- Switch to int32 type for xsd int fields [f74a896d]
- Regenerate against 6.0u2 wsdl [074494df]
- Include license header in generated files [ce9314c4]
- Add pointer field white list to generator [957c8827]
- Change pool recusive destroy to children destroy [2c1d1950]
- Add dvs.portgroup.info command [5d34409f]
- Update docs [216031c3]
- Remove govc-test pools in teardown hook [f7dfcc98]
- Simplify pool destroy test [556a9b17]
- Add folder management to vm.create [4e47b140]
- Update test ESX IP in Drone secrets file [7c33bcb3]
- Regenerate Drone secrets file [1b6ec477]
- Implemented the ablitiy to tail the vSphere event stream - govc tail and force flag added to events command [f64ea833]
- Including github.com/davecgh/go-spew/spew in go get [fd7d320f]
- Including github.com/davecgh/go-spew/spew in go get [1d4efec0]
- The -dump option now requests a recursive traversal as -json does [424d3611]
- Added new -dump output flag for pretty printing underlying objects using davecgh/go-spew [b45747f3]
- Run govc tests against ESX using Drone [a243716c]
- Double quotes network name to prevent space in name from failing the tests [fb75c63e]
- test_helper.bash updated to conditionally set env variables [564944ba]
- Added new govc vm.disk.create -mode option for selecting one the VirtualDiskMode types [c9c6e38f]
- Add -net flag to device.info command [6922c88b]
- Fix VirtualDeviceList.CreateFloppy [dff2c197]
- Ran gofmt on create.go [c7d8cd3e]
- Fix issue with optional UnitNumber (v2) [e077bcf5]
- Added arguments to govc vm.disk.create for thick provisioning and eager scrubbing, as requested in issue [#254](https://github.com/vmware/govmomi/issues/254) [539ad504]
- Handle import statement for types too [e66c6df9]
- Remove hardcoded urn:vim25 value from vim_wsdl.rb [265d8bdb]

<a name="v0.4.0"></a>
## [Release v0.4.0](https://github.com/vmware/govmomi/compare/v0.3.0...v0.4.0)

> Release Date: 2016-02-26

### 📖 Commits

- Release 0.4.0 [b3d202ab]
- Fix vm.change's ExtraConfig values being truncated at equal signs [749da321]
- Add switch to specify protocol version in SOAPAction header [13fbc59d]
- Update CHANGELOG [07013a97]
- Allow vm.create to take datastore cluster argument [bfe414fe]
- Include reference to datastore in CreateDisk [dda71761]
- Make NewKey function public [855abdb3]
- Use custom datastore flags in vm.create [d0031106]
- Modify govc's vm.create to create VM in one shot [306b613d]
- Add extra datastore arguments to vm.create [e96130b4]
- Add datastore cluster methods to finder [0a2da16d]
- Allow StoragePod type to be traversed [c69e9bc1]
- added explicit path during clone [4d2ea3f4]
- Update missing property whitelist [3d8eb102]
- re-generate vim25 using 6.0 Update 1b (vimbase [#3024326](https://github.com/vmware/govmomi/issues/3024326)) [779ae0a1]
- Handle import statements same as include [53c29f6a]
- Update govc.el URL [a738f89d]
- Doc updates [da2a249e]
- govc.el: minor fixes for distribution as a package [47e46425]
- handle GOVC_TEST_URL=user:pass[@IP](https://github.com/IP) pattern [8459ceb9]
- Add Emacs interface to govc [3b669760]
- Update README to include Drone build status and local build instructions [7ec8028d]
- Add config for Drone CI build [2ec65fbe]
- introduce Datastore.Type() [5437c466]
- introduce IsVC method and start using it [983571af]
- Introduce AttachedClusterHosts [0732f137]
- start using new helper functions for govc/flags [18945281]
- Add some common functions to find/finder.go [044d904a]
- Support vapp in pool.info command [534dabbd]
- Fix bats tests [4d9c6c72]
- Add -p and -a options to govc datastore.ls command [5e04d5ca]
- Added check for missing ovf deployment section [33963263]

<a name="v0.3.0"></a>
## [Release v0.3.0](https://github.com/vmware/govmomi/compare/v0.2.0...v0.3.0)

> Release Date: 2016-01-15

### 📖 Commits

- Mark 0.3.0 in change log [501f6106]
- Update contributors [83a26512]
- Print os.Args[0] in error messages [995d970f]
- Move stat function to object.Datastore [0a4c9782]
- Support VirtualApp in the lister [8a0d4217]
- Support empty folder in SearchFlag.VirtualMachines [82734ef3]
- Add support for custom session keep alive handler [f64f878f]
- Use OptionalBool for ExpandableReservation [2d498658]
- Script to capture vpxd traffic on VCSA [ac9a39b0]
- Script to capture and decrypt hostd SOAP traffic [3f473628]
- Move govc url.Parse wrapper to soap.ParseURL [eccc3e21]
- Don't assume sshClient firewall rule is disabled [e1031f44]
- Let the lister recurse into a ComputeHost [cd5d8baa]
- Specify the new entity's name upon import [b601a586]
- Explicitly instantiate and register flags [a5e26981]
- Parameterize datastore in VM tests [aca77c67]
- Pass context to command and flag functions [37324472]
- Minor optimization to encoding usage [6f955173]
- Create VMFS datastore with datastore.create [0f4aee8b]
- Add host storage commands [ec724783]
- Run license script [debdd854]
- Fix license script to work for uncommitted files [64022512]
- Remove host reference from HostFirewallSystem [5cb0c344]
- Change the comment that mentions ha-datacenter [4fb4052a]
- Let the ESXi to figure out datastore name [b76ad0eb]
- Add helper method to get VM power state [918188dc]
- Add permissions.{ls,set,remove} commands [29a2f027]
- Add DatacenterFlag.ManagedObjects helper [f27787a1]
- Option to disable API version check [0e629647]
- Add commands to add and remove datastores [42d899d0]
- Check host state in Datastore.AttachedHosts [369e0e7f]
- Test that vm.info -r prints mo names [7adf8375]
- Change ComputeResource.Hosts to return HostSystem [3198242e]
- Support property collection for embedded types [b34f346e]
- Fix vm nested hv option [8035c180]
- Update copyright years in code headers [b1d9d3c2]
- Add dvs commands [c99e7bac]
- Support DVS lookup in finder [c30b7f17]
- Embed Reference interface in NetworkReference [094fbdfe]
- Add DVS helpers [0657cf76]
- Add host.vnic.{service,info} commands [6e96a1db]
- Add VsanSystem and VirtualNicManager wrappers [ae6b0b77]
- Add vsan flags to cluster.change command [24297494]
- Add license.assigned.list id flag [4088502d]
- Add cluster.add license flag [d089489e]
- Add vm.change options to set hv/mmu [31ee6e03]
- Refactor host.add command to use HostConnectFlag [a414852e]
- Add cluster.{create,change,add} commands [51543392]
- Add cluster related host commands [8262e1da]
- Add HostConnectFlag [2443b364]
- Add object.HostSystem methods [8ae7da82]
- Add finder.Folder method [0f630dd9]
- Add bash function to save/load GOVC environments [7cd5fbb5]
- Add object.Common.Destroy method [12f26c21]
- Add ComputeResource.Reconfigure method [2ab8aa59]
- Add flags.NewOptionalBool [5f47f155]
- Add -feature flag to license list commands [25fe42b2]
- Add license.InfoList list wrapper [2e6c0476]
- Add license assignment commands [ef7371af]
- Add license.AssignmentManager [5005e6e4]
- Use object.Common in license.Manager [69a23bd4]
- Rename receiver variable for consistency [dbce3faf]
- Pass pointer to bps uint64 in last progress report [80705c11]
- VirtualMachine: Add Customize function on object.VirtualMachine [26e77c8e]
- Add license.decode command [c2a78973]
- Add DistributedVirtualPortgroup support to vm.info [b3a7e07e]
- Fix KeepAlive [1b11ad02]
- Add HostFirewallSystem wrapper [3ecfd0db]
- KeepAlive support with certificate based login [9ded9c1a]
- Add DiagnosticManager and logs commands [cf2a879b]
- Update README.md [7b14760a]
- Export Datastore.ServiceTicket method [ad694500]
- Added a method to create snapshot of a virtual machine [76690239]
- Use service ticket for datastore file access [6d4932af]
- Fix vcsa ssh config [5fcc29f6]
- Retry on empty result from property collector [ac390ec8]
- Add methods for client certificate based auth [f3041b2c]
- Add extension manager and govc commands [b9edc663]
- Fix key composition in building OVF spec [9057659c]
- Move OVF environment related code to env{,test}.go [f56f6e80]
- Add minimal doc to ovf package [b33c9aef]
- Added verbose option to the import.spec feature [3d40aefb]
- change for looking up a VM using instanceUUID [1df0a81d]
- Introduce govc vapp.{info|destroy|power} [5f4d36cd]
- Handle the import.spec case where no spec file is provided [88795252]
- Add inventory path to govc info commands [bcdc53fb]
- Collect govc host and pool info in one call [305371a8]
- Relax the convention around importing an ova [bfd47026]
- don't start goroutine while context is nil [3742a8aa]

<a name="v0.2.0"></a>
## [Release v0.2.0](https://github.com/vmware/govmomi/compare/prerelease-v0.1.0-73-gfc131d4...v0.2.0)

> Release Date: 2015-09-15

### ⏮ Reverts

- Add Host information to vm.info

### 📖 Commits

- Mark 0.2.0 in change log [b3315079]
- Add mode argument to release script [cc3bcbee]
- Build govc with new cross compilation facilities [ae4a6e53]
- Derive CONTRIBUTORS from commit history [4708d165]
- Move contrib/ -> scripts/ [00909f48]
- Capitalization [a0f4f799]
- Split import functionality into independent flags [13baa0e4]
- Added ovf.Property output to import.spec [6363d0e2]
- Update change log [7af121df]
- Fix event.Manager category cache [f9deb385]
- Avoid tabwriter in events command [7f0a892d]
- Use vm.power force flag for hard shutdown/reboot [29601b46]
- Add VirtualDiskManager CreateVirtualDisk wrapper [ea833cf5]
- Interative clean up of bats testing [bfabd01d]
- Clean up of vcsa creation script [7cba62d9]
- Add serial port URI info to device.info output [631d6228]
- Add -json support to device.info command [0b31dcff]
- Add govc vm.info resources option [54e324d1]
- Add helper method to wait for virtual machine power state. [9cc5d8f5]
- Remove superfluous math.Pow calculations [9ddd6337]
- Added common method of humanizing byte strings [5272b1e9]
- Add helper method to check if VMware Tools is running in the guest OS. [3145d146]
- Misc clean up [e4f4c737]
- Add host name to vm.info [01f2aed0]
- Use property.Collector.Retrieve() in vm.info [f24ec75a]
- Renamed vm.info VmInfos back to VirtualMachines [a779c3b7]
- Revert "Add Host information to vm.info" [2900f2ff]
- Add -hints option to host.esxcli command [2a567478]
- Add options to importing an ovf or and ova file [1f0708e2]
- Only retrieve "currentSession" property [debde780]
- Update CONTRIBUTORS [b5187c16]
- Added the ability to specify ovf properties during deployment [3e4ced8c]
- Introduce more VirtualApp methods [688a6b18]
- Add flag to specify destination folder for import.ovf and import.ova [b1f0cb0c]
- Add check for error reading ova file [c9fcf1ce]
- clone vmware/rbvmomi repo if it's missing [edb0a2cf]
- use e.Object.Reference().Type as suggested by Doug [40c26fc6]
- introduce CreateVApp and CreateChildVM_Task [c1442f95]
- add VirtualAppList and VirtualApp methods to Finder [25405362]
- Add CustomFieldsManager wrapper and cli commands [121f075c]
- include VirtualApp in ls -l output [dd016de3]
- Provide ability to override url username and password [b5db4d6d]
- Add OVF unmarshalling [11d5ae9c]
- Update travis.yml for new infra [135569e7]
- Make govet stop complaining [822432eb]
- Add datastore.info cli command [baf9149e]
- Add serial port matcher to SelectByBackingInfo [2b93c199]
- Merge branch 'gavrie-master' [26ba22de]
- Add Host information to vm.info [62591576]
- Add methods for useful properties to VirtualMachine [a90019ab]
- Add Relocate method to VirtualMachine [502963c4]
- Add String method to objects for pretty printing [7f4b6d38]
- Add events helpers and cli command [99f57f16]
- Update CONTRIBUTORS [4c989ac3]
- Update to vim25/6.0 API [ad7d1917]
- Add net.address flag [ad39adb9]
- Add Usage for host.esxcli [1acf418c]
- Modify archive.go bug [2e00fdb1]

<a name="prerelease-v0.1.0-73-gfc131d4"></a>
## [Release prerelease-v0.1.0-73-gfc131d4](https://github.com/vmware/govmomi/compare/prerelease-v0.1.0-62-g7734772...prerelease-v0.1.0-73-gfc131d4)

> Release Date: 2015-07-13

### 📖 Commits

- Add command to add host to datacenter [e01555f9]
- Stop returning children from `ManagedObjectList` [efbd3293]
- Update CONTRIBUTORS [d16670f5]
- Mention GOVC_USERNAME and GOVC_PASSWORD in CHANGELOG [97fbf898]
- Add test to check for flag name collisions [8766bda0]
- Remove flags for overriding username and password [791b3365]
- include GOVC_USERNAME and GOVC_PASSWORD in govc README [85957949]
- Export variables in release script [8584259a]

<a name="prerelease-v0.1.0-62-g7734772"></a>
## [Release prerelease-v0.1.0-62-g7734772](https://github.com/vmware/govmomi/compare/prerelease-v0.1.0-52-g871f5d4...prerelease-v0.1.0-62-g7734772)

> Release Date: 2015-07-06

### 📖 Commits

- Add test for GOVC_USERNAME and GOVC_PASSWORD [14889008]
- Only run license tests against evaluation license [c0a984cd]
- Allow override of username and password [293ac813]
- Add extraConfig option to vm.change and vm.info [e053bdf2]
- Update CONTRIBUTORS [1dec0695]
- Add missing types to list.ToElement [985291d5]

<a name="prerelease-v0.1.0-52-g871f5d4"></a>
## [Release prerelease-v0.1.0-52-g871f5d4](https://github.com/vmware/govmomi/compare/v0.1.0...prerelease-v0.1.0-52-g871f5d4)

> Release Date: 2015-06-16

### ⏮ Reverts

- Fix git dirty status error in build script

### 📖 Commits

- Add script to create a draft prerelease [871f5d4f]
- Revert "Fix git dirty status error in build script" [8bec13f7]
- Only use annotated tags to describe a version [c825a3c7]
- Retry twice on temporary network errors in govc [66320cb0]
- Add retry functionality to vim25 package [67be5f1d]
- Add method to destroy a compute resource [fba0548b]
- Add methods to add standalone or clustered hosts [2add2f7a]
- Add ability to create, read and modify clusters [de297fcb]
- Change finder functions to no longer take varargs [f10480af]
- Fix resource pool creation/modification [4bc93a66]
- Rename persist flag to persist-session [b434a9a8]
- Ignore ManagedObjectNotFound in list results [d85ad215]
- Add example that lists datastores [4c497373]
- Update govc CHANGELOG [5d153787]
- Add flag to toggle persisting session to disk [0165e2de]
- Add Mevan to CONTRIBUTORS [8acb2f28]
- Ignore missing environmentBrowser field [add15217]
- Fix error when using SDRS datastores [447d18cd]
- Find ComputeResource objects with find package [e85f6d59]
- Test package only depends on vim25 [55f984e8]
- Drop omitempty tag from optional pointer fields [dbe47230]
- Interpret negative values for unsigned fields [749f0bfa]
- Update CHANGELOG [49a34992]
- Update code to work with bool pointer fields [263780f3]
- Make optional bool fields pointers [93aad8da]
- Return errors for unexpected HTTP statuses [b7c51f61]
- Abort client tests on errors [62ca329a]
- Rename LICENSE file [ae345e7f]
- Add govc CHANGELOG [a783a8c6]
- Add commands to configure the autostart manager [ba707586]
- Re-enable search index test [af6a188e]
- Update govc README [ceea450c]
- Fix git dirty status error in build script [ea5c9a52]

<a name="v0.1.0"></a>
## [Release v0.1.0](https://github.com/vmware/govmomi/compare/test...v0.1.0)

> Release Date: 2015-03-17

### 📖 Commits

- Cross-compile govc using gox [477dcaf9]
- Add version variable that can be set by the linker [8593d9c7]
- Add CHANGELOG [fb38ca45]
- Add package docs to client.go [76f8f1a1]
- Use context.Context in client in root package [27bf35df]
- Comment out broken test [f3b8162f]
- Drop the _gen filename suffix [a1d9d1e7]
- Add context.Context argument to object package [91650a1f]
- Use vim25.Client throughout codebase [1814113a]
- Move property retrieval functions to property package [b977114e]
- Add lightweight client structure to vim25 package [8c3243d8]
- Add context.Context argument to find/list packages [ec4b5b85]
- Make Wait function in property package standalone [7eecfbc7]
- Add keep alive for soap.RoundTripper [6c1982c8]
- Return nil UserSession when not authenticated [1324d1f0]
- Comments for task.Wait [ae7ea3dd]
- Add context parameter to object.Task functions [a53a6b2c]
- Move functionality to wait for task to new package [f6f44097]
- Move Ancestors function to vim25/mo [ad2303cf]
- Move PropertyCollector to new property package [fb9e1439]
- Move Reference to vim25/mo [a6618591]
- Bind virtual machine to guest operation wrappers [bfdb90f1]
- Move HasFault to vim25/types [ec0c16a7]
- Move wrappers for managed objects to object package [683ca537]
- Add GetServiceContent function to vim25/soap [223a07f8]
- Decouple factory functions from client [25b07674]
- Move SessionManager to new session package [b96cf609]
- Return on error in SessionManager [ea8d5d11]
- Mutate copy of parameter instead of parameter itself [7d58a49e]
- Marshal soap.Client instead of govmomi.Client [e158fd95]
- Embed soap.Client in govmomi.Client [1336ad45]
- Work with pointer to url.URL [15cfd514]
- Move guest related wrappers to new guest package [be2936f8]
- Move LicenseManager to new license package [b772ba28]
- Move EventManager to new event package [7ac1477f]
- Retrieve dependencies before running test [2053e065]
- Add context.Context argument to RoundTripper [2d14321e]
- Include type of request in summarized debug log [64f716b2]
- Store reference to http.Transport [40249c87]
- Move debugging code in soap.Client to own struct [ac77f0c5]
- Loosen .ovf match in ova.import [c8fab31b]
- And further fixing the merge... go fmt. [9f685e92]
- Merge remote-tracking branch 'upstream/master' into event_manager [8dbb438b]
- created session manager wrapper [e57a557c]
- Change return pattern in CreateDatacenter [5525d5c6]
- Update contributors [8acd5512]
- Coding style consistency [7138d375]
- added SessionIsActive to Client [951e9194]
- Add CreateFolder method [2211e73d]
- Add Login/Logout functions to client struct [eef40cc0]
- Update contributors [3c7dea04]
- Fixed error when attempting to access datastore [9c4a9202]
- Add PropertiesN function on client struct [05ee0e62]
- Adding EventManager so that events can be queried for [01ee2fd5]
- Restrict permissions on session file [8d10cfc7]
- Key session file off of entire URL [88b5d03c]
- Error types for getter functions on finder [9354d314]
- Add description for pool.create [a30287dc]
- Prefix option list in help output [77466af0]
- Create multiple resource pools through pool.create [cbb8d0b2]
- Add usage and description for pool.destroy [8d4699d8]
- Change pool.change to take multiple arguments [2e195a92]
- Add usage and description for pool.info [38e4a2b2]
- Add usage and description for pool.create [2f286768]
- Set insert_key = false [413fa901]
- Update travis.yml [d6c2b33e]
- Adding CustomizationSpecManager [b878c20a]
- Add vm mark as vm and mark as template features [7c8f3e56]
- Update contributors [033d02e9]
- Add cpu and memory usage to host.info [18919172]
- Adding the RegisterVM task. [b29f93c1]
- Add error types for Finder [e6bf8bb5]
- Support multiple hosts in host.info command [852578b9]
- Set InventoryPath field [f1899c63]
- Add InventoryPath field [3a5c1cf3]
- Add resource pool cli commands [624f21a4]
- Add ResourcePool wrapper methods [4c7cd61f]
- Include ResourcePool in ls -l output [761d43e5]
- Support nested resource pools in lister [d2daf706]
- Add vm.change cli command [4d9d9a72]
- bats fixup: destroy datacenter [e6ebcd7f]
- Disable vcsa box password expiration [65838131]
- Add CONTRIBUTORS file [7a6e737b]
- Issue [#192](https://github.com/vmware/govmomi/issues/192): HostSystem doesn't seem to be returning the correct host. [1cbe968d]
- fix a problem of ignored https_proxy environment variable with https scheme [116a4044]
- Add create and destroy datacenter to govc. [df423c32]
- Usage for devices.{cdrom,floppy}.* [035bd12c]
- make storage resource manager [68e50dd3]
- Specify default network in test helper [b28d6f42]
- Fix boot order test [4b388e67]
- Expand vm.vnc command [4414a07e]
- rename the session file for windows naming check [e329e6e7]
- use filepath for filesystem related path operations [706520fa]
- Add -f flag to datastore.rm [ceb35f13]
- Default VM memory to 1GiB [6498890f]
- Include description for device.cdrom commands [591b74f4]
- Add usage to device.cdrom.insert [815f0286]
- Flag description casing [f2209c2b]
- Add usage to import commands [5e52668c]
- Expand datastore.ls [23cf4d35]
- Expose underlying VimFault through Fault() function [bca8ef73]
- Add Usage() function to subset of commands [90edb2bc]
- Implement subset of license manager [afdc145a]
- Add net.adapter option to network flag [14765d07]
- Add CreateEthernetCard method [18c2cce0]
- Don't run vm.destroy if there is no input [9b2730f0]
- Add new ops to vm.power command [611ced85]
- Add VM power ops [6cd9f466]
- Work on README [7918063c]
- Check minimum API version from client flag [db17cddd]
- Don't run datastore.rm if there is no input [df075430]
- Move environment variables names into constants [e49a6d57]
- Add device.scsi command [2cfe267f]
- Support scsi disk.controller type in vm.create [6df44c1a]
- Add CreateSCSIController method [39a60bbf]
- Rename vm.create disk.adapter to disk.controller [136fabe5]
- Change disk related commands to use new helpers [9c51314c]
- Add VirtualDisk support to device helpers [b0c895e5]
- Add helpers for creating disks [a00f4545]
- Add FindDiskController helper [16283936]
- Add VirtualDeviceList.FindSCSIController method [dda056dc]
- FindByBackingInfo -> SelectByBackingInfo [5402017a]
- Add vm disk related bats tests [0ff5759c]
- Output disk file backing in device.info [8f1e183a]
- Remove datastore test files [e7cfba4b]
- Use DeviceAdd helper in vm.network.add command [6b883be5]
- Use device name in vm.network.change command [eb5881ae]
- Remove vm.network.remove command [b7503468]
- Add vm.network.change cli command [0b81619a]
- Use VirtualDeviceList helpers in vm.network.remove [0af5c4cf]
- Add VirtualDeviceList FindByBackingInfo method [94c62da0]
- Move govc resource finders to govmomi/find package [c247b80c]
- Add vm.info bats test [5f0c8dd4]
- mv govc/flags/list -> govmomi/list [5d99454d]
- Fix HostSystem.ResourcePool with cluster parent [028bd3ff]
- Add ls bats test [48e25166]
- Add host bats test [c5f24bce]
- Add default GOVC_HOST to vcsim_env [f965c9ad]
- Add network flag required test [77fc8ade]
- Add wrapper to manually run govc against vcsim [b1236bf8]
- Fix network device.remove test [68831a1f]
- Default vcsim box to 4G memory [4649bf1f]
- Simplify vcsim_env helper [b3f71333]
- Answer pending VM questions from govc [2ca11cde]
- Move govc/test Vagrant boxes [b6c3ff31]
- Change network flag to use NetworkReference [b1b5b26e]
- Add network bats test [83f49af7]
- Add NetworkReference interface [a8ffa576]
- Add vcsim_env helper [6fe62e29]
- Fix collapse_ws helper [a616817d]
- Add DistributedVirtualPortgroup constructor [0614961e]
- Cache esxcli command info [1ddf6801]
- Add table formatter to esxcli command [c713b974]
- Include esxcli method info in response [fd19a011]
- Explicit exit status check in assert_failure [3c9a436f]
- Collapse whitespace in assert_line helper [5a63bc06]
- Change vm.ip -esxcli to wait for ip [c9bd4312]
- boot order test fixups [e97e5604]
- 32M is plenty of memory for ttylinux [0e128e0d]
- Add test cleanup script [85ded933]
- Add device.serial cli commands [2bc707e7]
- Add serial port device related helpers [17fb283a]
- Add device.boot tests [d9b846d1]
- Add device.floppy cli commands [b5a21e4e]
- Add floppy device related helpers [d1d39fc3]
- Refactor disk logic into disk.go [1e2c54c0]
- Fix attach disk error checks [9dff8e74]
- Add vm.disk.attach [0f352ec3]
- Refactor vm.disk.add to vm.disk.create [bdd7b37b]
- Add govc functional tests [ae2e990e]
- Fix alignment for 32-bit go [a707fae6]
- Default cli client url.User field [13274292]
- Add device.boot cli command [17df67ad]
- Add device.ls -boot option [3c345ad7]
- Add boot order related VirtualDeviceList helpers [3b25234c]
- Add VirtualMachine BootOptions wrappers [f996c7d0]
- Add some DeviceType constants [4f3b935b]
- Add VirtualDeviceList.Type method [86f90c52]
- Output MAC Address in device.info [5f3b95d7]
- Add VirtualMachineList.PrimaryMacAddress helper [58c3c64e]
- Fix import.ovf with relative ovf source path [67fea291]
- Support non-disk files in import.ovf [22602029]
- Add Upload.Headers field [92175548]
- Fix import.ova command [f095536d]
- Add device related govc commands [5093303a]
- Add device list related helpers [18644254]
- Add device list helpers [6803033e]
- Switch to BaseOptionValue for vm extra config [4f8cd87c]
- Regenerate types [76662657]
- Generate interface types for all base types [46ec389f]
- Remove Client param from ResourcePool methods [f78df469]
- Add Client reference to ResourcePool [ca3cd417]
- Add Client reference to Network [ffc306cc]
- Remove Client param from HttpNfcLease methods [c1138fc4]
- Add Client reference to HttpNfcLease [6f983a49]
- Remove Client param from HostSystem methods [d2d566d0]
- Add Client reference to HostSystem [60bf1770]
- Remove Client param from HostDatastoreBrowser methods [e32542c1]
- Add Client reference to HostDatastoreBrowser [8956959a]
- Remove Client param from Folder methods [79e7da1d]
- Add Client reference to Folder [68b3e6dc]
- Remove Client param from Datastore methods [da5b8ec0]
- Add Client reference to Datastore [f89dd25a]
- Remove Client param from Datacenter methods [1b372efa]
- Add Client reference to Datacenter [ce320403]
- Remove Client param from VirtualMachine methods [b99a9529]
- Add Client reference to VirtualMachine [eb700d65]
- Remove config check from esxcli.GuestInfo.IpAddress [673485e4]
- Add VCSA Vagrant box [667df16a]
- Use single consistent pattern to populate FlagSet [66b7daab]
- Export NewReference function [8fa06b5a]
- Check if info is nil before using it [a4e11a3a]
- Add ManagedObject wrappers [8bbe7361]
- Add vm.ip -esxcli option [9d5df71d]
- Add esxcli helper for guest related info [1818a2a6]
- Use vim.CLIInfo for esxcli command flags and help [ac6efdc9]
- Remove Cdrom function from disk flag [5b9b34bc]
- Use new esxcli command parser [01d201ee]
- New esxcli command parser [7531d60e]
- Refactor esxcli to esxcli.Executor [a27c9bd5]
- Refactor unmarshal [fdb2d2d0]
- Add esxcli related types and methods [2dd9910d]
- Add IsoFlag [aad819e8]
- Handle empty values in esxcli [df11fc04]
- Fix default network in NetworkFlag [6ceff6a4]
- Add DistributedVirtualPortgroup wrapper [bc39649d]
- Add DVS support to NetworkFlag [a7eb1d1e]
- Support DistributedVirtualPortgroup in lister [71898a73]
- Regenerate mo types [1cf31f03]
- Generate mo types regardless of props [1e7c1957]
- tasks are no longer generated [549a2712]
- Remove unused DiskFlag.Copy method [fcf2cd94]
- Add DiskFlag adpater option [e494c312]
