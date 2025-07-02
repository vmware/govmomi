// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package authentication

import (
	"context"
	"net/http"

	"github.com/vmware/govmomi/vapi/rest"
)

// Manager extends rest.Client, adding authentication related methods.
type Manager struct {
	*rest.Client
}

// NewManager creates a new Manager instance with the given client.
func NewManager(client *rest.Client) *Manager {
	return &Manager{
		Client: client,
	}
}

type TokenIssueSpec struct {
	SubjectToken       string `json:"subject_token"`
	SubjectTokenType   string `json:"subject_token_type"`
	GrantType          string `json:"grant_type"`
	ActorToken         string `json:"actor_token,omitempty"`
	ActorTokenType     string `json:"actor_token_type,omitempty"`
	RequestedTokenType string `json:"requested_token_type,omitempty"`
	Resource           string `json:"resource,omitempty"`
	Scope              string `json:"scope,omitempty"`
	Audience           string `json:"audience,omitempty"`
}

type TokenInfo struct {
	AccessToken     string `json:"access_token"`
	TokenType       string `json:"token_type"`
	ExpiresIn       int    `json:"expires_in,omitempty"`
	IssuedTokenType string `json:"issued_token_type,omitempty"`
	RefreshToken    string `json:"refresh_token,omitempty"`
	Scope           string `json:"scope,omitempty"`
}

func (c *Manager) Issue(ctx context.Context, token TokenIssueSpec) (*TokenInfo, error) {
	url := c.Resource("/vcenter/tokenservice/token-exchange")

	var res TokenInfo

	spec := struct {
		Spec TokenIssueSpec `json:"spec"`
	}{Spec: token}

	err := c.Do(ctx, url.Request(http.MethodPost, spec), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
