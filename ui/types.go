package ui

import (
	"pixl/swatch"
	"pixl/types"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *types.State
	Swatches   []*swatch.Swatch
}
