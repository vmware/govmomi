// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import (
	"testing"

	"github.com/vmware/govmomi/vim25/types"
)

func TestGuestIDToCIMOSType(t *testing.T) {
	tests := []struct {
		name     string
		guestID  types.VirtualMachineGuestOsIdentifier
		expected CIMOSType
	}{
		// Windows Desktop
		{
			name:     "DOS Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDosGuest,
			expected: CIMOSTypeMSDOS,
		},
		{
			name:     "Windows 3.1 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWin31Guest,
			expected: CIMOSTypeWIN3x,
		},
		{
			name:     "Windows 95 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWin95Guest,
			expected: CIMOSTypeWIN95,
		},
		{
			name:     "Windows 98 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWin98Guest,
			expected: CIMOSTypeWIN98,
		},
		{
			name:     "Windows ME Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinMeGuest,
			expected: CIMOSTypeWindowsMe,
		},
		{
			name:     "Windows NT Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinNTGuest,
			expected: CIMOSTypeWINNT,
		},
		{
			name:     "Windows 2000 Pro Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWin2000ProGuest,
			expected: CIMOSTypeWindows2000,
		},
		{
			name:     "Windows 2000 Server Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWin2000ServGuest,
			expected: CIMOSTypeWindows2000,
		},
		{
			name:     "Windows 2000 Advanced Server Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWin2000AdvServGuest,
			expected: CIMOSTypeWindows2000,
		},
		{
			name:     "Windows XP Home Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinXPHomeGuest,
			expected: CIMOSTypeWindowsXP,
		},
		{
			name:     "Windows XP Pro Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinXPProGuest,
			expected: CIMOSTypeWindowsXP,
		},
		{
			name:     "Windows XP Pro 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinXPPro64Guest,
			expected: CIMOSTypeWindowsXP64Bit,
		},
		{
			name:     "Windows Vista Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinVistaGuest,
			expected: CIMOSTypeWindowsVista,
		},
		{
			name:     "Windows Vista 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinVista64Guest,
			expected: CIMOSTypeWindowsVista64Bit,
		},
		{
			name:     "Windows 7 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindows7Guest,
			expected: CIMOSTypeMicrosoftWindows7,
		},
		{
			name:     "Windows 7 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindows7_64Guest,
			expected: CIMOSTypeMicrosoftWindows7,
		},
		{
			name:     "Windows 8 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindows8Guest,
			expected: CIMOSTypeMicrosoftWindows8,
		},
		{
			name:     "Windows 8 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindows8_64Guest,
			expected: CIMOSTypeMicrosoftWindows8,
		},
		{
			name:     "Windows 9 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindows9Guest,
			expected: CIMOSTypeOther,
		},
		{
			name:     "Windows 9 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindows9_64Guest,
			expected: CIMOSTypeOther64Bit,
		},
		{
			name:     "Windows 11 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindows11_64Guest,
			expected: CIMOSTypeOther64Bit,
		},
		{
			name:     "Windows 12 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindows12_64Guest,
			expected: CIMOSTypeOther64Bit,
		},
		{
			name:     "Windows Hyper-V Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindowsHyperVGuest,
			expected: CIMOSTypeOther64Bit,
		},

		// Windows Server
		{
			name:     "Windows Server 2003 Enterprise Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinNetEnterpriseGuest,
			expected: CIMOSTypeMicrosoftWindowsServer2003,
		},
		{
			name:     "Windows Server 2003 Datacenter Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinNetDatacenterGuest,
			expected: CIMOSTypeMicrosoftWindowsServer2003,
		},
		{
			name:     "Windows Server 2003 Business Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinNetBusinessGuest,
			expected: CIMOSTypeMicrosoftWindowsServer2003,
		},
		{
			name:     "Windows Server 2003 Standard Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinNetStandardGuest,
			expected: CIMOSTypeMicrosoftWindowsServer2003,
		},
		{
			name:     "Windows Server 2003 Web Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinNetWebGuest,
			expected: CIMOSTypeMicrosoftWindowsServer2003,
		},
		{
			name:     "Windows Server 2003 Enterprise 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinNetEnterprise64Guest,
			expected: CIMOSTypeMicrosoftWindowsServer2003_64Bit,
		},
		{
			name:     "Windows Server 2003 Datacenter 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinNetDatacenter64Guest,
			expected: CIMOSTypeMicrosoftWindowsServer2003_64Bit,
		},
		{
			name:     "Windows Server 2003 Standard 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinNetStandard64Guest,
			expected: CIMOSTypeMicrosoftWindowsServer2003_64Bit,
		},
		{
			name:     "Windows Longhorn Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinLonghornGuest,
			expected: CIMOSTypeMicrosoftWindowsServer2008,
		},
		{
			name:     "Windows Longhorn 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWinLonghorn64Guest,
			expected: CIMOSTypeMicrosoftWindowsServer2008_64Bit,
		},
		{
			name:     "Windows 7 Server 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindows7Server64Guest,
			expected: CIMOSTypeMicrosoftWindowsServer2008R2,
		},
		{
			name:     "Windows 8 Server 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindows8Server64Guest,
			expected: CIMOSTypeMicrosoftWindowsServer2012,
		},
		{
			name:     "Windows 9 Server 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindows9Server64Guest,
			expected: CIMOSTypeMicrosoftWindowsServer2012R2,
		},
		{
			name:     "Windows Server 2019 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindows2019srv_64Guest,
			expected: CIMOSTypeOther64Bit,
		},
		{
			name:     "Windows Server 2019 Next 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindows2019srvNext_64Guest,
			expected: CIMOSTypeOther64Bit,
		},
		{
			name:     "Windows Server 2022 Next 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierWindows2022srvNext_64Guest,
			expected: CIMOSTypeOther64Bit,
		},

		// Linux - Red Hat
		{
			name:     "Red Hat Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRedhatGuest,
			expected: CIMOSTypeRedHatEnterpriseLinux,
		},
		{
			name:     "RHEL 2 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRhel2Guest,
			expected: CIMOSTypeRedHatEnterpriseLinux,
		},
		{
			name:     "RHEL 3 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRhel3Guest,
			expected: CIMOSTypeRedHatEnterpriseLinux,
		},
		{
			name:     "RHEL 3 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRhel3_64Guest,
			expected: CIMOSTypeRedHatEnterpriseLinux64Bit,
		},
		{
			name:     "RHEL 4 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRhel4Guest,
			expected: CIMOSTypeRedHatEnterpriseLinux,
		},
		{
			name:     "RHEL 4 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRhel4_64Guest,
			expected: CIMOSTypeRedHatEnterpriseLinux64Bit,
		},
		{
			name:     "RHEL 5 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRhel5Guest,
			expected: CIMOSTypeRedHatEnterpriseLinux,
		},
		{
			name:     "RHEL 5 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRhel5_64Guest,
			expected: CIMOSTypeRedHatEnterpriseLinux64Bit,
		},
		{
			name:     "RHEL 6 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRhel6Guest,
			expected: CIMOSTypeRedHatEnterpriseLinux,
		},
		{
			name:     "RHEL 6 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRhel6_64Guest,
			expected: CIMOSTypeRedHatEnterpriseLinux64Bit,
		},
		{
			name:     "RHEL 7 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRhel7Guest,
			expected: CIMOSTypeRedHatEnterpriseLinux,
		},
		{
			name:     "RHEL 7 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRhel7_64Guest,
			expected: CIMOSTypeRedHatEnterpriseLinux64Bit,
		},
		{
			name:     "RHEL 8 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRhel8_64Guest,
			expected: CIMOSTypeRedHatEnterpriseLinux64Bit,
		},
		{
			name:     "RHEL 9 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRhel9_64Guest,
			expected: CIMOSTypeRedHatEnterpriseLinux64Bit,
		},
		{
			name:     "RHEL 10 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRhel10_64Guest,
			expected: CIMOSTypeUnknown,
		},

		// Linux - CentOS
		{
			name:     "CentOS Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierCentosGuest,
			expected: CIMOSTypeUnknown,
		},
		{
			name:     "CentOS 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierCentos64Guest,
			expected: CIMOSTypeUnknown,
		},
		{
			name:     "CentOS 6 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierCentos6Guest,
			expected: CIMOSTypeCentOS32bit,
		},
		{
			name:     "CentOS 6 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierCentos6_64Guest,
			expected: CIMOSTypeCentOS64bit,
		},
		{
			name:     "CentOS 7 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierCentos7Guest,
			expected: CIMOSTypeCentOS32bit,
		},
		{
			name:     "CentOS 7 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierCentos7_64Guest,
			expected: CIMOSTypeCentOS64bit,
		},
		{
			name:     "CentOS 8 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierCentos8_64Guest,
			expected: CIMOSTypeCentOS64bit,
		},
		{
			name:     "CentOS 9 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierCentos9_64Guest,
			expected: CIMOSTypeCentOS64bit,
		},

		// Linux - Oracle
		{
			name:     "Oracle Linux Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOracleLinuxGuest,
			expected: CIMOSTypeOracle32bit,
		},
		{
			name:     "Oracle Linux 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOracleLinux64Guest,
			expected: CIMOSTypeOracle64bit,
		},
		{
			name:     "Oracle Linux 6 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOracleLinux6Guest,
			expected: CIMOSTypeOracle32bit,
		},
		{
			name:     "Oracle Linux 6 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOracleLinux6_64Guest,
			expected: CIMOSTypeOracle64bit,
		},
		{
			name:     "Oracle Linux 7 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOracleLinux7Guest,
			expected: CIMOSTypeOracle32bit,
		},
		{
			name:     "Oracle Linux 7 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOracleLinux7_64Guest,
			expected: CIMOSTypeOracle64bit,
		},
		{
			name:     "Oracle Linux 8 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOracleLinux8_64Guest,
			expected: CIMOSTypeOracle64bit,
		},
		{
			name:     "Oracle Linux 9 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOracleLinux9_64Guest,
			expected: CIMOSTypeOracle64bit,
		},
		{
			name:     "Oracle Linux 10 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOracleLinux10_64Guest,
			expected: CIMOSTypeUnknown,
		},

		// Linux - SUSE
		{
			name:     "SUSE Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSuseGuest,
			expected: CIMOSTypeSUSE,
		},
		{
			name:     "SUSE 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSuse64Guest,
			expected: CIMOSTypeSUSE64Bit,
		},
		{
			name:     "SLES Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSlesGuest,
			expected: CIMOSTypeSLES,
		},
		{
			name:     "SLES 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSles64Guest,
			expected: CIMOSTypeSLES64Bit,
		},
		{
			name:     "SLES 10 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSles10Guest,
			expected: CIMOSTypeSLES,
		},
		{
			name:     "SLES 10 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSles10_64Guest,
			expected: CIMOSTypeSLES64Bit,
		},
		{
			name:     "SLES 11 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSles11Guest,
			expected: CIMOSTypeSLES,
		},
		{
			name:     "SLES 11 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSles11_64Guest,
			expected: CIMOSTypeSLES64Bit,
		},
		{
			name:     "SLES 12 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSles12Guest,
			expected: CIMOSTypeSLES,
		},
		{
			name:     "SLES 12 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSles12_64Guest,
			expected: CIMOSTypeSLES64Bit,
		},
		{
			name:     "SLES 15 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSles15_64Guest,
			expected: CIMOSTypeSLES64Bit,
		},
		{
			name:     "SLES 16 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSles16_64Guest,
			expected: CIMOSTypeSLES64Bit,
		},
		{
			name:     "OpenSUSE Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOpensuseGuest,
			expected: CIMOSTypeSUSE,
		},
		{
			name:     "OpenSUSE 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOpensuse64Guest,
			expected: CIMOSTypeSUSE64Bit,
		},

		// Linux - Ubuntu
		{
			name:     "Ubuntu Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierUbuntuGuest,
			expected: CIMOSTypeUbuntu,
		},
		{
			name:     "Ubuntu 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierUbuntu64Guest,
			expected: CIMOSTypeUbuntu64Bit,
		},

		// Linux - Debian
		{
			name:     "Debian 4 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian4Guest,
			expected: CIMOSTypeDebian,
		},
		{
			name:     "Debian 4 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian4_64Guest,
			expected: CIMOSTypeDebian64Bit,
		},
		{
			name:     "Debian 5 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian5Guest,
			expected: CIMOSTypeDebian,
		},
		{
			name:     "Debian 5 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian5_64Guest,
			expected: CIMOSTypeDebian64Bit,
		},
		{
			name:     "Debian 6 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian6Guest,
			expected: CIMOSTypeDebian,
		},
		{
			name:     "Debian 6 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian6_64Guest,
			expected: CIMOSTypeDebian64Bit,
		},
		{
			name:     "Debian 7 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian7Guest,
			expected: CIMOSTypeDebian,
		},
		{
			name:     "Debian 7 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian7_64Guest,
			expected: CIMOSTypeDebian64Bit,
		},
		{
			name:     "Debian 8 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian8Guest,
			expected: CIMOSTypeDebian,
		},
		{
			name:     "Debian 8 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian8_64Guest,
			expected: CIMOSTypeDebian64Bit,
		},
		{
			name:     "Debian 9 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian9Guest,
			expected: CIMOSTypeDebian,
		},
		{
			name:     "Debian 9 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian9_64Guest,
			expected: CIMOSTypeDebian64Bit,
		},
		{
			name:     "Debian 10 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian10Guest,
			expected: CIMOSTypeDebian,
		},
		{
			name:     "Debian 10 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian10_64Guest,
			expected: CIMOSTypeDebian64Bit,
		},
		{
			name:     "Debian 11 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian11Guest,
			expected: CIMOSTypeDebian,
		},
		{
			name:     "Debian 11 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian11_64Guest,
			expected: CIMOSTypeDebian64Bit,
		},
		{
			name:     "Debian 12 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian12Guest,
			expected: CIMOSTypeDebian,
		},
		{
			name:     "Debian 12 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian12_64Guest,
			expected: CIMOSTypeDebian64Bit,
		},
		{
			name:     "Debian 13 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian13Guest,
			expected: CIMOSTypeUnknown,
		},
		{
			name:     "Debian 13 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDebian13_64Guest,
			expected: CIMOSTypeUnknown,
		},

		// Linux - Fedora
		{
			name:     "Fedora Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFedoraGuest,
			expected: CIMOSTypeLINUX,
		},
		{
			name:     "Fedora 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFedora64Guest,
			expected: CIMOSTypeLinux64Bit,
		},

		// Linux - Mandrake/Mandriva
		{
			name:     "Mandrake Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierMandrakeGuest,
			expected: CIMOSTypeMandriva,
		},
		{
			name:     "Mandriva Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierMandrivaGuest,
			expected: CIMOSTypeMandriva,
		},
		{
			name:     "Mandriva 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierMandriva64Guest,
			expected: CIMOSTypeMandriva64Bit,
		},

		// Linux - TurboLinux
		{
			name:     "TurboLinux Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierTurboLinuxGuest,
			expected: CIMOSTypeTurboLinux,
		},
		{
			name:     "TurboLinux 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierTurboLinux64Guest,
			expected: CIMOSTypeTurboLinux64Bit,
		},

		// Linux - Asianux
		{
			name:     "Asianux 3 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierAsianux3Guest,
			expected: CIMOSTypeLINUX,
		},
		{
			name:     "Asianux 3 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierAsianux3_64Guest,
			expected: CIMOSTypeLinux64Bit,
		},
		{
			name:     "Asianux 4 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierAsianux4Guest,
			expected: CIMOSTypeLINUX,
		},
		{
			name:     "Asianux 4 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierAsianux4_64Guest,
			expected: CIMOSTypeLinux64Bit,
		},
		{
			name:     "Asianux 5 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierAsianux5_64Guest,
			expected: CIMOSTypeLinux64Bit,
		},
		{
			name:     "Asianux 7 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierAsianux7_64Guest,
			expected: CIMOSTypeLinux64Bit,
		},
		{
			name:     "Asianux 8 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierAsianux8_64Guest,
			expected: CIMOSTypeLinux64Bit,
		},
		{
			name:     "Asianux 9 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierAsianux9_64Guest,
			expected: CIMOSTypeLinux64Bit,
		},

		// Linux - Other specific distributions
		{
			name:     "CoreOS 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierCoreos64Guest,
			expected: CIMOSTypeLinux64Bit,
		},
		{
			name:     "VMware Photon 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierVmwarePhoton64Guest,
			expected: CIMOSTypeLinux64Bit,
		},
		{
			name:     "Amazon Linux 2 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierAmazonlinux2_64Guest,
			expected: CIMOSTypeLinux64Bit,
		},
		{
			name:     "Amazon Linux 3 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierAmazonlinux3_64Guest,
			expected: CIMOSTypeLinux64Bit,
		},
		{
			name:     "Rocky Linux 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierRockylinux_64Guest,
			expected: CIMOSTypeLinux64Bit,
		},
		{
			name:     "Miracle Linux 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierMiraclelinux_64Guest,
			expected: CIMOSTypeUnknown,
		},
		{
			name:     "Pardus 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierPardus_64Guest,
			expected: CIMOSTypeUnknown,
		},

		// Linux - Generic
		{
			name:     "Generic Linux Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierGenericLinuxGuest,
			expected: CIMOSTypeLINUX,
		},
		{
			name:     "Other 24x Linux Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOther24xLinuxGuest,
			expected: CIMOSTypeLinux24x,
		},
		{
			name:     "Other 24x Linux 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOther24xLinux64Guest,
			expected: CIMOSTypeLinux24x64Bit,
		},
		{
			name:     "Other 26x Linux Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOther26xLinuxGuest,
			expected: CIMOSTypeLinux26x,
		},
		{
			name:     "Other 26x Linux 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOther26xLinux64Guest,
			expected: CIMOSTypeLinux26x64Bit,
		},
		{
			name:     "Other 3x Linux Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOther3xLinuxGuest,
			expected: CIMOSTypeLINUX,
		},
		{
			name:     "Other 3x Linux 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOther3xLinux64Guest,
			expected: CIMOSTypeLinux64Bit,
		},
		{
			name:     "Other 4x Linux Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOther4xLinuxGuest,
			expected: CIMOSTypeLINUX,
		},
		{
			name:     "Other 4x Linux 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOther4xLinux64Guest,
			expected: CIMOSTypeLinux64Bit,
		},
		{
			name:     "Other 5x Linux Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOther5xLinuxGuest,
			expected: CIMOSTypeLINUX,
		},
		{
			name:     "Other 5x Linux 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOther5xLinux64Guest,
			expected: CIMOSTypeLinux64Bit,
		},
		{
			name:     "Other 6x Linux Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOther6xLinuxGuest,
			expected: CIMOSTypeLINUX,
		},
		{
			name:     "Other 6x Linux 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOther6xLinux64Guest,
			expected: CIMOSTypeLinux64Bit,
		},
		{
			name:     "Other 7x Linux Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOther7xLinuxGuest,
			expected: CIMOSTypeUnknown,
		},
		{
			name:     "Other 7x Linux 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOther7xLinux64Guest,
			expected: CIMOSTypeUnknown,
		},
		{
			name:     "Other Linux Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOtherLinuxGuest,
			expected: CIMOSTypeLINUX,
		},
		{
			name:     "Other Linux 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOtherLinux64Guest,
			expected: CIMOSTypeLinux64Bit,
		},

		// FreeBSD
		{
			name:     "FreeBSD Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFreebsdGuest,
			expected: CIMOSTypeFreeBSD,
		},
		{
			name:     "FreeBSD 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFreebsd64Guest,
			expected: CIMOSTypeFreeBSD64Bit,
		},
		{
			name:     "FreeBSD 11 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFreebsd11Guest,
			expected: CIMOSTypeFreeBSD,
		},
		{
			name:     "FreeBSD 11 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFreebsd11_64Guest,
			expected: CIMOSTypeFreeBSD64Bit,
		},
		{
			name:     "FreeBSD 12 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFreebsd12Guest,
			expected: CIMOSTypeFreeBSD,
		},
		{
			name:     "FreeBSD 12 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFreebsd12_64Guest,
			expected: CIMOSTypeFreeBSD64Bit,
		},
		{
			name:     "FreeBSD 13 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFreebsd13Guest,
			expected: CIMOSTypeFreeBSD,
		},
		{
			name:     "FreeBSD 13 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFreebsd13_64Guest,
			expected: CIMOSTypeFreeBSD64Bit,
		},
		{
			name:     "FreeBSD 14 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFreebsd14Guest,
			expected: CIMOSTypeFreeBSD,
		},
		{
			name:     "FreeBSD 14 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFreebsd14_64Guest,
			expected: CIMOSTypeFreeBSD64Bit,
		},
		{
			name:     "FreeBSD 15 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFreebsd15Guest,
			expected: CIMOSTypeUnknown,
		},
		{
			name:     "FreeBSD 15 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFreebsd15_64Guest,
			expected: CIMOSTypeUnknown,
		},

		// macOS / Darwin
		{
			name:     "Darwin Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwinGuest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin64Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 10 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin10Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 10 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin10_64Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 11 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin11Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 11 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin11_64Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 12 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin12_64Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 13 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin13_64Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 14 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin14_64Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 15 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin15_64Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 16 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin16_64Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 17 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin17_64Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 18 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin18_64Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 19 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin19_64Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 20 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin20_64Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 21 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin21_64Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 22 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin22_64Guest,
			expected: CIMOSTypeMACOS,
		},
		{
			name:     "Darwin 23 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierDarwin23_64Guest,
			expected: CIMOSTypeMACOS,
		},

		// Solaris
		{
			name:     "Solaris 6 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSolaris6Guest,
			expected: CIMOSTypeSolaris,
		},
		{
			name:     "Solaris 7 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSolaris7Guest,
			expected: CIMOSTypeSolaris,
		},
		{
			name:     "Solaris 8 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSolaris8Guest,
			expected: CIMOSTypeSolaris,
		},
		{
			name:     "Solaris 9 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSolaris9Guest,
			expected: CIMOSTypeSolaris,
		},
		{
			name:     "Solaris 10 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSolaris10Guest,
			expected: CIMOSTypeSolaris,
		},
		{
			name:     "Solaris 10 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSolaris10_64Guest,
			expected: CIMOSTypeSolaris64Bit,
		},
		{
			name:     "Solaris 11 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSolaris11_64Guest,
			expected: CIMOSTypeSolaris64Bit,
		},

		// Netware / Novell
		{
			name:     "Netware 4 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierNetware4Guest,
			expected: CIMOSTypeNetWare,
		},
		{
			name:     "Netware 5 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierNetware5Guest,
			expected: CIMOSTypeNetWare,
		},
		{
			name:     "Netware 6 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierNetware6Guest,
			expected: CIMOSTypeNetWare,
		},
		{
			name:     "NLD 9 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierNld9Guest,
			expected: CIMOSTypeNovellLinuxDesktop,
		},
		{
			name:     "OES Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOesGuest,
			expected: CIMOSTypeNovellOES,
		},
		{
			name:     "SJDS Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierSjdsGuest,
			expected: CIMOSTypeUnknown,
		},

		// VMware
		{
			name:     "VMkernel Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierVmkernelGuest,
			expected: CIMOSTypeVMwareESXi,
		},
		{
			name:     "VMkernel 5 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierVmkernel5Guest,
			expected: CIMOSTypeVMwareESXi,
		},
		{
			name:     "VMkernel 6 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierVmkernel6Guest,
			expected: CIMOSTypeVMwareESXi,
		},
		{
			name:     "VMkernel 65 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierVmkernel65Guest,
			expected: CIMOSTypeVMwareESXi,
		},
		{
			name:     "VMkernel 7 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierVmkernel7Guest,
			expected: CIMOSTypeVMwareESXi,
		},
		{
			name:     "VMkernel 8 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierVmkernel8Guest,
			expected: CIMOSTypeVMwareESXi,
		},

		// OS/2
		{
			name:     "OS/2 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOs2Guest,
			expected: CIMOSTypeOS2,
		},
		{
			name:     "eComStation Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierEComStationGuest,
			expected: CIMOSTypeeComStation32bitx,
		},
		{
			name:     "eComStation 2 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierEComStation2Guest,
			expected: CIMOSTypeeComStation32bitx,
		},

		// Unix
		{
			name:     "UnixWare 7 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierUnixWare7Guest,
			expected: CIMOSTypeSCOUnixWare,
		},
		{
			name:     "OpenServer 5 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOpenServer5Guest,
			expected: CIMOSTypeUnknown,
		},
		{
			name:     "OpenServer 6 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOpenServer6Guest,
			expected: CIMOSTypeUnknown,
		},

		// Other
		{
			name:     "Other Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierOtherGuest,
			expected: CIMOSTypeOther,
		},
		{
			name:     "Other Guest 64",
			guestID:  types.VirtualMachineGuestOsIdentifierOtherGuest64,
			expected: CIMOSTypeOther64Bit,
		},
		{
			name:     "CRX Pod 1 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierCrxPod1Guest,
			expected: CIMOSTypeOther,
		},
		{
			name:     "Fusion OS 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierFusionos_64Guest,
			expected: CIMOSTypeUnknown,
		},
		{
			name:     "ProLinux 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierProlinux_64Guest,
			expected: CIMOSTypeUnknown,
		},
		{
			name:     "KylinLinux 64 Guest",
			guestID:  types.VirtualMachineGuestOsIdentifierKylinlinux_64Guest,
			expected: CIMOSTypeUnknown,
		},

		// Custom string tests (not enum constants)
		{
			name:     "Unknown Guest ID",
			guestID:  types.VirtualMachineGuestOsIdentifier("unknownGuestID"),
			expected: CIMOSTypeUnknown,
		},
		{
			name:     "Empty Guest ID",
			guestID:  types.VirtualMachineGuestOsIdentifier(""),
			expected: CIMOSTypeUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GuestIDToCIMOSType(tt.guestID)
			if result != tt.expected {
				t.Errorf("GuestIDToCIMOSType(%q) = %v, want %v", tt.guestID, result, tt.expected)
			}
		})
	}
}
