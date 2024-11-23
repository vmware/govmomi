/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package esx

import (
	"github.com/vmware/govmomi/cli/esx"
)

// CommandInfo generated via `govc host.esxcli.model -c -dump`
var CommandInfo = []esx.CommandInfo{
	{
		CommandInfoItem: esx.CommandInfoItem{Name: "hardware.clock", DisplayName: "clock", Help: "Interaction with the hardware clock."},
		Method: []esx.CommandInfoMethod{
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "get", DisplayName: "get", Help: "Display the current hardware clock time."},
				Param:           nil,
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: "simple"},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "set", DisplayName: "set", Help: "Set the hardware clock time. Any missing parameters will default to the current time."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "day", DisplayName: "day", Help: "Day"},
						Aliases:         []string{"-d", "--day"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "hour", DisplayName: "hour", Help: "Hour"},
						Aliases:         []string{"-H", "--hour"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "min", DisplayName: "min", Help: "Minute"},
						Aliases:         []string{"-m", "--min"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "month", DisplayName: "month", Help: "Month"},
						Aliases:         []string{"-M", "--month"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "sec", DisplayName: "sec", Help: "Second"},
						Aliases:         []string{"-s", "--sec"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "year", DisplayName: "year", Help: "Year"},
						Aliases:         []string{"-y", "--year"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: ""},
				},
			},
		},
	},
	{
		CommandInfoItem: esx.CommandInfoItem{Name: "hardware.platform", DisplayName: "platform", Help: "Platform information."},
		Method: []esx.CommandInfoMethod{
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "get", DisplayName: "get", Help: "Get information about the platform"},
				Param:           nil,
				Hints: esx.CommandInfoHints{
					{Key: "fields:PlatformGet", Value: "UUID,Product Name,Vendor Name,Serial Number,Enclosure Serial Number,BIOS Asset Tag,IPMI Supported"},
					{Key: "formatter", Value: "simple"},
					{Key: "header:PlatformGet", Value: "Platform Information"},
				},
			},
		},
	},
	{
		CommandInfoItem: esx.CommandInfoItem{Name: "iscsi.software", DisplayName: "software", Help: "Operations that can be performed on software iSCSI"},
		Method: []esx.CommandInfoMethod{
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "get", DisplayName: "get", Help: "Software iSCSI information."},
				Param:           nil,
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: "simple"},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "set", DisplayName: "set", Help: "Enable or disable software iSCSI."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "enabled", DisplayName: "enabled", Help: "Enable or disable the module."},
						Aliases:         []string{"-e", "--enabled"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "name", DisplayName: "name", Help: "The iSCSI initiator name.\nThe initiator name must not be specified when disabling software iSCSI."},
						Aliases:         []string{"-n", "--name"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: "simple"},
				},
			},
		},
	},
	{
		CommandInfoItem: esx.CommandInfoItem{Name: "network.firewall", DisplayName: "firewall", Help: "A set of commands for firewall related operations"},
		Method: []esx.CommandInfoMethod{
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "get", DisplayName: "get", Help: "Get the firewall status."},
				Param:           nil,
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: "simple"},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "load", DisplayName: "load", Help: "Load firewall module and rulesets configuration."},
				Param:           nil,
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: ""},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "refresh", DisplayName: "refresh", Help: "Load ruleset configuration for firewall."},
				Param:           nil,
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: ""},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "set", DisplayName: "set", Help: "Set firewall enabled status and default action."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "defaultaction", DisplayName: "default-action", Help: "Set to true to set defaultaction PASS, set to false to DROP."},
						Aliases:         []string{"-d", "--default-action"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "enabled", DisplayName: "enabled", Help: "Set to true to enable the firewall, set to false to disable the firewall."},
						Aliases:         []string{"-e", "--enabled"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: ""},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "unload", DisplayName: "unload", Help: "Allow unload firewall module."},
				Param:           nil,
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: ""},
				},
			},
		},
	},
	{
		CommandInfoItem: esx.CommandInfoItem{Name: "network.ip.connection", DisplayName: "connection", Help: "List active tcpip connections"},
		Method: []esx.CommandInfoMethod{
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "list", DisplayName: "list", Help: "List active TCP/IP connections"},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "netstack", DisplayName: "netstack", Help: "The network stack instance; if unspecified, use the default netstack instance"},
						Aliases:         []string{"-N", "--netstack"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "type", DisplayName: "type", Help: "Connection type: [ip, tcp, udp, all]"},
						Aliases:         []string{"-t", "--type"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "fields:IpConnection", Value: "Proto,Recv Q,Send Q,Local Address,Foreign Address,State,World ID,CC Algo,World Name"},
					{Key: "formatter", Value: "table"},
				},
			},
		},
	},
	{
		CommandInfoItem: esx.CommandInfoItem{Name: "network.nic.ring.current", DisplayName: "current", Help: "Commands to access current NIC RX/TX ring buffer parameters"},
		Method: []esx.CommandInfoMethod{
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "get", DisplayName: "get", Help: "Get current RX/TX ring buffer parameters of a NIC"},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "nicname", DisplayName: "nic-name", Help: "The name of the NIC whose current RX/TX ring buffer parameters should be retrieved."},
						Aliases:         []string{"-n", "--nic-name"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "fields:RingInfo", Value: "RX,RX Mini,RX Jumbo,TX"},
					{Key: "formatter", Value: "simple"},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "set", DisplayName: "set", Help: "Set current RX/TX ring buffer parameters of a NIC"},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "nicname", DisplayName: "nic-name", Help: "The name of the NIC whose current RX/TX ring buffer parameters should be set."},
						Aliases:         []string{"-n", "--nic-name"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "rx", DisplayName: "rx", Help: "Number of ring entries for the RX ring."},
						Aliases:         []string{"-r", "--rx"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "rxjumbo", DisplayName: "rx-jumbo", Help: "Number of ring entries for the RX jumbo ring."},
						Aliases:         []string{"-j", "--rx-jumbo"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "rxmini", DisplayName: "rx-mini", Help: "Number of ring entries for the RX mini ring."},
						Aliases:         []string{"-m", "--rx-mini"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "tx", DisplayName: "tx", Help: "Number of ring entries for the TX ring."},
						Aliases:         []string{"-t", "--tx"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: ""},
				},
			},
		},
	},
	{
		CommandInfoItem: esx.CommandInfoItem{Name: "network.nic.ring.preset", DisplayName: "preset", Help: "Commands to access preset maximums for NIC RX/TX ring buffer parameters."},
		Method: []esx.CommandInfoMethod{
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "get", DisplayName: "get", Help: "Get preset maximums for RX/TX ring buffer parameters of a NIC."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "nicname", DisplayName: "nic-name", Help: "The name of the NIC whose preset maximums for RX/TX ring buffer parameters should be retrieved."},
						Aliases:         []string{"-n", "--nic-name"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "fields:RingInfo", Value: "Max RX,Max RX Mini,Max RX Jumbo,Max TX"},
					{Key: "formatter", Value: "simple"},
				},
			},
		},
	},
	{
		CommandInfoItem: esx.CommandInfoItem{Name: "network.vm", DisplayName: "vm", Help: "A set of commands for VM related operations"},
		Method: []esx.CommandInfoMethod{
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "list", DisplayName: "list", Help: "List networking information for the VM's that have active ports."},
				Param:           nil,
				Hints: esx.CommandInfoHints{
					{Key: "fields:VM", Value: "World ID,Name,Num Ports,Networks"},
					{Key: "formatter", Value: "table"},
				},
			},
		},
	},
	{
		CommandInfoItem: esx.CommandInfoItem{Name: "software.vib", DisplayName: "vib", Help: "Install, update, remove, or display individual VIB packages"},
		Method: []esx.CommandInfoMethod{
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "get", DisplayName: "get", Help: "Displays detailed information about one or more installed VIBs on the host and the managed DPU(s)."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "rebootingimage", DisplayName: "rebooting-image", Help: "Displays information for the ESXi image which becomes active after a reboot, or nothing if the pending-reboot image has not been created yet. If not specified, information from the current ESXi image in memory will be returned."},
						Aliases:         []string{"--rebooting-image"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "vibname", DisplayName: "vibname", Help: "Specifies one or more installed VIBs to display more information about. If this option is not specified, then all of the installed VIBs will be displayed. Must be one of the following forms: name, name:version, vendor:name, or vendor:name:version."},
						Aliases:         []string{"-n", "--vibname"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "fields:VIBExt", Value: "Name,Version,Type,Vendor,Acceptance Level,Summary,Description,ReferenceURLs,Creation Date,Depends,Conflicts,Replaces,Provides,Maintenance Mode Required,Hardware Platforms Required,Live Install Allowed,Live Remove Allowed,Stateless Ready,Overlay,Tags,Payloads,Platforms"},
					{Key: "formatter", Value: "simple"},
					{Key: "header:VIBExt", Value: "%ID%"},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "install", DisplayName: "install", Help: "Installs VIB packages from a URL or depot. VIBs may be installed, upgraded, or downgraded. WARNING: If your installation requires a reboot, you need to disable HA first."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "depot", DisplayName: "depot", Help: "Specifies full remote URLs of the depot index.xml or server file path pointing to an offline bundle .zip file."},
						Aliases:         []string{"-d", "--depot"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "dryrun", DisplayName: "dry-run", Help: "Performs a dry-run only. Report the VIB-level operations that would be performed, but do not change anything in the system."},
						Aliases:         []string{"--dry-run"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "force", DisplayName: "force", Help: "Bypasses checks for package dependencies, conflicts, obsolescence, and acceptance levels. Really not recommended unless you know what you are doing. Use of this option will result in a warning being displayed in vSphere Web Client.  Use this option only when instructed to do so by VMware Technical Support."},
						Aliases:         []string{"-f", "--force"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "maintenancemode", DisplayName: "maintenance-mode", Help: "Pretends that maintenance mode is in effect. Otherwise, installation will stop for live installs that require maintenance mode. This flag has no effect for reboot required remediations."},
						Aliases:         []string{"--maintenance-mode"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "noliveinstall", DisplayName: "no-live-install", Help: "Forces an install to /altbootbank even if the VIBs are eligible for live installation or removal. Will cause installation to be skipped on PXE-booted hosts."},
						Aliases:         []string{"--no-live-install"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "nosigcheck", DisplayName: "no-sig-check", Help: "Bypasses acceptance level verification, including signing. Use of this option poses a large security risk and will result in a SECURITY ALERT warning being displayed in vSphere Web Client."},
						Aliases:         []string{"--no-sig-check"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "proxy", DisplayName: "proxy", Help: "Specifies a proxy server to use for HTTP, FTP, and HTTPS connections. The format is proxy-url:port."},
						Aliases:         []string{"--proxy"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "vibname", DisplayName: "vibname", Help: "Specifies VIBs from a depot, using one of the following forms: name, name:version, vendor:name, or vendor:name:version."},
						Aliases:         []string{"-n", "--vibname"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "viburl", DisplayName: "viburl", Help: "Specifies one or more URLs to VIB packages to install. http:, https:, ftp:, and file: are all supported. If 'file:' is used, then the file path must be an absolute path on the ESXi host."},
						Aliases:         []string{"-v", "--viburl"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "fields:DpuInstallationResult", Value: "DPU ID,Message,VIBs Installed,VIBs Removed,VIBs Skipped"},
					{Key: "fields:InstallationResult", Value: "Message,VIBs Installed,VIBs Removed,VIBs Skipped,Reboot Required,DPU Results"},
					{Key: "formatter", Value: "simple"},
					{Key: "header:InstallationResult", Value: "Installation Result"},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "list", DisplayName: "list", Help: "Lists the installed VIB packages on the host and the managed DPU(s)."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "rebootingimage", DisplayName: "rebooting-image", Help: "Displays information for the ESXi image which becomes active after a reboot, or nothing if the pending-reboot image has not been created yet. If not specified, information from the current ESXi image in memory will be returned."},
						Aliases:         []string{"--rebooting-image"},
						Flag:            true,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "fields:VIBSummaryExt", Value: "Name,Version,Vendor,Acceptance Level,Install Date,Platforms"},
					{Key: "formatter", Value: "table"},
					{Key: "header:VIBSummaryExt", Value: "%ID%"},
					{Key: "show-header", Value: "true"},
					{Key: "table-columns", Value: "Name,Version,Vendor,Acceptance Level,Install Date,Platforms"},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "remove", DisplayName: "remove", Help: "Removes VIB packages from the host. WARNING: If your installation requires a reboot, you need to disable HA first."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "dryrun", DisplayName: "dry-run", Help: "Performs a dry-run only. Report the VIB-level operations that would be performed, but do not change anything in the system."},
						Aliases:         []string{"--dry-run"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "force", DisplayName: "force", Help: "Bypasses checks for package dependencies, conflicts, obsolescence, and acceptance levels. Really not recommended unless you know what you are doing. Use of this option will result in a warning being displayed in vSphere Web Client.  Use this option only when instructed to do so by VMware Technical Support."},
						Aliases:         []string{"-f", "--force"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "maintenancemode", DisplayName: "maintenance-mode", Help: "Pretends that maintenance mode is in effect. Otherwise, remove will stop for live removes that require maintenance mode. This flag has no effect for reboot required remediations."},
						Aliases:         []string{"--maintenance-mode"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "noliveinstall", DisplayName: "no-live-install", Help: "Forces an remove to /altbootbank even if the VIBs are eligible for live removal. Will cause installation to be skipped on PXE-booted hosts."},
						Aliases:         []string{"--no-live-install"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "vibname", DisplayName: "vibname", Help: "Specifies one or more VIBs on the host to remove. Must be one of the following forms: name, name:version, vendor:name, vendor:name:version."},
						Aliases:         []string{"-n", "--vibname"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "fields:DpuInstallationResult", Value: "DPU ID,Message,VIBs Installed,VIBs Removed,VIBs Skipped"},
					{Key: "fields:InstallationResult", Value: "Message,VIBs Installed,VIBs Removed,VIBs Skipped,Reboot Required,DPU Results"},
					{Key: "formatter", Value: "simple"},
					{Key: "header:InstallationResult", Value: "Removal Result"},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "update", DisplayName: "update", Help: "Update installed VIBs to newer VIB packages. No new VIBs will be installed, only updates. WARNING: If your installation requires a reboot, you need to disable HA first."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "depot", DisplayName: "depot", Help: "Specifies full remote URLs of the depot index.xml or server file path pointing to an offline bundle .zip file."},
						Aliases:         []string{"-d", "--depot"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "dryrun", DisplayName: "dry-run", Help: "Performs a dry-run only. Report the VIB-level operations that would be performed, but do not change anything in the system."},
						Aliases:         []string{"--dry-run"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "force", DisplayName: "force", Help: "Bypasses checks for package dependencies, conflicts, obsolescence, and acceptance levels. Really not recommended unless you know what you are doing.  Use of this option will result in a warning being displayed in vSphere Web Client.  Use this option only when instructed to do so by VMware Technical Support."},
						Aliases:         []string{"-f", "--force"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "maintenancemode", DisplayName: "maintenance-mode", Help: "Pretends that maintenance mode is in effect. Otherwise, installation will stop for live installs that require maintenance mode. This flag has no effect for reboot required remediations."},
						Aliases:         []string{"--maintenance-mode"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "noliveinstall", DisplayName: "no-live-install", Help: "Forces an install to /altbootbank even if the VIBs are eligible for live installation or removal. Will cause installation to be skipped on PXE-booted hosts."},
						Aliases:         []string{"--no-live-install"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "nosigcheck", DisplayName: "no-sig-check", Help: "Bypasses acceptance level verification, including signing. Use of this option poses a large security risk and will result in a SECURITY ALERT warning being displayed in vSphere Web Client."},
						Aliases:         []string{"--no-sig-check"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "proxy", DisplayName: "proxy", Help: "Specifies a proxy server to use for HTTP, FTP, and HTTPS connections. The format is proxy-url:port."},
						Aliases:         []string{"--proxy"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "vibname", DisplayName: "vibname", Help: "Specifies VIBs from a depot, using one of the following forms: name, name:version, vendor:name, or vendor:name:version. VIB packages which are not updates will be skipped."},
						Aliases:         []string{"-n", "--vibname"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "viburl", DisplayName: "viburl", Help: "Specifies one or more URLs to VIB packages to update to. http:, https:, ftp:, and file: are all supported. VIB packages which are not updates will be skipped."},
						Aliases:         []string{"-v", "--viburl"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "fields:DpuInstallationResult", Value: "DPU ID,Message,VIBs Installed,VIBs Removed,VIBs Skipped"},
					{Key: "fields:InstallationResult", Value: "Message,VIBs Installed,VIBs Removed,VIBs Skipped,Reboot Required,DPU Results"},
					{Key: "formatter", Value: "simple"},
					{Key: "header:InstallationResult", Value: "Installation Result"},
				},
			},
		},
	},
	{
		CommandInfoItem: esx.CommandInfoItem{Name: "system.hostname", DisplayName: "hostname", Help: "Operations pertaining the network name of the ESX host."},
		Method: []esx.CommandInfoMethod{
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "get", DisplayName: "get", Help: "Get the host, domain or fully qualified name of the ESX host."},
				Param:           nil,
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: "simple"},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "set", DisplayName: "set", Help: "This command allows the user to set the hostname, domain name or fully qualified domain name of the ESX host."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "domain", DisplayName: "domain", Help: "The domain name to set for the ESX host. This option is mutually exclusive with the --fqdn option."},
						Aliases:         []string{"-d", "--domain"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "fqdn", DisplayName: "fqdn", Help: "Set the fully qualified domain name of the ESX host."},
						Aliases:         []string{"-f", "--fqdn"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "host", DisplayName: "host", Help: "The host name to set for the ESX host. This name should not contain the DNS domain name of the host and can only contain letters, numbers and '-'. NOTE this is not the fully qualified name, that can be set with the --fqdn option. This option is mutually exclusive with the --fqdn option."},
						Aliases:         []string{"-H", "--host"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: ""},
				},
			},
		},
	},
	{
		CommandInfoItem: esx.CommandInfoItem{Name: "system.settings.advanced", DisplayName: "advanced", Help: "The advanced settings are a set of VMkernel options that specific configuration settings to be modified. These options are typically in place for specific workarounds or debugging."},
		Method: []esx.CommandInfoMethod{
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "add", DisplayName: "add", Help: "Add a user defined advanced option to the /UserVars/ advanced option tree."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "description", DisplayName: "description", Help: "Description of the new option."},
						Aliases:         []string{"-d", "--description"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "hidden", DisplayName: "hidden", Help: "Whether the option is hidden."},
						Aliases:         []string{"-H", "--hidden"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "hostspecific", DisplayName: "host-specific", Help: "This indicates that the value of this option is always unique to a host."},
						Aliases:         []string{"-O", "--host-specific"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "impact", DisplayName: "impact", Help: "This specifies the impact on the host when the value of the option is modified: \n    maintenance-mode: This indicates that the host must be in maintenance mode before the option value is modified.\n    reboot: This indicates that the host must be rebooted for the option value to take effect.\n"},
						Aliases:         []string{"-I", "--impact"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "intdefault", DisplayName: "int-default", Help: "The default value of the new option (integer option only, required)."},
						Aliases:         []string{"-i", "--int-default"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "max", DisplayName: "max", Help: "The maximum allowed value (integer option only, required)."},
						Aliases:         []string{"-M", "--max"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "min", DisplayName: "min", Help: "The minimum allowed value (integer option only, required)."},
						Aliases:         []string{"-m", "--min"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "option", DisplayName: "option", Help: "The name of the new option. Valid characters: letters, digits and underscore."},
						Aliases:         []string{"-o", "--option"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "stringdefault", DisplayName: "string-default", Help: "The default value of the new option (string option only). An empty string is assumed if not specified."},
						Aliases:         []string{"-s", "--string-default"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "type", DisplayName: "type", Help: "The type of the new option. Supported values: \n    integer: Advanced option with integer value.\n    string: Advanced option with string value.\n"},
						Aliases:         []string{"-t", "--type"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: ""},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "list", DisplayName: "list", Help: "List the advanced options available from the VMkernel."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "delta", DisplayName: "delta", Help: "Only display options whose values differ from their default."},
						Aliases:         []string{"-d", "--delta"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "option", DisplayName: "option", Help: "Only get the information for a single VMkernel advanced option."},
						Aliases:         []string{"-o", "--option"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "tree", DisplayName: "tree", Help: "Limit the list of advanced option to a specific sub tree."},
						Aliases:         []string{"-t", "--tree"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "fields:SettingsAdvancedOption", Value: "Path,Type,Int Value,Default Int Value,Min Value,Max Value,String Value,Default String Value,Valid Characters,Description,Host Specific,Impact"},
					{Key: "formatter", Value: "simple"},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "remove", DisplayName: "remove", Help: "Remove a user defined advanced option from the /UserVars/ advanced option tree."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "option", DisplayName: "option", Help: "The name of the option to remove (without the /UserVars/ prefix as it is implied)."},
						Aliases:         []string{"-o", "--option"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: ""},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "set", DisplayName: "set", Help: "Set the value of an advanced option."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "default", DisplayName: "default", Help: "Reset the option to its default value."},
						Aliases:         []string{"-d", "--default"},
						Flag:            true,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "intvalue", DisplayName: "int-value", Help: "If the option is an integer value use this option."},
						Aliases:         []string{"-i", "--int-value"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "option", DisplayName: "option", Help: "The name of the option to set the value of. Example: \"/Misc/HostName\""},
						Aliases:         []string{"-o", "--option"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "stringvalue", DisplayName: "string-value", Help: "If the option is a string use this option."},
						Aliases:         []string{"-s", "--string-value"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: ""},
				},
			},
		},
	},
	{
		CommandInfoItem: esx.CommandInfoItem{Name: "system.stats.uptime", DisplayName: "uptime", Help: "System uptime, in microseconds"},
		Method: []esx.CommandInfoMethod{
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "get", DisplayName: "get", Help: "Display the number of microseconds the system has been running."},
				Param:           nil,
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: "simple"},
				},
			},
		},
	},
	{
		CommandInfoItem: esx.CommandInfoItem{Name: "vm.process", DisplayName: "process", Help: "Operations on running virtual machine processes"},
		Method: []esx.CommandInfoMethod{
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "kill", DisplayName: "kill", Help: "Used to forcibly kill Virtual Machines that are stuck and not responding to normal stop operations."},
				Param: []esx.CommandInfoParam{
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "type", DisplayName: "type", Help: "The type of kill operation to attempt. There are three types of VM kills that can be attempted: [soft, hard, force]. Users should always attempt 'soft' kills first, which will give the VMX process a chance to shutdown cleanly (like kill or kill -SIGTERM). If that does not work move to 'hard' kills which will shutdown the process immediately (like kill -9 or kill -SIGKILL). 'force' should be used as a last resort attempt to kill the VM. If all three fail then a reboot is required."},
						Aliases:         []string{"-t", "--type"},
						Flag:            false,
					},
					{
						CommandInfoItem: esx.CommandInfoItem{Name: "worldid", DisplayName: "world-id", Help: "The World ID of the Virtual Machine to kill. This can be obtained from the 'vm process list' command"},
						Aliases:         []string{"-w", "--world-id"},
						Flag:            false,
					},
				},
				Hints: esx.CommandInfoHints{
					{Key: "formatter", Value: ""},
				},
			},
			{
				CommandInfoItem: esx.CommandInfoItem{Name: "list", DisplayName: "list", Help: "List the virtual machines on this system. This command currently will only list running VMs on the system."},
				Param:           nil,
				Hints: esx.CommandInfoHints{
					{Key: "fields:VirtualMachine", Value: "World ID,Process ID,VMX Cartel ID,UUID,Display Name,Config File"},
					{Key: "formatter", Value: "simple"},
					{Key: "header:VirtualMachine", Value: "%Display Name%"},
				},
			},
		},
	},
}
