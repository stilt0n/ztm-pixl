package util

import (
	"image"
	"image/color"
)

// Go doesn't have support for sets so we're using a map to
// empty values
func GetImageColors(img image.Image) map[color.Color]struct{} {
	colors := make(map[color.Color]struct{})
	var empty struct{}
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			colors[img.At(x, y)] = empty
		}
	}
	return colors
}
