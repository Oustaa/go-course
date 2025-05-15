package main

import (
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for snapshotTicker != nil || saveAfter != nil {
		select {
		case _, ok := <-snapshotTicker:
			if !ok {
				snapshotTicker = nil
				continue
			}
			takeSnapshot(logChan)
		case _, ok := <-saveAfter:
			if !ok {
				saveAfter = nil
				continue
			}
			saveSnapshot(logChan)
			return
		default:
			waitForData(logChan)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}
