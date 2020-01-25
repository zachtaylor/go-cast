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

// Length returns the number of items in the map
func (d *Dict) Length() int {
	return len(d.get())
}

// Size returns a guess at the number of bytes needed to encode the map
func (d Dict) Size() (int int) {
	d.addSize(&int)
	return
}

func (d Dict) addSize(int *int) {
	*int = 2 * len(d) // colons and separators
	for k, v := range d {
		*int += Size(k)
		*int += Size(v)
	}
	*int += *int % growFactor // round up to multiple of 32
	return
}

func (d Dict) encode(sb *StringBuilder, start string, sep, end byte, keyEncoder, valEncoder func(interface{}) string) {
	sb.Grow(d.Size())
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
}

// String implements fmt.Stringer
func (d Dict) String() (string string) {
	if d == nil {
		return "nil"
	}
	sb := poolStringBuilder.Get().(*StringBuilder)
	d.encode(sb, `map[`, ' ', ']', String, String)
	string = sb.String()
	sb.Reset()
	poolStringBuilder.Put(sb)
	return
}

// EncodeJSON writes Dict to string in a JSON value format
func (d Dict) EncodeJSON() (string string) {
	if d == nil {
		return "null"
	}
	sb := poolStringBuilder.Get().(*StringBuilder)
	d.encode(sb, `{`, ',', '}', EncodeJSON, EncodeJSON)
	string = sb.String()
	sb.Reset()
	poolStringBuilder.Put(sb)
	return
}
