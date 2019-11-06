package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func TestCastString(t *testing.T) {
	ans := []interface{}{"hello world", map[int]int{20: 19}, "0", 1, true, cast.Stringer(`foobar`)}
	anss := `[hello world map[20:19] 0 1 true foobar]`
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
