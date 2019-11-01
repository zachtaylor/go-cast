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
		"dat1": "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"dat2": "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"dat3": "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
	}
}

func BenchmarkLoadedMapString(b *testing.B) {
	data := MakeData()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%v", data)
	}
}
func BenchmarkLoadedCastJSONString(b *testing.B) {
	data := cast.JSON(MakeData())
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}
