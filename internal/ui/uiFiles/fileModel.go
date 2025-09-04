package uiFiles

import (
	indexer "hrubos.dev/collectorsden/internal/indexer"
	logger "hrubos.dev/collectorsden/internal/logger"
)

type FileModel struct {
	root string

	childrenCache map[string][]string // reason: keeps the tree queries inexpensive (drive cost)
}

func NewFileModel() *FileModel {
	fm := &FileModel{
		root: "./",
        childrenCache: make(map[string][]string),
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
			uid = fm.root
		}
		if cached, ok := fm.childrenCache[uid]; ok {
			return cached
		}

		files := indexer.GetFiles(uid)
		fm.childrenCache[uid] = files
		return files
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
