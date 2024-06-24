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

package simulator

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"net"

	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

type HostCertificateManager struct {
	mo.HostCertificateManager

	Host *mo.HostSystem
}

func (m *HostCertificateManager) init(r *Registry) {
	for _, obj := range r.objects {
		if h, ok := obj.(*HostSystem); ok {
			if h.ConfigManager.CertificateManager.Value == m.Self.Value {
				m.Host = &h.HostSystem
			}
		}
	}
}

func NewHostCertificateManager(h *mo.HostSystem) *HostCertificateManager {
	m := &HostCertificateManager{Host: h}

	_ = m.InstallServerCertificate(SpoofContext(), &types.InstallServerCertificate{
		Cert: string(m.Host.Config.Certificate),
	})

	return m
}

func (m *HostCertificateManager) InstallServerCertificate(ctx *Context, req *types.InstallServerCertificate) soap.HasFault {
	body := new(methods.InstallServerCertificateBody)

	var info object.HostCertificateInfo
	cert := []byte(req.Cert)
	_, err := info.FromPEM(cert)
	if err != nil {
		body.Fault_ = Fault(err.Error(), new(types.HostConfigFault))
		return body
	}

	m.CertificateInfo = info.HostCertificateManagerCertificateInfo

	m.Host.Config.Certificate = cert

	body.Res = new(types.InstallServerCertificateResponse)

	return body
}

func (m *HostCertificateManager) GenerateCertificateSigningRequest(ctx *Context, req *types.GenerateCertificateSigningRequest) soap.HasFault {
	block, _ := pem.Decode(m.Host.Config.Certificate)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic(err)
	}

	csr := x509.CertificateRequest{
		Subject:            cert.Subject,
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	if req.UseIpAddressAsCommonName {
		csr.IPAddresses = []net.IP{net.ParseIP(m.Host.Summary.ManagementServerIp)}
	} else {
		csr.DNSNames = []string{m.Host.Name}
	}

	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	der, _ := x509.CreateCertificateRequest(rand.Reader, &csr, key)
	var buf bytes.Buffer
	err = pem.Encode(&buf, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: der})
	if err != nil {
		panic(err)
	}

	return &methods.GenerateCertificateSigningRequestBody{
		Res: &types.GenerateCertificateSigningRequestResponse{
			Returnval: buf.String(),
		},
	}
}
