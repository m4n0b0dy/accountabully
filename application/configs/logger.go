package configs

import (
	"fmt"
	"log"
	"time"
)

func LogInfo(message string) {
	timestamp := time.Now().Format(time.RFC3339)
	logMessage := fmt.Sprintf("[%s] INFO: %s", timestamp, message)
	log.Println(logMessage)
}

func LogError(err error) {
	timestamp := time.Now().Format(time.RFC3339)
	logMessage := fmt.Sprintf("[%s] ERROR: %s", timestamp, err)
	log.Println(logMessage)
}
