package main

import (
	"fmt"
	"time"
)
const (
	logInfo = "INFO"
	logWarning= "WARNING"
	logError = "ERROR"
)

type logEntry struct {
	time time.Time
	severity string
	message string
}
var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func main() {
		go logger()

		logCh <- logEntry{time.Now(), logInfo, "App is starting"}
		logCh <- logEntry{time.Now(), logInfo, "App is running"}
		logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
		time.Sleep(100 * time.Millisecond)
		// an anonimous struc with empty fields
		doneCh <- struct{}{}
	}

func logger() {
	for {
		select {
		case entry := <- logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-12-31"), entry.severity, entry.message)
		case <- doneCh:
			break
		}
	}
}