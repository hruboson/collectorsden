package moduleFiles

import (
	"hrubos.dev/collectorsden/internal/database"
	indexer "hrubos.dev/collectorsden/internal/indexer"
	logger "hrubos.dev/collectorsden/internal/logger"
)

type Model struct {
	root string

	nodeCache map[string]indexer.Node // reason: keeps the tree queries inexpensive (drive cost)
}

func NewModel() *Model {
	m := &Model{
		root: "./",
        nodeCache: make(map[string]indexer.Node),
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

		struids := []string{}

		node, ok := m.nodeCache[uid]
		folder, ok := node.(*indexer.Folder)
		if ok {
			// iterate over children
			for _, child := range folder.GetChildren() {
				struids = append(struids, child.GetPath())
			}
		}

		return struids
	}

	isBranch = func(uid string) bool {
		if uid == "" {
			return false
		}

		node, ok := m.nodeCache[uid]
		if ok {
			//logger.Log("Cache hit for " + uid, logger.CatModel)
			return node.Type() == indexer.FOLDER
		}

		if indexer.IsDir(uid){
			//logger.Log("Cache miss for " + uid, logger.CatModel)
			
			// add uid to cache and its children
			folder := indexer.NewFolder(uid, nil)
			m.nodeCache[uid] = folder
			for _, child := range folder.GetChildren() {
				m.nodeCache[child.GetPath()] = child
			}

			return true
		} else {
			file := indexer.NewFile(uid, nil)
			m.nodeCache[uid] = file

			return false
		}
	}

	getName = func(uid string) string {
		return uid
	}

	return
}

func (m *Model) CheckNode(uid string, checked bool){
	node := m.nodeCache[uid]
	if(checked){
		database.StoreNode(node)
	}else{
		database.RemoveNode(node)
	}
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

func (m *Model) GetNodeFromUID(uid string) indexer.Node {
	return m.nodeCache[uid]
}
