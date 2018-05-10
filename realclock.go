package chronos

import "time"

type realClock struct{}

// NewClock returns a clock that mimics the time package's functions.
func NewClock() Clock {
	return &realClock{}
}

func (*realClock) Now() time.Time {
	return time.Now()
}

func (*realClock) Sleep(d time.Duration) {
	time.Sleep(d)
}

func (*realClock) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

func (*realClock) AfterFunc(d time.Duration, f func()) Timer {
	return &realTimer{time.AfterFunc(d, f)}
}

func (*realClock) NewTimer(d time.Duration) Timer {
	return &realTimer{time.NewTimer(d)}
}

func (*realClock) NewTicker(d time.Duration) Ticker {
	return &realTicker{time.NewTicker(d)}
}

type realTimer struct {
	innerTimer *time.Timer
}

func (timer *realTimer) C() <-chan time.Time {
	return timer.innerTimer.C
}

func (timer *realTimer) Stop() bool {
	return timer.innerTimer.Stop()
}

func (timer *realTimer) Reset(d time.Duration) bool {
	return timer.innerTimer.Reset(d)
}

type realTicker struct {
	innerTicker *time.Ticker
}

func (ticker *realTicker) C() <-chan time.Time {
	return ticker.innerTicker.C
}

func (ticker *realTicker) Stop() {
	ticker.innerTicker.Stop()
}
