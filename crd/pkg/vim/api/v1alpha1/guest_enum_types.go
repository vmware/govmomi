// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import "strings"

// VirtualMachineGuestOsFamily represents a possible guest family type.
// It corresponds to vim.vm.GuestOsDescriptor.GuestOsFamily.
type VirtualMachineGuestOsFamily string

const (
	// VirtualMachineGuestOsFamilyDarwin indicates an Apple macOS/Darwin guest OS.
	VirtualMachineGuestOsFamilyDarwin VirtualMachineGuestOsFamily = "Darwin"

	// VirtualMachineGuestOsFamilyLinux indicates a Linux guest OS.
	VirtualMachineGuestOsFamilyLinux VirtualMachineGuestOsFamily = "Linux"

	// VirtualMachineGuestOsFamilyNetware indicates a Novell NetWare guest OS.
	VirtualMachineGuestOsFamilyNetware VirtualMachineGuestOsFamily = "Netware"

	// VirtualMachineGuestOsFamilyOther indicates a guest OS not covered by any
	// specific family.
	VirtualMachineGuestOsFamilyOther VirtualMachineGuestOsFamily = "Other"

	// VirtualMachineGuestOsFamilySolaris indicates a Sun/Oracle Solaris guest OS.
	VirtualMachineGuestOsFamilySolaris VirtualMachineGuestOsFamily = "Solaris"

	// VirtualMachineGuestOsFamilyWindows indicates a Microsoft Windows guest OS.
	VirtualMachineGuestOsFamilyWindows VirtualMachineGuestOsFamily = "Windows"
)

// ToVimType returns the Vim identifier for the VirtualMachineGuestOsFamily.
func (t VirtualMachineGuestOsFamily) ToVimType() string {
	switch t {
	case VirtualMachineGuestOsFamilyDarwin:
		return "darwinGuestFamily"
	case VirtualMachineGuestOsFamilyLinux:
		return "linuxGuest"
	case VirtualMachineGuestOsFamilyNetware:
		return "netwareGuest"
	case VirtualMachineGuestOsFamilyOther:
		return "otherGuestFamily"
	case VirtualMachineGuestOsFamilySolaris:
		return "solarisGuest"
	case VirtualMachineGuestOsFamilyWindows:
		return "windowsGuest"
	}
	return string(t)
}

// FromVimType returns the VirtualMachineGuestOsFamily from the Vim
// identifier.
func (t *VirtualMachineGuestOsFamily) FromVimType(s string) {
	switch s {
	case "darwinGuestFamily":
		*t = VirtualMachineGuestOsFamilyDarwin
	case "linuxGuest":
		*t = VirtualMachineGuestOsFamilyLinux
	case "netwareGuest":
		*t = VirtualMachineGuestOsFamilyNetware
	case "otherGuestFamily":
		*t = VirtualMachineGuestOsFamilyOther
	case "solarisGuest":
		*t = VirtualMachineGuestOsFamilySolaris
	case "windowsGuest":
		*t = VirtualMachineGuestOsFamilyWindows
	}
	*t = VirtualMachineGuestOsFamily(s)
}

// VirtualMachineGuestOsIdentifier identifies a guest operating system.
// It corresponds to vim.vm.GuestOsDescriptor.GuestOsIdentifier.
type VirtualMachineGuestOsIdentifier string

