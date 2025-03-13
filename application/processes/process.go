package processes

import (
	"github.com/go-vgo/robotgo"
)

type Process struct {
	Name       string
	Pid        int
	Action     string
	FromActive bool
}

func CreateFromActiveWindow() *Process {
	return &Process{
		Name:       robotgo.GetTitle(),
		Pid:        robotgo.GetPid(),
		FromActive: true,
	}
}
