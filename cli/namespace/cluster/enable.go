// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cli/storage/policy"
	"github.com/vmware/govmomi/vapi/namespace"
)

type enableCluster struct {
	ControlPlaneDNSSearchDomains           string
	ImageStoragePolicy                     string
	NcpClusterNetworkSpec                  workloadNetwork
	ControlPlaneManagementNetwork          masterManagementNetwork
	ControlPlaneDNSNames                   string
	ControlPlaneNTPServers                 string
	EphemeralStoragePolicy                 string
	DefaultImageRepository                 string
	ServiceCidr                            string
	LoginBanner                            string
	SizeHint                               string
	WorkerDNS                              string
	DefaultImageRegistry                   string
	ControlPlaneDNS                        string
	NetworkProvider                        string
	ControlPlaneStoragePolicy              string
	DefaultKubernetesServiceContentLibrary string

	*flags.ClusterFlag
}

type masterManagementNetwork struct {
	Mode         string
	FloatingIP   string
	AddressRange *namespace.AddressRange
	Network      string
}

type workloadNetwork struct {
	NsxEdgeCluster string
	PodCidrs       string
	EgressCidrs    string
	Switch         string
	IngressCidrs   string
}

type objectReferences struct {
	Cluster                   string
	Network                   string
	ImageStoragePolicy        string
	ControlPlaneStoragePolicy string
	EphemeralStoragePolicy    string
	WorkerNetworkSwitch       string
	EdgeCluster               string
}

func init() {
	newEnableCluster := &enableCluster{
		ControlPlaneManagementNetwork: masterManagementNetwork{
			AddressRange: &namespace.AddressRange{},
		},
	}
	cli.Register("namespace.cluster.enable", newEnableCluster)
}

