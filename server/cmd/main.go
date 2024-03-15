package main

import (
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/app"
	csvfilewatcher "github.com/shahsuman438/SALES-DASH/CORE-API/pkg/csvFileWatcher"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/utils/logger"
)

func main() {
	done := make(chan bool)

	// Run the CSV watcher for the first folder in a goroutine
	go func() {
		dir := "sales"
		err := csvfilewatcher.WatchCSVFiles(dir)
		if err != nil {
			logger.Error("unable to watch csv file in sales folder", err)
		}
	}()

	// Run the CSV watcher for the second folder in a goroutine
	go func() {
		dir := "products"
		err := csvfilewatcher.WatchCSVFiles(dir)
		if err != nil {
			logger.Error("unable to watch csv file in products folder", err)
		}
	}()

	// Run the API in the main goroutine
	app.Start()

	// Notify that all watcher goroutines have completed
	done <- true
	done <- true

	// Wait for both watcher goroutines to complete before exiting
	<-done
	<-done
}
