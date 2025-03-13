package processes

import (
	"accountabully/application/configs"
)

func (p *Process) DoAction() {
	switch p.Action {
	case configs.ActionMinimize:
		p.Minimize()
	case configs.ActionClose:
		p.Close()
	case configs.ActionWarn:
		p.Warn()
	case configs.ActionNone:
		return
	}
}
