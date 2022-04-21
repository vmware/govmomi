# govc usage

This document is generated from `govc -h` and `govc $cmd -h` commands.

The following common options are filtered out in this document,
but appear via `govc $cmd -h`:

```
  -cert=                    Certificate [GOVC_CERTIFICATE]
  -debug=false              Store debug logs [GOVC_DEBUG]
  -trace=false              Write SOAP/REST traffic to stderr
  -verbose=false            Write request/response data to stderr
  -dump=false               Enable output dump
  -json=false               Enable JSON output
  -xml=false                Enable XML output
  -k=false                  Skip verification of server certificate [GOVC_INSECURE]
  -key=                     Private key [GOVC_PRIVATE_KEY]
  -persist-session=true     Persist session to disk [GOVC_PERSIST_SESSION]
  -tls-ca-certs=            TLS CA certificates file [GOVC_TLS_CA_CERTS]
  -tls-known-hosts=         TLS known hosts file [GOVC_TLS_KNOWN_HOSTS]
  -u=                       ESX or vCenter URL [GOVC_URL]
  -vim-namespace=urn:vim25  Vim namespace [GOVC_VIM_NAMESPACE]
  -vim-version=6.0          Vim version [GOVC_VIM_VERSION]
  -dc=                      Datacenter [GOVC_DATACENTER]
  -host.dns=                Find host by FQDN
  -host.ip=                 Find host by IP address
  -host.ipath=              Find host by inventory path
  -host.uuid=               Find host by UUID
  -vm.dns=                  Find VM by FQDN
  -vm.ip=                   Find VM by IP address
  -vm.ipath=                Find VM by inventory path
  -vm.path=                 Find VM by path to .vmx file
  -vm.uuid=                 Find VM by UUID
```

