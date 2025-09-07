package moduleFiles

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"hrubos.dev/collectorsden/internal/indexer"
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
	onCheck func(name string, checked bool),
	getNodeFromUID func(uid string) indexer.Node,
	setIndexedCheck func(uid string) bool,
) {

	logger.Log("Binding functions to tree", logger.CatView)

	v.treeWidget.ChildUIDs = func(uid widget.TreeNodeID) []widget.TreeNodeID {
		return childUIDs(uid)
	}
	v.treeWidget.IsBranch = func(uid widget.TreeNodeID) bool {
		return isBranch(uid)
	}
	v.treeWidget.CreateNode = func(branch bool) fyne.CanvasObject {
		//TODO refactor this, just add icon to files, the rest will be same for both branch and file
		if !branch {
			return container.NewHBox(
				widget.NewIcon(theme.FileIcon()),
				widget.NewCheck("", nil),
				widget.NewLabel(""),
			)
		}
		return container.NewHBox(
			widget.NewLabel(""),
			widget.NewCheck("", nil),
		)
	}
	v.treeWidget.UpdateNode = func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
		hbox, ok := node.(*fyne.Container)
		if !ok {
			panic(1) //TODO better error message/code
		}

		//TODO refactor this, just add icon to files, the rest will be same for both branch and file
		if !branch {
			//! warning: these need to be in the exact order of the .CreateNode container, its quite ugly hack but it works
			icon := hbox.Objects[0].(*widget.Icon)
			check := hbox.Objects[1].(*widget.Check)
			label := hbox.Objects[2].(*widget.Label)

			name := getName(uid)
			nodeName := getNodeFromUID(name).Name()
			label.SetText(nodeName)
			icon.SetResource(theme.FileIcon())

			// this fixes onchanged firing during re-rendering of the tree
			check.OnChanged = nil
			check.SetChecked(setIndexedCheck(uid))
			check.OnChanged = func(checked bool) {
				onCheck(name, checked)
			}
		} else {
			//! warning: these need to be in the exact order of the .CreateNode container, its quite ugly hack but it works
			label := hbox.Objects[0].(*widget.Label)
			check := hbox.Objects[1].(*widget.Check)

			name := getName(uid)
			node := getNodeFromUID(name)
			path := node.GetPath()
			label.SetText(node.Name())

			// this fixes onchanged firing during re-rendering of the tree
			check.OnChanged = nil
			check.SetChecked(setIndexedCheck(uid))
			check.OnChanged = func(checked bool) {
				onCheck(path, checked)
			}
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
