// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package volume

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/vmware/govmomi/vim25/types"
)

func writeTempCSV(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "volumes.csv")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestRelocateEntriesFromFile_NoHeader(t *testing.T) {
	path := writeTempCSV(t, ""+
		"f75989dc-95b9-4db7-af96-8583f24bc59d\n"+
		"a1111111-95b9-4db7-af96-8583f24bc59d\n")

	var warn bytes.Buffer
	entries, err := relocateEntriesFromFile(path, &warn)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d: %+v", len(entries), entries)
	}
	if entries[0].VolumeID != "f75989dc-95b9-4db7-af96-8583f24bc59d" {
		t.Errorf("unexpected first volume ID: %q", entries[0].VolumeID)
	}
	if warn.Len() != 0 {
		t.Errorf("expected no warning, got %q", warn.String())
	}
}

func TestRelocateEntriesFromFile_WithDatastore(t *testing.T) {
	path := writeTempCSV(t, ""+
		"f75989dc-95b9-4db7-af96-8583f24bc59d,vsanDatastore\n"+
		"a1111111-95b9-4db7-af96-8583f24bc59d, otherDatastore \n")

	var warn bytes.Buffer
	entries, err := relocateEntriesFromFile(path, &warn)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d: %+v", len(entries), entries)
	}
	if entries[0].Datastore != "vsanDatastore" {
		t.Errorf("unexpected datastore for entry 0: %q", entries[0].Datastore)
	}
	if entries[1].Datastore != "otherDatastore" {
		t.Errorf("unexpected datastore for entry 1 (should be trimmed): %q", entries[1].Datastore)
	}
}

func TestRelocateEntriesFromFile_HeaderRowSkippedAndWarned(t *testing.T) {
	path := writeTempCSV(t, ""+
		"volumeId,datastore\n"+
		"f75989dc-95b9-4db7-af96-8583f24bc59d,vsanDatastore\n")

	var warn bytes.Buffer
	entries, err := relocateEntriesFromFile(path, &warn)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d: %+v", len(entries), entries)
	}
	if warn.Len() == 0 {
		t.Fatal("expected a warning noting the header row was skipped, got none")
	}
	if !strings.Contains(warn.String(), "volumeId") {
		t.Errorf("expected warning to mention the skipped row, got %q", warn.String())
	}
}

// TestRelocateEntriesFromFile_NonUUIDFirstRowWarns documents the
// header-detection heuristic's failure mode: a real, non-UUID volume ID in
// row 1 is still (mis)treated as a header and dropped, but the caller must
// now be told about it via warn instead of the row silently vanishing.
func TestRelocateEntriesFromFile_NonUUIDFirstRowWarns(t *testing.T) {
	path := writeTempCSV(t, ""+
		"not-a-uuid-volume-id\n"+
		"a1111111-95b9-4db7-af96-8583f24bc59d\n")

	var warn bytes.Buffer
	entries, err := relocateEntriesFromFile(path, &warn)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry (row 1 misdetected as header), got %d: %+v", len(entries), entries)
	}
	if !strings.Contains(warn.String(), "not-a-uuid-volume-id") {
		t.Fatalf("expected a warning naming the skipped row, got %q", warn.String())
	}
}

func TestRelocateEntriesFromFile_Empty(t *testing.T) {
	path := writeTempCSV(t, "\n\n")

	var warn bytes.Buffer
	_, err := relocateEntriesFromFile(path, &warn)
	if err == nil {
		t.Fatal("expected an error for a file with no volume IDs")
	}
}

func TestRelocateEntriesFromFile_MissingFile(t *testing.T) {
	var warn bytes.Buffer
	_, err := relocateEntriesFromFile(filepath.Join(t.TempDir(), "missing.csv"), &warn)
	if err == nil {
		t.Fatal("expected an error for a missing file")
	}
}

func TestFaultOpID(t *testing.T) {
	fault := &types.MethodFault{
		FaultMessage: []types.LocalizableMessage{
			{
				Key: "some.key",
				Arg: []types.KeyAnyValue{
					{Key: "other", Value: "ignored"},
					{Key: "opId", Value: "abc-123"},
				},
			},
		},
	}

	if id := faultOpID(fault); id != "abc-123" {
		t.Errorf("expected opId abc-123, got %q", id)
	}
}

func TestFaultOpID_NilFault(t *testing.T) {
	if id := faultOpID(nil); id != "" {
		t.Errorf("expected empty opId for nil fault, got %q", id)
	}
}

func TestFaultOpID_NoOpIDArg(t *testing.T) {
	fault := &types.MethodFault{
		FaultMessage: []types.LocalizableMessage{
			{Key: "some.key", Arg: []types.KeyAnyValue{{Key: "other", Value: "ignored"}}},
		},
	}

	if id := faultOpID(fault); id != "" {
		t.Errorf("expected empty opId, got %q", id)
	}
}

func TestRelocateResultCounts(t *testing.T) {
	out := &relocateResult{
		Results: []relocateVolumeResult{
			{VolumeID: "a", Status: "OK"},
			{VolumeID: "b", Status: "FAILED", Error: "boom"},
			{VolumeID: "c", Status: "OK"},
		},
	}

	succeeded, failed := out.counts()
	if succeeded != 2 || failed != 1 {
		t.Errorf("expected 2 succeeded, 1 failed, got %d succeeded, %d failed", succeeded, failed)
	}
}
