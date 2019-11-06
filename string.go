package cast

import "strconv"

// String cast any value to string
func String(arg interface{}) string {
	switch v := arg.(type) {
	case string:
		return v
	case IStringer:
		return v.String()
	case []byte:
		return StringBytes(v)
	case error:
		return v.Error()
	case bool:
		return StringB(v)
	case int:
		return StringI(v)
	case uint:
		return StringUI(v)
	case int64:
		return StringI64(v)
	case float64:
		return StringF(v)
	default:
		return ReflectString(v)
	}
}

// StringB casts bool to string
func StringB(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// StringI casts int to string
func StringI(i int) string {
	return strconv.Itoa(i)
}

// StringUI casts uint to string
func StringUI(ui uint) string {
	return StringI(int(ui))
}

// StringI64 casts int64 to string
func StringI64(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}

// StringF casts float64 to string
func StringF(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
