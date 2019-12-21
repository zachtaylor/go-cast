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
