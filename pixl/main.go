package main

import (
	"image/color"
	"pixl/swatch"
	"pixl/types"
	"pixl/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")

	state := types.State{
		BrushColor:          color.NRGBA{255, 255, 255, 255},
		SelectedSwatchIndex: 0,
	}

	appInit := ui.AppInit{
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlWindow.ShowAndRun()
}
