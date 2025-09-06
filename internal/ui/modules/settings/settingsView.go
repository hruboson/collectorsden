package moduleSettings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"hrubos.dev/collectorsden/internal/config"
	"hrubos.dev/collectorsden/internal/ui/components"
)

// Implement fyne.widget
var _ fyne.Widget = (*View)(nil)

type View struct {
	widget.BaseWidget
	container *fyne.Container

	themeToggle *widget.Check
	exportButton *widget.Button
	exportEntry *widget.Entry
}

func NewView() *View {
	v := &View{
		themeToggle: widget.NewCheck("Dark Theme", nil),
		exportButton: widget.NewButton("Export database", nil),
		exportEntry: widget.NewEntry(),
	}

	v.themeToggle.SetChecked(config.DarkThemeOn)

	v.container = container.NewVBox(
		widget.NewLabel("Select theme:"),
		v.themeToggle,
		components.NewThreeFourthOneFourth(
			v.exportEntry,
			v.exportButton,
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

func (v *View) SetThemeToggleChangeHandler(f func(bool)) {
	v.themeToggle.OnChanged = f
}

func (v *View) SetExportButtonOnClick(f func()){
	v.exportButton.OnTapped = f
}

// ----- Text setters -----

func (v *View) SetExportEntryPlaceHolder(s string){
	v.exportEntry.SetPlaceHolder(s)
}
