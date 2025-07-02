// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmdk

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
	"text/template"

	"github.com/vmware/govmomi/units"
)

type Descriptor struct {
	Encoding  string            `json:"encoding"`
	Version   int               `json:"version"`
	CID       DiskContentID     `json:"cid"`
	ParentCID DiskContentID     `json:"parentCID"`
	Type      string            `json:"type"`
	Extent    []Extent          `json:"extent"`
	DDB       map[string]string `json:"ddb"`
}

type DiskContentID uint32

func (cid DiskContentID) String() string {
	return fmt.Sprintf("%0.8x", uint32(cid))
}

type Extent struct {
	Type       string `json:"type"`
	Permission string `json:"permission"`
	Size       int64  `json:"size"`
	Info       string `json:"info"`
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
		token := bytes.Trim(scanner.Bytes(), "\x00")
		line := strings.TrimSpace(string(token))
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
		if k := strings.TrimPrefix(key, "ddb."); k != key {
			d.DDB[k] = val
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
	size, err := strconv.ParseInt(s[0], 10, 64)
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

# Extent description ({{ cap }} capacity){{range .Extent }}
{{ .Permission }} {{ .Size }} {{ .Type }} "{{ .Info }}"{{end}}

# The Disk Data Base
#DDB{{ range $key, $val := .DDB }}
ddb.{{ $key }} = "{{ $val }}"{{end}}
`

func (d *Descriptor) Write(w io.Writer) error {
	t, err := template.New("vmdk").Funcs(template.FuncMap{
		"cap": func() string {
			return units.ByteSize(d.Capacity()).String()
		},
	}).Parse(descriptor)
	if err != nil {
		return err
	}
	return t.Execute(w, d)
}

// Capacity in bytes of the vmdk
func (d *Descriptor) Capacity() int64 {
	var size int64

	for i := range d.Extent {
		size += d.Extent[i].Size * SectorSize
	}

	return size
}
