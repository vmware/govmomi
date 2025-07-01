// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package object

import "testing"

func TestHostCertificateManagerCertificateInfo(t *testing.T) {
	subject := "emailAddress=vmca@vmware.com,CN=w2-xlr8-autoroot-esx004.eng.vmware.com,OU=VMware Engineering,O=VMware,L=Palo Alto,ST=California,C=US"

	var info HostCertificateInfo
	name := info.toName(subject)
	s := info.fromName(name)
	if subject != s {
		t.Errorf("%s != %s", s, subject)
	}
}
