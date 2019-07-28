package models

type Host struct {
	Ports      []int      `json:"ports"`
	Vulns      []string   `json:"vulns"`
	LastUpdate string     `json:"last_update"`
	Services   []*Service `json:"data"`
	Location
	HostInfo
}

// common fields for "/host/{ip}" and "/host/search" endpoint results
type HostInfo struct {
	// The IP address of the host as a string
	IPstr string `json:"ip_str"`

	// The autonomous system number (ex. "AS4837")
	ASN *string `json:"asn,omitempty"`

	// The operating system that powers the device
	OS *string `json:"os"`

	// The name of the organization that is assigned the IP space for this device
	Org *string `json:"org"`

	// The ISP that is providing the organization with the IP space for this device.
	// Consider this the "parent" of the organization in terms of IP ownership
	ISP *string `json:"isp"`

	// An array of strings containing all of the hostnames that have been assigned to the IP address for this device
	Hostnames []string `json:"hostnames"`

	// List of tags that describe the characteristics of the device
	Tags []string `json:"tags,omitempty"`

	// Raw HTML of response
	HTML string `json:"html,omitempty"`
}
