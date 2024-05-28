# Stopwatch

![Test](https://github.com/thetechpanda/mutex/actions/workflows/go.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thetechpanda/stopwatch)](https://goreportcard.com/report/github.com/thetechpanda/stopwatch)
[![Go Reference](https://pkg.go.dev/badge/github.com/thetechpanda/stopwatch.svg)](https://pkg.go.dev/github.com/thetechpanda/stopwatch)
[![Release](https://img.shields.io/github/release/thetechpanda/stopwatch.svg?style=flat-square)](https://github.com/thetechpanda/stopwatch/releases)
![Dependencies](https://img.shields.io/badge/Go_Dependencies-_None_-green.svg)

A dead simple stop watch.

## Key Features

It's easy, it also offers a globally available stop watch instance initialised at package load.

## Motivation

Needed a simple stop watch.

## Usage

```go
package main

import (
	"fmt"

	"github.com/thetechpanda/stopwatch"
)

func main() {
    fmt.Printf("time: %s\n", stopwatch.Elapsed().String())
}

```

## Installation

```bash
go get github.com/thetechpanda/stopwatch
```

## Contributing

Contributions are welcome and very much appreciated!

Feel free to open an issue or submit a pull request.

## License

The `stopwatch` package is released under the MIT License. See the [LICENSE](LICENSE) file for details.
