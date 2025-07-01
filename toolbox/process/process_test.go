// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package process

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/vmware/govmomi/toolbox/vix"
)

func checkGOOS(t *testing.T) {
	switch runtime.GOOS {
	case "linux", "darwin":
	default:
		t.Skipf("GOOS=%s", runtime.GOOS)
	}
}

func TestProcessFunction(t *testing.T) {
	checkGOOS(t)

	m := NewManager()
	var pids []int64

	for i := 0; i <= 2; i++ {
		r := &vix.StartProgramRequest{
			ProgramPath: "test",
			Arguments:   strconv.Itoa(i),
		}

		pid, _ := m.Start(r, NewFunc(func(_ context.Context, arg string) error {
			rc, _ := strconv.Atoi(arg)
			if rc == 0 {
				return nil

			}
			return &Error{Err: errors.New("fail"), ExitCode: int32(rc)}
		}))

		if pid == 0 {
			t.Fatalf("no pid")
		}

		pids = append(pids, pid)
	}

	m.wg.Wait()

	_ = m.ListProcesses(pids)

	for i, pid := range pids {
		p := m.entries[pid]
		if p.ExitCode != int32(i) {
			t.Errorf("%d: %d != %d", pid, p.ExitCode, i)
		}
	}
}

func TestProcessCommand(t *testing.T) {
	checkGOOS(t)

	m := NewManager()
	var pids []int64

	for i := 0; i <= 2; i++ {
		r := &vix.StartProgramRequest{
			ProgramPath: shell,
			Arguments:   fmt.Sprintf(`-c "exit %d"`, i),
		}

		pid, err := m.Start(r, New())
		if err != nil {
			t.Fatal(err)
		}
		pids = append(pids, pid)
	}

	m.wg.Wait()

	_ = m.ListProcesses(nil)

	for i, pid := range pids {
		p := m.entries[pid]
		if p.ExitCode != int32(i) {
			t.Errorf("%d: %d != %d", pid, p.ExitCode, i)
		}
	}

	r := &vix.StartProgramRequest{
		ProgramPath: shell,
	}

	shell = "/enoent/enoent"
	_, err := m.Start(r, New())
	if err == nil {
		t.Error("expected error")
	}
	shell = r.ProgramPath

	r.ProgramPath = "/enoent/enoent"
	_, err = m.Start(r, New())
	if err == nil {
		t.Error("expected error")
	}
}

func TestProcessKill(t *testing.T) {
	checkGOOS(t)

	m := NewManager()
	var pids []int64

	procs := []struct {
		r *vix.StartProgramRequest
		p *Process
	}{
		{
			&vix.StartProgramRequest{
				ProgramPath: "test",
				Arguments:   "none",
			},
			NewFunc(func(ctx context.Context, _ string) error {
				select {
				case <-ctx.Done():
					return &Error{Err: ctx.Err(), ExitCode: 42}
				case <-time.After(time.Minute):
				}

				return nil
			}),
		},
		{
			&vix.StartProgramRequest{
				ProgramPath: shell,
				Arguments:   fmt.Sprintf(`-c "while true; do sleep 1; done"`),
			},
			New(),
		},
	}

	for _, test := range procs {
		pid, err := m.Start(test.r, test.p)
		if err != nil {
			t.Fatal(err)
		}

		pids = append(pids, pid)
	}

	for {
		b := m.ListProcesses(pids)
		if bytes.Count(b, []byte("<proc>")) == len(pids) {
			break
		}

		<-time.After(time.Millisecond * 100)
	}

	for _, pid := range pids {
		if !m.Kill(pid) {
			t.Errorf("kill %d", pid)
		}
	}

	m.wg.Wait()

	for _, pid := range pids {
		p := m.entries[pid]

		if p.ExitCode == 0 {
			t.Errorf("%s: exit=%d", p.Name, p.ExitCode)
		}
	}

	if m.Kill(-1) {
		t.Error("kill -1")
	}
}

func TestProcessRemove(t *testing.T) {
	checkGOOS(t)

	m := NewManager()

	m.expire = time.Millisecond

	r := &vix.StartProgramRequest{
		ProgramPath: "test",
	}

	pid, _ := m.Start(r, NewFunc(func(_ context.Context, arg string) error {
		return nil
	}))

	m.wg.Wait()

	<-time.After(m.expire * 20)
	// pid should be removed by now
	b := m.ListProcesses([]int64{pid})
	if len(b) != 0 {
		t.Error("expected 0 processes")
	}
}

func TestEscapeXML(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{`echo "foo bar" > /dev/null`, "echo %22foo bar%22 %3E /dev/null"},
	}

	for i, test := range tests {
		e := EscapeXML.Replace(test.in)
		if e != test.out {
			t.Errorf("%d: %s != %s", i, e, test.out)
		}
	}
}

func TestProcessError(t *testing.T) {
	fault := errors.New("fail")
	var err error = &Error{Err: fault}

	if err.Error() != fault.Error() {
		t.Fatal()
	}
}

func TestProcessIO(t *testing.T) {
	checkGOOS(t)

	m := NewManager()

	r := &vix.StartProgramRequest{
		ProgramPath: "/bin/date",
	}

	p := New().WithIO()

	_, err := m.Start(r, p)
	if err != nil {
		t.Fatal(err)
	}

	m.wg.Wait()

	var buf bytes.Buffer

	_, _ = io.Copy(&buf, p.IO.Out)

	if buf.Len() == 0 {
		t.Error("no data")
	}
}

type testRoundTripper struct {
	*Process
}

func (c *testRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Connection", "close") // we need the server to close the connection after 1 request

	err := req.Write(c.IO.In.Writer)
	if err != nil {
		return nil, err
	}

	_ = c.IO.In.Close()

	<-c.ctx.Done()

	return http.ReadResponse(bufio.NewReader(c.IO.Out), req)
}

func TestProcessRoundTripper(t *testing.T) {
	checkGOOS(t)

	echo := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = r.Write(w)
	}))

	u, _ := url.Parse(echo.URL)

	m := NewManager()

	r := &vix.StartProgramRequest{
		ProgramPath: "http.RoundTrip",
		Arguments:   u.Host,
	}

	p := NewRoundTrip()

	_, err := m.Start(r, p)
	if err != nil {
		t.Fatal(err)
	}

	res, err := (&http.Client{Transport: &testRoundTripper{p}}).Get(echo.URL)
	if err != nil {
		t.Logf("Err: %s", p.IO.Err.String())
		t.Fatal(err)
	}

	if res.ContentLength == 0 {
		t.Errorf("len=%d", res.ContentLength)
	}
}
