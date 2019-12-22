package cast

import "reflect"

// Kind = reflect.Kind
type Kind = reflect.Kind

// Value = reflect.Value
type Value = reflect.Value

// Slice = reflect.Slice
const Slice = reflect.Slice

// Reflect uses reflect to any kind
func Reflect(arg interface{}, kind Kind) (reflect.Value, bool) {
	val := reflect.ValueOf(arg)
	return val, val.Kind() == kind
}

// ReflectArray uses reflect to Array
func ReflectArray(arg interface{}) (out Array, ok bool) {
	var val reflect.Value
	if val, ok = Reflect(arg, Slice); ok {
		len := val.Len()
		out = make(Array, len)
		for i := 0; i < len; i++ {
			out[i] = val.Index(i).Interface()
		}
	}
	return
}

// ReflectDict uses reflect to Dict
func ReflectDict(arg interface{}) (out Dict, ok bool) {
	var val reflect.Value
	if val, ok = Reflect(arg, reflect.Map); ok {
		out = make(Dict)
		for _, k := range val.MapKeys() {
			out[k.Interface()] = val.MapIndex(k).Interface()
		}
	}
	return
}
