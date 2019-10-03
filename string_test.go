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

func BenchmarkStringFmtSprint(b *testing.B) {
	data := "string"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(data)
	}
}

func BenchmarkStringCastSprint(b *testing.B) {
	data := "string"
	for i := 0; i < b.N; i++ {
		_ = cast.String(data)
	}
}

func BenchmarkNilMapFmtSprint(b *testing.B) {
	data := interface{}(nil)
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(data)
	}
}

func BenchmarkNilMapCastString(b *testing.B) {
	data := interface{}(nil)
	for i := 0; i < b.N; i++ {
		_ = cast.String(data)
	}
}
