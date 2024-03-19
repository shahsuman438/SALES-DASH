package main

import (
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/app"
	csvfilewatcher "github.com/shahsuman438/SALES-DASH/CORE-API/pkg/csvFileWatcher"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/utils/logger"
)

func main() {
	done := make(chan bool)

	go startWatcher("sales", done)
	go startWatcher("products", done)

	app.Start()

	<-done
	<-done
}

func startWatcher(dir string, done chan<- bool) {
	defer func() {
		done <- true
	}()

	if err := csvfilewatcher.WatchCSVFiles(dir); err != nil {
		logger.Error("Unable to watch csv files in", err)
	}
}
