// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	vim "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vsan"
	"github.com/vmware/govmomi/vsan/methods"
	"github.com/vmware/govmomi/vsan/types"
)

type stretch struct {
	*flags.DatacenterFlag

	WitnessHost            string
	FirstFaultDomainHosts  string
	SecondFaultDomainHosts string
	FirstFaultDomainName   string
	SecondFaultDomainName  string
	PreferredFaultDomain   string
}

func init() {
	cli.Register("cluster.stretch", &stretch{})
}

func (cmd *stretch) Usage() string {
	return "CLUSTER"
}

func (cmd *stretch) Description() string {
	return `Convert a vSAN cluster into a stretched cluster

The vSAN cluster is converted to a stretched cluster with a witness host
specified by the 'witness' flag.  The datastore hosts are placed into one
of two fault domains that are specified in each host list. The name of the
preferred fault domain can be specified by the 'preferred-fault-domain' flag.

Examples:
  govc cluster.stretch -dc remote-site-1 \
    -witness /dc-name/host/192.168.112.2 \
    -first-fault-domain-hosts 192.168.113.121 \
    -second-fault-domain-hosts 192.168.113.45,192.168.113.70 \
    cluster-name`
}

func (cmd *stretch) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.StringVar(&cmd.WitnessHost, "witness", "", "Witness host for the stretched cluster")
	f.StringVar(&cmd.FirstFaultDomainHosts, "first-fault-domain-hosts", "", "Hosts to place in the first fault domain")
	f.StringVar(&cmd.SecondFaultDomainHosts, "second-fault-domain-hosts", "", "Hosts to place in the second fault domain")
	f.StringVar(&cmd.FirstFaultDomainName, "first-fault-domain-name", "Primary", "Name of the first fault domain")
	f.StringVar(&cmd.SecondFaultDomainName, "second-fault-domain-name", "Secondary", "Name of the second fault domain")
	f.StringVar(&cmd.PreferredFaultDomain, "preferred-fault-domain", "Primary", "Name of the preferred fault domain")
}

func (cmd *stretch) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}

	if cmd.WitnessHost == "" ||
		cmd.FirstFaultDomainHosts == "" ||
		cmd.FirstFaultDomainName == "" ||
		cmd.SecondFaultDomainHosts == "" ||
		cmd.SecondFaultDomainName == "" ||
		cmd.PreferredFaultDomain == "" {
		return flag.ErrHelp
	}

	return nil
}

func (cmd *stretch) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	client, err := cmd.Client()
	if err != nil {
		return err
	}

	vsanClient, err := vsan.NewClient(ctx, client)
	if err != nil {
		return err
	}

	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	clusterResource, err := finder.ClusterComputeResource(ctx, f.Arg(0))
	if err != nil {
		return err
	}

	witnessHost, err := finder.HostSystem(ctx, cmd.WitnessHost)
	if err != nil {
		return err
	}

	faultDomainConfig, err := cmd.buildFaultDomainConfig(ctx)
	if err != nil {
		return err
	}

	req := types.VSANVcConvertToStretchedCluster{
		This:              vsan.VsanVcStretchedClusterSystem,
		Cluster:           clusterResource.Reference(),
		FaultDomainConfig: *faultDomainConfig,
		WitnessHost:       witnessHost.Reference(),
		PreferredFd:       cmd.PreferredFaultDomain,
		DiskMapping:       nil,
	}

	res, err := methods.VSANVcConvertToStretchedCluster(ctx, vsanClient, &req)
	if err != nil {
		return err
	}

	logger := cmd.ProgressLogger("stretching cluster... ")
	defer logger.Wait()

	task := object.NewTask(client, res.Returnval)
	_, err = task.WaitForResult(ctx, logger)
	return err
}

func (cmd *stretch) buildFaultDomainConfig(ctx context.Context) (*types.VimClusterVSANStretchedClusterFaultDomainConfig, error) {
	var faultDomainConfig types.VimClusterVSANStretchedClusterFaultDomainConfig
	var err error

	faultDomainConfig.FirstFdName = cmd.FirstFaultDomainName
	faultDomainConfig.FirstFdHosts, err = cmd.getManagedObjectRefs(cmd.FirstFaultDomainHosts, ctx)
	if err != nil {
		return nil, err
	}

	if len(faultDomainConfig.FirstFdHosts) == 0 {
		return nil, fmt.Errorf("no hosts for fault domain %q", cmd.FirstFaultDomainName)
	}

	faultDomainConfig.SecondFdName = cmd.SecondFaultDomainName
	faultDomainConfig.SecondFdHosts, err = cmd.getManagedObjectRefs(cmd.SecondFaultDomainHosts, ctx)
	if err != nil {
		return nil, err
	}

	if len(faultDomainConfig.SecondFdHosts) == 0 {
		return nil, fmt.Errorf("no hosts for fault domain %q", cmd.SecondFaultDomainName)
	}

	return &faultDomainConfig, nil
}

func (cmd *stretch) getManagedObjectRefs(domainHosts string, ctx context.Context) ([]vim.ManagedObjectReference, error) {
	finder, err := cmd.Finder()
	if err != nil {
		return nil, err
	}

	var refs []vim.ManagedObjectReference
	for _, host := range strings.Split(domainHosts, ",") {
		h, err := finder.HostSystem(ctx, host)
		if err != nil {
			return nil, err
		}

		refs = append(refs, h.Reference())
	}

	return refs, nil
}
