package services

type SNMP struct {
	// Contact information
	Contact string `json:"contact"`

	// Description of the device
	Description string `json:"description"`

	// Physical location
	Location *string `json:"location"`

	// Name given to the device by the owner
	Name *string `json:"name"`
}
