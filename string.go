package cast

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
		return Sprint(arg)
	}
}

// StringB casts bool to string
func StringB(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
