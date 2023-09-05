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
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type ls struct {
	*flags.VirtualMachineFlag
	*flags.OutputFlag

	alg string
}

func init() {
	cli.Register("vm.tpm2.cert.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {

	cmd.VirtualMachineFlag, ctx = flags.NewVirtualMachineFlag(ctx)
	cmd.VirtualMachineFlag.Register(ctx, f)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.OutputFlag.Register(ctx, f)

	f.StringVar(&cmd.alg, "G", "",
		`Public key algorithm. Either "rsa", "ecc", or "ecdsa"`)
}

func (cmd *ls) Description() string {
	return `List endorsement key certificates.

Examples:
  govc vm.tpm2.cert.ls -vm VM
  govc vm.tpm2.cert.ls -vm VM -G ecc`
}

func (cmd *ls) Process(ctx context.Context) error {

	if err := cmd.VirtualMachineFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	var alg x509.PublicKeyAlgorithm
	switch strings.ToLower(cmd.alg) {
	case "":
		alg = x509.UnknownPublicKeyAlgorithm
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

	var result lsResult
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

		var skipCert bool
		switch {
		case alg == x509.RSA && cert.PublicKeyAlgorithm != x509.RSA:
			skipCert = true
		case alg == x509.ECDSA && cert.PublicKeyAlgorithm != x509.ECDSA:
			skipCert = true
		}
		if skipCert {
			continue
		}

		fingerprint := md5.Sum(cert.Raw)
		var buf bytes.Buffer
		for i, f := range fingerprint {
			if i > 0 {
				fmt.Fprintf(&buf, ":")
			}
			fmt.Fprintf(&buf, "%02X", f)
		}

		info := certInfo{Fingerprint: buf.String()}
		switch cert.PublicKeyAlgorithm {
		case x509.RSA:
			info.Algorithm = "rsa"
		case x509.ECDSA:
			info.Algorithm = "ecc"
		}
		result = append(result, info)
	}

	return cmd.WriteResult(result)
}

type certInfo struct {
	Algorithm   string `json:"algorithm"`
	Fingerprint string `json:"fingerprint"`
}

type lsResult []certInfo

func (r lsResult) Write(w io.Writer) error {
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "Algorithm\tFingerprint\n")
	for i := range r {
		fmt.Fprintf(tw, "%s\t%s\n", r[i].Algorithm, r[i].Fingerprint)
	}
	return tw.Flush()
}
