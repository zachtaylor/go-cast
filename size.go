package cast

// Sizer is used by the encoder to guess at the number of bytes to alloc
type Sizer interface {
	Size() int
}

// Size returns a guess at the number of bytes needed to encode a value for approximation
func Size(arg interface{}) (int int) {
	addSize(arg, &int)
	return
}

// addSize adds the size of the arg to the int pointer
func addSize(arg interface{}, num *int) {
	switch v := arg.(type) {
	case string:
		*num += len(v)
	case IStringer:
		*num += 32
	case []byte:
		*num += len(v)
	case error:
		*num += len(v.Error())
	case Sizer:
		*num += v.Size()
	case Array:
		v.addSize(num)
	case []interface{}:
		Array(v).addSize(num)
	case Dict:
		v.addSize(num)
	case map[interface{}]interface{}:
		Dict(v).addSize(num)
	case JSON:
		v.addSize(num)
	case map[string]interface{}:
		JSON(v).addSize(num)
	default:
		*num += 4
	}
}
