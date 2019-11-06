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

// Stringer casts any value to fmt.Stringer
func Stringer(any interface{}) IStringer {
	return x{any}
}

type x struct {
	x interface{}
}

func (x x) String() string {
	return String(x.x)
}

// StringerB casts bool to IStringer
type StringerB bool

func (b StringerB) String() string {
	if b {
		return "true"
	}
	return "false"
}

// StringerE casts error to IStringer
func StringerE(err error) IStringer {
	return StringerFunc(func() string {
		return err.Error()
	})
}

// StringerFunc casts func to IStringer
type StringerFunc func() string

func (f StringerFunc) String() string {
	return f()
}

// StringerI casts int to IStringer
type StringerI int

func (i StringerI) String() string {
	return StringI(int(i))
}

// StringerS casts string to IStringer
type StringerS string

func (s StringerS) String() string {
	return string(s)
}
