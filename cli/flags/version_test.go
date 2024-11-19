/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

func TestDevelopmentVersion(t *testing.T) {
	if !isDevelopmentVersion("6.5.x") {
		t.Error("expected true")
	}

	if !isDevelopmentVersion("r4A70F") {
		t.Error("expected true")
	}

	if isDevelopmentVersion("6.5") {
		t.Error("expected false")
	}
}
