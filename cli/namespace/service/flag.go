// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package service

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"os"

	"github.com/vmware/govmomi/vapi/namespace"
)

// Common processing of flags for service and service versions creation/update commands.
type ServiceVersionFlag struct {
	SpecType        string
	TrustedProvider bool
	AcceptEULA      bool
	content         string
}

func (svf *ServiceVersionFlag) Register(ctx context.Context, f *flag.FlagSet) {
	f.StringVar(&svf.SpecType, "spec-type", "carvel", "Type of Spec: vsphere (deprecated) or carvel")
	f.BoolVar(&svf.TrustedProvider, "trusted", false, "Define if this is a trusted provider (only applicable for vSphere spec type)")
	f.BoolVar(&svf.AcceptEULA, "accept-eula", false, "Auto accept EULA (only required for vSphere spec type)")
}

func (svf *ServiceVersionFlag) Process(ctx context.Context) error {
	if svf.SpecType != "vsphere" && svf.SpecType != "carvel" {
		return fmt.Errorf("Invalid type: '%v', only 'vsphere' and 'carvel' specs are supported", svf.SpecType)
	}
	return nil
}

// SupervisorServiceVersionSpec returns a spec for a supervisor service version definition
func (svf *ServiceVersionFlag) SupervisorServiceVersionSpec(manifestFile string) (namespace.SupervisorService, error) {
	service := namespace.SupervisorService{}
	manifest, err := os.ReadFile(manifestFile)
	if err != nil {
		return service, fmt.Errorf("failed to read manifest file: %s", err)
	}

	content := base64.StdEncoding.EncodeToString(manifest)
	if svf.SpecType == "carvel" {
		service.CarvelService = &namespace.SupervisorServicesCarvelSpec{
			VersionSpec: namespace.CarvelVersionCreateSpec{
				Content: content,
			},
		}
	} else {
		service.VsphereService = &namespace.SupervisorServicesVSphereSpec{
			VersionSpec: namespace.SupervisorServicesVSphereVersionCreateSpec{
				Content:         content,
				AcceptEula:      svf.AcceptEULA,
				TrustedProvider: svf.TrustedProvider,
			},
		}
	}
	return service, nil
}
