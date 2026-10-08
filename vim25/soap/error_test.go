// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package soap

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"testing"

	"github.com/vmware/govmomi/vim25/types"
)

func TestIsCertificateUntrusted(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "tls.CertificateVerificationError",
			err: x509.HostnameError{
				Certificate: &x509.Certificate{},
				Host:        "1.2.3.4",
			},
			want: true,
		},
		{
			name: "tls.CertificateVerificationError",
			err: &tls.CertificateVerificationError{
				UnverifiedCertificates: []*x509.Certificate{
					&x509.Certificate{},
				},
				Err: x509.HostnameError{
					Certificate: &x509.Certificate{},
					Host:        "5.6.7.8",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCertificateUntrusted(tt.err); got != tt.want {
				t.Errorf("IsCertificateUntrusted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrappedFaults(t *testing.T) {
	sf := &Fault{String: "boom"}
	sf.Detail.Fault = types.NotFound{}

	err := fmt.Errorf("outer: %w", WrapSoapFault(sf))
	if !IsSoapFault(err) {
		t.Fatal("expected IsSoapFault on wrapped error")
	}
	if ToSoapFault(err) != sf {
		t.Fatal("expected ToSoapFault to return the wrapped fault")
	}
	if IsSoapFault(fmt.Errorf("plain")) || ToSoapFault(fmt.Errorf("plain")) != nil {
		t.Fatal("unexpected soap fault for plain error")
	}

	vf := &types.NotFound{}
	err = fmt.Errorf("outer: %w", WrapVimFault(vf))
	if !IsVimFault(err) || ToVimFault(err) != types.BaseMethodFault(vf) {
		t.Fatal("expected vim fault from wrapped error")
	}
	if IsVimFault(fmt.Errorf("plain")) || ToVimFault(fmt.Errorf("plain")) != nil {
		t.Fatal("unexpected vim fault for plain error")
	}
}
