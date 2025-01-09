# vector-go
[![Go Reference](https://pkg.go.dev/badge/github.com/HyperCodec/vector-go.svg)](https://pkg.go.dev/github.com/HyperCodec/vector-go)

A variable-length collection datatype that allocates as needed. It is based on the Rust implementation and aims to provide an efficient,
high-level API.

### Installation
To add this as a dependency, simply run `go get github.com/HyperCodec/vector-go`.

### Basic Example
```go
package main

import (
	"fmt"

	"github.com/HyperCodec/vector-go"
)

func main() {
    // create a vector with an initial capacity of 3 and an allocation amount of 5.
    v, _ := vector.EmptyWithCapacity[int](3, 5)

    v.PushBack(1)
    v.PushBack(2)
    v.PushBack(3)

    fmt.Println(v.Data())
}
```

### License
This project uses the `MIT` license.