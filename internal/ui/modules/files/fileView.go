package moduleFiles

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"hrubos.dev/collectorsden/internal/indexer"
	"hrubos.dev/collectorsden/internal/logger"
	"hrubos.dev/collectorsden/internal/ui/components"
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
	v := &View{
		treeWidget: widget.NewTree(nil, nil, nil, nil),
		rootDirEntryWidget: widget.NewEntry(),
		browserBtnWidget: widget.NewButton("Browse", nil),
		settingsBtnWidget: widget.NewButton("Settings", nil),
		statusLabel: widget.NewLabel(""),
	}

	// default values
	v.rootDirEntryWidget.SetPlaceHolder("Enter directory path...")

	v.container = container.NewBorder(
		container.NewBorder(
			nil,
			nil,
			nil,
	        container.NewHBox(v.browserBtnWidget, v.settingsBtnWidget),
			v.rootDirEntryWidget,
		),
		v.statusLabel,
		nil,
		nil,
		v.treeWidget,
	)

    v.ExtendBaseWidget(v) // Important so Fyne knows it's a widget

	return v
}

func (v *View) CreateRenderer() fyne.WidgetRenderer {
    return widget.NewSimpleRenderer(v.container)
}

// ----- Data setters -----

func (v *View) BindTree(
	childUIDs func(uid string) []string,
	isBranch func(uid string) bool,
	getName func(uid string) string,

	getNodeFromUID func(uid string) indexer.Node,

	onCheckFunction func(name string, checked bool),
	setIndexedCheck func(uid string) bool,
	defaultOpenFunction func(path string) error,
) {

	logger.Log("Binding functions to tree", logger.CatView)

	v.treeWidget.ChildUIDs = func(uid widget.TreeNodeID) []widget.TreeNodeID {
		return childUIDs(uid)
	}
	v.treeWidget.IsBranch = func(uid widget.TreeNodeID) bool {
		return isBranch(uid)
	}
	v.treeWidget.CreateNode = func(branch bool) fyne.CanvasObject {
		elements := []fyne.CanvasObject {
			widget.NewCheck("", nil),
			widget.NewLabel(""),
			layout.NewSpacer(),
			components.NewThemedIconButton("link", nil),
		}

		if !branch {
			elements = append(
				[]fyne.CanvasObject{widget.NewIcon(theme.FileIcon())}, 
				elements...
			)
		}
		
		return container.NewHBox(elements...)
	}
	v.treeWidget.UpdateNode = func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
		hbox, ok := node.(*fyne.Container)
		if !ok {
			err := errors.New("Could not convert fyne.CanvasObject to fyne.Container")
			logger.Fatal("Error while updating node in file tree", err, logger.CatView)
		}

		type branchWidgets struct {
			Icon             *widget.Icon   // only for files
			Check            *widget.Check
			Label            *widget.Label
			DefaultOpenButton *widget.Button
		}

		var widgets branchWidgets
		if branch {
			// branch layout
			widgets = branchWidgets{
				Check:             hbox.Objects[0].(*widget.Check),
				Label:             hbox.Objects[1].(*widget.Label),
				DefaultOpenButton: hbox.Objects[3].(*widget.Button),
			}
		} else {
			// file layout
			widgets = branchWidgets{
				Icon:              hbox.Objects[0].(*widget.Icon),
				Check:             hbox.Objects[1].(*widget.Check),
				Label:             hbox.Objects[2].(*widget.Label),
				DefaultOpenButton: hbox.Objects[4].(*widget.Button),
			}
		}

		name := getName(uid)
		nodeInfo := getNodeFromUID(name)
		labelText := nodeInfo.Name()
		checkTarget := name
		if branch {
			checkTarget = nodeInfo.GetPath()
		}

		widgets.Label.SetText(labelText)
		if widgets.Icon != nil { // icon only for files
			widgets.Icon.SetResource(theme.FileIcon())
		}

		widgets.DefaultOpenButton.OnTapped = func() {
			defaultOpenFunction(uid)
		}

		widgets.Check.OnChanged = nil // keep this... it fixes onchanged firing during re-rendering of the tree
		widgets.Check.SetChecked(setIndexedCheck(uid))
		widgets.Check.OnChanged = func(checked bool) {
			onCheckFunction(checkTarget, checked)
		}
	}
}

func (v *View) SwitchTreeRoot(root string){
	logger.Log("Switching tree root", logger.CatView)

	v.treeWidget.Root = root
	v.treeWidget.Refresh()
}

// ----- Callback setters -----

func (v *View) SetBrowseButtonOnTapped(f func()) {
	v.browserBtnWidget.OnTapped = f	
}

func (v *View) SetSettingsButtonOnTapped(f func()) {
	v.settingsBtnWidget.OnTapped = f
}

func (v *View) SetTreeWidgetOnSelected(f func(uid widget.TreeNodeID)) {
	v.treeWidget.OnSelected = f
}

func (v *View) SetEntryOnSubmitted(f func(text string)) {
	v.rootDirEntryWidget.OnSubmitted = f
}

// ----- Text/Status setters -----

func (v *View) RootDirEntryWidgetSetText(text string) {
	v.rootDirEntryWidget.SetText(text)
}

func (v *View) StatusLabelSetText(text string) {
	v.statusLabel.SetText(text)
}
