package uiFiles

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"hrubos.dev/collectorsden/internal/logger"
)

// Implement fyne.widget
var _ fyne.Widget = (*FileView)(nil)

type FileView struct {
	widget.BaseWidget
	treeWidget *widget.Tree
	rootDirEntryWidget *widget.Entry
	browserBtnWidget *widget.Button
	statusLabel *widget.Label
}

func NewFileView() *FileView {
	fv := &FileView{
		treeWidget: widget.NewTree(nil, nil, nil, nil),
		rootDirEntryWidget: widget.NewEntry(),
		browserBtnWidget: widget.NewButton("Browse", nil),
		statusLabel: widget.NewLabel(""),
	}

	// default values
	fv.rootDirEntryWidget.SetPlaceHolder("Enter directory path...")

	return fv
}

func (fv *FileView) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(
		container.NewBorder(
			nil,
			nil,
			nil,
			fv.browserBtnWidget,
			fv.rootDirEntryWidget,
		),
		fv.statusLabel,
		nil,
		nil,
		fv.treeWidget,
	)
	return widget.NewSimpleRenderer(c)
}

// ----- Data setters -----

func (fv *FileView) BindTree(
	childUIDs func(uid string) []string,
	isBranch func(uid string) bool,
	getName func(uid string) string,
) {

	logger.Log("Binding functions to tree", logger.CatView)

	fv.treeWidget.ChildUIDs = func(uid widget.TreeNodeID) []widget.TreeNodeID {
		return childUIDs(uid)
	}
	fv.treeWidget.IsBranch = func(uid widget.TreeNodeID) bool {
		return isBranch(uid)
	}
	fv.treeWidget.CreateNode = func(branch bool) fyne.CanvasObject {
		if !branch {
			return container.NewHBox(widget.NewIcon(theme.FileIcon()), widget.NewLabel(""))
		}
		return container.NewHBox(widget.NewLabel(""))
	}
	fv.treeWidget.UpdateNode = func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
		hbox := node.(*fyne.Container)
		if !branch {
			icon := hbox.Objects[0].(*widget.Icon)
			label := hbox.Objects[1].(*widget.Label)
			label.SetText(getName(uid))
			icon.SetResource(theme.FileIcon())
		} else {
			label := hbox.Objects[0].(*widget.Label)
			label.SetText(getName(uid))
		}
	}
}

func (fv *FileView) SwitchTreeRoot(root string){
	logger.Log("Switching tree root", logger.CatView)

	fv.treeWidget.Root = root
	fv.treeWidget.Refresh()
}

// ----- Callback setters

func (fv *FileView) SetBrowseButtonOnTapped(f func()) {
	fv.browserBtnWidget.OnTapped = f	
}

func (fv *FileView) SetTreeWidgetOnSelected(f func(uid widget.TreeNodeID)) {
	fv.treeWidget.OnSelected = f
}

// ----- Text setters -----

func (fv *FileView) RootDirEntryWidgetSetText(text string) {
	fv.rootDirEntryWidget.SetText(text)
}

func (fv *FileView) StatusLabelSetText(text string) {
	fv.statusLabel.SetText(text)
}

// ----- Getters -----
