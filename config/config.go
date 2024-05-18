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

type Filenames struct {
	MoE string
	MoT string
}

type Config struct {
	Sites Domains
	Filenames
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

	filenames := Filenames{
		MoE: "moe.xml",
		MoT: "mot.xml",
	}

	return Config{
		Sites:      domains,
		Filenames:  filenames,
		ICCTEmail:  os.Getenv("icct.email"),
		ICCTAuthor: os.Getenv("icct.author"),
	}
}
