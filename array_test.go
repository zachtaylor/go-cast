package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func TestArrayNil(t *testing.T) {
	var array cast.Array
	array.Add(`hello world`)
	if array.Length() < 1 {
		t.Log(`Array(nil).Add() fails`)
		t.Fail()
	}
}

func TestArrayString(t *testing.T) {
	ans := []interface{}{"hello world", 42, cast.Stringer(`foobar`), true}
	anss := `["hello world",42,foobar,true]`
	if str := cast.String(ans); str != anss {
		t.Log(`String()=` + str + ` expected=` + anss)
		t.Fail()
	}
	array := cast.SpreadArray(ans)
	if str := array.String(); str != anss {
		t.Log(`array.String=` + str + ` expected=` + anss)
		t.Fail()
	}
}

func TestSpreadArray(t *testing.T) {
	ans := []bool{true, true, false, false, true, false}
	if array := cast.SpreadArray(ans); array != nil {
		anslen := len(ans)
		actlen := array.Length()
		if actlen != anslen {
			t.Log(`array.length=` + cast.StringI(actlen) + ` expected=` + cast.StringI(anslen))
			t.Fail()
		}
		for i := 0; i < anslen; i++ {
			if act := array.Get(i); ans[i] != act {
				t.Log(`array[i]=` + cast.String(act) + ` expected=` + cast.StringB(ans[i]) + ` i=` + cast.StringI(i))
				t.Fail()
			}
		}
	}
}
