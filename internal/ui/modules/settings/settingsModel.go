package moduleSettings

import (
	db "hrubos.dev/collectorsden/internal/database"
	config "hrubos.dev/collectorsden/internal/config"
)

type Model struct {
}

func NewModel() *Model {
	fm := &Model{
	}

	return fm
}

func (m *Model) ExportDatabase() {
	err := db.Export()
	if err != nil {
		panic(err)
	}
}

// ----- Data setters -----


// ----- Data getters -----

func (m *Model) GetDefaultExportPath() string {
	return config.ExportPath + config.ExportFile
}
