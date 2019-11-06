package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func TestJSONCopy(t *testing.T) {
	data := cast.JSON{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"i1": 1,
		"i2": 2,
		"i3": 3,
		"b1": true,
		"b2": false,
	}
	data2 := data.Copy()
	for k, v := range data {
		if data2[k] != v {
			t.Log(`Expected`, v)
			t.Log(`Actual`, data2[k])
			t.Fail()
		}
	}
}
