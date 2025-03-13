package ui

import (
	bullier2 "accountabully/application/bullier"
	"accountabully/application/configs"
	repository2 "accountabully/application/repository"
	"database/sql"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type AppConfig struct {
	App         fyne.App
	MainWindow  fyne.Window
	Bully       *bullier2.Bully
	DB          repository2.Repository
	Rules       *bullier2.Rules
	StatusLabel *widget.Label
	Locked      bool
}

var myApp AppConfig

// hardCoreMode, turns bully on at startup, locks app on startup, and auto restarts bully after a set time
var hardCoreMode = true

func Run() {
	// create app
	fyneApp := app.NewWithID("com.github.accountabully")
	myApp.App = fyneApp

	fyne.CurrentApp().Settings().SetTheme(&CustomTheme{})

	// connect and migrate the database
	myApp.setupDB()

	// set the current rules
	myApp.setRules()

	// create the bully object
	myApp.Bully = bullier2.CreateBully()
	if hardCoreMode {
		// start bully automatically
		myApp.Bully.Start(myApp.Rules)
	}

	// create icon
	myApp.App.SetIcon(LogoIcon)

	// create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("Accountabully")
	myApp.MainWindow.Resize(fyne.NewSize(500, 400))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()
	myApp.MainWindow.SetCloseIntercept(func() {
		myApp.MainWindow.Hide()
	})
	myApp.MainWindow.SetTitle("Accountabully Rules Manager")

	// create the UI
	myApp.makeUI()

	// create the system tray
	myApp.makeSysTray()

	// run the app (starts only system tray)
	myApp.App.Run()
}

func (app *AppConfig) setupDB() {
	path := ""
	if os.Getenv("ACCOUNTABULLY_DB_PATH") != "" {
		path = os.Getenv("ACCOUNTABULLY_DB_PATH")
	} else {
		path = app.App.Storage().RootURI().Path() + "/accountabully.db"
	}
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		configs.LogError(err)
		log.Panic(err)
	}
	app.DB = repository2.NewSQLiteRepository(db)
	err = app.DB.Migrate(true)
	if err != nil {
		configs.LogError(err)
		log.Panic(err)
	}
}

func (app *AppConfig) setRules() {
	rules, err := app.DB.GetAllRules()
	if err != nil {
		configs.LogError(err)
		log.Panic(err)
	}
	app.Rules = &rules
}
