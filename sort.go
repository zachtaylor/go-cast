package cast

import "sort"

// ArrayI = sort.IntSlice
type ArrayI = sort.IntSlice

// ArrayS = sort.StringSlice
type ArrayS = sort.StringSlice

// SortInts uses sort.Ints
func SortInts(i []int) {
	sort.Ints(i)
}

// SortSlice uses sort.Slice
func SortSlice(slice interface{}, less func(i int, j int) bool) {
	sort.Slice(slice, less)
}

// SortStrings uses sort.Strings
func SortStrings(s []string) {
	sort.Strings(s)
}
