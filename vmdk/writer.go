// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package vmdk

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

const (
	// SparseMagicNumber is copied from
	// https://github.com/vmware/open-vmdk/blob/c977d2012d33cff8df3e809f20aa5df01e017f64/vmdk/vmware_vmdk.h#L48.
	SparseMagicNumber = 0x564d444b // "VMDK"

	// SparseVersionIncompatFlags is copied from
	// https://github.com/vmware/open-vmdk/blob/c977d2012d33cff8df3e809f20aa5df01e017f64/vmdk/vmware_vmdk.h#L51.
	SparseVersionIncompatFlags = 3

	// SparseFlagCompressed is copied from
	// https://github.com/vmware/open-vmdk/blob/c977d2012d33cff8df3e809f20aa5df01e017f64/vmdk/vmware_vmdk.h#L61.
	SparseFlagCompressed = (1 << 16)

	// SparseFlagEmbeddedLBA is copied from
	// https://github.com/vmware/open-vmdk/blob/c977d2012d33cff8df3e809f20aa5df01e017f64/vmdk/vmware_vmdk.h#L62.
	SparseFlagEmbeddedLBA = (1 << 17)

	// SparseFlagValidNewlineDetector is copied from
	// https://github.com/vmware/open-vmdk/blob/c977d2012d33cff8df3e809f20aa5df01e017f64/vmdk/vmware_vmdk.h#L57.
	SparseFlagValidNewlineDetector = (1 << 0)

	// GrainMarkerEndOfStream is copied from
	// https://github.com/vmware/open-vmdk/blob/c977d2012d33cff8df3e809f20aa5df01e017f64/vmdk/vmware_vmdk.h#L93.
	GrainMarkerEndOfStream = 0

	// GrainSize represents the size of a grain in sectors (128 sectors = 64KB).
	// This is the default grain size for VMDK sparse extents as specified in the VMDK specification.
	// The grain size must be a power of 2 and greater than 8 (4KB).
	GrainSize = 128

	// GrainDirectoryOffset is the offset in sectors where the Grain Directory starts.
	// This is calculated as 1 sector for the sparse extent header + 20 sectors for the embedded descriptor.
	GrainDirectoryOffset = 1 + 20
)

// SparseExtentHeaderOnDisk represents the on-disk structure of a sparse VMDK extent header.
// Corresponds to SparseExtentHeaderOnDisk in open-vmdk/vmdk/sparse.h
type SparseExtentHeaderOnDisk struct {
	MagicNumber        uint32
	Version            uint32
	Flags              uint32
	Capacity           uint64
	GrainSize          uint64
	DescriptorOffset   uint64
	DescriptorSize     uint64
	NumGTEsPerGT       uint32
	RGDOffset          uint64
	GDOffset           uint64
	Overhead           uint64
	UncleanShutdown    byte
	SingleEndLine      byte
	NonEndLine         byte
	DoubleEndLineChar1 byte
	DoubleEndLineChar2 byte
	CompressAlgorithm  uint16
	Padding            [433]byte
}

// SparseGrainLBAHeaderOnDisk represents the on-disk structure of a sparse grain LBA header.
// Corresponds to SparseGrainLBAHeaderOnDisk in open-vmdk/vmdk/sparse.h
type SparseGrainLBAHeaderOnDisk struct {
	Lba     uint64
	CmpSize uint32
}

// SparseMetaDataMarkerOnDisk represents the on-disk structure of a sparse metadata marker.
// Corresponds to SparseMetaDataMarkerOnDisk in open-vmdk/vmdk/sparse.h
type SparseMetaDataMarkerOnDisk struct {
	NumSectors uint64
	Size       uint32
	SectorType uint32
	Padding    [496]byte
}

// StreamOptimizedWriter creates stream-optimised VMDK files.
// It converts raw disk images into compressed VMDK format on-the-fly.
type StreamOptimizedWriter struct {
	file           *os.File
	header         SparseExtentHeaderOnDisk
	Capacity       int64
	GrainTables    []uint64 // stores individual grain table locations
	GrainDirectory []uint64 // stores grain table locations
}

