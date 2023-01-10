package pocketsmith

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func FetchJson(path string, result any) error {
	return doRequest("GET", path, nil, &result)
}

func PutJson(path string, body any, result any) error {
	return doRequest("PUT", path, body, &result)
}

func doRequest(method string, path string, body any, result any) error {
	request, err := createRequest(method, path, body)
	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(responseBody, &result)
}

func createRequest(method string, path string, body any) (*http.Request, error) {
	apiKey := os.Getenv("POCKETSMITH_API_KEY")
	apiUri := os.Getenv("POCKETSMITH_API_URI") + path

	var request *http.Request
	var err error

	if method == "POST" || method == "PUT" {
		requestBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		request, err = http.NewRequest(method, apiUri, bytes.NewBuffer(requestBody))
	} else {
		request, err = http.NewRequest(method, apiUri, nil)
	}

	if err != nil {
		return nil, err
	}

	request.Header.Add("X-Developer-Key", apiKey)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	return request, nil
}
