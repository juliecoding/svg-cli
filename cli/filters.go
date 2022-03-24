package cli

import (
	"fmt"

	svg "github.com/ajstarks/svgo"
)

type filt struct {
	description	string
	apply		applyFunc
}

type applyFunc func(svg.SVG)

type filterFunc func(svg.SVG)

func initFilters() map[string]filterFunc {
	return map[string]filterFunc{
		"blur": blur,
		"bw": bw,
		"carlton": carlton,
		"dawn":	dawn,
		"desaturate": desaturate,
		"day": day,
		"fuzzyTv": fuzzyTv,
		"ginza": ginza,
		"hueRotate": hueRotate,
		"instagram": instagram,
		"matrix": matrix,
		"montyPython": montyPython,
		"dusk": dusk,
		"pointLight": pointLight,
		"saturate": saturate,
		"sepia": sepia,
		"sunshine": sunshine,
	}
}

func initFilters2() map[string]filt {
	return map[string]filt{
		"blur": {
			"Makes image appear less sharp and defined",
			blur,
		},
		"bw": {
			"Changes image to black and white",
			bw,
		},
		"carlton": {
			"Did someone say happy dance?",
			carlton,
		},
		"dawn":	{
			"Animated. Gradually lightens and adds warm tones to image",
			dawn,
		},
		"desaturate": {
			"Mutes colors in an image",
			desaturate,
		},
		"day": {
			"Animated. Gradually brightens an image",
			day,
		},
		"fuzzyTv": {
			"Sharpens but distorts image",
			fuzzyTv,
		},
		"ginza": {
			"Based on the Instagram filter. Adds a warm glow to photos",
			ginza,
		},
		"hueRotate": {
			"Changes colors to their complement (e.g., red to green)",
			hueRotate,
		},
		"instagram": {
			"Makes viewers envious",
			instagram,
		},
		"matrix": {
			"Adds a cascading gif",
			matrix,
		},
		"montyPython": {
			"You may or may not expect it",
			montyPython,
		},
		"dusk": {
			"Animated. Gradually darkens and adds cool tones to image",
			dusk,
		},
		"pointLight": {
			"Adds a spotlight to upper corner of the image, a little like a lens flare",
			pointLight,
		},
		"saturate": {
			"Heightens colors",
			saturate,
		},
		"sepia": {
			"Makes image brown-toned",
			sepia,
		},
		"sunshine": {
			"Katrina's walking on it",
			sunshine,
		},
	}
}

func blur(canvas svg.SVG) {
	fs := svg.Filterspec{ In: "SourceGraphic", Result: "blur" }
	canvas.FeGaussianBlur(fs, 3, 3)
}

func bw(canvas svg.SVG) {
	fs := svg.Filterspec{ In: "SourceGraphic", Result: "bw" }
	canvas.FeColorMatrixSaturate(fs, 0.05)
}

func carlton(canvas svg.SVG) {
	canvas.FeImage("../assets/carlton.gif", "carlton")
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

func desaturate(canvas svg.SVG) {
	fs := svg.Filterspec{ In: "SourceGraphic", Result: "desaturate" }
	canvas.FeColorMatrixSaturate(fs, 0.3)
}

func day(canvas svg.SVG) {
	result := "day"
	floodId := "__dayFlood"
	floodSpec := svg.Filterspec{ Result: result }
	blendSpec := svg.Filterspec{ In2: result, In: "SourceGraphic" }
	canvas.FeFlood(floodSpec, "#e3b249", 0.1, fmt.Sprintf(`id="%s"`, floodId))
	canvas.FeBlend(blendSpec, "multiply")
	canvas.Animate(fmt.Sprintf("#%s", floodId), "flood-opacity", 0, 0.3, 6, 1)
	canvas.Animate(fmt.Sprintf("#%s", floodId), "flood-opacity", 0.3, 0.1, 9, 1)
}

func fuzzyTv(canvas svg.SVG) {
	fs := svg.Filterspec{ In: "SourceGraphic", Result: "fuzzyTv" }
	canvas.FeConvolveMatrix(fs, [9]int{-1, -1, -1, -1, 4, 0, 0, 0, -1})
}

func ginza(canvas svg.SVG) {
	fscm := svg.Filterspec{ In: "SourceGraphic", Result: "ginza" }

	matrix := [20]float64{
		0.8786, 0.1538, 0.4, 0, 0,
		0.0698, 0.9372, 0.0336, 0, 0,
		0.0544, 0.1068, 0.8262, 0, 0,
		0, 0, 0, 1, 0,
	}
	canvas.FeColorMatrix(fscm, matrix)

}

func hueRotate(canvas svg.SVG) {
	fs := svg.Filterspec{ In: "SourceGraphic", Result: "hueRotate" }
	canvas.FeColorMatrixHue(fs, 180)
}

func instagram(canvas svg.SVG) {
	fs := svg.Filterspec{ In: "SourceGraphic", Result: "instagram" }
	canvas.FeColorMatrixHue(fs, 20)
}

func matrix(canvas svg.SVG) {
	canvas.FeImage("../assets/matrixCascade.png", "matrix")
}

func montyPython(canvas svg.SVG) {
	cm := map[string]string{
		"red": "#ed190e",
		"orange": "#ed890e",
		"yellow": "#faf564",
		"green": "#20bf1b",
		"blue": "#1b70bf",
		"purple": "#781bbf",
		"pink": "#eb28a6",
	}
	fs := svg.Filterspec{ Result: "montyPython" }
	color := getUserInput("What is your favorite color?!")
	canvas.FeFlood(fs, cm[color], 0.5)
}

func dusk(canvas svg.SVG) {
	result := "dusk"
	floodId := "__duskFlood"
	floodSpec := svg.Filterspec{ Result: result }
	blendSpec := svg.Filterspec{ In2: result, In: "SourceGraphic" }
	canvas.FeFlood(floodSpec, "#090f36", 0.95, fmt.Sprintf(`id="%s"`, floodId))
	canvas.FeBlend(blendSpec, "multiply")
	canvas.Animate(fmt.Sprintf("#%s", floodId), "flood-opacity", 0, 0.95, 9, 1)
}

func pointLight(canvas svg.SVG) {
	fs := svg.Filterspec{ In: "SourceGraphic", Result: "pointLight" }
	canvas.FeSpecularLighting(fs, 0.5, 1.5, 80, "#F3F3F0")
	canvas.FePointLight(100, 100, 500)
	canvas.FeSpecEnd()
}

func saturate(canvas svg.SVG) {
	fs := svg.Filterspec{ In: "SourceGraphic", Result: "saturate" }
	canvas.FeColorMatrixSaturate(fs, 3)
}

func sepia(canvas svg.SVG) {
	fs := svg.Filterspec{ In: "SourceGraphic", Result: "sepia" }
	matrix := [20]float64{
		0.280, 0.5, 0.25, 0, 0,
		0.140, 0.4, 0.09, 0, 0,
		0.080, 0.3, 0.03, 0, 0,
		0, 0, 0, 1, 0,
	}
	canvas.FeColorMatrix(fs, matrix)
}

func sunshine(canvas svg.SVG) {
	fs := svg.Filterspec{ In: "SourceGraphic", Result: "moonbeam" }
	canvas.FeSpecularLighting(fs, 0.5, 1.5, 80, "#fff9c4")
	canvas.FePointLight(50, 50, 800)
	canvas.FePointLight(-50, -50, 800)
	canvas.FeSpecEnd()
}

// pointLight saturate sepia blur dawn desaturate dusk fuzzyTv ginza hueRotate matrix night carlton
