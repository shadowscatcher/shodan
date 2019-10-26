package models

// common Shodan error
type Error struct {
	Error string `json:"error"`
}

// Response with success indicator
type SimpleResponse struct {
	Success bool `json:"success"`
}

// Result of request to "/scan" endpoints
type Scan struct {
	Id          string `json:"id"`
	Count       int    `json:"count,omitempty"`
	CreditsLeft int    `json:"credits_left,omitempty"`
	Status      string `json:"status,omitempty"`
}

// Result of request to "/scans" endpoint
type ScanList struct {
	Total   int          `json:"total"`
	Matches []ScanReport `json:"matches"`
}

// Part of "/scans" result
type ScanReport struct {
	Id          string `json:"id"`
	Status      string `json:"status"`
	Created     string `json:"created"`
	Size        int    `json:"size"`
	StatusCheck string `json:"status_check"`
	CreditsLeft int    `json:"credits_left"`
	ApiKey      string `json:"api_key"`
}

// Result of request to "/account/profile" endpoint
type Profile struct {
	Member      bool   `json:"member"`
	Credits     int    `json:"credits"`
	DisplayName string `json:"display_name"`
	Created     string `json:"created"`
}

// Result of request to "/api-info" endpoint
type ApiInfo struct {
	ScanCredits  int            `json:"scan_credits"`
	UsageLimits  map[string]int `json:"usage_limits"`
	Plan         string         `json:"plan"`
	Https        bool           `json:"https"`
	Unlocked     bool           `json:"unlocked"`
	QueryCredits int            `json:"query_credits"`
	MonitoredIps int            `json:"monitored_ips"`
	UnlockedLeft int            `json:"unlocked_left"`
	Telnet       bool           `json:"telnet"`
}

// Result of request to "/org" endpoint
type Org struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Created     string      `json:"created"`
	Admins      []OrgMember `json:"admins"`
	Members     []OrgMember `json:"members"`
	UpgradeType string      `json:"upgrade_type"`
	Domains     []string    `json:"domains"`
	Logo        interface{} `json:"logo"`
}

type OrgMember struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Search query broken into tokens
type Tokens struct {
	Attributes map[string][]interface{} `json:"attributes"`
	Errors     []string                 `json:"errors"`
	String     string                   `json:"string"`
	Filters    []string                 `json:"filters"`
}

// List of the saved search queries
type SearchQueries struct {
	Total   int           `json:"total"`
	Matches []SearchQuery `json:"matches"`
}

// Saved search query
type SearchQuery struct {
	Votes       int      `json:"votes"`
	Description string   `json:"description"`
	Title       string   `json:"title"`
	Timestamp   string   `json:"timestamp"`
	Tags        []string `json:"tags"`
	Query       string   `json:"query"`
}

// Search query tags list
type QueryTags struct {
	Total   int        `json:"total"`
	Matches []QueryTag `json:"matches"`
}

// Search query tag
type QueryTag struct {
	Value interface{} `json:"value"`
	Count int         `json:"count"`
}

// Dataset description
type Dataset struct {
	Name        string `json:"name"`
	Scope       string `json:"scope"`
	Description string `json:"description"`
}

// File in dataset
type DatasetFile struct {
	Url       string `json:"url"`
	Timestamp int    `json:"timestamp"`
	Name      string `json:"name"`
	Size      int    `json:"size"`
}

// An object specifying the criteria that an alert should trigger. The only supported option at the moment is the "ip" filter
type Filter struct {

	// A list of IPs or network ranges defined using CIDR notation
	IP []string `json:"ip"`
}

// Alert on network events
type Alert struct {
	// The name to describe the network alert
	Name string `json:"name"`

	// An object specifying the criteria that an alert should trigger. The only supported option at the moment is the "ip" filter
	Filter Filter `json:"filters"`

	// Number of seconds that the alert should be active
	Expires *uint `json:"expires,omitempty"`
}

// information about a specific network alert
type AlertDetails struct {
	Alert
	Created    string `json:"created,omitempty"`
	Expiration string `json:"expiration,omitempty"`
	Id         string `json:"id,omitempty"`
	Triggers   map[string]struct {
		Ignore []struct {
			IP   string `json:"ip"`
			Port uint16 `json:"port"`
		} `json:"ignore,omitempty"`
	} `json:"triggers"`
}

// Trigger is an alert trigger - when it fires, a notification about event is sent
type Trigger struct {
	Name        string `json:"name"`
	Rule        string `json:"rule"`
	Description string `json:"description"`
}
