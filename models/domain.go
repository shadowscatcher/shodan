package models

type NsRecord struct {
	Subdomain string `json:"subdomain"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	LastSeen  string `json:"last_seen"`
}

type Domain struct {
	More       bool       `json:"more"`
	Domain     string     `json:"domain"`
	Tags       []string   `json:"tags"`
	Data       []NsRecord `json:"data"`
	Subdomains []string   `json:"subdomains"`
}
