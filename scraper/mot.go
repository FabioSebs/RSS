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
	"github.com/FabioSebs/RSS/utils"
	"github.com/gocolly/colly"
)

type MoTScraper struct {
	Collector    *colly.Collector
	Config       config.Config
	RSSGenerator generator.RSSGenerator
	RSS          entities.RSS
	RSSitems     entities.Item
}

func NewMoTScraper() WebScraper {
	var (
		env config.Config = config.NewConfig()
	)

	return &MoTScraper{
		Collector: colly.NewCollector(colly.AllowedDomains(
			env.PermittedURLs.MoT...,
		)),
		Config:       env,
		RSSGenerator: generator.NewRssGenerator(),
	}
}

func (g *MoTScraper) CollectorSetup() *colly.Collector {

	g.RSS = entities.RSS{
		Version: "2.0",
		Channel: entities.Channel{
			Title:          "Ministry of Transportation",
			Link:           g.Config.Domains.MoT,
			Description:    "Reports",
			ManagingEditor: g.Config.ICCTAuthor,
			PubDate:        time.Now(),
			Items:          nil, // needs scraping
		},
	}

	g.Collector.OnHTML("div.widget ul.comments-list", func(element *colly.HTMLElement) {

		element.ForEach("li", func(i int, h *colly.HTMLElement) {
			var (
				item entities.Item = entities.Item{
					Title:       h.ChildText("h3"),
					Link:        h.ChildAttr("h3 a", "href"),
					Description: h.ChildText("small"),
					PubDate:     time.Now(),
					Enclosure: entities.Enclosure{
						URL:  h.ChildAttr("div.alignleft a img", "src"),
						Type: "image/jpg",
					},
				}
			)
			if utils.ValidateTitle(item.Title) {
				g.RSS.Channel.Items = append(g.RSS.Channel.Items, item)
			}
		})
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

func (g *MoTScraper) WriteXML(rss entities.RSS) error {
	var (
		filename string = fmt.Sprintf("xml/%s", g.Config.Filenames.MoT)
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

func (g *MoTScraper) LaunchScraper(collector *colly.Collector) {
	if err := collector.Visit(g.Config.Domains.MoT); err != nil {
		fmt.Println("error occured: " + err.Error())
	}
	// Ensuring that the scraping process completes before the program exits
	collector.Wait()

	if len(g.RSS.Channel.Items) > 0 {
		if err := g.WriteXML(g.RSS); err != nil {
			log.Fatal(err)
		}
	}
}
