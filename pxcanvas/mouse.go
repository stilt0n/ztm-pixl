package pxcanvas

import (
	"pixl/pxcanvas/brush"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (pxCanvas *PxCanvas) Scrolled(event *fyne.ScrollEvent) {
	pxCanvas.scale(int(event.Scrolled.DY))
	pxCanvas.Refresh()
}

// This dispatches events when appropriate, whether the events actually
// result in a change to the canvas is determined in the hanlder functions
func (pxCanvas *PxCanvas) MouseMoved(event *desktop.MouseEvent) {
	if x, y := pxCanvas.MouseToCanvasXY(event); x != nil && y != nil {
		brush.TryBrush(pxCanvas.appState, pxCanvas, event)
		cursor := brush.Cursor(pxCanvas.PxCanvasConfig, pxCanvas.appState.BrushType, event, *x, *y)
		pxCanvas.renderer.SetCursor(cursor)
	} else {
		pxCanvas.renderer.SetCursor(make([]fyne.CanvasObject, 0))
	}
	// this will run when scroll wheel is pressed and mouse is moved
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, event)
	pxCanvas.Refresh()
	pxCanvas.mouseState.previousCoord = &event.PointEvent
}

// These two implementations are empty because we aren't using themn
// but they are required for the Hoverable interface
func (pxCanvas *PxCanvas) MouseIn(event *desktop.MouseEvent) {}
func (pxCanvas *PxCanvas) MouseOut()                         {}

func (pxCanvas *PxCanvas) MouseDown(event *desktop.MouseEvent) {
	brush.TryBrush(pxCanvas.appState, pxCanvas, event)
}

func (pxCanvas *PxCanvas) MouseUp(event *desktop.MouseEvent) {}
