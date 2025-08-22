package indexer

import (
	"path/filepath"
	"strings"
	"syscall"

	"github.com/karrick/godirwalk"

	_ "hrubos.dev/collectorsden/internal/logger"
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
				hiddenFolder, err := isHidden(osPathname)
				if err != nil || (currentDepth - baseDepth > i.MaxDepth || hiddenFolder) {
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

func isHidden(path string) (bool, error) {
    p, err := syscall.UTF16PtrFromString(path)
    if err != nil {
        return false, err
    }
    attrs, err := syscall.GetFileAttributes(p)
    if err != nil {
        return false, err
    }
    const FILE_ATTRIBUTE_HIDDEN = 0x2
    return attrs&FILE_ATTRIBUTE_HIDDEN != 0, nil
}
