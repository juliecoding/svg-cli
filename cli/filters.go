package cli

import (
	"fmt"

	"github.com/ajstarks/svgo"
)

func initFilters() map[string]filter {
	filters := make(map[string]filter)
	filters["dawn"]		= &Dawn{}
	filters["dusk"]		= &Dusk{}
	filters["night"]	= &Night{}

	return filters
}

type Filter struct {} // This should have the same structure/properties as the SVG element from lib

type filter interface { // JAK, look up correct naming convention
	Apply(canvas svg.SVG) svg.SVG
}

// These wouldn't need to be types if apply didn't work so differently for each one.
// Is a struct the appropriate data type?
type Dusk struct {}
type Dawn struct {}
type Night struct {}

func (d *Dusk) Apply(canvas svg.SVG) svg.SVG { // Could maybe make this void. Is that better?
	result := "floodOut"
	floodId := "__df"
	floodSpec := svg.Filterspec{ Result: result }
	blendSpec := svg.Filterspec{ In2: result, In: "SourceGraphic"}
	canvas.FeFlood(floodSpec, "#EFB2D1", 0, fmt.Sprintf(`id="%s"`, floodId))
	canvas.FeBlend(blendSpec, "multiply")
	canvas.Animate(fmt.Sprintf("#%s", floodId), "flood-opacity", 1, 0, 7, 1)
	return canvas
}

func (d *Dawn) Apply(canvas svg.SVG) svg.SVG {
	result1 := "duskFlood"
	floodId1 := "__df1"
	floodSpec1 := svg.Filterspec{ Result: result1 }
	blendSpec1 := svg.Filterspec{ In2: result1, In: "SourceGraphic" }
	canvas.FeFlood(floodSpec1, "#e3b249", 1, fmt.Sprintf(`id="%s"`, floodId1))
	canvas.FeBlend(blendSpec1, "multiply")
	canvas.Animate(fmt.Sprintf("#%s", floodId1), "flood-opacity", 0, 1, 7, 1)
	return canvas
}

func (n *Night) Apply(canvas svg.SVG) svg.SVG {
	result1 := "nightFlood"
	floodId1 := "__nf1"
	floodSpec1 := svg.Filterspec{ Result: result1 }
	blendSpec1 := svg.Filterspec{ In2: result1, In: "SourceGraphic" }
	canvas.FeFlood(floodSpec1, "#0c0b0f", 1, fmt.Sprintf(`id="%s"`, floodId1))
	canvas.FeBlend(blendSpec1, "multiply")
	canvas.Animate(fmt.Sprintf("#%s", floodId1), "flood-opacity", 0, 1, 7, 1)
	return canvas
}
