// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package session

import (
	"context"
	"flag"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/vapi/library"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag

	files bool
}

func init() {
	cli.Register("library.session.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.files, "i", false, "List session item files (with -json only)")
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *ls) Description() string {
	return `List library item update sessions.

Examples:
  govc library.session.ls
  govc library.session.ls -json | jq .`
}

type librarySession struct {
	*library.Session
	LibraryItemPath string `json:"library_item_path"`
}

type info struct {
	Sessions []librarySession `json:"sessions"`
	Files    map[string]any   `json:"files"`
	kind     string
}

func (i *info) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
	_, _ = fmt.Fprintln(tw, "ID\tItem\tType\tVersion\tProgress\tState\tExpires")

	for _, s := range i.Sessions {
		_, _ = fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%d\t%s\t%s\n",
			s.ID, s.LibraryItemPath, i.kind, s.LibraryItemContentVersion, s.ClientProgress, s.State,
			s.ExpirationTime.Format("2006-01-02 15:04"))
	}

	return tw.Flush()
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := library.NewManager(c)

	kinds := []struct {
		kind string
		list func(context.Context) ([]string, error)
		get  func(context.Context, string) (*library.Session, error)
	}{
		{"Update", m.ListLibraryItemUpdateSession, m.GetLibraryItemUpdateSession},
		{"Download", m.ListLibraryItemDownloadSession, m.GetLibraryItemDownloadSession},
	}

	for _, k := range kinds {
		ids, err := k.list(ctx)
		if err != nil {
			return err
		}
		if len(ids) == 0 {
			continue
		}
		i := &info{
			Files: make(map[string]any),
			kind:  k.kind,
		}

		for _, id := range ids {
			session, err := k.get(ctx, id)
			if err != nil {
				return err
			}
			var path string
			item, err := m.GetLibraryItem(ctx, session.LibraryItemID)
			if err == nil {
				// can only show library path if item exists
				lib, err := m.GetLibraryByID(ctx, item.LibraryID)
				if err != nil {
					return err
				}
				path = fmt.Sprintf("/%s/%s", lib.Name, item.Name)
			}
			i.Sessions = append(i.Sessions, librarySession{session, path})
			if !cmd.files {
				continue
			}
			if k.kind == "Update" {
				f, err := m.ListLibraryItemUpdateSessionFile(ctx, id)
				if err != nil {
					return err
				}
				i.Files[id] = f
			} else {
				f, err := m.ListLibraryItemDownloadSessionFile(ctx, id)
				if err != nil {
					return err
				}
				i.Files[id] = f
			}
		}

		err = cmd.WriteResult(i)
		if err != nil {
			return err
		}
	}
	return nil
}
