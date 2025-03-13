package ui

import (
	"accountabully/application/configs"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func updateToggleMenuItem(t *fyne.MenuItem, running bool) {
	if running {
		t.Label = "Pause Bully"
		t.Checked = true
		t.Icon = PauseIcon
	} else {
		t.Label = "Start Bully"
		t.Checked = false
		t.Icon = StartIcon
	}
}

func startBully(app *AppConfig, menu *fyne.Menu, toggleMenu *fyne.MenuItem) {
	app.Bully.Start(app.Rules)
	app.updateHeader("Bully has started")
	updateToggleMenuItem(toggleMenu, app.Bully.IsRunning())
	menu.Refresh()
}

func stopBully(app *AppConfig, lockedMenu *fyne.MenuItem, toggleMenu *fyne.MenuItem, showMenu *fyne.MenuItem, quitMenu *fyne.MenuItem, menu *fyne.Menu) {
	app.Bully.Stop()
	app.updateHeader("Bully has stopped")
	updateToggleMenuItem(toggleMenu, app.Bully.IsRunning())
	menu.Refresh()

	if hardCoreMode {
		// Restart bully and lock the app automatically after a set time
		go func() {
			time.Sleep(time.Duration(configs.BullyRestartTime) * time.Second)
			lockApp(app, lockedMenu, toggleMenu, showMenu, quitMenu, menu)
		}()
	}
}

func lockApp(app *AppConfig, lockedMenu *fyne.MenuItem, toggleMenu *fyne.MenuItem, showMenu *fyne.MenuItem, quitMenu *fyne.MenuItem, menu *fyne.Menu) {
	if !app.Bully.IsRunning() {
		startBully(app, menu, toggleMenu)
	}
	app.Locked = true
	lockedMenu.Checked = true
	lockedMenu.Label = "Unlock Bully"
	toggleMenu.Disabled = true
	showMenu.Disabled = true
	quitMenu.Disabled = true
	lockedMenu.Disabled = true
	go func() {
		// disable unlock for a set time
		time.Sleep(time.Duration(configs.BullyHardLockTime) * time.Second)
		lockedMenu.Disabled = false
		menu.Refresh()
	}()
	app.MainWindow.Hide()
	menu.Refresh()
	configs.LogInfo("App has been locked")
}

func unlockApp(app *AppConfig, lockedMenu *fyne.MenuItem, toggleMenu *fyne.MenuItem, showMenu *fyne.MenuItem, quitMenu *fyne.MenuItem, menu *fyne.Menu) {
	app.Locked = false
	lockedMenu.Checked = false
	lockedMenu.Label = "Lock Bully"
	toggleMenu.Disabled = false
	showMenu.Disabled = false
	quitMenu.Disabled = false
	menu.Refresh()
	configs.LogInfo("App has been unlocked")
}

func (app *AppConfig) makeSysTray() {
	if desk, ok := app.App.(desktop.App); ok {
		menu := fyne.NewMenu("Accountabully")
		var toggleMenu *fyne.MenuItem
		var showMenu *fyne.MenuItem
		var quitMenu *fyne.MenuItem
		var lockedMenu *fyne.MenuItem

		toggleMenu = fyne.NewMenuItem("Stop Unknown", func() {
			if app.Bully.IsRunning() {
				stopBully(app, lockedMenu, toggleMenu, showMenu, quitMenu, menu)
			} else {
				startBully(app, menu, toggleMenu)
			}
		})

		// only runs on startup
		updateToggleMenuItem(toggleMenu, app.Bully.IsRunning())
		showMenu = fyne.NewMenuItem("Show Rules", func() {
			app.MainWindow.Show()
			app.MainWindow.RequestFocus()
		})
		quitMenu = fyne.NewMenuItem("Quit", func() {
			app.App.Quit()
		})
		lockedMenu = fyne.NewMenuItem("Lock Bully", func() {
			if !app.Locked {
				lockApp(app, lockedMenu, toggleMenu, showMenu, quitMenu, menu)
			} else {
				unlockApp(app, lockedMenu, toggleMenu, showMenu, quitMenu, menu)
			}
		})
		if hardCoreMode {
			// start default locked
			lockApp(app, lockedMenu, toggleMenu, showMenu, quitMenu, menu)
		}

		menu.Items = append(menu.Items, toggleMenu, showMenu, lockedMenu, quitMenu)
		desk.SetSystemTrayMenu(menu)
	}
}
