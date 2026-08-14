// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmdk

import (
	"testing"
)

func TestDeriveImportTargets(t *testing.T) {
	// common local disk info
	disk := &Info{
		Name:       "local.vmdk",
		ImportName: "local",
	}

	tests := []struct {
		name       string
		params     ImportParams
		wantEntity string
		wantTarget string
		wantError  bool
	}{
		{
			name:       "Full Path",
			params:     ImportParams{Path: "windows/win2019/win2019.vmdk"},
			wantEntity: "win2019",
			wantTarget: "windows/win2019/win2019.vmdk",
			wantError:  false,
		},
		{
			name:       "Only vmdk name",
			params:     ImportParams{Path: "win2019.vmdk"},
			wantEntity: "win2019",
			wantTarget: "win2019/win2019.vmdk",
			wantError:  false,
		},
		{
			name:      "Only target path",
			params:    ImportParams{Path: "windows"},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity, target, err := deriveImportTargets(tt.params, disk)
			if tt.wantError && err == nil {
				t.Fatalf("error: got %v want %v", err, tt.wantError)
			}
			if entity != tt.wantEntity {
				t.Fatalf("entity: got %q want %q", entity, tt.wantEntity)
			}
			if target != tt.wantTarget {
				t.Fatalf("target: got %q want %q", target, tt.wantTarget)
			}
		})
	}
}
