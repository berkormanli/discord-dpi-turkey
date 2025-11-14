package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// CustomTheme implements a custom Fyne theme
type CustomTheme struct {
	isDark bool
}

// NewCustomTheme creates a new custom theme
func NewCustomTheme(isDark bool) fyne.Theme {
	return &CustomTheme{isDark: isDark}
}

func (t *CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if t.isDark {
		return theme.DarkTheme().Color(name, variant)
	}
	return theme.LightTheme().Color(name, variant)
}

func (t *CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	if t.isDark {
		return theme.DarkTheme().Font(style)
	}
	return theme.LightTheme().Font(style)
}

func (t *CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	if t.isDark {
		return theme.DarkTheme().Icon(name)
	}
	return theme.LightTheme().Icon(name)
}

func (t *CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	if t.isDark {
		return theme.DarkTheme().Size(name)
	}
	return theme.LightTheme().Size(name)
}
