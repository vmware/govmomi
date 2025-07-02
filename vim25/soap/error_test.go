// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package soap

import (
	"crypto/tls"
	"crypto/x509"
	"testing"
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
