package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/FabioSebs/RSS/config"
	"github.com/gojek/heimdall/v7/httpclient"
)

type Email struct {
	Recipients []string `json:"recipients"`
}

func SendRequestForEmail() {
	client := httpclient.NewClient()
	// Create an http.Request instance
	// Create the Email instance
	email := Email{
		Recipients: config.GetAllRecipients(), // Adjust as necessary
	}

	// Marshal the Email struct to JSON
	emailJSON, err := json.Marshal(email)
	if err != nil {
		panic(err)
	}

	emailBytes := bytes.NewBuffer(emailJSON)

	req, _ := http.NewRequest(http.MethodPost, "http://notifications:6000/v1/email?type=scrape", emailBytes)
	// Call the `Do` method, which has a similar interface to the `http.Do` method
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		// Handle non-OK responses
		panic("Request failed with status: " + res.Status)
	}
}
