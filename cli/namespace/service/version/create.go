// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package version

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"os"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/cli/namespace/service"
	"github.com/vmware/govmomi/vapi/namespace"
)

type create struct {
	*flags.ClientFlag
	*service.ServiceVersionFlag
}

func init() {
	cli.Register("namespace.service.version.create", &create{})
}

func (cmd *create) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.ServiceVersionFlag = &service.ServiceVersionFlag{}
	cmd.ServiceVersionFlag.Register(ctx, f)
}

func (cmd *create) Process(ctx context.Context) error {
	if err := cmd.ServiceVersionFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.ClientFlag.Process(ctx)
}

func (cmd *create) Description() string {
	return `Registers a new version for a registered vSphere Supervisor Service.

Examples:
  govc namespace.service.create my-service manifest-2.0.0.yaml`
}

func (cmd *create) Usage() string {
	return "SERVICE MANIFEST"
}

func (cmd *create) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	id := f.Arg(0)
	manifestFile := f.Arg(1)
	manifest, err := os.ReadFile(manifestFile)
	if err != nil {
		return fmt.Errorf("failed to read manifest file: %s", err)
	}

	content := base64.StdEncoding.EncodeToString(manifest)
	serviceVersion := namespace.SupervisorServiceVersion{}
	if cmd.ServiceVersionFlag.SpecType == "carvel" {
		serviceVersion.CarvelService = &namespace.CarvelVersionCreateSpec{
			Content: content,
		}
	} else {
		serviceVersion.VsphereService = &namespace.SupervisorServicesVSphereVersionCreateSpec{
			Content:         content,
			AcceptEula:      cmd.ServiceVersionFlag.AcceptEULA,
			TrustedProvider: cmd.ServiceVersionFlag.TrustedProvider,
		}
	}
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := namespace.NewManager(c)
	return m.CreateSupervisorServiceVersion(ctx, id, &serviceVersion)
}
