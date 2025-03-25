// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/library/finder"
	"github.com/vmware/govmomi/vim25/types"
)

type DatastoreFlag struct {
	common

	*DatacenterFlag

	Name string

	ds *object.Datastore
}

var datastoreFlagKey = flagKey("datastore")

// NewCustomDatastoreFlag creates and returns a new DatastoreFlag without
// trying to retrieve an existing one from the specified context.
func NewCustomDatastoreFlag(ctx context.Context) (*DatastoreFlag, context.Context) {
	v := &DatastoreFlag{}
	v.DatacenterFlag, ctx = NewDatacenterFlag(ctx)
	return v, ctx
}

func NewDatastoreFlag(ctx context.Context) (*DatastoreFlag, context.Context) {
	if v := ctx.Value(datastoreFlagKey); v != nil {
		return v.(*DatastoreFlag), ctx
	}

	v, ctx := NewCustomDatastoreFlag(ctx)
	ctx = context.WithValue(ctx, datastoreFlagKey, v)
	return v, ctx
}

func (f *DatastoreFlag) Register(ctx context.Context, fs *flag.FlagSet) {
	f.RegisterOnce(func() {
		f.DatacenterFlag.Register(ctx, fs)

		env := "GOVC_DATASTORE"
		value := os.Getenv(env)
		usage := fmt.Sprintf("Datastore [%s]", env)
		fs.StringVar(&f.Name, "ds", value, usage)
	})
}

func (f *DatastoreFlag) Process(ctx context.Context) error {
	return f.ProcessOnce(func() error {
		if err := f.DatacenterFlag.Process(ctx); err != nil {
			return err
		}
		return nil
	})
}

func (flag *DatastoreFlag) IsSet() bool {
	return flag.Name != ""
}

func (f *DatastoreFlag) Args(args []string) []object.DatastorePath {
	var files []object.DatastorePath

	for _, arg := range args {
		var p object.DatastorePath

		if p.FromString(arg) {
			f.Name = p.Datastore
		} else {
			p.Datastore = f.Name
			p.Path = arg
		}

		files = append(files, p)
	}

	return files
}

func (f *DatastoreFlag) Datastore() (*object.Datastore, error) {
	if f.ds != nil {
		return f.ds, nil
	}

	var p object.DatastorePath
	if p.FromString(f.Name) {
		// Example use case:
		//   -ds "$(govc object.collect -s vm/foo config.files.logDirectory)"
		f.Name = p.Datastore
	}

	finder, err := f.Finder()
	if err != nil {
		return nil, err
	}

	if f.ds, err = finder.DatastoreOrDefault(context.TODO(), f.Name); err != nil {
		return nil, err
	}

	return f.ds, nil
}

func (flag *DatastoreFlag) DatastoreIfSpecified() (*object.Datastore, error) {
	if flag.Name == "" {
		return nil, nil
	}
	return flag.Datastore()
}

func (f *DatastoreFlag) DatastorePath(name string) (string, error) {
	ds, err := f.Datastore()
	if err != nil {
		return "", err
	}

	return ds.Path(name), nil
}

func (f *DatastoreFlag) Stat(ctx context.Context, file string) (types.BaseFileInfo, error) {
	ds, err := f.Datastore()
	if err != nil {
		return nil, err
	}

	return ds.Stat(ctx, file)

}

func (f *DatastoreFlag) libraryPath(ctx context.Context, p string) (string, error) {
	vc, err := f.Client()
	if err != nil {
		return "", err
	}

	rc, err := f.RestClient()
	if err != nil {
		return "", err
	}

	m := library.NewManager(rc)

	r, err := finder.NewFinder(m).Find(ctx, p)
	if err != nil {
		return "", err
	}

	if len(r) != 1 {
		return "", fmt.Errorf("%s: %d found", p, len(r))
	}

	return finder.NewPathFinder(m, vc).Path(ctx, r[0])
}

// FileBacking converts the given file path for use as VirtualDeviceFileBackingInfo.FileName.
func (f *DatastoreFlag) FileBacking(ctx context.Context, file string, stat bool) (string, error) {
	u, err := url.Parse(file)
	if err != nil {
		return "", err
	}

	switch u.Scheme {
	case "library":
		return f.libraryPath(ctx, u.Path)
	case "ds":
		// datastore url, e.g. ds:///vmfs/volumes/$uuid/...
		return file, nil
	}

	var p object.DatastorePath
	if p.FromString(file) {
		// datastore is specified
		return file, nil
	}

	if stat {
		// Verify ISO exists
		if _, err := f.Stat(ctx, file); err != nil {
			return "", err
		}
	}

	return f.DatastorePath(file)
}
