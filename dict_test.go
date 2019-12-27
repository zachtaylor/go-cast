package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func TestDictNil(t *testing.T) {
	var dict cast.Dict
	dict.Set("hello", "world")
	if dict.Get("hello") != "world" {
		t.Log(`Dict(nil) failed Set()/Get()`)
		t.Fail()
	}
}

func TestDictString(t *testing.T) {
	dict := cast.Dict{
		"hello": cast.Stringer("world"),
	}
	dicts := dict.String()
	if dicts != "map[hello:world]" {
		t.Log(`Dict(string=>stringer) != [hello:world]:` + dicts)
		t.Fail()
	}
	dictjsons := dict.EncodeJSON()
	if dictjsons != `{"hello":world}` {
		t.Log(`Dict(string=>stringer) != {"hello":world}:` + dictjsons)
		t.Fail()
	}
}