const (
	// Windows

	// VirtualMachineGuestOsIdentifierDOS indicates MS-DOS.
	VirtualMachineGuestOsIdentifierDOS VirtualMachineGuestOsIdentifier = "DOS"

	// VirtualMachineGuestOsIdentifierWin31 indicates Windows 3.1.
	VirtualMachineGuestOsIdentifierWin31 VirtualMachineGuestOsIdentifier = "Win31"

	// VirtualMachineGuestOsIdentifierWin95 indicates Windows 95.
	VirtualMachineGuestOsIdentifierWin95 VirtualMachineGuestOsIdentifier = "Win95"

	// VirtualMachineGuestOsIdentifierWin98 indicates Windows 98.
	VirtualMachineGuestOsIdentifierWin98 VirtualMachineGuestOsIdentifier = "Win98"

	// VirtualMachineGuestOsIdentifierWinMe indicates Windows Millennium
	// Edition.
	VirtualMachineGuestOsIdentifierWinMe VirtualMachineGuestOsIdentifier = "WinMe"

	// VirtualMachineGuestOsIdentifierWinNT indicates Windows NT 4.
	VirtualMachineGuestOsIdentifierWinNT VirtualMachineGuestOsIdentifier = "WinNT"

	// VirtualMachineGuestOsIdentifierWin2000Pro indicates Windows 2000
	// Professional.
	VirtualMachineGuestOsIdentifierWin2000Pro VirtualMachineGuestOsIdentifier = "Win2000Pro"

	// VirtualMachineGuestOsIdentifierWin2000Serv indicates Windows 2000
	// Server.
	VirtualMachineGuestOsIdentifierWin2000Serv VirtualMachineGuestOsIdentifier = "Win2000Serv"

	// VirtualMachineGuestOsIdentifierWin2000AdvServ indicates Windows 2000
	// Advanced Server.
	VirtualMachineGuestOsIdentifierWin2000AdvServ VirtualMachineGuestOsIdentifier = "Win2000AdvServ"

	// VirtualMachineGuestOsIdentifierWinXPHome indicates Windows XP Home
	// Edition.
	VirtualMachineGuestOsIdentifierWinXPHome VirtualMachineGuestOsIdentifier = "WinXPHome"

	// VirtualMachineGuestOsIdentifierWinXPPro indicates Windows XP
	// Professional.
	VirtualMachineGuestOsIdentifierWinXPPro VirtualMachineGuestOsIdentifier = "WinXPPro"

	// VirtualMachineGuestOsIdentifierWinXPPro64 indicates Windows XP
	// Professional Edition (64 bit).
	VirtualMachineGuestOsIdentifierWinXPPro64 VirtualMachineGuestOsIdentifier = "WinXPPro64"

	// VirtualMachineGuestOsIdentifierWinNetWeb indicates Windows Server
	// 2003, Web Edition.
	VirtualMachineGuestOsIdentifierWinNetWeb VirtualMachineGuestOsIdentifier = "WinNetWeb"

	// VirtualMachineGuestOsIdentifierWinNetStandard indicates Windows
	// Server 2003, Standard Edition.
	VirtualMachineGuestOsIdentifierWinNetStandard VirtualMachineGuestOsIdentifier = "WinNetStandard"

	// VirtualMachineGuestOsIdentifierWinNetEnterprise indicates Windows
	// Server 2003, Enterprise Edition.
	VirtualMachineGuestOsIdentifierWinNetEnterprise VirtualMachineGuestOsIdentifier = "WinNetEnterprise"

	// VirtualMachineGuestOsIdentifierWinNetDatacenter indicates Windows
	// Server 2003, Datacenter Edition.
	VirtualMachineGuestOsIdentifierWinNetDatacenter VirtualMachineGuestOsIdentifier = "WinNetDatacenter"

	// VirtualMachineGuestOsIdentifierWinNetBusiness indicates Windows
	// Small Business Server 2003.
	VirtualMachineGuestOsIdentifierWinNetBusiness VirtualMachineGuestOsIdentifier = "WinNetBusiness"

	// VirtualMachineGuestOsIdentifierWinNetStandard64 indicates Windows
	// Server 2003, Standard Edition (64 bit).
	VirtualMachineGuestOsIdentifierWinNetStandard64 VirtualMachineGuestOsIdentifier = "WinNetStandard64"

	// VirtualMachineGuestOsIdentifierWinNetEnterprise64 indicates Windows
	// Server 2003, Enterprise Edition (64 bit).
	VirtualMachineGuestOsIdentifierWinNetEnterprise64 VirtualMachineGuestOsIdentifier = "WinNetEnterprise64"

	// VirtualMachineGuestOsIdentifierWinLonghorn indicates Windows
	// Longhorn.
	VirtualMachineGuestOsIdentifierWinLonghorn VirtualMachineGuestOsIdentifier = "WinLonghorn"

	// VirtualMachineGuestOsIdentifierWinLonghorn64 indicates Windows
	// Longhorn (64 bit).
	VirtualMachineGuestOsIdentifierWinLonghorn64 VirtualMachineGuestOsIdentifier = "WinLonghorn64"

	// VirtualMachineGuestOsIdentifierWinNetDatacenter64 indicates Windows
	// Server 2003, Datacenter Edition (64 bit).
	VirtualMachineGuestOsIdentifierWinNetDatacenter64 VirtualMachineGuestOsIdentifier = "WinNetDatacenter64"

	// VirtualMachineGuestOsIdentifierWinVista indicates Windows Vista.
	VirtualMachineGuestOsIdentifierWinVista VirtualMachineGuestOsIdentifier = "WinVista"

	// VirtualMachineGuestOsIdentifierWinVista64 indicates Windows Vista
	// (64 bit).
	VirtualMachineGuestOsIdentifierWinVista64 VirtualMachineGuestOsIdentifier = "WinVista64"

	// VirtualMachineGuestOsIdentifierWindows7 indicates Windows 7.
	VirtualMachineGuestOsIdentifierWindows7 VirtualMachineGuestOsIdentifier = "Windows7"

	// VirtualMachineGuestOsIdentifierWindows7x64 indicates Windows 7
	// (64 bit).
	VirtualMachineGuestOsIdentifierWindows7x64 VirtualMachineGuestOsIdentifier = "Windows7x64"

	// VirtualMachineGuestOsIdentifierWindows7Server64 indicates Windows
	// Server 2008 R2 (64 bit).
	VirtualMachineGuestOsIdentifierWindows7Server64 VirtualMachineGuestOsIdentifier = "Windows7Server64"

	// VirtualMachineGuestOsIdentifierWindows8 indicates Windows 8.
	VirtualMachineGuestOsIdentifierWindows8 VirtualMachineGuestOsIdentifier = "Windows8"

	// VirtualMachineGuestOsIdentifierWindows8x64 indicates Windows 8
	// (64 bit).
	VirtualMachineGuestOsIdentifierWindows8x64 VirtualMachineGuestOsIdentifier = "Windows8x64"

	// VirtualMachineGuestOsIdentifierWindows8Server64 indicates Windows
	// Server 2012 (64 bit).
	VirtualMachineGuestOsIdentifierWindows8Server64 VirtualMachineGuestOsIdentifier = "Windows8Server64"

	// VirtualMachineGuestOsIdentifierWindows9 indicates Windows 10.
	VirtualMachineGuestOsIdentifierWindows9 VirtualMachineGuestOsIdentifier = "Windows9"

	// VirtualMachineGuestOsIdentifierWindows9x64 indicates Windows 10
	// (64 bit).
	VirtualMachineGuestOsIdentifierWindows9x64 VirtualMachineGuestOsIdentifier = "Windows9x64"

	// VirtualMachineGuestOsIdentifierWindows9Server64 indicates Windows
	// Server 2016 (64 bit).
	VirtualMachineGuestOsIdentifierWindows9Server64 VirtualMachineGuestOsIdentifier = "Windows9Server64"

	// VirtualMachineGuestOsIdentifierWindows11x64 indicates Windows 11.
	VirtualMachineGuestOsIdentifierWindows11x64 VirtualMachineGuestOsIdentifier = "Windows11x64"

	// VirtualMachineGuestOsIdentifierWindows12x64 indicates Windows 12.
	VirtualMachineGuestOsIdentifierWindows12x64 VirtualMachineGuestOsIdentifier = "Windows12x64"

	// VirtualMachineGuestOsIdentifierWindowsHyperV indicates Windows
	// Hyper-V.
	VirtualMachineGuestOsIdentifierWindowsHyperV VirtualMachineGuestOsIdentifier = "WindowsHyperV"

	// VirtualMachineGuestOsIdentifierWindows2019Serverx64 indicates
	// Windows Server 2019.
	VirtualMachineGuestOsIdentifierWindows2019Serverx64 VirtualMachineGuestOsIdentifier = "Windows2019Serverx64"

	// VirtualMachineGuestOsIdentifierWindows2019ServerNextx64 indicates
	// Windows Server 2022.
	VirtualMachineGuestOsIdentifierWindows2019ServerNextx64 VirtualMachineGuestOsIdentifier = "Windows2019ServerNextx64"

	// VirtualMachineGuestOsIdentifierWindows2022ServerNextx64 indicates
	// Windows Server 2025.
	VirtualMachineGuestOsIdentifierWindows2022ServerNextx64 VirtualMachineGuestOsIdentifier = "Windows2022ServerNextx64"

	// FreeBSD

	// VirtualMachineGuestOsIdentifierFreeBSD indicates FreeBSD.
	VirtualMachineGuestOsIdentifierFreeBSD VirtualMachineGuestOsIdentifier = "FreeBSD"

	// VirtualMachineGuestOsIdentifierFreeBSD64 indicates FreeBSD x64.
	VirtualMachineGuestOsIdentifierFreeBSD64 VirtualMachineGuestOsIdentifier = "FreeBSD64"

	// VirtualMachineGuestOsIdentifierFreeBSD11 indicates FreeBSD 11.
	VirtualMachineGuestOsIdentifierFreeBSD11 VirtualMachineGuestOsIdentifier = "FreeBSD11"

	// VirtualMachineGuestOsIdentifierFreeBSD11x64 indicates FreeBSD 11
	// (64 bit).
	VirtualMachineGuestOsIdentifierFreeBSD11x64 VirtualMachineGuestOsIdentifier = "FreeBSD11x64"

	// VirtualMachineGuestOsIdentifierFreeBSD12 indicates FreeBSD 12.
	VirtualMachineGuestOsIdentifierFreeBSD12 VirtualMachineGuestOsIdentifier = "FreeBSD12"

	// VirtualMachineGuestOsIdentifierFreeBSD12x64 indicates FreeBSD 12
	// (64 bit).
	VirtualMachineGuestOsIdentifierFreeBSD12x64 VirtualMachineGuestOsIdentifier = "FreeBSD12x64"

	// VirtualMachineGuestOsIdentifierFreeBSD13 indicates FreeBSD 13.
	VirtualMachineGuestOsIdentifierFreeBSD13 VirtualMachineGuestOsIdentifier = "FreeBSD13"

	// VirtualMachineGuestOsIdentifierFreeBSD13x64 indicates FreeBSD 13
	// (64 bit).
	VirtualMachineGuestOsIdentifierFreeBSD13x64 VirtualMachineGuestOsIdentifier = "FreeBSD13x64"

	// VirtualMachineGuestOsIdentifierFreeBSD14 indicates FreeBSD 14.
	VirtualMachineGuestOsIdentifierFreeBSD14 VirtualMachineGuestOsIdentifier = "FreeBSD14"

	// VirtualMachineGuestOsIdentifierFreeBSD14x64 indicates FreeBSD 14
	// (64 bit).
	VirtualMachineGuestOsIdentifierFreeBSD14x64 VirtualMachineGuestOsIdentifier = "FreeBSD14x64"

	// VirtualMachineGuestOsIdentifierFreeBSD15 indicates FreeBSD 15.
	VirtualMachineGuestOsIdentifierFreeBSD15 VirtualMachineGuestOsIdentifier = "FreeBSD15"

	// VirtualMachineGuestOsIdentifierFreeBSD15x64 indicates FreeBSD 15
	// (64 bit).
	VirtualMachineGuestOsIdentifierFreeBSD15x64 VirtualMachineGuestOsIdentifier = "FreeBSD15x64"

	// Red Hat

	// VirtualMachineGuestOsIdentifierRedHat indicates Red Hat Linux 2.1.
	VirtualMachineGuestOsIdentifierRedHat VirtualMachineGuestOsIdentifier = "RedHat"

	// VirtualMachineGuestOsIdentifierRHEL2 indicates Red Hat Enterprise
	// Linux 2.
	VirtualMachineGuestOsIdentifierRHEL2 VirtualMachineGuestOsIdentifier = "RHEL2"

	// VirtualMachineGuestOsIdentifierRHEL3 indicates Red Hat Enterprise
	// Linux 3.
	VirtualMachineGuestOsIdentifierRHEL3 VirtualMachineGuestOsIdentifier = "RHEL3"

	// VirtualMachineGuestOsIdentifierRHEL3x64 indicates Red Hat Enterprise
	// Linux 3 (64 bit).
	VirtualMachineGuestOsIdentifierRHEL3x64 VirtualMachineGuestOsIdentifier = "RHEL3x64"

	// VirtualMachineGuestOsIdentifierRHEL4 indicates Red Hat Enterprise
	// Linux 4.
	VirtualMachineGuestOsIdentifierRHEL4 VirtualMachineGuestOsIdentifier = "RHEL4"

	// VirtualMachineGuestOsIdentifierRHEL4x64 indicates Red Hat Enterprise
	// Linux 4 (64 bit).
	VirtualMachineGuestOsIdentifierRHEL4x64 VirtualMachineGuestOsIdentifier = "RHEL4x64"

	// VirtualMachineGuestOsIdentifierRHEL5 indicates Red Hat Enterprise
	// Linux 5.
	VirtualMachineGuestOsIdentifierRHEL5 VirtualMachineGuestOsIdentifier = "RHEL5"

	// VirtualMachineGuestOsIdentifierRHEL5x64 indicates Red Hat Enterprise
	// Linux 5 (64 bit).
	VirtualMachineGuestOsIdentifierRHEL5x64 VirtualMachineGuestOsIdentifier = "RHEL5x64"

	// VirtualMachineGuestOsIdentifierRHEL6 indicates Red Hat Enterprise
	// Linux 6.
	VirtualMachineGuestOsIdentifierRHEL6 VirtualMachineGuestOsIdentifier = "RHEL6"

	// VirtualMachineGuestOsIdentifierRHEL6x64 indicates Red Hat Enterprise
	// Linux 6 (64 bit).
	VirtualMachineGuestOsIdentifierRHEL6x64 VirtualMachineGuestOsIdentifier = "RHEL6x64"

	// VirtualMachineGuestOsIdentifierRHEL7 indicates Red Hat Enterprise
	// Linux 7.
	VirtualMachineGuestOsIdentifierRHEL7 VirtualMachineGuestOsIdentifier = "RHEL7"

	// VirtualMachineGuestOsIdentifierRHEL7x64 indicates Red Hat Enterprise
	// Linux 7 (64 bit).
	VirtualMachineGuestOsIdentifierRHEL7x64 VirtualMachineGuestOsIdentifier = "RHEL7x64"

	// VirtualMachineGuestOsIdentifierRHEL8x64 indicates Red Hat Enterprise
	// Linux 8 (64 bit).
	VirtualMachineGuestOsIdentifierRHEL8x64 VirtualMachineGuestOsIdentifier = "RHEL8x64"

	// VirtualMachineGuestOsIdentifierRHEL9x64 indicates Red Hat Enterprise
	// Linux 9 (64 bit).
	VirtualMachineGuestOsIdentifierRHEL9x64 VirtualMachineGuestOsIdentifier = "RHEL9x64"

	// VirtualMachineGuestOsIdentifierRHEL10x64 indicates Red Hat Enterprise
	// Linux 10 (64 bit).
	VirtualMachineGuestOsIdentifierRHEL10x64 VirtualMachineGuestOsIdentifier = "RHEL10x64"

	// CentOS

	// VirtualMachineGuestOsIdentifierCentOS indicates CentOS 4/5.
	VirtualMachineGuestOsIdentifierCentOS VirtualMachineGuestOsIdentifier = "CentOS"

	// VirtualMachineGuestOsIdentifierCentOS64 indicates CentOS 4/5
	// (64 bit).
	VirtualMachineGuestOsIdentifierCentOS64 VirtualMachineGuestOsIdentifier = "CentOS64"

	// VirtualMachineGuestOsIdentifierCentOS6 indicates CentOS 6.
	VirtualMachineGuestOsIdentifierCentOS6 VirtualMachineGuestOsIdentifier = "CentOS6"

	// VirtualMachineGuestOsIdentifierCentOS6x64 indicates CentOS 6
	// (64 bit).
	VirtualMachineGuestOsIdentifierCentOS6x64 VirtualMachineGuestOsIdentifier = "CentOS6x64"

	// VirtualMachineGuestOsIdentifierCentOS7 indicates CentOS 7.
	VirtualMachineGuestOsIdentifierCentOS7 VirtualMachineGuestOsIdentifier = "CentOS7"

	// VirtualMachineGuestOsIdentifierCentOS7x64 indicates CentOS 7
	// (64 bit).
	VirtualMachineGuestOsIdentifierCentOS7x64 VirtualMachineGuestOsIdentifier = "CentOS7x64"

	// VirtualMachineGuestOsIdentifierCentOS8x64 indicates CentOS 8
	// (64 bit).
	VirtualMachineGuestOsIdentifierCentOS8x64 VirtualMachineGuestOsIdentifier = "CentOS8x64"

	// VirtualMachineGuestOsIdentifierCentOS9x64 indicates CentOS 9
	// (64 bit).
	VirtualMachineGuestOsIdentifierCentOS9x64 VirtualMachineGuestOsIdentifier = "CentOS9x64"

	// Oracle Linux

	// VirtualMachineGuestOsIdentifierOracleLinux indicates Oracle Linux
	// 4/5.
	VirtualMachineGuestOsIdentifierOracleLinux VirtualMachineGuestOsIdentifier = "OracleLinux"

	// VirtualMachineGuestOsIdentifierOracleLinux64 indicates Oracle Linux
	// 4/5 (64 bit).
	VirtualMachineGuestOsIdentifierOracleLinux64 VirtualMachineGuestOsIdentifier = "OracleLinux64"

	// VirtualMachineGuestOsIdentifierOracleLinux6 indicates Oracle Linux 6.
	VirtualMachineGuestOsIdentifierOracleLinux6 VirtualMachineGuestOsIdentifier = "OracleLinux6"

	// VirtualMachineGuestOsIdentifierOracleLinux6x64 indicates Oracle Linux
	// 6 (64 bit).
	VirtualMachineGuestOsIdentifierOracleLinux6x64 VirtualMachineGuestOsIdentifier = "OracleLinux6x64"

	// VirtualMachineGuestOsIdentifierOracleLinux7 indicates Oracle Linux 7.
	VirtualMachineGuestOsIdentifierOracleLinux7 VirtualMachineGuestOsIdentifier = "OracleLinux7"

	// VirtualMachineGuestOsIdentifierOracleLinux7x64 indicates Oracle Linux
	// 7 (64 bit).
	VirtualMachineGuestOsIdentifierOracleLinux7x64 VirtualMachineGuestOsIdentifier = "OracleLinux7x64"

	// VirtualMachineGuestOsIdentifierOracleLinux8x64 indicates Oracle Linux
	// 8 (64 bit).
	VirtualMachineGuestOsIdentifierOracleLinux8x64 VirtualMachineGuestOsIdentifier = "OracleLinux8x64"

	// VirtualMachineGuestOsIdentifierOracleLinux9x64 indicates Oracle Linux
	// 9 (64 bit).
	VirtualMachineGuestOsIdentifierOracleLinux9x64 VirtualMachineGuestOsIdentifier = "OracleLinux9x64"

	// VirtualMachineGuestOsIdentifierOracleLinux10x64 indicates Oracle
	// Linux 10 (64 bit).
	VirtualMachineGuestOsIdentifierOracleLinux10x64 VirtualMachineGuestOsIdentifier = "OracleLinux10x64"

	// SUSE

	// VirtualMachineGuestOsIdentifierSUSE indicates SUSE Linux.
	VirtualMachineGuestOsIdentifierSUSE VirtualMachineGuestOsIdentifier = "SUSE"

	// VirtualMachineGuestOsIdentifierSUSE64 indicates SUSE Linux (64 bit).
	VirtualMachineGuestOsIdentifierSUSE64 VirtualMachineGuestOsIdentifier = "SUSE64"

	// VirtualMachineGuestOsIdentifierSLES indicates SUSE Linux Enterprise
	// Server 9.
	VirtualMachineGuestOsIdentifierSLES VirtualMachineGuestOsIdentifier = "SLES"

	// VirtualMachineGuestOsIdentifierSLES64 indicates SUSE Linux Enterprise
	// Server 9 (64 bit).
	VirtualMachineGuestOsIdentifierSLES64 VirtualMachineGuestOsIdentifier = "SLES64"

	// VirtualMachineGuestOsIdentifierSLES10 indicates SUSE Linux Enterprise
	// Server 10.
	VirtualMachineGuestOsIdentifierSLES10 VirtualMachineGuestOsIdentifier = "SLES10"

	// VirtualMachineGuestOsIdentifierSLES10x64 indicates SUSE Linux
	// Enterprise Server 10 (64 bit).
	VirtualMachineGuestOsIdentifierSLES10x64 VirtualMachineGuestOsIdentifier = "SLES10x64"

	// VirtualMachineGuestOsIdentifierSLES11 indicates SUSE Linux Enterprise
	// Server 11.
	VirtualMachineGuestOsIdentifierSLES11 VirtualMachineGuestOsIdentifier = "SLES11"

	// VirtualMachineGuestOsIdentifierSLES11x64 indicates SUSE Linux
	// Enterprise Server 11 (64 bit).
	VirtualMachineGuestOsIdentifierSLES11x64 VirtualMachineGuestOsIdentifier = "SLES11x64"

	// VirtualMachineGuestOsIdentifierSLES12 indicates SUSE Linux Enterprise
	// Server 12.
	VirtualMachineGuestOsIdentifierSLES12 VirtualMachineGuestOsIdentifier = "SLES12"

	// VirtualMachineGuestOsIdentifierSLES12x64 indicates SUSE Linux
	// Enterprise Server 12 (64 bit).
	VirtualMachineGuestOsIdentifierSLES12x64 VirtualMachineGuestOsIdentifier = "SLES12x64"

	// VirtualMachineGuestOsIdentifierSLES15x64 indicates SUSE Linux
	// Enterprise Server 15 (64 bit).
	VirtualMachineGuestOsIdentifierSLES15x64 VirtualMachineGuestOsIdentifier = "SLES15x64"

	// VirtualMachineGuestOsIdentifierSLES16x64 indicates SUSE Linux
	// Enterprise Server 16 (64 bit).
	VirtualMachineGuestOsIdentifierSLES16x64 VirtualMachineGuestOsIdentifier = "SLES16x64"

	// Novell

	// VirtualMachineGuestOsIdentifierNLD9 indicates Novell Linux Desktop 9.
	VirtualMachineGuestOsIdentifierNLD9 VirtualMachineGuestOsIdentifier = "NLD9"

	// VirtualMachineGuestOsIdentifierOES indicates Open Enterprise Server.
	VirtualMachineGuestOsIdentifierOES VirtualMachineGuestOsIdentifier = "OES"

	// VirtualMachineGuestOsIdentifierSJDS indicates Sun Java Desktop System.
	VirtualMachineGuestOsIdentifierSJDS VirtualMachineGuestOsIdentifier = "SJDS"

	// Mandriva / Mandrake

	// VirtualMachineGuestOsIdentifierMandrake indicates Mandrake Linux.
	VirtualMachineGuestOsIdentifierMandrake VirtualMachineGuestOsIdentifier = "Mandrake"

	// VirtualMachineGuestOsIdentifierMandriva indicates Mandriva Linux.
	VirtualMachineGuestOsIdentifierMandriva VirtualMachineGuestOsIdentifier = "Mandriva"

	// VirtualMachineGuestOsIdentifierMandriva64 indicates Mandriva Linux
	// (64 bit).
	VirtualMachineGuestOsIdentifierMandriva64 VirtualMachineGuestOsIdentifier = "Mandriva64"

	// TurboLinux

	// VirtualMachineGuestOsIdentifierTurboLinux indicates Turbolinux.
	VirtualMachineGuestOsIdentifierTurboLinux VirtualMachineGuestOsIdentifier = "TurboLinux"

	// VirtualMachineGuestOsIdentifierTurboLinux64 indicates Turbolinux
	// (64 bit).
	VirtualMachineGuestOsIdentifierTurboLinux64 VirtualMachineGuestOsIdentifier = "TurboLinux64"

	// Ubuntu

	// VirtualMachineGuestOsIdentifierUbuntu indicates Ubuntu Linux.
	VirtualMachineGuestOsIdentifierUbuntu VirtualMachineGuestOsIdentifier = "Ubuntu"

	// VirtualMachineGuestOsIdentifierUbuntu64 indicates Ubuntu Linux
	// (64 bit).
	VirtualMachineGuestOsIdentifierUbuntu64 VirtualMachineGuestOsIdentifier = "Ubuntu64"

	// Debian

	// VirtualMachineGuestOsIdentifierDebian4 indicates Debian GNU/Linux 4.
	VirtualMachineGuestOsIdentifierDebian4 VirtualMachineGuestOsIdentifier = "Debian4"

	// VirtualMachineGuestOsIdentifierDebian4x64 indicates Debian GNU/Linux
	// 4 (64 bit).
	VirtualMachineGuestOsIdentifierDebian4x64 VirtualMachineGuestOsIdentifier = "Debian4x64"

	// VirtualMachineGuestOsIdentifierDebian5 indicates Debian GNU/Linux 5.
	VirtualMachineGuestOsIdentifierDebian5 VirtualMachineGuestOsIdentifier = "Debian5"

	// VirtualMachineGuestOsIdentifierDebian5x64 indicates Debian GNU/Linux
	// 5 (64 bit).
	VirtualMachineGuestOsIdentifierDebian5x64 VirtualMachineGuestOsIdentifier = "Debian5x64"

	// VirtualMachineGuestOsIdentifierDebian6 indicates Debian GNU/Linux 6.
	VirtualMachineGuestOsIdentifierDebian6 VirtualMachineGuestOsIdentifier = "Debian6"

	// VirtualMachineGuestOsIdentifierDebian6x64 indicates Debian GNU/Linux
	// 6 (64 bit).
	VirtualMachineGuestOsIdentifierDebian6x64 VirtualMachineGuestOsIdentifier = "Debian6x64"

	// VirtualMachineGuestOsIdentifierDebian7 indicates Debian GNU/Linux 7.
	VirtualMachineGuestOsIdentifierDebian7 VirtualMachineGuestOsIdentifier = "Debian7"

	// VirtualMachineGuestOsIdentifierDebian7x64 indicates Debian GNU/Linux
	// 7 (64 bit).
	VirtualMachineGuestOsIdentifierDebian7x64 VirtualMachineGuestOsIdentifier = "Debian7x64"

	// VirtualMachineGuestOsIdentifierDebian8 indicates Debian GNU/Linux 8.
	VirtualMachineGuestOsIdentifierDebian8 VirtualMachineGuestOsIdentifier = "Debian8"

	// VirtualMachineGuestOsIdentifierDebian8x64 indicates Debian GNU/Linux
	// 8 (64 bit).
	VirtualMachineGuestOsIdentifierDebian8x64 VirtualMachineGuestOsIdentifier = "Debian8x64"

	// VirtualMachineGuestOsIdentifierDebian9 indicates Debian GNU/Linux 9.
	VirtualMachineGuestOsIdentifierDebian9 VirtualMachineGuestOsIdentifier = "Debian9"

	// VirtualMachineGuestOsIdentifierDebian9x64 indicates Debian GNU/Linux
	// 9 (64 bit).
	VirtualMachineGuestOsIdentifierDebian9x64 VirtualMachineGuestOsIdentifier = "Debian9x64"

	// VirtualMachineGuestOsIdentifierDebian10 indicates Debian GNU/Linux
	// 10.
	VirtualMachineGuestOsIdentifierDebian10 VirtualMachineGuestOsIdentifier = "Debian10"

	// VirtualMachineGuestOsIdentifierDebian10x64 indicates Debian GNU/Linux
	// 10 (64 bit).
	VirtualMachineGuestOsIdentifierDebian10x64 VirtualMachineGuestOsIdentifier = "Debian10x64"

	// VirtualMachineGuestOsIdentifierDebian11 indicates Debian GNU/Linux
	// 11.
	VirtualMachineGuestOsIdentifierDebian11 VirtualMachineGuestOsIdentifier = "Debian11"

	// VirtualMachineGuestOsIdentifierDebian11x64 indicates Debian GNU/Linux
	// 11 (64 bit).
	VirtualMachineGuestOsIdentifierDebian11x64 VirtualMachineGuestOsIdentifier = "Debian11x64"

	// VirtualMachineGuestOsIdentifierDebian12 indicates Debian GNU/Linux
	// 12.
	VirtualMachineGuestOsIdentifierDebian12 VirtualMachineGuestOsIdentifier = "Debian12"

	// VirtualMachineGuestOsIdentifierDebian12x64 indicates Debian GNU/Linux
	// 12 (64 bit).
	VirtualMachineGuestOsIdentifierDebian12x64 VirtualMachineGuestOsIdentifier = "Debian12x64"

	// VirtualMachineGuestOsIdentifierDebian13 indicates Debian GNU/Linux
	// 13.
	VirtualMachineGuestOsIdentifierDebian13 VirtualMachineGuestOsIdentifier = "Debian13"

	// VirtualMachineGuestOsIdentifierDebian13x64 indicates Debian GNU/Linux
	// 13 (64 bit).
	VirtualMachineGuestOsIdentifierDebian13x64 VirtualMachineGuestOsIdentifier = "Debian13x64"

	// Asianux

	// VirtualMachineGuestOsIdentifierAsianux3 indicates Asianux Server 3.
	VirtualMachineGuestOsIdentifierAsianux3 VirtualMachineGuestOsIdentifier = "Asianux3"

	// VirtualMachineGuestOsIdentifierAsianux3x64 indicates Asianux Server
	// 3 (64 bit).
	VirtualMachineGuestOsIdentifierAsianux3x64 VirtualMachineGuestOsIdentifier = "Asianux3x64"

	// VirtualMachineGuestOsIdentifierAsianux4 indicates Asianux Server 4.
	VirtualMachineGuestOsIdentifierAsianux4 VirtualMachineGuestOsIdentifier = "Asianux4"

	// VirtualMachineGuestOsIdentifierAsianux4x64 indicates Asianux Server
	// 4 (64 bit).
	VirtualMachineGuestOsIdentifierAsianux4x64 VirtualMachineGuestOsIdentifier = "Asianux4x64"

	// VirtualMachineGuestOsIdentifierAsianux5x64 indicates Asianux Server
	// 5 (64 bit).
	VirtualMachineGuestOsIdentifierAsianux5x64 VirtualMachineGuestOsIdentifier = "Asianux5x64"

	// VirtualMachineGuestOsIdentifierAsianux7x64 indicates Asianux Server
	// 7 (64 bit).
	VirtualMachineGuestOsIdentifierAsianux7x64 VirtualMachineGuestOsIdentifier = "Asianux7x64"

	// VirtualMachineGuestOsIdentifierAsianux8x64 indicates Asianux Server
	// 8 (64 bit).
	VirtualMachineGuestOsIdentifierAsianux8x64 VirtualMachineGuestOsIdentifier = "Asianux8x64"

	// VirtualMachineGuestOsIdentifierAsianux9x64 indicates Asianux Server
	// 9 (64 bit).
	VirtualMachineGuestOsIdentifierAsianux9x64 VirtualMachineGuestOsIdentifier = "Asianux9x64"

	// VirtualMachineGuestOsIdentifierMiracleLinux64 indicates MIRACLE LINUX
	// (64 bit).
	VirtualMachineGuestOsIdentifierMiracleLinux64 VirtualMachineGuestOsIdentifier = "MiracleLinux64"

	// VirtualMachineGuestOsIdentifierPardus64 indicates Pardus (64 bit).
	VirtualMachineGuestOsIdentifierPardus64 VirtualMachineGuestOsIdentifier = "Pardus64"

	// OpenSUSE

	// VirtualMachineGuestOsIdentifierOpenSUSE indicates OpenSUSE Linux.
	VirtualMachineGuestOsIdentifierOpenSUSE VirtualMachineGuestOsIdentifier = "OpenSUSE"

	// VirtualMachineGuestOsIdentifierOpenSUSE64 indicates OpenSUSE Linux
	// (64 bit).
	VirtualMachineGuestOsIdentifierOpenSUSE64 VirtualMachineGuestOsIdentifier = "OpenSUSE64"

	// Fedora

	// VirtualMachineGuestOsIdentifierFedora indicates Fedora Linux.
	VirtualMachineGuestOsIdentifierFedora VirtualMachineGuestOsIdentifier = "Fedora"

	// VirtualMachineGuestOsIdentifierFedora64 indicates Fedora Linux
	// (64 bit).
	VirtualMachineGuestOsIdentifierFedora64 VirtualMachineGuestOsIdentifier = "Fedora64"

	// CoreOS / Photon

	// VirtualMachineGuestOsIdentifierCoreOS64 indicates CoreOS Linux
	// (64 bit).
	VirtualMachineGuestOsIdentifierCoreOS64 VirtualMachineGuestOsIdentifier = "CoreOS64"

	// VirtualMachineGuestOsIdentifierVMwarePhoton64 indicates VMware Photon
	// (64 bit).
	VirtualMachineGuestOsIdentifierVMwarePhoton64 VirtualMachineGuestOsIdentifier = "VMwarePhoton64"

	// Other Linux

	// VirtualMachineGuestOsIdentifierOther24xLinux indicates Linux 2.4.x
	// Kernel.
	VirtualMachineGuestOsIdentifierOther24xLinux VirtualMachineGuestOsIdentifier = "Other24xLinux"

	// VirtualMachineGuestOsIdentifierOther26xLinux indicates Linux 2.6.x
	// Kernel.
	VirtualMachineGuestOsIdentifierOther26xLinux VirtualMachineGuestOsIdentifier = "Other26xLinux"

	// VirtualMachineGuestOsIdentifierOtherLinux indicates Linux 2.2.x
	// Kernel.
	VirtualMachineGuestOsIdentifierOtherLinux VirtualMachineGuestOsIdentifier = "OtherLinux"

	// VirtualMachineGuestOsIdentifierOther3xLinux indicates Linux 3.x
	// Kernel.
	VirtualMachineGuestOsIdentifierOther3xLinux VirtualMachineGuestOsIdentifier = "Other3xLinux"

	// VirtualMachineGuestOsIdentifierOther4xLinux indicates Linux 4.x
	// Kernel.
	VirtualMachineGuestOsIdentifierOther4xLinux VirtualMachineGuestOsIdentifier = "Other4xLinux"

	// VirtualMachineGuestOsIdentifierOther5xLinux indicates Linux 5.x
	// Kernel.
	VirtualMachineGuestOsIdentifierOther5xLinux VirtualMachineGuestOsIdentifier = "Other5xLinux"

	// VirtualMachineGuestOsIdentifierOther6xLinux indicates Linux 6.x
	// Kernel.
	VirtualMachineGuestOsIdentifierOther6xLinux VirtualMachineGuestOsIdentifier = "Other6xLinux"

	// VirtualMachineGuestOsIdentifierOther7xLinux indicates Linux 7.x
	// Kernel.
	VirtualMachineGuestOsIdentifierOther7xLinux VirtualMachineGuestOsIdentifier = "Other7xLinux"

	// VirtualMachineGuestOsIdentifierGenericLinux indicates a generic Linux
	// guest.
	VirtualMachineGuestOsIdentifierGenericLinux VirtualMachineGuestOsIdentifier = "GenericLinux"

	// VirtualMachineGuestOsIdentifierOther24xLinux64 indicates Linux 2.4.x
	// Kernel (64 bit).
	VirtualMachineGuestOsIdentifierOther24xLinux64 VirtualMachineGuestOsIdentifier = "Other24xLinux64"

	// VirtualMachineGuestOsIdentifierOther26xLinux64 indicates Linux 2.6.x
	// Kernel (64 bit).
	VirtualMachineGuestOsIdentifierOther26xLinux64 VirtualMachineGuestOsIdentifier = "Other26xLinux64"

	// VirtualMachineGuestOsIdentifierOther3xLinux64 indicates Linux 3.x
	// Kernel (64 bit).
	VirtualMachineGuestOsIdentifierOther3xLinux64 VirtualMachineGuestOsIdentifier = "Other3xLinux64"

	// VirtualMachineGuestOsIdentifierOther4xLinux64 indicates Linux 4.x
	// Kernel (64 bit).
	VirtualMachineGuestOsIdentifierOther4xLinux64 VirtualMachineGuestOsIdentifier = "Other4xLinux64"

	// VirtualMachineGuestOsIdentifierOther5xLinux64 indicates Linux 5.x
	// Kernel (64 bit).
	VirtualMachineGuestOsIdentifierOther5xLinux64 VirtualMachineGuestOsIdentifier = "Other5xLinux64"

	// VirtualMachineGuestOsIdentifierOther6xLinux64 indicates Linux 6.x
	// Kernel (64 bit).
	VirtualMachineGuestOsIdentifierOther6xLinux64 VirtualMachineGuestOsIdentifier = "Other6xLinux64"

	// VirtualMachineGuestOsIdentifierOther7xLinux64 indicates Linux 7.x
	// Kernel (64 bit).
	VirtualMachineGuestOsIdentifierOther7xLinux64 VirtualMachineGuestOsIdentifier = "Other7xLinux64"

	// VirtualMachineGuestOsIdentifierOtherLinux64 indicates Linux 2.2.x
	// Kernel (64 bit).
	VirtualMachineGuestOsIdentifierOtherLinux64 VirtualMachineGuestOsIdentifier = "OtherLinux64"

	// Solaris

	// VirtualMachineGuestOsIdentifierSolaris6 indicates Solaris 6.
	VirtualMachineGuestOsIdentifierSolaris6 VirtualMachineGuestOsIdentifier = "Solaris6"

	// VirtualMachineGuestOsIdentifierSolaris7 indicates Solaris 7.
	VirtualMachineGuestOsIdentifierSolaris7 VirtualMachineGuestOsIdentifier = "Solaris7"

	// VirtualMachineGuestOsIdentifierSolaris8 indicates Solaris 8.
	VirtualMachineGuestOsIdentifierSolaris8 VirtualMachineGuestOsIdentifier = "Solaris8"

	// VirtualMachineGuestOsIdentifierSolaris9 indicates Solaris 9.
	VirtualMachineGuestOsIdentifierSolaris9 VirtualMachineGuestOsIdentifier = "Solaris9"

	// VirtualMachineGuestOsIdentifierSolaris10 indicates Solaris 10.
	VirtualMachineGuestOsIdentifierSolaris10 VirtualMachineGuestOsIdentifier = "Solaris10"

	// VirtualMachineGuestOsIdentifierSolaris10x64 indicates Solaris 10
	// (64 bit).
	VirtualMachineGuestOsIdentifierSolaris10x64 VirtualMachineGuestOsIdentifier = "Solaris10x64"

	// VirtualMachineGuestOsIdentifierSolaris11x64 indicates Solaris 11
	// (64 bit).
	VirtualMachineGuestOsIdentifierSolaris11x64 VirtualMachineGuestOsIdentifier = "Solaris11x64"

	// Chinese Linux distributions

	// VirtualMachineGuestOsIdentifierFusionOS64 indicates Fusion OS
	// (64 bit).
	VirtualMachineGuestOsIdentifierFusionOS64 VirtualMachineGuestOsIdentifier = "FusionOS64"

	// VirtualMachineGuestOsIdentifierProLinux64 indicates Pro Linux
	// (64 bit).
	VirtualMachineGuestOsIdentifierProLinux64 VirtualMachineGuestOsIdentifier = "ProLinux64"

	// VirtualMachineGuestOsIdentifierKylinLinux64 indicates Kylin Linux
	// (64 bit).
	VirtualMachineGuestOsIdentifierKylinLinux64 VirtualMachineGuestOsIdentifier = "KylinLinux64"

	// OS/2 and eComStation

	// VirtualMachineGuestOsIdentifierOS2 indicates OS/2.
	VirtualMachineGuestOsIdentifierOS2 VirtualMachineGuestOsIdentifier = "OS2"

	// VirtualMachineGuestOsIdentifierEComStation indicates eComStation 1.x.
	VirtualMachineGuestOsIdentifierEComStation VirtualMachineGuestOsIdentifier = "EComStation"

	// VirtualMachineGuestOsIdentifierEComStation2 indicates eComStation 2.x.
	VirtualMachineGuestOsIdentifierEComStation2 VirtualMachineGuestOsIdentifier = "EComStation2"

	// NetWare

	// VirtualMachineGuestOsIdentifierNetWare4 indicates Novell NetWare 4.
	VirtualMachineGuestOsIdentifierNetWare4 VirtualMachineGuestOsIdentifier = "NetWare4"

	// VirtualMachineGuestOsIdentifierNetWare5 indicates Novell NetWare 5.1.
	VirtualMachineGuestOsIdentifierNetWare5 VirtualMachineGuestOsIdentifier = "NetWare5"

	// VirtualMachineGuestOsIdentifierNetWare6 indicates Novell NetWare 6.x.
	VirtualMachineGuestOsIdentifierNetWare6 VirtualMachineGuestOsIdentifier = "NetWare6"

	// SCO

	// VirtualMachineGuestOsIdentifierOpenServer5 indicates SCO OpenServer 5.
	VirtualMachineGuestOsIdentifierOpenServer5 VirtualMachineGuestOsIdentifier = "OpenServer5"

	// VirtualMachineGuestOsIdentifierOpenServer6 indicates SCO OpenServer 6.
	VirtualMachineGuestOsIdentifierOpenServer6 VirtualMachineGuestOsIdentifier = "OpenServer6"

	// VirtualMachineGuestOsIdentifierUnixWare7 indicates SCO UnixWare 7.
	VirtualMachineGuestOsIdentifierUnixWare7 VirtualMachineGuestOsIdentifier = "UnixWare7"

	// macOS / Darwin

	// VirtualMachineGuestOsIdentifierDarwin indicates Apple macOS 10.5.
	VirtualMachineGuestOsIdentifierDarwin VirtualMachineGuestOsIdentifier = "Darwin"

	// VirtualMachineGuestOsIdentifierDarwin64 indicates Apple macOS 10.5
	// (64 bit).
	VirtualMachineGuestOsIdentifierDarwin64 VirtualMachineGuestOsIdentifier = "Darwin64"

	// VirtualMachineGuestOsIdentifierDarwin10 indicates Apple macOS 10.6.
	VirtualMachineGuestOsIdentifierDarwin10 VirtualMachineGuestOsIdentifier = "Darwin10"

	// VirtualMachineGuestOsIdentifierDarwin10x64 indicates Apple macOS
	// 10.6 (64 bit).
	VirtualMachineGuestOsIdentifierDarwin10x64 VirtualMachineGuestOsIdentifier = "Darwin10x64"

	// VirtualMachineGuestOsIdentifierDarwin11 indicates Apple macOS 10.7.
	VirtualMachineGuestOsIdentifierDarwin11 VirtualMachineGuestOsIdentifier = "Darwin11"

	// VirtualMachineGuestOsIdentifierDarwin11x64 indicates Apple macOS
	// 10.7 (64 bit).
	VirtualMachineGuestOsIdentifierDarwin11x64 VirtualMachineGuestOsIdentifier = "Darwin11x64"

	// VirtualMachineGuestOsIdentifierDarwin12x64 indicates Apple macOS
	// 10.8 (64 bit).
	VirtualMachineGuestOsIdentifierDarwin12x64 VirtualMachineGuestOsIdentifier = "Darwin12x64"

	// VirtualMachineGuestOsIdentifierDarwin13x64 indicates Apple macOS
	// 10.9 (64 bit).
	VirtualMachineGuestOsIdentifierDarwin13x64 VirtualMachineGuestOsIdentifier = "Darwin13x64"

	// VirtualMachineGuestOsIdentifierDarwin14x64 indicates Apple macOS
	// 10.10 (64 bit).
	VirtualMachineGuestOsIdentifierDarwin14x64 VirtualMachineGuestOsIdentifier = "Darwin14x64"

	// VirtualMachineGuestOsIdentifierDarwin15x64 indicates Apple macOS
	// 10.11 (64 bit).
	VirtualMachineGuestOsIdentifierDarwin15x64 VirtualMachineGuestOsIdentifier = "Darwin15x64"

	// VirtualMachineGuestOsIdentifierDarwin16x64 indicates Apple macOS
	// 10.12 (64 bit).
	VirtualMachineGuestOsIdentifierDarwin16x64 VirtualMachineGuestOsIdentifier = "Darwin16x64"

	// VirtualMachineGuestOsIdentifierDarwin17x64 indicates Apple macOS
	// 10.13 (64 bit).
	VirtualMachineGuestOsIdentifierDarwin17x64 VirtualMachineGuestOsIdentifier = "Darwin17x64"

	// VirtualMachineGuestOsIdentifierDarwin18x64 indicates Apple macOS
	// 10.14 (64 bit).
	VirtualMachineGuestOsIdentifierDarwin18x64 VirtualMachineGuestOsIdentifier = "Darwin18x64"

	// VirtualMachineGuestOsIdentifierDarwin19x64 indicates Apple macOS
	// 10.15 (64 bit).
	VirtualMachineGuestOsIdentifierDarwin19x64 VirtualMachineGuestOsIdentifier = "Darwin19x64"

	// VirtualMachineGuestOsIdentifierDarwin20x64 indicates Apple macOS
	// 11 (64 bit).
	VirtualMachineGuestOsIdentifierDarwin20x64 VirtualMachineGuestOsIdentifier = "Darwin20x64"

	// VirtualMachineGuestOsIdentifierDarwin21x64 indicates Apple macOS
	// 12 (64 bit).
	VirtualMachineGuestOsIdentifierDarwin21x64 VirtualMachineGuestOsIdentifier = "Darwin21x64"

	// VirtualMachineGuestOsIdentifierDarwin22x64 indicates Apple macOS
	// 13 (64 bit).
	VirtualMachineGuestOsIdentifierDarwin22x64 VirtualMachineGuestOsIdentifier = "Darwin22x64"

	// VirtualMachineGuestOsIdentifierDarwin23x64 indicates Apple macOS
	// 14 (64 bit).
	VirtualMachineGuestOsIdentifierDarwin23x64 VirtualMachineGuestOsIdentifier = "Darwin23x64"

	// VMkernel / ESX

	// VirtualMachineGuestOsIdentifierVMkernel indicates VMware ESX 4.x.
	VirtualMachineGuestOsIdentifierVMkernel VirtualMachineGuestOsIdentifier = "VMkernel"

	// VirtualMachineGuestOsIdentifierVMkernel5 indicates VMware ESXi 5.x.
	VirtualMachineGuestOsIdentifierVMkernel5 VirtualMachineGuestOsIdentifier = "VMkernel5"

	// VirtualMachineGuestOsIdentifierVMkernel6 indicates VMware ESXi 6.x.
	VirtualMachineGuestOsIdentifierVMkernel6 VirtualMachineGuestOsIdentifier = "VMkernel6"

	// VirtualMachineGuestOsIdentifierVMkernel65 indicates VMware ESXi 6.5.
	VirtualMachineGuestOsIdentifierVMkernel65 VirtualMachineGuestOsIdentifier = "VMkernel65"

	// VirtualMachineGuestOsIdentifierVMkernel7 indicates VMware ESXi 7.x.
	VirtualMachineGuestOsIdentifierVMkernel7 VirtualMachineGuestOsIdentifier = "VMkernel7"

	// VirtualMachineGuestOsIdentifierVMkernel8 indicates VMware ESXi 8.x.
	VirtualMachineGuestOsIdentifierVMkernel8 VirtualMachineGuestOsIdentifier = "VMkernel8"

	// VirtualMachineGuestOsIdentifierVMkernel9 indicates VMware ESXi 9.x.
	VirtualMachineGuestOsIdentifierVMkernel9 VirtualMachineGuestOsIdentifier = "VMkernel9"

	// Amazon Linux

	// VirtualMachineGuestOsIdentifierAmazonLinux2x64 indicates Amazon Linux
	// 2 (64 bit).
	VirtualMachineGuestOsIdentifierAmazonLinux2x64 VirtualMachineGuestOsIdentifier = "AmazonLinux2x64"

	// VirtualMachineGuestOsIdentifierAmazonLinux3x64 indicates Amazon Linux
	// 3 (64 bit).
	VirtualMachineGuestOsIdentifierAmazonLinux3x64 VirtualMachineGuestOsIdentifier = "AmazonLinux3x64"

	// CRX

	// VirtualMachineGuestOsIdentifierCRXPod1 indicates CRX Pod 1.
	VirtualMachineGuestOsIdentifierCRXPod1 VirtualMachineGuestOsIdentifier = "CRXPod1"

	// VirtualMachineGuestOsIdentifierCRXSys1 indicates CRX Sys 1.
	VirtualMachineGuestOsIdentifierCRXSys1 VirtualMachineGuestOsIdentifier = "CRXSys1"

	// Rocky Linux / AlmaLinux

	// VirtualMachineGuestOsIdentifierRockyLinux64 indicates Rocky Linux
	// (64 bit).
	VirtualMachineGuestOsIdentifierRockyLinux64 VirtualMachineGuestOsIdentifier = "RockyLinux64"

	// VirtualMachineGuestOsIdentifierAlmaLinux64 indicates AlmaLinux
	// (64 bit).
	VirtualMachineGuestOsIdentifierAlmaLinux64 VirtualMachineGuestOsIdentifier = "AlmaLinux64"

	// Other

	// VirtualMachineGuestOsIdentifierOther indicates Other Operating System.
	VirtualMachineGuestOsIdentifierOther VirtualMachineGuestOsIdentifier = "Other"

	// VirtualMachineGuestOsIdentifierOther64 indicates Other Operating
	// System (64 bit).
	VirtualMachineGuestOsIdentifierOther64 VirtualMachineGuestOsIdentifier = "Other64"
)

