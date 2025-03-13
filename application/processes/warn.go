package processes

import (
	"sync"
	"time"
)

var (
	warnTimers = make(map[int]*time.Timer)
	warnMu     sync.Mutex
)

func (p *Process) Warn() {
	// turning off log because has potential to be noisy
	//message := fmt.Sprintf("Accountabully warned process: %s (PID: %d)", p.Name, p.Pid)
	//configs.LogInfo(message)
	p.sendNotification("Warned")
}
