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
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/vmware/govmomi"
)

type cli struct {
	flag.FlagSet
	url string
}

var commands = map[string]Command{}

var Client *govmomi.Client

type Command interface {
	Parse([]string) error
	Run() error
}

func name(c Command) string {
	t := reflect.TypeOf(c).Elem()
	base := filepath.Base(t.PkgPath())
	if base == t.Name() {
		return t.Name()
	}
	return fmt.Sprintf("%s.%s", base, t.Name())
}

func Register(c Command) {
	commands[name(c)] = c
}

func (c *cli) Parse(args []string) error {
	c.StringVar(&c.url, "u", os.Getenv("GOVMOMI_URL"), "ESX or vCenter URL")

	if err := c.FlagSet.Parse(args); err != nil {
		return err
	}

	if c.url == "" {
		return flag.ErrHelp
	}

	u, err := url.Parse(c.url)
	if err != nil {
		return err
	}

	Client, err = govmomi.NewClient(*u)

	return err
}

func (c *cli) help() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	c.FlagSet.PrintDefaults()
	cmds := []string{}
	for name := range commands {
		cmds = append(cmds, name)
	}
	fmt.Fprintf(os.Stderr, "%s\n", strings.Join(cmds, "|"))
}

func Run(args []string) error {
	c := &cli{}
	c.Usage = c.help

	if err := c.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return nil
		}
		return err
	}

	args = c.Args()
	if len(args) == 0 {
		c.Usage()
		return nil
	}

	if cmd, ok := commands[args[0]]; ok {
		if err := cmd.Parse(args[1:]); err != nil {
			return err
		}
		return cmd.Run()
	}

	return nil
}
