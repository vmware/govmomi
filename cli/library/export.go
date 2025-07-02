// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vim25/soap"
)

type export struct {
	*flags.ClientFlag
	*flags.OutputFlag
	library.Item
}

func init() {
	cli.Register("library.export", &export{})
}

func (cmd *export) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *export) Usage() string {
	return "PATH [DEST]"
}

func (cmd *export) Description() string {
	return `Export library items.

If the given PATH is a library item, all files will be downloaded.

If the given PATH is a library item file, only that file will be downloaded.

By default, files are saved using the library item's file names to the current directory.
If DEST is given for a library item, files are saved there instead of the current directory.
If DEST is given for a library item file, the file will be saved with that name.
If DEST is '-', the file contents are written to stdout instead of saving to a file.

Examples:
  govc library.export library_name/item_name
  govc library.export library_name/item_name/file_name
  govc library.export library_name/item_name/*.ovf -`
}

func (cmd *export) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *export) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}
	cmd.KeepAlive(c)

	var names []string
	m := library.NewManager(c)
	res, err := flags.ContentLibraryResult(ctx, c, "", f.Arg(0))
	if err != nil {
		return err
	}

	switch t := res.GetResult().(type) {
	case library.Item:
		cmd.Item = t
	case library.File:
		names = []string{t.Name}
		cmd.Item = res.GetParent().GetResult().(library.Item)
	default:
		return fmt.Errorf("%q is a %T", f.Arg(0), t)
	}

	dst := f.Arg(1)
	one := len(names) == 1
	var log io.Writer = os.Stdout
	isStdout := one && dst == "-"
	if isStdout {
		log = io.Discard
	}

	session, err := m.CreateLibraryItemDownloadSession(ctx, library.Session{
		LibraryItemID: cmd.ID,
	})
	if err != nil {
		return err
	}

	if len(names) == 0 {
		files, err := m.ListLibraryItemDownloadSessionFile(ctx, session)
		if err != nil {
			return err
		}

		for _, file := range files {
			names = append(names, file.Name)
		}
	}

	for _, name := range names {
		_, err = m.PrepareLibraryItemDownloadSessionFile(ctx, session, name)
		if err != nil {
			return err
		}
	}

	download := func(src *url.URL, name string) error {
		p := soap.DefaultDownload

		if isStdout {
			s, _, err := c.Download(ctx, src, &p)
			if err != nil {
				return err
			}
			_, err = io.Copy(os.Stdout, s)
			_ = s.Close()
			return err
		}

		if cmd.OutputFlag.TTY {
			logger := cmd.ProgressLogger(fmt.Sprintf("Downloading %s... ", src.String()))
			defer logger.Wait()
			p.Progress = logger
		}

		if one && dst != "" {
			name = dst
		} else {
			name = filepath.Join(dst, name)
		}
		return c.DownloadFile(ctx, name, src, &p)
	}

	for _, name := range names {
		var info *library.DownloadFile
		_, _ = fmt.Fprintf(log, "Checking %s... ", name)
		for {
			info, err = m.GetLibraryItemDownloadSessionFile(ctx, session, name)
			if err != nil {
				return err
			}
			if info.Status == "PREPARED" {
				_, _ = fmt.Fprintln(log, info.Status)
				break // with this status we have a DownloadEndpoint.URI
			}
			time.Sleep(time.Second)
		}

		src, err := url.Parse(info.DownloadEndpoint.URI)
		if err != nil {
			return err
		}
		err = download(src, name)
		if err != nil {
			return err
		}
	}

	return nil
}
