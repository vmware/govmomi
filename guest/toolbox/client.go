/*
Copyright (c) 2017 VMware, Inc. All Rights Reserved.

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

package toolbox

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/vmware/govmomi/guest"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

// Client attempts to expose guest.OperationsManager as idiomatic Go interfaces
type Client struct {
	ProcessManager *guest.ProcessManager
	FileManager    *guest.FileManager
	Authentication types.BaseGuestAuthentication
	GuestFamily    types.VirtualMachineGuestOsFamily
}

// NewClient initializes a Client's ProcessManager, FileManager and GuestFamily
func NewClient(ctx context.Context, c *vim25.Client, vm mo.Reference, auth types.BaseGuestAuthentication) (*Client, error) {
	m := guest.NewOperationsManager(c, vm.Reference())

	pm, err := m.ProcessManager(ctx)
	if err != nil {
		return nil, err
	}

	fm, err := m.FileManager(ctx)
	if err != nil {
		return nil, err
	}

	family := ""
	var props mo.VirtualMachine
	pc := property.DefaultCollector(c)
	err = pc.RetrieveOne(context.Background(), vm.Reference(), []string{"guest.guestFamily", "guest.toolsInstallType"}, &props)
	if err != nil {
		return nil, err
	}

	if props.Guest != nil {
		family = props.Guest.GuestFamily
		if family == string(types.VirtualMachineGuestOsFamilyOtherGuestFamily) {
			if props.Guest.ToolsInstallType == string(types.VirtualMachineToolsInstallTypeGuestToolsTypeMSI) {
				// The case of Windows version not supported by the ESX version
				family = string(types.VirtualMachineGuestOsFamilyWindowsGuest)
			}
		}
	}

	return &Client{
		ProcessManager: pm,
		FileManager:    fm,
		Authentication: auth,
		GuestFamily:    types.VirtualMachineGuestOsFamily(family),
	}, nil
}

func (c *Client) rm(ctx context.Context, path string) {
	err := c.FileManager.DeleteFile(ctx, c.Authentication, path)
	if err != nil {
		log.Printf("rm %q: %s", path, err)
	}
}

func (c *Client) mktemp(ctx context.Context) (string, error) {
	return c.FileManager.CreateTemporaryFile(ctx, c.Authentication, "govmomi-", "", "")
}

type exitError struct {
	error
	exitCode int
}

func (e *exitError) ExitCode() int {
	return e.exitCode
}

// Run implements exec.Cmd.Run over vmx guest RPC against standard vmware-tools or toolbox.
func (c *Client) Run(ctx context.Context, cmd *exec.Cmd) error {
	if cmd.Stdin != nil {
		dst, err := c.mktemp(ctx)
		if err != nil {
			return err
		}

		defer c.rm(ctx, dst)

		var buf bytes.Buffer
		size, err := io.Copy(&buf, cmd.Stdin)
		if err != nil {
			return err
		}

		p := soap.DefaultUpload
		p.ContentLength = size
		attr := new(types.GuestPosixFileAttributes)

		err = c.Upload(ctx, &buf, dst, p, attr, true)
		if err != nil {
			return err
		}

		cmd.Args = append(cmd.Args, "<", dst)
	}

	output := []struct {
		io.Writer
		fd   string
		path string
	}{
		{cmd.Stdout, "1", ""},
		{cmd.Stderr, "2", ""},
	}

	for i, out := range output {
		if out.Writer == nil {
			continue
		}

		dst, err := c.mktemp(ctx)
		if err != nil {
			return err
		}

		defer c.rm(ctx, dst)

		cmd.Args = append(cmd.Args, out.fd+">", dst)
		output[i].path = dst
	}

	path := cmd.Path
	args := cmd.Args

	switch c.GuestFamily {
	case types.VirtualMachineGuestOsFamilyWindowsGuest:
		// Using 'cmd.exe /c' is required on Windows for i/o redirection
		path = "c:\\Windows\\System32\\cmd.exe"
		args = append([]string{"/c", cmd.Path}, args...)
	default:
		if !strings.ContainsAny(cmd.Path, "/") {
			// vmware-tools requires an absolute ProgramPath
			// Default to 'bash -c' as a convenience
			path = "/bin/bash"
			arg := "'" + strings.Join(append([]string{cmd.Path}, args...), " ") + "'"
			args = []string{"-c", arg}
		}
	}

	spec := types.GuestProgramSpec{
		ProgramPath:      path,
		Arguments:        strings.Join(args, " "),
		EnvVariables:     cmd.Env,
		WorkingDirectory: cmd.Dir,
	}

	pid, err := c.ProcessManager.StartProgram(ctx, c.Authentication, &spec)
	if err != nil {
		return err
	}

	rc := 0
	for {
		procs, err := c.ProcessManager.ListProcesses(ctx, c.Authentication, []int64{pid})
		if err != nil {
			return err
		}

		p := procs[0]
		if p.EndTime == nil {
			<-time.After(time.Second / 2)
			continue
		}

		rc = int(p.ExitCode)

		break
	}

	for _, out := range output {
		if out.Writer == nil {
			continue
		}

		f, _, err := c.Download(ctx, out.path)
		if err != nil {
			return err
		}

		_, err = io.Copy(out.Writer, f)
		_ = f.Close()
		if err != nil {
			return err
		}
	}

	if rc != 0 {
		return &exitError{fmt.Errorf("%s: exit %d", cmd.Path, rc), rc}
	}

	return nil
}

// archiveReader wraps an io.ReadCloser to support streaming download
// of a guest directory, stops reading once it sees the stream trailer.
// This is only useful when guest tools is the Go toolbox.
// The trailer is required since TransferFromGuest requires a Content-Length,
// which toolbox doesn't know ahead of time as the gzip'd tarball never touches the disk.
// We opted to wrap this here for now rather than guest.FileManager so
// DownloadFile can be also be used as-is to handle this use case.
type archiveReader struct {
	io.ReadCloser
}

var (
	gzipHeader    = []byte{0x1f, 0x8b, 0x08} // rfc1952 {ID1, ID2, CM}
	gzipHeaderLen = len(gzipHeader)
)

func (r *archiveReader) Read(buf []byte) (int, error) {
	nr, err := r.ReadCloser.Read(buf)

	// Stop reading if the last N bytes are the gzipTrailer
	if nr >= gzipHeaderLen {
		if bytes.Equal(buf[nr-gzipHeaderLen:nr], gzipHeader) {
			nr -= gzipHeaderLen
			err = io.EOF
		}
	}

	return nr, err
}

func isDir(src string) bool {
	u, err := url.Parse(src)
	if err != nil {
		return false
	}

	return strings.HasSuffix(u.Path, "/")
}

// Download initiates a file transfer from the guest
func (c *Client) Download(ctx context.Context, src string) (io.ReadCloser, int64, error) {
	vc := c.ProcessManager.Client()

	info, err := c.FileManager.InitiateFileTransferFromGuest(ctx, c.Authentication, src)
	if err != nil {
		return nil, 0, err
	}

	u, err := c.FileManager.TransferURL(ctx, info.Url)
	if err != nil {
		return nil, 0, err
	}

	p := soap.DefaultDownload

	f, n, err := vc.Download(ctx, u, &p)
	if err != nil {
		return nil, n, err
	}

	if strings.HasPrefix(src, "/archive:/") || isDir(src) {
		f = &archiveReader{ReadCloser: f} // look for the gzip trailer
	}

	return f, n, nil
}

// Upload transfers a file to the guest
func (c *Client) Upload(ctx context.Context, src io.Reader, dst string, p soap.Upload, attr types.BaseGuestFileAttributes, force bool) error {
	vc := c.ProcessManager.Client()

	var err error

	if p.ContentLength == 0 { // Content-Length is required
		switch r := src.(type) {
		case *bytes.Buffer:
			p.ContentLength = int64(r.Len())
		case *bytes.Reader:
			p.ContentLength = int64(r.Len())
		case *strings.Reader:
			p.ContentLength = int64(r.Len())
		case *os.File:
			info, serr := r.Stat()
			if serr != nil {
				return serr
			}

			p.ContentLength = info.Size()
		}

		if p.ContentLength == 0 { // os.File for example could be a device (stdin)
			buf := new(bytes.Buffer)

			p.ContentLength, err = io.Copy(buf, src)
			if err != nil {
				return err
			}

			src = buf
		}
	}

	url, err := c.FileManager.InitiateFileTransferToGuest(ctx, c.Authentication, dst, attr, p.ContentLength, force)
	if err != nil {
		return err
	}

	u, err := c.FileManager.TransferURL(ctx, url)
	if err != nil {
		return err
	}

	return vc.Client.Upload(ctx, src, u, &p)
}
