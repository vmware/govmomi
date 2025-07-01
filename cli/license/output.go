// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package license

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi/vim25/types"
)

type licenseOutput []types.LicenseManagerLicenseInfo

func (res licenseOutput) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 4, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "Key:\tEdition:\tUsed:\tTotal:\n")
	for _, v := range res {
		fmt.Fprintf(tw, "%s\t", v.LicenseKey)
		fmt.Fprintf(tw, "%s\t", v.EditionKey)
		fmt.Fprintf(tw, "%d\t", v.Used)
		fmt.Fprintf(tw, "%d\t", v.Total)
		fmt.Fprintf(tw, "\n")
	}
	return tw.Flush()
}
