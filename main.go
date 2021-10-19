package main

import (
	"os"
	"github.com/juliecoding/svg-cli/cli"
)

func main() {
    os.Exit(cli.Run(os.Args[1:]))
}
