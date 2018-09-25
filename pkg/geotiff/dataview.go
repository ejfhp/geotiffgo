package geotiff

import (
	"encoding/binary"
	"math"
)

//DataView allow to easily read data
type DataView []byte

//Uint8 read and return a uint8 from the underlying data starting at the given offset
func (d DataView) Uint8(offset uint, littleEndian bool) uint8 {
	return uint8(d[offset])
}

//Int8 read and return a int8 from the underlying data starting at the given offset
func (d DataView) Int8(offset uint, littleEndian bool) int8 {
	return int8(d[offset])
}

//Uint16 read and return a uint8 from the underlying data starting at the given offset
func (d DataView) Uint16(offset uint, littleEndian bool) uint16 {
	var decoding binary.ByteOrder
	if littleEndian {
		decoding = binary.LittleEndian
	} else {
		decoding = binary.BigEndian
	}
	return decoding.Uint16(d[offset:])
}

//Int16 read and return a uint8 from the underlying data starting at the given offset
func (d DataView) Int16(offset uint, littleEndian bool) int16 {
	var decoding binary.ByteOrder
	if littleEndian {
		decoding = binary.LittleEndian
	} else {
		decoding = binary.BigEndian
	}
	return int16(decoding.Uint16(d[offset:]))
}

//Uint32 read and return a uint32 from the underlying data starting at the given offset
func (d DataView) Uint32(offset uint, littleEndian bool) uint32 {
	var decoding binary.ByteOrder
	if littleEndian {
		decoding = binary.LittleEndian
	} else {
		decoding = binary.BigEndian
	}
	return decoding.Uint32(d[offset:])
}

//Int32 read and return a int32 from the underlying data starting at the given offset
func (d DataView) Int32(offset uint, littleEndian bool) int32 {
	var decoding binary.ByteOrder
	if littleEndian {
		decoding = binary.LittleEndian
	} else {
		decoding = binary.BigEndian
	}
	return int32(decoding.Uint32(d[offset:]))
}

//Uint64 read and return a uint64 from the underlying data starting at the given offset
func (d DataView) Uint64(offset uint, littleEndian bool) uint64 {
	var decoding binary.ByteOrder
	if littleEndian {
		decoding = binary.LittleEndian
	} else {
		decoding = binary.BigEndian
	}
	return decoding.Uint64(d[offset:])
}

//Int64 read and return a uint64 from the underlying data starting at the given offset
func (d DataView) Int64(offset uint, littleEndian bool) int64 {
	var decoding binary.ByteOrder
	if littleEndian {
		decoding = binary.LittleEndian
	} else {
		decoding = binary.BigEndian
	}
	return int64(decoding.Uint64(d[offset:]))
}

//Float32 read and return a float32 from the underlying data starting at the given offset
func (d DataView) Float32(offset uint, littleEndian bool) float32 {
	var decoding binary.ByteOrder
	if littleEndian {
		decoding = binary.LittleEndian
	} else {
		decoding = binary.BigEndian
	}
	bits := decoding.Uint32(d[offset:])
	return math.Float32frombits(bits)
}

//Float64 read and return a float64 from the underlying data starting at the given offset
func (d DataView) Float64(offset uint, littleEndian bool) float64 {
	var decoding binary.ByteOrder
	if littleEndian {
		decoding = binary.LittleEndian
	} else {
		decoding = binary.BigEndian
	}
	bits := decoding.Uint64(d[offset:])
	return math.Float64frombits(bits)
}
