package cast

import (
	"fmt"
	"strconv"
)

// Int cast any value to int
func Int(any interface{}) int {
	switch v := any.(type) {
	case int:
		return v
	case uint:
		return int(v)
	case int64:
		return int(v)
	case float64:
		return int(v)
	case string:
		return IntS(v)
	case fmt.Stringer:
		return IntS(v.String())
	case bool:
		if v {
			return 1
		}
		return -1
	default:
		return 0
	}
}

// IntS casts string to int
func IntS(s string) int {
	i, _ := strconv.ParseInt(s, 10, 0)
	return int(i)
}
