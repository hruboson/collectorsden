package moduleHome

import (
	"fyne.io/fyne/v2"

	"hrubos.dev/collectorsden/internal/ui/modules/store"
)

type Controller struct {
	*Model
	*View
	window fyne.Window
	app fyne.App
}

func NewController(m *Model, v *View, app fyne.App, window fyne.Window) *Controller {
	c := &Controller{
		Model: m,
		View: v,
		window: window,
		app: app,
	}

	c.View.SetOpenButtonOnTapped(c.tbd)

	return c
}

func (c  *Controller) tbd(){
	// for now just show file module
	// TODO also check if this frees the homeView,Model,Controller ... (hopefully it does)
	storeModel := moduleStore.NewModel()
	storeView := moduleStore.NewView()
	storeController := moduleStore.NewController(storeModel, storeView, c.app, c.window)

	c.window.SetContent(storeController.View)
}
