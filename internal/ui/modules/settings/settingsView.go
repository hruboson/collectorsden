package moduleSettings

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

	themeToggle *widget.Check
}

func NewView() *View {
	v := &View{
		btn: widget.NewButton("Btn", nil),
	}

	v.themeToggle = widget.NewCheck("Dark Theme", nil)
	v.themeToggle.SetChecked(true)

	v.container = container.NewVBox(
		widget.NewLabel("Select theme:"),
		v.themeToggle,
	)

    v.ExtendBaseWidget(v) // Important so Fyne knows it's a widget

	return v
}

func (v *View) CreateRenderer() fyne.WidgetRenderer {
    return widget.NewSimpleRenderer(v.container)
}

// ----- Data setters -----

// ----- Callback setters

func (v *View) SetThemeToggleChangeHandler(f func(bool)) {
	v.themeToggle.OnChanged = f
}

// ----- Text setters -----
