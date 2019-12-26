package cast

// Dict is just a map of any
type Dict map[interface{}]interface{}

// DictEncoder writes map to string
type DictEncoder func(*Dict) string

// get initializes the underlying data store if this was declared as nil
func (d *Dict) get() Dict {
	if *d == nil {
		*d = Dict{}
	}
	return *d
}

// Get returns the map value
func (d *Dict) Get(k interface{}) interface{} {
	return d.get()[k]
}

// Set sets an item in the map
func (d *Dict) Set(k, v interface{}) {
	d.get()[k] = v
}

// Unset deletes a key from the map
func (d *Dict) Unset(k interface{}) {
	delete(*d, k)
}

// Encode writes Dict to string
func (d Dict) Encode(start string, sep, end byte, keyEncoder, valEncoder func(interface{}) string) (str string) {
	sb := poolStringBuilder.Get().(*StringBuilder)
	sb.Grow(growFactor * len(d))
	sb.WriteString(start)
	first := true
	for k, v := range d {
		if !first {
			sb.WriteByte(sep)
		} else {
			first = false
		}
		sb.WriteString(keyEncoder(k))
		sb.WriteByte(':')
		sb.WriteString(valEncoder(v))
	}
	sb.WriteByte(end)
	str = sb.String()
	sb.Reset()
	poolStringBuilder.Put(sb)
	return
}

// String implements fmt.Stringer
func (d Dict) String() string {
	return d.Encode(`map[`, ' ', ']', String, String)
}

// EncodeJSON writes Dict to string in a JSON value format
func (d Dict) EncodeJSON() string {
	return d.Encode(`{`, ',', '}', EncodeJSON, EncodeJSON)
}
