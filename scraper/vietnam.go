package scraper

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/FabioSebs/RSS/config"
	"github.com/FabioSebs/RSS/entities"
	"github.com/FabioSebs/RSS/generator"
	"github.com/gocolly/colly"
)

type VietnamScraper struct {
	Collector    *colly.Collector
	Config       config.Config
	RSSGenerator generator.RSSGenerator
	RSS          entities.RSS
	RSSitems     entities.Item
}

func NewVietnamScraper() WebScraper {
	env := config.NewConfig()
	return &VietnamScraper{
		Collector:    colly.NewCollector(colly.AllowedDomains(env.PermittedURLs.Vietnam...)),
		Config:       env,
		RSSGenerator: generator.NewRssGenerator(),
	}
}

func (v *VietnamScraper) CollectorSetup() *colly.Collector {
	v.RSS = entities.RSS{
		Version:    "2.0",
		XMLNSMedia: "http://search.yahoo.com/mrss/",
		Channel: entities.Channel{
			Title:          "Vietnam News",
			Link:           v.Config.Domains.Antara,
			Description:    "Reports",
			ManagingEditor: v.Config.ICCTAuthor,
			PubDate:        time.Now(),
			Items:          nil,
		},
	}

	v.Collector.OnHTML("div.container div.list__listing div.list__flex div.list__main div.list__lflex div.list__lmain div.box-stream", func(element *colly.HTMLElement) {
		element.ForEach("div.box-stream-item", func(i int, h *colly.HTMLElement) {
			var (
				item = entities.Item{
					Title:       h.ChildText("div.box-stream-content h2 a"),
					Link:        v.Config.PermittedURLs.Vietnam[0] + h.ChildAttr("div.box-stream-content h2 a", "href"),
					Description: h.ChildText("div.box-stream-content p"),
					PubDate:     time.Now(),
					MediaThumbnail: entities.MediaThumbnail{
						URL: h.ChildAttr("a.box-stream-link-with-avatar img", "src"),
					},
				}
			)

			v.RSS.Channel.Items = append(v.RSS.Channel.Items, item)
		})
	})

	// Request Feedback
	v.Collector.OnRequest(func(r *colly.Request) {
		// g.Logger.WriteTrace(fmt.Sprintf("visiting url: %s", r.URL.String()))
		fmt.Printf("length of items = %d\n", len(v.RSS.Channel.Items))
		fmt.Println("request launched ...")
	})

	// Error Feedback
	v.Collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("error occured: " + err.Error())
	})

	return v.Collector
}

func (v *VietnamScraper) WriteXML(rss entities.RSS) error {
	var (
		filename string = fmt.Sprintf("xml/%s", v.Config.Filenames.Vietnam)
	)

	output, err := xml.MarshalIndent(rss, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling XML:", err)
		return err
	}

	// Add XML header
	output = []byte(xml.Header + string(output))

	// Write XML to file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(output)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Println("RSS feed successfully written")
	return nil
}

func (v *VietnamScraper) LaunchScraper(collector *colly.Collector) {
	if err := v.Collector.Visit(v.Config.Domains.Vietnam); err != nil {
		fmt.Println("error occured: " + err.Error())
	}
	// Ensuring that the scraping process completes before the program exits
	collector.Wait()

	if err := v.WriteXML(v.RSS); err != nil {
		log.Fatal(err)
	}
}
