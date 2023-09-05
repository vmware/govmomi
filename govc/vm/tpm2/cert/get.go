/*
Copyright (c) 2023-2023 VMware, Inc. All Rights Reserved.

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

package cert

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"flag"
	"fmt"
	"io"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type get struct {
	*flags.VirtualMachineFlag
	*flags.OutputFlag

	fingerprint string
}

func init() {
	cli.Register("vm.tpm2.cert.get", &get{})
}

func (cmd *get) Register(ctx context.Context, f *flag.FlagSet) {

	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.StringVar(&cmd.fingerprint, "fingerprint", "",
		`Fingerprint of cert to get. Use vm.tpm2.ls to see available certs."`)
}

func (cmd *get) Description() string {
	return `Get certificate by fingerprint.

Examples:
  govc vm.tpm2.cert.get -vm VM -fingerprint 41:5D:F1:AE:B9:F2:B1:22:9F:79:B7:FF:DA:55:5B:86`
}

func (cmd *get) Process(ctx context.Context) error {

	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *get) Run(ctx context.Context, f *flag.FlagSet) error {
	if cmd.fingerprint == "" {
		return flag.ErrHelp
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	// Get the VM's EK.
	var moVM mo.VirtualMachine
	if err := vm.Properties(
		ctx,
		vm.Reference(),
		[]string{"config.hardware.device"},
		&moVM); err != nil {
		return err
	}

	devices := object.VirtualDeviceList(moVM.Config.Hardware.Device)
	selectedDevices := devices.SelectByType(&types.VirtualTPM{})
	if len(selectedDevices) == 0 {
		return fmt.Errorf("no VirtualTPM devices found")
	}
	if len(selectedDevices) > 1 {
		return fmt.Errorf("multiple VirtualTPM devices found")
	}

	vtpmDev := selectedDevices[0].(*types.VirtualTPM)

	var result getResult
	for i := range vtpmDev.EndorsementKeyCertificate {
		// Use DecodeString as Decode complains about trailing data.
		derString := string(vtpmDev.EndorsementKeyCertificate[i])
		derData, err := base64.StdEncoding.DecodeString(derString)
		if err != nil {
			return err
		}
		cert, err := x509.ParseCertificate(derData)
		if err != nil {
			return err
		}

		fingerprint := md5.Sum(cert.Raw)
		var fingerprintBuf bytes.Buffer
		for i, f := range fingerprint {
			if i > 0 {
				fmt.Fprintf(&fingerprintBuf, ":")
			}
			fmt.Fprintf(&fingerprintBuf, "%02X", f)
		}

		if cmd.fingerprint == fingerprintBuf.String() {
			var pemBuf bytes.Buffer
			block := &pem.Block{
				Type:  "CERTIFICATE",
				Bytes: cert.Raw,
			}
			if err := pem.Encode(&pemBuf, block); err != nil {
				return err
			}
			result.PubCertPEM = pemBuf.Bytes()
			break
		}
	}

	if len(result.PubCertPEM) == 0 {
		return fmt.Errorf("fingerprint %s not found", cmd.fingerprint)
	}

	return cmd.WriteResult(result)
}

type getResult struct {
	PubCertPEM []byte `json:"pubCertPEM"`
}

func (r getResult) Write(w io.Writer) error {
	_, err := fmt.Fprint(w, string(r.PubCertPEM))
	return err
}
