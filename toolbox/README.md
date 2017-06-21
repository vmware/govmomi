# toolbox - VMware guest tools library for Go #

## Overview

The toolbox library is a lightweight, extensible framework for implementing VMware guest tools functionality.
The primary focus of the library is the implementation of VM guest RPC protocols, transport and dispatch.
These protocols are undocumented for the most part, but [open-vm-tools](https://github.com/vmware/open-vm-tools) serves
as a reference implementation.  The toolbox provides default implementations of the supported RPCs, which can be
overridden and/or extended by consumers.

## Supported features

Feature list from the perspective of vSphere public API interaction.  The properties, objects and methods listed are
relative to
the [VirtualMachine](http://pubs.vmware.com/vsphere-60/index.jsp#com.vmware.wssdk.apiref.doc/vim.VirtualMachine.html)
managed object type.

### guest.toolsVersionStatus property

The toolbox reports version as `guestToolsUnmanaged`.

See [ToolsVersionStatus](http://pubs.vmware.com/vsphere-60/index.jsp#com.vmware.wssdk.apiref.doc/vim.vm.GuestInfo.ToolsVersionStatus.html)

### guest.toolsRunningStatus and guest.guestState properties

The VMX determines these values based on the toolbox's response to the `ping` RPC.

### guest.ipAddress property

The VMX requests this value via the `Set_Option broadcastIP` RPC.

The default value can be overridden by setting the `Service.PrimaryIP` function.

See [vim.vm.GuestInfo](http://pubs.vmware.com/vsphere-60/index.jsp#com.vmware.wssdk.apiref.doc/vim.vm.GuestInfo.html)

### guest.net property

This data is pushed to the VMX using the `SendGuestInfo(INFO_IPADDRESS_V3)` RPC.

See [GuestNicInfo](http://pubs.vmware.com/vsphere-60/index.jsp?topic=%2Fcom.vmware.wssdk.apiref.doc%2Fvim.vm.GuestInfo.NicInfo.html).

### ShutdownGuest and RebootGuest methods

The [PowerCommandHandler](power.go) provides power hooks for customized guest shutdown and reboot.

### GuestAuthManager object

Guest operations can be authenticated using the `VixRelayedCommandHandler.Authenticate` hook.

See [vim.vm.guest.AuthManager](http://pubs.vmware.com/vsphere-60/index.jsp#com.vmware.wssdk.apiref.doc/vim.vm.guest.AuthManager.html)

### GuestFileManager object

| Method                          | Supported | Client Examples                                                                     |
|---------------------------------|-----------|-------------------------------------------------------------------------------------|
| ChangeFileAttributesInGuest     | Yes       | [chmod](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/chmod.go)       |
|                                 |           | [chown](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/chown.go)       |
|                                 |           | [touch](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/touch.go)       |
| CreateTemporaryDirectoryInGuest | Yes       | [mktemp](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/mktemp.go)  |
| CreateTemporaryFileInGuest      | Yes       | [mktemp](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/mktemp.go)     |
| DeleteDirectoryInGuest          | Yes       | [rmdir](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/rmdir.go)       |
| DeleteFileInGuest               | Yes       | [rm](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/rm.go)             |
| InitiateFileTransferFromGuest   | Yes       | [download](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/download.go) |
| InitiateFileTransferToGuest     | Yes       | [upload](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/upload.go)     |
| ListFilesInGuest                | Yes       | [ls](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/ls.go)             |
| MakeDirectoryInGuest            | Yes       | [mkdir](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/mkdir.go)       |
| MoveDirectoryInGuest            | Yes       | [mv](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/mv.go)             |
| MoveFileInGuest                 | Yes       | [mv](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/mv.go)             |

See [vim.vm.guest.FileManager](http://pubs.vmware.com/vsphere-60/index.jsp#com.vmware.wssdk.apiref.doc/vim.vm.guest.FileManager.html)

The toolbox provides support for transferring directories to and from guests as gzip'd tar streams, without writing the
tar file itself to the guest file system.  To disable this feature, set the `hgfs.FileServer.Archive` field to `false`.

### GuestProcessManager

The toolbox [ProcessManager](process.go) can manage both OS processes and Go functions running as goroutines.
Currently, the `ListProcessesInGuest` and `TerminateProcessInGuest` methods only apply those processes and goroutines
started by `StartProgramInGuest`.

| Method                         | Supported | Client Examples                                                                     |
|--------------------------------|-----------|-------------------------------------------------------------------------------------|
| ListProcessesInGuest           | Yes       | [ps](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/ps.go)             |
| ReadEnvironmentVariableInGuest | Yes       | [getenv](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/getenv.go)     |
| StartProgramInGuest            | Yes       | [start](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/start.go)       |
| TerminateProcessInGuest        | Yes       | [kill](https://github.com/vmware/govmomi/blob/master/govc/vm/guest/kill.go)         |

See [vim.vm.guest.ProcessManager](http://pubs.vmware.com/vsphere-60/index.jsp#com.vmware.wssdk.apiref.doc/vim.vm.guest.ProcessManager.html)

## Testing

The Go tests cover most of the toolbox code and can be run on any Linux or MacOSX machine, virtual or otherwise.

To test the toolbox with vSphere API interaction, it must be run inside a VM managed by vSphere without the standard
vmtoolsd running.

The [toolbox-test.sh](toolbox-test.sh) can be used to run the full suite of toolbox tests with vSphere API interaction.
Use the `-s` flag to start the standalone version of the toolbox and leave it running, to test vSphere interaction
without running the test suite.

## Consumers of the toolbox library

* [Toolbox example main](https://github.com/vmware/govmomi/blob/master/toolbox/toolbox/main.go
)

* [VIC tether toolbox extension](https://github.com/vmware/vic/blob/master/lib/tether/toolbox.go)

* [VIC container VM tether](https://github.com/vmware/vic/blob/main/cmd/tether/main_linux.go)

* [VIC container host tether](https://github.com/vmware/vic/blob/master/cmd/vic-init/main_linux.go)

## Supported guests

The toolbox guest RPC implementations tend to be simple and portable thanks to the Go standard library, but are only
supported on Linux currently.  Support for other guests, such as Windows, has been kept in mind but not yet tested.

## Supported vSphere Versions

The toolbox is supported with vSphere 6.0 and 6.5, but may function with older versions.
