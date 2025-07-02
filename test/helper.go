// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"context"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"testing"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

func HasDocker() bool {
	if runtime.GOOS != "linux" {
		return false
	}
	if _, err := exec.LookPath("docker"); err != nil {
		return false
	}
	return true
}

// URL parses the GOVMOMI_TEST_URL environment variable if set.
func URL() *url.URL {
	s := os.Getenv("GOVMOMI_TEST_URL")
	if s == "" {
		return nil
	}
	u, err := soap.ParseURL(s)
	if err != nil {
		panic(err)
	}
	return u
}

// NewAuthenticatedClient creates a new vim25.Client, authenticates the user
// specified in the test URL, and returns it.
func NewAuthenticatedClient(t *testing.T) *vim25.Client {
	u := URL()
	if u == nil {
		t.SkipNow()
	}

	soapClient := soap.NewClient(u, true)
	vimClient, err := vim25.NewClient(context.Background(), soapClient)
	if err != nil {
		t.Fatal(err)
	}

	req := types.Login{
		This: *vimClient.ServiceContent.SessionManager,
	}

	req.UserName = u.User.Username()
	if pw, ok := u.User.Password(); ok {
		req.Password = pw
	}

	_, err = methods.Login(context.Background(), vimClient, &req)
	if err != nil {
		t.Fatal(err)
	}

	return vimClient
}
