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

// ReadWriter = io.ReadWriter
type ReadWriter = io.ReadWriter

// ReadCloser = io.ReadCloser
type ReadCloser = io.ReadCloser

// ReadWriteCloser = io.ReadWriteCloser
type ReadWriteCloser = io.ReadWriteCloser

// Writer = io.Writer
type Writer = io.Writer

// Closer = io.Closer
type Closer = io.Closer

// WriteCloser = io.WriteCloser
type WriteCloser = io.WriteCloser

func WriteBytes(w Writer, b []byte) (int, error) {
	return w.Write(b)
}

func WriteString(w Writer, s string) (int, error) {
	return WriteBytes(w, BytesS(s))
}

func WriteN(w Writer, args ...interface{}) (int, error) {
	return WriteString(w, StringN(args...))
}
