package main

import (
	"github.com/FabioSebs/RSS/scraper"
)

func main() {
	// _, cancel := context.WithCancel(context.Background())

	// go routine for listening for os signal (os.Interrupt)
	// go func() {
	// 	signals := make(chan os.Signal, 1)
	// 	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	// 	<-signals // thread hangs until signal is found and written
	// 	cancel()  // sends message to ctx
	// }()

	// start server in seperate thread
	// go func() {
	// 	color.Println(color.Green("⇨ server up and running "))
	// 	server.RunServer()
	// }()

	// start scheduler in seperate thread
	// go jobs.RunScraper()

	// any other process
	scraper := scraper.NewThailandScraper()
	scraper.LaunchScraper(scraper.CollectorSetup())

	// main thread is waiting for os interrupt aka context cancel
	// <-ctx.Done()

}
