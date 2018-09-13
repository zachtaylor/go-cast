package cast

import (
	"fmt"
	"strconv"
)

func i2s(i int) string {
	return strconv.Itoa(i)
}
func ui2s(ui uint) string {
	return i642s(int64(ui))
}
func i642s(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}
func f2s(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
func s2i(s string) int {
	i, _ := strconv.ParseInt(s, 10, 0)
	return int(i)
}

// Stringer is alias for fmt.Stringer
type Stringer = fmt.Stringer

// String cast any value to string
func String(any interface{}) string {
	switch v := any.(type) {
	case string:
		return v
	case Stringer:
		return v.String()
	case int:
		return i2s(v)
	case uint:
		return ui2s(v)
	case int64:
		return i642s(v)
	case float64:
		return f2s(v)
	default:
		return ""
	}
}

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
		return s2i(v)
	default:
		return 0
	}
}

// Bool cast any value to bool
func Bool(any interface{}) bool {
	switch v := any.(type) {
	case bool:
		return v
	case int:
		return v > 0
	case float64:
		return v > 0
	case string:
		return v != ""
	default:
		return false
	}
}
