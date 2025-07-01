// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"strconv"
	"strings"
)

var (
	BuildVersion = "v0.0.0" // govc-test requires an (arbitrary) version set
	BuildCommit  string
	BuildDate    string
)

type version []int

func ParseVersion(s string) (version, error) {
	// remove any trailing "v" version identifier
	s = strings.TrimPrefix(s, "v")
	v := make(version, 0)

	ds := strings.Split(s, "-")
	ps := strings.Split(ds[0], ".")
	for _, p := range ps {
		i, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}

		v = append(v, i)
	}

	return v, nil
}

func (v version) Lte(u version) bool {
	lv := len(v)
	lu := len(u)

	for i := 0; i < lv; i++ {
		// Everything up to here has been equal and v has more elements than u.
		if i >= lu {
			return false
		}

		// Move to next digit if equal.
		if v[i] == u[i] {
			continue
		}

		return v[i] < u[i]
	}

	// Equal.
	return true
}
