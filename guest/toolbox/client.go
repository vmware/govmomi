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
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"

	"github.com/vmware/govmomi/guest"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

// Client attempts to expose guest.OperationsManager as idiomatic Go interfaces
type Client struct {
	ProcessManager *guest.ProcessManager
	FileManager    *guest.FileManager
	Authentication types.BaseGuestAuthentication
}

// RoundTrip implements http.RoundTripper over vmx guest RPC.
// This transport depends on govmomi/toolbox running in the VM guest and does not work with standard VMware tools.
// Using this transport makes it is possible to connect to HTTP endpoints that are bound to the VM's loopback address.
// Note that the toolbox's http.RoundTripper only supports the "http" scheme, "https" is not supported.
func (c *Client) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Scheme != "http" {
		return nil, fmt.Errorf("%q scheme not supported", req.URL.Scheme)
	}

	ctx := req.Context()

	req.Header.Set("Connection", "close") // we need the server to close the connection after 1 request

	spec := types.GuestProgramSpec{
		ProgramPath: "http.RoundTrip",
		Arguments:   req.URL.Host,
	}

	pid, err := c.ProcessManager.StartProgram(ctx, c.Authentication, &spec)
	if err != nil {
		return nil, err
	}

	dst := fmt.Sprintf("/proc/%d/stdin", pid)
	src := fmt.Sprintf("/proc/%d/stdout", pid)

	var buf bytes.Buffer
	err = req.Write(&buf)
	if err != nil {
		return nil, err
	}

	attr := new(types.GuestPosixFileAttributes)
	size := int64(buf.Len())

	url, err := c.FileManager.InitiateFileTransferToGuest(ctx, c.Authentication, dst, attr, size, true)
	if err != nil {
		return nil, err
	}

	vc := c.ProcessManager.Client()

	u, err := vc.ParseURL(url)
	if err != nil {
		return nil, err
	}

	p := soap.DefaultUpload
	p.ContentLength = size

	err = vc.Client.Upload(&buf, u, &p)
	if err != nil {
		return nil, err
	}

	info, err := c.FileManager.InitiateFileTransferFromGuest(ctx, c.Authentication, src)
	if err != nil {
		return nil, err
	}

	u, err = vc.ParseURL(info.Url)
	if err != nil {
		return nil, err
	}

	f, _, err := vc.Client.Download(u, &soap.DefaultDownload)
	if err != nil {
		return nil, err
	}

	return http.ReadResponse(bufio.NewReader(f), req)
}

// Run implements exec.Cmd.Run over vmx guest RPC.
func (c *Client) Run(ctx context.Context, cmd *exec.Cmd) error {
	vc := c.ProcessManager.Client()

	spec := types.GuestProgramSpec{
		ProgramPath:      cmd.Path,
		Arguments:        strings.Join(cmd.Args, " "),
		EnvVariables:     cmd.Env,
		WorkingDirectory: cmd.Dir,
	}

	pid, serr := c.ProcessManager.StartProgram(ctx, c.Authentication, &spec)
	if serr != nil {
		return serr
	}

	if cmd.Stdin != nil {
		dst := fmt.Sprintf("/proc/%d/stdin", pid)

		var buf bytes.Buffer
		size, err := io.Copy(&buf, cmd.Stdin)
		if err != nil {
			return err
		}

		attr := new(types.GuestPosixFileAttributes)

		url, err := c.FileManager.InitiateFileTransferToGuest(ctx, c.Authentication, dst, attr, size, true)
		if err != nil {
			return err
		}

		u, err := vc.ParseURL(url)
		if err != nil {
			return err
		}

		p := soap.DefaultUpload
		p.ContentLength = size

		err = vc.Client.Upload(&buf, u, &p)
		if err != nil {
			return err
		}
	}

	names := []string{"out", "err"}

	for i, w := range []io.Writer{cmd.Stdout, cmd.Stderr} {
		if w == nil {
			continue
		}

		src := fmt.Sprintf("/proc/%d/std%s", pid, names[i])

		info, err := c.FileManager.InitiateFileTransferFromGuest(ctx, c.Authentication, src)
		if err != nil {
			return err
		}

		u, err := vc.ParseURL(info.Url)
		if err != nil {
			return err
		}

		f, _, err := vc.Client.Download(u, &soap.DefaultDownload)
		if err != nil {
			return err
		}

		_, err = io.Copy(w, f)
		_ = f.Close()
		if err != nil {
			return err
		}
	}

	procs, err := c.ProcessManager.ListProcesses(ctx, c.Authentication, []int64{pid})
	if err != nil {
		return err
	}

	if len(procs) == 1 {
		rc := procs[0].ExitCode
		if rc != 0 {
			return fmt.Errorf("%s: exit %d", cmd.Path, rc)
		}
	}

	return nil
}
