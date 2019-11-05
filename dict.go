package cast

// Dict is just a map of any
type Dict map[interface{}]interface{}

// DictEncoder writes map to string
type DictEncoder func(*Dict) string

func (d *Dict) get() Dict {
	if d == nil {
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

// String uses the same format as fmt but it's faster when Dict is fmt.Stringer
func (d *Dict) String() string {
	var sb StringBuilder
	sb.WriteString(`map[`)
	first := true
	for k, v := range *d {
		if !first {
			sb.WriteByte(' ')
		}
		sb.WriteString(String(k))
		sb.WriteByte(':')
		sb.WriteString(String(v))
		first = false
	}
	sb.WriteByte(']')
	return sb.String()
}
