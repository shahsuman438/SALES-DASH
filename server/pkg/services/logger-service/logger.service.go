package loggerservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoggerPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func WriteLog(data *LoggerPayload) error {
	logServiceURL := "http://localhost:9999/log"
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unable to conect with logger services")
	}

	return nil
}
