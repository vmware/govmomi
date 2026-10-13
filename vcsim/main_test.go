// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import "testing"

func TestReleaseBuild(t *testing.T) {
	tests := []struct {
		name    string
		version string
		esx     bool
		want    string
	}{
		{name: "vCenter 9.0 GA", version: "9.0.0", want: "24755230"},
		{name: "ESXi 9.0 GA", version: "9.0.0", esx: true, want: "24755229"},
		{name: "vCenter 9.0.2", version: "9.0.2.0", want: "25148086"},
		{name: "ESXi 9.0.2", version: "9.0.2.0", esx: true, want: "25148076"},
		{name: "unknown release", version: "99.0.0"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, ok := releaseBuild(test.version, test.esx)
			if test.want == "" {
				if ok {
					t.Fatalf("releaseBuild(%q, %t) = %q, want no mapping", test.version, test.esx, got)
				}
				return
			}
			if !ok || got != test.want {
				t.Fatalf("releaseBuild(%q, %t) = %q, %t; want %q, true", test.version, test.esx, got, ok, test.want)
			}
		})
	}
}

func TestReleaseAPIVersions(t *testing.T) {
	tests := []struct {
		version     string
		wantServer  string
		wantHost    string
		wantSuccess bool
	}{
		{version: "9.0.0", wantServer: "9.0", wantHost: "9.0.0.0", wantSuccess: true},
		{version: "8.0.2", wantServer: "8.0", wantHost: "8.0.2.0", wantSuccess: true},
		{version: "9.0.0.0100", wantServer: "9.0", wantHost: "9.0.0.0", wantSuccess: true},
		{version: "invalid"},
	}

	for _, test := range tests {
		t.Run(test.version, func(t *testing.T) {
			server, host, ok := releaseAPIVersions(test.version)
			if server != test.wantServer || host != test.wantHost || ok != test.wantSuccess {
				t.Fatalf("releaseAPIVersions(%q) = %q, %q, %t; want %q, %q, %t", test.version, server, host, ok, test.wantServer, test.wantHost, test.wantSuccess)
			}
		})
	}
}
