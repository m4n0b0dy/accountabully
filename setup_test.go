package main

import (
	"accountabully/application/repository"
	"accountabully/application/ui"
	"fyne.io/fyne/v2/test"
	"os"
	"testing"
)

var testApp ui.AppConfig

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.MainWindow = a.NewWindow("test")
	testApp.DB = repository.NewTestRepository()
	testApp.Bully = nil
	os.Exit(m.Run())
}

// todo
