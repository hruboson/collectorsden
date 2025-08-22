package main

import (
	logger "hrubos.dev/collectorsden/internal/logger"
	ui "hrubos.dev/collectorsden/internal/ui"
)

func main(){
	logger.Init("app.log")
	logger.Log("Starting app", logger.CatApp)

	ui.Run()

	logger.Log("Closing app", logger.CatApp)
	logger.Close()
}
