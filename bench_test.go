package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func BenchmarkBuiltinStringBytes(b *testing.B) {
	x := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	for i := 0; i < b.N; i++ {
		_ = string(x)
	}
}

func BenchmarkCastStringBytes(b *testing.B) {
	x := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	for i := 0; i < b.N; i++ {
		_ = cast.StringBytes(x)
	}
}

func BenchmarkBuiltinBytesS(b *testing.B) {
	x := "hello world"
	for i := 0; i < b.N; i++ {
		_ = []byte(x)
	}
}

func BenchmarkCastBytesS(b *testing.B) {
	x := "hello world"
	for i := 0; i < b.N; i++ {
		_ = cast.BytesS(x)
	}
}

func BenchmarkLoadedMapSprintf(b *testing.B) {
	data := MakeDataMap()
	for i := 0; i < b.N; i++ {
		cast.Sprintf("%v", data)
	}
}

func BenchmarkLoadedJSONString(b *testing.B) {
	data := cast.JSON(MakeDataMap())
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}

func BenchmarkStringSprint(b *testing.B) {
	data := "stringdata123456789123456789123456789"
	for i := 0; i < b.N; i++ {
		_ = cast.Sprint(data)
	}
}

func BenchmarkStringString(b *testing.B) {
	data := "stringdata123456789123456789123456789"
	for i := 0; i < b.N; i++ {
		_ = cast.String(data)
	}
}
