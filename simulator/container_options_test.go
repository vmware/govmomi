// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"testing"
)

const containerOptionJSON = "{\n  \"message\": \"a 'quote', spaces and a \\\"double quote\\\"\",\n  \"template\": \"$HOME ${SERVICE_NAME} $5.00\",\n  \"enabled\": true\n}"

func TestContainerCreateStructuredOptions(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("requires a POSIX shell")
	}
	if !json.Valid([]byte(containerOptionJSON)) {
		t.Fatal("invalid JSON fixture")
	}
	dir := t.TempDir()
	capture := filepath.Join(dir, "arguments")
	script := "#!/bin/sh\nif [ \"$1\" = create ]; then printf '%s\\000' \"$@\" > \"$OPTION_ARGUMENTS\"; if [ \"$OPTION_FAIL\" = 1 ]; then printf '%s\\n' \"$@\" >&2; exit 64; fi; printf 'created-id\\n'; fi\n"
	if err := os.WriteFile(filepath.Join(dir, "docker"), []byte(script), 0700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("OPTION_ARGUMENTS", capture)
	mount := filepath.Join(dir, "a directory with 'quotes' and $dollars") + ":/guest path:ro"
	sentinel := filepath.Join(dir, "unevaluated")
	literal := "$(printf evaluated > " + shellQuote(sentinel) + ")"
	env := []string{"VMX_GUESTINFO_APPLICATION_CONFIG=" + containerOptionJSON, "APP_STATIC_LOCAL_PROFILE=1", "EMPTY_VALUE=", "SHELL_TEMPLATE=" + literal}
	c, err := create(&Context{Map: NewRegistry()}, "option-test", "uid", []string{"none"}, []string{mount}, []string{"127.0.0.1:12345:80/tcp"}, env, "--read-only "+shellQuote("busybox:stable"), []string{"echo", shellQuote("ordinary text")})
	if err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(capture)
	if err != nil {
		t.Fatal(err)
	}
	raw := bytes.Split(bytes.TrimSuffix(data, []byte{0}), []byte{0})
	got := make([]string, len(raw))
	for i := range raw {
		got[i] = string(raw[i])
	}
	want := []string{"create", "--name", c.name, "--network", "none", "-v", mount, "-p", "127.0.0.1:12345:80/tcp"}
	for _, value := range env {
		want = append(want, "--env", value)
	}
	want = append(want, "--read-only", "busybox:stable", "echo", "ordinary text")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("engine arguments = %q, want %q", got, want)
	}
	if _, err := os.Lstat(sentinel); !os.IsNotExist(err) {
		t.Fatal("literal environment value was evaluated by the shell")
	}
	t.Setenv("OPTION_FAIL", "1")
	_, err = create(&Context{Map: NewRegistry()}, "failed-option-test", "uid", nil, nil, nil, env, shellQuote("busybox:stable"), nil)
	if err == nil || !strings.Contains(err.Error(), "exit status 64") || strings.Contains(err.Error(), "VMX_GUESTINFO_APPLICATION_CONFIG") {
		t.Fatalf("unsafe or missing failure classification: %v", err)
	}
}

type commandLog struct {
	sync.Mutex
	bytes.Buffer
}

func (b *commandLog) Write(p []byte) (int, error) {
	b.Lock()
	defer b.Unlock()
	return b.Buffer.Write(p)
}
func (b *commandLog) text() string { b.Lock(); defer b.Unlock(); return b.Buffer.String() }

func TestCommandErrorDiagnostics(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("requires a POSIX shell")
	}
	const data = "ordinary-configuration-value"
	command := exec.Command(shell, "-c", "printf '%s' '"+data+"' >&2; exit 64")
	_, err := command.Output()
	if err == nil {
		t.Fatal("expected command failure")
	}
	var captured commandLog
	previous := log.Writer()
	log.SetOutput(&captured)
	t.Cleanup(func() { log.SetOutput(previous) })
	got := commandError("container create", err)
	if strings.Contains(got.Error(), data) || strings.Contains(captured.text(), data) {
		t.Fatal("command diagnostics exposed configuration data")
	}
	if !strings.Contains(got.Error(), "container create failed") || !strings.Contains(got.Error(), "exit status 64") {
		t.Fatalf("missing safe operation/status: %v", got)
	}
	var exited *exec.ExitError
	if !errors.As(got, &exited) || exited.ExitCode() != 64 || len(exited.Stderr) != 0 {
		t.Fatal("exit classification did not discard captured stderr")
	}
	for _, cause := range []error{context.Canceled, context.DeadlineExceeded, exec.ErrNotFound, os.ErrPermission, os.ErrNotExist} {
		wrapped := fmt.Errorf("%s: %w", data, cause)
		safe := commandError("volume population", wrapped)
		if !errors.Is(safe, cause) || strings.Contains(safe.Error(), data) {
			t.Fatalf("unsafe or lost classification: %v", safe)
		}
	}
	unknown := commandError("container create", errors.New(data))
	if strings.Contains(unknown.Error(), data) || strings.Contains(captured.text(), data) {
		t.Fatal("unknown error text exposed configuration data")
	}
}
