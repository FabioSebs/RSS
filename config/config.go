package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Domains struct {
	MoE []string
	MoT []string
	//append more here

}

type Config struct {
	Sites      Domains
	ICCTEmail  string
	ICCTAuthor string
}

func NewConfig() Config {
	domains := Domains{
		MoE: []string{
			os.Getenv("domain.moe"),
			os.Getenv("domain.moe.allowed"),
			os.Getenv("domain.moe.allowed2"),
			os.Getenv("domain.moe.allowed3"),
			os.Getenv("domain.moe.allowed4"),
			os.Getenv("domain.moe.allowed5"),
			os.Getenv("domain.moe.allowed6"),
		},

		MoT: []string{os.Getenv("domain.mot")},
		//append more here
	}

	return Config{
		Sites:      domains,
		ICCTEmail:  os.Getenv("icct.email"),
		ICCTAuthor: os.Getenv("icct.author"),
	}
}
