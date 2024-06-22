package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/FabioSebs/RSS/scheduler/jobs"
	"github.com/FabioSebs/RSS/scraper"
	"github.com/FabioSebs/RSS/server"
	"github.com/labstack/gommon/color"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// go routine for listening for os signal (os.Interrupt)
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		<-signals // thread hangs until signal is found and written
		cancel()  // sends message to ctx
	}()

	// start server in seperate thread
	go func() {
		color.Println(color.Green("⇨ server up and running "))
		server.RunServer()
	}()

	// start scheduler in seperate thread
	go jobs.RunScraper()

	// any other process
	// s := scraper.NewMoTScraper()
	// s.LaunchScraper(s.CollectorSetup())

	// scraper2 := scraper.NewMoeScraper()
	// scraper2.LaunchScraper(scraper2.CollectorSetup())

	scraper3 := scraper.NewMoeScraper(false)
	scraper3.LaunchScraper(scraper3.CollectorSetup())

	// scraper4 := scraper.NewThailandScraper()
	// scraper4.LaunchScraper(scraper4.CollectorSetup())

	// scraper5 := scraper.NewVietnamScraper()
	// scraper5.LaunchScraper(scraper5.CollectorSetup())

	// scraper6 := scraper.NewAntaraScraper()
	// scraper6.LaunchScraper(scraper6.CollectorSetup())

	// main thread is waiting for os interrupt aka context cancel
	<-ctx.Done()
}
