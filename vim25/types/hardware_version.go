/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/mod/semver"
)

// hardwareVersions is a list of supported hardware versions from
// https://kb.vmware.com/s/article/1003746.
var hardwareVersions = []HardwareVersion{
	"vmx-3",
	"vmx-4",
	"vmx-6",
	"vmx-7",
	"vmx-8",
	"vmx-9",
	"vmx-10",
	"vmx-11",
	"vmx-12",
	"vmx-13",
	"vmx-14",
	"vmx-15",
	"vmx-16",
	"vmx-17",
	"vmx-18",
	"vmx-19",
	"vmx-20",
	"vmx-21",
}

var esxiVersions = []string{
	"8.0.2",
	"8.0",
	"7.0.2",
	"7.0.1",
	"7.0.0",
	"6.7.2",
	"6.7",
	"6.5",
	"6.0",
	"5.5",
	"5.1",
	"5.0",
	"4.0",
	"3",
	"2",
}

// GetESXiVersions returns a list of ESXi versions.
func GetESXiVersions() []string {
	dst := make([]string, len(esxiVersions))
	copy(dst, esxiVersions)
	return dst
}

// GetHardwareVersions returns a list of hardware versions.
func GetHardwareVersions() []HardwareVersion {
	dst := make([]HardwareVersion, len(hardwareVersions))
	copy(dst, hardwareVersions)
	return dst
}

// GetSupportedHardwareVersionsForESXi returns a list of the hardware versions
// supported by the specified ESXi version.
func GetSupportedHardwareVersionsForESXi(esxVersion string) []HardwareVersion {
	mv, ok := GetHardwareVersionForESXi(esxVersion)
	if !ok {
		return nil
	}
	var hv []HardwareVersion
	for i := range hardwareVersions {
		if hardwareVersions[i].Int() <= mv.Int() {
			hv = append(hv, hardwareVersions[i])
		}
	}
	return hv
}

// GetHardwareVersionForESXi returns the maximum supported hardware
// version for a given ESX version. Please note that ESX versions prior to 7.0
// did not use semantic versioning to denote update releases, ex. 6.7U2 is
// 6.7.0.BUILD instead of 6.7.2. For the purposes of this function, however,
// please use 6.7.2 for 6.7U2.
func GetHardwareVersionForESXi(esxVersion string) (HardwareVersion, bool) {
	if !strings.HasPrefix(esxVersion, "v") {
		esxVersion = "v" + esxVersion
	}

	if !semver.IsValid(esxVersion) {
		return "", false
	}

	// From https://kb.vmware.com/s/article/1003746
	switch {
	case semver.Compare(esxVersion, "v8.0.2") >= 0:
		return "vmx-21", true
	case semver.Compare(esxVersion, "v8.0") >= 0:
		return "vmx-20", true
	case semver.Compare(esxVersion, "v7.0.2") >= 0:
		return "vmx-19", true
	case semver.Compare(esxVersion, "v7.0.1") >= 0:
		return "vmx-18", true
	case semver.Compare(esxVersion, "v7.0.0") >= 0:
		return "vmx-17", true
	case semver.Compare(esxVersion, "v6.7.2") >= 0:
		return "vmx-15", true
	case semver.Compare(esxVersion, "v6.7") >= 0:
		return "vmx-14", true
	case semver.Compare(esxVersion, "v6.5") >= 0:
		return "vmx-13", true
	case semver.Compare(esxVersion, "v6.0") >= 0:
		return "vmx-11", true
	case semver.Compare(esxVersion, "v5.5") >= 0:
		return "vmx-10", true
	case semver.Compare(esxVersion, "v5.1") >= 0:
		return "vmx-9", true
	case semver.Compare(esxVersion, "v5.0") >= 0:
		return "vmx-8", true
	case semver.Compare(esxVersion, "v4.0") >= 0:
		return "vmx-7", true
	case semver.Compare(esxVersion, "v3") >= 0:
		return "vmx-4", true
	case semver.Compare(esxVersion, "v2") >= 0:
		return "vmx-3", true
	}
	return "", false
}

type HardwareVersion string

func (hv HardwareVersion) IsValid() bool {
	_, err := hv.int()
	return err == nil
}

// String returns the string value of the hardware version when IsValid is true,
// otherwise an empty string is returned.
func (hv HardwareVersion) String() string {
	i, err := hv.int()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("vmx-%d", i)
}

// Int returns the numerical value of the hardware version when IsValid is true,
// otherwise 0 is returned.
func (hv HardwareVersion) Int() int {
	i, _ := hv.int()
	return i
}

var vmxRe = regexp.MustCompile(`(?i)^vmx-(\d+)$`)

func (hv HardwareVersion) int() (int, error) {
	if m := vmxRe.FindStringSubmatch(string(hv)); len(m) > 0 {
		v, err := strconv.ParseInt(m[1], 10, 0)
		if err != nil {
			return 0, err
		}
		return int(v), nil
	}
	return 0, fmt.Errorf("invalid hardware version: %q", string(hv))
}
