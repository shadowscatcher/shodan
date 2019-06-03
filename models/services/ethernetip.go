package services

type EthernetIP struct {
	// Command code
	Command       int `json:"command"`
	CommandLength int `json:"command_length"`

	// Return status of the command
	CommandStatus int `json:"command_status"`

	// Device type
	DeviceType string `json:"device_type"`

	// Length of the encapsulation header
	EncapsulationLength int    `json:"encapsulation_length"`
	IP                  string `json:"ip"`

	// Number of items returned
	ItemCount int `json:"item_count"`

	// Encapsulation options
	Options int `json:"options"`

	// Numeric identifier for the product
	ProductCode int `json:"product_code"`

	// Product name
	ProductName       string `json:"product_name"`
	ProductNameLength int    `json:"product_name_length"`

	// Hex-encoded string containing the original response from the Ethernet/IP service
	Raw string `json:"raw"`

	// Major revision number
	RevisionMajor int `json:"revision_major"`

	// Minor revision number
	RevisionMinor int `json:"revision_minor"`

	// Encapsulation sender context
	SenderContext string `json:"sender_context"`

	// Serial number
	Serial int `json:"serial"`

	// Encapsulation session identifier
	Session int `json:"session"`

	// Hex-encoded string containing the socket address
	SocketAddr string `json:"socket_addr"`

	// Device state
	State int `json:"state"`

	// Numeric vendor identifier
	Status int `json:"status"`
	TypeID int `json:"type_id"`

	// string, but can be int sometimes
	VendorID interface{} `json:"vendor_id"` // can be int

	// Protocol version
	Version int `json:"version"`
}
