package cast

import "time"

// Duration = time.Duration
type Duration = time.Duration

// Time = time.Time
type Time = time.Time

// After returns time.After()
func After(d Duration) <-chan time.Time { return time.After(d) }

// Now returns time.Now()
func Now() Time { return time.Now() }

// Unix return time.Unix()
func Unix(sec int64, nsec int64) Time { return time.Unix(sec, nsec) }

// Millisecond = time.Millisecond
var Millisecond = time.Millisecond

// Second = time.Second
var Second = time.Second

// Minute = time.Minute
var Minute = time.Minute

// Hour = time.Hour
var Hour = time.Hour
