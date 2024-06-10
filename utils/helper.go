package utils

import (
	"os"

	"github.com/FabioSebs/RSS/entities"
)

func GetDomains() entities.Domains {
	return entities.Domains{
		MoE: os.Getenv("domain.moe"),
		MoT: os.Getenv("domain.mot"),
		Reg: []string{
			os.Getenv("domain.regulation.presiden"),
			os.Getenv("domain.regulation.pemerintah"),
			os.Getenv("domain.regulation.menteri"),
			os.Getenv("domain.regulation.geburnur"),
		},
		Antara:   os.Getenv("domain.antara"),
		Vietnam:  os.Getenv("domain.vietnam"),
		Thailand: os.Getenv("domain.thailand"),
	}
}

func GetAllowedDomains() entities.PermittedURLs {
	return entities.PermittedURLs{
		MoE: []string{
			os.Getenv("domain.moe.allowed"),
			os.Getenv("domain.moe.allowed2"),
			os.Getenv("domain.moe.allowed3"),
			os.Getenv("domain.moe.allowed4"),
			os.Getenv("domain.moe.allowed5"),
			os.Getenv("domain.moe.allowed6"),
		},
		MoT: []string{
			os.Getenv("domain.mot.allowed"),
			os.Getenv("domain.mot.allowed2"),
			os.Getenv("domain.mot.allowed3"),
			os.Getenv("domain.mot.allowed4"),
			os.Getenv("domain.mot.allowed5"),
			os.Getenv("domain.mot.allowed6"),
		},
		Reg: []string{
			os.Getenv("domain.regulation.allowed"),
			os.Getenv("domain.regulation.allowed2"),
			os.Getenv("domain.regulation.allowed3"),
			os.Getenv("domain.regulation.allowed4"),
			os.Getenv("domain.regulation.allowed5"),
			os.Getenv("domain.regulation.allowed6"),
		},
		Antara: []string{
			os.Getenv("domain.antara.allowed"),
			os.Getenv("domain.antara.allowed2"),
			os.Getenv("domain.antara.allowed3"),
			os.Getenv("domain.antara.allowed4"),
			os.Getenv("domain.antara.allowed5"),
			os.Getenv("domain.antara.allowed6"),
		},
		Vietnam: []string{
			os.Getenv("domain.vietnam.allowed"),
			os.Getenv("domain.vietnam.allowed2"),
			os.Getenv("domain.vietnam.allowed3"),
			os.Getenv("domain.vietnam.allowed4"),
			os.Getenv("domain.vietnam.allowed5"),
			os.Getenv("domain.vietnam.allowed6"),
		},
		Thailand: []string{
			os.Getenv("domain.thailand.allowed"),
			os.Getenv("domain.thailand.allowed2"),
			os.Getenv("domain.thailand.allowed3"),
			os.Getenv("domain.thailand.allowed4"),
			os.Getenv("domain.thailand.allowed5"),
			os.Getenv("domain.thailand.allowed6"),
		},
	}
}

func GetFilenames() entities.Filenames {
	return entities.Filenames{
		MoE:      "moe.xml",
		MoT:      "mot.xml",
		ReG:      "reg.xml",
		Antara:   "antara.xml",
		Vietnam:  "vietnam.xml",
		Thailand: "thailand.xml",
	}
}
