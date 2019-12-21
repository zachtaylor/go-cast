package cast

import "strconv"

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
