package bullier

import (
	configs2 "accountabully/application/configs"
	"context"
	"time"
)

type Bully struct {
	cancel  context.CancelFunc
	running bool
}

func CreateBully() *Bully {
	b := &Bully{}
	return b
}

func (b *Bully) Start(rules *Rules) {
	if b.IsRunning() {
		configs2.LogInfo("Bully is already running.")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	b.cancel = cancel
	b.running = true

	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Duration(configs2.PauseMilleSeconds) * time.Millisecond)
		defer ticker.Stop()

		configs2.LogInfo("Bully has started.")
		for {
			select {
			case <-ctx.Done():
				configs2.LogInfo("Bully has stopped.")
				b.running = false
				return
			case <-ticker.C:
				allProcesses := rules.CheckAllProcesses()
				allProcesses = append(allProcesses, rules.CheckActiveWindow())

				for _, process := range allProcesses {
					process.DoAction()
				}

			}
		}
	}(ctx)
}

func (b *Bully) Stop() {
	if !b.IsRunning() {
		configs2.LogInfo("Bully is not running.")
		return
	}

	b.cancel()
	b.running = false
}

func (b *Bully) IsRunning() bool {
	return b.running
}
