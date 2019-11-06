package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func TestEncodeJSON(t *testing.T) {
	ans := []interface{}{"hello world", map[int]bool{20: true}, map[string]string{"x": "19"}, cast.Stringer(`foobar`)}
	anss := `["hello world",{20:true},{"x":"19"},foobar]`
	if str := cast.EncodeJSON(ans); str != anss {
		t.Log(`Expected ` + anss)
		t.Log(`Received ` + str)
		t.Fail()
	}
}
