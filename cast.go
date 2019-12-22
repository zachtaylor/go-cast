package cast // import "ztaylor.me/cast"

// poolStringBuilder is a global variable for pooling StringBuilder
var poolStringBuilder = Pool{
	New: func() interface{} {
		return &StringBuilder{}
	},
}

// Pool StringBuilder
//
// sb := poolStringBuilder.Get().(StringBuilder)
// ...
// sb.Reset()
// poolStringBuilder.Put(sb)
func init() {
	// pool first string builder in package scope
	poolStringBuilder.Put(poolStringBuilder.New())
}
