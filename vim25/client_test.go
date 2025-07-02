// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vim25

import (
	"context"
	"encoding/json"
	"net/url"
	"os"
	"testing"

	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

// Duplicated to prevent cyclic dependency...
func testURL(t *testing.T) *url.URL {
	s := os.Getenv("GOVMOMI_TEST_URL")
	if s == "" {
		t.SkipNow()
	}
	u, err := soap.ParseURL(s)
	if err != nil {
		panic(err)
	}
	return u
}

func sessionLogin(t *testing.T, c *Client) {
	req := types.Login{
		This: *c.ServiceContent.SessionManager,
	}

	u := testURL(t).User
	req.UserName = u.Username()
	if pw, ok := u.Password(); ok {
		req.Password = pw
	}

	_, err := methods.Login(context.Background(), c, &req)
	if err != nil {
		t.Fatal(err)
	}
}

func sessionCheck(t *testing.T, c *Client) {
	var mgr mo.SessionManager

	err := mo.RetrieveProperties(context.Background(), c, c.ServiceContent.PropertyCollector, *c.ServiceContent.SessionManager, &mgr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientSerialization(t *testing.T) {
	var c1, c2 *Client

	soapClient := soap.NewClient(testURL(t), true)
	c1, err := NewClient(context.Background(), soapClient)
	if err != nil {
		t.Fatal(err)
	}

	// Login
	sessionLogin(t, c1)
	sessionCheck(t, c1)

	// Serialize/deserialize
	b, err := json.Marshal(c1)
	if err != nil {
		t.Fatal(err)
	}
	c2 = &Client{}
	err = json.Unmarshal(b, c2)
	if err != nil {
		t.Fatal(err)
	}

	// Check the session is still valid
	sessionCheck(t, c2)
}
