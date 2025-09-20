package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"strconv"

	config "hrubos.dev/collectorsden/internal/config"
	db "hrubos.dev/collectorsden/internal/database"
	logger "hrubos.dev/collectorsden/internal/logger"
	ui "hrubos.dev/collectorsden/internal/ui"
)

func main(){
	config.DebugBuild, _ = strconv.ParseBool(config.DebugBuildStr)

	if config.DebugBuild {
		go func() {
			http.ListenAndServe("localhost:6420", nil)
		}()
	}

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