// ToVimType returns the vSphere API identifier string for the
// VirtualMachineGuestOsIdentifier.
func (t VirtualMachineGuestOsIdentifier) ToVimType() string {
	s := string(t)
	switch {
	case s == "DOS",
		strings.HasPrefix(s, "Win"):
		return toVimTypeWindows(t)
	case strings.HasPrefix(s, "FreeBSD"):
		return toVimTypeFreeBSD(t)
	case strings.HasPrefix(s, "Solaris"):
		return toVimTypeSolaris(t)
	case strings.HasPrefix(s, "Darwin"):
		return toVimTypeDarwin(t)
	case strings.HasPrefix(s, "VMkernel"):
		return toVimTypeVMkernel(t)
	case strings.HasPrefix(s, "NetWare"):
		return toVimTypeNetWare(t)
	case strings.HasPrefix(s, "OpenServer"),
		strings.HasPrefix(s, "UnixWare"):
		return toVimTypeSCO(t)
	case s == "OS2",
		strings.HasPrefix(s, "EComStation"):
		return toVimTypeOS2(t)
	case strings.HasPrefix(s, "CRX"):
		return toVimTypeCRX(t)
	case s == "Other",
		s == "Other64":
		return toVimTypeOtherOS(t)
	default:
		return toVimTypeLinux(t)
	}
}

