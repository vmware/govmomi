// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package guest

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/vmware/govmomi/vim25/types"
)

type AuthFlag struct {
	auth types.NamePasswordAuthentication
	proc bool
}

func newAuthFlag(ctx context.Context) (*AuthFlag, context.Context) {
	return &AuthFlag{}, ctx
}

func (flag *AuthFlag) String() string {
	return fmt.Sprintf("%s:%s", flag.auth.Username, strings.Repeat("x", len(flag.auth.Password)))
}

func (flag *AuthFlag) Set(s string) error {
	c := strings.SplitN(s, ":", 2)
	if len(c) > 0 {
		flag.auth.Username = c[0]
		if len(c) > 1 {
			flag.auth.Password = c[1]
		}
	}

	return nil
}

func (flag *AuthFlag) Register(ctx context.Context, f *flag.FlagSet) {
	env := "GOVC_GUEST_LOGIN"
	value := os.Getenv(env)
	err := flag.Set(value)
	if err != nil {
		fmt.Printf("could not set guest login values: %v", err)
	}
	usage := fmt.Sprintf("Guest VM credentials (<user>:<password>) [%s]", env)
	f.Var(flag, "l", usage)
	if flag.proc {
		f.BoolVar(&flag.auth.GuestAuthentication.InteractiveSession, "i", false, "Interactive session")
	}
}

func (flag *AuthFlag) Process(ctx context.Context) error {
	if flag.auth.Username == "" {
		return fmt.Errorf("guest login username must not be empty")
	}

	return nil
}

func (flag *AuthFlag) Auth() types.BaseGuestAuthentication {
	return &flag.auth
}
