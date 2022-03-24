package cli

import (
	"flag"
	"strings"
)

type appConfig struct {
	selected selectedFilters
	in	string
	out	string
}

// Because this is a custom flag Var, we have to implement String and Set methods
type selectedFilters []string

func (s *selectedFilters) String() string {
	return strings.Join(*s, " ")
}

func (s *selectedFilters) Set(str string) error {
	if str == "" {
		return nil
	}
	spl := strings.Split(str, " ")
	*s = append(*s, spl...)
	return nil
}

func getConfig() appConfig {
	var ac appConfig
	flag.StringVar(&ac.out, "out", "./out/out.svg", "Path to SVG output")
	flag.StringVar(&ac.in, "in", "", "Path to access input image")
	flag.Var(&ac.selected, "filters", "Name of filters from the filters.json file to apply to the input file, separated by spaces.\nPossible values are: blur, bw, carlton, desaturate, day, fuzzyTv, ginza, hueRotate, instagram, matrix, montyPython, dusk, pointLight, saturate, sepia, sunshine")
	// After all flags are defined, calling Parse parses the command line input into the defined flags.
	flag.Parse()
	ac.in = getValidInput(ac.in)
	ac.out = getValidOutput(ac.out)
	// ac.selected = getValidSelected(ac.selected)

	return ac
}
