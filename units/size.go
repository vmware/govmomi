// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package units

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type ByteSize int64

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func (b ByteSize) String() string {
	switch {
	case b >= EB:
		return fmt.Sprintf("%.1fEB", float32(b)/EB)
	case b >= PB:
		return fmt.Sprintf("%.1fPB", float32(b)/PB)
	case b >= TB:
		return fmt.Sprintf("%.1fTB", float32(b)/TB)
	case b >= GB:
		return fmt.Sprintf("%.1fGB", float32(b)/GB)
	case b >= MB:
		return fmt.Sprintf("%.1fMB", float32(b)/MB)
	case b >= KB:
		return fmt.Sprintf("%.1fKB", float32(b)/KB)
	}
	return fmt.Sprintf("%dB", b)
}

type FileSize int64

func (b FileSize) String() string {
	switch {
	case b >= EB:
		return fmt.Sprintf("%.1fE", float32(b)/EB)
	case b >= PB:
		return fmt.Sprintf("%.1fP", float32(b)/PB)
	case b >= TB:
		return fmt.Sprintf("%.1fT", float32(b)/TB)
	case b >= GB:
		return fmt.Sprintf("%.1fG", float32(b)/GB)
	case b >= MB:
		return fmt.Sprintf("%.1fM", float32(b)/MB)
	case b >= KB:
		return fmt.Sprintf("%.1fK", float32(b)/KB)
	}
	return fmt.Sprintf("%d", b)
}

var bytesRegexp = regexp.MustCompile(`^(?i)(\d+)([BKMGTPE]?)(ib|b)?$`)

func (b *ByteSize) Set(s string) error {
	m := bytesRegexp.FindStringSubmatch(s)
	if len(m) == 0 {
		return errors.New("invalid byte value")
	}

	i, err := strconv.ParseInt(m[1], 10, 64)
	if err != nil {
		return err
	}
	var multiplier int64 = 1

	switch m[2] {
	case "B", "b", "":
	case "K", "k":
		multiplier = KB
	case "M", "m":
		multiplier = MB
	case "G", "g":
		multiplier = GB
	case "T", "t":
		multiplier = TB
	case "P", "p":
		multiplier = PB
	case "E", "e":
		multiplier = EB
	default:
		return errors.New("invalid byte suffix")
	}

	if i > math.MaxInt64/multiplier {
		return errors.New("byte value out of range")
	}

	*b = ByteSize(i * multiplier)

	return nil
}
