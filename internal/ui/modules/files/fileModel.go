package moduleFiles

import (
	"strconv"

	indexer "hrubos.dev/collectorsden/internal/indexer"
	logger "hrubos.dev/collectorsden/internal/logger"
)

type Model struct {
	root string

	childrenCache map[string][]string // reason: keeps the tree queries inexpensive (drive cost)
}

func NewModel() *Model {
	m := &Model{
		root: "./",
        childrenCache: make(map[string][]string),
	}

	return m
}

// TreeData returns functions that can be bound to a widget.Tree
func (m *Model) TreeData() (
	childUIDs func(uid string) []string,
	isBranch func(uid string) bool,
	getName func(uid string) string,
) {
	childUIDs = func(uid string) []string {
		if uid == "" {
			uid = m.root
		}
		if cached, ok := m.childrenCache[uid]; ok {
			return cached
		}

		files := indexer.GetFiles(uid)
		m.childrenCache[uid] = files
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

func (m *Model) CheckNode(name string, checked bool){
	logger.Log(name + ": " + strconv.FormatBool(checked), logger.CatModel)
}

// ----- Data setters -----
func (m *Model) SetRoot(root string) {
	logger.Log("Tree root is now " + root, logger.CatModel)
	m.root = root
}

// ----- Data getters -----
func (m *Model) GetRoot() string {
	return m.root
}
