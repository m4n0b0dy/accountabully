package bullier

import (
	"accountabully/application/processes"
	"github.com/go-vgo/robotgo"
	"regexp"
	"strings"
)

type Rule struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Action string `json:"action"`
	Active bool   `json:"active"`
}

type Rules []Rule

func (r *Rules) CheckAllProcesses() []*processes.Process {
	var blockedProcesses []*processes.Process
	for _, rule := range *r {
		if !rule.Active {
			continue
		}

		ids, err := robotgo.FindIds(rule.Name)
		if err != nil {
			continue
		}
		for _, id := range ids {
			name, err := robotgo.FindName(id)
			if err != nil {
				continue
			}
			blockedProcesses = append(blockedProcesses, &processes.Process{
				Name:       name,
				Pid:        id,
				Action:     rule.Action,
				FromActive: false,
			})
		}
	}
	return blockedProcesses
}

func (r *Rules) CheckActiveWindow() *processes.Process {
	activeWindowProcess := processes.CreateFromActiveWindow()
	lCase := strings.ToLower(activeWindowProcess.Name)

	for _, rule := range *r {
		if !rule.Active {
			continue
		}
		rLCase := strings.ToLower(rule.Name)

		if strings.Contains(lCase, rLCase) {
			activeWindowProcess.Action = rule.Action
			return activeWindowProcess
		}

		pattern := "(?i)" + rLCase
		matched, err := regexp.MatchString(pattern, lCase)
		if err != nil {
			continue
		}
		if matched {
			activeWindowProcess.Action = rule.Action
			return activeWindowProcess
		}
	}
	return activeWindowProcess
}
