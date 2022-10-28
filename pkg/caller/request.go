package caller

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Get[T any](url string, headers map[string]string) (*T, error) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	for key, element := range headers {
		req.Header.Add(key, element)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var response T

	json.NewDecoder(res.Body).Decode(&response)

	return &response, nil
}

func Post[T any](url string, data map[string]string, headers map[string]string) (*T, error) {

	jsonData, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(jsonData)

	req, err := http.NewRequest("POST", url, buf)

	if err != nil {
		return nil, err
	}

	for key, element := range headers {
		req.Header.Add(key, element)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	var response T

	json.NewDecoder(res.Body).Decode(&response)

	return &response, nil
}
