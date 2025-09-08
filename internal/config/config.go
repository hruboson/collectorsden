package config

import "fyne.io/fyne/v2"

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
