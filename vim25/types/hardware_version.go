// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"regexp"
	"strconv"
)

// HardwareVersion is a VMX hardware version.
//
// Please see https://knowledge.broadcom.com/external/article/315655 for more
// information on when a hardware version was introduced.
//
// Please refer to https://knowledge.broadcom.com/external/article/312100 for
// what ESX/ESXi versions support which hardware versions.
type HardwareVersion uint8

const (
	invalidHardwareVersion HardwareVersion = 0
)

const (
	// VMX3 was introduced in ESX 2.x.
	VMX3 HardwareVersion = iota + 3

	// VMX4 was introduced in ESX 3.x.
	VMX4

	vmx5 // invalid

	// VMX6 was introduced in Workstation 6.0.x and is not supported by
	// ESX/ESXi per https://knowledge.broadcom.com/external/article/312100.
	VMX6

	// VMX7 was introduced in ESXi 4.x.
	VMX7

	// VMX8 was introduced in ESXi 5.0.
	VMX8

	// VMX9 was introduced in ESXi 5.1.
	VMX9

	// VMX10 was introduced in ESXi 5.5.
	VMX10

	// VMX11 was introduced in ESXi 6.0.
	VMX11

	// VMX12 was introduced in Workstation 12.x. and is not supported by
	// ESX/ESXi per https://knowledge.broadcom.com/external/article/312100.
	VMX12

	// VMX13 was introduced in ESXi 6.5.
	VMX13

	// VMX14 was introduced in ESXi 6.7.
	VMX14

	// VMX15 was introduced in ESXi 6.7 U2.
	VMX15

	// VMX16 was introduced in Workstation 15.x and is not supported by
	// ESX/ESXi per https://knowledge.broadcom.com/external/article/312100.
	VMX16

	// VMX17 was introduced in ESXi 7.0.
	VMX17

	// VMX18 was introduced in ESXi 7.0 U1 (7.0.1).
	VMX18

	// VMX19 was introduced in ESXi 7.0 U2 (7.0.2).
	VMX19

	// VMX20 was introduced in ESXi 8.0.
	VMX20

	// VMX21 was introduced in ESXi 8.0 U2 (8.0.2).
	VMX21

	// VMX22 was introduced in ESX 9.0.
	VMX22
)

const (
	// MinValidHardwareVersion is the minimum, valid hardware version supported
	// by VMware hypervisors in the wild.
	MinValidHardwareVersion = VMX3

	// MaxValidHardwareVersion is the maximum, valid hardware version supported
	// by VMware hypervisors in the wild.
	MaxValidHardwareVersion = VMX22
)

// IsSupported returns true if the hardware version is known to and supported by
// GoVmomi's generated types.
func (hv HardwareVersion) IsSupported() bool {
	return hv.IsValid() &&
		hv != vmx5 &&
		hv >= MinValidHardwareVersion &&
		hv <= MaxValidHardwareVersion
}

// IsValid returns true if the hardware version is not valid.
// Unlike IsSupported, this function returns true as long as the hardware
// version is greater than 0.
// For example, the result of parsing "abc" or "vmx-abc" is an invalid hardware
// version, whereas the result of parsing "vmx-99" is valid, just not supported.
func (hv HardwareVersion) IsValid() bool {
	return hv != invalidHardwareVersion
}

func (hv HardwareVersion) String() string {
	if hv.IsValid() {
		return fmt.Sprintf("vmx-%d", hv)
	}
	return ""
}

func (hv HardwareVersion) MarshalText() ([]byte, error) {
	return []byte(hv.String()), nil
}

func (hv *HardwareVersion) UnmarshalText(text []byte) error {
	v, err := ParseHardwareVersion(string(text))
	if err != nil {
		return err
	}
	*hv = v
	return nil
}

var (
	vmxRe        = regexp.MustCompile(`(?i)^vmx-(\d+)$`)
	vmxNumOnlyRe = regexp.MustCompile(`^(\d+)$`)
)

// MustParseHardwareVersion parses the provided string into a hardware version.
func MustParseHardwareVersion(s string) HardwareVersion {
	v, err := ParseHardwareVersion(s)
	if err != nil {
		panic(err)
	}
	return v
}

// ParseHardwareVersion parses the provided string into a hardware version.
// Supported formats include vmx-123 or 123. Please note that the parser will
// only return an error if the supplied version does not match the supported
// formats.
// Once parsed, use the function IsSupported to determine if the hardware
// version falls into the range of versions known to GoVmomi.
func ParseHardwareVersion(s string) (HardwareVersion, error) {
	if m := vmxRe.FindStringSubmatch(s); len(m) > 0 {
		u, err := strconv.ParseUint(m[1], 10, 8)
		if err != nil {
			return invalidHardwareVersion, fmt.Errorf(
				"failed to parse %s from %q as uint8: %w", m[1], s, err)
		}
		return HardwareVersion(u), nil
	} else if m := vmxNumOnlyRe.FindStringSubmatch(s); len(m) > 0 {
		u, err := strconv.ParseUint(m[1], 10, 8)
		if err != nil {
			return invalidHardwareVersion, fmt.Errorf(
				"failed to parse %s as uint8: %w", m[1], err)
		}
		return HardwareVersion(u), nil
	}
	return invalidHardwareVersion, fmt.Errorf("invalid version: %q", s)
}

var hardwareVersions []HardwareVersion

func init() {
	for i := MinValidHardwareVersion; i <= MaxValidHardwareVersion; i++ {
		if i.IsSupported() {
			hardwareVersions = append(hardwareVersions, i)
		}
	}
}

// GetHardwareVersions returns a list of hardware versions.
func GetHardwareVersions() []HardwareVersion {
	dst := make([]HardwareVersion, len(hardwareVersions))
	copy(dst, hardwareVersions)
	return dst
}
