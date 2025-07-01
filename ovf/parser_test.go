// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidInteger(t *testing.T) {
	t.Run("Valid integer terms", func(t *testing.T) {
		testCases := []string{
			"1",
			"1024",
			"1000",
			"1000000000",
		}

		for _, tc := range testCases {
			t.Run(tc, func(t *testing.T) {
				assert.True(t, validIntegerString(tc))
			})
		}
	})

	// while these all may not necessarily be invalid integer terms mathematically, they are either not valid per
	// DSP0004, or not valid per our use-case
	t.Run("Invalid integer terms", func(t *testing.T) {
		testCases := []string{
			"05",
			"1.5",
			"2^10",
			"-10",
			"1,000",
			"1,000,000",
		}

		for _, tc := range testCases {
			t.Run(tc, func(t *testing.T) {
				assert.False(t, validIntegerString(tc))
			})
		}
	})
}

func TestValidExponent(t *testing.T) {
	t.Run("Acceptable exponential terms", func(t *testing.T) {
		testCases := []string{
			"10^3",
			"2^10",
		}

		for _, tc := range testCases {
			t.Run(tc, func(t *testing.T) {
				assert.True(t, validExponentString(tc))
			})
		}
	})

	// while valid exponential terms mathematically, they are either not valid per DSP0004, or not valid per our
	// use-case
	t.Run("Unacceptable exponential terms", func(t *testing.T) {
		testCases := []string{
			"1",
			"1024",
			"05",
			"1.5",
			"-10",
			"0.5^2",
			"-10^3",
			"-2^10",
			"1e6",
		}

		for _, tc := range testCases {
			t.Run(tc, func(t *testing.T) {
				assert.False(t, validExponentString(tc))
			})
		}
	})
}

func TestValidByteUnitString(t *testing.T) {
	// Due to size constraints and our use-case, we only accept up to gigabyte
	t.Run("Acceptable SI decimal prefixes", func(t *testing.T) {
		testCases := []string{
			"byte",
			"kilobyte",
			"megabyte",
			"gigabyte",
		}

		for _, tc := range testCases {
			t.Run(tc, func(t *testing.T) {
				assert.True(t, validByteUnitString(tc))
			})
			// plural is valid
			plural := fmt.Sprintf("%ss", tc)
			t.Run(plural, func(t *testing.T) {
				assert.True(t, validByteUnitString(plural))
			})
			// function expects an already lowercased string, so capitalization is not valid
			capitalized := strings.ToUpper(tc)
			t.Run(capitalized, func(t *testing.T) {
				assert.False(t, validByteUnitString(capitalized))
			})
		}
	})

	// too large and/or currently don't fit use-case
	t.Run("Unacceptable SI decimal prefixes", func(t *testing.T) {
		testCases := []string{
			"terabyte",
			"petabyte",
			"exabyte",
			"zettabyte",
			"yottabyte",
		}

		for _, tc := range testCases {
			t.Run(tc, func(t *testing.T) {
				assert.False(t, validByteUnitString(tc))
			})
		}
	})

	// due to size constraints and our use-case, we only accept up to gibibyte
	t.Run("Acceptable IEC binary prefixes", func(t *testing.T) {
		testCases := []string{
			"kibibyte",
			"mebibyte",
			"gibibyte",
		}

		for _, tc := range testCases {
			t.Run(tc, func(t *testing.T) {
				assert.True(t, validByteUnitString(tc))
			})
			// plural is valid
			plural := fmt.Sprintf("%ss", tc)
			t.Run(plural, func(t *testing.T) {
				assert.True(t, validByteUnitString(plural))
			})
			// function expects an already-lowercased string, so capitalization is not valid
			capitalized := strings.ToUpper(tc)
			t.Run(capitalized, func(t *testing.T) {
				assert.False(t, validByteUnitString(capitalized))
			})
		}
	})

	// too large and/or currently don't fit use-case
	t.Run("Unacceptable IEC binary prefixes", func(t *testing.T) {
		testCases := []string{
			"tebibyte",
			"pebibyte",
			"exbibyte",
			"zebibyte",
			"yobibyte",
		}

		for _, tc := range testCases {
			t.Run(tc, func(t *testing.T) {
				assert.False(t, validByteUnitString(tc))
			})
		}
	})
}