func (cmd *enableCluster) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	// this brings in some additional flags
	// Datacenter flag, which we do want in order to configure a Finder we use to do some lookups
	// ClientFlag, which we do need to control auth and connection details
	// OutputFlag, which doesn't have any effect in this case as there's no response to output
	cmd.ClusterFlag.Register(ctx, f)
	// Descriptions are mostly extracted from:
	// https://vmware.github.io/vsphere-automation-sdk-rest/vsphere/operations/com/vmware/vcenter/namespace_management/clusters.enable-operation.html
	f.StringVar(&cmd.SizeHint, "size", "",
		"The size of the Kubernetes API server and the worker nodes. Value is one of: TINY, SMALL, MEDIUM, LARGE.")
	f.StringVar(&cmd.ServiceCidr, "service-cidr", "",
		"CIDR block from which Kubernetes allocates service cluster IP addresses. Shouldn't overlap with pod, ingress or egress CIDRs")
	f.StringVar(&cmd.NetworkProvider, "network-provider", "NSXT_CONTAINER_PLUGIN",
		"Optional. Provider of cluster networking for this vSphere Namespaces cluster. Currently only value supported is: NSXT_CONTAINER_PLUGIN.")
	f.StringVar(&cmd.NcpClusterNetworkSpec.PodCidrs, "pod-cidrs", "",
		"CIDR blocks from which Kubernetes allocates pod IP addresses. Comma-separated list. Shouldn't overlap with service, ingress or egress CIDRs.")
	f.StringVar(&cmd.NcpClusterNetworkSpec.IngressCidrs, "workload-network.ingress-cidrs", "",
		"CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer. Comma-separated list. Shouldn't overlap with pod, service or egress CIDRs.")
	f.StringVar(&cmd.NcpClusterNetworkSpec.EgressCidrs, "workload-network.egress-cidrs", "",
		"CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs. Comma-separated list. Shouldn't overlap with pod, service or ingress CIDRs.")
	f.StringVar(&cmd.NcpClusterNetworkSpec.Switch, "workload-network.switch", "",
		"vSphere Distributed Switch used to connect this cluster.")
	f.StringVar(&cmd.NcpClusterNetworkSpec.NsxEdgeCluster, "workload-network.edge-cluster", "",
		"NSX Edge Cluster to be used for Kubernetes Services of type LoadBalancer, Kubernetes Ingresses, and NSX SNAT.")
	f.StringVar(&cmd.ControlPlaneManagementNetwork.Network, "mgmt-network.network", "",
		"Identifier for the management network.")
	f.StringVar(&cmd.ControlPlaneManagementNetwork.Mode, "mgmt-network.mode", "STATICRANGE",
		"IPv4 address assignment modes. Value is one of: DHCP, STATICRANGE")
	f.StringVar(&cmd.ControlPlaneManagementNetwork.FloatingIP, "mgmt-network.floating-IP", "",
		"Optional. The Floating IP used by the HA master cluster in the when network Mode is DHCP.")
	f.StringVar(&cmd.ControlPlaneManagementNetwork.AddressRange.StartingAddress, "mgmt-network.starting-address", "",
		"Denotes the start of the IP range to be used. Optional, but required with network mode STATICRANGE.")
	f.IntVar(&cmd.ControlPlaneManagementNetwork.AddressRange.AddressCount, "mgmt-network.address-count", 5,
		"The number of IP addresses in the management range. Optional, but required with network mode STATICRANGE.")
	f.StringVar(&cmd.ControlPlaneManagementNetwork.AddressRange.SubnetMask, "mgmt-network.subnet-mask", "",
		"Subnet mask of the management network. Optional, but required with network mode STATICRANGE.")
	f.StringVar(&cmd.ControlPlaneManagementNetwork.AddressRange.Gateway, "mgmt-network.gateway", "",
		"Gateway to be used for the management IP range")
	f.StringVar(&cmd.ControlPlaneDNS, "control-plane-dns", "",
		"Comma-separated list of DNS server IP addresses to use on Kubernetes API server, specified in order of preference.")
	f.StringVar(&cmd.ControlPlaneDNSNames, "control-plane-dns-names", "",
		"Comma-separated list of DNS names to associate with the Kubernetes API server. These DNS names are embedded in the TLS certificate presented by the API server.")
	f.StringVar(&cmd.ControlPlaneDNSSearchDomains, "control-plane-dns-search-domains", "",
		"Comma-separated list of domains to be searched when trying to lookup a host name on Kubernetes API server, specified in order of preference.")
	f.StringVar(&cmd.ControlPlaneNTPServers, "control-plane-ntp-servers", "",
		"Optional. Comma-separated list of NTP server DNS names or IP addresses to use on Kubernetes API server, specified in order of preference. If unset, VMware Tools based time synchronization is enabled.")
	f.StringVar(&cmd.WorkerDNS, "worker-dns", "",
		"Comma-separated list of DNS server IP addresses to use on the worker nodes, specified in order of preference.")
	f.StringVar(&cmd.ControlPlaneStoragePolicy, "control-plane-storage-policy", "",
		"Storage Policy associated with Kubernetes API server.")
	f.StringVar(&cmd.EphemeralStoragePolicy, "ephemeral-storage-policy", "",
		"Storage Policy associated with ephemeral disks of all the Kubernetes Pods in the cluster.")
	f.StringVar(&cmd.ImageStoragePolicy, "image-storage-policy", "",
		"Storage Policy to be used for container images.")
	f.StringVar(&cmd.LoginBanner, "login-banner", "",
		"Optional. Disclaimer to be displayed prior to login via the Kubectl plugin.")
	// documented API is currently ambiguous with these duplicated fields, need to wait for this to be resolved
	//f.StringVar(&cmd.DefaultImageRegistry, "default-image-registry", "",
	//  "Optional. Default image registry to use when unspecified in the container image name. Defaults to Docker Hub.")
	//f.StringVar(&cmd.DefaultImageRepository, "default-image-repository", "",
	//  "Optional. Default image registry to use when unspecified in the container image name. Defaults to Docker Hub.")
	// TODO
	// f.StringVar(&cmd.DefaultKubernetesServiceContentLibrary, "default-kubernetes-service-content-library", "",
	//  "Optional. Content Library which holds the VM Images for vSphere Kubernetes Service. This Content Library should be subscribed to VMware's hosted vSphere Kubernetes Service Repository.")
}

func (cmd *enableCluster) Description() string {
	return `Enable vSphere Namespaces on the cluster.
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
    -image-storage-policy "vSAN Default Storage Policy"`
}

