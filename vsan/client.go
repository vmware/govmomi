/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package vsan

import (
	"context"
	"errors"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	vimtypes "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vsan/methods"
	vsantypes "github.com/vmware/govmomi/vsan/types"
)

// Namespace and Path constants
const (
	Namespace = "vsan"
	Path      = "/vsanHealth"
)

// Creates the vsan cluster config system instance. This is to be queried from vsan health.
var (
	VsanVcClusterConfigSystemInstance = vimtypes.ManagedObjectReference{
		Type:  "VsanVcClusterConfigSystem",
		Value: "vsan-cluster-config-system",
	}
	VsanPerformanceManagerInstance = vimtypes.ManagedObjectReference{
		Type:  "VsanPerformanceManager",
		Value: "vsan-performance-manager",
	}
	VsanQueryObjectIdentitiesInstance = vimtypes.ManagedObjectReference{
		Type:  "VsanObjectSystem",
		Value: "vsan-cluster-object-system",
	}
	VsanPropertyCollectorInstance = vimtypes.ManagedObjectReference{
		Type:  "PropertyCollector",
		Value: "vsan-property-collector",
	}
)

// Client used for accessing vsan health APIs.
type Client struct {
	*soap.Client

	Vim25Client *vim25.Client
}

// NewClient creates a new VsanHealth client
func NewClient(ctx context.Context, c *vim25.Client) (*Client, error) {
	sc := c.Client.NewServiceClient(Path, Namespace)
	return &Client{sc, c}, nil
}

// VsanClusterGetConfig calls the Vsan health's VsanClusterGetConfig API.
func (c *Client) VsanClusterGetConfig(ctx context.Context, cluster vimtypes.ManagedObjectReference) (*vsantypes.VsanConfigInfoEx, error) {
	req := vsantypes.VsanClusterGetConfig{
		This:    VsanVcClusterConfigSystemInstance,
		Cluster: cluster,
	}

	res, err := methods.VsanClusterGetConfig(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return &res.Returnval, nil
}

// VsanPerfQueryPerf calls the vsan performance manager API
func (c *Client) VsanPerfQueryPerf(ctx context.Context, cluster *vimtypes.ManagedObjectReference, qSpecs []vsantypes.VsanPerfQuerySpec) ([]vsantypes.VsanPerfEntityMetricCSV, error) {
	req := vsantypes.VsanPerfQueryPerf{
		This:       VsanPerformanceManagerInstance,
		Cluster:    cluster,
		QuerySpecs: qSpecs,
	}

	res, err := methods.VsanPerfQueryPerf(ctx, c, &req)
	if err != nil {
		return nil, err
	}
	return res.Returnval, nil
}

// VsanQueryObjectIdentities return host uuid
func (c *Client) VsanQueryObjectIdentities(ctx context.Context, cluster vimtypes.ManagedObjectReference) (*vsantypes.VsanObjectIdentityAndHealth, error) {
	req := vsantypes.VsanQueryObjectIdentities{
		This:    VsanQueryObjectIdentitiesInstance,
		Cluster: &cluster,
	}

	res, err := methods.VsanQueryObjectIdentities(ctx, c, &req)

	if err != nil {
		return nil, err
	}

	return res.Returnval, nil
}

// VsanHostGetConfig returns the config of host's vSAN system.
func (c *Client) VsanHostGetConfig(ctx context.Context, vsanSystem vimtypes.ManagedObjectReference) (*vsantypes.VsanHostConfigInfoEx, error) {
	req := vimtypes.RetrievePropertiesEx{
		SpecSet: []vimtypes.PropertyFilterSpec{{
			ObjectSet: []vimtypes.ObjectSpec{{
				Obj: vsanSystem}},
			PropSet: []vimtypes.PropertySpec{{
				Type:    "HostVsanSystem",
				PathSet: []string{"config"}}}}},
		This: VsanPropertyCollectorInstance}

	res, err := methods.RetrievePropertiesEx(ctx, c, &req)
	if err != nil {
		return nil, err
	}

	var property vimtypes.DynamicProperty
	if res != nil && res.Returnval != nil {
		for _, obj := range res.Returnval.Objects {
			for _, prop := range obj.PropSet {
				if prop.Name == "config" {
					property = prop
					break
				}
			}
		}
	}

	switch cfg := property.Val.(type) {
	case vimtypes.VsanHostConfigInfo:
		return &vsantypes.VsanHostConfigInfoEx{VsanHostConfigInfo: cfg}, nil
	case vsantypes.VsanHostConfigInfoEx:
		return &cfg, nil
	default:
		return nil, errors.New("host vSAN config not found")
	}
}
