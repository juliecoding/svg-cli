package cli

import (
	svg "github.com/ajstarks/svgo"
	// uuid "github.com/google/uuid"
)

type dimensions struct {
	width  int
	height int
}

type xcanvas struct {
	canvas	*svg.SVG
	d		dimensions
}

// JAK, try to streamline what you're doing with the img path so you don't have to pass it around so much.
// The imperative nature of the lib is a challenge.
// Do more error handling.
func (x *xcanvas) createCanvas(selectedFilters []string, imgPath string) error {
	x.canvas.Start(x.d.width, x.d.height)
	if len(selectedFilters) > 0 {
		x.applyFilters(selectedFilters, imgPath)
	} else {
		x.canvas.Image(0, 0, x.d.width, x.d.height, imgPath)
	}
	x.canvas.End()
	return nil
}

func (x *xcanvas) applyFilters(selectedFilters []string, imgPath string) {
	x.canvas.Def()
	ids := x.addFilters()
	x.canvas.DefEnd()
	x.addGroups(ids, imgPath)
}

func (x *xcanvas) addFilters() []string {
	var ids []string
	return ids
}

func (x *xcanvas) addGroups(ids []string, imgPath string) {
	// Add the opening group tags
	x.canvas.Image(0, 0, x.d.width, x.d.height, imgPath)
	// Add the closing group tags
}