/*
Copyright (c) 2020 VMware, Inc. All Rights Reserved.

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

package cluster

import (
	"context"
	"flag"
	"io"
	"mime"
	"net/http"
	"os"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/namespace"
)

type logs struct {
	*flags.ClusterFlag
}

func init() {
	cli.Register("namespace.logs.download", &logs{})
}

func (cmd *logs) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClusterFlag, ctx = flags.NewClusterFlag(ctx)
	cmd.ClusterFlag.Register(ctx, f)
}

func (cmd *logs) Usage() string {
	return "[NAME]"
}

func (cmd *logs) Description() string {
	return `Download namespace cluster support bundle.

If NAME name is "-", bundle is written to stdout.

See also: govc logs.download

Examples:
  govc namespace.logs.download -cluster k8s
  govc namespace.logs.download -cluster k8s - | tar -xvf -
  govc namespace.logs.download -cluster k8s logs.tar`
}

func (cmd *logs) Run(ctx context.Context, f *flag.FlagSet) error {
	c, err := cmd.RestClient()
	if err != nil {
		return err
	}

	cluster, err := cmd.Cluster()
	if err != nil {
		return err
	}

	id := cluster.Reference().Value

	name := f.Arg(0)

	m := namespace.NewManager(c)

	bundle, err := m.CreateSupportBundle(ctx, id)
	if err != nil {
		return err
	}

	req, err := m.SupportBundleRequest(ctx, bundle)
	if err != nil {
		return err
	}

	return c.Client.Do(ctx, req, func(res *http.Response) error {
		if name == "" {
			d := res.Header.Get("Content-Disposition")
			_, params, err := mime.ParseMediaType(d)
			if err == nil {
				name = params["filename"]
			}
		}

		var w io.Writer

		if name == "-" {
			w = os.Stdout
		} else {
			f, err := os.Create(name)
			if err != nil {
				return err
			}
			defer f.Close()
			w = f
		}

		_, err = io.Copy(w, res.Body)
		return err
	})
}
