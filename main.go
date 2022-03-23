package main

import (
	"os"

	// Decoding any particular image format requires the prior registration of a decoder function.
	// Registration is typically automatic as a side effect of initializing that format's package,
	// so here we're using _ to import the package purely for its initialization side effects.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/juliecoding/svg-cli/cli"
)

func main() {
	// Conventionally, for os.Exit, code zero indicates success,
	// non-zero an error.
    os.Exit(cli.Run())
}
