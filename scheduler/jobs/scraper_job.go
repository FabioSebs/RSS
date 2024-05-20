package jobs

import (
	"time"

	"github.com/FabioSebs/RSS/scheduler"
	"github.com/FabioSebs/RSS/scraper"
)

func RunScraper() {
	scheduler := scheduler.NewScheduler()
	scheduler.CreateJob(time.Hour*24, LaunchAllScrapers)
	scheduler.Start()
}

func LaunchAllScrapers() {
	mot := scraper.NewMoTScraper()
	mot.LaunchScraper(mot.CollectorSetup())

	moe := scraper.NewMoeScraper()
	moe.LaunchScraper(moe.CollectorSetup())

	reg := scraper.NewRegScraper()
	reg.LaunchScraper(reg.CollectorSetup())
}