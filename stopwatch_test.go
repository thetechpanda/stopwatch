package stopwatch_test

import (
	"testing"
	"time"

	"github.com/thetechpanda/stopwatch"
)

func TestStopWatch(t *testing.T) {
	time.Sleep(99 * time.Millisecond)
	if elapsed := stopwatch.Elapsed(); elapsed == 0 {
		t.Fatal("time isn't passing")
	} else {
		t.Logf("elapsed: %s", elapsed.String())
	}

	sw := stopwatch.New()
	time.Sleep(99 * time.Millisecond)
	if elapsed := sw.Elapsed(); elapsed == 0 {
		t.Fatal("time isn't passing")
	} else {
		t.Logf("elapsed: %s", elapsed.String())
	}
}
