package stopwatch_test

import (
	"fmt"
	"sync"
	"time"

	"github.com/thetechpanda/stopwatch"
)

func ExampleStopWatch() {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("elapsed: %s", stopwatch.Elapsed().String())

	sw := stopwatch.New()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("elapsed: %s", sw.Elapsed().String())
}

func ExampleClock() {
	var wg sync.WaitGroup
	clock := stopwatch.NewClock(time.Second)
	for i := range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-clock.Tick()
			fmt.Printf("hi from %d\n", i)
		}()
	}
	wg.Wait()
	clock.Stop()
}
