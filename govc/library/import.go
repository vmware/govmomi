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
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/govc/importx"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/library/finder"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25/soap"
)

type item struct {
	*flags.ClientFlag
	*flags.OutputFlag
	library.Item

	manifest bool
	pull     bool
}

func init() {
	cli.Register("library.import", &item{})
}

func (cmd *item) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.StringVar(&cmd.Name, "n", "", "Library item name")
	f.StringVar(&cmd.Type, "t", "", "Library item type")
	f.BoolVar(&cmd.manifest, "m", false, "Require ova manifest")
	f.BoolVar(&cmd.pull, "pull", false, "Pull library item from http endpoint")
}

func (cmd *item) Usage() string {
	return "LIBRARY ITEM"
}

func (cmd *item) Description() string {
	return `Import library items.

Examples:
  govc library.import library_name file.ova
  govc library.import library_name file.ovf
  govc library.import library_name file.iso
  govc library.import library_name/item_name file.ova # update existing item
  govc library.import library_name http://example.com/file.ovf # download and push to vCenter
  govc library.import -pull library_name http://example.com/file.ova # direct pull from vCenter`
}

func (cmd *item) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *item) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 2 {
		return flag.ErrHelp
	}

	file := f.Arg(1)
	base := filepath.Base(file)
	ext := filepath.Ext(base)
	mf := strings.Replace(base, ext, ".mf", 1)
	kind := ""
	client, err := cmd.Client()
	if err != nil {
		return err
	}
	opener := importx.Opener{Client: client}
	archive := &importx.ArchiveFlag{Archive: &importx.FileArchive{Path: file, Opener: opener}}

	manifest := make(map[string]*library.Checksum)
	if cmd.Name == "" {
		cmd.Name = strings.TrimSuffix(base, ext)
	}

	switch ext {
	case ".ova":
		archive.Archive = &importx.TapeArchive{Path: file, Opener: opener}
		base = "*.ovf"
		mf = "*.mf"
		kind = "ovf"
	case ".ovf":
		kind = "ovf"
	case ".iso":
		kind = "iso"
	}

	if cmd.Type == "" {
		cmd.Type = kind
	}

	if !cmd.pull && cmd.Type == "ovf" {
		f, _, err := archive.Open(mf)
		if err == nil {
			sums, err := library.ReadManifest(f)
			_ = f.Close()
			if err != nil {
				return err
			}
			manifest = sums
		} else {
			msg := fmt.Sprintf("manifest %q: %s", mf, err)
			if cmd.manifest {
				return errors.New(msg)
			}
			fmt.Fprintln(os.Stderr, msg)
		}
	}

	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		m := library.NewManager(c)
		res, err := finder.NewFinder(m).Find(ctx, f.Arg(0))
		if err != nil {
			return err
		}

		if len(res) != 1 {
			return fmt.Errorf("%q matches %d items", f.Arg(0), len(res))
		}

		switch t := res[0].GetResult().(type) {
		case library.Library:
			cmd.LibraryID = t.ID
			cmd.ID, err = m.CreateLibraryItem(ctx, cmd.Item)
			if err != nil {
				return err
			}
		case library.Item:
			cmd.Item = t
		default:
			return fmt.Errorf("%q is a %T", f.Arg(0), t)
		}

		session, err := m.CreateLibraryItemUpdateSession(ctx, library.Session{
			LibraryItemID: cmd.ID,
		})

		if cmd.pull {
			_, err = m.AddLibraryItemFileFromURI(ctx, session, filepath.Base(file), file)
			if err != nil {
				return err
			}

			return m.WaitOnLibraryItemUpdateSession(ctx, session, 3*time.Second, nil)
		}

		upload := func(name string) error {
			f, size, err := archive.Open(name)
			if err != nil {
				return err
			}
			defer f.Close()

			if e, ok := f.(*importx.TapeArchiveEntry); ok {
				name = e.Name // expand path.Match's (e.g. "*.ovf" -> "name.ovf")
			}

			info := library.UpdateFile{
				Name:       name,
				SourceType: "PUSH",
				Checksum:   manifest[name],
				Size:       &size,
			}

			update, err := m.AddLibraryItemFile(ctx, session, info)
			if err != nil {
				return err
			}

			p := soap.DefaultUpload
			p.Headers = map[string]string{
				"vmware-api-session-id": session,
			}
			p.ContentLength = size
			u, err := url.Parse(update.UploadEndpoint.URI)
			if err != nil {
				return err
			}
			if cmd.OutputFlag.TTY {
				logger := cmd.ProgressLogger(fmt.Sprintf("Uploading %s... ", name))
				p.Progress = logger
				defer logger.Wait()
			}
			return c.Upload(ctx, f, u, &p)
		}

		if err = upload(base); err != nil {
			return err
		}

		if cmd.Type == "ovf" {
			o, err := archive.ReadOvf(base)
			if err != nil {
				return err
			}

			e, err := archive.ReadEnvelope(o)
			if err != nil {
				return fmt.Errorf("failed to parse ovf: %s", err)
			}

			for i := range e.References {
				if err = upload(e.References[i].Href); err != nil {
					return err
				}
			}
		}

		return m.CompleteLibraryItemUpdateSession(ctx, session)
	})
}