func (cmd *enableCluster) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	// Cluster object reference lookup
	cluster, err := cmd.Cluster()
	if err != nil {
		return err
	}
	id := cluster.Reference().Value

	finder, err := cmd.ClusterFlag.Finder()
	if err != nil {
		return err
	}

	// Network object reference lookup.
	networkRef, err := finder.Network(ctx, cmd.ControlPlaneManagementNetwork.Network)
	if err != nil {
		return err
	}

	// Storage policy object references lookup
	pbmc, err := cmd.PbmClient()
	if err != nil {
		return fmt.Errorf("error creating client for storage policy lookup: %s", err)
	}

	storagePolicyNames := make(map[string]string)
	storagePolicyNames["control-plane"] = cmd.ControlPlaneStoragePolicy
	storagePolicyNames["ephemeral"] = cmd.EphemeralStoragePolicy
	storagePolicyNames["image"] = cmd.ImageStoragePolicy
	// keep track of names we looked up, so we don't repeat lookups
	visited := make(map[string]string)
	storagePolicyRefs := make(map[string]string)
	for k, v := range storagePolicyNames {
		if _, exists := visited[v]; !exists {
			policies, err := policy.ListProfiles(ctx, pbmc, v)
			if err != nil {
				return fmt.Errorf("error looking up storage policy %q: %s", v, err)
			} else if len(policies) != 1 {
				return fmt.Errorf("could not find a unique storage policy ID for query %q", v)
			}
			visited[v] = policies[0].GetPbmProfile().ProfileId.UniqueId
			storagePolicyRefs[k] = policies[0].GetPbmProfile().ProfileId.UniqueId
		} else {
			storagePolicyRefs[k] = visited[v]
		}

	}

	// DVS Object reference lookup
	// We need an id returned from the namespace lookup here, not a regular managed object reference.
	// Similar approach in powerCLI here:
	// https://github.com/lamw/PowerCLI-Example-Scripts/blob/7e4b9b9c93c5ffaa0ac2fefa8e02e5f751c044b7/Modules/VMware.WorkloadManagement/VMware.WorkloadManagement.psm1#L123
	// Note that the data model returned means we get no chance to choose the switch by name.
	// We assume there's just one switch per cluster and bail out otherwise.
	m := namespace.NewManager(c)
	clusterId := cluster.Reference().Value
	switches, err := m.ListCompatibleDistributedSwitches(ctx, clusterId)
	if err != nil {
		return fmt.Errorf("error in compatible switch lookup: %s", err)
	} else if len(switches) != 1 {
		return fmt.Errorf("expected to find 1 namespace compatible switch in cluster %q, found %d",
			clusterId, len(switches))
	}

	switchId := switches[0].DistributedSwitch

	edgeClusterDisplayName := cmd.NcpClusterNetworkSpec.NsxEdgeCluster
	edgeClusters, err := m.ListCompatibleEdgeClusters(ctx, clusterId, switchId)
	if err != nil {
		return fmt.Errorf("error in compatible edge cluster lookup: %s", err)
	}

	matchingEdgeClusters := make([]string, 0)
	for _, v := range edgeClusters {
		if v.DisplayName == edgeClusterDisplayName {
			matchingEdgeClusters = append(matchingEdgeClusters, v.EdgeCluster)
		}
	}
	if len(matchingEdgeClusters) != 1 {
		return fmt.Errorf("Didn't find unique match for edge cluster %q, found %d objects",
			edgeClusterDisplayName, len(matchingEdgeClusters))
	}

	resolvedObjectRefs := objectReferences{
		Cluster:                   cluster.Reference().Value,
		Network:                   networkRef.Reference().Value,
		ControlPlaneStoragePolicy: storagePolicyRefs["control-plane"],
		EphemeralStoragePolicy:    storagePolicyRefs["ephemeral"],
		ImageStoragePolicy:        storagePolicyRefs["image"],
		WorkerNetworkSwitch:       switchId,
		EdgeCluster:               matchingEdgeClusters[0],
	}

	enableClusterSpec, err := cmd.toVapiSpec(resolvedObjectRefs)
	if err != nil {
		return err
	}

	err = m.EnableCluster(ctx, id, enableClusterSpec)
	if err != nil {
		return err
	}

	return nil
}

