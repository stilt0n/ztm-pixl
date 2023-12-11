// Mouseable interface gives mouse down and mouse up events
// Scrollable gives Scrolled(*ScrollEvent)
// Hoverable gives mouse in, mouse moved and mouse out
package swatch

import "fyne.io/fyne/v2/driver/desktop"

func (swatch *Swatch) MouseDown(event *desktop.MouseEvent) {
	swatch.clickHandler(swatch)
	swatch.Selected = true
	swatch.Refresh()
}

// Required for interface but noop
func (swatch *Swatch) MouseUp(event *desktop.MouseEvent) {}
