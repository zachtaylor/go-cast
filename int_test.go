package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func TestInt(t *testing.T) {
	if cast.Int(true) == cast.Int(false) {
		t.Log(`cast.Int(true) == cast.Int(false)`)
		t.Fail()
	}
	if cast.Int(nil) == cast.Int(false) {
		t.Log(`cast.Int(nil) == cast.Int(false)`)
		t.Fail()
	}
}
