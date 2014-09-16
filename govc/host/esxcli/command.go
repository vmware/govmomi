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

package esxcli

import (
	"flag"
	"fmt"
	"strings"

	"github.com/vmware/govmomi/vim25/types"
)

type Command struct {
	name []string
	args []string
}

func NewCommand(args []string) *Command {
	c := &Command{}

	for i, arg := range args {
		if strings.HasPrefix(arg, "-") {
			c.args = args[i:]
			break
		} else {
			c.name = append(c.name, arg)
		}
	}

	return c
}

func (c *Command) Namespace() string {
	return strings.Join(c.name[:len(c.name)-1], ".")
}

func (c *Command) Name() string {
	return c.name[len(c.name)-1]
}

func (c *Command) Method() string {
	return "vim.EsxCLI." + strings.Join(c.name, ".")
}

func (c *Command) Moid() string {
	return "ha-cli-handler-" + strings.Join(c.name[:len(c.name)-1], "-")
}

// Parse generates a flag.FlagSet based on the given []types.DynamicTypeMgrParamTypeInfo and
// returns arguments for use with methods.ExecuteSoap
func (c *Command) Parse(params []types.DynamicTypeMgrParamTypeInfo) ([]types.ReflectManagedMethodExecuterSoapArgument, error) {
	flags := flag.NewFlagSet(c.Method(), flag.ExitOnError)
	vals := make([]string, len(params))

	for i, p := range params {
		v := &vals[i]
		flags.StringVar(v, p.Name, "", p.Name)
		flags.StringVar(v, p.Name[:1], "", p.Name)
	}

	err := flags.Parse(c.args)
	if err != nil {
		return nil, err
	}

	args := []types.ReflectManagedMethodExecuterSoapArgument{}

	for i, p := range params {
		if vals[i] == "" {
			continue
		}
		key := p.Name
		val := vals[i]
		args = append(args, types.ReflectManagedMethodExecuterSoapArgument{
			Name: key,
			Val:  fmt.Sprintf("<%s>%s</%s>", key, val, key),
		})
	}

	return args, nil
}
