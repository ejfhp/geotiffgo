package geotiff

// DataSlice contains a 'slice' of data
type DataSlice struct {
	dataView     DataView
	offset       uint
	littleEndian bool
	bigTiff      bool
}

// NewDataSlice create a new instance of dataSlice
func NewDataSlice(dataView DataView, offset uint, littleEndian bool, bigTiff bool) *DataSlice {
	return &DataSlice{dataView, offset, littleEndian, bigTiff}
}

// Offset returns the offset value
func (d *DataSlice) Offset() uint {
	return d.offset
}

// SliceTop returns the top offset of the DataSlice
func (d *DataSlice) SliceTop() uint {
	return d.offset + uint(len(d.dataView))
}

// LittleEndian returns the littleEndian value
func (d *DataSlice) LittleEndian() bool {
	return d.littleEndian
}

// BigTiff returns the bigTiff value
func (d *DataSlice) BigTiff() bool {
	return d.bigTiff
}

// Buffer returns the encapsulated DataView
func (d *DataSlice) Buffer() DataView {
	return d.dataView
}

// Covers returns true id the slice covers the given interval
func (d *DataSlice) Covers(offset uint, length uint) bool {
	return d.offset <= offset && d.SliceTop() >= offset+length
}

// ReadUint8 read a uint8 at offset
func (d *DataSlice) ReadUint8(offset uint) uint8 {
	return d.dataView.Uint8(offset-d.offset, d.littleEndian)
}

// ReadInt8 read a int8 at offset
func (d *DataSlice) ReadInt8(offset uint) int8 {
	return d.dataView.Int8(offset-d.offset, d.littleEndian)
}

// ReadUint16 read a uint16 at offset
func (d *DataSlice) ReadUint16(offset uint) uint16 {
	return d.dataView.Uint16(offset-d.offset, d.littleEndian)
}

// ReadInt16 read a int16 at offset
func (d *DataSlice) ReadInt16(offset uint) int16 {
	return d.dataView.Int16(offset-d.offset, d.littleEndian)
}

// ReadUint32 read a uint32 at offset
func (d *DataSlice) ReadUint32(offset uint) uint32 {
	return d.dataView.Uint32(offset-d.offset, d.littleEndian)
}

// ReadInt32 read a int32 at offset
func (d *DataSlice) ReadInt32(offset uint) int32 {
	return d.dataView.Int32(offset-d.offset, d.littleEndian)
}

// ReadUint64 read a uint64 at offset
func (d *DataSlice) ReadUint64(offset uint) uint64 {
	return d.dataView.Uint64(offset-d.offset, d.littleEndian)
}

// ReadInt64 read a int64 at offset
func (d *DataSlice) ReadInt64(offset uint) int64 {
	return d.dataView.Int64(offset-d.offset, d.littleEndian)
}

// ReadFloat32 read a float32 at offset
func (d *DataSlice) ReadFloat32(offset uint) float32 {
	return d.dataView.Float32(offset, d.littleEndian)
}

// ReadFloat64 read a float64 at offset
func (d *DataSlice) ReadFloat64(offset uint) float64 {
	return d.dataView.Float64(offset, d.littleEndian)
}

// ReadOffset returns uint64 in case of bigTiff otherwise uint32
func (d *DataSlice) ReadOffset(offset uint) uint {
	if d.bigTiff {
		return uint(d.ReadUint64(offset))
	}
	return uint(d.ReadUint32(offset))
}
