package moduleTemplate

import (
	"fyne.io/fyne/v2"
)

type Controller struct {
	*Model
	*View
	window fyne.Window
	app fyne.App
}

func NewController(fm *Model, fv *View, app fyne.App, window fyne.Window) *Controller {
	fc := &Controller{
		Model: fm,
		View: fv,
		window: window,
		app: app,
	}

	return fc
}
