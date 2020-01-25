package cast_test

import (
	"fmt"
	"testing"

	"ztaylor.me/cast"
)

func BenchmarkSTDFmtSprintfBuiltinMap1(b *testing.B) {
	data := MakeDataMap1()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%v", data)
	}
}

func BenchmarkCastStringBuiltinMap1(b *testing.B) {
	data := MakeDataMap1()
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

func BenchmarkSTDFmtSprintfBuiltinMap2(b *testing.B) {
	data := MakeDataMap2()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%v", data)
	}
}

func BenchmarkCastStringBuiltinMap2(b *testing.B) {
	data := MakeDataMap2()
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

func BenchmarkSTDFmtSprintfBuiltinMap3(b *testing.B) {
	data := MakeDataMap3()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%v", data)
	}
}

func BenchmarkCastStringBuiltinMap3(b *testing.B) {
	data := MakeDataMap3()
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
