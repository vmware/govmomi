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

package item

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vapi/library"
	"github.com/vmware/govmomi/vapi/rest"
)

type upload struct {
	*flags.ClientFlag
	*flags.OutputFlag
	item        library.Item
	description string
}

func init() {
	cli.Register("library.item.upload", &upload{})
}

func (cmd *upload) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag, ctx = flags.NewClientFlag(ctx)
	cmd.OutputFlag, ctx = flags.NewOutputFlag(ctx)
	cmd.ClientFlag.Register(ctx, f)
	cmd.OutputFlag.Register(ctx, f)
	f.StringVar(&cmd.item.Description, "d", "", "Description of library")
}

func (cmd *upload) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *upload) Description() string {
	return `List library items.

Examples:
  govc library.item.upload library_name file_name
  govc library.item.upload library_name file_name -json | jq .`
}

type uploadResult []library.Item

func (cmd *upload) Run(ctx context.Context, f *flag.FlagSet) error {
	return cmd.WithRestClient(ctx, func(c *rest.Client) error {
		m := library.NewManager(c)
		var sessionSpec library.UpdateSession
		var updateFileInfo library.UpdateFile
		var err error

		if f.NArg() != 2 {
			return flag.ErrHelp
		}

		libname := f.Arg(0)
		filename := f.Arg(1)

		// Make sure the file can me opened and calculate the md5 checksum
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		size := fileInfo.Size()

		hash := md5.New()
		_, err = io.Copy(hash, file)
		if err != nil {
			return err
		}
		md5String := hex.EncodeToString(hash.Sum(nil)[:16])

		lib, err := m.GetLibraryByName(ctx, libname)
		if err != nil {
			return err
		}

		// Create a library item
		cmd.item.Name = filepath.Base(filename)
		cmd.item.LibraryID = lib.ID
		cmd.item.Type = "iso"

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

		fmt.Printf("Upload uri: %s\n", addFileInfo.UploadEndpoint.URI)

		transferFile, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer transferFile.Close()

		req, err := http.NewRequest("PUT", addFileInfo.UploadEndpoint.URI, transferFile)
		if err != nil {
			return err
		}
		req.Header.Set("vmware-api-session-id", sessionID)

		err = c.Do(ctx, req, nil)
		if err != nil {
			return err
		}

		// Complete the session
		return m.CompleteLibraryItemUpdateSession(ctx, sessionID)
	})
}
