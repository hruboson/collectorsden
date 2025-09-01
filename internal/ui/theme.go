package ui

import (
    "image/color"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/theme"
)

type darkTheme struct {
    fyne.Theme
}

// Overriding colors
func (m *darkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
    if name == theme.ColorNameBackground {
        return color.NRGBA{R: 30, G: 30, B: 30, A: 255} // dark gray background
    }
    return theme.DefaultTheme().Color(name, variant)
}

// Keep same fonts
func (d *darkTheme) Font(style fyne.TextStyle) fyne.Resource {
	return resourceAssetsFontsRobotoRegularTtf
}

// Overriding sizes
func (m *darkTheme) Size(name fyne.ThemeSizeName) float32 {
    if name == theme.SizeNameText {
        return theme.DefaultTheme().Size(name) + 2 // larger text
    }
    return theme.DefaultTheme().Size(name)
}

// Overriding icons
func (m *darkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	switch name {
		case theme.IconNameMoveDown: return theme.FolderOpenIcon()
		case theme.IconNameNavigateNext: return theme.FolderIcon()
		default: return theme.DefaultTheme().Icon(name)
	}
}
