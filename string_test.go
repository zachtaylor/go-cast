package cast_test

import (
	"fmt"
	"testing"

	"ztaylor.me/cast"
)

func TestMapCastString(t *testing.T) {
	data := map[string]interface{}{}
	if cast.String(data) == "" {
		t.Log(`cast.String(map) == ""`)
		t.Fail()
	}
}

func TestBoolNilCastString(t *testing.T) {
	if cast.String(true) == cast.String(false) {
		t.Log(`cast.String(true) == cast.String(false)`)
		t.Fail()
	}

	if cast.String(nil) == cast.String(false) {
		t.Log(`cast.String(nil) == cast.String(false)`)
		t.Fail()
	}

	if cast.String([]byte{'f', 'o', 'o'}) == cast.String([]byte{'b', 'a', 'r'}) {
		t.Log(`cast.String([]byte)`)
		t.Fail()
	}
}

func TestSliceString(t *testing.T) {
	bs := []bool{true, true, true, false, false, true, false}
	if cast.String(bs) == "" {
		t.Log(`cast.String(boolslice) == ""`)
		t.Fail()
	}
}

func BenchmarkStringFmtSprint(b *testing.B) {
	data := "stringdata123456789123456789123456789"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(data)
	}
}

func BenchmarkStringCastString(b *testing.B) {
	data := "stringdata123456789123456789123456789"
	for i := 0; i < b.N; i++ {
		_ = cast.String(data)
	}
}
