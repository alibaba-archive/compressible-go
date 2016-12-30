# compressible-go
[![Build Status](https://travis-ci.org/teambition/compressible-go.svg?branch=master)](https://travis-ci.org/teambition/compressible-go)
[![Coverage Status](https://coveralls.io/repos/github/teambition/compressible-go/badge.svg?branch=master)](https://coveralls.io/github/teambition/compressible-go?branch=master)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/teambition/compressible-go/master/LICENSE)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/teambition/compressible-go)

Compressible Content-Type / mime checking for Go.

## Installation

```
go get -u github.com/teambition/compressible-go
```

## Demo

```go
import (
  "github.com/teambition/compressible-go"
)

fmt.Println(compressible.Is("text/html"))
// -> true

fmt.Println(compressible.Is("image/png"))
// -> false

var wt compressible.WithTrashold = 1024

fmt.Println(wt.Compressible("text/html", 1023))
// -> false
```

**Work with gear:**
```go
package main

import (
	"github.com/teambition/compressible-go"
	"github.com/teambition/gear"
	"github.com/teambition/gear/middleware/static"
)

func main() {
	app := gear.New()
	app.Set("AppCompress", compressible.WithThreshold(1024))

	// Add a static middleware
	app.Use(static.New(static.Options{
		Root:   "./",
		Prefix: "/",
	}))
	app.Error(app.Listen(":3000")) // http://127.0.0.1:3000/
}
```

## Documentation

The docs can be found at [godoc.org](https://godoc.org/github.com/teambition/compressible-go), as usual.

## License

MIT
