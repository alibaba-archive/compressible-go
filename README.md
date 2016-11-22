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
  compressible "github.com/teambition/compressible-go"
)
```

```go
fmt.Println(compressible.Is("text/html"))
// -> true

fmt.Println(compressible.Is("image/png"))
// -> false

var wt compressible.WithTrashold = 1024

fmt.Println(wt.Compressible("text/html", 1023))
// -> false
```

## Documentation

The docs can be found at [godoc.org](https://godoc.org/github.com/teambition/compressible-go), as usual.

## License

MIT
