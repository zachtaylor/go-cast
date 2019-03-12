package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func BenchmarkStringBytes(b *testing.B) {
	x := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	for i := 0; i < b.N; i++ {
		_ = cast.StringBytes(x)
	}
}

func BenchmarkDefaultStringBytes(b *testing.B) {
	x := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	for i := 0; i < b.N; i++ {
		_ = string(x)
	}
}

func BenchmarkBytesS(b *testing.B) {
	x := "hello world"
	for i := 0; i < b.N; i++ {
		_ = cast.BytesS(x)
	}
}

func BenchmarkDefaultBytesS(b *testing.B) {
	x := "hello world"
	for i := 0; i < b.N; i++ {
		_ = []byte(x)
	}
}