// toVimTypeWindows handles DOS and all Windows variants.
func toVimTypeWindows(t VirtualMachineGuestOsIdentifier) string {
	s := string(t)
	switch {
	case s == "DOS":
		return "dosGuest"
	case strings.HasPrefix(s, "Windows"):
		return toVimTypeWindowsModern(t)
	default:
		return toVimTypeWindowsLegacy(t)
	}
}

// toVimTypeWindowsLegacy handles Win3.1 through WinVista.
func toVimTypeWindowsLegacy(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierWin31:
		return "win31Guest"
	case VirtualMachineGuestOsIdentifierWin95:
		return "win95Guest"
	case VirtualMachineGuestOsIdentifierWin98:
		return "win98Guest"
	case VirtualMachineGuestOsIdentifierWinMe:
		return "winMeGuest"
	case VirtualMachineGuestOsIdentifierWinNT:
		return "winNTGuest"
	case VirtualMachineGuestOsIdentifierWin2000Pro:
		return "win2000ProGuest"
	case VirtualMachineGuestOsIdentifierWin2000Serv:
		return "win2000ServGuest"
	case VirtualMachineGuestOsIdentifierWin2000AdvServ:
		return "win2000AdvServGuest"
	case VirtualMachineGuestOsIdentifierWinXPHome:
		return "winXPHomeGuest"
	case VirtualMachineGuestOsIdentifierWinXPPro:
		return "winXPProGuest"
	case VirtualMachineGuestOsIdentifierWinXPPro64:
		return "winXPPro64Guest"
	case VirtualMachineGuestOsIdentifierWinNetWeb:
		return "winNetWebGuest"
	case VirtualMachineGuestOsIdentifierWinNetStandard:
		return "winNetStandardGuest"
	case VirtualMachineGuestOsIdentifierWinNetEnterprise:
		return "winNetEnterpriseGuest"
	case VirtualMachineGuestOsIdentifierWinNetDatacenter:
		return "winNetDatacenterGuest"
	case VirtualMachineGuestOsIdentifierWinNetBusiness:
		return "winNetBusinessGuest"
	case VirtualMachineGuestOsIdentifierWinNetStandard64:
		return "winNetStandard64Guest"
	case VirtualMachineGuestOsIdentifierWinNetEnterprise64:
		return "winNetEnterprise64Guest"
	case VirtualMachineGuestOsIdentifierWinLonghorn:
		return "winLonghornGuest"
	case VirtualMachineGuestOsIdentifierWinLonghorn64:
		return "winLonghorn64Guest"
	case VirtualMachineGuestOsIdentifierWinNetDatacenter64:
		return "winNetDatacenter64Guest"
	case VirtualMachineGuestOsIdentifierWinVista:
		return "winVistaGuest"
	case VirtualMachineGuestOsIdentifierWinVista64:
		return "winVista64Guest"
	}
	return string(t)
}

