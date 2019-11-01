package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func TestArrayString(t *testing.T) {
	ans := []interface{}{"hello world", 42, cast.Stringer(`foobar`), true}
	anss := `["hello world",42,foobar,true]`
	if str := cast.String(ans); str != anss {
		t.Log(`String()=` + str + ` expected=` + anss)
		t.Fail()
	}
}
