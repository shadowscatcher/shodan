package services

type DB2 struct {
	// External name for the database server
	ExternalName string `json:"external_name"`

	// Hardware/operating system information
	ServerPlatform string `json:"server_platform"`

	// Name of the server
	InstanceName string `json:"instance_name"`

	Version string `json:"db2_version"`
}
