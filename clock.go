package stopwatch

import (
	"sync"
	"time"
)

// Clock is a ticker that can be listened by multiple go routine. Does so by closing the receiving channel.
// Clock is not very precise, performance will decrease as listeners increase or rate gets lower than 20ms (apple m2 max tested).
type Clock struct {
	ticker *time.Ticker
	mu     sync.Mutex
	ticks  chan struct{}
}

// NewClock creates a clock
func NewClock(rate time.Duration) *Clock {
	c := &Clock{
		ticker: time.NewTicker(rate),
		ticks:  make(chan struct{}),
	}
	c.ticks = make(chan struct{})
	go func() {
		for {
			<-c.ticker.C
			c.mu.Lock()
			close(c.ticks)
			c.ticks = make(chan struct{})
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
func (c *Clock) Tick() <-chan struct{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ticks
}

// Reset stops a ticker and resets its period to the specified duration.
// The next tick will arrive after the new period elapses.
func (c *Clock) Reset(rate time.Duration) {
	c.ticker.Reset(rate)
}
