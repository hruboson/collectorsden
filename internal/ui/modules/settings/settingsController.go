package moduleSettings

import (
	"fyne.io/fyne/v2"

	"hrubos.dev/collectorsden/internal/config"
	"hrubos.dev/collectorsden/internal/logger"
	themes "hrubos.dev/collectorsden/internal/ui/themes"
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

	c.View.SetThemeToggleChangeHandler(c.themeSwitcherLogic)
	c.View.SetExportButtonOnClick(c.Model.ExportDatabase)

	c.View.SetExportEntryPlaceHolder(c.Model.GetDefaultExportPath())

	return c
}

// TODO with more themes move this to settingsModal.go
func (c *Controller) themeSwitcherLogic(enabled bool) {
	logger.Log("Switching theme", logger.CatUI)
	if enabled {
		c.app.Settings().SetTheme(themes.NewDarkTheme())
		config.DarkThemeOn = true
	} else {
		c.app.Settings().SetTheme(themes.NewLightTheme())
		config.DarkThemeOn = false
	}
}
