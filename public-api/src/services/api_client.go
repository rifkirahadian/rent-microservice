package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func ApiGet(url string, queryParams ...map[string]string) (interface{}, int, error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}

	// Add query parameters to the request if provided
	if len(queryParams) > 0 {
		query := req.URL.Query()
		for key, value := range queryParams[0] {
			query.Add(key, value)
		}
		req.URL.RawQuery = query.Encode()
	}

	// Send the request
	response, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer response.Body.Close()

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}

	// Parse response body as JSON
	var jsonResponse interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return nil, 0, err
	}

	return jsonResponse, response.StatusCode, nil
}

// postDataToAPI makes a POST request to an external API with the specified data and returns the response body and status code.
func ApiPost(url string, data interface{}) (interface{}, int, error) {
	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, 0, err
	}

	// Make POST request to external API
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, 0, err
	}
	defer response.Body.Close()

	// Read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}

	var jsonResponse interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return nil, 0, err
	}

	return jsonResponse, response.StatusCode, nil
}
