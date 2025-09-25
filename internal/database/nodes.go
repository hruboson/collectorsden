package database

import (
	"hrubos.dev/collectorsden/internal/indexer"
	logger "hrubos.dev/collectorsden/internal/logger"
)

//TODO
// maybe make different functions for folders and files (categories and content) and only call files in leaf nodes of tree

func StoreNode(node indexer.Node) error {
	var err error
	if node.Type() == indexer.FILE {
		var tryExisting Content
		texErr := db.One("Fullpath", node.GetPath(), &tryExisting)

		// node already in db
		if texErr == nil {
			return nil
		}

		entry := &Content{
			FullPath: node.GetPath(),
			Name: node.Name(),
		}
		err = db.Save(entry)
	} else if node.Type() == indexer.FOLDER {
		//var tryExisting Category
		//texErr := db.One("Fullpath", node.GetPath(), &tryExisting)

		entry := &Category{
			FullPath: node.GetPath(),
			Name: node.Name(),
		}
		err = db.Save(entry)
	}

	// check errors
	if err != nil {
		logger.Fatal("Error while storing node", err, logger.CatDB)
		return err
	}

	logger.Log("Successfully stored node " + node.Name(), logger.CatDB)

	return nil
}

func RemoveNode(node indexer.Node) error {
	var err error
	
	if node.Type() == indexer.FILE {
		var content Content
		err = db.One("FullPath", node.GetPath(), &content)

		// check errors
		if err != nil {
			logger.Fatal("Error while removing file", err, logger.CatDB)
			return err
		}

		err = db.DeleteStruct(&content)
	} else if node.Type() == indexer.FOLDER {
		var category Category
		err = db.One("FullPath", node.GetPath(), &category)

		// check errors
		if err != nil {
			logger.Fatal("Error while removing category", err, logger.CatDB)
			return err
		}

		err = db.DeleteStruct(&category)
	}

	// check delete errors
	if err != nil {
		logger.Fatal("Error while removing node", err, logger.CatDB)
		return err
	}

	logger.Log("Successfully removed node " + node.Name(), logger.CatDB)

	return nil
}

func AllNodes() []indexer.Node {
	nodes := make([]indexer.Node, 0)

	categories := make([]Category, 0)
	content := make([]Content, 0)
	err := db.All(&categories)
	if err != nil { panic(err) } //TODO
	err = db.All(&content)
	if err != nil { panic(err) } //TODO

	for _, cat := range categories {
		nodes = append(nodes, indexer.NewFolder(cat.FullPath, nil))
	}

	for _, con := range content {
		nodes = append(nodes, indexer.NewFile(con.FullPath, nil))
	}

	return nodes
}

func AllCategories() []Category {
	categories := make([]Category, 0)
	err := db.All(&categories)
	if err != nil { panic(err) } //TODO
	return categories
}
