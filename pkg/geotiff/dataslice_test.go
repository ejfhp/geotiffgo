package geotiff

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 0, false, false)
	if ds == nil {
		t.Errorf("dataslice is nil")
	}
}

func TestOffset(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 0, false, false)
	offset := ds.Offset()
	if offset != ds.offset {
		t.Errorf("offset is not correct")
	}
}

func TestSliceTop(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 3, false, false)
	top := ds.SliceTop()
	if top != 7 {
		t.Errorf("SliceTop is not correct")
	}
}

func TestLittleEndian(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 0, true, false)
	le := ds.LittleEndian()
	if le != ds.littleEndian {
		t.Errorf("LittleEndian is not correct")
	}
}

func TestBigTiff(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 0, true, true)
	le := ds.BigTiff()
	if le != ds.bigTiff {
		t.Errorf("BigTiff is not correct")
	}
}
func TestBuffer(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 2, true, true)
	buf := ds.Buffer()
	if !reflect.DeepEqual(buf, dw) {
		t.Errorf("Buffer is not correct")
	}
}

func TestCovers(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 2, true, true)
	cover := ds.Covers(1, 4)
	if cover == true {
		t.Errorf("Cover is not correct")
	}
	cover = ds.Covers(3, 3)
	if cover != true {
		t.Errorf("Cover is not correct")
	}
}

func TestReadUint8(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 1, false, true)
	ui := ds.ReadUint8(3)
	if ui != 19 {
		t.Errorf("Read is not correct: %d\n", ui)
	}
}

func TestReadInt8(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 1, false, true)
	ui := ds.ReadInt8(3)
	if ui != 19 {
		t.Errorf("Read is not correct: %d\n", ui)
	}
}

func TestReadUint16(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 0, false, true)
	ui := ds.ReadUint16(1)
	if ui != 4115 {
		t.Errorf("Read is not correct: %d\n", ui)
	}
}

func TestReadInt16(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 0, true, true)
	ui := ds.ReadInt16(1)
	if ui != 4880 {
		t.Errorf("Read is not correct: %d\n", ui)
	}
}

func TestReadUint32(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xfa, 0x01, 0x01, 0x10, 0x13, 0xfa, 0x01, 0x10, 0x13, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 2, false, true)
	ui := ds.ReadUint32(2)
	if ui != 17830906 {
		t.Errorf("Read is not correct: %d\n", ui)
	}
}

func TestReadInt32(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xfa, 0x01, 0x01, 0x10, 0x13, 0xfa, 0x01, 0x10, 0x13, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 0, true, true)
	ui := ds.ReadInt32(1)
	if ui != 33166096 {
		t.Errorf("Read is not correct: %d\n", ui)
	}
}

func TestReadUint64(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xfa, 0x01, 0x01, 0x10, 0x13, 0xfa, 0x01, 0x10, 0x13, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 2, false, true)
	ui := ds.ReadUint64(2)
	if ui != 76583158144897043 {
		t.Errorf("Read is not correct: %d\n", ui)
	}
}

func TestReadInt64(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xfa, 0x01, 0x01, 0x10, 0x13, 0xfa, 0x01, 0x10, 0x13, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 0, true, true)
	ui := ds.ReadInt64(1)
	if ui != -426979943155887344 {
		t.Errorf("Read is not correct: %d\n", ui)
	}
}

func TestReadFloat32(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xfa, 0x01, 0x01, 0x10, 0x13, 0xfa, 0x01, 0x10, 0x13, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 0, true, true)
	ui := ds.ReadFloat32(1)
	if ui != 9.186285e-38 {
		t.Errorf("Read is not correct: %v\n", ui)
	}
}

func TestReadFloat64(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xfa, 0x01, 0x01, 0x10, 0x13, 0xfa, 0x01, 0x10, 0x13, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 0, true, true)
	ui := ds.ReadFloat64(1)
	if ui != -1.0813248704204464e+280 {
		t.Errorf("Read is not correct: %v\n", ui)
	}
}
