package cast

import "strings"

// StringBuilder is an alias for strings.Builder
type StringBuilder = strings.Builder

// StringReader = strings.Reader
type StringReader = strings.Reader

// Replacer = strings.Replacer
type Replacer = strings.Replacer

// NewReplacer usesstrings.Replace
func NewReplacer(oldnew ...string) *Replacer { return strings.NewReplacer(oldnew...) }

// Contains uses strings.Contains
func Contains(a, b string) bool { return strings.Contains(a, b) }

// InCharset returns len(Trim()) < 1
func InCharset(a, b string) bool { return len(Trim(a, b)) < 1 }

// Split is an alias for strings.Split
func Split(s, sep string) []string { return strings.Split(s, sep) }

// replacer is used by EscapeString
var replacer = NewReplacer(`\`, `\\`, `"`, `\"`)

// EscapeString escapes '\' and '"' characters
func EscapeString(s string) string { return replacer.Replace(s) }

// LastIndex is an alias for strings.LastIndex
func LastIndex(a, b string) int { return strings.LastIndex(a, b) }

// Trim uses strings.Trim
func Trim(a, b string) string { return strings.Trim(a, b) }

// OutCharset returns len(Trim()) > 0
func OutCharset(a, b string) bool { return len(Trim(a, b)) > 0 }

// NewReader uses strings.NewReader
func NewReader(s string) *StringReader { return strings.NewReader(s) }
