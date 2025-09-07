package moduleFiles

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	indexer "hrubos.dev/collectorsden/internal/indexer"
	logger "hrubos.dev/collectorsden/internal/logger"
	moduleSettings "hrubos.dev/collectorsden/internal/ui/modules/settings"
)

type Controller struct {
	*Model
	*View
	window fyne.Window
	app fyne.App
}

func NewController(m *Model, v *View, app fyne.App, window fyne.Window) *Controller {
	c := &Controller{
		Model: m,
		View: v,
		window: window,
		app: app,
	}

	c.View.SetBrowseButtonOnTapped(c.browseFiles)
	c.View.SetSettingsButtonOnTapped(c.openSettingsWindow)
	c.View.SetEntryOnSubmitted(c.onEntrySubmit)
	c.View.SetTreeWidgetOnSelected(func(uid widget.TreeNodeID) {
		c.View.StatusLabelSetText(uid)
	})


	c.bindFileTree()

	return c
}

func (c *Controller) browseFiles(){
	callback := func(uri fyne.ListableURI, err error) {
		if uri != nil {
			c.Model.SetRoot(uri.Path())
			c.View.RootDirEntryWidgetSetText(uri.Path())

			c.View.SwitchTreeRoot(c.Model.GetRoot())
            c.View.StatusLabelSetText(uri.Path())
		}
	}

	size := c.window.Canvas().Size()
	width := size.Width
	height := size.Height
	dialogWidth := width * 0.66
	dialogHeight := height * 0.66

	folderDialog := dialog.NewFolderOpen(callback, c.window)
	folderDialog.Resize(fyne.NewSize(dialogWidth, dialogHeight))
	folderDialog.Show()
}

func (c *Controller) bindFileTree() {
	childUIDs, isBranch, getName := c.Model.TreeData()
	c.View.BindTree(childUIDs, isBranch, getName, c.Model.CheckNode, c.getNodeFromUID)
}

func (c *Controller) openSettingsWindow() {
    settingsWindow := fyne.CurrentApp().NewWindow("Settings")

	settingsModel := moduleSettings.NewModel()
	settingsView := moduleSettings.NewView()
	settingsController := moduleSettings.NewController(settingsModel, settingsView, c.app, c.window)

	settingsWindow.SetContent(settingsController.View)
    settingsWindow.Resize(fyne.NewSize(400, 300))
	settingsWindow.CenterOnScreen()

	logger.Log("Opening settings window", logger.CatUI)
    settingsWindow.Show()
}

func (c *Controller) onEntrySubmit(text string) {
	c.Model.SetRoot(text)
	c.View.RootDirEntryWidgetSetText(text)

	c.View.SwitchTreeRoot(c.Model.GetRoot())
	c.View.StatusLabelSetText(text)
}

func (c *Controller) getNodeFromUID(uid string) indexer.Node {
	return c.Model.GetNodeFromUID(uid)
}
