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
	// JAK, use more sophisticated retry's to allow more than 2 attempts
	// Try thrice, then print error and return 1
	d, err := getDimensions(a.config.in)
	if (err != nil) {
		a.config.in = getValidInput("")
		d, err = getDimensions(a.config.in)
	}

	w, err := getWriter(a.config.out)
	if (err != nil) {
		a.config.out = getValidOutput("")
		w, err = getWriter(a.config.out)
	}

	var x = xcanvas{ svg.New(w), d }
	err = x.createCanvas(a.config.selected, a.config.in)
	if err != nil {
		output("Bummer! I'm quitting", err)
		return 1
	}
	browser.OpenFile(a.config.out)

	// Success int
	return 0
}
