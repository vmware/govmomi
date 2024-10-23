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
