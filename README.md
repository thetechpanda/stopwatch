# Stopwatch

![Test](https://github.com/thetechpanda/mutex/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thetechpanda/stopwatch)](https://goreportcard.com/report/github.com/thetechpanda/stopwatch)
[![Go Reference](https://pkg.go.dev/badge/github.com/thetechpanda/stopwatch.svg)](https://pkg.go.dev/github.com/thetechpanda/stopwatch)
[![Release](https://img.shields.io/github/release/thetechpanda/stopwatch.svg?style=flat-square)](https://github.com/thetechpanda/stopwatch/releases)
![Dependencies](https://img.shields.io/badge/Go_Dependencies-_None_-green.svg)

A dead simple `stopwatch`, with a sprinkle of `clock`.

## Key Features

Stopwatch it's easy, offers a globally available stop watch instance initialised at package load.

Clock it's easy too, but bare bone. Is not very precise.

## Motivation

Needed a simple `stopwatch`. Then I didn't know where to put `clock`.

## Usage

```go
package main

import (
	"fmt"

	"github.com/thetechpanda/stopwatch"
)

func main() {
	// stopwatch
    fmt.Printf("time: %s\n", stopwatch.Elapsed().String())

	// clock
	clock := stopwatch.NewClock(time.Second)
	for i := range 1000 {
		go func() {
			<-clock.Tick()
			fmt.Printf("hi from %d\n", i)
		}()
	}
	// enjoy the flood
}

```

## Installation

```bash
go get github.com/thetechpanda/stopwatch
```

## Todo

- [ ] Clock benchmarks (how do I even benchmark it?)

## Contributing

Contributions are welcome and very much appreciated!

Feel free to open an issue or submit a pull request.

## License

The `stopwatch` package is released under the MIT License. See the [LICENSE](LICENSE) file for details.
