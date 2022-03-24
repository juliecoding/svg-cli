package cli

import (
	"flag"
	"strings"
)

type appConfig struct {
	in			string
	out			string
	selectedStr	string
	selected	[]string
}

func getConfig() appConfig {
	outDesc			:= "Path to SVG output"
	inDesc			:= "Path to access input image"
	selectedStrDesc := "Name of filters from the filters.json file to apply to the input file, separated by spaces.\nPossible values are: blur, bw, carlton, desaturate, day, fuzzyTv, ginza, hueRotate, instagram, matrix, montyPython, dusk, pointLight, saturate, sepia, sunshine"

	var ac appConfig
	flag.StringVar(&ac.out, "out", "./out/out.svg", outDesc)
	flag.StringVar(&ac.in, "in", "", inDesc)
	flag.StringVar(&ac.selectedStr, "filters", "", selectedStrDesc)

	// After all flags are defined, calling Parse parses the command line input into the defined flags.
	flag.Parse()

	ac.in = getValidInput(ac.in)
	ac.out = getValidOutput(ac.out)
	ac.selected = confirmSelectedFilters(strings.Split(ac.selectedStr, " "))

	return ac
}
