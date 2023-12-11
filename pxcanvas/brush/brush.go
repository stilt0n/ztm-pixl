package brush

import (
	"pixl/types"

	"fyne.io/fyne/v2/driver/desktop"
)

const (
	Pixel = iota
)

// This will allow for implementing more brush types in the future
func TryBrush(appState *types.State, canvas types.Brushable, event *desktop.MouseEvent) bool {
	switch {
	case appState.BrushType == Pixel:
		return TryPaintPixel(appState, canvas, event)
	default:
		return false
	}
}

func TryPaintPixel(appState *types.State, canvas types.Brushable, event *desktop.MouseEvent) bool {
	x, y := canvas.MouseToCanvasXY(event)
	if x != nil && y != nil && event.Button == desktop.MouseButtonPrimary {
		canvas.SetColor(appState.BrushColor, *x, *y)
		return true
	}
	return false
}
