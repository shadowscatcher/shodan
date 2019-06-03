package services

type SMB struct {
	// Whether or not the service allows anonymous connections
	Anonymous bool `json:"anonymous"`

	// List of features that the service supports
	Capabilities []string `json:"capabilities"`

	// Operating system
	OS string `json:"os,omitempty"`

	// Hex-encoded list of response packets that were received during the banner grabbing
	Raw []string `json:"raw"`

	// List of directories and files that are shared
	Shares []SmbShare `json:"shares,omitempty"`

	// Highest SMB version that was negotiated
	SmbVersion int `json:"smb_version"`

	// Name of the software that powers the service
	Software string `json:"software,omitempty"`
}

type SmbShare struct {
	Comments  string    `json:"comments"`
	Files     []SmbFile `json:"files,omitempty"`
	Name      string    `json:"name"`
	Special   bool      `json:"special"`
	Temporary bool      `json:"temporary"`
	Type      string    `json:"type"`
}

type SmbFile struct {
	Directory bool   `json:"directory"`
	Name      string `json:"name"`
	ReadOnly  bool   `json:"read-only"`
	Size      int    `json:"size"`
}
