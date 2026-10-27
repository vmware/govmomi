// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmdk

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// readStreamOptimized decodes the stream-optimized VMDK at path and returns
// the disk contents it represents, up to capacity bytes.
func readStreamOptimized(t *testing.T, path string, capacity int64) []byte {
	t.Helper()

	f, err := os.Open(path)
	require.NoError(t, err)
	defer f.Close()

	var header SparseExtentHeaderOnDisk
	require.NoError(t, binary.Read(f, binary.LittleEndian, &header))
	assert.Equal(t, uint32(SparseMagicNumber), header.MagicNumber)

	grains := getGrains(capacity)
	disk := make([]byte, grains*GrainSize*SectorSize)

	// Grain Directory: one uint32 sector offset per grain table.
	_, err = f.Seek(int64(header.GDOffset*SectorSize), io.SeekStart)
	require.NoError(t, err)
	grainDirectory := make([]uint32, getGrainTables(capacity))
	require.NoError(t, binary.Read(f, binary.LittleEndian, &grainDirectory))

	for i, tableOffset := range grainDirectory {
		_, err = f.Seek(int64(tableOffset)*SectorSize, io.SeekStart)
		require.NoError(t, err)
		entries := make([]uint32, header.NumGTEsPerGT)
		require.NoError(t, binary.Read(f, binary.LittleEndian, &entries))
		for j, entry := range entries {
			grainIndex := uint64(i)*uint64(header.NumGTEsPerGT) + uint64(j)
			if grainIndex >= grains {
				break
			}
			if entry == 0 {
				continue // unallocated grain, reads back as zeroes
			}
			_, err = f.Seek(int64(entry)*SectorSize, io.SeekStart)
			require.NoError(t, err)
			var grainHeader SparseGrainLBAHeaderOnDisk
			require.NoError(t, binary.Read(f, binary.LittleEndian, &grainHeader))
			assert.Equal(t, grainIndex*GrainSize, grainHeader.Lba)
			compressed := make([]byte, grainHeader.CmpSize)
			_, err = io.ReadFull(f, compressed)
			require.NoError(t, err)
			zr, err := zlib.NewReader(bytes.NewReader(compressed))
			require.NoError(t, err)
			grain := disk[grainIndex*GrainSize*SectorSize:(grainIndex+1)*GrainSize*SectorSize]
			_, err = io.ReadFull(zr, grain)
			require.NoError(t, err)
			require.NoError(t, zr.Close())
		}
	}

	return disk[:capacity]
}

// partialFinalGrainData returns capacity bytes where only the final,
// partial grain holds non-zero data.
func partialFinalGrainData(capacity int64) []byte {
	data := make([]byte, capacity)
	grainSize := int64(GrainSize * SectorSize)
	for i := (capacity - 1) / grainSize * grainSize; i < capacity; i++ {
		data[i] = 'A'
	}
	return data
}

// multiTableGrainData returns capacity bytes with non-zero grains spread
// across grain table boundaries (grain indexes 0, 511, 512 and 1023).
func multiTableGrainData(capacity int64) []byte {
	data := make([]byte, capacity)
	grainSize := int64(GrainSize * SectorSize)
	for _, g := range []int64{0, 511, 512, 1023} {
		start := g * grainSize
		for i := start; i < start+grainSize; i++ {
			data[i] = byte('A' + g%26)
		}
	}
	return data
}

func TestStreamOptimizedWriterRoundTrip(t *testing.T) {
	tests := []struct {
		name     string
		capacity int64
		data     []byte
	}{
		{"all zero", 1024 * 1024, make([]byte, 1024*1024)},
		{"partial final grain", 1024*1024 + 512, partialFinalGrainData(1024*1024 + 512)},
		{"multiple grain tables", 64 * 1024 * 1024, multiTableGrainData(64 * 1024 * 1024)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()

			imgPath := filepath.Join(tmpDir, "input.raw")
			require.NoError(t, os.WriteFile(imgPath, tt.data, 0o600))

			vmdkPath := filepath.Join(tmpDir, "output.vmdk")
			writer, err := NewStreamOptimizedWriter(vmdkPath, tt.capacity)
			require.NoError(t, err)

			img, err := os.Open(imgPath)
			require.NoError(t, err)
			require.NoError(t, writer.Write(img))
			require.NoError(t, img.Close())
			require.NoError(t, writer.Close())

			assert.Equal(t, tt.data, readStreamOptimized(t, vmdkPath, tt.capacity))
		})
	}
}

func TestStreamOptimizedWriterFinalSectorIsPreserved(t *testing.T) {
	// Two disks differing only in their final sector must produce
	// different VMDK files: the final sector used to be discarded,
	// which made both conversions byte identical.
	capacity := int64(1024*1024 + 512)

	writeDisk := func(t *testing.T, last byte) string {
		t.Helper()
		tmpDir := t.TempDir()
		data := make([]byte, capacity)
		data[capacity-1] = last
		imgPath := filepath.Join(tmpDir, "input.raw")
		require.NoError(t, os.WriteFile(imgPath, data, 0o600))
		vmdkPath := filepath.Join(tmpDir, "output.vmdk")
		writer, err := NewStreamOptimizedWriter(vmdkPath, capacity)
		require.NoError(t, err)
		img, err := os.Open(imgPath)
		require.NoError(t, err)
		require.NoError(t, writer.Write(img))
		require.NoError(t, img.Close())
		require.NoError(t, writer.Close())
		return vmdkPath
	}

	a := writeDisk(t, 'A')
	b := writeDisk(t, 'B')

	aData, err := os.ReadFile(a)
	require.NoError(t, err)
	bData, err := os.ReadFile(b)
	require.NoError(t, err)
	assert.NotEqual(t, aData, bData)

	// And the decoded final sector must match the input.
	decoded := readStreamOptimized(t, a, capacity)
	assert.Equal(t, byte('A'), decoded[capacity-1])
}

func TestStreamOptimizedWriterRejectsExcessInput(t *testing.T) {
	tmpDir := t.TempDir()
	writer, err := NewStreamOptimizedWriter(filepath.Join(tmpDir, "output.vmdk"), 1024*1024)
	require.NoError(t, err)

	data := make([]byte, 1024*1024+1)
	err = writer.Write(bytes.NewReader(data))
	assert.Error(t, err)
	require.NoError(t, writer.Abort())
}
