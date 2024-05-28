package stopwatch

import "time"

// StopWatch is a simple stopwatch
type StopWatch struct {
	t0 time.Time
}

// New creates a new stopwatch
func New() *StopWatch {
	return &StopWatch{
		t0: time.Now(),
	}
}

// Elapsed returns the elapsed time
func (sw *StopWatch) Elapsed() time.Duration {
	return time.Since(sw.t0)
}

// Reset resets the stopwatch
func (sw *StopWatch) Reset() time.Duration {
	d := time.Since(sw.t0)
	sw.t0 = time.Now()
	return d
}

// String returns the elapsed time as a string
func (sw *StopWatch) String() string {
	return sw.Elapsed().String()
}

var stopWatch = New()

// Elapsed returns the elapsed time.
// The stopwatch used is a global instance and starts when the package is loaded.
func Elapsed() time.Duration {
	return stopWatch.Elapsed()
}

// Reset resets the stopwatch and returns the elapsed time.
func Reset() time.Duration {
	return stopWatch.Reset()
}

// String returns the elapsed time as a string
func String() string {
	return stopWatch.String()
}
