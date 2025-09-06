package moduleFiles

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	logger "hrubos.dev/collectorsden/internal/logger"
	moduleSettings "hrubos.dev/collectorsden/internal/ui/modules/settings"
)

type Controller struct {
	*Model
	*View
	window fyne.Window
	app fyne.App
}

func NewController(fm *Model, fv *View, app fyne.App, window fyne.Window) *Controller {
	fc := &Controller{
		Model: fm,
		View: fv,
		window: window,
		app: app,
	}

	fc.View.SetBrowseButtonOnTapped(fc.browseFiles)
	fc.View.SetSettingsButtonOnTapped(fc.openSettingsWindow)
	fc.View.SetEntryOnSubmitted(fc.onEntrySubmit)
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

func (fc *Controller) openSettingsWindow() {
    settingsWindow := fyne.CurrentApp().NewWindow("Settings")

	settingsModel := moduleSettings.NewModel()
	settingsView := moduleSettings.NewView()
	settingsController := moduleSettings.NewController(settingsModel, settingsView, fc.app, fc.window)

	settingsWindow.SetContent(settingsController.View)
    settingsWindow.Resize(fyne.NewSize(400, 300))
	settingsWindow.CenterOnScreen()

	logger.Log("Opening settings window", logger.CatUI)
    settingsWindow.Show()
}

func (fc *Controller) onEntrySubmit(text string){
	fc.Model.SetRoot(text)
	fc.View.RootDirEntryWidgetSetText(text)

	fc.View.SwitchTreeRoot(fc.Model.GetRoot())
	fc.View.StatusLabelSetText(text)
}
