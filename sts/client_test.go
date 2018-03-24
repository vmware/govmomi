/*
Copyright (c) 2018 VMware, Inc. All Rights Reserved.

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

package sts

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
)

func TestIssueHOK(t *testing.T) {
	url := os.Getenv("GOVC_TEST_URL")
	if url == "" {
		t.SkipNow()
	}

	u, err := soap.ParseURL(url)
	if err != nil {
		t.Fatal(err)
	}
	u.User = nil // avoid govmomi.NewClient Login() call

	// solution user must exist, for example run on vCenter VM:
	// % /usr/lib/vmware-vmafd/bin/dir-cli service create --name govmomi --cert /etc/vmware-vpx/ssl/govmomi.crt --ssoadminrole Administrator --ssogroups SolutionUsers
	cert, err := tls.LoadX509KeyPair("govmomi.crt", "govmomi.key")
	if err != nil {
		t.Skip(err.Error())
	}

	ctx := context.Background()

	c, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		t.Fatal(err)
	}

	if !c.IsVC() {
		t.SkipNow()
	}

	sts, err := NewClient(ctx, c.Client)
	if err != nil {
		t.Fatal(err)
	}

	req := TokenRequest{
		Certificate: &cert,
		Delegatable: true,
	}

	s, err := sts.Issue(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	m := session.NewManager(c.Client)

	header := soap.Header{
		Security: s,
	}

	err = m.LoginByToken(c.WithHeader(ctx, header))
	if err != nil {
		t.Fatal(err)
	}

	now, err := methods.GetCurrentTime(ctx, c)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("current time=%s", now)
}

func TestIssueBearer(t *testing.T) {
	url := os.Getenv("GOVC_TEST_URL")
	if url == "" {
		t.SkipNow()
	}

	u, err := soap.ParseURL(url)
	if err != nil {
		t.Fatal(err)
	}
	user := u.User
	u.User = nil // avoid govmomi.NewClient Login() call

	ctx := context.Background()

	c, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		t.Fatal(err)
	}

	if !c.IsVC() {
		t.SkipNow()
	}

	sts, err := NewClient(ctx, c.Client)
	if err != nil {
		t.Fatal(err)
	}

	// Test that either Certificate or Userinfo is set.
	_, err = sts.Issue(ctx, TokenRequest{})
	if err == nil {
		t.Error("expected error")
	}

	req := TokenRequest{
		Userinfo: user,
	}

	s, err := sts.Issue(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	m := session.NewManager(c.Client)

	header := soap.Header{
		Security: s,
	}

	err = m.LoginByToken(c.WithHeader(ctx, header))
	if err != nil {
		t.Fatal(err)
	}

	now, err := methods.GetCurrentTime(ctx, c)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("current time=%s", now)
}

func TestIssueActAs(t *testing.T) {
	url := os.Getenv("GOVC_TEST_URL")
	if url == "" {
		t.SkipNow()
	}

	u, err := soap.ParseURL(url)
	if err != nil {
		t.Fatal(err)
	}
	user := u.User
	u.User = nil // avoid govmomi.NewClient Login() call

	// solution user must exist, for example run on vCenter VM:
	// % /usr/lib/vmware-vmafd/bin/dir-cli service create --name govmomi --cert /etc/vmware-vpx/ssl/govmomi.crt --ssoadminrole Administrator --ssogroups SolutionUsers
	cert, err := tls.LoadX509KeyPair("govmomi.crt", "govmomi.key")
	if err != nil {
		t.Skip(err.Error())
	}

	ctx := context.Background()

	c, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		t.Fatal(err)
	}

	if !c.IsVC() {
		t.SkipNow()
	}

	sts, err := NewClient(ctx, c.Client)
	if err != nil {
		t.Fatal(err)
	}

	req := TokenRequest{
		Delegatable: true,
		Userinfo:    user,
	}

	s, err := sts.Issue(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	req = TokenRequest{
		ActAs:       s.Token,
		Certificate: &cert,
	}

	s, err = sts.Issue(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	m := session.NewManager(c.Client)

	header := soap.Header{
		Security: s,
	}

	err = m.LoginByToken(c.WithHeader(ctx, header))
	if err != nil {
		t.Fatal(err)
	}

	now, err := methods.GetCurrentTime(ctx, c)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("current time=%s", now)
}
