package httpclient

import (
	"fmt"
	"log"

	"github.com/FabioSebs/RSS/config"
	// Import resty into your code and refer it as `resty`.
	"github.com/go-resty/resty/v2"
)

type Email struct {
	Recipients []string `json:"recipients"`
}

func SendRequestForEmail() {
	var (
		request Email = Email{
			Recipients: config.GetAllRecipients(),
		}

		client = resty.New()
		err    error
	)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post("http://notifications:6000/v1/email?type=scrape")

	if resp == nil {
		fmt.Println(err.Error())
	}

	log.Println("Email request sent successfully")
}
