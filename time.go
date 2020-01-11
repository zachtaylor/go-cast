package cast

import "time"

// Duration = time.Duration
type Duration = time.Duration

// Ticker = time.Ticker
type Ticker = time.Ticker

// Time = time.Time
type Time = time.Time

// Timer = time.Timer
type Timer = time.Timer

// After returns time.After()
func After(d Duration) <-chan time.Time {
	return time.After(d)
}

// NewTicker creates a time.Ticker
func NewTicker(d Duration) *Ticker {
	return time.NewTicker(d)
}

// NewTimer creates a time.Timer
func NewTimer(d Duration) *Timer {
	return time.NewTimer(d)
}

// Now returns time.Now()
func Now() Time {
	return time.Now()
}

// Unix return time.Unix()
func Unix(sec int64, nsec int64) Time {
	return time.Unix(sec, nsec)
}

// Millisecond = time.Millisecond
var Millisecond = time.Millisecond

// Second = time.Second
var Second = time.Second

// Minute = time.Minute
var Minute = time.Minute

// Hour = time.Hour
var Hour = time.Hour
