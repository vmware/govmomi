/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package importx

import (
	"archive/tar"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/vmware/govmomi/govc/cli"
)

type ova struct {
	*ovf
}

func init() {
	cmd := &ova{
		ovf: newOvf(),
	}

	cli.Register("import.ova", cmd)
}

func (cmd *ova) Run(f *flag.FlagSet) error {
	file, err := cmd.Prepare(f)

	if err != nil {
		return err
	}

	return cmd.Import(file)
}

// ImportOVA extracts a .ova file to a temporary directory,
// then imports as it would a .ovf file.
func (cmd *ova) Import(i importable) error {
	var ovf importable

	f, err := os.Open(string(i))
	if err != nil {
		return err
	}
	defer f.Close()

	dir, err := ioutil.TempDir("", "govc-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	r := tar.NewReader(f)
	for {
		h, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		path := filepath.Join(dir, h.Name)
		entry, err := os.Create(path)
		if err != nil {
			return err
		}

		fmt.Printf("Extracting %s...\n", h.Name)

		if _, err := io.Copy(entry, r); err != nil {
			_ = entry.Close()
			return err
		}

		if err := entry.Close(); err != nil {
			return err
		}

		if filepath.Ext(path) == ".ovf" {
			ovf = importable(path)
		}
	}

	return cmd.ovf.Import(ovf)
}
