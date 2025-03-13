package processes

import (
	"accountabully/application/configs"
	"fmt"
	"github.com/gen2brain/beeep"
	"sync"
	"time"
)

var (
	notificationTimers = make(map[string]*time.Timer)
	notificationMu     sync.Mutex
)

var LogoPath = "C:\\Program Files\\Accountabully\\logo.ico"

func (p *Process) sendNotification(actionTitle string) {
	uniqueKey := fmt.Sprintf("%d-%s", p.Pid, actionTitle)
	notificationMu.Lock()
	defer notificationMu.Unlock()

	if _, exists := notificationTimers[uniqueKey]; exists {
		return
	}
	err := beeep.Notify(fmt.Sprintf("Accountabully %s", actionTitle), p.Name, LogoPath)
	if err != nil {
		configs.LogError(err)
	}

	notificationTimers[uniqueKey] = time.AfterFunc(time.Second*time.Duration(configs.NotificationRateLimitSeconds), func() {
		notificationMu.Lock()
		defer notificationMu.Unlock()
		delete(notificationTimers, uniqueKey)
	})
}
