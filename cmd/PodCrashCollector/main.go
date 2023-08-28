package main

import (
	"log"
	"os"
	"time"

	// other necessary imports
	"github.com/fsnotify/fsnotify"
	pcc "github.com/rnemeth90/PodCrashCollector"
)

var (
	directory = os.Getenv("CDH_UPLOAD_PATH")
	logger    = log.New(os.Stdout, "", log.LstdFlags)
)

func main() {
	logger.Println("[LocalUploadPath]:", directory)

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.Mkdir(directory, 0755)
	}

	// Assuming there's a WatchFiles function available (we'll define it later)
	WatchFiles(directory)

	logger.Println("Press CTRL+C to quit.")
	select {} // Infinite block until program is terminated
}

func WatchFiles(directory string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("File Created:", event.Name)
					fileWatcherOnCreated(event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(directory)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

const fileCheckInterval = 5 * time.Second

func fileWatcherOnCreated(filePath string) {
	log.Println("[Detected]:", filePath)

	// Check if file is in use
	for pcc.IsFileInUse(filePath) {
		log.Printf("[Sleeping]: %s is locked. Sleeping for 5 seconds.", filePath)
		time.Sleep(fileCheckInterval)
	}

	log.Printf("[Uploading]: %s", filePath)
	uploader := &pcc.AzureStorageUploader{}
	// Assuming a PlatformOperations function or interface exists
	platformOps := pcc.GetPlatformOperations()

	if uploader.UploadFile(platformOps, filePath, false) {
		log.Printf("[UploadSuccess]: %s", filePath)
		log.Printf("[DeletingFile]: %s", filePath)
		pcc.DeleteFile(filePath, 5, 1000)
	}
}

