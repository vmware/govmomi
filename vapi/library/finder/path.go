// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package finder

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
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

func (f *PathFinder) convertPath(
	ctx context.Context,
	dc *object.Datacenter,
	ds mo.Datastore,
	path string) (string, error) {

	if v := ds.Capability.TopLevelDirectoryCreateSupported; v != nil && *v {
		return path, nil
	}

	if dc == nil {
		entities, err := mo.Ancestors(
			ctx,
			f.c,
			f.c.ServiceContent.PropertyCollector,
			ds.Self)
		if err != nil {
			return "", fmt.Errorf("failed to find ancestors: %w", err)
		}
		for _, entity := range entities {
			if entity.Self.Type == "Datacenter" {
				dc = object.NewDatacenter(f.c, entity.Self)
				break
			}
		}
	}

	if dc == nil {
		return "", errors.New("failed to find datacenter")
	}

	m := object.NewDatastoreNamespaceManager(f.c)
	return m.ConvertNamespacePathToUuidPath(ctx, dc, path)
}

// ResolveLibraryItemStorage transforms the StorageURIs field in the provided
// storage items from a datastore URL, ex.
// "ds:///vmfs/volumes/DATASTORE_UUID/contentlib-LIB_UUID/ITEM_UUID/file.vmdk",
// to the format that includes the datastore name, ex.
// "[DATASTORE_NAME] contentlib-LIB_UUID/ITEM_UUID/file.vmdk".
//
// If datastoreMap is provided, then it will be updated with the datastores
// involved in the resolver. The properties name, summary.url, and
// capability.topLevelDirectoryCreateSupported will be available after the
// resolver completes.
//
// If a storage item resides on a datastore that does not support the creation
// of top-level directories, then this means the datastore is vSAN and the
// storage item path needs to be further converted. If this occurs, then the
// datacenter to which the datastore belongs is required. If the datacenter
// parameter is non-nil, it is used, otherwise the datacenter for each datastore
// is resolved as needed. It is much more efficient to send in the datacenter if
// it is known ahead of time that the content library is stored on a vSAN
// datastore.
func (f *PathFinder) ResolveLibraryItemStorage(
	ctx context.Context,
	datacenter *object.Datacenter,
	datastoreMap map[string]mo.Datastore,
	storage []library.Storage) error {

	// TODO:
	// - reuse PathFinder.cache
	// - the transform here isn't Content Library specific, but is currently
	//   the only known use case
	var ids []types.ManagedObjectReference

	if datastoreMap == nil {
		datastoreMap = map[string]mo.Datastore{}
	}

	// Currently ContentLibrary only supports a single storage backing, but this
	// future proofs things.
	for _, item := range storage {
		id := item.StorageBacking.DatastoreID
		if _, ok := datastoreMap[id]; ok {
			continue
		}
		datastoreMap[id] = mo.Datastore{}
		ids = append(
			ids,
			types.ManagedObjectReference{Type: "Datastore", Value: id})
	}

	var (
		datastores []mo.Datastore
		pc         = property.DefaultCollector(f.c)
		props      = []string{
			"name",
			"summary.url",
			"capability.topLevelDirectoryCreateSupported",
		}
	)

	if err := pc.Retrieve(ctx, ids, props, &datastores); err != nil {
		return err
	}

	for i := range datastores {
		datastoreMap[datastores[i].Self.Value] = datastores[i]
	}

	for _, item := range storage {
		ds := datastoreMap[item.StorageBacking.DatastoreID]
		dsURL := ds.Summary.Url

		for i := range item.StorageURIs {
			szURI := item.StorageURIs[i]
			uri, err := url.Parse(szURI)
			if err != nil {
				return fmt.Errorf(
					"failed to parse storage URI %q: %w", szURI, err)
			}

			uri.OmitHost = false            // `ds://` required for ConvertNamespacePathToUuidPath()
			uri.Path = path.Clean(uri.Path) // required for ConvertNamespacePathToUuidPath()
			uri.RawQuery = ""

			uriPath := uri.String()
			u, err := f.convertPath(ctx, datacenter, ds, uriPath)
			if err != nil {
				return fmt.Errorf("failed to convert path %q: %w", uriPath, err)
			}
			u = strings.TrimPrefix(u, dsURL)
			u = strings.TrimPrefix(u, "/")

			item.StorageURIs[i] = (&object.DatastorePath{
				Datastore: ds.Name,
				Path:      u,
			}).String()
		}
	}

	return nil
}
