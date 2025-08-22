package indexer

type Folder struct {
	Foldername string

	Parent Node
	Nodes []Node
}

func (f *Folder) Accept(v Visitor) { v.visitFolder(f) }

func (f *Folder) GetFoldername() string { return f.Foldername }
func (f *Folder) Name() string { return f.GetFoldername() }

func (f *Folder) SetParent(p Node) { f.Parent = p }
func (f *Folder) GetParent() Node  { return f.Parent }
