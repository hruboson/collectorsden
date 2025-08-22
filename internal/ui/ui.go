package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	logger "hrubos.dev/collectorsden/internal/logger"
	indexer "hrubos.dev/collectorsden/internal/indexer"
)

func Run(){
	logger.Log("Starting UI", logger.CatUI)

	a := app.New()
	w := a.NewWindow("File Structure")

	loadingLabel := widget.NewLabel("Indexing files, please wait...")
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
	}()

    w.ShowAndRun() // Keep the app running until user closes
	logger.Log("Closing UI", logger.CatUI)
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
