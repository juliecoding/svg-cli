package cli

import (
	svg "github.com/ajstarks/svgo"
	browser "github.com/pkg/browser"
)

type app struct {
	filters map[string]filterFunc
	config  appConfig
}

func (a *app) op() int {
	d, err := getDimensions(a.config.in)
	if err != nil {
		a.config.in = getValidInput("")
		d, err = getDimensions(a.config.in)
		if err != nil {
			output("Exiting due to issues with input file", err)
			return 1
		}
	}

	w, err := initWriter(a.config.out)
	if err != nil {
		a.config.out = getValidOutput("")
		w, err = initWriter(a.config.out)
		if err != nil {
			output("Exiting due to issues with output file", err)
			return 1
		}
	}

	var x = xcanvas{ svg.New(w), d }
	err = x.drawCanvas(a.config.selected, a.config.in)
	if err != nil {
		output("Bummer! Exiting due to an error encountered while creating your SVG", err)
		return 1
	}
	browser.OpenFile(a.config.out)

	// Success int
	return 0
}
