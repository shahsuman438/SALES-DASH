package main

import (
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/csvfilewatcher"
	"github.com/stretchr/testify/assert"
)

// Mock function to simulate sales and product file processing
func mockProcessSalesFile(fileName string) error {
	// Simulate successful processing of sales file
	return nil
}

func mockProcessProductFile(fileName string) error {
	// Simulate successful processing of product file
	return nil
}

func TestWatchCSVFiles(t *testing.T) {
	// Create a mock watcher and simulate events
	watcher := &fsnotify.Watcher{
		Events: make(chan fsnotify.Event),
		Errors: make(chan error),
	}

	// Use a goroutine to simulate file events
	go func() {
		time.Sleep(100 * time.Millisecond) // Simulate delay for the file watcher

		// Simulate CSV file creation in "sales" directory
		watcher.Events <- fsnotify.Event{
			Name: "sales_test.csv",
			Op:   fsnotify.Create,
		}

		// Simulate CSV file creation in "products" directory
		watcher.Events <- fsnotify.Event{
			Name: "products_test.csv",
			Op:   fsnotify.Create,
		}

		close(watcher.Events)
	}()

	// Call the function under test and inject the mock functions
	err := csvfilewatcher.WatchCSVFiles("sales", mockProcessSalesFile, mockProcessProductFile)
	assert.NoError(t, err)
}
