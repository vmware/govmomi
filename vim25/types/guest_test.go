// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package types

import "testing"

func TestGuestIDToFamily(t *testing.T) {
	tests := []struct {
		name     string
		guestID  VirtualMachineGuestOsIdentifier
		expected VirtualMachineGuestOsFamily
	}{
		// Windows family tests
		{
			name:     "DOS Guest",
			guestID:  VirtualMachineGuestOsIdentifierDosGuest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "Windows 31 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWin31Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 95 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWin95Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 98 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWin98Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows ME Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinMeGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows NT Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinNTGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 2000 Pro Guest",
			guestID:  VirtualMachineGuestOsIdentifierWin2000ProGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 2000 Server Guest",
			guestID:  VirtualMachineGuestOsIdentifierWin2000ServGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 2000 Advanced Server Guest",
			guestID:  VirtualMachineGuestOsIdentifierWin2000AdvServGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows XP Home Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinXPHomeGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows XP Pro Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinXPProGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows XP Pro 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinXPPro64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Server 2003 Web Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinNetWebGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Server 2003 Standard Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinNetStandardGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Server 2003 Enterprise Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinNetEnterpriseGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Server 2003 Datacenter Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinNetDatacenterGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Server 2003 Business Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinNetBusinessGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Server 2003 Standard 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinNetStandard64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Server 2003 Enterprise 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinNetEnterprise64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Longhorn Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinLonghornGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Longhorn 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinLonghorn64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Server 2003 Datacenter 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinNetDatacenter64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Vista Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinVistaGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Vista 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWinVista64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 7 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindows7Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 7 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindows7_64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Server 2008 R2 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindows7Server64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 8 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindows8Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 8 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindows8_64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 8 Server 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindows8Server64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 9 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindows9Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 9 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindows9_64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 9 Server 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindows9Server64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 11 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindows11_64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows 12 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindows12_64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Hyper-V Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindowsHyperVGuest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Server 2019 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindows2019srv_64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Server 2022 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindows2019srvNext_64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Windows Server 2025 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierWindows2022srvNext_64Guest,
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},

		// FreeBSD family tests (should be Other)
		{
			name:     "FreeBSD Guest",
			guestID:  VirtualMachineGuestOsIdentifierFreebsdGuest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "FreeBSD 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierFreebsd64Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "FreeBSD 11 Guest",
			guestID:  VirtualMachineGuestOsIdentifierFreebsd11Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "FreeBSD 11 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierFreebsd11_64Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "FreeBSD 12 Guest",
			guestID:  VirtualMachineGuestOsIdentifierFreebsd12Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "FreeBSD 12 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierFreebsd12_64Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "FreeBSD 13 Guest",
			guestID:  VirtualMachineGuestOsIdentifierFreebsd13Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "FreeBSD 13 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierFreebsd13_64Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "FreeBSD 14 Guest",
			guestID:  VirtualMachineGuestOsIdentifierFreebsd14Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "FreeBSD 14 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierFreebsd14_64Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "FreeBSD 15 Guest",
			guestID:  VirtualMachineGuestOsIdentifierFreebsd15Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "FreeBSD 15 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierFreebsd15_64Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},

		// Red Hat Linux family tests
		{
			name:     "Red Hat Guest",
			guestID:  VirtualMachineGuestOsIdentifierRedhatGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "RHEL 2 Guest",
			guestID:  VirtualMachineGuestOsIdentifierRhel2Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "RHEL 3 Guest",
			guestID:  VirtualMachineGuestOsIdentifierRhel3Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "RHEL 3 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierRhel3_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "RHEL 4 Guest",
			guestID:  VirtualMachineGuestOsIdentifierRhel4Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "RHEL 4 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierRhel4_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "RHEL 5 Guest",
			guestID:  VirtualMachineGuestOsIdentifierRhel5Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "RHEL 5 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierRhel5_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "RHEL 6 Guest",
			guestID:  VirtualMachineGuestOsIdentifierRhel6Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "RHEL 6 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierRhel6_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "RHEL 7 Guest",
			guestID:  VirtualMachineGuestOsIdentifierRhel7Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "RHEL 7 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierRhel7_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "RHEL 8 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierRhel8_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "RHEL 9 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierRhel9_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "RHEL 10 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierRhel10_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},

		// CentOS Linux family tests
		{
			name:     "CentOS Guest",
			guestID:  VirtualMachineGuestOsIdentifierCentosGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "CentOS 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierCentos64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "CentOS 6 Guest",
			guestID:  VirtualMachineGuestOsIdentifierCentos6Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "CentOS 6 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierCentos6_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "CentOS 7 Guest",
			guestID:  VirtualMachineGuestOsIdentifierCentos7Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "CentOS 7 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierCentos7_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "CentOS 8 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierCentos8_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "CentOS 9 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierCentos9_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},

		// Oracle Linux family tests
		{
			name:     "Oracle Linux Guest",
			guestID:  VirtualMachineGuestOsIdentifierOracleLinuxGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Oracle Linux 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOracleLinux64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Oracle Linux 6 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOracleLinux6Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Oracle Linux 6 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOracleLinux6_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Oracle Linux 7 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOracleLinux7Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Oracle Linux 7 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOracleLinux7_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Oracle Linux 8 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOracleLinux8_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Oracle Linux 9 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOracleLinux9_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Oracle Linux 10 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOracleLinux10_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},

		// SUSE Linux family tests
		{
			name:     "SUSE Guest",
			guestID:  VirtualMachineGuestOsIdentifierSuseGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "SUSE 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSuse64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "SLES Guest",
			guestID:  VirtualMachineGuestOsIdentifierSlesGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "SLES 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSles64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "SLES 10 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSles10Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "SLES 10 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSles10_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "SLES 11 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSles11Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "SLES 11 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSles11_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "SLES 12 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSles12Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "SLES 12 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSles12_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "SLES 15 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSles15_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "SLES 16 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSles16_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},

		// OpenSUSE Linux family tests
		{
			name:     "OpenSUSE Guest",
			guestID:  VirtualMachineGuestOsIdentifierOpensuseGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "OpenSUSE 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOpensuse64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},

		// Mandrake/Mandriva Linux family tests
		{
			name:     "Mandrake Guest",
			guestID:  VirtualMachineGuestOsIdentifierMandrakeGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Mandriva Guest",
			guestID:  VirtualMachineGuestOsIdentifierMandrivaGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Mandriva 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierMandriva64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},

		// Ubuntu Linux family tests
		{
			name:     "Ubuntu Guest",
			guestID:  VirtualMachineGuestOsIdentifierUbuntuGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Ubuntu 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierUbuntu64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},

		// Debian Linux family tests
		{
			name:     "Debian 4 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian4Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 4 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian4_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 5 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian5Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 5 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian5_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 6 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian6Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 6 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian6_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 7 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian7Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 7 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian7_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 8 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian8Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 8 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian8_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 9 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian9Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 9 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian9_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 10 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian10Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 10 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian10_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 11 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian11Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 11 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian11_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 12 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian12Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 12 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian12_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 13 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian13Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Debian 13 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDebian13_64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},

		// Fedora Linux family tests
		{
			name:     "Fedora Guest",
			guestID:  VirtualMachineGuestOsIdentifierFedoraGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Fedora 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierFedora64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},

		// VMware Photon Linux family tests
		{
			name:     "VMware Photon 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierVmwarePhoton64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},

		// Other Linux family tests
		{
			name:     "CoreOS 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierCoreos64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other 24x Linux Guest",
			guestID:  VirtualMachineGuestOsIdentifierOther24xLinuxGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other 26x Linux Guest",
			guestID:  VirtualMachineGuestOsIdentifierOther26xLinuxGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other Linux Guest",
			guestID:  VirtualMachineGuestOsIdentifierOtherLinuxGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other 3x Linux Guest",
			guestID:  VirtualMachineGuestOsIdentifierOther3xLinuxGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other 4x Linux Guest",
			guestID:  VirtualMachineGuestOsIdentifierOther4xLinuxGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other 5x Linux Guest",
			guestID:  VirtualMachineGuestOsIdentifierOther5xLinuxGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other 6x Linux Guest",
			guestID:  VirtualMachineGuestOsIdentifierOther6xLinuxGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other 7x Linux Guest",
			guestID:  VirtualMachineGuestOsIdentifierOther7xLinuxGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Generic Linux Guest",
			guestID:  VirtualMachineGuestOsIdentifierGenericLinuxGuest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other 24x Linux 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOther24xLinux64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other 26x Linux 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOther26xLinux64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other 3x Linux 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOther3xLinux64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other 4x Linux 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOther4xLinux64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other 5x Linux 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOther5xLinux64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other 6x Linux 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOther6xLinux64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other 7x Linux 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOther7xLinux64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Other Linux 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOtherLinux64Guest,
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},

		// Solaris family tests
		{
			name:     "Solaris 6 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSolaris6Guest,
			expected: VirtualMachineGuestOsFamilySolarisGuest,
		},
		{
			name:     "Solaris 7 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSolaris7Guest,
			expected: VirtualMachineGuestOsFamilySolarisGuest,
		},
		{
			name:     "Solaris 8 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSolaris8Guest,
			expected: VirtualMachineGuestOsFamilySolarisGuest,
		},
		{
			name:     "Solaris 9 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSolaris9Guest,
			expected: VirtualMachineGuestOsFamilySolarisGuest,
		},
		{
			name:     "Solaris 10 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSolaris10Guest,
			expected: VirtualMachineGuestOsFamilySolarisGuest,
		},
		{
			name:     "Solaris 10 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSolaris10_64Guest,
			expected: VirtualMachineGuestOsFamilySolarisGuest,
		},
		{
			name:     "Solaris 11 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierSolaris11_64Guest,
			expected: VirtualMachineGuestOsFamilySolarisGuest,
		},

		// Netware family tests
		{
			name:     "Netware 4 Guest",
			guestID:  VirtualMachineGuestOsIdentifierNetware4Guest,
			expected: VirtualMachineGuestOsFamilyNetwareGuest,
		},
		{
			name:     "Netware 5 Guest",
			guestID:  VirtualMachineGuestOsIdentifierNetware5Guest,
			expected: VirtualMachineGuestOsFamilyNetwareGuest,
		},
		{
			name:     "Netware 6 Guest",
			guestID:  VirtualMachineGuestOsIdentifierNetware6Guest,
			expected: VirtualMachineGuestOsFamilyNetwareGuest,
		},

		// Darwin/macOS family tests
		{
			name:     "Darwin Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwinGuest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 10 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin10Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 10 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin10_64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 11 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin11Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 11 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin11_64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 12 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin12_64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 13 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin13_64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 14 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin14_64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 15 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin15_64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 16 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin16_64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 17 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin17_64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 18 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin18_64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 19 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin19_64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 20 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin20_64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 21 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin21_64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 22 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin22_64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Darwin 23 64 Guest",
			guestID:  VirtualMachineGuestOsIdentifierDarwin23_64Guest,
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},

		// Other family tests (various non-matching OS types)
		{
			name:     "OS/2 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOs2Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "eComStation Guest",
			guestID:  VirtualMachineGuestOsIdentifierEComStationGuest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "eComStation 2 Guest",
			guestID:  VirtualMachineGuestOsIdentifierEComStation2Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "OpenServer 5 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOpenServer5Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "OpenServer 6 Guest",
			guestID:  VirtualMachineGuestOsIdentifierOpenServer6Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "UnixWare 7 Guest",
			guestID:  VirtualMachineGuestOsIdentifierUnixWare7Guest,
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},

		// Custom string tests (not enum constants)
		{
			name:     "Custom Windows String",
			guestID:  VirtualMachineGuestOsIdentifier("winCustom"),
			expected: VirtualMachineGuestOsFamilyWindowsGuest,
		},
		{
			name:     "Custom Linux String",
			guestID:  VirtualMachineGuestOsIdentifier("customLinux"),
			expected: VirtualMachineGuestOsFamilyLinuxGuest,
		},
		{
			name:     "Custom Darwin String",
			guestID:  VirtualMachineGuestOsIdentifier("darwinCustom"),
			expected: VirtualMachineGuestOsFamilyDarwinGuestFamily,
		},
		{
			name:     "Custom Solaris String",
			guestID:  VirtualMachineGuestOsIdentifier("solarisCustom"),
			expected: VirtualMachineGuestOsFamilySolarisGuest,
		},
		{
			name:     "Custom Netware String",
			guestID:  VirtualMachineGuestOsIdentifier("netwareCustom"),
			expected: VirtualMachineGuestOsFamilyNetwareGuest,
		},
		{
			name:     "Unknown OS String",
			guestID:  VirtualMachineGuestOsIdentifier("unknownOS"),
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
		{
			name:     "Empty String",
			guestID:  VirtualMachineGuestOsIdentifier(""),
			expected: VirtualMachineGuestOsFamilyOtherGuestFamily,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.guestID.ToFamily()
			if result != tt.expected {
				t.Errorf("GuestIDToFamily(%q) = %q, want %q", tt.guestID, result, tt.expected)
			}
		})
	}
}
