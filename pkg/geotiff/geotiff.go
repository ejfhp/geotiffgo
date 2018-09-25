package geotiff

// https://github.com/geotiffjs/geotiff.js/

import (
	"fmt"
	"os"
)

// GeoTIFF represent a GeoTIFF file
type GeoTIFF struct {
	source         *os.File //The datasource to read from.
	littleEndian   bool     //Whether the image uses little endian.
	bigTiff        bool     //Whether the image uses bigTIFF conventions.
	firstIFDOffset uint     //The numeric byte-offset from the start of the image to the first IFD.
}

// FromFile instantiate a new GeoTIFF from a file
func FromFile(filePath string) (*GeoTIFF, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open GeoTIFF file %v due to %v", filePath, err)
	}
	headerData := make([]byte, 1024, 1024)
	_, err = file.Read(headerData)
	if err != nil {
		return nil, fmt.Errorf("cannot read from GeoTIFF file %v due to %v", filePath, err)
	}
	//Byte Order Mark
	dataView := DataView(headerData)
	bom := dataView.Uint16(0, false)
	var littleEndian bool
	if bom == 0x4949 {
		littleEndian = true
	} else if bom == 0x4d4d {
		littleEndian = false
	} else {
		return nil, fmt.Errorf("invalid Byte Order Value - BOM")
	}
	magicNumber := dataView.Uint16(2, littleEndian)
	var bigTiff bool
	if magicNumber == 42 {
		bigTiff = false
	} else if magicNumber == 43 {
		bigTiff = true
		offsetByteSize := dataView.Uint16(4, littleEndian)
		if offsetByteSize != 8 {
			return nil, fmt.Errorf("unsupported offset byte-size %v instead of 8", offsetByteSize)
		}
	} else {
		return nil, fmt.Errorf("invalid magic number: %v", magicNumber)
	}
	var firstIFDOffset uint
	if bigTiff {
		firstIFDOffset = uint(dataView.Uint64(8, littleEndian))
	} else {
		firstIFDOffset = uint(dataView.Uint32(4, littleEndian))
	}
	geoTiff := GeoTIFF{file, littleEndian, bigTiff, firstIFDOffset}
	return &geoTiff, nil
}

// GetSlice returns a DataSlice of data from the GeoTIFF file
func (g *GeoTIFF) GetSlice(offset uint, size uint) (*DataSlice, error) {
	buffer := make([]byte, size)
	_, err := g.source.ReadAt(buffer, int64(offset))
	if err != nil {
		return nil, fmt.Errorf("cannot read data from file due to %v", err)
	}
	dataSlice := NewDataSlice(buffer, offset, g.littleEndian, g.bigTiff)
	return dataSlice, nil
}
