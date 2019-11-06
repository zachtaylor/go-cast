package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func TestCastString(t *testing.T) {
	ans := []interface{}{"hello world", map[int]bool{20: true}, map[string]string{"x": "19"}, cast.Stringer(`foobar`)}
	anss := `[hello world map[20:true] map[x:19] foobar]`
	if str := cast.String(ans); str != anss {
		t.Log(`Expected ` + anss)
		t.Log(`Received ` + str)
		t.Fail()
	}
}

func TestCastStringBool(t *testing.T) {
	if "true" != cast.String(true) {
		t.Fail()
	} else if "false" != cast.String(false) {
		t.Fail()
	}
}
