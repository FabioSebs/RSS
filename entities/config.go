package entities

type AllowedDomains []string

type Domains struct {
	MoE string
	MoT string
	Reg []string
}

type PermittedURLs struct {
	MoE AllowedDomains
	MoT AllowedDomains
	Reg AllowedDomains
}

type Filenames struct {
	MoE string
	MoT string
	ReG string
}
