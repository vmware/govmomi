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

package tpm2

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"

	"github.com/google/go-tpm/tpm2"
)

type seal struct {
	*flags.VirtualMachineFlag
	*flags.OutputFlag

	input string
	alg   string
}

func init() {
	cli.Register("vm.tpm2.seal", &seal{})
}

func (cmd *seal) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.StringVar(&cmd.input, "in", "-",
		`Input data. Defaults to STDIN via "-"`)
	f.StringVar(&cmd.alg, "G", "rsa",
		`Public key algorithm. Either "rsa", "ecc", or "ecdsa"`)
}

func (cmd *seal) Description() string {
	return `Seal plain-text data to the VM's TPM2 endorsement key.

Examples:
  govc vm.tpm2.seal -vm VM -in plain.txt
  echo "Hello, world" | govc vm.tpm2.seal -vm VM
  echo "Seal with ECC." | govc vm.tpm2.seal -vm VM -G ecc`
}

func (cmd *seal) Process(ctx context.Context) error {
	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *seal) Run(ctx context.Context, f *flag.FlagSet) error {

	var alg x509.PublicKeyAlgorithm
	switch strings.ToLower(cmd.alg) {
	case "rsa":
		alg = x509.RSA
	case "ecc", "ecdsa":
		alg = x509.ECDSA
	default:
		return flag.ErrHelp
	}

	vm, err := cmd.VirtualMachine()
	if err != nil {
		return err
	}

	if vm == nil {
		return flag.ErrHelp
	}

	// Read the plain-text data.
	var plainTextFile *os.File
	if cmd.input == "-" {
		plainTextFile = os.Stdin
	} else {
		f, err := os.Open(cmd.input)
		if err != nil {
			return err
		}
		plainTextFile = f
		defer f.Close()
	}
	plainTextData, err := io.ReadAll(plainTextFile)
	if err != nil {
		return err
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

	var ekCert *x509.Certificate
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

		if alg == x509.RSA && cert.PublicKeyAlgorithm == x509.RSA {
			ekCert = cert
		} else if alg == x509.ECDSA && cert.PublicKeyAlgorithm == x509.ECDSA {
			ekCert = cert
		}
	}

	if ekCert == nil {
		return fmt.Errorf("unable to find ek certificate")
	}

	ek, err := tpm2.EKCertToTPMTPublic(*ekCert)
	if err != nil {
		return err
	}

	pub, priv, seed, err := tpm2.EKSeal(ek, plainTextData)
	if err != nil {
		return err
	}

	return cmd.WriteResult(sealResult{
		Public:  base64.StdEncoding.EncodeToString(tpm2.Marshal(pub)),
		Private: base64.StdEncoding.EncodeToString(tpm2.Marshal(priv)),
		Seed:    base64.StdEncoding.EncodeToString(tpm2.Marshal(seed)),
	})
}

type sealResult struct {
	Public  string `json:"public"`
	Private string `json:"private"`
	Seed    string `json:"seed"`
}

func (r sealResult) Write(w io.Writer) error {
	_, err := fmt.Fprintf(
		w,
		"%s@@NULL@@%s@@NULL@@%s",
		r.Public,
		r.Private,
		r.Seed,
	)
	return err
}
