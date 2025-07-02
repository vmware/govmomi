// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package volume

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cns"
	"github.com/vmware/govmomi/cns/types"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/units"
	vim "github.com/vmware/govmomi/vim25/types"
)

type disk struct {
	*flags.DatastoreFlag
	*flags.StorageProfileFlag

	size units.ByteSize

	spec types.CnsVolumeCreateSpec

	meta types.CnsKubernetesEntityMetadata

	disk string
}

func init() {
	cli.Register("volume.create", &disk{}, true)
}

func (cmd *disk) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatastoreFlag, ctx = flags.NewDatastoreFlag(ctx)
	cmd.DatastoreFlag.Register(ctx, f)

	cmd.StorageProfileFlag, ctx = flags.NewStorageProfileFlag(ctx)
	cmd.StorageProfileFlag.Register(ctx, f)

	_ = cmd.size.Set("10G")
	f.Var(&cmd.size, "size", "Size of new volume")

	f.StringVar(&cmd.spec.VolumeType, "type", string(types.CnsVolumeTypeBlock), "Volume type")

	f.StringVar(&cmd.spec.Metadata.ContainerCluster.ClusterType, "cluster-type", string(types.CnsClusterTypeKubernetes), "Cluster type")
	f.StringVar(&cmd.spec.Metadata.ContainerCluster.ClusterFlavor, "cluster-flavor", string(types.CnsClusterFlavorVanilla), "Cluster flavor")
	f.StringVar(&cmd.spec.Metadata.ContainerCluster.ClusterDistribution, "cluster-distro", "KUBERNETES", "Cluster distribution")
	f.StringVar(&cmd.spec.Metadata.ContainerCluster.ClusterId, "cluster-id", "", "Cluster ID")
	f.StringVar(&cmd.spec.Metadata.ContainerCluster.VSphereUser, "vsphere-user", "", "vSphere user")

	f.StringVar(&cmd.meta.EntityName, "entity-name", "", "Entity name")
	f.StringVar(&cmd.meta.Namespace, "entity-namespace", "default", "Entity namespace")
	f.StringVar(&cmd.meta.EntityType, "entity-type", string(types.CnsKubernetesEntityTypePVC), "Entity type")
	f.Var((*keyValue)(&cmd.meta.Labels), "label", "Label")

	f.StringVar(&cmd.disk, "disk-id", "", "Backing disk id")
}

func (cmd *disk) Process(ctx context.Context) error {
	if err := cmd.DatastoreFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.StorageProfileFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *disk) Usage() string {
	return "NAME"
}

func (cmd *disk) Description() string {
	return `Create volume NAME on DS.

Examples:
  govc volume.create -size 10G my-volume`
}

func (cmd *disk) Run(ctx context.Context, f *flag.FlagSet) error {
	cmd.spec.Name = f.Arg(0)
	if cmd.spec.Name == "" {
		return flag.ErrHelp
	}

	vc, err := cmd.Client()
	if err != nil {
		return err
	}

	c, err := cmd.CnsClient()
	if err != nil {
		return err
	}

	ds, err := cmd.DatastoreIfSpecified()
	if err != nil {
		return err
	}

	if ds != nil {
		cmd.spec.Datastores = []vim.ManagedObjectReference{ds.Reference()}
	}

	profiles, err := cmd.StorageProfileList(ctx)
	if err != nil {
		return err
	}

	for _, id := range profiles {
		cmd.spec.Profile = append(cmd.spec.Profile, &vim.VirtualMachineDefinedProfileSpec{ProfileId: id})
	}

	cmd.spec.BackingObjectDetails = &types.CnsBlockBackingDetails{
		BackingDiskId: cmd.disk,
		CnsBackingObjectDetails: types.CnsBackingObjectDetails{
			CapacityInMb: int64(cmd.size) / units.MB,
		},
	}

	if cmd.meta.EntityName == "" {
		cmd.meta.EntityName = cmd.spec.Name
	}
	cmd.spec.Metadata.EntityMetadata = append(cmd.spec.Metadata.EntityMetadata, &cmd.meta)

	if cmd.spec.Metadata.ContainerCluster.VSphereUser == "" {
		s, err := session.NewManager(vc).UserSession(ctx)
		if err != nil {
			return err
		}
		cmd.spec.Metadata.ContainerCluster.VSphereUser = s.UserName
	}

	task, err := c.CreateVolume(ctx, []types.CnsVolumeCreateSpec{cmd.spec})
	if err != nil {
		return err
	}

	info, err := cns.GetTaskInfo(ctx, task)
	if err != nil {
		return err
	}

	res, err := cns.GetTaskResult(ctx, info)
	if err != nil {
		return err
	}

	fmt.Println(res.GetCnsVolumeOperationResult().VolumeId.Id)

	return nil
}
