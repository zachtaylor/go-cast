package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func TestBool(t *testing.T) {
	if cast.Bool(nil) {
		t.Log(`cast.Bool(nil)!=false`)
		t.Fail()
	}

	if cast.Bool("") {
		t.Log(`cast.Bool("")!=false`)
		t.Fail()
	}

	if cast.Bool("false") {
		t.Log(`cast.Bool("false")!=false`)
		t.Fail()
	}

	if !cast.Bool([]byte{'h', 'i'}) {
		t.Log(`cast.Bool([]byte{'h', 'i'})==false`)
		t.Fail()
	}
}
