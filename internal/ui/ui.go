package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	config "hrubos.dev/collectorsden/internal/config"
	logger "hrubos.dev/collectorsden/internal/logger"

	bundled "hrubos.dev/collectorsden/internal/ui/bundled"
	themes "hrubos.dev/collectorsden/internal/ui/themes"

	//moduleFiles "hrubos.dev/collectorsden/internal/ui/modules/files"
	moduleHome "hrubos.dev/collectorsden/internal/ui/modules/home"
)

var INITIAL_WINDOW_WIDTH float32 = 1200
var INITIAL_WINDOW_HEIGHT float32 = 600

func Run(){
	logger.Log("Starting UI", logger.CatUI)

	app := app.NewWithID("hrubos.dev/collectorsden")
	config.AppSettings = app.Settings()
	config.AppSettings.SetTheme(themes.NewDarkTheme())
	app.SetIcon(bundled.ResourceAssetsImgIconPng)
	mainWindow := app.NewWindow("tree")

	homeModel := moduleHome.NewModel()
	homeView := moduleHome.NewView()
	homeController := moduleHome.NewController(homeModel, homeView, app, mainWindow)
	
	mainWindow.SetContent(homeController.View)
	mainWindow.Resize(fyne.NewSize(INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT))
	mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()

	logger.Log("Closing UI", logger.CatUI)
}
