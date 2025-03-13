package processes

import (
	configs2 "accountabully/application/configs"
	"fmt"
	"os/exec"
	"syscall"

	"github.com/go-vgo/robotgo"
)

func windowsCloseByPid(pid int) error {
	cmd := exec.Command("taskkill", "/F", "/FI", fmt.Sprintf("PID eq %d", pid))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to close process with PID %d: %v", pid, err)
	}
	return nil
}

func (p *Process) Close() {
	defer func() error {
		if r := recover(); r != nil {
			if configs2.OperatingSystem == "windows" {
				return windowsCloseByPid(p.Pid)
			}
		}
		return nil
	}()
	configs2.LogInfo(fmt.Sprintf("Accountabully closed process: %s (PID: %d)", p.Name, p.Pid))
	p.sendNotification("Closed")
	err := robotgo.Kill(p.Pid)
	if err != nil {
		configs2.LogError(err)
	}
}
