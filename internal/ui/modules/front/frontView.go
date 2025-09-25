package moduleTemplate

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"hrubos.dev/collectorsden/internal/database"
	_ "hrubos.dev/collectorsden/internal/logger"
)

// Implement fyne.widget
var _ fyne.Widget = (*View)(nil)

type View struct {
	widget.BaseWidget

	list *widget.List

	container *fyne.Container
	btnRefresh *widget.Button
}

func NewView() *View {
	v := &View{
		btnRefresh: widget.NewButton("Refresh", nil),
		list: widget.NewList(nil, nil, nil),
	}

	v.container = container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		v.list,
	)

    v.ExtendBaseWidget(v) // Important so Fyne knows it's a widget

	return v
}

func (v *View) CreateRenderer() fyne.WidgetRenderer {
    return widget.NewSimpleRenderer(v.container)
}

// ----- Callback setters

func (v *View) BindList(
	getCategory func(i int) database.Category,
	getLength func() int,
){
	v.list.Length = 
		func() int {
			return getLength()
		}
	v.list.CreateItem = 
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		}

		//TODO check for null
	v.list.UpdateItem = 
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(getCategory(i).FullPath)
		}
}

func (v *View) SetBrowseButtonOnTapped(f func()) {
	v.btnRefresh.OnTapped = f	
}

// ----- Text setters -----
