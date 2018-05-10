package chronos

import "time"

// Clock encapsulates all functions exposed by the standard time package.
type Clock interface {
	// Now returns the current local time.
	Now() time.Time
	// Sleep pauses the current goroutine for at least the duration d.
	// A negative or zero duration causes Sleep to return immediately.
	Sleep(d time.Duration)
	// After waits for the duration to elapse and then sends the current time
	// on the returned channel.
	After(d time.Duration) <-chan time.Time
	// AfterFunc waits for the duration to elapse and then calls f
	// in its own goroutine. It returns a Timer that can
	// be used to cancel the call using its Stop method.
	AfterFunc(d time.Duration, f func()) Timer
	// NewTimer creates a new Timer that will send
	// the current time on its channel after at least duration d.
	NewTimer(d time.Duration) Timer
	// NewTicker returns a new Ticker containing a channel that will send the
	// time with a period specified by the duration argument.
	NewTicker(d time.Duration) Ticker
}

// The Timer type represents a single event.
// When the Timer expires, the current time will be sent on C.
type Timer interface {
	C() <-chan time.Time
	// Stop prevents the Timer from firing.
	// It returns true if the call stops the timer, false if the timer has already
	// expired or been stopped.
	Stop() bool
	// Reset changes the timer to expire after duration d.
	// It returns true if the timer had been active, false if the timer had
	// expired or been stopped.
	Reset(d time.Duration) bool
}

// NewTicker returns a new Ticker containing a channel that will send the
// time with a period specified by the duration argument.
type Ticker interface {
	C() <-chan time.Time
	// Stop turns off a ticker. After Stop, no more ticks will be sent.
	// Stop does not close the channel, to prevent a read from the channel succeeding
	// incorrectly.
	Stop()
}
