package cast // import "ztaylor.me/cast"

// poolStringBuilder is a global variable for pooling StringBuilder
var poolStringBuilder = Pool{
	New: func() interface{} {
		return StringBuilder{}
	},
}

func init() {
	// just gonna pool a string builder in package scope
	poolStringBuilder.Put(poolStringBuilder.New())
	poolStringBuilder.Put(poolStringBuilder.New())
}
