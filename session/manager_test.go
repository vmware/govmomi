/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package session

import (
	"net/url"
	"testing"

	"github.com/vmware/govmomi/test"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

var serviceInstance = types.ManagedObjectReference{
	Type:  "ServiceInstance",
	Value: "ServiceInstance",
}

func getServiceContent(r soap.RoundTripper) (types.ServiceContent, error) {
	req := types.RetrieveServiceContent{
		This: serviceInstance,
	}

	res, err := methods.RetrieveServiceContent(context.TODO(), r, &req)
	if err != nil {
		return types.ServiceContent{}, err
	}

	return res.Returnval, nil
}

func sessionClient(u *url.URL, t *testing.T) *Manager {
	soapClient := soap.NewClient(u, true)
	serviceContent, err := getServiceContent(soapClient)
	if err != nil {
		t.Error(err)
	}

	return NewManager(soapClient, serviceContent)
}

func TestLogin(t *testing.T) {
	u := test.URL()
	if u == nil {
		t.SkipNow()
	}

	session := sessionClient(u, t)
	err := session.Login(context.Background(), u.User)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestLogout(t *testing.T) {
	u := test.URL()
	if u == nil {
		t.SkipNow()
	}

	session := sessionClient(u, t)
	err := session.Login(context.Background(), u.User)
	if err != nil {
		t.Error("Login Error: ", err)
	}

	err = session.Logout(context.Background())
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	err = session.Logout(context.Background())
	if err == nil {
		t.Errorf("Expected NotAuthenticated, got nil")
	}
}

func TestSessionIsActive(t *testing.T) {
	u := test.URL()
	if u == nil {
		t.SkipNow()
	}

	session := sessionClient(u, t)

	// Skip test against ESXi -- SessionIsActive is not implemented
	if session.serviceContent.About.ApiType != "VirtualCenter" {
		t.Skipf("Talking to %s instead of %s", session.serviceContent.About.ApiType, "VirtualCenter")
	}

	err := session.Login(context.Background(), u.User)
	if err != nil {
		t.Error("Login Error: ", err)
	}

	active, err := session.SessionIsActive(context.Background())
	if err != nil || !active {
		t.Errorf("Expected %t, got %t", true, active)
		t.Errorf("Expected nil, got %v", err)
	}

	session.Logout(context.Background())

	active, err = session.SessionIsActive(context.Background())
	if err == nil || active {
		t.Errorf("Expected %t, got %t", false, active)
		t.Errorf("Expected NotAuthenticated, got %v", err)
	}
}
