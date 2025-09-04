package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	logger "hrubos.dev/collectorsden/internal/logger"

	themes "hrubos.dev/collectorsden/internal/ui/themes"
	bundled "hrubos.dev/collectorsden/internal/ui/bundled"
	uiFiles "hrubos.dev/collectorsden/internal/ui/uiFiles"
)

var INITIAL_WINDOW_WIDTH float32 = 1200
var INITIAL_WINDOW_HEIGHT float32 = 600

func Run(){
	logger.Log("Starting UI", logger.CatUI)

	app := app.NewWithID("hrubos.dev/collectorsden")
	app.Settings().SetTheme(themes.NewDarkTheme())
	app.SetIcon(bundled.ResourceAssetsImgIconPng)
	mainWindow := app.NewWindow("tree")

	fileModel := uiFiles.NewModel()
	fileView := uiFiles.NewView()
	fileController := uiFiles.NewController(fileModel, fileView, app, mainWindow)

	mainWindow.SetContent(fileController.View)
	mainWindow.Resize(fyne.NewSize(INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT))
	mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()

	logger.Log("Closing UI", logger.CatUI)
}
