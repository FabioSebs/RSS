package entities

type AllowedDomains []string

type Domains struct {
	MoE      string
	MoT      string
	Reg      []string
	Antara   string
	Vietnam  string
	Thailand string
	// add more sources here
}

type PermittedURLs struct {
	MoE      AllowedDomains
	MoT      AllowedDomains
	Reg      AllowedDomains
	Antara   AllowedDomains
	Vietnam  AllowedDomains
	Thailand AllowedDomains
}

type Filenames struct {
	MoE      string
	MoT      string
	ReG      string
	Antara   string
	Vietnam  string
	Thailand string
}
