// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/simulator/esx"
	"github.com/vmware/govmomi/vim25/types"
)

func TestHostDatastoreSystem(t *testing.T) {
	s := New(NewServiceInstance(NewContext(), esx.ServiceContent, esx.RootFolder))

	ts := s.NewServer()
	defer ts.Close()

	ctx := context.Background()

	c, err := govmomi.NewClient(ctx, ts.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	host := object.NewHostSystem(c.Client, esx.HostSystem.Reference())

	dss, err := host.ConfigManager().DatastoreSystem(ctx)
	if err != nil {
		t.Error(err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	spec := types.HostNasVolumeSpec{
		Type:       string(types.HostFileSystemVolumeFileSystemTypeNFS),
		RemoteHost: "localhost",
	}

	tests := []func(string) (*object.Datastore, error){
		func(dir string) (*object.Datastore, error) {
			spec.LocalPath = dir
			spec.RemotePath = dir
			return dss.CreateNasDatastore(ctx, spec)
		},
		func(dir string) (*object.Datastore, error) {
			return dss.CreateLocalDatastore(ctx, filepath.Base(dir), dir)
		},
	}

	for _, create := range tests {
		for _, fail := range []bool{false, true} {
			if fail {
				_, err = create(pwd)
				if err == nil {
					t.Error("expected error")
				}

				// TODO: hds.Remove(ds)
				pwd = filepath.Join(pwd, "esx")
			} else {
				_, err = create(pwd)
				if err != nil {
					t.Error(err)
				}
			}
		}
	}

	for _, create := range tests {
		for _, dir := range []string{"./enoent", "host_datastore_system.go"} {
			_, err = create(dir)

			if err == nil {
				t.Error("expected error")
			}
		}
	}
}

func TestCreateNasDatastoreValidation(t *testing.T) {
	s := New(NewServiceInstance(NewContext(), esx.ServiceContent, esx.RootFolder))

	ts := s.NewServer()
	defer ts.Close()

	ctx := context.Background()

	c, err := govmomi.NewClient(ctx, ts.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	host := object.NewHostSystem(c.Client, esx.HostSystem.Reference())

	dss, err := host.ConfigManager().DatastoreSystem(ctx)
	if err != nil {
		t.Error(err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		spec types.HostNasVolumeSpec
	}{
		{
			"RemotePath is not specified",
			types.HostNasVolumeSpec{
				Type:       string(types.HostFileSystemVolumeFileSystemTypeNFS),
				LocalPath:  pwd,
				RemoteHost: "localhost",
			},
		},
		{
			"RemoteHost is not specified",
			types.HostNasVolumeSpec{
				Type:       string(types.HostFileSystemVolumeFileSystemTypeNFS),
				LocalPath:  pwd,
				RemotePath: pwd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := dss.CreateNasDatastore(ctx, tt.spec)

			if err == nil {
				t.Error("expected error")
			}
		})
	}
}
