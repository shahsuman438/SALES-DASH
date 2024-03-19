package csvfilewatcher

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/file"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/utils/logger"
)

func WatchCSVFiles(dir string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	if err := addDirectoryToWatcher(watcher, dir); err != nil {
		return err
	}

	processEvents(watcher, dir)

	return nil
}

func addDirectoryToWatcher(watcher *fsnotify.Watcher, dir string) error {
	return filepath.Walk("data/"+dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
}

func processEvents(watcher *fsnotify.Watcher, dir string) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				if filepath.Ext(event.Name) == ".csv" {
					logger.Info(fmt.Sprintf("New file detected: %s", event.Name))
					if err := processFile(dir, event.Name); err != nil {
						logger.Error("Error processing file in", err)
					}
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			logger.Error("Error occurred in", err)
		}
	}
}

func processFile(dir, fileName string) error {
	processFunc := file.ProcessSalesFile
	if dir == "products" {
		processFunc = file.ProcessProductFiles
	}

	return processFunc(fileName)
}
