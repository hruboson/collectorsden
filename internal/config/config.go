package config

// ----- Theme -----
var DarkThemeOn bool = true

// ----- Database -----
const (
    DBFile     		= "main.db"
	DBSecondaryFile = "secondary.db" // for imports
	AppFolder		= "collectorsden"
)

var ExportFile = "export.json"
var ExportPath string
