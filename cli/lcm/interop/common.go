// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package interop

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/lcm"
)

// newManager builds an lcm.Manager from the standard govc client flags.
// The LCM REST API uses HTTP Basic Auth; credentials and the base URL are
// taken directly from the session configuration.
func newManager(f *flags.ClientFlag) *lcm.Manager {
	u := f.Session.URL
	baseURL := &url.URL{Scheme: "https", Host: u.Host}

	username := ""
	password := ""
	if u.User != nil {
		username = u.User.Username()
		password, _ = u.User.Password()
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: f.Session.Insecure}, //nolint:gosec
		},
	}

	return lcm.NewManager(baseURL, username, password, httpClient)
}
