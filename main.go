package main

import (
	"github.com/FabioSebs/RSS/scraper"
)

func main() {
	mot := scraper.NewMoTScraper()
	mot.LaunchScraper(mot.CollectorSetup())

	moe := scraper.NewMoeScraper()
	moe.LaunchScraper(moe.CollectorSetup())

	reg := scraper.NewRegScraper()
	reg.LaunchScraper(reg.CollectorSetup())
	// server.RunServer()
}

// TODO: Regen on request or cron job + cli command
// TODO: Improve folder structure
