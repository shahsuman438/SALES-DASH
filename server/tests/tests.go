package tests

import (
	"fmt"
	// "net/http/httptest"
)

var (
	// server *httptest.Server
	green = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	red   = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue  = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	reset = string([]byte{27, 91, 48, 109})
)

// func setup() {
// 	app.Start()
// }

// func tearDown() {
// 	server.Close()
// }

func testPassed(message string) {
	fmt.Printf("%s%s %s\n", green, message, reset)
}

func testFailed(err error) {
	fmt.Printf("%s%s %s\n", red, err.Error(), reset)
}

func logTestInformation(message string) {
	fmt.Printf("%s%s %s\n", blue, message, reset)
}
