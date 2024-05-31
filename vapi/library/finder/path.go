/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package finder

import (
	"context"
	"fmt"
	"net/url"
	"path"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

// PathFinder is used to find the Datastore path of a library.Library, library.Item or library.File.
type PathFinder struct {
	m     *library.Manager
	c     *vim25.Client
	cache map[string]string
}

// NewPathFinder creates a new PathFinder instance.
func NewPathFinder(m *library.Manager, c *vim25.Client) *PathFinder {
	return &PathFinder{
		m:     m,
		c:     c,
		cache: make(map[string]string),
	}
}

// Path returns the absolute datastore path for a Library, Item or File.
// The cache is used by DatastoreName().
func (f *PathFinder) Path(ctx context.Context, r FindResult) (string, error) {
	switch l := r.GetResult().(type) {
	case library.Library:
		id := ""
		if len(l.Storage) != 0 {
			var err error
			id, err = f.datastoreName(ctx, l.Storage[0].DatastoreID)
			if err != nil {
				return "", err
			}
		}
		return fmt.Sprintf("[%s] contentlib-%s", id, l.ID), nil
	case library.Item:
		p, err := f.Path(ctx, r.GetParent())
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s/%s", p, l.ID), nil
	case library.File:
		return f.getFileItemPath(ctx, r)
	default:
		return "", fmt.Errorf("unsupported type=%T", l)
	}
}

// getFileItemPath returns the absolute datastore path for a library.File
func (f *PathFinder) getFileItemPath(ctx context.Context, r FindResult) (string, error) {
	name := r.GetName()

	dir, err := f.Path(ctx, r.GetParent())
	if err != nil {
		return "", err
	}

	p := path.Join(dir, name)

	lib := r.GetParent().GetParent().GetResult().(library.Library)
	if len(lib.Storage) == 0 {
		return p, nil
	}

	// storage file name includes a uuid, for example:
	//  "ubuntu-14.04.6-server-amd64.iso" -> "ubuntu-14.04.6-server-amd64_0653e3f3-b4f4-41fb-9b72-c4102450e3dc.iso"
	s, err := f.m.GetLibraryItemStorage(ctx, r.GetParent().GetID(), name)
	if err != nil {
		return p, err
	}
	// Currently there can only be 1 storage URI
	if len(s) == 0 {
		return p, nil
	}

	uris := s[0].StorageURIs
	if len(uris) == 0 {
		return p, nil
	}
	u, err := url.Parse(uris[0])
	if err != nil {
		return p, err
	}

	return path.Join(dir, path.Base(u.Path)), nil
}

// datastoreName returns the Datastore.Name for the given id.
func (f *PathFinder) datastoreName(ctx context.Context, id string) (string, error) {
	if name, ok := f.cache[id]; ok {
		return name, nil
	}

	obj := types.ManagedObjectReference{
		Type:  "Datastore",
		Value: id,
	}

	ds := object.NewDatastore(f.c, obj)
	name, err := ds.ObjectName(ctx)
	if err != nil {
		return "", err
	}

	f.cache[id] = name
	return name, nil
}
