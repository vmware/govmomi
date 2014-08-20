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

package flags

import (
	"encoding/json"
	"flag"
	"io"
	"os"
)

type OutputWriter interface {
	WriteTo(io.Writer) error
}

type OutputFlag struct {
	JSON bool
	TTY  bool
}

func (flag *OutputFlag) Register(f *flag.FlagSet) {
	f.BoolVar(&flag.JSON, "json", false, "Enable JSON output")
}

func (flag *OutputFlag) Process() error {
	if !flag.JSON {
		// Assume we have a tty if not outputting JSON
		flag.TTY = true
	}

	return nil
}

func (flag *OutputFlag) Write(b []byte) (int, error) {
	if !flag.TTY {
		return 0, nil
	}

	n, err := os.Stdout.Write(b)
	os.Stdout.Sync()
	return n, err
}

func (flag *OutputFlag) WriteResult(result OutputWriter) error {
	var err error
	var out = os.Stdout

	if flag.JSON {
		enc := json.NewEncoder(out)
		err = enc.Encode(result)
	} else {
		err = result.WriteTo(out)
	}

	return err
}
