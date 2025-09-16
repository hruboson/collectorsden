package database

import (
	"encoding/json"
	"fmt"
	"os"
    "path/filepath"

	"github.com/asdine/storm/v3"

	config "hrubos.dev/collectorsden/internal/config"
	logger "hrubos.dev/collectorsden/internal/logger"
)

var db *storm.DB

func Init() {
	var err error
	db, err = storm.Open(config.DBFile)
	if err != nil {
		panic(err)
	}
}

func Close() {
	if db != nil {
		db.Close()
	}
	logger.Log("Successfully close DB", logger.CatDB)
}

func Export() error {
	absPath, err := filepath.Abs(filepath.Join(config.ExportPath, config.ExportFile))
	if err != nil {
		return err
	}

	logger.Log("Exporting database to " + absPath + "...", logger.CatDB)

	// Export all tables
	var categories []Category

	if err := db.All(&categories); err != nil {
		return err
	}

	// Wrap export in a top-level object
	export := map[string]any{
		"categories": categories,
	}

	// Marshal to human-readable JSON
	data, err := json.MarshalIndent(export, "", "  ")
	if err != nil {
		return err
	}

	// Create folder if it doesn't exist
	if err := os.MkdirAll(config.ExportPath, 0755); err != nil {
		return err
	}

	// Save to file
	if err := os.WriteFile(absPath, data, 0644); err != nil {
		return err
	}

	logger.Log("Export to " + absPath + " completed", logger.CatDB)
	return nil
}

func Import() error {
	// Read JSON file
	data, err := os.ReadFile(config.ExportFile)
	if err != nil {
		return err
	}

	// Decode into struct
	var importData struct {
		Categories []Category `json:"categories"`
	}
	if err := json.Unmarshal(data, &importData); err != nil {
		return err
	}

	// Insert tables into db
	for _, c := range importData.Categories {
		if err := db.Save(&c); err != nil {
			return err
		}
	}

	// Log
	var categories []Category
	if err := db.All(&categories); err != nil {
		return err
	}

	logger.Log("Imported categories: ", logger.CatDB)
	for _, c := range categories {
		logger.Log(fmt.Sprintf("  %+v\n", c), logger.CatDB)
	}

	return nil
}
