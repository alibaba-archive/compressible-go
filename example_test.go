package compressible_test

import (
	"fmt"

	compressible "github.com/teambition/compressible-go"
)

func Example() {
	fmt.Println(compressible.Is("text/html"))
	// -> true

	fmt.Println(compressible.Is("image/png"))
	// -> false

	var wt compressible.WithTrashold = 1024

	fmt.Println(wt.Compressible("text/html", 1023))
	// -> false
}
