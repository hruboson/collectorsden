package main

import (
	"os"
	"path/filepath"

	logger "hrubos.dev/collectorsden/internal/logger"
	ui "hrubos.dev/collectorsden/internal/ui"
	db "hrubos.dev/collectorsden/internal/database"
	config "hrubos.dev/collectorsden/internal/config"
)

func main(){
	logger.Init("app.log")
	logger.Log("Starting app", logger.CatApp)

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	config.ExportPath = filepath.Join(home, config.AppFolder)

	db.Init()
	ui.Run()
	db.Close()

	logger.Log("Closing app", logger.CatApp)
	logger.Close()
}
