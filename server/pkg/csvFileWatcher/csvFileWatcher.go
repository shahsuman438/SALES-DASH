package csvfilewatcher

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	loggerservice "github.com/shahsuman438/SALES-DASH/server/pkg/services/logger-service"
	"github.com/shahsuman438/SALES-DASH/server/pkg/utils/logger"
)

func WatchCSVFiles(dir string, processSalesFile func(string) error, processProductFile func(string) error) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	if err := addDirectoryToWatcher(watcher, dir); err != nil {
		return err
	}

	processEvents(watcher, dir, processSalesFile, processProductFile)

	return nil
}

func addDirectoryToWatcher(watcher *fsnotify.Watcher, dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
}

func processEvents(watcher *fsnotify.Watcher, dir string, processSalesFile func(string) error, processProductFile func(string) error) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				if filepath.Ext(event.Name) == ".csv" {
					logger.Info(fmt.Sprintf("New file detected: %s", event.Name))
					loggerservice.WriteLog(&loggerservice.LoggerPayload{Name: "SYSTEM", Data: "new file detected."})
					if err := processFile(dir, event.Name, processSalesFile, processProductFile); err != nil {
						logger.Error("Error processing file error:", err)
					}
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			logger.Error("Error occurred error:", err)
		}
	}
}

func processFile(dir, fileName string, processSalesFile func(string) error, processProductFile func(string) error) error {
	if dir == "data/products" {
		loggerservice.WriteLog(&loggerservice.LoggerPayload{Name: "SYSTEM", Data: fmt.Sprintf("file:%s, start processing for product", fileName)})
		return processProductFile(fileName)
	}
	loggerservice.WriteLog(&loggerservice.LoggerPayload{Name: "SYSTEM", Data: fmt.Sprintf("file:%s, start processing for sales", fileName)})
	return processSalesFile(fileName)
}
