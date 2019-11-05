package cast_test

import "ztaylor.me/cast"

func MakeDataArray() []interface{} {
	return []interface{}{
		"hello world",
		"1",
		[]int{20, 19},
		map[interface{}]interface{}{
			"stringer": cast.Stringer("foobar stringer"),
		},
	}
}

func MakeDataMap() map[string]interface{} {
	return map[string]interface{}{
		"k1":   "v1",
		"k2":   "v2",
		"k3":   "v3",
		"i1":   1,
		"i2":   2,
		"i3":   3,
		"b1":   true,
		"b2":   false,
		"map":  map[string]interface{}{},
		"json": cast.JSON{},
		"dat1": "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"dat2": "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"dat3": "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
	}
}
