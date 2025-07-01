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
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag

	r bool
	s bool
}

func init() {
	cli.Register("session.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.r, "r", false, "List cached REST session (if any)")
	f.BoolVar(&cmd.s, "S", false, "List current SOAP session")
}

func (cmd *ls) Description() string {
	return `List active sessions.

All SOAP sessions are listed by default. The '-S' flag will limit this list to the current session.

Examples:
  govc session.ls
  govc session.ls -json | jq -r .currentSession.key`
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

type sessionInfo struct {
	cmd *ls
	now *time.Time
	mo.SessionManager
}

func (s *sessionInfo) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 4, 0, 2, ' ', 0)

	fmt.Fprintf(tw, "Key\t")
	fmt.Fprintf(tw, "Name\t")
	fmt.Fprintf(tw, "Created\t")
	fmt.Fprintf(tw, "Idle\t")
	fmt.Fprintf(tw, "Host\t")
	fmt.Fprintf(tw, "Agent\t")
	fmt.Fprintf(tw, "\t")
	fmt.Fprint(tw, "\n")

	for _, v := range s.SessionList {
		idle := "  ."
		if v.Key != s.CurrentSession.Key && v.IpAddress != "-" {
			since := s.now.Sub(v.LastActiveTime)
			if since > time.Hour {
				idle = "old"
			} else {
				idle = (time.Duration(since.Seconds()) * time.Second).String()
			}
		}
		fmt.Fprintf(tw, "%s\t", v.Key)
		fmt.Fprintf(tw, "%s\t", v.UserName)
		fmt.Fprintf(tw, "%s\t", v.LoginTime.Format("2006-01-02 15:04"))
		fmt.Fprintf(tw, "%s\t", idle)
		fmt.Fprintf(tw, "%s\t", v.IpAddress)
		fmt.Fprintf(tw, "%s\t", v.UserAgent)
		fmt.Fprint(tw, "\n")
	}

	return tw.Flush()
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	var m mo.SessionManager
	var props []string
	pc := property.DefaultCollector(c)
	if cmd.s {
		props = []string{"currentSession"}
	}

	err = pc.RetrieveOne(ctx, *c.ServiceContent.SessionManager, props, &m)
	if err != nil {
		return err
	}

	if cmd.s {
		m.SessionList = []types.UserSession{*m.CurrentSession}
	}

	now, err := methods.GetCurrentTime(ctx, c)
	if err != nil {
		return err
	}

	// The REST API doesn't include a way to get the complete list of sessions, only the current session.
	if cmd.r {
		rc := new(rest.Client)
		ok, _ := cmd.Session.Load(ctx, rc, cmd.ConfigureTLS)
		if ok {
			rs, err := rc.Session(ctx)
			if err != nil {
				return err
			}
			m.SessionList = append(m.SessionList, types.UserSession{
				Key:            rc.SessionID(),
				UserName:       rs.User + " (REST)",
				LoginTime:      rs.Created,
				LastActiveTime: rs.LastAccessed,
				IpAddress:      "-",
				UserAgent:      c.UserAgent,
			})
		}
	}

	return cmd.WriteResult(&sessionInfo{cmd, now, m})
}
