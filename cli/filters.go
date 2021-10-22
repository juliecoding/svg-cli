package cli

import (
	"fmt"

	svg "github.com/ajstarks/svgo"
)

type filterFunc func(svg.SVG)

func initFilters() map[string]filterFunc {
	return map[string]filterFunc{
		"blur": blur,
		"bw": bw,
		"dawn":	dawn,
		// "duotone": duotone, // Try to get user input on this one
		"dusk": dusk,
		"hueRotate": hueRotate,
		// "matrix": matrix,
		// "moodring": moodring,
		"night": night,
		// "storm": storm,
	}
}

func blur(canvas svg.SVG) {

}


func bw(canvas svg.SVG) {
	fs := svg.Filterspec{ In: "SourceGraphic", Result: "bw" }
	canvas.FeColorMatrixSaturate(fs, 0.05)
}

func dawn(canvas svg.SVG) {
	result := "dawn"
	floodId := "__dawnFlood"
	floodSpec := svg.Filterspec{ Result: result }
	blendSpec := svg.Filterspec{ In2: result, In: "SourceGraphic"}
	canvas.FeFlood(floodSpec, "#EFB2D1", 0, fmt.Sprintf(`id="%s"`, floodId))
	canvas.FeBlend(blendSpec, "multiply")
	canvas.Animate(fmt.Sprintf("#%s", floodId), "flood-opacity", 0.1, 0.5, .5, 1)
	canvas.Animate(fmt.Sprintf("#%s", floodId), "flood-opacity", 0.5, 0, 3, 1)
}

func dusk(canvas svg.SVG) {
	result := "dusk"
	floodId := "__duskFlood"
	floodSpec := svg.Filterspec{ Result: result }
	blendSpec := svg.Filterspec{ In2: result, In: "SourceGraphic" }
	canvas.FeFlood(floodSpec, "#e3b249", 0.1, fmt.Sprintf(`id="%s"`, floodId))
	canvas.FeBlend(blendSpec, "multiply")
	canvas.Animate(fmt.Sprintf("#%s", floodId), "flood-opacity", 0, 0.3, 6, 1)
	canvas.Animate(fmt.Sprintf("#%s", floodId), "flood-opacity", 0.3, 0.1, 9, 1)
}

func hueRotate(canvas svg.SVG) {
	fs := svg.Filterspec{ In: "SourceGraphic", Result: "hueRotate" }
	canvas.FeColorMatrixHue(fs, 180)
}

func night(canvas svg.SVG) {
	result := "night"
	floodId := "__nightFlood"
	floodSpec := svg.Filterspec{ Result: result }
	blendSpec := svg.Filterspec{ In2: result, In: "SourceGraphic" }
	canvas.FeFlood(floodSpec, "#090f36", 0.95, fmt.Sprintf(`id="%s"`, floodId))
	canvas.FeBlend(blendSpec, "multiply")
	canvas.Animate(fmt.Sprintf("#%s", floodId), "flood-opacity", 0, 0.95, 9, 1)
}
