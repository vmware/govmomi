// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"reflect"

	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

// Minimal set of internal types and methods:
// - Fetch() - used by ovftool to collect various managed object properties
// - RetrieveInternalContent() - used by ovftool to obtain a reference to NfcService (which it does not use by default)

func init() {
	types.Add("Fetch", reflect.TypeOf((*Fetch)(nil)).Elem())
}

type Fetch struct {
	This types.ManagedObjectReference `xml:"_this"`
	Prop string                       `xml:"prop"`
}

type FetchResponse struct {
	Returnval types.AnyType `xml:"returnval,omitempty,typeattr"`
}

type FetchBody struct {
	Req    *Fetch         `xml:"Fetch,omitempty"`
	Res    *FetchResponse `xml:"FetchResponse,omitempty"`
	Fault_ *soap.Fault    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *FetchBody) Fault() *soap.Fault { return b.Fault_ }

func init() {
	types.Add("RetrieveInternalContent", reflect.TypeOf((*RetrieveInternalContent)(nil)).Elem())
}

type RetrieveInternalContent struct {
	This types.ManagedObjectReference `xml:"_this"`
}

type RetrieveInternalContentResponse struct {
	Returnval InternalServiceInstanceContent `xml:"returnval"`
}

type RetrieveInternalContentBody struct {
	Res    *RetrieveInternalContentResponse `xml:"RetrieveInternalContentResponse,omitempty"`
	Fault_ *soap.Fault                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

func (b *RetrieveInternalContentBody) Fault() *soap.Fault { return b.Fault_ }

type InternalServiceInstanceContent struct {
	types.DynamicData

	NfcService types.ManagedObjectReference `xml:"nfcService"`
}
