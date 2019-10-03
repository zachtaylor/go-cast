package cast

// JSON is basic KV with value conversion getters
type JSON map[string]interface{}

// Get returns the value for the key
//
// To check for existence specifically, use map indexing instead ojson this func
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

func (json JSON) String() string {
	var sb StringBuilder
	var size int
	var k string
	var v interface{}
	for k, _ = range json {
		size += len(k)
		size += 12
	}
	sb.Grow(size)
	sb.WriteByte('{')
	first := true
	for k, v = range json {
		if !first {
			sb.WriteByte(',')
		} else {
			first = false
		}
		sb.WriteByte('"')
		sb.WriteString(k)

		switch x := v.(type) {
		case string:
			sb.WriteString(`":"`)
			sb.WriteString(x)
			sb.WriteByte('"')
		default:
			sb.WriteString(`":`)
			sb.WriteString(String(v))
		}
	}
	sb.WriteByte('}')
	return sb.String()
}

// Copy returns a new JSON that is shallow-copied
func (json JSON) Copy() JSON {
	JSON := JSON{}
	for k, v := range json {
		JSON[k] = v
	}
	return JSON
}
