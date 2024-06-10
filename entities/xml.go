package entities

import (
	"encoding/xml"
	"time"
)

type RSS struct {
	XMLName    xml.Name `xml:"rss"`
	XMLNSMedia string   `xml:"xmlns:media,attr"`
	Version    string   `xml:"version,attr"`
	Channel    Channel  `xml:"channel"`
}

type Channel struct {
	Title          string    `xml:"title"`
	Link           string    `xml:"link"`
	Description    string    `xml:"description"`
	ManagingEditor string    `xml:"managingEditor"`
	PubDate        time.Time `xml:"pubDate"`
	Items          []Item    `xml:"item"`
}

type Item struct {
	Title          string         `xml:"title"`
	Link           string         `xml:"link"`
	Description    string         `xml:"description"`
	PubDate        time.Time      `xml:"pubDate"`
	MediaThumbnail MediaThumbnail `xml:"media:thumbnail,omitempty"`
}

type MediaThumbnail struct {
	URL    string `xml:"url,attr"`
	Width  string `xml:"width,attr,omitempty"`
	Height string `xml:"height,attr,omitempty"`
}
