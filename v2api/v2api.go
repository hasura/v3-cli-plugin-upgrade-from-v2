package v2api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchV2Metadata(v3URL string, adminSecret string) map[string]interface{} {
	requestBody := `{"type":"export_metadata","version":2,"args":{}}`
	result := FetchV2Common(v3URL, requestBody, adminSecret)
	return result
}

func FetchV2InternalState(v3URL string, adminSecret string) map[string]interface{} {
	requestBody := `{"type":"dump_internal_state","version":1,"args":{}}`
	result := FetchV2Common(v3URL, requestBody, adminSecret)
	return result
}

// Common fetching logic - returns a generic JSON payload
func FetchV2Common(v3URL string, payload string, adminSecret string) map[string]interface{} {

	// Get the metadata API URL
	apiURL := fmt.Sprintf("%s/v1/metadata", v3URL)

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		errmsg := fmt.Sprintf("Error creating request: %s", err)
		panic(errmsg)
	}

	// Add the authorization header and content-type to the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", adminSecret)

	// Make the request
	response, err := client.Do(req)
	if err != nil {
		errmsg := fmt.Sprintf("Error making the request: %s", err)
		panic(errmsg)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errmsg := fmt.Sprintf("Error: Non-200 status code received - %s\n", response.Status)
		panic(errmsg)
	}

	// Read the response body
	body, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		errmsg := fmt.Sprintf("Error reading the response body: %s", err)
		panic(errmsg)
	}

	// Marshall the response into a JSON map[string]interface{} structure
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		errmsg := fmt.Sprintf("Response was not valid JSON: %s", err)
		panic(errmsg)
	}

	return result
}
