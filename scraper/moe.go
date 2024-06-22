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
	gtranslate "github.com/gilang-as/google-translate"
	"github.com/gocolly/colly"
)

type MoeScraper struct {
	Collector    *colly.Collector
	Config       config.Config
	RSSGenerator generator.RSSGenerator
	RSS          entities.RSS
	RSSitems     entities.Item
	English      bool
	// Logger    logger.Logger
}

func NewMoeScraper(english bool) WebScraper {
	env := config.NewConfig()
	return &MoeScraper{
		Collector: colly.NewCollector(colly.AllowedDomains(
			env.PermittedURLs.MoE...,
		//append more sites here
		)),
		Config:       env,
		RSSGenerator: generator.NewRssGenerator(),
		English:      english,
		// Logger: logger.NewLogger(),
	}
}

func (g *MoeScraper) CollectorSetup() *colly.Collector {

	g.RSS = entities.RSS{
		Version: "2.0",
		Channel: entities.Channel{
			Title:          "Ministry of Energy",
			Link:           g.Config.Domains.MoE,
			Description:    "Reports",
			ManagingEditor: g.Config.ICCTAuthor,
			PubDate:        time.Now(),
			Items:          nil, // needs scraping
		},
	}

	g.Collector.OnHTML("div.container div.list-berita", func(element *colly.HTMLElement) {
		// scraper logic goes here
		element.ForEach(" div.col-md-4 div.berita-item", func(i int, h *colly.HTMLElement) {
			var (
				item = entities.Item{
					Title:       h.ChildText("h4.title"),
					Link:        h.ChildAttr("a.btn-download", "href"),
					Description: h.ChildText("p.post-time"),
					PubDate:     time.Now(),
					Enclosure: entities.Enclosure{
						URL:  "https://www.esdm.go.id" + h.ChildAttr("img", "src"),
						Type: "image/jpg",
					},
				}
			)

			if g.English {
				titleTrans, err := gtranslate.Translator(gtranslate.Translate{
					Text: item.Title,
					//From: "id",
					To: "en",
				})
				if err != nil {
					fmt.Println(err.Error())
				}
				item.Title = titleTrans.Text

				if utils.ValidateTitle(item.Title) {
					g.RSS.Channel.Items = append(g.RSS.Channel.Items, item)
				}
			}

			if utils.ValidateTitleID(item.Title) {
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
func (g *MoeScraper) WriteXML(rss entities.RSS) error {
	var (
		filename string = fmt.Sprintf("xml/%s", g.Config.Filenames.MoE)
	)

	if g.English {
		filename = fmt.Sprintf("xml/%s", g.Config.Filenames.MoEEnglish)
	}

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

func (g *MoeScraper) LaunchScraper(collector *colly.Collector) {
	if err := collector.Visit(g.Config.Domains.MoE); err != nil {
		fmt.Println("error occured: " + err.Error())
	}
	// Ensuring that the scraping process completes before the program exits
	collector.Wait()

	if err := g.WriteXML(g.RSS); err != nil {
		log.Fatal(err)
	}

}
