package database

import (
	"hrubos.dev/collectorsden/internal/indexer"
	logger "hrubos.dev/collectorsden/internal/logger"
)

func StoreNode(node indexer.Node) error {
	var err error
	if node.Type() == indexer.FILE {
		entry := &Content{}
		err = db.Save(entry)
	} else if node.Type() == indexer.FOLDER {
		entry := &Category{
			Folder: node.GetPath(),
			Name: node.Name(),
		}
		err = db.Save(entry)
	}

	// check errors
	if err != nil {
		logger.Fatal(err, logger.CatDB)
		return err
	}

	return nil
}

func RemoveNode(node indexer.Node) error {
	return nil
}
