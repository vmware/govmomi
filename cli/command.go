// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/vim25/types"
)

type HasFlags interface {
	// Register may be called more than once and should be idempotent.
	Register(ctx context.Context, f *flag.FlagSet)

	// Process may be called more than once and should be idempotent.
	Process(ctx context.Context) error
}

type Command interface {
	HasFlags

	Run(ctx context.Context, f *flag.FlagSet) error
}

func generalHelp(w io.Writer, filter string) {
	var cmds, matches []string
	for name := range commands {
		cmds = append(cmds, name)

		if filter != "" && strings.Contains(name, filter) {
			matches = append(matches, name)
		}
	}

	if len(matches) == 0 {
		fmt.Fprintf(w, `Usage: %[1]s <COMMAND> [COMMON OPTIONS] [PATH]...

govmomi is a Go library for interacting with VMware vSphere APIs (ESXi and/or
vCenter Server).
It is licensed under the Apache License, Version 2.0

%[1]s is the CLI for govmomi.

The available commands are listed below. A detailed description of each
command can be displayed with "govc <COMMAND> -h". The description of all
commands can be also found at https://github.com/vmware/govmomi/blob/main/govc/USAGE.md.

Examples:
  show usage of a command:       govc <COMMAND> -h
  show toplevel structure:       govc ls
  show datacenter summary:       govc datacenter.info
  show all VMs:                  govc find -type m
  upload a ISO file:             govc datastore.upload -ds datastore1 ./config.iso vm-name/config.iso

Common options:
  -h                        Show this message
  -cert=                    Certificate [GOVC_CERTIFICATE]
  -debug=false              Store debug logs [GOVC_DEBUG]
  -trace=false              Write SOAP/REST traffic to stderr
  -verbose=false            Write request/response data to stderr
  -dump=false               Enable output dump
  -json=false               Enable JSON output
  -xml=false                Enable XML output
  -k=false                  Skip verification of server certificate [GOVC_INSECURE]
  -key=                     Private key [GOVC_PRIVATE_KEY]
  -persist-session=true     Persist session to disk [GOVC_PERSIST_SESSION]
  -tls-ca-certs=            TLS CA certificates file [GOVC_TLS_CA_CERTS]
  -tls-known-hosts=         TLS known hosts file [GOVC_TLS_KNOWN_HOSTS]
  -u=                       ESX or vCenter URL [GOVC_URL]
  -vim-namespace=urn:vim25  Vim namespace [GOVC_VIM_NAMESPACE]
  -vim-version=6.0          Vim version [GOVC_VIM_VERSION]
  -dc=                      Datacenter [GOVC_DATACENTER]
  -host.dns=                Find host by FQDN
  -host.ip=                 Find host by IP address
  -host.ipath=              Find host by inventory path
  -host.uuid=               Find host by UUID
  -vm.dns=                  Find VM by FQDN
  -vm.ip=                   Find VM by IP address
  -vm.ipath=                Find VM by inventory path
  -vm.path=                 Find VM by path to .vmx file
  -vm.uuid=                 Find VM by UUID

Available commands:
`, filepath.Base(os.Args[0]))

	} else {
		fmt.Fprintf(w, "%s: command '%s' not found, did you mean:\n", os.Args[0], filter)
		cmds = matches
	}

	sort.Strings(cmds)
	for _, name := range cmds {
		fmt.Fprintf(w, "  %s\n", name)
	}
}

func commandHelp(w io.Writer, name string, cmd Command, f *flag.FlagSet) {
	type HasUsage interface {
		Usage() string
	}

	fmt.Fprintf(w, "Usage: %s %s [OPTIONS]", os.Args[0], name)
	if u, ok := cmd.(HasUsage); ok {
		fmt.Fprintf(w, " %s", u.Usage())
	}
	fmt.Fprintf(w, "\n")

	type HasDescription interface {
		Description() string
	}

	if u, ok := cmd.(HasDescription); ok {
		fmt.Fprintf(w, "\n%s\n", u.Description())
	}

	n := 0
	f.VisitAll(func(_ *flag.Flag) {
		n += 1
	})

	if n > 0 {
		fmt.Fprintf(w, "\nOptions:\n")
		tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)
		f.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(tw, "\t-%s=%s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		tw.Flush()
	}
}

func clientLogout(ctx context.Context, cmd Command) error {
	type logout interface {
		Logout(context.Context) error
	}

	if l, ok := cmd.(logout); ok {
		return l.Logout(ctx)
	}

	return nil
}

func Run(args []string) int {
	hw := os.Stderr
	rc := 1
	hwrc := func(arg string) {
		arg = strings.TrimLeft(arg, "-")
		if arg == "h" || arg == "help" {
			hw = os.Stdout
			rc = 0
		}
	}

	var err error

	if len(args) == 0 {
		generalHelp(hw, "")
		return rc
	}

	// Look up real command name in aliases table.
	name, ok := aliases[args[0]]
	if !ok {
		name = args[0]
	}

	cmd, ok := commands[name]
	if !ok {
		hwrc(name)
		generalHelp(hw, name)
		return rc
	}

	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	ctx := context.Background()

	if id := os.Getenv("GOVC_OPERATION_ID"); id != "" {
		ctx = context.WithValue(ctx, types.ID{}, id)
	}

	cmd.Register(ctx, fs)

	if err = fs.Parse(args[1:]); err != nil {
		goto error
	}

	if err = cmd.Process(ctx); err != nil {
		goto error
	}

	if err = cmd.Run(ctx, fs); err != nil {
		goto error
	}

	if err = clientLogout(ctx, cmd); err != nil {
		goto error
	}

	return 0

error:
	if err == flag.ErrHelp {
		if len(args) == 2 {
			hwrc(args[1])
		}
		commandHelp(hw, args[0], cmd, fs)
	} else {
		if x, ok := err.(interface{ ExitCode() int }); ok {
			// propagate exit code, e.g. from guest.run
			rc = x.ExitCode()
		} else {
			w, ok := cmd.(interface{ WriteError(error) bool })
			if ok {
				ok = w.WriteError(err)
			}
			if !ok {
				fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
			}
		}
	}

	_ = clientLogout(ctx, cmd)

	return rc
}
