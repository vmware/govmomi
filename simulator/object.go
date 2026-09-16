// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"bytes"
	"strings"

	"github.com/vmware/govmomi/internal"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vim25/xml"
)

func SetCustomValue(ctx *Context, req *types.SetCustomValue) soap.HasFault {
	body := &methods.SetCustomValueBody{}

	cfm := ctx.Map.CustomFieldsManager()

	_, field := cfm.findByNameType(req.Key, req.This.Type)
	if field == nil {
		body.Fault_ = Fault("", &types.InvalidArgument{InvalidProperty: "key"})
		return body
	}

	res := cfm.SetField(ctx, &types.SetField{
		This:   cfm.Reference(),
		Entity: req.This,
		Key:    field.Key,
		Value:  req.Value,
	})

	if res.Fault() != nil {
		body.Fault_ = res.Fault()
		return body
	}

	body.Res = &types.SetCustomValueResponse{}
	return body
}

func GetCustomFieldsAvailable(ctx *Context, ref *types.ManagedObjectReference) []types.CustomFieldDef {
	// CustomFieldsManager is not available in ESX
	if ctx.Map.IsESX() {
		return nil
	}

	cfm := ctx.Map.CustomFieldsManager()
	var filtered []types.CustomFieldDef
	for _, f := range cfm.Field {
		if f.ManagedObjectType == "" || f.ManagedObjectType == ref.Type {
			filtered = append(filtered, f)
		}
	}
	return filtered
}

// newUUID returns a stable UUID string based on input s
func newUUID(s string) string {
	return internal.OID(s).String()
}

// newDVSUuid returns a stable DVS UUID based on input s, formatted the way real
// vCenter/ESXi format it on the wire: 16 space-separated hex byte pairs with a
// dash between the 8th and 9th pair (e.g. "50 13 a2 63 0a a6 77 65-37 e2 20 e6 2b 8f a2 f6").
// This package's usual newUUID format (a plain dashed UUID) is never what a
// real vCenter emits for a DVS UUID -- any client that validates or parses
// the field against that specific shape (e.g. matching HostProxySwitch.DvsUuid
// against DVSSummary.Uuid) would only ever see the wrong shape from vcsim.
func newDVSUuid(s string) string {
	hex := strings.ReplaceAll(newUUID(s), "-", "")
	var b strings.Builder
	for i := 0; i < len(hex); i += 2 {
		if i > 0 {
			if i == 16 {
				b.WriteByte('-')
			} else {
				b.WriteByte(' ')
			}
		}
		b.WriteString(hex[i : i+2])
	}
	return b.String()
}

// deepCopy uses xml encode/decode to copy src to dst
func deepCopy(src, dst any) {
	b, err := xml.Marshal(src)
	if err != nil {
		panic(err)
	}

	dec := xml.NewDecoder(bytes.NewReader(b))
	dec.TypeFunc = types.TypeFunc()
	err = dec.Decode(dst)
	if err != nil {
		panic(err)
	}
}
