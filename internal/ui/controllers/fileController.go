package controllers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"hrubos.dev/collectorsden/internal/ui/models"
	"hrubos.dev/collectorsden/internal/ui/views"
	"hrubos.dev/collectorsden/internal/logger"
)

type FileController struct {
	*models.FileModel
	*views.FileView
	window fyne.Window
}

func NewFileController(fm *models.FileModel, fv *views.FileView, window fyne.Window) *FileController {
	fc := &FileController{
		FileModel: fm,
		FileView: fv,
		window: window,
	}

	fc.FileView.SetBrowseButtonOnTapped(fc.browseFiles)
	fc.FileView.SetTreeWidgetOnSelected(func(uid widget.TreeNodeID) {
		fc.FileView.StatusLabelSetText(uid)
	})


	fc.bindFileTree()

	return fc
}

func (fc *FileController) browseFiles(){
	callback := func(uri fyne.ListableURI, err error) {
		if uri != nil {
			fc.FileModel.SetRoot(uri.Path())
			fc.FileView.RootDirEntryWidgetSetText(uri.Path())

			fc.FileView.SwitchTreeRoot(fc.FileModel.GetRoot())
            fc.FileView.StatusLabelSetText(uri.Path())
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

func (fc *FileController) bindFileTree() {
	logger.Log("Updating file tree data", logger.CatController)
	childUIDs, isBranch, getName := fc.FileModel.TreeData()
	fc.FileView.BindTree(childUIDs, isBranch, getName)
}
