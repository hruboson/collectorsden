package main

import (
	logger "hrubos.dev/collectorsden/internal/logger"
	ui "hrubos.dev/collectorsden/internal/ui"
	db "hrubos.dev/collectorsden/internal/database"
)

func main(){
	logger.Init("app.log")
	logger.Log("Starting app", logger.CatApp)

	db.Init()
	ui.Run()

	logger.Log("Closing app", logger.CatApp)
	logger.Close()
}
