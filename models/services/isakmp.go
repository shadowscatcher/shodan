package services

type ISAKMP struct {
	Aggressive   *ISAKMP `json:"aggressive,omitempty"`
	ExchangeType int     `json:"exchange_type"`

	// Flags that are enabled/ disabled for the connection
	Flags IsakmpFlags `json:"flags"`

	// Hex-encoded security parameter index for the initiator
	InitiatorSPI string `json:"initiator_spi"`

	// Size of the ISAKMP packet
	Length int `json:"length"`

	// Hex-encoded ID for the message
	MsgID string `json:"msg_id"`

	// Next payload type sent after the initiation
	NextPayload int `json:"next_payload"`

	// Hex-encoded security parameter index for the responder
	ResponderSPI string `json:"responder_spi"`

	// List of vendor IDs that can be used to fingerprint the service
	VendorIds []string `json:"vendor_ids"`

	// Protocol version
	Version string `json:"version"`
}

type IsakmpFlags struct {
	Authentication bool `json:"authentication"`
	Commit         bool `json:"commit"`
	Encryption     bool `json:"encryption"`
}
