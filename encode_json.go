package cast

// EncodeJSON writes a value as a JSON value
func EncodeJSON(arg interface{}) (str string) {
	switch v := arg.(type) {
	case string:
		str = `"` + EscapeString(v) + `"`
	case IStringer:
		str = v.String()
	case []byte:
		str = StringBytes(v)
	case error:
		str = v.Error()
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

type encoder interface {
	EncodeJSON() string
}

// EncodeJSON builds a JSON string representation of this slice
//
// See `EncodeJSON`
func (a *Array) EncodeJSON() string {
	var sb StringBuilder
	sb.WriteByte('[')
	for k, v := range *a {
		if k > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(EncodeJSON(v))
	}
	sb.WriteByte(']')
	return sb.String()
}

// EncodeJSON builds a JSON string representation of this map
//
// See `EncodeJSON`
func (d *Dict) EncodeJSON() string {
	var sb StringBuilder
	var k, v interface{}
	sb.WriteByte('{')
	first := true
	for k, v = range *d {
		if !first {
			sb.WriteByte(',')
		} else {
			first = false
		}
		sb.WriteString(EncodeJSON(k))
		sb.WriteByte(':')
		sb.WriteString(EncodeJSON(v))
	}
	sb.WriteByte('}')
	return sb.String()
}
