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

type ThailandScraper struct {
	Collector    *colly.Collector
	Config       config.Config
	RSSGenerator generator.RSSGenerator
	RSS          entities.RSS
	RSSitems     entities.Item
}

func NewThailandScraper() WebScraper {
	env := config.NewConfig()
	return &ThailandScraper{
		Collector:    colly.NewCollector(colly.AllowedDomains(env.PermittedURLs.Thailand...)),
		Config:       env,
		RSSGenerator: generator.NewRssGenerator(),
	}
}

func (t *ThailandScraper) CollectorSetup() *colly.Collector {
	t.RSS = entities.RSS{
		Version: "2.0",
		Channel: entities.Channel{
			Title:          "Thailand News",
			Link:           t.Config.Domains.Thailand,
			Description:    "Reports",
			ManagingEditor: t.Config.ICCTAuthor,
			PubDate:        time.Now(),
			Items:          nil,
		},
	}

	t.Collector.OnHTML("div.container-wrapper div.mag-box-container ul#posts-container", func(element *colly.HTMLElement) {
		element.ForEach("li.post-item", func(i int, h *colly.HTMLElement) {
			var (
				item = entities.Item{
					Title:       h.ChildText("div.post-details h2.post-title a"),
					Link:        h.ChildAttr("a", "href"),
					Description: "News",
					PubDate:     time.Now(),
					Enclosure: entities.Enclosure{
						URL:  h.ChildAttr("a.post-thumb img", "data-src"),
						Type: "image/png",
					},
				}
			)
			t.RSS.Channel.Items = append(t.RSS.Channel.Items, item)
		})
	})

	// Request Feedback
	t.Collector.OnRequest(func(r *colly.Request) {
		// g.Logger.WriteTrace(fmt.Sprintf("visiting url: %s", r.URL.String()))
		fmt.Printf("length of items = %d\n", len(t.RSS.Channel.Items))
		fmt.Println("request launched ...")
	})

	// Error Feedback
	t.Collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("error occured: " + err.Error())
	})

	return t.Collector
}

func (t *ThailandScraper) WriteXML(rss entities.RSS) error {
	var (
		filename string = fmt.Sprintf("xml/%s", t.Config.Filenames.Thailand)
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

func (t *ThailandScraper) LaunchScraper(collector *colly.Collector) {
	for i := 1; i <= 20; i += 1 {
		var (
			url string = t.Config.Domains.Thailand
		)

		if i == 1 {
			fmt.Println("visiting : " + url)
			if err := collector.Visit(url); err != nil {
				fmt.Println("error occured: " + err.Error())
			}
		} else {
			url = fmt.Sprintf("%s/page/%d", url, i)
			fmt.Println("visiting : " + url)
			if err := collector.Visit(url); err != nil {
				fmt.Println("error occured: " + err.Error())
			}
		}

	}

	// Ensuring that the scraping process completes before the program exits
	collector.Wait()

	if err := t.WriteXML(t.RSS); err != nil {
		log.Fatal(err)
	}
}
