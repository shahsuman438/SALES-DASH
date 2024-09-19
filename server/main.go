package main

import (
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/app"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/csvfilewatcher"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/file"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/utils/logger"
)

func main() {
	// Create a channel to signal when both watchers are done
	done := make(chan bool)

	// Start file watchers for the "sales" and "products" directories
	go startWatcher("sales", file.ProcessSalesFile, file.ProcessProductFiles, done)
	go startWatcher("products", file.ProcessSalesFile, file.ProcessProductFiles, done)

	// Start the application
	app.Start()

	// Wait for both watchers to complete (just an example, in production you might want continuous watching)
	<-done
	<-done
}

func startWatcher(dir string, processSalesFile func(string) error, processProductFile func(string) error, done chan<- bool) {
	defer func() {
		done <- true
	}()

	// Call WatchCSVFiles with appropriate file processing functions
	if err := csvfilewatcher.WatchCSVFiles(dir, processSalesFile, processProductFile); err != nil {
		logger.Error("Unable to watch csv files in", err)
	}
}
