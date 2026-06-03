// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package library

import (
	"strings"
	"testing"
)

func TestReadManifest(t *testing.T) {
	tests := []struct {
		name     string
		manifest string
		want     map[string]Checksum
	}{
		{
			name:     "compact",
			manifest: "SHA1(name.vmdk)= 344a536fab6782622b6beb923798e84134bd4cbd",
			want: map[string]Checksum{
				"name.vmdk": {
					Algorithm: "SHA1",
					Checksum:  "344a536fab6782622b6beb923798e84134bd4cbd",
				},
			},
		},
		{
			name:     "spaced",
			manifest: "SHA1(name.vmdk) = 344a536fab6782622b6beb923798e84134bd4cbd",
			want: map[string]Checksum{
				"name.vmdk": {
					Algorithm: "SHA1",
					Checksum:  "344a536fab6782622b6beb923798e84134bd4cbd",
				},
			},
		},
		{
			name: "mixed multi-line",
			manifest: `SHA1(ttylinux-pc_i486-16.1.ovf)= 344a536fab6782622b6beb923798e84134bd4cbd
SHA1(ttylinux-pc_i486-16.1-disk1.vmdk) = ed64564a37366bfe1c93af80e2ead0cbd398c3d3`,
			want: map[string]Checksum{
				"ttylinux-pc_i486-16.1.ovf": {
					Algorithm: "SHA1",
					Checksum:  "344a536fab6782622b6beb923798e84134bd4cbd",
				},
				"ttylinux-pc_i486-16.1-disk1.vmdk": {
					Algorithm: "SHA1",
					Checksum:  "ed64564a37366bfe1c93af80e2ead0cbd398c3d3",
				},
			},
		},
		{
			name:     "sha256",
			manifest: "SHA256(name.ovf) = 7437f1b90e40ced4f27a33997c15028204b72eb3",
			want: map[string]Checksum{
				"name.ovf": {
					Algorithm: "SHA256",
					Checksum:  "7437f1b90e40ced4f27a33997c15028204b72eb3",
				},
			},
		},
		{
			name:     "filename with parens",
			manifest: "SHA256(foo (1).ovf)= 7437f1b90e40ced4f27a33997c15028204b72eb3",
			want: map[string]Checksum{
				"foo (1).ovf": {
					Algorithm: "SHA256",
					Checksum:  "7437f1b90e40ced4f27a33997c15028204b72eb3",
				},
			},
		},
		{
			name:     "filename with equals",
			manifest: "SHA1(foo=bar.vmdk)= 344a536fab6782622b6beb923798e84134bd4cbd",
			want: map[string]Checksum{
				"foo=bar.vmdk": {
					Algorithm: "SHA1",
					Checksum:  "344a536fab6782622b6beb923798e84134bd4cbd",
				},
			},
		},
		{
			name: "blank and malformed lines",
			manifest: `
SHA1(name.vmdk = 344a536fab6782622b6beb923798e84134bd4cbd
`,
			want: map[string]Checksum{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sums, err := ReadManifest(strings.NewReader(tc.manifest))
			if err != nil {
				t.Fatal(err)
			}

			if len(sums) != len(tc.want) {
				t.Fatalf("len(sums) = %d; want %d", len(sums), len(tc.want))
			}

			for name, want := range tc.want {
				got, ok := sums[name]
				if !ok {
					t.Fatalf("missing checksum for %q", name)
				}
				if got.Algorithm != want.Algorithm {
					t.Errorf("Algorithm = %q; want %q", got.Algorithm, want.Algorithm)
				}
				if got.Checksum != want.Checksum {
					t.Errorf("Checksum = %q; want %q", got.Checksum, want.Checksum)
				}
			}
		})
	}
}
