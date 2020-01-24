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

func BenchmarkSTDJsonMarshalBuiltinMap1(b *testing.B) {
	data := MakeDataMap1()
	for i := 0; i < b.N; i++ {
		json.Marshal(data)
	}
}

func BenchmarkSTDFmtSprintfBuiltinMap1(b *testing.B) {
	data := MakeDataMap1()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%v", data)
	}
}

func BenchmarkCastDictEncodeJSON1(b *testing.B) {
	data, _ := cast.ReflectDict(MakeDataMap1())
	for i := 0; i < b.N; i++ {
		data.EncodeJSON()
	}
}

func BenchmarkCastStringBuiltinMap1(b *testing.B) {
	data := MakeDataMap1()
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}

func BenchmarkCastJSONString1(b *testing.B) {
	data := cast.JSON(MakeDataMap1())
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}

func BenchmarkCastDictString1(b *testing.B) {
	data, _ := cast.ReflectDict(MakeDataMap1())
	for i := 0; i < b.N; i++ {
		data.String()
	}
}

// map block 2

func BenchmarkSTDJsonMarshalBuiltinMap2(b *testing.B) {
	data := MakeDataMap2()
	for i := 0; i < b.N; i++ {
		json.Marshal(data)
	}
}

func BenchmarkSTDFmtSprintfBuiltinMap2(b *testing.B) {
	data := MakeDataMap2()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%v", data)
	}
}

func BenchmarkCastDictEncodeJSON2(b *testing.B) {
	data, _ := cast.ReflectDict(MakeDataMap2())
	for i := 0; i < b.N; i++ {
		data.EncodeJSON()
	}
}

func BenchmarkCastStringBuiltinMap2(b *testing.B) {
	data := MakeDataMap2()
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}
func BenchmarkCastJSONString2(b *testing.B) {
	data := cast.JSON(MakeDataMap2())
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}

func BenchmarkCastDictString2(b *testing.B) {
	data, _ := cast.ReflectDict(MakeDataMap2())
	for i := 0; i < b.N; i++ {
		data.String()
	}
}

// map block 3

func BenchmarkSTDJsonMarshalBuiltinMap3(b *testing.B) {
	data := MakeDataMap3()
	for i := 0; i < b.N; i++ {
		json.Marshal(data)
	}
}

func BenchmarkSTDFmtSprintfBuiltinMap3(b *testing.B) {
	data := MakeDataMap3()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%v", data)
	}
}

func BenchmarkCastDictEncodeJSON3(b *testing.B) {
	data, _ := cast.ReflectDict(MakeDataMap3())
	for i := 0; i < b.N; i++ {
		data.EncodeJSON()
	}
}

func BenchmarkCastStringBuiltinMap3(b *testing.B) {
	data := MakeDataMap3()
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}
func BenchmarkCastJSONString3(b *testing.B) {
	data := cast.JSON(MakeDataMap3())
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}

func BenchmarkCastDictString3(b *testing.B) {
	data, _ := cast.ReflectDict(MakeDataMap3())
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
