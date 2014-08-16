/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package cli

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"text/tabwriter"
)

type HasFlags interface {
	Register(f *flag.FlagSet)
	Process() error
}

type Command interface {
	HasFlags

	Run(f *flag.FlagSet) error
}

type NoFlags struct{}

func (n *NoFlags) Register(f *flag.FlagSet) {}

func (n *NoFlags) Process() error { return nil }

var hasFlagsType = reflect.TypeOf((*HasFlags)(nil)).Elem()

func RegisterCommand(h HasFlags, f *flag.FlagSet) {
	Walk(h, hasFlagsType, func(v interface{}) error {
		v.(HasFlags).Register(f)
		return nil
	})
}

func ProcessCommand(h HasFlags) error {
	err := Walk(h, hasFlagsType, func(v interface{}) error {
		err := v.(HasFlags).Process()
		return err
	})
	return err
}

func generalHelp() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

	cmds := []string{}
	for name := range commands {
		cmds = append(cmds, name)
	}

	sort.Strings(cmds)

	for _, name := range cmds {
		fmt.Fprintf(os.Stderr, "  %s\n", name)
	}
}

func commandHelp(name string, f *flag.FlagSet) {
	fmt.Fprintf(os.Stderr, "Usage of %s %s:\n", os.Args[0], name)

	n := 0
	f.VisitAll(func(_ *flag.Flag) {
		n += 1
	})

	if n == 0 {
		fmt.Fprintf(os.Stderr, "  (no flags)\n")
	} else {
		tw := tabwriter.NewWriter(os.Stderr, 2, 0, 2, ' ', 0)

		type IsBoolFlagger interface {
			IsBoolFlag() bool
		}

		f.VisitAll(func(f *flag.Flag) {
			if b, ok := f.Value.(IsBoolFlagger); ok && b.IsBoolFlag() {
				fmt.Fprintf(tw, "\t-%s\t%s\n", f.Name, f.Usage)
				return
			}

			fmt.Fprintf(tw, "\t-%s=%s\t%s\n", f.Name, f.DefValue, f.Usage)
		})

		tw.Flush()
	}
}

func Run(args []string) int {
	if len(args) == 0 {
		generalHelp()
		return 1
	}

	cmd, ok := commands[args[0]]
	if !ok {
		generalHelp()
		return 1
	}

	f := flag.NewFlagSet("", flag.ContinueOnError)
	f.SetOutput(ioutil.Discard)

	RegisterCommand(cmd, f)

	if err := f.Parse(args[1:]); err != nil {
		if err == flag.ErrHelp {
			commandHelp(args[0], f)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
		return 1
	}

	if err := ProcessCommand(cmd); err != nil {
		if err == flag.ErrHelp {
			commandHelp(args[0], f)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
		return 1
	}

	if err := cmd.Run(f); err != nil {
		if err == flag.ErrHelp {
			commandHelp(args[0], f)
		} else {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
		return 1
	}

	return 0
}
