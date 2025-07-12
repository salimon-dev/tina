package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func SendNexusRequest(path string, method string, body []byte) ([]byte, error) {
	url := os.Getenv("NEXUS_BASE_URL") + path

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	apiKey := os.Getenv("API_KEY")
	req.Header.Set("Authorization", "API-KEY "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return responseBody, errors.New(fmt.Sprintf("Request failed with status code %d", res.StatusCode))
	}

	return responseBody, nil
}