func (cmd *enableCluster) toVapiSpec(refs objectReferences) (*namespace.EnableClusterSpec, error) {

	podCidrs, err := splitCidrList(splitCommaSeparatedList(cmd.NcpClusterNetworkSpec.PodCidrs))
	if err != nil {
		return nil, fmt.Errorf("invalid workload-network.pod-cidrs value: %s", err)
	}
	egressCidrs, err := splitCidrList(splitCommaSeparatedList(cmd.NcpClusterNetworkSpec.EgressCidrs))
	if err != nil {
		return nil, fmt.Errorf("invalid workload-network.egress-cidrs value: %s", err)
	}
	ingressCidrs, err := splitCidrList(splitCommaSeparatedList(cmd.NcpClusterNetworkSpec.IngressCidrs))
	if err != nil {
		return nil, fmt.Errorf("invalid workload-network.ingress-cidrs value: %s", err)
	}
	serviceCidr, err := splitCidr(cmd.ServiceCidr)
	if err != nil {
		return nil, fmt.Errorf("invalid service-cidr value: %s", err)
	}
	var masterManagementNetwork *namespace.MasterManagementNetwork
	if (cmd.ControlPlaneManagementNetwork.Mode != "") ||
		(cmd.ControlPlaneManagementNetwork.FloatingIP != "") ||
		(cmd.ControlPlaneManagementNetwork.Network != "") {
		masterManagementNetwork = &namespace.MasterManagementNetwork{}
		masterManagementNetwork.AddressRange = cmd.ControlPlaneManagementNetwork.AddressRange
		masterManagementNetwork.FloatingIP = cmd.ControlPlaneManagementNetwork.FloatingIP
		ipam := namespace.IpAssignmentModeFromString(cmd.ControlPlaneManagementNetwork.Mode)
		masterManagementNetwork.Mode = &ipam
		masterManagementNetwork.Network = cmd.ControlPlaneManagementNetwork.Network
	}
	if masterManagementNetwork != nil {
		if (masterManagementNetwork.AddressRange.SubnetMask == "") &&
			(masterManagementNetwork.AddressRange.StartingAddress == "") &&
			(masterManagementNetwork.AddressRange.Gateway == "") &&
			(masterManagementNetwork.AddressRange.AddressCount == 0) {
			masterManagementNetwork.AddressRange = nil
		}
		masterManagementNetwork.Network = refs.Network
	}

	sh := namespace.SizingHintFromString(cmd.SizeHint)
	np := namespace.ClusterNetworkProviderFromString(cmd.NetworkProvider)

	spec := namespace.EnableClusterSpec{
		MasterDNSSearchDomains: splitCommaSeparatedList(cmd.ControlPlaneDNSSearchDomains),
		ImageStorage:           namespace.ImageStorageSpec{StoragePolicy: refs.ImageStoragePolicy},
		NcpClusterNetworkSpec: &namespace.NcpClusterNetworkSpec{
			NsxEdgeCluster:           refs.EdgeCluster,
			PodCidrs:                 podCidrs,
			EgressCidrs:              egressCidrs,
			ClusterDistributedSwitch: refs.WorkerNetworkSwitch,
			IngressCidrs:             ingressCidrs,
		},
		MasterManagementNetwork:                masterManagementNetwork,
		MasterDNSNames:                         splitCommaSeparatedList(cmd.ControlPlaneDNSNames),
		MasterNTPServers:                       splitCommaSeparatedList(cmd.ControlPlaneNTPServers),
		EphemeralStoragePolicy:                 refs.EphemeralStoragePolicy,
		DefaultImageRepository:                 cmd.DefaultImageRepository,
		ServiceCidr:                            serviceCidr,
		LoginBanner:                            cmd.LoginBanner,
		SizeHint:                               &sh,
		WorkerDNS:                              splitCommaSeparatedList(cmd.WorkerDNS),
		DefaultImageRegistry:                   nil,
		MasterDNS:                              splitCommaSeparatedList(cmd.ControlPlaneDNS),
		NetworkProvider:                        &np,
		MasterStoragePolicy:                    refs.ControlPlaneStoragePolicy,
		DefaultKubernetesServiceContentLibrary: cmd.DefaultKubernetesServiceContentLibrary,
	}
	return &spec, nil
}

func splitCidr(input string) (*namespace.Cidr, error) {
	parts := strings.Split(input, "/")
	if !(len(parts) == 2) || parts[0] == "" {
		return nil, fmt.Errorf("invalid cidr %q supplied, needs to be in form '192.168.1.0/24'", input)
	}

	prefix, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	result := namespace.Cidr{
		Address: parts[0],
		Prefix:  prefix,
	}
	return &result, nil
}

func splitCidrList(input []string) ([]namespace.Cidr, error) {
	var result []namespace.Cidr
	for i, cidrIn := range input {
		cidr, err := splitCidr(cidrIn)
		if err != nil {
			return nil, fmt.Errorf("parsing cidr %q in list position %d : %q", cidrIn, i, err)
		}
		result = append(result, *cidr)
	}
	return result, nil
}

func splitCommaSeparatedList(cslist string) []string {
	return deleteEmpty(strings.Split(cslist, ","))
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
