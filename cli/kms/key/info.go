// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strconv"
	"text/tabwriter"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/crypto"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.ClientFlag
	*flags.OutputFlag

	provider string

	avail, vms, hosts, other bool
}

func init() {
	cli.Register("kms.key.info", &info{}, true)
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.StringVar(&cmd.provider, "p", "", "Provider ID")
	f.BoolVar(&cmd.avail, "a", true, "Show key availability")
	f.BoolVar(&cmd.vms, "vms", false, "Show VMs using keys")
	f.BoolVar(&cmd.hosts, "hosts", false, "Show hosts using keys")
	f.BoolVar(&cmd.other, "other", false, "Show 3rd party using keys")
}

func (cmd *info) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return cmd.OutputFlag.Process(ctx)
}

func (cmd *info) Usage() string {
	return "ID..."
}

func (cmd *info) Description() string {
	return `Crypto key info.

Examples:
  govc kms.key.info -p my-kp ID
  govc kms.key.info -p my-kp -vms ID`
}

type infoResult struct {
	Status []types.CryptoManagerKmipCryptoKeyStatus `json:"status"`
	ctx    context.Context
	c      *vim25.Client
}

func (r *infoResult) writeObj(w io.Writer, refs []types.ManagedObjectReference) {
	for _, ref := range refs {
		p, err := find.InventoryPath(r.ctx, r.c, ref)
		if err != nil {
			p = ref.String()
		}
		fmt.Fprintf(w, "  %s\t%s\n", ref.Value, p)
	}
}

func (r *infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	for _, s := range r.Status {
		avail := "-"
		reason := s.Reason
		if s.KeyAvailable != nil {
			avail = strconv.FormatBool(*s.KeyAvailable)
			if *s.KeyAvailable && reason == "" {
				reason = "KeyStateAvailable"
			}
		}
		pid := ""
		if p := s.KeyId.ProviderId; p != nil {
			pid = p.Id
		}
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n",
			pid, s.KeyId.KeyId, avail, reason)
		r.writeObj(tw, s.EncryptedVMs)
		r.writeObj(tw, s.AffectedHosts)
	}

	return tw.Flush()
}

func argsToKeys(p string, args []string) []types.CryptoKeyId {
	ids := make([]types.CryptoKeyId, len(args))

	var provider *types.KeyProviderId

	if p != "" {
		provider = &types.KeyProviderId{Id: p}
	}

	for i, id := range args {
		ids[i] = types.CryptoKeyId{
			KeyId:      id,
			ProviderId: provider,
		}
	}

	return ids
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	n := f.NArg()
	if n == 0 {
		return flag.ErrHelp
	}

	ids := argsToKeys(cmd.provider, f.Args())

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := crypto.GetManagerKmip(c)
	if err != nil {
		return err
	}

	check := int32(0)

	opts := []struct {
		enabled bool
		val     int32
	}{
		{cmd.avail, crypto.CheckKeyAvailable},
		{cmd.vms, crypto.CheckKeyUsedByVms},
		{cmd.hosts, crypto.CheckKeyUsedByHosts},
		{cmd.other, crypto.CheckKeyUsedByOther},
	}

	for _, opt := range opts {
		if opt.enabled {
			check = check | opt.val
		}
	}

	res, err := m.QueryCryptoKeyStatus(ctx, ids, check)
	if err != nil {
		return err
	}

	return cmd.WriteResult(&infoResult{res, ctx, c})
}
