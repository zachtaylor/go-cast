package cast_test

import (
	"encoding/json"
	"fmt"
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

// encoding races

// map format

func BenchmarkSTDFmtSprintfBuiltinMap(b *testing.B) {
	data := MakeDataMap()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%v", data)
	}
}

func BenchmarkSTDJsonMarshalBuiltinMap(b *testing.B) {
	data := MakeDataMap()
	for i := 0; i < b.N; i++ {
		json.Marshal(data)
	}
}

func BenchmarkCastDictEncodeJSON(b *testing.B) {
	data, _ := cast.ReflectDict(MakeDataMap())
	for i := 0; i < b.N; i++ {
		data.EncodeJSON()
	}
}

func BenchmarkCastStringBuiltinMap(b *testing.B) {
	data := MakeDataMap()
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}

func BenchmarkCastJSONString(b *testing.B) {
	data := cast.JSON(MakeDataMap())
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}

func BenchmarkCastDictString(b *testing.B) {
	data, _ := cast.ReflectDict(MakeDataMap())
	for i := 0; i < b.N; i++ {
		data.String()
	}
}

// array format

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
