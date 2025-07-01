// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

import (
	"context"
	"errors"
	"net/http"

	"github.com/vmware/govmomi/vapi/internal"
)

const (
	OvfDefaultSecurityPolicy = "OVF default policy"
)

// ContentSecurityPoliciesInfo contains information on security policies that can
// be used to describe security for content library items.
type ContentSecurityPoliciesInfo struct {
	// ItemTypeRules are rules governing the policy.
	ItemTypeRules map[string]string `json:"item_type_rules"`
	// Name is a human-readable identifier identifying the policy.
	Name string `json:"name"`
	// Policy is the unique identifier for a policy.
	Policy string `json:"policy"`
}

// ListSecurityPolicies lists security policies
func (c *Manager) ListSecurityPolicies(ctx context.Context) ([]ContentSecurityPoliciesInfo, error) {
	url := c.Resource(internal.SecurityPoliciesPath)
	var res []ContentSecurityPoliciesInfo
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

func (c *Manager) DefaultOvfSecurityPolicy(ctx context.Context) (string, error) {
	res, err := c.ListSecurityPolicies(ctx)

	if err != nil {
		return "", err
	}

	for _, policy := range res {
		if policy.Name == OvfDefaultSecurityPolicy {
			return policy.Policy, nil
		}
	}

	return "", errors.New("failed to find default ovf security policy")
}
