// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package simulator

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

// TestContainerStructuredOptionsRuntime uses an explicitly selected local image.
// It never pulls an image, starts a service, or joins an external network.
func TestContainerStructuredOptionsRuntime(t *testing.T) {
	if os.Getenv("VCSIM_OPTIONS_RUNTIME") != "1" {
		t.Skip("set VCSIM_OPTIONS_RUNTIME=1 with a local immutable shell image")
	}
	image := os.Getenv("VCSIM_OPTIONS_RUNTIME_IMAGE")
	if !strings.Contains(image, "@sha256:") {
		t.Fatal("an immutable local image reference is required")
	}
	if !json.Valid([]byte(containerOptionJSON)) {
		t.Fatal("invalid ordinary JSON fixture")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := exec.CommandContext(ctx, "docker", "image", "inspect", image).Run(); err != nil {
		t.Fatal("selected local image unavailable")
	}
	marker := uuid.NewString()
	name := "vcsim-options-" + marker
	label := "vcsim.option-test=" + marker
	dir := filepath.Join(t.TempDir(), "ordinary 'quotes' and $spaces")
	if err := os.Mkdir(dir, 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "value"), []byte("mounted bytes\n"), 0600); err != nil {
		t.Fatal(err)
	}
	sentinel := filepath.Join(t.TempDir(), "unevaluated")
	literal := "$(printf evaluated > " + shellQuote(sentinel) + ")"
	id := ""
	t.Cleanup(func() {
		cleanup, stop := context.WithTimeout(context.Background(), 15*time.Second)
		defer stop()
		if id == "" {
			body, err := exec.CommandContext(cleanup, "docker", "inspect", "--format", `{{.Id}} {{index .Config.Labels "vcsim.option-test"}}`, name).Output()
			if err != nil {
				t.Error("owned container identity is uncertain after failed create")
			}
			if err == nil {
				fields := strings.Fields(string(body))
				if len(fields) == 2 && fields[1] == marker {
					id = fields[0]
				} else {
					t.Error("owned container identity did not match after failed create")
				}
			}
		}
		if id != "" {
			if err := exec.CommandContext(cleanup, "docker", "rm", "-v", "-f", id).Run(); err != nil {
				t.Error("owned option-test container cleanup failed")
			}
		}
		left, err := exec.CommandContext(cleanup, "docker", "ps", "-a", "-q", "--filter", "label="+label).Output()
		if err != nil || len(bytes.TrimSpace(left)) != 0 {
			t.Error("owned option-test container absence not confirmed")
		}
	})
	run := []string{"docker", "create", "--pull=never", "--name", name, "--label", label, "--network", "none", "--volume", dir + ":/fixture:ro", "--env", "VMX_GUESTINFO_APPLICATION_CONFIG=" + containerOptionJSON, "--env", "APP_STATIC_LOCAL_PROFILE=1", "--env", "EMPTY_VALUE=", "--env", "SHELL_TEMPLATE=" + literal}
	suffix := []string{shellQuote(image)}
	for _, arg := range []string{"sh", "-c", `printf '%s\000' "$VMX_GUESTINFO_APPLICATION_CONFIG" "$APP_STATIC_LOCAL_PROFILE" "${EMPTY_VALUE+x}" "$EMPTY_VALUE" "$SHELL_TEMPLATE"; cat /fixture/value`} {
		suffix = append(suffix, shellQuote(arg))
	}
	command := containerCreateCommand(run, suffix)
	// Attach the same deadline to the generated shell command without changing
	// its argv; only this test-owned create is executed.
	command = exec.CommandContext(ctx, command.Path, command.Args[1:]...)
	out, err := command.Output()
	if err != nil {
		t.Fatal(commandError("container create", err))
	}
	id = strings.TrimSpace(string(out))
	if id == "" {
		t.Fatal("container engine returned no identity")
	}
	out, err = exec.CommandContext(ctx, "docker", "start", "-a", id).Output()
	if err != nil {
		t.Fatal(commandError("container start", err))
	}
	want := []byte(strings.Join([]string{containerOptionJSON, "1", "x", "", literal}, "\x00") + "\x00mounted bytes\n")
	if _, statErr := os.Lstat(sentinel); !os.IsNotExist(statErr) {
		t.Fatal("literal environment value was evaluated by the shell")
	}
	if !bytes.Equal(out, want) {
		t.Fatalf("container configuration roundtrip differs: got %q, want %q", out, want)
	}
}
