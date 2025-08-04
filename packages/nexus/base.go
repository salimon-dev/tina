package nexus

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getNexusBaseURL() string {
	return os.Getenv("NEXUS_BASE_URL")
}

func GetAccessToken() string {
	return os.Getenv("NEXUS_ACCESS_TOKEN")
}

func GetUsername() string {
	return os.Getenv("NEXUS_USERNAME")
}

func SendHttpRequest(method string, path string, body []byte) ([]byte, error) {
	url := fmt.Sprintf("%s%s", getNexusBaseURL(), path)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "API-KEY "+GetAccessToken())
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
		fmt.Println(string(responseBody))
		return nil, fmt.Errorf("Request failed with status code %d", res.StatusCode)
	}
	return responseBody, nil

}
