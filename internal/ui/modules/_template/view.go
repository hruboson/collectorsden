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
	fv := &View{
		btn: widget.NewButton("Btn", nil)
	}

	fv.container = container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		fv.btn,
	)

    fv.ExtendBaseWidget(fv) // Important so Fyne knows it's a widget

	return fv
}

func (fv *View) CreateRenderer() fyne.WidgetRenderer {
    return widget.NewSimpleRenderer(fv.container)
}

// ----- Data setters -----

// ----- Callback setters

func (fv *View) SetBrowseButtonOnTapped(f func()) {
	fv.browserBtnWidget.OnTapped = f	
}

// ----- Text setters -----
