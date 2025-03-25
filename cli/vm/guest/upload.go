// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"
	"io"
	"os"
	"path/filepath"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/vim25/soap"
)

type upload struct {
	*GuestFlag
	*FileAttrFlag

	overwrite bool
}

func init() {
	cli.Register("guest.upload", &upload{})
}

func (cmd *upload) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.GuestFlag, ctx = newGuestFlag(ctx)
	cmd.GuestFlag.Register(ctx, f)
	cmd.FileAttrFlag, ctx = newFileAttrFlag(ctx)
	cmd.FileAttrFlag.Register(ctx, f)

	f.BoolVar(&cmd.overwrite, "f", false, "If set, the guest destination file is clobbered")
}

func (cmd *upload) Usage() string {
	return "SOURCE DEST"
}

func (cmd *upload) Description() string {
	return `Copy SOURCE from the local system to DEST in the guest VM.

If SOURCE name is "-", read source from stdin.

Examples:
  govc guest.upload -l user:pass -vm=my-vm ~/.ssh/id_rsa.pub /home/$USER/.ssh/authorized_keys
  cowsay "have a great day" | govc guest.upload -l user:pass -vm=my-vm - /etc/motd
  tar -cf- foo/ | govc guest.run -d - tar -C /tmp -xf- # upload a directory`
}

func (cmd *upload) Process(ctx context.Context) error {
	if err := cmd.GuestFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.FileAttrFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *upload) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	c, err := cmd.Toolbox(ctx)
	if err != nil {
		return err
	}

	src := f.Arg(0)
	dst := f.Arg(1)

	p := soap.DefaultUpload

	var r io.Reader = os.Stdin

	if src != "-" {
		f, err := os.Open(filepath.Clean(src))
		if err != nil {
			return err
		}
		defer f.Close()

		r = f

		if cmd.OutputFlag.TTY {
			logger := cmd.ProgressLogger("Uploading... ")
			p.Progress = logger
			defer logger.Wait()
		}
	}

	return c.Upload(ctx, r, dst, p, cmd.Attr(), cmd.overwrite)
}
