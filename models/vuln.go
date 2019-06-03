package models

type Vulnerability struct {
	// Common Vulnerability Scoring System value
	CVSS interface{} `json:"cvss"`

	// List of URLs that are related to the vulnerability
	References []string `json:"references"`

	// A description of the vulnerability
	Summary string `json:"summary"`

	// Whether or not the vulnerability has been verified by the Shodan crawlers
	Verified bool `json:"verified"`
}
