package cast // import "ztaylor.me/cast"

import "reflect"

func cast(arg interface{}, kind reflect.Kind) (reflect.Value, bool) {
	val := reflect.ValueOf(arg)
	return val, val.Kind() == kind
}

func castSlice(arg interface{}) (out []interface{}, ok bool) {
	var slice reflect.Value
	if slice, ok = cast(arg, reflect.Slice); ok {
		len := slice.Len()
		out = make([]interface{}, len)
		for i := 0; i < len; i++ {
			out[i] = slice.Index(i).Interface()
		}
	}
	return
}
