package services

type SSH struct {
	// Cipher used during negotiation
	Cipher string `json:"cipher"`

	// Fingerprint for the service
	Fingerprint string `json:"fingerprint"`
	Hassh       string `json:"hassh"`

	// Key exchange algorithms that are supported by the server
	Kex SshKex `json:"kex"`

	// SSH key of the service
	Key string `json:"key"`

	// Message authentication code algorithm
	Mac string `json:"mac"`

	// Key type
	Type string `json:"type"`
}

type SshKex struct {
	CompressionAlgorithms   []string `json:"compression_algorithms"`
	EncryptionAlgorithms    []string `json:"encryption_algorithms"`
	KexAlgorithms           []string `json:"kex_algorithms"`
	KexFollows              bool     `json:"kex_follows"`
	Languages               []string `json:"languages"`
	MacAlgorithms           []string `json:"mac_algorithms"`
	ServerHostKeyAlgorithms []string `json:"server_host_key_algorithms"`
	Unused                  int      `json:"unused"`
}
