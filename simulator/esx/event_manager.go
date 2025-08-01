// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package esx

import "github.com/vmware/govmomi/vim25/types"

// EventInfo is the default template for the EventManager description.eventInfo property.
// Capture method:
// govc object.collect -s -dump EventManager:ha-eventmgr description.eventInfo
// The captured list has been manually pruned and FullFormat fields changed to use Go's template variable syntax.
var EventInfo = []types.EventDescriptionEventDetail{
	{
		Key:         "UserLoginSessionEvent",
		Description: "User login",
		Category:    "info",
		FullFormat:  "User {{.UserName}}@{{.IpAddress}} logged in as {{.UserAgent}}",
	},
	{
		Key:         "UserLogoutSessionEvent",
		Description: "User logout",
		Category:    "info",
		FullFormat:  "User {{.UserName}}@{{.IpAddress}} logged out (login time: {{.LoginTime}}, number of API invocations: {{.CallCount}}, user agent: {{.UserAgent}})",
	},
	{
		Key:         "DatacenterCreatedEvent",
		Description: "Datacenter created",
		Category:    "info",
		FullFormat:  "Created datacenter {{.Datacenter.Name}} in folder {{.Parent.Name}}",
	},
	{
		Key:         "DatastoreFileMovedEvent",
		Description: "File or directory moved to datastore",
		Category:    "info",
		FullFormat:  "Move of file or directory {{.SourceFile}} from {{.SourceDatastore.Name}} to {{.Datastore.Name}} as {{.TargetFile}}",
	},
	{
		Key:         "DatastoreFileCopiedEvent",
		Description: "File or directory copied to datastore",
		Category:    "info",
		FullFormat:  "Copy of file or directory {{.SourceFile}} from {{.SourceDatastore.Name}} to {{.Datastore.Name}} as {{.TargetFile}}",
	},
	{
		Key:         "DatastoreFileDeletedEvent",
		Description: "File or directory deleted",
		Category:    "info",
		FullFormat:  "Deletion of file or directory {{.TargetFile}} from {{.Datastore.Name}} was initiated",
	},
	{
		Key:         "EnteringMaintenanceModeEvent",
		Description: "Entering maintenance mode",
		Category:    "info",
		FullFormat:  "Host {{.Host.Name}} in {{.Datacenter.Name}} has started to enter maintenance mode",
	},
	{
		Key:         "EnteredMaintenanceModeEvent",
		Description: "Entered maintenance mode",
		Category:    "info",
		FullFormat:  "Host {{.Host.Name}} in {{.Datacenter.Name}} has entered maintenance mode",
	},
	{
		Key:         "ExitMaintenanceModeEvent",
		Description: "Exit maintenance mode",
		Category:    "info",
		FullFormat:  "Host {{.Host.Name}} in {{.Datacenter.Name}} has exited maintenance mode",
	},
	{
		Key:         "HostRemovedEvent",
		Description: "Host removed",
		FullFormat:  "Removed host {{.Host.Name}} in {{.Datacenter.Name}}",
		Category:    "info",
	},
	{
		Key:         "VmSuspendedEvent",
		Description: "VM suspended",
		Category:    "info",
		FullFormat:  "{{.Vm.Name}} on {{.Host.Name}} in {{.Datacenter.Name}} is suspended",
	},
	{
		Key:         "VmMigratedEvent",
		Description: "VM migrated",
		Category:    "info",
		FullFormat:  "Migration of virtual machine {{.Vm.Name}} from {{.SourceHost.Name}}, {{.SourceDatastore.Name}} to {{.Host.Name}}, {{.Ds.Name}} completed",
	},
	{
		Key:         "VmBeingMigratedEvent",
		Description: "VM migrating",
		Category:    "info",
		FullFormat:  "Relocating {{.Vm.Name}} from {{.Host.Name}, {{.Ds.Name}} in {{.Datacenter.Name}} to {{.DestHost.Name}, {{.DestDatastore.Name}} in {{.DestDatacenter.Name}}",
	},
	{
		Key:         "VmMacAssignedEvent",
		Description: "VM MAC assigned",
		Category:    "info",
		FullFormat:  "New MAC address ({{.Mac}}) assigned to adapter {{.Adapter}} for {{.Vm.Name}}",
	},
	{
		Key:         "VmRegisteredEvent",
		Description: "VM registered",
		Category:    "info",
		FullFormat:  "Registered {{.Vm.Name}} on {{.Host.Name}} in {{.Datacenter.Name}}",
	},
	{
		Key:         "VmReconfiguredEvent",
		Description: "VM reconfigured",
		Category:    "info",
		FullFormat:  "Reconfigured {{.Vm.Name}} on {{.Host.Name}} in {{.Datacenter.Name}}",
	},
	{
		Key:         "VmGuestRebootEvent",
		Description: "Guest reboot",
		Category:    "info",
		FullFormat:  "Guest OS reboot for {{.Vm.Name}} on {{.Host.Name}} in {{.Datacenter.Name}}",
	},
	{
		Key:         "VmBeingClonedEvent",
		Description: "VM being cloned",
		Category:    "info",
		FullFormat:  "Cloning {{.Vm.Name}} on host {{.Host.Name}} in {{.Datacenter.Name}} to {{.DestName}} on host {{.DestHost.Name}}",
	},
	{
		Key:         "VmClonedEvent",
		Description: "VM cloned",
		Category:    "info",
		FullFormat:  "Clone of {{.SourceVm.Name}} completed",
	},
	{
		Key:         "VmBeingDeployedEvent",
		Description: "Deploying VM",
		Category:    "info",
		FullFormat:  "Deploying {{.Vm.Name}} on host {{.Host.Name}} in {{.Datacenter.Name}} from template {{.SrcTemplate.Name}}",
	},
	{
		Key:         "VmDeployedEvent",
		Description: "VM deployed",
		Category:    "info",
		FullFormat:  "Template {{.SrcTemplate.Name}} deployed on host {{.Host.Name}}",
	},
	{
		Key:         "VmInstanceUuidAssignedEvent",
		Description: "Assign a new instance UUID",
		Category:    "info",
		FullFormat:  "Assign a new instance UUID ({{.InstanceUuid}}) to {{.Vm.Name}}",
	},
	{
		Key:         "VmPoweredOnEvent",
		Description: "VM powered on",
		Category:    "info",
		FullFormat:  "{{.Vm.Name}} on {{.Host.Name}} in {{.Datacenter.Name}} is powered on",
	},
	{
		Key:         "VmStartingEvent",
		Description: "VM starting",
		Category:    "info",
		FullFormat:  "{{.Vm.Name}} on host {{.Host.Name}} in {{.Datacenter.Name}} is starting",
	},
	{
		Key:         "VmStoppingEvent",
		Description: "VM stopping",
		Category:    "info",
		FullFormat:  "{{.Vm.Name}} on host {{.Host.Name}} in {{.Datacenter.Name}} is stopping",
	},
	{
		Key:         "VmSuspendingEvent",
		Description: "VM being suspended",
		Category:    "info",
		FullFormat:  "{{.Vm.Name}} on {{.Host.Name}} in {{.Datacenter.Name}} is being suspended",
	},
	{
		Key:         "VmResumingEvent",
		Description: "VM resuming",
		Category:    "info",
		FullFormat:  "{{.Vm.Name}} on {{.Host.Name}} in {{.Datacenter.Name}} is resumed",
	},
	{
		Key:         "VmBeingCreatedEvent",
		Description: "Creating VM",
		Category:    "info",
		FullFormat:  "Creating {{.Vm.Name}} on host {{.Host.Name}} in {{.Datacenter.Name}}",
	},
	{
		Key:         "VmCreatedEvent",
		Description: "VM created",
		Category:    "info",
		FullFormat:  "Created virtual machine {{.Vm.Name}} on {{.Host.Name}} in {{.Datacenter.Name}}",
	},
	{
		Key:         "VmRemovedEvent",
		Description: "VM removed",
		Category:    "info",
		FullFormat:  "Removed {{.Vm.Name}} on {{.Host.Name}} from {{.Datacenter.Name}}",
	},
	{
		Key:         "VmResettingEvent",
		Description: "VM resetting",
		Category:    "info",
		FullFormat:  "{{.Vm.Name}} on {{.Host.Name}} in {{.Datacenter.Name}} is reset",
	},
	{
		Key:         "VmGuestShutdownEvent",
		Description: "Guest OS shut down",
		Category:    "info",
		FullFormat:  "Guest OS shut down for {{.Vm.Name}} on {{.Host.Name}} in {{.Datacenter.Name}}",
	},
	{
		Key:         "VmUuidAssignedEvent",
		Description: "VM UUID assigned",
		Category:    "info",
		FullFormat:  "Assigned new BIOS UUID ({{.Uuid}}) to {{.Vm.Name}} on {{.Host.Name}} in {{.Datacenter.Name}}",
	},
	{
		Key:         "VmPoweredOffEvent",
		Description: "VM powered off",
		Category:    "info",
		FullFormat:  "{{.Vm.Name}} on {{.Host.Name}} in {{.Datacenter.Name}} is powered off",
	},
	{
		Key:         "VmRelocatedEvent",
		Description: "VM relocated",
		Category:    "info",
		FullFormat:  "Completed the relocation of the virtual machine",
	},
	{
		Key:         "CustomizationFailed",
		Description: "An error occurred during customization",
		Category:    "info",
		FullFormat:  "An error occurred during customization on VM {{.Vm.Name}}",
	},
	{
		Key:         "CustomizationStartedEvent",
		Description: "Started customization",
		Category:    "info",
		FullFormat:  "Started customization of VM {{.Vm.Name}}",
	},
	{
		Key:         "CustomizationSucceeded",
		Description: "Customization succeeded",
		Category:    "info",
		FullFormat:  "Customization of VM {{.Vm.Name}} succeeded",
	},
	{
		Key:         "CustomizationNetworkSetupFailed",
		Description: "Cannot complete customization network setup",
		Category:    "error",
		FullFormat:  "An error occurred while setting up network properties of the guest OS",
	},
	{
		Key:         "DrsVmMigratedEvent",
		Description: "DRS VM migrated",
		Category:    "info",
		FullFormat:  "DRS migrated {{.Vm.Name}} from {{.SourceHost.Name}} to {{.Host.Name}} in cluster {{.ComputeResource.Name}} in {{.Datacenter.Name}}",
	},
	{
		Key:         "DrsVmPoweredOnEvent",
		Description: "DRS VM powered on",
		Category:    "info",
		FullFormat:  "DRS powered On {{.Vm.Name}} on {{.Host.Name}} in {{.Datacenter.Name}}",
	},
	{
		Key:         "DvsCreatedEvent",
		Description: "vSphere Distributed Switch created",
		Category:    "info",
		FullFormat:  "A vSphere Distributed Switch {{.Dvs.Name}} was created in {{.Datacenter.Name}}.",
	},
	{
		Key:         "DvsDestroyedEvent",
		Description: "vSphere Distributed Switch deleted",
		Category:    "info",
		FullFormat:  "vSphere Distributed Switch {{.Dvs.Name}} in {{.Datacenter.Name}} was deleted.",
	},
	{
		Key:         "DvsReconfiguredEvent",
		Description: "vSphere Distributed Switch reconfigured",
		Category:    "info",
		FullFormat:  "The vSphere Distributed Switch {{.Dvs.Name}} in {{.Datacenter.Name}} was reconfigured.",
	},
	{
		Key:         "DvsHostJoinedEvent",
		Description: "Host joined the vSphere Distributed Switch",
		Category:    "info",
		FullFormat:  "The host {{.HostJoined.Name}} joined the vSphere Distributed Switch {{.Dvs.Name}} in {{.Datacenter.Name}}.",
	},
	{
		Key:         "DvsHostLeftEvent",
		Description: "Host left vSphere Distributed Switch",
		Category:    "info",
		FullFormat:  "The host {{.HostLeft.Name}} left the vSphere Distributed Switch {{.Dvs.Name}} in {{.Datacenter.Name}}.",
	},
	{
		Key:         "DVPortgroupCreatedEvent",
		Description: "dvPort group created",
		Category:    "info",
		FullFormat:  "dvPort group {{.Net.Name}} in {{.Datacenter.Name}} was added to switch {{.Dvs.Name}}.",
	},
	{
		Key:         "DVPortgroupDestroyedEvent",
		Description: "dvPort group deleted",
		Category:    "info",
		FullFormat:  "dvPort group {{.Net.Name}} in {{.Datacenter.Name}} was deleted.",
	},
}
