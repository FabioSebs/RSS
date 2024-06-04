package httpclient

import (
	"bytes"
	"encoding/json"
	"log"
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

	req, err := http.NewRequest(http.MethodPost, "http://notifications:6000/v1/email?type=scrape", emailBytes)
	if err != nil {
		log.Println("Oh no request went wrong!")
		log.Println(err.Error())
	}
	// Call the `Do` method, which has a similar interface to the `http.Do` method

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println("Oh no request went wrong!")
		log.Println(err.Error())
	}

	log.Println("Email request sent successfully")
	defer res.Body.Close()

}
