/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.
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

package internal_test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

func TestHostSystemManagementIPs(t *testing.T) {
	ips := internal.HostSystemManagementIPs(esx.HostSystem.Config.VirtualNicManagerInfo.NetConfig)

	if len(ips) != 1 {
		t.Fatalf("no mgmt ip found")
	}
	if ips[0].String() != "127.0.0.1" {
		t.Fatalf("Expected management ip %s, got %s", "127.0.0.1", ips[0].String())
	}
}

func TestUsingVCEnvoySidecar(t *testing.T) {
	t.Run("VC HTTPS port", func(t *testing.T) {
		scheme := "https"
		hostname := "my-vcenter"
		port := 443
		u := &url.URL{Scheme: scheme, Host: fmt.Sprintf("%s:%d", hostname, port)}
		client := &vim25.Client{Client: soap.NewClient(u, true)}
		usingSidecar := internal.UsingEnvoySidecar(client)
		require.False(t, usingSidecar)
	})
	t.Run("Envoy sidecar", func(t *testing.T) {
		scheme := "http"
		hostname := "localhost"
		port := 1080
		u := &url.URL{Scheme: scheme, Host: fmt.Sprintf("%s:%d", hostname, port)}
		client := &vim25.Client{Client: soap.NewClient(u, true)}
		usingSidecar := internal.UsingEnvoySidecar(client)
		require.True(t, usingSidecar)
	})
}