// NewStreamOptimizedWriter creates a new StreamOptimizedWriter for writing VMDK files.
// It takes the output filename and the total capacity of the virtual disk in bytes as inputs
// It returns a pointer to the StreamOptimizedWriter and an error if any.
//
// Corresponds to StreamOptimized_Create() in open-vmdk/vmdk/sparse.c
func NewStreamOptimizedWriter(filename string, capacity int64) (*StreamOptimizedWriter, error) {
	outputFile, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to create output file: %v", err)
	}

	return &StreamOptimizedWriter{
		file:     outputFile,
		Capacity: capacity,
		header: SparseExtentHeaderOnDisk{
			MagicNumber:        SparseMagicNumber,
			Version:            3,
			Capacity:           uint64(capacity / SectorSize),
			DescriptorOffset:   1,
			GrainSize:          GrainSize,
			CompressAlgorithm:  1, // zlib
			SingleEndLine:      '\n',
			NonEndLine:         ' ',
			DoubleEndLineChar1: '\r',
			DoubleEndLineChar2: '\n',
			DescriptorSize:     20,
			NumGTEsPerGT:       512,
			Flags:              SparseFlagCompressed | SparseFlagEmbeddedLBA | SparseFlagValidNewlineDetector,
			Overhead:           1 + 20 + getGrainDirectorySize(capacity) + (getGrainTables(capacity) * 4),
			GDOffset:           GrainDirectoryOffset,
		},
		GrainDirectory: calculateGrainTableLocations(capacity),
		GrainTables:    make([]uint64, getGrains(capacity)),
	}, nil
}

// getGrains calculates the number of grains needed for the given capacity.
// The result is rounded up so a final partial grain is covered by a grain
// table entry instead of being silently dropped.
func getGrains(capacity int64) uint64 {
	grainSize := int64(GrainSize * SectorSize)
	if capacity <= 0 {
		return 0
	}
	return uint64((capacity-1)/grainSize + 1)
}

// getGrainTables calculates the number of Grain Tables needed for the given capacity.
func getGrainTables(capacity int64) uint64 {
	grains := getGrains(capacity)
	grainTables := grains / uint64(512)
	if grains%uint64(512) != 0 {
		grainTables++
	}
	return grainTables
}

// getGrainDirectorySize calculates the size of the Grain Directory in sectors.
func getGrainDirectorySize(capacity int64) uint64 {
	grainTables := getGrainTables(capacity)
	grainDirectory := (grainTables * 4) / 512
	if (grainTables*4)%512 != 0 {
		grainDirectory++
	}
	return grainDirectory
}

