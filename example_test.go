package compressible_test

import (
	"fmt"

	"github.com/teambition/compressible-go"
	"github.com/teambition/gear"
	"github.com/teambition/gear/middleware/static"
)

func Example() {
	fmt.Println(compressible.Is("text/html"))
	// -> true

	fmt.Println(compressible.Is("image/png"))
	// -> false

	var wt compressible.WithThreshold = 1024

	fmt.Println(wt.Compressible("text/html", 1023))
	// -> false
}

func ExampleWithGear() {
	app := gear.New()
	app.Set("AppCompress", compressible.WithThreshold(1024))

	// Add a static middleware
	app.Use(static.New(static.Options{
		Root:   "./",
		Prefix: "/",
	}))
	app.Error(app.Listen(":3000")) // http://127.0.0.1:3000/
}
