package services

type Lantronix struct {
	// Gateway of the Lantronix device
	Gateway *string `json:"gateway"`

	// Main IP address that the service is listening on
	IP *string `json:"ip"`

	// MAC address
	Mac string `json:"mac"`

	// Password for the device
	Password *string `json:"password"`

	// Device type
	Type *string `json:"type"`

	// Lantronix firmware version
	Version string `json:"version"`
}
