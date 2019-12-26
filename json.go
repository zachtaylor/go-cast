package cast

import "encoding/json"

// JSON is basic KV with value conversion getters
type JSON map[string]interface{}

// EncodeJSON writes a value as a JSON value
func EncodeJSON(arg interface{}) (str string) {
	switch v := arg.(type) {
	case string:
		str = `"` + EscapeString(v) + `"`
	case IStringer:
		str = v.String()
	case []byte:
		str = `"` + EscapeString(StringBytes(v)) + `"`
	case error:
		str = `"error: ` + v.Error() + `"`
	case bool:
		str = StringB(v)
	case int:
		str = StringI(v)
	case uint:
		str = StringUI(v)
	case int64:
		str = StringI64(v)
	case float64:
		str = StringF(v)
	case Array:
		return v.EncodeJSON()
	case []interface{}:
		str = Array(v).EncodeJSON()
	case Dict:
		str = v.EncodeJSON()
	case map[interface{}]interface{}:
		d := Dict(v)
		str = d.EncodeJSON()
	case JSON:
		str = v.String()
	case map[string]interface{}:
		str = JSON(v).String()
	default:
		if array, ok := ReflectArray(arg); ok {
			str = array.EncodeJSON()
		} else if dict, ok := ReflectDict(arg); ok {
			str = dict.EncodeJSON()
		} else {
			str = Sprint(arg)
		}
	}
	return
}

// Get returns the value for the key, use map indexing for existance 'ok'
func (json JSON) Get(k string) interface{} {
	return json[k]
}

// GetS returns a string wrapped from Get return
func (json JSON) GetS(k string) string {
	return String(json[k])
}

// GetI returns an int wrapped from Get return
func (json JSON) GetI(k string) int {
	return Int(json[k])
}

// GetB returns a bool wrapped from Get return
func (json JSON) GetB(k string) bool {
	return Bool(json[k])
}

// String encodes JSON to string
func (json JSON) String() (str string) {
	sb := poolStringBuilder.Get().(*StringBuilder)
	sb.Grow(growFactor * len(json))
	var k string
	var v interface{}
	sb.WriteByte('{')
	first := true
	for k, v = range json {
		if !first {
			sb.WriteByte(',')
		} else {
			first = false
		}
		sb.WriteByte('"')
		sb.WriteString(EscapeString(k))
		sb.WriteString(`":`)
		sb.WriteString(EncodeJSON(v))
	}
	sb.WriteByte('}')
	str = sb.String()
	sb.Reset()
	poolStringBuilder.Put(sb)
	return
}

// Copy returns a new JSON that is shallow-copied
func (json JSON) Copy() (j JSON) {
	if json != nil {
		j = JSON{}
		for k, v := range json {
			j[k] = v
		}
	}
	return
}

// DecodeJSON wraps encoding/json.Decoder.Decode
func DecodeJSON(r Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
