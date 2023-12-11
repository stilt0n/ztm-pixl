package ui

import (
	"pixl/pxcanvas"
	"pixl/swatch"
	"pixl/types"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *types.State
	Swatches   []*swatch.Swatch
}
