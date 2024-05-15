package entities

import "time"

type RSSData struct {
	Title       string
	Link        string
	Description string
	Author      string
	Email       string
	Created     time.Time
}
