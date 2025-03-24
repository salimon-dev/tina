package openai

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getOpenAiToken() string {
	return os.Getenv("OPEN_AI_KEY")
}

func SendRequest(method string, path string, body []byte) ([]byte, error) {
	url := "https://api.openai.com" + path
	accessToken := getOpenAiToken()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
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
