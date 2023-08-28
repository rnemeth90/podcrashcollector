package podcrashcollector

import (
	"os"
	"time"
)

func IsFileInUse(path string) bool {
	_, err := os.OpenFile(path, os.O_RDWR|os.O_EXCL, 0666)
	if err != nil {
		if os.IsPermission(err) {
			return true
		}
	}
	return false
}

func DeleteFile(path string, maxRetries int, delayMilliseconds int) bool {
	for i := 0; i < maxRetries; i++ {
		err := os.Remove(path)
		if err == nil {
			return true
		}
		time.Sleep(time.Duration(delayMilliseconds) * time.Millisecond)
	}
	return false
}

