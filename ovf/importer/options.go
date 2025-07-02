// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package importer

import (
	"encoding/json"

	"github.com/vmware/govmomi/ovf"
	"github.com/vmware/govmomi/vim25/types"
)

type KeyValue struct {
	Key   string
	Value string
}

// case insensitive for Key + Value
func (kv *KeyValue) UnmarshalJSON(b []byte) error {
	e := struct {
		types.KeyValue
		Key   *string
		Value *string
	}{
		types.KeyValue{}, &kv.Key, &kv.Value,
	}

	err := json.Unmarshal(b, &e)
	if err != nil {
		return err
	}

	if kv.Key == "" {
		kv.Key = e.KeyValue.Key // "key"
	}

	if kv.Value == "" {
		kv.Value = e.KeyValue.Value // "value"
	}

	return nil
}

type Property struct {
	KeyValue
	Spec *ovf.Property `json:",omitempty"`
}

type Network struct {
	Name    string
	Network string
}

type Options struct {
	AllDeploymentOptions []string `json:",omitempty"`
	Deployment           string   `json:",omitempty"`

	AllDiskProvisioningOptions []string `json:",omitempty"`
	DiskProvisioning           string

	AllIPAllocationPolicyOptions []string `json:",omitempty"`
	IPAllocationPolicy           string

	AllIPProtocolOptions []string `json:",omitempty"`
	IPProtocol           string

	PropertyMapping []Property `json:",omitempty"`

	NetworkMapping []Network `json:",omitempty"`

	Annotation string `json:",omitempty"`

	MarkAsTemplate bool
	PowerOn        bool
	InjectOvfEnv   bool
	WaitForIP      bool
	Name           *string
}
