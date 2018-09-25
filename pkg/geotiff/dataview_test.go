package geotiff

import (
	"testing"
)

func TestUint8(t *testing.T) {
	data := DataView{0x01, 0x10, 0x13, 0xff}
	expected := uint8(1)
	u := data.Uint8(0, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	expected = uint8(1)
	u = data.Uint8(0, true)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	expected = uint8(16)
	u = data.Uint8(1, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
}

func TestInt8(t *testing.T) {
	data := DataView{0xFC, 0xD3, 0x0D}
	expected := int8(-4)
	u := data.Int8(0, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	expected = int8(-45)
	u = data.Int8(1, true)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	expected = int8(13)
	u = data.Int8(2, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
}

func TestUint16(t *testing.T) {
	data := DataView{0x11, 0x12, 0x13, 0xff}
	expected := uint16(4370)
	u := data.Uint16(0, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	expected = uint16(4625)
	u = data.Uint16(0, true)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	expected = uint16(5119)
	u = data.Uint16(2, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
}
func TestInt16(t *testing.T) {
	data := DataView{0xD3, 0xC3, 0x00, 0x12, 0x11, 0x12, 0x13, 0xFF}
	expected := int16(-15616)
	u := data.Int16(1, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	data = DataView{0xFF, 0x13, 0x12, 0x11, 0x12, 0x00, 0xC3, 0xD3}
	expected = int16(-15616)
	u = data.Int16(5, true)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	data = DataView{0xFF, 0x13, 0x12, 0x11, 0x12, 0x00, 0xC3, 0xD3}
	expected = int16(4882)
	u = data.Int16(1, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
}

func TestUint32(t *testing.T) {
	data := DataView{0x11, 0x12, 0x13, 0xff, 0x00}
	expected := uint32(286397439)
	u := data.Uint32(0, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	expected = uint32(4279439889)
	u = data.Uint32(0, true)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	expected = uint32(303300352)
	u = data.Uint32(1, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
}

func TestInt32(t *testing.T) {
	data := DataView{0xF1, 0x12, 0x13, 0xff, 0x00}
	expected := int32(-250473473)
	u := data.Int32(0, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	expected = int32(-15527183)
	u = data.Int32(0, true)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	expected = int32(303300352)
	u = data.Int32(1, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
}

func TestUint64(t *testing.T) {
	data := DataView{0xD3, 0xC3, 0x00, 0x12, 0x11, 0x12, 0x13, 0xFF}
	expected := uint64(15259040040057181183)
	u := data.Uint64(0, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	data = DataView{0xFF, 0x13, 0x12, 0x11, 0x12, 0x00, 0xC3, 0xD3}
	expected = uint64(15259040040057181183)
	u = data.Uint64(0, true)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
}

func TestInt64(t *testing.T) {
	data := DataView{0xFD, 0xE1, 0xE3, 0xD6, 0xD8, 0x9D, 0x78, 0x01}
	expected := int64(-152590400457181183)
	u := data.Int64(0, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	data = DataView{0x01, 0x78, 0x2B, 0xF5, 0x18, 0x45, 0xD3, 0xFF}
	expected = int64(-12590400457181183)
	u = data.Int64(0, true)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	data = DataView{0x00, 0x2C, 0xBA, 0xE7, 0x0A, 0xD4, 0x87, 0xFF}
	expected = int64(12590400457181183)
	u = data.Int64(0, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
}

func TestFloat32(t *testing.T) {
	data := DataView{0xF1, 0x12, 0x13, 0xff, 0x00}
	expected := float32(-7.233438e+29)
	u := data.Float32(0, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	expected = float32(-1.9549486e+38)
	u = data.Float32(0, true)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
	expected = float32(4.6699333e-28)
	u = data.Float32(1, false)
	if u != expected {
		t.Errorf("expected value is %v but got %v", expected, u)
	}
}