// toVimTypeWindowsModern handles Windows 7 through Windows Server 2025.
func toVimTypeWindowsModern(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierWindows7:
		return "windows7Guest"
	case VirtualMachineGuestOsIdentifierWindows7x64:
		return "windows7_64Guest"
	case VirtualMachineGuestOsIdentifierWindows7Server64:
		return "windows7Server64Guest"
	case VirtualMachineGuestOsIdentifierWindows8:
		return "windows8Guest"
	case VirtualMachineGuestOsIdentifierWindows8x64:
		return "windows8_64Guest"
	case VirtualMachineGuestOsIdentifierWindows8Server64:
		return "windows8Server64Guest"
	case VirtualMachineGuestOsIdentifierWindows9:
		return "windows9Guest"
	case VirtualMachineGuestOsIdentifierWindows9x64:
		return "windows9_64Guest"
	case VirtualMachineGuestOsIdentifierWindows9Server64:
		return "windows9Server64Guest"
	case VirtualMachineGuestOsIdentifierWindows11x64:
		return "windows11_64Guest"
	case VirtualMachineGuestOsIdentifierWindows12x64:
		return "windows12_64Guest"
	case VirtualMachineGuestOsIdentifierWindowsHyperV:
		return "windowsHyperVGuest"
	case VirtualMachineGuestOsIdentifierWindows2019Serverx64:
		return "windows2019srv_64Guest"
	case VirtualMachineGuestOsIdentifierWindows2019ServerNextx64:
		return "windows2019srvNext_64Guest"
	case VirtualMachineGuestOsIdentifierWindows2022ServerNextx64:
		return "windows2022srvNext_64Guest"
	}
	return string(t)
}

func toVimTypeFreeBSD(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierFreeBSD:
		return "freebsdGuest"
	case VirtualMachineGuestOsIdentifierFreeBSD64:
		return "freebsd64Guest"
	case VirtualMachineGuestOsIdentifierFreeBSD11:
		return "freebsd11Guest"
	case VirtualMachineGuestOsIdentifierFreeBSD11x64:
		return "freebsd11_64Guest"
	case VirtualMachineGuestOsIdentifierFreeBSD12:
		return "freebsd12Guest"
	case VirtualMachineGuestOsIdentifierFreeBSD12x64:
		return "freebsd12_64Guest"
	case VirtualMachineGuestOsIdentifierFreeBSD13:
		return "freebsd13Guest"
	case VirtualMachineGuestOsIdentifierFreeBSD13x64:
		return "freebsd13_64Guest"
	case VirtualMachineGuestOsIdentifierFreeBSD14:
		return "freebsd14Guest"
	case VirtualMachineGuestOsIdentifierFreeBSD14x64:
		return "freebsd14_64Guest"
	case VirtualMachineGuestOsIdentifierFreeBSD15:
		return "freebsd15Guest"
	case VirtualMachineGuestOsIdentifierFreeBSD15x64:
		return "freebsd15_64Guest"
	}
	return string(t)
}

// toVimTypeLinux dispatches to a distro-specific helper.
func toVimTypeLinux(t VirtualMachineGuestOsIdentifier) string {
	s := string(t)
	switch {
	case s == "RedHat",
		strings.HasPrefix(s, "RHEL"):
		return toVimTypeLinuxRHEL(t)
	case strings.HasPrefix(s, "CentOS"):
		return toVimTypeLinuxCentOS(t)
	case strings.HasPrefix(s, "OracleLinux"):
		return toVimTypeLinuxOracle(t)
	case strings.HasPrefix(s, "SUSE"),
		strings.HasPrefix(s, "SLES"),
		strings.HasPrefix(s, "OpenSUSE"):
		return toVimTypeLinuxSUSE(t)
	case strings.HasPrefix(s, "Debian"):
		return toVimTypeLinuxDebian(t)
	case strings.HasPrefix(s, "Asianux"),
		strings.HasPrefix(s, "MiracleLinux"),
		strings.HasPrefix(s, "Pardus"):
		return toVimTypeLinuxAsianux(t)
	case strings.HasPrefix(s, "Other"),
		s == "GenericLinux":
		return toVimTypeLinuxOtherKernel(t)
	default:
		return toVimTypeLinuxMisc(t)
	}
}

func toVimTypeLinuxRHEL(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierRedHat:
		return "redhatGuest"
	case VirtualMachineGuestOsIdentifierRHEL2:
		return "rhel2Guest"
	case VirtualMachineGuestOsIdentifierRHEL3:
		return "rhel3Guest"
	case VirtualMachineGuestOsIdentifierRHEL3x64:
		return "rhel3_64Guest"
	case VirtualMachineGuestOsIdentifierRHEL4:
		return "rhel4Guest"
	case VirtualMachineGuestOsIdentifierRHEL4x64:
		return "rhel4_64Guest"
	case VirtualMachineGuestOsIdentifierRHEL5:
		return "rhel5Guest"
	case VirtualMachineGuestOsIdentifierRHEL5x64:
		return "rhel5_64Guest"
	case VirtualMachineGuestOsIdentifierRHEL6:
		return "rhel6Guest"
	case VirtualMachineGuestOsIdentifierRHEL6x64:
		return "rhel6_64Guest"
	case VirtualMachineGuestOsIdentifierRHEL7:
		return "rhel7Guest"
	case VirtualMachineGuestOsIdentifierRHEL7x64:
		return "rhel7_64Guest"
	case VirtualMachineGuestOsIdentifierRHEL8x64:
		return "rhel8_64Guest"
	case VirtualMachineGuestOsIdentifierRHEL9x64:
		return "rhel9_64Guest"
	case VirtualMachineGuestOsIdentifierRHEL10x64:
		return "rhel10_64Guest"
	}
	return string(t)
}

func toVimTypeLinuxCentOS(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierCentOS:
		return "centosGuest"
	case VirtualMachineGuestOsIdentifierCentOS64:
		return "centos64Guest"
	case VirtualMachineGuestOsIdentifierCentOS6:
		return "centos6Guest"
	case VirtualMachineGuestOsIdentifierCentOS6x64:
		return "centos6_64Guest"
	case VirtualMachineGuestOsIdentifierCentOS7:
		return "centos7Guest"
	case VirtualMachineGuestOsIdentifierCentOS7x64:
		return "centos7_64Guest"
	case VirtualMachineGuestOsIdentifierCentOS8x64:
		return "centos8_64Guest"
	case VirtualMachineGuestOsIdentifierCentOS9x64:
		return "centos9_64Guest"
	}
	return string(t)
}

