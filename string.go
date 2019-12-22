package cast

// String casts any value to string
func String(arg interface{}) (str string) {
	switch v := arg.(type) {
	case string:
		str = v
	case IStringer:
		str = v.String()
	case []byte:
		str = StringBytes(v)
	case error:
		str = v.Error()
	case bool:
		str = StringB(v)
	case int:
		str = StringI(v)
	case uint:
		str = StringUI(v)
	case int64:
		str = StringI64(v)
	case float64:
		str = StringF(v)
	default:
		str = Sprint(arg)
	}
	return
}

// StringN cast any number of values to string
func StringN(args ...interface{}) (str string) {
	if args == nil || len(args) < 1 {
		// skip
	} else if len(args) == 1 {
		str = String(args[0])
	} else {
		sb, first := poolStringBuilder.Get().(*StringBuilder), true
		sb.Grow(32 * len(args))
		for _, arg := range args {
			if first {
				first = false
			} else {
				sb.WriteByte(' ')
			}
			sb.WriteString(String(arg))
		}
		str = sb.String()
		sb.Reset()
		poolStringBuilder.Put(sb)
	}
	return
}

// StringB casts bool to string
func StringB(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
