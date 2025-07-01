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

func TestESXiVersion(t *testing.T) {

	var uniqueExpectedVersions []ESXiVersion

	type testCase struct {
		name                    string
		in                      string
		expectedIsValid         bool
		expectedVersion         ESXiVersion
		expectedString          string
		expectedHardwareVersion HardwareVersion
	}

	testCasesForMajorVersion := func(
		major int,
		expectedVersion ESXiVersion,
		expectedString string,
		expectedHardwareVersion HardwareVersion) []testCase {

		uniqueExpectedVersions = append(uniqueExpectedVersions, MustParseESXiVersion(expectedString))

		szMajor := strconv.Itoa(major)
		return []testCase{
			{
				name:                    szMajor,
				in:                      szMajor,
				expectedIsValid:         true,
				expectedVersion:         expectedVersion,
				expectedString:          expectedString,
				expectedHardwareVersion: expectedHardwareVersion,
			},
			{
				name:                    szMajor + ".0",
				in:                      szMajor + ".0",
				expectedIsValid:         true,
				expectedVersion:         expectedVersion,
				expectedString:          expectedString,
				expectedHardwareVersion: expectedHardwareVersion,
			},
			{
				name:                    szMajor + ".0.0",
				in:                      szMajor + ".0.0",
				expectedIsValid:         true,
				expectedVersion:         expectedVersion,
				expectedString:          expectedString,
				expectedHardwareVersion: expectedHardwareVersion,
			},
		}
	}

	testCasesForMajorMinorVersion := func(
		major, minor int,
		expectedVersion ESXiVersion,
		expectedString string,
		expectedHardwareVersion HardwareVersion) []testCase {

		uniqueExpectedVersions = append(uniqueExpectedVersions, MustParseESXiVersion(expectedString))

		szMajor := strconv.Itoa(major)
		szMinor := strconv.Itoa(minor)

		return []testCase{
			{
				name:                    szMajor + "." + szMinor,
				in:                      szMajor + "." + szMinor,
				expectedIsValid:         true,
				expectedVersion:         expectedVersion,
				expectedString:          expectedString,
				expectedHardwareVersion: expectedHardwareVersion,
			},
			{
				name:                    szMajor + "." + szMinor + ".0",
				in:                      szMajor + "." + szMinor + ".0",
				expectedIsValid:         true,
				expectedVersion:         expectedVersion,
				expectedString:          expectedString,
				expectedHardwareVersion: expectedHardwareVersion,
			},
		}
	}

	testCasesForMajorMinorUpdateVersion := func(
		major, minor, update int,
		expectedVersion ESXiVersion,
		expectedString string,
		expectedHardwareVersion HardwareVersion) []testCase {

		uniqueExpectedVersions = append(uniqueExpectedVersions, MustParseESXiVersion(expectedString))

		szMajor := strconv.Itoa(major)
		szMinor := strconv.Itoa(minor)
		szUpdate := strconv.Itoa(update)

		return []testCase{
			{
				name:                    szMajor + "." + szMinor + "." + szUpdate,
				in:                      szMajor + "." + szMinor + "." + szUpdate,
				expectedIsValid:         true,
				expectedVersion:         expectedVersion,
				expectedString:          expectedString,
				expectedHardwareVersion: expectedHardwareVersion,
			},
			{
				name:                    szMajor + "." + szMinor + "u" + szUpdate,
				in:                      szMajor + "." + szMinor + "u" + szUpdate,
				expectedIsValid:         true,
				expectedVersion:         expectedVersion,
				expectedString:          expectedString,
				expectedHardwareVersion: expectedHardwareVersion,
			},
			{
				name:                    szMajor + "." + szMinor + "U" + szUpdate,
				in:                      szMajor + "." + szMinor + "U" + szUpdate,
				expectedIsValid:         true,
				expectedVersion:         expectedVersion,
				expectedString:          expectedString,
				expectedHardwareVersion: expectedHardwareVersion,
			},
			{
				name:                    szMajor + "." + szMinor + " u" + szUpdate,
				in:                      szMajor + "." + szMinor + " u" + szUpdate,
				expectedIsValid:         true,
				expectedVersion:         expectedVersion,
				expectedString:          expectedString,
				expectedHardwareVersion: expectedHardwareVersion,
			},
			{
				name:                    szMajor + "." + szMinor + " U" + szUpdate,
				in:                      szMajor + "." + szMinor + " U" + szUpdate,
				expectedIsValid:         true,
				expectedVersion:         expectedVersion,
				expectedString:          expectedString,
				expectedHardwareVersion: expectedHardwareVersion,
			},
		}
	}

	testCases := []testCase{
		{
			name: "EmptyString",
			in:   "",
		},
		{
			name: "InvalidMajorVersion",
			in:   "1",
		},
		{
			name: "ValidMajorInvalidMinorVersion",
			in:   "2.1",
		},
		{
			name: "ValidMajorValidMinorInvalidPatchVersion",
			in:   "7.0.5",
		},
		{
			name: "ValidMajorValidMinorInvalidUpdateVersion",
			in:   "7.0U5",
		},
		{
			name: "ValidMajorValidMinorValidPatchInvalidUpdateVersion",
			in:   "7.0.0U5",
		},
	}

	testCases = append(testCases, testCasesForMajorVersion(2, ESXi2000, "2", VMX3)...)
	testCases = append(testCases, testCasesForMajorVersion(3, ESXi3000, "3", VMX4)...)
	testCases = append(testCases, testCasesForMajorVersion(4, ESXi4000, "4", VMX7)...)
	testCases = append(testCases, testCasesForMajorVersion(5, ESXi5000, "5.0", VMX8)...)
	testCases = append(testCases, testCasesForMajorMinorVersion(5, 1, ESXi5100, "5.1", VMX9)...)
	testCases = append(testCases, testCasesForMajorMinorVersion(5, 5, ESXi5500, "5.5", VMX10)...)
	testCases = append(testCases, testCasesForMajorVersion(6, ESXi6000, "6.0", VMX11)...)
	testCases = append(testCases, testCasesForMajorMinorVersion(6, 5, ESXi6500, "6.5", VMX13)...)
	testCases = append(testCases, testCasesForMajorMinorVersion(6, 7, ESXi6700, "6.7", VMX14)...)
	testCases = append(testCases, testCasesForMajorMinorUpdateVersion(6, 7, 2, ESXi6720, "6.7.2", VMX15)...)
	testCases = append(testCases, testCasesForMajorVersion(7, ESXi7000, "7.0", VMX17)...)
	testCases = append(testCases, testCasesForMajorMinorUpdateVersion(7, 0, 1, ESXi7010, "7.0.1", VMX18)...)
	testCases = append(testCases, testCasesForMajorMinorUpdateVersion(7, 0, 2, ESXi7020, "7.0.2", VMX19)...)
	testCases = append(testCases, testCasesForMajorVersion(8, ESXi8000, "8.0", VMX20)...)
	testCases = append(testCases, testCasesForMajorMinorUpdateVersion(8, 0, 1, ESXi8010, "8.0.1", VMX20)...)
	testCases = append(testCases, testCasesForMajorMinorUpdateVersion(8, 0, 2, ESXi8020, "8.0.2", VMX21)...)

	t.Run("GetESXiVersions", func(t *testing.T) {
		a, e := uniqueExpectedVersions, GetESXiVersions()
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

	t.Run("ParseESXiVersion", func(t *testing.T) {
		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				v, err := ParseESXiVersion(tc.in)
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
					if a, e := v.IsHardwareVersionSupported(tc.expectedHardwareVersion), true; a != e {
						t.Errorf("is hardware version supported failed for %s", tc.expectedHardwareVersion)
					}
					if a, e := v.HardwareVersion(), tc.expectedHardwareVersion; a != e {
						t.Errorf("unexpected hardware version: a=%s, e=%s", a, e)
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
					v    ESXiVersion
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
