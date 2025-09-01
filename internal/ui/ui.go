package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	indexer "hrubos.dev/collectorsden/internal/indexer"
	logger "hrubos.dev/collectorsden/internal/logger"
)

var INITIAL_WINDOW_WIDTH float32 = 1200
var INITIAL_WINDOW_HEIGHT float32 = 600
var WINDOW_WIDTH float32 = INITIAL_WINDOW_WIDTH
var WINDOW_HEIGHT float32 = INITIAL_WINDOW_HEIGHT

func Run(){
	logger.Log("Starting UI", logger.CatUI)

	app := app.NewWithID("hrubos.dev/collectorsden")
	app.Settings().SetTheme(&darkTheme{})
	mainWindow := app.NewWindow("tree")

	rootDirEntryWidget := newRootDirEntry()

	tree := newFileTree()
	topBar := newTopBar(rootDirEntryWidget, tree, mainWindow)
	statusLabel := widget.NewLabel("")

	content := container.NewBorder(topBar, statusLabel, nil, nil, tree)

	mainWindow.SetContent(content)
	mainWindow.Resize(fyne.NewSize(INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT))
	mainWindow.ShowAndRun()

	logger.Log("Closing UI", logger.CatUI)
}


// --------------------- Components ---------------------

func newRootDirEntry() *widget.Entry {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter directory path...")
	return entry
}

func newFileTree() *widget.Tree {
	tree := widget.NewTree(nil, nil, nil, nil)
	return tree
}

func newTopBar(entry *widget.Entry, tree *widget.Tree, window fyne.Window) *fyne.Container {
	browseBtn := widget.NewButton("Browse", func() {
		callback := func(uri fyne.ListableURI, err error) {
			if uri != nil {
				entry.SetText(uri.Path())
				refreshTreeFromEntry(tree, entry)
			}
		}

		size := window.Canvas().Size()
		width := size.Width
		height := size.Height
		dialogWidth := width * 0.66
		dialogHeight := height * 0.66 

		folderDialog := dialog.NewFolderOpen(callback, window)  
		folderDialog.Resize(fyne.NewSize(dialogWidth, dialogHeight))
		folderDialog.Show()
	})

	entry.OnSubmitted = func(_ string) {
		refreshTreeFromEntry(tree, entry)
	}

	return container.NewBorder(nil, nil, nil, browseBtn, entry)
}

// --------------------- Logic ---------------------

func refreshTreeFromEntry(tree *widget.Tree, entry *widget.Entry) {
	root := entry.Text
	tree.Root = root
	refreshTree(tree, root)
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
		if !branch{
			return container.NewHBox(
				widget.NewIcon(theme.FileIcon()),
				widget.NewLabel(""),
			)
		}else{
			return container.NewHBox(
				widget.NewLabel(""),
			)
		}
	}

	// Set name for each node
	tree.UpdateNode = func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
		hbox := node.(*fyne.Container)
		if !branch {
			icon := hbox.Objects[0].(*widget.Icon)
			label := hbox.Objects[1].(*widget.Label)

			label.SetText(indexer.GetFileName(uid))
			icon.SetResource(theme.FileIcon())
		} else {
			label := hbox.Objects[0].(*widget.Label)
			label.SetText(indexer.GetFileName(uid))
		}
	}

	// If node is dir make branch
	tree.IsBranch = func(uid widget.TreeNodeID) (ok bool) {
		return indexer.IsDir(uid)
	}

	tree.Refresh()
}
