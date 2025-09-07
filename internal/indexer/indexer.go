package indexer

import (
	"os"

	logger "hrubos.dev/collectorsden/internal/logger"
)

type Indexer struct {
	MaxDepth int
	IndexHiddenFiles bool

	FileStructure FileStructure
}

func NewIndexer() *Indexer {
	return &Indexer{
		MaxDepth:    		3,
		IndexHiddenFiles: 	false,
	}
}

type FileStructure struct {
	Nodes []Node
	LastUpdated int
}

func (fs *FileStructure) Print(indent string) {
	for _, node := range fs.Nodes { // iterate all top-level nodes
		printNode(node, indent)
	}
}

/**********************
* PRIMITIVE FUNCTIONS *
**********************/

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		logger.Log("Error while determining whether " + path + " is folder", logger.CatIndexer)
		return false
	}
	return fi.IsDir()
}

// DEPRECATED

/*func GetFileName(path string) string {
	fi, err := os.Stat(path)
	if err != nil {
		return ""
	}
	return fi.Name()
}

func GetFiles(path string) (list []string) {
	fi, err := os.ReadDir(path)
	if err != nil {
		msg := fmt.Sprintf("Error while geting file from %s", path)
		logger.Log(msg, logger.CatIndexer)
		return
	}

	for _, file := range fi {
		filename := path + string(filepath.Separator) + file.Name()
		if(isHidden(filename)){
			continue
		}
		list = append(list, filename)
	}

	logger.Log("Retrieving files from " + path, logger.CatIndexer)
	return
}*/

func isHidden(path string) bool {
	// TODO multiplatform syscall
	/*
    p, err := syscall.UTF16PtrFromString(path)
    if err != nil {
		msg := fmt.Sprintf("Invalid path string '%s': contains NUL characters", path)
		logger.Log(msg, logger.CatIndexer)
        return false
    }
    attrs, err := syscall.GetFileAttributes(p)
    if err != nil {
		msg := fmt.Sprintf("Failed to get file attributes for '%s': %v (while determining if the file is hidden)", path, err)
		logger.Log(msg, logger.CatIndexer)
        return false
    }

    const FILE_ATTRIBUTE_HIDDEN = 0x2
    return attrs&FILE_ATTRIBUTE_HIDDEN != 0*/
	return false
}
