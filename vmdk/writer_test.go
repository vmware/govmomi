package vmdk

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"os"
	"path/filepath"
	"testing"
)

var (
	diskSize           = int64(10485760) // 10 MB in KB
	numberOfGrains     = uint64(160)     // Each grain is 128 sectors × 512 bytes = 65536 bytes So 10485760 / 65536 = 160 grains
	grainTables        = uint64(1)       // Each grain table holds 512 entries So 160 / 512 = 0.3125, which rounds up to 1 grain table
	grainDirectorySize = uint64(1)       // 1 grain table × 4 bytes per entry = 4 bytes So 4 bytes / 512 bytes per sector = 0.0078, which rounds up to 1 sector
)

func TestGetGrains(t *testing.T) {
	result := getGrains(diskSize)

	if result != numberOfGrains {
		t.Errorf("getGrains(%d) = %d; want %d", diskSize, result, numberOfGrains)
	}

}

func TestGetGrainTables(t *testing.T) {
	result := getGrainTables(diskSize)
	if result != grainTables {
		t.Errorf("getGrainTables(%d) = %d; want %d", diskSize, result, grainTables)
	}
}

func TestGetGrainDirectorySize(t *testing.T) {
	result := getGrainDirectorySize(diskSize)
	if result != grainDirectorySize {
		t.Errorf("getGrainDirectorySize(%d) = %d; want %d", diskSize, result, 1)
	}
}

