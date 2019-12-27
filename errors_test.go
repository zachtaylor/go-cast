package cast_test

import (
	"testing"

	"ztaylor.me/cast"
)

func TestErrorEncoding(t *testing.T) {
	err := &cast.Error{`hello`, cast.NewError(nil, `world`)}
	expecteds := `error: "hello: world"`
	if actual := cast.String(err); actual != expecteds {
		t.Log("Expected", expecteds)
		t.Log("Actual", actual)
		t.Fail()
	}
	expectedjsons := `"error: hello: world"`
	if actual := cast.EncodeJSON(err); actual != expectedjsons {
		t.Log("Expected", expectedjsons)
		t.Log("Actual", actual)
		t.Fail()
	}
}
