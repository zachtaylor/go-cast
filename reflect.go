package cast

import "reflect"

// Reflect uses reflect to any kind, but it always finishes, so you must check kind with ok
func Reflect(arg interface{}, kind reflect.Kind) (reflect.Value, bool) {
	val := reflect.ValueOf(arg)
	return val, val.Kind() == kind
}

// ReflectArray uses reflect to Array
func ReflectArray(arg interface{}) (out Array, ok bool) {
	var val reflect.Value
	if val, ok = Reflect(arg, reflect.Slice); ok {
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
