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

// Add adds items to the items
func (a *Array) Add(items ...interface{}) {
	if items == nil {
		return
	}
	for _, item := range items {
		*a = append(*a, item)
	}
}

// // String isn't faster than fmt.Sprintf([]interface{})
// func (a *Array) String() (str string) {
// 	sb := poolStringBuilder.Get().(StringBuilder)
// 	sb.WriteByte('[')
// 	for k, v := range *a {
// 		if k > 0 {
// 			sb.WriteByte(' ')
// 		}
// 		sb.WriteString(String(v))
// 	}
// 	sb.WriteByte(']')
// 	str = sb.String()
// 	sb.Reset()
// 	poolStringBuilder.Put(sb)
// 	return
// }
