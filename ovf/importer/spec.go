// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package importer

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/vmware/govmomi/ovf"
	"github.com/vmware/govmomi/vim25/types"
)

var (
	allDiskProvisioningOptions = types.OvfCreateImportSpecParamsDiskProvisioningType("").Strings()

	allIPAllocationPolicyOptions = types.VAppIPAssignmentInfoIpAllocationPolicy("").Strings()

	allIPProtocolOptions = types.VAppIPAssignmentInfoProtocols("").Strings()
)

func SpecMap(e *ovf.Envelope, hidden, verbose bool) (res []Property) {
	if e == nil || e.VirtualSystem == nil {
		return nil
	}

	for _, p := range e.VirtualSystem.Product {
		for i, v := range p.Property {
			if v.UserConfigurable == nil {
				continue
			}
			if !*v.UserConfigurable && !hidden {
				continue
			}

			d := ""
			if v.Default != nil {
				d = *v.Default
			}

			// vSphere only accept True/False as boolean values for some reason
			if v.Type == "boolean" {
				d = cases.Title(language.Und).String(d)
			}

			np := Property{KeyValue: KeyValue{Key: p.Key(v), Value: d}}

			if verbose {
				np.Spec = &p.Property[i]
			}

			res = append(res, np)
		}
	}

	return
}

func Spec(fpath string, a Archive, hidden, verbose bool) (*Options, error) {
	e := &ovf.Envelope{}
	if fpath != "" {
		d, err := ReadOvf(fpath, a)
		if err != nil {
			return nil, err
		}

		if e, err = ReadEnvelope(d); err != nil {
			return nil, err
		}
	}

	var deploymentOptions []string
	if e.DeploymentOption != nil && e.DeploymentOption.Configuration != nil {
		// add default first
		for _, c := range e.DeploymentOption.Configuration {
			if c.Default != nil && *c.Default {
				deploymentOptions = append(deploymentOptions, c.ID)
			}
		}

		for _, c := range e.DeploymentOption.Configuration {
			if c.Default == nil || !*c.Default {
				deploymentOptions = append(deploymentOptions, c.ID)
			}
		}
	}

	o := Options{
		DiskProvisioning:   string(types.OvfCreateImportSpecParamsDiskProvisioningTypeFlat),
		IPAllocationPolicy: string(types.VAppIPAssignmentInfoIpAllocationPolicyDhcpPolicy),
		IPProtocol:         string(types.VAppIPAssignmentInfoProtocolsIPv4),
		MarkAsTemplate:     false,
		PowerOn:            false,
		WaitForIP:          false,
		InjectOvfEnv:       false,
		PropertyMapping:    SpecMap(e, hidden, verbose),
	}

	if deploymentOptions != nil {
		o.Deployment = deploymentOptions[0]
	}

	if e.VirtualSystem != nil && e.VirtualSystem.Annotation != nil {
		o.Annotation = e.VirtualSystem.Annotation.Annotation
	}

	if e.Network != nil {
		for _, net := range e.Network.Networks {
			o.NetworkMapping = append(o.NetworkMapping, Network{net.Name, ""})
		}
	}

	if verbose {
		if deploymentOptions != nil {
			o.AllDeploymentOptions = deploymentOptions
		}
		o.AllDiskProvisioningOptions = allDiskProvisioningOptions
		o.AllIPAllocationPolicyOptions = allIPAllocationPolicyOptions
		o.AllIPProtocolOptions = allIPProtocolOptions
	}

	return &o, nil
}
