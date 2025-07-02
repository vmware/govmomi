// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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
	*ServiceVersionFlag
}

func init() {
	cli.Register("namespace.service.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.ServiceVersionFlag = &ServiceVersionFlag{}
	cmd.ServiceVersionFlag.Register(ctx, f)
}

func (cmd *create) Description() string {
	return `Registers a vSphere Supervisor Service version on vCenter for a new service.
A service version can be registered once on vCenter and then be installed on multiple vSphere Supervisors managed by this vCenter.
A vSphere Supervisor Service contains a list of service versions; this call will create a service and its first version.

Examples:
  govc namespace.service.create manifest.yaml`
}

func (cmd *create) Usage() string {
	return "MANIFEST"
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.ServiceVersionFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}

	manifestFile := f.Arg(0)

	manifest, err := os.ReadFile(manifestFile)
	if err != nil {
		return fmt.Errorf("failed to read manifest file: %s", err)
	}
	content := base64.StdEncoding.EncodeToString(manifest)

	service := namespace.SupervisorService{}
	if cmd.ServiceVersionFlag.SpecType == "carvel" {
		service.CarvelService = &namespace.SupervisorServicesCarvelSpec{
			VersionSpec: namespace.CarvelVersionCreateSpec{
				Content: content,
			},
		}
	} else {
		service.VsphereService = &namespace.SupervisorServicesVSphereSpec{
			VersionSpec: namespace.SupervisorServicesVSphereVersionCreateSpec{
				Content:         content,
				AcceptEula:      cmd.ServiceVersionFlag.AcceptEULA,
				TrustedProvider: cmd.ServiceVersionFlag.TrustedProvider,
			},
		}
	}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := namespace.NewManager(c)
	return m.CreateSupervisorService(ctx, &service)
}
