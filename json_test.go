package cast_test

import (
	"fmt"
	"testing"

	"ztaylor.me/cast"
)

func MakeData() map[string]interface{} {
	return map[string]interface{}{
		"k1":   "v1",
		"k2":   "v2",
		"k3":   "v3",
		"i1":   1,
		"i2":   2,
		"i3":   3,
		"b1":   true,
		"b2":   false,
		"map":  map[string]interface{}{},
		"json": cast.JSON{},
	}
}

func BenchmarkLoadedMapString(b *testing.B) {
	data := MakeData()
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}
func BenchmarkLoadedJSONString(b *testing.B) {
	data := cast.JSON(MakeData())
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}

func BenchmarkNilJSONFmtSprint(b *testing.B) {
	data := cast.JSON(nil)
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(data)
	}
}

func BenchmarkNilJSONCastString(b *testing.B) {
	data := cast.JSON(nil)
	for i := 0; i < b.N; i++ {
		_ = cast.String(data)
	}
}
