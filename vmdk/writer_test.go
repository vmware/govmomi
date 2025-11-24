// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmdk

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	diskSize           = int64(10485760) // 10 MB in KB
	numberOfGrains     = uint64(160)     // Each grain is 128 sectors × 512 bytes = 65536 bytes So 10485760 / 65536 = 160 grains
	grainTables        = uint64(1)       // Each grain table holds 512 entries So 160 / 512 = 0.3125, which rounds up to 1 grain table
	grainDirectorySize = uint64(1)       // 1 grain table × 4 bytes per entry = 4 bytes So 4 bytes / 512 bytes per sector = 0.0078, which rounds up to 1 sector
)

func TestGetGrains(t *testing.T) {
	result := getGrains(diskSize)
	assert.Equal(t, numberOfGrains, result)
}

func TestGetGrainTables(t *testing.T) {
	result := getGrainTables(diskSize)
	assert.Equal(t, grainTables, result)
}

func TestGetGrainDirectorySize(t *testing.T) {
	result := getGrainDirectorySize(diskSize)
	assert.Equal(t, grainDirectorySize, result)
}

func TestCalculateGrainTableLocations(t *testing.T) {
	expected := []uint64{22} // Overhead is 22 sectors for 10MB disk
	result := calculateGrainTableLocations(diskSize)

	assert.Equal(t, expected, result)
}

type zeroReader struct{}

func (z *zeroReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 0
	}
	return len(p), nil
}

func TestNewStreamOptimizedWriter(t *testing.T) {
	// case 1: valid parameters
	tmpDir := t.TempDir()
	imgPath := filepath.Join(tmpDir, "test.img")
	imgFile, err := os.Create(imgPath)
	require.NoError(t, err, "Failed to create test img file")

	_, err = io.CopyN(imgFile, rand.Reader, 1024*1024) // Create a 1MB random img file
	require.NoError(t, err, "Failed to create test img file")
	imgFile.Close()

	writer, err := NewStreamOptimizedWriter(filepath.Join(tmpDir, "output.vmdk"), 1024*1024)
	require.NoError(t, err, "Failed to create StreamOptimizedWriter")

	file, err := os.Open(imgPath)
	require.NoError(t, err, "Failed to open test img file")
	defer file.Close()

	err = writer.Write(file)
	require.NoError(t, err, "Failed to write data to VMDK")

	err = writer.Close()
	require.NoError(t, err, "Failed to close VMDK writer")

	outputInfo, err := os.Open(filepath.Join(tmpDir, "output.vmdk"))
	require.NoError(t, err, "Failed to open output VMDK file")
	defer outputInfo.Close()

	magicBytes := make([]byte, 4)
	_, err = io.ReadFull(outputInfo, magicBytes)
	require.NoError(t, err, "Failed to read output VMDK file")
	assert.Equal(t, uint32(SparseMagicNumber), binary.LittleEndian.Uint32(magicBytes), "Output VMDK magic number mismatch")

	// Case 2: Create a 1MB img file with zeroes
	imgPath2 := filepath.Join(tmpDir, "test2.img")
	imgFile2, err := os.Create(imgPath2)
	require.NoError(t, err, "Failed to create test img file")

	zeros := io.LimitReader(&zeroReader{}, 1024*1024)
	_, err = io.CopyN(imgFile2, zeros, 1024*1024) // Create a 1MB img file with zeroes
	require.NoError(t, err, "Failed to create test2 img file")
	imgFile2.Close()

	writer2, err := NewStreamOptimizedWriter(filepath.Join(tmpDir, "output2.vmdk"), 1024*1024)
	require.NoError(t, err, "Failed to create StreamOptimizedWriter")

	file2, err := os.Open(imgPath2)
	require.NoError(t, err, "Failed to open test img file")
	defer file2.Close()

	err = writer2.Write(file2)
	require.NoError(t, err, "Failed to write data to VMDK")

	err = writer2.Close()
	require.NoError(t, err, "Failed to close VMDK writer")

	outputInfo2, err := os.Open(filepath.Join(tmpDir, "output2.vmdk"))
	require.NoError(t, err, "Failed to open output VMDK file")
	defer outputInfo2.Close()

	magicBytes2 := make([]byte, 4)
	_, err = io.ReadFull(outputInfo2, magicBytes2)
	require.NoError(t, err, "Failed to read output VMDK file")
	assert.Equal(t, uint32(SparseMagicNumber), binary.LittleEndian.Uint32(magicBytes2), "Output VMDK magic number mismatch")

	// Verify sizes
	info, err := outputInfo.Stat()
	require.NoError(t, err, "Failed to stat output VMDK file")

	info2, err := outputInfo2.Stat()
	require.NoError(t, err, "Failed to stat output2 VMDK file")

	assert.Greater(t, info.Size(), info2.Size(), "Output VMDK size with random data should be greater than output2 VMDK size with zeros")

	// Case 3: Disk size is not a multiple of sector size
	imgPath3 := filepath.Join(tmpDir, "test3.img")
	imgFile3, err := os.Create(imgPath3)
	require.NoError(t, err, "Failed to create test img file")

	_, err = io.CopyN(imgFile3, rand.Reader, 10000) // Create a 100KB img file with zeroes
	require.NoError(t, err, "Failed to create test3 img file")
	imgFile3.Close()

	writer3, err := NewStreamOptimizedWriter(filepath.Join(tmpDir, "output3.vmdk"), 10000)
	require.NoError(t, err, "Failed to create StreamOptimizedWriter")

	file3, err := os.Open(imgPath3)
	require.NoError(t, err, "Failed to open test img file")
	defer file3.Close()

	err = writer3.Write(file3)
	require.NoError(t, err, "Failed to write data to VMDK")

	err = writer3.Close()
	require.NoError(t, err, "Failed to close VMDK writer")

	outputInfo3, err := os.Open(filepath.Join(tmpDir, "output3.vmdk"))
	require.NoError(t, err, "Failed to open output VMDK file")
	defer outputInfo3.Close()

	magicBytes3 := make([]byte, 4)
	_, err = io.ReadFull(outputInfo3, magicBytes3)
	require.NoError(t, err, "Failed to read output VMDK file")
	assert.Equal(t, uint32(SparseMagicNumber), binary.LittleEndian.Uint32(magicBytes3), "Output VMDK magic number mismatch")
}