func TestCalculateGrainTableLocations(t *testing.T) {
	expected := []uint64{22} // Overhead is 22 sectors for 10MB disk
	result := calculateGrainTableLocations(diskSize)

	if len(result) != len(expected) {
		t.Fatalf("calculateGrainTableLocations(%d) returned %d entries; want %d", diskSize, len(result), len(expected))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("calculateGrainTableLocations(%d)[%d] = %d; want %d", diskSize, i, result[i], v)
		}
	}
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
	if err != nil {
		t.Fatalf("Failed to create test img file: %v", err)
	}
	_, err = io.CopyN(imgFile, rand.Reader, 1024*1024) // Create a 1MB random img file
	if err != nil {
		t.Fatalf("Failed to create test img file: %v", err)
	}
	imgFile.Close()

	writer, err := NewStreamOptimizedWriter(filepath.Join(tmpDir, "output.vmdk"), 1024*1024)
	if err != nil {
		t.Fatalf("Failed to create StreamOptimizedWriter: %v", err)
	}
	file, err := os.Open(imgPath)
	if err != nil {
		t.Fatalf("Failed to open test img file: %v", err)
	}
	defer file.Close()
	if err := writer.Write(file); err != nil {
		t.Fatalf("Failed to write data to VMDK: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("Failed to close VMDK writer: %v", err)
	}
	outputInfo, err := os.Open(filepath.Join(tmpDir, "output.vmdk"))
	if err != nil {
		t.Fatalf("Failed to open output VMDK file: %v", err)
	}
	defer outputInfo.Close()
	magicBytes := make([]byte, 4)
	_, err = io.ReadFull(outputInfo, magicBytes)
	if err != nil {
		t.Fatalf("Failed to read output VMDK file: %v", err)
	}
	if binary.LittleEndian.Uint32(magicBytes) != MagicNumber {
		t.Errorf("Output VMDK magic number = %x; want %x", binary.LittleEndian.Uint32(magicBytes), MagicNumber)
	}

	// Case 2: Create a 1MB img file with zeroes
	imgPath2 := filepath.Join(tmpDir, "test2.img")
	imgFile2, err := os.Create(imgPath2)
	if err != nil {
		t.Fatalf("Failed to create test img file: %v", err)
	}
	zeros := io.LimitReader(&zeroReader{}, 1024*1024)
	_, err = io.CopyN(imgFile2, zeros, 1024*1024) // Create a 1MB img file with zeroes
	if err != nil {
		t.Fatalf("Failed to create test2 img file: %v", err)
	}
	imgFile2.Close()

	writer2, err := NewStreamOptimizedWriter(filepath.Join(tmpDir, "output2.vmdk"), 1024*1024)
	if err != nil {
		t.Fatalf("Failed to create StreamOptimizedWriter: %v", err)
	}
	file2, err := os.Open(imgPath2)
	if err != nil {
		t.Fatalf("Failed to open test img file: %v", err)
	}
	defer file.Close()
	if err := writer2.Write(file2); err != nil {
		t.Fatalf("Failed to write data to VMDK: %v", err)
	}
	if err := writer2.Close(); err != nil {
		t.Fatalf("Failed to close VMDK writer: %v", err)
	}
	outputInfo2, err := os.Open(filepath.Join(tmpDir, "output2.vmdk"))
	if err != nil {
		t.Fatalf("Failed to open output VMDK file: %v", err)
	}
	defer outputInfo2.Close()
	magicBytes2 := make([]byte, 4)
	_, err = io.ReadFull(outputInfo2, magicBytes2)
	if err != nil {
		t.Fatalf("Failed to read output VMDK file: %v", err)
	}
	if binary.LittleEndian.Uint32(magicBytes2) != MagicNumber {
		t.Errorf("Output VMDK magic number = %x; want %x", binary.LittleEndian.Uint32(magicBytes2), MagicNumber)
	}

	// Verify sizes
	info, err := outputInfo.Stat()
	if err != nil {
		t.Fatalf("Failed to stat output VMDK file: %v", err)
	}
	info2, err := outputInfo2.Stat()
	if err != nil {
		t.Fatalf("Failed to stat output2 VMDK file: %v", err)
	}
	if info.Size() <= info2.Size() {
		t.Errorf("Output VMDK size with random data %d should be greater than output2 VMDK size with zeros %d", info.Size(), info2.Size())
	}

	// Case 3: Disk size is not a multiple of sector size
	imgPath3 := filepath.Join(tmpDir, "test3.img")
	imgFile3, err := os.Create(imgPath3)
	if err != nil {
		t.Fatalf("Failed to create test img file: %v", err)
	}
	_, err = io.CopyN(imgFile3, rand.Reader, 10000) // Create a 100KB img file with zeroes
	if err != nil {
		t.Fatalf("Failed to create test2 img file: %v", err)
	}
	imgFile3.Close()

	writer3, err := NewStreamOptimizedWriter(filepath.Join(tmpDir, "output3.vmdk"), 10000)
	if err != nil {
		t.Fatalf("Failed to create StreamOptimizedWriter: %v", err)
	}
	file3, err := os.Open(imgPath3)
	if err != nil {
		t.Fatalf("Failed to open test img file: %v", err)
	}
	defer file.Close()
	if err := writer3.Write(file3); err != nil {
		t.Fatalf("Failed to write data to VMDK: %v", err)
	}
	if err := writer3.Close(); err != nil {
		t.Fatalf("Failed to close VMDK writer: %v", err)
	}
	outputInfo3, err := os.Open(filepath.Join(tmpDir, "output3.vmdk"))
	if err != nil {
		t.Fatalf("Failed to open output VMDK file: %v", err)
	}
	defer outputInfo3.Close()
	magicBytes3 := make([]byte, 4)
	_, err = io.ReadFull(outputInfo3, magicBytes3)
	if err != nil {
		t.Fatalf("Failed to read output VMDK file: %v", err)
	}
	if binary.LittleEndian.Uint32(magicBytes3) != MagicNumber {
		t.Errorf("Output VMDK magic number = %x; want %x", binary.LittleEndian.Uint32(magicBytes3), MagicNumber)
	}

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
			if result != tt.want {
				t.Errorf("getGrains(%d) = %d; want %d", tt.capacity, result, tt.want)
			}
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
			if result != tt.want {
				t.Errorf("getGrainTables(%d) = %d; want %d", tt.capacity, result, tt.want)
			}
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
			if result != tt.want {
				t.Errorf("getGrainDirectorySize(%d) = %d; want %d", tt.capacity, result, tt.want)
			}
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
			if len(result) != len(tt.want) {
				t.Fatalf("calculateGrainTableLocations(%d) returned %d entries; want %d", tt.capacity, len(result), len(tt.want))
			}
			for i, v := range tt.want {
				if result[i] != v {
					t.Errorf("calculateGrainTableLocations(%d)[%d] = %d; want %d", tt.capacity, i, result[i], v)
				}
			}
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
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewStreamOptimizedWriter(%s, %d) error = %v; wantErr %v", tt.filename, tt.capacity, err, tt.wantErr)
			}
			if err == nil {
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
			if result != tt.want {
				t.Errorf("isZeroed(%v) = %v; want %v", tt.data, result, tt.want)
			}
		})
	}
}
