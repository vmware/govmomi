// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"context"
	"flag"
)

type EmptyFlag struct{}

func (flag *EmptyFlag) Register(ctx context.Context, f *flag.FlagSet) {
}

func (flag *EmptyFlag) Process(ctx context.Context) error {
	return nil
}