<details><summary>Contents</summary>

 - [about](#about)
 - [about.cert](#aboutcert)
 - [cluster.add](#clusteradd)
 - [cluster.change](#clusterchange)
 - [cluster.create](#clustercreate)
 - [cluster.group.change](#clustergroupchange)
 - [cluster.group.create](#clustergroupcreate)
 - [cluster.group.ls](#clustergroupls)
 - [cluster.group.remove](#clustergroupremove)
 - [cluster.module.create](#clustermodulecreate)
 - [cluster.module.ls](#clustermodulels)
 - [cluster.module.rm](#clustermodulerm)
 - [cluster.module.vm.add](#clustermodulevmadd)
 - [cluster.module.vm.rm](#clustermodulevmrm)
 - [cluster.override.change](#clusteroverridechange)
 - [cluster.override.info](#clusteroverrideinfo)
 - [cluster.override.remove](#clusteroverrideremove)
 - [cluster.rule.change](#clusterrulechange)
 - [cluster.rule.create](#clusterrulecreate)
 - [cluster.rule.info](#clusterruleinfo)
 - [cluster.rule.ls](#clusterrulels)
 - [cluster.rule.remove](#clusterruleremove)
 - [cluster.stretch](#clusterstretch)
 - [cluster.usage](#clusterusage)
 - [datacenter.create](#datacentercreate)
 - [datacenter.info](#datacenterinfo)
 - [datastore.cluster.change](#datastoreclusterchange)
 - [datastore.cluster.info](#datastoreclusterinfo)
 - [datastore.cp](#datastorecp)
 - [datastore.create](#datastorecreate)
 - [datastore.disk.create](#datastorediskcreate)
 - [datastore.disk.inflate](#datastorediskinflate)
 - [datastore.disk.info](#datastorediskinfo)
 - [datastore.disk.shrink](#datastorediskshrink)
 - [datastore.download](#datastoredownload)
 - [datastore.info](#datastoreinfo)
 - [datastore.ls](#datastorels)
 - [datastore.maintenance.enter](#datastoremaintenanceenter)
 - [datastore.maintenance.exit](#datastoremaintenanceexit)
 - [datastore.mkdir](#datastoremkdir)
 - [datastore.mv](#datastoremv)
 - [datastore.remove](#datastoreremove)
 - [datastore.rm](#datastorerm)
 - [datastore.tail](#datastoretail)
 - [datastore.upload](#datastoreupload)
 - [datastore.vsan.dom.ls](#datastorevsandomls)
 - [datastore.vsan.dom.rm](#datastorevsandomrm)
 - [device.boot](#deviceboot)
 - [device.cdrom.add](#devicecdromadd)
 - [device.cdrom.eject](#devicecdromeject)
 - [device.cdrom.insert](#devicecdrominsert)
 - [device.connect](#deviceconnect)
 - [device.disconnect](#devicedisconnect)
 - [device.floppy.add](#devicefloppyadd)
 - [device.floppy.eject](#devicefloppyeject)
 - [device.floppy.insert](#devicefloppyinsert)
 - [device.info](#deviceinfo)
 - [device.ls](#devicels)
 - [device.pci.add](#devicepciadd)
 - [device.pci.ls](#devicepcils)
 - [device.pci.remove](#devicepciremove)
 - [device.remove](#deviceremove)
 - [device.scsi.add](#devicescsiadd)
 - [device.serial.add](#deviceserialadd)
 - [device.serial.connect](#deviceserialconnect)
 - [device.serial.disconnect](#deviceserialdisconnect)
 - [device.usb.add](#deviceusbadd)
 - [disk.create](#diskcreate)
 - [disk.ls](#diskls)
 - [disk.register](#diskregister)
 - [disk.rm](#diskrm)
 - [disk.snapshot.create](#disksnapshotcreate)
 - [disk.snapshot.ls](#disksnapshotls)
 - [disk.snapshot.rm](#disksnapshotrm)
 - [disk.tags.attach](#disktagsattach)
 - [disk.tags.detach](#disktagsdetach)
 - [dvs.add](#dvsadd)
 - [dvs.change](#dvschange)
 - [dvs.create](#dvscreate)
 - [dvs.portgroup.add](#dvsportgroupadd)
 - [dvs.portgroup.change](#dvsportgroupchange)
 - [dvs.portgroup.info](#dvsportgroupinfo)
 - [env](#env)
 - [events](#events)
 - [export.ovf](#exportovf)
 - [extension.info](#extensioninfo)
 - [extension.register](#extensionregister)
 - [extension.setcert](#extensionsetcert)
 - [extension.unregister](#extensionunregister)
 - [fields.add](#fieldsadd)
 - [fields.info](#fieldsinfo)
 - [fields.ls](#fieldsls)
 - [fields.rename](#fieldsrename)
 - [fields.rm](#fieldsrm)
 - [fields.set](#fieldsset)
 - [find](#find)
 - [firewall.ruleset.find](#firewallrulesetfind)
 - [folder.create](#foldercreate)
 - [folder.info](#folderinfo)
 - [guest.chmod](#guestchmod)
 - [guest.chown](#guestchown)
 - [guest.df](#guestdf)
 - [guest.download](#guestdownload)
 - [guest.getenv](#guestgetenv)
 - [guest.kill](#guestkill)
 - [guest.ls](#guestls)
 - [guest.mkdir](#guestmkdir)
 - [guest.mktemp](#guestmktemp)
 - [guest.mv](#guestmv)
 - [guest.ps](#guestps)
 - [guest.rm](#guestrm)
 - [guest.rmdir](#guestrmdir)
 - [guest.run](#guestrun)
 - [guest.start](#gueststart)
 - [guest.touch](#guesttouch)
 - [guest.upload](#guestupload)
 - [host.account.create](#hostaccountcreate)
 - [host.account.remove](#hostaccountremove)
 - [host.account.update](#hostaccountupdate)
 - [host.add](#hostadd)
 - [host.autostart.add](#hostautostartadd)
 - [host.autostart.configure](#hostautostartconfigure)
 - [host.autostart.info](#hostautostartinfo)
 - [host.autostart.remove](#hostautostartremove)
 - [host.cert.csr](#hostcertcsr)
 - [host.cert.import](#hostcertimport)
 - [host.cert.info](#hostcertinfo)
 - [host.date.change](#hostdatechange)
 - [host.date.info](#hostdateinfo)
 - [host.disconnect](#hostdisconnect)
 - [host.esxcli](#hostesxcli)
 - [host.info](#hostinfo)
 - [host.maintenance.enter](#hostmaintenanceenter)
 - [host.maintenance.exit](#hostmaintenanceexit)
 - [host.option.ls](#hostoptionls)
 - [host.option.set](#hostoptionset)
 - [host.portgroup.add](#hostportgroupadd)
 - [host.portgroup.change](#hostportgroupchange)
 - [host.portgroup.info](#hostportgroupinfo)
 - [host.portgroup.remove](#hostportgroupremove)
 - [host.reconnect](#hostreconnect)
 - [host.remove](#hostremove)
 - [host.service](#hostservice)
 - [host.service.ls](#hostservicels)
 - [host.shutdown](#hostshutdown)
 - [host.storage.info](#hoststorageinfo)
 - [host.storage.mark](#hoststoragemark)
 - [host.storage.partition](#hoststoragepartition)
 - [host.vnic.change](#hostvnicchange)
 - [host.vnic.info](#hostvnicinfo)
 - [host.vnic.service](#hostvnicservice)
 - [host.vswitch.add](#hostvswitchadd)
 - [host.vswitch.info](#hostvswitchinfo)
 - [host.vswitch.remove](#hostvswitchremove)
 - [import.ova](#importova)
 - [import.ovf](#importovf)
 - [import.spec](#importspec)
 - [import.vmdk](#importvmdk)
 - [library.checkin](#librarycheckin)
 - [library.checkout](#librarycheckout)
 - [library.clone](#libraryclone)
 - [library.cp](#librarycp)
 - [library.create](#librarycreate)
 - [library.deploy](#librarydeploy)
 - [library.export](#libraryexport)
 - [library.import](#libraryimport)
 - [library.info](#libraryinfo)
 - [library.ls](#libraryls)
 - [library.publish](#librarypublish)
 - [library.rm](#libraryrm)
 - [library.session.ls](#librarysessionls)
 - [library.session.rm](#librarysessionrm)
 - [library.subscriber.create](#librarysubscribercreate)
 - [library.subscriber.info](#librarysubscriberinfo)
 - [library.subscriber.ls](#librarysubscriberls)
 - [library.subscriber.rm](#librarysubscriberrm)
 - [library.sync](#librarysync)
 - [library.update](#libraryupdate)
 - [library.vmtx.info](#libraryvmtxinfo)
 - [license.add](#licenseadd)
 - [license.assign](#licenseassign)
 - [license.assigned.ls](#licenseassignedls)
 - [license.decode](#licensedecode)
 - [license.label.set](#licenselabelset)
 - [license.ls](#licensels)
 - [license.remove](#licenseremove)
 - [logs](#logs)
 - [logs.download](#logsdownload)
 - [logs.ls](#logsls)
 - [ls](#ls)
 - [metric.change](#metricchange)
 - [metric.info](#metricinfo)
 - [metric.interval.change](#metricintervalchange)
 - [metric.interval.info](#metricintervalinfo)
 - [metric.ls](#metricls)
 - [metric.reset](#metricreset)
 - [metric.sample](#metricsample)
 - [namespace.cluster.disable](#namespaceclusterdisable)
 - [namespace.cluster.enable](#namespaceclusterenable)
 - [namespace.cluster.ls](#namespaceclusterls)
 - [namespace.logs.download](#namespacelogsdownload)
 - [namespace.service.activate](#namespaceserviceactivate)
 - [namespace.service.create](#namespaceservicecreate)
 - [namespace.service.deactivate](#namespaceservicedeactivate)
 - [namespace.service.info](#namespaceserviceinfo)
 - [namespace.service.ls](#namespaceservicels)
 - [namespace.service.rm](#namespaceservicerm)
 - [object.collect](#objectcollect)
 - [object.destroy](#objectdestroy)
 - [object.method](#objectmethod)
 - [object.mv](#objectmv)
 - [object.reload](#objectreload)
 - [object.rename](#objectrename)
 - [object.save](#objectsave)
 - [option.ls](#optionls)
 - [option.set](#optionset)
 - [permissions.ls](#permissionsls)
 - [permissions.remove](#permissionsremove)
 - [permissions.set](#permissionsset)
 - [pool.change](#poolchange)
 - [pool.create](#poolcreate)
 - [pool.destroy](#pooldestroy)
 - [pool.info](#poolinfo)
 - [role.create](#rolecreate)
 - [role.ls](#rolels)
 - [role.remove](#roleremove)
 - [role.update](#roleupdate)
 - [role.usage](#roleusage)
 - [session.login](#sessionlogin)
 - [session.logout](#sessionlogout)
 - [session.ls](#sessionls)
 - [session.rm](#sessionrm)
 - [snapshot.create](#snapshotcreate)
 - [snapshot.remove](#snapshotremove)
 - [snapshot.revert](#snapshotrevert)
 - [snapshot.tree](#snapshottree)
 - [sso.group.create](#ssogroupcreate)
 - [sso.group.ls](#ssogroupls)
 - [sso.group.rm](#ssogrouprm)
 - [sso.group.update](#ssogroupupdate)
 - [sso.idp.ls](#ssoidpls)
 - [sso.service.ls](#ssoservicels)
 - [sso.user.create](#ssousercreate)
 - [sso.user.id](#ssouserid)
 - [sso.user.ls](#ssouserls)
 - [sso.user.rm](#ssouserrm)
 - [sso.user.update](#ssouserupdate)
 - [storage.policy.create](#storagepolicycreate)
 - [storage.policy.info](#storagepolicyinfo)
 - [storage.policy.ls](#storagepolicyls)
 - [storage.policy.rm](#storagepolicyrm)
 - [tags.attach](#tagsattach)
 - [tags.attached.ls](#tagsattachedls)
 - [tags.category.create](#tagscategorycreate)
 - [tags.category.info](#tagscategoryinfo)
 - [tags.category.ls](#tagscategoryls)
 - [tags.category.rm](#tagscategoryrm)
 - [tags.category.update](#tagscategoryupdate)
 - [tags.create](#tagscreate)
 - [tags.detach](#tagsdetach)
 - [tags.info](#tagsinfo)
 - [tags.ls](#tagsls)
 - [tags.rm](#tagsrm)
 - [tags.update](#tagsupdate)
 - [task.cancel](#taskcancel)
 - [tasks](#tasks)
 - [tree](#tree)
 - [vapp.destroy](#vappdestroy)
 - [vapp.power](#vapppower)
 - [vcsa.access.consolecli.get](#vcsaaccessconsolecliget)
 - [vcsa.access.consolecli.set](#vcsaaccessconsolecliset)
 - [vcsa.access.dcui.get](#vcsaaccessdcuiget)
 - [vcsa.access.dcui.set](#vcsaaccessdcuiset)
 - [vcsa.access.shell.get](#vcsaaccessshellget)
 - [vcsa.access.shell.set](#vcsaaccessshellset)
 - [vcsa.access.ssh.get](#vcsaaccesssshget)
 - [vcsa.access.ssh.set](#vcsaaccesssshset)
 - [vcsa.log.forwarding.info](#vcsalogforwardinginfo)
 - [vcsa.net.proxy.info](#vcsanetproxyinfo)
 - [vcsa.shutdown.cancel](#vcsashutdowncancel)
 - [vcsa.shutdown.get](#vcsashutdownget)
 - [vcsa.shutdown.poweroff](#vcsashutdownpoweroff)
 - [vcsa.shutdown.reboot](#vcsashutdownreboot)
 - [version](#version)
 - [vm.change](#vmchange)
 - [vm.clone](#vmclone)
 - [vm.console](#vmconsole)
 - [vm.create](#vmcreate)
 - [vm.customize](#vmcustomize)
 - [vm.destroy](#vmdestroy)
 - [vm.disk.attach](#vmdiskattach)
 - [vm.disk.change](#vmdiskchange)
 - [vm.disk.create](#vmdiskcreate)
 - [vm.guest.tools](#vmguesttools)
 - [vm.info](#vminfo)
 - [vm.instantclone](#vminstantclone)
 - [vm.ip](#vmip)
 - [vm.keystrokes](#vmkeystrokes)
 - [vm.markastemplate](#vmmarkastemplate)
 - [vm.markasvm](#vmmarkasvm)
 - [vm.migrate](#vmmigrate)
 - [vm.network.add](#vmnetworkadd)
 - [vm.network.change](#vmnetworkchange)
 - [vm.option.info](#vmoptioninfo)
 - [vm.power](#vmpower)
 - [vm.question](#vmquestion)
 - [vm.rdm.attach](#vmrdmattach)
 - [vm.rdm.ls](#vmrdmls)
 - [vm.register](#vmregister)
 - [vm.unregister](#vmunregister)
 - [vm.upgrade](#vmupgrade)
 - [vm.vnc](#vmvnc)
 - [volume.ls](#volumels)
 - [volume.rm](#volumerm)
 - [vsan.change](#vsanchange)
 - [vsan.info](#vsaninfo)

</details>

## about

```
Usage: govc about [OPTIONS]

Display About info for HOST.

System information including the name, type, version, and build number.

Examples:
  govc about
  govc about -json | jq -r .About.ProductLineId

Options:
  -c=false               Include client info
  -l=false               Include service content
```

## about.cert

```
Usage: govc about.cert [OPTIONS]

Display TLS certificate info for HOST.

If the HOST certificate cannot be verified, about.cert will return with exit code 60 (as curl does).
If the '-k' flag is provided, about.cert will return with exit code 0 in this case.
The SHA1 thumbprint can also be used as '-thumbprint' for the 'host.add' and 'cluster.add' commands.

Examples:
  govc about.cert -k -json | jq -r .ThumbprintSHA1
  govc about.cert -k -show | sudo tee /usr/local/share/ca-certificates/host.crt
  govc about.cert -k -thumbprint | tee -a ~/.govmomi/known_hosts

Options:
  -show=false            Show PEM encoded server certificate only
  -thumbprint=false      Output host hash and thumbprint only
```

## cluster.add

```
Usage: govc cluster.add [OPTIONS]

Add HOST to CLUSTER.

The host is added to the cluster specified by the 'cluster' flag.

Examples:
  thumbprint=$(govc about.cert -k -u host.example.com -thumbprint | awk '{print $2}')
  govc cluster.add -cluster ClusterA -hostname host.example.com -username root -password pass -thumbprint $thumbprint
  govc cluster.add -cluster ClusterB -hostname 10.0.6.1 -username root -password pass -noverify

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
  -connect=true          Immediately connect to host
  -force=false           Force when host is managed by another VC
  -hostname=             Hostname or IP address of the host
  -license=              Assign license key
  -noverify=false        Accept host thumbprint without verification
  -password=             Password of administration account on the host
  -thumbprint=           SHA-1 thumbprint of the host's SSL certificate
  -username=             Username of administration account on the host
```

## cluster.change

```
Usage: govc cluster.change [OPTIONS] CLUSTER...

Change configuration of the given clusters.

Examples:
  govc cluster.change -drs-enabled -vsan-enabled -vsan-autoclaim ClusterA
  govc cluster.change -drs-enabled=false ClusterB
  govc cluster.change -drs-vmotion-rate=4 ClusterC

Options:
  -drs-enabled=<nil>     Enable DRS
  -drs-mode=             DRS behavior for virtual machines: manual, partiallyAutomated, fullyAutomated
  -drs-vmotion-rate=0    Aggressiveness of vMotions (1-5)
  -ha-enabled=<nil>      Enable HA
  -vsan-autoclaim=<nil>  Autoclaim storage on cluster hosts
  -vsan-enabled=<nil>    Enable vSAN
```

## cluster.create

```
Usage: govc cluster.create [OPTIONS] CLUSTER

Create CLUSTER in datacenter.

The cluster is added to the folder specified by the 'folder' flag. If not given,
this defaults to the host folder in the specified or default datacenter.

Examples:
  govc cluster.create ClusterA
  govc cluster.create -folder /dc2/test-folder ClusterB

Options:
  -folder=               Inventory folder [GOVC_FOLDER]
```

## cluster.group.change

```
Usage: govc cluster.group.change [OPTIONS] NAME...

Set cluster group members.

Examples:
  govc cluster.group.change -name my_group vm_a vm_b vm_c # set
  govc cluster.group.change -name my_group vm_a vm_b vm_c $(govc cluster.group.ls -name my_group) vm_d # add
  govc cluster.group.ls -name my_group | grep -v vm_b | xargs govc cluster.group.change -name my_group vm_a vm_b vm_c # remove

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
  -name=                 Cluster group name
```

## cluster.group.create

```
Usage: govc cluster.group.create [OPTIONS]

Create cluster group.

One of '-vm' or '-host' must be provided to specify the group type.

Examples:
  govc cluster.group.create -name my_vm_group -vm vm_a vm_b vm_c
  govc cluster.group.create -name my_host_group -host host_a host_b host_c

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
  -host=false            Create cluster Host group
  -name=                 Cluster group name
  -vm=false              Create cluster VM group
```

## cluster.group.ls

```
Usage: govc cluster.group.ls [OPTIONS]

List cluster groups and group members.

Examples:
  govc cluster.group.ls -cluster my_cluster
  govc cluster.group.ls -cluster my_cluster -l | grep ClusterHostGroup
  govc cluster.group.ls -cluster my_cluster -name my_group

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
  -l=false               Long listing format
  -name=                 Cluster group name
```

## cluster.group.remove

```
Usage: govc cluster.group.remove [OPTIONS]

Remove cluster group.

Examples:
  govc cluster.group.remove -cluster my_cluster -name my_group

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
  -name=                 Cluster group name
```

## cluster.module.create

```
Usage: govc cluster.module.create [OPTIONS]

Create cluster module.

This command will output the ID of the new module.

Examples:
  govc cluster.module.create -cluster my_cluster

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
```

## cluster.module.ls

```
Usage: govc cluster.module.ls [OPTIONS]

List cluster modules.

When -id is specified, that module's members are listed.

Examples:
  govc cluster.module.ls
  govc cluster.module.ls -json | jq .
  govc cluster.module.ls -id module_id

Options:
  -id=                   Module ID
```

## cluster.module.rm

```
Usage: govc cluster.module.rm [OPTIONS] ID

Delete cluster module ID.

Examples:
  govc cluster.module.rm module_id

Options:
```

## cluster.module.vm.add

```
Usage: govc cluster.module.vm.add [OPTIONS] VM...

Add VM(s) to a cluster module.

Examples:
  govc cluster.module.vm.add -id module_id $vm...

Options:
  -id=                   Module ID
```

## cluster.module.vm.rm

```
Usage: govc cluster.module.vm.rm [OPTIONS] VM...

Remove VM(s) from a cluster module.

Examples:
  govc cluster.module.vm.rm -id module_id $vm...

Options:
  -id=                   Module ID
```

## cluster.override.change

```
Usage: govc cluster.override.change [OPTIONS]

Change cluster VM overrides.

Examples:
  govc cluster.override.change -cluster cluster_1 -vm vm_1 -drs-enabled=false
  govc cluster.override.change -cluster cluster_1 -vm vm_2 -drs-enabled -drs-mode fullyAutomated
  govc cluster.override.change -cluster cluster_1 -vm vm_3 -ha-restart-priority high
  govc cluster.override.change -cluster cluster_1 -vm vm_4 -ha-additional-delay 30
  govc cluster.override.change -cluster cluster_1 -vm vm_5 -ha-ready-condition poweredOn

Options:
  -cluster=               Cluster [GOVC_CLUSTER]
  -drs-enabled=<nil>      Enable DRS
  -drs-mode=              DRS behavior for virtual machines: manual, partiallyAutomated, fullyAutomated
  -ha-additional-delay=0  HA Additional Delay
  -ha-ready-condition=    HA VM Ready Condition (Start next priority VMs when): poweredOn, guestHbStatusGreen, appHbStatusGreen, useClusterDefault
  -ha-restart-priority=   HA restart priority: disabled, lowest, low, medium, high, highest
  -vm=                    Virtual machine [GOVC_VM]
```

## cluster.override.info

```
Usage: govc cluster.override.info [OPTIONS]

Cluster VM overrides info.

Examples:
  govc cluster.override.info
  govc cluster.override.info -json

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
```

## cluster.override.remove

```
Usage: govc cluster.override.remove [OPTIONS]

Remove cluster VM overrides.

Examples:
  govc cluster.override.remove -cluster cluster_1 -vm vm_1

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
  -vm=                   Virtual machine [GOVC_VM]
```

## cluster.rule.change

```
Usage: govc cluster.rule.change [OPTIONS] NAME...

Change cluster rule.

Examples:
  govc cluster.rule.change -cluster my_cluster -name my_rule -enable=false

Options:
  -cluster=                 Cluster [GOVC_CLUSTER]
  -enable=<nil>             Enable rule
  -host-affine-group=       Host affine group name
  -host-anti-affine-group=  Host anti-affine group name
  -l=false                  Long listing format
  -mandatory=<nil>          Enforce rule compliance
  -name=                    Cluster rule name
  -vm-group=                VM group name
```

## cluster.rule.create

```
Usage: govc cluster.rule.create [OPTIONS] NAME...

Create cluster rule.

Rules are not enabled by default, use the 'enable' flag to enable upon creation or cluster.rule.change after creation.

One of '-affinity', '-anti-affinity', '-depends' or '-vm-host' must be provided to specify the rule type.

With '-affinity' or '-anti-affinity', at least 2 vm NAME arguments must be specified.

With '-depends', vm group NAME and vm group dependency NAME arguments must be specified.

With '-vm-host', use the '-vm-group' flag combined with the '-host-affine-group' and/or '-host-anti-affine-group' flags.

Examples:
  govc cluster.rule.create -name pod1 -enable -affinity vm_a vm_b vm_c
  govc cluster.rule.create -name pod2 -enable -anti-affinity vm_d vm_e vm_f
  govc cluster.rule.create -name pod3 -enable -mandatory -vm-host -vm-group my_vms -host-affine-group my_hosts
  govc cluster.rule.create -name pod4 -depends vm_group_app vm_group_db

Options:
  -affinity=false           Keep Virtual Machines Together
  -anti-affinity=false      Separate Virtual Machines
  -cluster=                 Cluster [GOVC_CLUSTER]
  -depends=false            Virtual Machines to Virtual Machines
  -enable=<nil>             Enable rule
  -host-affine-group=       Host affine group name
  -host-anti-affine-group=  Host anti-affine group name
  -l=false                  Long listing format
  -mandatory=<nil>          Enforce rule compliance
  -name=                    Cluster rule name
  -vm-group=                VM group name
  -vm-host=false            Virtual Machines to Hosts
```

## cluster.rule.info

```
Usage: govc cluster.rule.info [OPTIONS]

Provides detailed infos about cluster rules, their types and rule members.

Examples:
  govc cluster.rule.info -cluster my_cluster
  govc cluster.rule.info -cluster my_cluster -name my_rule

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
  -l=false               Long listing format
  -name=                 Cluster rule name
```

## cluster.rule.ls

```
Usage: govc cluster.rule.ls [OPTIONS]

List cluster rules and rule members.

Examples:
  govc cluster.rule.ls -cluster my_cluster
  govc cluster.rule.ls -cluster my_cluster -name my_rule
  govc cluster.rule.ls -cluster my_cluster -l
  govc cluster.rule.ls -cluster my_cluster -name my_rule -l

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
  -l=false               Long listing format
  -name=                 Cluster rule name
```

## cluster.rule.remove

```
Usage: govc cluster.rule.remove [OPTIONS]

Remove cluster rule.

Examples:
  govc cluster.group.remove -cluster my_cluster -name my_rule

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
  -l=false               Long listing format
  -name=                 Cluster rule name
```

## cluster.stretch

```
Usage: govc cluster.stretch [OPTIONS] CLUSTER

Convert a vSAN cluster into a stretched cluster

The vSAN cluster is converted to a stretched cluster with a witness host
specified by the 'witness' flag.  The datastore hosts are placed into one
of two fault domains that are specified in each host list. The name of the
preferred fault domain can be specified by the 'preferred-fault-domain' flag.

Examples:
  govc cluster.stretch -dc remote-site-1 \
    -witness /dc-name/host/192.168.112.2 \
    -first-fault-domain-hosts 192.168.113.121 \
    -second-fault-domain-hosts 192.168.113.45,192.168.113.70 \
    cluster-name

Options:
  -first-fault-domain-hosts=           Hosts to place in the first fault domain
  -first-fault-domain-name=Primary     Name of the first fault domain
  -preferred-fault-domain=Primary      Name of the preferred fault domain
  -second-fault-domain-hosts=          Hosts to place in the second fault domain
  -second-fault-domain-name=Secondary  Name of the second fault domain
  -witness=                            Witness host for the stretched cluster
```

## cluster.usage

```
Usage: govc cluster.usage [OPTIONS] CLUSTER

Cluster resource usage summary.

Examples:
  govc cluster.usage ClusterName
  govc cluster.usage -S ClusterName # summarize shared storage only
  govc cluster.usage -json ClusterName | jq -r .CPU.Summary.Usage

Options:
  -S=false               Exclude host local storage
```

## datacenter.create

```
Usage: govc datacenter.create [OPTIONS] NAME...

Create datacenter NAME.

Examples:
  govc datacenter.create MyDC # create
  govc object.destroy /MyDC   # delete

Options:
  -folder=               Inventory folder [GOVC_FOLDER]
```

## datacenter.info

```
Usage: govc datacenter.info [OPTIONS] [PATH]...

Options:
```

## datastore.cluster.change

```
Usage: govc datastore.cluster.change [OPTIONS] CLUSTER...

Change configuration of the given datastore clusters.

Examples:
  govc datastore.cluster.change -drs-enabled ClusterA
  govc datastore.cluster.change -drs-enabled=false ClusterB

Options:
  -drs-enabled=<nil>     Enable Storage DRS
  -drs-mode=             Storage DRS behavior: manual, automated
```

## datastore.cluster.info

```
Usage: govc datastore.cluster.info [OPTIONS] [PATH]...

Display datastore cluster info.

Examples:
  govc datastore.cluster.info
  govc datastore.cluster.info MyDatastoreCluster

Options:
```

## datastore.cp

```
Usage: govc datastore.cp [OPTIONS] SRC DST

Copy SRC to DST on DATASTORE.

Examples:
  govc datastore.cp foo/foo.vmx foo/foo.vmx.old
  govc datastore.cp -f my.vmx foo/foo.vmx
  govc datastore.cp disks/disk1.vmdk disks/disk2.vmdk
  govc datastore.cp disks/disk1.vmdk -dc-target DC2 disks/disk2.vmdk
  govc datastore.cp disks/disk1.vmdk -ds-target NFS-2 disks/disk2.vmdk

Options:
  -dc-target=            Datacenter destination (defaults to -dc)
  -ds=                   Datastore [GOVC_DATASTORE]
  -ds-target=            Datastore destination (defaults to -ds)
  -f=false               If true, overwrite any identically named file at the destination
  -t=true                Use file type to choose disk or file manager
```

## datastore.create

```
Usage: govc datastore.create [OPTIONS] HOST...

Create datastore on HOST.

Examples:
  govc datastore.create -type nfs -name nfsDatastore -remote-host 10.143.2.232 -remote-path /share cluster1
  govc datastore.create -type vmfs -name vmfsDatastore -disk=mpx.vmhba0:C0:T0:L0 cluster1
  govc datastore.create -type local -name localDatastore -path /var/datastore host1

Options:
  -disk=                 Canonical name of disk (VMFS only)
  -force=false           Ignore DuplicateName error if datastore is already mounted on a host
  -host=                 Host system [GOVC_HOST]
  -mode=readOnly         Access mode for the mount point (readOnly|readWrite)
  -name=                 Datastore name
  -password=             Password to use when connecting (CIFS only)
  -path=                 Local directory path for the datastore (local only)
  -remote-host=          Remote hostname of the NAS datastore
  -remote-path=          Remote path of the NFS mount point
  -type=                 Datastore type (NFS|NFS41|CIFS|VMFS|local)
  -username=             Username to use when connecting (CIFS only)
  -version=<nil>         VMFS major version
```

## datastore.disk.create

```
Usage: govc datastore.disk.create [OPTIONS] VMDK

Create VMDK on DS.

Examples:
  govc datastore.mkdir disks
  govc datastore.disk.create -size 24G disks/disk1.vmdk
  govc datastore.disk.create disks/parent.vmdk disk/child.vmdk

Options:
  -a=lsiLogic            Disk adapter
  -d=thin                Disk format
  -ds=                   Datastore [GOVC_DATASTORE]
  -f=false               Force
  -size=10.0GB           Size of new disk
  -uuid=                 Disk UUID
```

## datastore.disk.inflate

```
Usage: govc datastore.disk.inflate [OPTIONS] VMDK

Inflate VMDK on DS.

Examples:
  govc datastore.disk.inflate disks/disk1.vmdk

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
```

## datastore.disk.info

```
Usage: govc datastore.disk.info [OPTIONS] VMDK

Query VMDK info on DS.

Examples:
  govc datastore.disk.info disks/disk1.vmdk

Options:
  -c=false               Chain format
  -d=false               Include datastore in output
  -ds=                   Datastore [GOVC_DATASTORE]
  -p=true                Include parents
  -uuid=false            Include disk UUID
```

## datastore.disk.shrink

```
Usage: govc datastore.disk.shrink [OPTIONS] VMDK

Shrink VMDK on DS.

Examples:
  govc datastore.disk.shrink disks/disk1.vmdk

Options:
  -copy=<nil>            Perform shrink in-place mode if false, copy-shrink mode otherwise
  -ds=                   Datastore [GOVC_DATASTORE]
```

## datastore.download

```
Usage: govc datastore.download [OPTIONS] SOURCE DEST

Copy SOURCE from DS to DEST on the local system.

If DEST name is "-", source is written to stdout.

Examples:
  govc datastore.download vm-name/vmware.log ./local.log
  govc datastore.download vm-name/vmware.log - | grep -i error

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
  -host=                 Host system [GOVC_HOST]
```

## datastore.info

```
Usage: govc datastore.info [OPTIONS] [PATH]...

Display info for Datastores.

Examples:
  govc datastore.info
  govc datastore.info vsanDatastore
  # info on Datastores shared between cluster hosts:
  govc object.collect -s -d " " /dc1/host/k8s-cluster host | xargs govc datastore.info -H
  # info on Datastores shared between VM hosts:
  govc ls /dc1/vm/*k8s* | xargs -n1 -I% govc object.collect -s % summary.runtime.host | xargs govc datastore.info -H

Options:
  -H=false               Display info for Datastores shared between hosts
```

## datastore.ls

```
Usage: govc datastore.ls [OPTIONS] [FILE]...

Options:
  -R=false               List subdirectories recursively
  -a=false               Do not ignore entries starting with .
  -ds=                   Datastore [GOVC_DATASTORE]
  -l=false               Long listing format
  -p=false               Append / indicator to directories
```

## datastore.maintenance.enter

```
Usage: govc datastore.maintenance.enter [OPTIONS] DATASTORE

Put DATASTORE in maintenance mode.

Examples:
  govc datastore.cluster.change -drs-mode automated my-datastore-cluster # automatically schedule Storage DRS migration
  govc datastore.maintenance.enter -ds my-datastore-cluster/datastore1
  # no virtual machines can be powered on and no provisioning operations can be performed on the datastore during this time
  govc datastore.maintenance.exit -ds my-datastore-cluster/datastore1

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
```

## datastore.maintenance.exit

```
Usage: govc datastore.maintenance.exit [OPTIONS] DATASTORE

Take DATASTORE out of maintenance mode.

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
```

## datastore.mkdir

```
Usage: govc datastore.mkdir [OPTIONS] DIRECTORY

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
  -namespace=false       Return uuid of namespace created on vsan datastore
  -p=false               Create intermediate directories as needed
```

## datastore.mv

```
Usage: govc datastore.mv [OPTIONS] SRC DST

Move SRC to DST on DATASTORE.

Examples:
  govc datastore.mv foo/foo.vmx foo/foo.vmx.old
  govc datastore.mv -f my.vmx foo/foo.vmx

Options:
  -dc-target=            Datacenter destination (defaults to -dc)
  -ds=                   Datastore [GOVC_DATASTORE]
  -ds-target=            Datastore destination (defaults to -ds)
  -f=false               If true, overwrite any identically named file at the destination
  -t=true                Use file type to choose disk or file manager
```

## datastore.remove

```
Usage: govc datastore.remove [OPTIONS] HOST...

Remove datastore from HOST.

Examples:
  govc datastore.remove -ds nfsDatastore cluster1
  govc datastore.remove -ds nasDatastore host1 host2 host3

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
  -host=                 Host system [GOVC_HOST]
```

## datastore.rm

```
Usage: govc datastore.rm [OPTIONS] FILE

Remove FILE from DATASTORE.

Examples:
  govc datastore.rm vm/vmware.log
  govc datastore.rm vm
  govc datastore.rm -f images/base.vmdk

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
  -f=false               Force; ignore nonexistent files and arguments
  -namespace=false       Path is uuid of namespace on vsan datastore
  -t=true                Use file type to choose disk or file manager
```

## datastore.tail

```
Usage: govc datastore.tail [OPTIONS] PATH

Output the last part of datastore files.

Examples:
  govc datastore.tail -n 100 vm-name/vmware.log
  govc datastore.tail -n 0 -f vm-name/vmware.log

Options:
  -c=-1                  Output the last NUM bytes
  -ds=                   Datastore [GOVC_DATASTORE]
  -f=false               Output appended data as the file grows
  -host=                 Host system [GOVC_HOST]
  -n=10                  Output the last NUM lines
```

## datastore.upload

```
Usage: govc datastore.upload [OPTIONS] SOURCE DEST

Copy SOURCE from the local system to DEST on DS.

If SOURCE name is "-", read source from stdin.

Examples:
  govc datastore.upload -ds datastore1 ./config.iso vm-name/config.iso
  genisoimage ... | govc datastore.upload -ds datastore1 - vm-name/config.iso

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
```

## datastore.vsan.dom.ls

```
Usage: govc datastore.vsan.dom.ls [OPTIONS] [UUID]...

List vSAN DOM objects in DS.

Examples:
  govc datastore.vsan.dom.ls
  govc datastore.vsan.dom.ls -ds vsanDatastore -l
  govc datastore.vsan.dom.ls -l d85aa758-63f5-500a-3150-0200308e589c

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
  -l=false               Long listing
  -o=false               List orphan objects
```

## datastore.vsan.dom.rm

```
Usage: govc datastore.vsan.dom.rm [OPTIONS] UUID...

Remove vSAN DOM objects in DS.

Examples:
  govc datastore.vsan.dom.rm d85aa758-63f5-500a-3150-0200308e589c
  govc datastore.vsan.dom.rm -f d85aa758-63f5-500a-3150-0200308e589c
  govc datastore.vsan.dom.ls -o | xargs govc datastore.vsan.dom.rm

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
  -f=false               Force delete
  -v=false               Print deleted UUIDs to stdout, failed to stderr
```

## device.boot

```
Usage: govc device.boot [OPTIONS]

Configure VM boot settings.

Examples:
  govc device.boot -vm $vm -delay 1000 -order floppy,cdrom,ethernet,disk
  govc device.boot -vm $vm -order - # reset boot order
  govc device.boot -vm $vm -firmware efi -secure
  govc device.boot -vm $vm -firmware bios -secure=false

Options:
  -delay=0               Delay in ms before starting the boot sequence
  -firmware=             Firmware type [bios|efi]
  -order=                Boot device order [-,floppy,cdrom,ethernet,disk]
  -retry=false           If true, retry boot after retry-delay
  -retry-delay=0         Delay in ms before a boot retry
  -secure=<nil>          Enable EFI secure boot
  -setup=false           If true, enter BIOS setup on next boot
  -vm=                   Virtual machine [GOVC_VM]
```

## device.cdrom.add

```
Usage: govc device.cdrom.add [OPTIONS]

Add CD-ROM device to VM.

Examples:
  govc device.cdrom.add -vm $vm
  govc device.ls -vm $vm | grep ide-
  govc device.cdrom.add -vm $vm -controller ide-200
  govc device.info cdrom-*

Options:
  -controller=           IDE controller name
  -vm=                   Virtual machine [GOVC_VM]
```

## device.cdrom.eject

```
Usage: govc device.cdrom.eject [OPTIONS]

Eject media from CD-ROM device.

If device is not specified, the first CD-ROM device is used.

Examples:
  govc device.cdrom.eject -vm vm-1
  govc device.cdrom.eject -vm vm-1 -device floppy-1

Options:
  -device=               CD-ROM device name
  -vm=                   Virtual machine [GOVC_VM]
```

## device.cdrom.insert

```
Usage: govc device.cdrom.insert [OPTIONS] ISO

Insert media on datastore into CD-ROM device.

If device is not specified, the first CD-ROM device is used.

Examples:
  govc device.cdrom.insert -vm vm-1 -device cdrom-3000 images/boot.iso

Options:
  -device=               CD-ROM device name
  -ds=                   Datastore [GOVC_DATASTORE]
  -vm=                   Virtual machine [GOVC_VM]
```

## device.connect

```
Usage: govc device.connect [OPTIONS] DEVICE...

Connect DEVICE on VM.

Examples:
  govc device.connect -vm $name cdrom-3000

Options:
  -vm=                   Virtual machine [GOVC_VM]
```

## device.disconnect

```
Usage: govc device.disconnect [OPTIONS] DEVICE...

Disconnect DEVICE on VM.

Examples:
  govc device.disconnect -vm $name cdrom-3000

Options:
  -vm=                   Virtual machine [GOVC_VM]
```

## device.floppy.add

```
Usage: govc device.floppy.add [OPTIONS]

Add floppy device to VM.

Examples:
  govc device.floppy.add -vm $vm
  govc device.info floppy-*

Options:
  -vm=                   Virtual machine [GOVC_VM]
```

## device.floppy.eject

```
Usage: govc device.floppy.eject [OPTIONS]

Eject image from floppy device.

If device is not specified, the first floppy device is used.

Examples:
  govc device.floppy.eject -vm vm-1

Options:
  -device=               Floppy device name
  -vm=                   Virtual machine [GOVC_VM]
```

## device.floppy.insert

```
Usage: govc device.floppy.insert [OPTIONS] IMG

Insert IMG on datastore into floppy device.

If device is not specified, the first floppy device is used.

Examples:
  govc device.floppy.insert -vm vm-1 vm-1/config.img

Options:
  -device=               Floppy device name
  -ds=                   Datastore [GOVC_DATASTORE]
  -vm=                   Virtual machine [GOVC_VM]
```

## device.info

```
Usage: govc device.info [OPTIONS] [DEVICE]...

Device info for VM.

Examples:
  govc device.info -vm $name
  govc device.info -vm $name disk-*
  govc device.info -vm $name -json disk-* | jq -r .Devices[].Backing.Uuid
  govc device.info -vm $name -json 'disk-*' | jq -r .Devices[].Backing.FileName # vmdk path
  govc device.info -vm $name -json ethernet-0 | jq -r .Devices[].MacAddress

Options:
  -net=                  Network [GOVC_NETWORK]
  -net.adapter=e1000     Network adapter type
  -net.address=          Network hardware address
  -vm=                   Virtual machine [GOVC_VM]
```

## device.ls

```
Usage: govc device.ls [OPTIONS]

List devices for VM.

Examples:
  govc device.ls -vm $name
  govc device.ls -vm $name disk-*
  govc device.ls -vm $name -json | jq '.Devices[].Name'

Options:
  -boot=false            List devices configured in the VM's boot options
  -vm=                   Virtual machine [GOVC_VM]
```

## device.pci.add

```
Usage: govc device.pci.add [OPTIONS] PCI_ADDRESS...

Add PCI Passthrough device to VM.

Examples:
  govc device.pci.ls -vm $vm
  govc device.pci.add -vm $vm $pci_address
  govc device.info -vm $vm

Assuming vm name is helloworld, list command has below output

$ govc device.pci.ls -vm helloworld
System ID                             Address       Vendor Name Device Name
5b087ce4-ce46-72c0-c7c2-28ac9e22c3c2  0000:60:00.0  Pensando    Ethernet Controller 1
5b087ce4-ce46-72c0-c7c2-28ac9e22c3c2  0000:61:00.0  Pensando    Ethernet Controller 2

To add only 'Ethernet Controller 1', command should be as below. No output upon success.

$ govc device.pci.add -vm helloworld 0000:60:00.0

To add both 'Ethernet Controller 1' and 'Ethernet Controller 2', command should be as below.
No output upon success.

$ govc device.pci.add -vm helloworld 0000:60:00.0 0000:61:00.0

$ govc device.info -vm helloworld
...
Name:               pcipassthrough-13000
  Type:             VirtualPCIPassthrough
  Label:            PCI device 0
  Summary:
  Key:              13000
  Controller:       pci-100
  Unit number:      18
Name:               pcipassthrough-13001
  Type:             VirtualPCIPassthrough
  Label:            PCI device 1
  Summary:
  Key:              13001
  Controller:       pci-100
  Unit number:      19

Options:
  -vm=                   Virtual machine [GOVC_VM]
```

## device.pci.ls

```
Usage: govc device.pci.ls [OPTIONS]

List allowed PCI passthrough devices that could be attach to VM.

Examples:
  govc device.pci.ls -vm VM

Options:
  -vm=                   Virtual machine [GOVC_VM]
```

## device.pci.remove

```
Usage: govc device.pci.remove [OPTIONS] <PCI ADDRESS>...

Remove PCI Passthrough device from VM.

Examples:
  govc device.info -vm $vm
  govc device.pci.remove -vm $vm $pci_address
  govc device.info -vm $vm

Assuming vm name is helloworld, device info command has below output

$ govc device.info -vm helloworld
...
Name:               pcipassthrough-13000
  Type:             VirtualPCIPassthrough
  Label:            PCI device 0
  Summary:
  Key:              13000
  Controller:       pci-100
  Unit number:      18
Name:               pcipassthrough-13001
  Type:             VirtualPCIPassthrough
  Label:            PCI device 1
  Summary:
  Key:              13001
  Controller:       pci-100
  Unit number:      19

To remove only 'pcipassthrough-13000', command should be as below. No output upon success.

$ govc device.pci.remove -vm helloworld pcipassthrough-13000

To remove both 'pcipassthrough-13000' and 'pcipassthrough-13001', command should be as below.
No output upon success.

$ govc device.pci.remove -vm helloworld pcipassthrough-13000 pcipassthrough-13001

Options:
  -vm=                   Virtual machine [GOVC_VM]
```

## device.remove

```
Usage: govc device.remove [OPTIONS] DEVICE...

Remove DEVICE from VM.

Examples:
  govc device.remove -vm $name cdrom-3000
  govc device.remove -vm $name disk-1000
  govc device.remove -vm $name -keep disk-*

Options:
  -keep=false            Keep files in datastore
  -vm=                   Virtual machine [GOVC_VM]
```

## device.scsi.add

```
Usage: govc device.scsi.add [OPTIONS]

Add SCSI controller to VM.

Examples:
  govc device.scsi.add -vm $vm
  govc device.scsi.add -vm $vm -type pvscsi
  govc device.info -vm $vm {lsi,pv}*

Options:
  -hot=false             Enable hot-add/remove
  -sharing=noSharing     SCSI sharing
  -type=lsilogic         SCSI controller type (lsilogic|buslogic|pvscsi|lsilogic-sas)
  -vm=                   Virtual machine [GOVC_VM]
```

## device.serial.add

```
Usage: govc device.serial.add [OPTIONS]

Add serial port to VM.

Examples:
  govc device.serial.add -vm $vm
  govc device.info -vm $vm serialport-*

Options:
  -vm=                   Virtual machine [GOVC_VM]
```

## device.serial.connect

```
Usage: govc device.serial.connect [OPTIONS] URI

Connect service URI to serial port.

If "-" is given as URI, connects file backed device with file name of
device name + .log suffix in the VM Config.Files.LogDirectory.

Defaults to the first serial port if no DEVICE is given.

Examples:
  govc device.ls | grep serialport-
  govc device.serial.connect -vm $vm -device serialport-8000 telnet://:33233
  govc device.info -vm $vm serialport-*
  govc device.serial.connect -vm $vm "[datastore1] $vm/console.log"
  govc device.serial.connect -vm $vm -
  govc datastore.tail -f $vm/serialport-8000.log

Options:
  -client=false          Use client direction
  -device=               serial port device name
  -vm=                   Virtual machine [GOVC_VM]
  -vspc-proxy=           vSPC proxy URI
```

## device.serial.disconnect

```
Usage: govc device.serial.disconnect [OPTIONS]

Disconnect service URI from serial port.

Examples:
  govc device.ls | grep serialport-
  govc device.serial.disconnect -vm $vm -device serialport-8000
  govc device.info -vm $vm serialport-*

Options:
  -device=               serial port device name
  -vm=                   Virtual machine [GOVC_VM]
```

## device.usb.add

```
Usage: govc device.usb.add [OPTIONS]

Add USB device to VM.

Examples:
  govc device.usb.add -vm $vm
  govc device.usb.add -type xhci -vm $vm
  govc device.info usb*

Options:
  -auto=true             Enable ability to hot plug devices
  -ehci=true             Enable enhanced host controller interface (USB 2.0)
  -type=usb              USB controller type (usb|xhci)
  -vm=                   Virtual machine [GOVC_VM]
```

## disk.create

```
Usage: govc disk.create [OPTIONS] NAME

Create disk NAME on DS.

Examples:
  govc disk.create -size 24G my-disk

Options:
  -datastore-cluster=    Datastore cluster [GOVC_DATASTORE_CLUSTER]
  -ds=                   Datastore [GOVC_DATASTORE]
  -keep=<nil>            Keep disk after VM is deleted
  -pool=                 Resource pool [GOVC_RESOURCE_POOL]
  -size=10.0GB           Size of new disk
```

## disk.ls

```
Usage: govc disk.ls [OPTIONS] [ID]...

List disk IDs on DS.

Examples:
  govc disk.ls
  govc disk.ls -l -T
  govc disk.ls -l e9b06a8b-d047-4d3c-b15b-43ea9608b1a6
  govc disk.ls -c k8s-region -t us-west-2

Options:
  -L=false               Print disk backing path instead of disk name
  -R=false               Reconcile the datastore inventory info
  -T=false               List attached tags
  -c=                    Query tag category
  -ds=                   Datastore [GOVC_DATASTORE]
  -l=false               Long listing format
  -t=                    Query tag name
```

## disk.register

```
Usage: govc disk.register [OPTIONS] PATH [NAME]

Register existing disk on DS.

Examples:
  govc disk.register disks/disk1.vmdk my-disk

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
```

## disk.rm

```
Usage: govc disk.rm [OPTIONS] ID

Remove disk ID on DS.

Examples:
  govc disk.rm ID

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
```

## disk.snapshot.create

```
Usage: govc disk.snapshot.create [OPTIONS] ID DESC

Create snapshot of ID on DS.

Examples:
  govc disk.snapshot.create b9fe5f17-3b87-4a03-9739-09a82ddcc6b0 my-disk-snapshot

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
```

## disk.snapshot.ls

```
Usage: govc disk.snapshot.ls [OPTIONS] ID

List snapshots for disk ID on DS.

Examples:
  govc snapshot.disk.ls -l 9b06a8b-d047-4d3c-b15b-43ea9608b1a6

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
  -l=false               Long listing format
```

## disk.snapshot.rm

```
Usage: govc disk.snapshot.rm [OPTIONS] ID SID

Remove disk ID snapshot ID on DS.

Examples:
  govc disk.snapshot.rm ffe6a398-eb8e-4eaa-9118-e1f16b8b8e3c ecbca542-0a25-4127-a585-82e4047750d6

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
```

## disk.tags.attach

```
Usage: govc disk.tags.attach [OPTIONS] NAME ID

Attach tag NAME to disk ID.

Examples:
  govc disk.tags.attach -c k8s-region k8s-region-us $id

Options:
  -c=                    Tag category
```

## disk.tags.detach

```
Usage: govc disk.tags.detach [OPTIONS] NAME ID

Detach tag NAME from disk ID.

Examples:
  govc disk.tags.detach -c k8s-region k8s-region-us $id

Options:
  -c=                    Tag category
```

## dvs.add

```
Usage: govc dvs.add [OPTIONS] HOST...

Add hosts to DVS.

Examples:
  govc dvs.add -dvs dvsName -pnic vmnic1 hostA hostB hostC

Options:
  -dvs=                  DVS path
  -host=                 Host system [GOVC_HOST]
  -pnic=vmnic0           Name of the host physical NIC
```

## dvs.change

```
Usage: govc dvs.change [OPTIONS] DVS

Change DVS (DistributedVirtualSwitch) in datacenter.

Examples:
  govc dvs.change -product-version 5.5.0 DSwitch
  govc dvs.change -mtu 9000 DSwitch
  govc dvs.change -discovery-protocol [lldp|cdp] DSwitch

Options:
  -discovery-protocol=   Link Discovery Protocol
  -mtu=0                 DVS Max MTU
  -product-version=      DVS product version
```

## dvs.create

```
Usage: govc dvs.create [OPTIONS] DVS

Create DVS (DistributedVirtualSwitch) in datacenter.

The dvs is added to the folder specified by the 'folder' flag. If not given,
this defaults to the network folder in the specified or default datacenter.

Examples:
  govc dvs.create DSwitch
  govc dvs.create -product-version 5.5.0 DSwitch
  govc dvs.create -mtu 9000 DSwitch
  govc dvs.create -discovery-protocol [lldp|cdp] DSwitch

Options:
  -discovery-protocol=   Link Discovery Protocol
  -folder=               Inventory folder [GOVC_FOLDER]
  -mtu=0                 DVS Max MTU
  -num-uplinks=0         Number of Uplinks
  -product-version=      DVS product version
```

## dvs.portgroup.add

```
Usage: govc dvs.portgroup.add [OPTIONS] NAME

Add portgroup to DVS.

The '-type' options are defined by the dvs.DistributedVirtualPortgroup.PortgroupType API.
The UI labels '-type' as "Port binding" with the following choices:
    "Static binding":  earlyBinding
    "Dynanic binding": lateBinding
    "No binding":      ephemeral

The '-auto-expand' option is labeled in the UI as "Port allocation".
The default value is false, behaves as the UI labeled "Fixed" choice.
When given '-auto-expand=true', behaves as the UI labeled "Elastic" choice.

Examples:
  govc dvs.create DSwitch
  govc dvs.portgroup.add -dvs DSwitch -type earlyBinding -nports 16 ExternalNetwork
  govc dvs.portgroup.add -dvs DSwitch -type ephemeral InternalNetwork
  govc object.destroy network/InternalNetwork # remove the portgroup

Options:
  -auto-expand=<nil>     Ignore the limit on the number of ports
  -dvs=                  DVS path
  -nports=128            Number of ports
  -type=earlyBinding     Portgroup type (earlyBinding|lateBinding|ephemeral)
  -vlan=0                VLAN ID
  -vlan-mode=vlan        vlan mode (vlan|trunking)
  -vlan-range=0-4094     VLAN Ranges with comma delimited
```

## dvs.portgroup.change

```
Usage: govc dvs.portgroup.change [OPTIONS] PATH

Change DVS portgroup configuration.

Examples:
  govc dvs.portgroup.change -nports 26 ExternalNetwork
  govc dvs.portgroup.change -vlan 3214 ExternalNetwork

Options:
  -auto-expand=<nil>     Ignore the limit on the number of ports
  -nports=0              Number of ports
  -type=earlyBinding     Portgroup type (earlyBinding|lateBinding|ephemeral)
  -vlan=0                VLAN ID
  -vlan-mode=vlan        vlan mode (vlan|trunking)
  -vlan-range=0-4094     VLAN Ranges with comma delimited
```

## dvs.portgroup.info

```
Usage: govc dvs.portgroup.info [OPTIONS] DVS

Portgroup info for DVS.

Examples:
  govc dvs.portgroup.info DSwitch
  govc dvs.portgroup.info -pg InternalNetwork DSwitch
  govc find / -type DistributedVirtualSwitch | xargs -n1 govc dvs.portgroup.info

Options:
  -active=false          Filter by port active or inactive status
  -connected=false       Filter by port connected or disconnected status
  -count=0               Number of matches to return (0 = unlimited)
  -inside=true           Filter by port inside or outside status
  -pg=                   Distributed Virtual Portgroup
  -r=false               Show DVS rules
  -uplinkPort=false      Filter for uplink ports
  -vlan=0                Filter by VLAN ID (0 = unfiltered)
```

## env

```
Usage: govc env [OPTIONS]

Output the environment variables for this client.

If credentials are included in the url, they are split into separate variables.
Useful as bash scripting helper to parse GOVC_URL.

Options:
  -x=false               Output variables for each GOVC_URL component
```

## events

```
Usage: govc events [OPTIONS] [PATH]...

Display events.

Examples:
  govc events vm/my-vm1 vm/my-vm2
  govc events /dc1/vm/* /dc2/vm/*
  govc events -type VmPoweredOffEvent -type VmPoweredOnEvent
  govc ls -t HostSystem host/* | xargs govc events | grep -i vsan

Options:
  -f=false               Follow event stream
  -force=false           Disable number objects to monitor limit
  -l=false               Long listing format
  -n=25                  Output the last N events
  -type=[]               Include only the specified event types
```

## export.ovf

```
Usage: govc export.ovf [OPTIONS] DIR

Export VM.

Examples:
  govc export.ovf -vm $vm DIR

Options:
  -f=false               Overwrite existing
  -i=false               Include image files (*.{iso,img})
  -name=                 Specifies target name (defaults to source name)
  -prefix=true           Prepend target name to image filenames if missing
  -sha=0                 Generate manifest using SHA 1, 256, 512 or 0 to skip
  -snapshot=             Specifies a snapshot to export from (supports running VMs)
  -vm=                   Virtual machine [GOVC_VM]
```

## extension.info

```
Usage: govc extension.info [OPTIONS] [KEY]...

Options:
```

## extension.register

```
Usage: govc extension.register [OPTIONS]

Options:
  -update=false          Update extension
```

## extension.setcert

```
Usage: govc extension.setcert [OPTIONS] ID

Set certificate for the extension ID.

The '-cert-pem' option can be one of the following:
'-' : Read the certificate from stdin
'+' : Generate a new key pair and save locally to ID.crt and ID.key
... : Any other value is passed as-is to ExtensionManager.SetCertificate

Examples:
  govc extension.setcert -cert-pem + -org Example com.example.extname

Options:
  -cert-pem=-            PEM encoded certificate
  -org=VMware            Organization for generated certificate
```

## extension.unregister

```
Usage: govc extension.unregister [OPTIONS]

Options:
```

## fields.add

```
Usage: govc fields.add [OPTIONS] NAME

Add a custom field type with NAME.

Examples:
  govc fields.add my-field-name # adds a field to all managed object types
  govc fields.add -type VirtualMachine my-vm-field-name # adds a field to the VirtualMachine type

Options:
  -type=                 Managed object type
```

## fields.info

```
Usage: govc fields.info [OPTIONS] PATH...

Display custom field values for PATH.

Also known as "Custom Attributes".

Examples:
  govc fields.info vm/*
  govc fields.info -n my-field-name vm/*

Options:
  -n=                    Filter by custom field name
```

## fields.ls

```
Usage: govc fields.ls [OPTIONS]

List custom field definitions.

Examples:
  govc fields.ls
  govc fields.ls -type VirtualMachine

Options:
  -type=                 Filter by a Managed Object Type
```

## fields.rename

```
Usage: govc fields.rename [OPTIONS] KEY NAME

Options:
```

## fields.rm

```
Usage: govc fields.rm [OPTIONS] KEY...

Options:
```

## fields.set

```
Usage: govc fields.set [OPTIONS] KEY VALUE PATH...

Set custom field values for PATH.

Examples:
  govc fields.set my-field-name field-value vm/my-vm
  govc fields.set -add my-new-global-field-name field-value vm/my-vm
  govc fields.set -add -type VirtualMachine my-new-vm-field-name field-value vm/my-vm

Options:
  -add=false             Adds the field if it does not exist. Use the -type flag to specify the managed object type to which the field is added. Using -add and omitting -kind causes a new, global field to be created if a field with the provided name does not already exist.
  -type=                 Managed object type on which to add the field if it does not exist. This flag is ignored unless -add=true
```

## find

```
Usage: govc find [OPTIONS] [ROOT] [KEY VAL]...

Find managed objects.

ROOT can be an inventory path or ManagedObjectReference.
ROOT defaults to '.', an alias for the root folder or DC if set.

Optional KEY VAL pairs can be used to filter results against object instance properties.
Use the govc 'object.collect' command to view possible object property keys.

The '-type' flag value can be a managed entity type or one of the following aliases:

  a    VirtualApp
  c    ClusterComputeResource
  d    Datacenter
  f    Folder
  g    DistributedVirtualPortgroup
  h    HostSystem
  m    VirtualMachine
  n    Network
  o    OpaqueNetwork
  p    ResourcePool
  r    ComputeResource
  s    Datastore
  w    DistributedVirtualSwitch

Examples:
  govc find
  govc find -l / # include object type in output
  govc find /dc1 -type c
  govc find vm -name my-vm-*
  govc find . -type n
  govc find -p /folder-a/dc-1/host/folder-b/cluster-a -type Datacenter # prints /folder-a/dc-1
  govc find . -type m -runtime.powerState poweredOn
  govc find . -type m -datastore $(govc find -i datastore -name vsanDatastore)
  govc find . -type s -summary.type vsan
  govc find . -type s -customValue *:prod # Key:Value
  govc find . -type h -hardware.cpuInfo.numCpuCores 16

Options:
  -i=false               Print the managed object reference
  -l=false               Long listing format
  -maxdepth=-1           Max depth
  -name=*                Resource name
  -p=false               Find parent objects
  -type=[]               Resource type
```

## firewall.ruleset.find

```
Usage: govc firewall.ruleset.find [OPTIONS]

Find firewall rulesets matching the given rule.

For a complete list of rulesets: govc host.esxcli network firewall ruleset list
For a complete list of rules:    govc host.esxcli network firewall ruleset rule list

Examples:
  govc firewall.ruleset.find -direction inbound -port 22
  govc firewall.ruleset.find -direction outbound -port 2377

Options:
  -c=true                Check if esx firewall is enabled
  -direction=outbound    Direction
  -enabled=true          Find enabled rule sets if true, disabled if false
  -host=                 Host system [GOVC_HOST]
  -port=0                Port
  -proto=tcp             Protocol
  -type=dst              Port type
```

## folder.create

```
Usage: govc folder.create [OPTIONS] PATH...

Create folder with PATH.

Examples:
  govc folder.create /dc1/vm/folder-foo
  govc object.mv /dc1/vm/vm-foo-* /dc1/vm/folder-foo
  govc folder.create -pod /dc1/datastore/sdrs
  govc object.mv /dc1/datastore/iscsi-* /dc1/datastore/sdrs

Options:
  -pod=false             Create folder(s) of type StoragePod (DatastoreCluster)
```

## folder.info

```
Usage: govc folder.info [OPTIONS] [PATH]...

Options:
```

## guest.chmod

```
Usage: govc guest.chmod [OPTIONS] MODE FILE

Change FILE MODE on VM.

Examples:
  govc guest.chmod -vm $name 0644 /var/log/foo.log

Options:
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.chown

```
Usage: govc guest.chown [OPTIONS] UID[:GID] FILE

Change FILE UID and GID on VM.

Examples:
  govc guest.chown -vm $name UID[:GID] /var/log/foo.log

Options:
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.df

```
Usage: govc guest.df [OPTIONS]

Report file system disk space usage.

Examples:
  govc guest.df -vm $name

Options:
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.download

```
Usage: govc guest.download [OPTIONS] SOURCE DEST

Copy SOURCE from the guest VM to DEST on the local system.

If DEST name is "-", source is written to stdout.

Examples:
  govc guest.download -l user:pass -vm=my-vm /var/log/my.log ./local.log
  govc guest.download -l user:pass -vm=my-vm /etc/motd -
  tar -cf- foo/ | govc guest.run -d - tar -C /tmp -xf-
  govc guest.run tar -C /tmp -cf- foo/ | tar -C /tmp -xf- # download directory

Options:
  -f=false               If set, the local destination file is clobbered
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.getenv

```
Usage: govc guest.getenv [OPTIONS] [NAME]...

Read NAME environment variables from VM.

Examples:
  govc guest.getenv -vm $name
  govc guest.getenv -vm $name HOME

Options:
  -i=false               Interactive session
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.kill

```
Usage: govc guest.kill [OPTIONS]

Kill process ID on VM.

Examples:
  govc guest.kill -vm $name -p 12345

Options:
  -i=false               Interactive session
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -p=[]                  Process ID
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.ls

```
Usage: govc guest.ls [OPTIONS] PATH

List PATH files in VM.

Examples:
  govc guest.ls -vm $name /tmp

Options:
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -s=false               Simple path only listing
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.mkdir

```
Usage: govc guest.mkdir [OPTIONS] PATH

Create directory PATH in VM.

Examples:
  govc guest.mkdir -vm $name /tmp/logs
  govc guest.mkdir -vm $name -p /tmp/logs/foo/bar

Options:
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -p=false               Create intermediate directories as needed
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.mktemp

```
Usage: govc guest.mktemp [OPTIONS]

Create a temporary file or directory in VM.

Examples:
  govc guest.mktemp -vm $name
  govc guest.mktemp -vm $name -d
  govc guest.mktemp -vm $name -t myprefix
  govc guest.mktemp -vm $name -p /var/tmp/$USER

Options:
  -d=false               Make a directory instead of a file
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -p=                    If specified, create relative to this directory
  -s=                    Suffix
  -t=                    Prefix
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.mv

```
Usage: govc guest.mv [OPTIONS] SOURCE DEST

Move (rename) files in VM.

Examples:
  govc guest.mv -vm $name /tmp/foo.sh /tmp/bar.sh
  govc guest.mv -vm $name -n /tmp/baz.sh /tmp/bar.sh

Options:
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -n=false               Do not overwrite an existing file
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.ps

```
Usage: govc guest.ps [OPTIONS]

List processes in VM.

By default, unless the '-e', '-p' or '-U' flag is specified, only processes owned
by the '-l' flag user are displayed.

The '-x' and '-X' flags only apply to processes started by vmware-tools,
such as those started with the govc guest.start command.

Examples:
  govc guest.ps -vm $name
  govc guest.ps -vm $name -e
  govc guest.ps -vm $name -p 12345
  govc guest.ps -vm $name -U root

Options:
  -U=                    Select by process UID
  -X=false               Wait for process to exit
  -e=false               Select all processes
  -i=false               Interactive session
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -p=[]                  Select by process ID
  -vm=                   Virtual machine [GOVC_VM]
  -x=false               Output exit time and code
```

## guest.rm

```
Usage: govc guest.rm [OPTIONS] PATH

Remove file PATH in VM.

Examples:
  govc guest.rm -vm $name /tmp/foo.log

Options:
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.rmdir

```
Usage: govc guest.rmdir [OPTIONS] PATH

Remove directory PATH in VM.

Examples:
  govc guest.rmdir -vm $name /tmp/empty-dir
  govc guest.rmdir -vm $name -r /tmp/non-empty-dir

Options:
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -r=false               Recursive removal
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.run

```
Usage: govc guest.run [OPTIONS] PATH [ARG]...

Run program PATH in VM and display output.

The guest.run command starts a program in the VM with i/o redirected, waits for the process to exit and
propagates the exit code to the govc process exit code.  Note that stdout and stderr are redirected by default,
stdin is only redirected when the '-d' flag is specified.

Note that vmware-tools requires program PATH to be absolute.
If PATH is not absolute and vm guest family is Windows,
guest.run changes the command to: 'c:\\Windows\\System32\\cmd.exe /c "PATH [ARG]..."'
Otherwise the command is changed to: '/bin/bash -c "PATH [ARG]..."'

Examples:
  govc guest.run -vm $name ifconfig
  govc guest.run -vm $name ifconfig eth0
  cal | govc guest.run -vm $name -d - cat
  govc guest.run -vm $name -d "hello $USER" cat
  govc guest.run -vm $name curl -s :invalid: || echo $? # exit code 6
  govc guest.run -vm $name -e FOO=bar -e BIZ=baz -C /tmp env
  govc guest.run -vm $name -l root:mypassword ntpdate -u pool.ntp.org
  govc guest.run -vm $name powershell C:\\network_refresh.ps1

Options:
  -C=                    The absolute path of the working directory for the program to start
  -d=                    Input data string. A value of '-' reads from OS stdin
  -e=[]                  Set environment variables
  -i=false               Interactive session
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.start

```
Usage: govc guest.start [OPTIONS] PATH [ARG]...

Start program in VM.

The process can have its status queried with govc guest.ps.
When the process completes, its exit code and end time will be available for 5 minutes after completion.

Examples:
  govc guest.start -vm $name /bin/mount /dev/hdb1 /data
  pid=$(govc guest.start -vm $name /bin/long-running-thing)
  govc guest.ps -vm $name -p $pid -X

Options:
  -C=                    The absolute path of the working directory for the program to start
  -e=[]                  Set environment variable (key=val)
  -i=false               Interactive session
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.touch

```
Usage: govc guest.touch [OPTIONS] FILE

Change FILE times on VM.

Examples:
  govc guest.touch -vm $name /var/log/foo.log
  govc guest.touch -vm $name -d "$(date -d '1 day ago')" /var/log/foo.log

Options:
  -a=false               Change only the access time
  -c=false               Do not create any files
  -d=                    Use DATE instead of current time
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -vm=                   Virtual machine [GOVC_VM]
```

## guest.upload

```
Usage: govc guest.upload [OPTIONS] SOURCE DEST

Copy SOURCE from the local system to DEST in the guest VM.

If SOURCE name is "-", read source from stdin.

Examples:
  govc guest.upload -l user:pass -vm=my-vm ~/.ssh/id_rsa.pub /home/$USER/.ssh/authorized_keys
  cowsay "have a great day" | govc guest.upload -l user:pass -vm=my-vm - /etc/motd
  tar -cf- foo/ | govc guest.run -d - tar -C /tmp -xf- # upload a directory

Options:
  -f=false               If set, the guest destination file is clobbered
  -gid=<nil>             Group ID
  -l=:                   Guest VM credentials [GOVC_GUEST_LOGIN]
  -perm=0                File permissions
  -uid=<nil>             User ID
  -vm=                   Virtual machine [GOVC_VM]
```

## host.account.create

```
Usage: govc host.account.create [OPTIONS]

Create local account on HOST.

Examples:
  govc host.account.create -id $USER -password password-for-esx60

Options:
  -description=          The description of the specified account
  -host=                 Host system [GOVC_HOST]
  -id=                   The ID of the specified account
  -password=             The password for the specified account id
```

## host.account.remove

```
Usage: govc host.account.remove [OPTIONS]

Remove local account on HOST.

Examples:
  govc host.account.remove -id $USER

Options:
  -description=          The description of the specified account
  -host=                 Host system [GOVC_HOST]
  -id=                   The ID of the specified account
  -password=             The password for the specified account id
```

## host.account.update

```
Usage: govc host.account.update [OPTIONS]

Update local account on HOST.

Examples:
  govc host.account.update -id root -password password-for-esx60

Options:
  -description=          The description of the specified account
  -host=                 Host system [GOVC_HOST]
  -id=                   The ID of the specified account
  -password=             The password for the specified account id
```

## host.add

```
Usage: govc host.add [OPTIONS]

Add host to datacenter.

The host is added to the folder specified by the 'folder' flag. If not given,
this defaults to the host folder in the specified or default datacenter.

Examples:
  thumbprint=$(govc about.cert -k -u host.example.com -thumbprint | awk '{print $2}')
  govc host.add -hostname host.example.com -username root -password pass -thumbprint $thumbprint
  govc host.add -hostname 10.0.6.1 -username root -password pass -noverify

Options:
  -connect=true          Immediately connect to host
  -folder=               Inventory folder [GOVC_FOLDER]
  -force=false           Force when host is managed by another VC
  -hostname=             Hostname or IP address of the host
  -noverify=false        Accept host thumbprint without verification
  -password=             Password of administration account on the host
  -thumbprint=           SHA-1 thumbprint of the host's SSL certificate
  -username=             Username of administration account on the host
```

## host.autostart.add

```
Usage: govc host.autostart.add [OPTIONS] VM...

Options:
  -host=                      Host system [GOVC_HOST]
  -start-action=powerOn       Start Action
  -start-delay=-1             Start Delay
  -start-order=-1             Start Order
  -stop-action=systemDefault  Stop Action
  -stop-delay=-1              Stop Delay
  -wait=systemDefault         Wait for Hearbeat Setting (systemDefault|yes|no)
```

## host.autostart.configure

```
Usage: govc host.autostart.configure [OPTIONS]

Options:
  -enabled=<nil>             Enable autostart
  -host=                     Host system [GOVC_HOST]
  -start-delay=0             Start delay
  -stop-action=              Stop action
  -stop-delay=0              Stop delay
  -wait-for-heartbeat=<nil>  Wait for hearbeat
```

## host.autostart.info

```
Usage: govc host.autostart.info [OPTIONS]

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.autostart.remove

```
Usage: govc host.autostart.remove [OPTIONS] VM...

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.cert.csr

```
Usage: govc host.cert.csr [OPTIONS]

Generate a certificate-signing request (CSR) for HOST.

Options:
  -host=                 Host system [GOVC_HOST]
  -ip=false              Use IP address as CN
```

## host.cert.import

```
Usage: govc host.cert.import [OPTIONS] FILE

Install SSL certificate FILE on HOST.

If FILE name is "-", read certificate from stdin.

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.cert.info

```
Usage: govc host.cert.info [OPTIONS]

Display SSL certificate info for HOST.

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.date.change

```
Usage: govc host.date.change [OPTIONS]

Change date and time for HOST.

Examples:
  govc host.date.change -date "$(date -u)"
  govc host.date.change -server time.vmware.com
  govc host.service enable ntpd
  govc host.service start ntpd

Options:
  -date=                 Update the date/time on the host
  -host=                 Host system [GOVC_HOST]
  -server=               IP or FQDN for NTP server(s)
  -tz=                   Change timezone of the host
```

## host.date.info

```
Usage: govc host.date.info [OPTIONS]

Display date and time info for HOST.

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.disconnect

```
Usage: govc host.disconnect [OPTIONS]

Disconnect HOST from vCenter.

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.esxcli

```
Usage: govc host.esxcli [OPTIONS] COMMAND [ARG]...

Invoke esxcli command on HOST.

Output is rendered in table form when possible, unless disabled with '-hints=false'.

Examples:
  govc host.esxcli network ip connection list
  govc host.esxcli system settings advanced set -o /Net/GuestIPHack -i 1
  govc host.esxcli network firewall ruleset set -r remoteSerialPort -e true
  govc host.esxcli network firewall set -e false
  govc host.esxcli hardware platform get

Options:
  -hints=true            Use command info hints when formatting output
  -host=                 Host system [GOVC_HOST]
```

## host.info

```
Usage: govc host.info [OPTIONS]

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.maintenance.enter

```
Usage: govc host.maintenance.enter [OPTIONS] HOST...

Put HOST in maintenance mode.

While this task is running and when the host is in maintenance mode,
no VMs can be powered on and no provisioning operations can be performed on the host.

Options:
  -evacuate=false        Evacuate powered off VMs
  -host=                 Host system [GOVC_HOST]
  -timeout=0             Timeout
```

## host.maintenance.exit

```
Usage: govc host.maintenance.exit [OPTIONS] HOST...

Take HOST out of maintenance mode.

This blocks if any concurrent running maintenance-only host configurations operations are being performed.
For example, if VMFS volumes are being upgraded.

The 'timeout' flag is the number of seconds to wait for the exit maintenance mode to succeed.
If the timeout is less than or equal to zero, there is no timeout.

Options:
  -host=                 Host system [GOVC_HOST]
  -timeout=0             Timeout
```

## host.option.ls

```
Usage: govc host.option.ls [OPTIONS] [NAME]

List option with the given NAME.

If NAME ends with a dot, all options for that subtree are listed.

Examples:
  govc host.option.ls
  govc host.option.ls Config.HostAgent.
  govc host.option.ls Config.HostAgent.plugins.solo.enableMob

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.option.set

```
Usage: govc host.option.set [OPTIONS] NAME VALUE

Set option NAME to VALUE.

Examples:
  govc host.option.set Config.HostAgent.plugins.solo.enableMob true
  govc host.option.set Config.HostAgent.log.level verbose

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.portgroup.add

```
Usage: govc host.portgroup.add [OPTIONS] NAME

Add portgroup to HOST.

Examples:
  govc host.portgroup.add -vswitch vSwitch0 -vlan 3201 bridge

Options:
  -host=                 Host system [GOVC_HOST]
  -vlan=0                VLAN ID
  -vswitch=              vSwitch Name
```

## host.portgroup.change

```
Usage: govc host.portgroup.change [OPTIONS] NAME

Change configuration of HOST portgroup NAME.

Examples:
  govc host.portgroup.change -allow-promiscuous -forged-transmits -mac-changes "VM Network"
  govc host.portgroup.change -vswitch-name vSwitch1 "Management Network"

Options:
  -allow-promiscuous=<nil>  Allow promiscuous mode
  -forged-transmits=<nil>   Allow forged transmits
  -host=                    Host system [GOVC_HOST]
  -mac-changes=<nil>        Allow MAC changes
  -name=                    Portgroup name
  -vlan-id=-1               VLAN ID
  -vswitch-name=            vSwitch name
```

## host.portgroup.info

```
Usage: govc host.portgroup.info [OPTIONS]

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.portgroup.remove

```
Usage: govc host.portgroup.remove [OPTIONS] NAME

Remove portgroup from HOST.

Examples:
  govc host.portgroup.remove bridge

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.reconnect

```
Usage: govc host.reconnect [OPTIONS]

Reconnect HOST to vCenter.

This command can also be used to change connection properties (hostname, fingerprint, username, password),
without disconnecting the host.

Options:
  -force=false           Force when host is managed by another VC
  -host=                 Host system [GOVC_HOST]
  -hostname=             Hostname or IP address of the host
  -noverify=false        Accept host thumbprint without verification
  -password=             Password of administration account on the host
  -sync-state=false      Sync state
  -thumbprint=           SHA-1 thumbprint of the host's SSL certificate
  -username=             Username of administration account on the host
```

## host.remove

```
Usage: govc host.remove [OPTIONS] HOST...

Remove HOST from vCenter.

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.service

```
Usage: govc host.service [OPTIONS] ACTION ID

Apply host service ACTION to service ID.

Where ACTION is one of: start, stop, restart, status, enable, disable

Examples:
  govc host.service enable TSM-SSH
  govc host.service start TSM-SSH

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.service.ls

```
Usage: govc host.service.ls [OPTIONS]

List HOST services.

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.shutdown

```
Usage: govc host.shutdown [OPTIONS] HOST...

Shutdown HOST.

Options:
  -f=false               Force shutdown when host is not in maintenance mode
  -host=                 Host system [GOVC_HOST]
  -r=false               Reboot host
```

## host.storage.info

```
Usage: govc host.storage.info [OPTIONS]

Show HOST storage system information.

Examples:
  govc find / -type h | xargs -n1 govc host.storage.info -unclaimed -host

Options:
  -host=                 Host system [GOVC_HOST]
  -refresh=false         Refresh the storage system provider
  -rescan=false          Rescan all host bus adapters
  -rescan-vmfs=false     Rescan for new VMFSs
  -t=lun                 Type (hba,lun)
  -unclaimed=false       Only show disks that can be used as new VMFS datastores
```

## host.storage.mark

```
Usage: govc host.storage.mark [OPTIONS] DEVICE_PATH

Mark device at DEVICE_PATH.

Options:
  -host=                 Host system [GOVC_HOST]
  -local=<nil>           Mark as local
  -ssd=<nil>             Mark as SSD
```

## host.storage.partition

```
Usage: govc host.storage.partition [OPTIONS] DEVICE_PATH

Show partition table for device at DEVICE_PATH.

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.vnic.change

```
Usage: govc host.vnic.change [OPTIONS] DEVICE

Change a virtual nic DEVICE.

Examples:
  govc host.vnic.change -host hostname -mtu 9000 vmk1

Options:
  -host=                 Host system [GOVC_HOST]
  -mtu=0                 vmk MTU
```

## host.vnic.info

```
Usage: govc host.vnic.info [OPTIONS]

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.vnic.service

```
Usage: govc host.vnic.service [OPTIONS] SERVICE DEVICE


Enable or disable service on a virtual nic device.

Where SERVICE is one of: vmotion|faultToleranceLogging|vSphereReplication|vSphereReplicationNFC|management|vsan|vSphereProvisioning
Where DEVICE is one of: vmk0|vmk1|...

Examples:
  govc host.vnic.service -host hostname -enable vsan vmk0
  govc host.vnic.service -host hostname -enable=false vmotion vmk1

Options:
  -enable=true           Enable service
  -host=                 Host system [GOVC_HOST]
```

## host.vswitch.add

```
Usage: govc host.vswitch.add [OPTIONS] NAME

Options:
  -host=                 Host system [GOVC_HOST]
  -mtu=0                 MTU
  -nic=                  Bridge nic device
  -ports=128             Number of ports
```

## host.vswitch.info

```
Usage: govc host.vswitch.info [OPTIONS]

Options:
  -host=                 Host system [GOVC_HOST]
```

## host.vswitch.remove

```
Usage: govc host.vswitch.remove [OPTIONS] NAME

Options:
  -host=                 Host system [GOVC_HOST]
```

## import.ova

```
Usage: govc import.ova [OPTIONS] PATH_TO_OVA

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
  -folder=               Inventory folder [GOVC_FOLDER]
  -host=                 Host system [GOVC_HOST]
  -name=                 Name to use for new entity
  -options=              Options spec file path for VM deployment
  -pool=                 Resource pool [GOVC_RESOURCE_POOL]
```

## import.ovf

```
Usage: govc import.ovf [OPTIONS] PATH_TO_OVF

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
  -folder=               Inventory folder [GOVC_FOLDER]
  -host=                 Host system [GOVC_HOST]
  -name=                 Name to use for new entity
  -options=              Options spec file path for VM deployment
  -pool=                 Resource pool [GOVC_RESOURCE_POOL]
```

## import.spec

```
Usage: govc import.spec [OPTIONS] PATH_TO_OVF_OR_OVA

Options:
```

## import.vmdk

```
Usage: govc import.vmdk [OPTIONS] PATH_TO_VMDK [REMOTE_DIRECTORY]

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
  -folder=               Inventory folder [GOVC_FOLDER]
  -force=false           Overwrite existing disk
  -pool=                 Resource pool [GOVC_RESOURCE_POOL]
```

## library.checkin

```
Usage: govc library.checkin [OPTIONS] PATH

Check in VM to Content Library item PATH.

Note: this command requires vCenter 7.0 or higher.

Examples:
  govc library.checkin -vm my-vm my-content/template-vm-item

Options:
  -m=                    Check in message
  -vm=                   Virtual machine [GOVC_VM]
```

## library.checkout

```
Usage: govc library.checkout [OPTIONS] PATH NAME

Check out Content Library item PATH to vm NAME.

Note: this command requires vCenter 7.0 or higher.

Examples:
  govc library.checkout -cluster my-cluster my-content/template-vm-item my-vm

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
  -folder=               Inventory folder [GOVC_FOLDER]
  -host=                 Host system [GOVC_HOST]
  -pool=                 Resource pool [GOVC_RESOURCE_POOL]
```

## library.clone

```
Usage: govc library.clone [OPTIONS] PATH NAME

Clone VM to Content Library PATH.

By default, clone as a VM template (requires vCenter version 6.7U1 or higher).
Clone as an OVF when the '-ovf' flag is specified.

Examples:
  govc library.clone -vm template-vm my-content template-vm-item
  govc library.clone -ovf -vm template-vm my-content ovf-item

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
  -ds=                   Datastore [GOVC_DATASTORE]
  -e=false               Include extra configuration
  -folder=               Inventory folder [GOVC_FOLDER]
  -host=                 Host system [GOVC_HOST]
  -m=false               Preserve MAC-addresses on network adapters
  -ovf=false             Clone as OVF (default is VM Template)
  -pool=                 Resource pool [GOVC_RESOURCE_POOL]
  -profile=              Storage profile
  -vm=                   Virtual machine [GOVC_VM]
```

## library.cp

```
Usage: govc library.cp [OPTIONS] SRC DST

Copy SRC library item to DST library.
Examples:
  govc library.cp /my-content/my-item /my-other-content
  govc library.cp -n my-item2 /my-content/my-item /my-other-content

Options:
  -n=                    Library item name
```

## library.create

```
Usage: govc library.create [OPTIONS] NAME

Create library.

Examples:
  govc library.create library_name
  govc library.create -sub http://server/path/lib.json library_name
  govc library.create -pub library_name

Options:
  -d=                    Description of library
  -ds=                   Datastore [GOVC_DATASTORE]
  -pub=<nil>             Publish library
  -pub-password=         Publication password
  -pub-username=         Publication username
  -sub=                  Subscribe to library URL
  -sub-autosync=true     Automatic synchronization
  -sub-ondemand=false    Download content on demand
  -sub-password=         Subscription password
  -sub-username=         Subscription username
  -thumbprint=           SHA-1 thumbprint of the host's SSL certificate
```

## library.deploy

```
Usage: govc library.deploy [OPTIONS] TEMPLATE [NAME]

Deploy library OVF template.

Examples:
  govc library.deploy /library_name/ovf_template vm_name
  govc library.export /library_name/ovf_template/*.ovf # save local copy of .ovf
  govc import.spec *.ovf > deploy.json # generate options from .ovf
  # edit deploy.json as needed
  govc library.deploy -options deploy.json /library_name/ovf_template

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
  -folder=               Inventory folder [GOVC_FOLDER]
  -host=                 Host system [GOVC_HOST]
  -options=              Options spec file path for VM deployment
  -pool=                 Resource pool [GOVC_RESOURCE_POOL]
  -profile=              Storage profile
```

## library.export

```
Usage: govc library.export [OPTIONS] PATH [DEST]

Export library items.

If the given PATH is a library item, all files will be downloaded.

If the given PATH is a library item file, only that file will be downloaded.

By default, files are saved using the library item's file names to the current directory.
If DEST is given for a library item, files are saved there instead of the current directory.
If DEST is given for a library item file, the file will be saved with that name.
If DEST is '-', the file contents are written to stdout instead of saving to a file.

Examples:
  govc library.export library_name/item_name
  govc library.export library_name/item_name/file_name
  govc library.export library_name/item_name/*.ovf -

Options:
```

## library.import

```
Usage: govc library.import [OPTIONS] LIBRARY ITEM

Import library items.

Examples:
  govc library.import library_name file.ova
  govc library.import library_name file.ovf
  govc library.import library_name file.iso
  govc library.import library_id file.iso # Use library id if multiple libraries have the same name
  govc library.import library_name/item_name file.ova # update existing item
  govc library.import library_name http://example.com/file.ovf # download and push to vCenter
  govc library.import -pull library_name http://example.com/file.ova # direct pull from vCenter

Options:
  -m=false               Require ova manifest
  -n=                    Library item name
  -pull=false            Pull library item from http endpoint
  -t=                    Library item type
```

## library.info

```
Usage: govc library.info [OPTIONS]

Display library information.

Examples:
  govc library.info
  govc library.info /lib1
  govc library.info -l /lib1 | grep Size:
  govc library.info /lib1/item1
  govc library.info /lib1/item1/
  govc library.info */
  govc device.cdrom.insert -vm $vm -device cdrom-3000 $(govc library.info -L /lib1/item1/file1)
  govc library.info -json | jq .
  govc library.info /lib1/item1 -json | jq .

Options:
  -L=false               List Datastore path only
  -U=false               List pub/sub URL(s) only
  -l=false               Long listing format
```

## library.ls

```
Usage: govc library.ls [OPTIONS]

List libraries, items, and files.

Examples:
  govc library.ls
  govc library.ls /lib1
  govc library.ls /lib1/item1
  govc library.ls /lib1/item1/
  govc library.ls */
  govc library.ls -json | jq .
  govc library.ls /lib1/item1 -json | jq .

Options:
```

## library.publish

```
Usage: govc library.publish [OPTIONS] NAME|ITEM [SUBSCRIPTION-ID]...

Publish library NAME or ITEM to subscribers.

If no subscriptions are specified, then publishes the library to all its subscribers.
See 'govc library.subscriber.ls' to get a list of subscription IDs.

Examples:
  govc library.publish /my-library
  govc library.publish /my-library subscription-id1 subscription-id2
  govc library.publish /my-library/my-item
  govc library.publish /my-library/my-item subscription-id1 subscription-id2

Options:
```

## library.rm

```
Usage: govc library.rm [OPTIONS] NAME

Delete library or item NAME.

Examples:
  govc library.rm /library_name
  govc library.rm library_id # Use library id if multiple libraries have the same name
  govc library.rm /library_name/item_name

Options:
```

## library.session.ls

```
Usage: govc library.session.ls [OPTIONS]

List library item update sessions.

Examples:
  govc library.session.ls
  govc library.session.ls -json | jq .

Options:
```

## library.session.rm

```
Usage: govc library.session.rm [OPTIONS]

Remove a library item update session.

Examples:
  govc library.session.rm session_id

Options:
  -f=false               Cancel session if active
```

## library.subscriber.create

```
Usage: govc library.subscriber.create [OPTIONS] PUBLISHED-LIBRARY SUBSCRIPTION-LIBRARY

Create library subscriber.

Examples:
  govc library.subscriber.create -cluster my-cluster published-library subscription-library

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
  -folder=               Inventory folder [GOVC_FOLDER]
  -host=                 Host system [GOVC_HOST]
  -net=                  Network [GOVC_NETWORK]
  -net.adapter=e1000     Network adapter type
  -net.address=          Network hardware address
  -pool=                 Resource pool [GOVC_RESOURCE_POOL]
```

## library.subscriber.info

```
Usage: govc library.subscriber.info [OPTIONS] PUBLISHED-LIBRARY SUBSCRIPTION-ID

Library subscriber info.

Examples:
  id=$(govc library.subscriber.ls | grep my-library-name | awk '{print $1}')
  govc library.subscriber.info published-library-name $id

Options:
```

## library.subscriber.ls

```
Usage: govc library.subscriber.ls [OPTIONS]

List library subscriptions.

Examples:
  govc library.subscriber.ls library-name

Options:
```

## library.subscriber.rm

```
Usage: govc library.subscriber.rm [OPTIONS] SUBSCRIPTION-ID

Delete subscription of the published library.

The subscribed library associated with the subscription will not be deleted.

Examples:
  id=$(govc library-subscriber.ls | grep my-library-name | awk '{print $1}')
  govc library.subscriber.rm $id

Options:
```

## library.sync

```
Usage: govc library.sync [OPTIONS] NAME|ITEM

Sync library NAME or ITEM.

Examples:
  govc library.sync subscribed-library
  govc library.sync subscribed-library/item
  govc library.sync -vmtx local-library subscribed-library # convert subscribed OVFs to local VMTX

Options:
  -folder=               Inventory folder [GOVC_FOLDER]
  -pool=                 Resource pool [GOVC_RESOURCE_POOL]
  -vmtx=                 Sync subscribed library to local library as VM Templates
```

## library.update

```
Usage: govc library.update [OPTIONS] PATH

Update library or item PATH.

Examples:
  govc library.update -d "new library description" -n "new-name" my-library
  govc library.update -d "new item description" -n "new-item-name" my-library/my-item

Options:
  -d=                    Library or item description
  -n=                    Library or item name
```

## library.vmtx.info

```
Usage: govc library.vmtx.info [OPTIONS]

Display VMTX template details

Examples:
  govc library.vmtx.info /library_name/vmtx_template_name

Options:
```

## license.add

```
Usage: govc license.add [OPTIONS] KEY...

Options:
```

## license.assign

```
Usage: govc license.assign [OPTIONS] KEY

Assign licenses to HOST or CLUSTER.

Examples:
  govc license.assign $VCSA_LICENSE_KEY
  govc license.assign -host a_host.example.com $ESX_LICENSE_KEY
  govc license.assign -cluster a_cluster $VSAN_LICENSE_KEY

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
  -host=                 Host system [GOVC_HOST]
  -name=                 Display name
  -remove=false          Remove assignment
```

## license.assigned.ls

```
Usage: govc license.assigned.ls [OPTIONS]

Options:
  -id=                   Entity ID
```

## license.decode

```
Usage: govc license.decode [OPTIONS] KEY...

Options:
  -feature=              List licenses with given feature
```

## license.label.set

```
Usage: govc license.label.set [OPTIONS] LICENSE KEY VAL

Set license labels.

Examples:
  govc license.label.set 00000-00000-00000-00000-00000 team cnx # add/set label
  govc license.label.set 00000-00000-00000-00000-00000 team ""  # remove label
  govc license.ls -json | jq '.[] | select(.Labels[].Key == "team") | .LicenseKey'

Options:
```

## license.ls

```
Usage: govc license.ls [OPTIONS]

Options:
  -feature=              List licenses with given feature
```

## license.remove

```
Usage: govc license.remove [OPTIONS] KEY...

Options:
```

## logs

```
Usage: govc logs [OPTIONS]

View VPX and ESX logs.

The '-log' option defaults to "hostd" when connected directly to a host or
when connected to VirtualCenter and a '-host' option is given.  Otherwise,
the '-log' option defaults to "vpxd:vpxd.log".  The '-host' option is ignored
when connected directly to a host.  See 'govc logs.ls' for other '-log' options.

Examples:
  govc logs -n 1000 -f
  govc logs -host esx1
  govc logs -host esx1 -log vmkernel

Options:
  -f=false               Follow log file changes
  -host=                 Host system [GOVC_HOST]
  -log=                  Log file key
  -n=25                  Output the last N log lines
```

## logs.download

```
Usage: govc logs.download [OPTIONS] [PATH]...

Generate diagnostic bundles.

A diagnostic bundle includes log files and other configuration information.

Use PATH to include a specific set of hosts to include.

Examples:
  govc logs.download
  govc logs.download host-a host-b

Options:
  -default=false         Specifies if the bundle should include the default server
```

## logs.ls

```
Usage: govc logs.ls [OPTIONS]

List diagnostic log keys.

Examples:
  govc logs.ls
  govc logs.ls -host host-a

Options:
  -host=                 Host system [GOVC_HOST]
```

## ls

```
Usage: govc ls [OPTIONS] [PATH]...

List inventory items.

Examples:
  govc ls -l '*'
  govc ls -t ClusterComputeResource host
  govc ls -t Datastore host/ClusterA/* | grep -v local | xargs -n1 basename | sort | uniq

Options:
  -L=false               Follow managed object references
  -i=false               Print the managed object reference
  -l=false               Long listing format
  -t=                    Object type
```

## metric.change

```
Usage: govc metric.change [OPTIONS] NAME...

Change counter NAME levels.

Examples:
  govc metric.change -level 1 net.bytesRx.average net.bytesTx.average

Options:
  -device-level=0        Level for the per device counter
  -i=real                Interval ID (real|day|week|month|year)
  -level=0               Level for the aggregate counter
```

## metric.info

```
Usage: govc metric.info [OPTIONS] PATH [NAME]...

Metric info for NAME.

If PATH is a value other than '-', provider summary and instance list are included
for the given object type.

If NAME is not specified, all available metrics for the given INTERVAL are listed.
An object PATH must be provided in this case.

Examples:
  govc metric.info vm/my-vm
  govc metric.info -i 300 vm/my-vm
  govc metric.info - cpu.usage.average
  govc metric.info /dc1/host/cluster cpu.usage.average

Options:
  -g=                    Show info for a specific Group
  -i=real                Interval ID (real|day|week|month|year)
```

## metric.interval.change

```
Usage: govc metric.interval.change [OPTIONS]

Change historical metric intervals.

Examples:
  govc metric.interval.change -i 300 -level 2
  govc metric.interval.change -i 86400 -enabled=false

Options:
  -enabled=<nil>         Enable or disable
  -i=real                Interval ID (real|day|week|month|year)
  -level=0               Level
```

## metric.interval.info

```
Usage: govc metric.interval.info [OPTIONS]

List historical metric intervals.

Examples:
  govc metric.interval.info
  govc metric.interval.info -i 300

Options:
  -i=real                Interval ID (real|day|week|month|year)
```

## metric.ls

```
Usage: govc metric.ls [OPTIONS] PATH

List available metrics for PATH.

The default output format is the metric name.
The '-l' flag includes the metric description.
The '-L' flag includes the metric units, instance count (if any) and description.
The instance count is prefixed with a single '@'.
If no aggregate is provided for the metric, instance count is prefixed with two '@@'.

Examples:
  govc metric.ls /dc1/host/cluster1
  govc metric.ls datastore/*
  govc metric.ls -L -g CPU /dc1/host/cluster1/host1
  govc metric.ls vm/* | grep mem. | xargs govc metric.sample vm/*

Options:
  -L=false               Longer listing format
  -g=                    List a specific Group
  -i=real                Interval ID (real|day|week|month|year)
  -l=false               Long listing format
```

## metric.reset

```
Usage: govc metric.reset [OPTIONS] NAME...

Reset counter NAME to the default level of data collection.

Examples:
  govc metric.reset net.bytesRx.average net.bytesTx.average

Options:
  -i=real                Interval ID (real|day|week|month|year)
```

## metric.sample

```
Usage: govc metric.sample [OPTIONS] PATH... NAME...

Sample for object PATH of metric NAME.

Interval ID defaults to 20 (realtime) if supported, otherwise 300 (5m interval).

By default, INSTANCE '*' samples all instances and the aggregate counter.
An INSTANCE value of '-' will only sample the aggregate counter.
An INSTANCE value other than '*' or '-' will only sample the given instance counter.

If PLOT value is set to '-', output a gnuplot script.  If non-empty with another
value, PLOT will pipe the script to gnuplot for you.  The value is also used to set
the gnuplot 'terminal' variable, unless the value is that of the DISPLAY env var.
Only 1 metric NAME can be specified when the PLOT flag is set.

Examples:
  govc metric.sample host/cluster1/* cpu.usage.average
  govc metric.sample -plot .png host/cluster1/* cpu.usage.average | xargs open
  govc metric.sample vm/* net.bytesTx.average net.bytesTx.average
  govc metric.sample -instance vmnic0 vm/* net.bytesTx.average
  govc metric.sample -instance - vm/* net.bytesTx.average

Options:
  -d=30                  Limit object display name to D chars
  -i=real                Interval ID (real|day|week|month|year)
  -instance=*            Instance
  -n=5                   Max number of samples
  -plot=                 Plot data using gnuplot
  -t=false               Include sample times
```

## namespace.cluster.disable

```
Usage: govc namespace.cluster.disable [OPTIONS]

Disables vSphere Namespaces on the specified cluster.

Examples:
  govc namespace.cluster.disable -cluster "Workload-Cluster"

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
```

## namespace.cluster.enable

```
Usage: govc namespace.cluster.enable [OPTIONS]

Enable vSphere Namespaces on the cluster.
This operation sets up Kubernetes instance for the cluster along with worker nodes.

Examples:
  govc namespace.cluster.enable \
    -cluster "Workload-Cluster" \
    -service-cidr 10.96.0.0/23 \
    -pod-cidrs 10.244.0.0/20 \
    -control-plane-dns 10.10.10.10 \
    -control-plane-dns-names wcp.example.com \
    -workload-network.egress-cidrs 10.0.0.128/26 \
    -workload-network.ingress-cidrs "10.0.0.64/26" \
    -workload-network.switch VDS \
    -workload-network.edge-cluster Edge-Cluster-1 \
    -size TINY   \
    -mgmt-network.network "DVPG-Management Network" \
    -mgmt-network.gateway 10.0.0.1 \
    -mgmt-network.starting-address 10.0.0.45 \
    -mgmt-network.subnet-mask 255.255.255.0 \
    -ephemeral-storage-policy "vSAN Default Storage Policy" \
    -control-plane-storage-policy "vSAN Default Storage Policy" \
    -image-storage-policy "vSAN Default Storage Policy"

Options:
  -cluster=                                Cluster [GOVC_CLUSTER]
  -control-plane-dns=                      Comma-separated list of DNS server IP addresses to use on Kubernetes API server, specified in order of preference.
  -control-plane-dns-names=                Comma-separated list of DNS names to associate with the Kubernetes API server. These DNS names are embedded in the TLS certificate presented by the API server.
  -control-plane-dns-search-domains=       Comma-separated list of domains to be searched when trying to lookup a host name on Kubernetes API server, specified in order of preference.
  -control-plane-ntp-servers=              Optional. Comma-separated list of NTP server DNS names or IP addresses to use on Kubernetes API server, specified in order of preference. If unset, VMware Tools based time synchronization is enabled.
  -control-plane-storage-policy=           Storage policy associated with Kubernetes API server.
  -ephemeral-storage-policy=               Storage policy associated with ephemeral disks of all the Kubernetes Pods in the cluster.
  -image-storage-policy=                   Storage policy to be used for container images.
  -login-banner=                           Optional. Disclaimer to be displayed prior to login via the Kubectl plugin.
  -mgmt-network.address-count=5            The number of IP addresses in the management range. Optional, but required with network mode STATICRANGE.
  -mgmt-network.floating-IP=               Optional. The Floating IP used by the HA master cluster in the when network Mode is DHCP.
  -mgmt-network.gateway=                   Gateway to be used for the management IP range
  -mgmt-network.mode=STATICRANGE           IPv4 address assignment modes. Value is one of: DHCP, STATICRANGE
  -mgmt-network.network=                   Identifier for the management network.
  -mgmt-network.starting-address=          Denotes the start of the IP range to be used. Optional, but required with network mode STATICRANGE.
  -mgmt-network.subnet-mask=               Subnet mask of the management network. Optional, but required with network mode STATICRANGE.
  -network-provider=NSXT_CONTAINER_PLUGIN  Optional. Provider of cluster networking for this vSphere Namespaces cluster. Currently only value supported is: NSXT_CONTAINER_PLUGIN.
  -pod-cidrs=                              CIDR blocks from which Kubernetes allocates pod IP addresses. Comma-separated list. Shouldn't overlap with service, ingress or egress CIDRs.
  -service-cidr=                           CIDR block from which Kubernetes allocates service cluster IP addresses. Shouldn't overlap with pod, ingress or egress CIDRs
  -size=                                   The size of the Kubernetes API server and the worker nodes. Value is one of: TINY, SMALL, MEDIUM, LARGE.
  -worker-dns=                             Comma-separated list of DNS server IP addresses to use on the worker nodes, specified in order of preference.
  -workload-network.edge-cluster=          NSX Edge Cluster to be used for Kubernetes Services of type LoadBalancer, Kubernetes Ingresses, and NSX SNAT.
  -workload-network.egress-cidrs=          CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs. Comma-separated list. Shouldn't overlap with pod, service or ingress CIDRs.
  -workload-network.ingress-cidrs=         CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer. Comma-separated list. Shouldn't overlap with pod, service or egress CIDRs.
  -workload-network.switch=                vSphere Distributed Switch used to connect this cluster.
```

## namespace.cluster.ls

```
Usage: govc namespace.cluster.ls [OPTIONS]

List namepace enabled clusters.

Examples:
  govc namespace.cluster.ls
  govc namespace.cluster.ls -l
  govc namespace.cluster.ls -json | jq .

Options:
  -l=false               Long listing format
```

## namespace.logs.download

```
Usage: govc namespace.logs.download [OPTIONS] [NAME]

Download namespace cluster support bundle.

If NAME name is "-", bundle is written to stdout.

See also: govc logs.download

Examples:
  govc namespace.logs.download -cluster k8s
  govc namespace.logs.download -cluster k8s - | tar -xvf -
  govc namespace.logs.download -cluster k8s logs.tar

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
```

## namespace.service.activate

```
Usage: govc namespace.service.activate [OPTIONS] NAME...

Activates a vSphere Namespace Supervisor Service.

Examples:
  govc namespace.service.activate my-supervisor-service other-supervisor-service

Options:
```

## namespace.service.create

```
Usage: govc namespace.service.create [OPTIONS] MANIFEST

Creates a vSphere Namespace Supervisor Service.

Examples:
  govc namespace.service.create manifest.yaml

Options:
```

## namespace.service.deactivate

```
Usage: govc namespace.service.deactivate [OPTIONS] NAME...

Deactivates a vSphere Namespace Supervisor Service.

Examples:
  govc namespace.service.deactivate my-supervisor-service other-supervisor-service

Options:
```

## namespace.service.info

```
Usage: govc namespace.service.info [OPTIONS] NAME

Gets information of a specific supervisor service.

Examples:
  govc namespace.service.info my-supervisor-service
  govc namespace.service.info -json my-supervisor-service | jq .

Options:
```

## namespace.service.ls

```
Usage: govc namespace.service.ls [OPTIONS]

List namepace registered supervisor services.

Examples:
  govc namespace.service.ls
  govc namespace.service.ls -l
  govc namespace.service.ls -json | jq .

Options:
  -l=false               Long listing format
```

## namespace.service.rm

```
Usage: govc namespace.service.rm [OPTIONS] NAME...

Removes a vSphere Namespace Supervisor Service.

Examples:
  govc namespace.service.rm my-supervisor-service other-supervisor-service

Options:
```

## object.collect

```
Usage: govc object.collect [OPTIONS] [MOID] [PROPERTY]...

Collect managed object properties.

MOID can be an inventory path or ManagedObjectReference.
MOID defaults to '-', an alias for 'ServiceInstance:ServiceInstance' or the root folder if a '-type' flag is given.

If a '-type' flag is given, properties are collected using a ContainerView object where MOID is the root of the view.

By default only the current property value(s) are collected.  To wait for updates, use the '-n' flag or
specify a property filter.  A property filter can be specified by prefixing the property name with a '-',
followed by the value to match.

The '-R' flag sets the Filter using the given XML encoded request, which can be captured by 'vcsim -trace' for example.
It can be useful for replaying property filters created by other clients and converting filters to Go code via '-O -dump'.

The '-type' flag value can be a managed entity type or one of the following aliases:

  a    VirtualApp
  c    ClusterComputeResource
  d    Datacenter
  f    Folder
  g    DistributedVirtualPortgroup
  h    HostSystem
  m    VirtualMachine
  n    Network
  o    OpaqueNetwork
  p    ResourcePool
  r    ComputeResource
  s    Datastore
  w    DistributedVirtualSwitch

Examples:
  govc object.collect - content
  govc object.collect -s HostSystem:ha-host hardware.systemInfo.uuid
  govc object.collect -s /ha-datacenter/vm/foo overallStatus
  govc object.collect -s /ha-datacenter/vm/foo -guest.guestOperationsReady true # property filter
  govc object.collect -type m / name runtime.powerState # collect properties for multiple objects
  govc object.collect -json -n=-1 EventManager:ha-eventmgr latestEvent | jq .
  govc object.collect -json -s $(govc object.collect -s - content.perfManager) description.counterType | jq .
  govc object.collect -R create-filter-request.xml # replay filter
  govc object.collect -R create-filter-request.xml -O # convert filter to Go code
  govc object.collect -s vm/my-vm summary.runtime.host | xargs govc ls -L # inventory path of VM's host
  govc object.collect -dump -o "network/VM Network" # output Managed Object structure as Go code
  govc object.collect -json $vm config | \ # use -json + jq to search array elements
    jq -r '.[] | select(.Val.Hardware.Device[].MacAddress == "00:0c:29:0c:73:c0") | .Val.Name'

Options:
  -O=false               Output the CreateFilter request itself
  -R=                    Raw XML encoded CreateFilter request
  -d=,                   Delimiter for array values
  -n=0                   Wait for N property updates
  -o=false               Output the structure of a single Managed Object
  -s=false               Output property value only
  -type=[]               Resource type.  If specified, MOID is used for a container view root
  -wait=0s               Max wait time for updates
```

## object.destroy

```
Usage: govc object.destroy [OPTIONS] PATH...

Destroy managed objects.

Examples:
  govc object.destroy /dc1/network/dvs /dc1/host/cluster

Options:
```

## object.method

```
Usage: govc object.method [OPTIONS] PATH...

Enable or disable methods for managed objects.

Examples:
  govc object.method -name Destroy_Task -enable=false /dc1/vm/foo
  govc object.collect /dc1/vm/foo disabledMethod | grep --color Destroy_Task
  govc object.method -name Destroy_Task -enable /dc1/vm/foo

Options:
  -enable=true           Enable method
  -name=                 Method name
  -reason=               Reason for disabling method
  -source=govc           Source ID
```

## object.mv

```
Usage: govc object.mv [OPTIONS] PATH... FOLDER

Move managed entities to FOLDER.

Examples:
  govc folder.create /dc1/host/example
  govc object.mv /dc2/host/*.example.com /dc1/host/example

Options:
```

## object.reload

```
Usage: govc object.reload [OPTIONS] PATH...

Reload managed object state.

Examples:
  govc datastore.upload $vm.vmx $vm/$vm.vmx
  govc object.reload /dc1/vm/$vm

Options:
```

## object.rename

```
Usage: govc object.rename [OPTIONS] PATH NAME

Rename managed objects.

Examples:
  govc object.rename /dc1/network/dvs1 Switch1

Options:
```

## object.save

```
Usage: govc object.save [OPTIONS] [PATH]

Save managed objects.

By default, the object tree and all properties are saved, starting at PATH.
PATH defaults to ServiceContent, but can be specified to save a subset of objects.
The primary use case for this command is to save inventory from a live vCenter and
load it into a vcsim instance.

Examples:
  govc object.save -d my-vcenter
  vcsim -load my-vcenter

Options:
  -1=false               Save ROOT only, without its children
  -d=                    Save objects in directory
  -f=false               Remove existing object directory
  -folder=               Inventory folder [GOVC_FOLDER]
  -r=true                Include children of the container view root
  -type=[]               Resource types to save.  Defaults to all types
  -v=false               Verbose output
```

## option.ls

```
Usage: govc option.ls [OPTIONS] [NAME]

List option with the given NAME.

If NAME ends with a dot, all options for that subtree are listed.

Examples:
  govc option.ls
  govc option.ls config.vpxd.sso.
  govc option.ls config.vpxd.sso.sts.uri

Options:
```

## option.set

```
Usage: govc option.set [OPTIONS] NAME VALUE

Set option NAME to VALUE.

Examples:
  govc option.set log.level info
  govc option.set logger.Vsan verbose

Options:
```

## permissions.ls

```
Usage: govc permissions.ls [OPTIONS] [PATH]...

List the permissions defined on or effective on managed entities.

Examples:
  govc permissions.ls
  govc permissions.ls /dc1/host/cluster1

Options:
  -a=true                Include inherited permissions defined by parent entities
  -i=false               Use moref instead of inventory path
```

## permissions.remove

```
Usage: govc permissions.remove [OPTIONS] [PATH]...

Removes a permission rule from managed entities.

Examples:
  govc permissions.remove -principal root
  govc permissions.remove -principal $USER@vsphere.local /dc1/host/cluster1

Options:
  -f=false               Ignore NotFound fault if permission for this entity and user or group does not exist
  -group=false           True, if principal refers to a group name; false, for a user name
  -i=false               Use moref instead of inventory path
  -principal=            User or group for which the permission is defined
```

## permissions.set

```
Usage: govc permissions.set [OPTIONS] [PATH]...

Set the permissions managed entities.

Examples:
  govc permissions.set -principal root -role Admin
  govc permissions.set -principal $USER@vsphere.local -role Admin /dc1/host/cluster1

Options:
  -group=false           True, if principal refers to a group name; false, for a user name
  -i=false               Use moref instead of inventory path
  -principal=            User or group for which the permission is defined
  -propagate=true        Whether or not this permission propagates down the hierarchy to sub-entities
  -role=Admin            Permission role name
```

## pool.change

```
Usage: govc pool.change [OPTIONS] POOL...

Change the configuration of one or more resource POOLs.

POOL may be an absolute or relative path to a resource pool or a (clustered)
compute host. If it resolves to a compute host, the associated root resource
pool is returned. If a relative path is specified, it is resolved with respect
to the current datacenter's "host" folder (i.e. /ha-datacenter/host).

Paths to nested resource pools must traverse through the root resource pool of
the selected compute host, i.e. "compute-host/Resources/nested-pool".

The same globbing rules that apply to the "ls" command apply here. For example,
POOL may be specified as "*/Resources/*" to expand to all resource pools that
are nested one level under the root resource pool, on all (clustered) compute
hosts in the current datacenter.

Options:
  -cpu.expandable=<nil>   CPU expandable reservation
  -cpu.limit=<nil>        CPU limit in MHz
  -cpu.reservation=<nil>  CPU reservation in MHz
  -cpu.shares=            CPU shares level or number
  -mem.expandable=<nil>   Memory expandable reservation
  -mem.limit=<nil>        Memory limit in MB
  -mem.reservation=<nil>  Memory reservation in MB
  -mem.shares=            Memory shares level or number
  -name=                  Resource pool name
```

## pool.create

```
Usage: govc pool.create [OPTIONS] POOL...

Create one or more resource POOLs.

POOL may be an absolute or relative path to a resource pool. The parent of the
specified POOL must be an existing resource pool. If a relative path is
specified, it is resolved with respect to the current datacenter's "host"
folder (i.e. /ha-datacenter/host). The basename of the specified POOL is used
as the name for the new resource pool.

The same globbing rules that apply to the "ls" command apply here. For example,
the path to the parent resource pool in POOL may be specified as "*/Resources"
to expand to the root resource pools on all (clustered) compute hosts in the
current datacenter.

For example:
  */Resources/test             Create resource pool "test" on all (clustered)
                               compute hosts in the current datacenter.
  somehost/Resources/*/nested  Create resource pool "nested" in every
                               resource pool that is a direct descendant of
                               the root resource pool on "somehost".

Options:
  -cpu.expandable=true   CPU expandable reservation
  -cpu.limit=-1          CPU limit in MHz
  -cpu.reservation=0     CPU reservation in MHz
  -cpu.shares=normal     CPU shares level or number
  -mem.expandable=true   Memory expandable reservation
  -mem.limit=-1          Memory limit in MB
  -mem.reservation=0     Memory reservation in MB
  -mem.shares=normal     Memory shares level or number
```

## pool.destroy

```
Usage: govc pool.destroy [OPTIONS] POOL...

Destroy one or more resource POOLs.

POOL may be an absolute or relative path to a resource pool or a (clustered)
compute host. If it resolves to a compute host, the associated root resource
pool is returned. If a relative path is specified, it is resolved with respect
to the current datacenter's "host" folder (i.e. /ha-datacenter/host).

Paths to nested resource pools must traverse through the root resource pool of
the selected compute host, i.e. "compute-host/Resources/nested-pool".

The same globbing rules that apply to the "ls" command apply here. For example,
POOL may be specified as "*/Resources/*" to expand to all resource pools that
are nested one level under the root resource pool, on all (clustered) compute
hosts in the current datacenter.

Options:
  -children=false        Remove all children pools
```

## pool.info

```
Usage: govc pool.info [OPTIONS] POOL...

Retrieve information about one or more resource POOLs.

POOL may be an absolute or relative path to a resource pool or a (clustered)
compute host. If it resolves to a compute host, the associated root resource
pool is returned. If a relative path is specified, it is resolved with respect
to the current datacenter's "host" folder (i.e. /ha-datacenter/host).

Paths to nested resource pools must traverse through the root resource pool of
the selected compute host, i.e. "compute-host/Resources/nested-pool".

The same globbing rules that apply to the "ls" command apply here. For example,
POOL may be specified as "*/Resources/*" to expand to all resource pools that
are nested one level under the root resource pool, on all (clustered) compute
hosts in the current datacenter.

Options:
  -a=false               List virtual app resource pools
  -p=true                List resource pools
```

## role.create

```
Usage: govc role.create [OPTIONS] NAME [PRIVILEGE]...

Create authorization role.

Optionally populate the role with the given PRIVILEGE(s).

Examples:
  govc role.create MyRole
  govc role.create NoDC $(govc role.ls Admin | grep -v Datacenter.)

Options:
  -i=false               Use moref instead of inventory path
```

## role.ls

```
Usage: govc role.ls [OPTIONS] [NAME]

List authorization roles.

If NAME is provided, list privileges for the role.

Examples:
  govc role.ls
  govc role.ls Admin

Options:
  -i=false               Use moref instead of inventory path
```

## role.remove

```
Usage: govc role.remove [OPTIONS] NAME

Remove authorization role.

Examples:
  govc role.remove MyRole
  govc role.remove MyRole -force

Options:
  -force=false           Force removal if role is in use
  -i=false               Use moref instead of inventory path
```

## role.update

```
Usage: govc role.update [OPTIONS] NAME [PRIVILEGE]...

Update authorization role.

Set, Add or Remove role PRIVILEGE(s).

Examples:
  govc role.update MyRole $(govc role.ls Admin | grep VirtualMachine.)
  govc role.update -r MyRole $(govc role.ls Admin | grep VirtualMachine.GuestOperations.)
  govc role.update -a MyRole $(govc role.ls Admin | grep Datastore.)
  govc role.update -name RockNRole MyRole

Options:
  -a=false               Add given PRIVILEGE(s)
  -i=false               Use moref instead of inventory path
  -name=                 Change role name
  -r=false               Remove given PRIVILEGE(s)
```

## role.usage

```
Usage: govc role.usage [OPTIONS] NAME...

List usage for role NAME.

Examples:
  govc role.usage
  govc role.usage Admin

Options:
  -i=false               Use moref instead of inventory path
```

## session.login

```
Usage: govc session.login [OPTIONS] [PATH]

Session login.

The session.login command is optional, all other govc commands will auto login when given credentials.
The session.login command can be used to:
- Persist a session without writing to disk via the '-cookie' flag
- Acquire a clone ticket
- Login using a clone ticket
- Login using a vCenter Extension certificate
- Issue a SAML token
- Renew a SAML token
- Login using a SAML token
- Avoid passing credentials to other govc commands
- Send an authenticated raw HTTP request

The session.login command can be used for authenticated curl-style HTTP requests when a PATH arg is given.
PATH may also contain a query string. The '-u' flag (GOVC_URL) is used for the URL scheme, host and port.
The request method (-X) defaults to GET. When set to POST, PUT or PATCH, a request body must be provided via stdin.

Examples:
  govc session.login -u root:password@host # Creates a cached session in ~/.govmomi/sessions
  govc session.ls -u root@host # Use the cached session with another command
  ticket=$(govc session.login -u root@host -clone)
  govc session.login -u root@host -ticket $ticket
  govc session.login -u host -extension com.vmware.vsan.health -cert rui.crt -key rui.key
  token=$(govc session.login -u host -cert user.crt -key user.key -issue) # HoK token
  bearer=$(govc session.login -u user:pass@host -issue) # Bearer token
  token=$(govc session.login -u host -cert user.crt -key user.key -issue -token "$bearer")
  govc session.login -u host -cert user.crt -key user.key -token "$token"
  token=$(govc session.login -u host -cert user.crt -key user.key -renew -lifetime 24h -token "$token")
  # HTTP requests
  govc session.login -r -X GET /api/vcenter/namespace-management/clusters | jq .
  govc session.login -r -X POST /rest/vcenter/cluster/modules <<<'{"spec": {"cluster": "domain-c9"}}'

Options:
  -X=                    HTTP method
  -clone=false           Acquire clone ticket
  -cookie=               Set HTTP cookie for an existing session
  -extension=            Extension name
  -issue=false           Issue SAML token
  -l=false               Output session cookie
  -lifetime=10m0s        SAML token lifetime
  -r=false               REST login
  -renew=false           Renew SAML token
  -ticket=               Use clone ticket for login
  -token=                Use SAML token for login or as issue identity
```

## session.logout

```
Usage: govc session.logout [OPTIONS]

Logout the current session.

The session.logout command can be used to end the current persisted session.
The session.rm command can be used to remove sessions other than the current session.

Examples:
  govc session.logout

Options:
  -r=false               REST logout
```

## session.ls

```
Usage: govc session.ls [OPTIONS]

List active sessions.

All SOAP sessions are listed by default. The '-S' flag will limit this list to the current session.

Examples:
  govc session.ls
  govc session.ls -json | jq -r .CurrentSession.Key

Options:
  -S=false               List current SOAP session
  -r=false               List cached REST session (if any)
```

## session.rm

```
Usage: govc session.rm [OPTIONS] KEY...

Remove active sessions.

Examples:
  govc session.ls | grep root
  govc session.rm 5279e245-e6f1-4533-4455-eb94353b213a

Options:
```

## snapshot.create

```
Usage: govc snapshot.create [OPTIONS] NAME

Create snapshot of VM with NAME.

Examples:
  govc snapshot.create -vm my-vm happy-vm-state

Options:
  -d=                    Snapshot description
  -m=true                Include memory state
  -q=false               Quiesce guest file system
  -vm=                   Virtual machine [GOVC_VM]
```

## snapshot.remove

```
Usage: govc snapshot.remove [OPTIONS] NAME

Remove snapshot of VM with given NAME.

NAME can be the snapshot name, tree path, moid or '*' to remove all snapshots.

Examples:
  govc snapshot.remove -vm my-vm happy-vm-state

Options:
  -c=true                Consolidate disks
  -r=false               Remove snapshot children
  -vm=                   Virtual machine [GOVC_VM]
```

## snapshot.revert

```
Usage: govc snapshot.revert [OPTIONS] [NAME]

Revert to snapshot of VM with given NAME.

If NAME is not provided, revert to the current snapshot.
Otherwise, NAME can be the snapshot name, tree path or moid.

Examples:
  govc snapshot.revert -vm my-vm happy-vm-state

Options:
  -s=false               Suppress power on
  -vm=                   Virtual machine [GOVC_VM]
```

## snapshot.tree

```
Usage: govc snapshot.tree [OPTIONS]

List VM snapshots in a tree-like format.

The command will exit 0 with no output if VM does not have any snapshots.

Examples:
  govc snapshot.tree -vm my-vm
  govc snapshot.tree -vm my-vm -D -i -d

Options:
  -C=false               Print the current snapshot name only
  -D=false               Print the snapshot creation date
  -c=true                Print the current snapshot
  -d=false               Print the snapshot description
  -f=false               Print the full path prefix for snapshot
  -i=false               Print the snapshot id
  -s=false               Print the snapshot size
  -vm=                   Virtual machine [GOVC_VM]
```

## sso.group.create

```
Usage: govc sso.group.create [OPTIONS] NAME

Create SSO group.

Examples:
  govc sso.group.create NAME

Options:
  -d=                    Group description
```

## sso.group.ls

```
Usage: govc sso.group.ls [OPTIONS]

List SSO groups.

Examples:
  govc sso.group.ls -s

Options:
```

## sso.group.rm

```
Usage: govc sso.group.rm [OPTIONS] NAME

Remove SSO group.

Examples:
  govc sso.group.rm NAME

Options:
```

## sso.group.update

```
Usage: govc sso.group.update [OPTIONS]

Update SSO group.

Examples:
  govc sso.group.update -d "Group description" NAME
  govc sso.group.update -a user1 NAME
  govc sso.group.update -r user2 NAME
  govc sso.group.update -g -a group1 NAME
  govc sso.group.update -g -r group2 NAME

Options:
  -a=                    Add user/group to group
  -d=                    Group description
  -g=false               Add/Remove group from group
  -r=                    Remove user/group from group
```

## sso.idp.ls

```
Usage: govc sso.idp.ls [OPTIONS]

List SSO identity provider sources.

Examples:
  govc sso.idp.ls
  govc sso.idp.ls -json

Options:
```

## sso.service.ls

```
Usage: govc sso.service.ls [OPTIONS]

List platform services.

Examples:
  govc sso.service.ls
  govc sso.service.ls -t vcenterserver -P vmomi
  govc sso.service.ls -t cs.identity
  govc sso.service.ls -t cs.identity -P wsTrust -U
  govc sso.service.ls -t cs.identity -json | jq -r .[].ServiceEndpoints[].Url

Options:
  -P=                    Endpoint protocol
  -T=                    Endpoint type
  -U=false               List endpoint URL(s) only
  -l=false               Long listing format
  -n=                    Node ID
  -p=                    Service product
  -s=                    Site ID
  -t=                    Service type
```

## sso.user.create

```
Usage: govc sso.user.create [OPTIONS] NAME

Create SSO users.

Examples:
  govc sso.user.create -C "$(cat cert.pem)" -A -R Administrator NAME # solution user
  govc sso.user.create -p password NAME # person user

Options:
  -A=<nil>               ActAsUser role for solution user WSTrust
  -C=                    Certificate for solution user
  -R=                    Role for solution user (RegularUser|Administrator)
  -d=                    User description
  -f=                    First name
  -l=                    Last name
  -m=                    Email address
  -p=                    Password
```

## sso.user.id

```
Usage: govc sso.user.id [OPTIONS] NAME

Print SSO user and group IDs.

Examples:
  govc sso.user.id
  govc sso.user.id Administrator
  govc sso.user.id -json Administrator

Options:
```

## sso.user.ls

```
Usage: govc sso.user.ls [OPTIONS]

List SSO users.

Examples:
  govc sso.user.ls -s

Options:
  -s=false               List solution users
```

## sso.user.rm

```
Usage: govc sso.user.rm [OPTIONS] NAME

Remove SSO users.

Examples:
  govc sso.user.rm NAME

Options:
```

## sso.user.update

```
Usage: govc sso.user.update [OPTIONS] NAME

Update SSO users.

Examples:
  govc sso.user.update -C "$(cat cert.pem)" NAME
  govc sso.user.update -p password NAME

Options:
  -A=<nil>               ActAsUser role for solution user WSTrust
  -C=                    Certificate for solution user
  -R=                    Role for solution user (RegularUser|Administrator)
  -d=                    User description
  -f=                    First name
  -l=                    Last name
  -m=                    Email address
  -p=                    Password
```

## storage.policy.create

```
Usage: govc storage.policy.create [OPTIONS] NAME

Create VM Storage Policy.

Examples:
  govc storage.policy.create -category my_cat -tag my_tag MyStoragePolicy # Tag based placement

Options:
  -category=             Category
  -d=                    Description
  -tag=                  Tag
```

## storage.policy.info

```
Usage: govc storage.policy.info [OPTIONS] [NAME]

VM Storage Policy info.

Examples:
  govc storage.policy.info
  govc storage.policy.info "vSAN Default Storage Policy"
  govc storage.policy.info -c -s

Options:
  -c=false               Check VM Compliance
  -s=false               Check Storage Compatibility
```

## storage.policy.ls

```
Usage: govc storage.policy.ls [OPTIONS] [NAME]

VM Storage Policy listing.

Examples:
  govc storage.policy.ls
  govc storage.policy.ls "vSAN Default Storage Policy"
  govc storage.policy.ls -i "vSAN Default Storage Policy"

Options:
  -i=false               List policy ID only
```

## storage.policy.rm

```
Usage: govc storage.policy.rm [OPTIONS] ID

Remove Storage Policy ID.

Examples:
  govc storage.policy.rm "my policy name"
  govc storage.policy.rm af7935ab-466d-4b0e-af3c-4ec6bce2112f

Options:
```

## tags.attach

```
Usage: govc tags.attach [OPTIONS] NAME PATH

Attach tag NAME to object PATH.

Examples:
  govc tags.attach k8s-region-us /dc1
  govc tags.attach -c k8s-region us-ca1 /dc1/host/cluster1

Options:
  -c=                    Tag category
```

## tags.attached.ls

```
Usage: govc tags.attached.ls [OPTIONS] NAME

List attached tags or objects.

Examples:
  govc tags.attached.ls k8s-region-us
  govc tags.attached.ls -json k8s-zone-us-ca1 | jq .
  govc tags.attached.ls -r /dc1/host/cluster1
  govc tags.attached.ls -json -r /dc1 | jq .

Options:
  -l=false               Long listing format
  -r=false               List tags attached to resource
```

## tags.category.create

```
Usage: govc tags.category.create [OPTIONS] NAME

Create tag category.

This command will output the ID of the new tag category.

Examples:
  govc tags.category.create -d "Kubernetes region" -t Datacenter k8s-region
  govc tags.category.create -d "Kubernetes zone" k8s-zone

Options:
  -d=                    Description
  -m=false               Allow multiple tags per object
  -t=[]                  Object types
```

## tags.category.info

```
Usage: govc tags.category.info [OPTIONS] [NAME]

Display category info.

If NAME is provided, display info for only that category.
Otherwise display info for all categories.

Examples:
  govc tags.category.info
  govc tags.category.info k8s-zone

Options:
```

## tags.category.ls

```
Usage: govc tags.category.ls [OPTIONS]

List all categories.

Examples:
  govc tags.category.ls
  govc tags.category.ls -json | jq .

Options:
```

## tags.category.rm

```
Usage: govc tags.category.rm [OPTIONS] NAME

Delete category NAME.

Fails if category is used by any tag, unless the '-f' flag is provided.

Examples:
  govc tags.category.rm k8s-region
  govc tags.category.rm -f k8s-zone

Options:
  -f=false               Delete tag regardless of attached objects
```

## tags.category.update

```
Usage: govc tags.category.update [OPTIONS] NAME

Update category.

The '-t' flag can only be used to add new object types.  Removing category types is not supported by vCenter.

Examples:
  govc tags.category.update -n k8s-vcp-region -d "Kubernetes VCP region" k8s-region
  govc tags.category.update -t ClusterComputeResource k8s-zone

Options:
  -d=                    Description
  -m=<nil>               Allow multiple tags per object
  -n=                    Name of category
  -t=[]                  Object types
```

## tags.create

```
Usage: govc tags.create [OPTIONS] NAME

Create tag.

The '-c' option to specify a tag category is required.
This command will output the ID of the new tag.

Examples:
  govc tags.create -d "Kubernetes Zone US CA1" -c k8s-zone k8s-zone-us-ca1

Options:
  -c=                    Category name
  -d=                    Description of tag
```

## tags.detach

```
Usage: govc tags.detach [OPTIONS] NAME PATH

Detach tag NAME from object PATH.

Examples:
  govc tags.detach k8s-region-us /dc1
  govc tags.detach -c k8s-region us-ca1 /dc1/host/cluster1

Options:
  -c=                    Tag category
```

## tags.info

```
Usage: govc tags.info [OPTIONS] NAME

Display tags info.

If NAME is provided, display info for only that tag.  Otherwise display info for all tags.

Examples:
  govc tags.info
  govc tags.info k8s-zone-us-ca1
  govc tags.info -c k8s-zone

Options:
  -C=true                Display category name instead of ID
  -c=                    Category name
```

## tags.ls

```
Usage: govc tags.ls [OPTIONS]

List tags.

Examples:
  govc tags.ls
  govc tags.ls -c k8s-zone
  govc tags.ls -json | jq .
  govc tags.ls -c k8s-region -json | jq .

Options:
  -c=                    Category name
```

## tags.rm

```
Usage: govc tags.rm [OPTIONS] NAME

Delete tag NAME.

Fails if tag is attached to any object, unless the '-f' flag is provided.

Examples:
  govc tags.rm k8s-zone-us-ca1
  govc tags.rm -f -c k8s-zone us-ca2

Options:
  -c=                    Tag category
  -f=false               Delete tag regardless of attached objects
```

## tags.update

```
Usage: govc tags.update [OPTIONS] NAME

Update tag.

Examples:
  govc tags.update -d "K8s zone US-CA1" k8s-zone-us-ca1
  govc tags.update -d "K8s zone US-CA1" -c k8s-zone us-ca1

Options:
  -c=                    Tag category
  -d=                    Description of tag
  -n=                    Name of tag
```

## task.cancel

```
Usage: govc task.cancel [OPTIONS] ID...

Cancel tasks.

Examples:
  govc task.cancel task-759

Options:
```

## tasks

```
Usage: govc tasks [OPTIONS] [PATH]

Display info for recent tasks.

When a task has completed, the result column includes the task duration on success or
error message on failure.  If a task is still in progress, the result column displays
the completion percentage and the task ID.  The task ID can be used as an argument to
the 'task.cancel' command.

By default, all recent tasks are included (via TaskManager), but can be limited by PATH
to a specific inventory object.

Examples:
  govc tasks
  govc tasks -f
  govc tasks -f /dc1/host/cluster1

Options:
  -f=false               Follow recent task updates
  -l=false               Use long task description
  -n=25                  Output the last N tasks
```

## tree

```
Usage: govc tree [OPTIONS] [PATH]

List contents of the inventory in a tree-like format.

Examples:
  govc tree -C /
  govc tree /datacenter/vm

Options:
  -C=false               Colorize output
  -L=0                   Max display depth of the inventory tree
  -l=false               Follow runtime references (e.g. HostSystem VMs)
  -p=false               Print the object type
```

## vapp.destroy

```
Usage: govc vapp.destroy [OPTIONS] VAPP...

Options:
```

## vapp.power

```
Usage: govc vapp.power [OPTIONS]

Options:
  -force=false           Force (If force is false, the shutdown order in the vApp is executed. If force is true, all virtual machines are powered-off (regardless of shutdown order))
  -off=false             Power off
  -on=false              Power on
  -suspend=false         Power suspend
  -vapp.ipath=           Find vapp by inventory path
```

## vcsa.access.consolecli.get

```
Usage: govc vcsa.access.consolecli.get [OPTIONS]

Get enabled state of the console-based controlled CLI (TTY1).

Note: This command requires vCenter 7.0.2 or higher.

Examples:
govc vcsa.access.consolecli.get

Options:
```

## vcsa.access.consolecli.set

```
Usage: govc vcsa.access.consolecli.set [OPTIONS]

Set enabled state of the console-based controlled CLI (TTY1).

Note: This command requires vCenter 7.0.2 or higher.

Examples:
# Enable Console CLI
govc vcsa.access.consolecli.set -enabled=true

# Disable Console CLI
govc vcsa.access.consolecli.set -enabled=false

Options:
  -enabled=false         Enable Console CLI.
```

## vcsa.access.dcui.get

```
Usage: govc vcsa.access.dcui.get [OPTIONS]

Get enabled state of Direct Console User Interface (DCUI TTY2).

Note: This command requires vCenter 7.0.2 or higher.

Examples:
govc vcsa.access.dcui.get

Options:
```

## vcsa.access.dcui.set

```
Usage: govc vcsa.access.dcui.set [OPTIONS]

Set enabled state of Direct Console User Interface (DCUI TTY2).

Note: This command requires vCenter 7.0.2 or higher.

Examples:
# Enable DCUI
govc vcsa.access.dcui.set -enabled=true

# Disable DCUI
govc vcsa.access.dcui.set -enabled=false

Options:
  -enabled=false         Enable Direct Console User Interface (DCUI TTY2).
```

## vcsa.access.shell.get

```
Usage: govc vcsa.access.shell.get [OPTIONS]

Get enabled state of BASH, that is, access to BASH from within the controlled CLI.

Note: This command requires vCenter 7.0.2 or higher.

Examples:
govc vcsa.access.shell.get

Options:
```

## vcsa.access.shell.set

```
Usage: govc vcsa.access.shell.set [OPTIONS]

Set enabled state of BASH, that is, access to BASH from within the controlled CLI.

Note: This command requires vCenter 7.0.2 or higher.

Examples:
# Enable Shell
govc vcsa.access.shell.set -enabled=true -timeout=240

# Disable Shell
govc vcsa.access.shell.set -enabled=false

Options:
  -enabled=false         Enable BASH, that is, access to BASH from within the controlled CLI.
  -timeout=0             The timeout (in seconds) specifies how long you enable the Shell access. The maximum timeout is 86400 seconds(1 day).
```

## vcsa.access.ssh.get

```
Usage: govc vcsa.access.ssh.get [OPTIONS]

Get enabled state of the SSH-based controlled CLI.

Note: This command requires vCenter 7.0.2 or higher.

Examples:
govc vcsa.access.ssh.get

Options:
```

## vcsa.access.ssh.set

```
Usage: govc vcsa.access.ssh.set [OPTIONS]

Set enabled state of the SSH-based controlled CLI.

Note: This command requires vCenter 7.0.2 or higher.

Examples:
# Enable SSH
govc vcsa.access.ssh.set -enabled=true

# Disable SSH
govc vcsa.access.ssh.set -enabled=false

Options:
  -enabled=false         Enable SSH-based controlled CLI.
```

## vcsa.log.forwarding.info

```
Usage: govc vcsa.log.forwarding.info [OPTIONS]

Retrieve the VC Appliance log forwarding configuration

Examples:
  govc vcsa.log.forwarding.info

Options:
```

## vcsa.net.proxy.info

```
Usage: govc vcsa.net.proxy.info [OPTIONS]

Retrieve the VC networking proxy configuration

Examples:
  govc vcsa.net.proxy.info

Options:
```

## vcsa.shutdown.cancel

```
Usage: govc vcsa.shutdown.cancel [OPTIONS]

Cancel pending shutdown action.

Note: This command requires vCenter 7.0.2 or higher.

Examples:
govc vcsa.shutdown.cancel

Options:
```

## vcsa.shutdown.get

```
Usage: govc vcsa.shutdown.get [OPTIONS]

Get details about the pending shutdown action.

Note: This command requires vCenter 7.0.2 or higher.

Examples:
govc vcsa.shutdown.get

Options:
```

## vcsa.shutdown.poweroff

```
Usage: govc vcsa.shutdown.poweroff [OPTIONS] REASON

Power off the appliance.

Note: This command requires vCenter 7.0.2 or higher.

Examples:
govc vcsa.shutdown.poweroff -delay 10 "powering off for maintenance"

Options:
  -delay=0               Minutes after which poweroff should start.
```

## vcsa.shutdown.reboot

```
Usage: govc vcsa.shutdown.reboot [OPTIONS] REASON

Reboot the appliance.

Note: This command requires vCenter 7.0.2 or higher.

Examples:
govc vcsa.shutdown.reboot -delay 10 "rebooting for maintenance"

Options:
  -delay=0               Minutes after which reboot should start.
```

## version

```
Usage: govc version [OPTIONS]

Options:
  -l=false   Print detailed govc version information
  -require=  Require govc version >= this value
```

## vm.change

```
Usage: govc vm.change [OPTIONS]

Change VM configuration.

To add ExtraConfig variables that can read within the guest, use the 'guestinfo.' prefix.

Examples:
  govc vm.change -vm $vm -mem.reservation 2048
  govc vm.change -vm $vm -e smc.present=TRUE -e ich7m.present=TRUE
  # Enable both cpu and memory hotplug on a guest:
  govc vm.change -vm $vm -cpu-hot-add-enabled -memory-hot-add-enabled
  govc vm.change -vm $vm -e guestinfo.vmname $vm
  # Read the contents of a file and use them as ExtraConfig value
  govc vm.change -vm $vm -f guestinfo.data="$(realpath .)/vmdata.config"
  # Read the variable set above inside the guest:
  vmware-rpctool "info-get guestinfo.vmname"
  govc vm.change -vm $vm -latency high
  govc vm.change -vm $vm -latency normal
  govc vm.change -vm $vm -uuid 4139c345-7186-4924-a842-36b69a24159b
  govc vm.change -vm $vm -scheduled-hw-upgrade-policy always

Options:
  -annotation=                   VM description
  -c=0                           Number of CPUs
  -cpu-hot-add-enabled=<nil>     Enable CPU hot add
  -cpu.limit=<nil>               CPU limit in MHz
  -cpu.reservation=<nil>         CPU reservation in MHz
  -cpu.shares=                   CPU shares level or number
  -e=[]                          ExtraConfig. <key>=<value>
  -f=[]                          ExtraConfig. <key>=<absolute path to file>
  -g=                            Guest OS
  -latency=                      Latency sensitivity (low|normal|high)
  -m=0                           Size in MB of memory
  -mem.limit=<nil>               Memory limit in MB
  -mem.reservation=<nil>         Memory reservation in MB
  -mem.shares=                   Memory shares level or number
  -memory-hot-add-enabled=<nil>  Enable memory hot add
  -memory-pin=<nil>              Reserve all guest memory
  -name=                         Display name
  -nested-hv-enabled=<nil>       Enable nested hardware-assisted virtualization
  -scheduled-hw-upgrade-policy=  Schedule hardware upgrade policy (onSoftPowerOff|never|always)
  -sync-time-with-host=<nil>     Enable SyncTimeWithHost
  -uuid=                         BIOS UUID
  -vm=                           Virtual machine [GOVC_VM]
  -vpmc-enabled=<nil>            Enable CPU performance counters
```

## vm.clone

```
Usage: govc vm.clone [OPTIONS] NAME

Clone VM or template to NAME.

Examples:
  govc vm.clone -vm template-vm new-vm
  govc vm.clone -vm template-vm -link new-vm
  govc vm.clone -vm template-vm -snapshot s-name new-vm
  govc vm.clone -vm template-vm -link -snapshot s-name new-vm
  govc vm.clone -vm template-vm -cluster cluster1 new-vm # use compute cluster placement
  govc vm.clone -vm template-vm -datastore-cluster dscluster new-vm # use datastore cluster placement
  govc vm.clone -vm template-vm -snapshot $(govc snapshot.tree -vm template-vm -C) new-vm
  govc vm.clone -vm template-vm -template new-template # clone a VM template
  govc vm.clone -vm=/ClusterName/vm/FolderName/VM_templateName -on=true -host=myesxi01 -ds=datastore01 myVM_name

Options:
  -annotation=           VM description
  -c=0                   Number of CPUs
  -cluster=              Use cluster for VM placement via DRS
  -customization=        Customization Specification Name
  -datastore-cluster=    Datastore cluster [GOVC_DATASTORE_CLUSTER]
  -ds=                   Datastore [GOVC_DATASTORE]
  -folder=               Inventory folder [GOVC_FOLDER]
  -force=false           Create VM if vmx already exists
  -host=                 Host system [GOVC_HOST]
  -link=false            Creates a linked clone from snapshot or source VM
  -m=0                   Size in MB of memory
  -net=                  Network [GOVC_NETWORK]
  -net.adapter=e1000     Network adapter type
  -net.address=          Network hardware address
  -on=true               Power on VM
  -pool=                 Resource pool [GOVC_RESOURCE_POOL]
  -snapshot=             Snapshot name to clone from
  -template=false        Create a Template
  -vm=                   Virtual machine [GOVC_VM]
  -waitip=false          Wait for VM to acquire IP address
```

## vm.console

```
Usage: govc vm.console [OPTIONS] VM

Generate console URL or screen capture for VM.

One of VMRC, VMware Player, VMware Fusion or VMware Workstation must be installed to
open VMRC console URLs.

Examples:
  govc vm.console my-vm
  govc vm.console -capture screen.png my-vm  # screen capture
  govc vm.console -capture - my-vm | display # screen capture to stdout
  open $(govc vm.console my-vm)              # MacOSX VMRC
  open $(govc vm.console -h5 my-vm)          # MacOSX H5
  xdg-open $(govc vm.console my-vm)          # Linux VMRC
  xdg-open $(govc vm.console -h5 my-vm)      # Linux H5

Options:
  -capture=              Capture console screen shot to file
  -h5=false              Generate HTML5 UI console link
  -vm=                   Virtual machine [GOVC_VM]
  -wss=false             Generate WebSocket console link
```

## vm.create

```
Usage: govc vm.create [OPTIONS] NAME

Create VM.

For a list of possible '-g' IDs, use 'govc vm.option.info' or see:
https://code.vmware.com/apis/358/vsphere/doc/vim.vm.GuestOsDescriptor.GuestOsIdentifier.html

Examples:
  govc vm.create -on=false vm-name
  govc vm.create -cluster cluster1 vm-name # use compute cluster placement
  govc vm.create -datastore-cluster dscluster vm-name # use datastore cluster placement
  govc vm.create -m 2048 -c 2 -g freebsd64Guest -net.adapter vmxnet3 -disk.controller pvscsi vm-name

Options:
  -annotation=           VM description
  -c=1                   Number of CPUs
  -cluster=              Use cluster for VM placement via DRS
  -datastore-cluster=    Datastore cluster [GOVC_DATASTORE_CLUSTER]
  -disk=                 Disk path (to use existing) OR size (to create new, e.g. 20GB)
  -disk-datastore=       Datastore for disk file
  -disk.controller=scsi  Disk controller type
  -ds=                   Datastore [GOVC_DATASTORE]
  -firmware=bios         Firmware type [bios|efi]
  -folder=               Inventory folder [GOVC_FOLDER]
  -force=false           Create VM if vmx already exists
  -g=otherGuest          Guest OS ID
  -host=                 Host system [GOVC_HOST]
  -iso=                  ISO path
  -iso-datastore=        Datastore for ISO file
  -link=true             Link specified disk
  -m=1024                Size in MB of memory
  -net=                  Network [GOVC_NETWORK]
  -net.adapter=e1000     Network adapter type
  -net.address=          Network hardware address
  -on=true               Power on VM
  -pool=                 Resource pool [GOVC_RESOURCE_POOL]
  -version=              ESXi hardware version [5.0|5.5|6.0|6.5|6.7|7.0]
```

## vm.customize

```
Usage: govc vm.customize [OPTIONS] [NAME]

Customize VM.

Optionally specify a customization spec NAME.

The '-ip', '-netmask' and '-gateway' flags are for static IP configuration.
If the VM has multiple NICs, an '-ip' and '-netmask' must be specified for each.

The '-dns-server' and '-dns-suffix' flags can be specified multiple times.

Windows -tz value requires the Index (hex): https://support.microsoft.com/en-us/help/973627/microsoft-time-zone-index-values

Examples:
  govc vm.customize -vm VM NAME
  govc vm.customize -vm VM -name my-hostname -ip dhcp
  govc vm.customize -vm VM -gateway GATEWAY -ip NEWIP -netmask NETMASK -dns-server DNS1,DNS2 NAME
  # Multiple -ip without -mac are applied by vCenter in the order in which the NICs appear on the bus
  govc vm.customize -vm VM -ip 10.0.0.178 -netmask 255.255.255.0 -ip 10.0.0.162 -netmask 255.255.255.0
  # Multiple -ip with -mac are applied by vCenter to the NIC with the given MAC address
  govc vm.customize -vm VM -mac 00:50:56:be:dd:f8 -ip 10.0.0.178 -netmask 255.255.255.0 -mac 00:50:56:be:60:cf -ip 10.0.0.162 -netmask 255.255.255.0
  # Dual stack IPv4/IPv6, single NIC
  govc vm.customize -vm VM -ip 10.0.0.1 -netmask 255.255.255.0 -ip6 '2001:db8::1/64' -name my-hostname NAME
  # DHCPv6, single NIC
  govc vm.customize -vm VM -ip6 dhcp6 NAME
  # Static IPv6, three NICs, last one with two addresses
  govc vm.customize -vm VM -ip6 2001:db8::1/64 -ip6 2001:db8::2/64 -ip6 2001:db8::3/64,2001:db8::4/64 NAME
  govc vm.customize -vm VM -auto-login 3 NAME
  govc vm.customize -vm VM -prefix demo NAME
  govc vm.customize -vm VM -tz America/New_York NAME

Options:
  -auto-login=0          Number of times the VM should automatically login as an administrator
  -dns-server=[]         DNS server list
  -dns-suffix=[]         DNS suffix list
  -domain=               Domain name
  -gateway=[]            Gateway
  -ip=[]                 IPv4 address
  -ip6=[]                IPv6 addresses with optional netmask (defaults to /64), separated by comma
  -mac=[]                MAC address
  -name=                 Host name
  -netmask=[]            Netmask
  -prefix=               Host name generator prefix
  -type=Linux            Customization type if spec NAME is not specified (Linux|Windows)
  -tz=                   Time zone
  -vm=                   Virtual machine [GOVC_VM]
```

## vm.destroy

```
Usage: govc vm.destroy [OPTIONS] VM...

Power off and delete VM.

When a VM is destroyed, any attached virtual disks are also deleted.
Use the 'device.remove -vm VM -keep disk-*' command to detach and
keep disks if needed, prior to calling vm.destroy.

Examples:
  govc vm.destroy my-vm

Options:
```

## vm.disk.attach

```
Usage: govc vm.disk.attach [OPTIONS]

Attach existing disk to VM.

A delta disk is created by default, where changes are persisted. Specifying '-link=false' will persist to the same disk.

Examples:
  govc vm.disk.attach -vm $name -disk $name/disk1.vmdk
  govc device.info -vm $name disk-* # 'File' field is where changes are persisted. 'Parent' field is set when '-link=true'
  govc vm.disk.attach -vm $name -disk $name/shared.vmdk -link=false -sharing sharingMultiWriter
  govc device.remove -vm $name -keep disk-* # detach disk(s)

Options:
  -controller=           Disk controller
  -disk=                 Disk path name
  -ds=                   Datastore [GOVC_DATASTORE]
  -link=true             Link specified disk
  -mode=                 Disk mode (persistent|nonpersistent|undoable|independent_persistent|independent_nonpersistent|append)
  -persist=true          Persist attached disk
  -sharing=              Sharing (sharingNone|sharingMultiWriter)
  -vm=                   Virtual machine [GOVC_VM]
```

## vm.disk.change

```
Usage: govc vm.disk.change [OPTIONS]

Change some properties of a VM's DISK

In particular, you can change the DISK mode, and the size (as long as it is bigger)

Examples:
  govc vm.disk.change -vm VM -disk.key 2001 -size 10G
  govc vm.disk.change -vm VM -disk.label "BDD disk" -size 10G
  govc vm.disk.change -vm VM -disk.name "hard-1000-0" -size 12G
  govc vm.disk.change -vm VM -disk.filePath "[DS] VM/VM-1.vmdk" -mode nonpersistent

Options:
  -disk.filePath=        Disk file name
  -disk.key=0            Disk unique key
  -disk.label=           Disk label
  -disk.name=            Disk name
  -mode=                 Disk mode (persistent|nonpersistent|undoable|independent_persistent|independent_nonpersistent|append)
  -sharing=              Sharing (sharingNone|sharingMultiWriter)
  -size=0B               New disk size
  -vm=                   Virtual machine [GOVC_VM]
```

## vm.disk.create

```
Usage: govc vm.disk.create [OPTIONS]

Create disk and attach to VM.

Examples:
  govc vm.disk.create -vm $name -name $name/disk1 -size 10G
  govc vm.disk.create -vm $name -name $name/disk2 -size 10G -eager -thick -sharing sharingMultiWriter

Options:
  -controller=           Disk controller
  -ds=                   Datastore [GOVC_DATASTORE]
  -eager=false           Eagerly scrub new disk
  -mode=persistent       Disk mode (persistent|nonpersistent|undoable|independent_persistent|independent_nonpersistent|append)
  -name=                 Name for new disk
  -sharing=              Sharing (sharingNone|sharingMultiWriter)
  -size=10.0GB           Size of new disk
  -thick=false           Thick provision new disk
  -vm=                   Virtual machine [GOVC_VM]
```

## vm.guest.tools

```
Usage: govc vm.guest.tools [OPTIONS] VM...

Manage guest tools in VM.

Examples:
  govc vm.guest.tools -mount VM
  govc vm.guest.tools -unmount VM
  govc vm.guest.tools -upgrade -options "opt1 opt2" VM

Options:
  -mount=false           Mount tools CD installer in the guest
  -options=              Installer options
  -unmount=false         Unmount tools CD installer in the guest
  -upgrade=false         Upgrade tools in the guest
```

## vm.info

```
Usage: govc vm.info [OPTIONS] VM...

Display info for VM.

The '-r' flag displays additional info for CPU, memory and storage usage,
along with the VM's Datastores, Networks and PortGroups.

Examples:
  govc vm.info $vm
  govc vm.info -r $vm | grep Network:
  govc vm.info -json $vm
  govc find . -type m -runtime.powerState poweredOn | xargs govc vm.info

Options:
  -e=false               Show ExtraConfig
  -g=true                Show general summary
  -r=false               Show resource summary
  -t=false               Show ToolsConfigInfo
  -waitip=false          Wait for VM to acquire IP address
```

## vm.instantclone

```
Usage: govc vm.instantclone [OPTIONS] NAME

Instant Clone VM to NAME.

Examples:
  govc vm.instantclone -vm source-vm new-vm
  # Configure ExtraConfig variables on a guest VM:
  govc vm.instantclone -vm source-vm -e guestinfo.ipaddress=192.168.0.1 -e guestinfo.netmask=255.255.255.0 new-vm
  # Read the variable set above inside the guest:
  vmware-rpctool "info-get guestinfo.ipaddress"
  vmware-rpctool "info-get guestinfo.netmask"

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
  -e=[]                  ExtraConfig. <key>=<value>
  -folder=               Inventory folder [GOVC_FOLDER]
  -net=                  Network [GOVC_NETWORK]
  -net.adapter=e1000     Network adapter type
  -net.address=          Network hardware address
  -pool=                 Resource pool [GOVC_RESOURCE_POOL]
  -vm=                   Virtual machine [GOVC_VM]
```

## vm.ip

```
Usage: govc vm.ip [OPTIONS] VM...

List IPs for VM.

By default the vm.ip command depends on vmware-tools to report the 'guest.ipAddress' field and will
wait until it has done so.  This value can also be obtained using:

  govc vm.info -json $vm | jq -r .VirtualMachines[].Guest.IpAddress

When given the '-a' flag, only IP addresses for which there is a corresponding virtual nic are listed.
If there are multiple nics, the listed addresses will be comma delimited.  The '-a' flag depends on
vmware-tools to report the 'guest.net' field and will wait until it has done so for all nics.
Note that this list includes IPv6 addresses if any, use '-v4' to filter them out.  IP addresses reported
by tools for which there is no virtual nic are not included, for example that of the 'docker0' interface.

These values can also be obtained using:

  govc vm.info -json $vm | jq -r .VirtualMachines[].Guest.Net[].IpConfig.IpAddress[].IpAddress

When given the '-n' flag, filters '-a' behavior to the nic specified by MAC address or device name.

The 'esxcli' flag does not require vmware-tools to be installed, but does require the ESX host to
have the /Net/GuestIPHack setting enabled.

The 'wait' flag default to 1hr (original default was infinite).  If a VM does not obtain an IP within
the wait time, the command will still exit with status 0.

Examples:
  govc vm.ip $vm
  govc vm.ip -wait 5m $vm
  govc vm.ip -a -v4 $vm
  govc vm.ip -n 00:0c:29:57:7b:c3 $vm
  govc vm.ip -n ethernet-0 $vm
  govc host.esxcli system settings advanced set -o /Net/GuestIPHack -i 1
  govc vm.ip -esxcli $vm

Options:
  -a=false               Wait for an IP address on all NICs
  -esxcli=false          Use esxcli instead of guest tools
  -n=                    Wait for IP address on NIC, specified by device name or MAC
  -v4=false              Only report IPv4 addresses
  -wait=1h0m0s           Wait time for the VM obtain an IP address
```

## vm.keystrokes

```
Usage: govc vm.keystrokes [OPTIONS] VM

Send Keystrokes to VM.

Examples:
 Default Scenario
  govc vm.keystrokes -vm $vm -s "root" 	# writes 'root' to the console
  govc vm.keystrokes -vm $vm -c 0x15 	# writes an 'r' to the console
  govc vm.keystrokes -vm $vm -r 1376263 # writes an 'r' to the console
  govc vm.keystrokes -vm $vm -c 0x28 	# presses ENTER on the console
  govc vm.keystrokes -vm $vm -c 0x4c -la=true -lc=true 	# sends CTRL+ALT+DEL to console
  govc vm.keystrokes -vm $vm -c 0x15,KEY_ENTER # writes an 'r' to the console and press ENTER

List of available aliases:
KEY_0, KEY_1, KEY_102ND, KEY_2, KEY_3, KEY_4, KEY_5, KEY_6, KEY_7, KEY_8, KEY_9, KEY_A, KEY_AGAIN, KEY_APOSTROPHE, KEY_B, KEY_BACKSLASH, KEY_BACKSPACE, KEY_C, KEY_CAPSLOCK, KEY_COMMA, KEY_COMPOSE, KEY_COPY, KEY_CUT, KEY_D, KEY_DELETE, KEY_DOT, KEY_DOWN, KEY_E, KEY_END, KEY_ENTER, KEY_EQUAL, KEY_ERR_OVF, KEY_ESC, KEY_F, KEY_F1, KEY_F10, KEY_F11, KEY_F12, KEY_F13, KEY_F14, KEY_F15, KEY_F16, KEY_F17, KEY_F18, KEY_F19, KEY_F2, KEY_F20, KEY_F21, KEY_F22, KEY_F23, KEY_F24, KEY_F3, KEY_F4, KEY_F5, KEY_F6, KEY_F7, KEY_F8, KEY_F9, KEY_FIND, KEY_FRONT, KEY_G, KEY_GRAVE, KEY_H, KEY_HANGEUL, KEY_HANJA, KEY_HASHTILDE, KEY_HELP, KEY_HENKAN, KEY_HIRAGANA, KEY_HOME, KEY_I, KEY_INSERT, KEY_J, KEY_K, KEY_KATAKANA, KEY_KATAKANAHIRAGANA, KEY_KP0, KEY_KP1, KEY_KP2, KEY_KP3, KEY_KP4, KEY_KP5, KEY_KP6, KEY_KP7, KEY_KP8, KEY_KP9, KEY_KPASTERISK, KEY_KPCOMMA, KEY_KPDOT, KEY_KPENTER, KEY_KPEQUAL, KEY_KPJPCOMMA, KEY_KPLEFTPAREN, KEY_KPMINUS, KEY_KPPLUS, KEY_KPRIGHTPAREN, KEY_KPSLASH, KEY_L, KEY_LEFT, KEY_LEFTALT, KEY_LEFTBRACE, KEY_LEFTCTRL, KEY_LEFTMETA, KEY_LEFTSHIFT, KEY_M, KEY_MEDIA_BACK, KEY_MEDIA_CALC, KEY_MEDIA_COFFEE, KEY_MEDIA_EDIT, KEY_MEDIA_EJECTCD, KEY_MEDIA_FIND, KEY_MEDIA_FORWARD, KEY_MEDIA_MUTE, KEY_MEDIA_NEXTSONG, KEY_MEDIA_PLAYPAUSE, KEY_MEDIA_PREVIOUSSONG, KEY_MEDIA_REFRESH, KEY_MEDIA_SCROLLDOWN, KEY_MEDIA_SCROLLUP, KEY_MEDIA_SLEEP, KEY_MEDIA_STOP, KEY_MEDIA_STOPCD, KEY_MEDIA_VOLUMEDOWN, KEY_MEDIA_VOLUMEUP, KEY_MEDIA_WWW, KEY_MINUS, KEY_MOD_LALT, KEY_MOD_LCTRL, KEY_MOD_LMETA, KEY_MOD_LSHIFT, KEY_MOD_RALT, KEY_MOD_RCTRL, KEY_MOD_RMETA, KEY_MOD_RSHIFT, KEY_MUHENKAN, KEY_MUTE, KEY_N, KEY_NONE, KEY_NUMLOCK, KEY_O, KEY_OPEN, KEY_P, KEY_PAGEDOWN, KEY_PAGEUP, KEY_PASTE, KEY_PAUSE, KEY_POWER, KEY_PROPS, KEY_Q, KEY_R, KEY_RIGHT, KEY_RIGHTALT, KEY_RIGHTBRACE, KEY_RIGHTCTRL, KEY_RIGHTMETA, KEY_RIGHTSHIFT, KEY_RO, KEY_S, KEY_SCROLLLOCK, KEY_SEMICOLON, KEY_SLASH, KEY_SPACE, KEY_STOP, KEY_SYSRQ, KEY_T, KEY_TAB, KEY_U, KEY_UNDO, KEY_UP, KEY_V, KEY_VOLUMEDOWN, KEY_VOLUMEUP, KEY_W, KEY_X, KEY_Y, KEY_YEN, KEY_Z, KEY_ZENKAKUHANKAKU


Options:
  -c=                    USB HID Codes (hex) or aliases, comma separated
  -la=false              Enable/Disable Left Alt
  -lc=false              Enable/Disable Left Control
  -lg=false              Enable/Disable Left Gui
  -ls=false              Enable/Disable Left Shift
  -r=0                   Raw USB HID Code Value (int32)
  -ra=false              Enable/Disable Right Alt
  -rc=false              Enable/Disable Right Control
  -rg=false              Enable/Disable Right Gui
  -rs=false              Enable/Disable Right Shift
  -s=                    Raw String to Send
  -vm=                   Virtual machine [GOVC_VM]
```

## vm.markastemplate

```
Usage: govc vm.markastemplate [OPTIONS] VM...

Mark VM as a virtual machine template.

Examples:
  govc vm.markastemplate $name

Options:
```

## vm.markasvm

```
Usage: govc vm.markasvm [OPTIONS] VM...

Mark VM template as a virtual machine.

Examples:
  govc vm.markasvm -host host1 $name
  govc vm.markasvm -host host1 $name

Options:
  -host=                 Host system [GOVC_HOST]
  -pool=                 Resource pool [GOVC_RESOURCE_POOL]
```

## vm.migrate

```
Usage: govc vm.migrate [OPTIONS] VM...

Migrates VM to a specific resource pool, host or datastore.

Examples:
  govc vm.migrate -host another-host vm-1 vm-2 vm-3
  govc vm.migrate -pool another-pool vm-1 vm-2 vm-3
  govc vm.migrate -ds another-ds vm-1 vm-2 vm-3

Options:
  -ds=                       Datastore [GOVC_DATASTORE]
  -folder=                   Inventory folder [GOVC_FOLDER]
  -host=                     Host system [GOVC_HOST]
  -pool=                     Resource pool [GOVC_RESOURCE_POOL]
  -priority=defaultPriority  The task priority
```

## vm.network.add

```
Usage: govc vm.network.add [OPTIONS]

Add network adapter to VM.

Examples:
  govc vm.network.add -vm $vm -net "VM Network" -net.adapter e1000e
  govc vm.network.add -vm $vm -net SwitchName/PortgroupName
  govc device.info -vm $vm ethernet-*

Options:
  -net=                  Network [GOVC_NETWORK]
  -net.adapter=e1000     Network adapter type
  -net.address=          Network hardware address
  -vm=                   Virtual machine [GOVC_VM]
```

## vm.network.change

```
Usage: govc vm.network.change [OPTIONS] DEVICE

Change network DEVICE configuration.

Note that '-net' is currently required with '-net.address', even when not changing the VM network.

Examples:
  govc vm.network.change -vm $vm -net PG2 ethernet-0
  govc vm.network.change -vm $vm -net PG2 -net.address 00:00:0f:2e:5d:69 ethernet-0 # set to manual MAC address
  govc vm.network.change -vm $vm -net PG2 -net.address - ethernet-0 # set to generated MAC address
  govc device.info -vm $vm ethernet-*

Options:
  -net=                  Network [GOVC_NETWORK]
  -net.adapter=e1000     Network adapter type
  -net.address=          Network hardware address
  -vm=                   Virtual machine [GOVC_VM]
```

## vm.option.info

```
Usage: govc vm.option.info [OPTIONS] [GUEST_ID]...

VM config options for CLUSTER.

The config option data contains information about the execution environment for a VM
in the given CLUSTER, and optionally for a specific HOST.

By default, supported guest OS IDs and full name are listed.

Examples:
  govc vm.option.info -cluster C0
  govc vm.option.info -cluster C0 -dump ubuntu64Guest
  govc vm.option.info -cluster C0 -json | jq .GuestOSDescriptor[].Id
  govc vm.option.info -host my_hostname
  govc vm.option.info -vm my_vm

Options:
  -cluster=              Cluster [GOVC_CLUSTER]
  -host=                 Host system [GOVC_HOST]
  -vm=                   Virtual machine [GOVC_VM]
```

## vm.power

```
Usage: govc vm.power [OPTIONS] NAME...

Invoke VM power operations.

Examples:
  govc vm.power -on VM1 VM2 VM3
  govc vm.power -on -M VM1 VM2 VM3
  govc vm.power -off -force VM1

Options:
  -M=false               Use Datacenter.PowerOnMultiVM method instead of VirtualMachine.PowerOnVM
  -force=false           Force (ignore state error and hard shutdown/reboot if tools unavailable)
  -off=false             Power off
  -on=false              Power on
  -r=false               Reboot guest
  -reset=false           Power reset
  -s=false               Shutdown guest
  -suspend=false         Power suspend
  -wait=true             Wait for the operation to complete
```

## vm.question

```
Usage: govc vm.question [OPTIONS]

Options:
  -answer=               Answer to question
  -vm=                   Virtual machine [GOVC_VM]
```

## vm.rdm.attach

```
Usage: govc vm.rdm.attach [OPTIONS]

Attach DEVICE to VM with RDM.

Examples:
  govc vm.rdm.attach -vm VM -device /vmfs/devices/disks/naa.000000000000000000000000000000000

Options:
  -device=               Device Name
  -vm=                   Virtual machine [GOVC_VM]
```

## vm.rdm.ls

```
Usage: govc vm.rdm.ls [OPTIONS]

List available devices that could be attach to VM with RDM.

Examples:
  govc vm.rdm.ls -vm VM

Options:
  -vm=                   Virtual machine [GOVC_VM]
```

## vm.register

```
Usage: govc vm.register [OPTIONS] VMX

Add an existing VM to the inventory.

VMX is a path to the vm config file, relative to DATASTORE.

Examples:
  govc vm.register path/name.vmx
  govc vm.register -template -host $host path/name.vmx

Options:
  -ds=                   Datastore [GOVC_DATASTORE]
  -folder=               Inventory folder [GOVC_FOLDER]
  -host=                 Host system [GOVC_HOST]
  -name=                 Name of the VM
  -pool=                 Resource pool [GOVC_RESOURCE_POOL]
  -template=false        Mark VM as template
```

## vm.unregister

```
Usage: govc vm.unregister [OPTIONS] VM...

Remove VM from inventory without removing any of the VM files on disk.

Options:
```

## vm.upgrade

```
Usage: govc vm.upgrade [OPTIONS]

Upgrade VMs to latest hardware version

Examples:
  govc vm.upgrade -vm $vm_name
  govc vm.upgrade -version=$version -vm $vm_name
  govc vm.upgrade -version=$version -vm.uuid $vm_uuid

Options:
  -version=0             Target vm hardware version, by default -- latest available
  -vm=                   Virtual machine [GOVC_VM]
```

## vm.vnc

```
Usage: govc vm.vnc [OPTIONS] VM...

Enable or disable VNC for VM.

Port numbers are automatically chosen if not specified.

If neither -enable or -disable is specified, the current state is returned.

Examples:
  govc vm.vnc -enable -password 1234 $vm | awk '{print $2}' | xargs open

Options:
  -disable=false         Disable VNC
  -enable=false          Enable VNC
  -password=             VNC password
  -port=-1               VNC port (-1 for auto-select)
  -port-range=5900-5999  VNC port auto-select range
```

## volume.ls

```
Usage: govc volume.ls [OPTIONS] [ID...]

List CNS volumes.

Examples:
  govc volume.ls
  govc volume.ls -l
  govc volume.ls -ds vsanDatastore
  govc volume.ls df86393b-5ae0-4fca-87d0-b692dbc67d45
  govc disk.ls -l $(govc volume.ls -L pvc-9744a4ff-07f4-43c4-b8ed-48ea7a528734)

Options:
  -L=false               List volume disk or file backing ID only
  -ds=                   Datastore [GOVC_DATASTORE]
  -i=false               List volume ID only
  -l=false               Long listing format
```

## volume.rm

```
Usage: govc volume.rm [OPTIONS] ID

Remove CNS volume.

Note: if volume.rm returns not found errors,
consider using 'govc disk.ls -R' to reconcile the datastore inventory.

Examples:
  govc volume.rm f75989dc-95b9-4db7-af96-8583f24bc59d

Options:
```

## vsan.change

```
Usage: govc vsan.change [OPTIONS] CLUSTER

Change vSAN configuration.

Examples:
  govc vsan.change -unmap-enabled ClusterA # enable unmap
  govc vsan.change -unmap-enabled=false ClusterA # disable unmap

Options:
  -unmap-enabled=<nil>   Enable Unmap
```

## vsan.info

```
Usage: govc vsan.info [OPTIONS] CLUSTER...

Display vSAN configuration.

Examples:
  govc vsan.info
  govc vsan.info ClusterA
  govc vsan.info -json

Options:
```