func toVimTypeLinuxOracle(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierOracleLinux:
		return "oracleLinuxGuest"
	case VirtualMachineGuestOsIdentifierOracleLinux64:
		return "oracleLinux64Guest"
	case VirtualMachineGuestOsIdentifierOracleLinux6:
		return "oracleLinux6Guest"
	case VirtualMachineGuestOsIdentifierOracleLinux6x64:
		return "oracleLinux6_64Guest"
	case VirtualMachineGuestOsIdentifierOracleLinux7:
		return "oracleLinux7Guest"
	case VirtualMachineGuestOsIdentifierOracleLinux7x64:
		return "oracleLinux7_64Guest"
	case VirtualMachineGuestOsIdentifierOracleLinux8x64:
		return "oracleLinux8_64Guest"
	case VirtualMachineGuestOsIdentifierOracleLinux9x64:
		return "oracleLinux9_64Guest"
	case VirtualMachineGuestOsIdentifierOracleLinux10x64:
		return "oracleLinux10_64Guest"
	}
	return string(t)
}

// toVimTypeLinuxSUSE handles SUSE, SLES, and OpenSUSE.
func toVimTypeLinuxSUSE(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierSUSE:
		return "suseGuest"
	case VirtualMachineGuestOsIdentifierSUSE64:
		return "suse64Guest"
	case VirtualMachineGuestOsIdentifierSLES:
		return "slesGuest"
	case VirtualMachineGuestOsIdentifierSLES64:
		return "sles64Guest"
	case VirtualMachineGuestOsIdentifierSLES10:
		return "sles10Guest"
	case VirtualMachineGuestOsIdentifierSLES10x64:
		return "sles10_64Guest"
	case VirtualMachineGuestOsIdentifierSLES11:
		return "sles11Guest"
	case VirtualMachineGuestOsIdentifierSLES11x64:
		return "sles11_64Guest"
	case VirtualMachineGuestOsIdentifierSLES12:
		return "sles12Guest"
	case VirtualMachineGuestOsIdentifierSLES12x64:
		return "sles12_64Guest"
	case VirtualMachineGuestOsIdentifierSLES15x64:
		return "sles15_64Guest"
	case VirtualMachineGuestOsIdentifierSLES16x64:
		return "sles16_64Guest"
	case VirtualMachineGuestOsIdentifierOpenSUSE:
		return "opensuseGuest"
	case VirtualMachineGuestOsIdentifierOpenSUSE64:
		return "opensuse64Guest"
	}
	return string(t)
}

func toVimTypeLinuxDebian(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierDebian4:
		return "debian4Guest"
	case VirtualMachineGuestOsIdentifierDebian4x64:
		return "debian4_64Guest"
	case VirtualMachineGuestOsIdentifierDebian5:
		return "debian5Guest"
	case VirtualMachineGuestOsIdentifierDebian5x64:
		return "debian5_64Guest"
	case VirtualMachineGuestOsIdentifierDebian6:
		return "debian6Guest"
	case VirtualMachineGuestOsIdentifierDebian6x64:
		return "debian6_64Guest"
	case VirtualMachineGuestOsIdentifierDebian7:
		return "debian7Guest"
	case VirtualMachineGuestOsIdentifierDebian7x64:
		return "debian7_64Guest"
	case VirtualMachineGuestOsIdentifierDebian8:
		return "debian8Guest"
	case VirtualMachineGuestOsIdentifierDebian8x64:
		return "debian8_64Guest"
	case VirtualMachineGuestOsIdentifierDebian9:
		return "debian9Guest"
	case VirtualMachineGuestOsIdentifierDebian9x64:
		return "debian9_64Guest"
	case VirtualMachineGuestOsIdentifierDebian10:
		return "debian10Guest"
	case VirtualMachineGuestOsIdentifierDebian10x64:
		return "debian10_64Guest"
	case VirtualMachineGuestOsIdentifierDebian11:
		return "debian11Guest"
	case VirtualMachineGuestOsIdentifierDebian11x64:
		return "debian11_64Guest"
	case VirtualMachineGuestOsIdentifierDebian12:
		return "debian12Guest"
	case VirtualMachineGuestOsIdentifierDebian12x64:
		return "debian12_64Guest"
	case VirtualMachineGuestOsIdentifierDebian13:
		return "debian13Guest"
	case VirtualMachineGuestOsIdentifierDebian13x64:
		return "debian13_64Guest"
	}
	return string(t)
}

// toVimTypeLinuxAsianux handles Asianux, MIRACLE LINUX, and Pardus.
func toVimTypeLinuxAsianux(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierAsianux3:
		return "asianux3Guest"
	case VirtualMachineGuestOsIdentifierAsianux3x64:
		return "asianux3_64Guest"
	case VirtualMachineGuestOsIdentifierAsianux4:
		return "asianux4Guest"
	case VirtualMachineGuestOsIdentifierAsianux4x64:
		return "asianux4_64Guest"
	case VirtualMachineGuestOsIdentifierAsianux5x64:
		return "asianux5_64Guest"
	case VirtualMachineGuestOsIdentifierAsianux7x64:
		return "asianux7_64Guest"
	case VirtualMachineGuestOsIdentifierAsianux8x64:
		return "asianux8_64Guest"
	case VirtualMachineGuestOsIdentifierAsianux9x64:
		return "asianux9_64Guest"
	case VirtualMachineGuestOsIdentifierMiracleLinux64:
		return "miraclelinux_64Guest"
	case VirtualMachineGuestOsIdentifierPardus64:
		return "pardus_64Guest"
	}
	return string(t)
}

// toVimTypeLinuxOtherKernel handles generic/other Linux kernel variants.
func toVimTypeLinuxOtherKernel(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierOther24xLinux:
		return "other24xLinuxGuest"
	case VirtualMachineGuestOsIdentifierOther26xLinux:
		return "other26xLinuxGuest"
	case VirtualMachineGuestOsIdentifierOtherLinux:
		return "otherLinuxGuest"
	case VirtualMachineGuestOsIdentifierOther3xLinux:
		return "other3xLinuxGuest"
	case VirtualMachineGuestOsIdentifierOther4xLinux:
		return "other4xLinuxGuest"
	case VirtualMachineGuestOsIdentifierOther5xLinux:
		return "other5xLinuxGuest"
	case VirtualMachineGuestOsIdentifierOther6xLinux:
		return "other6xLinuxGuest"
	case VirtualMachineGuestOsIdentifierOther7xLinux:
		return "other7xLinuxGuest"
	case VirtualMachineGuestOsIdentifierGenericLinux:
		return "genericLinuxGuest"
	case VirtualMachineGuestOsIdentifierOther24xLinux64:
		return "other24xLinux64Guest"
	case VirtualMachineGuestOsIdentifierOther26xLinux64:
		return "other26xLinux64Guest"
	case VirtualMachineGuestOsIdentifierOther3xLinux64:
		return "other3xLinux64Guest"
	case VirtualMachineGuestOsIdentifierOther4xLinux64:
		return "other4xLinux64Guest"
	case VirtualMachineGuestOsIdentifierOther5xLinux64:
		return "other5xLinux64Guest"
	case VirtualMachineGuestOsIdentifierOther6xLinux64:
		return "other6xLinux64Guest"
	case VirtualMachineGuestOsIdentifierOther7xLinux64:
		return "other7xLinux64Guest"
	case VirtualMachineGuestOsIdentifierOtherLinux64:
		return "otherLinux64Guest"
	}
	return string(t)
}

// toVimTypeLinuxMisc handles the remaining Linux distributions.
func toVimTypeLinuxMisc(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	// Novell
	case VirtualMachineGuestOsIdentifierNLD9:
		return "nld9Guest"
	case VirtualMachineGuestOsIdentifierOES:
		return "oesGuest"
	case VirtualMachineGuestOsIdentifierSJDS:
		return "sjdsGuest"
	// Mandriva / Mandrake
	case VirtualMachineGuestOsIdentifierMandrake:
		return "mandrakeGuest"
	case VirtualMachineGuestOsIdentifierMandriva:
		return "mandrivaGuest"
	case VirtualMachineGuestOsIdentifierMandriva64:
		return "mandriva64Guest"
	// TurboLinux
	case VirtualMachineGuestOsIdentifierTurboLinux:
		return "turboLinuxGuest"
	case VirtualMachineGuestOsIdentifierTurboLinux64:
		return "turboLinux64Guest"
	// Ubuntu
	case VirtualMachineGuestOsIdentifierUbuntu:
		return "ubuntuGuest"
	case VirtualMachineGuestOsIdentifierUbuntu64:
		return "ubuntu64Guest"
	// Fedora
	case VirtualMachineGuestOsIdentifierFedora:
		return "fedoraGuest"
	case VirtualMachineGuestOsIdentifierFedora64:
		return "fedora64Guest"
	// CoreOS / Photon
	case VirtualMachineGuestOsIdentifierCoreOS64:
		return "coreos64Guest"
	case VirtualMachineGuestOsIdentifierVMwarePhoton64:
		return "vmwarePhoton64Guest"
	// Chinese Linux distributions
	case VirtualMachineGuestOsIdentifierFusionOS64:
		return "fusionos_64Guest"
	case VirtualMachineGuestOsIdentifierProLinux64:
		return "prolinux_64Guest"
	case VirtualMachineGuestOsIdentifierKylinLinux64:
		return "kylinlinux_64Guest"
	// Amazon Linux
	case VirtualMachineGuestOsIdentifierAmazonLinux2x64:
		return "amazonlinux2_64Guest"
	case VirtualMachineGuestOsIdentifierAmazonLinux3x64:
		return "amazonlinux3_64Guest"
	// Rocky Linux / AlmaLinux
	case VirtualMachineGuestOsIdentifierRockyLinux64:
		return "rockylinux_64Guest"
	case VirtualMachineGuestOsIdentifierAlmaLinux64:
		return "almalinux_64Guest"
	}
	return string(t)
}

func toVimTypeSolaris(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierSolaris6:
		return "solaris6Guest"
	case VirtualMachineGuestOsIdentifierSolaris7:
		return "solaris7Guest"
	case VirtualMachineGuestOsIdentifierSolaris8:
		return "solaris8Guest"
	case VirtualMachineGuestOsIdentifierSolaris9:
		return "solaris9Guest"
	case VirtualMachineGuestOsIdentifierSolaris10:
		return "solaris10Guest"
	case VirtualMachineGuestOsIdentifierSolaris10x64:
		return "solaris10_64Guest"
	case VirtualMachineGuestOsIdentifierSolaris11x64:
		return "solaris11_64Guest"
	}
	return string(t)
}

func toVimTypeDarwin(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierDarwin:
		return "darwinGuest"
	case VirtualMachineGuestOsIdentifierDarwin64:
		return "darwin64Guest"
	case VirtualMachineGuestOsIdentifierDarwin10:
		return "darwin10Guest"
	case VirtualMachineGuestOsIdentifierDarwin10x64:
		return "darwin10_64Guest"
	case VirtualMachineGuestOsIdentifierDarwin11:
		return "darwin11Guest"
	case VirtualMachineGuestOsIdentifierDarwin11x64:
		return "darwin11_64Guest"
	case VirtualMachineGuestOsIdentifierDarwin12x64:
		return "darwin12_64Guest"
	case VirtualMachineGuestOsIdentifierDarwin13x64:
		return "darwin13_64Guest"
	case VirtualMachineGuestOsIdentifierDarwin14x64:
		return "darwin14_64Guest"
	case VirtualMachineGuestOsIdentifierDarwin15x64:
		return "darwin15_64Guest"
	case VirtualMachineGuestOsIdentifierDarwin16x64:
		return "darwin16_64Guest"
	case VirtualMachineGuestOsIdentifierDarwin17x64:
		return "darwin17_64Guest"
	case VirtualMachineGuestOsIdentifierDarwin18x64:
		return "darwin18_64Guest"
	case VirtualMachineGuestOsIdentifierDarwin19x64:
		return "darwin19_64Guest"
	case VirtualMachineGuestOsIdentifierDarwin20x64:
		return "darwin20_64Guest"
	case VirtualMachineGuestOsIdentifierDarwin21x64:
		return "darwin21_64Guest"
	case VirtualMachineGuestOsIdentifierDarwin22x64:
		return "darwin22_64Guest"
	case VirtualMachineGuestOsIdentifierDarwin23x64:
		return "darwin23_64Guest"
	}
	return string(t)
}

func toVimTypeVMkernel(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierVMkernel:
		return "vmkernelGuest"
	case VirtualMachineGuestOsIdentifierVMkernel5:
		return "vmkernel5Guest"
	case VirtualMachineGuestOsIdentifierVMkernel6:
		return "vmkernel6Guest"
	case VirtualMachineGuestOsIdentifierVMkernel65:
		return "vmkernel65Guest"
	case VirtualMachineGuestOsIdentifierVMkernel7:
		return "vmkernel7Guest"
	case VirtualMachineGuestOsIdentifierVMkernel8:
		return "vmkernel8Guest"
	case VirtualMachineGuestOsIdentifierVMkernel9:
		return "vmkernel9Guest"
	}
	return string(t)
}

func toVimTypeNetWare(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierNetWare4:
		return "netware4Guest"
	case VirtualMachineGuestOsIdentifierNetWare5:
		return "netware5Guest"
	case VirtualMachineGuestOsIdentifierNetWare6:
		return "netware6Guest"
	}
	return string(t)
}

