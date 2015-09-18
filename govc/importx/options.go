/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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

package importx

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	"github.com/vmware/govmomi/ovf"
	"github.com/vmware/govmomi/vim25/types"
)

type Property struct {
	types.KeyValue
	Spec *ovf.Property `json:",omitempty"`
}

type Options struct {
	AllDeploymentOptions []string `json:",omitempty"`
	Deployment           string

	AllDiskProvisioningOptions []string `json:",omitempty"`
	DiskProvisioning           string

	AllIPAllocationPolicyOptions []string `json:",omitempty"`
	IPAllocationPolicy           string

	AllIPProtocolOptions []string `json:",omitempty"`
	IPProtocol           string

	PropertyMapping []Property `json:",omitempty"`

	PowerOn      bool
	InjectOvfEnv bool
	WaitForIP    bool
}

type OptionsFlag struct {
	Options Options

	path string
}

func (flag *OptionsFlag) Register(f *flag.FlagSet) {
	f.StringVar(&flag.path, "options", "", "Options spec file path for VM deployment")
}

func (flag *OptionsFlag) Process() error {
	if len(flag.path) > 0 {
		f, err := os.Open(flag.path)
		if err != nil {
			return err
		}
		defer f.Close()

		o, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(o, &flag.Options); err != nil {
			return err
		}
	}

	return nil
}