// Write reads raw disk data from the provided io.Reader, compresses it, and writes it to the VMDK file.
// It handles grain allocation, compression, and updates the Grain Tables and Grain Directory accordingly.
// Corresponds to StreamOptimizedPwrite() in open-vmdk/vmdk/sparse.c
func (w *StreamOptimizedWriter) Write(data io.Reader) error {
	err := w.writeMetadata()
	if err != nil {
		return fmt.Errorf("failed to write metadata: %v", err)
	}
	var compressedBuf bytes.Buffer
	zlibWriter := zlib.NewWriter(&compressedBuf)

	grainIndex := 0
	for grainIndex < len(w.GrainTables) {
		// Read a grain of raw data
		// Write the compressed data to the header
		// Track locations in GrainTables and GrainDirectory
		rawGrain := make([]byte, GrainSize*SectorSize)
		_, err := io.ReadFull(data, rawGrain)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.ErrUnexpectedEOF {
			return fmt.Errorf("failed to read grain data: %v", err)
		}
		// A short read (io.ErrUnexpectedEOF) is a partial final grain: the
		// buffer is allocated zeroed, so the bytes past the input are
		// already zero padding within the declared capacity.

		// Check if the grain is all zeroes
		if isZeroed(rawGrain) {
			// Mark grain as unallocated in GrainTables
			w.GrainTables[grainIndex] = 0
			grainIndex++
			continue
		}

		// Compress it using zlib
		compressedBuf.Reset()
		zlibWriter.Reset(&compressedBuf)
		_, err = zlibWriter.Write(rawGrain)
		if err != nil {
			w.Abort()
			return fmt.Errorf("failed to compress grain data: %v", err)
		}
		zlibWriter.Close()
		compressedData := compressedBuf.Bytes()

		// Grain records must start on a sector boundary: pad the gap left
		// by the previous variable-sized record with zeroes.
		nextOffset, err := w.file.Seek(0, io.SeekCurrent)
		if err != nil {
			w.Abort()
			return fmt.Errorf("failed to get current file offset: %v", err)
		}
		if aligned := alignToSector(nextOffset); aligned > nextOffset {
			padding := make([]byte, aligned-nextOffset)
			if _, err := w.file.Write(padding); err != nil {
				w.Abort()
				return fmt.Errorf("failed to write grain alignment padding: %v", err)
			}
			nextOffset = aligned
		}

		// Write SparseGrainLBAHeaderOnDisk
		sparseGrainLBAHeaderOnDisk := SparseGrainLBAHeaderOnDisk{
			Lba:     uint64(grainIndex) * uint64(GrainSize),
			CmpSize: uint32(len(compressedBuf.Bytes())),
		}
		err = binary.Write(w.file, binary.LittleEndian, &sparseGrainLBAHeaderOnDisk)
		if err != nil {
			w.Abort()
			return fmt.Errorf("failed to write SparseGrainLBAHeaderOnDisk: %v", err)
		}

		w.GrainTables[grainIndex] = uint64(nextOffset) / SectorSize

		// Write compressed data to file
		_, err = w.file.Write(compressedData)
		if err != nil {
			w.Abort()
			return fmt.Errorf("failed to write compressed grain data: %v", err)
		}

		grainIndex++
	}
	if grainIndex >= len(w.GrainTables) {
		// Every grain within the declared capacity has been written.
		// Trailing input would be silently truncated, so reject it.
		var probe [1]byte
		if _, err := io.ReadFull(data, probe[:]); err == nil {
			return fmt.Errorf("input exceeds disk capacity of %d bytes", w.Capacity)
		}
	}
	return nil
}

// isZeroed checks if the given byte slice is entirely zeroes.
// Corresponds to the zero-checking logic in StreamOptimizedPwrite() in open-vmdk/vmdk/sparse.c
func isZeroed(data []byte) bool {
	for _, b := range data {
		if b != 0 {
			return false
		}
	}
	return true
}

// calculateGrainTableLocations calculates the locations of Grain Tables in the VMDK file.
// Corresponds to the Grain Table location calculation in StreamOptimized_Create() in open-vmdk/vmdk/sparse.c
func calculateGrainTableLocations(capacity int64) []uint64 {
	// Calculate where the first grain table starts
	grainDirectoryStart := uint64(GrainDirectoryOffset)
	grainDirectorySize := getGrainDirectorySize(capacity)
	grainDirectoryEnd := grainDirectoryStart + grainDirectorySize - 1
	grainTableStart := grainDirectoryEnd + 1
	// Loop through the number of grain tables
	grainTableLocations := make([]uint64, getGrainTables(capacity))
	for i := uint64(0); i < getGrainTables(capacity); i++ {
		grainTableLocations[i] = grainTableStart + (i * 4)
	}

	// Add each grain table's location to the slice (each one is 4 sectors apart)
	return grainTableLocations
}

// writeMetadata writes the SparseExtentHeaderOnDisk, Descriptor, and reserves space for Grain Directory and Grain Tables.
func (w *StreamOptimizedWriter) writeMetadata() error {
	// Write SparseExtentHeaderOnDisk
	err := binary.Write(w.file, binary.LittleEndian, &w.header)
	if err != nil {
		return fmt.Errorf("failed to write SparseMetaDataMarkerOnDisk: %v", err)
	}

	// Write the descriptor
	extent := Extent{
		Size: w.Capacity / SectorSize,
		Type: "SPARSE",
	}
	descriptor := NewDescriptor(extent)
	buf := &bytes.Buffer{}
	err = descriptor.Write(buf)
	if err != nil {
		return fmt.Errorf("failed to write descriptor: %v", err)
	}
	n, err := w.file.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to write descriptor to file: %v", err)
	}
	// Add padding to make descriptor size 20 sectors
	paddingSize := (20 * SectorSize) - n
	if paddingSize > 0 {
		padding := make([]byte, paddingSize)
		_, err = w.file.Write(padding)
		if err != nil {
			return fmt.Errorf("failed to write padding to file: %v", err)
		}
	}

	// Reserve space for Grain Directory and Grain Tables
	overhead := make([]byte, (w.header.Overhead-1-20)*SectorSize)
	if _, err := w.file.Write(overhead); err != nil {
		w.Abort()
		return fmt.Errorf("failed to write overhead space: %v", err)
	}

	return nil
}