func toVimTypeSCO(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierOpenServer5:
		return "openServer5Guest"
	case VirtualMachineGuestOsIdentifierOpenServer6:
		return "openServer6Guest"
	case VirtualMachineGuestOsIdentifierUnixWare7:
		return "unixWare7Guest"
	}
	return string(t)
}

func toVimTypeOS2(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierOS2:
		return "os2Guest"
	case VirtualMachineGuestOsIdentifierEComStation:
		return "eComStationGuest"
	case VirtualMachineGuestOsIdentifierEComStation2:
		return "eComStation2Guest"
	}
	return string(t)
}

func toVimTypeCRX(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierCRXPod1:
		return "crxPod1Guest"
	case VirtualMachineGuestOsIdentifierCRXSys1:
		return "crxSys1Guest"
	}
	return string(t)
}

func toVimTypeOtherOS(t VirtualMachineGuestOsIdentifier) string {
	switch t {
	case VirtualMachineGuestOsIdentifierOther:
		return "otherGuest"
	case VirtualMachineGuestOsIdentifierOther64:
		return "otherGuest64"
	}
	return string(t)
}

// FromVimType sets the VirtualMachineGuestOsIdentifier from the vSphere API
// identifier string.
func (t *VirtualMachineGuestOsIdentifier) FromVimType(s string) {
	switch {
	case s == "dosGuest",
		strings.HasPrefix(s, "win"):
		fromVimTypeWindows(t, s)
	case strings.HasPrefix(s, "freebsd"):
		fromVimTypeFreeBSD(t, s)
	case strings.HasPrefix(s, "solaris"):
		fromVimTypeSolaris(t, s)
	case strings.HasPrefix(s, "darwin"):
		fromVimTypeDarwin(t, s)
	case strings.HasPrefix(s, "vmkernel"):
		fromVimTypeVMkernel(t, s)
	case strings.HasPrefix(s, "netware"):
		fromVimTypeNetWare(t, s)
	case strings.HasPrefix(s, "openServer"),
		strings.HasPrefix(s, "unixWare"):
		fromVimTypeSCO(t, s)
	case s == "os2Guest",
		strings.HasPrefix(s, "eComStation"):
		fromVimTypeOS2(t, s)
	case strings.HasPrefix(s, "crx"):
		fromVimTypeCRX(t, s)
	case s == "otherGuest",
		s == "otherGuest64":
		fromVimTypeOtherOS(t, s)
	default:
		fromVimTypeLinux(t, s)
	}
}

// fromVimTypeWindows handles DOS and all Windows variants.
func fromVimTypeWindows(t *VirtualMachineGuestOsIdentifier, s string) {
	switch {
	case s == "dosGuest":
		*t = VirtualMachineGuestOsIdentifierDOS
	case strings.HasPrefix(s, "windows"):
		fromVimTypeWindowsModern(t, s)
	default:
		fromVimTypeWindowsLegacy(t, s)
	}
}

// fromVimTypeWindowsLegacy handles win3.1 through winVista.
func fromVimTypeWindowsLegacy(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "win31Guest":
		*t = VirtualMachineGuestOsIdentifierWin31
	case "win95Guest":
		*t = VirtualMachineGuestOsIdentifierWin95
	case "win98Guest":
		*t = VirtualMachineGuestOsIdentifierWin98
	case "winMeGuest":
		*t = VirtualMachineGuestOsIdentifierWinMe
	case "winNTGuest":
		*t = VirtualMachineGuestOsIdentifierWinNT
	case "win2000ProGuest":
		*t = VirtualMachineGuestOsIdentifierWin2000Pro
	case "win2000ServGuest":
		*t = VirtualMachineGuestOsIdentifierWin2000Serv
	case "win2000AdvServGuest":
		*t = VirtualMachineGuestOsIdentifierWin2000AdvServ
	case "winXPHomeGuest":
		*t = VirtualMachineGuestOsIdentifierWinXPHome
	case "winXPProGuest":
		*t = VirtualMachineGuestOsIdentifierWinXPPro
	case "winXPPro64Guest":
		*t = VirtualMachineGuestOsIdentifierWinXPPro64
	case "winNetWebGuest":
		*t = VirtualMachineGuestOsIdentifierWinNetWeb
	case "winNetStandardGuest":
		*t = VirtualMachineGuestOsIdentifierWinNetStandard
	case "winNetEnterpriseGuest":
		*t = VirtualMachineGuestOsIdentifierWinNetEnterprise
	case "winNetDatacenterGuest":
		*t = VirtualMachineGuestOsIdentifierWinNetDatacenter
	case "winNetBusinessGuest":
		*t = VirtualMachineGuestOsIdentifierWinNetBusiness
	case "winNetStandard64Guest":
		*t = VirtualMachineGuestOsIdentifierWinNetStandard64
	case "winNetEnterprise64Guest":
		*t = VirtualMachineGuestOsIdentifierWinNetEnterprise64
	case "winLonghornGuest":
		*t = VirtualMachineGuestOsIdentifierWinLonghorn
	case "winLonghorn64Guest":
		*t = VirtualMachineGuestOsIdentifierWinLonghorn64
	case "winNetDatacenter64Guest":
		*t = VirtualMachineGuestOsIdentifierWinNetDatacenter64
	case "winVistaGuest":
		*t = VirtualMachineGuestOsIdentifierWinVista
	case "winVista64Guest":
		*t = VirtualMachineGuestOsIdentifierWinVista64
	}
}

// fromVimTypeWindowsModern handles windows7 through windows2022.
func fromVimTypeWindowsModern(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "windows7Guest":
		*t = VirtualMachineGuestOsIdentifierWindows7
	case "windows7_64Guest":
		*t = VirtualMachineGuestOsIdentifierWindows7x64
	case "windows7Server64Guest":
		*t = VirtualMachineGuestOsIdentifierWindows7Server64
	case "windows8Guest":
		*t = VirtualMachineGuestOsIdentifierWindows8
	case "windows8_64Guest":
		*t = VirtualMachineGuestOsIdentifierWindows8x64
	case "windows8Server64Guest":
		*t = VirtualMachineGuestOsIdentifierWindows8Server64
	case "windows9Guest":
		*t = VirtualMachineGuestOsIdentifierWindows9
	case "windows9_64Guest":
		*t = VirtualMachineGuestOsIdentifierWindows9x64
	case "windows9Server64Guest":
		*t = VirtualMachineGuestOsIdentifierWindows9Server64
	case "windows11_64Guest":
		*t = VirtualMachineGuestOsIdentifierWindows11x64
	case "windows12_64Guest":
		*t = VirtualMachineGuestOsIdentifierWindows12x64
	case "windowsHyperVGuest":
		*t = VirtualMachineGuestOsIdentifierWindowsHyperV
	case "windows2019srv_64Guest":
		*t = VirtualMachineGuestOsIdentifierWindows2019Serverx64
	case "windows2019srvNext_64Guest":
		*t = VirtualMachineGuestOsIdentifierWindows2019ServerNextx64
	case "windows2022srvNext_64Guest":
		*t = VirtualMachineGuestOsIdentifierWindows2022ServerNextx64
	}
}

func fromVimTypeFreeBSD(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "freebsdGuest":
		*t = VirtualMachineGuestOsIdentifierFreeBSD
	case "freebsd64Guest":
		*t = VirtualMachineGuestOsIdentifierFreeBSD64
	case "freebsd11Guest":
		*t = VirtualMachineGuestOsIdentifierFreeBSD11
	case "freebsd11_64Guest":
		*t = VirtualMachineGuestOsIdentifierFreeBSD11x64
	case "freebsd12Guest":
		*t = VirtualMachineGuestOsIdentifierFreeBSD12
	case "freebsd12_64Guest":
		*t = VirtualMachineGuestOsIdentifierFreeBSD12x64
	case "freebsd13Guest":
		*t = VirtualMachineGuestOsIdentifierFreeBSD13
	case "freebsd13_64Guest":
		*t = VirtualMachineGuestOsIdentifierFreeBSD13x64
	case "freebsd14Guest":
		*t = VirtualMachineGuestOsIdentifierFreeBSD14
	case "freebsd14_64Guest":
		*t = VirtualMachineGuestOsIdentifierFreeBSD14x64
	case "freebsd15Guest":
		*t = VirtualMachineGuestOsIdentifierFreeBSD15
	case "freebsd15_64Guest":
		*t = VirtualMachineGuestOsIdentifierFreeBSD15x64
	}
}

// fromVimTypeLinux dispatches to a distro-specific helper.
func fromVimTypeLinux(t *VirtualMachineGuestOsIdentifier, s string) {
	switch {
	case s == "redhatGuest",
		strings.HasPrefix(s, "rhel"):
		fromVimTypeLinuxRHEL(t, s)
	case strings.HasPrefix(s, "centos"):
		fromVimTypeLinuxCentOS(t, s)
	case strings.HasPrefix(s, "oracle"):
		fromVimTypeLinuxOracle(t, s)
	case strings.HasPrefix(s, "suse"),
		strings.HasPrefix(s, "sles"),
		strings.HasPrefix(s, "opensuse"):
		fromVimTypeLinuxSUSE(t, s)
	case strings.HasPrefix(s, "debian"):
		fromVimTypeLinuxDebian(t, s)
	case strings.HasPrefix(s, "asianux"),
		strings.HasPrefix(s, "miraclelinux"),
		strings.HasPrefix(s, "pardus"):
		fromVimTypeLinuxAsianux(t, s)
	case strings.HasPrefix(s, "other"),
		s == "genericLinuxGuest":
		fromVimTypeLinuxOtherKernel(t, s)
	default:
		fromVimTypeLinuxMisc(t, s)
	}
}

func fromVimTypeLinuxRHEL(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "redhatGuest":
		*t = VirtualMachineGuestOsIdentifierRedHat
	case "rhel2Guest":
		*t = VirtualMachineGuestOsIdentifierRHEL2
	case "rhel3Guest":
		*t = VirtualMachineGuestOsIdentifierRHEL3
	case "rhel3_64Guest":
		*t = VirtualMachineGuestOsIdentifierRHEL3x64
	case "rhel4Guest":
		*t = VirtualMachineGuestOsIdentifierRHEL4
	case "rhel4_64Guest":
		*t = VirtualMachineGuestOsIdentifierRHEL4x64
	case "rhel5Guest":
		*t = VirtualMachineGuestOsIdentifierRHEL5
	case "rhel5_64Guest":
		*t = VirtualMachineGuestOsIdentifierRHEL5x64
	case "rhel6Guest":
		*t = VirtualMachineGuestOsIdentifierRHEL6
	case "rhel6_64Guest":
		*t = VirtualMachineGuestOsIdentifierRHEL6x64
	case "rhel7Guest":
		*t = VirtualMachineGuestOsIdentifierRHEL7
	case "rhel7_64Guest":
		*t = VirtualMachineGuestOsIdentifierRHEL7x64
	case "rhel8_64Guest":
		*t = VirtualMachineGuestOsIdentifierRHEL8x64
	case "rhel9_64Guest":
		*t = VirtualMachineGuestOsIdentifierRHEL9x64
	case "rhel10_64Guest":
		*t = VirtualMachineGuestOsIdentifierRHEL10x64
	}
}

func fromVimTypeLinuxCentOS(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "centosGuest":
		*t = VirtualMachineGuestOsIdentifierCentOS
	case "centos64Guest":
		*t = VirtualMachineGuestOsIdentifierCentOS64
	case "centos6Guest":
		*t = VirtualMachineGuestOsIdentifierCentOS6
	case "centos6_64Guest":
		*t = VirtualMachineGuestOsIdentifierCentOS6x64
	case "centos7Guest":
		*t = VirtualMachineGuestOsIdentifierCentOS7
	case "centos7_64Guest":
		*t = VirtualMachineGuestOsIdentifierCentOS7x64
	case "centos8_64Guest":
		*t = VirtualMachineGuestOsIdentifierCentOS8x64
	case "centos9_64Guest":
		*t = VirtualMachineGuestOsIdentifierCentOS9x64
	}
}

func fromVimTypeLinuxOracle(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "oracleLinuxGuest":
		*t = VirtualMachineGuestOsIdentifierOracleLinux
	case "oracleLinux64Guest":
		*t = VirtualMachineGuestOsIdentifierOracleLinux64
	case "oracleLinux6Guest":
		*t = VirtualMachineGuestOsIdentifierOracleLinux6
	case "oracleLinux6_64Guest":
		*t = VirtualMachineGuestOsIdentifierOracleLinux6x64
	case "oracleLinux7Guest":
		*t = VirtualMachineGuestOsIdentifierOracleLinux7
	case "oracleLinux7_64Guest":
		*t = VirtualMachineGuestOsIdentifierOracleLinux7x64
	case "oracleLinux8_64Guest":
		*t = VirtualMachineGuestOsIdentifierOracleLinux8x64
	case "oracleLinux9_64Guest":
		*t = VirtualMachineGuestOsIdentifierOracleLinux9x64
	case "oracleLinux10_64Guest":
		*t = VirtualMachineGuestOsIdentifierOracleLinux10x64
	}
}

