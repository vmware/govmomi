// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package associations

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
)

const (
	basePath         = "/api/vcenter/consumption-domains/zones/cluster"
	associationsPath = basePath + "/%s/associations"
)

// Manager extends rest.Client, adding vSphere Zone association related methods.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

// Status
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/vcenter/data-structures/ConsumptionDomains_Zones_Cluster_Associations_Status
type Status struct {
	Success bool `json:"success"`
}

// AddAssociations
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/vcenter/api/vcenter/consumption-domains/zones/cluster/zone/associations__action=add/post
func (c *Manager) AddAssociations(zone string, cluster string) error {
	path := c.Resource(fmt.Sprintf(associationsPath, zone)).WithParam("action", "add")
	req := path.Request(http.MethodPost, []string{cluster})
	var res Status
	if err := c.Do(context.Background(), req, &res); err != nil {
		return err
	}

	// This endpoint does not always return and error upon failure.
	// In such cases we need to parse the response to figure out the actual status.

	if !res.Success {
		return errors.New("unable to add associations")
	}

	return nil
}

// RemoveAssociations
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/vcenter/api/vcenter/consumption-domains/zones/cluster/zone/associations__action=remove/post/
func (c *Manager) RemoveAssociations(zone string, cluster string) error {
	path := c.Resource(fmt.Sprintf(associationsPath, zone)).WithParam("action", "remove")
	req := path.Request(http.MethodPost, []string{cluster})
	return c.Do(context.Background(), req, nil)
}

// GetAssociations
// https://developer.broadcom.com/xapis/vsphere-automation-api/latest/vcenter/api/vcenter/consumption-domains/zones/cluster/zone/associations/get/
func (c *Manager) GetAssociations(zone string) ([]string, error) {
	path := c.Resource(fmt.Sprintf(associationsPath, zone))
	req := path.Request(http.MethodGet)
	var res []string
	return res, c.Do(context.Background(), req, &res)
}
