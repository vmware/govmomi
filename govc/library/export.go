/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.

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

package library

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/library/finder"
	"github.com/vmware/govmomi/vapi/rest"
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
	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		var names []string
		m := library.NewManager(c)
		res, err := finder.NewFinder(m).Find(ctx, f.Arg(0))
		if err != nil {
			return err
		}

		if len(res) != 1 {
			return fmt.Errorf("%q matches %d items", f.Arg(0), len(res))
		}

		switch t := res[0].GetResult().(type) {
		case library.Item:
			cmd.Item = t
		case library.File:
			names = []string{t.Name}
			cmd.Item = res[0].GetParent().GetResult().(library.Item)
		default:
			return fmt.Errorf("%q is a %T", f.Arg(0), t)
		}

		dst := f.Arg(1)
		one := len(names) == 1
		var log io.Writer = os.Stdout
		isStdout := one && dst == "-"
		if isStdout {
			log = ioutil.Discard
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
			var info *library.DownloadFileInfo
			_, _ = fmt.Fprintf(log, "Checking %s... ", name)
			for {
				info, err = m.GetLibraryItemDownloadSessionFile(ctx, session, name)
				if err != nil {
					return err
				}
				if info.Status == "ERROR" {
					_, _ = fmt.Fprintln(log, info.Status)
					return fmt.Errorf("preparing %s: %v", name, info.ErrorMessage)
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
	})
}
