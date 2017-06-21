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
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/vmware/govmomi/toolbox/vix"
)

var (
	xmlEscape *strings.Replacer

	shell = "/bin/sh"

	defaultOwner = os.Getenv("USER")
)

func init() {
	// See: VixToolsEscapeXMLString
	chars := []string{
		`"`,
		"%",
		"&",
		"'",
		"<",
		">",
	}

	replace := make([]string, 0, len(chars)*2)

	for _, c := range chars {
		replace = append(replace, c)
		replace = append(replace, url.QueryEscape(c))
	}

	xmlEscape = strings.NewReplacer(replace...)

	// See procMgrPosix.c:ProcMgrStartProcess:
	// Prefer bash -c as is uses exec() to replace itself,
	// whereas bourne shell does a fork & exec, so two processes are started.
	if sh, err := exec.LookPath("bash"); err != nil {
		shell = sh
	}

	if defaultOwner == "" {
		defaultOwner = "toolbox"
	}
}

// ProcessState is the toolbox representation of the GuestProcessInfo type
type ProcessState struct {
	Name      string
	Args      string
	Owner     string
	Pid       int64
	ExitCode  int32
	StartTime int64
	EndTime   int64
}

func (s *ProcessState) toXML() string {
	const format = "<proc>" +
		"<cmd>%s</cmd>" +
		"<name>%s</name>" +
		"<pid>%d</pid>" +
		"<user>%s</user>" +
		"<start>%d</start>" +
		"<eCode>%d</eCode>" +
		"<eTime>%d</eTime>" +
		"</proc>"

	name := filepath.Base(s.Name)

	argv := []string{s.Name}

	if len(s.Args) != 0 {
		argv = append(argv, xmlEscape.Replace(s.Args))
	}

	args := strings.Join(argv, " ")

	exit := atomic.LoadInt32(&s.ExitCode)
	end := atomic.LoadInt64(&s.EndTime)

	return fmt.Sprintf(format, name, args, s.Pid, s.Owner, s.StartTime, exit, end)
}

// Process managed by the ProcessManager.
type Process struct {
	ProcessState

	Start func(context.Context, *vix.StartProgramRequest) (int64, error)
	Wait  func() error
	Kill  context.CancelFunc

	ctx context.Context
}

// ProcessError can be returned by the Process.Wait function to propagate ExitCode to ProcessState.
type ProcessError struct {
	Err      error
	ExitCode int32
}

func (e *ProcessError) Error() string {
	return e.Err.Error()
}

// ProcessManager manages processes within the guest.
// See: http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.vm.guest.ProcessManager.html
type ProcessManager struct {
	wg      sync.WaitGroup
	mu      sync.Mutex
	expire  time.Duration
	entries map[int64]*Process
	pids    sync.Pool
}

// NewProcessManager creates a new ProcessManager instance.
func NewProcessManager() *ProcessManager {
	// We use pseudo PIDs that don't conflict with OS PIDs, so they can live in the same table.
	// For the pseudo PIDs, we use a sync.Pool rather than a plain old counter to avoid the unlikely,
	// but possible wrapping should such a counter exceed MaxInt64.
	pid := int64(32768) // TODO: /proc/sys/kernel/pid_max

	return &ProcessManager{
		expire:  time.Minute * 5,
		entries: make(map[int64]*Process),
		pids: sync.Pool{
			New: func() interface{} {
				return atomic.AddInt64(&pid, 1)
			},
		},
	}
}

// Start calls the Process.Start function, returning the pid on success or an error.
// A goroutine is started that calls the Process.Wait function.  After Process.Wait has
// returned, the ProcessState EndTime and ExitCode fields are set.  The process state can be
// queried via ListProcessesInGuest until it is removed, 5 minutes after Wait returns.
func (m *ProcessManager) Start(r *vix.StartProgramRequest, p *Process) (int64, error) {
	p.Name = r.ProgramPath
	p.Args = r.Arguments

	// Owner is cosmetic, but useful for example with: govc guest.ps -U $uid
	if p.Owner == "" {
		p.Owner = defaultOwner
	}

	p.StartTime = time.Now().Unix()

	p.ctx, p.Kill = context.WithCancel(context.Background())

	pid, err := p.Start(p.ctx, r)
	if err != nil {
		return -1, err
	}

	if pid == 0 {
		p.Pid = m.pids.Get().(int64) // pseudo pid for funcs
	} else {
		p.Pid = pid
	}

	m.mu.Lock()
	m.entries[p.Pid] = p
	m.mu.Unlock()

	m.wg.Add(1)
	go func() {
		werr := p.Wait()

		atomic.StoreInt64(&p.EndTime, time.Now().Unix())

		if werr != nil {
			rc := int32(1)
			if xerr, ok := werr.(*ProcessError); ok {
				rc = xerr.ExitCode
			}

			atomic.StoreInt32(&p.ExitCode, rc)
		}

		m.wg.Done()

		// See: http://pubs.vmware.com/vsphere-65/topic/com.vmware.wssdk.apiref.doc/vim.vm.guest.ProcessManager.ProcessInfo.html
		// "If the process was started using StartProgramInGuest then the process completion time
		//  will be available if queried within 5 minutes after it completes."
		<-time.After(m.expire)

		m.mu.Lock()
		delete(m.entries, p.Pid)
		m.mu.Unlock()

		if pid == 0 {
			m.pids.Put(p.Pid) // pseudo pid can be reused now
		}
	}()

	return p.Pid, nil
}

