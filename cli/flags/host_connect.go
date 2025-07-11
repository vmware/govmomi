// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"context"
	"flag"
	"fmt"
	"net/url"

	"github.com/vmware/govmomi/fault"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

type HostConnectFlag struct {
	common

	types.HostConnectSpec

	noverify bool
}

var hostConnectFlagKey = flagKey("hostConnect")

func NewHostConnectFlag(ctx context.Context) (*HostConnectFlag, context.Context) {
	if v := ctx.Value(hostConnectFlagKey); v != nil {
		return v.(*HostConnectFlag), ctx
	}

	v := &HostConnectFlag{}
	ctx = context.WithValue(ctx, hostConnectFlagKey, v)
	return v, ctx
}

func (flag *HostConnectFlag) Register(ctx context.Context, f *flag.FlagSet) {
	flag.RegisterOnce(func() {
		f.StringVar(&flag.HostName, "hostname", "", "Hostname or IP address of the host")
		f.StringVar(&flag.UserName, "username", "", "Username of administration account on the host")
		f.StringVar(&flag.Password, "password", "", "Password of administration account on the host")
		f.StringVar(&flag.SslThumbprint, "thumbprint", "", "SHA-1 thumbprint of the host's SSL certificate")
		f.BoolVar(&flag.Force, "force", false, "Force when host is managed by another VC")

		f.BoolVar(&flag.noverify, "noverify", false, "Accept host thumbprint without verification")
	})
}

func (flag *HostConnectFlag) Process(ctx context.Context) error {
	return nil
}

// Spec attempts to fill in SslThumbprint if empty.
// First checks GOVC_TLS_KNOWN_HOSTS, if not found and noverify=true then
// use object.HostCertificateInfo to get the thumbprint.
func (flag *HostConnectFlag) Spec(c *vim25.Client) types.HostConnectSpec {
	spec := flag.HostConnectSpec

	if spec.SslThumbprint == "" {
		spec.SslThumbprint = c.Thumbprint(spec.HostName)

		if spec.SslThumbprint == "" && flag.noverify {
			var info object.HostCertificateInfo
			t := c.DefaultTransport()
			_ = info.FromURL(&url.URL{Host: spec.HostName}, t.TLSClientConfig)
			spec.SslThumbprint = info.ThumbprintSHA1
		}
	}

	return spec
}

// Fault checks if error is SSLVerifyFault, including the thumbprint if so
func (flag *HostConnectFlag) Fault(err error) error {
	var verify *types.SSLVerifyFault
	if _, ok := fault.As(err, &verify); ok {
		return fmt.Errorf("%s thumbprint=%s", err, verify.Thumbprint)
	}

	return err
}
