package cli

import (
	"fmt"
	"image"
	"os"

	// Decoding any particular image format requires the prior registration of a decoder function.
	// Registration is typically automatic as a side effect of initializing that format's package,
	// so here we're using _ to import the package purely for its initialization side effects.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	svg "github.com/ajstarks/svgo"
)


type app struct {
	filters	map[string]filter
	config appConfig
	canvas *svg.SVG // JAK, is syntax correct??
}


type dimensions struct {
	width int
	height int
}


func (a *app) op() {
	a.filters = initFilters()
	a.config = getConfig()
	d := a.getDimensions()
	w := a.getWriter()
	a.canvas = svg.New(w)
	a.canvas.Start(d.width, d.height)
	a.canvas.Def()
	a.applyFilters()
	a.canvas.DefEnd()

	if len(a.config.selected) > 0 {
		a.canvas.Image(0, 0, d.width, d.height, a.config.in, `filter="url(#__filters)`)
	} else {
		a.canvas.Image(0, 0, d.width, d.height, a.config.in)
	}
	a.canvas.End()
	// Could have the program automatically open the file in a browser
}


func (a *app) applyFilters() {
	if len(a.config.selected) == 0 {
		return
	}
	a.canvas.Filter("__filters")
	// Loop through selected filters and apply them
	for _, s := range a.config.selected {
		// Check if object has key
		filt, ok := a.filters[s]
		if ok {
			filt.Apply(*a.canvas)
		} else {
			outputError(fmt.Sprintf("Filter '%s' does not exist; skipping...", s), nil) // JAK may be worth creating output func that outputError calls or vice versa
		}
	}
	a.canvas.Fend()
}


func (a *app) getWriter() *os.File {
	w, err := os.Create(a.config.out)
	if err != nil {
		// JAK, add retry logic
		outputError(fmt.Sprintf("There was an issue creating the file %s", a.config.out), err)
		a.config.out = getUserInput("Please enter a valid filepath for your output (without quotes): ")
	}
	return w
}


// Would like to dissociate this from the app.
func (a *app) getDimensions() dimensions {
	var d dimensions
	f, errO := os.Open(a.config.in)
	if errO != nil {
		// JAK, add retry logic
		outputError(fmt.Sprintf("There was an issue opening the input file at %s", a.config.in), errO)
		a.config.in = getUserInput("Please enter a valid input filepath (without quotes): ")
	}
	img, _, errD := image.DecodeConfig(f)
	if errD != nil {
		// JAK, add retry logic
		outputError(fmt.Sprintf("There was an issue decoding the input file at %s", a.config.in), errD)
		a.config.in = getUserInput("Please enter a valid input filepath (without quotes): ")
	}
	d.width = img.Width
	d.height = img.Height
	f.Close()
	return d
}
