// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmdk_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi/crypto"
	"github.com/vmware/govmomi/vmdk"
)

func TestDescriptor(t *testing.T) {
	t.Run("Write", func(t *testing.T) {
		desc := &vmdk.Descriptor{
			Version:  1,
			Encoding: "UTF-8",
			CID:      123,
			Type:     "vmfs",
			Extent: []vmdk.Extent{{
				Type:       "VMFS",
				Permission: "RW",
				Size:       1024,
				Info:       "test-flat.vmdk",
			}},
			DDB: map[string]string{
				"adapterType":      "lsilogic",
				"virtualHWVersion": "14",
			},
			EncryptionKeys: &crypto.KeyLocator{
				Type: crypto.KeyLocatorTypeFQID,
				Indirect: &crypto.KeyLocatorIndirect{
					Type:     crypto.KeyLocatorTypeFQID,
					UniqueID: "my-unique-id",
					FQID: crypto.KeyLocatorFQIDParams{
						KeyServerID: "production-server",
						KeyID:       "encryption-key-001",
					},
				},
			},
		}

		var buf bytes.Buffer

		err := desc.Write(&buf)
		if err != nil {
			t.Fatal(err)
		}

		parsed, err := vmdk.ParseDescriptor(&buf)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, desc, parsed, cmp.Diff(desc, parsed))
	})

	t.Run("Parse", func(t *testing.T) {
		act, err := vmdk.ParseDescriptor(strings.NewReader(descriptorText))
		assert.NoError(t, err)

		exp := &vmdk.Descriptor{
			Version:  1,
			Encoding: "UTF-8",
			CID:      123,
			Type:     "vmfs",
			Extent: []vmdk.Extent{{
				Type:       "VMFS",
				Permission: "RW",
				Size:       1024,
				Info:       "test-flat.vmdk",
			}},
			DDB: map[string]string{
				"adapterType":      "lsilogic",
				"virtualHWVersion": "14",
			},
			EncryptionKeys: &crypto.KeyLocator{
				Type: crypto.KeyLocatorTypeFQID,
				Indirect: &crypto.KeyLocatorIndirect{
					Type:     crypto.KeyLocatorTypeFQID,
					UniqueID: "my-unique-id",
					FQID: crypto.KeyLocatorFQIDParams{
						KeyServerID: "production-server",
						KeyID:       "encryption-key-001",
					},
				},
			},
		}

		assert.Equal(t, exp, act, cmp.Diff(exp, act))
	})
}

const descriptorText = `# Disk DescriptorFile
version=1
encoding="UTF-8"
CID=0000007b
parentCID=00000000
createType="vmfs"
encryptionKeys="vmware:key/fqid/my%2dunique%2did/production%2dserver/encryption%2dkey%2d001"

# Extent description (512.0KB capacity)
RW 1024 VMFS "test-flat.vmdk"

# The Disk Data Base
#DDB
ddb.adapterType = "lsilogic"
ddb.virtualHWVersion = "14"`
