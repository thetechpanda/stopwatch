package stopwatch_test

import (
	"fmt"
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
