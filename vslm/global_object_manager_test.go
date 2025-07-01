// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vslm_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/simulator"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vslm"
	_ "github.com/vmware/govmomi/vslm/simulator"
	vso "github.com/vmware/govmomi/vslm/types"
)

func TestList(t *testing.T) {
	model := simulator.VPX()
	model.Datastore = 4

	now := time.Now().Format(time.RFC3339Nano)

	simulator.Test(func(ctx context.Context, vc *vim25.Client) {
		datastores, err := find.NewFinder(vc).DatastoreList(ctx, "*")
		if err != nil {
			t.Fatal(err)
		}

		c, err := vslm.NewClient(ctx, vc)
		if err != nil {
			t.Fatal(err)
		}

		m := vslm.NewGlobalObjectManager(c)

		namespaces := []string{"foo", "bar", "baz"}

		dsIDs := make([]string, len(datastores))

		for i, ds := range datastores {
			dsIDs[i] = strings.SplitN(ds.Reference().Value, "-", 2)[1]

			for j, ns := range namespaces {
				spec := types.VslmCreateSpec{
					Name:         fmt.Sprintf("%s-disk-%d", ns, i+j),
					CapacityInMB: int64(i+j) * 10,
					BackingSpec: &types.VslmCreateSpecDiskFileBackingSpec{
						VslmCreateSpecBackingSpec: types.VslmCreateSpecBackingSpec{
							Datastore: ds.Reference(),
						},
					},
				}
				t.Logf("CreateDisk %s (%dMB) on datastore-%s", spec.Name, spec.CapacityInMB, dsIDs[i])
				task, err := m.CreateDisk(ctx, spec)
				if err != nil {
					t.Fatal(err)
				}
				disk, err := task.Wait(ctx, time.Hour)
				if err != nil {
					t.Fatal(err)
				}

				id := disk.(types.VStorageObject).Config.Id

				metadata := []types.KeyValue{
					{Key: "namespace", Value: ns},
					{Key: "name", Value: spec.Name},
				}

				task, err = m.UpdateMetadata(ctx, id, metadata, nil)
				if err != nil {
					t.Fatal(err)
				}
				_, err = task.Wait(ctx, time.Hour)
				if err != nil {
					t.Fatal(err)
				}
			}
		}

		tests := []struct {
			expect int
			query  []vso.VslmVsoVStorageObjectQuerySpec
		}{
			{model.Datastore * len(namespaces), nil},
			{-1, []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    "invalid",
				QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumEquals),
				QueryValue:    []string{"any"},
			}}},
			{-1, []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumName),
				QueryOperator: "invalid",
				QueryValue:    []string{"any"},
			}}},
			{-1, []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumName),
				QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumEquals),
				QueryValue:    nil,
			}}},
			{-1, []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumName),
				QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumEquals),
				QueryValue:    []string{"one", "two"},
			}}},
			{-1, []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumCapacity),
				QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumStartsWith),
				QueryValue:    []string{"10"},
			}}},
			{-1, []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumCapacity),
				QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumGreaterThan),
				QueryValue:    []string{"ten"},
			}}},
			{3, []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumCapacity),
				QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumGreaterThan),
				QueryValue:    []string{"30"},
			}}},
			{0, []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumCapacity),
				QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumGreaterThan),
				QueryValue:    []string{"5000"},
			}}},
			{len(namespaces), []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumDatastoreMoId),
				QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumEquals),
				QueryValue:    []string{dsIDs[0]},
			}}},
			{model.Datastore, []vso.VslmVsoVStorageObjectQuerySpec{
				{
					QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumMetadataKey),
					QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumEquals),
					QueryValue:    []string{"namespace"},
				},
				{
					QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumMetadataValue),
					QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumEquals),
					QueryValue:    []string{namespaces[0]},
				},
			}},
			{model.Datastore, []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumName),
				QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumStartsWith),
				QueryValue:    []string{namespaces[1]},
			}}},
			{model.Datastore, []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumName),
				QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumStartsWith),
				QueryValue:    []string{namespaces[1]},
			}}},
			{model.Datastore * len(namespaces), []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumName),
				QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumContains),
				QueryValue:    []string{"disk"},
			}}},
			{0, []vso.VslmVsoVStorageObjectQuerySpec{
				{
					QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumName),
					QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumStartsWith),
					QueryValue:    []string{namespaces[0]},
				},
				{
					QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumName),
					QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumStartsWith),
					QueryValue:    []string{namespaces[1]},
				},
			}},
			{model.Datastore * len(namespaces), []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumCreateTime),
				QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumGreaterThan),
				QueryValue:    []string{now},
			}}},
			{0, []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumCreateTime),
				QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumLessThan),
				QueryValue:    []string{now},
			}}},
			{-1, []vso.VslmVsoVStorageObjectQuerySpec{{
				QueryField:    string(vso.VslmVsoVStorageObjectQuerySpecQueryFieldEnumCreateTime),
				QueryOperator: string(vso.VslmVsoVStorageObjectQuerySpecQueryOperatorEnumContains),
				QueryValue:    []string{now},
			}}},
		}

		for _, test := range tests {
			if test.expect > 2 {
				vslm.DefaultMaxResults = test.expect / 2 // test pagination
			}
			t.Run(queryString(test.query), func(t *testing.T) {
				res, err := m.List(ctx, test.query...)
				if test.expect == -1 {
					if err == nil {
						t.Error("expected error")
					}
				} else {
					if err != nil {
						t.Fatal(err)
					}

					if len(res.Id) != test.expect {
						t.Errorf("expected %d, got: %d", test.expect, len(res.Id))
					}
				}
			})
		}
	}, model)
}

func queryString(query []vso.VslmVsoVStorageObjectQuerySpec) string {
	if query == nil {
		return "no query"
	}
	res := make([]string, len(query))
	for i, q := range query {
		res[i] = fmt.Sprintf("%s.%s=%s",
			q.QueryField, q.QueryOperator,
			strings.Join(q.QueryValue, ","))
	}
	return strings.Join(res, " and ")
}
