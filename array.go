package cast

// Array is used to shuffle slices of data, and has extra controls suitable for encoding
type Array struct {
	Items   []interface{}
	Encoder ArrayEncoder
}

// ArrayEncoder writes array to []byte
type ArrayEncoder func(*Array) string

// defaultArrayEncoder writes arrays with [x,x,...,x], escapes string literals in "quotes"
func defaultArrayEncoder(a *Array) string {
	var sb StringBuilder
	sb.WriteByte('[')
	for k, v := range a.Items {
		if k > 0 {
			sb.WriteByte(',')
		}
		switch v.(type) {
		case string:
			sb.WriteByte('"')
			sb.WriteString(String(v))
			sb.WriteByte('"')
		default:
			sb.WriteString(String(v))
		}
	}
	sb.WriteByte(']')
	return sb.String()
}

// Arr creates a new Array with literal `[]interface{}`
func Arr(slice []interface{}) *Array {
	return &Array{slice, defaultArrayEncoder}
}

// NewArray takes any number of arguments of any type
//
// it's actually illegal to use the spread operator with any type besides `[]interface{}`, see `SpreadArray`
func NewArray(items ...interface{}) *Array {
	return Arr(items)
}

// SpreadArray casts any slice (w/reflect) into NewArray
func SpreadArray(slice interface{}) (a *Array) {
	if v, ok := castSlice(slice); ok {
		a = NewArray(v...)
	}
	return
}

// Get returns the slice
func (a *Array) Get(i int) interface{} {
	return a.Items[i]
}

// Length returns len(a.items)
func (a *Array) Length() int {
	return len(a.Items)
}

// Add adds items to the items
func (a *Array) Add(items ...interface{}) {
	if items == nil {
		return
	}
	for _, item := range items {
		a.Items = append(a.Items, item)
	}
}

// String uses the internal encoder if available, else package-level String
func (a *Array) String() string {
	if a.Encoder == nil {
		return defaultArrayEncoder(a)
	}
	return a.Encoder(a)
}
