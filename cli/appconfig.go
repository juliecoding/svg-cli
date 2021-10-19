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
	spl := strings.Split(str, " ")
	*s = append(*s, spl...)
	return nil
}


func getConfig() appConfig {
	var ac appConfig
	flag.StringVar(&ac.out, "out", "./out/img.svg", "Path to SVG output")
	flag.StringVar(&ac.in, "in", "", "Path to access input image")
	flag.Var(&ac.selected, "filters", "Name of filters from the filters.json file to apply to the input file.")
	flag.Parse()
	ac.validateOut()
	ac.validateIn()

	return ac
}


func (ac *appConfig) validateIn() {
	// Check if it's a real file and we can open it.
	msgEmpty := "Please provide an input filepath:"
	// msgBad := fmt.Sprintf("The input filepath you provided,\n    %q\n , doesn't look valid. Please provide a different filepath:", ac.in)
	if ac.in == "" {
		ac.in = getUserInput(msgEmpty)
		// JAK implement retry logic
		ac.validateIn()
	}
	// JAK, Run an fs "valid" check
	// ac.in = getUserInput(msgBad)
}


func (ac *appConfig) validateOut() {
	msgEmpty := "Please provide an output filepath:"
	if ac.out == "" {
		// JAK implement retry logic
		ac.out = getUserInput(msgEmpty)
		ac.validateOut()
	}
	// JAK, if it doesn't end with .svg (or maybe .xml), reject it
	// JAK, add option to change to avoid overwrites
}
