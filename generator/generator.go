package generator

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/FabioSebs/RSS/config"
	"github.com/FabioSebs/RSS/entities"
	"github.com/gopherlibs/feedhub/feedhub"
)

///////////////////////////////////////////////////////////////////////////
////////////////////////GENERATE XML///////////////////////////////////////
///////////////////////////////////////////////////////////////////////////

type RSSGenerator struct {
	Config config.Config
}

func NewRssGenerator() RSSGenerator {
	return RSSGenerator{
		Config: config.NewConfig(),
	}
}

func (r *RSSGenerator) InitializeMoEFeed() *feedhub.Feed {
	feed := &feedhub.Feed{
		Title:       "Ministry of Energy",
		Link:        &feedhub.Link{Href: r.Config.Sites.MoE[0]},
		Description: "reports of the MoE",
		Author: &feedhub.Author{
			Name:  r.Config.ICCTAuthor,
			Email: r.Config.ICCTEmail,
		},
		Created: time.Now(),
	}
	return feed
}

func (r *RSSGenerator) AddToFeed(rss []entities.RSSData) []*feedhub.Item {
	var (
		items []*feedhub.Item
	)

	for idx := 0; idx < len(rss); idx++ {
		var (
			item feedhub.Item
		)

		item.Title = rss[idx].Title
		item.Link = &feedhub.Link{Href: rss[idx].Link}
		item.Description = rss[idx].Description
		item.Author = &feedhub.Author{Name: rss[idx].Author, Email: rss[idx].Email}
		item.Created = time.Now()

		items = append(items, &item)
	}

	return items
}

func (r *RSSGenerator) ConvertStringToXML(xmlStr string) error {
	// Write the XML string to a file
	err := ioutil.WriteFile("moe.xml", []byte(xmlStr), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return err
	}
	return nil
}
