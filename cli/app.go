package cli

import (
	"fmt"
	"image"
	"os"

	svg "github.com/ajstarks/svgo"
	browser "github.com/pkg/browser"
)


type app struct {
	filters map[string]filterFunc
	config appConfig
	canvas *svg.SVG
}

type dimensions struct {
	width int
	height int
}

func (a *app) op() {
	d := a.getDimensions()
	w := a.getWriter()
	a.canvas = svg.New(w)
	a.canvas.Start(d.width, d.height)
	a.canvas.Def()
	a.applyFilters()
	a.canvas.DefEnd()

	if len(a.config.selected) > 0 {
		a.canvas.Image(0, 0, d.width, d.height, a.config.in, `filter="url(#__filters)"`)
	} else {
		a.canvas.Image(0, 0, d.width, d.height, a.config.in)
	}
	a.canvas.End()
	browser.OpenFile(a.config.out)
}


func (a *app) applyFilters() {
	if len(a.config.selected) == 0 {
		return
	}
	a.canvas.Filter("__filters")
	// Loop through selected filters and apply them
	for _, f := range a.config.selected {
		// Check if filter map has key
		_, ok := a.filters[f]
		if ok {
			a.filters[f](*a.canvas)
		} else {
			output(fmt.Sprintf("Filter '%s' does not exist; skipping...", f), nil)
		}
	}
	// SourceGraphic should be the "in" value of the 1st node
	a.canvas.FeMerge(append([]string{"SourceGraphic"}, a.config.selected...))
	a.canvas.Fend()
}


func (a *app) getWriter() *os.File {
	w, err := os.Create(a.config.out)
	if err != nil {
		output(fmt.Sprintf("There was an issue creating the file %s", a.config.out), err)
		a.config.out = getUserInput("Please enter a valid filepath for your output (without quotes): ")
	}
	return w
}

func (a *app) getDimensions() dimensions {
	var d dimensions
	f, errO := os.Open(a.config.in)
	if errO != nil {
		output(fmt.Sprintf("There was an issue opening the input file at %s", a.config.in), errO)
		a.config.in = getUserInput("Please enter a valid input filepath (without quotes): ")
	}
	img, _, errD := image.DecodeConfig(f)
	if errD != nil {
		output(fmt.Sprintf("There was an issue decoding the input file at %s", a.config.in), errD)
		a.config.in = getUserInput("Please enter a valid input filepath (without quotes): ")
	}
	d.width = img.Width
	d.height = img.Height
	f.Close()
	return d
}
