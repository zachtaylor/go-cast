package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func TestString(t *testing.T) {
	if cast.String(true) == cast.String(false) {
		t.Log(`cast.String(true) == cast.String(false)`)
		t.Fail()
	}

	if cast.String(nil) == cast.String(false) {
		t.Log(`cast.String(nil) == cast.String(false)`)
		t.Fail()
	}

	if cast.String([]byte{'f', 'o', 'o'}) == cast.String([]byte{'b', 'a', 'r'}) {
		t.Log(`cast.String([]byte)`)
		t.Fail()
	}
}
