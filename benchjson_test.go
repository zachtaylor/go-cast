package cast_test

import (
	"encoding/json"
	"testing"

	"ztaylor.me/cast"
)

func BenchmarkSTDJsonMarshalBuiltinMap1(b *testing.B) {
	data := MakeDataMap1()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(data)
	}
}

func BenchmarkCastDictEncodeJSON1(b *testing.B) {
	data, _ := cast.ReflectDict(MakeDataMap1())
	for i := 0; i < b.N; i++ {
		_ = data.EncodeJSON()
	}
}

func BenchmarkCastJSONString1(b *testing.B) {
	data := cast.JSON(MakeDataMap1())
	for i := 0; i < b.N; i++ {
		_ = data.String()
	}
}

func BenchmarkSTDJsonMarshalBuiltinMap2(b *testing.B) {
	data := MakeDataMap2()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(data)
	}
}

func BenchmarkCastDictEncodeJSON2(b *testing.B) {
	data, _ := cast.ReflectDict(MakeDataMap2())
	for i := 0; i < b.N; i++ {
		_ = data.EncodeJSON()
	}
}

func BenchmarkCastJSONString2(b *testing.B) {
	data := cast.JSON(MakeDataMap2())
	for i := 0; i < b.N; i++ {
		_ = data.String()
	}
}

func BenchmarkSTDJsonMarshalBuiltinMap3(b *testing.B) {
	data := MakeDataMap3()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(data)
	}
}

func BenchmarkCastDictEncodeJSON3(b *testing.B) {
	data, _ := cast.ReflectDict(MakeDataMap3())
	for i := 0; i < b.N; i++ {
		_ = data.EncodeJSON()
	}
}

func BenchmarkCastJSONString3(b *testing.B) {
	data := cast.JSON(MakeDataMap3())
	for i := 0; i < b.N; i++ {
		_ = data.String()
	}
}
