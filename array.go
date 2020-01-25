package cast

// Array is a slice of any
type Array []interface{}

// ArrayEncoder writes array to string
type ArrayEncoder func(*Array) string

// NewArray takes any number of arguments of any type
//
// it's actually illegal to use the spread operator with any type besides `[]interface{}`, see `ReflectArray`
func NewArray(items ...interface{}) Array {
	return Array(items)
}

// Get returns the slice value
func (a *Array) Get(i int) (val interface{}) {
	if a != nil {
		val = (*a)[i]
	}
	return
}

// Length returns len(a.items)
func (a Array) Length() int {
	return len(a)
}

// Size returns a guess at the number of bytes needed to encode the array
func (a Array) Size() (int int) {
	a.addSize(&int)
	return
}

func (a Array) addSize(int *int) {
	*int += 2 + len(a) // brackets and spaces
	for _, v := range a {
		*int += Size(v)
	}
	*int += *int % growFactor // round up to multiple of 32
	return
}

// Add adds items to the items
func (a *Array) Add(items ...interface{}) {
	if items == nil {
		return
	}
	for _, item := range items {
		*a = append(*a, item)
	}
}

// Encode writes Array to string
func (a Array) Encode(sep byte, encoder func(interface{}) string) (str string) {
	sb := poolStringBuilder.Get().(*StringBuilder)
	a.encode(sb, sep, encoder)
	str = sb.String()
	sb.Reset()
	poolStringBuilder.Put(sb)
	return
}

func (a Array) encode(sb *StringBuilder, sep byte, encoder func(interface{}) string) {
	sb.Grow(a.Size())
	sb.WriteByte('[')
	for k, v := range a {
		if k > 0 {
			sb.WriteByte(sep)
		}
		sb.WriteString(encoder(v))
	}
	sb.WriteByte(']')
}

// String is faster than fmt.Sprintf
func (a Array) String() string {
	if a == nil {
		return "[]"
	}
	return a.Encode(' ', String)
}

// EncodeJSON builds a JSON string representation of this slice
func (a Array) EncodeJSON() string {
	if a == nil {
		return "[]"
	}
	return a.Encode(',', EncodeJSON)
}
