package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	_ "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	logger "hrubos.dev/collectorsden/internal/logger"
	indexer "hrubos.dev/collectorsden/internal/indexer"
)

func Run(){
	logger.Log("Starting UI", logger.CatUI)

	mainApp := app.New()
	mainWindow := mainApp.NewWindow("tree")

	root := "E:\\Archive"

	lb := widget.NewLabel("")

	dirEntry := widget.NewEntry()
	dirEntry.Text = root

	tree := widget.NewTree(nil, nil, nil, nil)
	tree.Root = root
	refreshTree(tree, dirEntry.Text)

	tree.OnSelected = func(uid widget.TreeNodeID) {
		lb.SetText(uid)
	}

	mainWindow.SetContent(tree)
	mainWindow.Resize(fyne.NewSize(550, 450))
	mainWindow.ShowAndRun()

	logger.Log("Closing UI", logger.CatUI)
}


func refreshTree(tree *widget.Tree, root string) {
	tree.Root = root

	tree.ChildUIDs = func(uid widget.TreeNodeID) (c []widget.TreeNodeID) {
		if uid == "" {
			c = indexer.GetFiles(root)
		} else {
			c = indexer.GetFiles(uid)
		}
		return
	}

	// Create new Node (Node is Label component)
	tree.CreateNode = func(branch bool) (o fyne.CanvasObject) {
		return widget.NewLabel("")
	}

	// Set name for each node
	tree.UpdateNode = func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
		l := node.(*widget.Label)
		l.SetText(indexer.GetFileName(uid))
	}

	// If node is dir make branch
	tree.IsBranch = func(uid widget.TreeNodeID) (ok bool) {
		return indexer.IsDir(uid)
	}
	
	tree.Refresh()
}

// Recursive helper to render FileStructure as text
func renderFileStructure(fs *indexer.FileStructure, indent string) string {
	result := ""
	for _, node := range fs.Nodes {
		result += renderNode(node, indent)
	}
	return result
}

func renderNode(node indexer.Node, indent string) string {
	switch n := node.(type) {
	case *indexer.Folder:
		str := indent + "+ Folder: " + n.GetFoldername() + "\n"
		for _, child := range n.Nodes {
			str += renderNode(child, indent+"  ")
		}
		return str
	case *indexer.File:
		return indent + "- File: " + n.GetFiletype() + " " + n.Filename + "\n"
	default:
		return indent + "? Unknown node type\n"
	}
}


// LOADING VERSION IN Ui.Run()
/*a := app.New()
w := a.NewWindow("File Structure")

/*loadingLabel := widget.NewLabel("Indexing files, please wait...")
scroll := container.NewScroll(loadingLabel)
w.SetContent(scroll)
w.Resize(fyne.NewSize(600, 400))

indx := indexer.NewIndexer()

go func() {
	err := indx.TreeDir("G:\\Projects")
	if err != nil {
		logger.Fatal(err)
	}

	// Prepare the textual tree
	text := renderFileStructure(&indx.FileStructure, "")

	fyne.Do(func(){
		// Update UI on the main thread
		w.Canvas().Refresh(scroll) // optional
		w.SetContent(container.NewScroll(widget.NewLabel(text)))
	})
}()*/
