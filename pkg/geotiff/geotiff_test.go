package geotiff

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestFromFile(t *testing.T) {
	testTiffPath := filepath.Join("testdata", "very_small.tif")
	gt, err := FromFile(testTiffPath)
	b := make([]byte, 1, 1)
	n, err := gt.source.Read(b)
	if err != nil {
		t.Errorf("source is closed? %v", err)
	} else {
		fmt.Printf("Read 1 byte from source: %d\n", n)
	}
	if err != nil {
		t.Errorf("failed due to %v", err)
	}
	if gt == nil {
		t.Errorf("failed because image is nil")
	}
	if !gt.littleEndian {
		t.Errorf("failed because tiff should be LittleEndian")
	}
	fmt.Printf("Tiff is LittleEndian: %t\n", gt.littleEndian)
	fmt.Printf("Tiff firsta page offset: %d\n", gt.firstIFDOffset)
}

func TestGetSlice(t *testing.T) {
	testTiffPath := filepath.Join("testdata", "very_small.tif")
	gt, err := FromFile(testTiffPath)
	dataSlice, err := gt.GetSlice(10, 10)
	if err != nil {
		t.Errorf("failed due to %v", err)
	}
	if dataSlice.Offset() != 10 {
		t.Errorf("wrong offset")
	}
	if dataSlice.SliceTop() != 20 {
		t.Errorf("wrong offset")
	}
}
