/*
Copyright (c) 2019-2024 VMware, Inc. All Rights Reserved.

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

package library

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/vmware/govmomi/vapi/internal"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25/soap"
)

// TransferEndpoint provides information on the source of a library item file.
type TransferEndpoint struct {
	URI                      string `json:"uri,omitempty"`
	SSLCertificate           string `json:"ssl_certificate,omitempty"`
	SSLCertificateThumbprint string `json:"ssl_certificate_thumbprint,omitempty"`
}

type ProbeResult struct {
	Status         string                    `json:"status"`
	SSLThumbprint  string                    `json:"ssl_thumbprint,omitempty"`
	SSLCertificate string                    `json:"ssl_certificate,omitempty"`
	ErrorMessages  []rest.LocalizableMessage `json:"error_messages,omitempty"`
}

// UpdateFile is the specification for the updatesession
// operations file:add, file:get, and file:list.
type UpdateFile struct {
	BytesTransferred int64                    `json:"bytes_transferred,omitempty"`
	Checksum         *Checksum                `json:"checksum_info,omitempty"`
	ErrorMessage     *rest.LocalizableMessage `json:"error_message,omitempty"`
	Name             string                   `json:"name"`
	Size             int64                    `json:"size,omitempty"`
	SourceEndpoint   *TransferEndpoint        `json:"source_endpoint,omitempty"`
	SourceType       string                   `json:"source_type"`
	Status           string                   `json:"status,omitempty"`
	UploadEndpoint   *TransferEndpoint        `json:"upload_endpoint,omitempty"`
}

// FileValidationError contains the validation error of a file in the update session
type FileValidationError struct {
	Name         string                  `json:"name"`
	ErrorMessage rest.LocalizableMessage `json:"error_message"`
}

// UpdateFileValidation contains the result of validating the files in the update session
type UpdateFileValidation struct {
	HasErrors    bool                  `json:"has_errors"`
	MissingFiles []string              `json:"missing_files,omitempty"`
	InvalidFiles []FileValidationError `json:"invalid_files,omitempty"`
}

// AddLibraryItemFile adds a file
func (c *Manager) AddLibraryItemFile(ctx context.Context, sessionID string, updateFile UpdateFile) (*UpdateFile, error) {
	url := c.Resource(internal.LibraryItemUpdateSessionFile).WithID(sessionID).WithAction("add")
	spec := struct {
		FileSpec UpdateFile `json:"file_spec"`
	}{updateFile}
	var res UpdateFile
	err := c.Do(ctx, url.Request(http.MethodPost, spec), &res)
	if err != nil {
		return nil, err
	}
	if res.Status == "ERROR" {
		return nil, res.ErrorMessage
	}
	return &res, nil
}

// AddLibraryItemFileFromURI adds a file from a remote URI.
func (c *Manager) AddLibraryItemFileFromURI(ctx context.Context, sessionID, name, uri string, checksum ...Checksum) (*UpdateFile, error) {
	source := &TransferEndpoint{
		URI: uri,
	}

	file := UpdateFile{
		Name:           name,
		SourceType:     "PULL",
		SourceEndpoint: source,
	}

	if len(checksum) == 1 && checksum[0].Checksum != "" {
		file.Checksum = &checksum[0]
	} else if len(checksum) > 1 {
		return nil, fmt.Errorf("expected 0 or 1 checksum, got %d", len(checksum))
	}

	if res, err := c.Head(uri); err == nil {
		file.Size = res.ContentLength
		if res.TLS != nil {
			source.SSLCertificateThumbprint = soap.ThumbprintSHA1(res.TLS.PeerCertificates[0])
		}
	} else {
		res, err := c.ProbeTransferEndpoint(ctx, *source)
		if err != nil {
			return nil, err
		}
		if res.SSLCertificate != "" {
			source.SSLCertificate = res.SSLCertificate
		} else {
			source.SSLCertificateThumbprint = res.SSLThumbprint
		}
	}

	return c.AddLibraryItemFile(ctx, sessionID, file)
}

// GetLibraryItemUpdateSessionFile retrieves information about a specific file
// that is a part of an update session.
func (c *Manager) GetLibraryItemUpdateSessionFile(ctx context.Context, sessionID string, fileName string) (*UpdateFile, error) {
	url := c.Resource(internal.LibraryItemUpdateSessionFile).WithID(sessionID).WithAction("get")
	spec := struct {
		Name string `json:"file_name"`
	}{fileName}
	var res UpdateFile
	return &res, c.Do(ctx, url.Request(http.MethodPost, spec), &res)
}

// ListLibraryItemUpdateSessionFile lists all files in the library item associated with the update session
func (c *Manager) ListLibraryItemUpdateSessionFile(ctx context.Context, sessionID string) ([]UpdateFile, error) {
	url := c.Resource(internal.LibraryItemUpdateSessionFile).WithParam("update_session_id", sessionID)
	var res []UpdateFile
	return res, c.Do(ctx, url.Request(http.MethodGet), &res)
}

// ValidateLibraryItemUpdateSessionFile validates all files in the library item associated with the update session
func (c *Manager) ValidateLibraryItemUpdateSessionFile(ctx context.Context, sessionID string) (*UpdateFileValidation, error) {
	url := c.Resource(internal.LibraryItemUpdateSessionFile).WithID(sessionID).WithAction("validate")
	var res UpdateFileValidation
	return &res, c.Do(ctx, url.Request(http.MethodPost), &res)
}

// RemoveLibraryItemUpdateSessionFile requests a file to be removed. The file will only be effectively removed when the update session is completed.
func (c *Manager) RemoveLibraryItemUpdateSessionFile(ctx context.Context, sessionID string, fileName string) error {
	url := c.Resource(internal.LibraryItemUpdateSessionFile).WithID(sessionID).WithAction("remove")
	spec := struct {
		Name string `json:"file_name"`
	}{fileName}
	return c.Do(ctx, url.Request(http.MethodPost, spec), nil)
}

func (c *Manager) ProbeTransferEndpoint(ctx context.Context, endpoint TransferEndpoint) (*ProbeResult, error) {
	url := c.Resource(internal.LibraryItemUpdateSessionFile).WithAction("probe")
	spec := struct {
		SourceEndpoint TransferEndpoint `json:"source_endpoint"`
	}{endpoint}
	var res ProbeResult
	return &res, c.Do(ctx, url.Request(http.MethodPost, spec), &res)
}

// ReadManifest converts an ovf manifest to a map of file name -> Checksum.
func ReadManifest(m io.Reader) (map[string]*Checksum, error) {
	// expected format: openssl sha1 *.{ovf,vmdk}
	c := make(map[string]*Checksum)

	scanner := bufio.NewScanner(m)
	for scanner.Scan() {
		line := strings.SplitN(scanner.Text(), ")=", 2)
		if len(line) != 2 {
			continue
		}
		name := strings.SplitN(line[0], "(", 2)
		if len(name) != 2 {
			continue
		}
		sum := &Checksum{
			Algorithm: strings.TrimSpace(name[0]),
			Checksum:  strings.TrimSpace(line[1]),
		}
		c[name[1]] = sum
	}

	return c, scanner.Err()
}
