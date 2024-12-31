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

package vmdk

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
	"text/template"
)

type Descriptor struct {
	Encoding  string
	Version   int
	CID       DiskContentID
	ParentCID DiskContentID
	Type      string
	Extent    []Extent
	DDB       map[string]string
}

type DiskContentID uint32

func (cid DiskContentID) String() string {
	return fmt.Sprintf("%0.8x", uint32(cid))
}

type Extent struct {
	Type       string
	Permission string
	Size       uint64
	Info       string
}

func NewDescriptor(extent ...Extent) *Descriptor {
	for i := range extent {
		if extent[i].Type == "" {
			extent[i].Type = "VMFS"
		}
		if extent[i].Permission == "" {
			extent[i].Permission = "RW"
		}
	}
	return &Descriptor{
		Version:  1,
		Encoding: "UTF-8",
		Type:     "vmfs",
		DDB:      map[string]string{},
		Extent:   extent,
	}
}

func ParseDescriptor(r io.Reader) (*Descriptor, error) {
	d := NewDescriptor()

	scanner := bufio.NewScanner(r)

	// NOTE: not doing any validation currently, or using this function yet.
	// Will add validation as needed when use-cases are implemented.
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if d.parseExtent(line) {
			continue
		}

		s := strings.SplitN(line, "=", 2)

		if len(s) != 2 {
			continue
		}

		key, val := strings.TrimSpace(s[0]), strings.TrimSpace(s[1])
		val = strings.Trim(val, `"`)
		if strings.HasPrefix(key, "ddb") {
			d.DDB[key] = val
			continue
		}

		switch strings.ToLower(key) {
		case "encoding":
			d.Encoding = val
		case "version":
			d.Version, _ = strconv.Atoi(val)
		case "cid":
			_, _ = fmt.Sscanf(val, "%x", &d.CID)
		case "parentcid":
			_, _ = fmt.Sscanf(val, "%x", &d.ParentCID)
		case "createType":
			d.Type = val
		}
	}

	return d, scanner.Err()
}

var permissions = []string{"RDONLY", "RW", "NOACCESS"}

func (d *Descriptor) parseExtent(line string) bool {
	// Each extent is defined by a line following this pattern:
	// perm size type "%s"

	s := strings.SplitN(line, " ", 2)

	if len(s) != 2 || !slices.Contains(permissions, s[0]) {
		return false
	}

	x := Extent{
		Permission: s[0],
	}

	s = strings.SplitN(s[1], " ", 2)
	size, err := strconv.ParseUint(s[0], 10, 64)
	if len(s) != 2 || err != nil {
		return false
	}

	x.Size = size

	s = strings.SplitN(s[1], " ", 2)
	x.Type = s[0]

	if len(s) == 2 {
		x.Info = strings.Trim(s[1], `"`)
	}

	d.Extent = append(d.Extent, x)

	return true
}

var descriptor = `# Disk DescriptorFile
version={{ .Version }}
encoding="{{ .Encoding }}"
CID={{ .CID }}
parentCID={{ .ParentCID }}
createType="{{ .Type }}"

# Extent description{{range .Extent }}
{{ .Permission }} {{ .Size }} {{ .Type }} "{{ .Info }}"{{end}}

# The Disk Data Base
#DDB{{ range $key, $val := .DDB }}
{{ $key }} = "{{ $val }}"{{end}}
`

func (d *Descriptor) Write(w io.Writer) error {
	t, err := template.New("vmdk").Parse(descriptor)
	if err != nil {
		return err
	}
	return t.Execute(w, d)
}
