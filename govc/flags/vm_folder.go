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
)

type VmFolderFlag struct {
	*DatacenterFlag

	folder *govmomi.Folder
	// TODO: optional name + lookup. using default dc vm folder for now.
}

func (f *VmFolderFlag) Register(fs *flag.FlagSet) {
}

func (f *VmFolderFlag) Process() error {
	return nil
}

func (f *VmFolderFlag) VmFolder() (*govmomi.Folder, error) {
	if f.folder != nil {
		return f.folder, nil
	}

	c, err := f.Client()
	if err != nil {
		return nil, err
	}

	dc, err := f.Datacenter()
	if err != nil {
		return nil, err
	}

	folders, err := dc.Folders(c)
	if err != nil {
		return nil, err
	}

	f.folder = &folders.VmFolder

	return f.folder, nil
}
