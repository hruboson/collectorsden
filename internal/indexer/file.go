package indexer

type File struct {
	Filename string
	Filetype string // todo change type later

	Parent Node
}

func (f *File) Accept(v Visitor) { v.visitFile(f) }

func (f *File) GetFiletype() string { return "FILE" } // todo
func (f *File) GetFilename() string { return f.Filename }
func (f *File) Name() string { return f.GetFilename() }

func (f *File) SetParent(p Node)   { f.Parent = p }
func (f *File) GetParent() Node    { return f.Parent }
