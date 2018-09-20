package geotiff

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFieldTagNames(t *testing.T) {
	freeOffset := FieldTagNames[0x0120]
	if freeOffset != "FreeOffsets" {
		t.Errorf("got the wrong value")
	}
}

func TestFieldTags(t *testing.T) {
	freeOffset := FieldTags["FreeOffsets"]
	if freeOffset != 0x0120 {
		t.Errorf("got the wrong value")
	}
}

func TestFieldTypeNames(t *testing.T) {
	undefined := FieldTypeNames[0x0007]
	if undefined != "UNDEFINED" {
		t.Errorf("got the wrong value")
	}
}

func TestFieldTypes(t *testing.T) {
	undefined := FieldTypes["UNDEFINED"]
	if undefined != 0x0007 {
		t.Errorf("got the wrong value")
	}
}

func TestGeoKeyNames(t *testing.T) {
	wgs84 := GeoKeyNames[2062]
	if wgs84 != "GeogTOWGS84GeoKey" {
		t.Errorf("got the wrong value")
	}
}

func TestGeoKeys(t *testing.T) {
	wgs84 := GeoKeys["GeogTOWGS84GeoKey"]
	if wgs84 != 2062 {
		t.Errorf("got the wrong value")
	}
}

func TestGetFieldTypeLength(t *testing.T) {
	testVal := map[uint16]uint8{0x0001: 1, 0x0003: 2, 0x000B: 4, 0x0009: 4, 0x0012: 8}
	for f, n := range testVal {
		val, err := getFieldTypeLength(f)
		if err != nil {
			t.Errorf("Field Type not recognized: %v", f)
		}
		if val != n {
			t.Errorf("Got the wrong length: %v instead of %v", val, n)
		}
	}
}

func TestGetValues(t *testing.T) {
	dw := DataView{0x01, 0x10, 0x13, 0xff}
	ds := NewDataSlice(dw, 0, false, false)
	values, err := GetValues(ds, FieldTypes["BYTE"], 2, 2)
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	for _, v := range values {
		fmt.Printf("V is: %v, %T\n", v, v)
		fmt.Printf("Type: %v\n", reflect.TypeOf(v).Name())
		if reflect.TypeOf(v).Kind() != reflect.Uint8 {
			t.Errorf("type should be uint8 but is %T", v)
		}
	}
}
