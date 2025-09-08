package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"

	"hrubos.dev/collectorsden/internal/config"
)

func NewThemedIconButton(iconName fyne.ThemeIconName, tapped func()) *widget.Button {
	btn := widget.NewButtonWithIcon("", theme.DefaultTheme().Icon(iconName), tapped)

    btn.SetIcon(config.AppSettings.Theme().Icon(iconName))

	// Listen for theme changes and update the icon
	fyne.CurrentApp().Settings().AddListener(func(s fyne.Settings) {
		btn.SetIcon(config.AppSettings.Theme().Icon(iconName))
	})

	return btn
}
