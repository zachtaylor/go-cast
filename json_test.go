package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func TestJSONString(t *testing.T) {
	json := cast.JSON{
		"hello": cast.Stringer("world"),
	}
	jsons := json.String()
	if jsons != `{"hello":world}` {
		t.Fail()
	}
}

func TestEncodeJSON(t *testing.T) {
	ans := []interface{}{"hello world", map[int]bool{20: true}, map[string]string{"x": "19"}, cast.Stringer(`foobar`)}
	anss := `["hello world",{20:true},{"x":"19"},foobar]`
	if str := cast.EncodeJSON(ans); str != anss {
		t.Log(`Expected ` + anss)
		t.Log(`Received ` + str)
		t.Fail()
	}
}

func TestJSONCopy(t *testing.T) {
	data := cast.JSON{
		"k": "s",
		"i": 1,
		"b": true,
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
