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

package namespace

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type namespaceFlag struct {
	*flags.StorageProfileFlag

	library flags.StringList
	vmclass flags.StringList
	storage []string
}

func (ns *namespaceFlag) Register(ctx context.Context, f *flag.FlagSet) {
	ns.StorageProfileFlag, ctx = flags.NewStorageProfileFlag(ctx, "storage")
	ns.StorageProfileFlag.Register(ctx, f)

	f.Var(&ns.library, "library", "Content library IDs to associate with the vSphere Namespace.")
	f.Var(&ns.vmclass, "vmclass", "Virtual machine class IDs to associate with the vSphere Namespace.")
}

func (ns *namespaceFlag) Process(ctx context.Context) error {
	if err := ns.StorageProfileFlag.Process(ctx); err != nil {
		return err
	}

	rc, err := ns.RestClient()
	if err != nil {
		return err
	}

	for i, name := range ns.library {
		l, err := flags.ContentLibrary(ctx, rc, name)
		if err == nil {
			ns.library[i] = l.ID
		}
	}

	ns.storage, err = ns.StorageProfileList(ctx)

	return err
}

func (ns *namespaceFlag) storageSpec() []namespace.StorageSpec {
	s := make([]namespace.StorageSpec, len(ns.storage))
	for i := range ns.storage {
		s[i].Policy = ns.storage[i]
	}
	return s
}

func (ns *namespaceFlag) vmServiceSpec() namespace.VmServiceSpec {
	return namespace.VmServiceSpec{
		ContentLibraries: ns.library,
		VmClasses:        ns.vmclass,
	}
}
