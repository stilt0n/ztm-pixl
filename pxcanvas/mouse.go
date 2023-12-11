package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (pxCanvas *PxCanvas) Scrolled(event *fyne.ScrollEvent) {
	pxCanvas.scale(int(event.Scrolled.DY))
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseMoved(event *desktop.MouseEvent) {
	// this will run when scroll wheel is pressed and mouse is moved
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, event)
	pxCanvas.Refresh()
	pxCanvas.mouseState.previousCoord = &event.PointEvent
}

// These two implementations are empty because we aren't using themn
// but they are required for the Hoverable interface
func (pxCanvas *PxCanvas) MouseIn(event *desktop.MouseEvent) {}
func (pxCanvas *PxCanvas) MouseOut()                         {}
