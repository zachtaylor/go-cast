package cast

import "fmt"

// IStringer = fmt.Stringer
type IStringer = fmt.Stringer

// Sprint is a shortcut for fmt.Sprint
func Sprint(a ...interface{}) string {
	return fmt.Sprint(a...)
}

// Sprintf is a shortcut for fmt.Sprintf
func Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// Fprintf is a shortcut for fmt.Fprintf
func Fprintf(w Writer, format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(w, format, args...)
}
