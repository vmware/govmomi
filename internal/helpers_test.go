// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package internal_test

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
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

func TestClientUsingEnvoyHostGateway(t *testing.T) {
	prefix := "hgw"
	suffix := ".sock"
	randBytes := make([]byte, 16)
	_, err := rand.Read(randBytes)
	require.NoError(t, err)

	testSocketPath := filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes)+suffix)

	l, err := net.Listen("unix", testSocketPath)
	require.NoError(t, err)
	handler := &testHTTPServer{
		expectedURL: "http://localhost/foo",
		response:    "Hello, Unix socket!",
		t:           t,
	}
	server := http.Server{
		Handler: handler,
	}
	go server.Serve(l)
	defer server.Close()
	defer l.Close()

	// First make sure the test server works fine, since we're starting a goroutine.
	unixDialer := func(proto, addr string) (conn net.Conn, err error) {
		return net.Dial("unix", testSocketPath)
	}
	tr := &http.Transport{
		Dial: unixDialer,
	}
	client := &http.Client{Transport: tr}

	require.Eventually(t, func() bool {
		_, err := client.Get(handler.expectedURL)
		return err == nil
	}, 15*time.Second, 1*time.Second, "Expected test HTTP server to be up")

	envVar := "VCENTER_ENVOY_HOST_GATEWAY"
	oldValue := os.Getenv(envVar)
	defer os.Setenv(envVar, oldValue)
	os.Setenv(envVar, testSocketPath)

	// Build a new client using the test unix socket.
	vc := &vim25.Client{Client: soap.NewClient(&url.URL{}, true)}
	newClient := internal.ClientWithEnvoyHostGateway(vc)

	// An HTTP request made using the new client should hit the server listening on the Unix socket.
	resp, err := newClient.Get(handler.expectedURL)

	// ...but should successfully connect to the Unix socket set up for testing.
	require.NoError(t, err)
	response, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	require.Equal(t, response, []byte(handler.response))
}

type testHTTPServer struct {
	expectedURL string
	response    string
	t           *testing.T
}

func (t *testHTTPServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	require.Equal(t.t, "/foo", req.URL.Path)
	resp.Write([]byte(t.response))
}

func TestRewriteURLForHostGateway(t *testing.T) {
	testURL, err := url.Parse("https://foo.bar/baz?query_param=1")
	require.NoError(t, err)

	hostMoref := types.ManagedObjectReference{
		Type:  "HostSystem",
		Value: "host-123",
	}
	result := internal.HostGatewayTransferURL(testURL, hostMoref)
	require.Equal(t, "localhost", result.Host)
	require.Equal(t, "/hgw/host-123/baz", result.Path)
	values := url.Values{"query_param": []string{"1"}}
	require.Equal(t, values, result.Query())
}

func TestSoapArgument(t *testing.T) {
	arg := internal.ReflectManagedMethodExecuterSoapArgument{
		Name: "vibname",
		Val:  "<vibname>crx</vibname>",
	}

	val := arg.Value()
	if len(val) != 1 || val[0] != "crx" {
		t.Errorf("val=%s", val)
	}
}

func TestEsxcliName(t *testing.T) {
	name := internal.EsxcliName("vim.EsxCLI.software.vib.get")
	if name != "VimEsxCLISoftwareVibGet" {
		t.Errorf("name=%s", name)
	}
}
