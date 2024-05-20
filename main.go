package main

import (
	"github.com/FabioSebs/RSS/scheduler/jobs"
	"github.com/FabioSebs/RSS/server"
)

func main() {
	jobs.RunScraper()
	server.RunServer()
}

// TODO: Regen on request or cron job + cli command
// TODO: Improve folder structure
