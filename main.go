package main

import (
	"accountabully/application/configs"
	"accountabully/application/ui"
	"github.com/go-vgo/robotgo"
)

func main() {
	pids, err := robotgo.FindIds("accountabully.exe")
	if err != nil {
		panic(err)
	}
	if len(pids) <= 1 {
		// one for itself
		ui.Run()
	} else {
		configs.LogInfo("accountabully.exe is already running")
	}

}
