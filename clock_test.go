package stopwatch_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/thetechpanda/stopwatch"
)

func TestClock(t *testing.T) {
	rate := 15 * time.Millisecond
	var wg sync.WaitGroup
	var total atomic.Int64
	var loops, hoops int64 = 500, 500 // runs 500 times, raises 500 listener each time

	clock := stopwatch.NewClock(rate)
	hoop := func(then time.Time) {
		defer wg.Done()
		<-clock.Tick()
		total.Add(int64(time.Since(then)))
	}
	for range loops {
		then := time.Now()
		for range hoops {
			wg.Add(1)
			go hoop(then)
		}
		wg.Wait()
	}
	clock.Stop()
	avg := float64(total.Load() / loops / hoops)
	ratio := 1.1 // 10% = all receiver registered a tick up to +10% after
	if avg > float64(rate)*ratio {
		t.Fatalf("average registered clock rate is higher than the rate +10%% = %s", time.Duration(avg))
	}
	t.Logf("average registered clock rate = %s", time.Duration(avg))
}
