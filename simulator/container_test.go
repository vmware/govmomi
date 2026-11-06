// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestEncodeDMI(t *testing.T) {
	id := uuid.Must(uuid.Parse("423e36c8-723d-2b24-8590-0eb1d9180aa2"))

	tests := []struct {
		f      func(uuid.UUID) string
		expect string
	}{
		{productSerial, "VMware-42 3e 36 c8 72 3d 2b 24-85 90 0e b1 d9 18 0a a2"},
		{productUUID, "C8363E42-3D72-242B-8590-0EB1D9180AA2"},
	}

	for _, test := range tests {
		val := test.f(id)
		if val != test.expect {
			t.Errorf("%q != %q", val, test.expect)
		}
	}
}

// TestContainerRemove exercises the command boundary without a container engine.
func TestContainerRemove(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("requires a POSIX shell")
	}

	for _, tt := range []struct {
		name     string
		volumes  string
		networks string
		fail     string
	}{
		{name: "distinct resources", volumes: "volume-a\nvolume-b\n", networks: "network-a\nnetwork-b\n"},
		{name: "network only", networks: "network-a\n"},
		{name: "volume only", volumes: "volume-a\n"},
		{name: "no resources"},
		{name: "container removal failure", fail: "rm"},
		{name: "volume listing failure", fail: "volume ls"},
		{name: "volume removal failure", volumes: "volume-a\n", networks: "network-a\n", fail: "volume rm"},
		{name: "network listing failure", fail: "network ls"},
		{name: "network removal failure", volumes: "volume-a\n", networks: "network-a\n", fail: "network rm"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			log := filepath.Join(dir, "commands")
			script := `#!/bin/sh
case "$1" in
  rm) step=rm ;;
  *) step="$1 $2" ;;
esac
printf '%s' "$1" >> "$COMMAND_LOG"
shift
printf '\t%s' "$@" >> "$COMMAND_LOG"
printf '\n' >> "$COMMAND_LOG"
if [ "$step" = "$FAIL_COMMAND" ]; then exit 1; fi
case "$step" in
  "volume ls") printf '%s' "$VOLUMES" ;;
  "network ls") printf '%s' "$NETWORKS" ;;
esac
`
			if err := os.WriteFile(filepath.Join(dir, "docker"), []byte(script), 0700); err != nil {
				t.Fatal(err)
			}
			t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
			t.Setenv("COMMAND_LOG", log)
			t.Setenv("VOLUMES", tt.volumes)
			t.Setenv("NETWORKS", tt.networks)
			t.Setenv("FAIL_COMMAND", tt.fail)

			canceled := 0
			c := &container{name: "owned-container", id: "owned-id", cancelWatch: func() { canceled++ }}
			// Retain an unrelated watch to ensure removal touches only its own entry.
			other := &container{id: "other-id"}
			eventWatch.Lock()
			if eventWatch.watches == nil {
				eventWatch.watches = make(map[string]*container)
			}
			eventWatch.watches[c.id] = c
			eventWatch.watches[other.id] = other
			eventWatch.Unlock()
			t.Cleanup(func() {
				eventWatch.Lock()
				delete(eventWatch.watches, "owned-id")
				delete(eventWatch.watches, "other-id")
				eventWatch.Unlock()
			})

			checkCommands := func(containerFailure bool) {
				t.Helper()
				data, err := os.ReadFile(log)
				if err != nil {
					t.Fatal(err)
				}
				want := []string{"rm\t-v\t-f\towned-id"}
				if !containerFailure {
					want = append(want, "volume\tls\t-q\t--filter\tlabel=container=owned-container\t--filter\tlabel="+deleteWithContainer)
					if ids := strings.Fields(tt.volumes); len(ids) > 0 && tt.fail != "volume ls" {
						want = append(want, "volume\trm\t-f\t"+strings.Join(ids, "\t"))
					}
					want = append(want, "network\tls\t-q\t--filter\tlabel=container=owned-container\t--filter\tlabel="+deleteWithContainer)
					if ids := strings.Fields(tt.networks); len(ids) > 0 && tt.fail != "network ls" {
						want = append(want, "network\trm\t-f\t"+strings.Join(ids, "\t"))
					}
				}
				if got := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n"); !reflect.DeepEqual(got, want) {
					t.Fatalf("commands = %q, want %q", got, want)
				}
			}

			err := c.remove(nil)
			if (err != nil) != (tt.fail != "") {
				t.Fatalf("remove error = %v, failing command = %q", err, tt.fail)
			}
			checkCommands(tt.fail == "rm")
			if tt.fail != "" {
				if c.id != "owned-id" || canceled != 0 {
					t.Fatalf("failed cleanup lost ownership: id=%q, cancellations=%d", c.id, canceled)
				}
				eventWatch.Lock()
				watched := eventWatch.watches[c.id] == c
				eventWatch.Unlock()
				if !watched {
					t.Fatal("failed cleanup lost its watch")
				}
				t.Setenv("FAIL_COMMAND", "")
				if err := c.remove(nil); err != nil {
					t.Fatalf("retry: %v", err)
				}
			}
			if c.id != "" || canceled != 1 {
				t.Fatalf("successful cleanup: id=%q, cancellations=%d", c.id, canceled)
			}
			eventWatch.Lock()
			owned := eventWatch.watches["owned-id"]
			unrelated := eventWatch.watches[other.id]
			eventWatch.Unlock()
			if owned != nil || unrelated != other {
				t.Fatal("cleanup changed the wrong watches")
			}
			before, err := os.ReadFile(log)
			if err != nil {
				t.Fatal(err)
			}
			if err := c.remove(nil); err != nil {
				t.Fatal(err)
			}
			after, err := os.ReadFile(log)
			if err != nil {
				t.Fatal(err)
			}
			if string(before) != string(after) || canceled != 1 {
				t.Fatal("repeated cleanup ran commands or canceled again")
			}
		})
	}
}
