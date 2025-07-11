// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"
	"io"
	"os"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/progress"
)

type download struct {
	*GuestFlag

	overwrite bool
}

func init() {
	cli.Register("guest.download", &download{})
}

func (cmd *download) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.GuestFlag, ctx = newGuestFlag(ctx)
	cmd.GuestFlag.Register(ctx, f)

	f.BoolVar(&cmd.overwrite, "f", false, "If set, the local destination file is clobbered")
}

func (cmd *download) Usage() string {
	return "SOURCE DEST"
}

func (cmd *download) Description() string {
	return `Copy SOURCE from the guest VM to DEST on the local system.

If DEST name is "-", source is written to stdout.

Examples:
  govc guest.download -l user:pass -vm=my-vm /var/log/my.log ./local.log
  govc guest.download -l user:pass -vm=my-vm /etc/motd -
  tar -cf- foo/ | govc guest.run -d - tar -C /tmp -xf-
  govc guest.run tar -C /tmp -cf- foo/ | tar -C /tmp -xf- # download directory`
}

func (cmd *download) Process(ctx context.Context) error {
	if err := cmd.GuestFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *download) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	src := f.Arg(0)
	dst := f.Arg(1)

	_, err := os.Stat(dst)
	if err == nil && !cmd.overwrite {
		return os.ErrExist
	}

	c, err := cmd.Toolbox(ctx)
	if err != nil {
		return err
	}

	s, n, err := c.Download(ctx, src)
	if err != nil {
		return err
	}

	if dst == "-" {
		_, err = io.Copy(os.Stdout, s)
		return err
	}

	var p progress.Sinker

	if cmd.OutputFlag.TTY {
		logger := cmd.ProgressLogger("Downloading... ")
		p = logger
		defer logger.Wait()
	}

	return c.ProcessManager.Client().WriteFile(ctx, dst, s, n, p, nil)
}
