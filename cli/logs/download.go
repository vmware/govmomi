// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package logs

import (
	"context"
	"flag"
	"fmt"
	"path"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type download struct {
	*flags.DatacenterFlag

	IncludeDefault bool
}

func init() {
	cli.Register("logs.download", &download{})
}

func (cmd *download) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)

	f.BoolVar(&cmd.IncludeDefault, "default", false, "Specifies if the bundle should include the default server")
}

func (cmd *download) Process(ctx context.Context) error {
	if err := cmd.DatacenterFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *download) Usage() string {
	return "[PATH]..."
}

func (cmd *download) Description() string {
	return `Generate diagnostic bundles.

A diagnostic bundle includes log files and other configuration information.

Use PATH to include a specific set of hosts to include.

Examples:
  govc logs.download
  govc logs.download host-a host-b`
}

func (cmd *download) DownloadFile(c *vim25.Client, b string) error {
	u, err := c.Client.ParseURL(b)
	if err != nil {
		return err
	}

	dst := path.Base(u.Path)
	p := soap.DefaultDownload
	if cmd.OutputFlag.TTY {
		logger := cmd.ProgressLogger(fmt.Sprintf("Downloading %s... ", dst))
		defer logger.Wait()
		p.Progress = logger
	}

	return c.Client.DownloadFile(context.Background(), dst, u, &p)
}

func (cmd *download) GenerateLogBundles(m *object.DiagnosticManager, host []*object.HostSystem) ([]types.DiagnosticManagerBundleInfo, error) {
	ctx := context.TODO()
	logger := cmd.ProgressLogger("Generating log bundles... ")
	defer logger.Wait()

	task, err := m.GenerateLogBundles(ctx, cmd.IncludeDefault, host)
	if err != nil {
		return nil, err
	}

	r, err := task.WaitForResult(ctx, logger)
	if err != nil {
		return nil, err
	}

	return r.Result.(types.ArrayOfDiagnosticManagerBundleInfo).DiagnosticManagerBundleInfo, nil
}

func (cmd *download) Run(ctx context.Context, f *flag.FlagSet) error {
	finder, err := cmd.Finder()
	if err != nil {
		return err
	}

	var host []*object.HostSystem

	for _, arg := range f.Args() {
		hs, err := finder.HostSystemList(ctx, arg)
		if err != nil {
			return err
		}

		host = append(host, hs...)
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m := object.NewDiagnosticManager(c)

	bundles, err := cmd.GenerateLogBundles(m, host)
	if err != nil {
		return err
	}

	for _, bundle := range bundles {
		err := cmd.DownloadFile(c, bundle.Url)
		if err != nil {
			return err
		}
	}

	return nil
}
