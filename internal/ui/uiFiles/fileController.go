package uiFiles

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Controller struct {
	*Model
	*View
	window fyne.Window
}

func NewController(fm *Model, fv *View, window fyne.Window) *Controller {
	fc := &Controller{
		Model: fm,
		View: fv,
		window: window,
	}

	fc.View.SetBrowseButtonOnTapped(fc.browseFiles)
	fc.View.SetTreeWidgetOnSelected(func(uid widget.TreeNodeID) {
		fc.View.StatusLabelSetText(uid)
	})


	fc.bindFileTree()

	return fc
}

func (fc *Controller) browseFiles(){
	callback := func(uri fyne.ListableURI, err error) {
		if uri != nil {
			fc.Model.SetRoot(uri.Path())
			fc.View.RootDirEntryWidgetSetText(uri.Path())

			fc.View.SwitchTreeRoot(fc.Model.GetRoot())
            fc.View.StatusLabelSetText(uri.Path())
		}
	}

	size := fc.window.Canvas().Size()
	width := size.Width
	height := size.Height
	dialogWidth := width * 0.66
	dialogHeight := height * 0.66

	folderDialog := dialog.NewFolderOpen(callback, fc.window)
	folderDialog.Resize(fyne.NewSize(dialogWidth, dialogHeight))
	folderDialog.Show()
}

func (fc *Controller) bindFileTree() {
	childUIDs, isBranch, getName := fc.Model.TreeData()
	fc.View.BindTree(childUIDs, isBranch, getName)
}
