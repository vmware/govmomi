// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package license

import (
	"context"
	"flag"
	"fmt"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/license"
	"github.com/vmware/govmomi/vim25/types"
)

type add struct {
	*flags.ClientFlag
	*flags.OutputFlag
}

func init() {
	cli.Register("license.add", &add{})
}

func (cmd *add) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *add) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *add) Usage() string {
	return "KEY..."
}

func (cmd *add) Run(ctx context.Context, f *flag.FlagSet) error {
	client, err := cmd.Client()
	if err != nil {
		return err
	}

	m := license.NewManager(client)

	// From the vSphere 5.5 documentation:
	//
	//     To specify the edition type and any optional functions, use
	//     updateLicense for ESX Server and addLicense follow by
	//     LicenseAssingmentManager.updateAssignedLicense for VirtualCenter.
	//
	var addFunc func(ctx context.Context, key string, labels map[string]string) (types.LicenseManagerLicenseInfo, error)
	switch t := client.ServiceContent.About.ApiType; t {
	case "HostAgent":
		addFunc = m.Update
	case "VirtualCenter":
		addFunc = m.Add
	default:
		return fmt.Errorf("unsupported ApiType: %s", t)
	}

	result := make(licenseOutput, 0)
	for _, v := range f.Args() {
		license, err := addFunc(ctx, v, nil)
		if err != nil {
			return err
		}

		result = append(result, license)
	}

	return cmd.WriteResult(licenseOutput(result))
}
