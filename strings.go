package cast

import "strings"

// StringBuilder is an alias for strings.Builder
type StringBuilder = strings.Builder

var replacer = strings.NewReplacer(`\`, `\\`, `"`, `\"`)

// Contains uses strings.Contains
func Contains(a, b string) bool { return strings.Contains(a, b) }

// InCharset returns len(Trim()) < 1
func InCharset(a, b string) bool { return len(Trim(a, b)) < 1 }

// EscapeString escapes '\' and '"' characters
func EscapeString(s string) string { return replacer.Replace(s) }

// LastIndex is an alias for strings.LastIndex
func LastIndex(a, b string) int { return strings.LastIndex(a, b) }

// Trim uses strings.Trim
func Trim(a, b string) string { return strings.Trim(a, b) }

// OutCharset returns len(Trim()) > 0
func OutCharset(a, b string) bool { return len(Trim(a, b)) > 0 }
