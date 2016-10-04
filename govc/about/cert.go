/*
Copyright (c) 2016 VMware, Inc. All Rights Reserved.

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

package about

import (
	"context"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type cert struct {
	*flags.ClientFlag
	*flags.OutputFlag

	show       bool
	thumbprint bool
}

func init() {
	cli.Register("about.cert", &cert{})
}

func (cmd *cert) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.BoolVar(&cmd.show, "show", false, "Show PEM encoded server certificate only")
	f.BoolVar(&cmd.thumbprint, "thumbprint", false, "Output host hash and thumbprint only")
}

func (cmd *cert) Description() string {
	return `Display SSL certificate info for HOST.

Examples:
  govc about.cert -k -json | jq -r .ThumbprintSHA1
  govc about.cert -k -show | sudo tee /usr/local/share/ca-certificates/host.crt
  govc about.cert -k -thumbprint | tee -a ~/.govmomi/known_hosts`
}

func (cmd *cert) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

type certResult struct {
	cmd  *cert
	err  error
	peer *x509.Certificate
}

func (r *certResult) Write(w io.Writer) error {
	if r.cmd.show {
		return pem.Encode(w, &pem.Block{Type: "CERTIFICATE", Bytes: r.peer.Raw})
	}

	info := new(object.HostCertificateInfo).FromCertificate(r.peer)

	if r.cmd.thumbprint {
		u := r.cmd.URLWithoutPassword()
		_, err := fmt.Fprintf(w, "%x %s\n", sha1.Sum([]byte(u.Host)), info.ThumbprintSHA1)
		return err
	}

	info.Status = string(types.HostCertificateManagerCertificateInfoCertificateStatusGood)
	if r.err != nil {
		info.Status = fmt.Sprintf("ERROR %s", r.err)
	}

	return r.cmd.WriteResult(info)
}

func (cmd *cert) Run(ctx context.Context, f *flag.FlagSet) error {
	u := cmd.URLWithoutPassword()
	c := soap.NewClient(u, false)

	if err := cmd.SetRootCAs(c); err != nil {
		return err
	}

	r := certResult{cmd: cmd}

	c.SetDialTLS(func(err error, state tls.ConnectionState) error {
		r.peer = state.PeerCertificates[0]
		r.err = err
		return nil
	})

	_, err := vim25.NewClient(ctx, c) // invoke DialTLS()
	if err != nil && err != r.err {
		return err
	}

	if r.err != nil && r.cmd.IsSecure() {
		cmd.Out = os.Stderr
		// using same exit code as curl:
		defer os.Exit(60)
	}

	return r.Write(cmd.Out)
}
