package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	logger "hrubos.dev/collectorsden/internal/logger"

	controllers "hrubos.dev/collectorsden/internal/ui/controllers"
	models "hrubos.dev/collectorsden/internal/ui/models"
	views "hrubos.dev/collectorsden/internal/ui/views"
)

var INITIAL_WINDOW_WIDTH float32 = 1200
var INITIAL_WINDOW_HEIGHT float32 = 600

func Run(){
	logger.Log("Starting UI", logger.CatUI)

	app := app.NewWithID("hrubos.dev/collectorsden")
	app.Settings().SetTheme(&darkTheme{})
	mainWindow := app.NewWindow("tree")

	fileModel := models.NewFileModel()
	fileView := views.NewFileView()
	fileController := controllers.NewFileController(fileModel, fileView, mainWindow)

	mainWindow.SetContent(fileController.FileView)
	mainWindow.Resize(fyne.NewSize(INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT))
	mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()

	logger.Log("Closing UI", logger.CatUI)
}
