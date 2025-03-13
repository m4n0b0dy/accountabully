package processes

import (
	configs2 "accountabully/application/configs"
	"fmt"
	"github.com/go-vgo/robotgo"
)

func winBackupMinimize() {
	_ = robotgo.KeyDown("cmd")
	_ = robotgo.KeyTap("down")
	_ = robotgo.KeyTap("down")
	_ = robotgo.KeyUp("cmd")
}
func (p *Process) Minimize() {
	robotgo.MinWindow(p.Pid)
	// flaky behavior, so we need to run again on active window (only works on active window) and check if still up
	if p.FromActive {
		robotgo.MilliSleep(100)
		activeWindow := CreateFromActiveWindow()
		if activeWindow.Name == p.Name {
			if configs2.OperatingSystem == "windows" {
				winBackupMinimize()
			}
		}
	}
	configs2.LogInfo(fmt.Sprintf("Accountabully minimized process: %s (PID: %d)", p.Name, p.Pid))
	p.sendNotification("Minimized")
}