// Kill cancels the Process Context.
// Returns true if pid exists in the process table, false otherwise.
func (m *ProcessManager) Kill(pid int64) bool {
	m.mu.Lock()
	entry, ok := m.entries[pid]
	m.mu.Unlock()

	if ok {
		entry.Kill()
		return true
	}

	return false
}

// ListProcesses marshals the ProcessState for the given pids.
// If no pids are specified, all current processes are included.
// The return value can be used for responding to a VixMsgListProcessesExRequest.
func (m *ProcessManager) ListProcesses(pids []int64) []byte {
	w := new(bytes.Buffer)

	m.mu.Lock()

	if len(pids) == 0 {
		for _, p := range m.entries {
			_, _ = w.WriteString(p.toXML())
		}
	} else {
		for _, id := range pids {
			p, ok := m.entries[id]
			if !ok {
				continue
			}

			_, _ = w.WriteString(p.toXML())
		}
	}

	m.mu.Unlock()

	return w.Bytes()
}

type processFunc struct {
	wg sync.WaitGroup

	run func(context.Context, string) error

	err error
}

// NewProcessFunc creates a new Process, where the Start function calls the given run function within a goroutine.
// The Wait function waits for the goroutine to finish and returns the error returned by run.
// The run ctx param may be used to return early via the ProcessManager.Kill method.
// The run args command is that of the VixMsgStartProgramRequest.Arguments field.
func NewProcessFunc(run func(ctx context.Context, args string) error) *Process {
	f := &processFunc{run: run}

	return &Process{
		Start: f.start,
		Wait:  f.wait,
	}
}

func (f *processFunc) start(ctx context.Context, r *vix.StartProgramRequest) (int64, error) {
	f.wg.Add(1)

	go func() {
		f.err = f.run(ctx, r.Arguments)

		f.wg.Done()
	}()

	return 0, nil
}

func (f *processFunc) wait() error {
	f.wg.Wait()
	return f.err
}

type processCmd struct {
	cmd *exec.Cmd
}

// NewProcess creates a new Process, where the Start function use exec.CommandContext to create and start the process.
// The Wait function waits for the process to finish and returns the error returned by exec.Cmd.Wait().
// Prior to Wait returning, the exec.Cmd.Wait() error is used to set the ProcessState.ExitCode, if error is of type exec.ExitError.
// The ctx param may be used to kill the process via the ProcessManager.Kill method.
// The VixMsgStartProgramRequest param fields are mapped to the exec.Cmd counterpart fields.
// Processes are started within a sub-shell, allowing for i/o redirection, just as with the C version of vmware-tools.
func NewProcess() *Process {
	c := new(processCmd)

	return &Process{
		Start: c.start,
		Wait:  c.wait,
	}
}

func (c *processCmd) start(ctx context.Context, r *vix.StartProgramRequest) (int64, error) {
	name, err := exec.LookPath(r.ProgramPath)
	if err != nil {
		return -1, err
	}
	// #nosec: Subprocess launching with variable
	// Note that processCmd is currently used only for testing.
	c.cmd = exec.CommandContext(ctx, shell, "-c", fmt.Sprintf("%s %s", name, r.Arguments))
	c.cmd.Dir = r.WorkingDir
	c.cmd.Env = r.EnvVars

	err = c.cmd.Start()
	if err != nil {
		return -1, err
	}

	return int64(c.cmd.Process.Pid), nil
}

func (c *processCmd) wait() error {
	err := c.cmd.Wait()
	if err != nil {
		xerr := &ProcessError{
			Err:      err,
			ExitCode: 1,
		}

		if x, ok := err.(*exec.ExitError); ok {
			if status, ok := x.Sys().(syscall.WaitStatus); ok {
				xerr.ExitCode = int32(status.ExitStatus())
			}
		}

		return xerr
	}

	return nil
}
