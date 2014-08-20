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

package list

import (
	"path"
	"path/filepath"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/types"
)

type Recurser struct {
	Client *govmomi.Client

	All bool // Fetch complete objects for leaf nodes
}

func (r Recurser) Recurse(root types.ManagedObjectReference, prefix string, parts []string) ([]Element, error) {
	k := Lister{
		Client:    r.Client,
		Reference: root,
		Prefix:    prefix,
	}

	if r.All && len(parts) < 2 {
		k.All = true
	}

	in, err := k.List()
	if err != nil {
		return nil, err
	}

	// This folder is a leaf as far as the glob goes.
	if len(parts) == 0 {
		return in, nil
	}

	pattern := parts[0]
	parts = parts[1:]

	var out []Element
	for _, e := range in {
		ref := e.Object.Reference()
		matched, err := filepath.Match(pattern, path.Base(e.Path))
		if err != nil {
			return nil, err
		}

		if !matched {
			continue
		}

		// Include non-traversable leaf elements in result.
		// (consider the pattern "./vm/my-vm-*")
		if len(parts) == 0 && !traversable(ref) {
			out = append(out, e)
			continue
		}

		nres, err := r.Recurse(ref, path.Join(prefix, path.Base(e.Path)), parts)
		if err != nil {
			return nil, err
		}

		out = append(out, nres...)
	}

	return out, nil
}
