package indexer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/karrick/godirwalk"

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

// Populates the FileStructure in Indexer struct
func (i *Indexer) TreeDir(directory string) error {
	baseDepth := len(strings.Split(filepath.Clean(directory), string(filepath.Separator)))

	root := &Folder{
		Foldername: filepath.Base(directory),
		Nodes: []Node{},
	}

	folders := map[string]*Folder{
		directory: root,
	}

	err := godirwalk.Walk(directory, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if de.IsDir() {
				currentDepth := len(strings.Split(filepath.Clean(osPathname), string(filepath.Separator)))
				hiddenFolder := isHidden(osPathname)
				if currentDepth - baseDepth > i.MaxDepth || hiddenFolder {
					return godirwalk.SkipThis
				}
			}

			parentPath := filepath.Dir(osPathname)
			parentFolder, ok := folders[parentPath]
			if !ok {
				// should not happen unless walking strange symlinks
				return nil
			}

			if de.IsDir() {
				folder := &Folder{
					Foldername: de.Name(),
					Nodes:    []Node{},
				}
				folder.SetParent(parentFolder)

				parentFolder.Nodes = append(parentFolder.Nodes, folder)
				folders[osPathname] = folder // add to map for children
			} else {
				file := &File{
					Filename: de.Name(),
					Filetype: filepath.Ext(de.Name()),
				}
				file.SetParent(parentFolder)

				parentFolder.Nodes = append(parentFolder.Nodes, file)
			}

			return nil
		},
	})
	
	i.FileStructure.Nodes = root.Nodes
	return err
}

/**********************
* PRIMITIVE FUNCTIONS *
**********************/

// return files from path
func GetFileName(path string) string {
	fi, err := os.Stat(path)
	if err != nil {
		return ""
	}
	return fi.Name()
}

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func GetFiles(path string) (list []string) {
	fi, err := os.ReadDir(path)
	if err != nil {
		msg := fmt.Sprintf("Error while geting file from %s", path)
		logger.Log(logger.CatIndexer, msg)
		return
	}

	for _, file := range fi {
		filename := path + string(filepath.Separator) + file.Name()
		if(isHidden(filename)){
			continue
		}
		list = append(list, filename)
	}
	return
}

func isHidden(path string) bool {
    p, err := syscall.UTF16PtrFromString(path)
    if err != nil {
		msg := fmt.Sprintf("Invalid path string '%s': contains NUL characters", path)
		logger.Log(logger.CatIndexer, msg)
        return false
    }
    attrs, err := syscall.GetFileAttributes(p)
    if err != nil {
		msg := fmt.Sprintf("Failed to get file attributes for '%s': %v (while determining if the file is hidden)", path, err)
		logger.Log(logger.CatIndexer, msg)
        return false
    }
    const FILE_ATTRIBUTE_HIDDEN = 0x2
    return attrs&FILE_ATTRIBUTE_HIDDEN != 0
}
