/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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
		for _, a := range e.VirtualSystem.Annotation {
			o.Annotation += a.Annotation
		}
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
