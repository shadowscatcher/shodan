package models

// Shodan crawler
type Shodan struct {
	// Unique ID that identifies the Shodan crawler
	Crawler string `json:"crawler"`

	// Unique ID for this banner
	Id *string `json:"id"`

	// Name of the Shodan module used by the crawler to generate this banner
	Module string `json:"module,omitempty"`

	// [NOT DOCUMENTED]
	Ptr bool `json:"ptr"`

	// Configuration options used during the data collection
	Options CrawlerOptions `json:"options"`
}

type CrawlerOptions struct {
	// Hostname to use when sending web requests
	Hostname string `json:"hostname,omitempty"`

	// ID of the banner that triggered the scan for this port
	Referrer string `json:"referrer,omitempty"`

	// ID of the scan
	Scan string `json:"scan,omitempty"`
}
