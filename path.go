package cast

import "path"

// SplitPath uses path.Split
func SplitPath(s string) (dir, file string) { return path.Split(s) }
