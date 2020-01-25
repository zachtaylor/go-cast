package cast_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"ztaylor.me/cast"
)

func BenchmarkSTDFmtSprintBuiltinSlice(b *testing.B) {
	data := MakeDataSlice()
	for i := 0; i < b.N; i++ {
		fmt.Sprint(data...)
	}
}

func BenchmarkSTDJsonMarshalBuiltinSlice(b *testing.B) {
	data := MakeDataSlice()
	for i := 0; i < b.N; i++ {
		json.Marshal(data)
	}
}

func BenchmarkCastStringBuiltinSlice(b *testing.B) {
	data := MakeDataSlice()
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}

func BenchmarkCastStringCastArray(b *testing.B) {
	data := cast.Array(MakeDataSlice())
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}

func BenchmarkCastStringNBuiltinSlice(b *testing.B) {
	data := MakeDataSlice()
	for i := 0; i < b.N; i++ {
		cast.StringN(data...)
	}
}
