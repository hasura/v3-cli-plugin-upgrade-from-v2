package v2api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchV2Info(v3URL string, adminSecret string) string {

	// Get the metadata API URL
	apiURL := fmt.Sprintf("%s/v1/metadata", v3URL)
	fmt.Println("API URL:", apiURL)

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	requestBody := `{"type":"export_metadata","version":2,"args":{}}`
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		errmsg := fmt.Sprintf("Error creating request: %s", err)
		panic(errmsg)
	}

	// Add the authorization header to the request
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
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errmsg := fmt.Sprintf("Error reading the response body: %s", err)
		panic(errmsg)
	}

	// Process the response as needed
	fmt.Println(string(body))
	fmt.Println("Status code:", response.StatusCode)

	return "So much info TODO."
}
