package ui

import (
	"fyne.io/fyne/v2/container"
)

func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)
	// Border has main central container
	// and then left, top, bottom and right sections
	// order is top, bottom, left, right, ...center-content
	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker, app.PixlCanvas)
	app.PixlWindow.SetContent(appLayout)
}
