package services

type Rsync struct {
	// Whether the server requires authentication or not
	Authentication bool `json:"authentication"`

	// List of modules (folders) where the key is the name of the module and the value is the optional description
	Modules map[string]string `json:"modules"`
}
