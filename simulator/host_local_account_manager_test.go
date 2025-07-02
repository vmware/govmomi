// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

func TestHostLocalAccountManager(t *testing.T) {
	ctx := context.Background()
	m := ESX()

	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	ts := m.Service.NewServer()
	defer ts.Close()

	c, err := govmomi.NewClient(ctx, ts.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	ref := types.ManagedObjectReference{Type: "HostLocalAccountManager", Value: "ha-localacctmgr"}

	createUserReq := &types.CreateUser{
		This: ref,
		User: &types.HostAccountSpec{
			Id: "userid",
		},
	}

	_, err = methods.CreateUser(ctx, c.Client, createUserReq)
	if err != nil {
		t.Fatal(err)
	}

	_, err = methods.CreateUser(ctx, c.Client, createUserReq)
	if err == nil {
		t.Fatal("expect err; got nil")
	}

	updateUserReq := &types.UpdateUser{
		This: ref,
		User: &types.HostAccountSpec{
			Id: "userid",
		},
	}

	_, err = methods.UpdateUser(ctx, c.Client, updateUserReq)
	if err != nil {
		t.Fatal(err)
	}

	removeUserReq := &types.RemoveUser{
		This:     ref,
		UserName: "userid",
	}

	_, err = methods.RemoveUser(ctx, c.Client, removeUserReq)
	if err != nil {
		t.Fatal(err)
	}

	_, err = methods.RemoveUser(ctx, c.Client, removeUserReq)
	if err == nil {
		t.Fatal("expect err; got nil")
	}
}
