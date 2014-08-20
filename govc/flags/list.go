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

package flags

import (
	"flag"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/govc/flags/list"
)

type ListRelativeFunc func() (govmomi.Reference, error)

type ListFlag struct {
	*DatastoreFlag
	*OutputFlag
}

func (l *ListFlag) Register(f *flag.FlagSet) {}

func (l *ListFlag) Process() error { return nil }

func (l *ListFlag) ListSlice(args []string, tl bool, fn ListRelativeFunc) ([]list.Element, error) {
	var out []list.Element

	for _, arg := range args {
		es, err := l.List(arg, tl, fn)
		if err != nil {
			return nil, err
		}

		out = append(out, es...)
	}

	return out, nil
}

func (l *ListFlag) List(arg string, tl bool, fn ListRelativeFunc) ([]list.Element, error) {
	c, err := l.Client()
	if err != nil {
		return nil, err
	}

	root := list.Element{
		Path:   "/",
		Object: c.RootFolder(),
	}

	parts := list.ToParts(arg)

	if len(parts) > 0 {
		switch parts[0] {
		case "..": // Relative to datacenter, back to root
			// Remove every occurance of ..
			for len(parts) > 0 && parts[0] == ".." {
				parts = parts[1:]
			}
		case ".": // Relative to whatever
			rootObj, err := fn()
			if err != nil {
				return nil, err
			}

			root.Path = "/" + rootObj.Reference().Value
			root.Object = rootObj
			parts = parts[1:]
		}
	}

	r := list.Recurser{
		Client:        c,
		All:           l.JSON,
		TraverseLeafs: tl,
	}

	es, err := r.Recurse(root, parts)
	if err != nil {
		return nil, err
	}

	return es, nil
}
