package main

import "hrubos.dev/collectorsden/internal/ui"
import "hrubos.dev/collectorsden/internal/indexer"
import "hrubos.dev/collectorsden/internal/logger"

func main(){
	logger.Init("app.log")
	logger.Log("Starting app", logger.CatApp)


	indexer.TreeDir("E:\\Archive", 3)
	if 0 == 1{
		ui.Run()
	}
	

	logger.Log("Closing app", logger.CatApp)
	logger.Close()
}
