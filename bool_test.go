package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func TestCastBool(t *testing.T) {
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

	if cast.Bool([]byte{'f', 'a', 'l', 's', 'e'}) {
		t.Log(`cast.Bool([]byte{'f', 'a', 'l', 's', 'e'})!=false`)
		t.Fail()
	}
}
