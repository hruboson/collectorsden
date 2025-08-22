package main

import (
	indexerPackage "hrubos.dev/collectorsden/internal/indexer"
	"hrubos.dev/collectorsden/internal/logger"
	"hrubos.dev/collectorsden/internal/ui"
)

func main(){
	logger.Init("app.log")
	logger.Log("Starting app", logger.CatApp)

	indexer := indexerPackage.NewIndexer()


	err := indexer.TreeDir("G:\\Projects")
	if err != nil {
		logger.Fatal(err)
	}

	indexer.FileStructure.Print("")
	ui.Run()

	logger.Log("Closing app", logger.CatApp)
	logger.Close()
}
