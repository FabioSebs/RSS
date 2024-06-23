package jobs

import (
	"time"

	"github.com/FabioSebs/RSS/httpclient"
	"github.com/FabioSebs/RSS/scheduler"
	"github.com/FabioSebs/RSS/scraper"
)

func RunScraper() {
	scheduler := scheduler.NewScheduler()
	scheduler.CreateJob(time.Minute*3, LaunchAllScrapers)
	scheduler.Start()
}

func LaunchAllScrapers() {
	mot := scraper.NewMoTScraper(true)
	mot.LaunchScraper(mot.CollectorSetup())

	mot2 := scraper.NewMoTScraper(false)
	mot2.LaunchScraper(mot.CollectorSetup())

	moe := scraper.NewMoeScraper(true)
	moe.LaunchScraper(moe.CollectorSetup())

	moe2 := scraper.NewMoeScraper(false)
	moe2.LaunchScraper(moe.CollectorSetup())

	// reg := scraper.NewRegScraper()
	// reg.LaunchScraper(reg.CollectorSetup())

	viet := scraper.NewVietnamScraper()
	viet.LaunchScraper(viet.CollectorSetup())

	thailand := scraper.NewThailandScraper()
	thailand.LaunchScraper(thailand.CollectorSetup())

	antara := scraper.NewAntaraScraper()
	antara.LaunchScraper(antara.CollectorSetup())

	httpclient.SendRequestForEmail()
}
