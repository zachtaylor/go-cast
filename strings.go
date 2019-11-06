package cast

import "strings"

// StringBuilder is an alias for strings.Builder
type StringBuilder = strings.Builder

// // ReplaceAll is an alias for strings.ReplaceAll
// func ReplaceAll(s, old, new string) string {
// 	return strings.ReplaceAll(s, old, new)
// }

var replacer = strings.NewReplacer(`\`, `\\`, `"`, `\"`)

// EscapeString escapes '\' and '"' characters
func EscapeString(s string) string {
	return replacer.Replace(s)
}
