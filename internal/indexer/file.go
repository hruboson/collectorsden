package indexer

import "path/filepath"

type File struct {
	fullpath string
	filename string
	filetype string // todo change type later

	parent Node
}

func NewFile(path string, parent Node) *File {
	return &File{
		fullpath: path,
		filename: filepath.Base(path),
		filetype: filepath.Ext(path), // TODO check cross-platform
		parent:   parent,
	}
}

func (f *File) Accept(v Visitor) { v.visitFile(f) }

func (f *File) Type() NodeType { return FILE } // TODO
func (f *File) GetFilename() string { return f.filename }
func (f *File) GetPath() string { return f.fullpath }
func (f *File) Name() string        { return f.GetFilename() }

func (f *File) SetParent(p Node) { f.parent = p }
func (f *File) GetParent() Node  { return f.parent }
