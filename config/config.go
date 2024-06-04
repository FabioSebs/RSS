package config

import (
	"os"

	"github.com/FabioSebs/RSS/entities"
	"github.com/FabioSebs/RSS/utils"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	entities.Domains
	entities.PermittedURLs
	entities.Filenames
	ICCTEmail  string
	ICCTAuthor string
}

func NewConfig() Config {
	return Config{
		Domains:       utils.GetDomains(),
		PermittedURLs: utils.GetAllowedDomains(),
		Filenames:     utils.GetFilenames(),
		ICCTEmail:     os.Getenv("icct.email"),
		ICCTAuthor:    os.Getenv("icct.author"),
	}
}

func GetAllRecipients() []string {
	return []string{
		"fabrzytech@gmail.com",
		"sebsphotosofficial@gmail.com",
	}
}
