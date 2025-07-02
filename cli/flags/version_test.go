// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package flags

import (
	"reflect"
	"testing"
)

func TestParseVersion(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name    string
		args    args
		want    version
		wantErr bool
	}{
		{name: "5.5.5.5", args: args{version: "5.5.5.5"}, want: version{5, 5, 5, 5}, wantErr: false},
		{name: "0.24.0", args: args{version: "0.24.0"}, want: version{0, 24, 0}, wantErr: false},
		{name: "v0.24.0", args: args{version: "v0.24.0"}, want: version{0, 24, 0}, wantErr: false},
		{name: "v0.24.0-next", args: args{version: "v0.24.0-next"}, want: version{0, 24, 0}, wantErr: false},
		{name: "0.10x", args: args{version: "0.10x"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseVersion(tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLte(t *testing.T) {
	v1, err := ParseVersion("5.5")
	if err != nil {
		panic(err)
	}

	v2, err := ParseVersion("5.6")
	if err != nil {
		panic(err)
	}

	if !v1.Lte(v1) {
		t.Errorf("Expected 5.5 <= 5.5")
	}

	if !v1.Lte(v2) {
		t.Errorf("Expected 5.5 <= 5.6")
	}

	if v2.Lte(v1) {
		t.Errorf("Expected not 5.6 <= 5.5")
	}
}
