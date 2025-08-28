// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"regexp"
)

var (
	isGuestIDWindowsRx = regexp.MustCompile(`(?i)^win.+$`)
	isGuestIDDarwinRx  = regexp.MustCompile(`(?i)^darwin.+$`)
	isGuestIDLinuxRx   = regexp.MustCompile(`(?i)((^.*linux.*$)|(^(centos|coreos|debian|fedora|mandrake|mandriva|opensuse|redhat|rhel|suse|sles|ubuntu|vmwarephoton).+$))`)
	isGuestIDNetwareRx = regexp.MustCompile(`(?i)^netware.+$`)
	isGuestIDSolarisRx = regexp.MustCompile(`(?i)^solaris.+$`)
)

// ToFamily returns the family to which the provided guest identifier belongs.
func (g VirtualMachineGuestOsIdentifier) ToFamily() VirtualMachineGuestOsFamily {
	return GuestIDToFamily(g)
}

// GuestIDToFamily returns the family to which the provided guest identifier
// belongs.
func GuestIDToFamily[T ~string](guestID T) VirtualMachineGuestOsFamily {
	szGuestID := string(guestID)
	switch {
	case isGuestIDDarwinRx.MatchString(szGuestID):
		return VirtualMachineGuestOsFamilyDarwinGuestFamily
	case isGuestIDLinuxRx.MatchString(szGuestID):
		return VirtualMachineGuestOsFamilyLinuxGuest
	case isGuestIDSolarisRx.MatchString(szGuestID):
		return VirtualMachineGuestOsFamilySolarisGuest
	case isGuestIDNetwareRx.MatchString(szGuestID):
		return VirtualMachineGuestOsFamilyNetwareGuest
	case isGuestIDWindowsRx.MatchString(szGuestID):
		return VirtualMachineGuestOsFamilyWindowsGuest
	default:
		return VirtualMachineGuestOsFamilyOtherGuestFamily
	}
}
