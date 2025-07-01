// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package kms

import (
	"context"
	"flag"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/crypto"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
)

type trust struct {
	*flags.ClientFlag

	client types.UploadClientCert
	server types.UploadKmipServerCert
}

func init() {
	cli.Register("kms.trust", &trust{})
}

func (cmd *trust) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)

	f.StringVar(&cmd.client.Certificate, "client-cert", "", "Client public certificate")
	f.StringVar(&cmd.client.PrivateKey, "client-key", "", "Client private key")
	f.StringVar(&cmd.server.Certificate, "server-cert", "", "Server public certificate")
}

func (cmd *trust) Usage() string {
	return "NAME"
}

func (cmd *trust) Description() string {
	return `Establish trust between KMS and vCenter.

Examples:
  # "Make vCenter Trust KMS"
  govc kms.trust -server-cert "$(govc about.cert -show)" my-kp

  # "Make KMS Trust vCenter" -> "KMS certificate and private key"
  govc kms.trust -client-cert "$(cat crt.pem) -client-key "$(cat key.pem) my-kp

  # "Download the vCenter certificate and upload it to the KMS"
  govc about.cert -show > vcenter-cert.pem`
}

func (cmd *trust) Run(ctx context.Context, f *flag.FlagSet) error {
	id := f.Arg(0)
	if id == "" {
		return flag.ErrHelp
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := crypto.GetManagerKmip(c)
	if err != nil {
		return err
	}

	if cmd.client.Certificate != "" {
		cmd.client.This = m.Reference()
		cmd.client.Cluster.Id = id
		_, err = methods.UploadClientCert(ctx, c, &cmd.client)
		if err != nil {
			return err
		}
	}

	if cmd.server.Certificate != "" {
		cmd.server.This = m.Reference()
		cmd.server.Cluster.Id = id
		_, err = methods.UploadKmipServerCert(ctx, c, &cmd.server)
		if err != nil {
			return err
		}
	}

	return nil
}
