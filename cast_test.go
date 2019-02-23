package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func TestCastStringBool(t *testing.T) {
	if cast.String(true) == cast.String(false) {
		t.Log(`cast.String(bool)`)
		t.Fail()
	}
}

func TestCastIntBool(t *testing.T) {
	if cast.Int(true) == cast.Int(false) {
		t.Log(`cast.Int(bool)`)
		t.Fail()
	}
}

func TestCastBoolFalse(t *testing.T) {
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
}