func TestGetGrains_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		capacity int64
		want     uint64
	}{
		{"10MB", 10485760, 160},
		{"1MB", 1048576, 16},
		{"1GB", 1073741824, 16384},
		{"1 grain", 65536, 1},
		{"0 bytes", 0, 0},
		{"1TB", 1099511627776, 16777216},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getGrains(tt.capacity)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestGetGrainTables_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		capacity int64
		want     uint64
	}{
		{"10MB", 10485760, 1},
		{"1MB", 1048576, 1},
		{"1GB", 1073741824, 32},
		{"1 grain", 65536, 1},
		{"0 bytes", 0, 0},
		{"1TB", 1099511627776, 32768},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getGrainTables(tt.capacity)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestGetGrainDirectorySize_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		capacity int64
		want     uint64
	}{
		{"10MB", 10485760, 1},
		{"1MB", 1048576, 1},
		{"1GB", 1073741824, 1},
		{"1 grain", 65536, 1},
		{"0 bytes", 0, 0},
		{"1TB", 1099511627776, 256},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getGrainDirectorySize(tt.capacity)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestCalculateGrainTableLocations_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		capacity int64
		want     []uint64
	}{
		{"10MB", 10485760, []uint64{22}},
		{"1MB", 1048576, []uint64{22}},
		{"1GB", 1073741824, []uint64{22, 26, 30, 34, 38, 42, 46, 50, 54, 58, 62, 66, 70, 74, 78, 82, 86, 90, 94, 98, 102, 106, 110, 114, 118, 122, 126, 130, 134, 138, 142, 146}},
		{"1 grain", 65536, []uint64{22}},
		{"0 bytes", 0, []uint64{}},
		{"1TB", 1099511627776, func() []uint64 {
			locations := make([]uint64, 32768)
			for i := uint64(0); i < 32768; i++ {
				locations[i] = 277 + (i * 4)
			}
			return locations
		}()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateGrainTableLocations(tt.capacity)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestNewStreamOptimizedWriter_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		capacity int64
		wantErr  bool
	}{
		{"Valid case", "output1.vmdk", 1048576, false},
		{"Zero capacity", "output2.vmdk", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			writer, err := NewStreamOptimizedWriter(filepath.Join(tmpDir, tt.filename), tt.capacity)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				writer.Close()
			}
		})
	}
}

func TestIsZeroed(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want bool
	}{
		{"All zeroes", []byte{0, 0, 0, 0}, true},
		{"Some non-zeroes", []byte{0, 1, 0, 0}, false},
		{"All non-zeroes", []byte{1, 1, 1, 1}, false},
		{"Empty slice", []byte{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isZeroed(tt.data)
			assert.Equal(t, tt.want, result)
		})
	}
}
