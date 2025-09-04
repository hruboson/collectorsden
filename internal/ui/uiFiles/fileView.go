package uiFiles

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"hrubos.dev/collectorsden/internal/logger"
)

// Implement fyne.widget
var _ fyne.Widget = (*View)(nil)

type View struct {
	widget.BaseWidget
	container *fyne.Container
	treeWidget *widget.Tree
	rootDirEntryWidget *widget.Entry
	browserBtnWidget *widget.Button
	settingsBtnWidget *widget.Button
	statusLabel *widget.Label
}

func NewView() *View {
	fv := &View{
		treeWidget: widget.NewTree(nil, nil, nil, nil),
		rootDirEntryWidget: widget.NewEntry(),
		browserBtnWidget: widget.NewButton("Browse", nil),
		settingsBtnWidget: widget.NewButton("Settings", nil),
		statusLabel: widget.NewLabel(""),
	}

	// default values
	fv.rootDirEntryWidget.SetPlaceHolder("Enter directory path...")

	fv.container = container.NewBorder(
		container.NewBorder(
			nil,
			nil,
			nil,
	        container.NewHBox(fv.browserBtnWidget, fv.settingsBtnWidget),
			fv.rootDirEntryWidget,
		),
		fv.statusLabel,
		nil,
		nil,
		fv.treeWidget,
	)

    fv.ExtendBaseWidget(fv) // Important so Fyne knows it's a widget

	return fv
}

func (fv *View) CreateRenderer() fyne.WidgetRenderer {
    return widget.NewSimpleRenderer(fv.container)
}

// ----- Data setters -----

func (fv *View) BindTree(
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

func (fv *View) SwitchTreeRoot(root string){
	logger.Log("Switching tree root", logger.CatView)

	fv.treeWidget.Root = root
	fv.treeWidget.Refresh()
}

// ----- Callback setters

func (fv *View) SetBrowseButtonOnTapped(f func()) {
	fv.browserBtnWidget.OnTapped = f	
}

func (fv *View) SetSettingsButtonOnTapped(f func()) {
	fv.settingsBtnWidget.OnTapped = f
}

func (fv *View) SetTreeWidgetOnSelected(f func(uid widget.TreeNodeID)) {
	fv.treeWidget.OnSelected = f
}

// ----- Text setters -----

func (fv *View) RootDirEntryWidgetSetText(text string) {
	fv.rootDirEntryWidget.SetText(text)
}

func (fv *View) StatusLabelSetText(text string) {
	fv.statusLabel.SetText(text)
}

// ----- Getters -----
