package stopwatch

import (
	"sync"
	"time"
)

// Clock is a ticker that can be listened by multiple go routine. Does so by closing the receiving channel.
// Clock is not very precise, but it is ok for the synchronisation of a few goroutine.
type Clock struct {
	ticker *time.Ticker
	mu     sync.Mutex
	ticks  chan time.Time
}

// NewClock creates a clock
func NewClock(rate time.Duration) *Clock {
	c := &Clock{
		ticker: time.NewTicker(rate),
		ticks:  make(chan time.Time),
	}
	c.ticks = make(chan time.Time)
	go func() {
		for {
			<-c.ticker.C
			c.mu.Lock()
			close(c.ticks)
			c.ticks = make(chan time.Time)
			c.mu.Unlock()
		}
	}()
	return c
}

// Stop stops the clock, after that no more ticks are sent.
func (c *Clock) Stop() {
	c.ticker.Stop()
}

// Tick returns a receive only channel that will be closed on the next tick.
// The channel returned by this function should not be passed around.
// The only action performed on the channel is close().
func (c *Clock) Tick() <-chan time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ticks
}

// Reset stops a ticker and resets its period to the specified duration.
// The next tick will arrive after the new period elapses.
func (c *Clock) Reset(rate time.Duration) {
	c.ticker.Reset(rate)
}
