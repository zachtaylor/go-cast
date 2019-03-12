package cast

import "fmt"

// Stringer casts any value to fmt.Stringer
func Stringer(any interface{}) fmt.Stringer {
	return x{any}
}

type x struct {
	x interface{}
}

func (x x) String() string {
	return String(x.x)
}

// StringerB casts bool to fmt.Stringer
type StringerB bool

func (b StringerB) String() string {
	if b {
		return "true"
	}
	return "false"
}

// StringerE casts error to fmt.Stringer
func StringerE(err error) fmt.Stringer {
	return StringerFunc(func() string {
		return err.Error()
	})
}

// StringerFunc casts func to fmt.Stringer
type StringerFunc func() string

func (f StringerFunc) String() string {
	return f()
}

// StringerI casts int to fmt.Stringer
type StringerI int

func (i StringerI) String() string {
	return StringI(int(i))
}

// StringerS casts string to fmt.Stringer
type StringerS string

func (s StringerS) String() string {
	return string(s)
}
