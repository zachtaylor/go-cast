package cast

// JSON is basic KV with value conversion getters
type JSON map[string]interface{}

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
	sb := poolStringBuilder.Get().(StringBuilder)
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
