package cast

import (
	"encoding/json"
	"io"
)

// DecodeJSON wraps encoding/json.Decoder.Decode
func DecodeJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