// fromVimTypeLinuxSUSE handles suseGuest, sles*, and opensuse*.
func fromVimTypeLinuxSUSE(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "suseGuest":
		*t = VirtualMachineGuestOsIdentifierSUSE
	case "suse64Guest":
		*t = VirtualMachineGuestOsIdentifierSUSE64
	case "slesGuest":
		*t = VirtualMachineGuestOsIdentifierSLES
	case "sles64Guest":
		*t = VirtualMachineGuestOsIdentifierSLES64
	case "sles10Guest":
		*t = VirtualMachineGuestOsIdentifierSLES10
	case "sles10_64Guest":
		*t = VirtualMachineGuestOsIdentifierSLES10x64
	case "sles11Guest":
		*t = VirtualMachineGuestOsIdentifierSLES11
	case "sles11_64Guest":
		*t = VirtualMachineGuestOsIdentifierSLES11x64
	case "sles12Guest":
		*t = VirtualMachineGuestOsIdentifierSLES12
	case "sles12_64Guest":
		*t = VirtualMachineGuestOsIdentifierSLES12x64
	case "sles15_64Guest":
		*t = VirtualMachineGuestOsIdentifierSLES15x64
	case "sles16_64Guest":
		*t = VirtualMachineGuestOsIdentifierSLES16x64
	case "opensuseGuest":
		*t = VirtualMachineGuestOsIdentifierOpenSUSE
	case "opensuse64Guest":
		*t = VirtualMachineGuestOsIdentifierOpenSUSE64
	}
}

func fromVimTypeLinuxDebian(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "debian4Guest":
		*t = VirtualMachineGuestOsIdentifierDebian4
	case "debian4_64Guest":
		*t = VirtualMachineGuestOsIdentifierDebian4x64
	case "debian5Guest":
		*t = VirtualMachineGuestOsIdentifierDebian5
	case "debian5_64Guest":
		*t = VirtualMachineGuestOsIdentifierDebian5x64
	case "debian6Guest":
		*t = VirtualMachineGuestOsIdentifierDebian6
	case "debian6_64Guest":
		*t = VirtualMachineGuestOsIdentifierDebian6x64
	case "debian7Guest":
		*t = VirtualMachineGuestOsIdentifierDebian7
	case "debian7_64Guest":
		*t = VirtualMachineGuestOsIdentifierDebian7x64
	case "debian8Guest":
		*t = VirtualMachineGuestOsIdentifierDebian8
	case "debian8_64Guest":
		*t = VirtualMachineGuestOsIdentifierDebian8x64
	case "debian9Guest":
		*t = VirtualMachineGuestOsIdentifierDebian9
	case "debian9_64Guest":
		*t = VirtualMachineGuestOsIdentifierDebian9x64
	case "debian10Guest":
		*t = VirtualMachineGuestOsIdentifierDebian10
	case "debian10_64Guest":
		*t = VirtualMachineGuestOsIdentifierDebian10x64
	case "debian11Guest":
		*t = VirtualMachineGuestOsIdentifierDebian11
	case "debian11_64Guest":
		*t = VirtualMachineGuestOsIdentifierDebian11x64
	case "debian12Guest":
		*t = VirtualMachineGuestOsIdentifierDebian12
	case "debian12_64Guest":
		*t = VirtualMachineGuestOsIdentifierDebian12x64
	case "debian13Guest":
		*t = VirtualMachineGuestOsIdentifierDebian13
	case "debian13_64Guest":
		*t = VirtualMachineGuestOsIdentifierDebian13x64
	}
}

// fromVimTypeLinuxAsianux handles asianux*, miraclelinux*, and pardus*.
func fromVimTypeLinuxAsianux(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "asianux3Guest":
		*t = VirtualMachineGuestOsIdentifierAsianux3
	case "asianux3_64Guest":
		*t = VirtualMachineGuestOsIdentifierAsianux3x64
	case "asianux4Guest":
		*t = VirtualMachineGuestOsIdentifierAsianux4
	case "asianux4_64Guest":
		*t = VirtualMachineGuestOsIdentifierAsianux4x64
	case "asianux5_64Guest":
		*t = VirtualMachineGuestOsIdentifierAsianux5x64
	case "asianux7_64Guest":
		*t = VirtualMachineGuestOsIdentifierAsianux7x64
	case "asianux8_64Guest":
		*t = VirtualMachineGuestOsIdentifierAsianux8x64
	case "asianux9_64Guest":
		*t = VirtualMachineGuestOsIdentifierAsianux9x64
	case "miraclelinux_64Guest":
		*t = VirtualMachineGuestOsIdentifierMiracleLinux64
	case "pardus_64Guest":
		*t = VirtualMachineGuestOsIdentifierPardus64
	}
}

// fromVimTypeLinuxOtherKernel handles generic/other Linux kernel variants.
func fromVimTypeLinuxOtherKernel(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "other24xLinuxGuest":
		*t = VirtualMachineGuestOsIdentifierOther24xLinux
	case "other26xLinuxGuest":
		*t = VirtualMachineGuestOsIdentifierOther26xLinux
	case "otherLinuxGuest":
		*t = VirtualMachineGuestOsIdentifierOtherLinux
	case "other3xLinuxGuest":
		*t = VirtualMachineGuestOsIdentifierOther3xLinux
	case "other4xLinuxGuest":
		*t = VirtualMachineGuestOsIdentifierOther4xLinux
	case "other5xLinuxGuest":
		*t = VirtualMachineGuestOsIdentifierOther5xLinux
	case "other6xLinuxGuest":
		*t = VirtualMachineGuestOsIdentifierOther6xLinux
	case "other7xLinuxGuest":
		*t = VirtualMachineGuestOsIdentifierOther7xLinux
	case "genericLinuxGuest":
		*t = VirtualMachineGuestOsIdentifierGenericLinux
	case "other24xLinux64Guest":
		*t = VirtualMachineGuestOsIdentifierOther24xLinux64
	case "other26xLinux64Guest":
		*t = VirtualMachineGuestOsIdentifierOther26xLinux64
	case "other3xLinux64Guest":
		*t = VirtualMachineGuestOsIdentifierOther3xLinux64
	case "other4xLinux64Guest":
		*t = VirtualMachineGuestOsIdentifierOther4xLinux64
	case "other5xLinux64Guest":
		*t = VirtualMachineGuestOsIdentifierOther5xLinux64
	case "other6xLinux64Guest":
		*t = VirtualMachineGuestOsIdentifierOther6xLinux64
	case "other7xLinux64Guest":
		*t = VirtualMachineGuestOsIdentifierOther7xLinux64
	case "otherLinux64Guest":
		*t = VirtualMachineGuestOsIdentifierOtherLinux64
	}
}

// fromVimTypeLinuxMisc handles the remaining Linux distributions.
func fromVimTypeLinuxMisc(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	// Novell
	case "nld9Guest":
		*t = VirtualMachineGuestOsIdentifierNLD9
	case "oesGuest":
		*t = VirtualMachineGuestOsIdentifierOES
	case "sjdsGuest":
		*t = VirtualMachineGuestOsIdentifierSJDS
	// Mandriva / Mandrake
	case "mandrakeGuest":
		*t = VirtualMachineGuestOsIdentifierMandrake
	case "mandrivaGuest":
		*t = VirtualMachineGuestOsIdentifierMandriva
	case "mandriva64Guest":
		*t = VirtualMachineGuestOsIdentifierMandriva64
	// TurboLinux
	case "turboLinuxGuest":
		*t = VirtualMachineGuestOsIdentifierTurboLinux
	case "turboLinux64Guest":
		*t = VirtualMachineGuestOsIdentifierTurboLinux64
	// Ubuntu
	case "ubuntuGuest":
		*t = VirtualMachineGuestOsIdentifierUbuntu
	case "ubuntu64Guest":
		*t = VirtualMachineGuestOsIdentifierUbuntu64
	// Fedora
	case "fedoraGuest":
		*t = VirtualMachineGuestOsIdentifierFedora
	case "fedora64Guest":
		*t = VirtualMachineGuestOsIdentifierFedora64
	// CoreOS / Photon
	case "coreos64Guest":
		*t = VirtualMachineGuestOsIdentifierCoreOS64
	case "vmwarePhoton64Guest":
		*t = VirtualMachineGuestOsIdentifierVMwarePhoton64
	// Chinese Linux distributions
	case "fusionos_64Guest":
		*t = VirtualMachineGuestOsIdentifierFusionOS64
	case "prolinux_64Guest":
		*t = VirtualMachineGuestOsIdentifierProLinux64
	case "kylinlinux_64Guest":
		*t = VirtualMachineGuestOsIdentifierKylinLinux64
	// Amazon Linux
	case "amazonlinux2_64Guest":
		*t = VirtualMachineGuestOsIdentifierAmazonLinux2x64
	case "amazonlinux3_64Guest":
		*t = VirtualMachineGuestOsIdentifierAmazonLinux3x64
	// Rocky Linux / AlmaLinux
	case "rockylinux_64Guest":
		*t = VirtualMachineGuestOsIdentifierRockyLinux64
	case "almalinux_64Guest":
		*t = VirtualMachineGuestOsIdentifierAlmaLinux64
	}
}

func fromVimTypeSolaris(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "solaris6Guest":
		*t = VirtualMachineGuestOsIdentifierSolaris6
	case "solaris7Guest":
		*t = VirtualMachineGuestOsIdentifierSolaris7
	case "solaris8Guest":
		*t = VirtualMachineGuestOsIdentifierSolaris8
	case "solaris9Guest":
		*t = VirtualMachineGuestOsIdentifierSolaris9
	case "solaris10Guest":
		*t = VirtualMachineGuestOsIdentifierSolaris10
	case "solaris10_64Guest":
		*t = VirtualMachineGuestOsIdentifierSolaris10x64
	case "solaris11_64Guest":
		*t = VirtualMachineGuestOsIdentifierSolaris11x64
	}
}

func fromVimTypeDarwin(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "darwinGuest":
		*t = VirtualMachineGuestOsIdentifierDarwin
	case "darwin64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin64
	case "darwin10Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin10
	case "darwin10_64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin10x64
	case "darwin11Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin11
	case "darwin11_64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin11x64
	case "darwin12_64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin12x64
	case "darwin13_64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin13x64
	case "darwin14_64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin14x64
	case "darwin15_64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin15x64
	case "darwin16_64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin16x64
	case "darwin17_64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin17x64
	case "darwin18_64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin18x64
	case "darwin19_64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin19x64
	case "darwin20_64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin20x64
	case "darwin21_64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin21x64
	case "darwin22_64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin22x64
	case "darwin23_64Guest":
		*t = VirtualMachineGuestOsIdentifierDarwin23x64
	}
}

func fromVimTypeVMkernel(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "vmkernelGuest":
		*t = VirtualMachineGuestOsIdentifierVMkernel
	case "vmkernel5Guest":
		*t = VirtualMachineGuestOsIdentifierVMkernel5
	case "vmkernel6Guest":
		*t = VirtualMachineGuestOsIdentifierVMkernel6
	case "vmkernel65Guest":
		*t = VirtualMachineGuestOsIdentifierVMkernel65
	case "vmkernel7Guest":
		*t = VirtualMachineGuestOsIdentifierVMkernel7
	case "vmkernel8Guest":
		*t = VirtualMachineGuestOsIdentifierVMkernel8
	case "vmkernel9Guest":
		*t = VirtualMachineGuestOsIdentifierVMkernel9
	}
}

func fromVimTypeNetWare(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "netware4Guest":
		*t = VirtualMachineGuestOsIdentifierNetWare4
	case "netware5Guest":
		*t = VirtualMachineGuestOsIdentifierNetWare5
	case "netware6Guest":
		*t = VirtualMachineGuestOsIdentifierNetWare6
	}
}

func fromVimTypeSCO(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "openServer5Guest":
		*t = VirtualMachineGuestOsIdentifierOpenServer5
	case "openServer6Guest":
		*t = VirtualMachineGuestOsIdentifierOpenServer6
	case "unixWare7Guest":
		*t = VirtualMachineGuestOsIdentifierUnixWare7
	}
}

func fromVimTypeOS2(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "os2Guest":
		*t = VirtualMachineGuestOsIdentifierOS2
	case "eComStationGuest":
		*t = VirtualMachineGuestOsIdentifierEComStation
	case "eComStation2Guest":
		*t = VirtualMachineGuestOsIdentifierEComStation2
	}
}

func fromVimTypeCRX(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "crxPod1Guest":
		*t = VirtualMachineGuestOsIdentifierCRXPod1
	case "crxSys1Guest":
		*t = VirtualMachineGuestOsIdentifierCRXSys1
	}
}

func fromVimTypeOtherOS(t *VirtualMachineGuestOsIdentifier, s string) {
	switch s {
	case "otherGuest":
		*t = VirtualMachineGuestOsIdentifierOther
	case "otherGuest64":
		*t = VirtualMachineGuestOsIdentifierOther64
	}
}
