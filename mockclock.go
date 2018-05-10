package chronos

import "time"

// MockClock exposes configuration points for all functions provided by the Clock interface.
type MockClock struct {
	Now       MockNow
	Sleep     MockSleep
	After     MockAfter
	AfterFunc MockAfterFunc
	Timer     MockTimer
	Ticker    MockTicker
}

// NewMock returns a mock instance that can be used to mock the Clock interface's methods.
func NewMock() *MockClock {
	return &MockClock{}
}

// Clock returns a Clock instance that can be injected wherever a mock is expected.
func (mock *MockClock) Clock() Clock {
	return &mockClock{mock}
}

type mockClock struct {
	mock *MockClock
}

func (clock *mockClock) Now() time.Time {
	return clock.mock.Now.patch()
}

func (clock *mockClock) Sleep(d time.Duration) {
	clock.mock.Sleep.patch(d)
}

func (clock *mockClock) After(d time.Duration) <-chan time.Time {
	return clock.mock.After.patch(d)
}

func (clock *mockClock) AfterFunc(d time.Duration, f func()) Timer {
	return clock.mock.AfterFunc.patch(d, f)
}

func (clock *mockClock) NewTimer(d time.Duration) Timer {
	return clock.mock.Timer.patch(d)
}

func (clock *mockClock) NewTicker(d time.Duration) Ticker {
	return clock.mock.Ticker.patch(d)
}

type MockNow struct {
	patch func() time.Time
}

// Patch patches the Now function.
func (mock *MockNow) Patch(patch func() time.Time) {
	mock.patch = patch
}

type MockSleep struct {
	patch func(time.Duration)
}

// Patch patches the Sleep function.
func (mock *MockSleep) Patch(patch func(time.Duration)) {
	mock.patch = patch
}

type MockAfter struct {
	patch func(time.Duration) <-chan time.Time
}

// Patch patches the After function.
func (mock *MockAfter) Patch(patch func(time.Duration) <-chan time.Time) {
	mock.patch = patch
}

type MockAfterFunc struct {
	patch func(time.Duration, func()) Timer
}

// Patch patches the AfterFunc function.
func (mock *MockAfterFunc) Patch(patch func(time.Duration, func()) Timer) {
	mock.patch = patch
}

type MockTimer struct {
	patch func(time.Duration) Timer
}

// Patch patches the NewTimer function.
func (mock *MockTimer) Patch(patch func(time.Duration) Timer) {
	mock.patch = patch
}

type MockTicker struct {
	patch func(time.Duration) Ticker
}

// Patch patches the NewTicker function.
func (mock *MockTicker) Patch(patch func(time.Duration) Ticker) {
	mock.patch = patch
}
