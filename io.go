package cast

import "io"

// Copy is io.Copy
func Copy(dst Writer, src Reader) (written int64, err error) {
	return io.Copy(dst, src)
}

// EOF = io.EOF
var EOF = io.EOF

// Reader = io.Reader
type Reader = io.Reader

// Writer = io.Writer
type Writer = io.Writer

// ReadWriter = io.ReadWriter
type ReadWriter = io.ReadWriter
