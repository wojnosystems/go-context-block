# Overview

Helper methods for contexts.

This is just a small library of things I tend to write over and over again with contexts.

# Installation

`go get -u github.com/wojnosystems/go-context-block`

# Examples

```go
package main

import (
	"context"
	"github.com/wojnosystems/go-context-block/contextBlock"
)

func main() {
	contextBlock.WithTimeout(context.Background(), 10*time.Second, func(ctx context.Context) {
        // use the context in something that needs it
	})
}
```
