package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/shahsuman438/SALES-DASH/server/pkg/csvfilewatcher"
	"github.com/stretchr/testify/assert"
)

// Mock processing function for sales file
func mockProcessSalesFile(fileName string) error {
	fmt.Println("Processing sales file:", fileName)
	return nil
}

// Mock processing function for product file
func mockProcessProductFile(fileName string) error {
	fmt.Println("Processing product file:", fileName)
	return nil
}

func TestWatchCSVFiles(t *testing.T) {
	// Create temporary directories for testing
	tempDirSales, err := os.MkdirTemp("", "sales")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDirSales) // Clean up after test

	tempDirProducts, err := os.MkdirTemp("", "products")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDirProducts) // Clean up after test

	// Start watching the sales and products directories in separate go routines
	go func() {
		err := csvfilewatcher.WatchCSVFiles(tempDirSales, mockProcessSalesFile, mockProcessProductFile)
		assert.NoError(t, err)
	}()

	go func() {
		err := csvfilewatcher.WatchCSVFiles(tempDirProducts, mockProcessSalesFile, mockProcessProductFile)
		assert.NoError(t, err)
	}()

	// Give the watcher some time to initialize
	time.Sleep(100 * time.Millisecond)

	// Create a new CSV file in the sales directory
	salesFile := filepath.Join(tempDirSales, "sales_test.csv")
	err = os.WriteFile(salesFile, []byte("test,data"), 0644)
	assert.NoError(t, err)

	// Create a new CSV file in the products directory
	productsFile := filepath.Join(tempDirProducts, "products_test.csv")
	err = os.WriteFile(productsFile, []byte("test,data"), 0644)
	assert.NoError(t, err)

	// Allow time for the file watcher to detect the files
	time.Sleep(1 * time.Second)
}
