package cast

import "fmt"

// Bool cast any value to bool
func Bool(any interface{}) bool {
	switch v := any.(type) {
	case nil:
		return false
	case bool:
		return v
	case string:
		return BoolS(v)
	case fmt.Stringer:
		return BoolS(v.String())
	case []byte:
		return len(v) > 0
	case error:
		return v != nil
	case int:
		return v > 0
	case uint:
		return v > 0
	case float64:
		return v > 0
	default:
		return false
	}
}

// BoolS casts string to bool
func BoolS(s string) bool {
	return s != "" && s != "false"
}
