package uiFiles

import (
	indexer "hrubos.dev/collectorsden/internal/indexer"
	logger "hrubos.dev/collectorsden/internal/logger"
)

type FileModel struct {
	root string
}

func NewFileModel() *FileModel {
	fm := &FileModel{
		root: "./",
	}

	return fm
}

// TreeData returns functions that can be bound to a widget.Tree
func (fm *FileModel) TreeData() (
	childUIDs func(uid string) []string,
	isBranch func(uid string) bool,
	getName func(uid string) string,
) {
	childUIDs = func(uid string) []string {
		if uid == "" {
			return indexer.GetFiles(fm.root)
		}
		return indexer.GetFiles(uid)
	}

	isBranch = func(uid string) bool {
		return indexer.IsDir(uid)
	}

	getName = func(uid string) string {
		return indexer.GetFileName(uid)
	}

	return
}

// ----- Data setters -----
func (fm *FileModel) SetRoot(root string) {
	logger.Log("Tree root is now " + root, logger.CatModel)
	fm.root = root
}

// ----- Data getters -----
func (fm *FileModel) GetRoot() string {
	return fm.root
}
