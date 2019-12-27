package cast // import "ztaylor.me/cast"

var growFactor = 32 // tunes Pool StringBuilder preparatory Grow behavior

// poolStringBuilder is a global variable for pooling StringBuilder
//
// sb := poolStringBuilder.Get().(StringBuilder)
// ...
// sb.Reset()
// poolStringBuilder.Put(sb)
var poolStringBuilder = Pool{
	New: func() interface{} {
		return &StringBuilder{}
	},
}
