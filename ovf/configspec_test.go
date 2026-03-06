// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package ovf

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vmware/govmomi/vim25/types"
)

func TestEnvelopeToConfigSpec(t *testing.T) {

	t.Run("Strict", func(t *testing.T) {
		t.Run("Unsupported ResourceType", func(t *testing.T) {
			e := testEnvelope(t, "fixtures/unsupported-resourcetype.ovf")
			_, err := e.ToConfigSpecWithOptions(ToConfigSpecOptions{Strict: true})
			if assert.Error(t, err) {
				assert.Contains(t, err.Error(), "unsupported resource type")
				if ovfErr, ok := AsErrUnsupportedItem(err); assert.True(t, ok) {
					assert.Equal(t, ovfErr.Index, 2)
					assert.Equal(t, ovfErr.InstanceID, "3")
					assert.Equal(t, ovfErr.Name, "invalidResourceType")
					assert.Equal(t, ovfErr.ResourceType, CIMResourceType(35))
					assert.Equal(t, ovfErr.ResourceSubType, "")
				}
			}
		})
		t.Run("Unsupported ResourceSubType", func(t *testing.T) {
			e := testEnvelope(t, "fixtures/unsupported-resourcesubtype.ovf")
			_, err := e.ToConfigSpecWithOptions(ToConfigSpecOptions{Strict: true})
			if assert.Error(t, err) {
				assert.Contains(t, err.Error(), "unsupported resource subtype")
				if ovfErr, ok := AsErrUnsupportedItem(err); assert.True(t, ok) {
					assert.Equal(t, ovfErr.Index, 2)
					assert.Equal(t, ovfErr.InstanceID, "3")
					assert.Equal(t, ovfErr.Name, "invalidResourceSubType")
					assert.Equal(t, ovfErr.ResourceType, CIMResourceType(1))
					assert.Equal(t, ovfErr.ResourceSubType, "invalidresourcesubtype")
				}
			}
		})
	})

	t.Run("VirtualSystemCollection", func(t *testing.T) {
		t.Run("No index", func(t *testing.T) {
			e := testEnvelope(t, "fixtures/virtualsystemcollection.ovf")
			_, err := e.ToConfigSpec()
			assert.Error(t, err, "no VirtualSystem")
		})

		t.Run("Index 0", func(t *testing.T) {
			e := testEnvelope(t, "fixtures/virtualsystemcollection.ovf")
			configSpec, err := e.ToConfigSpecWithOptions(ToConfigSpecOptions{
				VirtualSystemCollectionIndex: types.New(int(0)),
			})
			assert.NoError(t, err)
			assert.NotEmpty(t, configSpec)

			// Verify configSpec matches first VirtualSystem (id="storage server")
			assert.Equal(t, "storage server", configSpec.Name) // Uses VirtualSystem id attribute
			assert.Equal(t, int64(1), configSpec.MemoryMB)     // 1 GByte specified as VirtualQuantity=1 with AllocationUnits="byte*2^30"
			assert.Equal(t, int32(1), configSpec.NumCPUs)      // 1 virtual CPU

			// Verify VApp configuration from ProductSection
			if va, ok := configSpec.VAppConfig.(*types.VAppConfigSpec); assert.True(t, ok) {
				// Check product info
				if assert.Len(t, va.Product, 1) {
					assert.Equal(t, "The Great Appliance", va.Product[0].Info.Name)
					assert.Equal(t, "Some Great Corporation", va.Product[0].Info.Vendor)
					assert.Equal(t, "13.00", va.Product[0].Info.Version)
					assert.Equal(t, "13.00-b5", va.Product[0].Info.FullVersion)
					assert.Equal(t, "http://www.somegreatcorporation.com/greatappliance", va.Product[0].Info.ProductUrl)
					assert.Equal(t, "http://www.somegreatcorporation.com/", va.Product[0].Info.VendorUrl)
				}

				// Check properties
				if assert.Len(t, va.Property, 2) {
					// Check adminemail property
					assert.Equal(t, "adminemail", va.Property[0].Info.Id)
					assert.Equal(t, "string", va.Property[0].Info.Type)
					assert.Equal(t, "Email address of administrator", va.Property[0].Info.Description)
					assert.Equal(t, "", va.Property[0].Info.DefaultValue) // No default value in OVF

					assert.NotNil(t, va.Property[0].Info.UserConfigurable)
					assert.Equal(t, false, *va.Property[0].Info.UserConfigurable)

					// Check app_ip property
					assert.Equal(t, "app_ip", va.Property[1].Info.Id)
					assert.Equal(t, "string", va.Property[1].Info.Type)
					assert.Equal(t, "The IP address of this appliance", va.Property[1].Info.Description)
					assert.Equal(t, "192.168.0.10", va.Property[1].Info.DefaultValue)
					assert.NotNil(t, va.Property[1].Info.UserConfigurable)
					assert.True(t, *va.Property[1].Info.UserConfigurable)
					assert.Empty(t, va.Property[1].Info.Value)

					assert.NotNil(t, va.Property[1].Info.UserConfigurable)
					assert.Equal(t, true, *va.Property[1].Info.UserConfigurable)
				}
			}
		})

		t.Run("Index 1", func(t *testing.T) {
			e := testEnvelope(t, "fixtures/virtualsystemcollection.ovf")
			configSpec, err := e.ToConfigSpecWithOptions(ToConfigSpecOptions{
				VirtualSystemCollectionIndex: types.New(int(1)),
			})
			assert.NoError(t, err)
			assert.NotEmpty(t, configSpec)

			// Verify configSpec matches second VirtualSystem (id="web-server")
			assert.Equal(t, "web-server", configSpec.Name) // Uses VirtualSystem id attribute
			assert.Equal(t, int64(1), configSpec.MemoryMB) // 1 GByte specified as VirtualQuantity=1 with AllocationUnits="byte*2^30"
			assert.Equal(t, int32(1), configSpec.NumCPUs)  // 1 virtual CPU

			// Verify VApp configuration from ProductSection
			if va, ok := configSpec.VAppConfig.(*types.VAppConfigSpec); assert.True(t, ok) {
				// Check product info
				if assert.Len(t, va.Product, 1) {
					assert.Equal(t, "The Great Appliance", va.Product[0].Info.Name)
					assert.Equal(t, "Some Great Corporation", va.Product[0].Info.Vendor)
					assert.Equal(t, "13.00", va.Product[0].Info.Version)
					assert.Equal(t, "13.00-b5", va.Product[0].Info.FullVersion)
					assert.Equal(t, "http://www.somegreatcorporation.com/greatappliance", va.Product[0].Info.ProductUrl)
					assert.Equal(t, "http://www.somegreatcorporation.com/", va.Product[0].Info.VendorUrl)
				}

				// Check properties
				if assert.Len(t, va.Property, 2) {
					// Check adminemail property
					assert.Equal(t, "adminemail", va.Property[0].Info.Id)
					assert.Equal(t, "string", va.Property[0].Info.Type)
					assert.Equal(t, "Email address of administrator", va.Property[0].Info.Description)
					assert.Equal(t, "", va.Property[0].Info.DefaultValue) // No default value in OVF

					assert.NotNil(t, va.Property[0].Info.UserConfigurable)
					assert.Equal(t, false, *va.Property[0].Info.UserConfigurable)

					// Check app_ip property
					assert.Equal(t, "app_ip", va.Property[1].Info.Id)
					assert.Equal(t, "string", va.Property[1].Info.Type)
					assert.Equal(t, "The IP address of this appliance", va.Property[1].Info.Description)
					assert.Equal(t, "192.168.0.10", va.Property[1].Info.DefaultValue)
					assert.Empty(t, va.Property[1].Info.Value)

					assert.NotNil(t, va.Property[1].Info.UserConfigurable)
					assert.Equal(t, false, *va.Property[1].Info.UserConfigurable)
				}
			}
		})
	})

	t.Run("vApp property type parsing", func(t *testing.T) {
		e := testEnvelope(t, "fixtures/properties.ovf")
		cs, err := e.ToConfigSpec()
		require.NoError(t, err)

		va, ok := cs.VAppConfig.(*types.VAppConfigSpec)
		require.True(t, ok)

		findProp := func(id string) *types.VAppPropertyInfo {
			for i := range va.Property {
				if va.Property[i].Info.Id == id {
					return va.Property[i].Info
				}
			}
			return nil
		}

		t.Run("boolean default value is normalised to canonical form", func(t *testing.T) {
			p := findProp("enable_ssh")
			require.NotNil(t, p, "enable_ssh property not found in VAppConfigSpec")
			// OVF envelope has ovf:value="false" (lowercase); vSphere API requires "False".
			assert.Equal(t, "False", p.DefaultValue)
		})

		t.Run("empty default value is preserved as empty string", func(t *testing.T) {
			p := findProp("ntp-server")
			require.NotNil(t, p, "ntp-server property not found in VAppConfigSpec")
			// OVF envelope has no ovf:value; empty default must pass through without error.
			assert.Equal(t, "", p.DefaultValue)
		})

		t.Run("whitespace-padded default value is trimmed", func(t *testing.T) {
			p := findProp("whitespace_int")
			require.NotNil(t, p, "whitespace_int property not found in VAppConfigSpec")
			// OVF envelope has ovf:value="  42  "; trimmed value must be stored.
			assert.Equal(t, "42", p.DefaultValue)
		})

		t.Run("OVF qualifiers in VAppPropertyInfo.Type per DSP0243 9.5.1",
			func(t *testing.T) {
				// ntp-server: MinLen(1),MaxLen(65535) => string(1..65535)
				p := findProp("ntp-server")
				require.NotNil(t, p)
				assert.Equal(t, "string(1..65535)", p.Type)

				// nfs_mount: MinLen(0),MaxLen(65535) => string(0..65535)
				p = findProp("nfs_mount")
				require.NotNil(t, p)
				assert.Equal(t, "string(0..65535)", p.Type)

				// vmname has no qualifiers => plain "string"
				p = findProp("vmname")
				require.NotNil(t, p)
				assert.Equal(t, "string", p.Type)

				// whitespace_int: ovf:type="int" (no qualifiers) => "int"
				p = findProp("whitespace_int")
				require.NotNil(t, p)
				assert.Equal(t, "int", p.Type)
			})
	})

	t.Run("category grouping (DSP0243 9.5.1)", func(t *testing.T) {
		e := testEnvelope(t, "fixtures/category-grouping.ovf")
		findProp := func(va *types.VAppConfigSpec, id string) *types.VAppPropertyInfo {
			for i := range va.Property {
				if va.Property[i].Info.Id == id {
					return va.Property[i].Info
				}
			}
			return nil
		}
		// Fixture has DeploymentOptionSection (default, extended), ProductSection with categories
		// Net, Storage, Security and properties with/without ovf:configuration.

		t.Run("default config: categories and config filtering", func(t *testing.T) {
			cs, err := e.ToConfigSpec()
			require.NoError(t, err)
			va, ok := cs.VAppConfig.(*types.VAppConfigSpec)
			require.True(t, ok)
			// Only common properties (no ovf:configuration or configuration matches default).
			require.Len(t, va.Property, 3, "default config should have 3 vApp properties")
			p := findProp(va, "net_common")
			require.NotNil(t, p)
			assert.Equal(t, "Net", p.Category)
			p = findProp(va, "storage_common")
			require.NotNil(t, p)
			assert.Equal(t, "Storage", p.Category)
			p = findProp(va, "auth_mode")
			require.NotNil(t, p)
			assert.Equal(t, "Security", p.Category)
			assert.Nil(t, findProp(va, "net_extended"), "extended-only property should be absent in default config")
			assert.Nil(t, findProp(va, "storage_extended"), "extended-only property should be absent in default config")
			assert.Nil(t, findProp(va, "extended_secret"), "extended-only property should be absent in default config")
		})

		t.Run("extended config: all categories and config-specific properties", func(t *testing.T) {
			cs, err := e.ToConfigSpecWithOptions(ToConfigSpecOptions{
				DeploymentConfiguration: "extended",
			})
			require.NoError(t, err)
			va, ok := cs.VAppConfig.(*types.VAppConfigSpec)
			require.True(t, ok)
			require.Len(t, va.Property, 6, "extended config should have 6 vApp properties")
			// Net category: net_common, net_extended
			p := findProp(va, "net_common")
			require.NotNil(t, p)
			assert.Equal(t, "Net", p.Category)
			p = findProp(va, "net_extended")
			require.NotNil(t, p)
			assert.Equal(t, "Net", p.Category)
			// Storage category: storage_common, storage_extended
			p = findProp(va, "storage_common")
			require.NotNil(t, p)
			assert.Equal(t, "Storage", p.Category)
			p = findProp(va, "storage_extended")
			require.NotNil(t, p)
			assert.Equal(t, "Storage", p.Category)
			// Security category: auth_mode, extended_secret
			p = findProp(va, "auth_mode")
			require.NotNil(t, p)
			assert.Equal(t, "Security", p.Category)
			p = findProp(va, "extended_secret")
			require.NotNil(t, p)
			assert.Equal(t, "Security", p.Category)
		})
	})

	t.Run("OVF property types and qualifiers to vAppPropertyInfo.Type", func(t *testing.T) {
		e := testEnvelope(t, "fixtures/property-types.ovf")
		cs, err := e.ToConfigSpec()
		require.NoError(t, err)
		va, ok := cs.VAppConfig.(*types.VAppConfigSpec)
		require.True(t, ok)

		findProp := func(id string) *types.VAppPropertyInfo {
			for i := range va.Property {
				if va.Property[i].Info.Id == id {
					return va.Property[i].Info
				}
			}
			return nil
		}

		// All OVF property type and qualifier combinations and expected vApp Type.
		expect := map[string]string{
			// string: no qualifiers
			"p_string_plain": "string",
			// string: MinLen only => string(min..)
			"p_string_min_only": "string(5..)",
			// string: MaxLen only => string(0..max) (minLen defaults to 0)
			"p_string_max_only": "string(0..100)",
			// string: MinLen + MaxLen => string(min..max)
			"p_string_min_max": "string(1..253)",
			// string: ValueMap => string["choice1", "choice2", ...]
			"p_string_valuemap":        `string["a", "b", "c"]`,
			"p_string_valuemap_single": `string["only"]`,
			// password: no qualifiers
			"p_password_plain": "password",
			// password: MinLen only => password(min..)
			"p_password_min_only": "password(8..)",
			// boolean
			"p_boolean": "boolean",
			// int: no qualifiers
			"p_int_plain": "int",
			// int: ValueMap numeric => int(min..max)
			"p_int_valuemap": "int(1..5)",
			// OVF integer types (DSP0243 Table 6) → int
			"p_uint8": "int", "p_sint8": "int", "p_uint16": "int", "p_sint16": "int",
			"p_uint32": "int", "p_sint32": "int", "p_uint64": "int", "p_sint64": "int",
			// real32, real64 → real
			"p_real32": "real", "p_real64": "real",
			// Case insensitivity: String, Boolean (DSP0243 Table 6)
			"p_String_capital": "string", "p_Boolean_capital": "boolean",
		}

		for id, wantType := range expect {
			p := findProp(id)
			require.NotNil(t, p, "property %q not found in VAppConfigSpec", id)
			assert.Equal(t, wantType, p.Type,
				"OVF property %q: expected vAppPropertyInfo.Type %q, got %q",
				id, wantType, p.Type)
		}
	})

	t.Run("Photon 5", func(t *testing.T) {
		e := testEnvelope(t, "fixtures/photon5.ovf")
		cs, err := e.ToConfigSpec()
		assert.NoError(t, err)
		assert.NotEmpty(t, cs)

		if testing.Verbose() {
			var w bytes.Buffer
			enc := types.NewJSONEncoder(&w)
			enc.SetIndent("", "    ")
			assert.NoError(t, enc.Encode(cs))
			t.Logf("\n\nconfigSpec=%s\n\n", w.String())
		}
	})

	t.Run("Ubuntu 24.10", func(t *testing.T) {
		e := testEnvelope(t, "fixtures/ubuntu24.10.ovf")
		cs, err := e.ToConfigSpec()
		assert.NoError(t, err)
		assert.NotEmpty(t, cs)

		if testing.Verbose() {
			var w bytes.Buffer
			enc := types.NewJSONEncoder(&w)
			enc.SetIndent("", "    ")
			assert.NoError(t, enc.Encode(cs))
			t.Logf("\n\nconfigSpec=%s\n\n", w.String())
		}
	})

	t.Run("Mixed file-backed and empty disks", func(t *testing.T) {
		// Per DMTF OVF 2.1.1 (DSP0243): disks with ovf:fileRef in DiskSection are file-backed;
		// disks without fileRef are empty (created with capacity only). ConfigSpec indicates
		// file-backed disks via non-empty DiskBackingInfo.FileName (see commit 3c3ed53).
		e := testEnvelope(t, "fixtures/mixed-disks.ovf")
		cs, err := e.ToConfigSpec()
		require.NoError(t, err)
		require.NotEmpty(t, cs.DeviceChange)

		var disks []*types.VirtualDisk
		for _, dc := range cs.DeviceChange {
			if d, ok := dc.GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); ok {
				disks = append(disks, d)
			}
		}
		require.Len(t, disks, 2, "fixture has one file-backed and one empty disk")

		db1, ok := disks[0].Backing.(*types.VirtualDiskFlatVer2BackingInfo)
		require.True(t, ok)
		assert.NotEmpty(t, db1.FileName, "first disk (vmdisk1) is file-backed: FileName must be non-empty")
		assert.Equal(t, "os.vmdk", db1.FileName, "fileRef from OVF becomes backing FileName (CreateDisk appends .vmdk)")
		assert.Equal(t, int64(20*1024*1024*1024), disks[0].CapacityInBytes)

		db2, ok := disks[1].Backing.(*types.VirtualDiskFlatVer2BackingInfo)
		require.True(t, ok)
		assert.Empty(t, db2.FileName, "second disk (vmdisk2) has no fileRef: FileName must be empty (empty disk)")
		assert.Equal(t, int64(10*1024*1024*1024), disks[1].CapacityInBytes)

		// FileOperation: file-backed disks (non-empty FileName) get FileOperation="" (existing file);
		// empty disks get FileOperation=Create (see ToConfigSpecWithOptions / toHardware).
		for _, dc := range cs.DeviceChange {
			spec := dc.GetVirtualDeviceConfigSpec()
			if d, ok := spec.Device.(*types.VirtualDisk); ok {
				b, ok := d.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
				require.True(t, ok)
				if b.FileName != "" {
					assert.Equal(t, types.VirtualDeviceConfigSpecFileOperation(""), spec.FileOperation,
						"file-backed disk must have empty FileOperation (existing file)")
				} else {
					assert.Equal(t, types.VirtualDeviceConfigSpecFileOperationCreate, spec.FileOperation,
						"empty disk must have FileOperation Create")
				}
			}
		}
	})

	t.Run("File-backed disk name from path.Base(File.Href)", func(t *testing.T) {
		// When a disk has ovf:fileRef, backing FileName is set from path.Base(File.Href)
		// so that file-backed disks are distinguished from empty disks.
		e := testEnvelope(t, "fixtures/file-ref-path.ovf")
		cs, err := e.ToConfigSpec()
		require.NoError(t, err)
		require.NotEmpty(t, cs.DeviceChange)

		var disks []*types.VirtualDisk
		for _, dc := range cs.DeviceChange {
			if d, ok := dc.GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); ok {
				disks = append(disks, d)
			}
		}
		require.Len(t, disks, 1)
		db, ok := disks[0].Backing.(*types.VirtualDiskFlatVer2BackingInfo)
		require.True(t, ok)
		assert.Equal(t, "data.vmdk", db.FileName,
			"backing FileName must be path.Base(File.Href), i.e. data.vmdk from subdir/data.vmdk")
	})

	t.Run("Disk capacity from property reference (DSP0243 9.1)", func(t *testing.T) {
		// Empty disk with ovf:capacity="${disksize}"; property disksize=5, units byte*2^30 → 5 GiB.
		e := testEnvelope(t, "fixtures/disk-capacity-property.ovf")
		cs, err := e.ToConfigSpec()
		require.NoError(t, err)
		require.NotEmpty(t, cs.DeviceChange)
		var disks []*types.VirtualDisk
		for _, dc := range cs.DeviceChange {
			if d, ok := dc.GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); ok {
				disks = append(disks, d)
			}
		}
		require.Len(t, disks, 1)
		assert.Equal(t, int64(5*1024*1024*1024), disks[0].CapacityInBytes,
			"capacity from property disksize=5 with byte*2^30 units must be 5 GiB")
	})

	t.Run("Property key invalid character (DSP0243 9.5.1)", func(t *testing.T) {
		// ovf:key shall not contain period or colon.
		e := testEnvelope(t, "fixtures/invalid-property-key.ovf")
		_, err := e.ToConfigSpec()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid character")
		assert.Contains(t, err.Error(), "invalid.key")
	})

	t.Run("uber.ovf: comprehensive OVF parsing and ConfigSpec", func(t *testing.T) {
		e := testEnvelope(t, "fixtures/uber.ovf")

		findProp := func(va *types.VAppConfigSpec, id string) *types.VAppPropertyInfo {
			if va == nil {
				return nil
			}
			for i := range va.Property {
				if va.Property[i].Info.Id == id {
					return va.Property[i].Info
				}
			}
			return nil
		}

		t.Run("envelope parsing", func(t *testing.T) {
			require.NotNil(t, e.VirtualSystem)
			require.Nil(t, e.VirtualSystemCollection)
			require.NotEmpty(t, e.References, "References with File for file-backed disk")
			require.NotNil(t, e.Disk)
			require.Len(t, e.Disk.Disks, 2, "vmdisk1 file-backed, vmdisk2 property capacity")
			require.NotNil(t, e.Network)
			require.Len(t, e.Network.Networks, 2)
			require.NotNil(t, e.DeploymentOption)
			require.Len(t, e.DeploymentOption.Configuration, 2)
			require.NotEmpty(t, e.VirtualSystem.Product)
			ps := e.VirtualSystem.Product[0]
			require.NotEmpty(t, ps.Items, "category grouping: Items preserved")
			require.NotEmpty(t, ps.Property)
			require.NotNil(t, e.VirtualSystem.Annotation)
			require.NotNil(t, e.VirtualSystem.OperatingSystem)
			require.Len(t, e.VirtualSystem.VirtualHardware, 1)
			hw := e.VirtualSystem.VirtualHardware[0]
			require.NotEmpty(t, hw.Item)
			require.Equal(t, "uber-vm", e.VirtualSystem.ID)
		})

		t.Run("default config: ConfigSpec and VAppConfig", func(t *testing.T) {
			cs, err := e.ToConfigSpec()
			require.NoError(t, err)
			assert.Equal(t, "uber-vm", cs.Name)
			assert.NotEmpty(t, cs.GuestId)
			assert.Equal(t, int32(1), cs.NumCPUs)
			assert.Equal(t, int64(512), cs.MemoryMB)
			require.NotEmpty(t, cs.DeviceChange)

			var disks []*types.VirtualDisk
			for _, dc := range cs.DeviceChange {
				if d, ok := dc.GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); ok {
					disks = append(disks, d)
				}
			}
			require.Len(t, disks, 1, "default config: only vmdisk1 (second disk is extended-only)")
			db, ok := disks[0].Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			require.True(t, ok)
			assert.Equal(t, "uber-os.vmdk", db.FileName)
			assert.Equal(t, int64(20*1024*1024*1024), disks[0].CapacityInBytes)

			va, ok := cs.VAppConfig.(*types.VAppConfigSpec)
			require.True(t, ok)
			// default: net_common, hostname, storage_common, disksize, auth_mode, enable_ssh (no extended-only)
			require.Len(t, va.Property, 6)
			assert.Equal(t, "Net", findProp(va, "net_common").Category)
			assert.Equal(t, "Net", findProp(va, "hostname").Category)
			assert.Equal(t, "Storage", findProp(va, "storage_common").Category)
			assert.Equal(t, "Storage", findProp(va, "disksize").Category)
			assert.Equal(t, "Security", findProp(va, "auth_mode").Category)
			assert.Equal(t, "Security", findProp(va, "enable_ssh").Category)
			assert.Nil(t, findProp(va, "net_extended"))
			assert.Nil(t, findProp(va, "storage_extended"))
			assert.Nil(t, findProp(va, "extended_secret"))
		})

		t.Run("extended config: all properties, categories, and two disks", func(t *testing.T) {
			cs, err := e.ToConfigSpecWithOptions(ToConfigSpecOptions{
				DeploymentConfiguration: "extended",
			})
			require.NoError(t, err)
			va, ok := cs.VAppConfig.(*types.VAppConfigSpec)
			require.True(t, ok)
			require.Len(t, va.Property, 9)
			assert.Equal(t, "Net", findProp(va, "net_extended").Category)
			assert.Equal(t, "Storage", findProp(va, "storage_extended").Category)
			assert.Equal(t, "Security", findProp(va, "extended_secret").Category)

			var disks []*types.VirtualDisk
			for _, dc := range cs.DeviceChange {
				if d, ok := dc.GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); ok {
					disks = append(disks, d)
				}
			}
			require.Len(t, disks, 2, "extended config: vmdisk1 + vmdisk2 (second disk is extended-only)")
		})

		t.Run("disks default config: one file-backed disk", func(t *testing.T) {
			cs, err := e.ToConfigSpec()
			require.NoError(t, err)
			var disks []*types.VirtualDisk
			for _, dc := range cs.DeviceChange {
				if d, ok := dc.GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); ok {
					disks = append(disks, d)
				}
			}
			require.Len(t, disks, 1)
			db, ok := disks[0].Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			require.True(t, ok)
			assert.NotEmpty(t, db.FileName, "vmdisk1 is file-backed")
			assert.Equal(t, "uber-os.vmdk", db.FileName)
			assert.Equal(t, int64(20*1024*1024*1024), disks[0].CapacityInBytes)
		})

		t.Run("disks extended config: two disks (file-backed + property-backed capacity)", func(t *testing.T) {
			cs, err := e.ToConfigSpecWithOptions(ToConfigSpecOptions{
				DeploymentConfiguration: "extended",
			})
			require.NoError(t, err)
			var disks []*types.VirtualDisk
			for _, dc := range cs.DeviceChange {
				if d, ok := dc.GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); ok {
					disks = append(disks, d)
				}
			}
			require.Len(t, disks, 2)

			db1, ok := disks[0].Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			require.True(t, ok)
			assert.NotEmpty(t, db1.FileName, "vmdisk1 is file-backed")
			assert.Equal(t, "uber-os.vmdk", db1.FileName)
			assert.Equal(t, int64(20*1024*1024*1024), disks[0].CapacityInBytes)

			db2, ok := disks[1].Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			require.True(t, ok)
			assert.Empty(t, db2.FileName, "vmdisk2 is empty disk with property capacity")
			assert.Equal(t, int64(5*1024*1024*1024), disks[1].CapacityInBytes,
				"capacity from property disksize=5 GiB (DSP0243 9.1)")
		})

		t.Run("property types and qualifiers", func(t *testing.T) {
			cs, err := e.ToConfigSpec()
			require.NoError(t, err)
			va, ok := cs.VAppConfig.(*types.VAppConfigSpec)
			require.True(t, ok)
			p := findProp(va, "enable_ssh")
			require.NotNil(t, p)
			assert.Equal(t, "False", p.DefaultValue, "boolean normalised to canonical form")
			p = findProp(va, "hostname")
			require.NotNil(t, p)
			assert.Equal(t, "string(1..253)", p.Type, "MinLen(1),MaxLen(253) -> string(1..253)")
		})

		t.Run("invalid deployment config returns error", func(t *testing.T) {
			_, err := e.ToConfigSpecWithOptions(ToConfigSpecOptions{
				DeploymentConfiguration: "nonexistent",
			})
			require.Error(t, err)
			assert.Contains(t, err.Error(), "not found in DeploymentOptionSection")
		})
	})

	t.Run("DeploymentConfiguration", func(t *testing.T) {
		e := testEnvelope(t, "fixtures/configspec.ovf")

		t.Run("explicit frontend includes frontend-only elements", func(t *testing.T) {
			cs, err := e.ToConfigSpecWithOptions(ToConfigSpecOptions{
				DeploymentConfiguration: "frontend",
			})
			require.NoError(t, err)
			va, ok := cs.VAppConfig.(*types.VAppConfigSpec)
			require.True(t, ok)
			require.Len(t, va.Property, 8, "frontend config has 8 vApp properties")
			ids := make([]string, len(va.Property))
			for i := range va.Property {
				ids[i] = va.Property[i].Info.Id
			}
			assert.Contains(t, ids, "frontend_ip")
			assert.Contains(t, ids, "frontend_gateway")
			assert.Len(t, cs.DeviceChange, 22, "frontend config has 22 devices (extra NIC)")
		})

		t.Run("invalid name returns error", func(t *testing.T) {
			_, err := e.ToConfigSpecWithOptions(ToConfigSpecOptions{
				DeploymentConfiguration: "nonexistent",
			})
			require.Error(t, err)
			assert.Contains(t, err.Error(), "not found in DeploymentOptionSection")
		})
	})

	t.Run("DeploymentConfigs fixture (small/medium/large)", func(t *testing.T) {
		e := testEnvelope(t, "fixtures/deployment-configs.ovf")

		findProp := func(va *types.VAppConfigSpec, id string) *types.VAppPropertyInfo {
			if va == nil {
				return nil
			}
			for i := range va.Property {
				if va.Property[i].Info.Id == id {
					return va.Property[i].Info
				}
			}
			return nil
		}

		t.Run("small config: 2 CPU, 2 GiB RAM, 1 disk 1 GiB file-backed", func(t *testing.T) {
			cs, err := e.ToConfigSpecWithOptions(ToConfigSpecOptions{
				DeploymentConfiguration: "small",
			})
			require.NoError(t, err)
			assert.Equal(t, int32(2), cs.NumCPUs)
			assert.Equal(t, int64(2048), cs.MemoryMB)
			var disks []*types.VirtualDisk
			for _, dc := range cs.DeviceChange {
				if d, ok := dc.GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); ok {
					disks = append(disks, d)
				}
			}
			require.Len(t, disks, 1)
			assert.Equal(t, int64(1*1024*1024*1024), disks[0].CapacityInBytes)
			db, ok := disks[0].Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			require.True(t, ok)
			assert.Equal(t, "system-and-model-small.vmdk", db.FileName)

			va, _ := cs.VAppConfig.(*types.VAppConfigSpec)
			require.NotNil(t, va)
			assert.NotNil(t, findProp(va, "hostname"))
			assert.NotNil(t, findProp(va, "tier_small"))
			assert.NotNil(t, findProp(va, "size_preset"))
			assert.Nil(t, findProp(va, "tier_medium"))
			assert.Nil(t, findProp(va, "tier_large"))
		})

		t.Run("medium config (default): 4 CPU, 4 GiB RAM, 2 disks 2 GiB", func(t *testing.T) {
			cs, err := e.ToConfigSpec()
			require.NoError(t, err)
			assert.Equal(t, int32(4), cs.NumCPUs)
			assert.Equal(t, int64(4096), cs.MemoryMB)
			var disks []*types.VirtualDisk
			for _, dc := range cs.DeviceChange {
				if d, ok := dc.GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); ok {
					disks = append(disks, d)
				}
			}
			require.Len(t, disks, 2)
			assert.Equal(t, int64(2*1024*1024*1024), disks[0].CapacityInBytes)
			assert.Equal(t, int64(2*1024*1024*1024), disks[1].CapacityInBytes)
			db0, ok := disks[0].Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			require.True(t, ok)
			assert.Equal(t, "system-and-model-medium.vmdk", db0.FileName)
			db1, ok := disks[1].Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			require.True(t, ok)
			assert.Empty(t, db1.FileName)

			va, _ := cs.VAppConfig.(*types.VAppConfigSpec)
			require.NotNil(t, va)
			assert.NotNil(t, findProp(va, "tier_medium"))
			assert.Nil(t, findProp(va, "tier_small"))
			assert.Nil(t, findProp(va, "tier_large"))
		})

		t.Run("large config: 8 CPU, 8 GiB RAM, 3 disks 4 GiB", func(t *testing.T) {
			cs, err := e.ToConfigSpecWithOptions(ToConfigSpecOptions{
				DeploymentConfiguration: "large",
			})
			require.NoError(t, err)
			assert.Equal(t, int32(8), cs.NumCPUs)
			assert.Equal(t, int64(8192), cs.MemoryMB)
			var disks []*types.VirtualDisk
			for _, dc := range cs.DeviceChange {
				if d, ok := dc.GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); ok {
					disks = append(disks, d)
				}
			}
			require.Len(t, disks, 3)
			for i := range disks {
				assert.Equal(t, int64(4*1024*1024*1024), disks[i].CapacityInBytes)
			}
			db0, _ := disks[0].Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			db1, _ := disks[1].Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			db2, _ := disks[2].Backing.(*types.VirtualDiskFlatVer2BackingInfo)
			assert.Equal(t, "system-large.vmdk", db0.FileName)
			assert.Equal(t, "model-large.vmdk", db1.FileName)
			assert.Empty(t, db2.FileName)

			va, _ := cs.VAppConfig.(*types.VAppConfigSpec)
			require.NotNil(t, va)
			assert.NotNil(t, findProp(va, "tier_large"))
			assert.Nil(t, findProp(va, "tier_small"))
			assert.Nil(t, findProp(va, "tier_medium"))
		})

		t.Run("property qualifiers embedded in VAppPropertyInfo.Type", func(t *testing.T) {
			cs, err := e.ToConfigSpec()
			require.NoError(t, err)
			va, ok := cs.VAppConfig.(*types.VAppConfigSpec)
			require.True(t, ok)

			p := findProp(va, "hostname")
			require.NotNil(t, p)
			assert.Equal(t, "string(1..253)", p.Type,
				"hostname has MinLen(1),MaxLen(253) => string(1..253)")

			p = findProp(va, "size_preset")
			require.NotNil(t, p)
			assert.Contains(t, p.Type, "string[",
				"size_preset has ValueMap => string[\"small\", ...]")
			assert.Contains(t, p.Type, "small")
			assert.Contains(t, p.Type, "medium")
			assert.Contains(t, p.Type, "large")

			p = findProp(va, "tier_medium")
			require.NotNil(t, p)
			assert.Equal(t, "string(1..64)", p.Type,
				"tier_medium has MinLen(1),MaxLen(64) => string(1..64)")
		})

		t.Run("empty DeploymentConfiguration uses default (medium)", func(t *testing.T) {
			cs, err := e.ToConfigSpecWithOptions(ToConfigSpecOptions{})
			require.NoError(t, err)
			assert.Equal(t, int32(4), cs.NumCPUs)
			assert.Equal(t, int64(4096), cs.MemoryMB)
		})
	})

	t.Run("Large", func(t *testing.T) {
		e := testEnvelope(t, "fixtures/configspec.ovf")
		cs, err := e.ToConfigSpec()
		assert.NoError(t, err)
		assert.NotEmpty(t, cs)

		if testing.Verbose() {
			var w bytes.Buffer
			enc := types.NewJSONEncoder(&w)
			enc.SetIndent("", "    ")
			assert.NoError(t, enc.Encode(cs))
			t.Logf("\n\nconfigSpec=%s\n\n", w.String())
		}

		assert.Equal(t, "haproxy", cs.Name)
		assert.Equal(t, int32(2), cs.NumCPUs)
		assert.Equal(t, int32(2), *cs.NumCoresPerSocket)
		assert.Equal(t, int64(4096), cs.MemoryMB)
		assert.Equal(t, "vmx-13", cs.Version)

		if assert.Len(t, cs.DeviceChange, 21) {

			var scsiController1Key int32
			if d, ok := cs.DeviceChange[0].GetVirtualDeviceConfigSpec().Device.(*types.VirtualLsiLogicController); assert.True(t, ok) {
				scsiController1Key = d.Key
				assert.Equal(t, int32(0), d.BusNumber)
				if assert.NotNil(t, d.SlotInfo) {
					si, ok := d.SlotInfo.(*types.VirtualDevicePciBusSlotInfo)
					assert.True(t, ok)
					assert.Equal(t, int32(128), si.PciSlotNumber)
				}
			}

			if d, ok := cs.DeviceChange[1].GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); assert.True(t, ok) {
				assert.Equal(t, scsiController1Key, d.ControllerKey)
				if assert.NotNil(t, d.UnitNumber) {
					assert.Equal(t, int32(0), *d.UnitNumber)
				}
				db, ok := d.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
				assert.True(t, ok)
				assert.Equal(t, string(types.VirtualDiskModePersistent), db.DiskMode)
				if assert.NotNil(t, db.ThinProvisioned) {
					assert.True(t, *db.ThinProvisioned)
				}
				assert.Equal(t, int64(20*1024*1024*1024), d.CapacityInBytes)
				// Disk is file-backed (vmdisk1 has ovf:fileRef in configspec.ovf); backing must have non-empty FileName.
				assert.NotEmpty(t, db.FileName, "file-backed disk should have non-empty FileName")
			}

			if bd, ok := cs.DeviceChange[2].GetVirtualDeviceConfigSpec().Device.(types.BaseVirtualEthernetCard); assert.True(t, ok) {
				d := bd.GetVirtualEthernetCard()
				if assert.NotNil(t, d.UnitNumber) {
					assert.Equal(t, int32(2), *d.UnitNumber)
				}
				if assert.NotNil(t, d.Connectable) {
					assert.True(t, d.Connectable.AllowGuestControl)
				}
				if assert.NotNil(t, d.WakeOnLanEnabled) {
					assert.False(t, *d.WakeOnLanEnabled)
				}
				if assert.NotNil(t, d.SlotInfo) {
					si, ok := d.SlotInfo.(*types.VirtualDevicePciBusSlotInfo)
					assert.True(t, ok)
					assert.Equal(t, int32(160), si.PciSlotNumber)
				}
			}

			if bd, ok := cs.DeviceChange[3].GetVirtualDeviceConfigSpec().Device.(types.BaseVirtualEthernetCard); assert.True(t, ok) {
				d := bd.GetVirtualEthernetCard()
				if assert.NotNil(t, d.UnitNumber) {
					assert.Equal(t, int32(3), *d.UnitNumber)
				}
				if assert.NotNil(t, d.Connectable) {
					assert.True(t, d.Connectable.AllowGuestControl)
				}
				if assert.NotNil(t, d.WakeOnLanEnabled) {
					assert.False(t, *d.WakeOnLanEnabled)
				}
				if assert.NotNil(t, d.SlotInfo) {
					si, ok := d.SlotInfo.(*types.VirtualDevicePciBusSlotInfo)
					assert.True(t, ok)
					assert.Equal(t, int32(192), si.PciSlotNumber)
				}
				if assert.NotNil(t, d.UptCompatibilityEnabled) {
					assert.False(t, *d.UptCompatibilityEnabled)
				}
			}

			if d, ok := cs.DeviceChange[4].GetVirtualDeviceConfigSpec().Device.(*types.VirtualMachineVideoCard); assert.True(t, ok) {
				if assert.NotNil(t, d.Enable3DSupport) {
					assert.False(t, *d.Enable3DSupport)
				}
				assert.Equal(t, int64(262144), d.GraphicsMemorySizeInKB)
				if assert.NotNil(t, d.UseAutoDetect) {
					assert.False(t, *d.UseAutoDetect)
				}
				assert.Equal(t, int64(4096), d.VideoRamSizeInKB)
				assert.Equal(t, int32(1), d.NumDisplays)
				assert.Equal(t, "automatic", d.Use3dRenderer)
			}

			var ideControllerKey int32
			if d, ok := cs.DeviceChange[5].GetVirtualDeviceConfigSpec().Device.(*types.VirtualIDEController); assert.True(t, ok) {
				ideControllerKey = d.Key
				assert.Equal(t, int32(1), d.BusNumber)
			}

			if d, ok := cs.DeviceChange[6].GetVirtualDeviceConfigSpec().Device.(*types.VirtualIDEController); assert.True(t, ok) {
				assert.Equal(t, int32(0), d.BusNumber)
			}

			if d, ok := cs.DeviceChange[7].GetVirtualDeviceConfigSpec().Device.(*types.VirtualCdrom); assert.True(t, ok) {
				if assert.NotNil(t, d.UnitNumber) {
					assert.Equal(t, int32(0), *d.UnitNumber)
				}
				assert.Equal(t, ideControllerKey, d.ControllerKey)
			}

			var sioControllerKey int32
			if d, ok := cs.DeviceChange[8].GetVirtualDeviceConfigSpec().Device.(*types.VirtualSIOController); assert.True(t, ok) {
				sioControllerKey = d.Key
			}

			if d, ok := cs.DeviceChange[9].GetVirtualDeviceConfigSpec().Device.(*types.VirtualFloppy); assert.True(t, ok) {
				if assert.NotNil(t, d.UnitNumber) {
					assert.Equal(t, int32(0), *d.UnitNumber)
				}
				assert.Equal(t, sioControllerKey, d.ControllerKey)
			}

			if d, ok := cs.DeviceChange[10].GetVirtualDeviceConfigSpec().Device.(*types.VirtualMachineVMCIDevice); assert.True(t, ok) {
				if assert.NotNil(t, d.AllowUnrestrictedCommunication) {
					assert.False(t, *d.AllowUnrestrictedCommunication)
				}
			}

			var scsiController2Key int32
			if d, ok := cs.DeviceChange[11].GetVirtualDeviceConfigSpec().Device.(*types.ParaVirtualSCSIController); assert.True(t, ok) {
				scsiController2Key = d.Key
				assert.Equal(t, int32(1), d.BusNumber)
				assert.Nil(t, d.SlotInfo)
				assert.Len(t, d.Device, 1)
			}

			if d, ok := cs.DeviceChange[12].GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); assert.True(t, ok) {
				assert.Equal(t, scsiController2Key, d.ControllerKey)
				if assert.NotNil(t, d.UnitNumber) {
					assert.Equal(t, int32(0), *d.UnitNumber)
				}
				db, ok := d.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
				assert.True(t, ok)
				assert.Equal(t, string(types.VirtualDiskModePersistent), db.DiskMode)
				if assert.NotNil(t, db.ThinProvisioned) {
					assert.True(t, *db.ThinProvisioned)
				}
				assert.Equal(t, int64(10*1024*1024*1024), d.CapacityInBytes)
				// Disk has no HostResource (capacity-only item); created as empty disk, so FileName is empty.
				assert.Empty(t, db.FileName, "empty disk should have empty FileName")
			}

			var sataController1Key int32
			if d, ok := cs.DeviceChange[13].GetVirtualDeviceConfigSpec().Device.(*types.VirtualAHCIController); assert.True(t, ok) {
				sataController1Key = d.Key
				assert.Equal(t, int32(0), d.BusNumber)
				assert.Nil(t, d.SlotInfo)
			}

			if d, ok := cs.DeviceChange[14].GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); assert.True(t, ok) {
				assert.Equal(t, sataController1Key, d.ControllerKey)
				if assert.NotNil(t, d.UnitNumber) {
					assert.Equal(t, int32(0), *d.UnitNumber)
				}
				db, ok := d.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
				assert.True(t, ok)
				assert.Equal(t, string(types.VirtualDiskModePersistent), db.DiskMode)
				if assert.NotNil(t, db.ThinProvisioned) {
					assert.True(t, *db.ThinProvisioned)
				}
				assert.Equal(t, int64(10*1024*1024*1024), d.CapacityInBytes)
			}

			var sataController2Key int32
			if d, ok := cs.DeviceChange[15].GetVirtualDeviceConfigSpec().Device.(*types.VirtualAHCIController); assert.True(t, ok) {
				sataController2Key = d.Key
				assert.Equal(t, int32(1), d.BusNumber)
				assert.Nil(t, d.SlotInfo)
			}

			if d, ok := cs.DeviceChange[16].GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); assert.True(t, ok) {
				assert.Equal(t, sataController2Key, d.ControllerKey)
				if assert.NotNil(t, d.UnitNumber) {
					assert.Equal(t, int32(0), *d.UnitNumber)
				}
				db, ok := d.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
				assert.True(t, ok)
				assert.Equal(t, string(types.VirtualDiskModePersistent), db.DiskMode)
				if assert.NotNil(t, db.ThinProvisioned) {
					assert.True(t, *db.ThinProvisioned)
				}
				assert.Equal(t, int64(10*1024*1024*1024), d.CapacityInBytes)
			}

			var nvmeController1Key int32
			if d, ok := cs.DeviceChange[17].GetVirtualDeviceConfigSpec().Device.(*types.VirtualNVMEController); assert.True(t, ok) {
				nvmeController1Key = d.Key
				assert.Equal(t, int32(0), d.BusNumber)
				assert.Nil(t, d.SlotInfo)
			}

			if d, ok := cs.DeviceChange[18].GetVirtualDeviceConfigSpec().Device.(*types.VirtualDisk); assert.True(t, ok) {
				assert.Equal(t, nvmeController1Key, d.ControllerKey)
				if assert.NotNil(t, d.UnitNumber) {
					assert.Equal(t, int32(0), *d.UnitNumber)
				}
				db, ok := d.Backing.(*types.VirtualDiskFlatVer2BackingInfo)
				assert.True(t, ok)
				assert.Equal(t, string(types.VirtualDiskModePersistent), db.DiskMode)
				if assert.NotNil(t, db.ThinProvisioned) {
					assert.True(t, *db.ThinProvisioned)
				}
				assert.Equal(t, int64(10*1024*1024*1024), d.CapacityInBytes)
			}

			if d, ok := cs.DeviceChange[19].GetVirtualDeviceConfigSpec().Device.(*types.VirtualUSBController); assert.True(t, ok) {
				if assert.NotNil(t, d.AutoConnectDevices) {
					assert.True(t, *d.AutoConnectDevices)
				}
				if assert.NotNil(t, d.EhciEnabled) {
					assert.False(t, *d.EhciEnabled)
				}
			}

			if d, ok := cs.DeviceChange[20].GetVirtualDeviceConfigSpec().Device.(*types.VirtualUSBXHCIController); assert.True(t, ok) {
				if assert.NotNil(t, d.AutoConnectDevices) {
					assert.True(t, *d.AutoConnectDevices)
				}
			}
		}

		assert.ElementsMatch(t, cs.ExtraConfig, []types.BaseOptionValue{
			&types.OptionValue{
				Key:   "guest_rpc.auth.cloud-init.set",
				Value: "FALSE",
			},
		})

		if assert.NotNil(t, cs.Flags) {
			if assert.NotNil(t, cs.Flags.VbsEnabled) {
				assert.False(t, *cs.Flags.VbsEnabled)
			}
			if assert.NotNil(t, cs.Flags.VvtdEnabled) {
				assert.False(t, *cs.Flags.VvtdEnabled)
			}
		}

		assert.Equal(t, "bios", cs.Firmware)

		if assert.NotNil(t, cs.BootOptions) {
			if assert.NotNil(t, cs.BootOptions.EfiSecureBootEnabled) {
				assert.False(t, *cs.BootOptions.EfiSecureBootEnabled)
			}
		}

		if assert.NotNil(t, cs.CpuHotAddEnabled) {
			assert.False(t, *cs.CpuHotAddEnabled)
		}
		if assert.NotNil(t, cs.CpuHotRemoveEnabled) {
			assert.False(t, *cs.CpuHotRemoveEnabled)
		}
		if assert.NotNil(t, cs.MemoryHotAddEnabled) {
			assert.False(t, *cs.MemoryHotAddEnabled)
		}
		if assert.NotNil(t, cs.NestedHVEnabled) {
			assert.False(t, *cs.NestedHVEnabled)
		}
		if assert.NotNil(t, cs.VirtualICH7MPresent) {
			assert.False(t, *cs.VirtualICH7MPresent)
		}

		assert.Equal(t, int32(1), cs.SimultaneousThreads)

		if assert.NotNil(t, cs.VPMCEnabled) {
			assert.False(t, *cs.VPMCEnabled)
		}

		if assert.NotNil(t, cs.CpuAllocation) {
			if assert.NotNil(t, cs.CpuAllocation.Shares) {
				assert.Equal(t, &types.SharesInfo{
					Shares: 2000,
					Level:  types.SharesLevelNormal,
				}, cs.CpuAllocation.Shares)
			}
		}

		if assert.NotNil(t, cs.PowerOpInfo) {
			assert.Equal(t, &types.VirtualMachineDefaultPowerOpInfo{
				PowerOffType:  "soft",
				ResetType:     "soft",
				SuspendType:   "hard",
				StandbyAction: "checkpoint",
			}, cs.PowerOpInfo)
		}

		if assert.NotNil(t, cs.Tools) {
			assert.Equal(t, &types.ToolsConfigInfo{
				SyncTimeWithHost:        types.NewBool(false),
				SyncTimeWithHostAllowed: types.NewBool(true),
				AfterPowerOn:            types.NewBool(true),
				AfterResume:             types.NewBool(true),
				BeforeGuestShutdown:     types.NewBool(true),
				BeforeGuestStandby:      types.NewBool(true),
				ToolsUpgradePolicy:      "manual",
			}, cs.Tools)
		}

		if va, ok := cs.VAppConfig.(*types.VAppConfigSpec); assert.True(t, ok) {
			if assert.Len(t, va.Product, 4) {
				assert.ElementsMatch(t,
					[]types.VAppProductSpec{
						{
							ArrayUpdateSpec: types.ArrayUpdateSpec{
								Operation: types.ArrayUpdateOperationAdd,
							},
							Info: &types.VAppProductInfo{
								Key:         0,
								Name:        "HAProxy for the Load Balancer API v0.2.0",
								Vendor:      "VMware Inc.",
								Version:     "v0.2.0",
								FullVersion: "v0.2.0",
								ProductUrl:  "https://vmware.com",
								VendorUrl:   "https://vmware.com",
							},
						},
						{
							ArrayUpdateSpec: types.ArrayUpdateSpec{
								Operation: types.ArrayUpdateOperationAdd,
							},
							Info: &types.VAppProductInfo{
								Key:     1,
								ClassId: "appliance",
							},
						},
						{
							ArrayUpdateSpec: types.ArrayUpdateSpec{
								Operation: types.ArrayUpdateOperationAdd,
							},
							Info: &types.VAppProductInfo{
								Key:     2,
								ClassId: "network",
							},
						},
						{
							ArrayUpdateSpec: types.ArrayUpdateSpec{
								Operation: types.ArrayUpdateOperationAdd,
							},
							Info: &types.VAppProductInfo{
								Key:     3,
								ClassId: "loadbalance",
							},
						},
					},
					va.Product,
				)
			}
			if assert.Len(t, va.Property, 6) {
				assert.ElementsMatch(t,
					[]types.VAppPropertySpec{

						// default
						{
							ArrayUpdateSpec: types.ArrayUpdateSpec{
								Operation: types.ArrayUpdateOperationAdd,
							},
							Info: &types.VAppPropertyInfo{
								Key:              0,
								Category:         "Load Balancer API",
								Id:               "BUILD_TIMESTAMP",
								Type:             "string",
								UserConfigurable: types.NewBool(false),
								DefaultValue:     "1615488399",
							},
						},
						{
							ArrayUpdateSpec: types.ArrayUpdateSpec{
								Operation: types.ArrayUpdateOperationAdd,
							},
							Info: &types.VAppPropertyInfo{
								Key:              1,
								Category:         "Load Balancer API",
								Id:               "BUILD_DATE",
								Type:             "string",
								UserConfigurable: types.NewBool(false),
								DefaultValue:     "2021-03-11T18:46:39Z",
							},
						},

						// appliance (ovf:password="true" => Type "password")
						{
							ArrayUpdateSpec: types.ArrayUpdateSpec{
								Operation: types.ArrayUpdateOperationAdd,
							},
							Info: &types.VAppPropertyInfo{
								Key:              2,
								Category:         "1. Appliance Configuration",
								ClassId:          "appliance",
								Id:               "root_pwd",
								Label:            "1.1. Root Password",
								Type:             "password",
								UserConfigurable: types.NewBool(true),
								Description:      "The initial password of the root user. Subsequent changes of password should be performed in operating system. (6-128 characters)",
							},
						},

						// network
						{
							ArrayUpdateSpec: types.ArrayUpdateSpec{
								Operation: types.ArrayUpdateOperationAdd,
							},
							Info: &types.VAppPropertyInfo{
								Key:              3,
								Category:         "2. Network Config",
								ClassId:          "network",
								Id:               "hostname",
								Label:            "2.1. Host Name",
								Type:             "string",
								UserConfigurable: types.NewBool(true),
								DefaultValue:     "haproxy.local",
								Description:      "The host name. A fully-qualified domain name is also supported.",
							},
						},

						// loadbalance
						{
							ArrayUpdateSpec: types.ArrayUpdateSpec{
								Operation: types.ArrayUpdateOperationAdd,
							},
							Info: &types.VAppPropertyInfo{
								Key:              4,
								Category:         "3. Load Balancing",
								ClassId:          "loadbalance",
								Id:               "service_ip_range",
								Label:            "3.1. Load Balancer IP Ranges, comma-separated in CIDR format (Eg 1.2.3.4/28,5.6.7.8/28)",
								Type:             "string",
								UserConfigurable: types.NewBool(true),
								Description:      "The IP ranges the load balancer will use for Kubernetes Services and Control Planes. The Appliance will currently respond to ALL the IPs in these ranges whether they're assigned or not. As such, these ranges must not overlap with the IPs assigned for the appliance or any other VMs on the network.",
							},
						},
						{
							ArrayUpdateSpec: types.ArrayUpdateSpec{
								Operation: types.ArrayUpdateOperationAdd,
							},
							Info: &types.VAppPropertyInfo{
								Key:              5,
								Category:         "3. Load Balancing",
								ClassId:          "loadbalance",
								Id:               "dataplane_port",
								Label:            "3.2. Dataplane API Management Port",
								Type:             "int",
								UserConfigurable: types.NewBool(true),
								DefaultValue:     "5556",
								Description:      "Specifies the port on which the Dataplane API will be advertized on the Management Network.",
							},
						},
					},
					va.Property,
				)
			}
		}
	})
}
