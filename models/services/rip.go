package services

type RIP struct {
	// List of hosts/routes served by the daemon
	Addresses []RipAddress `json:"addresses"`

	// Command type (1 – Request, 2 – Response)
	Command int `json:"command"`

	// RIP version used by the service
	Version int `json:"version"`
}

type RipAddress struct {
	Addr string `json:"addr"`
	// Note: can be int (0)
	Family  interface{} `json:"family"`
	Metric  int         `json:"metric"`
	NextHop interface{} `json:"next_hop"`
	Subnet  *string     `json:"subnet"`
	Tag     *int        `json:"tag"`
}
