// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	"github.com/vmware/govmomi/cli"
	_ "github.com/vmware/govmomi/cli/about"
	_ "github.com/vmware/govmomi/cli/alarm"
	_ "github.com/vmware/govmomi/cli/cluster"
	_ "github.com/vmware/govmomi/cli/cluster/draft"
	_ "github.com/vmware/govmomi/cli/cluster/draft/baseimage"
	_ "github.com/vmware/govmomi/cli/cluster/draft/component"
	_ "github.com/vmware/govmomi/cli/cluster/group"
	_ "github.com/vmware/govmomi/cli/cluster/module"
	_ "github.com/vmware/govmomi/cli/cluster/override"
	_ "github.com/vmware/govmomi/cli/cluster/rule"
	_ "github.com/vmware/govmomi/cli/cluster/vlcm"
	_ "github.com/vmware/govmomi/cli/datacenter"
	_ "github.com/vmware/govmomi/cli/datastore"
	_ "github.com/vmware/govmomi/cli/datastore/cluster"
	_ "github.com/vmware/govmomi/cli/datastore/disk"
	_ "github.com/vmware/govmomi/cli/datastore/maintenance"
	_ "github.com/vmware/govmomi/cli/datastore/vsan"
	_ "github.com/vmware/govmomi/cli/device"
	_ "github.com/vmware/govmomi/cli/device/cdrom"
	_ "github.com/vmware/govmomi/cli/device/clock"
	_ "github.com/vmware/govmomi/cli/device/floppy"
	_ "github.com/vmware/govmomi/cli/device/model"
	_ "github.com/vmware/govmomi/cli/device/pci"
	_ "github.com/vmware/govmomi/cli/device/scsi"
	_ "github.com/vmware/govmomi/cli/device/serial"
	_ "github.com/vmware/govmomi/cli/device/usb"
	_ "github.com/vmware/govmomi/cli/disk"
	_ "github.com/vmware/govmomi/cli/disk/metadata"
	_ "github.com/vmware/govmomi/cli/disk/snapshot"
	_ "github.com/vmware/govmomi/cli/dvs"
	_ "github.com/vmware/govmomi/cli/dvs/portgroup"
	_ "github.com/vmware/govmomi/cli/env"
	_ "github.com/vmware/govmomi/cli/events"
	_ "github.com/vmware/govmomi/cli/export"
	_ "github.com/vmware/govmomi/cli/extension"
	_ "github.com/vmware/govmomi/cli/fields"
	_ "github.com/vmware/govmomi/cli/folder"
	_ "github.com/vmware/govmomi/cli/gpu"
	_ "github.com/vmware/govmomi/cli/host"
	_ "github.com/vmware/govmomi/cli/host/account"
	_ "github.com/vmware/govmomi/cli/host/autostart"
	_ "github.com/vmware/govmomi/cli/host/cert"
	_ "github.com/vmware/govmomi/cli/host/date"
	_ "github.com/vmware/govmomi/cli/host/esxcli"
	_ "github.com/vmware/govmomi/cli/host/firewall"
	_ "github.com/vmware/govmomi/cli/host/maintenance"
	_ "github.com/vmware/govmomi/cli/host/option"
	_ "github.com/vmware/govmomi/cli/host/portgroup"
	_ "github.com/vmware/govmomi/cli/host/service"
	_ "github.com/vmware/govmomi/cli/host/storage"
	_ "github.com/vmware/govmomi/cli/host/tpm"
	_ "github.com/vmware/govmomi/cli/host/vnic"
	_ "github.com/vmware/govmomi/cli/host/vswitch"
	_ "github.com/vmware/govmomi/cli/importx"
	_ "github.com/vmware/govmomi/cli/kms"
	_ "github.com/vmware/govmomi/cli/kms/key"
	_ "github.com/vmware/govmomi/cli/library"
	_ "github.com/vmware/govmomi/cli/library/policy"
	_ "github.com/vmware/govmomi/cli/library/session"
	_ "github.com/vmware/govmomi/cli/library/subscriber"
	_ "github.com/vmware/govmomi/cli/library/trust"
	_ "github.com/vmware/govmomi/cli/license"
	_ "github.com/vmware/govmomi/cli/logs"
	_ "github.com/vmware/govmomi/cli/ls"
	_ "github.com/vmware/govmomi/cli/metric"
	_ "github.com/vmware/govmomi/cli/metric/interval"
	_ "github.com/vmware/govmomi/cli/namespace"
	_ "github.com/vmware/govmomi/cli/namespace/cluster"
	_ "github.com/vmware/govmomi/cli/namespace/service"
	_ "github.com/vmware/govmomi/cli/namespace/service/version"
	_ "github.com/vmware/govmomi/cli/namespace/vmclass"
	_ "github.com/vmware/govmomi/cli/object"
	_ "github.com/vmware/govmomi/cli/option"
	_ "github.com/vmware/govmomi/cli/permissions"
	_ "github.com/vmware/govmomi/cli/pool"
	_ "github.com/vmware/govmomi/cli/role"
	_ "github.com/vmware/govmomi/cli/session"
	_ "github.com/vmware/govmomi/cli/sso/group"
	_ "github.com/vmware/govmomi/cli/sso/idp"
	_ "github.com/vmware/govmomi/cli/sso/lpp"
	_ "github.com/vmware/govmomi/cli/sso/service"
	_ "github.com/vmware/govmomi/cli/sso/user"
	_ "github.com/vmware/govmomi/cli/storage/policy"
	_ "github.com/vmware/govmomi/cli/tags"
	_ "github.com/vmware/govmomi/cli/tags/association"
	_ "github.com/vmware/govmomi/cli/tags/category"
	_ "github.com/vmware/govmomi/cli/task"
	_ "github.com/vmware/govmomi/cli/vapp"
	_ "github.com/vmware/govmomi/cli/vcsa/access/consolecli"
	_ "github.com/vmware/govmomi/cli/vcsa/access/dcui"
	_ "github.com/vmware/govmomi/cli/vcsa/access/shell"
	_ "github.com/vmware/govmomi/cli/vcsa/access/ssh"
	_ "github.com/vmware/govmomi/cli/vcsa/log"
	_ "github.com/vmware/govmomi/cli/vcsa/proxy"
	_ "github.com/vmware/govmomi/cli/vcsa/shutdown"
	_ "github.com/vmware/govmomi/cli/version"
	_ "github.com/vmware/govmomi/cli/vlcm/depot/content/baseimages"
	_ "github.com/vmware/govmomi/cli/vlcm/depot/offline"
	_ "github.com/vmware/govmomi/cli/vm"
	_ "github.com/vmware/govmomi/cli/vm/check"
	_ "github.com/vmware/govmomi/cli/vm/dataset"
	_ "github.com/vmware/govmomi/cli/vm/dataset/entry"
	_ "github.com/vmware/govmomi/cli/vm/disk"
	_ "github.com/vmware/govmomi/cli/vm/guest"
	_ "github.com/vmware/govmomi/cli/vm/network"
	_ "github.com/vmware/govmomi/cli/vm/option"
	_ "github.com/vmware/govmomi/cli/vm/policy"
	_ "github.com/vmware/govmomi/cli/vm/rdm"
	_ "github.com/vmware/govmomi/cli/vm/snapshot"
	_ "github.com/vmware/govmomi/cli/vm/target"
	_ "github.com/vmware/govmomi/cli/volume"
	_ "github.com/vmware/govmomi/cli/volume/snapshot"
	_ "github.com/vmware/govmomi/cli/vsan"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
