/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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

package flags

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/library/finder"
	"github.com/vmware/govmomi/vapi/rest"
)

// errContentLibraryMatch is an error returned when a query returns more than one result.
type errContentLibraryMatch struct {
	// Type is the type of object being queried.
	Type string

	// Key is the key used to perform the query.
	Key string

	// Val is the value used to perform the query.
	Val string

	// Count is the number of objects returned.
	Count int
}

// Error returns the error string.
func (e errContentLibraryMatch) Error() string {
	kind := e.Type
	if kind == "" {
		kind = "library|item"
	}
	hint := ""
	if e.Count > 1 {
		hint = fmt.Sprintf(" (use %q ID instead of NAME)", kind)
	}
	return fmt.Sprintf("%q=%q matches %d items%s", e.Key, e.Val, e.Count, hint)
}

func ContentLibraryResult(ctx context.Context, c *rest.Client, kind string, path string) (finder.FindResult, error) {
	res, err := finder.NewFinder(library.NewManager(c)).Find(ctx, path)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, errContentLibraryMatch{Type: kind, Key: "path", Val: path, Count: len(res)}
	}
	return res[0], nil
}

// ContentLibrary attempts to find a content library with the given path,
// asserting 1 match of type library.Library.
func ContentLibrary(ctx context.Context, c *rest.Client, path string) (*library.Library, error) {
	r, err := ContentLibraryResult(ctx, c, "library", path)
	if err != nil {
		return nil, err
	}
	lib, ok := r.GetResult().(library.Library)
	if !ok {
		return nil, fmt.Errorf("%q is a %T", path, r)
	}
	return &lib, nil
}

// ContentLibraryItem attempts to find a content library with the given path,
// asserting 1 match of type library.Item.
func ContentLibraryItem(ctx context.Context, c *rest.Client, path string) (*library.Item, error) {
	r, err := ContentLibraryResult(ctx, c, "item", path)
	if err != nil {
		return nil, err
	}
	item, ok := r.GetResult().(library.Item)
	if !ok {
		return nil, fmt.Errorf("%q is a %T", path, r)
	}
	return &item, nil
}
