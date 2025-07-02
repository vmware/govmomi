// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

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

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/ovf/importer"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vim25/soap"
)

type item struct {
	*flags.ClientFlag
	*flags.OutputFlag
	library.Item
	library.Checksum

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
	f.StringVar(&cmd.Checksum.Checksum, "c", "", "Checksum value to verify the pulled library item")
	f.StringVar(&cmd.Checksum.Algorithm, "a", "SHA256", "Algorithm used to calculate the checksum. Possible values are: SHA1, MD5, SHA256 (default), SHA512")
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
  govc library.import library_id file.iso # Use library id if multiple libraries have the same name
  govc library.import library_name/item_name file.ova # update existing item
  govc library.import library_name http://example.com/file.ovf # download and push to vCenter
  govc library.import -pull library_name http://example.com/file.ova # direct pull from vCenter
  govc library.import -pull -c=<checksum> -a=<SHA1|MD5|SHA256|SHA512> library_name http://example.com/file.ova # direct pull from vCenter with checksum validation`
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

	// Checksums are verified after the file is uploaded to a server.
	// Check the algorithm and fail early if it's not supported.
	if cmd.pull && cmd.Checksum.Checksum != "" {
		switch cmd.Checksum.Algorithm {
		case "SHA1", "MD5", "SHA256", "SHA512":
		default:
			return fmt.Errorf("invalid checksum algorithm: %s", cmd.Checksum.Algorithm)
		}
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
	opener := importer.Opener{Client: client}
	var archive importer.Archive
	archive = &importer.FileArchive{Path: file, Opener: opener}

	manifest := make(map[string]*library.Checksum)
	if cmd.Name == "" {
		cmd.Name = strings.TrimSuffix(base, ext)
	}

	switch ext {
	case ".ova":
		archive = &importer.TapeArchive{Path: file, Opener: opener}
		base = "*.ovf"
		mf = "*.mf"
		kind = library.ItemTypeOVF
	case ".ovf":
		kind = library.ItemTypeOVF
	case ".iso":
		kind = library.ItemTypeISO
	}

	if cmd.Type == "" {
		cmd.Type = kind
	}

	if !cmd.pull && cmd.Type == library.ItemTypeOVF {
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

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}
	cmd.KeepAlive(c)

	m := library.NewManager(c)
	res, err := flags.ContentLibraryResult(ctx, c, "", f.Arg(0))
	if err != nil {
		return err
	}

	switch t := res.GetResult().(type) {
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
	if err != nil {
		return err
	}
	if cmd.pull {
		_, err = m.AddLibraryItemFileFromURI(ctx, session, filepath.Base(file), file, cmd.Checksum)
		if err != nil {
			return err
		}

		err = m.CompleteLibraryItemUpdateSession(ctx, session)
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

		if e, ok := f.(*importer.TapeArchiveEntry); ok {
			name = e.Name // expand path.Match's (e.g. "*.ovf" -> "name.ovf")
		}

		info := library.UpdateFile{
			Name:       name,
			SourceType: "PUSH",
			Checksum:   manifest[name],
			Size:       size,
		}

		update, err := m.AddLibraryItemFile(ctx, session, info)
		if err != nil {
			return err
		}

		p := soap.DefaultUpload
		p.ContentLength = size
		u, err := url.Parse(update.UploadEndpoint.URI)
		if err != nil {
			return err
		}
		if cmd.TTY {
			logger := cmd.ProgressLogger(fmt.Sprintf("Uploading %s... ", name))
			p.Progress = logger
			defer logger.Wait()
		}
		return c.Upload(ctx, f, u, &p)
	}

	if err = upload(base); err != nil {
		return err
	}

	if cmd.Type == library.ItemTypeOVF {
		o, err := importer.ReadOvf(base, archive)
		if err != nil {
			return err
		}

		e, err := importer.ReadEnvelope(o)
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
}
