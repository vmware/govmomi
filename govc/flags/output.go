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
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/vmware/govmomi/vim25/soap"
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

// Log outputs the specified string, prefixed with the current time.
// A newline is not automatically added. If the specified string
// starts with a '\r', the current line is cleared first.
func (flag *OutputFlag) Log(s string) (int, error) {
	if len(s) > 0 && s[0] == '\r' {
		flag.Write([]byte{'\r', 033, '[', 'K'})
		s = s[1:]
	}

	return flag.WriteString(time.Now().Format("[02-01-06 15:04:05] ") + s)
}

func (flag *OutputFlag) Write(b []byte) (int, error) {
	if !flag.TTY {
		return 0, nil
	}

	n, err := os.Stdout.Write(b)
	os.Stdout.Sync()
	return n, err
}

func (flag *OutputFlag) WriteString(s string) (int, error) {
	return flag.Write([]byte(s))
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

const (
	KiB = 1024
	MiB = 1024 * KiB
	GiB = 1024 * MiB
)

func (flag *OutputFlag) ProgressLogger(prefix string, ch chan soap.Progress) *sync.WaitGroup {
	var wg sync.WaitGroup

	go func() {
		var bps int64
		var pos int64
		var f float32
		var err error

		tick := time.NewTicker(1 * time.Second)
		defer wg.Done()

		for done := false; !done && err == nil; {
			select {
			case p, ok := <-ch:
				if !ok {
					done = true
					break
				}

				bps += (p.Pos - pos)
				pos = p.Pos
				f = float32(p.Pos) / float32(p.Size)
				err = p.Error
			case <-tick.C:
				b := ""
				switch {
				case bps >= GiB:
					b = fmt.Sprintf("%.1fGiB", float32(bps)/float32(GiB))
				case bps >= MiB:
					b = fmt.Sprintf("%.1fMiB", float32(bps)/float32(MiB))
				case bps >= KiB:
					b = fmt.Sprintf("%.1fKiB", float32(bps)/float32(KiB))
				default:
					b = fmt.Sprintf("%.1fB", bps)
				}

				flag.Log(fmt.Sprintf("\r%s(%.0f%%, %s/s)", prefix, 100*f, b))
				bps = 0
			}
		}

		if err != nil && err != io.EOF {
			flag.Log(fmt.Sprintf("\r%sError: %s\n", prefix, err))
		} else {
			flag.Log(fmt.Sprintf("\r%sOK\n", prefix))
		}
	}()

	wg.Add(1)

	return &wg
}
