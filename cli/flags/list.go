// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import "fmt"

type StringList []string

func (l *StringList) String() string {
	return fmt.Sprint(*l)
}

func (l *StringList) Set(value string) error {
	*l = append(*l, value)
	return nil
}
