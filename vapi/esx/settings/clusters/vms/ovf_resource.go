// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vms

// LocationType defines the supported locations of OVF packages.
type LocationType string

const (
	RemoteFile LocationType = "REMOTE_FILE"
	LocalFile  LocationType = "LOCAL_FILE"
)

type SslCertificateValidation string

const (
	SslCertificateValidationEnabled  SslCertificateValidation = "ENABLED"
	SslCertificateValidationDisabled SslCertificateValidation = "DISABLED"
)

type AuthenticationScheme string

const (
	None            AuthenticationScheme = "NONE"
	VmwareSessionId AuthenticationScheme = "VMWARE_SESSION_ID"
)

// OvfResource  contains fields that describe the location of an OVF package
// and a configuration for its download.
type OvfResource struct {

	// LocationType of OVF package.
	LocationType LocationType `json:"location_type"`

	// Url to the file server or the local VC file system where the OVF package
	// can be downloaded. The supported URI schemes are http, https, and file.
	Url string `json:"url"`

	// SslCertificateValidation configuration for SSL Certificate validation of
	// the URL specified by the Url.
	SslCertificateValidation SslCertificateValidation `json:"ssl_certificate_validation"`

	// Certificate that is to be trusted by vLCM when downloading the OVF
	// package from a file server.
	Certificate string `json:"certificate,omitempty"`

	// AuthenticationScheme is the authentication scheme needed to access the OVF URL.
	AuthenticationScheme AuthenticationScheme `json:"authentication_scheme"`
}