// alignToSector returns the next sector boundary at or after offset.
func alignToSector(offset int64) int64 {
	return (offset + int64(SectorSize) - 1) / int64(SectorSize) * int64(SectorSize)
}

// Close finalizes the VMDK file by writing Grain Tables, Grain Directory, and Metadata Marker.
// And marking the End of Stream.
// Corresponds to StreamOptimizedClose() in open-vmdk/vmdk/sparse.c
func (w *StreamOptimizedWriter) Close() error {
	// The end-of-stream marker must start on a sector boundary, just like
	// the grain records written by Write.
	offset, err := w.file.Seek(0, io.SeekCurrent)
	if err != nil {
		return fmt.Errorf("failed to get current file offset: %v", err)
	}
	if aligned := alignToSector(offset); aligned > offset {
		padding := make([]byte, aligned-offset)
		if _, err := w.file.Write(padding); err != nil {
			return fmt.Errorf("failed to write end-of-stream alignment padding: %v", err)
		}
	}

	// Write Metadata Marker
	metadataMarker := SparseMetaDataMarkerOnDisk{
		NumSectors: 0,
		Size:       0,
		SectorType: GrainMarkerEndOfStream,
	}
	err = binary.Write(w.file, binary.LittleEndian, &metadataMarker)
	if err != nil {
		return fmt.Errorf("failed to write SparseMetaDataMarkerOnDisk: %v", err)
	}

	// Write the Grain Directory to the reserved space we created in
	// writeMetadata. Entries are uint32 sector offsets, matching the 4
	// bytes per entry reserved by the layout calculations.
	// seek back to where the Grain Directory starts
	_, err = w.file.Seek(int64(w.header.GDOffset*SectorSize), io.SeekStart)
	if err != nil {
		return fmt.Errorf("failed to seek to Grain Directory: %v", err)
	}
	// Write Grain Directory
	for _, grainDirOffset := range w.GrainDirectory {
		err := binary.Write(w.file, binary.LittleEndian, uint32(grainDirOffset))
		if err != nil {
			return fmt.Errorf("failed to write Grain Directory: %v", err)
		}
	}
	// Write each Grain Table at its reserved location. GrainDirectory
	// holds the sector offset of every table; each table occupies
	// NumGTEsPerGT uint32 entries within its reserved sectors, the
	// remainder stays zero filled from writeMetadata.
	for i, grainTableOffset := range w.GrainDirectory {
		if _, err := w.file.Seek(int64(grainTableOffset*SectorSize), io.SeekStart); err != nil {
			return fmt.Errorf("failed to seek to Grain Table %d: %v", i, err)
		}
		start := i * int(w.header.NumGTEsPerGT)
		end := start + int(w.header.NumGTEsPerGT)
		if end > len(w.GrainTables) {
			end = len(w.GrainTables)
		}
		for _, grainTableEntry := range w.GrainTables[start:end] {
			err := binary.Write(w.file, binary.LittleEndian, uint32(grainTableEntry))
			if err != nil {
				return fmt.Errorf("failed to write Grain Table %d: %v", i, err)
			}
		}
	}

	return w.file.Close()
}

// Abort cleans up the VMDK file in case of an error during writing.
// It closes and removes the partially written file.
// Corresponds to StreamOptimizedAbort() in open-vmdk/vmdk/sparse.c
func (w *StreamOptimizedWriter) Abort() error {
	if err := w.file.Close(); err != nil {
		return fmt.Errorf("failed to close file during abort: %v", err)
	}
	if err := os.Remove(w.file.Name()); err != nil {
		return fmt.Errorf("failed to remove file during abort: %v", err)
	}
	return nil
}
