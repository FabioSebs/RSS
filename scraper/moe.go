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

type MoeScraper struct {
	Collector    *colly.Collector
	Config       config.Config
	RSSGenerator generator.RSSGenerator
	// Logger    logger.Logger
}

func NewMoeScraper() WebScraper {
	env := config.NewConfig()
	return &MoeScraper{
		Collector: colly.NewCollector(colly.AllowedDomains(
			env.PermittedURLs.MoE...,
		//append more sites here
		)),
		Config:       env,
		RSSGenerator: generator.NewRssGenerator(),
		// Logger: logger.NewLogger(),
	}
}

func (g *MoeScraper) CollectorSetup() *colly.Collector {
	g.Collector.OnHTML("div.container", func(element *colly.HTMLElement) {
		var (
			pubTime time.Time = time.Now()

			rss entities.RSS = entities.RSS{
				Version: "2.0",
				Channel: entities.Channel{
					Title:          "Ministry of Energy",
					Link:           g.Config.Domains.MoE,
					Description:    "Reports",
					ManagingEditor: g.Config.ICCTAuthor,
					PubDate:        pubTime,
					Items:          nil, // needs scraping
				},
			}
		)

		// scraper logic goes here
		element.ForEach("div.list-berita div.col-md-4 div.berita-item", func(i int, h *colly.HTMLElement) {
			var (
				item entities.Item = entities.Item{
					Title:       h.ChildText("h4.title"),
					Link:        h.ChildAttr("a.btn-download", "href"),
					Description: h.ChildText("p.post-time"),
					PubDate:     time.Now(),
				}
			)

			rss.Channel.Items = append(rss.Channel.Items, item)
		})

		if len(rss.Channel.Items) > 0 {
			if err := g.WriteXML(rss); err != nil {
				log.Fatal(err)
			}
		}
	})

	// Request Feedback
	g.Collector.OnRequest(func(r *colly.Request) {
		// g.Logger.WriteTrace(fmt.Sprintf("visiting url: %s", r.URL.String()))
		fmt.Println("visiting... " + r.URL.Host)
	})

	// Error Feedback
	g.Collector.OnError(func(_ *colly.Response, err error) {
		// g.Logger.WriteError(fmt.Sprintf("error: %s", err.Error()))
		fmt.Println("error occured: " + err.Error())
	})
	return g.Collector
}
func (g *MoeScraper) WriteXML(rss entities.RSS) error {
	output, err := xml.MarshalIndent(rss, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling XML:", err)
		return err
	}

	// Add XML header
	output = []byte(xml.Header + string(output))

	// Write XML to file
	file, err := os.Create(g.Config.Filenames.MoE)
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

func (g *MoeScraper) LaunchScraper(collector *colly.Collector) {
	if err := collector.Visit(g.Config.Domains.MoE); err != nil {
		fmt.Println("error occured: " + err.Error())
	}
	// Ensuring that the scraping process completes before the program exits
	collector.Wait()
}