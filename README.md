# compressible-go
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
d := compressible.Default{}

d.SetTrashold(1024)

fmt.Println(d.Compressible("text/html", 1024))
// -> true

fmt.Println(d.Compressible("text/html", 1000))
// -> false

fmt.Println(d.Compressible("image/png", 1024))
// -> false

fmt.Println(d.Compressible("text/boobar", 1024))
// -> true
```

## Documentation

The docs can be found at [godoc.org](https://godoc.org/github.com/teambition/compressible-go), as usual.

## License

MIT
