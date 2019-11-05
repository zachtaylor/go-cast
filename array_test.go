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

func TestReflectArray(t *testing.T) {
	ans := []bool{true, true, false, false, true, false}
	if array, ok := cast.ReflectArray(ans); ok {
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
