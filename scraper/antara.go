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

type AntaraScraper struct {
	Collector    *colly.Collector
	Config       config.Config
	RSSGenerator generator.RSSGenerator
	RSS          entities.RSS
	RSSitems     entities.Item
}

func NewAntaraScraper() WebScraper {
	env := config.NewConfig()
	return &AntaraScraper{
		Collector:    colly.NewCollector(colly.AllowedDomains(env.PermittedURLs.Antara...)),
		Config:       env,
		RSSGenerator: generator.NewRssGenerator(),
	}
}

func (a *AntaraScraper) CollectorSetup() *colly.Collector {
	// init rss
	a.RSS = entities.RSS{
		Version:    "2.0",
		XMLNSMedia: "http://search.yahoo.com/mrss/",
		Channel: entities.Channel{
			Title:          "Antara News Jakarta",
			Link:           a.Config.Domains.Antara,
			Description:    "Reports",
			ManagingEditor: a.Config.ICCTAuthor,
			PubDate:        time.Now(),
			Items:          nil,
		},
	}

	a.Collector.OnHTML("div.container div.row div.col-md-8 div.wrapper__list__article", func(element *colly.HTMLElement) {
		element.ForEach("div.card__post div.row", func(i int, h *colly.HTMLElement) {
			item := entities.Item{
				Title:       h.ChildText("div.col-md-7 div.card__post__body div.card__post__content div.card__post__title h2.post_title a"),
				Link:        h.ChildAttr("div.col-md-7 div.card__post__body div.card__post__content div.card__post__title h2.post_title a", "href"),
				Description: h.ChildText("div.col-md-7 div.card__post__body div.card__post__content p"),
				PubDate:     time.Now(),
				MediaThumbnail: entities.MediaThumbnail{
					URL: "https://cdn.antaranews.com/cache/1200x800/2019/01/Logo-LKBN-Antara_fea.png",
				}}

			a.RSS.Channel.Items = append(a.RSS.Channel.Items, item)
		})
	})

	// Request Feedback
	a.Collector.OnRequest(func(r *colly.Request) {
		// g.Logger.WriteTrace(fmt.Sprintf("visiting url: %s", r.URL.String()))
		fmt.Printf("Length of Items : %d", len(a.RSS.Channel.Items))
		fmt.Println("request sent")
	})

	// Error Feedback
	a.Collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("error occured: " + err.Error())
	})

	return a.Collector
}

func (a *AntaraScraper) WriteXML(rss entities.RSS) error {
	var (
		filename string = fmt.Sprintf("xml/%s", a.Config.Filenames.Antara)
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

func (a *AntaraScraper) LaunchScraper(collector *colly.Collector) {
	for i := 1; i <= 15; i++ {
		var (
			url string = a.Config.Domains.Antara
		)

		if i == 1 {
			fmt.Println("visiting.. " + url)
			if err := collector.Visit(url); err != nil {
				fmt.Println("error occured: " + err.Error())
				break
			}
		} else {
			url = fmt.Sprintf("%s/%d", url, i)
			fmt.Println("visiting.. " + url)
			if err := collector.Visit(url); err != nil {
				fmt.Println("error occured: " + err.Error())
				break
			}
		}

		time.Sleep(time.Second)
	}

	// Ensuring that the scraping process completes before the program exits
	collector.Wait()
	if err := a.WriteXML(a.RSS); err != nil {
		log.Fatal(err)
	}

}
