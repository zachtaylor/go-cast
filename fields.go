package cast

import "fmt"

// Fields is basic KV that attaches package-level conversion funcs
type Fields map[string]interface{}

// Val returns the value for the key
//
// To check for existence specifically, use map indexing instead of this func
func (f Fields) Val(k string) interface{} {
	return f[k]
}

// Sval returns a string wrapped from Val return
func (f Fields) Sval(k string) string {
	return String(f[k])
}

// Ival returns an int wrapped from Val return
func (f Fields) Ival(k string) int {
	return Int(f[k])
}

// Bval returns a bool wrapped from Val return
func (f Fields) Bval(k string) bool {
	return Bool(f[k])
}

func (f Fields) String() string {
	return fmt.Sprintf("%v", map[string]interface{}(f))
}

// Copy returns a new Fields that is shallow-copied
func (f Fields) Copy() Fields {
	Fields := Fields{}
	for k, v := range f {
		Fields[k] = v
	}
	return Fields
}
