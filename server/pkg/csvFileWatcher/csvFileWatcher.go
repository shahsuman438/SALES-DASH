package csvfilewatcher

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/file"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/utils/logger"
)

// WatchCSVFiles watches for new CSV files in the specified directory
func WatchCSVFiles(dir string) error {
	// Create a new watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	// Add the directory to the watcher
	err = filepath.Walk("data/"+dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Process events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					// A new file was created
					if filepath.Ext(event.Name) == ".csv" {
						logger.Info(fmt.Sprintf("New file detected: %s", event.Name))
						// Process file here
						if dir == "products" {
							err := file.ProcessProductFiles(event.Name)
							if err != nil {
								logger.Error("Error reading File", err)
								continue
							}
						}
						if dir == "sales" {
							err := file.ProcessSalesFile(event.Name)
							if err != nil {
								logger.Error("Error reading File", err)
								continue
							}
						}
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Error:", err)
			}
		}
	}()

	// Wait for program termination
	<-make(chan struct{})

	return nil
}
