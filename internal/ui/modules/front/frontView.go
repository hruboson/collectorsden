package moduleTemplate

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Implement fyne.widget
var _ fyne.Widget = (*View)(nil)

type View struct {
	widget.BaseWidget
	container *fyne.Container
	btn *widget.Button
}

func NewView() *View {
	v := &View{
		btn: widget.NewButton("Btn", nil),
	}

	v.container = container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		v.btn,
	)

    v.ExtendBaseWidget(v) // Important so Fyne knows it's a widget

	return v
}

func (v *View) CreateRenderer() fyne.WidgetRenderer {
    return widget.NewSimpleRenderer(v.container)
}

// ----- Data setters -----

// ----- Callback setters

func (v *View) SetBrowseButtonOnTapped(f func()) {
	v.browserBtnWidget.OnTapped = f	
}

// ----- Text setters -----
