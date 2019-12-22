package cast

import "bytes"

// Buffer = bytes.Buffer
type Buffer = bytes.Buffer

// NewBuffer uses bytes.NewBufferString
func NewBuffer(s string) *Buffer {
	return bytes.NewBufferString(s)
}

// StringBytes casts []byte to string
//
// outperforms `string(b)` by ~95%
func StringBytes(b []byte) string {
	return *(*string)(Pointer(&b))
}

// BytesS casts string to []byte
//
// outperforms `[]byte(s)` by ~95%
func BytesS(s string) []byte {
	return *(*[]byte)(Pointer(&s))
}
