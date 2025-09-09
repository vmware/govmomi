// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func TestParseHardwareVersion(t *testing.T) {
	testCases := []struct {
		name                string
		in                  string
		expectedIsValid     bool
		expectedIsSupported bool
		expectedVersion     HardwareVersion
		expectedString      string
	}{
		{
			name:                "EmptyString",
			in:                  "",
			expectedIsValid:     false,
			expectedIsSupported: false,
			expectedVersion:     0,
			expectedString:      "",
		},
		{
			name:                "PrefixSansNumber",
			in:                  "vmx-",
			expectedIsValid:     false,
			expectedIsSupported: false,
			expectedVersion:     0,
			expectedString:      "",
		},
		{
			name:                "NumberSansPrefix",
			in:                  "13",
			expectedIsValid:     true,
			expectedIsSupported: true,
			expectedVersion:     VMX13,
			expectedString:      "vmx-13",
		},
		{
			name:                "PrefixAndNumber",
			in:                  "vmx-13",
			expectedIsValid:     true,
			expectedIsSupported: true,
			expectedVersion:     VMX13,
			expectedString:      "vmx-13",
		},
		{
			name:                "UpperPrefixAndNumber",
			in:                  "VMX-18",
			expectedIsValid:     true,
			expectedIsSupported: true,
			expectedVersion:     VMX18,
			expectedString:      "vmx-18",
		},
		{
			name:                "vmx-99",
			in:                  "vmx-99",
			expectedIsValid:     true,
			expectedIsSupported: false,
			expectedVersion:     HardwareVersion(99),
			expectedString:      "vmx-99",
		},
		{
			name:                "ETooLarge",
			in:                  "vmx-512",
			expectedIsValid:     false,
			expectedIsSupported: false,
			expectedVersion:     0,
			expectedString:      "",
		},
		{
			name:                "ETooLarge2",
			in:                  "512",
			expectedIsValid:     false,
			expectedIsSupported: false,
			expectedVersion:     0,
			expectedString:      "",
		},
	}

	t.Run("ParseHardwareVersion", func(t *testing.T) {
		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				v, err := ParseHardwareVersion(tc.in)
				if err != nil && tc.expectedIsValid {
					t.Fatalf("unexpected error: %v", err)
				}
				if a, e := v.IsValid(), tc.expectedIsValid; a != e {
					t.Errorf("unexpected invalid value: a=%v, e=%v", a, e)
				}
				if a, e := v.IsSupported(), tc.expectedIsSupported; a != e {
					t.Errorf("unexpected supported value: a=%v, e=%v", a, e)
				}
				if a, e := v, tc.expectedVersion; a != e {
					t.Errorf("unexpected value: a=%v, e=%v", a, e)
				}
				if a, e := v.String(), tc.expectedString; a != e {
					t.Errorf("unexpected string: a=%v, e=%v", a, e)
				}
			})
		}
	})

}

