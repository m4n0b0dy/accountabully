package configs

import "runtime"

var PauseMilleSeconds = 100
var NotificationRateLimitSeconds = 30
var BullyRestartTime = 120
var BullyHardLockTime = 600

var OperatingSystem = runtime.GOOS

const (
	ActionMinimize = "minimize"
	ActionClose    = "close"
	ActionNone     = "none"
	ActionWarn     = "warn"
)
