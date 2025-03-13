package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

// CustomTheme defines a custom theme
type CustomTheme struct{}

func (t *CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 255, G: 240, B: 240, A: 255} // Pink
	case theme.ColorNamePrimary:
		return color.RGBA{R: 0, G: 0, B: 0, A: 255} // Black
	case theme.ColorNameButton:
		return color.RGBA{R: 255, G: 255, B: 255, A: 255} // White
	case theme.ColorNameHover:
		return color.RGBA{R: 230, G: 230, B: 230, A: 255} // White
	case theme.ColorNameForeground: // text
		return color.RGBA{R: 0, G: 0, B: 0, A: 255} // Black
	case theme.ColorNameFocus:
		return color.RGBA{R: 255, G: 255, B: 255, A: 255} // White
	case theme.ColorNamePlaceHolder:
		return color.RGBA{R: 0, G: 0, B: 0, A: 255} // Black
	case theme.ColorNameInputBackground:
		return color.RGBA{R: 255, G: 255, B: 255, A: 255} // White
	case theme.ColorNameSelection: // disabling for now
		return color.RGBA{R: 230, G: 230, B: 230, A: 255} // White
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (t *CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t *CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
