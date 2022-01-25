/*
Copyright (c) 2022 VMware, Inc. All Rights Reserved.

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

package policy

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/appliance/update/policy"
)

type set struct {
	*flags.ClientFlag
	checkCertificate bool
	customURL        string
	username         string
	password         string
	autostage        bool
	minute           int
	hour             int
}

func init() {
	cli.Register("vcsa.update.policy.set", &set{})
}

func (cmd *set) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.BoolVar(&cmd.autostage,
		"auto_stage",
		false,
		"Automatically stage the latest update if available.")

	f.BoolVar(&cmd.checkCertificate,
		"certificate_check",
		true,
		"Indicates whether certificates will be checked during patching. "+
			"Warning: Setting this field to false will result in an insecure connection to update repository "+
			"which can potentially put the appliance security at risk.")

	f.StringVar(&cmd.customURL,
		"custom_URL",
		"",
		"Current appliance update repository URL. If unset then default URL is assumed.")

	f.StringVar(&cmd.username,
		"username",
		"",
		"Username for the update repository If unset username will not be used to login.")

	f.StringVar(&cmd.password,
		"password",
		"",
		"Password for the update repository password If unset password will not be used to login.")

	f.IntVar(&cmd.hour,
		"hour",
		0,
		"Hour 0-24")

	f.IntVar(&cmd.minute,
		"minute",
		0,
		"Minute 0-59")

}

func (cmd *set) Description() string {
	return `Sets the automatic update checking and staging policy.
Examples:
  govc vcsa.update.policy.set`
}

func (cmd *set) Usage() string {
	return "[DAY]"
}

func (cmd *set) Run(ctx context.Context, f *flag.FlagSet) error {
	if f.NArg() != 1 {
		return flag.ErrHelp
	}
	var config policy.Config

	config.AutoStage = cmd.autostage
	config.CertificateCheck = cmd.checkCertificate
	config.CustomURL = cmd.customURL
	config.Username = cmd.username
	config.Password = cmd.password
	config.CheckSchedule = []policy.CheckSchedule{{
		Day:    f.Arg(0),
		Hour:   cmd.hour,
		Minute: cmd.minute,
	}}

	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	m := policy.NewManager(c)

	err = m.Set(ctx, config)
	if err != nil {
		return err
	}

	return nil
}
