/*
Copyright (c) 2021 VMware, Inc. All Rights Reserved.

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

package service

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"os"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type create struct {
	*flags.ClientFlag

	specType        string
	trustedProvider bool
	acceptEULA      bool
}

func init() {
	cli.Register("namespace.service.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.specType, "spec-type", "vsphere", "Type of Spec: only vsphere is supported right now")
	f.BoolVar(&cmd.trustedProvider, "trusted", false, "Define if this is a trusted provider")
	f.BoolVar(&cmd.acceptEULA, "accept-eula", false, "Auto accept EULA")
}

func (cmd *create) Description() string {
	return `Creates a vSphere Namespace Supervisor Service.

Examples:
  govc namespace.service.create manifest.yaml`
}

func (cmd *create) Usage() string {
	return "MANIFEST"
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	manifest := f.Args()
	if len(manifest) != 1 {
		return flag.ErrHelp
	}

	if cmd.specType != "vsphere" {
		return fmt.Errorf("only vsphere specs are accepted right now")
	}

	manifestFile, err := os.ReadFile(manifest[0])
	if err != nil {
		return fmt.Errorf("failed to read manifest file: %s", err)
	}

	content := base64.StdEncoding.EncodeToString(manifestFile)
	service := namespace.SupervisorService{
		VsphereService: namespace.SupervisorServicesVSphereSpec{
			VersionSpec: namespace.SupervisorServicesVSphereVersionCreateSpec{
				Content:         content,
				AcceptEula:      cmd.acceptEULA,
				TrustedProvider: cmd.trustedProvider,
			},
		},
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := namespace.NewManager(c)
	return m.CreateSupervisorService(ctx, &service)

}
