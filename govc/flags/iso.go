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

	"github.com/vmware/govmomi/vim25/types"
)

type IsoFlag struct {
	*DatastoreFlag

	name string
}

func (f *IsoFlag) Register(fs *flag.FlagSet) {
	fs.StringVar(&f.name, "iso", "", "Path to ISO")
}

func (f *IsoFlag) Process() error {
	return nil
}

func (f *IsoFlag) IsSet() bool {
	return f.name != ""
}

func (f *IsoFlag) Path() (string, error) {
	return f.DatastorePath(f.name)
}

func (f *IsoFlag) Controller() types.BaseVirtualDevice {
	controller := &types.VirtualIDEController{
		VirtualController: types.VirtualController{
			VirtualDevice: types.VirtualDevice{
				Key: 200,
			},
		},
	}

	return controller
}

func (f *IsoFlag) Device() (*types.VirtualCdrom, error) {
	ds, err := f.Datastore()
	if err != nil {
		return nil, err
	}

	_, err = f.Stat(f.name)
	if err != nil {
		return nil, err
	}

	device := &types.VirtualCdrom{
		VirtualDevice: types.VirtualDevice{
			Key:           -2,
			ControllerKey: 200,
			UnitNumber:    -1,
			Backing: &types.VirtualCdromIsoBackingInfo{
				VirtualDeviceFileBackingInfo: types.VirtualDeviceFileBackingInfo{
					FileName: ds.Path(f.name),
				},
			},
		},
	}

	return device, nil
}
