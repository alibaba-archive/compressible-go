package compressible_test

import (
	"fmt"

	compressible "github.com/teambition/compressible-go"
)

func ExampleDefaultCompressible() {
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
}
