package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (app *AppConfig) makeUI() {
	var top, bottom, left, right, center fyne.CanvasObject

	app.makeHeader()
	list := app.makeRulesList()
	addRule := app.makeAddRule()
	center = container.NewBorder(addRule, nil, nil, nil, list)
	bottom = app.StatusLabel

	appContent := container.NewBorder(top, bottom, left, right, center)

	// add container to window
	app.MainWindow.SetContent(appContent)
}

func (app *AppConfig) makeHeader() {
	app.StatusLabel = widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
}

func (app *AppConfig) updateHeader(text string) {
	app.StatusLabel.SetText(text)
}
