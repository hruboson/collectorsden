package config

import "fyne.io/fyne/v2"

var DebugBuildStr = "false" // should be changed when running go build... see Makefile
var DebugBuild = false // set in main

// ----- Theme -----
var DarkThemeOn bool = true
var AppSettings fyne.Settings

// ----- Database -----
const (
    DBFile     		= "main.db"
	DBSecondaryFile = "secondary.db" // for imports
	AppFolder		= "collectorsden"
)

var ExportFile = "export.json"
var ExportPath string
