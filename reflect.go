package cast

import "reflect"

// Kind = reflect.Kind
type Kind = reflect.Kind

// Value = reflect.Value
type Value = reflect.Value

// Slice = reflect.Slice
const Slice = reflect.Slice

// Map = reflect.Map
const Map = reflect.Map

// Type = reflect.Type
type Type = reflect.Type

// TypeOf is an alias for reflect.TypeOf
func TypeOf(arg interface{}) Type { return reflect.TypeOf(arg) }

// KindOf is an alias for TypeOf().Kind()
func KindOf(arg interface{}) Kind { return TypeOf(arg).Kind() }

// Reflect uses reflect to any kind
func Reflect(arg interface{}) Value { return reflect.ValueOf(arg) }

// ReflectArray uses reflect to Array
func ReflectArray(arg interface{}) (out Array, ok bool) {
	if val := Reflect(arg); val.Kind() == Slice {
		len := val.Len()
		out = make(Array, len)
		for i := 0; i < len; i++ {
			out[i] = val.Index(i).Interface()
		}
		ok = true
	}
	return
}

// ReflectDict uses reflect to Dict
func ReflectDict(arg interface{}) (out Dict, ok bool) {
	if val := Reflect(arg); val.Kind() == Map {
		out = make(Dict)
		for _, k := range val.MapKeys() {
			out[k.Interface()] = val.MapIndex(k).Interface()
		}
		ok = true
	}
	return
}
