/*
Copyright (c) 2018 VMware, Inc. All Rights Reserved.

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
	"archive/tar"
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/ovf"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/rest"
)

type ova struct {
	*flags.ClientFlag
	*flags.OutputFlag
	item library.Item
}

func init() {
	cli.Register("library.ova", &ova{})
}

func (cmd *ova) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *ova) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ova) Description() string {
	return `Upload and OVA file.

Examples:
  govc library.ova library_name file.ova`
}

type ovaResult []library.Library

func (r ovaResult) Write(w io.Writer) error {
	for i := range r {
		fmt.Fprintln(w, r[i].Name)
	}
	return nil
}

// OVAFile is a wrapper around the tar reader
type OVAFile struct {
	filename string
	file     *os.File
	tarFile  *tar.Reader
}

// NewOVAFile creates a new OVA file reader
func NewOVAFile(filename string) (*OVAFile, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	tarFile := tar.NewReader(f)
	return &OVAFile{filename: filename, file: f, tarFile: tarFile}, nil
}

// Find looks for a filename match in the OVA file
func (of *OVAFile) Find(filename string) (*tar.Header, error) {
	for {
		header, err := of.tarFile.Next()
		if err == io.EOF {
			return nil, err
		}
		if header.Name == filename {
			return header, nil
		}
	}
}

// Next returns the next header in the OVA file
func (of *OVAFile) Next() (*tar.Header, error) {
	return of.tarFile.Next()
}

// Read reads from the current file in the OVA file
func (of *OVAFile) Read(b []byte) (int, error) {
	if of.tarFile == nil {
		return 0, io.EOF
	}
	return of.tarFile.Read(b)
}

// Close will close the file associated with the OVA file
func (of *OVAFile) Close() error {
	return of.file.Close()
}

// getOVAFileInfo opens an OVA, finds the file entry, and returns both the size and md5 checksum
func getOVAFileInfo(ovafile string, filename string) (int64, string, error) {
	of, err := NewOVAFile(ovafile)
	hdr, err := of.Find(filename)

	hash := md5.New()
	_, err = io.Copy(hash, of)
	if err != nil {
		return 0, "", err
	}
	md5String := hex.EncodeToString(hash.Sum(nil)[:16])

	return hdr.Size, md5String, nil
}

// uploadFile will upload a single file from an OVA using the sessionID provided
func uploadFile(ctx context.Context, m *library.Manager, sessionID string, ovafile string, filename string) error {
	var updateFileInfo library.UpdateFile

	fmt.Printf("Uploading %s from %s\n", filename, ovafile)
	size, md5String, _ := getOVAFileInfo(ovafile, filename)

	// Get the URI for the file upload

	updateFileInfo.Name = filename
	updateFileInfo.Size = &size
	updateFileInfo.SourceType = "PUSH"
	updateFileInfo.Checksum = &library.Checksum{
		Algorithm: "MD5",
		Checksum:  md5String,
	}

	addFileInfo, err := m.AddLibraryItemFile(ctx, sessionID, updateFileInfo)
	if err != nil {
		return err
	}

	of, err := NewOVAFile(ovafile)
	if err != nil {
		return err
	}
	defer of.Close()

	// Setup to point to the OVA file to be transferred
	_, err = of.Find(filename)

	req, err := http.NewRequest("PUT", addFileInfo.UploadEndpoint.URI, of)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("vmware-api-session-id", sessionID)

	// TODO: this likely should be done using the rest client api
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (cmd *ova) Run(ctx context.Context, f *flag.FlagSet) error {
	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		m := library.NewManager(c)
		var err error
		var fileList []string
		var sessionSpec library.UpdateSession

		if f.NArg() != 2 {
			return flag.ErrHelp
		}

		libname := f.Arg(0)
		ovafilename := f.Arg(1)

		if !strings.HasSuffix(ovafilename, ".ova") {
			return fmt.Errorf("Filename does not end in '.ova'")
		}

		o, err := NewOVAFile(ovafilename)
		if err != nil {
			return err
		}
		hdr, err := o.Next()

		// Per the OVA spec, the first file must be the .ovf file.
		if !strings.HasSuffix(hdr.Name, ".ovf") {
			return fmt.Errorf("OVA file must have the .ovf file first, found %s", hdr.Name)
		}
		fmt.Printf("Found OVF file: %s\n", hdr.Name)

		// Parse the OVF file to find the file references
		e, err := ovf.Unmarshal(o)
		if err != nil {
			return err
		}

		// Create list of files to upload
		fileList = append(fileList, hdr.Name)
		for _, ref := range e.References {
			fmt.Printf("Found OVF file reference: %s\n", ref.Href)
			_, err := o.Find(ref.Href)
			if err != nil {
				return fmt.Errorf("Could not find OVF referenced file: %s", ref.Href)
			}
			fileList = append(fileList, ref.Href)
		}

		// Find the library ID
		library, err := m.GetLibraryByName(ctx, libname)
		if err != nil {
			return err
		}

		// Create a library item
		cmd.item.Name = filepath.Base(ovafilename)
		cmd.item.LibraryID = library.ID
		cmd.item.Type = "ovf"
		cmd.item.Description = "Testing 1 2 3"

		itemID, err := m.CreateLibraryItem(ctx, cmd.item)
		if err != nil {
			return err
		}
		fmt.Printf("Library item id: %s\n", itemID)

		// Create the update session to use for uploading the file
		sessionSpec.LibraryItemID = itemID
		sessionID, err := m.CreateLibraryItemUpdateSession(ctx, sessionSpec)
		if err != nil {
			return err
		}
		fmt.Printf("Update session: %s\n", sessionID)

		// Upload all the files
		for _, filename := range fileList {
			err = uploadFile(ctx, m, sessionID, ovafilename, filename)
			if err != nil {
				return err
			}
		}

		// Complete the session
		err = m.CompleteLibraryItemUpdateSession(ctx, sessionID)
		if err != nil {
			return err
		}
		return nil
	})
}