func TestHardwareVersion(t *testing.T) {

	var uniqueExpectedVersions []HardwareVersion

	type testCase struct {
		name                string
		in                  string
		expectedIsValid     bool
		expectedIsSupported bool
		expectedVersion     HardwareVersion
		expectedString      string
	}

	testCasesForVersion := func(
		version int,
		expectedVersion HardwareVersion,
		expectedString string) []testCase {

		uniqueExpectedVersions = append(uniqueExpectedVersions, MustParseHardwareVersion(expectedString))

		szVersion := strconv.Itoa(version)
		return []testCase{
			{
				name:                szVersion,
				in:                  szVersion,
				expectedIsValid:     true,
				expectedIsSupported: true,
				expectedVersion:     expectedVersion,
				expectedString:      expectedString,
			},
			{
				name:                "vmx-" + szVersion,
				in:                  "vmx-" + szVersion,
				expectedIsValid:     true,
				expectedIsSupported: true,
				expectedVersion:     expectedVersion,
				expectedString:      expectedString,
			},
			{
				name:                "VMX-" + szVersion,
				in:                  "VMX-" + szVersion,
				expectedIsValid:     true,
				expectedIsSupported: true,
				expectedVersion:     expectedVersion,
				expectedString:      expectedString,
			},
		}
	}

	testCases := []testCase{
		{
			name:                "EmptyString",
			in:                  "",
			expectedIsValid:     false,
			expectedIsSupported: false,
			expectedVersion:     0,
			expectedString:      "",
		},
		{
			name:                "PrefixSansVersion",
			in:                  "vmx-",
			expectedIsValid:     false,
			expectedIsSupported: false,
			expectedVersion:     0,
			expectedString:      "",
		},
		{
			name:                "PrefixAndInvalidVersion",
			in:                  "vmx-0",
			expectedIsValid:     false,
			expectedIsSupported: false,
			expectedVersion:     0,
			expectedString:      "",
		},
		{
			name:                "UnsupportedVersion",
			in:                  "1",
			expectedIsValid:     true,
			expectedIsSupported: false,
			expectedVersion:     HardwareVersion(1),
			expectedString:      "vmx-1",
		},
	}

	testCases = append(testCases, testCasesForVersion(3, VMX3, "vmx-3")...)
	testCases = append(testCases, testCasesForVersion(4, VMX4, "vmx-4")...)
	testCases = append(testCases, testCasesForVersion(6, VMX6, "vmx-6")...)
	testCases = append(testCases, testCasesForVersion(7, VMX7, "vmx-7")...)
	testCases = append(testCases, testCasesForVersion(8, VMX8, "vmx-8")...)
	testCases = append(testCases, testCasesForVersion(9, VMX9, "vmx-9")...)
	testCases = append(testCases, testCasesForVersion(10, VMX10, "vmx-10")...)
	testCases = append(testCases, testCasesForVersion(11, VMX11, "vmx-11")...)
	testCases = append(testCases, testCasesForVersion(12, VMX12, "vmx-12")...)
	testCases = append(testCases, testCasesForVersion(13, VMX13, "vmx-13")...)
	testCases = append(testCases, testCasesForVersion(14, VMX14, "vmx-14")...)
	testCases = append(testCases, testCasesForVersion(15, VMX15, "vmx-15")...)
	testCases = append(testCases, testCasesForVersion(16, VMX16, "vmx-16")...)
	testCases = append(testCases, testCasesForVersion(17, VMX17, "vmx-17")...)
	testCases = append(testCases, testCasesForVersion(18, VMX18, "vmx-18")...)
	testCases = append(testCases, testCasesForVersion(19, VMX19, "vmx-19")...)
	testCases = append(testCases, testCasesForVersion(20, VMX20, "vmx-20")...)
	testCases = append(testCases, testCasesForVersion(21, VMX21, "vmx-21")...)
	testCases = append(testCases, testCasesForVersion(22, VMX22, "vmx-22")...)

	t.Run("GetHardwareVersions", func(t *testing.T) {
		a, e := uniqueExpectedVersions, GetHardwareVersions()
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
		sort.Slice(e, func(i, j int) bool { return e[i] < e[j] })
		if a, e := len(a), len(e); a != e {
			t.Errorf("unexpected number of versions: a=%d, e=%d", a, e)
		}
		for i := range a {
			if a[i] != e[i] {
				t.Errorf("unexpected version: i=%d, a=%s, e=%s", i, a, e)
			}
		}
	})

	t.Run("ParseHardwareVersion", func(t *testing.T) {
		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				v, err := ParseHardwareVersion(tc.in)
				if err != nil && tc.expectedIsValid {
					t.Fatalf("unexpected error: %v", err)
				}
				if a, e := v.IsValid(), tc.expectedIsValid; a != e {
					t.Errorf("unexpected invalid value: a=%v, e=%v", a, e)
				}
				if tc.expectedIsValid {
					if a, e := v, tc.expectedVersion; a != e {
						t.Errorf("unexpected value: a=%v, e=%v", a, e)
					}
					if a, e := v.String(), tc.expectedString; a != e {
						t.Errorf("unexpected string: a=%v, e=%v", a, e)
					}
				}
			})
		}
	})
	t.Run("MarshalText", func(t *testing.T) {
		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				data, err := tc.expectedVersion.MarshalText()
				if err != nil {
					t.Fatalf("unexpected error marshaling text: %v", err)
				}
				if a, e := string(data), tc.expectedString; a != e {
					t.Errorf("unexpected data marshaling text: %s", a)
				}
			})
		}
	})

	t.Run("UnmarshalText", func(t *testing.T) {
		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				var (
					data = []byte(tc.in)
					v    HardwareVersion
				)
				if err := v.UnmarshalText(data); err != nil {
					if !tc.expectedIsValid {
						if a, e := err.Error(), fmt.Sprintf("invalid version: %q", tc.in); a != e {
							t.Errorf("unexpected error unmarshaling text: %q", a)
						}
					} else {
						t.Errorf("unexpected error unmarshaling text: %v", err)
					}
				} else if a, e := v, v; a != e {
					t.Errorf("unexpected data unmarshaling text: %s", a)
				}
			})
		}
	})
}
