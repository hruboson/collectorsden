package moduleHome

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"hrubos.dev/collectorsden/internal/ui/bundled"
)

// Implement fyne.widget
var _ fyne.Widget = (*View)(nil)

type View struct {
	widget.BaseWidget
	container *fyne.Container

	btnNew *widget.Button
	btnOpen *widget.Button
	btnLastFront *widget.Button

	appLogo *canvas.Image
}

func NewView() *View {
	// TODO quick open last stores (under buttons)
	v := &View{
		btnNew: widget.NewButton("New Store", nil),
		btnOpen: widget.NewButton("Open existing", nil),
		btnLastFront: widget.NewButton("Open last Front", nil),
		
		appLogo: canvas.NewImageFromResource(bundled.ResourceAssetsImgIconPng),
	}

 	v.appLogo.FillMode = canvas.ImageFillOriginal

	v.container = container.New(
		layout.NewCenterLayout(),
		container.NewBorder(
			v.appLogo,
			nil,
			container.NewHBox(
				v.btnNew, 
				v.btnOpen,
				v.btnLastFront,
			),
			nil,
			nil,
		),
	)

    v.ExtendBaseWidget(v) // Important so Fyne knows it's a widget

	return v
}

func (v *View) CreateRenderer() fyne.WidgetRenderer {
    return widget.NewSimpleRenderer(v.container)
}

// ----- Data setters -----

// ----- Callback setters

func (v *View) SetOpenButtonOnTapped(f func()) {
	v.btnOpen.OnTapped = f	
}

// ----- Text setters -----
