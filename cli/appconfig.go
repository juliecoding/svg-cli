package cli

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
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
	msgInvalid := fmt.Sprintf("The input filepath you provided, %q, doesn't look valid.\nPlease provide a different filepath:", ac.in)
	msgAbsolute := fmt.Sprintf("An error occurred converting the relative input path %q to an absolute path.\nConsider providing an absolute path to the input file:", ac.in)
	if ac.in == "" {
		ac.in = getUserInput(msgEmpty)
		// JAK implement retry logic
		ac.validateIn()
	}
	if !fs.ValidPath(ac.in) {
		ac.in = getUserInput(msgInvalid)
	}
	abs, err := filepath.Abs(ac.in)
	if (err != nil) {
		ac.in = getUserInput(msgAbsolute)
	} else {
		ac.in = abs
		fmt.Println("HEEEY" + ac.in)
	}
}


func (ac *appConfig) validateOut() {
	msgEmpty := "Please provide an output filepath:"
	if ac.out == "" {
		// JAK implement retry logic
		ac.out = getUserInput(msgEmpty)
		ac.validateOut()
	}
	extension := filepath.Ext(ac.out)
	if extension != "svg" && extension != "xml" {
		outputError("Changing output file extension to '.svg'", nil)
		ac.out = ac.out + ".svg"
	}
	// JAK, add option to change to avoid overwrites
}
