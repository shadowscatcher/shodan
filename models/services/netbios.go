package services

type Netbios struct {
	// MAC address of the network interface
	MAC string `json:"mac"`

	// List of NetBIOS names
	Names []NetbiosName `json:"names"`

	// List of additional networks/interfaces that the service is listening on
	Networks []string `json:"networks"`

	// List of hex-encoded response packets
	Raw []string `json:"raw"`

	// Name of the server running NetBIOS
	Servername string `json:"servername"`

	// Name of the user running the service
	Username *string `json:"username"`
}

type NetbiosName struct {
	Flags  int    `json:"flags"`
	Name   string `json:"name"`
	Suffix int    `json:"suffix"`
}
