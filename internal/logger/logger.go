package logger

import (
	"io"
	"log"
	"os"
	"strconv"
)

var debugBuildValue string = "false" // should be changed when running go build... see Makefile
var debugEnabled bool
var file *os.File

// Log message categories
const (
	CatUI = "UI"
	CatApp = "APP"
	CatIndexer = "INDXR"
	CatOther = ""
)

func Init(logFilePath string) error {
	debugEnabled, _ = strconv.ParseBool(debugBuildValue)

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
	if !debugEnabled {
		return
	}

	if category != "" {
		log.Println("[" + category + "]" + " " + msg)
	}else{
		log.Println(msg)
	}
}

func Fatal(err error, category ...string){
	if err == nil {
		return
	}

	// Determine category
	cat := ""
	if len(category) > 0 {
		cat = category[0]
	}

	// Log the error
	if cat != "" {
		log.Println("[FATAL][" + cat + "] " + err.Error())
	} else {
		log.Println("[FATAL] " + err.Error())
	}

	// Flush & close file
	Close()

	// Exit program
	os.Exit(1)
}
