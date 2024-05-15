package scraper

import (
	"github.com/gocolly/colly"
)

type WebScraper interface {
	CollectorSetup() *colly.Collector
	// GetReviewsConcurrently(*colly.Collector) ([]EV, time.Duration)
	LaunchScraper(collector *colly.Collector)
}

type GoCollyProgram struct {
	Collector *colly.Collector
	// Config    config.Config
	// Logger    logger.Logger
}

func NewWebScraper() WebScraper {
	// env := config.NewConfig()

	return &GoCollyProgram{
		// Collector: colly.NewCollector(colly.AllowedDomains(
		// 	env.AllowedDomains...,
		// )),
		// Config: env,
		// Logger: logger.NewLogger(),
	}
}

func (g *GoCollyProgram) CollectorSetup() *colly.Collector {
	g.Collector.OnHTML("div.facetwp-template ", func(element *colly.HTMLElement) {
		// scraper goes here

	})

	// Request Feedback
	g.Collector.OnRequest(func(r *colly.Request) {
		// g.Logger.WriteTrace(fmt.Sprintf("visiting url: %s", r.URL.String()))
		//logging
	})

	// Error Feedback
	g.Collector.OnError(func(_ *colly.Response, err error) {
		// g.Logger.WriteError(fmt.Sprintf("error: %s", err.Error()))
	})
	return g.Collector
}

func (g *GoCollyProgram) LaunchScraper(collector *colly.Collector) {
	// start := time.Now()
	collector.Visit("")
}

// func writeRSS(data []Publication, fname string) {
// 	cardata, err := json.MarshalIndent(data, "", " ")
// 	if err != nil {
// 		log.Println("Unable to create json file")
// 		return
// 	}

// 	if err = ioutil.WriteFile(fmt.Sprintf("%s.json", fname), cardata, 0644); err != nil {
// 		log.Println("unable to write to json file")
// 	}
// }
