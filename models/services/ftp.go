package services

type FTP struct {
	// True if anonymous access is allowed
	Anonymous bool `json:"anonymous"`

	// List of features and the valid parameters that are supported by the FTP server
	Features map[string]FtpFeature `json:"features"`

	// Numeric hash of the list of supported ftp.features
	FeaturesHash *int `json:"features_hash"`
}

type FtpFeature struct {
	Parameters []string `json:"parameters"`
}
