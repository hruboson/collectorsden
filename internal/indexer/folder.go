package indexer

import (
	"fmt"
	"os"
	"path/filepath"

	logger "hrubos.dev/collectorsden/internal/logger"
)

type Folder struct {
	fullpath string
	name string

	parent Node
	children  []Node
}

func NewFolder(path string, parent Node) *Folder {
	return &Folder{
		fullpath: path,
		name: filepath.Base(path),
		parent:     parent,
		children:      []Node{},
	}
}

func (f *Folder) Accept(v Visitor) { v.visitFolder(f) }

func (f *Folder) GetPath() string { return f.fullpath }
func (f *Folder) Name() string          { return f.name }

func (f *Folder) SetParent(p Node) { f.parent = p }
func (f *Folder) GetParent() Node  { return f.parent }

func (f *Folder) Type() NodeType { return FOLDER }

// Return child Nodes of a folder
func (f *Folder) GetChildren() []Node {
	// if folder does not yet have children
	if len(f.children) <= 0 {
		entries, err := os.ReadDir(f.GetPath())
		if err != nil {
			msg := fmt.Sprintf("Error while reading %s", f.GetPath())
			logger.Log(msg, logger.CatIndexer)
			return nil
		}


		logger.Log("Retrieved children from " + f.GetPath(), logger.CatIndexer)
		for _, entry := range entries {
			fullpath := filepath.Join(f.GetPath(), entry.Name())
			if isHidden(fullpath) {
				continue
			}

			if entry.IsDir() {
				folder := NewFolder(fullpath, f.parent)
				f.children = append(f.children, folder)
			} else {
				file := NewFile(fullpath, f.parent)
				f.children = append(f.children, file)
			}
		}
	}

	return f.children
}