func TestValidUnit(t *testing.T) {
	t.Run("Valid unit strings", func(t *testing.T) {
		testCases := []string{
			"",
			"1024",
			"2^10",
			"10^3",
			"byte",
			"1024*byte",
			"1024 * byte",
			"byte*1024",
			"byte * 1024",
			"2^10*byte",
			"2^10 * bytes",
			"byte*2^10",
			"byte * 2^10",
			"1024 * 1024*1024",
			"2^10*2^10 * 2^10",
			"2^10 * 1024 * 2^10",
			"2^10*1024 * byte",
			"byte*2^10 * 1024",
			"1000*byte * 1000 * 1000",
			"gigabyte",
			"Gigabyte",
			"1024 * 1024 * kilobyte",
			"Byte",
			"1024*Byte",
			"1024 * BYTE",
			"BYTE*1024",
			"Byte * 1024",
			"2^10*Byte",
			"2^10 * BYTE",
			"BYTE*2^10",
			"Byte * 2^10",
			"1000000",
			"1000000000*byte",
		}

		for _, tc := range testCases {
			t.Run(tc, func(t *testing.T) {
				assert.True(t, validCapacityString(tc))
			})
		}
	})

	// either these do not abide by DSP0004, or they do not make sense in the context of our use-case
	t.Run("Invalid unit strings", func(t *testing.T) {
		testCases := []string{
			"1000*",
			"*1024",
			"* 1000 * byte",
			"byte * 2^30 *",
			"1000byte",
			"1024*1024*",
			"byte*byte",
			"2^10/1024",
			"1024*1024 / 1024",
			"2 ^ 10",
			"512 + 512",
			"2048 - 1024",
			"1,000,000",
			"1,000,000,000*byte",
		}

		for _, tc := range testCases {
			t.Run(tc, func(t *testing.T) {
				assert.False(t, validCapacityString(tc))
			})
		}
	})
}

func TestParseCapacityAllocationUnits(t *testing.T) {
	t.Run("Valid capacity allocation units multiplier is correctly parsed", func(t *testing.T) {
		testCases := []struct {
			s        string
			expected int64
		}{
			{"", 1},
			{"byte", 1},
			{"kilobyte", 1000},
			{"kibibyte", 1024},
			{"megabyte", 1000 * 1000},
			{"mebibyte", 1024 * 1024},
			{"gigabyte", 1000 * 1000 * 1000},
			{"gibibyte", 1024 * 1024 * 1024},
			{"Byte", 1},
			{"Kilobyte", 1000},
			{"Kibibyte", 1024},
			{"Megabyte", 1000 * 1000},
			{"Mebibyte", 1024 * 1024},
			{"Gigabyte", 1000 * 1000 * 1000},
			{"Gibibyte", 1024 * 1024 * 1024},
			{"BYTE", 1},
			{"KILOBYTE", 1000},
			{"KIBIBYTE", 1024},
			{"MEGABYTE", 1000 * 1000},
			{"MEBIBYTE", 1024 * 1024},
			{"GIGABYTE", 1000 * 1000 * 1000},
			{"GIBIBYTE", 1024 * 1024 * 1024},
			{"10^3", 1000},
			{"1024", 1024},
			{"1000 * byte", 1000},
			{"2^10 * byte", 1024},
			{"byte * 2^10", 1024},
			{"10^9*byte", 1000 * 1000 * 1000},
			{"10^3 * megabyte", 1000 * 1000 * 1000},
			{"kibibyte * 2^10 * 1024", 1024 * 1024 * 1024},
			{"byte*2^10*2^10*1024", 1024 * 1024 * 1024},
			{"1000*byte*1000*10^3", 1000 * 1000 * 1000},
		}

		for _, tc := range testCases {
			t.Run(tc.s, func(t *testing.T) {
				assert.Equal(t, tc.expected, ParseCapacityAllocationUnits(tc.s))
			})
		}
	})

	// either these do not abide by DSP0004, or they do not make sense in the context of our use-case
	t.Run("Invalid capacity allocation units should return zero", func(t *testing.T) {
		testCases := []string{
			"bit", // the only unit, valid per DSP0004, that we care about is byte; this makes more sense than checking meter
			"nibble",
			"dataword",
			"byte*byte",
			"1000*",
			"*1024",
			"* 1000 * byte",
			"byte * 2^30 *",
			"1000byte",
			"1024*1024*",
			"byte*byte",
			"2^10/1024",
			"1024*1024 / 1024",
			"2 ^ 10",
			"512 + 512",
			"2048 - 1024",
		}

		for _, tc := range testCases {
			t.Run(tc, func(t *testing.T) {
				assert.Equal(t, int64(0), ParseCapacityAllocationUnits(tc))
			})
		}
	})
}
