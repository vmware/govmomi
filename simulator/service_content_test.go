// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"testing"

	"github.com/vmware/govmomi"
)

// TestServiceContentAboutVersion asserts that ServiceContent.About's Version and
// ApiVersion are independently settable on the Model and are what a connected client
// observes. These are the fields the vcsim -version and -api-version flags bind to.
// Version defaults to a hardcoded 6.5.0 in vpx.ServiceContent, so before -version existed
// a client saw that while apiVersion reported whatever -api-version was set to.
func TestServiceContentAboutVersion(t *testing.T) {
	const (
		version    = "9.0.0.0"
		apiVersion = "9.0.0.0"
	)

	m := VPX()
	m.ServiceContent.About.Version = version
	m.ServiceContent.About.ApiVersion = apiVersion

	if err := m.Create(); err != nil {
		t.Fatal(err)
	}
	defer m.Remove()

	s := m.Service.NewServer()
	defer s.Close()

	ctx := context.Background()

	c, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	about := c.Client.ServiceContent.About
	if about.Version != version {
		t.Errorf("About.Version = %q, want %q", about.Version, version)
	}
	if about.ApiVersion != apiVersion {
		t.Errorf("About.ApiVersion = %q, want %q", about.ApiVersion, apiVersion)
	}
}
