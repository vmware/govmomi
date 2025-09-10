// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package tpm

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/vmware/govmomi/cli"
	"github.com/vmware/govmomi/cli/flags"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type info struct {
	*flags.DatacenterFlag
}

func init() {
	cli.Register("host.tpm.info", &info{})
}

func (cmd *info) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.DatacenterFlag, ctx = flags.NewDatacenterFlag(ctx)
	cmd.DatacenterFlag.Register(ctx, f)
}

func (cmd *info) Description() string {
	return `Trusted Platform Module summary.

Examples:
  govc host.tpm.info
  govc host.tpm.info -json`
}

type TrustedPlatformModule struct {
	Name            string                                    `json:"name"`
	Supported       bool                                      `json:"supported"`
	Version         string                                    `json:"version,omitempty"`
	TxtEnabled      bool                                      `json:"txtEnabled,omitempty"`
	Attestation     *types.HostTpmAttestationInfo             `json:"attestation,omitempty"`
	StateEncryption *types.HostRuntimeInfoStateEncryptionInfo `json:"stateEncryption,omitempty"`
}

func HostTrustedPlatformModule(ctx context.Context, c *vim25.Client, root types.ManagedObjectReference) ([]TrustedPlatformModule, error) {
	v, err := view.NewManager(c).CreateContainerView(ctx, root, []string{"HostSystem"}, true)
	if err != nil {
		return nil, err
	}

	defer v.Destroy(ctx)

	props := []string{
		"name",
		"summary.tpmAttestation",
		"summary.runtime.stateEncryption",
		"capability.tpmSupported",
		"capability.tpmVersion",
		"capability.txtEnabled",
	}

	var hosts []mo.HostSystem
	err = v.Retrieve(ctx, []string{"HostSystem"}, props, &hosts)
	if err != nil {
		return nil, err
	}

	tpm := make([]TrustedPlatformModule, len(hosts))

	b := func(v *bool) bool {
		if v == nil {
			return false
		}
		return *v
	}

	for i, host := range hosts {
		m := TrustedPlatformModule{
			Name:        host.Name,
			Attestation: host.Summary.TpmAttestation,
		}
		if host.Capability != nil {
			m.Supported = host.Capability.TpmSupported
			m.Version = host.Capability.TpmVersion
			m.TxtEnabled = b(host.Capability.TxtEnabled)
		}
		if host.Summary.Runtime != nil {
			m.StateEncryption = host.Summary.Runtime.StateEncryption
		}
		tpm[i] = m
	}

	return tpm, nil
}

func (cmd *info) Run(ctx context.Context, f *flag.FlagSet) error {
	dc, err := cmd.DatacenterIfSpecified()
	if err != nil {
		return err
	}
	c, err := cmd.Client()
	if err != nil {
		return err
	}

	root := c.ServiceContent.RootFolder
	if dc != nil {
		root = dc.Reference()
	}

	tpm, err := HostTrustedPlatformModule(ctx, c, root)
	if err != nil {
		return err
	}

	return cmd.WriteResult(infoResult(tpm))
}

type infoResult []TrustedPlatformModule

func (r infoResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 2, 0, 2, ' ', 0)

	fields := []string{"Name", "Attestation", "Last Verified", "TPM version", "TXT", "Message"}
	fmt.Fprintln(tw, strings.Join(fields, "\t"))

	for _, h := range r {
		if h.Supported {
			fields = []string{
				h.Name,
				string(h.Attestation.Status),
				h.Attestation.Time.Format(time.RFC3339),
				h.Version,
				strconv.FormatBool(h.TxtEnabled),
			}
			if m := h.Attestation.Message; m != nil {
				fields = append(fields, m.Message)
			}
		} else {
			fields = []string{h.Name, "N/A", "N/A", "N/A", "N/A"}
		}
		fmt.Fprintln(tw, strings.Join(fields, "\t"))
	}

	return tw.Flush()
}
