package cast_test

import (
	"fmt"
	"testing"

	"ztaylor.me/cast"
)

func BenchmarkLoadedMapFmtSprintf(b *testing.B) {
	data := MakeDataMap()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%v", data)
	}
}
func BenchmarkLoadedCastJSONString(b *testing.B) {
	data := cast.JSON(MakeDataMap())
	for i := 0; i < b.N; i++ {
		cast.String(data)
	}
}
