package logger

import (
	"io"
	"log"
	"os"

	"hrubos.dev/collectorsden/internal/config"
)

var file *os.File

// Log message categories
const (
	CatUI = "UI"
	CatApp = "APP"
	CatIndexer = "INDXR"
	CatDB = "DB"
	CatController = "CONT"
	CatModel = "MODL"
	CatView = "VIEW"
	CatOther = ""
)

func Init(logFilePath string) error {
	// Open or create log file
	var err error
	file, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	// Write logs to both stdout and file
	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)

	return nil
}

func Close() {
	if file != nil {
		file.Close()
	}
}

func Log(msg string, category string){
	if !config.DebugBuild {
		return
	}

	if category != "" {
		log.Println("[" + category + "]" + " " + msg)
	}else{
		log.Println(msg)
	}
}

func Fatal(msg string, err error, category string){
	if err == nil {
		return
	}

	log.Println("[FATAL][" + category + "] " + msg + "\n" + err.Error())

	// Flush & close log file
	Close()

	// Exit program
	os.Exit(1)
}
