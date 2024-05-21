package scraper

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/FabioSebs/RSS/config"
	"github.com/FabioSebs/RSS/entities"
	"github.com/FabioSebs/RSS/generator"
	"github.com/gocolly/colly"
)

var (
	counter   int      = 0
	filenames []string = []string{
		"xml/reg_presiden.xml",
		"xml/reg_pemerintah.xml",
		"xml/reg_menteri.xml",
		"xml/reg_geburnur.xml",
	}

	excel   generator.ExcelGen = generator.NewExcelGen()
	columns []string           = []string{
		"Title",
		"Link",
		"Status",
		"PDF",
	}
)

type RegScraper struct {
	Collector    *colly.Collector
	Config       config.Config
	RSSGenerator generator.RSSGenerator
}

func NewRegScraper() WebScraper {
	var (
		env config.Config = config.NewConfig()
	)

	return &RegScraper{
		Collector: colly.NewCollector(colly.AllowedDomains(
			env.PermittedURLs.Reg...,
		)),
		Config:       env,
		RSSGenerator: generator.NewRssGenerator(),
	}
}

func (g *RegScraper) CollectorSetup() *colly.Collector {
	g.Collector.OnHTML("div.app-container div.rounded-4", func(element *colly.HTMLElement) {
		var (
			pubTime time.Time = time.Now()

			rss entities.RSS = entities.RSS{
				Version: "2.0",
				Channel: entities.Channel{
					Title:          "Database of Regulations (ID)",
					Link:           g.Config.Domains.Reg[0],
					Description:    "Regulations",
					ManagingEditor: g.Config.ICCTAuthor,
					PubDate:        pubTime,
					Items:          nil, // needs scraping
				},
			}
		)

		// add items to rss / scrape data
		element.ForEach("div.row.mb-8 div.col-12 div.card div.card-body div.flex-grow-1", func(i int, h *colly.HTMLElement) {
			var (
				item entities.Item = entities.Item{
					Title:       h.ChildText("div.row.g-4.g-xl-9.mb-4 div.col-lg-8"),
					Link:        g.Config.PermittedURLs.Reg[2] + h.ChildAttr("a", "href"),
					Description: cleanString(h.ChildText("a")),
					PubDate:     time.Now(),
				}

				excelRow []string = []string{
					item.Title,
					item.Link,
					h.ChildText("div div.row.g-4.g-xl-9.mb-4 div.col-lg-2 span"),
					g.Config.PermittedURLs.Reg[2] + h.ChildAttr("div.border-top.border-gray-300.pt-4.mt-4 div.row.g-4.g-xl-9 div ul li a", "href"),
				}
			)

			excel.SetValues(i+2, counter+1, excelRow) // (rowNo, sheetNo, values)
			rss.Channel.Items = append(rss.Channel.Items, item)
		})

		if len(rss.Channel.Items) > 0 {
			if err := g.WriteXML(rss); err != nil {
				log.Fatal(err)
			}
		}

		excel.NewSheet(fmt.Sprintf("Sheet%d", counter+2)) // have sheet 5 but its ok
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

func (g *RegScraper) WriteXML(rss entities.RSS) error {
	output, err := xml.MarshalIndent(rss, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling XML:", err)
		return err
	}

	// Add XML header
	output = []byte(xml.Header + string(output))

	// Write XML to file
	file, err := os.Create(filenames[counter])
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

func (g *RegScraper) LaunchScraper(collector *colly.Collector) {
	excel.NewSheet("Sheet1")

	for idx := 0; idx < len(g.Config.Domains.Reg); idx++ {
		excel.SetColumns(idx+1, columns)
		if err := collector.Visit(g.Config.Domains.Reg[idx]); err != nil {
			fmt.Println("error occured: " + err.Error())
		}
		counter++ // 0-3
	}

	counter = 0
	excel.SaveFile("regulations.xlsx")
	collector.Wait()
}

func cleanString(s string) string {
	// Remove HTML entities, newlines, tabs, and extra spaces
	re := regexp.MustCompile(`&#xA;|\n|\t`)
	s = re.ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)
	s = strings.Join(strings.Fields(s), " ") // Replace multiple spaces with a single space
	return s
}
