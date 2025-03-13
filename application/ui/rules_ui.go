package ui

import (
	"accountabully/application/bullier"
	"accountabully/application/configs"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strings"
)

var NextActionMap = map[string]string{
	"minimize": "close",
	"close":    "warn",
	"warn":     "minimize",
}

func createActionLabelText(rule bullier.Rule) string {
	if rule.Active {
		return fmt.Sprintf("Bully will %s %s", rule.Action, rule.Name)
	} else {
		return fmt.Sprintf("Bully will not %s %s (inactive)", rule.Action, rule.Name)
	}
}

func (app *AppConfig) refetchRules() {
	// longer term, rules should be in charge of the db connection and keeping itself up to date
	freshRules, err := app.DB.GetAllRules()
	if err != nil {
		configs.LogError(err)
		app.updateHeader("Error refetching rules")
	}
	*(app.Rules) = freshRules
}

func updateActivateActionIcon(activateAction *widget.ToolbarAction, active bool) {
	if active {
		activateAction.SetIcon(DisableIcon)
	} else {
		activateAction.SetIcon(ReEnableIcon)
	}
}

func (app *AppConfig) ruleToolBar(i int) *fyne.Container {
	rule := (*app.Rules)[i]

	ruleLabel := widget.NewLabel(createActionLabelText(rule))
	var activateAction *widget.ToolbarAction
	activateAction = widget.NewToolbarAction(nil, func() {
		updateRule := bullier.Rule{
			ID:     rule.ID,
			Name:   rule.Name,
			Action: rule.Action,
			Active: !rule.Active,
		}
		err := app.DB.UpdateRule(updateRule)
		if err != nil {
			configs.LogError(err)
			app.updateHeader("Error updating rule")
		}
		app.refetchRules()
		updateActivateActionIcon(activateAction, updateRule.Active)

		ruleLabel.SetText(createActionLabelText(rule))
		end := "active"
		if !updateRule.Active {
			end = "inactive"
		}
		msg := fmt.Sprintf("Rule %s set to %s", rule.Name, end)
		configs.LogInfo(msg)
		app.updateHeader(msg)
	})
	updateActivateActionIcon(activateAction, rule.Active) // set default

	cycleAction := widget.NewToolbarAction(SwitchIcon, func() {
		updateRule := bullier.Rule{
			ID:     rule.ID,
			Name:   rule.Name,
			Action: NextActionMap[rule.Action],
			Active: rule.Active,
		}
		err := app.DB.UpdateRule(updateRule)
		if err != nil {
			configs.LogError(err)
			app.updateHeader("Error updating rule")
		}
		app.refetchRules()

		ruleLabel.SetText(createActionLabelText(rule))
		msg := fmt.Sprintf("Rule %s set to %s", rule.Name, updateRule.Action)
		configs.LogInfo(msg)
		app.updateHeader(msg)
	})
	deleteAction := widget.NewToolbarAction(DeleteIcon, func() {
		err := app.DB.DeleteRule(rule.ID)
		if err != nil {
			configs.LogError(err)
			app.updateHeader("Error deleting rule")
		}
		app.refetchRules()
		msg := fmt.Sprintf("Rule %s deleted", rule.Name)
		configs.LogInfo(msg)
		app.updateHeader(msg)
	})

	return container.NewHBox(widget.NewToolbar(activateAction, cycleAction, deleteAction), ruleLabel)

}

func (app *AppConfig) makeRulesList() *widget.List {
	return widget.NewList(
		func() int {
			return len(*(app.Rules))
		},
		func() fyne.CanvasObject {
			label := widget.NewLabel("")
			toolbar := widget.NewToolbar()
			return container.NewHBox(label, toolbar)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[1] = app.ruleToolBar(i)
		},
	)
}

func (app *AppConfig) makeAddRule() *fyne.Container {
	nameEntry := widget.NewEntry()
	nameEntry.PlaceHolder = "Type the name of the program to bully"
	addButton := widget.NewButton("Add Rule", func() {
		text := strings.ToLower(nameEntry.Text)
		if text == "" {
			msg := "Rule name cannot be empty"
			configs.LogInfo(msg)
			app.updateHeader(msg)
			return
		}
		// check name against current rules
		for _, rule := range *(app.Rules) {
			if rule.Name == text {
				msg := fmt.Sprintf("Rule %s already exists", text)
				configs.LogInfo(msg)
				app.updateHeader(msg)

				return
			}
		}

		newRule := bullier.Rule{
			Name:   text,
			Action: "warn",
			Active: true,
		}
		err := app.DB.InsertRule(newRule)
		if err != nil {
			configs.LogError(err)
			app.updateHeader("Error inserting rule")
			return
		}
		app.refetchRules()
		msg := fmt.Sprintf("Rule %s added", newRule.Name)
		configs.LogInfo(msg)
		app.updateHeader(msg)
	})
	return container.NewBorder(nil, nil, nil, addButton, nameEntry)
}
